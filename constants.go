package aifinitsdk

const (
	DefaultBaseURL = "https://ainfinit.mtm.mn"
)

// Device Manage
const (
	Put_DeviceCoolingCommand        = "/facade/open/vending_machine/deviceCoolingCommand"
	Put_DeviceSetting               = "/facade/open/vending_machine/setting"
	Post_DeviceActivation           = "/facade/open/vending_machine/bind"
	Get_VendingMachineList          = "/facade/open/vending_machine/infoPage"
	Get_VendingMachineInfo          = "/facade/open/vending_machine/info"
	Put_UpdateVendingMachineInfo    = "/facade/open/vending_machine/info"
	Get_VendingMachineDeviceDetail  = "/facade/open/vending_machine/deviceInfo"
	Post_VendingMachinePeopleFlow   = "/facade/open/vending_machine/peopleFlow"
	Put_VendingMachineDeviceControl = "/facade/open/vending_machine/control"
	Put_VendingMachineDeviceSetting = "/facade/open/vending_machine/setting"
)

// Product Manage
const (
	Get_ProductLastInfo          = "/facade/open/goods/latestInfo"
	Get_ProductList              = "/facade/open/goods/page"
	Get_ProductDetail            = "/facade/open/goods/%s" // %s is itemCode
	Post_ProductMutualExclusion  = "/facade/open/goods/getGoodsExclusionInfo"
	Post_NewProductApplication   = "/facade/open/goodsApply"
	Get_ProductApplicationList   = "/facade/open/goodsApply/page"
	Get_ProductApplicationDetail = "/facade/open/goodsApply/%s" // %s is itemCode
	Put_UpdateProductAppication  = "/facade/open/goodsApply"
)

const (
	Put_OpenDoor         = "/open/operation/vending_machine/open"
	Get_ListOrders       = "/facade/open/order/page"
	Put_AddNewGoods      = "/facade/open/replenish/items"
	Get_SoldGoods        = "/facade/open/replenish/items"
	Post_UpdateSoldGoods = "/facade/open/replenish/items"
	Del_DeleteGoods      = "/facade/open/replenish/items"

	Get_SearchOpenDoor      = "/facade/open/vending_machine"
	Get_OrderVideo          = "/facade/open/order/video"
	Post_ProductPriceUpdate = "/facade/open/replenish/replaceVmItemsPrice"
)

const (
	Post_AdvertisementMaterialApply   = "/facade/open/materials/sourceMaterialsApply"
	Get_AdvertisementMaterialPage     = "/facade/open/materials/sourceMaterialsPage"
	Get_AdvertisementMaterialDetail   = "/facade/open/materials/sourceMaterialsDetail/{id}"
	Del_AdvertisementMaterialDelete   = "/facade/open/materials/sourceMaterialsDelete/{id}"
	Post_AdvertisementAdAddition      = "/facade/open/materials"
	Put_AdvertisementAdUpdate         = "/facade/open/materials"
	Del_AdvertisementAdDelete         = "/facade/open/materials/{id}"
	Get_AdvertisementAdPage           = "/facade/open/materials/page"
	Get_AdvertisementAdDetail         = "/facade/open/materials/detail/{id}"
	Get_AdvertisementAdDetailByVmCode = "/facade/open/materials/detail"
	Put_AdvertisementAssociatedToVm   = "/facade/open/materials/bind/{id}"
	Get_AdvertisementAdAssociatedToVm = "/facade/open/materials/updatePromotionStatus/{promotionId}/{status}"
	Get_AdvertisementVmPromotion      = "/facade/open/materials/getVmPromotion"
)
