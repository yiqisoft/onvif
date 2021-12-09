package media2

type GetProfilesFunction struct{}

func (function *GetProfilesFunction) Request() interface{} {
	return &GetProfiles{}
}

func (function *GetProfilesFunction) Response() interface{} {
	return &GetProfilesResponse{}
}

type GetAnalyticsConfigurationsFunction struct{}

func (function *GetAnalyticsConfigurationsFunction) Request() interface{} {
	return &GetAnalyticsConfigurations{}
}

func (function *GetAnalyticsConfigurationsFunction) Response() interface{} {
	return &GetAnalyticsConfigurationsResponse{}
}

type AddConfigurationFunction struct{}

func (function *AddConfigurationFunction) Request() interface{} {
	return &AddConfiguration{}
}

func (function *AddConfigurationFunction) Response() interface{} {
	return &AddConfigurationResponse{}
}

type RemoveConfigurationFunction struct{}

func (function *RemoveConfigurationFunction) Request() interface{} {
	return &RemoveConfiguration{}
}

func (function *RemoveConfigurationFunction) Response() interface{} {
	return &RemoveConfigurationResponse{}
}
