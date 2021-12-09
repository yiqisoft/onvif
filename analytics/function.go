package analytics

type GetSupportedAnalyticsModulesFunction struct{}

func (function *GetSupportedAnalyticsModulesFunction) Request() interface{} {
	return &GetSupportedAnalyticsModules{}
}

func (function *GetSupportedAnalyticsModulesFunction) Response() interface{} {
	return &GetSupportedAnalyticsModulesResponse{}
}

type GetAnalyticsModulesFunction struct{}

func (function *GetAnalyticsModulesFunction) Request() interface{} {
	return &GetAnalyticsModules{}
}

func (function *GetAnalyticsModulesFunction) Response() interface{} {
	return &GetAnalyticsModulesResponse{}
}

type CreateAnalyticsModulesFunction struct{}

func (function *CreateAnalyticsModulesFunction) Request() interface{} {
	return &CreateAnalyticsModules{}
}

func (function *CreateAnalyticsModulesFunction) Response() interface{} {
	return &CreateAnalyticsModulesResponse{}
}

type DeleteAnalyticsModulesFunction struct{}

func (function *DeleteAnalyticsModulesFunction) Request() interface{} {
	return &DeleteAnalyticsModules{}
}

func (function *DeleteAnalyticsModulesFunction) Response() interface{} {
	return &DeleteAnalyticsModulesResponse{}
}

type GetAnalyticsModuleOptionsFunction struct{}

func (function *GetAnalyticsModuleOptionsFunction) Request() interface{} {
	return &GetAnalyticsModuleOptions{}
}

func (function *GetAnalyticsModuleOptionsFunction) Response() interface{} {
	return &GetAnalyticsModuleOptionsResponse{}
}

type ModifyAnalyticsModulesFunction struct{}

func (function *ModifyAnalyticsModulesFunction) Request() interface{} {
	return &ModifyAnalyticsModules{}
}

func (function *ModifyAnalyticsModulesFunction) Response() interface{} {
	return &ModifyAnalyticsModulesResponse{}
}

type GetSupportedRulesFunction struct{}

func (function *GetSupportedRulesFunction) Request() interface{} {
	return &GetSupportedRules{}
}

func (function *GetSupportedRulesFunction) Response() interface{} {
	return &GetSupportedRulesResponse{}
}

type GetRulesFunction struct{}

func (function *GetRulesFunction) Request() interface{} {
	return &GetRules{}
}

func (function *GetRulesFunction) Response() interface{} {
	return &GetRulesResponse{}
}

type CreateRulesFunction struct{}

func (function *CreateRulesFunction) Request() interface{} {
	return &CreateRules{}
}

func (function *CreateRulesFunction) Response() interface{} {
	return &CreateRulesResponse{}
}

type DeleteRulesFunction struct{}

func (function *DeleteRulesFunction) Request() interface{} {
	return &DeleteRules{}
}

func (function *DeleteRulesFunction) Response() interface{} {
	return &DeleteRulesResponse{}
}

type GetRuleOptionsFunction struct{}

func (function *GetRuleOptionsFunction) Request() interface{} {
	return &GetRuleOptions{}
}

func (function *GetRuleOptionsFunction) Response() interface{} {
	return &GetRuleOptionsResponse{}
}

type ModifyRulesFunction struct{}

func (function *ModifyRulesFunction) Request() interface{} {
	return &ModifyRules{}
}

func (function *ModifyRulesFunction) Response() interface{} {
	return &ModifyRulesResponse{}
}
