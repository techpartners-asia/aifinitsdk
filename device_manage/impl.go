package aifinitsdk_device

import (
	"fmt"
	"strconv"
	"time"

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
	return &vendingMachineManageClient{
		Client: client,
		Resty:  resty.New(),
		code:   code,
	}
}

func (c *vendingMachineManageClient) Update(request *UpdateRequest) (*UpdateResponse, error) {
	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var result UpdateResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetBody(request).SetResult(&result).
		Post(aifinitsdk_constants.Put_UpdateVendingMachineInfo)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
	}

	return &result, nil
}

func (c *vendingMachineManageClient) Detail() (*DetailResponse, error) {
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

	return &result, nil
}

func (c *vendingMachineManageClient) DeviceInfo() (*DeviceInfoResult, error) {
	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}
	fmt.Println(c.code)

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

	return &result, nil
}

func (c *vendingMachineManageClient) PeopleFlow(request *PeopleFlowRequest) (*PeopleFlowResponse, error) {
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

	return &result, nil
}

func (c *vendingMachineManageClient) List(request *ListRequest) (*ListResponse, error) {
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

	return &result, nil
}
func (c *vendingMachineManageClient) Control(request *ControlRequest) (*ControlResponse, error) {
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

	return &result, nil
}

func (c *vendingMachineManageClient) Activation() {}
func (c *vendingMachineManageClient) Alarm()      {}
func (c *vendingMachineManageClient) Setting()    {}
