package onvif

import (
	"fmt"

	"github.com/IOTechSystems/onvif/analytics"
	"github.com/IOTechSystems/onvif/device"
	"github.com/IOTechSystems/onvif/event"
	"github.com/IOTechSystems/onvif/media"
	"github.com/IOTechSystems/onvif/media2"
	"github.com/IOTechSystems/onvif/ptz"
)

var DeviceFunctionMap = map[string]Function{
	// Network Configuration
	GetHostname:              &device.GetHostnameFunction{},
	SetHostname:              &device.SetHostnameFunction{},
	GetDNS:                   &device.GetDNSFunction{},
	SetDNS:                   &device.SetDNSFunction{},
	GetNetworkInterfaces:     &device.GetNetworkInterfacesFunction{},
	SetNetworkInterfaces:     &device.SetNetworkInterfacesFunction{},
	GetNetworkProtocols:      &device.GetNetworkProtocolsFunction{},
	SetNetworkProtocols:      &device.SetNetworkProtocolsFunction{},
	GetNetworkDefaultGateway: &device.GetNetworkDefaultGatewayFunction{},
	SetNetworkDefaultGateway: &device.SetNetworkDefaultGatewayFunction{},
	// System Function
	GetDeviceInformation:    &device.GetDeviceInformationFunction{},
	GetSystemDateAndTime:    &device.GetSystemDateAndTimeFunction{},
	SetSystemDateAndTime:    &device.SetSystemDateAndTimeFunction{},
	SetSystemFactoryDefault: &device.SetSystemFactoryDefaultFunction{},
	SystemReboot:            &device.SystemRebootFunction{},
	// User Handling
	GetUsers:    &device.GetUsersFunction{},
	SetUser:     &device.SetUserFunction{},
	CreateUsers: &device.CreateUsersFunction{},
	DeleteUsers: &device.DeleteUsersFunction{},
	// Auto Discovery
	GetDiscoveryMode: &device.GetDiscoveryModeFunction{},
	SetDiscoveryMode: &device.SetDiscoveryModeFunction{},
	GetScopes:        &device.GetScopesFunction{},
	SetScopes:        &device.SetScopesFunction{},
	AddScopes:        &device.AddScopesFunction{},
	RemoveScopes:     &device.RemoveScopesFunction{},

	GetEndpointReference: &device.GetEndpointReferenceFunction{},
}

var MediaFunctionMap = map[string]Function{
	// Metadata Configuration
	GetMetadataConfiguration:            &media.GetMetadataConfigurationFunction{},
	GetMetadataConfigurations:           &media.GetMetadataConfigurationsFunction{},
	AddMetadataConfiguration:            &media.AddMetadataConfigurationFunction{},
	RemoveMetadataConfiguration:         &media.RemoveMetadataConfigurationFunction{},
	SetMetadataConfiguration:            &media.SetMetadataConfigurationFunction{},
	GetCompatibleMetadataConfigurations: &media.GetCompatibleMetadataConfigurationsFunction{},
	GetMetadataConfigurationOptions:     &media.GetMetadataConfigurationOptionsFunction{},
	// Video Streaming
	GetProfiles:    &media.GetProfilesFunction{},
	GetStreamUri:   &media.GetStreamUriFunction{},
	GetSnapshotUri: &media.GetSnapshotUriFunction{},
	// Video Encoder Configuration
	GetVideoEncoderConfiguration:        &media.GetVideoEncoderConfigurationFunction{},
	GetVideoEncoderConfigurations:       &media.GetVideoEncoderConfigurationsFunction{},
	SetVideoEncoderConfiguration:        &media.SetVideoEncoderConfigurationFunction{},
	GetVideoEncoderConfigurationOptions: &media.GetVideoEncoderConfigurationOptionsFunction{},
	// PTZ Configuration
	AddPTZConfiguration:    &media.AddPTZConfigurationFunction{},
	RemovePTZConfiguration: &media.RemovePTZConfigurationFunction{},
}

var Media2FunctionMap = map[string]Function{
	GetProfiles:                &media2.GetProfilesFunction{},
	GetAnalyticsConfigurations: &media2.GetAnalyticsConfigurationsFunction{},
	AddConfiguration:           &media2.AddConfigurationFunction{},
	RemoveConfiguration:        &media2.RemoveConfigurationFunction{},
}

var PTZFunctionMap = map[string]Function{
	// PTZ Capabilities
	GetNodes: &ptz.GetNodesFunction{},
	GetNode:  &ptz.GetNodeFunction{},
	// PTZ Configuration
	GetConfigurations:       &ptz.GetConfigurationsFunction{},
	GetConfiguration:        &ptz.GetConfigurationFunction{},
	GetConfigurationOptions: &ptz.GetConfigurationOptionsFunction{},
	SetConfiguration:        &ptz.SetConfigurationFunction{},
	// PTZ Actuation
	AbsoluteMove:   &ptz.AbsoluteMoveFunction{},
	RelativeMove:   &ptz.RelativeMoveFunction{},
	ContinuousMove: &ptz.ContinuousMoveFunction{},
	Stop:           &ptz.StopFunction{},
	GetStatus:      &ptz.GetStatusFunction{},
	// PTZ Preset
	SetPreset:    &ptz.SetPresetFunction{},
	GetPresets:   &ptz.GetPresetsFunction{},
	GotoPreset:   &ptz.GotoPresetFunction{},
	RemovePreset: &ptz.RemovePresetFunction{},
	// PTZ Home Position
	GotoHomePosition: &ptz.GotoHomePositionFunction{},
	SetHomePosition:  &ptz.SetHomePositionFunction{},
	// PTZ Auxiliary Operations
	SendAuxiliaryCommand: &ptz.SendAuxiliaryCommandFunction{},
}

var EventFunctionMap = map[string]Function{
	GetEventProperties:          &event.GetEventPropertiesFunction{},
	CreatePullPointSubscription: &event.CreatePullPointSubscriptionFunction{},
	PullMessages:                &event.PullMessagesFunction{},
	Unsubscribe:                 &event.UnsubscribeFunction{},
	Subscribe:                   &event.SubscribeFunction{},
	Renew:                       &event.RenewFunction{},
}

var AnalyticsFunctionMap = map[string]Function{
	// Analytics Module
	GetSupportedAnalyticsModules: &analytics.GetSupportedAnalyticsModulesFunction{},
	GetAnalyticsModules:          &analytics.GetAnalyticsModulesFunction{},
	CreateAnalyticsModules:       &analytics.CreateAnalyticsModulesFunction{},
	DeleteAnalyticsModules:       &analytics.DeleteAnalyticsModulesFunction{},
	GetAnalyticsModuleOptions:    &analytics.GetAnalyticsModuleOptionsFunction{},
	ModifyAnalyticsModules:       &analytics.ModifyAnalyticsModulesFunction{},
	// Rule
	GetSupportedRules: &analytics.GetSupportedRulesFunction{},
	GetRules:          &analytics.GetRulesFunction{},
	CreateRules:       &analytics.CreateRulesFunction{},
	DeleteRules:       &analytics.DeleteRulesFunction{},
	GetRuleOptions:    &analytics.GetRuleOptionsFunction{},
	ModifyRules:       &analytics.ModifyRulesFunction{},
}

func FunctionByServiceAndFunctionName(serviceName, functionName string) (Function, error) {
	var function Function
	var exist bool
	switch serviceName {
	case DeviceWebService:
		function, exist = DeviceFunctionMap[functionName]
		if !exist {
			return nil, fmt.Errorf("the web service '%s'not support the function '%s'", serviceName, functionName)
		}
	case MediaWebService:
		function, exist = MediaFunctionMap[functionName]
		if !exist {
			return nil, fmt.Errorf("the web service '%s' not support the function '%s'", serviceName, functionName)
		}
	case Media2WebService:
		function, exist = Media2FunctionMap[functionName]
		if !exist {
			return nil, fmt.Errorf("the web service '%s' not support the function '%s'", serviceName, functionName)
		}
	case PTZWebService:
		function, exist = PTZFunctionMap[functionName]
		if !exist {
			return nil, fmt.Errorf("the web service '%s' not support the function '%s'", serviceName, functionName)
		}
	case EventWebService:
		function, exist = EventFunctionMap[functionName]
		if !exist {
			return nil, fmt.Errorf("the web service '%s' not support the function '%s'", serviceName, functionName)
		}
	case AnalyticsWebService:
		function, exist = AnalyticsFunctionMap[functionName]
		if !exist {
			return nil, fmt.Errorf("the web service '%s' not support the function '%s'", serviceName, functionName)
		}
	default:
		return nil, fmt.Errorf("not support the web service '%s'", serviceName)
	}
	return function, nil
}
