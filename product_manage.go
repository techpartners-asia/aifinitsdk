package aifinitsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	restyClient := client.GetRestyClient()
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
	if request.Product.Name == "" {
		return nil, NewAinfinitError(fmt.Errorf("name cannot be empty"))
	}

	if request.Product.Price <= 0 {
		return nil, NewAinfinitError(fmt.Errorf("price must be greater than 0"))
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
			req = req.SetFileReader("file", request.Product.ImgFileNames[i], bytes.NewReader(img))
		}
	}

	if len(request.Product.PhysicalImgFiles) != len(request.Product.PhysicalImgFileNames) {
		return nil, NewAinfinitError(fmt.Errorf("actual image files and names must be the same length"))
	}

	for i, img := range request.Product.PhysicalImgFiles {
		if img != nil {
			req = req.SetFileReader("files", request.Product.PhysicalImgFileNames[i], bytes.NewReader(img))
		}
	}

	if request.Product.WeightFile != nil {
		req = req.SetFileReader("weightFile", request.Product.WeightFileName, bytes.NewReader(request.Product.WeightFile))
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

	if request.Item.Id == 0 {
		return nil, NewAinfinitError(fmt.Errorf("id cannot be 0"))
	}

	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, NewAinfinitError(err)
	}

	req := c.Resty.R().
		SetHeader("Authorization", signature).
		SetHeader("Content-Type", "multipart/form-data")

	// Add the JSON data as a form field
	req = req.SetMultipartField("item", "", "application/json", strings.NewReader(request.Item.String()))

	if len(request.Item.ImgFiles) != len(request.Item.ImgFileNames) {
		return nil, NewAinfinitError(fmt.Errorf("image files and names must be the same length"))
	}

	for i, img := range request.Item.ImgFiles {
		if img != nil {
			req = req.SetFileReader("file", request.Item.ImgFileNames[i], bytes.NewReader(img))
		}
	}

	if len(request.Item.PhysicalImgFiles) != len(request.Item.PhysicalImgFileNames) {
		return nil, NewAinfinitError(fmt.Errorf("actual image files and names must be the same length"))
	}

	for i, img := range request.Item.PhysicalImgFiles {
		if img != nil {
			req = req.SetFileReader("files", request.Item.PhysicalImgFileNames[i], bytes.NewReader(img))
		}
	}

	if request.Item.WeightFile != nil {
		req = req.SetFileReader("weightFile", request.Item.WeightFileName, bytes.NewReader(request.Item.WeightFile))
	}

	var updateProductApplication *UpdateProductApplicationResponse
	resp, err := req.SetResult(&updateProductApplication).Put(Put_UpdateProductAppication)
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
	Id             int      `json:"id"`             // Application ID
	Name           string   `json:"name"`           // Product name
	Price          int      `json:"price"`          // Suggested retail price in cents
	Weight         int      `json:"weight"`         // Product weight in grams
	WeightVariance int      `json:"weightVariance"` // Acceptable weight variance in grams
	ImgUrl         string   `json:"imgUrl"`         // URL to the main product image
	ItemCode       string   `json:"itemCode"`       // Unique product code (only available after approval)
	CollType       int      `json:"collType"`       // Collection type: 1 - single item, 2 - collection/multiple items
	UpdateTime     string   `json:"updateTime"`     // Last update time in "YYYY-MM-DD HH:MM:SS" format
	CreateTime     string   `json:"createTime"`     // Creation time in "YYYY-MM-DD HH:MM:SS" format
	Status         int      `json:"status"`         // Product status: 1 - available, 2 - unavailable
	QrCodes        string   `json:"qrCodes"`        // Product barcode (e.g., "6934024500113")
	ItemCodes      []string `json:"itemCodes"`      // List of item codes included in this product if it's a collection
	ActualImgs     []string `json:"actualImgs"`     // List of URLs to actual/real product images
	WeightFile     string   `json:"weightFile"`     // Path or URL to a file containing detailed weight data
	ApplyStatus    int      `json:"applyStatus"`    // Application status: 1 - under review, 2 - approved, 3 - rejected
	ApplyTime      string   `json:"applyTime"`      // Application submission time in "YYYY-MM-DD HH:MM:SS" format
	RejectType     string   `json:"rejectType"`     // Rejection type: "1" - name non-compliant, "2" - barcode non-compliant, "3" - display image unclear, "4" - other
	RejectReason   string   `json:"rejectReason"`   // Rejection reason
	WeightImgUrl   string   `json:"weightImgUrl"`   // URL to the weight image
}

func (p *Product) String() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(bytes)
}

type UpdateProductApplication struct {
	Id      int    `json:"id"`
	Price   int    `json:"price"`
	Weight  int    `json:"weight"`
	QrCodes string `json:"qrCodes"`

	ImgFiles     [][]byte `json:"-"` // product image files
	ImgFileNames []string `json:"-"` // product image file names

	PhysicalImgFiles     [][]byte `json:"-"` // physical image files IMPORTANT: at least 2 and bar code clearly visible
	PhysicalImgFileNames []string `json:"-"` // physical image file names IMPORTANT: at least 2 and bar code clearly visible
	WeightFile           []byte   `json:"-"` // weight image file
	WeightFileName       string   `json:"-"` // weight image file name
}

func (u *UpdateProductApplication) String() string {
	bytes, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(bytes)
}

type NewProductApplication struct {
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Weight  float64 `json:"weight"`
	QrCodes string  `json:"qrCodes"`

	ImgFiles     [][]byte `json:"-"` // product image files
	ImgFileNames []string `json:"-"` // product image file names

	PhysicalImgFiles     [][]byte `json:"-"` // physical image files IMPORTANT: at least 2 and bar code clearly visible
	PhysicalImgFileNames []string `json:"-"` // physical image file names IMPORTANT: at least 2 and bar code clearly visible
	WeightFile           []byte   `json:"-"` // weight image file
	WeightFileName       string   `json:"-"` // weight image file name
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
	Item *UpdateProductApplication `json:"item,omitempty"`
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

func (e DeleteGoodsError) String() string {
	switch int(e) {
	case ErrDeleteGoodsSelfDealerNotExist:
		return "Self dealer does not exist"
	case ErrDeleteGoodsUnknownGoods:
		return "Unknown goods"
	case ErrDeleteGoodsNoOperatingPermissions:
		return "No operating permissions"
	default:
		return fmt.Sprintf("DeleteGoodsError: %d", e)
	}
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

func (e AddNewGoodsError) String() string {
	switch int(e) {
	case ErrAddNewGoodsTooManyGoods:
		return "Too many goods"
	case ErrAddNewGoodsDuplicateGoods:
		return "Duplicate goods"
	case ErrAddNewGoodsMutuallyExclusiveGoods:
		return "Mutually exclusive goods"
	case ErrAddNewGoodsDownloadedGoods:
		return "Downloaded goods"
	case ErrAddNewGoodsSelfDealerNotExist:
		return "Self dealer does not exist"
	case ErrAddNewGoodsUnknownGoods:
		return "Unknown goods"
	case ErrAddNewGoodsNoOperatingPermissions:
		return "No operating permissions"
	default:
		return fmt.Sprintf("AddNewGoodsError: %d", e)
	}
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

func (e ProductPriceUpdateError) String() string {
	switch int(e) {
	case ErrProductPriceUpdateVendingMachineDoesNotExistTargetGoods:
		return "Vending machine does not exist target goods"
	case ErrProductPriceUpdateThereAreDuplicateProducts:
		return "There are duplicate products"
	case ErrProductPriceUpdateThereAreDownloadedGoods:
		return "There are downloaded goods"
	case ErrProductPriceUpdateTheSelfDealerDoesNotExist:
		return "The self dealer does not exist"
	case ErrProductPriceUpdateThereAreUnknownProducts:
		return "There are unknown products"
	case ErrProductPriceUpdateNoOperatingPermissions:
		return "No operating permissions"
	default:
		return fmt.Sprintf("ProductPriceUpdateError: %d", e)
	}
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

func (e GetOrderVideoError) String() string {
	switch int(e) {
	case ErrGetOrderVideoSuccess:
		return "Success"
	case ErrGetOrderVideoNoOrderOrReplenishmentRecordsFound:
		return "No order or replenishment records found"
	case ErrGetOrderVideoTheOpeningRequestDoesNotExist:
		return "The opening request does not exist"
	default:
		return fmt.Sprintf("GetOrderVideoError: %d", e)
	}
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

func (e SearchOpenDoorError) String() string {
	switch int(e) {
	case ErrSearchOpenDoorSuccess:
		return "Open door success"
	case ErrSearchOpenDoorClosingSuccess:
		return "Closing success"
	case ErrSearchOpenDoorLastShoppingNotOver:
		return "Last shopping not over"
	case ErrSearchOpenDoorLastReplenishmentNotOver:
		return "Last replenishment not over"
	case ErrSearchOpenDoorEquipmentPoweredOff:
		return "Equipment powered off"
	case ErrSearchOpenDoorEquipmentInOperationMode:
		return "Equipment in operation mode"
	case ErrSearchOpenDoorDoorOpenedFailed:
		return "Door opened failed"
	case ErrSearchOpenDoorDeviceBackgroundProcess:
		return "Device background process"
	case ErrSearchOpenDoorDeviceReceivedMessageTimeout:
		return "Device received message timeout"
	case ErrSearchOpenDoorUnknownError:
		return "Unknown error"
	case ErrSearchOpenDoorDebuggingInformationIncorrect:
		return "Debugging information incorrect"
	case ErrSearchOpenDoorVerificationOfSaleOfPlanningProductsFailed:
		return "Verification of sale of planning products failed"
	case ErrSearchOpenDoorFailedDoorOpeningEquipmentSerialFailure:
		return "Failed door opening equipment serial failure"
	case ErrSearchOpenDoorEquipmentHeavyFaults:
		return "Equipment heavy faults"
	case ErrSearchOpenDoorDeviceCameraAllDropped:
		return "Device camera all dropped"
	case ErrSearchOpenDoorFailedDoorOpeningLocalRecognitionAlgorithmAbnormal:
		return "Failed door opening local recognition algorithm abnormal"
	case ErrSearchOpenDoorFailedToOpenTheDoorTheLockWasUnusual:
		return "Failed to open the door, the lock was unusual"
	case ErrSearchOpenDoorEquipmentPowerSupplyStatusError:
		return "Equipment power supply status error"
	case ErrSearchOpenDoorDoorLockAbnormalTheDoorIsOpen:
		return "Door lock abnormal, the door is open"
	case ErrSearchOpenDoorDoorLockAbnormalTheDoorIsClosed:
		return "Door lock abnormal, the door is closed"
	case ErrSearchOpenDoorDoorLockAbnormalTheDoorIsClosed2:
		return "Door lock abnormal, the door is closed (2)"
	case ErrSearchOpenDoorDoorLockAbnormalTheDoorIsClosed3:
		return "Door lock abnormal, the door is closed (3)"
	case ErrSearchOpenDoorEquipmentHasNotReportedTheResults:
		return "Equipment has not reported the results"
	case ErrSearchOpenDoorDoorRequestIdDoesNotExist:
		return "Door request ID does not exist"
	case ErrSearchOpenDoorTypeParameterTypeError:
		return "Type parameter type error"
	case ErrSearchOpenDoorTooManyOrdersInTheShop:
		return "Too many orders in the shop"
	case ErrSearchOpenDoorNoSearchPermissions:
		return "No search permissions"
	default:
		return fmt.Sprintf("SearchOpenDoorError: %d", e)
	}
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

func (e UpdateSoldGoodsError) String() string {
	switch int(e) {
	case ErrUpdateSoldGoodsTooManyGoods:
		return "Too many goods"
	case ErrUpdateSoldGoodsDuplicateGoods:
		return "Duplicate goods"
	case ErrUpdateSoldGoodsMutuallyExclusiveGoods:
		return "Mutually exclusive goods"
	case ErrUpdateSoldGoodsDownloadedGoods:
		return "Downloaded goods"
	case ErrUpdateSoldGoodsSelfDealerNotExist:
		return "Self dealer does not exist"
	case ErrUpdateSoldGoodsUnknownGoods:
		return "Unknown goods"
	case ErrUpdateSoldGoodsNoOperatingPermissions:
		return "No operating permissions"
	default:
		return fmt.Sprintf("UpdateSoldGoodsError: %d", e)
	}
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

func (e GetSoldGoodsError) String() string {
	switch int(e) {
	case ErrGetSoldGoodsSelfDealerNotExist:
		return "Self dealer does not exist"
	case ErrGetSoldGoodsSelfDealerNotBelongToMerchant:
		return "Self dealer not belong to merchant"
	default:
		return fmt.Sprintf("GetSoldGoodsError: %d", e)
	}
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

func (e OpenDoorError) String() string {
	switch int(e) {
	case ErrOpenDoorSuccess:
		return "Open door success"
	case ErrOpenDoorFailed:
		return "Open door failed"
	case ErrOpenDoorTimeout:
		return "Open door timeout"
	case ErrOpenDoorUnusualMachinePackage:
		return "Unusual machine package"
	case ErrOpenDoorOfflineEquipment:
		return "Offline equipment"
	case ErrOpenDoorSelfDealerNotInOperation:
		return "Self dealer not in operation"
	case ErrOpenDoorTooManyOrdersNotCompleted:
		return "Too many orders not completed"
	case ErrOpenDoorNonBusinessSelfSellerMachine:
		return "Non-business self seller machine"
	default:
		return fmt.Sprintf("OpenDoorError: %d", e)
	}
}
