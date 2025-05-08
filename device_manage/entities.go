package aifinitsdk_device

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

type VendingMachine struct {
	DeviceSn   string `json:"deviceSn"`
	ScanCode   string `json:"scanCode"`
	Name       string `json:"name"`
	Location   string `json:"location"`
	UpdateTime string `json:"updateTime"`
}

type PeopleFlow struct {
	Code          string `json:"code"`
	VisitorCount  int    `json:"visitorCount"`
	AggregateTime string `json:"aggregateTime"`
}
