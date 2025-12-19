package aifinitsdk

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"resty.dev/v3"
)

type AdvertisementManageClient interface {
	MaterialApply(request *SourceMaterialApplyRequest) (*SourceMaterialApplyResponse, error)
	MaterialPage(request *SourceMaterialPageRequest) (*SourceMaterialPageResponse, error)
	MaterialDetail(materialId int) (*SourceMaterialDetailResponse, error)
	MaterialDelete(materialId int) (*SourceMaterialDeleteResponse, error)

	AdAddition(request *AdAdditionRequest) (*AdAdditionResponse, error)
	AdPage(request *AdPageRequest) (*AdPageResponse, error)
	AdDetailByAdId(adId int) (*AdDetailResponse, error)
	AdDetailByVmCode(code string) (*AdDetailResponse, error)
	AdUpdate(request *AdUpdateRequest) (*AdUpdateResponse, error)
	AdDelete(adId int) (*AdDeleteResponse, error)

	AdAssociatedToVm(adId int, request *AdAssociatedToVmRequest) (*AdAssociatedToVmResponse, error)
	ControlAdStatus(promotionId int, status AdStatus) (*AdControlStatusResponse, error)
	GetVmPromotion(vmCode string) (*GetVmPromotionResponse, error)
	//callback
	MediaReviewNotify()
}

type advertisementManageClientImpl struct {
	Client Client
	Resty  *resty.Client
}

func NewAdvertisementManageClient(client Client) AdvertisementManageClient {
	restyClient := client.GetRestyClient()
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
		return nil, NewAinfinitError(err)
	}

	var result SourceMaterialApplyResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).
		SetBody(request.SourceMaterialList).
		SetResult(&result).
		Post(Post_AdvertisementMaterialApply)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(errors.New(resp.String()))
	}

	if !isSuccessStatus(result.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", result.Status, result.Message))
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
		return nil, NewAinfinitError(err)
	}

	var result SourceMaterialPageResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetBody(request).SetResult(&result).Get(Get_AdvertisementMaterialPage)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(errors.New(resp.String()))
	}

	if !isSuccessStatus(result.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", result.Status, result.Message))
	}

	return &result, nil
}

func (c *advertisementManageClientImpl) MaterialDetail(materialId int) (*SourceMaterialDetailResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"materialId": strconv.Itoa(materialId),
		}).Debug("Getting source material detail")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var result SourceMaterialDetailResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetPathParam("id", strconv.Itoa(materialId)).SetResult(&result).Get(Get_AdvertisementMaterialDetail)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(errors.New(resp.String()))
	}

	if !isSuccessStatus(result.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", result.Status, result.Message))
	}

	return &result, nil
}

func (c *advertisementManageClientImpl) MaterialDelete(materialId int) (*SourceMaterialDeleteResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"materialId": strconv.Itoa(materialId),
		}).Debug("Deleting source material")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var result SourceMaterialDeleteResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetPathParam("id", strconv.Itoa(materialId)).SetResult(&result).Delete(Del_AdvertisementMaterialDelete)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(errors.New(resp.String()))
	}

	if !isSuccessStatus(result.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", result.Status, result.Message))
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
		return nil, NewAinfinitError(err)
	}

	var result AdAdditionResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetBody(request).SetResult(&result).Post(Post_AdvertisementAdAddition)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(errors.New(resp.String()))
	}

	if !isSuccessStatus(result.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", result.Status, result.Message))
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
		return nil, NewAinfinitError(err)
	}

	var result AdPageResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetQueryParams(map[string]string{
		"page":      strconv.Itoa(request.Page),
		"page_size": strconv.Itoa(request.PageSize),
	}).SetResult(&result).Get(Get_AdvertisementAdPage)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(errors.New(resp.String()))
	}

	if !isSuccessStatus(result.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", result.Status, result.Message))
	}

	return &result, nil
}

func (c *advertisementManageClientImpl) AdDetailByAdId(adId int) (*AdDetailResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"adId": strconv.Itoa(adId),
		}).Debug("Getting ad detail by ad id")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var result AdDetailResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetPathParam("id", strconv.Itoa(adId)).SetResult(&result).Get(Get_AdvertisementAdDetail)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(errors.New(resp.String()))
	}

	if !isSuccessStatus(result.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", result.Status, result.Message))
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
		return nil, NewAinfinitError(err)
	}

	var result AdDetailResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetQueryParam("code", vmCode).SetResult(&result).Get(Get_AdvertisementAdDetailByVmCode)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(errors.New(resp.String()))
	}

	if !isSuccessStatus(result.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", result.Status, result.Message))
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
		return nil, NewAinfinitError(err)
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var result AdUpdateResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetBody(request.Ad).SetResult(&result).Put(Put_AdvertisementAdUpdate)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(errors.New(resp.String()))
	}

	if !isSuccessStatus(result.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", result.Status, result.Message))
	}

	return &result, nil
}

func (c *advertisementManageClientImpl) AdDelete(adId int) (*AdDeleteResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"adId": strconv.Itoa(adId),
		}).Debug("Deleting ad")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var result AdDeleteResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetPathParam("id", strconv.Itoa(adId)).SetResult(&result).Delete(Del_AdvertisementAdDelete)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(errors.New(resp.String()))
	}

	if !isSuccessStatus(result.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", result.Status, result.Message))
	}

	return &result, nil
}

func (c *advertisementManageClientImpl) AdAssociatedToVm(adId int, request *AdAssociatedToVmRequest) (*AdAssociatedToVmResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"adId": strconv.Itoa(adId),
		}).Debug("Getting ad associated to vm")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var result AdAssociatedToVmResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetBody(request).SetPathParam("id", strconv.Itoa(adId)).SetResult(&result).Put(Put_AdvertisementAssociatedToVm)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(errors.New(resp.String()))
	}

	if !isSuccessStatus(result.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", result.Status, result.Message))
	}

	return &result, nil
}

func (c *advertisementManageClientImpl) ControlAdStatus(promotionId int, status AdStatus) (*AdControlStatusResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"promotionId": strconv.Itoa(promotionId),
			"status":      status,
		}).Debug("Controlling ad status")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var result AdControlStatusResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetPathParam("promotionId", strconv.Itoa(promotionId)).
		SetPathParam("status", strconv.Itoa(int(status))).SetResult(&result).Put(Get_AdvertisementAdAssociatedToVm)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(errors.New(resp.String()))
	}

	if !isSuccessStatus(result.Status) {
		return nil, ConvertAdvertisementError(result.Status, result.Message)
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
		return nil, NewAinfinitError(err)
	}

	var result GetVmPromotionResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetQueryParam("code", vmCode).SetResult(&result).Get(Get_AdvertisementVmPromotion)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(errors.New(resp.String()))
	}

	if !isSuccessStatus(result.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", result.Status, result.Message))
	}

	return &result, nil
}

func (c *advertisementManageClientImpl) MediaReviewNotify() {}

type FileType int

const (
	FileTypeImageResource FileType = 1
	FileTypeVideoResource FileType = 2
)

type BusinessType int

const (
	BusinessTypePublicService BusinessType = 1
	BusinessTypeCommercial    BusinessType = 2
)

type AdStatus int

const (
	AdStatusUnderReview AdStatus = 1 // Under review
	AdStatusApproved    AdStatus = 2 // Approved
	AdStatusRejected    AdStatus = 3 // Rejected
)

type SourceMaterial struct {
	Id         int      `json:"id,omitempty"`
	FileUrl    string   `json:"fileUrl,omitempty"`
	FileType   FileType `json:"fileType,omitempty"`
	Name       string   `json:"name,omitempty"`
	Status     int      `json:"status,omitempty"`
	CreateTime string   `json:"createTime,omitempty"`
}

type ImgRel struct {
	Id                int      `json:"id"`
	Priority          int      `json:"priority"` // 1-100 smallest to largest
	PromotionId       int      `json:"promotionId"`
	FileType          FileType `json:"fileType"`
	FileUrl           string   `json:"fileUrl"`
	SourceMaterialsId int      `json:"sourceMaterialsId"`
}

type Vm struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Ad struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	BusinessType int      `json:"businessType"`
	Duration     int      `json:"duration"`
	Status       int      `json:"status"`
	CreateTime   string   `json:"createTime"`
	UpdateTime   string   `json:"updateTime"`
	ImgRelList   []ImgRel `json:"imgRelList"`
	VmList       []Vm     `json:"vmList"`
}

const (
	ErrCodeSuccess   = 200
	ErrCodeNotFound  = 404
	ErrCodeForbidden = 403
)

const (
	ErrCodeSourceMaterialNotFound     = 4440
	ErrCodeSourceMaterialNotAllowed   = 4441
	ErrCodeSourceMaterialDoesNotExist = 4448
)

const (
	ErrCodeAdvertisementNotFound     = 4450
	ErrCodeAdvertisementNotAllowed   = 4451
	ErrCodeAdvertisementInvalidInput = 4452
	ErrCodeAdRemoveNotAllowed        = 4446
)

const (
	ErrAdDetailNotFound   = 4440
	ErrAdDetailNotAllowed = 4441
)

type ErrAdDetail int

func (e ErrAdDetail) Error() string {
	return fmt.Sprintf("ErrAdDetail: %d", e)
}

func ConvertAdDetailError(code int, message string) error {
	switch code {
	case ErrAdDetailNotFound:
		return fmt.Errorf("AdDetailNotFound: %d, message: %s", code, message)
	case ErrAdDetailNotAllowed:
		return fmt.Errorf("AdDetailNotAllowed: %d, message: %s", code, message)
	default:
		return fmt.Errorf("AdDetailError: %d, message: %s", code, message)
	}
}

type SourceMaterialError int

func (e SourceMaterialError) Error() string {
	return fmt.Sprintf("SourceMaterialError: %d", e)
}

func ConvertSourceMaterialError(code int, message string) error {
	switch code {
	case ErrCodeSourceMaterialNotFound:
		return fmt.Errorf("SourceMaterialNotFound: %d, message: %s", code, message)
	case ErrCodeSourceMaterialNotAllowed:
		return fmt.Errorf("SourceMaterialNotAllowed: %d, message: %s", code, message)
	case ErrCodeSourceMaterialDoesNotExist:
		return fmt.Errorf("SourceMaterialDoesNotExist: %d, message: %s", code, message)
	default:
		return fmt.Errorf("SourceMaterialError: %d, message: %s", code, message)
	}
}

type AdvertisementError int

func (e AdvertisementError) Error() string {
	return fmt.Sprintf("AdvertisementError: %d", e)
}

func ConvertAdvertisementError(code int, message string) error {
	switch code {
	case ErrCodeAdvertisementNotFound:
		return fmt.Errorf("AdvertisementNotFound: %d, message: %s", code, message)
	case ErrCodeAdvertisementNotAllowed:
		return fmt.Errorf("AdvertisementNotAllowed: %d, message: %s", code, message)
	case ErrCodeAdvertisementInvalidInput:
		return fmt.Errorf("AdvertisementInvalidInput: %d, message: %s", code, message)
	case ErrCodeAdRemoveNotAllowed:
		return fmt.Errorf("AdRemoveNotAllowed: %d, message: %s", code, message)
	default:
		return fmt.Errorf("AdvertisementError: %d, message: %s", code, message)
	}
}

type SourceMaterialApplyRequest struct {
	SourceMaterialList []SourceMaterial `json:"source_material_list"`
}

type SourceMaterialPageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

type AdAdditionRequest struct {
	Name         string `json:"name,omitempty"`
	BusinessType int    `json:"businessType,omitempty"`
	Duration     int    `json:"duration,omitempty"`
	ImgRelList   []struct {
		Priority          int `json:"priority,omitempty"`
		SourceMaterialsId int `json:"sourceMaterialsId,omitempty"`
	} `json:"imgRelList,omitempty"`
}

type AdPageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

type AdUpdateRequest struct {
	Ad struct {
		Id           int      `json:"id,omitempty" validate:"required"`
		Name         string   `json:"name,omitempty"`
		BusinessType int      `json:"businessType,omitempty"`
		Duration     int      `json:"duration,omitempty"`
		ImgRelList   []ImgRel `json:"imgRelList,omitempty"`
		VmList       []Vm     `json:"vmList,omitempty"`
	} `json:"ad,omitempty"`
}

type AdAssociatedToVmRequest struct {
	VmList       []string `json:"vmList,omitempty"`       // vmCode that provided from vmdetail
	ScanCodeList []string `json:"scanCodeList,omitempty"` // ScanCodeList
}

type SourceMaterialApplyResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Count   int    `json:"count"`
	Result  []struct {
		Id int `json:"id"`
	} `json:"result"`
	Ok bool `json:"ok"`
}

type SourceMaterialPageResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Count   int    `json:"count"`
	Data    struct {
		Total int              `json:"total"`
		Rows  []SourceMaterial `json:"rows"`
	} `json:"data"`
}

type SourceMaterialDetailResponse struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    SourceMaterial `json:"data"`
}

type SourceMaterialDeleteResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

type AdAdditionResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
	Ok bool `json:"ok"`
}

type AdPageResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Total int  `json:"total"`
		Rows  []Ad `json:"rows"`
	} `json:"data"`
	Ok bool `json:"ok"`
}

type AdDetailResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    *Ad    `json:"data,omitempty"`
}

type AdUpdateResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
	Ok bool `json:"ok"`
}

type AdDeleteResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

type AdAssociatedToVmResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type AdControlStatusResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

type GetVmPromotionResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    *Ad    `json:"data,omitempty"`
}
