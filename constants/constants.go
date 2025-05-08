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
	Put_OpenDoor            = BaseURL + "/open/operation/vending_machine/open"
	Get_SoldGoods           = BaseURL + "/facade/open/replenish/items"
	Put_UpdateSoldGoods     = BaseURL + "/facade/open/replenish/items"
	Get_SearchOpenDoor      = BaseURL + "/open/operation/vending_machine/open"
	Get_OrderVideo          = BaseURL + "/facade/open/order/video"
	Post_ProductPriceUpdate = BaseURL + "/facade/open/replenish/replaceVmItemsPrice"
)
