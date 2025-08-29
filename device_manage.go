package aifinitsdk

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"resty.dev/v3"
)

// AinfinitError represents an error with the ainfinit tag
type AinfinitError struct {
	Err error
}

func (e *AinfinitError) Error() string {
	return fmt.Sprintf("[ainfinit] %v", e.Err)
}

// NewAinfinitError creates a new AinfinitError
func NewAinfinitError(err error) error {
	return &AinfinitError{Err: err}
}

type VendingMachineManageClient interface {
	Activation(machineCode string, request *DeviceActivationRequest) (*DeviceActivationResponse, error)
	List(request *ListMachineRequest) (*ListMachineResponse, error)
	DeviceInfo(machineCode string) (*DeviceInfoResponse, error)
	MachineDetail(machineCode string) (*MachineDetailResponse, error)
	PeopleFlow(request *DevicePeopleFlowRequest, machineCode string) (*DevicePeopleFlowResponse, error)
	Update(request *DeviceUpdateRequest, machineCode string) (*DeviceUpdateResponse, error)
	Control(request *DeviceControlRequest, machineCode string) (*DeviceControlResponse, error)

	Alarm(machineCode string)
	Setting(request SettingRequest, machineCode string) (*SettingResponse, error)
	RefrigerationControl(request RefrigerationControlRequest, machineCode string) (*RefrigerationControlResponse, error)
}

type vendingMachineManageClient struct {
	Client Client
	Resty  *resty.Client
}

func NewDeviceClient(client Client) VendingMachineManageClient {
	restyClient := resty.New()
	if client.RestyDebug() {
		restyClient.SetDebug(true)
	}

	return &vendingMachineManageClient{
		Client: client,
		Resty:  restyClient,
	}
}

func (c *vendingMachineManageClient) Update(request *DeviceUpdateRequest, machineCode string) (*DeviceUpdateResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request": request,
			"code":    machineCode,
		}).Debug("Updating vending machine")
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return nil, NewAinfinitError(err)
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var result DeviceUpdateResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetQueryParam("code", machineCode).SetBody(request).SetResult(&result).
		Put(Put_UpdateVendingMachineInfo)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String()))
	}

	if !isSuccessStatus(result.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", result.Status, result.Message))
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", result),
		}).Debug("Updated vending machine successfully")
	}

	return &result, nil
}

func (c *vendingMachineManageClient) DeviceInfo(machineCode string) (*DeviceInfoResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithField("code", machineCode).Debug("Getting vending machine details")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var result DeviceInfoResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetResult(&result).
		SetQueryParam("code", machineCode).
		Get(Get_VendingMachineInfo)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String()))
	}

	if !isSuccessStatus(result.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", result.Status, result.Message))
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", result),
		}).Debug("Got vending machine details successfully")
	}

	return &result, nil
}

func (c *vendingMachineManageClient) MachineDetail(machineCode string) (*MachineDetailResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithField("code", machineCode).Debug("Getting device info")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var result MachineDetailResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetResult(&result).
		SetQueryParam("code", machineCode).
		Get(Get_VendingMachineDeviceDetail)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String()))
	}

	if !isSuccessStatus(result.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", result.Status, result.Message))
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", result),
		}).Debug("Got device info successfully")
	}

	return &result, nil
}

func (c *vendingMachineManageClient) PeopleFlow(request *DevicePeopleFlowRequest, machineCode string) (*DevicePeopleFlowResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request": request,
		}).Debug("Getting people flow data")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var result DevicePeopleFlowResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetBody(request).SetResult(&result).
		Post(Post_VendingMachinePeopleFlow)
	if err != nil {
		return nil, NewAinfinitError(err)
	}
	if resp.IsError() {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String()))
	}

	if !isSuccessStatus(result.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", result.Status, result.Message))
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", result),
		}).Debug("Got people flow data successfully")
	}

	return &result, nil
}

func (c *vendingMachineManageClient) List(request *ListMachineRequest) (*ListMachineResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request": request,
		}).Debug("Listing vending machines")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var result ListMachineResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetResult(&result).
		SetQueryParam("page", strconv.Itoa(request.Page)).
		SetQueryParam("limit", strconv.Itoa(request.Limit)).
		SetQueryParam("nameOf", request.NameOf).
		Get(Get_VendingMachineList)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String()))
	}

	if !isSuccessStatus(result.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", result.Status, result.Message))
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", result),
		}).Debug("Listed vending machines successfully")
	}

	return &result, nil
}

func (c *vendingMachineManageClient) Control(request *DeviceControlRequest, machineCode string) (*DeviceControlResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request": request,
			"code":    machineCode,
		}).Debug("Controlling vending machine")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var result DeviceControlResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetQueryParam("code", machineCode).SetBody(request).SetResult(&result).
		Put(Put_VendingMachineDeviceControl)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String()))
	}

	if !isSuccessStatus(result.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", result.Status, result.Message))
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", result),
		}).Debug("Controlled vending machine successfully")
	}

	return &result, nil
}

func (c *vendingMachineManageClient) Activation(machineCode string, request *DeviceActivationRequest) (*DeviceActivationResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithField("request", request).Debug("Activating vending machine")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var result DeviceActivationResponse
	_, err = c.Resty.R().SetHeader("Authorization", signature).SetQueryParam("code", machineCode).SetBody(request).SetResult(&result).
		Post(Post_DeviceActivation)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", result),
		}).Debug("Activated vending machine successfully")
	}

	return &result, nil
}
func (c *vendingMachineManageClient) Alarm(machineCode string) {}

func (c *vendingMachineManageClient) Setting(request SettingRequest, machineCode string) (*SettingResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithField("request", request).Debug("Setting vending machine")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	info, err := c.DeviceInfo(machineCode)
	if err != nil {
		return nil, err
	}

	var result SettingResponse
	_, err = c.Resty.R().SetHeader("Authorization", signature).SetQueryParam("scanCode", machineCode).SetQueryParam("deviseSn", info.Data.DeviceSn).SetBody(request).SetResult(&result).
		Put(Put_DeviceSetting)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", result),
		}).Debug("Setting vending machine successfully")
	}

	return &result, nil

}

func (c *vendingMachineManageClient) RefrigerationControl(request RefrigerationControlRequest, machineCode string) (*RefrigerationControlResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithField("request", request).Debug("RefrigerationControl requested")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var result RefrigerationControlResponse
	_, err = c.Resty.R().SetHeader("Authorization", signature).SetQueryParams(map[string]string{
		"vmCode":      request.VmCode,
		"comprEnable": strconv.Itoa(request.ComprEnable),
		"temp":        strconv.Itoa(request.Temp),
		"tempMode":    strconv.Itoa(request.TempMode),
	}).SetBody(request).SetResult(&result).
		Put(Put_DeviceSetting)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", result),
		}).Debug("Setting vending machine successfully")
	}

	return &result, nil

}

type RefrigerationControlRequest struct {
	VmCode      string `json:"replVideoUploadFlag"` // Restocking video upload, 0 = No, 1 = Yes
	ComprEnable int    `json:"comprEnable"`         // thermostat switch: 1 on, 0 off
	Temp        int    `json:"temp"`                // refrigeration:(-28~-18), heating (30-50)
	TempMode    int    `json:"tempMode"`            // 0: normal temp, 10: refrigeration, 11: refrigeration energy saving, 20: heating mode, 21: heating energy saving

}

type RefrigerationControlResponse struct {
	Status  uint   `json:"status"`
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

type SettingRequest struct {
	ReplVideoUploadFlag uint `json:"replVideoUploadFlag"` // Restocking video upload, 0 = No, 1 = Yes
}

type SettingResponse struct {
	Status  uint   `json:"status"`
	Message string `json:"message"`
}

type DeviceActivationRequest struct {
	Name          string `json:"name"`
	Location      string `json:"location"`
	ScanCode      string `json:"scanCode"`
	ContactNumber string `json:"contactNumber"`
}

type DeviceActivationResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Device represents a vending machine device with its status and capabilities
type Device struct {
	// Code is the unique identifier for the vending machine
	Code string `json:"code"`
	// BytesTotal represents total storage space in MB
	BytesTotal int `json:"bytesTotal"`
	// BytesFree represents available storage space in MB
	BytesFree int `json:"bytesFree"`
	// ClientVersion represents the current version of the client software
	ClientVersion string `json:"clientVersion"`
	// CameraCount represents the number of available cameras
	CameraCount int `json:"cameraCount"`
	// GravityCount represents the number of available gravity sensors
	GravityCount int `json:"gravityCount"`
	// Light represents the light status (0: normal, 1: abnormal)
	Light int `json:"light"`
	// Detector represents the local detection status (0: normal, 1: abnormal)
	Detector int `json:"detector"`
	// GravitySensor represents the gravity sensor status (0: normal, 1: abnormal)
	GravitySensor int `json:"gravitySensor"`
	// SerialPort represents the serial port connection status (0: normal, 1: abnormal)
	SerialPort int `json:"serialPort"`
	// SerialPortDataFormat represents the serial port data format status (0: normal, 1: abnormal)
	SerialPortDataFormat int `json:"serialPortDataFormat"`
	// NetworkInfo contains information about network connections
	NetworkInfo struct {
		// Mobile represents 4G connection status (1: connected, 0: disconnected)
		Mobile int `json:"mobile"`
		// Wifi represents WiFi connection status (1: connected, 0: disconnected)
		Wifi int `json:"wifi"`
		// Etherne represents ethernet connection status (1: connected, 0: disconnected)
		Etherne int `json:"ethernet"`
		// SignalWifi represents WiFi signal strength
		SignalWifi int `json:"signalWifi"`
		// SignalMobile represents 4G signal strength
		SignalMobile int `json:"signalMobile"`
	} `json:"networkInfo"`
	// CardInfo contains information about the SIM card
	CardInfo struct {
		// Iccid is the Integrated Circuit Card Identifier
		Iccid string `json:"iccid"`
		// Carrier represents the mobile carrier (CMCC: China Mobile, CUCC: China Unicom, CTCC: China Telecom)
		Carrier string `json:"carrier"`
		// Imsi is the International Mobile Subscriber Identity
		Imsi string `json:"imsi"`
	} `json:"cardInfo"`
	// Ccid is the Chip Card ID (empty string if not reported)
	Ccid string `json:"ccid"`
	// PowerStatus represents the power source (1: mains power, 2: UPS)
	PowerStatus int `json:"powerStatus"`
	// DeviceUpdateTimestamp represents the last heartbeat timestamp
	DeviceUpdateTimestamp int64 `json:"deviceUpdateTimestamp"`
	// OnlineStatus represents the network status (1: online, 0: offline, calculated based on heartbeat timestamp)
	OnlineStatus int `json:"onlineStatus"`
	// Temperature represents the current temperature
	Temperature int `json:"temperature"`
	// TargetTemp represents the target temperature setting
	TargetTemp int `json:"targetTemp"`
	// Volume represents the current volume level
	Volume int `json:"volume"`
	// EngineOn represents the compressor status (1: on, 0: off)
	EngineOn int `json:"engineOn"`
}

type VendingMachine struct {
	DeviceSn   string `json:"deviceSn"`
	ScanCode   string `json:"scanCode"`
	Name       string `json:"name"`
	Location   string `json:"location"`
	UpdateTime string `json:"updateTime"`
}

type PeopleFlow struct {
	Code          string `json:"code"`
	VisitorCount  int    `json:"visitorCount"`
	AggregateTime string `json:"aggregateTime"`
}

type ListMachineRequest struct {
	Page   int    `json:"page,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	NameOf string `json:"nameOf,omitempty"`
}

type DevicePeopleFlowRequest struct {
	Field          string   `json:"field,omitempty"`
	StartTimeStamp int64    `json:"startTimestamp,omitempty"`
	EndTimeStamp   int64    `json:"endTimestamp,omitempty"`
	Codes          []string `json:"codes,omitempty"`
}

type DeviceUpdateRequest struct {
	Name          string `json:"name,omitempty" validate:"required"`
	Code          string `json:"code,omitempty"`
	ScanCode      string `json:"scanCode,omitempty"`
	ContactNumber string `json:"contactNumber,omitempty"`
	Location      string `json:"location,omitempty"`
	Volume        int    `json:"volume,omitempty"`
	AdVolume      int    `json:"adVolume,omitempty"`
	Temp          int    `json:"temp,omitempty"`
	EngineOn      int    `json:"engineOn,omitempty"`
}

type DeviceControlRequest struct {
	Volume   int `json:"volume,omitempty"`   // 0 ~ 100
	AdVolume int `json:"adVolume,omitempty"` // 0 ~ 100
	Temp     int `json:"temp,omitempty"`     // -30 ~ 20
	EngineOn int `json:"engineOn,omitempty"` // 0 | 1
}
type MachineDetailResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Device `json:"data"`
}

type ListMachineResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Total int              `json:"total"`
		Rows  []VendingMachine `json:"rows"`
	}
}

type DevicePeopleFlowResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Result  []PeopleFlow `json:"result"`
	Count   int          `json:"count"`
}

type DeviceInfoResponse struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    DeviceInfoData `json:"data"`
}

type DeviceInfoData struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	ScanCode      string `json:"scanCode"`
	DeviceSn      string `json:"deviceSn"`
	ContactNumber string `json:"contactNumber"`
	Location      string `json:"location"`
	UpdateTime    string `json:"updateTime"`
}

type DeviceUpdateResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type DeviceControlResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
