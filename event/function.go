package event

type GetEventPropertiesFunction struct{}

func (function *GetEventPropertiesFunction) Request() interface{} {
	return &GetEventProperties{}
}

func (function *GetEventPropertiesFunction) Response() interface{} {
	return &GetEventPropertiesResponse{}
}

type CreatePullPointSubscriptionFunction struct{}

func (function *CreatePullPointSubscriptionFunction) Request() interface{} {
	return &CreatePullPointSubscription{}
}

func (function *CreatePullPointSubscriptionFunction) Response() interface{} {
	return &CreatePullPointSubscriptionResponse{}
}

type PullMessagesFunction struct{}

func (function *PullMessagesFunction) Request() interface{} {
	return &PullMessages{}
}

func (function *PullMessagesFunction) Response() interface{} {
	return &PullMessagesResponse{}
}

type UnsubscribeFunction struct{}

func (function *UnsubscribeFunction) Request() interface{} {
	return &Unsubscribe{}
}

func (function *UnsubscribeFunction) Response() interface{} {
	return &UnsubscribeResponse{}
}

type SubscribeFunction struct{}

func (function *SubscribeFunction) Request() interface{} {
	return &Subscribe{}
}

func (function *SubscribeFunction) Response() interface{} {
	return &SubscribeResponse{}
}

type RenewFunction struct{}

func (function *RenewFunction) Request() interface{} {
	return &Renew{}
}

func (function *RenewFunction) Response() interface{} {
	return &RenewResponse{}
}
