package aifinitsdk_operation

import (
	"fmt"
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

func NewOperationClientImpl(client aifinitsdk.Client, deviceCode string) *OperationClientImpl {
	return &OperationClientImpl{
		Client:     client,
		Resty:      resty.New(),
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
	resp, err := c.Resty.R().SetHeader("Authorization", signature).
		SetQueryParam("code", c.DeviceCode).
		SetQueryParams(map[string]string{
			"type":           fmt.Sprintf("%d", request.Type),
			"requestId":      request.RequestID,
			"userCode":       request.UserCode,
			"localTimestamp": fmt.Sprintf("%d", request.LocalTimeStamp),
		}).SetResult(&openDoorResponse).Put(aifinitsdk_constants.Put_OpenDoor)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, ConvertOpenDoorError(resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": openDoorResponse,
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
			"response": getSoldGoodsResponse,
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
			"response": updateSoldGoodsResponse,
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
			"response": searchOpenDoorResponse,
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
			"response": getOrderVideoResponse,
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
			"response": productPriceUpdateResponse,
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
		SetBody(request).SetResult(&addNewGoodsResponse).Put(aifinitsdk_constants.Put_AddNewGoods)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, ConvertAddNewGoodsError(resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": addNewGoodsResponse,
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

	var deleteGoodsResponse *DeleteGoodsResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).
		SetQueryParam("code", c.DeviceCode).
		SetBody(request).SetResult(&deleteGoodsResponse).Delete(aifinitsdk_constants.Del_DeleteGoods)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, ConvertDeleteGoodsError(resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": deleteGoodsResponse,
		}).Debug("Deleted goods successfully")
	}

	return deleteGoodsResponse, nil
}
