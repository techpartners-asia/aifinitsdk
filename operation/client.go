package aifinitsdk_operation

type OperationClient interface {
	//2.2.3.11
	AddNewGoods(request *AddNewGoodsRequest) (*AddNewGoodsResponse, error)
	//2.2.3.12
	DeleteGoods(request *DeleteGoodsRequest) (*DeleteGoodsResponse, error)

	//2.2.3.1
	GetSoldGoods() (*GetSoldGoodsResponse, error)
	// zaragdsan baraanuudiig niitedni shinechlene
	//2.2.3.2
	UpdateSoldGoods(request *UpdateSoldGoodsRequest) (*UpdateSoldGoodsResponse, error)
	//2.2.3.3
	OpenDoor(request *OpenDoorRequest) (*OpenDoorResponse, error)
	// zaragdsan baraag haalga ongoilgoh requesteer avah
	//2.2.3.4
	SearchOpenDoorRequest(request *SearchOpenDoorRequest) (*SearchOpenDoorResponse, error)
	//zaragdsan baraanii jagsaalt
	//2.2.3.6
	ListOrder(request *ListOrderRequest) (*ListOrderResponse, error)
	// orderiin video avah
	//2.2.3.8
	GetOrderVideo(request *GetOrderVideoRequest) (*GetOrderVideoResponse, error)
	// product price update
	//2.2.3.10
	ProductPriceUpdate(request *ProductPriceUpdateRequest) (*ProductPriceUpdateResponse, error)
}
