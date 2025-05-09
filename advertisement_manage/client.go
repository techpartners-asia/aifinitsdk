package advertisementmanage

type AdvertisementManageClient interface {
	MaterialApply(request *SourceMaterialApplyRequest) (*SourceMaterialApplyResponse, error)
	MaterialPage(request *SourceMaterialPageRequest) (*SourceMaterialPageResponse, error)
	MaterialDetail(materialId string) (*SourceMaterialDetailResponse, error)
	MaterialDelete(materialId string) (*SourceMaterialDeleteResponse, error)
	AdAddition(request *AdAdditionRequest) (*AdAdditionResponse, error)
	AdPage(request *AdPageRequest) (*AdPageResponse, error)
	AdDetailByAdId(adId int) (*AdDetailResponse, error)
	AdDetailByVmCode(code string) (*AdDetailResponse, error)
	AdUpdate(request *AdUpdateRequest) (*AdUpdateResponse, error)
	AdDelete(adId int) (*AdDeleteResponse, error)
	AdAssociatedToVm(adId int, request *AdAssociatedToVmRequest) (*AdAssociatedToVmResponse, error)
	ControlAdStatus(promotionId int, status AdStatus) (*AdControlStatusResponse, error)
	GetVmPromotion(vmCode string) (*GetVmPromotionResponse, error)
	//callback
	MediaReviewNotify()
}
