package aifinitsdk_constants

const (
	BaseURL = "https://open.ainfinit.com"
)
const (
	Post_DeviceActivation           = BaseURL + "/facade/open/vending_machine/deviceActivation"
	Get_VendingMachineList          = BaseURL + "/facade/open/vending_machine/infoPage"
	Get_VendingMachineInfo          = BaseURL + "/facade/open/vending_machine/info"
	Put_UpdateVendingMachineInfo    = BaseURL + "/facade/open/vending_machine/info"
	Get_VendingMachineDeviceInfo    = BaseURL + "/facade/open/vending_machine/deviceInfo"
	Post_VendingMachinePeopleFlow   = BaseURL + "/facade/open/vending_machine/peopleFlow"
	Put_VendingMachineDeviceControl = BaseURL + "/facade/open/vending_machine/control"
	Put_VendingMachineDeviceSetting = BaseURL + "/facade/open/vending_machine/setting"
)
const (
	Get_ProductLastInfo          = BaseURL + "/facade/open/goods/latestInfo"
	Get_ProductList              = BaseURL + "/facade/open/goods/page"
	Get_ProductDetail            = BaseURL + "/facade/open/goods/%s" // %s is itemCode
	Post_ProductMutualExclusion  = BaseURL + "/facade/open/goods/getGoodsExclusionInfo"
	Post_NewProductApplication   = BaseURL + "/facade/open/goodsApply"
	Get_ProductApplicationList   = BaseURL + "/facade/open/goodsApply/page"
	Get_ProductApplicationDetail = BaseURL + "/facade/open/goodsApply/%s" // %s is itemCode
	Put_UpdateProductAppication  = BaseURL + "/facade/open/goodsApply"
)

const (
	Put_OpenDoor = BaseURL + "/open/operation/vending_machine/open"

	Put_AddNewGoods      = BaseURL + "/facade/open/replenish/items"
	Get_SoldGoods        = BaseURL + "/facade/open/replenish/items"
	Post_UpdateSoldGoods = BaseURL + "/facade/open/replenish/items"
	Del_DeleteGoods      = BaseURL + "/facade/open/replenish/items"

	Get_SearchOpenDoor      = BaseURL + "/facade/open/vending_machine"
	Get_OrderVideo          = BaseURL + "/facade/open/order/video"
	Post_ProductPriceUpdate = BaseURL + "/facade/open/replenish/replaceVmItemsPrice"
)

const (
	Post_AdvertisementMaterialApply   = BaseURL + "/facade/open/materials/sourceMaterialsApply"
	Get_AdvertisementMaterialPage     = BaseURL + "/facade/open/materials/sourceMaterialsPage"
	Get_AdvertisementMaterialDetail   = BaseURL + "/facade/open/materials/sourceMaterialsDetail/{id}"
	Del_AdvertisementMaterialDelete   = BaseURL + "/facade/open/materials/sourceMaterialsDelete/{id}"
	Post_AdvertisementAdAddition      = BaseURL + "/facade/open/materials"
	Put_AdvertisementAdUpdate         = BaseURL + "/facade/open/materials"
	Del_AdvertisementAdDelete         = BaseURL + "/facade/open/materials/{id}"
	Get_AdvertisementAdPage           = BaseURL + "/facade/open/materials/page"
	Get_AdvertisementAdDetail         = BaseURL + "/facade/open/materials/detail/{id}"
	Get_AdvertisementAdDetailByVmCode = BaseURL + "/facade/open/materials/detail"
	Put_AdvertisementAssociatedToVm   = BaseURL + "/facade/open/materials/bind/{id}"
	Get_AdvertisementAdAssociatedToVm = BaseURL + "/facade/open/materials/updatePromotionStatus/{promotionId}/{status}"
	Get_AdvertisementVmPromotion      = BaseURL + "/facade/open/materials/getVmPromotion"
)
