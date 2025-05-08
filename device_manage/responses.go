package aifinitsdk_device

type DeviceInfoResult struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Device `json:"data"`
}

type ListResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Total int              `json:"total"`
		Rows  []VendingMachine `json:"rows"`
	}
}

type PeopleFlowResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Result  []PeopleFlow `json:"result"`
	Count   int          `json:"count"`
}

type DetailResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    DetailData `json:"data"`
}

type DetailData struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	ScanCode      string `json:"scanCode"`
	DeviceSn      string `json:"deviceSn"`
	ContactNumber string `json:"contactNumber"`
	Location      string `json:"location"`
	UpdateTime    string `json:"updateTime"`
}

type UpdateResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ControlResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
