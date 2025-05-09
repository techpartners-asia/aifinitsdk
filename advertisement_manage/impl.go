package advertisementmanage

import (
	"errors"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"github.com/techpartners-asia/aifinitsdk"
	aifinitsdk_constants "github.com/techpartners-asia/aifinitsdk/constants"
	"resty.dev/v3"
)

type advertisementManageClientImpl struct {
	Client aifinitsdk.Client
	Resty  *resty.Client
}

func NewAdvertisementManageClient(client aifinitsdk.Client) AdvertisementManageClient {
	restyClient := resty.New()
	if client.RestyDebug() {
		restyClient.SetDebug(true)
	}
	return &advertisementManageClientImpl{
		Client: client,
		Resty:  restyClient,
	}
}

func (c *advertisementManageClientImpl) MaterialApply(request *SourceMaterialApplyRequest) (*SourceMaterialApplyResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request": request,
		}).Debug("Applying source material")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var result SourceMaterialApplyResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetBody(request).SetResult(&result).Post(aifinitsdk_constants.Post_AdvertisementMaterialApply)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, errors.New(resp.String())
	}

	return &result, nil
}

func (c *advertisementManageClientImpl) MaterialPage(request *SourceMaterialPageRequest) (*SourceMaterialPageResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request": request,
		}).Debug("Getting source material page")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var result SourceMaterialPageResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetBody(request).SetResult(&result).Post(aifinitsdk_constants.Get_AdvertisementMaterialPage)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, errors.New(resp.String())
	}

	return &result, nil
}

func (c *advertisementManageClientImpl) MaterialDetail(materialId string) (*SourceMaterialDetailResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"materialId": materialId,
		}).Debug("Getting source material detail")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var result SourceMaterialDetailResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetPathParam("id", materialId).SetResult(&result).Get(aifinitsdk_constants.Get_AdvertisementMaterialDetail)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, errors.New(resp.String())
	}

	return &result, nil
}

func (c *advertisementManageClientImpl) MaterialDelete(materialId string) (*SourceMaterialDeleteResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"materialId": materialId,
		}).Debug("Deleting source material")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var result SourceMaterialDeleteResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetPathParam("id", materialId).SetResult(&result).Delete(aifinitsdk_constants.Del_AdvertisementMaterialDelete)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, errors.New(resp.String())
	}

	return &result, nil
}

func (c *advertisementManageClientImpl) AdAddition(request *AdAdditionRequest) (*AdAdditionResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request": request,
		}).Debug("Adding ad")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var result AdAdditionResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetBody(request).SetResult(&result).Post(aifinitsdk_constants.Post_AdvertisementAdAddition)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, errors.New(resp.String())
	}

	return &result, nil
}

func (c *advertisementManageClientImpl) AdPage(request *AdPageRequest) (*AdPageResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request": request,
		}).Debug("Getting ad page")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var result AdPageResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetQueryParams(map[string]string{
		"page":     strconv.Itoa(request.Page),
		"pageSize": strconv.Itoa(request.PageSize),
	}).SetResult(&result).Get(aifinitsdk_constants.Get_AdvertisementAdPage)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, errors.New(resp.String())
	}

	return &result, nil
}

func (c *advertisementManageClientImpl) AdDetailByAdId(adId int) (*AdDetailResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"adId": adId,
		}).Debug("Getting ad detail by ad id")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var result AdDetailResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetPathParam("id", strconv.Itoa(adId)).SetResult(&result).Get(aifinitsdk_constants.Get_AdvertisementAdDetail)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, errors.New(resp.String())
	}

	return &result, nil
}

func (c *advertisementManageClientImpl) AdDetailByVmCode(vmCode string) (*AdDetailResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"vmCode": vmCode,
		}).Debug("Getting ad detail by vm code")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var result AdDetailResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetQueryParam("code", vmCode).SetResult(&result).Get(aifinitsdk_constants.Get_AdvertisementAdDetailByVmCode)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, errors.New(resp.String())
	}

	return &result, nil
}

func (c *advertisementManageClientImpl) AdUpdate(request *AdUpdateRequest) (*AdUpdateResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"request": request,
		}).Debug("Updating ad")
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return nil, err
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var result AdUpdateResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetBody(request).SetResult(&result).Put(aifinitsdk_constants.Put_AdvertisementAdUpdate)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, errors.New(resp.String())
	}

	return &result, nil
}

func (c *advertisementManageClientImpl) AdDelete(adId int) (*AdDeleteResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"adId": adId,
		}).Debug("Deleting ad")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var result AdDeleteResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetPathParam("id", strconv.Itoa(adId)).SetResult(&result).Delete(aifinitsdk_constants.Del_AdvertisementAdDelete)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, errors.New(resp.String())
	}

	return &result, nil
}

func (c *advertisementManageClientImpl) AdAssociatedToVm(adId int, request *AdAssociatedToVmRequest) (*AdAssociatedToVmResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"adId": adId,
		}).Debug("Getting ad associated to vm")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var result AdAssociatedToVmResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetBody(request).SetPathParam("id", strconv.Itoa(adId)).SetResult(&result).Put(aifinitsdk_constants.Put_AdvertisementAssociatedToVm)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, errors.New(resp.String())
	}

	return &result, nil
}

func (c *advertisementManageClientImpl) ControlAdStatus(promotionId int, status AdStatus) (*AdControlStatusResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"promotionId": promotionId,
			"status":      status,
		}).Debug("Controlling ad status")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var result AdControlStatusResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetPathParam("promotionId", strconv.Itoa(promotionId)).SetPathParam("status", strconv.Itoa(int(status))).SetResult(&result).Put(aifinitsdk_constants.Get_AdvertisementAdAssociatedToVm)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, errors.New(resp.String())
	}

	return &result, nil
}

func (c *advertisementManageClientImpl) GetVmPromotion(vmCode string) (*GetVmPromotionResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"vmCode": vmCode,
		}).Debug("Getting vm promotion")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var result GetVmPromotionResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetQueryParam("code", vmCode).SetResult(&result).Get(aifinitsdk_constants.Get_AdvertisementVmPromotion)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, errors.New(resp.String())
	}

	return &result, nil
}

func (c *advertisementManageClientImpl) MediaReviewNotify() {}
