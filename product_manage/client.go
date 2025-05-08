package aifinitsdk_product

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
