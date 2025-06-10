package aifinitsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"resty.dev/v3"
)

type OperationClient interface {
	//mashind baraa nemne
	//2.2.3.11
	AddGoods(request *AddNewGoodsRequest, machineCode string) (*AddNewGoodsResponse, error)
	//2.2.3.12
	DeleteGoods(request *DeleteGoodsRequest, machineCode string) (*DeleteGoodsResponse, error)

	//machine dotorh baraanuudiin jagsaalt
	//2.2.3.1
	ListGoods(machineCode string) (*GetMachineGoodsResponse, error)
	// zaragdsan baraanuudiig niitedni shinechlene
	//2.2.3.2
	UpdateGoods(request *UpdateGoodsRequest, machineCode string) (*UpdateSoldGoodsResponse, error)
	//2.2.3.3
	OpenDoor(request *OpenDoorRequest, machineCode string) (*OpenDoorResponse, error)
	// zaragdsan baraag haalga ongoilgoh requesteer avah
	//2.2.3.4
	OpenDoorReqDetail(request *OpenDoorDetailRequest, machineCode string) (*OpenDoorDetailResponse, error)
	//zaragdsan baraanii jagsaalt
	//2.2.3.6
	ListOrders(request *ListOrderRequest, machineCode string) (*ListOrderResponse, error)
	// orderiin video avah
	//2.2.3.8
	GetOrderVideo(request *GetOrderVideoRequest, machineCode string) (*GetOrderVideoResponse, error)
	// product price update
	//2.2.3.10
	UpdateGoodsPrice(request *UpdateGoodsPriceRequest, machineCode string) (*ProductPriceUpdateResponse, error)
}

type OperationClientImpl struct {
	Client Client
	Resty  *resty.Client
}

func NewOperationClientImpl(client Client) OperationClient {
	restyClient := resty.New()

	if client.RestyDebug() {
		restyClient.SetDebug(true)
	}

	return &OperationClientImpl{
		Client: client,
		Resty:  restyClient,
	}
}

func (c *OperationClientImpl) OpenDoor(request *OpenDoorRequest, machineCode string) (*OpenDoorResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request":     request,
			"device_code": machineCode,
		}).Debug("Opening door")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var openDoorResponse *OpenDoorResponse
	req := c.Resty.R().SetHeader("Authorization", signature).
		SetQueryParam("code", machineCode).
		SetResult(&openDoorResponse)

	if request.Type != 0 {
		req.SetQueryParam("type", fmt.Sprintf("%d", request.Type))
	}

	if request.LocalTimeStamp != 0 {
		req.SetQueryParam("localTimestamp", fmt.Sprintf("%d", request.LocalTimeStamp))
	}

	if request.UserCode != "" {
		req.SetQueryParam("userCode", request.UserCode)
	}

	if request.RequestID != "" {
		req.SetQueryParam("requestId", request.RequestID)
	}

	resp, err := req.Put(Put_OpenDoor)

	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, ConvertOpenDoorError(resp.StatusCode(), resp.String())
	}

	if !isSuccessStatus(openDoorResponse.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", openDoorResponse.Status, openDoorResponse.Message))
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", openDoorResponse),
		}).Debug("Door opened successfully")
	}

	return openDoorResponse, nil
}

func (c *OperationClientImpl) ListGoods(machineCode string) (*GetMachineGoodsResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithField("device_code", machineCode).Debug("Getting sold goods")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var getSoldGoodsResponse *GetMachineGoodsResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).
		SetQueryParam("code", machineCode).
		SetResult(&getSoldGoodsResponse).Get(Get_SoldGoods)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, ConvertGetSoldGoodsError(resp.StatusCode(), resp.String())
	}

	if !isSuccessStatus(getSoldGoodsResponse.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", getSoldGoodsResponse.Status, getSoldGoodsResponse.Message))
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", getSoldGoodsResponse),
		}).Debug("Got sold goods successfully")
	}

	return getSoldGoodsResponse, nil
}

func (c *OperationClientImpl) UpdateGoods(request *UpdateGoodsRequest, machineCode string) (*UpdateSoldGoodsResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request":     request,
			"device_code": machineCode,
		}).Debug("Updating sold goods")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var updateSoldGoodsResponse *UpdateSoldGoodsResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).
		SetQueryParam("code", machineCode).
		SetBody(request).SetResult(&updateSoldGoodsResponse).Post(Post_UpdateSoldGoods)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, ConvertUpdateSoldGoodsError(resp.StatusCode(), resp.String())
	}

	if !isSuccessStatus(updateSoldGoodsResponse.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", updateSoldGoodsResponse.Status, updateSoldGoodsResponse.Message))
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", updateSoldGoodsResponse),
		}).Debug("Updated sold goods successfully")
	}

	return updateSoldGoodsResponse, nil
}

func (c *OperationClientImpl) OpenDoorReqDetail(request *OpenDoorDetailRequest, machineCode string) (*OpenDoorDetailResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request":     request,
			"device_code": machineCode,
		}).Debug("Searching open door")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var searchOpenDoorResponse *OpenDoorDetailResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).
		SetQueryParam("code", machineCode).
		SetQueryParams(map[string]string{
			"type":      fmt.Sprintf("%d", request.Type),
			"requestId": request.RequestID,
		}).SetResult(&searchOpenDoorResponse).Get(Get_SearchOpenDoor)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, ConvertSearchOpenDoorError(resp.StatusCode(), resp.String())
	}

	if !isSuccessStatus(int(searchOpenDoorResponse.Status)) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", searchOpenDoorResponse.Status, searchOpenDoorResponse.Message))
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", searchOpenDoorResponse),
		}).Debug("Searched open door successfully")
	}

	return searchOpenDoorResponse, nil
}

func (c *OperationClientImpl) GetOrderVideo(request *GetOrderVideoRequest, machineCode string) (*GetOrderVideoResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request":     request,
			"device_code": machineCode,
		}).Debug("Getting order video")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var getOrderVideoResponse *GetOrderVideoResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).
		SetQueryParam("code", machineCode).
		SetQueryParams(map[string]string{
			"type":      fmt.Sprintf("%d", request.Type),
			"requestId": request.RequestID,
		}).SetResult(&getOrderVideoResponse).Get(Get_OrderVideo)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, ConvertGetOrderVideoError(resp.StatusCode(), resp.String())
	}

	if !isSuccessStatus(getOrderVideoResponse.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", getOrderVideoResponse.Status, getOrderVideoResponse.Message))
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", getOrderVideoResponse),
		}).Debug("Got order video successfully")
	}

	return getOrderVideoResponse, nil
}

func (c *OperationClientImpl) UpdateGoodsPrice(request *UpdateGoodsPriceRequest, machineCode string) (*ProductPriceUpdateResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request":     request,
			"device_code": machineCode,
		}).Debug("Updating product price")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var productPriceUpdateResponse *ProductPriceUpdateResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).
		SetQueryParam("code", machineCode).
		SetBody(request).SetResult(&productPriceUpdateResponse).Post(Post_ProductPriceUpdate)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, ConvertProductPriceUpdateError(resp.StatusCode(), resp.String())
	}

	if !isSuccessStatus(productPriceUpdateResponse.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", productPriceUpdateResponse.Status, productPriceUpdateResponse.Message))
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", productPriceUpdateResponse),
		}).Debug("Updated product price successfully")
	}

	return productPriceUpdateResponse, nil
}

func (c *OperationClientImpl) AddGoods(request *AddNewGoodsRequest, machineCode string) (*AddNewGoodsResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request":     request,
			"device_code": machineCode,
		}).Debug("Adding new goods")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var addNewGoodsResponse *AddNewGoodsResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).
		SetQueryParam("code", machineCode).
		SetBody(request.Items).SetResult(&addNewGoodsResponse).Put(Put_AddNewGoods)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, ConvertAddNewGoodsError(resp.StatusCode(), resp.String())
	}

	if !isSuccessStatus(addNewGoodsResponse.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", addNewGoodsResponse.Status, addNewGoodsResponse.Message))
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", addNewGoodsResponse),
		}).Debug("Added new goods successfully")
	}

	return addNewGoodsResponse, nil
}

func (c *OperationClientImpl) DeleteGoods(request *DeleteGoodsRequest, machineCode string) (*DeleteGoodsResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request":     request,
			"device_code": machineCode,
		}).Debug("Deleting goods")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	url := fmt.Sprintf("%s?code=%s", Del_DeleteGoods, machineCode)
	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	req.Header.Set("Authorization", signature)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, NewAinfinitError(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.StatusCode >= 400 {
		return nil, ConvertDeleteGoodsError(resp.StatusCode, string(body))
	}

	var deleteGoodsResponse *DeleteGoodsResponse
	if err := json.Unmarshal(body, &deleteGoodsResponse); err != nil {
		return nil, NewAinfinitError(err)
	}

	if !isSuccessStatus(deleteGoodsResponse.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", deleteGoodsResponse.Status, deleteGoodsResponse.Message))
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", deleteGoodsResponse),
		}).Debug("Deleted goods successfully")
	}

	return deleteGoodsResponse, nil
}

// ListOrder implements OperationClient.
func (c *OperationClientImpl) ListOrders(request *ListOrderRequest, machineCode string) (*ListOrderResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request":     request,
			"device_code": machineCode,
		}).Debug("Listing orders")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var listOrderResponse *ListOrderResponse
	query := c.Resty.R().SetHeader("Authorization", signature).
		SetQueryParam("code", machineCode)
		// SetQueryParam("beginTime", fmt.Sprintf("%d", request.BeginTime)).
		// SetQueryParam("endTime", fmt.Sprintf("%d", request.EndTime)).
		// SetQueryParam("page", fmt.Sprintf("%d", request.Page)).
		// SetQueryParam("limit", fmt.Sprintf("%d", request.Limit)).

	if request.BeginTime != 0 {
		query.SetQueryParam("beginTime", fmt.Sprintf("%d", request.BeginTime))
	}

	if request.EndTime != 0 {
		query.SetQueryParam("endTime", fmt.Sprintf("%d", request.EndTime))
	}

	if request.Page != 0 {
		query.SetQueryParam("page", fmt.Sprintf("%d", request.Page))
	}

	if request.Limit != 0 {
		query.SetQueryParam("limit", fmt.Sprintf("%d", request.Limit))
	}

	resp, err := query.SetResult(&listOrderResponse).Get(Get_ListOrders)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String()))
	}

	if !isSuccessStatus(listOrderResponse.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", listOrderResponse.Status, listOrderResponse.Message))
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", listOrderResponse),
		}).Debug("Listed orders successfully")
	}

	return listOrderResponse, nil
}

type OrderGoods struct {
	ItemCode  string `json:"itemCode"`  // Product code
	ItemName  string `json:"itemName"`  // Product name
	ItemPrice int    `json:"itemPrice"` // Commodity prices
	Count     int    `json:"count"`     // Quantity of goods
}

type OrderCallbackResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Goods struct {
	ItemCode      string `json:"itemCode,omitempty"`
	ActualPrice   int    `json:"actualPrice,omitempty"`
	OriginalPrice int    `json:"originalPrice,omitempty"`
	Count         int    `json:"count,omitempty"`
}

type Order struct {
	TradeRequestId  string  `json:"tradeRequestId,omitempty"`
	OrderCode       string  `json:"orderCode,omitempty"`
	VmCode          string  `json:"vmCode,omitempty"`
	MachineId       int     `json:"machineId,omitempty"`
	UserCode        string  `json:"userCode,omitempty"`
	HandleStatus    int     `json:"handleStatus,omitempty"`
	ShopMove        int     `json:"shopMove,omitempty"`
	TotalFee        int     `json:"totalFee,omitempty"`
	OpenDoorTime    int64   `json:"openDoorTime,omitempty"`
	CloseDoorTime   int64   `json:"closeDoorTime,omitempty"`
	OpenDoorWeight  int     `json:"openDoorWeight,omitempty"`
	CloseDoorWeight int     `json:"closeDoorWeight,omitempty"`
	OrderGoodsList  []Goods `json:"orderGoodsList,omitempty"`
}

type OpenDoorType int

const (
	OpenDoorForShopping      OpenDoorType = 1
	OpenDoorForReplenishment OpenDoorType = 2
)

type VideoStatus int

const (
	VideoStatusPendingUpload     VideoStatus = -1
	VideoStatusUploadComplete    VideoStatus = 0
	VideoStatusVideoDoesNotExist VideoStatus = 1
	VideoStatusNetworkError      VideoStatus = 2
	VideoStatusUploadInProgress  VideoStatus = 3
)

type OpenDoorRequest struct {
	Type           OpenDoorType `json:"type,omitempty" validate:"required"`      // 1: shopping, 2: replenishment
	RequestID      string       `json:"requestId,omitempty" validate:"required"` // oruulj ogno
	UserCode       string       `json:"userCode,omitempty"`
	LocalTimeStamp int64        `json:"localTimestamp,omitempty"`
}

type UpdateGoodsRequest []Goods

type OpenDoorDetailRequest struct {
	Type      OpenDoorType `json:"type,omitempty"`
	RequestID string       `json:"requestId,omitempty"`
}

type ListOrderRequest struct {
	BeginTime int64 `json:"beginTime,omitempty"`
	EndTime   int64 `json:"endTime,omitempty"`
	Page      int   `json:"page,omitempty"`  //default 1
	Limit     int   `json:"limit,omitempty"` //default 10 max 50
}

type GetOrderVideoRequest struct {
	RequestID string       `json:"requestId,omitempty"`
	Type      OpenDoorType `json:"type,omitempty"`
}

type UpdateGoodsPriceRequest struct {
	VmCodes []string `json:"vmCodes,omitempty"`
	Items   []Goods  `json:"items,omitempty"`
}

type AddNewGoodsRequest struct {
	Items []Goods `json:"items,omitempty"`
}

type DeleteGoodsRequest struct {
	ItemCodes []string `json:"itemCodes,omitempty"`
}

type OpenDoorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		OrderCode string `json:"orderCode"`
	} `json:"data"`
}

type GetMachineGoodsResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Result  []Goods `json:"result"`
	Count   int     `json:"count"`
}

type SearchOpenDoorData struct {
	TradeRequestId  string  `json:"tradeRequestId"`
	OrderCode       string  `json:"orderCode"`
	VmCode          string  `json:"vmCode"`
	MachineId       int     `json:"machineId"`
	HandleStatus    int     `json:"handleStatus"`
	TotalFee        int     `json:"totalFee"`
	OpenDoorTime    int64   `json:"openDoorTime"`
	CloseDoorTime   int64   `json:"closeDoorTime"`
	OpenDoorWeight  int     `json:"openDoorWeight"`
	CloseDoorWeight int     `json:"closeDoorWeight"`
	OrderGoodsList  []Goods `json:"orderGoodsList"`
	ShopMove        int     `json:"shopMove"`
	ScanCode        string  `json:"scanCode"`
}

type OpenDoorDetailResponse struct {
	Status  DoorOpenCloseStatus `json:"status"`
	Message string              `json:"message"`
	Data    SearchOpenDoorData  `json:"data"`
}

type ListOrderResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Total int     `json:"total"`
		Rows  []Order `json:"rows"`
	} `json:"data"`
}

type GetOrderVideoResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		OrderCode   string      `json:"orderCode"`
		VideoUrl    string      `json:"videoUrl"`
		VideoURLs   []string    `json:"videoUrls"`
		VideoStatus VideoStatus `json:"videoStatus"`
	} `json:"data"`
}

type ProductPriceUpdateResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type AddNewGoodsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type DeleteGoodsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Ok      *bool  `json:"ok,omitempty"`
}

// Door Open/Close Status Code
type DoorOpenCloseStatus int

const (
	// Success Status Codes
	DoorOpenCloseStatusOpened DoorOpenCloseStatus = 201 // Door opened successfully
	DoorOpenCloseStatusClosed DoorOpenCloseStatus = 202 // Door closed successfully

	// Shopping/Restocking Conflict Status Codes
	DoorOpenCloseStatusShoppingNotFinished   DoorOpenCloseStatus = 2031 // Failed to open - previous shopping not finished
	DoorOpenCloseStatusRestockingNotFinished DoorOpenCloseStatus = 2032 // Failed to open - previous restocking not finished

	// Device State Status Codes
	DoorOpenCloseStatusPowerOff          DoorOpenCloseStatus = 2033 // Failed to open - device power off, running on UPS
	DoorOpenCloseStatusMaintenanceMode   DoorOpenCloseStatus = 2034 // Failed to open - device in maintenance mode
	DoorOpenCloseStatusBackgroundProcess DoorOpenCloseStatus = 204  // Failed to open - device background process active

	// Timeout and Error Status Codes
	DoorOpenCloseStatusTimeout      DoorOpenCloseStatus = 503 // Failed to open - device response timeout (20s)
	DoorOpenCloseStatusNoResult     DoorOpenCloseStatus = 504 // Failed to open - no result reported within 5 minutes
	DoorOpenCloseStatusUnknownError DoorOpenCloseStatus = 505 // Failed to open - unknown error
	DoorOpenCloseStatusCalibration  DoorOpenCloseStatus = 506 // Failed to open - calibration error

	// Hardware and Verification Status Codes
	DoorOpenCloseStatusProductVerification DoorOpenCloseStatus = 5050 // Failed to open - product verification failed
	DoorOpenCloseStatusSerialPortFault     DoorOpenCloseStatus = 5051 // Failed to open - serial port fault
	DoorOpenCloseStatusWeightSensorFault   DoorOpenCloseStatus = 5052 // Failed to open - weight sensor fault
	DoorOpenCloseStatusCamerasOffline      DoorOpenCloseStatus = 5053 // Failed to open - all cameras offline
	DoorOpenCloseStatusAlgorithmError      DoorOpenCloseStatus = 5054 // Failed to open - local recognition algorithm error
	DoorOpenCloseStatusDoorLockError       DoorOpenCloseStatus = 5055 // Failed to open - door lock error
	DoorOpenCloseStatusPowerStatusError    DoorOpenCloseStatus = 5056 // Failed to open - device power status error

	// Door Lock State Status Codes
	DoorOpenCloseStatusDoorOpenLockOpen     DoorOpenCloseStatus = 5057 // Door lock error - door open + lock open
	DoorOpenCloseStatusDoorClosedLockOpen   DoorOpenCloseStatus = 5058 // Door lock error - door closed + lock open
	DoorOpenCloseStatusDoorOpenLockClosed   DoorOpenCloseStatus = 5059 // Door lock error - door open + lock closed
	DoorOpenCloseStatusDoorClosedLockClosed DoorOpenCloseStatus = 5060 // Door lock error - door closed + lock closed

	// Request Status Codes
	DoorOpenCloseStatusNoResultYet     DoorOpenCloseStatus = 404   // No result reported by device yet
	DoorOpenCloseStatusRequestNotFound DoorOpenCloseStatus = 42404 // Request ID does not exist
	DoorOpenCloseStatusInvalidType     DoorOpenCloseStatus = 40005 // Invalid parameter: type
	DoorOpenCloseStatusTooManyOrders   DoorOpenCloseStatus = 40526 // Too many shopping orders in progress
	DoorOpenCloseStatusNoPermission    DoorOpenCloseStatus = 42403 // No permission to query
)
