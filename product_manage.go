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
	if c.Client.IsDebug() {
		logrus.Debug("Getting last info")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var lastInfo *LastInfoResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetResult(&lastInfo).Get(Get_ProductLastInfo)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String()))
	}

	if !isSuccessStatus(lastInfo.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", lastInfo.Status, lastInfo.Message))
	}

	if c.Client.IsDebug() {
		logrus.WithFields(logrus.Fields{
			"response": fmt.Sprintf("%+v", lastInfo),
		}).Debug("Got last info successfully")
	}

	return lastInfo, nil
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
		return nil, NewAinfinitError(err)
	}

	var products *ProductListResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetResult(&products).Get(Get_ProductList)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String()))
	}

	if !isSuccessStatus(products.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", products.Status, products.Message))
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
		return nil, NewAinfinitError(err)
	}

	var product *ProductDetailResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetResult(&product).Get(fmt.Sprintf(Get_ProductDetail, itemCode))
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String()))
	}

	if !isSuccessStatus(product.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", product.Status, product.Message))
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
		return nil, NewAinfinitError(err)
	}

	var mutualExclusion *MutualExclusionResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetBody(request).SetResult(&mutualExclusion).Post(Post_ProductMutualExclusion)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String()))
	}

	if !isSuccessStatus(mutualExclusion.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", mutualExclusion.Status, mutualExclusion.Message))
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
		return nil, NewAinfinitError(fmt.Errorf("request cannot be nil"))
	}
	if request.Product == nil {
		return nil, NewAinfinitError(fmt.Errorf("product cannot be nil"))
	}

	if c.Client.IsDebug() {
		logrus.WithField("request", request).Debug("Creating new product application")
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	var newProductApplication *NewProductApplicationResponse
	req := c.Resty.R().
		SetHeader("Authorization", signature).
		SetHeader("Content-Type", "multipart/form-data")

	// Add the JSON data as a form field
	req = req.SetMultipartField("item", "", "application/json", strings.NewReader(request.Product.String()))

	if len(request.Product.ImgFiles) != len(request.Product.ImgFileNames) {
		return nil, NewAinfinitError(fmt.Errorf("image files and names must be the same length"))
	}

	for i, img := range request.Product.ImgFiles {
		if img != nil {
			req = req.SetFileReader("file", request.Product.ImgFileNames[i], img)
		}
	}

	if len(request.Product.PhysicalImgFiles) != len(request.Product.PhysicalImgFileNames) {
		return nil, NewAinfinitError(fmt.Errorf("actual image files and names must be the same length"))
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
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String()))
	}

	if !isSuccessStatus(newProductApplication.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", newProductApplication.Status, newProductApplication.Message))
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
		return nil, NewAinfinitError(err)
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
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String()))
	}

	if !isSuccessStatus(listProductApplication.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", listProductApplication.Status, listProductApplication.Message))
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
		return nil, NewAinfinitError(err)
	}

	var detailProductApplication *DetailProductApplicationResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetResult(&detailProductApplication).Get(fmt.Sprintf(Get_ProductApplicationDetail, itemCode))
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String()))
	}

	if !isSuccessStatus(detailProductApplication.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", detailProductApplication.Status, detailProductApplication.Message))
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
		return nil, NewAinfinitError(err)
	}

	var updateProductApplication *UpdateProductApplicationResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetBody(request).SetResult(&updateProductApplication).Put(Put_UpdateProductAppication)
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	if resp.IsError() {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String()))
	}

	if !isSuccessStatus(updateProductApplication.Status) {
		return nil, NewAinfinitError(fmt.Errorf("status: %d, message: %s", updateProductApplication.Status, updateProductApplication.Message))
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
	Id             int      `json:"id"`             // Unique product ID
	Name           string   `json:"name"`           // Product name
	Price          int      `json:"price"`          // Suggested retail price, in cents
	Weight         int      `json:"weight"`         // Product weight in grams
	WeightVariance int      `json:"weightVariance"` // Acceptable weight variance in grams
	ImgUrl         string   `json:"imgUrl"`         // URL to the main product image
	ItemCode       string   `json:"itemCode"`       // Unique product code
	CollType       int      `json:"collType"`       // Collection type: 1 - single item, 2 - collection/multiple items
	UpdateTime     string   `json:"updateTime"`     // Last update time in "YYYY-MM-DD HH:MM:SS" format
	CreateTime     string   `json:"createTime"`     // Creation time in "YYYY-MM-DD HH:MM:SS" format
	Status         int      `json:"status"`         // Product status: 1 - available, 2 - unavailable
	QrCodes        string   `json:"qrCodes"`        // Comma-separated list of barcodes (e.g., "6920180209601,6920180209602")
	ItemCodes      []string `json:"itemCodes"`      // List of item codes included in this product if it's a collection
	ActualImgs     []string `json:"actualImgs"`     // List of URLs to actual/real product images
	WeightFile     string   `json:"weightFile"`     // Path or URL to a file containing detailed weight data (if any)
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
		return NewAinfinitError(fmt.Errorf("DeleteGoodsSelfDealerNotExist: %d, message: %s", code, message))
	case ErrDeleteGoodsUnknownGoods:
		return NewAinfinitError(fmt.Errorf("DeleteGoodsUnknownGoods: %d, message: %s", code, message))
	case ErrDeleteGoodsNoOperatingPermissions:
		return NewAinfinitError(fmt.Errorf("DeleteGoodsNoOperatingPermissions: %d, message: %s", code, message))
	default:
		return NewAinfinitError(fmt.Errorf("DeleteGoodsError: %d, message: %s", code, message))
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
		return NewAinfinitError(fmt.Errorf("AddNewGoodsTooManyGoods: %d, message: %s", code, message))
	case ErrAddNewGoodsDuplicateGoods:
		return NewAinfinitError(fmt.Errorf("AddNewGoodsDuplicateGoods: %d, message: %s", code, message))
	case ErrAddNewGoodsMutuallyExclusiveGoods:
		return NewAinfinitError(fmt.Errorf("AddNewGoodsMutuallyExclusiveGoods: %d, message: %s", code, message))
	case ErrAddNewGoodsDownloadedGoods:
		return NewAinfinitError(fmt.Errorf("AddNewGoodsDownloadedGoods: %d, message: %s", code, message))
	case ErrAddNewGoodsSelfDealerNotExist:
		return NewAinfinitError(fmt.Errorf("AddNewGoodsSelfDealerNotExist: %d, message: %s", code, message))
	case ErrAddNewGoodsUnknownGoods:
		return NewAinfinitError(fmt.Errorf("AddNewGoodsUnknownGoods: %d, message: %s", code, message))
	case ErrAddNewGoodsNoOperatingPermissions:
		return NewAinfinitError(fmt.Errorf("AddNewGoodsNoOperatingPermissions: %d, message: %s", code, message))
	default:
		return NewAinfinitError(fmt.Errorf("AddNewGoodsError: %d, message: %s", code, message))
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
		return NewAinfinitError(fmt.Errorf("ProductPriceUpdateVendingMachineDoesNotExistTargetGoods: %d, message: %s", code, message))
	case ErrProductPriceUpdateThereAreDuplicateProducts:
		return NewAinfinitError(fmt.Errorf("ProductPriceUpdateThereAreDuplicateProducts: %d, message: %s", code, message))
	case ErrProductPriceUpdateThereAreDownloadedGoods:
		return NewAinfinitError(fmt.Errorf("ProductPriceUpdateThereAreDownloadedGoods: %d, message: %s", code, message))
	case ErrProductPriceUpdateTheSelfDealerDoesNotExist:
		return NewAinfinitError(fmt.Errorf("ProductPriceUpdateTheSelfDealerDoesNotExist: %d, message: %s", code, message))
	case ErrProductPriceUpdateThereAreUnknownProducts:
		return NewAinfinitError(fmt.Errorf("ProductPriceUpdateThereAreUnknownProducts: %d, message: %s", code, message))
	case ErrProductPriceUpdateNoOperatingPermissions:
		return NewAinfinitError(fmt.Errorf("ProductPriceUpdateNoOperatingPermissions: %d, message: %s", code, message))
	default:
		return NewAinfinitError(fmt.Errorf("ProductPriceUpdateError: %d, message: %s", code, message))
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
		return NewAinfinitError(fmt.Errorf("GetOrderVideoSuccess: %d, message: %s", code, message))
	case ErrGetOrderVideoNoOrderOrReplenishmentRecordsFound:
		return NewAinfinitError(fmt.Errorf("GetOrderVideoNoOrderOrReplenishmentRecordsFound: %d, message: %s", code, message))
	case ErrGetOrderVideoTheOpeningRequestDoesNotExist:
		return NewAinfinitError(fmt.Errorf("GetOrderVideoTheOpeningRequestDoesNotExist: %d, message: %s", code, message))
	default:
		return NewAinfinitError(fmt.Errorf("GetOrderVideoError: %d, message: %s", code, message))
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
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorSuccess: %d, message: %s", code, message))
	case ErrSearchOpenDoorClosingSuccess:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorClosingSuccess: %d, message: %s", code, message))
	case ErrSearchOpenDoorLastShoppingNotOver:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorLastShoppingNotOver: %d, message: %s", code, message))
	case ErrSearchOpenDoorLastReplenishmentNotOver:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorLastReplenishmentNotOver: %d, message: %s", code, message))
	case ErrSearchOpenDoorEquipmentPoweredOff:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorEquipmentPoweredOff: %d, message: %s", code, message))
	case ErrSearchOpenDoorEquipmentInOperationMode:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorEquipmentInOperationMode: %d, message: %s", code, message))
	case ErrSearchOpenDoorDoorOpenedFailed:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorDoorOpenedFailed: %d, message: %s", code, message))
	case ErrSearchOpenDoorDeviceBackgroundProcess:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorDeviceBackgroundProcess: %d, message: %s", code, message))
	case ErrSearchOpenDoorDeviceReceivedMessageTimeout:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorDeviceReceivedMessageTimeout: %d, message: %s", code, message))
	case ErrSearchOpenDoorUnknownError:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorUnknownError: %d, message: %s", code, message))
	case ErrSearchOpenDoorDebuggingInformationIncorrect:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorDebuggingInformationIncorrect: %d, message: %s", code, message))
	case ErrSearchOpenDoorVerificationOfSaleOfPlanningProductsFailed:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorVerificationOfSaleOfPlanningProductsFailed: %d, message: %s", code, message))
	case ErrSearchOpenDoorFailedDoorOpeningEquipmentSerialFailure:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorFailedDoorOpeningEquipmentSerialFailure: %d, message: %s", code, message))
	case ErrSearchOpenDoorEquipmentHeavyFaults:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorEquipmentHeavyFaults: %d, message: %s", code, message))
	case ErrSearchOpenDoorDeviceCameraAllDropped:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorDeviceCameraAllDropped: %d, message: %s", code, message))
	case ErrSearchOpenDoorFailedDoorOpeningLocalRecognitionAlgorithmAbnormal:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorFailedDoorOpeningLocalRecognitionAlgorithmAbnormal: %d, message: %s", code, message))
	case ErrSearchOpenDoorFailedToOpenTheDoorTheLockWasUnusual:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorFailedToOpenTheDoorTheLockWasUnusual: %d, message: %s", code, message))
	case ErrSearchOpenDoorEquipmentPowerSupplyStatusError:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorEquipmentPowerSupplyStatusError: %d, message: %s", code, message))
	case ErrSearchOpenDoorDoorLockAbnormalTheDoorIsOpen:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorDoorLockAbnormalTheDoorIsOpen: %d, message: %s", code, message))
	case ErrSearchOpenDoorDoorLockAbnormalTheDoorIsClosed:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorDoorLockAbnormalTheDoorIsClosed: %d, message: %s", code, message))
	case ErrSearchOpenDoorDoorLockAbnormalTheDoorIsClosed2:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorDoorLockAbnormalTheDoorIsClosed2: %d, message: %s", code, message))
	case ErrSearchOpenDoorDoorLockAbnormalTheDoorIsClosed3:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorDoorLockAbnormalTheDoorIsClosed3: %d, message: %s", code, message))
	case ErrSearchOpenDoorEquipmentHasNotReportedTheResults:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorEquipmentHasNotReportedTheResults: %d, message: %s", code, message))
	case ErrSearchOpenDoorDoorRequestIdDoesNotExist:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorDoorRequestIdDoesNotExist: %d, message: %s", code, message))
	case ErrSearchOpenDoorTypeParameterTypeError:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorTypeParameterTypeError: %d, message: %s", code, message))
	case ErrSearchOpenDoorTooManyOrdersInTheShop:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorTooManyOrdersInTheShop: %d, message: %s", code, message))
	case ErrSearchOpenDoorNoSearchPermissions:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorNoSearchPermissions: %d, message: %s", code, message))
	default:
		return NewAinfinitError(fmt.Errorf("SearchOpenDoorError: %d, message: %s", code, message))
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
		return NewAinfinitError(fmt.Errorf("UpdateSoldGoodsTooManyGoods: %d, message: %s", code, message))
	case ErrUpdateSoldGoodsDuplicateGoods:
		return NewAinfinitError(fmt.Errorf("UpdateSoldGoodsDuplicateGoods: %d, message: %s", code, message))
	case ErrUpdateSoldGoodsMutuallyExclusiveGoods:
		return NewAinfinitError(fmt.Errorf("UpdateSoldGoodsMutuallyExclusiveGoods: %d, message: %s", code, message))
	case ErrUpdateSoldGoodsDownloadedGoods:
		return NewAinfinitError(fmt.Errorf("UpdateSoldGoodsDownloadedGoods: %d, message: %s", code, message))
	case ErrUpdateSoldGoodsSelfDealerNotExist:
		return NewAinfinitError(fmt.Errorf("UpdateSoldGoodsSelfDealerNotExist: %d, message: %s", code, message))
	case ErrUpdateSoldGoodsUnknownGoods:
		return NewAinfinitError(fmt.Errorf("UpdateSoldGoodsUnknownGoods: %d, message: %s", code, message))
	case ErrUpdateSoldGoodsNoOperatingPermissions:
		return NewAinfinitError(fmt.Errorf("UpdateSoldGoodsNoOperatingPermissions: %d, message: %s", code, message))
	default:
		return NewAinfinitError(fmt.Errorf("UpdateSoldGoodsError: %d, message: %s", code, message))
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
		return NewAinfinitError(fmt.Errorf("GetSoldGoodsSelfDealerNotExist: %d, message: %s", code, message))
	case ErrGetSoldGoodsSelfDealerNotBelongToMerchant:
		return NewAinfinitError(fmt.Errorf("GetSoldGoodsSelfDealerNotBelongToMerchant: %d, message: %s", code, message))
	default:
		return NewAinfinitError(fmt.Errorf("GetSoldGoodsError: %d, message: %s", code, message))
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
		return NewAinfinitError(fmt.Errorf("OpenDoorSuccess: %d, message: %s", code, message))
	case ErrOpenDoorFailed:
		return NewAinfinitError(fmt.Errorf("OpenDoorFailed: %d, message: %s", code, message))
	case ErrOpenDoorTimeout:
		return NewAinfinitError(fmt.Errorf("OpenDoorTimeout: %d, message: %s", code, message))
	case ErrOpenDoorUnusualMachinePackage:
		return NewAinfinitError(fmt.Errorf("OpenDoorUnusualMachinePackage: %d, message: %s", code, message))
	case ErrOpenDoorOfflineEquipment:
		return NewAinfinitError(fmt.Errorf("OpenDoorOfflineEquipment: %d, message: %s", code, message))
	case ErrOpenDoorSelfDealerNotInOperation:
		return NewAinfinitError(fmt.Errorf("OpenDoorSelfDealerNotInOperation: %d, message: %s", code, message))
	case ErrOpenDoorTooManyOrdersNotCompleted:
		return NewAinfinitError(fmt.Errorf("OpenDoorTooManyOrdersNotCompleted: %d, message: %s", code, message))
	case ErrOpenDoorNonBusinessSelfSellerMachine:
		return NewAinfinitError(fmt.Errorf("OpenDoorNonBusinessSelfSellerMachine: %d, message: %s", code, message))
	default:
		return NewAinfinitError(fmt.Errorf("OpenDoorError: %d, message: %s", code, message))
	}
}
