package aifinitsdk_device

// Device represents a vending machine device with its status and capabilities
type Device struct {
	// Code is the unique identifier for the vending machine
	Code string `json:"code"`
	// BytesTotal represents total storage space in MB
	BytesTotal int `json:"bytesTotal"`
	// BytesFree represents available storage space in MB
	BytesFree int `json:"bytesFree"`
	// ClientVersion represents the current version of the client software
	ClientVersion string `json:"clientVersion"`
	// CameraCount represents the number of available cameras
	CameraCount int `json:"cameraCount"`
	// GravityCount represents the number of available gravity sensors
	GravityCount int `json:"gravityCount"`
	// Light represents the light status (0: normal, 1: abnormal)
	Light int `json:"light"`
	// Detector represents the local detection status (0: normal, 1: abnormal)
	Detector int `json:"detector"`
	// GravitySensor represents the gravity sensor status (0: normal, 1: abnormal)
	GravitySensor int `json:"gravitySensor"`
	// SerialPort represents the serial port connection status (0: normal, 1: abnormal)
	SerialPort int `json:"serialPort"`
	// SerialPortDataFormat represents the serial port data format status (0: normal, 1: abnormal)
	SerialPortDataFormat int `json:"serialPortDataFormat"`
	// NetworkInfo contains information about network connections
	NetworkInfo struct {
		// Mobile represents 4G connection status (1: connected, 0: disconnected)
		Mobile int `json:"mobile"`
		// Wifi represents WiFi connection status (1: connected, 0: disconnected)
		Wifi int `json:"wifi"`
		// Etherne represents ethernet connection status (1: connected, 0: disconnected)
		Etherne int `json:"ethernet"`
		// SignalWifi represents WiFi signal strength
		SignalWifi int `json:"signalWifi"`
		// SignalMobile represents 4G signal strength
		SignalMobile int `json:"signalMobile"`
	} `json:"networkInfo"`
	// CardInfo contains information about the SIM card
	CardInfo struct {
		// Iccid is the Integrated Circuit Card Identifier
		Iccid string `json:"iccid"`
		// Carrier represents the mobile carrier (CMCC: China Mobile, CUCC: China Unicom, CTCC: China Telecom)
		Carrier string `json:"carrier"`
		// Imsi is the International Mobile Subscriber Identity
		Imsi string `json:"imsi"`
	} `json:"cardInfo"`
	// Ccid is the Chip Card ID (empty string if not reported)
	Ccid string `json:"ccid"`
	// PowerStatus represents the power source (1: mains power, 2: UPS)
	PowerStatus int `json:"powerStatus"`
	// DeviceUpdateTimestamp represents the last heartbeat timestamp
	DeviceUpdateTimestamp int64 `json:"deviceUpdateTimestamp"`
	// OnlineStatus represents the network status (1: online, 0: offline, calculated based on heartbeat timestamp)
	OnlineStatus int `json:"onlineStatus"`
	// Temperature represents the current temperature
	Temperature int `json:"temperature"`
	// TargetTemp represents the target temperature setting
	TargetTemp int `json:"targetTemp"`
	// Volume represents the current volume level
	Volume int `json:"volume"`
	// EngineOn represents the compressor status (1: on, 0: off)
	EngineOn int `json:"engineOn"`
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
