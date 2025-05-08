package aifinitsdk_product

import (
	"encoding/json"
)

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

type ProductListResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Total int       `json:"total"`
		Rows  []Product `json:"rows"`
	}
}

type ProductListPage struct {
	Page        int    `json:"page"`
	Limit       int    `json:"limit"`
	UpdatedTime string `json:"updatedTime"`
	GoodsName   string `json:"goodsName"`
	QrCodes     string `json:"qrCodes"`
}

type LastInfoResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    LastInfo `json:"data"`
}

type LastInfo struct {
	Count          int   `json:"count"`
	LastUpdateTime int64 `json:"lastUpdateTime"`
}

type ProductDetailResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Product `json:"data"`
}

type MutualExclusionRequest struct {
	ItemCodes []string `json:"item_codes"`
}

type MutualExclusionResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Total int      `json:"total"`
		Rows  []string `json:"rows"`
	}
}

type NewProductApplicationRequest struct {
	Product *Product `json:"item"`
}

type NewProductApplicationResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    int    `json:"data"`
}

type ListProductApplicationParams struct {
	Page        int    `json:"page"`
	PageSize    int    `json:"pageSize"`
	ApplyStatus int    `json:"applyStatus"` // 1: review, 2: pass, 3: reject
	GoodsName   string `json:"goodsName"`
	QrCodes     string `json:"qrCodes"`
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

type UpdateProductApplicationRequest struct {
	Item *Product `json:"item"`
}

type UpdateProductApplicationResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
