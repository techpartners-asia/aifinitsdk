package aifinitsdk_operation

type Good struct {
	ItemCode      string `json:"itemCode"`
	ActualPrice   int    `json:"actualPrice"`
	OriginalPrice int    `json:"originalPrice"`
}

type Order struct {
	TradeRequestId  string `json:"tradeRequestId,omitempty"`
	OrderCode       string `json:"orderCode,omitempty"`
	VmCode          string `json:"vmCode,omitempty"`
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

type OpenDoorType int

const (
	OpenDoorForShopping      OpenDoorType = 1
	OpenDoorForReplenishment OpenDoorType = 2
)

type VideoStatus int

const (
	VideoStatusPendingUpload     VideoStatus = -1
	VideoStatusUploadComplete    VideoStatus = 0
	VideoStatusVideoDoesNotExist VideoStatus = 1
	VideoStatusNetworkError      VideoStatus = 2
	VideoStatusUploadInProgress  VideoStatus = 3
)
