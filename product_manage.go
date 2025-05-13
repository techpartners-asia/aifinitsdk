package aifinitsdk

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"resty.dev/v3"
)

type ProductManageClient interface {
	LastInfo() (*LastInfoResponse, error)
	ProductList(page, limit int) (*ProductListResponse, error)
	ProductDetail(itemCode string) (*ProductDetailResponse, error)
	MutualExclusion(request *MutualExclusionRequest) (*MutualExclusionResponse, error)
	NewProductApplication(request *NewProductApplicationRequest) (*NewProductApplicationResponse, error)
	ListProductApplication(params *ListProductApplicationParams) (*ListProductApplicationResponse, error)
	DetailProductApplication(itemCode string) (*DetailProductApplicationResponse, error)
	UpdateProductApplication(itemCode string, request *UpdateProductApplicationRequest) (*UpdateProductApplicationResponse, error)
}

type ProductClient struct {
	Client Client
	Resty  *resty.Client
}

func NewProductClient(client Client) ProductManageClient {
	restyClient := resty.New()
	if client.RestyDebug() {
		restyClient.SetDebug(true)
	}

	return &ProductClient{
		Client: client,
		Resty:  restyClient,
	}
}

func (c *ProductClient) LastInfo() (*LastInfoResponse, error) {
	panic("unimplemented")
}

func (c *ProductClient) ProductList(page, limit int) (*ProductListResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"page":  page,
			"limit": limit,
		}).Debug("Getting product list")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var products *ProductListResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetResult(&products).Get(Get_ProductList)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", products),
		}).Debug("Got product list successfully")
	}

	return products, nil
}

func (c *ProductClient) ProductDetail(itemCode string) (*ProductDetailResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithField("item_code", itemCode).Debug("Getting product detail")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var product *ProductDetailResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetResult(&product).Get(fmt.Sprintf(Get_ProductDetail, itemCode))
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", product),
		}).Debug("Got product detail successfully")
	}

	return product, nil
}

func (c *ProductClient) MutualExclusion(request *MutualExclusionRequest) (*MutualExclusionResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithField("request", request).Debug("Getting product mutual exclusion")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var mutualExclusion *MutualExclusionResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetBody(request).SetResult(&mutualExclusion).Post(Post_ProductMutualExclusion)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", mutualExclusion),
		}).Debug("Got product mutual exclusion successfully")
	}

	return mutualExclusion, nil
}

// multipart/form-data
// item - product info, required
// file - product images
// files - physical map of the goods, at least 2 and the bar code clearly visible
// weightFile - Weight of pictures
func (c *ProductClient) NewProductApplication(request *NewProductApplicationRequest) (*NewProductApplicationResponse, error) {
	if request == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if request.Product == nil {
		return nil, fmt.Errorf("product cannot be nil")
	}

	if c.Client.IsDebug() {
		logrus.WithField("request", request).Debug("Creating new product application")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var newProductApplication *NewProductApplicationResponse
	req := c.Resty.R().
		SetHeader("Authorization", signature).
		SetMultipartField("item", "", "application/json", strings.NewReader(request.Product.String()))

	if len(request.Product.ImgFiles) != len(request.Product.ImgFileNames) {
		return nil, fmt.Errorf("image files and names must be the same length")
	}

	for i, img := range request.Product.ImgFiles {
		if img != nil {
			req = req.SetFileReader("file", request.Product.ImgFileNames[i], img)
		}
	}

	if len(request.Product.PhysicalImgFiles) != len(request.Product.PhysicalImgFileNames) {
		return nil, fmt.Errorf("actual image files and names must be the same length")
	}

	for i, img := range request.Product.PhysicalImgFiles {
		if img != nil {
			req = req.SetFileReader("files", request.Product.PhysicalImgFileNames[i], img)
		}
	}

	if request.Product.WeightFile != nil {
		req = req.SetFileReader("weightFile", request.Product.WeightFileName, request.Product.WeightFile)
	}

	resp, err := req.SetResult(&newProductApplication).Post(Post_NewProductApplication)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", newProductApplication),
		}).Debug("Created new product application successfully")
	}

	return newProductApplication, nil
}

func (c *ProductClient) ListProductApplication(params *ListProductApplicationParams) (*ListProductApplicationResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithField("params", params).Debug("Listing product applications")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var listProductApplication *ListProductApplicationResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetQueryParams(map[string]string{
		"page":        fmt.Sprintf("%d", params.Page),
		"pageSize":    fmt.Sprintf("%d", params.PageSize),
		"applyStatus": fmt.Sprintf("%d", params.ApplyStatus),
		"goodsName":   params.GoodsName,
		"qrCodes":     params.QrCodes,
	}).SetResult(&listProductApplication).Get(Get_ProductApplicationList)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", listProductApplication),
		}).Debug("Listed product applications successfully")
	}

	return listProductApplication, nil
}

func (c *ProductClient) DetailProductApplication(itemCode string) (*DetailProductApplicationResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithField("item_code", itemCode).Debug("Getting product application detail")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var detailProductApplication *DetailProductApplicationResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetResult(&detailProductApplication).Get(fmt.Sprintf(Get_ProductApplicationDetail, itemCode))
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", detailProductApplication),
		}).Debug("Got product application detail successfully")
	}

	return detailProductApplication, nil
}

func (c *ProductClient) UpdateProductApplication(itemCode string, request *UpdateProductApplicationRequest) (*UpdateProductApplicationResponse, error) {
	if c.Client.IsDebug() {
		logrus.WithField("request", request).Debug("Updating product application")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var updateProductApplication *UpdateProductApplicationResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetBody(request).SetResult(&updateProductApplication).Put(fmt.Sprintf(Put_UpdateProductAppication, itemCode))
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", updateProductApplication),
		}).Debug("Updated product application successfully")
	}

	return updateProductApplication, nil
}

// ENTITIES

type Product struct {
	Id             int      `json:"id"`
	Name           string   `json:"name"`
	Price          int      `json:"price"`
	Weight         int      `json:"weight"`
	WeightVariance int      `json:"weightVariance"`
	ImgUrl         string   `json:"imgUrl"`
	ItemCode       string   `json:"itemCode"`
	CollType       int      `json:"collType"` // collection type: 1- single, 2- multiple
	UpdateTime     string   `json:"updateTime"`
	CreateTime     string   `json:"createTime"`
	Status         int      `json:"status"`
	QrCodes        string   `json:"qrCodes"`
	ItemCodes      []string `json:"itemCodes"`
	ActualImgs     []string `json:"actualImgs"`
	WeightFile     string   `json:"weightFile"`
}

func (p *Product) String() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(bytes)
}

type NewProductApplication struct {
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Weight  int    `json:"weight"`
	QrCodes string `json:"qrCodes"`

	ImgFiles     []io.Reader `json:"-"` // product image files
	ImgFileNames []string    `json:"-"` // product image file names

	PhysicalImgFiles     []io.Reader `json:"-"` // physical image files IMPORTANT: at least 2 and bar code clearly visible
	PhysicalImgFileNames []string    `json:"-"` // physical image file names IMPORTANT: at least 2 and bar code clearly visible

	WeightFile     io.Reader `json:"-"` // docs: weight of pictures
	WeightFileName string    `json:"-"` // docs: weight of pictures
}

func (n *NewProductApplication) String() string {
	bytes, err := json.Marshal(n)
	if err != nil {
		return ""
	}
	return string(bytes)
}

type LastInfo struct {
	Count          int   `json:"count"`
	LastUpdateTime int64 `json:"lastUpdateTime"`
}

type ProductListPage struct {
	Page        int    `json:"page,omitempty"`
	Limit       int    `json:"limit,omitempty"`
	UpdatedTime string `json:"updatedTime,omitempty"`
	GoodsName   string `json:"goodsName,omitempty"`
	QrCodes     string `json:"qrCodes,omitempty"`
}

type MutualExclusionRequest struct {
	ItemCodes []string `json:"item_codes,omitempty"`
}

type NewProductApplicationRequest struct {
	Product *NewProductApplication `json:"item,omitempty"`
}

type ListProductApplicationParams struct {
	Page        int    `json:"page,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
	ApplyStatus int    `json:"applyStatus,omitempty"` // 1: review, 2: pass, 3: reject
	GoodsName   string `json:"goodsName,omitempty"`
	QrCodes     string `json:"qrCodes,omitempty"`
}

type UpdateProductApplicationRequest struct {
	Item *Product `json:"item,omitempty"`
}

type ProductListResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Total int       `json:"total"`
		Rows  []Product `json:"rows"`
	}
}

type LastInfoResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    LastInfo `json:"data"`
}

type ProductDetailResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Product `json:"data"`
}

type MutualExclusionResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Total int      `json:"total"`
		Rows  []string `json:"rows"`
	}
}

type NewProductApplicationResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    int    `json:"data"`
}

type ListProductApplicationResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Total int       `json:"total"`
		Rows  []Product `json:"rows"`
	}
}

type DetailProductApplicationResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Product `json:"data"`
}

type UpdateProductApplicationResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type DeleteGoodsError int

const (
	ErrDeleteGoodsSelfDealerNotExist     int = 40506
	ErrDeleteGoodsUnknownGoods           int = 40507
	ErrDeleteGoodsNoOperatingPermissions int = 40531
)

func (e DeleteGoodsError) Error() string {
	return fmt.Sprintf("DeleteGoodsError: %d", e)
}

func ConvertDeleteGoodsError(code int, message string) error {
	switch code {
	case ErrDeleteGoodsSelfDealerNotExist:
		return fmt.Errorf("DeleteGoodsSelfDealerNotExist: %d, message: %s", code, message)
	case ErrDeleteGoodsUnknownGoods:
		return fmt.Errorf("DeleteGoodsUnknownGoods: %d, message: %s", code, message)
	case ErrDeleteGoodsNoOperatingPermissions:
		return fmt.Errorf("DeleteGoodsNoOperatingPermissions: %d, message: %s", code, message)
	default:
		return fmt.Errorf("DeleteGoodsError: %d, message: %s", code, message)
	}
}

type AddNewGoodsError int

const (
	ErrAddNewGoodsTooManyGoods           int = 10004
	ErrAddNewGoodsDuplicateGoods         int = 40502
	ErrAddNewGoodsMutuallyExclusiveGoods int = 40503
	ErrAddNewGoodsDownloadedGoods        int = 40504
	ErrAddNewGoodsSelfDealerNotExist     int = 40506
	ErrAddNewGoodsUnknownGoods           int = 40507
	ErrAddNewGoodsNoOperatingPermissions int = 40531
)

func (e AddNewGoodsError) Error() string {
	return fmt.Sprintf("AddNewGoodsError: %d", e)
}

func ConvertAddNewGoodsError(code int, message string) error {
	switch code {
	case ErrAddNewGoodsTooManyGoods:
		return fmt.Errorf("AddNewGoodsTooManyGoods: %d, message: %s", code, message)
	case ErrAddNewGoodsDuplicateGoods:
		return fmt.Errorf("AddNewGoodsDuplicateGoods: %d, message: %s", code, message)
	case ErrAddNewGoodsMutuallyExclusiveGoods:
		return fmt.Errorf("AddNewGoodsMutuallyExclusiveGoods: %d, message: %s", code, message)
	case ErrAddNewGoodsDownloadedGoods:
		return fmt.Errorf("AddNewGoodsDownloadedGoods: %d, message: %s", code, message)
	case ErrAddNewGoodsSelfDealerNotExist:
		return fmt.Errorf("AddNewGoodsSelfDealerNotExist: %d, message: %s", code, message)
	case ErrAddNewGoodsUnknownGoods:
		return fmt.Errorf("AddNewGoodsUnknownGoods: %d, message: %s", code, message)
	case ErrAddNewGoodsNoOperatingPermissions:
		return fmt.Errorf("AddNewGoodsNoOperatingPermissions: %d, message: %s", code, message)
	default:
		return fmt.Errorf("AddNewGoodsError: %d, message: %s", code, message)
	}
}

type ProductPriceUpdateError int

const (
	ErrProductPriceUpdateVendingMachineDoesNotExistTargetGoods int = 3501
	ErrProductPriceUpdateThereAreDuplicateProducts             int = 40502
	ErrProductPriceUpdateThereAreDownloadedGoods               int = 40504
	ErrProductPriceUpdateTheSelfDealerDoesNotExist             int = 40506
	ErrProductPriceUpdateThereAreUnknownProducts               int = 40507
	ErrProductPriceUpdateNoOperatingPermissions                int = 40531
)

func (e ProductPriceUpdateError) Error() string {
	return fmt.Sprintf("ProductPriceUpdateError: %d", e)
}

func ConvertProductPriceUpdateError(code int, message string) error {
	switch code {
	case ErrProductPriceUpdateVendingMachineDoesNotExistTargetGoods:
		return fmt.Errorf("ProductPriceUpdateVendingMachineDoesNotExistTargetGoods: %d, message: %s", code, message)
	case ErrProductPriceUpdateThereAreDuplicateProducts:
		return fmt.Errorf("ProductPriceUpdateThereAreDuplicateProducts: %d, message: %s", code, message)
	case ErrProductPriceUpdateThereAreDownloadedGoods:
		return fmt.Errorf("ProductPriceUpdateThereAreDownloadedGoods: %d, message: %s", code, message)
	case ErrProductPriceUpdateTheSelfDealerDoesNotExist:
		return fmt.Errorf("ProductPriceUpdateTheSelfDealerDoesNotExist: %d, message: %s", code, message)
	case ErrProductPriceUpdateThereAreUnknownProducts:
		return fmt.Errorf("ProductPriceUpdateThereAreUnknownProducts: %d, message: %s", code, message)
	case ErrProductPriceUpdateNoOperatingPermissions:
		return fmt.Errorf("ProductPriceUpdateNoOperatingPermissions: %d, message: %s", code, message)
	default:
		return fmt.Errorf("ProductPriceUpdateError: %d, message: %s", code, message)
	}
}

type GetOrderVideoError int

const (
	ErrGetOrderVideoSuccess                            int = 200
	ErrGetOrderVideoNoOrderOrReplenishmentRecordsFound int = 404
	ErrGetOrderVideoTheOpeningRequestDoesNotExist      int = 42404
)

func (e GetOrderVideoError) Error() string {
	return fmt.Sprintf("GetOrderVideoError: %d", e)
}

func ConvertGetOrderVideoError(code int, message string) error {
	switch code {
	case ErrGetOrderVideoSuccess:
		return fmt.Errorf("GetOrderVideoSuccess: %d, message: %s", code, message)
	case ErrGetOrderVideoNoOrderOrReplenishmentRecordsFound:
		return fmt.Errorf("GetOrderVideoNoOrderOrReplenishmentRecordsFound: %d, message: %s", code, message)
	case ErrGetOrderVideoTheOpeningRequestDoesNotExist:
		return fmt.Errorf("GetOrderVideoTheOpeningRequestDoesNotExist: %d, message: %s", code, message)
	default:
		return fmt.Errorf("GetOrderVideoError: %d, message: %s", code, message)
	}
}

type SearchOpenDoorError int

const (
	ErrSearchOpenDoorSuccess                                            int = 201
	ErrSearchOpenDoorClosingSuccess                                     int = 202
	ErrSearchOpenDoorLastShoppingNotOver                                int = 2031
	ErrSearchOpenDoorLastReplenishmentNotOver                           int = 2032
	ErrSearchOpenDoorEquipmentPoweredOff                                int = 2033
	ErrSearchOpenDoorEquipmentInOperationMode                           int = 2034
	ErrSearchOpenDoorDoorOpenedFailed                                   int = 204
	ErrSearchOpenDoorDeviceBackgroundProcess                            int = 503
	ErrSearchOpenDoorDeviceReceivedMessageTimeout                       int = 504
	ErrSearchOpenDoorUnknownError                                       int = 505
	ErrSearchOpenDoorDebuggingInformationIncorrect                      int = 506
	ErrSearchOpenDoorVerificationOfSaleOfPlanningProductsFailed         int = 5050
	ErrSearchOpenDoorFailedDoorOpeningEquipmentSerialFailure            int = 5051
	ErrSearchOpenDoorEquipmentHeavyFaults                               int = 5052
	ErrSearchOpenDoorDeviceCameraAllDropped                             int = 5053
	ErrSearchOpenDoorFailedDoorOpeningLocalRecognitionAlgorithmAbnormal int = 5054
	ErrSearchOpenDoorFailedToOpenTheDoorTheLockWasUnusual               int = 5055
	ErrSearchOpenDoorEquipmentPowerSupplyStatusError                    int = 5056
	ErrSearchOpenDoorDoorLockAbnormalTheDoorIsOpen                      int = 5057
	ErrSearchOpenDoorDoorLockAbnormalTheDoorIsClosed                    int = 5058
	ErrSearchOpenDoorDoorLockAbnormalTheDoorIsClosed2                   int = 5059
	ErrSearchOpenDoorDoorLockAbnormalTheDoorIsClosed3                   int = 5060
	ErrSearchOpenDoorEquipmentHasNotReportedTheResults                  int = 404
	ErrSearchOpenDoorDoorRequestIdDoesNotExist                          int = 42404
	ErrSearchOpenDoorTypeParameterTypeError                             int = 40005
	ErrSearchOpenDoorTooManyOrdersInTheShop                             int = 40526
	ErrSearchOpenDoorNoSearchPermissions                                int = 42403
)

func (e SearchOpenDoorError) Error() string {
	return fmt.Sprintf("SearchOpenDoorError: %d", e)
}

func ConvertSearchOpenDoorError(code int, message string) error {
	switch code {
	case ErrSearchOpenDoorSuccess:
		return fmt.Errorf("SearchOpenDoorSuccess: %d, message: %s", code, message)
	case ErrSearchOpenDoorClosingSuccess:
		return fmt.Errorf("SearchOpenDoorClosingSuccess: %d, message: %s", code, message)
	case ErrSearchOpenDoorLastShoppingNotOver:
		return fmt.Errorf("SearchOpenDoorLastShoppingNotOver: %d, message: %s", code, message)
	case ErrSearchOpenDoorLastReplenishmentNotOver:
		return fmt.Errorf("SearchOpenDoorLastReplenishmentNotOver: %d, message: %s", code, message)
	case ErrSearchOpenDoorEquipmentPoweredOff:
		return fmt.Errorf("SearchOpenDoorEquipmentPoweredOff: %d, message: %s", code, message)
	case ErrSearchOpenDoorEquipmentInOperationMode:
		return fmt.Errorf("SearchOpenDoorEquipmentInOperationMode: %d, message: %s", code, message)
	case ErrSearchOpenDoorDoorOpenedFailed:
		return fmt.Errorf("SearchOpenDoorDoorOpenedFailed: %d, message: %s", code, message)
	case ErrSearchOpenDoorDeviceBackgroundProcess:
		return fmt.Errorf("SearchOpenDoorDeviceBackgroundProcess: %d, message: %s", code, message)
	case ErrSearchOpenDoorDeviceReceivedMessageTimeout:
		return fmt.Errorf("SearchOpenDoorDeviceReceivedMessageTimeout: %d, message: %s", code, message)
	case ErrSearchOpenDoorUnknownError:
		return fmt.Errorf("SearchOpenDoorUnknownError: %d, message: %s", code, message)
	case ErrSearchOpenDoorDebuggingInformationIncorrect:
		return fmt.Errorf("SearchOpenDoorDebuggingInformationIncorrect: %d, message: %s", code, message)
	case ErrSearchOpenDoorVerificationOfSaleOfPlanningProductsFailed:
		return fmt.Errorf("SearchOpenDoorVerificationOfSaleOfPlanningProductsFailed: %d, message: %s", code, message)
	case ErrSearchOpenDoorFailedDoorOpeningEquipmentSerialFailure:
		return fmt.Errorf("SearchOpenDoorFailedDoorOpeningEquipmentSerialFailure: %d, message: %s", code, message)
	case ErrSearchOpenDoorEquipmentHeavyFaults:
		return fmt.Errorf("SearchOpenDoorEquipmentHeavyFaults: %d, message: %s", code, message)
	case ErrSearchOpenDoorDeviceCameraAllDropped:
		return fmt.Errorf("SearchOpenDoorDeviceCameraAllDropped: %d, message: %s", code, message)
	case ErrSearchOpenDoorFailedDoorOpeningLocalRecognitionAlgorithmAbnormal:
		return fmt.Errorf("SearchOpenDoorFailedDoorOpeningLocalRecognitionAlgorithmAbnormal: %d, message: %s", code, message)
	case ErrSearchOpenDoorFailedToOpenTheDoorTheLockWasUnusual:
		return fmt.Errorf("SearchOpenDoorFailedToOpenTheDoorTheLockWasUnusual: %d, message: %s", code, message)
	case ErrSearchOpenDoorEquipmentPowerSupplyStatusError:
		return fmt.Errorf("SearchOpenDoorEquipmentPowerSupplyStatusError: %d, message: %s", code, message)
	case ErrSearchOpenDoorDoorLockAbnormalTheDoorIsOpen:
		return fmt.Errorf("SearchOpenDoorDoorLockAbnormalTheDoorIsOpen: %d, message: %s", code, message)
	case ErrSearchOpenDoorDoorLockAbnormalTheDoorIsClosed:
		return fmt.Errorf("SearchOpenDoorDoorLockAbnormalTheDoorIsClosed: %d, message: %s", code, message)
	case ErrSearchOpenDoorDoorLockAbnormalTheDoorIsClosed2:
		return fmt.Errorf("SearchOpenDoorDoorLockAbnormalTheDoorIsClosed2: %d, message: %s", code, message)
	case ErrSearchOpenDoorDoorLockAbnormalTheDoorIsClosed3:
		return fmt.Errorf("SearchOpenDoorDoorLockAbnormalTheDoorIsClosed3: %d, message: %s", code, message)
	case ErrSearchOpenDoorEquipmentHasNotReportedTheResults:
		return fmt.Errorf("SearchOpenDoorEquipmentHasNotReportedTheResults: %d, message: %s", code, message)
	case ErrSearchOpenDoorDoorRequestIdDoesNotExist:
		return fmt.Errorf("SearchOpenDoorDoorRequestIdDoesNotExist: %d, message: %s", code, message)
	case ErrSearchOpenDoorTypeParameterTypeError:
		return fmt.Errorf("SearchOpenDoorTypeParameterTypeError: %d, message: %s", code, message)
	case ErrSearchOpenDoorTooManyOrdersInTheShop:
		return fmt.Errorf("SearchOpenDoorTooManyOrdersInTheShop: %d, message: %s", code, message)
	case ErrSearchOpenDoorNoSearchPermissions:
		return fmt.Errorf("SearchOpenDoorNoSearchPermissions: %d, message: %s", code, message)
	default:
		return fmt.Errorf("SearchOpenDoorError: %d, message: %s", code, message)
	}
}

type UpdateSoldGoodsError int

const (
	ErrUpdateSoldGoodsTooManyGoods           int = 10004
	ErrUpdateSoldGoodsDuplicateGoods         int = 40502
	ErrUpdateSoldGoodsMutuallyExclusiveGoods int = 40503
	ErrUpdateSoldGoodsDownloadedGoods        int = 40504
	ErrUpdateSoldGoodsSelfDealerNotExist     int = 40506
	ErrUpdateSoldGoodsUnknownGoods           int = 40507
	ErrUpdateSoldGoodsNoOperatingPermissions int = 40531
)

func (e UpdateSoldGoodsError) Error() string {
	return fmt.Sprintf("UpdateSoldGoodsError: %d", e)
}

func ConvertUpdateSoldGoodsError(code int, message string) error {
	switch code {
	case ErrUpdateSoldGoodsTooManyGoods:
		return fmt.Errorf("UpdateSoldGoodsTooManyGoods: %d, message: %s", code, message)
	case ErrUpdateSoldGoodsDuplicateGoods:
		return fmt.Errorf("UpdateSoldGoodsDuplicateGoods: %d, message: %s", code, message)
	case ErrUpdateSoldGoodsMutuallyExclusiveGoods:
		return fmt.Errorf("UpdateSoldGoodsMutuallyExclusiveGoods: %d, message: %s", code, message)
	case ErrUpdateSoldGoodsDownloadedGoods:
		return fmt.Errorf("UpdateSoldGoodsDownloadedGoods: %d, message: %s", code, message)
	case ErrUpdateSoldGoodsSelfDealerNotExist:
		return fmt.Errorf("UpdateSoldGoodsSelfDealerNotExist: %d, message: %s", code, message)
	case ErrUpdateSoldGoodsUnknownGoods:
		return fmt.Errorf("UpdateSoldGoodsUnknownGoods: %d, message: %s", code, message)
	case ErrUpdateSoldGoodsNoOperatingPermissions:
		return fmt.Errorf("UpdateSoldGoodsNoOperatingPermissions: %d, message: %s", code, message)
	default:
		return fmt.Errorf("UpdateSoldGoodsError: %d, message: %s", code, message)
	}
}

type UpdateSoldGoodsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type GetSoldGoodsError int

const (
	ErrGetSoldGoodsSelfDealerNotExist            int = 40506
	ErrGetSoldGoodsSelfDealerNotBelongToMerchant int = 40531
)

func (e GetSoldGoodsError) Error() string {
	return fmt.Sprintf("GetSoldGoodsError: %d", e)
}

func ConvertGetSoldGoodsError(code int, message string) error {
	switch code {
	case ErrGetSoldGoodsSelfDealerNotExist:
		return fmt.Errorf("GetSoldGoodsSelfDealerNotExist: %d, message: %s", code, message)
	case ErrGetSoldGoodsSelfDealerNotBelongToMerchant:
		return fmt.Errorf("GetSoldGoodsSelfDealerNotBelongToMerchant: %d, message: %s", code, message)
	default:
		return fmt.Errorf("GetSoldGoodsError: %d, message: %s", code, message)
	}
}

type OpenDoorError int

const (
	ErrOpenDoorSuccess                      int = 200
	ErrOpenDoorFailed                       int = 400
	ErrOpenDoorTimeout                      int = 503
	ErrOpenDoorUnusualMachinePackage        int = 3501
	ErrOpenDoorOfflineEquipment             int = 10416
	ErrOpenDoorSelfDealerNotInOperation     int = 40525
	ErrOpenDoorTooManyOrdersNotCompleted    int = 40526
	ErrOpenDoorNonBusinessSelfSellerMachine int = 40531
)

func (e OpenDoorError) Error() string {
	return fmt.Sprintf("OpenDoorError: %d", e)
}

func ConvertOpenDoorError(code int, message string) error {
	switch code {
	case ErrOpenDoorSuccess:
		return fmt.Errorf("OpenDoorSuccess: %d, message: %s", code, message)
	case ErrOpenDoorFailed:
		return fmt.Errorf("OpenDoorFailed: %d, message: %s", code, message)
	case ErrOpenDoorTimeout:
		return fmt.Errorf("OpenDoorTimeout: %d, message: %s", code, message)
	case ErrOpenDoorUnusualMachinePackage:
		return fmt.Errorf("OpenDoorUnusualMachinePackage: %d, message: %s", code, message)
	case ErrOpenDoorOfflineEquipment:
		return fmt.Errorf("OpenDoorOfflineEquipment: %d, message: %s", code, message)
	case ErrOpenDoorSelfDealerNotInOperation:
		return fmt.Errorf("OpenDoorSelfDealerNotInOperation: %d, message: %s", code, message)
	case ErrOpenDoorTooManyOrdersNotCompleted:
		return fmt.Errorf("OpenDoorTooManyOrdersNotCompleted: %d, message: %s", code, message)
	case ErrOpenDoorNonBusinessSelfSellerMachine:
		return fmt.Errorf("OpenDoorNonBusinessSelfSellerMachine: %d, message: %s", code, message)
	default:
		return fmt.Errorf("OpenDoorError: %d, message: %s", code, message)
	}
}
