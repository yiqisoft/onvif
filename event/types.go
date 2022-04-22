package event

import (
	"encoding/xml"
	"fmt"
	"reflect"

	"github.com/IOTechSystems/onvif/event/topic"
	"github.com/IOTechSystems/onvif/xsd"
	mv "github.com/clbanning/mxj/v2"
)

//Address Alias
type Address xsd.String

//CurrentTime alias
type CurrentTime xsd.DateTime //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
//TerminationTime alias
type TerminationTime xsd.DateTime //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
//FixedTopicSet alias
type FixedTopicSet xsd.Boolean //wsnt http://docs.oasis-open.org/wsn/b-2.xsd

//Documentation alias
type Documentation xsd.AnyType //wstop http://docs.oasis-open.org/wsn/t-1.xsd

//TopicExpressionDialect alias
type TopicExpressionDialect xsd.AnyURI

//Message alias
type Message xsd.AnyType

//ActionType for AttributedURIType
type ActionType AttributedURIType

//AttributedURIType in ws-addr
type AttributedURIType xsd.AnyURI //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd

//AbsoluteOrRelativeTimeType <xsd:union memberTypes="xsd:dateTime xsd:duration"/>
type AbsoluteOrRelativeTimeType struct { //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
	xsd.DateTime
	xsd.Duration
}

//EndpointReferenceType in ws-addr
type EndpointReferenceType struct { //wsa http://www.w3.org/2005/08/addressing/ws-addr.xsd
	Address             AttributedURIType `xml:"wsa:Address"`
	ReferenceParameters *ReferenceParametersType
	Metadata            *MetadataType `xml:"Metadata"`
}

// FilterType struct
type FilterType struct {
	TopicExpression *TopicExpressionType `xml:"wsnt:TopicExpression,omitempty"`
	MessageContent  *QueryExpressionType `xml:"wsnt:MessageContent,omitempty"`
}

//EndpointReference alais
type EndpointReference EndpointReferenceType

//ReferenceParametersType in ws-addr
type ReferenceParametersType struct { //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd
	//Here can be anyAttribute
}

//Metadata in ws-addr
type Metadata MetadataType //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd

//MetadataType in ws-addr
type MetadataType struct { //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd

	//Here can be anyAttribute
}

//TopicSet alias
type TopicSet map[string]interface{} //wstop http://docs.oasis-open.org/wsn/t-1.xsd

type Node struct {
	XMLName xml.Name
	Content []byte `xml:",innerxml"`
}

func (n *TopicSet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	node := Node{}
	err := d.DecodeElement(&node, &start)
	if err != nil {
		return err
	}
	wrapper := "root" // The TopicSet is an array, we need to wrap with a tag for XML parsing
	c := fmt.Sprintf("<%s>%s</%s>", wrapper, node.Content, wrapper)
	result, err := mv.NewMapXmlSeq([]byte(c))
	if err != nil {
		return err
	}
	if result[wrapper] != nil && reflect.ValueOf(result[wrapper]).Kind() == reflect.Map {
		*n = (result[wrapper]).(map[string]interface{})
	}
	return nil
}

//TopicSetType alias
type TopicSetType struct { //wstop http://docs.oasis-open.org/wsn/t-1.xsd
	//ExtensibleDocumented

	//here can be any element
	RuleEngine *topic.RuleEngine `json:"tns:RuleEngine,omitempty" xml:",omitempty"`
}

//ExtensibleDocumented struct
type ExtensibleDocumented struct { //wstop http://docs.oasis-open.org/wsn/t-1.xsd
	Documentation Documentation //к xsd-документе documentation с маленькой буквы начинается
	//here can be anyAttribute
}

//ProducerReference Alias
type ProducerReference EndpointReferenceType

//SubscriptionReference Alias
type SubscriptionReference EndpointReferenceType

//NotificationMessageHolderType Alias
type NotificationMessageHolderType struct {
	SubscriptionReference SubscriptionReference //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
	Topic                 Topic
	ProducerReference     ProducerReference
	Message               MessageBody
}

type MessageBody struct {
	Message MessageDescription
}

type MessageDescription struct {
	PropertyOperation xsd.AnyType `xml:"PropertyOperation,attr"`
	Source            Source      `json:",omitempty" xml:",omitempty"`
	Data              Data        `json:",omitempty" xml:",omitempty"`
}

type Source struct {
	SimpleItem []SimpleItem `json:",omitempty" xml:",omitempty"`
}

type Data struct {
	SimpleItem []SimpleItem `json:",omitempty" xml:",omitempty"`
}

type SimpleItem struct {
	Name  xsd.AnyType `xml:"Name,attr"`
	Value xsd.AnyType `xml:"Value,attr"`
}

//NotificationMessage Alias
type NotificationMessage NotificationMessageHolderType //wsnt http://docs.oasis-open.org/wsn/b-2.xsd

//QueryExpressionType struct for wsnt:MessageContent
type QueryExpressionType struct { //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
	MessageKind xsd.String `xml:",chardata"` // boolean(ncex:Producer="15")
}

//MessageContentType Alias
type MessageContentType QueryExpressionType

//QueryExpression Alias
type QueryExpression QueryExpressionType

//TopicExpressionType struct for wsnt:TopicExpression
type TopicExpressionType struct { //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
	TopicKinds xsd.String `xml:",chardata"`
}

//Topic Alias
type Topic TopicExpressionType

// Capabilities of event
type Capabilities struct { //tev
	WSSubscriptionPolicySupport                   xsd.Boolean `xml:"WSSubscriptionPolicySupport,attr"`
	WSPullPointSupport                            xsd.Boolean `xml:"WSPullPointSupport,attr"`
	WSPausableSubscriptionManagerInterfaceSupport xsd.Boolean `xml:"WSPausableSubscriptionManagerInterfaceSupport,attr"`
	MaxNotificationProducers                      xsd.Int     `xml:"MaxNotificationProducers,attr"`
	MaxPullPoints                                 xsd.Int     `xml:"MaxPullPoints,attr"`
	PersistentNotificationStorage                 xsd.Boolean `xml:"PersistentNotificationStorage,attr"`
}

//ResourceUnknownFault response type
type ResourceUnknownFault struct {
}

//InvalidFilterFault response type
type InvalidFilterFault struct {
}

//TopicExpressionDialectUnknownFault response type
type TopicExpressionDialectUnknownFault struct {
}

//InvalidTopicExpressionFault response type
type InvalidTopicExpressionFault struct {
}

//TopicNotSupportedFault response type
type TopicNotSupportedFault struct {
}

//InvalidProducerPropertiesExpressionFault response type
type InvalidProducerPropertiesExpressionFault struct {
}

//InvalidMessageContentExpressionFault response type
type InvalidMessageContentExpressionFault struct {
}

//UnacceptableInitialTerminationTimeFault response type
type UnacceptableInitialTerminationTimeFault struct {
}

//UnrecognizedPolicyRequestFault response type
type UnrecognizedPolicyRequestFault struct {
}

//UnsupportedPolicyRequestFault response type
type UnsupportedPolicyRequestFault struct {
}

//NotifyMessageNotSupportedFault response type
type NotifyMessageNotSupportedFault struct {
}

//SubscribeCreationFailedFault response type
type SubscribeCreationFailedFault struct {
}
