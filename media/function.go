package media

type GetMetadataConfigurationFunction struct{}

func (function *GetMetadataConfigurationFunction) Request() interface{} {
	return &GetMetadataConfiguration{}
}

func (function *GetMetadataConfigurationFunction) Response() interface{} {
	return &GetMetadataConfigurationResponse{}
}

type GetMetadataConfigurationsFunction struct{}

func (function *GetMetadataConfigurationsFunction) Request() interface{} {
	return &GetMetadataConfigurations{}
}

func (function *GetMetadataConfigurationsFunction) Response() interface{} {
	return &GetMetadataConfigurationsResponse{}
}

type AddMetadataConfigurationFunction struct{}

func (function *AddMetadataConfigurationFunction) Request() interface{} {
	return &AddMetadataConfiguration{}
}

func (function *AddMetadataConfigurationFunction) Response() interface{} {
	return &AddMetadataConfigurationResponse{}
}

type RemoveMetadataConfigurationFunction struct{}

func (function *RemoveMetadataConfigurationFunction) Request() interface{} {
	return &RemoveMetadataConfiguration{}
}

func (function *RemoveMetadataConfigurationFunction) Response() interface{} {
	return &RemoveMetadataConfigurationResponse{}
}

type SetMetadataConfigurationFunction struct{}

func (function *SetMetadataConfigurationFunction) Request() interface{} {
	return &SetMetadataConfiguration{}
}

func (function *SetMetadataConfigurationFunction) Response() interface{} {
	return &SetMetadataConfigurationResponse{}
}

type GetCompatibleMetadataConfigurationsFunction struct{}

func (function *GetCompatibleMetadataConfigurationsFunction) Request() interface{} {
	return &GetCompatibleMetadataConfigurations{}
}

func (function *GetCompatibleMetadataConfigurationsFunction) Response() interface{} {
	return &GetCompatibleMetadataConfigurationsResponse{}
}

type GetMetadataConfigurationOptionsFunction struct{}

func (function *GetMetadataConfigurationOptionsFunction) Request() interface{} {
	return &GetMetadataConfigurationOptions{}
}

func (function *GetMetadataConfigurationOptionsFunction) Response() interface{} {
	return &GetMetadataConfigurationOptionsResponse{}
}

type GetProfilesFunction struct{}

func (function *GetProfilesFunction) Request() interface{} {
	return &GetProfiles{}
}

func (function *GetProfilesFunction) Response() interface{} {
	return &GetProfilesResponse{}
}

type GetStreamUriFunction struct{}

func (function *GetStreamUriFunction) Request() interface{} {
	return &GetStreamUri{}
}

func (function *GetStreamUriFunction) Response() interface{} {
	return &GetStreamUriResponse{}
}

type GetVideoEncoderConfigurationFunction struct{}

func (function *GetVideoEncoderConfigurationFunction) Request() interface{} {
	return &GetVideoEncoderConfiguration{}
}

func (function *GetVideoEncoderConfigurationFunction) Response() interface{} {
	return &GetVideoEncoderConfigurationResponse{}
}

type SetVideoEncoderConfigurationFunction struct{}

func (function *SetVideoEncoderConfigurationFunction) Request() interface{} {
	return &SetVideoEncoderConfiguration{}
}

func (function *SetVideoEncoderConfigurationFunction) Response() interface{} {
	return &SetVideoEncoderConfigurationResponse{}
}

type GetVideoEncoderConfigurationOptionsFunction struct{}

func (function *GetVideoEncoderConfigurationOptionsFunction) Request() interface{} {
	return &GetVideoEncoderConfigurationOptions{}
}

func (function *GetVideoEncoderConfigurationOptionsFunction) Response() interface{} {
	return &GetVideoEncoderConfigurationOptionsResponse{}
}
