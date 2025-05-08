package aifinitsdk_operation

type OpenDoorRequest struct {
	Type           OpenDoorType `json:"type,omitempty" validate:"required"`      // 1: shopping, 2: replenishment
	RequestID      string       `json:"requestId,omitempty" validate:"required"` // oruulj ogno
	UserCode       string       `json:"userCode,omitempty"`
	LocalTimeStamp int64        `json:"localTimestamp,omitempty"`
}

type UpdateSoldGoodsRequest []Good

type SearchOpenDoorRequest struct {
	Type      OpenDoorType `json:"type,omitempty"`
	RequestID string       `json:"requestId,omitempty"`
}

type ListOrderRequest struct {
	BeginTime int64 `json:"beginTime,omitempty"`
	EndTime   int64 `json:"endTime,omitempty"`
	Page      int   `json:"page,omitempty"`  //default 1
	Limit     int   `json:"limit,omitempty"` //default 10 max 50
}

type GetOrderVideoRequest struct {
	RequestID string       `json:"requestId,omitempty"`
	Type      OpenDoorType `json:"type,omitempty"`
}

type ProductPriceUpdateRequest struct {
	VmCodes []string `json:"vmCodes,omitempty"`
	Items   []Good   `json:"items,omitempty"`
}

type AddNewGoodsRequest struct {
	Items []Good `json:"items,omitempty"`
}

type DeleteGoodsRequest struct {
	ItemCodes []string `json:"itemCodes,omitempty"`
}
