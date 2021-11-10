package device

type GetHostnameFunction struct{}

func (function *GetHostnameFunction) Request() interface{} {
	return &GetHostname{}
}

func (function *GetHostnameFunction) Response() interface{} {
	return &GetHostnameResponse{}
}

type SetHostnameFunction struct{}

func (function *SetHostnameFunction) Request() interface{} {
	return &SetHostname{}
}

func (function *SetHostnameFunction) Response() interface{} {
	return &SetHostnameResponse{}
}

type GetDNSFunction struct{}

func (function *GetDNSFunction) Request() interface{} {
	return &GetDNS{}
}

func (function *GetDNSFunction) Response() interface{} {
	return &GetDNSResponse{}
}

type SetDNSFunction struct{}

func (function *SetDNSFunction) Request() interface{} {
	return &SetDNS{}
}

func (function *SetDNSFunction) Response() interface{} {
	return &SetDNSResponse{}
}

type GetNetworkInterfacesFunction struct{}

func (function *GetNetworkInterfacesFunction) Request() interface{} {
	return &GetNetworkInterfaces{}
}

func (function *GetNetworkInterfacesFunction) Response() interface{} {
	return &GetNetworkInterfacesResponse{}
}

type SetNetworkInterfacesFunction struct{}

func (function *SetNetworkInterfacesFunction) Request() interface{} {
	return &SetNetworkInterfaces{}
}

func (function *SetNetworkInterfacesFunction) Response() interface{} {
	return &SetNetworkInterfacesResponse{}
}

type GetNetworkProtocolsFunction struct{}

func (function *GetNetworkProtocolsFunction) Request() interface{} {
	return &GetNetworkProtocols{}
}

func (function *GetNetworkProtocolsFunction) Response() interface{} {
	return &GetNetworkProtocolsResponse{}
}

type SetNetworkProtocolsFunction struct{}

func (function *SetNetworkProtocolsFunction) Request() interface{} {
	return &SetNetworkProtocols{}
}

func (function *SetNetworkProtocolsFunction) Response() interface{} {
	return &SetNetworkProtocolsResponse{}
}

type GetNetworkDefaultGatewayFunction struct{}

func (function *GetNetworkDefaultGatewayFunction) Request() interface{} {
	return &GetNetworkDefaultGateway{}
}

func (function *GetNetworkDefaultGatewayFunction) Response() interface{} {
	return &GetNetworkDefaultGatewayResponse{}
}

type SetNetworkDefaultGatewayFunction struct{}

func (function *SetNetworkDefaultGatewayFunction) Request() interface{} {
	return &SetNetworkDefaultGateway{}
}

func (function *SetNetworkDefaultGatewayFunction) Response() interface{} {
	return &SetNetworkDefaultGatewayResponse{}
}

type GetDeviceInformationFunction struct{}

func (function *GetDeviceInformationFunction) Request() interface{} {
	return &GetDeviceInformation{}
}

func (function *GetDeviceInformationFunction) Response() interface{} {
	return &GetDeviceInformationResponse{}
}

type GetSystemDateAndTimeFunction struct{}

func (function *GetSystemDateAndTimeFunction) Request() interface{} {
	return &GetSystemDateAndTime{}
}

func (function *GetSystemDateAndTimeFunction) Response() interface{} {
	return &GetSystemDateAndTimeResponse{}
}

type SetSystemDateAndTimeFunction struct{}

func (function *SetSystemDateAndTimeFunction) Request() interface{} {
	return &SetSystemDateAndTime{}
}

func (function *SetSystemDateAndTimeFunction) Response() interface{} {
	return &SetSystemDateAndTimeResponse{}
}

type SetSystemFactoryDefaultFunction struct{}

func (function *SetSystemFactoryDefaultFunction) Request() interface{} {
	return &SetSystemFactoryDefault{}
}

func (function *SetSystemFactoryDefaultFunction) Response() interface{} {
	return &SetSystemFactoryDefaultResponse{}
}

type SystemRebootFunction struct{}

func (function *SystemRebootFunction) Request() interface{} {
	return &SystemReboot{}
}

func (function *SystemRebootFunction) Response() interface{} {
	return &SystemRebootResponse{}
}

type GetUsersFunction struct{}

func (function *GetUsersFunction) Request() interface{} {
	return &GetUsers{}
}

func (function *GetUsersFunction) Response() interface{} {
	return &GetUsersResponse{}
}

type CreateUsersFunction struct{}

func (function *CreateUsersFunction) Request() interface{} {
	return &CreateUsers{}
}

func (function *CreateUsersFunction) Response() interface{} {
	return &CreateUsersResponse{}
}

type DeleteUsersFunction struct{}

func (function *DeleteUsersFunction) Request() interface{} {
	return &DeleteUsers{}
}

func (function *DeleteUsersFunction) Response() interface{} {
	return &DeleteUsersResponse{}
}

type SetUserFunction struct{}

func (function *SetUserFunction) Request() interface{} {
	return &SetUser{}
}

func (function *SetUserFunction) Response() interface{} {
	return &SetUserResponse{}
}

type GetDiscoveryModeFunction struct{}

func (function *GetDiscoveryModeFunction) Request() interface{} {
	return &GetDiscoveryMode{}
}

func (function *GetDiscoveryModeFunction) Response() interface{} {
	return &GetDiscoveryModeResponse{}
}

type SetDiscoveryModeFunction struct{}

func (function *SetDiscoveryModeFunction) Request() interface{} {
	return &SetDiscoveryMode{}
}

func (function *SetDiscoveryModeFunction) Response() interface{} {
	return &SetDiscoveryModeResponse{}
}

type GetScopesFunction struct{}

func (function *GetScopesFunction) Request() interface{} {
	return &GetScopes{}
}

func (function *GetScopesFunction) Response() interface{} {
	return &GetScopesResponse{}
}

type SetScopesFunction struct{}

func (function *SetScopesFunction) Request() interface{} {
	return &SetScopes{}
}

func (function *SetScopesFunction) Response() interface{} {
	return &SetScopesResponse{}
}

type AddScopesFunction struct{}

func (function *AddScopesFunction) Request() interface{} {
	return &AddScopes{}
}

func (function *AddScopesFunction) Response() interface{} {
	return &AddScopesResponse{}
}

type RemoveScopesFunction struct{}

func (function *RemoveScopesFunction) Request() interface{} {
	return &RemoveScopes{}
}

func (function *RemoveScopesFunction) Response() interface{} {
	return &RemoveScopesResponse{}
}
