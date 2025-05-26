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
	List(request *ListMachineRequest) (*ListMachineResponse, error)
	Detail(machineCode string) (*DeviceDetailResponse, error)
	DeviceInfo(machineCode string) (*DeviceInfoResult, error)
	PeopleFlow(request *DevicePeopleFlowRequest, machineCode string) (*DevicePeopleFlowResponse, error)
	Update(request *DeviceUpdateRequest, machineCode string) (*DeviceUpdateResponse, error)
	Control(request *DeviceControlRequest, machineCode string) (*DeviceControlResponse, error)

	Activation(machineCode string)
	Alarm(machineCode string)
	Setting(machineCode string)
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

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", result),
		}).Debug("Updated vending machine successfully")
	}

	return &result, nil
}

func (c *vendingMachineManageClient) Detail(machineCode string) (*DeviceDetailResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithField("code", machineCode).Debug("Getting vending machine details")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var result DeviceDetailResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetResult(&result).
		SetQueryParam("code", machineCode).
		Get(Get_VendingMachineInfo)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String()))
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", result),
		}).Debug("Got vending machine details successfully")
	}

	return &result, nil
}

func (c *vendingMachineManageClient) DeviceInfo(machineCode string) (*DeviceInfoResult, error) {
	if c.Client.IsDebug() {
		logrus.WithField("code", machineCode).Debug("Getting device info")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var result DeviceInfoResult
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetResult(&result).
		SetQueryParam("code", machineCode).
		Get(Get_VendingMachineDeviceInfo)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String()))
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

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", result),
		}).Debug("Controlled vending machine successfully")
	}

	return &result, nil
}

func (c *vendingMachineManageClient) Activation(machineCode string) {}
func (c *vendingMachineManageClient) Alarm(machineCode string)      {}
func (c *vendingMachineManageClient) Setting(machineCode string)    {}

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
type DeviceInfoResult struct {
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

type DeviceDetailResponse struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    DeviceDetailData `json:"data"`
}

type DeviceDetailData struct {
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
