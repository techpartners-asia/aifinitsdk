package aifinitsdk_device

type DeviceInfoResult struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Device `json:"data"`
}

type VendingMachine struct {
	DeviceSn   string `json:"deviceSn"`
	ScanCode   string `json:"scanCode"`
	Name       string `json:"name"`
	Location   string `json:"location"`
	UpdateTime string `json:"updateTime"`
}

type Device struct {
	Code                 string `json:"code"`
	BytesTotal           int    `json:"bytesTotal"`
	BytesFree            int    `json:"bytesFree"`
	ClientVersion        string `json:"clientVersion"`
	CameraCount          int    `json:"cameraCount"`
	GravityCount         int    `json:"gravityCount"`
	Light                int    `json:"light"`
	Detector             int    `json:"detector"`
	GravitySensor        int    `json:"gravitySensor"`
	SerialPort           int    `json:"serialPort"`
	SerialPortDataFormat int    `json:"serialPortDataFormat"`
	NetworkInfo          struct {
		Mobile       int `json:"mobile"`
		Wifi         int `json:"wifi"`
		Etherne      int `json:"ethernet"`
		SignalWifi   int `json:"signalWifi"`
		SignalMobile int `json:"signalMobile"`
	} `json:"networkInfo"`
	CardInfo struct {
		Iccid   string `json:"iccid"`
		Carrier string `json:"carrier"`
		Imsi    string `json:"imsi"`
	} `json:"cardInfo"`
	Ccid                  string `json:"ccid"`
	PowerStatus           int    `json:"powerStatus"`
	DeviceUpdateTimestamp int64  `json:"deviceUpdateTimestamp"`
	OnlineStatus          int    `json:"onlineStatus"`
	Temperature           int    `json:"temperature"`
	TargetTemp            int    `json:"targetTemp"`
	Volume                int    `json:"volume"`
	EngineOn              int    `json:"engineOn"`
}

type ListRequest struct {
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	NameOf string `json:"nameOf"`
}

type ListResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Total int              `json:"total"`
		Rows  []VendingMachine `json:"rows"`
	}
}

type PeopleFlowRequest struct {
	Field          *string  `json:"field,omitempty"`
	StartTimeStamp *int64   `json:"startTimestamp,omitempty"`
	EndTimeStamp   *int64   `json:"endTimestamp,omitempty"`
	Codes          []string `json:"codes,omitempty"`
}

type PeopleFlow struct {
	Code          string `json:"code"`
	VisitorCount  int    `json:"visitorCount"`
	AggregateTime string `json:"aggregateTime"`
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

type UpdateRequest struct {
	Name          *string `json:"name,omitempty"`
	Code          *string `json:"code,omitempty"`
	ScanCode      *string `json:"scanCode,omitempty"`
	ContactNumber *string `json:"contactNumber,omitempty"`
	Location      *string `json:"location,omitempty"`
	Volume        *int    `json:"volume,omitempty"`
	AdVolume      *int    `json:"adVolume,omitempty"`
	Temp          *int    `json:"temp,omitempty"`
	EngineOn      *int    `json:"engineOn,omitempty"`
}

type UpdateResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ControlRequest struct {
	Volume   *int `json:"volume,omitempty"`
	AdVolume *int `json:"adVolume,omitempty"`
	Temp     *int `json:"temp,omitempty"`
	EngineOn *int `json:"engineOn,omitempty"`
}

type ControlResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
