package aifinitsdk_product

import (
	"fmt"
	"strings"
	"time"

	aifinitsdk "github.com/techpartners-asia/aifinitsdk"
	aifinitsdk_constants "github.com/techpartners-asia/aifinitsdk/constants"
	"resty.dev/v3"
)

type ProductClient struct {
	Client aifinitsdk.Client
	Resty  *resty.Client
}

func NewProductClient(client aifinitsdk.Client) *ProductClient {
	return &ProductClient{
		Client: client,
		Resty:  resty.New(),
	}
}

func (c *ProductClient) GetProductList(page, limit int) (*ProductListResponse, error) {
	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var products *ProductListResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetResult(&products).Get(aifinitsdk_constants.Get_ProductList)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
	}

	return products, nil
}
func (c *ProductClient) GetProductDetail(itemCode string) (*ProductDetailResponse, error) {
	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var product *ProductDetailResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetResult(&product).Get(fmt.Sprintf(aifinitsdk_constants.Get_ProductDetail, itemCode))
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
	}

	return product, nil
}

func (c *ProductClient) GetProductMutualExclusion(request *MutualExclusionRequest) (*MutualExclusionResponse, error) {
	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var mutualExclusion *MutualExclusionResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetBody(request).SetResult(&mutualExclusion).Post(aifinitsdk_constants.Post_ProductMutualExclusion)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
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
		req = req.SetFileReader("file", request.Product.ImgFileNames[i], img)
	}

	if len(request.Product.PhysicalImgFiles) != len(request.Product.PhysicalImgFileNames) {
		return nil, fmt.Errorf("actual image files and names must be the same length")
	}

	for i, img := range request.Product.PhysicalImgFiles {
		if img != nil {
			req = req.SetFileReader("files", request.Product.PhysicalImgFileNames[i], img)
		}
	}

	if len(request.Product.WeightFiles) != len(request.Product.WeightFileNames) {
		return nil, fmt.Errorf("weight files and names must be the same length")
	}

	for i, weight := range request.Product.WeightFiles {
		req = req.SetFileReader("weightFile", request.Product.WeightFileNames[i], weight)
	}

	resp, err := req.SetResult(&newProductApplication).Post(aifinitsdk_constants.Post_NewProductApplication)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
	}

	return newProductApplication, nil
}

func (c *ProductClient) ListProductApplication(params *ListProductApplicationParams) (*ListProductApplicationResponse, error) {
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
	}).SetResult(&listProductApplication).Get(aifinitsdk_constants.Get_ProductApplicationList)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
	}

	return listProductApplication, nil
}

func (c *ProductClient) DetailProductApplication(itemCode string) (*DetailProductApplicationResponse, error) {
	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var detailProductApplication *DetailProductApplicationResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetResult(&detailProductApplication).Get(fmt.Sprintf(aifinitsdk_constants.Get_ProductApplicationDetail, itemCode))
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
	}

	return detailProductApplication, nil
}

func (c *ProductClient) UpdateProductApplication(request *UpdateProductApplicationRequest) (*UpdateProductApplicationResponse, error) {
	signature, err := c.Client.GetSignature(time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}

	var updateProductApplication *UpdateProductApplicationResponse
	resp, err := c.Resty.R().SetHeader("Authorization", signature).SetBody(request).SetResult(&updateProductApplication).Put(aifinitsdk_constants.Put_UpdateProductAppication)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("status: %d, message: %s", resp.StatusCode(), resp.String())
	}

	return updateProductApplication, nil
}
