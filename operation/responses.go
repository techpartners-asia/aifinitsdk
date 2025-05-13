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

type SearchOpenDoorData struct {
	TradeRequestId  string `json:"tradeRequestId"`
	OrderCode       string `json:"orderCode"`
	VmCode          string `json:"vmCode"`
	MachineId       int    `json:"machineId"`
	HandleStatus    int    `json:"handleStatus"`
	TotalFee        int    `json:"totalFee"`
	OpenDoorTime    int64  `json:"openDoorTime"`
	CloseDoorTime   int64  `json:"closeDoorTime"`
	OpenDoorWeight  int    `json:"openDoorWeight"`
	CloseDoorWeight int    `json:"closeDoorWeight"`
	OrderGoodsList  []Good `json:"orderGoodsList"`
	ShopMove        int    `json:"shopMove"`
	ScanCode        string `json:"scanCode"`
}

type SearchOpenDoorResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    SearchOpenDoorData `json:"data"`
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
