package onvif

import (
	"github.com/IOTechSystems/onvif/device"
	"github.com/IOTechSystems/onvif/media"
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
	GetProfiles:  &media.GetProfilesFunction{},
	GetStreamUri: &media.GetStreamUriFunction{},
	// Video Encoder Configuration
	GetVideoEncoderConfiguration:        &media.GetVideoEncoderConfigurationFunction{},
	SetVideoEncoderConfiguration:        &media.SetVideoEncoderConfigurationFunction{},
	GetVideoEncoderConfigurationOptions: &media.GetVideoEncoderConfigurationOptionsFunction{},
	// PTZ Configuration
	AddPTZConfiguration:    &media.AddPTZConfigurationFunction{},
	RemovePTZConfiguration: &media.RemovePTZConfigurationFunction{},
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
