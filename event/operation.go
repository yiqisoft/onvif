package event

import (
	"github.com/IOTechSystems/onvif/xsd"
)

//GetServiceCapabilities action
type GetServiceCapabilities struct {
	XMLName string `xml:"tev:GetServiceCapabilities"`
}

//GetServiceCapabilitiesResponse type
type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities
}

//SubscriptionPolicy action
type SubscriptionPolicy struct { //tev http://www.onvif.org/ver10/events/wsdl
	ChangedOnly xsd.Boolean `xml:"ChangedOnly,attr"`
	string
}

//Subscribe action for subscribe event topic
type Subscribe struct { //http://docs.oasis-open.org/wsn/b-2.xsd
	XMLName            struct{}               `xml:"wsnt:Subscribe"`
	ConsumerReference  *EndpointReferenceType `xml:"wsnt:ConsumerReference"`
	Filter             *FilterType            `xml:"wsnt:Filter"`
	SubscriptionPolicy *xsd.String            `xml:"wsnt:SubscriptionPolicy"`
	TerminationTime    *xsd.String            `xml:"wsnt:TerminationTime"`
}

//SubscribeResponse message for subscribe event topic
type SubscribeResponse struct { //http://docs.oasis-open.org/wsn/b-2.xsd
	SubscriptionReference SubscriptionReferenceResponse
	CurrentTime           *xsd.String
	TerminationTime       *xsd.String
}

//Renew action for refresh event topic subscription
type Renew struct { //http://docs.oasis-open.org/wsn/b-2.xsd
	XMLName         string     `xml:"wsnt:Renew"`
	TerminationTime xsd.String `xml:"wsnt:TerminationTime"`
}

//RenewResponse for Renew action
type RenewResponse struct { //http://docs.oasis-open.org/wsn/b-2.xsd
	TerminationTime TerminationTime `xml:"wsnt:TerminationTime"`
	CurrentTime     CurrentTime     `xml:"wsnt:CurrentTime"`
}

//Unsubscribe action for Unsubscribe event topic
type Unsubscribe struct { //http://docs.oasis-open.org/wsn/b-2.xsd
	XMLName string `xml:"tev:Unsubscribe"`
	Any     string
}

//UnsubscribeResponse message for Unsubscribe event topic
type UnsubscribeResponse struct { //http://docs.oasis-open.org/wsn/b-2.xsd
	Any string
}

//CreatePullPointSubscription action
//BUG(r) Bad AbsoluteOrRelativeTimeType type
type CreatePullPointSubscription struct {
	XMLName                string      `xml:"tev:CreatePullPointSubscription,omitempty"`
	Filter                 *FilterType `xml:"tev:Filter,omitempty"`
	InitialTerminationTime *xsd.String `xml:"tev:InitialTerminationTime,omitempty"`
	SubscriptionPolicy     *xsd.String `xml:"tev:SubscriptionPolicy,omitempty"`
}

//CreatePullPointSubscriptionResponse action
type CreatePullPointSubscriptionResponse struct {
	SubscriptionReference SubscriptionReferenceResponse
	CurrentTime           CurrentTime
	TerminationTime       TerminationTime
}

type SubscriptionReferenceResponse struct {
	Address             AttributedURIType
	ReferenceParameters *ReferenceParametersType
	Metadata            *MetadataType
}

//GetEventProperties action
type GetEventProperties struct {
	XMLName string `xml:"tev:GetEventProperties"`
}

//GetEventPropertiesResponse action
type GetEventPropertiesResponse struct {
	TopicNamespaceLocation          *xsd.AnyURI             `json:",omitempty" xml:",omitempty"`
	FixedTopicSet                   *FixedTopicSet          `json:",omitempty" xml:",omitempty"`
	TopicSet                        *TopicSet               `json:",omitempty" xml:",omitempty"`
	TopicExpressionDialect          *TopicExpressionDialect `json:",omitempty" xml:",omitempty"`
	MessageContentFilterDialect     *xsd.AnyURI             `json:",omitempty" xml:",omitempty"`
	ProducerPropertiesFilterDialect *xsd.AnyURI             `json:",omitempty" xml:",omitempty"`
	MessageContentSchemaLocation    *xsd.AnyURI             `json:",omitempty" xml:",omitempty"`
}

//Port type PullPointSubscription

//PullMessages Action
type PullMessages struct {
	XMLName      string       `xml:"tev:PullMessages"`
	Timeout      xsd.Duration `xml:"tev:Timeout"`
	MessageLimit xsd.Int      `xml:"tev:MessageLimit"`
}

//PullMessagesResponse response type
type PullMessagesResponse struct {
	CurrentTime         *xsd.String           `json:",omitempty" xml:",omitempty"`
	TerminationTime     *xsd.String           `json:",omitempty" xml:",omitempty"`
	NotificationMessage []NotificationMessage `json:",omitempty" xml:",omitempty"`
}

//PullMessagesFaultResponse response type
type PullMessagesFaultResponse struct {
	MaxTimeout      xsd.Duration
	MaxMessageLimit xsd.Int
}

//Seek action
type Seek struct {
	XMLName string       `xml:"tev:Seek"`
	UtcTime xsd.DateTime `xml:"tev:UtcTime"`
	Reverse xsd.Boolean  `xml:"tev:Reverse"`
}

//SeekResponse action
type SeekResponse struct {
}

//SetSynchronizationPoint action
type SetSynchronizationPoint struct {
	XMLName string `xml:"tev:SetSynchronizationPoint"`
}

//SetSynchronizationPointResponse action
type SetSynchronizationPointResponse struct {
}

// Notify type
type Notify struct {
	NotificationMessage []NotificationMessage `json:",omitempty" xml:",omitempty"`
}
