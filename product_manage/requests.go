package aifinitsdk_product

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
