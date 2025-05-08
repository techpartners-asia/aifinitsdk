package aifinitsdk_device

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	aifinitsdk "github.com/techpartners-asia/aifinitsdk"
	aifinitsdk_constants "github.com/techpartners-asia/aifinitsdk/constants"
	"resty.dev/v3"
)

type vendingMachineManageClient struct {
	Client aifinitsdk.Client
	Resty  *resty.Client
	code   string
}

func NewDeviceClient(client aifinitsdk.Client, code string) VendingMachineManageClient {
	restyClient := resty.New()
	if client.RestyDebug() {
		restyClient.SetDebug(true)
	}

	return &vendingMachineManageClient{
		Client: client,
		Resty:  restyClient,
		code:   code,
	}
}

func (c *vendingMachineManageClient) Update(request *UpdateRequest) (*UpdateResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request": request,
			"code":    c.code,
		}).Debug("Updating vending machine")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var result UpdateResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetBody(request).SetResult(&result).
		Put(aifinitsdk_constants.Put_UpdateVendingMachineInfo)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", result),
		}).Debug("Updated vending machine successfully")
	}

	return &result, nil
}

func (c *vendingMachineManageClient) Detail() (*DetailResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithField("code", c.code).Debug("Getting vending machine details")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var result DetailResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetResult(&result).
		SetQueryParam("code", c.code).
		Get(aifinitsdk_constants.Get_VendingMachineInfo)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", result),
		}).Debug("Got vending machine details successfully")
	}

	return &result, nil
}

func (c *vendingMachineManageClient) DeviceInfo() (*DeviceInfoResult, error) {
	if c.Client.IsDebug() {
		logrus.WithField("code", c.code).Debug("Getting device info")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var result DeviceInfoResult
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetResult(&result).
		SetQueryParam("code", c.code).
		Get(fmt.Sprintf("%s/facade/open/vending_machine/deviceInfo", aifinitsdk_constants.BaseURL))
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", result),
		}).Debug("Got device info successfully")
	}

	return &result, nil
}

func (c *vendingMachineManageClient) PeopleFlow(request *PeopleFlowRequest) (*PeopleFlowResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request": request,
			"code":    c.code,
		}).Debug("Getting people flow data")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var result PeopleFlowResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetBody(request).SetResult(&result).
		Post(aifinitsdk_constants.Post_VendingMachinePeopleFlow)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", result),
		}).Debug("Got people flow data successfully")
	}

	return &result, nil
}

func (c *vendingMachineManageClient) List(request *ListRequest) (*ListResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request": request,
		}).Debug("Listing vending machines")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var result ListResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetResult(&result).
		SetQueryParam("page", strconv.Itoa(request.Page)).
		SetQueryParam("limit", strconv.Itoa(request.Limit)).
		SetQueryParam("nameOf", request.NameOf).
		Get(fmt.Sprintf("%s/facade/open/vending_machine/infoPage", aifinitsdk_constants.BaseURL))
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", result),
		}).Debug("Listed vending machines successfully")
	}

	return &result, nil
}

func (c *vendingMachineManageClient) Control(request *ControlRequest) (*ControlResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request": request,
			"code":    c.code,
		}).Debug("Controlling vending machine")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var result ControlResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetBody(request).SetQueryParam("code", c.code).SetResult(&result).
		Put(aifinitsdk_constants.Put_VendingMachineDeviceControl)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", result),
		}).Debug("Controlled vending machine successfully")
	}

	return &result, nil
}

func (c *vendingMachineManageClient) Activation() {}
func (c *vendingMachineManageClient) Alarm()      {}
func (c *vendingMachineManageClient) Setting()    {}
