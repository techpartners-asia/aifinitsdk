package advertisementmanage

type SourceMaterialApplyResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Count   int    `json:"count"`
	Result  []struct {
		Id int `json:"id"`
	} `json:"result"`
	Ok bool `json:"ok"`
}

type SourceMaterialPageResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Count   int    `json:"count"`
	Data    struct {
		Total int              `json:"total"`
		Rows  []SourceMaterial `json:"rows"`
	} `json:"data"`
}

type SourceMaterialDetailResponse struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    SourceMaterial `json:"data"`
}

type SourceMaterialDeleteResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

type AdAdditionResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
	Ok bool `json:"ok"`
}

type AdPageResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Total int  `json:"total"`
		Rows  []Ad `json:"rows"`
	} `json:"data"`
	Ok bool `json:"ok"`
}

type AdDetailResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Ad     `json:"data"`
}

type AdUpdateResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
	Ok bool `json:"ok"`
}

type AdDeleteResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

type AdAssociatedToVmResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type AdControlStatusResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

type GetVmPromotionResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Ad     `json:"data"`
}
