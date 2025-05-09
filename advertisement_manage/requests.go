package advertisementmanage

type SourceMaterialApplyRequest struct {
	SourceMaterialList []SourceMaterial `json:"source_material_list"`
}

type SourceMaterialPageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

type AdAdditionRequest struct {
	Name         string `json:"name,omitempty"`
	BusinessType int    `json:"businessType,omitempty"`
	Duration     int    `json:"duration,omitempty"`
	ImgRelList   []struct {
		Priority          int `json:"priority,omitempty"`
		SourceMaterialsId int `json:"sourceMaterialsId,omitempty"`
	} `json:"imgRelList,omitempty"`
}

type AdPageRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

type AdUpdateRequest struct {
	Ad struct {
		Id           int      `json:"id,omitempty" validate:"required"`
		Name         string   `json:"name,omitempty"`
		BusinessType int      `json:"businessType,omitempty"`
		Duration     int      `json:"duration,omitempty"`
		ImgRelList   []ImgRel `json:"imgRelList,omitempty"`
		VmList       []Vm     `json:"vmList,omitempty"`
	} `json:"ad,omitempty"`
}

type AdAssociatedToVmRequest struct {
	VmList []string `json:"vmList,omitempty"`
}
