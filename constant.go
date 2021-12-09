package onvif

// Onvif WebService
const (
	DeviceWebService    = "Device"
	EventWebService     = "Event"
	MediaWebService     = "Media"
	Media2WebService    = "Media2"
	PTZWebService       = "PTZ"
	AnalyticsWebService = "Analytics"
)

// Onvif WebService function
const (
	// WebService - Device
	GetHostname              = "GetHostname"
	SetHostname              = "SetHostname"
	GetDNS                   = "GetDNS"
	SetDNS                   = "SetDNS"
	GetNetworkInterfaces     = "GetNetworkInterfaces"
	SetNetworkInterfaces     = "SetNetworkInterfaces"
	GetNetworkProtocols      = "GetNetworkProtocols"
	SetNetworkProtocols      = "SetNetworkProtocols"
	GetNetworkDefaultGateway = "GetNetworkDefaultGateway"
	SetNetworkDefaultGateway = "SetNetworkDefaultGateway"

	GetDeviceInformation    = "GetDeviceInformation"
	GetSystemDateAndTime    = "GetSystemDateAndTime"
	SetSystemDateAndTime    = "SetSystemDateAndTime"
	SetSystemFactoryDefault = "SetSystemFactoryDefault"
	SystemReboot            = "SystemReboot"

	GetUsers    = "GetUsers"
	CreateUsers = "CreateUsers"
	DeleteUsers = "DeleteUsers"
	SetUser     = "SetUser"

	GetDiscoveryMode = "GetDiscoveryMode"
	SetDiscoveryMode = "SetDiscoveryMode"
	GetScopes        = "GetScopes"
	SetScopes        = "SetScopes"
	AddScopes        = "AddScopes"
	RemoveScopes     = "RemoveScopes"

	// WebService - Media
	GetMetadataConfiguration            = "GetMetadataConfiguration"
	GetMetadataConfigurations           = "GetMetadataConfigurations"
	AddMetadataConfiguration            = "AddMetadataConfiguration"
	RemoveMetadataConfiguration         = "RemoveMetadataConfiguration"
	SetMetadataConfiguration            = "SetMetadataConfiguration"
	GetCompatibleMetadataConfigurations = "GetCompatibleMetadataConfigurations"
	GetMetadataConfigurationOptions     = "GetMetadataConfigurationOptions"
	GetProfiles                         = "GetProfiles"
	GetStreamUri                        = "GetStreamUri"
	GetVideoEncoderConfiguration        = "GetVideoEncoderConfiguration"
	GetVideoEncoderConfigurations       = "GetVideoEncoderConfigurations"
	SetVideoEncoderConfiguration        = "SetVideoEncoderConfiguration"
	GetVideoEncoderConfigurationOptions = "GetVideoEncoderConfigurationOptions"

	// WebService - Media2
	GetAnalyticsConfigurations = "GetAnalyticsConfigurations"
	AddConfiguration           = "AddConfiguration"
	RemoveConfiguration        = "RemoveConfiguration"

	// WebService - PTZ
	GetNodes                = "GetNodes"
	GetNode                 = "GetNode"
	GetConfigurations       = "GetConfigurations"
	GetConfiguration        = "GetConfiguration"
	GetConfigurationOptions = "GetConfigurationOptions"
	SetConfiguration        = "SetConfiguration"
	AddPTZConfiguration     = "AddPTZConfiguration"
	RemovePTZConfiguration  = "RemovePTZConfiguration"
	AbsoluteMove            = "AbsoluteMove"
	RelativeMove            = "RelativeMove"
	ContinuousMove          = "ContinuousMove"
	Stop                    = "Stop"
	GetStatus               = "GetStatus"
	SetPreset               = "SetPreset"
	GetPresets              = "GetPresets"
	GotoPreset              = "GotoPreset"
	RemovePreset            = "RemovePreset"
	GotoHomePosition        = "GotoHomePosition"
	SetHomePosition         = "SetHomePosition"
	SendAuxiliaryCommand    = "SendAuxiliaryCommand"

	// WebService - Event
	GetEventProperties          = "GetEventProperties"
	CreatePullPointSubscription = "CreatePullPointSubscription"
	PullMessages                = "PullMessages"
	Unsubscribe                 = "Unsubscribe"
	Subscribe                   = "Subscribe"
	Renew                       = "Renew"

	// WebService - Analytics
	GetSupportedAnalyticsModules = "GetSupportedAnalyticsModules"
	GetAnalyticsModules          = "GetAnalyticsModules"
	CreateAnalyticsModules       = "CreateAnalyticsModules"
	DeleteAnalyticsModules       = "DeleteAnalyticsModules"
	GetAnalyticsModuleOptions    = "GetAnalyticsModuleOptions"
	ModifyAnalyticsModules       = "ModifyAnalyticsModules"

	GetSupportedRules = "GetSupportedRules"
	GetRules          = "GetRules"
	CreateRules       = "CreateRules"
	DeleteRules       = "DeleteRules"
	GetRuleOptions    = "GetRuleOptions"
	ModifyRules       = "ModifyRules"
)

// Onvif Auth Mode
const (
	DigestAuth        = "digest"
	UsernameTokenAuth = "usernametoken"
	Both              = "both"
	NoAuth            = "none"
)
