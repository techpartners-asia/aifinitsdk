package aifinitsdk_operation

type OperationClient interface {
	GetSoldGoods() (*GetSoldGoodsResponse, error)
	// zaragdsan baraanuudiig niitedni shinechlene
	UpdateSoldGoods(request *UpdateSoldGoodsRequest) (*UpdateSoldGoodsResponse, error)
	OpenDoor(request *OpenDoorRequest) (*OpenDoorResponse, error)
	// zaragdsan baraag haalga ongoilgoh requesteer avah
	SearchOpenDoorRequest(request *SearchOpenDoorRequest) (*SearchOpenDoorResponse, error)
	//zaragdsan baraanii jagsaalt
	ListOrder(request *ListOrderRequest) (*ListOrderResponse, error)
	// orderiin video avah
	GetOrderVideo(request *GetOrderVideoRequest) (*GetOrderVideoResponse, error)
	// product price update
	ProductPriceUpdate(request *ProductPriceUpdateRequest) (*ProductPriceUpdateResponse, error)
}
