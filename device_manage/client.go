package aifinitsdk_device

type VendingMachineManageClient interface {
	Activation()
	List(request *ListRequest) (*ListResponse, error)
	Detail() (*DetailResponse, error)
	DeviceInfo() (*DeviceInfoResult, error)
	PeopleFlow(request *PeopleFlowRequest) (*PeopleFlowResponse, error)
	Update(request *UpdateRequest) (*UpdateResponse, error)
	Control(request *ControlRequest) (*ControlResponse, error)
	Alarm()
	Setting()
}
