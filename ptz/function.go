package ptz

type GetNodesFunction struct{}

func (function *GetNodesFunction) Request() interface{} {
	return &GetNodes{}
}

func (function *GetNodesFunction) Response() interface{} {
	return &GetNodesResponse{}
}

type GetNodeFunction struct{}

func (function *GetNodeFunction) Request() interface{} {
	return &GetNode{}
}

func (function *GetNodeFunction) Response() interface{} {
	return &GetNodeResponse{}
}

type GetConfigurationsFunction struct{}

func (function *GetConfigurationsFunction) Request() interface{} {
	return &GetConfigurations{}
}

func (function *GetConfigurationsFunction) Response() interface{} {
	return &GetConfigurationsResponse{}
}

type GetConfigurationFunction struct{}

func (function *GetConfigurationFunction) Request() interface{} {
	return &GetConfiguration{}
}

func (function *GetConfigurationFunction) Response() interface{} {
	return &GetConfigurationResponse{}
}

type GetConfigurationOptionsFunction struct{}

func (function *GetConfigurationOptionsFunction) Request() interface{} {
	return &GetConfigurationOptions{}
}

func (function *GetConfigurationOptionsFunction) Response() interface{} {
	return &GetConfigurationOptionsResponse{}
}

type SetConfigurationFunction struct{}

func (function *SetConfigurationFunction) Request() interface{} {
	return &SetConfiguration{}
}

func (function *SetConfigurationFunction) Response() interface{} {
	return &SetConfigurationResponse{}
}

type AbsoluteMoveFunction struct{}

func (function *AbsoluteMoveFunction) Request() interface{} {
	return &AbsoluteMove{}
}

func (function *AbsoluteMoveFunction) Response() interface{} {
	return &AbsoluteMoveResponse{}
}

type RelativeMoveFunction struct{}

func (function *RelativeMoveFunction) Request() interface{} {
	return &RelativeMove{}
}

func (function *RelativeMoveFunction) Response() interface{} {
	return &RelativeMoveResponse{}
}

type ContinuousMoveFunction struct{}

func (function *ContinuousMoveFunction) Request() interface{} {
	return &ContinuousMove{}
}

func (function *ContinuousMoveFunction) Response() interface{} {
	return &ContinuousMoveResponse{}
}

type StopFunction struct{}

func (function *StopFunction) Request() interface{} {
	return &Stop{}
}

func (function *StopFunction) Response() interface{} {
	return &StopResponse{}
}

type GetStatusFunction struct{}

func (function *GetStatusFunction) Request() interface{} {
	return &GetStatus{}
}

func (function *GetStatusFunction) Response() interface{} {
	return &GetStatusResponse{}
}

type SetPresetFunction struct{}

func (function *SetPresetFunction) Request() interface{} {
	return &SetPreset{}
}

func (function *SetPresetFunction) Response() interface{} {
	return &SetPresetResponse{}
}

type GetPresetsFunction struct{}

func (function *GetPresetsFunction) Request() interface{} {
	return &GetPresets{}
}

func (function *GetPresetsFunction) Response() interface{} {
	return &GetPresetsResponse{}
}

type GotoPresetFunction struct{}

func (function *GotoPresetFunction) Request() interface{} {
	return &GotoPreset{}
}

func (function *GotoPresetFunction) Response() interface{} {
	return &GotoPresetResponse{}
}

type RemovePresetFunction struct{}

func (function *RemovePresetFunction) Request() interface{} {
	return &RemovePreset{}
}

func (function *RemovePresetFunction) Response() interface{} {
	return &RemovePresetResponse{}
}

type GotoHomePositionFunction struct{}

func (function *GotoHomePositionFunction) Request() interface{} {
	return &GotoHomePosition{}
}

func (function *GotoHomePositionFunction) Response() interface{} {
	return &GotoHomePositionResponse{}
}

type SetHomePositionFunction struct{}

func (function *SetHomePositionFunction) Request() interface{} {
	return &SetHomePosition{}
}

func (function *SetHomePositionFunction) Response() interface{} {
	return &SetHomePositionResponse{}
}

type SendAuxiliaryCommandFunction struct{}

func (function *SendAuxiliaryCommandFunction) Request() interface{} {
	return &SendAuxiliaryCommand{}
}

func (function *SendAuxiliaryCommandFunction) Response() interface{} {
	return &SendAuxiliaryCommandResponse{}
}
