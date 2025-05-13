package aifinitsdk_operation

type OperationClient interface {
	//mashind baraa nemne
	//2.2.3.11
	AddGoods(request *AddNewGoodsRequest, machineCode string) (*AddNewGoodsResponse, error)
	//2.2.3.12
	DeleteGoods(request *DeleteGoodsRequest, machineCode string) (*DeleteGoodsResponse, error)

	//machine dotorh baraanuudiin jagsaalt
	//2.2.3.1
	ListGoods(machineCode string) (*GetSoldGoodsResponse, error)
	// zaragdsan baraanuudiig niitedni shinechlene
	//2.2.3.2
	UpdateGoods(request *UpdateGoodsRequest, machineCode string) (*UpdateSoldGoodsResponse, error)
	//2.2.3.3
	OpenDoor(request *OpenDoorRequest, machineCode string) (*OpenDoorResponse, error)
	// zaragdsan baraag haalga ongoilgoh requesteer avah
	//2.2.3.4
	GetSoldGoodsByRequestID(request *GetSoldGoodsByRequestIDRequest, machineCode string) (*SearchOpenDoorResponse, error)
	//zaragdsan baraanii jagsaalt
	//2.2.3.6
	ListOrders(request *ListOrderRequest, machineCode string) (*ListOrderResponse, error)
	// orderiin video avah
	//2.2.3.8
	GetOrderVideo(request *GetOrderVideoRequest, machineCode string) (*GetOrderVideoResponse, error)
	// product price update
	//2.2.3.10
	UpdateGoodsPrice(request *ProductPriceUpdateRequest, machineCode string) (*ProductPriceUpdateResponse, error)
}
