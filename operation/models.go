package aifinitsdk_operation

type OpenDoorType int

const (
	OpenDoorForShopping      OpenDoorType = 1
	OpenDoorForReplenishment OpenDoorType = 2
)

type OpenDoorRequest struct {
	Type           OpenDoorType `json:"type" validate:"required"`      // 1: shopping, 2: replenishment
	RequestID      string       `json:"requestId" validate:"required"` // oruulj ogno
	UserCode       string       `json:"userCode"`
	LocalTimeStamp int64        `json:"localTimestamp"`
}

type OpenDoorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		OrderCode string `json:"orderCode"`
	} `json:"data"`
}

type Good struct {
	ItemCode      string `json:"itemCode"`
	ActualPrice   int    `json:"actualPrice"`
	OriginalPrice int    `json:"originalPrice"`
}

type GetSoldGoodsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  []Good `json:"result"`
	Count   int    `json:"count"`
}

type UpdateSoldGoodsRequest []Good

type SearchOpenDoorRequest struct {
	Type      OpenDoorType `json:"type"`
	RequestID string       `json:"requestId"`
}

type Order struct {
	TradeRequestId  string `json:"tradeRequestId,omitempty"`
	OrderCode       string `json:"orderCode,omitempty"`
	VmCode          string `json:"vmCode"`
	MachineId       int    `json:"machineId,omitempty"`
	UserCode        string `json:"userCode,omitempty"`
	HandleStatus    int    `json:"handleStatus,omitempty"`
	ShopMove        int    `json:"shopMove,omitempty"`
	TotalFee        int    `json:"totalFee,omitempty"`
	OpenDoorTime    int64  `json:"openDoorTime"`
	CloseDoorTime   int64  `json:"closeDoorTime"`
	OpenDoorWeight  int    `json:"openDoorWeight"`
	CloseDoorWeight int    `json:"closeDoorWeight"`
	OrderGoodsList  []Good `json:"orderGoodsList,omitempty"`
}

type SearchOpenDoorResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Order `json:"data"`
}

type ListOrderRequest struct {
	BeginTime int64 `json:"beginTime"`
	EndTime   int64 `json:"endTime"`
	Page      int   `json:"page"`  //default 1
	Limit     int   `json:"limit"` //default 10 max 50
}

type ListOrderResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Total int     `json:"total"`
		Rows  []Order `json:"rows"`
	} `json:"data"`
}

type GetOrderVideoRequest struct {
	RequestID string       `json:"requestId"`
	Type      OpenDoorType `json:"type"`
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
type VideoStatus int

const (
	VideoStatusPendingUpload     VideoStatus = -1
	VideoStatusUploadComplete    VideoStatus = 0
	VideoStatusVideoDoesNotExist VideoStatus = 1
	VideoStatusNetworkError      VideoStatus = 2
	VideoStatusUploadInProgress  VideoStatus = 3
)

type ProductPriceUpdateRequest struct {
	VmCodes []string `json:"vmCodes"`
	Items   []Good   `json:"items"`
}

type ProductPriceUpdateResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type AddNewGoodsRequest struct {
	Items []Good
}

type AddNewGoodsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type DeleteGoodsRequest struct {
	ItemCodes []string `json:"itemCodes"`
}

type DeleteGoodsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
