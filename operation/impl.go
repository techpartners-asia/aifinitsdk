package aifinitsdk_operation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	aifinitsdk "github.com/techpartners-asia/aifinitsdk"
	aifinitsdk_constants "github.com/techpartners-asia/aifinitsdk/constants"
	"resty.dev/v3"
)

type OperationClientImpl struct {
	Client     aifinitsdk.Client
	Resty      *resty.Client
	DeviceCode string
}

func NewOperationClientImpl(client aifinitsdk.Client, deviceCode string) OperationClient {
	restyClient := resty.New()

	if client.RestyDebug() {
		restyClient.SetDebug(true)
	}

	return &OperationClientImpl{
		Client:     client,
		Resty:      restyClient,
		DeviceCode: deviceCode,
	}
}

func (c *OperationClientImpl) OpenDoor(request *OpenDoorRequest) (*OpenDoorResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request":     request,
			"device_code": c.DeviceCode,
		}).Debug("Opening door")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		return nil, err
	}

	var openDoorResponse *OpenDoorResponse
	req := c.Resty.R().SetHeader("Authorization", signature).
		SetQueryParam("code", c.DeviceCode).
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

	resp, err := req.Put(aifinitsdk_constants.Put_OpenDoor)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, ConvertOpenDoorError(resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", openDoorResponse),
		}).Debug("Door opened successfully")
	}

	return openDoorResponse, nil
}

func (c *OperationClientImpl) GetSoldGoods() (*GetSoldGoodsResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithField("device_code", c.DeviceCode).Debug("Getting sold goods")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var getSoldGoodsResponse *GetSoldGoodsResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).
		SetQueryParam("code", c.DeviceCode).
		SetResult(&getSoldGoodsResponse).Get(aifinitsdk_constants.Get_SoldGoods)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, ConvertGetSoldGoodsError(resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", getSoldGoodsResponse),
		}).Debug("Got sold goods successfully")
	}

	return getSoldGoodsResponse, nil
}

func (c *OperationClientImpl) UpdateSoldGoods(request *UpdateSoldGoodsRequest) (*UpdateSoldGoodsResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request":     request,
			"device_code": c.DeviceCode,
		}).Debug("Updating sold goods")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var updateSoldGoodsResponse *UpdateSoldGoodsResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).
		SetQueryParam("code", c.DeviceCode).
		SetBody(request).SetResult(&updateSoldGoodsResponse).Post(aifinitsdk_constants.Post_UpdateSoldGoods)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, ConvertUpdateSoldGoodsError(resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", updateSoldGoodsResponse),
		}).Debug("Updated sold goods successfully")
	}

	return updateSoldGoodsResponse, nil
}

func (c *OperationClientImpl) SearchOpenDoor(request *SearchOpenDoorRequest) (*SearchOpenDoorResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request":     request,
			"device_code": c.DeviceCode,
		}).Debug("Searching open door")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var searchOpenDoorResponse *SearchOpenDoorResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).
		SetQueryParam("code", c.DeviceCode).
		SetQueryParams(map[string]string{
			"type":      fmt.Sprintf("%d", request.Type),
			"requestId": request.RequestID,
		}).SetResult(&searchOpenDoorResponse).Post(aifinitsdk_constants.Get_SearchOpenDoor)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, ConvertSearchOpenDoorError(resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", searchOpenDoorResponse),
		}).Debug("Searched open door successfully")
	}

	return searchOpenDoorResponse, nil
}

func (c *OperationClientImpl) GetOrderVideo(request *GetOrderVideoRequest) (*GetOrderVideoResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request":     request,
			"device_code": c.DeviceCode,
		}).Debug("Getting order video")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var getOrderVideoResponse *GetOrderVideoResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).
		SetQueryParam("code", c.DeviceCode).
		SetQueryParams(map[string]string{
			"type":      fmt.Sprintf("%d", request.Type),
			"requestId": request.RequestID,
		}).SetResult(&getOrderVideoResponse).Post(aifinitsdk_constants.Get_OrderVideo)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, ConvertGetOrderVideoError(resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", getOrderVideoResponse),
		}).Debug("Got order video successfully")
	}

	return getOrderVideoResponse, nil
}

func (c *OperationClientImpl) ProductPriceUpdate(request *ProductPriceUpdateRequest) (*ProductPriceUpdateResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request":     request,
			"device_code": c.DeviceCode,
		}).Debug("Updating product price")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var productPriceUpdateResponse *ProductPriceUpdateResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).
		SetQueryParam("code", c.DeviceCode).
		SetBody(request).SetResult(&productPriceUpdateResponse).Post(aifinitsdk_constants.Post_ProductPriceUpdate)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, ConvertProductPriceUpdateError(resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", productPriceUpdateResponse),
		}).Debug("Updated product price successfully")
	}

	return productPriceUpdateResponse, nil
}

func (c *OperationClientImpl) AddNewGoods(request *AddNewGoodsRequest) (*AddNewGoodsResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request":     request,
			"device_code": c.DeviceCode,
		}).Debug("Adding new goods")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var addNewGoodsResponse *AddNewGoodsResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).
		SetQueryParam("code", c.DeviceCode).
		SetBody(request.Items).SetResult(&addNewGoodsResponse).Put(aifinitsdk_constants.Put_AddNewGoods)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, ConvertAddNewGoodsError(resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", addNewGoodsResponse),
		}).Debug("Added new goods successfully")
	}

	return addNewGoodsResponse, nil
}

func (c *OperationClientImpl) DeleteGoods(request *DeleteGoodsRequest) (*DeleteGoodsResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request":     request,
			"device_code": c.DeviceCode,
		}).Debug("Deleting goods")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s?code=%s", aifinitsdk_constants.Del_DeleteGoods, c.DeviceCode)
	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", signature)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, ConvertDeleteGoodsError(resp.StatusCode, string(body))
	}

	var deleteGoodsResponse *DeleteGoodsResponse
	if err := json.Unmarshal(body, &deleteGoodsResponse); err != nil {
		return nil, err
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", deleteGoodsResponse),
		}).Debug("Deleted goods successfully")
	}

	return deleteGoodsResponse, nil
}

// ListOrder implements OperationClient.
func (c *OperationClientImpl) ListOrder(request *ListOrderRequest) (*ListOrderResponse, error) {
	panic("unimplemented")
}

// SearchOpenDoorRequest implements OperationClient.
func (c *OperationClientImpl) SearchOpenDoorRequest(request *SearchOpenDoorRequest) (*SearchOpenDoorResponse, error) {
	panic("unimplemented")
}
