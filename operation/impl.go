package aifinitsdk_operation

import (
	"fmt"
	"time"

	aifinitsdk "git.techpartners.asia/mtm/thirdparty/aifinitsdk"
	aifinitsdk_constants "git.techpartners.asia/mtm/thirdparty/aifinitsdk/constants"
	"github.com/go-playground/validator/v10"
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

	return openDoorResponse, nil
}

func (c *OperationClientImpl) GetSoldGoods() (*GetSoldGoodsResponse, error) {
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

	return getSoldGoodsResponse, nil
}
func (c *OperationClientImpl) UpdateSoldGoods(request *UpdateSoldGoodsRequest) (*UpdateSoldGoodsResponse, error) {
	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var updateSoldGoodsResponse *UpdateSoldGoodsResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).
		SetQueryParam("code", c.DeviceCode).
		SetBody(request).SetResult(&updateSoldGoodsResponse).Put(aifinitsdk_constants.Put_UpdateSoldGoods)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, ConvertUpdateSoldGoodsError(resp.StatusCode(), resp.String())
	}

	return updateSoldGoodsResponse, nil
}

func (c *OperationClientImpl) SearchOpenDoor(request *SearchOpenDoorRequest) (*SearchOpenDoorResponse, error) {
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

	return searchOpenDoorResponse, nil
}

func (c *OperationClientImpl) GetOrderVideo(request *GetOrderVideoRequest) (*GetOrderVideoResponse, error) {
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

	return getOrderVideoResponse, nil
}

func (c *OperationClientImpl) ProductPriceUpdate(request *ProductPriceUpdateRequest) (*ProductPriceUpdateResponse, error) {
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

	return productPriceUpdateResponse, nil
}
