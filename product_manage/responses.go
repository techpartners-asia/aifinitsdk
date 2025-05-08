package aifinitsdk_product

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
