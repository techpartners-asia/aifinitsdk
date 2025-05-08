package aifinitsdk_operation

type OpenDoorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		OrderCode string `json:"orderCode"`
	} `json:"data"`
}

type GetSoldGoodsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  []Good `json:"result"`
	Count   int    `json:"count"`
}

type SearchOpenDoorResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Order `json:"data"`
}

type ListOrderResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Total int     `json:"total"`
		Rows  []Order `json:"rows"`
	} `json:"data"`
}

type GetOrderVideoResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		OrderCode   string      `json:"orderCode"`
		VideoUrl    string      `json:"videoUrl"`
		VideoURLs   []string    `json:"videoUrls"`
		VideoStatus VideoStatus `json:"videoStatus"`
	} `json:"data"`
}

type ProductPriceUpdateResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type AddNewGoodsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type DeleteGoodsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Ok      *bool  `json:"ok,omitempty"`
}
