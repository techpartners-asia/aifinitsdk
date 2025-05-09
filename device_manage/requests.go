package aifinitsdk_device

type ListRequest struct {
	Page   int    `json:"page,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	NameOf string `json:"nameOf,omitempty"`
}

type PeopleFlowRequest struct {
	Field          string   `json:"field,omitempty"`
	StartTimeStamp int64    `json:"startTimestamp,omitempty"`
	EndTimeStamp   int64    `json:"endTimestamp,omitempty"`
	Codes          []string `json:"codes,omitempty"`
}

type UpdateRequest struct {
	Name          string `json:"name,omitempty" validate:"required"`
	Code          string `json:"code,omitempty"`
	ScanCode      string `json:"scanCode,omitempty"`
	ContactNumber string `json:"contactNumber,omitempty"`
	Location      string `json:"location,omitempty"`
	Volume        int    `json:"volume,omitempty"`
	AdVolume      int    `json:"adVolume,omitempty"`
	Temp          int    `json:"temp,omitempty"`
	EngineOn      int    `json:"engineOn,omitempty"`
}

type ControlRequest struct {
	Volume   int `json:"volume,omitempty"`   // 0 ~ 100
	AdVolume int `json:"adVolume,omitempty"` // 0 ~ 100
	Temp     int `json:"temp,omitempty"`     // -30 ~ 20
	EngineOn int `json:"engineOn,omitempty"` // 0 | 1
}
