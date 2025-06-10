package aifinitsdk

// ------------------- 2.2.2.8 ------------------- //
// Alarm Notification
//
// Used to notify the third-party system of vending machine maintenance or operational exceptions.
// The notification type is determined by the 'action' parameter in the query:
// - client_warning: Maintenance Exception
// - operating_exception: Operational Exception
//
// Note: For maintenance exceptions (client_warning):
// - When temperature drops below 3°C, a low temperature alert will be triggered
// - Each vending machine can only trigger one low temperature alert per day
//
// Note: For operational exceptions (operating_exception):
// - Weight anomaly (exType=1) triggers two notifications: one at occurrence and one when video is uploaded
// - Both notifications share the same exId
// - Video availability varies by exType:
//   * exType=1: Always includes video
//   * exType=2,3: No video
//   * exType=4,5,6: May reference shopping/restocking videos

// Alarm Action Type
type AlarmAction string

const (
	AlarmActionClientWarning      AlarmAction = "client_warning"      // Maintenance Exception
	AlarmActionOperatingException AlarmAction = "operating_exception" // Operational Exception
)

// Maintenance Exception Status
type MaintenanceExceptionStatus int

const (
	MaintenanceExceptionStatusTriggered MaintenanceExceptionStatus = 0 // Exception triggered
	MaintenanceExceptionStatusRecovered MaintenanceExceptionStatus = 1 // Exception recovered
)

// Maintenance Exception Code
type MaintenanceExceptionCode int

const (
	MaintenanceExceptionCodeCameraIssue      MaintenanceExceptionCode = 1   // Camera Issue
	MaintenanceExceptionCodeHeavySensor      MaintenanceExceptionCode = 128 // Heavy Sensor
	MaintenanceExceptionCodeUPSPower         MaintenanceExceptionCode = 2   // UPS Power
	MaintenanceExceptionCodeOverheating      MaintenanceExceptionCode = 0   // Overheating
	MaintenanceExceptionCodeShelfMalfunction MaintenanceExceptionCode = 0   // Shelf Malfunction
	MaintenanceExceptionCodeLightCurtain     MaintenanceExceptionCode = 0   // Light Curtain Error
	MaintenanceExceptionCodePositionShift    MaintenanceExceptionCode = 0   // Position Shift
	MaintenanceExceptionCodeCardReader       MaintenanceExceptionCode = 0   // Card Reader Issue
	MaintenanceExceptionCodePowerOff         MaintenanceExceptionCode = 0   // Power Off
	MaintenanceExceptionCodeTooCold          MaintenanceExceptionCode = 0   // Too Cold (< 3°C)
	MaintenanceExceptionCodeLockState        MaintenanceExceptionCode = 0   // Lock State Issue
	MaintenanceExceptionCodeLockModules      MaintenanceExceptionCode = 0   // Lock Modules
	MaintenanceExceptionCodeNetwork          MaintenanceExceptionCode = 0   // Network
	MaintenanceExceptionCodeLight            MaintenanceExceptionCode = 0   // Light
	MaintenanceExceptionCodeSerialConnection MaintenanceExceptionCode = 0   // Serial Connection
	MaintenanceExceptionCodeDiskSpace        MaintenanceExceptionCode = 0   // Disk Space
)

// Operational Exception Type
type OperationalExceptionType int

const (
	OperationalExceptionTypeWeightAnomaly       OperationalExceptionType = 1 // Abnormal weight change (non-shopping)
	OperationalExceptionTypeDoorLockAnomaly     OperationalExceptionType = 2 // Door lock anomaly (non-shopping)
	OperationalExceptionTypeUPSPower            OperationalExceptionType = 3 // Switched to UPS power
	OperationalExceptionTypeShoppingLockTimeout OperationalExceptionType = 4 // Lock timeout after shopping
	OperationalExceptionTypeRestockLockTimeout  OperationalExceptionType = 5 // Lock timeout after restocking
	OperationalExceptionTypeShoppingTimeout     OperationalExceptionType = 6 // Shopping session timeout
	OperationalExceptionTypeForeignIntrusion    OperationalExceptionType = 7 // Foreign object intrusion
	OperationalExceptionTypeInventoryMismatch   OperationalExceptionType = 8 // Inventory mismatch
	OperationalExceptionTypeUnauthorizedDoor    OperationalExceptionType = 9 // Door opened without shopping
)

// Alarm Video Status
type AlarmVideoStatus int

const (
	AlarmVideoStatusNotUploaded  AlarmVideoStatus = -1 // Not uploaded
	AlarmVideoStatusSuccess      AlarmVideoStatus = 0  // Success
	AlarmVideoStatusNotFound     AlarmVideoStatus = 1  // Not found
	AlarmVideoStatusUploadFailed AlarmVideoStatus = 2  // Upload failed
)

// Maintenance Exception Notification
type MaintenanceExceptionNotificationCallbackRequest struct {
	ExCode     MaintenanceExceptionCode   `json:"exCode"`             // Exception code
	NotifyTime int64                      `json:"notifyTime"`         // Time of exception occurrence or recovery
	Status     MaintenanceExceptionStatus `json:"status"`             // 0: Triggered, 1: Recovered
	VmCode     string                     `json:"vmCode"`             // Vending machine code
	VmName     string                     `json:"vmName"`             // Vending machine name
	ScanCode   string                     `json:"scanCode,omitempty"` // QR sticker code
}

// Operational Exception Notification
type OperationalExceptionNotificationCallbackRequest struct {
	VmName        string                   `json:"vmName"`                  // Device name
	VmCode        string                   `json:"vmCode"`                  // Device code
	RequestID     string                   `json:"requestId,omitempty"`     // Door-open request ID (only for shopping-related alerts)
	ExID          string                   `json:"exId"`                    // Exception alarm ID
	ExType        OperationalExceptionType `json:"exType"`                  // Exception type
	ExDetail      string                   `json:"exDetail"`                // Detailed info (e.g., lockopen_doorclose)
	SendTime      int64                    `json:"sendTime"`                // Client-side exception timestamp (in ms)
	VideoURL      string                   `json:"videoUrl,omitempty"`      // Video URL (if applicable)
	VideoStatus   AlarmVideoStatus         `json:"videoStatus,omitempty"`   // Video status
	VideoSendTime int64                    `json:"videoSendTime,omitempty"` // Video upload timestamp (in ms)
	ScanCode      string                   `json:"scanCode,omitempty"`      // QR sticker code
}

// Product Application Review Status
type ProductApplicationReviewStatus int

const (
	ProductApplicationReviewStatusApproved ProductApplicationReviewStatus = 2 // Approved
	ProductApplicationReviewStatusRejected ProductApplicationReviewStatus = 3 // Rejected
)

// Product Application Reject Type
type ProductApplicationRejectType string

const (
	ProductApplicationRejectTypeNameNonCompliant    ProductApplicationRejectType = "1" // Name non-compliant
	ProductApplicationRejectTypeBarcodeNonCompliant ProductApplicationRejectType = "2" // Barcode non-compliant
	ProductApplicationRejectTypeImageUnclear        ProductApplicationRejectType = "3" // Image unclear
	ProductApplicationRejectTypeOther               ProductApplicationRejectType = "4" // Other
)

// ------------------- 2.2.1.10 ------------------- //
// Product Application Review Notification
//
// The system sends a notification to the third-party system after reviewing a new product application.
// This notification contains the review results including approval status and any rejection reasons.
//
// Note: The notification includes different fields based on the review status:
// - For approved applications: itemCode is provided
// - For rejected applications: rejectType and rejectReason are required

// Product Application Review Notification
type ProductApplicationReviewNotificationCallbackRequest struct {
	ID           int64                          `json:"id"`                     // New product application ID
	Status       ProductApplicationReviewStatus `json:"status"`                 // 2: Approved, 3: Rejected
	ItemCode     string                         `json:"itemCode,omitempty"`     // Product code (present only if approved)
	RejectType   ProductApplicationRejectType   `json:"rejectType,omitempty"`   // Rejection type
	RejectReason string                         `json:"rejectReason,omitempty"` // Reason for rejection
}

// ------------------- 2.2.4.13 ------------------- //
// Advertisement Online/Offline Notification
//
// The system sends a notification to the third-party system when an advertisement's status changes
// between online and offline. This notification contains the advertisement's ID, name, and new status.
//
// Note: The status field indicates the current state of the advertisement:
// - 1: Online - Advertisement is now active and visible
// - 2: Offline - Advertisement is now inactive and hidden

// Advertisement Online/Offline Status
type AdvertisementOnlineStatus int

const (
	AdvertisementOnlineStatusOnline  AdvertisementOnlineStatus = 1 // Online
	AdvertisementOnlineStatusOffline AdvertisementOnlineStatus = 2 // Offline
)

// Advertisement Online/Offline Notification
type AdvertisementOnlineNotificationCallbackRequest struct {
	ID     int                       `json:"id"`     // Advertisement ID
	Name   string                    `json:"name"`   // Advertisement name
	Status AdvertisementOnlineStatus `json:"status"` // 1: Online, 2: Offline
}

// ------------------- 2.2.4.5 ------------------- //
// Material Review Notification
//
// After advertising materials are submitted, JiandanGou (Simple Buy) operations personnel will review them.
// The review result will be sent as a notification.
//
// Note: The status field indicates the review result:
// - 2: Approved - Material has been approved
// - 3: Rejected - Material has been rejected (rejectReason will be provided)

// Material Review Status
type MaterialReviewStatus int

const (
	MaterialReviewStatusApproved MaterialReviewStatus = 2 // Approved
	MaterialReviewStatusRejected MaterialReviewStatus = 3 // Rejected
)

// Material Review Source
type MaterialReviewSource struct {
	ID int `json:"id"` // Material ID
}

// Material Review Notification
type MaterialReviewNotificationCallbackRequest struct {
	SourceMaterialsList []MaterialReviewSource `json:"sourceMaterialsList"`    // List of advertising materials
	Status              MaterialReviewStatus   `json:"status"`                 // Review result: 2 = approved, 3 = rejected
	RejectReason        string                 `json:"rejectReason,omitempty"` // Reason for rejection
}

// ------------------- 2.2.1.5 ------------------- //
// Product Change Notification
//
// Notifies third-party systems of product changes including new products,
// product updates, and product deletions.
//
// Note: The action parameter in the query determines the type of change:
// - add: New product added
// - update: Existing product updated
// - delete: Product deleted

// Product Change Action Type
type ProductChangeAction string

const (
	ProductChangeActionAdd    ProductChangeAction = "add"    // New product
	ProductChangeActionUpdate ProductChangeAction = "update" // Update product
	ProductChangeActionDelete ProductChangeAction = "delete" // Delete product
)

// Product Collection Type
type ProductCollectionType int

const (
	ProductCollectionTypeSingle ProductCollectionType = 1 // Single item
	ProductCollectionTypeBundle ProductCollectionType = 2 // Bundle
)

// Product Status
type ProductStatus int

const (
	ProductStatusListed   ProductStatus = 1 // Listed (available)
	ProductStatusUnlisted ProductStatus = 2 // Unlisted (unavailable)
)

// Product Change Notification
type ProductChangeNotificationCallbackRequest struct {
	Code      string                `json:"code"`      // Product code
	CollType  ProductCollectionType `json:"collType"`  // Product collection type: 1 = single item, 2 = bundle
	ImageUrl  string                `json:"imageUrl"`  // Product image URL
	ItemCodes []string              `json:"itemCodes"` // Codes of items in the bundle (empty for single items)
	Name      string                `json:"name"`      // Product name
	Price     int                   `json:"price"`     // Price in cents
	Status    ProductStatus         `json:"status"`    // 1 = listed (available), 2 = unlisted (unavailable)
	Weight    int                   `json:"weight"`    // Weight in grams
}

// ------------------- 2.2.3.7 ------------------- //
// Order Settlement Notification
//
// The system sends a notification to the third-party system after the user completes their shopping
// and closes the door. This notification contains the final shopping results.
//
// Note: Some orders may require additional processing:
// - Cloud-based recognition for complex cases
// - Manual handling for edge cases
// These special cases are tracked using the handleStatus field in the response.

type HandleStatus int

const (
	HandleStatusLocalSuccess HandleStatus = 1 // Local recognition is normal
	HandleStatusLocalFailure HandleStatus = 2 // Local identification failure
	HandleStatusCloudSuccess HandleStatus = 3 // Cloud Identification Completed
	HandleStatusCloudFailure HandleStatus = 4 // Cloud recognition failure
)

// IsFailure checks if the handle status is a failure
func (h HandleStatus) IsFailure() bool {
	return h == HandleStatusLocalFailure || h == HandleStatusCloudFailure
}

// IsSuccess checks if the handle status is a success
func (h HandleStatus) IsSuccess() bool {
	return h == HandleStatusLocalSuccess || h == HandleStatusCloudSuccess
}

func (h HandleStatus) String() string {
	statuses := map[HandleStatus]string{
		HandleStatusLocalSuccess: "Local Success",
		HandleStatusLocalFailure: "Local Failure",
		HandleStatusCloudSuccess: "Cloud Success",
		HandleStatusCloudFailure: "Cloud Failure",
	}
	if status, ok := statuses[h]; ok {
		return status
	}
	return "Unknown"
}

type AbnormalReason string

const (
	AbnormalReasonCameraEx        AbnormalReason = "CAMERA_EX"        // Camera anomaly
	AbnormalReasonGravityEx       AbnormalReason = "GRAVITY_EX"       // I don't feel important
	AbnormalReasonForeignInvasion AbnormalReason = "FOREIGN_INVASION" // Invasion of Foreign Objects
	AbnormalReasonUnknownItem     AbnormalReason = "UNKNOWN_ITEM"     // Unknown Products
	AbnormalReasonOther           AbnormalReason = "OTHER"            // Other
	AbnormalReasonUnfriendly      AbnormalReason = "UNFRIENDLY"       // Non-friendly operation
	AbnormalReasonVideoError      AbnormalReason = "VIDEO_ERROR"      // Video anomalies
	AbnormalReasonHardwareEx      AbnormalReason = "HARDWARE_EX"      // Failed algorithmic recognition
)

func (h AbnormalReason) String() string {
	reasons := map[AbnormalReason]string{
		AbnormalReasonCameraEx:        "Camera anomaly",
		AbnormalReasonGravityEx:       "I don't feel important",
		AbnormalReasonForeignInvasion: "Invasion of Foreign Objects",
		AbnormalReasonUnknownItem:     "Unknown Products",
		AbnormalReasonOther:           "Other",
		AbnormalReasonUnfriendly:      "Non-friendly operation",
		AbnormalReasonVideoError:      "Video anomalies",
		AbnormalReasonHardwareEx:      "Failed algorithmic recognition",
	}
	if reason, ok := reasons[h]; ok {
		return reason
	}
	return "Unknown"
}

type HardwareException string

const (
	HardwareExceptionCamera          HardwareException = "Camera"            // Camera is unusual
	HardwareExceptionGravity         HardwareException = "GRAVITY"           // Awesome feeling
	HardwareExceptionForeignInvasion HardwareException = "FOREIGN_INVASION"  // Invasion of Foreign Objects
	HardwareExceptionNetwork         HardwareException = "Network anomalies" // Network anomalies
	HardwareExceptionCrash           HardwareException = "CRASH"             // Unusual software exit
)

type ShopMove int

const (
	ShopMoveDoorNotOpen      ShopMove = 1 // The door is not open
	ShopMoveDoorOpenNoMove   ShopMove = 2 // The door is open, but no movement
	ShopMoveDoorOpenWithMove ShopMove = 3 // The door is open and there is movement
)

type OrderCallbackRequest struct {
	TradeRequestId  string            `json:"tradeRequestId"`            // Open the door to id
	OrderCode       string            `json:"orderCode"`                 // Order Code
	UserCode        string            `json:"userCode,omitempty"`        // User code
	VmCode          string            `json:"vmCode"`                    // Self-dealer equipment code
	HandleStatus    HandleStatus      `json:"handleStatus"`              // Identification of processing status
	AbnormalReasons []AbnormalReason  `json:"abnormalReasons,omitempty"` // Artificial handling of abnormal causes
	OpenDoorTime    int64             `json:"openDoorTime"`              // Opening the door time stamp
	OpenDoorWeight  int               `json:"openDoorWeight"`            // Total weight of the door
	CloseDoorTime   int64             `json:"closeDoorTime"`             // Closed time stamp
	CloseDoorWeight int               `json:"closeDoorWeight"`           // Total closing weight
	HardwareEx      HardwareException `json:"hardwareEx,omitempty"`      // Hardware exception
	ShopMove        ShopMove          `json:"shopMove"`                  // Door movement status
	VideoUrl        string            `json:"videoUrl,omitempty"`        // Shopping video address
	VideoUrls       []string          `json:"videoUrls,omitempty"`       // Shopping video address collection
	OrderGoodsList  []OrderGoods      `json:"orderGoodsList,omitempty"`  // List of products
	Candidates      []OrderGoods      `json:"candidates,omitempty"`      // Candidate product set
}

// END Order Callback Request |^

// ------------------- 2.2.3.5 ------------------- //
// Asynchronous Door Open/Close Notification
//
// Used to notify the third-party system of the door open/close result.
// The notification type is determined by the 'action' parameter:
// - trade_open: Shopping door open
// - trade_close: Shopping door close
// - replenish_open: Restocking door open
// - replenish_close: Restocking door close
//
// Note: The orderCode field is only present for shopping-related notifications
// (trade_open and trade_close actions).

// Door Open/Close Action Type
type DoorOpenCloseAction string

const (
	DoorOpenCloseActionTradeOpen      DoorOpenCloseAction = "trade_open"      // Shopping door open
	DoorOpenCloseActionTradeClose     DoorOpenCloseAction = "trade_close"     // Shopping door close
	DoorOpenCloseActionReplenishOpen  DoorOpenCloseAction = "replenish_open"  // Restocking door open
	DoorOpenCloseActionReplenishClose DoorOpenCloseAction = "replenish_close" // Restocking door close
)

// Door Open/Close Notification
type DoorOpenCloseNotificationCallbackRequest struct {
	OrderCode string `json:"orderCode,omitempty"` // Order number (only for shopping open/close)
	OpenType  int    `json:"openType"`            // 1: Shopping, 2: Restocking
	RequestID string `json:"requestId"`           // Door open request ID
	Status    int    `json:"status"`              // Door open/close result (see section 2.2.3.4)
	VmCode    string `json:"vmCode"`              // Vending machine code
}

// ------------------- 2.2.3.4 ------------------- //
// Query Door Open/Close Result
//
// This API is used to query the door open/close result based on a request ID.
// The response includes different fields based on the door operation scenario:
// - Shopping (type=1): Includes order details, user info, and goods list
// - Restocking (type=2): Includes only basic door operation info
//
// Note: The status codes indicate various success and failure scenarios,
// including device state, hardware issues, and permission problems.

// Door Open/Close Result - Shopping
type DoorOpenCloseShoppingResult struct {
	TradeRequestID  string       `json:"tradeRequestId"`  // Request ID
	OrderCode       string       `json:"orderCode"`       // Order number
	VmCode          string       `json:"vmCode"`          // Vending machine code
	MachineID       int          `json:"machineId"`       // Machine ID
	UserCode        string       `json:"userCode"`        // User code
	HandleStatus    int          `json:"handleStatus"`    // Handle status
	ShopMove        int          `json:"shopMove"`        // Shop movement status
	TotalFee        int          `json:"totalFee"`        // Total fee
	OpenDoorTime    int64        `json:"openDoorTime"`    // Door open time
	CloseDoorTime   int64        `json:"closeDoorTime"`   // Door close time
	OpenDoorWeight  int          `json:"openDoorWeight"`  // Weight when door opened
	CloseDoorWeight int          `json:"closeDoorWeight"` // Weight when door closed
	OrderGoodsList  []OrderGoods `json:"orderGoodsList"`  // List of ordered goods
}

// Door Open/Close Result - Restocking
type DoorOpenCloseRestockingResult struct {
	RequestID       string `json:"requestId"`       // Request ID
	VmCode          string `json:"vmCode"`          // Vending machine code
	OpenDoorTime    int64  `json:"openDoorTime"`    // Door open time
	CloseDoorTime   int64  `json:"closeDoorTime"`   // Door close time
	OpenDoorWeight  int    `json:"openDoorWeight"`  // Weight when door opened
	CloseDoorWeight int    `json:"closeDoorWeight"` // Weight when door closed
}

// Door Open/Close Response
type DoorOpenCloseResponse struct {
	Status  int         `json:"status"`  // Status code
	Message string      `json:"message"` // Status message
	Data    interface{} `json:"data"`    // Result data (Shopping or Restocking)
}
