package advertisementmanage

type SourceMaterialStatus int

type FileType int

type BusinessType int

type AdStatus int

const (
	SourceMaterialStatusUnderReview SourceMaterialStatus = 1
	SourceMaterialStatusApproved    SourceMaterialStatus = 2
	SourceMaterialStatusRejected    SourceMaterialStatus = 3
)

const (
	FileTypeImageResource FileType = 1
	FileTypeVideoResource FileType = 2
)

const (
	BusinessTypePublicService BusinessType = 1
	BusinessTypeCommercial    BusinessType = 2
)

const (
	AdStatusEnabled  AdStatus = 1
	AdStatusDisabled AdStatus = 2
)

type SourceMaterial struct {
	Id         int                  `json:"id,omitempty"`
	FileUrl    string               `json:"fileUrl,omitempty"`
	FileType   FileType             `json:"fileType,omitempty"`
	Name       string               `json:"name,omitempty"`
	Status     SourceMaterialStatus `json:"status,omitempty"`
	CreateTime string               `json:"createTime,omitempty"`
}

type ImgRel struct {
	Id                int      `json:"id"`
	Priority          int      `json:"priority"` // 1-100 smallest to largest
	PromotionId       int      `json:"promotionId"`
	FileType          FileType `json:"fileType"`
	FileUrl           string   `json:"fileUrl"`
	SourceMaterialsId int      `json:"sourceMaterialsId"`
}

type Vm struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Ad struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	BusinessType int      `json:"businessType"`
	Duration     int      `json:"duration"`
	Status       int      `json:"status"`
	CreateTime   string   `json:"createTime"`
	UpdateTime   string   `json:"updateTime"`
	ImgRelList   []ImgRel `json:"imgRelList"`
	VmList       []Vm     `json:"vmList"`
}
