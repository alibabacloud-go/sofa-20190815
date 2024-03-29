// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddMsConfigAttributesRequest struct {
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	Desc          *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceId    *int64  `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s AddMsConfigAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMsConfigAttributesRequest) GoString() string {
	return s.String()
}

func (s *AddMsConfigAttributesRequest) SetAttributeName(v string) *AddMsConfigAttributesRequest {
	s.AttributeName = &v
	return s
}

func (s *AddMsConfigAttributesRequest) SetDesc(v string) *AddMsConfigAttributesRequest {
	s.Desc = &v
	return s
}

func (s *AddMsConfigAttributesRequest) SetInstanceId(v string) *AddMsConfigAttributesRequest {
	s.InstanceId = &v
	return s
}

func (s *AddMsConfigAttributesRequest) SetResourceId(v int64) *AddMsConfigAttributesRequest {
	s.ResourceId = &v
	return s
}

type AddMsConfigAttributesResponseBody struct {
	Attribute     *AddMsConfigAttributesResponseBodyAttribute `json:"Attribute,omitempty" xml:"Attribute,omitempty" type:"Struct"`
	RequestId     *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                     `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                     `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s AddMsConfigAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddMsConfigAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *AddMsConfigAttributesResponseBody) SetAttribute(v *AddMsConfigAttributesResponseBodyAttribute) *AddMsConfigAttributesResponseBody {
	s.Attribute = v
	return s
}

func (s *AddMsConfigAttributesResponseBody) SetRequestId(v string) *AddMsConfigAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMsConfigAttributesResponseBody) SetResultCode(v string) *AddMsConfigAttributesResponseBody {
	s.ResultCode = &v
	return s
}

func (s *AddMsConfigAttributesResponseBody) SetResultMessage(v string) *AddMsConfigAttributesResponseBody {
	s.ResultMessage = &v
	return s
}

type AddMsConfigAttributesResponseBodyAttribute struct {
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	Desc          *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s AddMsConfigAttributesResponseBodyAttribute) String() string {
	return tea.Prettify(s)
}

func (s AddMsConfigAttributesResponseBodyAttribute) GoString() string {
	return s.String()
}

func (s *AddMsConfigAttributesResponseBodyAttribute) SetAttributeName(v string) *AddMsConfigAttributesResponseBodyAttribute {
	s.AttributeName = &v
	return s
}

func (s *AddMsConfigAttributesResponseBodyAttribute) SetDesc(v string) *AddMsConfigAttributesResponseBodyAttribute {
	s.Desc = &v
	return s
}

func (s *AddMsConfigAttributesResponseBodyAttribute) SetId(v int64) *AddMsConfigAttributesResponseBodyAttribute {
	s.Id = &v
	return s
}

func (s *AddMsConfigAttributesResponseBodyAttribute) SetInstanceId(v string) *AddMsConfigAttributesResponseBodyAttribute {
	s.InstanceId = &v
	return s
}

type AddMsConfigAttributesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddMsConfigAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddMsConfigAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMsConfigAttributesResponse) GoString() string {
	return s.String()
}

func (s *AddMsConfigAttributesResponse) SetHeaders(v map[string]*string) *AddMsConfigAttributesResponse {
	s.Headers = v
	return s
}

func (s *AddMsConfigAttributesResponse) SetStatusCode(v int32) *AddMsConfigAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMsConfigAttributesResponse) SetBody(v *AddMsConfigAttributesResponseBody) *AddMsConfigAttributesResponse {
	s.Body = v
	return s
}

type AddMsConfigResourcesRequest struct {
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	Desc       *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s AddMsConfigResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMsConfigResourcesRequest) GoString() string {
	return s.String()
}

func (s *AddMsConfigResourcesRequest) SetAppName(v string) *AddMsConfigResourcesRequest {
	s.AppName = &v
	return s
}

func (s *AddMsConfigResourcesRequest) SetAttributes(v string) *AddMsConfigResourcesRequest {
	s.Attributes = &v
	return s
}

func (s *AddMsConfigResourcesRequest) SetDesc(v string) *AddMsConfigResourcesRequest {
	s.Desc = &v
	return s
}

func (s *AddMsConfigResourcesRequest) SetInstanceId(v string) *AddMsConfigResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *AddMsConfigResourcesRequest) SetRegion(v string) *AddMsConfigResourcesRequest {
	s.Region = &v
	return s
}

func (s *AddMsConfigResourcesRequest) SetResourceId(v string) *AddMsConfigResourcesRequest {
	s.ResourceId = &v
	return s
}

type AddMsConfigResourcesResponseBody struct {
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resource      *AddMsConfigResourcesResponseBodyResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	ResultCode    *string                                   `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                   `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s AddMsConfigResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddMsConfigResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *AddMsConfigResourcesResponseBody) SetRequestId(v string) *AddMsConfigResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMsConfigResourcesResponseBody) SetResource(v *AddMsConfigResourcesResponseBodyResource) *AddMsConfigResourcesResponseBody {
	s.Resource = v
	return s
}

func (s *AddMsConfigResourcesResponseBody) SetResultCode(v string) *AddMsConfigResourcesResponseBody {
	s.ResultCode = &v
	return s
}

func (s *AddMsConfigResourcesResponseBody) SetResultMessage(v string) *AddMsConfigResourcesResponseBody {
	s.ResultMessage = &v
	return s
}

type AddMsConfigResourcesResponseBodyResource struct {
	AppName    *string                                               `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Attributes []*AddMsConfigResourcesResponseBodyResourceAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Desc       *string                                               `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Id         *int64                                                `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId *string                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region     *string                                               `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceId *string                                               `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s AddMsConfigResourcesResponseBodyResource) String() string {
	return tea.Prettify(s)
}

func (s AddMsConfigResourcesResponseBodyResource) GoString() string {
	return s.String()
}

func (s *AddMsConfigResourcesResponseBodyResource) SetAppName(v string) *AddMsConfigResourcesResponseBodyResource {
	s.AppName = &v
	return s
}

func (s *AddMsConfigResourcesResponseBodyResource) SetAttributes(v []*AddMsConfigResourcesResponseBodyResourceAttributes) *AddMsConfigResourcesResponseBodyResource {
	s.Attributes = v
	return s
}

func (s *AddMsConfigResourcesResponseBodyResource) SetDesc(v string) *AddMsConfigResourcesResponseBodyResource {
	s.Desc = &v
	return s
}

func (s *AddMsConfigResourcesResponseBodyResource) SetId(v int64) *AddMsConfigResourcesResponseBodyResource {
	s.Id = &v
	return s
}

func (s *AddMsConfigResourcesResponseBodyResource) SetInstanceId(v string) *AddMsConfigResourcesResponseBodyResource {
	s.InstanceId = &v
	return s
}

func (s *AddMsConfigResourcesResponseBodyResource) SetRegion(v string) *AddMsConfigResourcesResponseBodyResource {
	s.Region = &v
	return s
}

func (s *AddMsConfigResourcesResponseBodyResource) SetResourceId(v string) *AddMsConfigResourcesResponseBodyResource {
	s.ResourceId = &v
	return s
}

type AddMsConfigResourcesResponseBodyResourceAttributes struct {
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	Desc          *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s AddMsConfigResourcesResponseBodyResourceAttributes) String() string {
	return tea.Prettify(s)
}

func (s AddMsConfigResourcesResponseBodyResourceAttributes) GoString() string {
	return s.String()
}

func (s *AddMsConfigResourcesResponseBodyResourceAttributes) SetAttributeName(v string) *AddMsConfigResourcesResponseBodyResourceAttributes {
	s.AttributeName = &v
	return s
}

func (s *AddMsConfigResourcesResponseBodyResourceAttributes) SetDesc(v string) *AddMsConfigResourcesResponseBodyResourceAttributes {
	s.Desc = &v
	return s
}

func (s *AddMsConfigResourcesResponseBodyResourceAttributes) SetId(v int64) *AddMsConfigResourcesResponseBodyResourceAttributes {
	s.Id = &v
	return s
}

func (s *AddMsConfigResourcesResponseBodyResourceAttributes) SetInstanceId(v string) *AddMsConfigResourcesResponseBodyResourceAttributes {
	s.InstanceId = &v
	return s
}

type AddMsConfigResourcesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddMsConfigResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddMsConfigResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMsConfigResourcesResponse) GoString() string {
	return s.String()
}

func (s *AddMsConfigResourcesResponse) SetHeaders(v map[string]*string) *AddMsConfigResourcesResponse {
	s.Headers = v
	return s
}

func (s *AddMsConfigResourcesResponse) SetStatusCode(v int32) *AddMsConfigResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMsConfigResourcesResponse) SetBody(v *AddMsConfigResourcesResponseBody) *AddMsConfigResourcesResponse {
	s.Body = v
	return s
}

type CreateMqSofamqGroupRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupType  *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreateMqSofamqGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMqSofamqGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateMqSofamqGroupRequest) SetGroupId(v string) *CreateMqSofamqGroupRequest {
	s.GroupId = &v
	return s
}

func (s *CreateMqSofamqGroupRequest) SetGroupType(v string) *CreateMqSofamqGroupRequest {
	s.GroupType = &v
	return s
}

func (s *CreateMqSofamqGroupRequest) SetInstanceId(v string) *CreateMqSofamqGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateMqSofamqGroupRequest) SetRemark(v string) *CreateMqSofamqGroupRequest {
	s.Remark = &v
	return s
}

type CreateMqSofamqGroupResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMqSofamqGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMqSofamqGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMqSofamqGroupResponseBody) SetRequestId(v string) *CreateMqSofamqGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMqSofamqGroupResponseBody) SetResultCode(v string) *CreateMqSofamqGroupResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMqSofamqGroupResponseBody) SetResultMessage(v string) *CreateMqSofamqGroupResponseBody {
	s.ResultMessage = &v
	return s
}

type CreateMqSofamqGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMqSofamqGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMqSofamqGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMqSofamqGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateMqSofamqGroupResponse) SetHeaders(v map[string]*string) *CreateMqSofamqGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateMqSofamqGroupResponse) SetStatusCode(v int32) *CreateMqSofamqGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMqSofamqGroupResponse) SetBody(v *CreateMqSofamqGroupResponseBody) *CreateMqSofamqGroupResponse {
	s.Body = v
	return s
}

type CreateMqSofamqTopicRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MessageType *int64  `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	Remark      *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Topic       *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s CreateMqSofamqTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMqSofamqTopicRequest) GoString() string {
	return s.String()
}

func (s *CreateMqSofamqTopicRequest) SetInstanceId(v string) *CreateMqSofamqTopicRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateMqSofamqTopicRequest) SetMessageType(v int64) *CreateMqSofamqTopicRequest {
	s.MessageType = &v
	return s
}

func (s *CreateMqSofamqTopicRequest) SetRemark(v string) *CreateMqSofamqTopicRequest {
	s.Remark = &v
	return s
}

func (s *CreateMqSofamqTopicRequest) SetTopic(v string) *CreateMqSofamqTopicRequest {
	s.Topic = &v
	return s
}

type CreateMqSofamqTopicResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMqSofamqTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMqSofamqTopicResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMqSofamqTopicResponseBody) SetRequestId(v string) *CreateMqSofamqTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMqSofamqTopicResponseBody) SetResultCode(v string) *CreateMqSofamqTopicResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMqSofamqTopicResponseBody) SetResultMessage(v string) *CreateMqSofamqTopicResponseBody {
	s.ResultMessage = &v
	return s
}

type CreateMqSofamqTopicResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMqSofamqTopicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMqSofamqTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMqSofamqTopicResponse) GoString() string {
	return s.String()
}

func (s *CreateMqSofamqTopicResponse) SetHeaders(v map[string]*string) *CreateMqSofamqTopicResponse {
	s.Headers = v
	return s
}

func (s *CreateMqSofamqTopicResponse) SetStatusCode(v int32) *CreateMqSofamqTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMqSofamqTopicResponse) SetBody(v *CreateMqSofamqTopicResponseBody) *CreateMqSofamqTopicResponse {
	s.Body = v
	return s
}

type CreateRMSUnifiedAlarmRuleRequest struct {
	AlarmNodata             *int64                                          `json:"AlarmNodata,omitempty" xml:"AlarmNodata,omitempty"`
	Category                *string                                         `json:"Category,omitempty" xml:"Category,omitempty"`
	ChannelsRepeatList      []*string                                       `json:"ChannelsRepeatList,omitempty" xml:"ChannelsRepeatList,omitempty" type:"Repeated"`
	Emergency               *string                                         `json:"Emergency,omitempty" xml:"Emergency,omitempty"`
	EmergencyUrl            *string                                         `json:"EmergencyUrl,omitempty" xml:"EmergencyUrl,omitempty"`
	Level                   *int64                                          `json:"Level,omitempty" xml:"Level,omitempty"`
	Name                    *string                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	NotifyFiring            *int64                                          `json:"NotifyFiring,omitempty" xml:"NotifyFiring,omitempty"`
	NotifyNodata            *int64                                          `json:"NotifyNodata,omitempty" xml:"NotifyNodata,omitempty"`
	NotifyRecovered         *int64                                          `json:"NotifyRecovered,omitempty" xml:"NotifyRecovered,omitempty"`
	NotifyTarget            []*CreateRMSUnifiedAlarmRuleRequestNotifyTarget `json:"NotifyTarget,omitempty" xml:"NotifyTarget,omitempty" type:"Repeated"`
	NotifyTimeFilterJsonStr *string                                         `json:"NotifyTimeFilterJsonStr,omitempty" xml:"NotifyTimeFilterJsonStr,omitempty"`
	PendingHit              *int64                                          `json:"PendingHit,omitempty" xml:"PendingHit,omitempty"`
	RecoveredHit            *int64                                          `json:"RecoveredHit,omitempty" xml:"RecoveredHit,omitempty"`
	RuleConfig              *string                                         `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	SilenceTime             *int64                                          `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	Step                    *int64                                          `json:"Step,omitempty" xml:"Step,omitempty"`
	WorkspaceName           *string                                         `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s CreateRMSUnifiedAlarmRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRMSUnifiedAlarmRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRMSUnifiedAlarmRuleRequest) SetAlarmNodata(v int64) *CreateRMSUnifiedAlarmRuleRequest {
	s.AlarmNodata = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequest) SetCategory(v string) *CreateRMSUnifiedAlarmRuleRequest {
	s.Category = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequest) SetChannelsRepeatList(v []*string) *CreateRMSUnifiedAlarmRuleRequest {
	s.ChannelsRepeatList = v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequest) SetEmergency(v string) *CreateRMSUnifiedAlarmRuleRequest {
	s.Emergency = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequest) SetEmergencyUrl(v string) *CreateRMSUnifiedAlarmRuleRequest {
	s.EmergencyUrl = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequest) SetLevel(v int64) *CreateRMSUnifiedAlarmRuleRequest {
	s.Level = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequest) SetName(v string) *CreateRMSUnifiedAlarmRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequest) SetNotifyFiring(v int64) *CreateRMSUnifiedAlarmRuleRequest {
	s.NotifyFiring = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequest) SetNotifyNodata(v int64) *CreateRMSUnifiedAlarmRuleRequest {
	s.NotifyNodata = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequest) SetNotifyRecovered(v int64) *CreateRMSUnifiedAlarmRuleRequest {
	s.NotifyRecovered = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequest) SetNotifyTarget(v []*CreateRMSUnifiedAlarmRuleRequestNotifyTarget) *CreateRMSUnifiedAlarmRuleRequest {
	s.NotifyTarget = v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequest) SetNotifyTimeFilterJsonStr(v string) *CreateRMSUnifiedAlarmRuleRequest {
	s.NotifyTimeFilterJsonStr = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequest) SetPendingHit(v int64) *CreateRMSUnifiedAlarmRuleRequest {
	s.PendingHit = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequest) SetRecoveredHit(v int64) *CreateRMSUnifiedAlarmRuleRequest {
	s.RecoveredHit = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequest) SetRuleConfig(v string) *CreateRMSUnifiedAlarmRuleRequest {
	s.RuleConfig = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequest) SetSilenceTime(v int64) *CreateRMSUnifiedAlarmRuleRequest {
	s.SilenceTime = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequest) SetStep(v int64) *CreateRMSUnifiedAlarmRuleRequest {
	s.Step = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequest) SetWorkspaceName(v string) *CreateRMSUnifiedAlarmRuleRequest {
	s.WorkspaceName = &v
	return s
}

type CreateRMSUnifiedAlarmRuleRequestNotifyTarget struct {
	Subscriber       *string `json:"Subscriber,omitempty" xml:"Subscriber,omitempty"`
	SubscriberName   *string `json:"SubscriberName,omitempty" xml:"SubscriberName,omitempty"`
	SubscriberSource *string `json:"SubscriberSource,omitempty" xml:"SubscriberSource,omitempty"`
	SubscriberType   *string `json:"SubscriberType,omitempty" xml:"SubscriberType,omitempty"`
	SubscriberUuid   *string `json:"SubscriberUuid,omitempty" xml:"SubscriberUuid,omitempty"`
}

func (s CreateRMSUnifiedAlarmRuleRequestNotifyTarget) String() string {
	return tea.Prettify(s)
}

func (s CreateRMSUnifiedAlarmRuleRequestNotifyTarget) GoString() string {
	return s.String()
}

func (s *CreateRMSUnifiedAlarmRuleRequestNotifyTarget) SetSubscriber(v string) *CreateRMSUnifiedAlarmRuleRequestNotifyTarget {
	s.Subscriber = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequestNotifyTarget) SetSubscriberName(v string) *CreateRMSUnifiedAlarmRuleRequestNotifyTarget {
	s.SubscriberName = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequestNotifyTarget) SetSubscriberSource(v string) *CreateRMSUnifiedAlarmRuleRequestNotifyTarget {
	s.SubscriberSource = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequestNotifyTarget) SetSubscriberType(v string) *CreateRMSUnifiedAlarmRuleRequestNotifyTarget {
	s.SubscriberType = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleRequestNotifyTarget) SetSubscriberUuid(v string) *CreateRMSUnifiedAlarmRuleRequestNotifyTarget {
	s.SubscriberUuid = &v
	return s
}

type CreateRMSUnifiedAlarmRuleResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateRMSUnifiedAlarmRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRMSUnifiedAlarmRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRMSUnifiedAlarmRuleResponseBody) SetRequestId(v string) *CreateRMSUnifiedAlarmRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleResponseBody) SetResultCode(v string) *CreateRMSUnifiedAlarmRuleResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleResponseBody) SetResultMessage(v string) *CreateRMSUnifiedAlarmRuleResponseBody {
	s.ResultMessage = &v
	return s
}

type CreateRMSUnifiedAlarmRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRMSUnifiedAlarmRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRMSUnifiedAlarmRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRMSUnifiedAlarmRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRMSUnifiedAlarmRuleResponse) SetHeaders(v map[string]*string) *CreateRMSUnifiedAlarmRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleResponse) SetStatusCode(v int32) *CreateRMSUnifiedAlarmRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRMSUnifiedAlarmRuleResponse) SetBody(v *CreateRMSUnifiedAlarmRuleResponseBody) *CreateRMSUnifiedAlarmRuleResponse {
	s.Body = v
	return s
}

type DeleteMqSofamqGroupRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteMqSofamqGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMqSofamqGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteMqSofamqGroupRequest) SetGroupId(v string) *DeleteMqSofamqGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteMqSofamqGroupRequest) SetInstanceId(v string) *DeleteMqSofamqGroupRequest {
	s.InstanceId = &v
	return s
}

type DeleteMqSofamqGroupResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s DeleteMqSofamqGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMqSofamqGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMqSofamqGroupResponseBody) SetRequestId(v string) *DeleteMqSofamqGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMqSofamqGroupResponseBody) SetResultCode(v string) *DeleteMqSofamqGroupResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DeleteMqSofamqGroupResponseBody) SetResultMessage(v string) *DeleteMqSofamqGroupResponseBody {
	s.ResultMessage = &v
	return s
}

type DeleteMqSofamqGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteMqSofamqGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMqSofamqGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMqSofamqGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteMqSofamqGroupResponse) SetHeaders(v map[string]*string) *DeleteMqSofamqGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteMqSofamqGroupResponse) SetStatusCode(v int32) *DeleteMqSofamqGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMqSofamqGroupResponse) SetBody(v *DeleteMqSofamqGroupResponseBody) *DeleteMqSofamqGroupResponse {
	s.Body = v
	return s
}

type DeleteMqSofamqTopicRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s DeleteMqSofamqTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMqSofamqTopicRequest) GoString() string {
	return s.String()
}

func (s *DeleteMqSofamqTopicRequest) SetInstanceId(v string) *DeleteMqSofamqTopicRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteMqSofamqTopicRequest) SetTopic(v string) *DeleteMqSofamqTopicRequest {
	s.Topic = &v
	return s
}

type DeleteMqSofamqTopicResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s DeleteMqSofamqTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMqSofamqTopicResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMqSofamqTopicResponseBody) SetRequestId(v string) *DeleteMqSofamqTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMqSofamqTopicResponseBody) SetResultCode(v string) *DeleteMqSofamqTopicResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DeleteMqSofamqTopicResponseBody) SetResultMessage(v string) *DeleteMqSofamqTopicResponseBody {
	s.ResultMessage = &v
	return s
}

type DeleteMqSofamqTopicResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteMqSofamqTopicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMqSofamqTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMqSofamqTopicResponse) GoString() string {
	return s.String()
}

func (s *DeleteMqSofamqTopicResponse) SetHeaders(v map[string]*string) *DeleteMqSofamqTopicResponse {
	s.Headers = v
	return s
}

func (s *DeleteMqSofamqTopicResponse) SetStatusCode(v int32) *DeleteMqSofamqTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMqSofamqTopicResponse) SetBody(v *DeleteMqSofamqTopicResponseBody) *DeleteMqSofamqTopicResponse {
	s.Body = v
	return s
}

type DeleteMqSofamqTraceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QueryId    *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s DeleteMqSofamqTraceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMqSofamqTraceRequest) GoString() string {
	return s.String()
}

func (s *DeleteMqSofamqTraceRequest) SetInstanceId(v string) *DeleteMqSofamqTraceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteMqSofamqTraceRequest) SetQueryId(v string) *DeleteMqSofamqTraceRequest {
	s.QueryId = &v
	return s
}

type DeleteMqSofamqTraceResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s DeleteMqSofamqTraceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMqSofamqTraceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMqSofamqTraceResponseBody) SetRequestId(v string) *DeleteMqSofamqTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMqSofamqTraceResponseBody) SetResultCode(v string) *DeleteMqSofamqTraceResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DeleteMqSofamqTraceResponseBody) SetResultMessage(v string) *DeleteMqSofamqTraceResponseBody {
	s.ResultMessage = &v
	return s
}

type DeleteMqSofamqTraceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteMqSofamqTraceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMqSofamqTraceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMqSofamqTraceResponse) GoString() string {
	return s.String()
}

func (s *DeleteMqSofamqTraceResponse) SetHeaders(v map[string]*string) *DeleteMqSofamqTraceResponse {
	s.Headers = v
	return s
}

func (s *DeleteMqSofamqTraceResponse) SetStatusCode(v int32) *DeleteMqSofamqTraceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMqSofamqTraceResponse) SetBody(v *DeleteMqSofamqTraceResponseBody) *DeleteMqSofamqTraceResponse {
	s.Body = v
	return s
}

type DeleteMqSofamqWarnRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	WarnId     *int64  `json:"WarnId,omitempty" xml:"WarnId,omitempty"`
}

func (s DeleteMqSofamqWarnRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMqSofamqWarnRequest) GoString() string {
	return s.String()
}

func (s *DeleteMqSofamqWarnRequest) SetInstanceId(v string) *DeleteMqSofamqWarnRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteMqSofamqWarnRequest) SetWarnId(v int64) *DeleteMqSofamqWarnRequest {
	s.WarnId = &v
	return s
}

type DeleteMqSofamqWarnResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s DeleteMqSofamqWarnResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMqSofamqWarnResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMqSofamqWarnResponseBody) SetRequestId(v string) *DeleteMqSofamqWarnResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMqSofamqWarnResponseBody) SetResultCode(v string) *DeleteMqSofamqWarnResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DeleteMqSofamqWarnResponseBody) SetResultMessage(v string) *DeleteMqSofamqWarnResponseBody {
	s.ResultMessage = &v
	return s
}

type DeleteMqSofamqWarnResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteMqSofamqWarnResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMqSofamqWarnResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMqSofamqWarnResponse) GoString() string {
	return s.String()
}

func (s *DeleteMqSofamqWarnResponse) SetHeaders(v map[string]*string) *DeleteMqSofamqWarnResponse {
	s.Headers = v
	return s
}

func (s *DeleteMqSofamqWarnResponse) SetStatusCode(v int32) *DeleteMqSofamqWarnResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMqSofamqWarnResponse) SetBody(v *DeleteMqSofamqWarnResponseBody) *DeleteMqSofamqWarnResponse {
	s.Body = v
	return s
}

type DeleteMsConfigAttributesRequest struct {
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteMsConfigAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMsConfigAttributesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMsConfigAttributesRequest) SetId(v int64) *DeleteMsConfigAttributesRequest {
	s.Id = &v
	return s
}

func (s *DeleteMsConfigAttributesRequest) SetInstanceId(v string) *DeleteMsConfigAttributesRequest {
	s.InstanceId = &v
	return s
}

type DeleteMsConfigAttributesResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s DeleteMsConfigAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMsConfigAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMsConfigAttributesResponseBody) SetRequestId(v string) *DeleteMsConfigAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMsConfigAttributesResponseBody) SetResultCode(v string) *DeleteMsConfigAttributesResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DeleteMsConfigAttributesResponseBody) SetResultMessage(v string) *DeleteMsConfigAttributesResponseBody {
	s.ResultMessage = &v
	return s
}

type DeleteMsConfigAttributesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteMsConfigAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMsConfigAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMsConfigAttributesResponse) GoString() string {
	return s.String()
}

func (s *DeleteMsConfigAttributesResponse) SetHeaders(v map[string]*string) *DeleteMsConfigAttributesResponse {
	s.Headers = v
	return s
}

func (s *DeleteMsConfigAttributesResponse) SetStatusCode(v int32) *DeleteMsConfigAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMsConfigAttributesResponse) SetBody(v *DeleteMsConfigAttributesResponseBody) *DeleteMsConfigAttributesResponse {
	s.Body = v
	return s
}

type DeleteMsConfigResourcesRequest struct {
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteMsConfigResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMsConfigResourcesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMsConfigResourcesRequest) SetId(v int64) *DeleteMsConfigResourcesRequest {
	s.Id = &v
	return s
}

func (s *DeleteMsConfigResourcesRequest) SetInstanceId(v string) *DeleteMsConfigResourcesRequest {
	s.InstanceId = &v
	return s
}

type DeleteMsConfigResourcesResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s DeleteMsConfigResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMsConfigResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMsConfigResourcesResponseBody) SetRequestId(v string) *DeleteMsConfigResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMsConfigResourcesResponseBody) SetResultCode(v string) *DeleteMsConfigResourcesResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DeleteMsConfigResourcesResponseBody) SetResultMessage(v string) *DeleteMsConfigResourcesResponseBody {
	s.ResultMessage = &v
	return s
}

type DeleteMsConfigResourcesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteMsConfigResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMsConfigResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMsConfigResourcesResponse) GoString() string {
	return s.String()
}

func (s *DeleteMsConfigResourcesResponse) SetHeaders(v map[string]*string) *DeleteMsConfigResourcesResponse {
	s.Headers = v
	return s
}

func (s *DeleteMsConfigResourcesResponse) SetStatusCode(v int32) *DeleteMsConfigResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMsConfigResourcesResponse) SetBody(v *DeleteMsConfigResourcesResponseBody) *DeleteMsConfigResourcesResponse {
	s.Body = v
	return s
}

type DeleteRMSUnifiedAlarmRuleRequest struct {
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s DeleteRMSUnifiedAlarmRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRMSUnifiedAlarmRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRMSUnifiedAlarmRuleRequest) SetId(v int64) *DeleteRMSUnifiedAlarmRuleRequest {
	s.Id = &v
	return s
}

func (s *DeleteRMSUnifiedAlarmRuleRequest) SetWorkspaceName(v string) *DeleteRMSUnifiedAlarmRuleRequest {
	s.WorkspaceName = &v
	return s
}

type DeleteRMSUnifiedAlarmRuleResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s DeleteRMSUnifiedAlarmRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRMSUnifiedAlarmRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRMSUnifiedAlarmRuleResponseBody) SetRequestId(v string) *DeleteRMSUnifiedAlarmRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRMSUnifiedAlarmRuleResponseBody) SetResultCode(v string) *DeleteRMSUnifiedAlarmRuleResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DeleteRMSUnifiedAlarmRuleResponseBody) SetResultMessage(v string) *DeleteRMSUnifiedAlarmRuleResponseBody {
	s.ResultMessage = &v
	return s
}

type DeleteRMSUnifiedAlarmRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRMSUnifiedAlarmRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRMSUnifiedAlarmRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRMSUnifiedAlarmRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRMSUnifiedAlarmRuleResponse) SetHeaders(v map[string]*string) *DeleteRMSUnifiedAlarmRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRMSUnifiedAlarmRuleResponse) SetStatusCode(v int32) *DeleteRMSUnifiedAlarmRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRMSUnifiedAlarmRuleResponse) SetBody(v *DeleteRMSUnifiedAlarmRuleResponseBody) *DeleteRMSUnifiedAlarmRuleResponse {
	s.Body = v
	return s
}

type DescribeCasComputersRequest struct {
	AppServiceIdsRepeatList []*string `json:"AppServiceIdsRepeatList,omitempty" xml:"AppServiceIdsRepeatList,omitempty" type:"Repeated"`
	CurrentPage             *int64    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Name                    *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	PageSize                *int64    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Workspace               *string   `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s DescribeCasComputersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasComputersRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasComputersRequest) SetAppServiceIdsRepeatList(v []*string) *DescribeCasComputersRequest {
	s.AppServiceIdsRepeatList = v
	return s
}

func (s *DescribeCasComputersRequest) SetCurrentPage(v int64) *DescribeCasComputersRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCasComputersRequest) SetName(v string) *DescribeCasComputersRequest {
	s.Name = &v
	return s
}

func (s *DescribeCasComputersRequest) SetPageSize(v int64) *DescribeCasComputersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCasComputersRequest) SetWorkspace(v string) *DescribeCasComputersRequest {
	s.Workspace = &v
	return s
}

type DescribeCasComputersResponseBody struct {
	CurrentPage   *int64                                  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data          []*DescribeCasComputersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageSize      *int64                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                 `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	TotalCount    *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCasComputersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasComputersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCasComputersResponseBody) SetCurrentPage(v int64) *DescribeCasComputersResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCasComputersResponseBody) SetData(v []*DescribeCasComputersResponseBodyData) *DescribeCasComputersResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCasComputersResponseBody) SetPageSize(v int64) *DescribeCasComputersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCasComputersResponseBody) SetRequestId(v string) *DescribeCasComputersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCasComputersResponseBody) SetResultCode(v string) *DescribeCasComputersResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DescribeCasComputersResponseBody) SetResultMessage(v string) *DescribeCasComputersResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DescribeCasComputersResponseBody) SetTotalCount(v int64) *DescribeCasComputersResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCasComputersResponseBodyData struct {
	AppId                 *string                                          `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppServiceIds         []*string                                        `json:"AppServiceIds,omitempty" xml:"AppServiceIds,omitempty" type:"Repeated"`
	AssignedAppServiceIds []*string                                        `json:"AssignedAppServiceIds,omitempty" xml:"AssignedAppServiceIds,omitempty" type:"Repeated"`
	AutoRenew             *bool                                            `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod       *int64                                           `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	Bandwidth             *int64                                           `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	CommonImage           *bool                                            `json:"CommonImage,omitempty" xml:"CommonImage,omitempty"`
	Cpu                   *int64                                           `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CpuShared             *bool                                            `json:"CpuShared,omitempty" xml:"CpuShared,omitempty"`
	CreationTime          *string                                          `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DataDisks             []*DescribeCasComputersResponseBodyDataDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	DeployMode            *string                                          `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	Description           *string                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	ElasticIp             *string                                          `json:"ElasticIp,omitempty" xml:"ElasticIp,omitempty"`
	ExpiredTime           *string                                          `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	IaasId                *string                                          `json:"IaasId,omitempty" xml:"IaasId,omitempty"`
	IaasStatus            *string                                          `json:"IaasStatus,omitempty" xml:"IaasStatus,omitempty"`
	IaasType              *string                                          `json:"IaasType,omitempty" xml:"IaasType,omitempty"`
	Id                    *string                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageIaasId           *string                                          `json:"ImageIaasId,omitempty" xml:"ImageIaasId,omitempty"`
	ImageId               *string                                          `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName             *string                                          `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	Initialized           *bool                                            `json:"Initialized,omitempty" xml:"Initialized,omitempty"`
	InstanceChargeType    *string                                          `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	IoOptimized           *bool                                            `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	LastOpsOrderId        *string                                          `json:"LastOpsOrderId,omitempty" xml:"LastOpsOrderId,omitempty"`
	LastOpsType           *string                                          `json:"LastOpsType,omitempty" xml:"LastOpsType,omitempty"`
	Memory                *int64                                           `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Name                  *string                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	NetworkType           *string                                          `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Options               []*DescribeCasComputersResponseBodyDataOptions   `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
	Os                    *string                                          `json:"Os,omitempty" xml:"Os,omitempty"`
	OsBit                 *int64                                           `json:"OsBit,omitempty" xml:"OsBit,omitempty"`
	OsVersion             *string                                          `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	PaasStatus            *string                                          `json:"PaasStatus,omitempty" xml:"PaasStatus,omitempty"`
	Password              *string                                          `json:"Password,omitempty" xml:"Password,omitempty"`
	PrivateIp             *string                                          `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	ProviderId            *string                                          `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	PublicIp              *string                                          `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	RegionId              *string                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SerialNumber          *string                                          `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	SpecIaasId            *string                                          `json:"SpecIaasId,omitempty" xml:"SpecIaasId,omitempty"`
	Status                *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	SystemDisk            *DescribeCasComputersResponseBodyDataSystemDisk  `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	Tags                  []*DescribeCasComputersResponseBodyDataTags      `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TenantId              *string                                          `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TenantId2             *string                                          `json:"TenantId2,omitempty" xml:"TenantId2,omitempty"`
	ThreadsPerCore        *int64                                           `json:"ThreadsPerCore,omitempty" xml:"ThreadsPerCore,omitempty"`
	UtcCreate             *string                                          `json:"UtcCreate,omitempty" xml:"UtcCreate,omitempty"`
	UtcModified           *string                                          `json:"UtcModified,omitempty" xml:"UtcModified,omitempty"`
	VSwitchIaasId         *string                                          `json:"VSwitchIaasId,omitempty" xml:"VSwitchIaasId,omitempty"`
	VpcId                 *string                                          `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	WorkspaceId           *string                                          `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	ZoneId                *string                                          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeCasComputersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasComputersResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCasComputersResponseBodyData) SetAppId(v string) *DescribeCasComputersResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetAppServiceIds(v []*string) *DescribeCasComputersResponseBodyData {
	s.AppServiceIds = v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetAssignedAppServiceIds(v []*string) *DescribeCasComputersResponseBodyData {
	s.AssignedAppServiceIds = v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetAutoRenew(v bool) *DescribeCasComputersResponseBodyData {
	s.AutoRenew = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetAutoRenewPeriod(v int64) *DescribeCasComputersResponseBodyData {
	s.AutoRenewPeriod = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetBandwidth(v int64) *DescribeCasComputersResponseBodyData {
	s.Bandwidth = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetCommonImage(v bool) *DescribeCasComputersResponseBodyData {
	s.CommonImage = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetCpu(v int64) *DescribeCasComputersResponseBodyData {
	s.Cpu = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetCpuShared(v bool) *DescribeCasComputersResponseBodyData {
	s.CpuShared = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetCreationTime(v string) *DescribeCasComputersResponseBodyData {
	s.CreationTime = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetDataDisks(v []*DescribeCasComputersResponseBodyDataDataDisks) *DescribeCasComputersResponseBodyData {
	s.DataDisks = v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetDeployMode(v string) *DescribeCasComputersResponseBodyData {
	s.DeployMode = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetDescription(v string) *DescribeCasComputersResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetElasticIp(v string) *DescribeCasComputersResponseBodyData {
	s.ElasticIp = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetExpiredTime(v string) *DescribeCasComputersResponseBodyData {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetIaasId(v string) *DescribeCasComputersResponseBodyData {
	s.IaasId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetIaasStatus(v string) *DescribeCasComputersResponseBodyData {
	s.IaasStatus = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetIaasType(v string) *DescribeCasComputersResponseBodyData {
	s.IaasType = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetId(v string) *DescribeCasComputersResponseBodyData {
	s.Id = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetImageIaasId(v string) *DescribeCasComputersResponseBodyData {
	s.ImageIaasId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetImageId(v string) *DescribeCasComputersResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetImageName(v string) *DescribeCasComputersResponseBodyData {
	s.ImageName = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetInitialized(v bool) *DescribeCasComputersResponseBodyData {
	s.Initialized = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetInstanceChargeType(v string) *DescribeCasComputersResponseBodyData {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetIoOptimized(v bool) *DescribeCasComputersResponseBodyData {
	s.IoOptimized = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetLastOpsOrderId(v string) *DescribeCasComputersResponseBodyData {
	s.LastOpsOrderId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetLastOpsType(v string) *DescribeCasComputersResponseBodyData {
	s.LastOpsType = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetMemory(v int64) *DescribeCasComputersResponseBodyData {
	s.Memory = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetName(v string) *DescribeCasComputersResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetNetworkType(v string) *DescribeCasComputersResponseBodyData {
	s.NetworkType = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetOptions(v []*DescribeCasComputersResponseBodyDataOptions) *DescribeCasComputersResponseBodyData {
	s.Options = v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetOs(v string) *DescribeCasComputersResponseBodyData {
	s.Os = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetOsBit(v int64) *DescribeCasComputersResponseBodyData {
	s.OsBit = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetOsVersion(v string) *DescribeCasComputersResponseBodyData {
	s.OsVersion = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetPaasStatus(v string) *DescribeCasComputersResponseBodyData {
	s.PaasStatus = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetPassword(v string) *DescribeCasComputersResponseBodyData {
	s.Password = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetPrivateIp(v string) *DescribeCasComputersResponseBodyData {
	s.PrivateIp = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetProviderId(v string) *DescribeCasComputersResponseBodyData {
	s.ProviderId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetPublicIp(v string) *DescribeCasComputersResponseBodyData {
	s.PublicIp = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetRegionId(v string) *DescribeCasComputersResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetSerialNumber(v string) *DescribeCasComputersResponseBodyData {
	s.SerialNumber = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetSpecIaasId(v string) *DescribeCasComputersResponseBodyData {
	s.SpecIaasId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetStatus(v string) *DescribeCasComputersResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetSystemDisk(v *DescribeCasComputersResponseBodyDataSystemDisk) *DescribeCasComputersResponseBodyData {
	s.SystemDisk = v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetTags(v []*DescribeCasComputersResponseBodyDataTags) *DescribeCasComputersResponseBodyData {
	s.Tags = v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetTenantId(v string) *DescribeCasComputersResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetTenantId2(v string) *DescribeCasComputersResponseBodyData {
	s.TenantId2 = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetThreadsPerCore(v int64) *DescribeCasComputersResponseBodyData {
	s.ThreadsPerCore = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetUtcCreate(v string) *DescribeCasComputersResponseBodyData {
	s.UtcCreate = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetUtcModified(v string) *DescribeCasComputersResponseBodyData {
	s.UtcModified = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetVSwitchIaasId(v string) *DescribeCasComputersResponseBodyData {
	s.VSwitchIaasId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetVpcId(v string) *DescribeCasComputersResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetWorkspaceId(v string) *DescribeCasComputersResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyData) SetZoneId(v string) *DescribeCasComputersResponseBodyData {
	s.ZoneId = &v
	return s
}

type DescribeCasComputersResponseBodyDataDataDisks struct {
	Category           *string                                                `json:"Category,omitempty" xml:"Category,omitempty"`
	Computer           *DescribeCasComputersResponseBodyDataDataDisksComputer `json:"Computer,omitempty" xml:"Computer,omitempty" type:"Struct"`
	DeleteAutoSnapshot *bool                                                  `json:"DeleteAutoSnapshot,omitempty" xml:"DeleteAutoSnapshot,omitempty"`
	DeleteWithComputer *bool                                                  `json:"DeleteWithComputer,omitempty" xml:"DeleteWithComputer,omitempty"`
	Device             *string                                                `json:"Device,omitempty" xml:"Device,omitempty"`
	EnableAutoSnapshot *bool                                                  `json:"EnableAutoSnapshot,omitempty" xml:"EnableAutoSnapshot,omitempty"`
	IaasId             *string                                                `json:"IaasId,omitempty" xml:"IaasId,omitempty"`
	Id                 *string                                                `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageId            *string                                                `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Name               *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	Portable           *bool                                                  `json:"Portable,omitempty" xml:"Portable,omitempty"`
	ProviderId         *string                                                `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	RegionId           *string                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Size               *int64                                                 `json:"Size,omitempty" xml:"Size,omitempty"`
	Status             *string                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId           *string                                                `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Type               *string                                                `json:"Type,omitempty" xml:"Type,omitempty"`
	UtcCreate          *string                                                `json:"UtcCreate,omitempty" xml:"UtcCreate,omitempty"`
	UtcModified        *string                                                `json:"UtcModified,omitempty" xml:"UtcModified,omitempty"`
	WorkspaceId        *string                                                `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	ZoneId             *string                                                `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeCasComputersResponseBodyDataDataDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasComputersResponseBodyDataDataDisks) GoString() string {
	return s.String()
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetCategory(v string) *DescribeCasComputersResponseBodyDataDataDisks {
	s.Category = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetComputer(v *DescribeCasComputersResponseBodyDataDataDisksComputer) *DescribeCasComputersResponseBodyDataDataDisks {
	s.Computer = v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetDeleteAutoSnapshot(v bool) *DescribeCasComputersResponseBodyDataDataDisks {
	s.DeleteAutoSnapshot = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetDeleteWithComputer(v bool) *DescribeCasComputersResponseBodyDataDataDisks {
	s.DeleteWithComputer = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetDevice(v string) *DescribeCasComputersResponseBodyDataDataDisks {
	s.Device = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetEnableAutoSnapshot(v bool) *DescribeCasComputersResponseBodyDataDataDisks {
	s.EnableAutoSnapshot = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetIaasId(v string) *DescribeCasComputersResponseBodyDataDataDisks {
	s.IaasId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetId(v string) *DescribeCasComputersResponseBodyDataDataDisks {
	s.Id = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetImageId(v string) *DescribeCasComputersResponseBodyDataDataDisks {
	s.ImageId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetName(v string) *DescribeCasComputersResponseBodyDataDataDisks {
	s.Name = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetPortable(v bool) *DescribeCasComputersResponseBodyDataDataDisks {
	s.Portable = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetProviderId(v string) *DescribeCasComputersResponseBodyDataDataDisks {
	s.ProviderId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetRegionId(v string) *DescribeCasComputersResponseBodyDataDataDisks {
	s.RegionId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetSize(v int64) *DescribeCasComputersResponseBodyDataDataDisks {
	s.Size = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetStatus(v string) *DescribeCasComputersResponseBodyDataDataDisks {
	s.Status = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetTenantId(v string) *DescribeCasComputersResponseBodyDataDataDisks {
	s.TenantId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetType(v string) *DescribeCasComputersResponseBodyDataDataDisks {
	s.Type = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetUtcCreate(v string) *DescribeCasComputersResponseBodyDataDataDisks {
	s.UtcCreate = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetUtcModified(v string) *DescribeCasComputersResponseBodyDataDataDisks {
	s.UtcModified = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetWorkspaceId(v string) *DescribeCasComputersResponseBodyDataDataDisks {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisks) SetZoneId(v string) *DescribeCasComputersResponseBodyDataDataDisks {
	s.ZoneId = &v
	return s
}

type DescribeCasComputersResponseBodyDataDataDisksComputer struct {
	IaasId *string `json:"IaasId,omitempty" xml:"IaasId,omitempty"`
	Id     *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCasComputersResponseBodyDataDataDisksComputer) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasComputersResponseBodyDataDataDisksComputer) GoString() string {
	return s.String()
}

func (s *DescribeCasComputersResponseBodyDataDataDisksComputer) SetIaasId(v string) *DescribeCasComputersResponseBodyDataDataDisksComputer {
	s.IaasId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisksComputer) SetId(v string) *DescribeCasComputersResponseBodyDataDataDisksComputer {
	s.Id = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisksComputer) SetName(v string) *DescribeCasComputersResponseBodyDataDataDisksComputer {
	s.Name = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataDataDisksComputer) SetStatus(v string) *DescribeCasComputersResponseBodyDataDataDisksComputer {
	s.Status = &v
	return s
}

type DescribeCasComputersResponseBodyDataOptions struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCasComputersResponseBodyDataOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasComputersResponseBodyDataOptions) GoString() string {
	return s.String()
}

func (s *DescribeCasComputersResponseBodyDataOptions) SetKey(v string) *DescribeCasComputersResponseBodyDataOptions {
	s.Key = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataOptions) SetValue(v string) *DescribeCasComputersResponseBodyDataOptions {
	s.Value = &v
	return s
}

type DescribeCasComputersResponseBodyDataSystemDisk struct {
	Category           *string                                                 `json:"Category,omitempty" xml:"Category,omitempty"`
	Computer           *DescribeCasComputersResponseBodyDataSystemDiskComputer `json:"Computer,omitempty" xml:"Computer,omitempty" type:"Struct"`
	DeleteAutoSnapshot *bool                                                   `json:"DeleteAutoSnapshot,omitempty" xml:"DeleteAutoSnapshot,omitempty"`
	DeleteWithComputer *bool                                                   `json:"DeleteWithComputer,omitempty" xml:"DeleteWithComputer,omitempty"`
	Device             *string                                                 `json:"Device,omitempty" xml:"Device,omitempty"`
	EnableAutoSnapshot *bool                                                   `json:"EnableAutoSnapshot,omitempty" xml:"EnableAutoSnapshot,omitempty"`
	IaasId             *string                                                 `json:"IaasId,omitempty" xml:"IaasId,omitempty"`
	Id                 *string                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageId            *string                                                 `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Name               *string                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	Portable           *bool                                                   `json:"Portable,omitempty" xml:"Portable,omitempty"`
	ProviderId         *string                                                 `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	RegionId           *string                                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Size               *int64                                                  `json:"Size,omitempty" xml:"Size,omitempty"`
	Status             *string                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId           *string                                                 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Type               *string                                                 `json:"Type,omitempty" xml:"Type,omitempty"`
	UtcCreate          *string                                                 `json:"UtcCreate,omitempty" xml:"UtcCreate,omitempty"`
	UtcModified        *string                                                 `json:"UtcModified,omitempty" xml:"UtcModified,omitempty"`
	WorkspaceId        *string                                                 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	ZoneId             *string                                                 `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeCasComputersResponseBodyDataSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasComputersResponseBodyDataSystemDisk) GoString() string {
	return s.String()
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetCategory(v string) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.Category = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetComputer(v *DescribeCasComputersResponseBodyDataSystemDiskComputer) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.Computer = v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetDeleteAutoSnapshot(v bool) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.DeleteAutoSnapshot = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetDeleteWithComputer(v bool) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.DeleteWithComputer = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetDevice(v string) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.Device = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetEnableAutoSnapshot(v bool) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.EnableAutoSnapshot = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetIaasId(v string) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.IaasId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetId(v string) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.Id = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetImageId(v string) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.ImageId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetName(v string) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.Name = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetPortable(v bool) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.Portable = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetProviderId(v string) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.ProviderId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetRegionId(v string) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.RegionId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetSize(v int64) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.Size = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetStatus(v string) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.Status = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetTenantId(v string) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.TenantId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetType(v string) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.Type = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetUtcCreate(v string) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.UtcCreate = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetUtcModified(v string) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.UtcModified = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetWorkspaceId(v string) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDisk) SetZoneId(v string) *DescribeCasComputersResponseBodyDataSystemDisk {
	s.ZoneId = &v
	return s
}

type DescribeCasComputersResponseBodyDataSystemDiskComputer struct {
	IaasId *string `json:"IaasId,omitempty" xml:"IaasId,omitempty"`
	Id     *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCasComputersResponseBodyDataSystemDiskComputer) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasComputersResponseBodyDataSystemDiskComputer) GoString() string {
	return s.String()
}

func (s *DescribeCasComputersResponseBodyDataSystemDiskComputer) SetIaasId(v string) *DescribeCasComputersResponseBodyDataSystemDiskComputer {
	s.IaasId = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDiskComputer) SetId(v string) *DescribeCasComputersResponseBodyDataSystemDiskComputer {
	s.Id = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDiskComputer) SetName(v string) *DescribeCasComputersResponseBodyDataSystemDiskComputer {
	s.Name = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataSystemDiskComputer) SetStatus(v string) *DescribeCasComputersResponseBodyDataSystemDiskComputer {
	s.Status = &v
	return s
}

type DescribeCasComputersResponseBodyDataTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCasComputersResponseBodyDataTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasComputersResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *DescribeCasComputersResponseBodyDataTags) SetKey(v string) *DescribeCasComputersResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *DescribeCasComputersResponseBodyDataTags) SetValue(v string) *DescribeCasComputersResponseBodyDataTags {
	s.Value = &v
	return s
}

type DescribeCasComputersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCasComputersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCasComputersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasComputersResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasComputersResponse) SetHeaders(v map[string]*string) *DescribeCasComputersResponse {
	s.Headers = v
	return s
}

func (s *DescribeCasComputersResponse) SetStatusCode(v int32) *DescribeCasComputersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCasComputersResponse) SetBody(v *DescribeCasComputersResponseBody) *DescribeCasComputersResponse {
	s.Body = v
	return s
}

type DisableMqSofamqWarnRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	WarnId     *int64  `json:"WarnId,omitempty" xml:"WarnId,omitempty"`
}

func (s DisableMqSofamqWarnRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableMqSofamqWarnRequest) GoString() string {
	return s.String()
}

func (s *DisableMqSofamqWarnRequest) SetInstanceId(v string) *DisableMqSofamqWarnRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableMqSofamqWarnRequest) SetWarnId(v int64) *DisableMqSofamqWarnRequest {
	s.WarnId = &v
	return s
}

type DisableMqSofamqWarnResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s DisableMqSofamqWarnResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableMqSofamqWarnResponseBody) GoString() string {
	return s.String()
}

func (s *DisableMqSofamqWarnResponseBody) SetRequestId(v string) *DisableMqSofamqWarnResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableMqSofamqWarnResponseBody) SetResultCode(v string) *DisableMqSofamqWarnResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DisableMqSofamqWarnResponseBody) SetResultMessage(v string) *DisableMqSofamqWarnResponseBody {
	s.ResultMessage = &v
	return s
}

type DisableMqSofamqWarnResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableMqSofamqWarnResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableMqSofamqWarnResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableMqSofamqWarnResponse) GoString() string {
	return s.String()
}

func (s *DisableMqSofamqWarnResponse) SetHeaders(v map[string]*string) *DisableMqSofamqWarnResponse {
	s.Headers = v
	return s
}

func (s *DisableMqSofamqWarnResponse) SetStatusCode(v int32) *DisableMqSofamqWarnResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableMqSofamqWarnResponse) SetBody(v *DisableMqSofamqWarnResponseBody) *DisableMqSofamqWarnResponse {
	s.Body = v
	return s
}

type EnableMqSofamqWarnRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	WarnId     *int64  `json:"WarnId,omitempty" xml:"WarnId,omitempty"`
}

func (s EnableMqSofamqWarnRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableMqSofamqWarnRequest) GoString() string {
	return s.String()
}

func (s *EnableMqSofamqWarnRequest) SetInstanceId(v string) *EnableMqSofamqWarnRequest {
	s.InstanceId = &v
	return s
}

func (s *EnableMqSofamqWarnRequest) SetWarnId(v int64) *EnableMqSofamqWarnRequest {
	s.WarnId = &v
	return s
}

type EnableMqSofamqWarnResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s EnableMqSofamqWarnResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableMqSofamqWarnResponseBody) GoString() string {
	return s.String()
}

func (s *EnableMqSofamqWarnResponseBody) SetRequestId(v string) *EnableMqSofamqWarnResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableMqSofamqWarnResponseBody) SetResultCode(v string) *EnableMqSofamqWarnResponseBody {
	s.ResultCode = &v
	return s
}

func (s *EnableMqSofamqWarnResponseBody) SetResultMessage(v string) *EnableMqSofamqWarnResponseBody {
	s.ResultMessage = &v
	return s
}

type EnableMqSofamqWarnResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableMqSofamqWarnResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableMqSofamqWarnResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableMqSofamqWarnResponse) GoString() string {
	return s.String()
}

func (s *EnableMqSofamqWarnResponse) SetHeaders(v map[string]*string) *EnableMqSofamqWarnResponse {
	s.Headers = v
	return s
}

func (s *EnableMqSofamqWarnResponse) SetStatusCode(v int32) *EnableMqSofamqWarnResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableMqSofamqWarnResponse) SetBody(v *EnableMqSofamqWarnResponseBody) *EnableMqSofamqWarnResponse {
	s.Body = v
	return s
}

type GetMqSofamqConsumerJStackRequest struct {
	Cell       *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetMqSofamqConsumerJStackRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqConsumerJStackRequest) GoString() string {
	return s.String()
}

func (s *GetMqSofamqConsumerJStackRequest) SetCell(v string) *GetMqSofamqConsumerJStackRequest {
	s.Cell = &v
	return s
}

func (s *GetMqSofamqConsumerJStackRequest) SetClientId(v string) *GetMqSofamqConsumerJStackRequest {
	s.ClientId = &v
	return s
}

func (s *GetMqSofamqConsumerJStackRequest) SetGroupId(v string) *GetMqSofamqConsumerJStackRequest {
	s.GroupId = &v
	return s
}

func (s *GetMqSofamqConsumerJStackRequest) SetInstanceId(v string) *GetMqSofamqConsumerJStackRequest {
	s.InstanceId = &v
	return s
}

type GetMqSofamqConsumerJStackResponseBody struct {
	Data          *GetMqSofamqConsumerJStackResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                    `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                    `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GetMqSofamqConsumerJStackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqConsumerJStackResponseBody) GoString() string {
	return s.String()
}

func (s *GetMqSofamqConsumerJStackResponseBody) SetData(v *GetMqSofamqConsumerJStackResponseBodyData) *GetMqSofamqConsumerJStackResponseBody {
	s.Data = v
	return s
}

func (s *GetMqSofamqConsumerJStackResponseBody) SetRequestId(v string) *GetMqSofamqConsumerJStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMqSofamqConsumerJStackResponseBody) SetResultCode(v string) *GetMqSofamqConsumerJStackResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetMqSofamqConsumerJStackResponseBody) SetResultMessage(v string) *GetMqSofamqConsumerJStackResponseBody {
	s.ResultMessage = &v
	return s
}

type GetMqSofamqConsumerJStackResponseBodyData struct {
	ClientId *string                                            `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	Jstack   []*GetMqSofamqConsumerJStackResponseBodyDataJstack `json:"Jstack,omitempty" xml:"Jstack,omitempty" type:"Repeated"`
}

func (s GetMqSofamqConsumerJStackResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqConsumerJStackResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMqSofamqConsumerJStackResponseBodyData) SetClientId(v string) *GetMqSofamqConsumerJStackResponseBodyData {
	s.ClientId = &v
	return s
}

func (s *GetMqSofamqConsumerJStackResponseBodyData) SetJstack(v []*GetMqSofamqConsumerJStackResponseBodyDataJstack) *GetMqSofamqConsumerJStackResponseBodyData {
	s.Jstack = v
	return s
}

type GetMqSofamqConsumerJStackResponseBodyDataJstack struct {
	Thread    *string   `json:"Thread,omitempty" xml:"Thread,omitempty"`
	TrackList []*string `json:"TrackList,omitempty" xml:"TrackList,omitempty" type:"Repeated"`
}

func (s GetMqSofamqConsumerJStackResponseBodyDataJstack) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqConsumerJStackResponseBodyDataJstack) GoString() string {
	return s.String()
}

func (s *GetMqSofamqConsumerJStackResponseBodyDataJstack) SetThread(v string) *GetMqSofamqConsumerJStackResponseBodyDataJstack {
	s.Thread = &v
	return s
}

func (s *GetMqSofamqConsumerJStackResponseBodyDataJstack) SetTrackList(v []*string) *GetMqSofamqConsumerJStackResponseBodyDataJstack {
	s.TrackList = v
	return s
}

type GetMqSofamqConsumerJStackResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMqSofamqConsumerJStackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMqSofamqConsumerJStackResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqConsumerJStackResponse) GoString() string {
	return s.String()
}

func (s *GetMqSofamqConsumerJStackResponse) SetHeaders(v map[string]*string) *GetMqSofamqConsumerJStackResponse {
	s.Headers = v
	return s
}

func (s *GetMqSofamqConsumerJStackResponse) SetStatusCode(v int32) *GetMqSofamqConsumerJStackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMqSofamqConsumerJStackResponse) SetBody(v *GetMqSofamqConsumerJStackResponseBody) *GetMqSofamqConsumerJStackResponse {
	s.Body = v
	return s
}

type GetMqSofamqConsumerStatusRequest struct {
	Cell       *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	Detail     *bool   `json:"Detail,omitempty" xml:"Detail,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NeedJstack *bool   `json:"NeedJstack,omitempty" xml:"NeedJstack,omitempty"`
}

func (s GetMqSofamqConsumerStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqConsumerStatusRequest) GoString() string {
	return s.String()
}

func (s *GetMqSofamqConsumerStatusRequest) SetCell(v string) *GetMqSofamqConsumerStatusRequest {
	s.Cell = &v
	return s
}

func (s *GetMqSofamqConsumerStatusRequest) SetDetail(v bool) *GetMqSofamqConsumerStatusRequest {
	s.Detail = &v
	return s
}

func (s *GetMqSofamqConsumerStatusRequest) SetGroupId(v string) *GetMqSofamqConsumerStatusRequest {
	s.GroupId = &v
	return s
}

func (s *GetMqSofamqConsumerStatusRequest) SetInstanceId(v string) *GetMqSofamqConsumerStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetMqSofamqConsumerStatusRequest) SetNeedJstack(v bool) *GetMqSofamqConsumerStatusRequest {
	s.NeedJstack = &v
	return s
}

type GetMqSofamqConsumerStatusResponseBody struct {
	Data          *GetMqSofamqConsumerStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                    `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                    `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GetMqSofamqConsumerStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqConsumerStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetMqSofamqConsumerStatusResponseBody) SetData(v *GetMqSofamqConsumerStatusResponseBodyData) *GetMqSofamqConsumerStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBody) SetRequestId(v string) *GetMqSofamqConsumerStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBody) SetResultCode(v string) *GetMqSofamqConsumerStatusResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBody) SetResultMessage(v string) *GetMqSofamqConsumerStatusResponseBody {
	s.ResultMessage = &v
	return s
}

type GetMqSofamqConsumerStatusResponseBodyData struct {
	ConnectionSet              []*GetMqSofamqConsumerStatusResponseBodyDataConnectionSet              `json:"ConnectionSet,omitempty" xml:"ConnectionSet,omitempty" type:"Repeated"`
	ConsumeModel               *string                                                                `json:"ConsumeModel,omitempty" xml:"ConsumeModel,omitempty"`
	ConsumeTps                 *string                                                                `json:"ConsumeTps,omitempty" xml:"ConsumeTps,omitempty"`
	ConsumerConnectionInfoList []*GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList `json:"ConsumerConnectionInfoList,omitempty" xml:"ConsumerConnectionInfoList,omitempty" type:"Repeated"`
	DelayTime                  *int64                                                                 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	DetailInTopicList          []*GetMqSofamqConsumerStatusResponseBodyDataDetailInTopicList          `json:"DetailInTopicList,omitempty" xml:"DetailInTopicList,omitempty" type:"Repeated"`
	InstanceId                 *string                                                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LastTimestamp              *int64                                                                 `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Online                     *bool                                                                  `json:"Online,omitempty" xml:"Online,omitempty"`
	RebalanceOk                *bool                                                                  `json:"RebalanceOk,omitempty" xml:"RebalanceOk,omitempty"`
	SubscriptionSame           *bool                                                                  `json:"SubscriptionSame,omitempty" xml:"SubscriptionSame,omitempty"`
	TotalDiff                  *int64                                                                 `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
}

func (s GetMqSofamqConsumerStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqConsumerStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMqSofamqConsumerStatusResponseBodyData) SetConnectionSet(v []*GetMqSofamqConsumerStatusResponseBodyDataConnectionSet) *GetMqSofamqConsumerStatusResponseBodyData {
	s.ConnectionSet = v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyData) SetConsumeModel(v string) *GetMqSofamqConsumerStatusResponseBodyData {
	s.ConsumeModel = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyData) SetConsumeTps(v string) *GetMqSofamqConsumerStatusResponseBodyData {
	s.ConsumeTps = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyData) SetConsumerConnectionInfoList(v []*GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList) *GetMqSofamqConsumerStatusResponseBodyData {
	s.ConsumerConnectionInfoList = v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyData) SetDelayTime(v int64) *GetMqSofamqConsumerStatusResponseBodyData {
	s.DelayTime = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyData) SetDetailInTopicList(v []*GetMqSofamqConsumerStatusResponseBodyDataDetailInTopicList) *GetMqSofamqConsumerStatusResponseBodyData {
	s.DetailInTopicList = v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyData) SetInstanceId(v string) *GetMqSofamqConsumerStatusResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyData) SetLastTimestamp(v int64) *GetMqSofamqConsumerStatusResponseBodyData {
	s.LastTimestamp = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyData) SetOnline(v bool) *GetMqSofamqConsumerStatusResponseBodyData {
	s.Online = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyData) SetRebalanceOk(v bool) *GetMqSofamqConsumerStatusResponseBodyData {
	s.RebalanceOk = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyData) SetSubscriptionSame(v bool) *GetMqSofamqConsumerStatusResponseBodyData {
	s.SubscriptionSame = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyData) SetTotalDiff(v int64) *GetMqSofamqConsumerStatusResponseBodyData {
	s.TotalDiff = &v
	return s
}

type GetMqSofamqConsumerStatusResponseBodyDataConnectionSet struct {
	ClientAddr *string `json:"ClientAddr,omitempty" xml:"ClientAddr,omitempty"`
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	RemoteIp   *string `json:"RemoteIp,omitempty" xml:"RemoteIp,omitempty"`
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetMqSofamqConsumerStatusResponseBodyDataConnectionSet) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqConsumerStatusResponseBodyDataConnectionSet) GoString() string {
	return s.String()
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConnectionSet) SetClientAddr(v string) *GetMqSofamqConsumerStatusResponseBodyDataConnectionSet {
	s.ClientAddr = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConnectionSet) SetClientId(v string) *GetMqSofamqConsumerStatusResponseBodyDataConnectionSet {
	s.ClientId = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConnectionSet) SetLanguage(v string) *GetMqSofamqConsumerStatusResponseBodyDataConnectionSet {
	s.Language = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConnectionSet) SetRemoteIp(v string) *GetMqSofamqConsumerStatusResponseBodyDataConnectionSet {
	s.RemoteIp = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConnectionSet) SetVersion(v string) *GetMqSofamqConsumerStatusResponseBodyDataConnectionSet {
	s.Version = &v
	return s
}

type GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList struct {
	ClientId        *string                                                                               `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	Connection      *string                                                                               `json:"Connection,omitempty" xml:"Connection,omitempty"`
	ConsumeType     *string                                                                               `json:"ConsumeType,omitempty" xml:"ConsumeType,omitempty"`
	Jstack          []*GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListJstack          `json:"Jstack,omitempty" xml:"Jstack,omitempty" type:"Repeated"`
	Language        *string                                                                               `json:"Language,omitempty" xml:"Language,omitempty"`
	LastTimestamp   *int64                                                                                `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	MessageModel    *string                                                                               `json:"MessageModel,omitempty" xml:"MessageModel,omitempty"`
	RunningDataList []*GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList `json:"RunningDataList,omitempty" xml:"RunningDataList,omitempty" type:"Repeated"`
	StartTimestamp  *int64                                                                                `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	SubscriptionSet []*GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet `json:"SubscriptionSet,omitempty" xml:"SubscriptionSet,omitempty" type:"Repeated"`
	ThreadCount     *int64                                                                                `json:"ThreadCount,omitempty" xml:"ThreadCount,omitempty"`
	Version         *string                                                                               `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList) GoString() string {
	return s.String()
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetClientId(v string) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.ClientId = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetConnection(v string) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.Connection = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetConsumeType(v string) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.ConsumeType = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetJstack(v []*GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListJstack) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.Jstack = v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetLanguage(v string) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.Language = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetLastTimestamp(v int64) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.LastTimestamp = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetMessageModel(v string) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.MessageModel = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetRunningDataList(v []*GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.RunningDataList = v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetStartTimestamp(v int64) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.StartTimestamp = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetSubscriptionSet(v []*GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.SubscriptionSet = v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetThreadCount(v int64) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.ThreadCount = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetVersion(v string) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.Version = &v
	return s
}

type GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListJstack struct {
	Thread    *string   `json:"Thread,omitempty" xml:"Thread,omitempty"`
	TrackList []*string `json:"TrackList,omitempty" xml:"TrackList,omitempty" type:"Repeated"`
}

func (s GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListJstack) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListJstack) GoString() string {
	return s.String()
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListJstack) SetThread(v string) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListJstack {
	s.Thread = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListJstack) SetTrackList(v []*string) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListJstack {
	s.TrackList = v
	return s
}

type GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList struct {
	Diff               *int64  `json:"Diff,omitempty" xml:"Diff,omitempty"`
	FailedCountPerHour *int64  `json:"FailedCountPerHour,omitempty" xml:"FailedCountPerHour,omitempty"`
	FailedTps          *string `json:"FailedTps,omitempty" xml:"FailedTps,omitempty"`
	GroupId            *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OkTps              *string `json:"OkTps,omitempty" xml:"OkTps,omitempty"`
	Rt                 *string `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Topic              *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList) GoString() string {
	return s.String()
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList) SetDiff(v int64) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList {
	s.Diff = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList) SetFailedCountPerHour(v int64) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList {
	s.FailedCountPerHour = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList) SetFailedTps(v string) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList {
	s.FailedTps = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList) SetGroupId(v string) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList {
	s.GroupId = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList) SetOkTps(v string) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList {
	s.OkTps = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList) SetRt(v string) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList {
	s.Rt = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList) SetTopic(v string) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList {
	s.Topic = &v
	return s
}

type GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet struct {
	SubString  *string   `json:"SubString,omitempty" xml:"SubString,omitempty"`
	SubVersion *int64    `json:"SubVersion,omitempty" xml:"SubVersion,omitempty"`
	TagsSet    []*string `json:"TagsSet,omitempty" xml:"TagsSet,omitempty" type:"Repeated"`
	Topic      *string   `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet) GoString() string {
	return s.String()
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet) SetSubString(v string) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet {
	s.SubString = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet) SetSubVersion(v int64) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet {
	s.SubVersion = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet) SetTagsSet(v []*string) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet {
	s.TagsSet = v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet) SetTopic(v string) *GetMqSofamqConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet {
	s.Topic = &v
	return s
}

type GetMqSofamqConsumerStatusResponseBodyDataDetailInTopicList struct {
	DelayTime     *int64  `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	LastTimestamp *int64  `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Topic         *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	TotalDiff     *int64  `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
}

func (s GetMqSofamqConsumerStatusResponseBodyDataDetailInTopicList) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqConsumerStatusResponseBodyDataDetailInTopicList) GoString() string {
	return s.String()
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataDetailInTopicList) SetDelayTime(v int64) *GetMqSofamqConsumerStatusResponseBodyDataDetailInTopicList {
	s.DelayTime = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataDetailInTopicList) SetLastTimestamp(v int64) *GetMqSofamqConsumerStatusResponseBodyDataDetailInTopicList {
	s.LastTimestamp = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataDetailInTopicList) SetTopic(v string) *GetMqSofamqConsumerStatusResponseBodyDataDetailInTopicList {
	s.Topic = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponseBodyDataDetailInTopicList) SetTotalDiff(v int64) *GetMqSofamqConsumerStatusResponseBodyDataDetailInTopicList {
	s.TotalDiff = &v
	return s
}

type GetMqSofamqConsumerStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMqSofamqConsumerStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMqSofamqConsumerStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqConsumerStatusResponse) GoString() string {
	return s.String()
}

func (s *GetMqSofamqConsumerStatusResponse) SetHeaders(v map[string]*string) *GetMqSofamqConsumerStatusResponse {
	s.Headers = v
	return s
}

func (s *GetMqSofamqConsumerStatusResponse) SetStatusCode(v int32) *GetMqSofamqConsumerStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMqSofamqConsumerStatusResponse) SetBody(v *GetMqSofamqConsumerStatusResponseBody) *GetMqSofamqConsumerStatusResponse {
	s.Body = v
	return s
}

type GetMqSofamqDLQMessageByIdRequest struct {
	Cell       *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s GetMqSofamqDLQMessageByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqDLQMessageByIdRequest) GoString() string {
	return s.String()
}

func (s *GetMqSofamqDLQMessageByIdRequest) SetCell(v string) *GetMqSofamqDLQMessageByIdRequest {
	s.Cell = &v
	return s
}

func (s *GetMqSofamqDLQMessageByIdRequest) SetGroupId(v string) *GetMqSofamqDLQMessageByIdRequest {
	s.GroupId = &v
	return s
}

func (s *GetMqSofamqDLQMessageByIdRequest) SetInstanceId(v string) *GetMqSofamqDLQMessageByIdRequest {
	s.InstanceId = &v
	return s
}

func (s *GetMqSofamqDLQMessageByIdRequest) SetMsgId(v string) *GetMqSofamqDLQMessageByIdRequest {
	s.MsgId = &v
	return s
}

type GetMqSofamqDLQMessageByIdResponseBody struct {
	Data          *GetMqSofamqDLQMessageByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                    `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                    `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GetMqSofamqDLQMessageByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqDLQMessageByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetMqSofamqDLQMessageByIdResponseBody) SetData(v *GetMqSofamqDLQMessageByIdResponseBodyData) *GetMqSofamqDLQMessageByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetMqSofamqDLQMessageByIdResponseBody) SetRequestId(v string) *GetMqSofamqDLQMessageByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMqSofamqDLQMessageByIdResponseBody) SetResultCode(v string) *GetMqSofamqDLQMessageByIdResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetMqSofamqDLQMessageByIdResponseBody) SetResultMessage(v string) *GetMqSofamqDLQMessageByIdResponseBody {
	s.ResultMessage = &v
	return s
}

type GetMqSofamqDLQMessageByIdResponseBodyData struct {
	Body           *string                                                  `json:"Body,omitempty" xml:"Body,omitempty"`
	BodyCrc        *int64                                                   `json:"BodyCrc,omitempty" xml:"BodyCrc,omitempty"`
	BornHost       *string                                                  `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	BornTimestamp  *int64                                                   `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	InstanceId     *string                                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId          *string                                                  `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	PropertyList   []*GetMqSofamqDLQMessageByIdResponseBodyDataPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	ReconsumeTimes *int64                                                   `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	StoreHost      *string                                                  `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	StoreSize      *int64                                                   `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	StoreTimestamp *int64                                                   `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	Topic          *string                                                  `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetMqSofamqDLQMessageByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqDLQMessageByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMqSofamqDLQMessageByIdResponseBodyData) SetBody(v string) *GetMqSofamqDLQMessageByIdResponseBodyData {
	s.Body = &v
	return s
}

func (s *GetMqSofamqDLQMessageByIdResponseBodyData) SetBodyCrc(v int64) *GetMqSofamqDLQMessageByIdResponseBodyData {
	s.BodyCrc = &v
	return s
}

func (s *GetMqSofamqDLQMessageByIdResponseBodyData) SetBornHost(v string) *GetMqSofamqDLQMessageByIdResponseBodyData {
	s.BornHost = &v
	return s
}

func (s *GetMqSofamqDLQMessageByIdResponseBodyData) SetBornTimestamp(v int64) *GetMqSofamqDLQMessageByIdResponseBodyData {
	s.BornTimestamp = &v
	return s
}

func (s *GetMqSofamqDLQMessageByIdResponseBodyData) SetInstanceId(v string) *GetMqSofamqDLQMessageByIdResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetMqSofamqDLQMessageByIdResponseBodyData) SetMsgId(v string) *GetMqSofamqDLQMessageByIdResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *GetMqSofamqDLQMessageByIdResponseBodyData) SetPropertyList(v []*GetMqSofamqDLQMessageByIdResponseBodyDataPropertyList) *GetMqSofamqDLQMessageByIdResponseBodyData {
	s.PropertyList = v
	return s
}

func (s *GetMqSofamqDLQMessageByIdResponseBodyData) SetReconsumeTimes(v int64) *GetMqSofamqDLQMessageByIdResponseBodyData {
	s.ReconsumeTimes = &v
	return s
}

func (s *GetMqSofamqDLQMessageByIdResponseBodyData) SetStoreHost(v string) *GetMqSofamqDLQMessageByIdResponseBodyData {
	s.StoreHost = &v
	return s
}

func (s *GetMqSofamqDLQMessageByIdResponseBodyData) SetStoreSize(v int64) *GetMqSofamqDLQMessageByIdResponseBodyData {
	s.StoreSize = &v
	return s
}

func (s *GetMqSofamqDLQMessageByIdResponseBodyData) SetStoreTimestamp(v int64) *GetMqSofamqDLQMessageByIdResponseBodyData {
	s.StoreTimestamp = &v
	return s
}

func (s *GetMqSofamqDLQMessageByIdResponseBodyData) SetTopic(v string) *GetMqSofamqDLQMessageByIdResponseBodyData {
	s.Topic = &v
	return s
}

type GetMqSofamqDLQMessageByIdResponseBodyDataPropertyList struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetMqSofamqDLQMessageByIdResponseBodyDataPropertyList) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqDLQMessageByIdResponseBodyDataPropertyList) GoString() string {
	return s.String()
}

func (s *GetMqSofamqDLQMessageByIdResponseBodyDataPropertyList) SetName(v string) *GetMqSofamqDLQMessageByIdResponseBodyDataPropertyList {
	s.Name = &v
	return s
}

func (s *GetMqSofamqDLQMessageByIdResponseBodyDataPropertyList) SetValue(v string) *GetMqSofamqDLQMessageByIdResponseBodyDataPropertyList {
	s.Value = &v
	return s
}

type GetMqSofamqDLQMessageByIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMqSofamqDLQMessageByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMqSofamqDLQMessageByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqDLQMessageByIdResponse) GoString() string {
	return s.String()
}

func (s *GetMqSofamqDLQMessageByIdResponse) SetHeaders(v map[string]*string) *GetMqSofamqDLQMessageByIdResponse {
	s.Headers = v
	return s
}

func (s *GetMqSofamqDLQMessageByIdResponse) SetStatusCode(v int32) *GetMqSofamqDLQMessageByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMqSofamqDLQMessageByIdResponse) SetBody(v *GetMqSofamqDLQMessageByIdResponseBody) *GetMqSofamqDLQMessageByIdResponse {
	s.Body = v
	return s
}

type GetMqSofamqMessageByIdRequest struct {
	Cell       *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetMqSofamqMessageByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqMessageByIdRequest) GoString() string {
	return s.String()
}

func (s *GetMqSofamqMessageByIdRequest) SetCell(v string) *GetMqSofamqMessageByIdRequest {
	s.Cell = &v
	return s
}

func (s *GetMqSofamqMessageByIdRequest) SetInstanceId(v string) *GetMqSofamqMessageByIdRequest {
	s.InstanceId = &v
	return s
}

func (s *GetMqSofamqMessageByIdRequest) SetMsgId(v string) *GetMqSofamqMessageByIdRequest {
	s.MsgId = &v
	return s
}

func (s *GetMqSofamqMessageByIdRequest) SetTopic(v string) *GetMqSofamqMessageByIdRequest {
	s.Topic = &v
	return s
}

type GetMqSofamqMessageByIdResponseBody struct {
	Data          *GetMqSofamqMessageByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId     *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                 `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GetMqSofamqMessageByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqMessageByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetMqSofamqMessageByIdResponseBody) SetData(v *GetMqSofamqMessageByIdResponseBodyData) *GetMqSofamqMessageByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetMqSofamqMessageByIdResponseBody) SetRequestId(v string) *GetMqSofamqMessageByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMqSofamqMessageByIdResponseBody) SetResultCode(v string) *GetMqSofamqMessageByIdResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetMqSofamqMessageByIdResponseBody) SetResultMessage(v string) *GetMqSofamqMessageByIdResponseBody {
	s.ResultMessage = &v
	return s
}

type GetMqSofamqMessageByIdResponseBodyData struct {
	Body           *string                                               `json:"Body,omitempty" xml:"Body,omitempty"`
	BodyCrc        *int64                                                `json:"BodyCrc,omitempty" xml:"BodyCrc,omitempty"`
	BornHost       *string                                               `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	BornTimestamp  *int64                                                `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	InstanceId     *string                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId          *string                                               `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	PropertyList   []*GetMqSofamqMessageByIdResponseBodyDataPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	ReconsumeTimes *int64                                                `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	StoreHost      *string                                               `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	StoreSize      *int64                                                `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	StoreTimestamp *int64                                                `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	Topic          *string                                               `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetMqSofamqMessageByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqMessageByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMqSofamqMessageByIdResponseBodyData) SetBody(v string) *GetMqSofamqMessageByIdResponseBodyData {
	s.Body = &v
	return s
}

func (s *GetMqSofamqMessageByIdResponseBodyData) SetBodyCrc(v int64) *GetMqSofamqMessageByIdResponseBodyData {
	s.BodyCrc = &v
	return s
}

func (s *GetMqSofamqMessageByIdResponseBodyData) SetBornHost(v string) *GetMqSofamqMessageByIdResponseBodyData {
	s.BornHost = &v
	return s
}

func (s *GetMqSofamqMessageByIdResponseBodyData) SetBornTimestamp(v int64) *GetMqSofamqMessageByIdResponseBodyData {
	s.BornTimestamp = &v
	return s
}

func (s *GetMqSofamqMessageByIdResponseBodyData) SetInstanceId(v string) *GetMqSofamqMessageByIdResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetMqSofamqMessageByIdResponseBodyData) SetMsgId(v string) *GetMqSofamqMessageByIdResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *GetMqSofamqMessageByIdResponseBodyData) SetPropertyList(v []*GetMqSofamqMessageByIdResponseBodyDataPropertyList) *GetMqSofamqMessageByIdResponseBodyData {
	s.PropertyList = v
	return s
}

func (s *GetMqSofamqMessageByIdResponseBodyData) SetReconsumeTimes(v int64) *GetMqSofamqMessageByIdResponseBodyData {
	s.ReconsumeTimes = &v
	return s
}

func (s *GetMqSofamqMessageByIdResponseBodyData) SetStoreHost(v string) *GetMqSofamqMessageByIdResponseBodyData {
	s.StoreHost = &v
	return s
}

func (s *GetMqSofamqMessageByIdResponseBodyData) SetStoreSize(v int64) *GetMqSofamqMessageByIdResponseBodyData {
	s.StoreSize = &v
	return s
}

func (s *GetMqSofamqMessageByIdResponseBodyData) SetStoreTimestamp(v int64) *GetMqSofamqMessageByIdResponseBodyData {
	s.StoreTimestamp = &v
	return s
}

func (s *GetMqSofamqMessageByIdResponseBodyData) SetTopic(v string) *GetMqSofamqMessageByIdResponseBodyData {
	s.Topic = &v
	return s
}

type GetMqSofamqMessageByIdResponseBodyDataPropertyList struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetMqSofamqMessageByIdResponseBodyDataPropertyList) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqMessageByIdResponseBodyDataPropertyList) GoString() string {
	return s.String()
}

func (s *GetMqSofamqMessageByIdResponseBodyDataPropertyList) SetName(v string) *GetMqSofamqMessageByIdResponseBodyDataPropertyList {
	s.Name = &v
	return s
}

func (s *GetMqSofamqMessageByIdResponseBodyDataPropertyList) SetValue(v string) *GetMqSofamqMessageByIdResponseBodyDataPropertyList {
	s.Value = &v
	return s
}

type GetMqSofamqMessageByIdResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMqSofamqMessageByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMqSofamqMessageByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqMessageByIdResponse) GoString() string {
	return s.String()
}

func (s *GetMqSofamqMessageByIdResponse) SetHeaders(v map[string]*string) *GetMqSofamqMessageByIdResponse {
	s.Headers = v
	return s
}

func (s *GetMqSofamqMessageByIdResponse) SetStatusCode(v int32) *GetMqSofamqMessageByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMqSofamqMessageByIdResponse) SetBody(v *GetMqSofamqMessageByIdResponseBody) *GetMqSofamqMessageByIdResponse {
	s.Body = v
	return s
}

type GetMqSofamqTraceByMsgIdRequest struct {
	BeginTime  *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	Cell       *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetMqSofamqTraceByMsgIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqTraceByMsgIdRequest) GoString() string {
	return s.String()
}

func (s *GetMqSofamqTraceByMsgIdRequest) SetBeginTime(v int64) *GetMqSofamqTraceByMsgIdRequest {
	s.BeginTime = &v
	return s
}

func (s *GetMqSofamqTraceByMsgIdRequest) SetCell(v string) *GetMqSofamqTraceByMsgIdRequest {
	s.Cell = &v
	return s
}

func (s *GetMqSofamqTraceByMsgIdRequest) SetEndTime(v int64) *GetMqSofamqTraceByMsgIdRequest {
	s.EndTime = &v
	return s
}

func (s *GetMqSofamqTraceByMsgIdRequest) SetInstanceId(v string) *GetMqSofamqTraceByMsgIdRequest {
	s.InstanceId = &v
	return s
}

func (s *GetMqSofamqTraceByMsgIdRequest) SetMsgId(v string) *GetMqSofamqTraceByMsgIdRequest {
	s.MsgId = &v
	return s
}

func (s *GetMqSofamqTraceByMsgIdRequest) SetTopic(v string) *GetMqSofamqTraceByMsgIdRequest {
	s.Topic = &v
	return s
}

type GetMqSofamqTraceByMsgIdResponseBody struct {
	QueryId       *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GetMqSofamqTraceByMsgIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqTraceByMsgIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetMqSofamqTraceByMsgIdResponseBody) SetQueryId(v string) *GetMqSofamqTraceByMsgIdResponseBody {
	s.QueryId = &v
	return s
}

func (s *GetMqSofamqTraceByMsgIdResponseBody) SetRequestId(v string) *GetMqSofamqTraceByMsgIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMqSofamqTraceByMsgIdResponseBody) SetResultCode(v string) *GetMqSofamqTraceByMsgIdResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetMqSofamqTraceByMsgIdResponseBody) SetResultMessage(v string) *GetMqSofamqTraceByMsgIdResponseBody {
	s.ResultMessage = &v
	return s
}

type GetMqSofamqTraceByMsgIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMqSofamqTraceByMsgIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMqSofamqTraceByMsgIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqTraceByMsgIdResponse) GoString() string {
	return s.String()
}

func (s *GetMqSofamqTraceByMsgIdResponse) SetHeaders(v map[string]*string) *GetMqSofamqTraceByMsgIdResponse {
	s.Headers = v
	return s
}

func (s *GetMqSofamqTraceByMsgIdResponse) SetStatusCode(v int32) *GetMqSofamqTraceByMsgIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMqSofamqTraceByMsgIdResponse) SetBody(v *GetMqSofamqTraceByMsgIdResponseBody) *GetMqSofamqTraceByMsgIdResponse {
	s.Body = v
	return s
}

type GetMqSofamqTraceResultRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QueryId    *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s GetMqSofamqTraceResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqTraceResultRequest) GoString() string {
	return s.String()
}

func (s *GetMqSofamqTraceResultRequest) SetInstanceId(v string) *GetMqSofamqTraceResultRequest {
	s.InstanceId = &v
	return s
}

func (s *GetMqSofamqTraceResultRequest) SetQueryId(v string) *GetMqSofamqTraceResultRequest {
	s.QueryId = &v
	return s
}

type GetMqSofamqTraceResultResponseBody struct {
	Data          *GetMqSofamqTraceResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId     *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                 `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GetMqSofamqTraceResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqTraceResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetMqSofamqTraceResultResponseBody) SetData(v *GetMqSofamqTraceResultResponseBodyData) *GetMqSofamqTraceResultResponseBody {
	s.Data = v
	return s
}

func (s *GetMqSofamqTraceResultResponseBody) SetRequestId(v string) *GetMqSofamqTraceResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBody) SetResultCode(v string) *GetMqSofamqTraceResultResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBody) SetResultMessage(v string) *GetMqSofamqTraceResultResponseBody {
	s.ResultMessage = &v
	return s
}

type GetMqSofamqTraceResultResponseBodyData struct {
	CreateTime *int64                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceId *string                                            `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId      *string                                            `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	MsgKey     *string                                            `json:"MsgKey,omitempty" xml:"MsgKey,omitempty"`
	QueryId    *string                                            `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	Status     *string                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	Topic      *string                                            `json:"Topic,omitempty" xml:"Topic,omitempty"`
	TraceList  []*GetMqSofamqTraceResultResponseBodyDataTraceList `json:"TraceList,omitempty" xml:"TraceList,omitempty" type:"Repeated"`
	UpdateTime *int64                                             `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId     *string                                            `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetMqSofamqTraceResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqTraceResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMqSofamqTraceResultResponseBodyData) SetCreateTime(v int64) *GetMqSofamqTraceResultResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyData) SetInstanceId(v string) *GetMqSofamqTraceResultResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyData) SetMsgId(v string) *GetMqSofamqTraceResultResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyData) SetMsgKey(v string) *GetMqSofamqTraceResultResponseBodyData {
	s.MsgKey = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyData) SetQueryId(v string) *GetMqSofamqTraceResultResponseBodyData {
	s.QueryId = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyData) SetStatus(v string) *GetMqSofamqTraceResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyData) SetTopic(v string) *GetMqSofamqTraceResultResponseBodyData {
	s.Topic = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyData) SetTraceList(v []*GetMqSofamqTraceResultResponseBodyDataTraceList) *GetMqSofamqTraceResultResponseBodyData {
	s.TraceList = v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyData) SetUpdateTime(v int64) *GetMqSofamqTraceResultResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyData) SetUserId(v string) *GetMqSofamqTraceResultResponseBodyData {
	s.UserId = &v
	return s
}

type GetMqSofamqTraceResultResponseBodyDataTraceList struct {
	BornHost     *string                                                   `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	Cell         *string                                                   `json:"Cell,omitempty" xml:"Cell,omitempty"`
	CostTime     *int64                                                    `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	MsgId        *string                                                   `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	MsgKey       *string                                                   `json:"MsgKey,omitempty" xml:"MsgKey,omitempty"`
	PubGroupName *string                                                   `json:"PubGroupName,omitempty" xml:"PubGroupName,omitempty"`
	PubTime      *int64                                                    `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	Status       *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	SubList      []*GetMqSofamqTraceResultResponseBodyDataTraceListSubList `json:"SubList,omitempty" xml:"SubList,omitempty" type:"Repeated"`
	Tag          *string                                                   `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Topic        *string                                                   `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetMqSofamqTraceResultResponseBodyDataTraceList) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqTraceResultResponseBodyDataTraceList) GoString() string {
	return s.String()
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceList) SetBornHost(v string) *GetMqSofamqTraceResultResponseBodyDataTraceList {
	s.BornHost = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceList) SetCell(v string) *GetMqSofamqTraceResultResponseBodyDataTraceList {
	s.Cell = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceList) SetCostTime(v int64) *GetMqSofamqTraceResultResponseBodyDataTraceList {
	s.CostTime = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceList) SetMsgId(v string) *GetMqSofamqTraceResultResponseBodyDataTraceList {
	s.MsgId = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceList) SetMsgKey(v string) *GetMqSofamqTraceResultResponseBodyDataTraceList {
	s.MsgKey = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceList) SetPubGroupName(v string) *GetMqSofamqTraceResultResponseBodyDataTraceList {
	s.PubGroupName = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceList) SetPubTime(v int64) *GetMqSofamqTraceResultResponseBodyDataTraceList {
	s.PubTime = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceList) SetStatus(v string) *GetMqSofamqTraceResultResponseBodyDataTraceList {
	s.Status = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceList) SetSubList(v []*GetMqSofamqTraceResultResponseBodyDataTraceListSubList) *GetMqSofamqTraceResultResponseBodyDataTraceList {
	s.SubList = v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceList) SetTag(v string) *GetMqSofamqTraceResultResponseBodyDataTraceList {
	s.Tag = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceList) SetTopic(v string) *GetMqSofamqTraceResultResponseBodyDataTraceList {
	s.Topic = &v
	return s
}

type GetMqSofamqTraceResultResponseBodyDataTraceListSubList struct {
	Cell         *string                                                             `json:"Cell,omitempty" xml:"Cell,omitempty"`
	ClientList   []*GetMqSofamqTraceResultResponseBodyDataTraceListSubListClientList `json:"ClientList,omitempty" xml:"ClientList,omitempty" type:"Repeated"`
	FailCount    *int64                                                              `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	SubGroupName *string                                                             `json:"SubGroupName,omitempty" xml:"SubGroupName,omitempty"`
	SuccessCount *int64                                                              `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s GetMqSofamqTraceResultResponseBodyDataTraceListSubList) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqTraceResultResponseBodyDataTraceListSubList) GoString() string {
	return s.String()
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceListSubList) SetCell(v string) *GetMqSofamqTraceResultResponseBodyDataTraceListSubList {
	s.Cell = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceListSubList) SetClientList(v []*GetMqSofamqTraceResultResponseBodyDataTraceListSubListClientList) *GetMqSofamqTraceResultResponseBodyDataTraceListSubList {
	s.ClientList = v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceListSubList) SetFailCount(v int64) *GetMqSofamqTraceResultResponseBodyDataTraceListSubList {
	s.FailCount = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceListSubList) SetSubGroupName(v string) *GetMqSofamqTraceResultResponseBodyDataTraceListSubList {
	s.SubGroupName = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceListSubList) SetSuccessCount(v int64) *GetMqSofamqTraceResultResponseBodyDataTraceListSubList {
	s.SuccessCount = &v
	return s
}

type GetMqSofamqTraceResultResponseBodyDataTraceListSubListClientList struct {
	ClientHost     *string `json:"ClientHost,omitempty" xml:"ClientHost,omitempty"`
	CostTime       *int64  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	ReconsumeTimes *int64  `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubGroupName   *string `json:"SubGroupName,omitempty" xml:"SubGroupName,omitempty"`
	SubTime        *int64  `json:"SubTime,omitempty" xml:"SubTime,omitempty"`
}

func (s GetMqSofamqTraceResultResponseBodyDataTraceListSubListClientList) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqTraceResultResponseBodyDataTraceListSubListClientList) GoString() string {
	return s.String()
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceListSubListClientList) SetClientHost(v string) *GetMqSofamqTraceResultResponseBodyDataTraceListSubListClientList {
	s.ClientHost = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceListSubListClientList) SetCostTime(v int64) *GetMqSofamqTraceResultResponseBodyDataTraceListSubListClientList {
	s.CostTime = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceListSubListClientList) SetReconsumeTimes(v int64) *GetMqSofamqTraceResultResponseBodyDataTraceListSubListClientList {
	s.ReconsumeTimes = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceListSubListClientList) SetStatus(v string) *GetMqSofamqTraceResultResponseBodyDataTraceListSubListClientList {
	s.Status = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceListSubListClientList) SetSubGroupName(v string) *GetMqSofamqTraceResultResponseBodyDataTraceListSubListClientList {
	s.SubGroupName = &v
	return s
}

func (s *GetMqSofamqTraceResultResponseBodyDataTraceListSubListClientList) SetSubTime(v int64) *GetMqSofamqTraceResultResponseBodyDataTraceListSubListClientList {
	s.SubTime = &v
	return s
}

type GetMqSofamqTraceResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMqSofamqTraceResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMqSofamqTraceResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMqSofamqTraceResultResponse) GoString() string {
	return s.String()
}

func (s *GetMqSofamqTraceResultResponse) SetHeaders(v map[string]*string) *GetMqSofamqTraceResultResponse {
	s.Headers = v
	return s
}

func (s *GetMqSofamqTraceResultResponse) SetStatusCode(v int32) *GetMqSofamqTraceResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMqSofamqTraceResultResponse) SetBody(v *GetMqSofamqTraceResultResponseBody) *GetMqSofamqTraceResultResponse {
	s.Body = v
	return s
}

type GetMsConfigAttributesRequest struct {
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetMsConfigAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMsConfigAttributesRequest) GoString() string {
	return s.String()
}

func (s *GetMsConfigAttributesRequest) SetId(v int64) *GetMsConfigAttributesRequest {
	s.Id = &v
	return s
}

func (s *GetMsConfigAttributesRequest) SetInstanceId(v string) *GetMsConfigAttributesRequest {
	s.InstanceId = &v
	return s
}

type GetMsConfigAttributesResponseBody struct {
	Attribute     *GetMsConfigAttributesResponseBodyAttribute `json:"Attribute,omitempty" xml:"Attribute,omitempty" type:"Struct"`
	RequestId     *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                     `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                     `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GetMsConfigAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMsConfigAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *GetMsConfigAttributesResponseBody) SetAttribute(v *GetMsConfigAttributesResponseBodyAttribute) *GetMsConfigAttributesResponseBody {
	s.Attribute = v
	return s
}

func (s *GetMsConfigAttributesResponseBody) SetRequestId(v string) *GetMsConfigAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMsConfigAttributesResponseBody) SetResultCode(v string) *GetMsConfigAttributesResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetMsConfigAttributesResponseBody) SetResultMessage(v string) *GetMsConfigAttributesResponseBody {
	s.ResultMessage = &v
	return s
}

type GetMsConfigAttributesResponseBodyAttribute struct {
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	Desc          *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetMsConfigAttributesResponseBodyAttribute) String() string {
	return tea.Prettify(s)
}

func (s GetMsConfigAttributesResponseBodyAttribute) GoString() string {
	return s.String()
}

func (s *GetMsConfigAttributesResponseBodyAttribute) SetAttributeName(v string) *GetMsConfigAttributesResponseBodyAttribute {
	s.AttributeName = &v
	return s
}

func (s *GetMsConfigAttributesResponseBodyAttribute) SetDesc(v string) *GetMsConfigAttributesResponseBodyAttribute {
	s.Desc = &v
	return s
}

func (s *GetMsConfigAttributesResponseBodyAttribute) SetId(v int64) *GetMsConfigAttributesResponseBodyAttribute {
	s.Id = &v
	return s
}

func (s *GetMsConfigAttributesResponseBodyAttribute) SetInstanceId(v string) *GetMsConfigAttributesResponseBodyAttribute {
	s.InstanceId = &v
	return s
}

type GetMsConfigAttributesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMsConfigAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMsConfigAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMsConfigAttributesResponse) GoString() string {
	return s.String()
}

func (s *GetMsConfigAttributesResponse) SetHeaders(v map[string]*string) *GetMsConfigAttributesResponse {
	s.Headers = v
	return s
}

func (s *GetMsConfigAttributesResponse) SetStatusCode(v int32) *GetMsConfigAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMsConfigAttributesResponse) SetBody(v *GetMsConfigAttributesResponseBody) *GetMsConfigAttributesResponse {
	s.Body = v
	return s
}

type GetMsConfigResourcesRequest struct {
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetMsConfigResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMsConfigResourcesRequest) GoString() string {
	return s.String()
}

func (s *GetMsConfigResourcesRequest) SetId(v int64) *GetMsConfigResourcesRequest {
	s.Id = &v
	return s
}

func (s *GetMsConfigResourcesRequest) SetInstanceId(v string) *GetMsConfigResourcesRequest {
	s.InstanceId = &v
	return s
}

type GetMsConfigResourcesResponseBody struct {
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resource      *GetMsConfigResourcesResponseBodyResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	ResultCode    *string                                   `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                   `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GetMsConfigResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMsConfigResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *GetMsConfigResourcesResponseBody) SetRequestId(v string) *GetMsConfigResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMsConfigResourcesResponseBody) SetResource(v *GetMsConfigResourcesResponseBodyResource) *GetMsConfigResourcesResponseBody {
	s.Resource = v
	return s
}

func (s *GetMsConfigResourcesResponseBody) SetResultCode(v string) *GetMsConfigResourcesResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetMsConfigResourcesResponseBody) SetResultMessage(v string) *GetMsConfigResourcesResponseBody {
	s.ResultMessage = &v
	return s
}

type GetMsConfigResourcesResponseBodyResource struct {
	AppName    *string                                               `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Attributes []*GetMsConfigResourcesResponseBodyResourceAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Desc       *string                                               `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Id         *int64                                                `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId *string                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region     *string                                               `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceId *string                                               `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s GetMsConfigResourcesResponseBodyResource) String() string {
	return tea.Prettify(s)
}

func (s GetMsConfigResourcesResponseBodyResource) GoString() string {
	return s.String()
}

func (s *GetMsConfigResourcesResponseBodyResource) SetAppName(v string) *GetMsConfigResourcesResponseBodyResource {
	s.AppName = &v
	return s
}

func (s *GetMsConfigResourcesResponseBodyResource) SetAttributes(v []*GetMsConfigResourcesResponseBodyResourceAttributes) *GetMsConfigResourcesResponseBodyResource {
	s.Attributes = v
	return s
}

func (s *GetMsConfigResourcesResponseBodyResource) SetDesc(v string) *GetMsConfigResourcesResponseBodyResource {
	s.Desc = &v
	return s
}

func (s *GetMsConfigResourcesResponseBodyResource) SetId(v int64) *GetMsConfigResourcesResponseBodyResource {
	s.Id = &v
	return s
}

func (s *GetMsConfigResourcesResponseBodyResource) SetInstanceId(v string) *GetMsConfigResourcesResponseBodyResource {
	s.InstanceId = &v
	return s
}

func (s *GetMsConfigResourcesResponseBodyResource) SetRegion(v string) *GetMsConfigResourcesResponseBodyResource {
	s.Region = &v
	return s
}

func (s *GetMsConfigResourcesResponseBodyResource) SetResourceId(v string) *GetMsConfigResourcesResponseBodyResource {
	s.ResourceId = &v
	return s
}

type GetMsConfigResourcesResponseBodyResourceAttributes struct {
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	Desc          *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetMsConfigResourcesResponseBodyResourceAttributes) String() string {
	return tea.Prettify(s)
}

func (s GetMsConfigResourcesResponseBodyResourceAttributes) GoString() string {
	return s.String()
}

func (s *GetMsConfigResourcesResponseBodyResourceAttributes) SetAttributeName(v string) *GetMsConfigResourcesResponseBodyResourceAttributes {
	s.AttributeName = &v
	return s
}

func (s *GetMsConfigResourcesResponseBodyResourceAttributes) SetDesc(v string) *GetMsConfigResourcesResponseBodyResourceAttributes {
	s.Desc = &v
	return s
}

func (s *GetMsConfigResourcesResponseBodyResourceAttributes) SetId(v int64) *GetMsConfigResourcesResponseBodyResourceAttributes {
	s.Id = &v
	return s
}

func (s *GetMsConfigResourcesResponseBodyResourceAttributes) SetInstanceId(v string) *GetMsConfigResourcesResponseBodyResourceAttributes {
	s.InstanceId = &v
	return s
}

type GetMsConfigResourcesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMsConfigResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMsConfigResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMsConfigResourcesResponse) GoString() string {
	return s.String()
}

func (s *GetMsConfigResourcesResponse) SetHeaders(v map[string]*string) *GetMsConfigResourcesResponse {
	s.Headers = v
	return s
}

func (s *GetMsConfigResourcesResponse) SetStatusCode(v int32) *GetMsConfigResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMsConfigResourcesResponse) SetBody(v *GetMsConfigResourcesResponseBody) *GetMsConfigResourcesResponse {
	s.Body = v
	return s
}

type GrayPushMsConfigDataRequest struct {
	AttributeId *int64  `json:"AttributeId,omitempty" xml:"AttributeId,omitempty"`
	Data        *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Hosts       *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Operator    *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s GrayPushMsConfigDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GrayPushMsConfigDataRequest) GoString() string {
	return s.String()
}

func (s *GrayPushMsConfigDataRequest) SetAttributeId(v int64) *GrayPushMsConfigDataRequest {
	s.AttributeId = &v
	return s
}

func (s *GrayPushMsConfigDataRequest) SetData(v string) *GrayPushMsConfigDataRequest {
	s.Data = &v
	return s
}

func (s *GrayPushMsConfigDataRequest) SetHosts(v string) *GrayPushMsConfigDataRequest {
	s.Hosts = &v
	return s
}

func (s *GrayPushMsConfigDataRequest) SetInstanceId(v string) *GrayPushMsConfigDataRequest {
	s.InstanceId = &v
	return s
}

func (s *GrayPushMsConfigDataRequest) SetOperator(v string) *GrayPushMsConfigDataRequest {
	s.Operator = &v
	return s
}

type GrayPushMsConfigDataResponseBody struct {
	PushResult    []*GrayPushMsConfigDataResponseBodyPushResult `json:"PushResult,omitempty" xml:"PushResult,omitempty" type:"Repeated"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                       `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                       `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GrayPushMsConfigDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrayPushMsConfigDataResponseBody) GoString() string {
	return s.String()
}

func (s *GrayPushMsConfigDataResponseBody) SetPushResult(v []*GrayPushMsConfigDataResponseBodyPushResult) *GrayPushMsConfigDataResponseBody {
	s.PushResult = v
	return s
}

func (s *GrayPushMsConfigDataResponseBody) SetRequestId(v string) *GrayPushMsConfigDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrayPushMsConfigDataResponseBody) SetResultCode(v string) *GrayPushMsConfigDataResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GrayPushMsConfigDataResponseBody) SetResultMessage(v string) *GrayPushMsConfigDataResponseBody {
	s.ResultMessage = &v
	return s
}

type GrayPushMsConfigDataResponseBodyPushResult struct {
	Host    *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GrayPushMsConfigDataResponseBodyPushResult) String() string {
	return tea.Prettify(s)
}

func (s GrayPushMsConfigDataResponseBodyPushResult) GoString() string {
	return s.String()
}

func (s *GrayPushMsConfigDataResponseBodyPushResult) SetHost(v string) *GrayPushMsConfigDataResponseBodyPushResult {
	s.Host = &v
	return s
}

func (s *GrayPushMsConfigDataResponseBodyPushResult) SetSuccess(v bool) *GrayPushMsConfigDataResponseBodyPushResult {
	s.Success = &v
	return s
}

type GrayPushMsConfigDataResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GrayPushMsConfigDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrayPushMsConfigDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GrayPushMsConfigDataResponse) GoString() string {
	return s.String()
}

func (s *GrayPushMsConfigDataResponse) SetHeaders(v map[string]*string) *GrayPushMsConfigDataResponse {
	s.Headers = v
	return s
}

func (s *GrayPushMsConfigDataResponse) SetStatusCode(v int32) *GrayPushMsConfigDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GrayPushMsConfigDataResponse) SetBody(v *GrayPushMsConfigDataResponseBody) *GrayPushMsConfigDataResponse {
	s.Body = v
	return s
}

type ListMqSofamqGroupRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupType  *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNum    *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListMqSofamqGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqGroupRequest) GoString() string {
	return s.String()
}

func (s *ListMqSofamqGroupRequest) SetGroupId(v string) *ListMqSofamqGroupRequest {
	s.GroupId = &v
	return s
}

func (s *ListMqSofamqGroupRequest) SetGroupType(v string) *ListMqSofamqGroupRequest {
	s.GroupType = &v
	return s
}

func (s *ListMqSofamqGroupRequest) SetInstanceId(v string) *ListMqSofamqGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListMqSofamqGroupRequest) SetPageNum(v int64) *ListMqSofamqGroupRequest {
	s.PageNum = &v
	return s
}

func (s *ListMqSofamqGroupRequest) SetPageSize(v int64) *ListMqSofamqGroupRequest {
	s.PageSize = &v
	return s
}

type ListMqSofamqGroupResponseBody struct {
	Data          *ListMqSofamqGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId     *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                            `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                            `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMqSofamqGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListMqSofamqGroupResponseBody) SetData(v *ListMqSofamqGroupResponseBodyData) *ListMqSofamqGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListMqSofamqGroupResponseBody) SetRequestId(v string) *ListMqSofamqGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMqSofamqGroupResponseBody) SetResultCode(v string) *ListMqSofamqGroupResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMqSofamqGroupResponseBody) SetResultMessage(v string) *ListMqSofamqGroupResponseBody {
	s.ResultMessage = &v
	return s
}

type ListMqSofamqGroupResponseBodyData struct {
	Content  []*ListMqSofamqGroupResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	PageNum  *int64                                      `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int64                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMqSofamqGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMqSofamqGroupResponseBodyData) SetContent(v []*ListMqSofamqGroupResponseBodyDataContent) *ListMqSofamqGroupResponseBodyData {
	s.Content = v
	return s
}

func (s *ListMqSofamqGroupResponseBodyData) SetPageNum(v int64) *ListMqSofamqGroupResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListMqSofamqGroupResponseBodyData) SetPageSize(v int64) *ListMqSofamqGroupResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMqSofamqGroupResponseBodyData) SetTotal(v int64) *ListMqSofamqGroupResponseBodyData {
	s.Total = &v
	return s
}

type ListMqSofamqGroupResponseBodyDataContent struct {
	Cluster            *string `json:"Cluster,omitempty" xml:"Cluster,omitempty"`
	DeleteMark         *string `json:"DeleteMark,omitempty" xml:"DeleteMark,omitempty"`
	GmtCreate          *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified        *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GroupId            *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupType          *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Operator           *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	ReadEnable         *bool   `json:"ReadEnable,omitempty" xml:"ReadEnable,omitempty"`
	Remark             *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RetryPerm          *int64  `json:"RetryPerm,omitempty" xml:"RetryPerm,omitempty"`
	RetryReadQueueNum  *int64  `json:"RetryReadQueueNum,omitempty" xml:"RetryReadQueueNum,omitempty"`
	RetryWriteQueueNum *int64  `json:"RetryWriteQueueNum,omitempty" xml:"RetryWriteQueueNum,omitempty"`
	Scope              *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	Version            *int64  `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListMqSofamqGroupResponseBodyDataContent) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqGroupResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListMqSofamqGroupResponseBodyDataContent) SetCluster(v string) *ListMqSofamqGroupResponseBodyDataContent {
	s.Cluster = &v
	return s
}

func (s *ListMqSofamqGroupResponseBodyDataContent) SetDeleteMark(v string) *ListMqSofamqGroupResponseBodyDataContent {
	s.DeleteMark = &v
	return s
}

func (s *ListMqSofamqGroupResponseBodyDataContent) SetGmtCreate(v int64) *ListMqSofamqGroupResponseBodyDataContent {
	s.GmtCreate = &v
	return s
}

func (s *ListMqSofamqGroupResponseBodyDataContent) SetGmtModified(v int64) *ListMqSofamqGroupResponseBodyDataContent {
	s.GmtModified = &v
	return s
}

func (s *ListMqSofamqGroupResponseBodyDataContent) SetGroupId(v string) *ListMqSofamqGroupResponseBodyDataContent {
	s.GroupId = &v
	return s
}

func (s *ListMqSofamqGroupResponseBodyDataContent) SetGroupType(v string) *ListMqSofamqGroupResponseBodyDataContent {
	s.GroupType = &v
	return s
}

func (s *ListMqSofamqGroupResponseBodyDataContent) SetId(v int64) *ListMqSofamqGroupResponseBodyDataContent {
	s.Id = &v
	return s
}

func (s *ListMqSofamqGroupResponseBodyDataContent) SetInstanceId(v string) *ListMqSofamqGroupResponseBodyDataContent {
	s.InstanceId = &v
	return s
}

func (s *ListMqSofamqGroupResponseBodyDataContent) SetOperator(v string) *ListMqSofamqGroupResponseBodyDataContent {
	s.Operator = &v
	return s
}

func (s *ListMqSofamqGroupResponseBodyDataContent) SetReadEnable(v bool) *ListMqSofamqGroupResponseBodyDataContent {
	s.ReadEnable = &v
	return s
}

func (s *ListMqSofamqGroupResponseBodyDataContent) SetRemark(v string) *ListMqSofamqGroupResponseBodyDataContent {
	s.Remark = &v
	return s
}

func (s *ListMqSofamqGroupResponseBodyDataContent) SetRetryPerm(v int64) *ListMqSofamqGroupResponseBodyDataContent {
	s.RetryPerm = &v
	return s
}

func (s *ListMqSofamqGroupResponseBodyDataContent) SetRetryReadQueueNum(v int64) *ListMqSofamqGroupResponseBodyDataContent {
	s.RetryReadQueueNum = &v
	return s
}

func (s *ListMqSofamqGroupResponseBodyDataContent) SetRetryWriteQueueNum(v int64) *ListMqSofamqGroupResponseBodyDataContent {
	s.RetryWriteQueueNum = &v
	return s
}

func (s *ListMqSofamqGroupResponseBodyDataContent) SetScope(v string) *ListMqSofamqGroupResponseBodyDataContent {
	s.Scope = &v
	return s
}

func (s *ListMqSofamqGroupResponseBodyDataContent) SetVersion(v int64) *ListMqSofamqGroupResponseBodyDataContent {
	s.Version = &v
	return s
}

type ListMqSofamqGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMqSofamqGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMqSofamqGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqGroupResponse) GoString() string {
	return s.String()
}

func (s *ListMqSofamqGroupResponse) SetHeaders(v map[string]*string) *ListMqSofamqGroupResponse {
	s.Headers = v
	return s
}

func (s *ListMqSofamqGroupResponse) SetStatusCode(v int32) *ListMqSofamqGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMqSofamqGroupResponse) SetBody(v *ListMqSofamqGroupResponseBody) *ListMqSofamqGroupResponse {
	s.Body = v
	return s
}

type ListMqSofamqTopicRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNum    *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s ListMqSofamqTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqTopicRequest) GoString() string {
	return s.String()
}

func (s *ListMqSofamqTopicRequest) SetInstanceId(v string) *ListMqSofamqTopicRequest {
	s.InstanceId = &v
	return s
}

func (s *ListMqSofamqTopicRequest) SetPageNum(v int64) *ListMqSofamqTopicRequest {
	s.PageNum = &v
	return s
}

func (s *ListMqSofamqTopicRequest) SetPageSize(v int64) *ListMqSofamqTopicRequest {
	s.PageSize = &v
	return s
}

func (s *ListMqSofamqTopicRequest) SetTopic(v string) *ListMqSofamqTopicRequest {
	s.Topic = &v
	return s
}

type ListMqSofamqTopicResponseBody struct {
	Data          *ListMqSofamqTopicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId     *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                            `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                            `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMqSofamqTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqTopicResponseBody) GoString() string {
	return s.String()
}

func (s *ListMqSofamqTopicResponseBody) SetData(v *ListMqSofamqTopicResponseBodyData) *ListMqSofamqTopicResponseBody {
	s.Data = v
	return s
}

func (s *ListMqSofamqTopicResponseBody) SetRequestId(v string) *ListMqSofamqTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMqSofamqTopicResponseBody) SetResultCode(v string) *ListMqSofamqTopicResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMqSofamqTopicResponseBody) SetResultMessage(v string) *ListMqSofamqTopicResponseBody {
	s.ResultMessage = &v
	return s
}

type ListMqSofamqTopicResponseBodyData struct {
	Content  []*ListMqSofamqTopicResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	PageNum  *int64                                      `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int64                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMqSofamqTopicResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqTopicResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMqSofamqTopicResponseBodyData) SetContent(v []*ListMqSofamqTopicResponseBodyDataContent) *ListMqSofamqTopicResponseBodyData {
	s.Content = v
	return s
}

func (s *ListMqSofamqTopicResponseBodyData) SetPageNum(v int64) *ListMqSofamqTopicResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListMqSofamqTopicResponseBodyData) SetPageSize(v int64) *ListMqSofamqTopicResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMqSofamqTopicResponseBodyData) SetTotal(v int64) *ListMqSofamqTopicResponseBodyData {
	s.Total = &v
	return s
}

type ListMqSofamqTopicResponseBodyDataContent struct {
	Cluster       *string `json:"Cluster,omitempty" xml:"Cluster,omitempty"`
	GmtCreate     *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified   *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MessageType   *int64  `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	Operator      *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Perm          *int64  `json:"Perm,omitempty" xml:"Perm,omitempty"`
	ReadQueueNum  *int64  `json:"ReadQueueNum,omitempty" xml:"ReadQueueNum,omitempty"`
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Topic         *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	WriteQueueNum *int64  `json:"WriteQueueNum,omitempty" xml:"WriteQueueNum,omitempty"`
}

func (s ListMqSofamqTopicResponseBodyDataContent) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqTopicResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListMqSofamqTopicResponseBodyDataContent) SetCluster(v string) *ListMqSofamqTopicResponseBodyDataContent {
	s.Cluster = &v
	return s
}

func (s *ListMqSofamqTopicResponseBodyDataContent) SetGmtCreate(v int64) *ListMqSofamqTopicResponseBodyDataContent {
	s.GmtCreate = &v
	return s
}

func (s *ListMqSofamqTopicResponseBodyDataContent) SetGmtModified(v int64) *ListMqSofamqTopicResponseBodyDataContent {
	s.GmtModified = &v
	return s
}

func (s *ListMqSofamqTopicResponseBodyDataContent) SetId(v int64) *ListMqSofamqTopicResponseBodyDataContent {
	s.Id = &v
	return s
}

func (s *ListMqSofamqTopicResponseBodyDataContent) SetInstanceId(v string) *ListMqSofamqTopicResponseBodyDataContent {
	s.InstanceId = &v
	return s
}

func (s *ListMqSofamqTopicResponseBodyDataContent) SetMessageType(v int64) *ListMqSofamqTopicResponseBodyDataContent {
	s.MessageType = &v
	return s
}

func (s *ListMqSofamqTopicResponseBodyDataContent) SetOperator(v string) *ListMqSofamqTopicResponseBodyDataContent {
	s.Operator = &v
	return s
}

func (s *ListMqSofamqTopicResponseBodyDataContent) SetPerm(v int64) *ListMqSofamqTopicResponseBodyDataContent {
	s.Perm = &v
	return s
}

func (s *ListMqSofamqTopicResponseBodyDataContent) SetReadQueueNum(v int64) *ListMqSofamqTopicResponseBodyDataContent {
	s.ReadQueueNum = &v
	return s
}

func (s *ListMqSofamqTopicResponseBodyDataContent) SetRemark(v string) *ListMqSofamqTopicResponseBodyDataContent {
	s.Remark = &v
	return s
}

func (s *ListMqSofamqTopicResponseBodyDataContent) SetTopic(v string) *ListMqSofamqTopicResponseBodyDataContent {
	s.Topic = &v
	return s
}

func (s *ListMqSofamqTopicResponseBodyDataContent) SetWriteQueueNum(v int64) *ListMqSofamqTopicResponseBodyDataContent {
	s.WriteQueueNum = &v
	return s
}

type ListMqSofamqTopicResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMqSofamqTopicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMqSofamqTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqTopicResponse) GoString() string {
	return s.String()
}

func (s *ListMqSofamqTopicResponse) SetHeaders(v map[string]*string) *ListMqSofamqTopicResponse {
	s.Headers = v
	return s
}

func (s *ListMqSofamqTopicResponse) SetStatusCode(v int32) *ListMqSofamqTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMqSofamqTopicResponse) SetBody(v *ListMqSofamqTopicResponseBody) *ListMqSofamqTopicResponse {
	s.Body = v
	return s
}

type ListMqSofamqTraceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNum    *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryItem  *string `json:"QueryItem,omitempty" xml:"QueryItem,omitempty"`
}

func (s ListMqSofamqTraceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqTraceRequest) GoString() string {
	return s.String()
}

func (s *ListMqSofamqTraceRequest) SetInstanceId(v string) *ListMqSofamqTraceRequest {
	s.InstanceId = &v
	return s
}

func (s *ListMqSofamqTraceRequest) SetPageNum(v int64) *ListMqSofamqTraceRequest {
	s.PageNum = &v
	return s
}

func (s *ListMqSofamqTraceRequest) SetPageSize(v int64) *ListMqSofamqTraceRequest {
	s.PageSize = &v
	return s
}

func (s *ListMqSofamqTraceRequest) SetQueryItem(v string) *ListMqSofamqTraceRequest {
	s.QueryItem = &v
	return s
}

type ListMqSofamqTraceResponseBody struct {
	Data          *ListMqSofamqTraceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId     *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                            `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                            `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMqSofamqTraceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqTraceResponseBody) GoString() string {
	return s.String()
}

func (s *ListMqSofamqTraceResponseBody) SetData(v *ListMqSofamqTraceResponseBodyData) *ListMqSofamqTraceResponseBody {
	s.Data = v
	return s
}

func (s *ListMqSofamqTraceResponseBody) SetRequestId(v string) *ListMqSofamqTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMqSofamqTraceResponseBody) SetResultCode(v string) *ListMqSofamqTraceResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMqSofamqTraceResponseBody) SetResultMessage(v string) *ListMqSofamqTraceResponseBody {
	s.ResultMessage = &v
	return s
}

type ListMqSofamqTraceResponseBodyData struct {
	Content  []*ListMqSofamqTraceResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	PageNum  *int64                                      `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int64                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMqSofamqTraceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqTraceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMqSofamqTraceResponseBodyData) SetContent(v []*ListMqSofamqTraceResponseBodyDataContent) *ListMqSofamqTraceResponseBodyData {
	s.Content = v
	return s
}

func (s *ListMqSofamqTraceResponseBodyData) SetPageNum(v int64) *ListMqSofamqTraceResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListMqSofamqTraceResponseBodyData) SetPageSize(v int64) *ListMqSofamqTraceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMqSofamqTraceResponseBodyData) SetTotal(v int64) *ListMqSofamqTraceResponseBodyData {
	s.Total = &v
	return s
}

type ListMqSofamqTraceResponseBodyDataContent struct {
	Cell        *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId       *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	MsgKey      *string `json:"MsgKey,omitempty" xml:"MsgKey,omitempty"`
	QueryId     *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Topic       *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s ListMqSofamqTraceResponseBodyDataContent) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqTraceResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListMqSofamqTraceResponseBodyDataContent) SetCell(v string) *ListMqSofamqTraceResponseBodyDataContent {
	s.Cell = &v
	return s
}

func (s *ListMqSofamqTraceResponseBodyDataContent) SetGmtCreate(v int64) *ListMqSofamqTraceResponseBodyDataContent {
	s.GmtCreate = &v
	return s
}

func (s *ListMqSofamqTraceResponseBodyDataContent) SetGmtModified(v int64) *ListMqSofamqTraceResponseBodyDataContent {
	s.GmtModified = &v
	return s
}

func (s *ListMqSofamqTraceResponseBodyDataContent) SetInstanceId(v string) *ListMqSofamqTraceResponseBodyDataContent {
	s.InstanceId = &v
	return s
}

func (s *ListMqSofamqTraceResponseBodyDataContent) SetMsgId(v string) *ListMqSofamqTraceResponseBodyDataContent {
	s.MsgId = &v
	return s
}

func (s *ListMqSofamqTraceResponseBodyDataContent) SetMsgKey(v string) *ListMqSofamqTraceResponseBodyDataContent {
	s.MsgKey = &v
	return s
}

func (s *ListMqSofamqTraceResponseBodyDataContent) SetQueryId(v string) *ListMqSofamqTraceResponseBodyDataContent {
	s.QueryId = &v
	return s
}

func (s *ListMqSofamqTraceResponseBodyDataContent) SetStatus(v string) *ListMqSofamqTraceResponseBodyDataContent {
	s.Status = &v
	return s
}

func (s *ListMqSofamqTraceResponseBodyDataContent) SetTopic(v string) *ListMqSofamqTraceResponseBodyDataContent {
	s.Topic = &v
	return s
}

type ListMqSofamqTraceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMqSofamqTraceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMqSofamqTraceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqTraceResponse) GoString() string {
	return s.String()
}

func (s *ListMqSofamqTraceResponse) SetHeaders(v map[string]*string) *ListMqSofamqTraceResponse {
	s.Headers = v
	return s
}

func (s *ListMqSofamqTraceResponse) SetStatusCode(v int32) *ListMqSofamqTraceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMqSofamqTraceResponse) SetBody(v *ListMqSofamqTraceResponseBody) *ListMqSofamqTraceResponse {
	s.Body = v
	return s
}

type ListMqSofamqWarnRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNum    *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s ListMqSofamqWarnRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqWarnRequest) GoString() string {
	return s.String()
}

func (s *ListMqSofamqWarnRequest) SetGroupId(v string) *ListMqSofamqWarnRequest {
	s.GroupId = &v
	return s
}

func (s *ListMqSofamqWarnRequest) SetInstanceId(v string) *ListMqSofamqWarnRequest {
	s.InstanceId = &v
	return s
}

func (s *ListMqSofamqWarnRequest) SetPageNum(v int64) *ListMqSofamqWarnRequest {
	s.PageNum = &v
	return s
}

func (s *ListMqSofamqWarnRequest) SetPageSize(v int64) *ListMqSofamqWarnRequest {
	s.PageSize = &v
	return s
}

func (s *ListMqSofamqWarnRequest) SetTopic(v string) *ListMqSofamqWarnRequest {
	s.Topic = &v
	return s
}

type ListMqSofamqWarnResponseBody struct {
	Data          *ListMqSofamqWarnResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId     *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                           `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                           `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMqSofamqWarnResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqWarnResponseBody) GoString() string {
	return s.String()
}

func (s *ListMqSofamqWarnResponseBody) SetData(v *ListMqSofamqWarnResponseBodyData) *ListMqSofamqWarnResponseBody {
	s.Data = v
	return s
}

func (s *ListMqSofamqWarnResponseBody) SetRequestId(v string) *ListMqSofamqWarnResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMqSofamqWarnResponseBody) SetResultCode(v string) *ListMqSofamqWarnResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMqSofamqWarnResponseBody) SetResultMessage(v string) *ListMqSofamqWarnResponseBody {
	s.ResultMessage = &v
	return s
}

type ListMqSofamqWarnResponseBodyData struct {
	Content  []*ListMqSofamqWarnResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	PageNum  *int64                                     `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int64                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMqSofamqWarnResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqWarnResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMqSofamqWarnResponseBodyData) SetContent(v []*ListMqSofamqWarnResponseBodyDataContent) *ListMqSofamqWarnResponseBodyData {
	s.Content = v
	return s
}

func (s *ListMqSofamqWarnResponseBodyData) SetPageNum(v int64) *ListMqSofamqWarnResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListMqSofamqWarnResponseBodyData) SetPageSize(v int64) *ListMqSofamqWarnResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMqSofamqWarnResponseBodyData) SetTotal(v int64) *ListMqSofamqWarnResponseBodyData {
	s.Total = &v
	return s
}

type ListMqSofamqWarnResponseBodyDataContent struct {
	AlertTime   *string `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	Attribute   *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	BlockTime   *int64  `json:"BlockTime,omitempty" xml:"BlockTime,omitempty"`
	Contacts    *string `json:"Contacts,omitempty" xml:"Contacts,omitempty"`
	DelayTime   *int64  `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	Frequency   *int64  `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Operator    *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Threshold   *int64  `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Topic       *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	WarnLevel   *int64  `json:"WarnLevel,omitempty" xml:"WarnLevel,omitempty"`
	WarnStatus  *int64  `json:"WarnStatus,omitempty" xml:"WarnStatus,omitempty"`
	WarnType    *int64  `json:"WarnType,omitempty" xml:"WarnType,omitempty"`
}

func (s ListMqSofamqWarnResponseBodyDataContent) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqWarnResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListMqSofamqWarnResponseBodyDataContent) SetAlertTime(v string) *ListMqSofamqWarnResponseBodyDataContent {
	s.AlertTime = &v
	return s
}

func (s *ListMqSofamqWarnResponseBodyDataContent) SetAttribute(v string) *ListMqSofamqWarnResponseBodyDataContent {
	s.Attribute = &v
	return s
}

func (s *ListMqSofamqWarnResponseBodyDataContent) SetBlockTime(v int64) *ListMqSofamqWarnResponseBodyDataContent {
	s.BlockTime = &v
	return s
}

func (s *ListMqSofamqWarnResponseBodyDataContent) SetContacts(v string) *ListMqSofamqWarnResponseBodyDataContent {
	s.Contacts = &v
	return s
}

func (s *ListMqSofamqWarnResponseBodyDataContent) SetDelayTime(v int64) *ListMqSofamqWarnResponseBodyDataContent {
	s.DelayTime = &v
	return s
}

func (s *ListMqSofamqWarnResponseBodyDataContent) SetFrequency(v int64) *ListMqSofamqWarnResponseBodyDataContent {
	s.Frequency = &v
	return s
}

func (s *ListMqSofamqWarnResponseBodyDataContent) SetGmtCreate(v int64) *ListMqSofamqWarnResponseBodyDataContent {
	s.GmtCreate = &v
	return s
}

func (s *ListMqSofamqWarnResponseBodyDataContent) SetGmtModified(v int64) *ListMqSofamqWarnResponseBodyDataContent {
	s.GmtModified = &v
	return s
}

func (s *ListMqSofamqWarnResponseBodyDataContent) SetGroupId(v string) *ListMqSofamqWarnResponseBodyDataContent {
	s.GroupId = &v
	return s
}

func (s *ListMqSofamqWarnResponseBodyDataContent) SetId(v int64) *ListMqSofamqWarnResponseBodyDataContent {
	s.Id = &v
	return s
}

func (s *ListMqSofamqWarnResponseBodyDataContent) SetInstanceId(v string) *ListMqSofamqWarnResponseBodyDataContent {
	s.InstanceId = &v
	return s
}

func (s *ListMqSofamqWarnResponseBodyDataContent) SetOperator(v string) *ListMqSofamqWarnResponseBodyDataContent {
	s.Operator = &v
	return s
}

func (s *ListMqSofamqWarnResponseBodyDataContent) SetThreshold(v int64) *ListMqSofamqWarnResponseBodyDataContent {
	s.Threshold = &v
	return s
}

func (s *ListMqSofamqWarnResponseBodyDataContent) SetTopic(v string) *ListMqSofamqWarnResponseBodyDataContent {
	s.Topic = &v
	return s
}

func (s *ListMqSofamqWarnResponseBodyDataContent) SetWarnLevel(v int64) *ListMqSofamqWarnResponseBodyDataContent {
	s.WarnLevel = &v
	return s
}

func (s *ListMqSofamqWarnResponseBodyDataContent) SetWarnStatus(v int64) *ListMqSofamqWarnResponseBodyDataContent {
	s.WarnStatus = &v
	return s
}

func (s *ListMqSofamqWarnResponseBodyDataContent) SetWarnType(v int64) *ListMqSofamqWarnResponseBodyDataContent {
	s.WarnType = &v
	return s
}

type ListMqSofamqWarnResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMqSofamqWarnResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMqSofamqWarnResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqWarnResponse) GoString() string {
	return s.String()
}

func (s *ListMqSofamqWarnResponse) SetHeaders(v map[string]*string) *ListMqSofamqWarnResponse {
	s.Headers = v
	return s
}

func (s *ListMqSofamqWarnResponse) SetStatusCode(v int32) *ListMqSofamqWarnResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMqSofamqWarnResponse) SetBody(v *ListMqSofamqWarnResponseBody) *ListMqSofamqWarnResponse {
	s.Body = v
	return s
}

type ListMqSofamqWarnHistoryRequest struct {
	Cell       *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNum    *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	WarnId     *int64  `json:"WarnId,omitempty" xml:"WarnId,omitempty"`
}

func (s ListMqSofamqWarnHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqWarnHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListMqSofamqWarnHistoryRequest) SetCell(v string) *ListMqSofamqWarnHistoryRequest {
	s.Cell = &v
	return s
}

func (s *ListMqSofamqWarnHistoryRequest) SetGroupId(v string) *ListMqSofamqWarnHistoryRequest {
	s.GroupId = &v
	return s
}

func (s *ListMqSofamqWarnHistoryRequest) SetInstanceId(v string) *ListMqSofamqWarnHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *ListMqSofamqWarnHistoryRequest) SetPageNum(v int64) *ListMqSofamqWarnHistoryRequest {
	s.PageNum = &v
	return s
}

func (s *ListMqSofamqWarnHistoryRequest) SetPageSize(v int64) *ListMqSofamqWarnHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListMqSofamqWarnHistoryRequest) SetTopic(v string) *ListMqSofamqWarnHistoryRequest {
	s.Topic = &v
	return s
}

func (s *ListMqSofamqWarnHistoryRequest) SetWarnId(v int64) *ListMqSofamqWarnHistoryRequest {
	s.WarnId = &v
	return s
}

type ListMqSofamqWarnHistoryResponseBody struct {
	Data          *ListMqSofamqWarnHistoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId     *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                  `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                  `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMqSofamqWarnHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqWarnHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListMqSofamqWarnHistoryResponseBody) SetData(v *ListMqSofamqWarnHistoryResponseBodyData) *ListMqSofamqWarnHistoryResponseBody {
	s.Data = v
	return s
}

func (s *ListMqSofamqWarnHistoryResponseBody) SetRequestId(v string) *ListMqSofamqWarnHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMqSofamqWarnHistoryResponseBody) SetResultCode(v string) *ListMqSofamqWarnHistoryResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMqSofamqWarnHistoryResponseBody) SetResultMessage(v string) *ListMqSofamqWarnHistoryResponseBody {
	s.ResultMessage = &v
	return s
}

type ListMqSofamqWarnHistoryResponseBodyData struct {
	Content  []*ListMqSofamqWarnHistoryResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	PageNum  *int64                                            `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int64                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMqSofamqWarnHistoryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqWarnHistoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMqSofamqWarnHistoryResponseBodyData) SetContent(v []*ListMqSofamqWarnHistoryResponseBodyDataContent) *ListMqSofamqWarnHistoryResponseBodyData {
	s.Content = v
	return s
}

func (s *ListMqSofamqWarnHistoryResponseBodyData) SetPageNum(v int64) *ListMqSofamqWarnHistoryResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListMqSofamqWarnHistoryResponseBodyData) SetPageSize(v int64) *ListMqSofamqWarnHistoryResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMqSofamqWarnHistoryResponseBodyData) SetTotal(v int64) *ListMqSofamqWarnHistoryResponseBodyData {
	s.Total = &v
	return s
}

type ListMqSofamqWarnHistoryResponseBodyDataContent struct {
	Cell        *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Topic       *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	WarnId      *int64  `json:"WarnId,omitempty" xml:"WarnId,omitempty"`
	WarnInfo    *string `json:"WarnInfo,omitempty" xml:"WarnInfo,omitempty"`
}

func (s ListMqSofamqWarnHistoryResponseBodyDataContent) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqWarnHistoryResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListMqSofamqWarnHistoryResponseBodyDataContent) SetCell(v string) *ListMqSofamqWarnHistoryResponseBodyDataContent {
	s.Cell = &v
	return s
}

func (s *ListMqSofamqWarnHistoryResponseBodyDataContent) SetGmtCreate(v int64) *ListMqSofamqWarnHistoryResponseBodyDataContent {
	s.GmtCreate = &v
	return s
}

func (s *ListMqSofamqWarnHistoryResponseBodyDataContent) SetGmtModified(v int64) *ListMqSofamqWarnHistoryResponseBodyDataContent {
	s.GmtModified = &v
	return s
}

func (s *ListMqSofamqWarnHistoryResponseBodyDataContent) SetGroupId(v string) *ListMqSofamqWarnHistoryResponseBodyDataContent {
	s.GroupId = &v
	return s
}

func (s *ListMqSofamqWarnHistoryResponseBodyDataContent) SetInstanceId(v string) *ListMqSofamqWarnHistoryResponseBodyDataContent {
	s.InstanceId = &v
	return s
}

func (s *ListMqSofamqWarnHistoryResponseBodyDataContent) SetTopic(v string) *ListMqSofamqWarnHistoryResponseBodyDataContent {
	s.Topic = &v
	return s
}

func (s *ListMqSofamqWarnHistoryResponseBodyDataContent) SetWarnId(v int64) *ListMqSofamqWarnHistoryResponseBodyDataContent {
	s.WarnId = &v
	return s
}

func (s *ListMqSofamqWarnHistoryResponseBodyDataContent) SetWarnInfo(v string) *ListMqSofamqWarnHistoryResponseBodyDataContent {
	s.WarnInfo = &v
	return s
}

type ListMqSofamqWarnHistoryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMqSofamqWarnHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMqSofamqWarnHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMqSofamqWarnHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListMqSofamqWarnHistoryResponse) SetHeaders(v map[string]*string) *ListMqSofamqWarnHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListMqSofamqWarnHistoryResponse) SetStatusCode(v int32) *ListMqSofamqWarnHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMqSofamqWarnHistoryResponse) SetBody(v *ListMqSofamqWarnHistoryResponseBody) *ListMqSofamqWarnHistoryResponse {
	s.Body = v
	return s
}

type LogoutMsRegistryServiceRequest struct {
	InstanceId          *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ServerIpsRepeatList []*string `json:"ServerIpsRepeatList,omitempty" xml:"ServerIpsRepeatList,omitempty" type:"Repeated"`
}

func (s LogoutMsRegistryServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s LogoutMsRegistryServiceRequest) GoString() string {
	return s.String()
}

func (s *LogoutMsRegistryServiceRequest) SetInstanceId(v string) *LogoutMsRegistryServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *LogoutMsRegistryServiceRequest) SetServerIpsRepeatList(v []*string) *LogoutMsRegistryServiceRequest {
	s.ServerIpsRepeatList = v
	return s
}

type LogoutMsRegistryServiceResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s LogoutMsRegistryServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LogoutMsRegistryServiceResponseBody) GoString() string {
	return s.String()
}

func (s *LogoutMsRegistryServiceResponseBody) SetRequestId(v string) *LogoutMsRegistryServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *LogoutMsRegistryServiceResponseBody) SetResultCode(v string) *LogoutMsRegistryServiceResponseBody {
	s.ResultCode = &v
	return s
}

func (s *LogoutMsRegistryServiceResponseBody) SetResultMessage(v string) *LogoutMsRegistryServiceResponseBody {
	s.ResultMessage = &v
	return s
}

type LogoutMsRegistryServiceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LogoutMsRegistryServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LogoutMsRegistryServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s LogoutMsRegistryServiceResponse) GoString() string {
	return s.String()
}

func (s *LogoutMsRegistryServiceResponse) SetHeaders(v map[string]*string) *LogoutMsRegistryServiceResponse {
	s.Headers = v
	return s
}

func (s *LogoutMsRegistryServiceResponse) SetStatusCode(v int32) *LogoutMsRegistryServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *LogoutMsRegistryServiceResponse) SetBody(v *LogoutMsRegistryServiceResponseBody) *LogoutMsRegistryServiceResponse {
	s.Body = v
	return s
}

type PushMsConfigDataRequest struct {
	AttributeId *int64  `json:"AttributeId,omitempty" xml:"AttributeId,omitempty"`
	Cells       *string `json:"Cells,omitempty" xml:"Cells,omitempty"`
	Data        *string `json:"Data,omitempty" xml:"Data,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Operator    *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s PushMsConfigDataRequest) String() string {
	return tea.Prettify(s)
}

func (s PushMsConfigDataRequest) GoString() string {
	return s.String()
}

func (s *PushMsConfigDataRequest) SetAttributeId(v int64) *PushMsConfigDataRequest {
	s.AttributeId = &v
	return s
}

func (s *PushMsConfigDataRequest) SetCells(v string) *PushMsConfigDataRequest {
	s.Cells = &v
	return s
}

func (s *PushMsConfigDataRequest) SetData(v string) *PushMsConfigDataRequest {
	s.Data = &v
	return s
}

func (s *PushMsConfigDataRequest) SetInstanceId(v string) *PushMsConfigDataRequest {
	s.InstanceId = &v
	return s
}

func (s *PushMsConfigDataRequest) SetOperator(v string) *PushMsConfigDataRequest {
	s.Operator = &v
	return s
}

type PushMsConfigDataResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s PushMsConfigDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushMsConfigDataResponseBody) GoString() string {
	return s.String()
}

func (s *PushMsConfigDataResponseBody) SetRequestId(v string) *PushMsConfigDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushMsConfigDataResponseBody) SetResultCode(v string) *PushMsConfigDataResponseBody {
	s.ResultCode = &v
	return s
}

func (s *PushMsConfigDataResponseBody) SetResultMessage(v string) *PushMsConfigDataResponseBody {
	s.ResultMessage = &v
	return s
}

type PushMsConfigDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PushMsConfigDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushMsConfigDataResponse) String() string {
	return tea.Prettify(s)
}

func (s PushMsConfigDataResponse) GoString() string {
	return s.String()
}

func (s *PushMsConfigDataResponse) SetHeaders(v map[string]*string) *PushMsConfigDataResponse {
	s.Headers = v
	return s
}

func (s *PushMsConfigDataResponse) SetStatusCode(v int32) *PushMsConfigDataResponse {
	s.StatusCode = &v
	return s
}

func (s *PushMsConfigDataResponse) SetBody(v *PushMsConfigDataResponseBody) *PushMsConfigDataResponse {
	s.Body = v
	return s
}

type QueryMqSofamqConsumerAccumulateRequest struct {
	Cell       *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	Detail     *bool   `json:"Detail,omitempty" xml:"Detail,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QueryMqSofamqConsumerAccumulateRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqConsumerAccumulateRequest) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqConsumerAccumulateRequest) SetCell(v string) *QueryMqSofamqConsumerAccumulateRequest {
	s.Cell = &v
	return s
}

func (s *QueryMqSofamqConsumerAccumulateRequest) SetDetail(v bool) *QueryMqSofamqConsumerAccumulateRequest {
	s.Detail = &v
	return s
}

func (s *QueryMqSofamqConsumerAccumulateRequest) SetGroupId(v string) *QueryMqSofamqConsumerAccumulateRequest {
	s.GroupId = &v
	return s
}

func (s *QueryMqSofamqConsumerAccumulateRequest) SetInstanceId(v string) *QueryMqSofamqConsumerAccumulateRequest {
	s.InstanceId = &v
	return s
}

type QueryMqSofamqConsumerAccumulateResponseBody struct {
	Data          *QueryMqSofamqConsumerAccumulateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId     *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                          `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                          `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMqSofamqConsumerAccumulateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqConsumerAccumulateResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqConsumerAccumulateResponseBody) SetData(v *QueryMqSofamqConsumerAccumulateResponseBodyData) *QueryMqSofamqConsumerAccumulateResponseBody {
	s.Data = v
	return s
}

func (s *QueryMqSofamqConsumerAccumulateResponseBody) SetRequestId(v string) *QueryMqSofamqConsumerAccumulateResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMqSofamqConsumerAccumulateResponseBody) SetResultCode(v string) *QueryMqSofamqConsumerAccumulateResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMqSofamqConsumerAccumulateResponseBody) SetResultMessage(v string) *QueryMqSofamqConsumerAccumulateResponseBody {
	s.ResultMessage = &v
	return s
}

type QueryMqSofamqConsumerAccumulateResponseBodyData struct {
	ConsumeTps        *string                                                             `json:"ConsumeTps,omitempty" xml:"ConsumeTps,omitempty"`
	DelayTime         *int64                                                              `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	DetailInTopicList []*QueryMqSofamqConsumerAccumulateResponseBodyDataDetailInTopicList `json:"DetailInTopicList,omitempty" xml:"DetailInTopicList,omitempty" type:"Repeated"`
	LastTimestamp     *int64                                                              `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Online            *bool                                                               `json:"Online,omitempty" xml:"Online,omitempty"`
	TotalDiff         *int64                                                              `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
}

func (s QueryMqSofamqConsumerAccumulateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqConsumerAccumulateResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqConsumerAccumulateResponseBodyData) SetConsumeTps(v string) *QueryMqSofamqConsumerAccumulateResponseBodyData {
	s.ConsumeTps = &v
	return s
}

func (s *QueryMqSofamqConsumerAccumulateResponseBodyData) SetDelayTime(v int64) *QueryMqSofamqConsumerAccumulateResponseBodyData {
	s.DelayTime = &v
	return s
}

func (s *QueryMqSofamqConsumerAccumulateResponseBodyData) SetDetailInTopicList(v []*QueryMqSofamqConsumerAccumulateResponseBodyDataDetailInTopicList) *QueryMqSofamqConsumerAccumulateResponseBodyData {
	s.DetailInTopicList = v
	return s
}

func (s *QueryMqSofamqConsumerAccumulateResponseBodyData) SetLastTimestamp(v int64) *QueryMqSofamqConsumerAccumulateResponseBodyData {
	s.LastTimestamp = &v
	return s
}

func (s *QueryMqSofamqConsumerAccumulateResponseBodyData) SetOnline(v bool) *QueryMqSofamqConsumerAccumulateResponseBodyData {
	s.Online = &v
	return s
}

func (s *QueryMqSofamqConsumerAccumulateResponseBodyData) SetTotalDiff(v int64) *QueryMqSofamqConsumerAccumulateResponseBodyData {
	s.TotalDiff = &v
	return s
}

type QueryMqSofamqConsumerAccumulateResponseBodyDataDetailInTopicList struct {
	DelayTime     *int64  `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	LastTimestamp *int64  `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Topic         *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	TotalDiff     *int64  `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
}

func (s QueryMqSofamqConsumerAccumulateResponseBodyDataDetailInTopicList) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqConsumerAccumulateResponseBodyDataDetailInTopicList) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqConsumerAccumulateResponseBodyDataDetailInTopicList) SetDelayTime(v int64) *QueryMqSofamqConsumerAccumulateResponseBodyDataDetailInTopicList {
	s.DelayTime = &v
	return s
}

func (s *QueryMqSofamqConsumerAccumulateResponseBodyDataDetailInTopicList) SetLastTimestamp(v int64) *QueryMqSofamqConsumerAccumulateResponseBodyDataDetailInTopicList {
	s.LastTimestamp = &v
	return s
}

func (s *QueryMqSofamqConsumerAccumulateResponseBodyDataDetailInTopicList) SetTopic(v string) *QueryMqSofamqConsumerAccumulateResponseBodyDataDetailInTopicList {
	s.Topic = &v
	return s
}

func (s *QueryMqSofamqConsumerAccumulateResponseBodyDataDetailInTopicList) SetTotalDiff(v int64) *QueryMqSofamqConsumerAccumulateResponseBodyDataDetailInTopicList {
	s.TotalDiff = &v
	return s
}

type QueryMqSofamqConsumerAccumulateResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMqSofamqConsumerAccumulateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMqSofamqConsumerAccumulateResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqConsumerAccumulateResponse) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqConsumerAccumulateResponse) SetHeaders(v map[string]*string) *QueryMqSofamqConsumerAccumulateResponse {
	s.Headers = v
	return s
}

func (s *QueryMqSofamqConsumerAccumulateResponse) SetStatusCode(v int32) *QueryMqSofamqConsumerAccumulateResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMqSofamqConsumerAccumulateResponse) SetBody(v *QueryMqSofamqConsumerAccumulateResponseBody) *QueryMqSofamqConsumerAccumulateResponse {
	s.Body = v
	return s
}

type QueryMqSofamqConsumerConnectionRequest struct {
	Cell       *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QueryMqSofamqConsumerConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqConsumerConnectionRequest) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqConsumerConnectionRequest) SetCell(v string) *QueryMqSofamqConsumerConnectionRequest {
	s.Cell = &v
	return s
}

func (s *QueryMqSofamqConsumerConnectionRequest) SetGroupId(v string) *QueryMqSofamqConsumerConnectionRequest {
	s.GroupId = &v
	return s
}

func (s *QueryMqSofamqConsumerConnectionRequest) SetInstanceId(v string) *QueryMqSofamqConsumerConnectionRequest {
	s.InstanceId = &v
	return s
}

type QueryMqSofamqConsumerConnectionResponseBody struct {
	Data          *QueryMqSofamqConsumerConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId     *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                          `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                          `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMqSofamqConsumerConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqConsumerConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqConsumerConnectionResponseBody) SetData(v *QueryMqSofamqConsumerConnectionResponseBodyData) *QueryMqSofamqConsumerConnectionResponseBody {
	s.Data = v
	return s
}

func (s *QueryMqSofamqConsumerConnectionResponseBody) SetRequestId(v string) *QueryMqSofamqConsumerConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMqSofamqConsumerConnectionResponseBody) SetResultCode(v string) *QueryMqSofamqConsumerConnectionResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMqSofamqConsumerConnectionResponseBody) SetResultMessage(v string) *QueryMqSofamqConsumerConnectionResponseBody {
	s.ResultMessage = &v
	return s
}

type QueryMqSofamqConsumerConnectionResponseBodyData struct {
	ConnectionList []*QueryMqSofamqConsumerConnectionResponseBodyDataConnectionList `json:"ConnectionList,omitempty" xml:"ConnectionList,omitempty" type:"Repeated"`
}

func (s QueryMqSofamqConsumerConnectionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqConsumerConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqConsumerConnectionResponseBodyData) SetConnectionList(v []*QueryMqSofamqConsumerConnectionResponseBodyDataConnectionList) *QueryMqSofamqConsumerConnectionResponseBodyData {
	s.ConnectionList = v
	return s
}

type QueryMqSofamqConsumerConnectionResponseBodyDataConnectionList struct {
	ClientAddr *string `json:"ClientAddr,omitempty" xml:"ClientAddr,omitempty"`
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	RemoteIp   *string `json:"RemoteIp,omitempty" xml:"RemoteIp,omitempty"`
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s QueryMqSofamqConsumerConnectionResponseBodyDataConnectionList) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqConsumerConnectionResponseBodyDataConnectionList) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqConsumerConnectionResponseBodyDataConnectionList) SetClientAddr(v string) *QueryMqSofamqConsumerConnectionResponseBodyDataConnectionList {
	s.ClientAddr = &v
	return s
}

func (s *QueryMqSofamqConsumerConnectionResponseBodyDataConnectionList) SetClientId(v string) *QueryMqSofamqConsumerConnectionResponseBodyDataConnectionList {
	s.ClientId = &v
	return s
}

func (s *QueryMqSofamqConsumerConnectionResponseBodyDataConnectionList) SetLanguage(v string) *QueryMqSofamqConsumerConnectionResponseBodyDataConnectionList {
	s.Language = &v
	return s
}

func (s *QueryMqSofamqConsumerConnectionResponseBodyDataConnectionList) SetRemoteIp(v string) *QueryMqSofamqConsumerConnectionResponseBodyDataConnectionList {
	s.RemoteIp = &v
	return s
}

func (s *QueryMqSofamqConsumerConnectionResponseBodyDataConnectionList) SetVersion(v string) *QueryMqSofamqConsumerConnectionResponseBodyDataConnectionList {
	s.Version = &v
	return s
}

type QueryMqSofamqConsumerConnectionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMqSofamqConsumerConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMqSofamqConsumerConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqConsumerConnectionResponse) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqConsumerConnectionResponse) SetHeaders(v map[string]*string) *QueryMqSofamqConsumerConnectionResponse {
	s.Headers = v
	return s
}

func (s *QueryMqSofamqConsumerConnectionResponse) SetStatusCode(v int32) *QueryMqSofamqConsumerConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMqSofamqConsumerConnectionResponse) SetBody(v *QueryMqSofamqConsumerConnectionResponseBody) *QueryMqSofamqConsumerConnectionResponse {
	s.Body = v
	return s
}

type QueryMqSofamqConsumerTimespanRequest struct {
	Cell       *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s QueryMqSofamqConsumerTimespanRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqConsumerTimespanRequest) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqConsumerTimespanRequest) SetCell(v string) *QueryMqSofamqConsumerTimespanRequest {
	s.Cell = &v
	return s
}

func (s *QueryMqSofamqConsumerTimespanRequest) SetGroupId(v string) *QueryMqSofamqConsumerTimespanRequest {
	s.GroupId = &v
	return s
}

func (s *QueryMqSofamqConsumerTimespanRequest) SetInstanceId(v string) *QueryMqSofamqConsumerTimespanRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMqSofamqConsumerTimespanRequest) SetTopic(v string) *QueryMqSofamqConsumerTimespanRequest {
	s.Topic = &v
	return s
}

type QueryMqSofamqConsumerTimespanResponseBody struct {
	Data          *QueryMqSofamqConsumerTimespanResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId     *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                        `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                        `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMqSofamqConsumerTimespanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqConsumerTimespanResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqConsumerTimespanResponseBody) SetData(v *QueryMqSofamqConsumerTimespanResponseBodyData) *QueryMqSofamqConsumerTimespanResponseBody {
	s.Data = v
	return s
}

func (s *QueryMqSofamqConsumerTimespanResponseBody) SetRequestId(v string) *QueryMqSofamqConsumerTimespanResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMqSofamqConsumerTimespanResponseBody) SetResultCode(v string) *QueryMqSofamqConsumerTimespanResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMqSofamqConsumerTimespanResponseBody) SetResultMessage(v string) *QueryMqSofamqConsumerTimespanResponseBody {
	s.ResultMessage = &v
	return s
}

type QueryMqSofamqConsumerTimespanResponseBodyData struct {
	ConsumeTimestamp *int64  `json:"ConsumeTimestamp,omitempty" xml:"ConsumeTimestamp,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxTimestamp     *int64  `json:"MaxTimestamp,omitempty" xml:"MaxTimestamp,omitempty"`
	MinTimestamp     *int64  `json:"MinTimestamp,omitempty" xml:"MinTimestamp,omitempty"`
	Topic            *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s QueryMqSofamqConsumerTimespanResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqConsumerTimespanResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqConsumerTimespanResponseBodyData) SetConsumeTimestamp(v int64) *QueryMqSofamqConsumerTimespanResponseBodyData {
	s.ConsumeTimestamp = &v
	return s
}

func (s *QueryMqSofamqConsumerTimespanResponseBodyData) SetInstanceId(v string) *QueryMqSofamqConsumerTimespanResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *QueryMqSofamqConsumerTimespanResponseBodyData) SetMaxTimestamp(v int64) *QueryMqSofamqConsumerTimespanResponseBodyData {
	s.MaxTimestamp = &v
	return s
}

func (s *QueryMqSofamqConsumerTimespanResponseBodyData) SetMinTimestamp(v int64) *QueryMqSofamqConsumerTimespanResponseBodyData {
	s.MinTimestamp = &v
	return s
}

func (s *QueryMqSofamqConsumerTimespanResponseBodyData) SetTopic(v string) *QueryMqSofamqConsumerTimespanResponseBodyData {
	s.Topic = &v
	return s
}

type QueryMqSofamqConsumerTimespanResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMqSofamqConsumerTimespanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMqSofamqConsumerTimespanResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqConsumerTimespanResponse) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqConsumerTimespanResponse) SetHeaders(v map[string]*string) *QueryMqSofamqConsumerTimespanResponse {
	s.Headers = v
	return s
}

func (s *QueryMqSofamqConsumerTimespanResponse) SetStatusCode(v int32) *QueryMqSofamqConsumerTimespanResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMqSofamqConsumerTimespanResponse) SetBody(v *QueryMqSofamqConsumerTimespanResponseBody) *QueryMqSofamqConsumerTimespanResponse {
	s.Body = v
	return s
}

type QueryMqSofamqDLQMessageByGroupIdRequest struct {
	BeginTime  *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	Cell       *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNum    *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryMqSofamqDLQMessageByGroupIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqDLQMessageByGroupIdRequest) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqDLQMessageByGroupIdRequest) SetBeginTime(v int64) *QueryMqSofamqDLQMessageByGroupIdRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdRequest) SetCell(v string) *QueryMqSofamqDLQMessageByGroupIdRequest {
	s.Cell = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdRequest) SetEndTime(v int64) *QueryMqSofamqDLQMessageByGroupIdRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdRequest) SetGroupId(v string) *QueryMqSofamqDLQMessageByGroupIdRequest {
	s.GroupId = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdRequest) SetInstanceId(v string) *QueryMqSofamqDLQMessageByGroupIdRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdRequest) SetPageNum(v int64) *QueryMqSofamqDLQMessageByGroupIdRequest {
	s.PageNum = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdRequest) SetPageSize(v int64) *QueryMqSofamqDLQMessageByGroupIdRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdRequest) SetTaskId(v string) *QueryMqSofamqDLQMessageByGroupIdRequest {
	s.TaskId = &v
	return s
}

type QueryMqSofamqDLQMessageByGroupIdResponseBody struct {
	Data          *QueryMqSofamqDLQMessageByGroupIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                           `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                           `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMqSofamqDLQMessageByGroupIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqDLQMessageByGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBody) SetData(v *QueryMqSofamqDLQMessageByGroupIdResponseBodyData) *QueryMqSofamqDLQMessageByGroupIdResponseBody {
	s.Data = v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBody) SetRequestId(v string) *QueryMqSofamqDLQMessageByGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBody) SetResultCode(v string) *QueryMqSofamqDLQMessageByGroupIdResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBody) SetResultMessage(v string) *QueryMqSofamqDLQMessageByGroupIdResponseBody {
	s.ResultMessage = &v
	return s
}

type QueryMqSofamqDLQMessageByGroupIdResponseBodyData struct {
	Content  []*QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	PageNum  *int64                                                     `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int64                                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskId   *string                                                    `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Total    *int64                                                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryMqSofamqDLQMessageByGroupIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqDLQMessageByGroupIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBodyData) SetContent(v []*QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent) *QueryMqSofamqDLQMessageByGroupIdResponseBodyData {
	s.Content = v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBodyData) SetPageNum(v int64) *QueryMqSofamqDLQMessageByGroupIdResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBodyData) SetPageSize(v int64) *QueryMqSofamqDLQMessageByGroupIdResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBodyData) SetTaskId(v string) *QueryMqSofamqDLQMessageByGroupIdResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBodyData) SetTotal(v int64) *QueryMqSofamqDLQMessageByGroupIdResponseBodyData {
	s.Total = &v
	return s
}

type QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent struct {
	Body           *string                                                                `json:"Body,omitempty" xml:"Body,omitempty"`
	BodyCrc        *int64                                                                 `json:"BodyCrc,omitempty" xml:"BodyCrc,omitempty"`
	BornHost       *string                                                                `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	BornTimestamp  *int64                                                                 `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	InstanceId     *string                                                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId          *string                                                                `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	PropertyList   []*QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContentPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	ReconsumeTimes *int64                                                                 `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	StoreHost      *string                                                                `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	StoreSize      *int64                                                                 `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	StoreTimestamp *int64                                                                 `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	Topic          *string                                                                `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent) SetBody(v string) *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent {
	s.Body = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent) SetBodyCrc(v int64) *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent {
	s.BodyCrc = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent) SetBornHost(v string) *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent {
	s.BornHost = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent) SetBornTimestamp(v int64) *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent {
	s.BornTimestamp = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent) SetInstanceId(v string) *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent {
	s.InstanceId = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent) SetMsgId(v string) *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent {
	s.MsgId = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent) SetPropertyList(v []*QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContentPropertyList) *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent {
	s.PropertyList = v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent) SetReconsumeTimes(v int64) *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent {
	s.ReconsumeTimes = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent) SetStoreHost(v string) *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent {
	s.StoreHost = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent) SetStoreSize(v int64) *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent {
	s.StoreSize = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent) SetStoreTimestamp(v int64) *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent {
	s.StoreTimestamp = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent) SetTopic(v string) *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContent {
	s.Topic = &v
	return s
}

type QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContentPropertyList struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContentPropertyList) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContentPropertyList) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContentPropertyList) SetName(v string) *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContentPropertyList {
	s.Name = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContentPropertyList) SetValue(v string) *QueryMqSofamqDLQMessageByGroupIdResponseBodyDataContentPropertyList {
	s.Value = &v
	return s
}

type QueryMqSofamqDLQMessageByGroupIdResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMqSofamqDLQMessageByGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMqSofamqDLQMessageByGroupIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqDLQMessageByGroupIdResponse) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponse) SetHeaders(v map[string]*string) *QueryMqSofamqDLQMessageByGroupIdResponse {
	s.Headers = v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponse) SetStatusCode(v int32) *QueryMqSofamqDLQMessageByGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMqSofamqDLQMessageByGroupIdResponse) SetBody(v *QueryMqSofamqDLQMessageByGroupIdResponseBody) *QueryMqSofamqDLQMessageByGroupIdResponse {
	s.Body = v
	return s
}

type QueryMqSofamqGroupSubDetailRequest struct {
	Cell       *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QueryMqSofamqGroupSubDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqGroupSubDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqGroupSubDetailRequest) SetCell(v string) *QueryMqSofamqGroupSubDetailRequest {
	s.Cell = &v
	return s
}

func (s *QueryMqSofamqGroupSubDetailRequest) SetGroupId(v string) *QueryMqSofamqGroupSubDetailRequest {
	s.GroupId = &v
	return s
}

func (s *QueryMqSofamqGroupSubDetailRequest) SetInstanceId(v string) *QueryMqSofamqGroupSubDetailRequest {
	s.InstanceId = &v
	return s
}

type QueryMqSofamqGroupSubDetailResponseBody struct {
	Data          *QueryMqSofamqGroupSubDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId     *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                      `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                      `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMqSofamqGroupSubDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqGroupSubDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqGroupSubDetailResponseBody) SetData(v *QueryMqSofamqGroupSubDetailResponseBodyData) *QueryMqSofamqGroupSubDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryMqSofamqGroupSubDetailResponseBody) SetRequestId(v string) *QueryMqSofamqGroupSubDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMqSofamqGroupSubDetailResponseBody) SetResultCode(v string) *QueryMqSofamqGroupSubDetailResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMqSofamqGroupSubDetailResponseBody) SetResultMessage(v string) *QueryMqSofamqGroupSubDetailResponseBody {
	s.ResultMessage = &v
	return s
}

type QueryMqSofamqGroupSubDetailResponseBodyData struct {
	GroupId              *string                                                            `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MessageModel         *string                                                            `json:"MessageModel,omitempty" xml:"MessageModel,omitempty"`
	Online               *bool                                                              `json:"Online,omitempty" xml:"Online,omitempty"`
	SubscriptionDataList []*QueryMqSofamqGroupSubDetailResponseBodyDataSubscriptionDataList `json:"SubscriptionDataList,omitempty" xml:"SubscriptionDataList,omitempty" type:"Repeated"`
}

func (s QueryMqSofamqGroupSubDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqGroupSubDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqGroupSubDetailResponseBodyData) SetGroupId(v string) *QueryMqSofamqGroupSubDetailResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *QueryMqSofamqGroupSubDetailResponseBodyData) SetMessageModel(v string) *QueryMqSofamqGroupSubDetailResponseBodyData {
	s.MessageModel = &v
	return s
}

func (s *QueryMqSofamqGroupSubDetailResponseBodyData) SetOnline(v bool) *QueryMqSofamqGroupSubDetailResponseBodyData {
	s.Online = &v
	return s
}

func (s *QueryMqSofamqGroupSubDetailResponseBodyData) SetSubscriptionDataList(v []*QueryMqSofamqGroupSubDetailResponseBodyDataSubscriptionDataList) *QueryMqSofamqGroupSubDetailResponseBodyData {
	s.SubscriptionDataList = v
	return s
}

type QueryMqSofamqGroupSubDetailResponseBodyDataSubscriptionDataList struct {
	Online    *bool   `json:"Online,omitempty" xml:"Online,omitempty"`
	SubString *string `json:"SubString,omitempty" xml:"SubString,omitempty"`
	Topic     *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s QueryMqSofamqGroupSubDetailResponseBodyDataSubscriptionDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqGroupSubDetailResponseBodyDataSubscriptionDataList) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqGroupSubDetailResponseBodyDataSubscriptionDataList) SetOnline(v bool) *QueryMqSofamqGroupSubDetailResponseBodyDataSubscriptionDataList {
	s.Online = &v
	return s
}

func (s *QueryMqSofamqGroupSubDetailResponseBodyDataSubscriptionDataList) SetSubString(v string) *QueryMqSofamqGroupSubDetailResponseBodyDataSubscriptionDataList {
	s.SubString = &v
	return s
}

func (s *QueryMqSofamqGroupSubDetailResponseBodyDataSubscriptionDataList) SetTopic(v string) *QueryMqSofamqGroupSubDetailResponseBodyDataSubscriptionDataList {
	s.Topic = &v
	return s
}

type QueryMqSofamqGroupSubDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMqSofamqGroupSubDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMqSofamqGroupSubDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqGroupSubDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqGroupSubDetailResponse) SetHeaders(v map[string]*string) *QueryMqSofamqGroupSubDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryMqSofamqGroupSubDetailResponse) SetStatusCode(v int32) *QueryMqSofamqGroupSubDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMqSofamqGroupSubDetailResponse) SetBody(v *QueryMqSofamqGroupSubDetailResponseBody) *QueryMqSofamqGroupSubDetailResponse {
	s.Body = v
	return s
}

type QueryMqSofamqMessageByKeyRequest struct {
	Cell       *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Key        *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s QueryMqSofamqMessageByKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqMessageByKeyRequest) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqMessageByKeyRequest) SetCell(v string) *QueryMqSofamqMessageByKeyRequest {
	s.Cell = &v
	return s
}

func (s *QueryMqSofamqMessageByKeyRequest) SetInstanceId(v string) *QueryMqSofamqMessageByKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMqSofamqMessageByKeyRequest) SetKey(v string) *QueryMqSofamqMessageByKeyRequest {
	s.Key = &v
	return s
}

func (s *QueryMqSofamqMessageByKeyRequest) SetTopic(v string) *QueryMqSofamqMessageByKeyRequest {
	s.Topic = &v
	return s
}

type QueryMqSofamqMessageByKeyResponseBody struct {
	Data          []*QueryMqSofamqMessageByKeyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId     *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                      `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                      `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMqSofamqMessageByKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqMessageByKeyResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqMessageByKeyResponseBody) SetData(v []*QueryMqSofamqMessageByKeyResponseBodyData) *QueryMqSofamqMessageByKeyResponseBody {
	s.Data = v
	return s
}

func (s *QueryMqSofamqMessageByKeyResponseBody) SetRequestId(v string) *QueryMqSofamqMessageByKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMqSofamqMessageByKeyResponseBody) SetResultCode(v string) *QueryMqSofamqMessageByKeyResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMqSofamqMessageByKeyResponseBody) SetResultMessage(v string) *QueryMqSofamqMessageByKeyResponseBody {
	s.ResultMessage = &v
	return s
}

type QueryMqSofamqMessageByKeyResponseBodyData struct {
	Body           *string                                                  `json:"Body,omitempty" xml:"Body,omitempty"`
	BodyCrc        *int64                                                   `json:"BodyCrc,omitempty" xml:"BodyCrc,omitempty"`
	BornHost       *string                                                  `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	BornTimestamp  *int64                                                   `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	InstanceId     *string                                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId          *string                                                  `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	PropertyList   []*QueryMqSofamqMessageByKeyResponseBodyDataPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	ReconsumeTimes *int64                                                   `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	StoreHost      *string                                                  `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	StoreSize      *int64                                                   `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	StoreTimestamp *int64                                                   `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	Topic          *string                                                  `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s QueryMqSofamqMessageByKeyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqMessageByKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqMessageByKeyResponseBodyData) SetBody(v string) *QueryMqSofamqMessageByKeyResponseBodyData {
	s.Body = &v
	return s
}

func (s *QueryMqSofamqMessageByKeyResponseBodyData) SetBodyCrc(v int64) *QueryMqSofamqMessageByKeyResponseBodyData {
	s.BodyCrc = &v
	return s
}

func (s *QueryMqSofamqMessageByKeyResponseBodyData) SetBornHost(v string) *QueryMqSofamqMessageByKeyResponseBodyData {
	s.BornHost = &v
	return s
}

func (s *QueryMqSofamqMessageByKeyResponseBodyData) SetBornTimestamp(v int64) *QueryMqSofamqMessageByKeyResponseBodyData {
	s.BornTimestamp = &v
	return s
}

func (s *QueryMqSofamqMessageByKeyResponseBodyData) SetInstanceId(v string) *QueryMqSofamqMessageByKeyResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *QueryMqSofamqMessageByKeyResponseBodyData) SetMsgId(v string) *QueryMqSofamqMessageByKeyResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *QueryMqSofamqMessageByKeyResponseBodyData) SetPropertyList(v []*QueryMqSofamqMessageByKeyResponseBodyDataPropertyList) *QueryMqSofamqMessageByKeyResponseBodyData {
	s.PropertyList = v
	return s
}

func (s *QueryMqSofamqMessageByKeyResponseBodyData) SetReconsumeTimes(v int64) *QueryMqSofamqMessageByKeyResponseBodyData {
	s.ReconsumeTimes = &v
	return s
}

func (s *QueryMqSofamqMessageByKeyResponseBodyData) SetStoreHost(v string) *QueryMqSofamqMessageByKeyResponseBodyData {
	s.StoreHost = &v
	return s
}

func (s *QueryMqSofamqMessageByKeyResponseBodyData) SetStoreSize(v int64) *QueryMqSofamqMessageByKeyResponseBodyData {
	s.StoreSize = &v
	return s
}

func (s *QueryMqSofamqMessageByKeyResponseBodyData) SetStoreTimestamp(v int64) *QueryMqSofamqMessageByKeyResponseBodyData {
	s.StoreTimestamp = &v
	return s
}

func (s *QueryMqSofamqMessageByKeyResponseBodyData) SetTopic(v string) *QueryMqSofamqMessageByKeyResponseBodyData {
	s.Topic = &v
	return s
}

type QueryMqSofamqMessageByKeyResponseBodyDataPropertyList struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryMqSofamqMessageByKeyResponseBodyDataPropertyList) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqMessageByKeyResponseBodyDataPropertyList) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqMessageByKeyResponseBodyDataPropertyList) SetName(v string) *QueryMqSofamqMessageByKeyResponseBodyDataPropertyList {
	s.Name = &v
	return s
}

func (s *QueryMqSofamqMessageByKeyResponseBodyDataPropertyList) SetValue(v string) *QueryMqSofamqMessageByKeyResponseBodyDataPropertyList {
	s.Value = &v
	return s
}

type QueryMqSofamqMessageByKeyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMqSofamqMessageByKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMqSofamqMessageByKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqMessageByKeyResponse) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqMessageByKeyResponse) SetHeaders(v map[string]*string) *QueryMqSofamqMessageByKeyResponse {
	s.Headers = v
	return s
}

func (s *QueryMqSofamqMessageByKeyResponse) SetStatusCode(v int32) *QueryMqSofamqMessageByKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMqSofamqMessageByKeyResponse) SetBody(v *QueryMqSofamqMessageByKeyResponseBody) *QueryMqSofamqMessageByKeyResponse {
	s.Body = v
	return s
}

type QueryMqSofamqMessageByTopicRequest struct {
	BeginTime  *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	Cell       *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNum    *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s QueryMqSofamqMessageByTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqMessageByTopicRequest) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqMessageByTopicRequest) SetBeginTime(v int64) *QueryMqSofamqMessageByTopicRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicRequest) SetCell(v string) *QueryMqSofamqMessageByTopicRequest {
	s.Cell = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicRequest) SetEndTime(v int64) *QueryMqSofamqMessageByTopicRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicRequest) SetInstanceId(v string) *QueryMqSofamqMessageByTopicRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicRequest) SetPageNum(v int64) *QueryMqSofamqMessageByTopicRequest {
	s.PageNum = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicRequest) SetPageSize(v int64) *QueryMqSofamqMessageByTopicRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicRequest) SetTaskId(v string) *QueryMqSofamqMessageByTopicRequest {
	s.TaskId = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicRequest) SetTopic(v string) *QueryMqSofamqMessageByTopicRequest {
	s.Topic = &v
	return s
}

type QueryMqSofamqMessageByTopicResponseBody struct {
	Data          *QueryMqSofamqMessageByTopicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId     *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                      `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                      `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMqSofamqMessageByTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqMessageByTopicResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqMessageByTopicResponseBody) SetData(v *QueryMqSofamqMessageByTopicResponseBodyData) *QueryMqSofamqMessageByTopicResponseBody {
	s.Data = v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponseBody) SetRequestId(v string) *QueryMqSofamqMessageByTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponseBody) SetResultCode(v string) *QueryMqSofamqMessageByTopicResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponseBody) SetResultMessage(v string) *QueryMqSofamqMessageByTopicResponseBody {
	s.ResultMessage = &v
	return s
}

type QueryMqSofamqMessageByTopicResponseBodyData struct {
	Content  []*QueryMqSofamqMessageByTopicResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	PageNum  *int64                                                `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int64                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskId   *string                                               `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Total    *int64                                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryMqSofamqMessageByTopicResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqMessageByTopicResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqMessageByTopicResponseBodyData) SetContent(v []*QueryMqSofamqMessageByTopicResponseBodyDataContent) *QueryMqSofamqMessageByTopicResponseBodyData {
	s.Content = v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponseBodyData) SetPageNum(v int64) *QueryMqSofamqMessageByTopicResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponseBodyData) SetPageSize(v int64) *QueryMqSofamqMessageByTopicResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponseBodyData) SetTaskId(v string) *QueryMqSofamqMessageByTopicResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponseBodyData) SetTotal(v int64) *QueryMqSofamqMessageByTopicResponseBodyData {
	s.Total = &v
	return s
}

type QueryMqSofamqMessageByTopicResponseBodyDataContent struct {
	Body           *string                                                           `json:"Body,omitempty" xml:"Body,omitempty"`
	BodyCrc        *int64                                                            `json:"BodyCrc,omitempty" xml:"BodyCrc,omitempty"`
	BornHost       *string                                                           `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	BornTimestamp  *int64                                                            `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	InstanceId     *string                                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId          *string                                                           `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	PropertyList   []*QueryMqSofamqMessageByTopicResponseBodyDataContentPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	ReconsumeTimes *int64                                                            `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	StoreHost      *string                                                           `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	StoreSize      *int64                                                            `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	StoreTimestamp *int64                                                            `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	Topic          *string                                                           `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s QueryMqSofamqMessageByTopicResponseBodyDataContent) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqMessageByTopicResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqMessageByTopicResponseBodyDataContent) SetBody(v string) *QueryMqSofamqMessageByTopicResponseBodyDataContent {
	s.Body = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponseBodyDataContent) SetBodyCrc(v int64) *QueryMqSofamqMessageByTopicResponseBodyDataContent {
	s.BodyCrc = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponseBodyDataContent) SetBornHost(v string) *QueryMqSofamqMessageByTopicResponseBodyDataContent {
	s.BornHost = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponseBodyDataContent) SetBornTimestamp(v int64) *QueryMqSofamqMessageByTopicResponseBodyDataContent {
	s.BornTimestamp = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponseBodyDataContent) SetInstanceId(v string) *QueryMqSofamqMessageByTopicResponseBodyDataContent {
	s.InstanceId = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponseBodyDataContent) SetMsgId(v string) *QueryMqSofamqMessageByTopicResponseBodyDataContent {
	s.MsgId = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponseBodyDataContent) SetPropertyList(v []*QueryMqSofamqMessageByTopicResponseBodyDataContentPropertyList) *QueryMqSofamqMessageByTopicResponseBodyDataContent {
	s.PropertyList = v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponseBodyDataContent) SetReconsumeTimes(v int64) *QueryMqSofamqMessageByTopicResponseBodyDataContent {
	s.ReconsumeTimes = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponseBodyDataContent) SetStoreHost(v string) *QueryMqSofamqMessageByTopicResponseBodyDataContent {
	s.StoreHost = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponseBodyDataContent) SetStoreSize(v int64) *QueryMqSofamqMessageByTopicResponseBodyDataContent {
	s.StoreSize = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponseBodyDataContent) SetStoreTimestamp(v int64) *QueryMqSofamqMessageByTopicResponseBodyDataContent {
	s.StoreTimestamp = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponseBodyDataContent) SetTopic(v string) *QueryMqSofamqMessageByTopicResponseBodyDataContent {
	s.Topic = &v
	return s
}

type QueryMqSofamqMessageByTopicResponseBodyDataContentPropertyList struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryMqSofamqMessageByTopicResponseBodyDataContentPropertyList) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqMessageByTopicResponseBodyDataContentPropertyList) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqMessageByTopicResponseBodyDataContentPropertyList) SetName(v string) *QueryMqSofamqMessageByTopicResponseBodyDataContentPropertyList {
	s.Name = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponseBodyDataContentPropertyList) SetValue(v string) *QueryMqSofamqMessageByTopicResponseBodyDataContentPropertyList {
	s.Value = &v
	return s
}

type QueryMqSofamqMessageByTopicResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMqSofamqMessageByTopicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMqSofamqMessageByTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqMessageByTopicResponse) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqMessageByTopicResponse) SetHeaders(v map[string]*string) *QueryMqSofamqMessageByTopicResponse {
	s.Headers = v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponse) SetStatusCode(v int32) *QueryMqSofamqMessageByTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMqSofamqMessageByTopicResponse) SetBody(v *QueryMqSofamqMessageByTopicResponseBody) *QueryMqSofamqMessageByTopicResponse {
	s.Body = v
	return s
}

type QueryMqSofamqTraceByMsgKeyRequest struct {
	BeginTime  *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	Cell       *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgKey     *string `json:"MsgKey,omitempty" xml:"MsgKey,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s QueryMqSofamqTraceByMsgKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqTraceByMsgKeyRequest) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqTraceByMsgKeyRequest) SetBeginTime(v int64) *QueryMqSofamqTraceByMsgKeyRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryMqSofamqTraceByMsgKeyRequest) SetCell(v string) *QueryMqSofamqTraceByMsgKeyRequest {
	s.Cell = &v
	return s
}

func (s *QueryMqSofamqTraceByMsgKeyRequest) SetEndTime(v int64) *QueryMqSofamqTraceByMsgKeyRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMqSofamqTraceByMsgKeyRequest) SetInstanceId(v string) *QueryMqSofamqTraceByMsgKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMqSofamqTraceByMsgKeyRequest) SetMsgKey(v string) *QueryMqSofamqTraceByMsgKeyRequest {
	s.MsgKey = &v
	return s
}

func (s *QueryMqSofamqTraceByMsgKeyRequest) SetTopic(v string) *QueryMqSofamqTraceByMsgKeyRequest {
	s.Topic = &v
	return s
}

type QueryMqSofamqTraceByMsgKeyResponseBody struct {
	QueryId       *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMqSofamqTraceByMsgKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqTraceByMsgKeyResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqTraceByMsgKeyResponseBody) SetQueryId(v string) *QueryMqSofamqTraceByMsgKeyResponseBody {
	s.QueryId = &v
	return s
}

func (s *QueryMqSofamqTraceByMsgKeyResponseBody) SetRequestId(v string) *QueryMqSofamqTraceByMsgKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMqSofamqTraceByMsgKeyResponseBody) SetResultCode(v string) *QueryMqSofamqTraceByMsgKeyResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMqSofamqTraceByMsgKeyResponseBody) SetResultMessage(v string) *QueryMqSofamqTraceByMsgKeyResponseBody {
	s.ResultMessage = &v
	return s
}

type QueryMqSofamqTraceByMsgKeyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMqSofamqTraceByMsgKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMqSofamqTraceByMsgKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqTraceByMsgKeyResponse) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqTraceByMsgKeyResponse) SetHeaders(v map[string]*string) *QueryMqSofamqTraceByMsgKeyResponse {
	s.Headers = v
	return s
}

func (s *QueryMqSofamqTraceByMsgKeyResponse) SetStatusCode(v int32) *QueryMqSofamqTraceByMsgKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMqSofamqTraceByMsgKeyResponse) SetBody(v *QueryMqSofamqTraceByMsgKeyResponseBody) *QueryMqSofamqTraceByMsgKeyResponse {
	s.Body = v
	return s
}

type QueryMqSofamqTraceByTopicRequest struct {
	BeginTime  *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	Cell       *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s QueryMqSofamqTraceByTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqTraceByTopicRequest) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqTraceByTopicRequest) SetBeginTime(v int64) *QueryMqSofamqTraceByTopicRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryMqSofamqTraceByTopicRequest) SetCell(v string) *QueryMqSofamqTraceByTopicRequest {
	s.Cell = &v
	return s
}

func (s *QueryMqSofamqTraceByTopicRequest) SetEndTime(v int64) *QueryMqSofamqTraceByTopicRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMqSofamqTraceByTopicRequest) SetInstanceId(v string) *QueryMqSofamqTraceByTopicRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMqSofamqTraceByTopicRequest) SetTopic(v string) *QueryMqSofamqTraceByTopicRequest {
	s.Topic = &v
	return s
}

type QueryMqSofamqTraceByTopicResponseBody struct {
	QueryId       *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMqSofamqTraceByTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqTraceByTopicResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqTraceByTopicResponseBody) SetQueryId(v string) *QueryMqSofamqTraceByTopicResponseBody {
	s.QueryId = &v
	return s
}

func (s *QueryMqSofamqTraceByTopicResponseBody) SetRequestId(v string) *QueryMqSofamqTraceByTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMqSofamqTraceByTopicResponseBody) SetResultCode(v string) *QueryMqSofamqTraceByTopicResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMqSofamqTraceByTopicResponseBody) SetResultMessage(v string) *QueryMqSofamqTraceByTopicResponseBody {
	s.ResultMessage = &v
	return s
}

type QueryMqSofamqTraceByTopicResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMqSofamqTraceByTopicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMqSofamqTraceByTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMqSofamqTraceByTopicResponse) GoString() string {
	return s.String()
}

func (s *QueryMqSofamqTraceByTopicResponse) SetHeaders(v map[string]*string) *QueryMqSofamqTraceByTopicResponse {
	s.Headers = v
	return s
}

func (s *QueryMqSofamqTraceByTopicResponse) SetStatusCode(v int32) *QueryMqSofamqTraceByTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMqSofamqTraceByTopicResponse) SetBody(v *QueryMqSofamqTraceByTopicResponseBody) *QueryMqSofamqTraceByTopicResponse {
	s.Body = v
	return s
}

type QueryMsConfigAttributesRequest struct {
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceId    *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s QueryMsConfigAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigAttributesRequest) GoString() string {
	return s.String()
}

func (s *QueryMsConfigAttributesRequest) SetAppName(v string) *QueryMsConfigAttributesRequest {
	s.AppName = &v
	return s
}

func (s *QueryMsConfigAttributesRequest) SetAttributeName(v string) *QueryMsConfigAttributesRequest {
	s.AttributeName = &v
	return s
}

func (s *QueryMsConfigAttributesRequest) SetInstanceId(v string) *QueryMsConfigAttributesRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMsConfigAttributesRequest) SetRegion(v string) *QueryMsConfigAttributesRequest {
	s.Region = &v
	return s
}

func (s *QueryMsConfigAttributesRequest) SetResourceId(v string) *QueryMsConfigAttributesRequest {
	s.ResourceId = &v
	return s
}

type QueryMsConfigAttributesResponseBody struct {
	Attribute     *QueryMsConfigAttributesResponseBodyAttribute `json:"Attribute,omitempty" xml:"Attribute,omitempty" type:"Struct"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                       `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                       `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMsConfigAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMsConfigAttributesResponseBody) SetAttribute(v *QueryMsConfigAttributesResponseBodyAttribute) *QueryMsConfigAttributesResponseBody {
	s.Attribute = v
	return s
}

func (s *QueryMsConfigAttributesResponseBody) SetRequestId(v string) *QueryMsConfigAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMsConfigAttributesResponseBody) SetResultCode(v string) *QueryMsConfigAttributesResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMsConfigAttributesResponseBody) SetResultMessage(v string) *QueryMsConfigAttributesResponseBody {
	s.ResultMessage = &v
	return s
}

type QueryMsConfigAttributesResponseBodyAttribute struct {
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	Desc          *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QueryMsConfigAttributesResponseBodyAttribute) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigAttributesResponseBodyAttribute) GoString() string {
	return s.String()
}

func (s *QueryMsConfigAttributesResponseBodyAttribute) SetAttributeName(v string) *QueryMsConfigAttributesResponseBodyAttribute {
	s.AttributeName = &v
	return s
}

func (s *QueryMsConfigAttributesResponseBodyAttribute) SetDesc(v string) *QueryMsConfigAttributesResponseBodyAttribute {
	s.Desc = &v
	return s
}

func (s *QueryMsConfigAttributesResponseBodyAttribute) SetId(v int64) *QueryMsConfigAttributesResponseBodyAttribute {
	s.Id = &v
	return s
}

func (s *QueryMsConfigAttributesResponseBodyAttribute) SetInstanceId(v string) *QueryMsConfigAttributesResponseBodyAttribute {
	s.InstanceId = &v
	return s
}

type QueryMsConfigAttributesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMsConfigAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMsConfigAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigAttributesResponse) GoString() string {
	return s.String()
}

func (s *QueryMsConfigAttributesResponse) SetHeaders(v map[string]*string) *QueryMsConfigAttributesResponse {
	s.Headers = v
	return s
}

func (s *QueryMsConfigAttributesResponse) SetStatusCode(v int32) *QueryMsConfigAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMsConfigAttributesResponse) SetBody(v *QueryMsConfigAttributesResponseBody) *QueryMsConfigAttributesResponse {
	s.Body = v
	return s
}

type QueryMsConfigClientValuesRequest struct {
	AttributeId *int64  `json:"AttributeId,omitempty" xml:"AttributeId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ips         *string `json:"Ips,omitempty" xml:"Ips,omitempty"`
}

func (s QueryMsConfigClientValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigClientValuesRequest) GoString() string {
	return s.String()
}

func (s *QueryMsConfigClientValuesRequest) SetAttributeId(v int64) *QueryMsConfigClientValuesRequest {
	s.AttributeId = &v
	return s
}

func (s *QueryMsConfigClientValuesRequest) SetInstanceId(v string) *QueryMsConfigClientValuesRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMsConfigClientValuesRequest) SetIps(v string) *QueryMsConfigClientValuesRequest {
	s.Ips = &v
	return s
}

type QueryMsConfigClientValuesResponseBody struct {
	Clients       []*QueryMsConfigClientValuesResponseBodyClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                         `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                         `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMsConfigClientValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigClientValuesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMsConfigClientValuesResponseBody) SetClients(v []*QueryMsConfigClientValuesResponseBodyClients) *QueryMsConfigClientValuesResponseBody {
	s.Clients = v
	return s
}

func (s *QueryMsConfigClientValuesResponseBody) SetRequestId(v string) *QueryMsConfigClientValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMsConfigClientValuesResponseBody) SetResultCode(v string) *QueryMsConfigClientValuesResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMsConfigClientValuesResponseBody) SetResultMessage(v string) *QueryMsConfigClientValuesResponseBody {
	s.ResultMessage = &v
	return s
}

type QueryMsConfigClientValuesResponseBodyClients struct {
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Ip      *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMsConfigClientValuesResponseBodyClients) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigClientValuesResponseBodyClients) GoString() string {
	return s.String()
}

func (s *QueryMsConfigClientValuesResponseBodyClients) SetData(v string) *QueryMsConfigClientValuesResponseBodyClients {
	s.Data = &v
	return s
}

func (s *QueryMsConfigClientValuesResponseBodyClients) SetIp(v string) *QueryMsConfigClientValuesResponseBodyClients {
	s.Ip = &v
	return s
}

func (s *QueryMsConfigClientValuesResponseBodyClients) SetSuccess(v bool) *QueryMsConfigClientValuesResponseBodyClients {
	s.Success = &v
	return s
}

type QueryMsConfigClientValuesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMsConfigClientValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMsConfigClientValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigClientValuesResponse) GoString() string {
	return s.String()
}

func (s *QueryMsConfigClientValuesResponse) SetHeaders(v map[string]*string) *QueryMsConfigClientValuesResponse {
	s.Headers = v
	return s
}

func (s *QueryMsConfigClientValuesResponse) SetStatusCode(v int32) *QueryMsConfigClientValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMsConfigClientValuesResponse) SetBody(v *QueryMsConfigClientValuesResponseBody) *QueryMsConfigClientValuesResponse {
	s.Body = v
	return s
}

type QueryMsConfigClientsRequest struct {
	AttributeId *int64  `json:"AttributeId,omitempty" xml:"AttributeId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNum     *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryMsConfigClientsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigClientsRequest) GoString() string {
	return s.String()
}

func (s *QueryMsConfigClientsRequest) SetAttributeId(v int64) *QueryMsConfigClientsRequest {
	s.AttributeId = &v
	return s
}

func (s *QueryMsConfigClientsRequest) SetInstanceId(v string) *QueryMsConfigClientsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMsConfigClientsRequest) SetPageNum(v int64) *QueryMsConfigClientsRequest {
	s.PageNum = &v
	return s
}

func (s *QueryMsConfigClientsRequest) SetPageSize(v string) *QueryMsConfigClientsRequest {
	s.PageSize = &v
	return s
}

type QueryMsConfigClientsResponseBody struct {
	Clients       []*QueryMsConfigClientsResponseBodyClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	PageNum       *int64                                     `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize      *int64                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                    `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                    `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	TotalCount    *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryMsConfigClientsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigClientsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMsConfigClientsResponseBody) SetClients(v []*QueryMsConfigClientsResponseBodyClients) *QueryMsConfigClientsResponseBody {
	s.Clients = v
	return s
}

func (s *QueryMsConfigClientsResponseBody) SetPageNum(v int64) *QueryMsConfigClientsResponseBody {
	s.PageNum = &v
	return s
}

func (s *QueryMsConfigClientsResponseBody) SetPageSize(v int64) *QueryMsConfigClientsResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryMsConfigClientsResponseBody) SetRequestId(v string) *QueryMsConfigClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMsConfigClientsResponseBody) SetResultCode(v string) *QueryMsConfigClientsResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMsConfigClientsResponseBody) SetResultMessage(v string) *QueryMsConfigClientsResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryMsConfigClientsResponseBody) SetTotalCount(v int64) *QueryMsConfigClientsResponseBody {
	s.TotalCount = &v
	return s
}

type QueryMsConfigClientsResponseBodyClients struct {
	Cell     *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	Data     *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Ip       *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	PushData *string `json:"PushData,omitempty" xml:"PushData,omitempty"`
}

func (s QueryMsConfigClientsResponseBodyClients) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigClientsResponseBodyClients) GoString() string {
	return s.String()
}

func (s *QueryMsConfigClientsResponseBodyClients) SetCell(v string) *QueryMsConfigClientsResponseBodyClients {
	s.Cell = &v
	return s
}

func (s *QueryMsConfigClientsResponseBodyClients) SetData(v string) *QueryMsConfigClientsResponseBodyClients {
	s.Data = &v
	return s
}

func (s *QueryMsConfigClientsResponseBodyClients) SetIp(v string) *QueryMsConfigClientsResponseBodyClients {
	s.Ip = &v
	return s
}

func (s *QueryMsConfigClientsResponseBodyClients) SetPushData(v string) *QueryMsConfigClientsResponseBodyClients {
	s.PushData = &v
	return s
}

type QueryMsConfigClientsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMsConfigClientsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMsConfigClientsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigClientsResponse) GoString() string {
	return s.String()
}

func (s *QueryMsConfigClientsResponse) SetHeaders(v map[string]*string) *QueryMsConfigClientsResponse {
	s.Headers = v
	return s
}

func (s *QueryMsConfigClientsResponse) SetStatusCode(v int32) *QueryMsConfigClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMsConfigClientsResponse) SetBody(v *QueryMsConfigClientsResponseBody) *QueryMsConfigClientsResponse {
	s.Body = v
	return s
}

type QueryMsConfigDataRequest struct {
	AttributeId *int64  `json:"AttributeId,omitempty" xml:"AttributeId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QueryMsConfigDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigDataRequest) GoString() string {
	return s.String()
}

func (s *QueryMsConfigDataRequest) SetAttributeId(v int64) *QueryMsConfigDataRequest {
	s.AttributeId = &v
	return s
}

func (s *QueryMsConfigDataRequest) SetInstanceId(v string) *QueryMsConfigDataRequest {
	s.InstanceId = &v
	return s
}

type QueryMsConfigDataResponseBody struct {
	QueryResult   []*QueryMsConfigDataResponseBodyQueryResult `json:"QueryResult,omitempty" xml:"QueryResult,omitempty" type:"Repeated"`
	RequestId     *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                     `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                     `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMsConfigDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMsConfigDataResponseBody) SetQueryResult(v []*QueryMsConfigDataResponseBodyQueryResult) *QueryMsConfigDataResponseBody {
	s.QueryResult = v
	return s
}

func (s *QueryMsConfigDataResponseBody) SetRequestId(v string) *QueryMsConfigDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMsConfigDataResponseBody) SetResultCode(v string) *QueryMsConfigDataResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMsConfigDataResponseBody) SetResultMessage(v string) *QueryMsConfigDataResponseBody {
	s.ResultMessage = &v
	return s
}

type QueryMsConfigDataResponseBodyQueryResult struct {
	Cell *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s QueryMsConfigDataResponseBodyQueryResult) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigDataResponseBodyQueryResult) GoString() string {
	return s.String()
}

func (s *QueryMsConfigDataResponseBodyQueryResult) SetCell(v string) *QueryMsConfigDataResponseBodyQueryResult {
	s.Cell = &v
	return s
}

func (s *QueryMsConfigDataResponseBodyQueryResult) SetData(v string) *QueryMsConfigDataResponseBodyQueryResult {
	s.Data = &v
	return s
}

type QueryMsConfigDataResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMsConfigDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMsConfigDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigDataResponse) GoString() string {
	return s.String()
}

func (s *QueryMsConfigDataResponse) SetHeaders(v map[string]*string) *QueryMsConfigDataResponse {
	s.Headers = v
	return s
}

func (s *QueryMsConfigDataResponse) SetStatusCode(v int32) *QueryMsConfigDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMsConfigDataResponse) SetBody(v *QueryMsConfigDataResponseBody) *QueryMsConfigDataResponse {
	s.Body = v
	return s
}

type QueryMsConfigResourcesRequest struct {
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNum    *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryMsConfigResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigResourcesRequest) GoString() string {
	return s.String()
}

func (s *QueryMsConfigResourcesRequest) SetAppName(v string) *QueryMsConfigResourcesRequest {
	s.AppName = &v
	return s
}

func (s *QueryMsConfigResourcesRequest) SetInstanceId(v string) *QueryMsConfigResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMsConfigResourcesRequest) SetKeyword(v string) *QueryMsConfigResourcesRequest {
	s.Keyword = &v
	return s
}

func (s *QueryMsConfigResourcesRequest) SetPageNum(v int64) *QueryMsConfigResourcesRequest {
	s.PageNum = &v
	return s
}

func (s *QueryMsConfigResourcesRequest) SetPageSize(v string) *QueryMsConfigResourcesRequest {
	s.PageSize = &v
	return s
}

type QueryMsConfigResourcesResponseBody struct {
	PageNum       *int64                                         `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize      *int64                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources     []*QueryMsConfigResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	ResultCode    *string                                        `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                        `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	TotalCount    *int64                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryMsConfigResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMsConfigResourcesResponseBody) SetPageNum(v int64) *QueryMsConfigResourcesResponseBody {
	s.PageNum = &v
	return s
}

func (s *QueryMsConfigResourcesResponseBody) SetPageSize(v int64) *QueryMsConfigResourcesResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryMsConfigResourcesResponseBody) SetRequestId(v string) *QueryMsConfigResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMsConfigResourcesResponseBody) SetResources(v []*QueryMsConfigResourcesResponseBodyResources) *QueryMsConfigResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *QueryMsConfigResourcesResponseBody) SetResultCode(v string) *QueryMsConfigResourcesResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMsConfigResourcesResponseBody) SetResultMessage(v string) *QueryMsConfigResourcesResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryMsConfigResourcesResponseBody) SetTotalCount(v int64) *QueryMsConfigResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type QueryMsConfigResourcesResponseBodyResources struct {
	AppName    *string                                                  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Attributes []*QueryMsConfigResourcesResponseBodyResourcesAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Desc       *string                                                  `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Id         *int64                                                   `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId *string                                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region     *string                                                  `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceId *string                                                  `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s QueryMsConfigResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *QueryMsConfigResourcesResponseBodyResources) SetAppName(v string) *QueryMsConfigResourcesResponseBodyResources {
	s.AppName = &v
	return s
}

func (s *QueryMsConfigResourcesResponseBodyResources) SetAttributes(v []*QueryMsConfigResourcesResponseBodyResourcesAttributes) *QueryMsConfigResourcesResponseBodyResources {
	s.Attributes = v
	return s
}

func (s *QueryMsConfigResourcesResponseBodyResources) SetDesc(v string) *QueryMsConfigResourcesResponseBodyResources {
	s.Desc = &v
	return s
}

func (s *QueryMsConfigResourcesResponseBodyResources) SetId(v int64) *QueryMsConfigResourcesResponseBodyResources {
	s.Id = &v
	return s
}

func (s *QueryMsConfigResourcesResponseBodyResources) SetInstanceId(v string) *QueryMsConfigResourcesResponseBodyResources {
	s.InstanceId = &v
	return s
}

func (s *QueryMsConfigResourcesResponseBodyResources) SetRegion(v string) *QueryMsConfigResourcesResponseBodyResources {
	s.Region = &v
	return s
}

func (s *QueryMsConfigResourcesResponseBodyResources) SetResourceId(v string) *QueryMsConfigResourcesResponseBodyResources {
	s.ResourceId = &v
	return s
}

type QueryMsConfigResourcesResponseBodyResourcesAttributes struct {
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	Desc          *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QueryMsConfigResourcesResponseBodyResourcesAttributes) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigResourcesResponseBodyResourcesAttributes) GoString() string {
	return s.String()
}

func (s *QueryMsConfigResourcesResponseBodyResourcesAttributes) SetAttributeName(v string) *QueryMsConfigResourcesResponseBodyResourcesAttributes {
	s.AttributeName = &v
	return s
}

func (s *QueryMsConfigResourcesResponseBodyResourcesAttributes) SetDesc(v string) *QueryMsConfigResourcesResponseBodyResourcesAttributes {
	s.Desc = &v
	return s
}

func (s *QueryMsConfigResourcesResponseBodyResourcesAttributes) SetId(v int64) *QueryMsConfigResourcesResponseBodyResourcesAttributes {
	s.Id = &v
	return s
}

func (s *QueryMsConfigResourcesResponseBodyResourcesAttributes) SetInstanceId(v string) *QueryMsConfigResourcesResponseBodyResourcesAttributes {
	s.InstanceId = &v
	return s
}

type QueryMsConfigResourcesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMsConfigResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMsConfigResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMsConfigResourcesResponse) GoString() string {
	return s.String()
}

func (s *QueryMsConfigResourcesResponse) SetHeaders(v map[string]*string) *QueryMsConfigResourcesResponse {
	s.Headers = v
	return s
}

func (s *QueryMsConfigResourcesResponse) SetStatusCode(v int32) *QueryMsConfigResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMsConfigResourcesResponse) SetBody(v *QueryMsConfigResourcesResponseBody) *QueryMsConfigResourcesResponse {
	s.Body = v
	return s
}

type QueryRMSMetricsRequest struct {
	ContentType      *string                        `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	DsId             *string                        `json:"DsId,omitempty" xml:"DsId,omitempty"`
	End              *int64                         `json:"End,omitempty" xml:"End,omitempty"`
	FieldsRepeatList []*string                      `json:"FieldsRepeatList,omitempty" xml:"FieldsRepeatList,omitempty" type:"Repeated"`
	PeriodType       *string                        `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	Plugin           *string                        `json:"Plugin,omitempty" xml:"Plugin,omitempty"`
	Start            *int64                         `json:"Start,omitempty" xml:"Start,omitempty"`
	Where            []*QueryRMSMetricsRequestWhere `json:"Where,omitempty" xml:"Where,omitempty" type:"Repeated"`
	WorkspaceName    *string                        `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryRMSMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRMSMetricsRequest) GoString() string {
	return s.String()
}

func (s *QueryRMSMetricsRequest) SetContentType(v string) *QueryRMSMetricsRequest {
	s.ContentType = &v
	return s
}

func (s *QueryRMSMetricsRequest) SetDsId(v string) *QueryRMSMetricsRequest {
	s.DsId = &v
	return s
}

func (s *QueryRMSMetricsRequest) SetEnd(v int64) *QueryRMSMetricsRequest {
	s.End = &v
	return s
}

func (s *QueryRMSMetricsRequest) SetFieldsRepeatList(v []*string) *QueryRMSMetricsRequest {
	s.FieldsRepeatList = v
	return s
}

func (s *QueryRMSMetricsRequest) SetPeriodType(v string) *QueryRMSMetricsRequest {
	s.PeriodType = &v
	return s
}

func (s *QueryRMSMetricsRequest) SetPlugin(v string) *QueryRMSMetricsRequest {
	s.Plugin = &v
	return s
}

func (s *QueryRMSMetricsRequest) SetStart(v int64) *QueryRMSMetricsRequest {
	s.Start = &v
	return s
}

func (s *QueryRMSMetricsRequest) SetWhere(v []*QueryRMSMetricsRequestWhere) *QueryRMSMetricsRequest {
	s.Where = v
	return s
}

func (s *QueryRMSMetricsRequest) SetWorkspaceName(v string) *QueryRMSMetricsRequest {
	s.WorkspaceName = &v
	return s
}

type QueryRMSMetricsRequestWhere struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryRMSMetricsRequestWhere) String() string {
	return tea.Prettify(s)
}

func (s QueryRMSMetricsRequestWhere) GoString() string {
	return s.String()
}

func (s *QueryRMSMetricsRequestWhere) SetKey(v string) *QueryRMSMetricsRequestWhere {
	s.Key = &v
	return s
}

func (s *QueryRMSMetricsRequestWhere) SetValue(v string) *QueryRMSMetricsRequestWhere {
	s.Value = &v
	return s
}

type QueryRMSMetricsResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result        *string `json:"Result,omitempty" xml:"Result,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryRMSMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRMSMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRMSMetricsResponseBody) SetRequestId(v string) *QueryRMSMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRMSMetricsResponseBody) SetResult(v string) *QueryRMSMetricsResponseBody {
	s.Result = &v
	return s
}

func (s *QueryRMSMetricsResponseBody) SetResultCode(v string) *QueryRMSMetricsResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryRMSMetricsResponseBody) SetResultMessage(v string) *QueryRMSMetricsResponseBody {
	s.ResultMessage = &v
	return s
}

type QueryRMSMetricsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryRMSMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRMSMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRMSMetricsResponse) GoString() string {
	return s.String()
}

func (s *QueryRMSMetricsResponse) SetHeaders(v map[string]*string) *QueryRMSMetricsResponse {
	s.Headers = v
	return s
}

func (s *QueryRMSMetricsResponse) SetStatusCode(v int32) *QueryRMSMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRMSMetricsResponse) SetBody(v *QueryRMSMetricsResponseBody) *QueryRMSMetricsResponse {
	s.Body = v
	return s
}

type QueryRMSUnifiedAlarmEventRequest struct {
	AlarmLevel            *int64   `json:"AlarmLevel,omitempty" xml:"AlarmLevel,omitempty"`
	AlarmRuleId           *int64   `json:"AlarmRuleId,omitempty" xml:"AlarmRuleId,omitempty"`
	AlarmRuleUuid         *string  `json:"AlarmRuleUuid,omitempty" xml:"AlarmRuleUuid,omitempty"`
	AlarmStackInfoJsonStr *string  `json:"AlarmStackInfoJsonStr,omitempty" xml:"AlarmStackInfoJsonStr,omitempty"`
	AlarmStatusRepeatList []*int64 `json:"AlarmStatusRepeatList,omitempty" xml:"AlarmStatusRepeatList,omitempty" type:"Repeated"`
	AlarmTargetKeyword    *string  `json:"AlarmTargetKeyword,omitempty" xml:"AlarmTargetKeyword,omitempty"`
	AlarmTargetType       *string  `json:"AlarmTargetType,omitempty" xml:"AlarmTargetType,omitempty"`
	CurrentPage           *int64   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime               *int64   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Keyword               *string  `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageSize              *int64   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime             *int64   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                *int64   `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkspaceName         *string  `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryRMSUnifiedAlarmEventRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRMSUnifiedAlarmEventRequest) GoString() string {
	return s.String()
}

func (s *QueryRMSUnifiedAlarmEventRequest) SetAlarmLevel(v int64) *QueryRMSUnifiedAlarmEventRequest {
	s.AlarmLevel = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventRequest) SetAlarmRuleId(v int64) *QueryRMSUnifiedAlarmEventRequest {
	s.AlarmRuleId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventRequest) SetAlarmRuleUuid(v string) *QueryRMSUnifiedAlarmEventRequest {
	s.AlarmRuleUuid = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventRequest) SetAlarmStackInfoJsonStr(v string) *QueryRMSUnifiedAlarmEventRequest {
	s.AlarmStackInfoJsonStr = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventRequest) SetAlarmStatusRepeatList(v []*int64) *QueryRMSUnifiedAlarmEventRequest {
	s.AlarmStatusRepeatList = v
	return s
}

func (s *QueryRMSUnifiedAlarmEventRequest) SetAlarmTargetKeyword(v string) *QueryRMSUnifiedAlarmEventRequest {
	s.AlarmTargetKeyword = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventRequest) SetAlarmTargetType(v string) *QueryRMSUnifiedAlarmEventRequest {
	s.AlarmTargetType = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventRequest) SetCurrentPage(v int64) *QueryRMSUnifiedAlarmEventRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventRequest) SetEndTime(v int64) *QueryRMSUnifiedAlarmEventRequest {
	s.EndTime = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventRequest) SetKeyword(v string) *QueryRMSUnifiedAlarmEventRequest {
	s.Keyword = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventRequest) SetPageSize(v int64) *QueryRMSUnifiedAlarmEventRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventRequest) SetStartTime(v int64) *QueryRMSUnifiedAlarmEventRequest {
	s.StartTime = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventRequest) SetStatus(v int64) *QueryRMSUnifiedAlarmEventRequest {
	s.Status = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventRequest) SetWorkspaceName(v string) *QueryRMSUnifiedAlarmEventRequest {
	s.WorkspaceName = &v
	return s
}

type QueryRMSUnifiedAlarmEventResponseBody struct {
	CurrentPage           *int64                                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize              *int64                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId             *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode            *string                                                       `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage         *string                                                       `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	TotalCount            *int64                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UnifiedAlarmEventList []*QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList `json:"UnifiedAlarmEventList,omitempty" xml:"UnifiedAlarmEventList,omitempty" type:"Repeated"`
}

func (s QueryRMSUnifiedAlarmEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRMSUnifiedAlarmEventResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRMSUnifiedAlarmEventResponseBody) SetCurrentPage(v int64) *QueryRMSUnifiedAlarmEventResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBody) SetPageSize(v int64) *QueryRMSUnifiedAlarmEventResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBody) SetRequestId(v string) *QueryRMSUnifiedAlarmEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBody) SetResultCode(v string) *QueryRMSUnifiedAlarmEventResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBody) SetResultMessage(v string) *QueryRMSUnifiedAlarmEventResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBody) SetTotalCount(v int64) *QueryRMSUnifiedAlarmEventResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBody) SetUnifiedAlarmEventList(v []*QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList) *QueryRMSUnifiedAlarmEventResponseBody {
	s.UnifiedAlarmEventList = v
	return s
}

type QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList struct {
	AlarmContent       *string                                                                       `json:"AlarmContent,omitempty" xml:"AlarmContent,omitempty"`
	AlarmEventId       *string                                                                       `json:"AlarmEventId,omitempty" xml:"AlarmEventId,omitempty"`
	AlarmHistoryId     *int64                                                                        `json:"AlarmHistoryId,omitempty" xml:"AlarmHistoryId,omitempty"`
	AlarmHistoryList   []*QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList `json:"AlarmHistoryList,omitempty" xml:"AlarmHistoryList,omitempty" type:"Repeated"`
	AlarmLevel         *int64                                                                        `json:"AlarmLevel,omitempty" xml:"AlarmLevel,omitempty"`
	AlarmRecoverTime   *string                                                                       `json:"AlarmRecoverTime,omitempty" xml:"AlarmRecoverTime,omitempty"`
	AlarmRuleId        *int64                                                                        `json:"AlarmRuleId,omitempty" xml:"AlarmRuleId,omitempty"`
	AlarmStartTime     *string                                                                       `json:"AlarmStartTime,omitempty" xml:"AlarmStartTime,omitempty"`
	AlarmTarget        *string                                                                       `json:"AlarmTarget,omitempty" xml:"AlarmTarget,omitempty"`
	AlarmTargetDeleted *bool                                                                         `json:"AlarmTargetDeleted,omitempty" xml:"AlarmTargetDeleted,omitempty"`
	AlarmTargetType    *string                                                                       `json:"AlarmTargetType,omitempty" xml:"AlarmTargetType,omitempty"`
	AlarmType          *string                                                                       `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	AlarmUrl           *string                                                                       `json:"AlarmUrl,omitempty" xml:"AlarmUrl,omitempty"`
	AlarmUrlWithDomain *string                                                                       `json:"AlarmUrlWithDomain,omitempty" xml:"AlarmUrlWithDomain,omitempty"`
	DataSourceName     *string                                                                       `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	SilenceRemainTime  *int64                                                                        `json:"SilenceRemainTime,omitempty" xml:"SilenceRemainTime,omitempty"`
	Status             *int64                                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList) String() string {
	return tea.Prettify(s)
}

func (s QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList) GoString() string {
	return s.String()
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList) SetAlarmContent(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList {
	s.AlarmContent = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList) SetAlarmEventId(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList {
	s.AlarmEventId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList) SetAlarmHistoryId(v int64) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList {
	s.AlarmHistoryId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList) SetAlarmHistoryList(v []*QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList {
	s.AlarmHistoryList = v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList) SetAlarmLevel(v int64) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList {
	s.AlarmLevel = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList) SetAlarmRecoverTime(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList {
	s.AlarmRecoverTime = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList) SetAlarmRuleId(v int64) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList {
	s.AlarmRuleId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList) SetAlarmStartTime(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList {
	s.AlarmStartTime = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList) SetAlarmTarget(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList {
	s.AlarmTarget = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList) SetAlarmTargetDeleted(v bool) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList {
	s.AlarmTargetDeleted = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList) SetAlarmTargetType(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList {
	s.AlarmTargetType = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList) SetAlarmType(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList {
	s.AlarmType = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList) SetAlarmUrl(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList {
	s.AlarmUrl = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList) SetAlarmUrlWithDomain(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList {
	s.AlarmUrlWithDomain = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList) SetDataSourceName(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList {
	s.DataSourceName = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList) SetSilenceRemainTime(v int64) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList {
	s.SilenceRemainTime = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList) SetStatus(v int64) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventList {
	s.Status = &v
	return s
}

type QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList struct {
	AlarmContent       *string `json:"AlarmContent,omitempty" xml:"AlarmContent,omitempty"`
	AlarmDetail        *string `json:"AlarmDetail,omitempty" xml:"AlarmDetail,omitempty"`
	AlarmLevel         *int64  `json:"AlarmLevel,omitempty" xml:"AlarmLevel,omitempty"`
	AlarmRuleId        *int64  `json:"AlarmRuleId,omitempty" xml:"AlarmRuleId,omitempty"`
	AlarmRuleName      *string `json:"AlarmRuleName,omitempty" xml:"AlarmRuleName,omitempty"`
	AlarmTarget        *string `json:"AlarmTarget,omitempty" xml:"AlarmTarget,omitempty"`
	AlarmTargetDeleted *bool   `json:"AlarmTargetDeleted,omitempty" xml:"AlarmTargetDeleted,omitempty"`
	AlarmTargetType    *string `json:"AlarmTargetType,omitempty" xml:"AlarmTargetType,omitempty"`
	AlarmTime          *int64  `json:"AlarmTime,omitempty" xml:"AlarmTime,omitempty"`
	AlarmType          *string `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	AlarmUrl           *string `json:"AlarmUrl,omitempty" xml:"AlarmUrl,omitempty"`
	AlarmUrlWithDomain *string `json:"AlarmUrlWithDomain,omitempty" xml:"AlarmUrlWithDomain,omitempty"`
	DataSourceName     *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	EventId            *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified        *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Status             *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId           *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UniqueIdentity     *string `json:"UniqueIdentity,omitempty" xml:"UniqueIdentity,omitempty"`
	WorkspaceId        *int64  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) String() string {
	return tea.Prettify(s)
}

func (s QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) GoString() string {
	return s.String()
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetAlarmContent(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.AlarmContent = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetAlarmDetail(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.AlarmDetail = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetAlarmLevel(v int64) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.AlarmLevel = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetAlarmRuleId(v int64) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.AlarmRuleId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetAlarmRuleName(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.AlarmRuleName = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetAlarmTarget(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.AlarmTarget = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetAlarmTargetDeleted(v bool) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.AlarmTargetDeleted = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetAlarmTargetType(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.AlarmTargetType = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetAlarmTime(v int64) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.AlarmTime = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetAlarmType(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.AlarmType = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetAlarmUrl(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.AlarmUrl = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetAlarmUrlWithDomain(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.AlarmUrlWithDomain = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetDataSourceName(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.DataSourceName = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetEventId(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.EventId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetGmtCreate(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.GmtCreate = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetGmtModified(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.GmtModified = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetId(v int64) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.Id = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetStatus(v int64) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.Status = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetTenantId(v int64) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.TenantId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetUniqueIdentity(v string) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.UniqueIdentity = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList) SetWorkspaceId(v int64) *QueryRMSUnifiedAlarmEventResponseBodyUnifiedAlarmEventListAlarmHistoryList {
	s.WorkspaceId = &v
	return s
}

type QueryRMSUnifiedAlarmEventResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryRMSUnifiedAlarmEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRMSUnifiedAlarmEventResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRMSUnifiedAlarmEventResponse) GoString() string {
	return s.String()
}

func (s *QueryRMSUnifiedAlarmEventResponse) SetHeaders(v map[string]*string) *QueryRMSUnifiedAlarmEventResponse {
	s.Headers = v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponse) SetStatusCode(v int32) *QueryRMSUnifiedAlarmEventResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRMSUnifiedAlarmEventResponse) SetBody(v *QueryRMSUnifiedAlarmEventResponseBody) *QueryRMSUnifiedAlarmEventResponse {
	s.Body = v
	return s
}

type QueryRMSUnifiedAlarmNotifyHistoryRequest struct {
	AlarmRuleId           *int64                                                      `json:"AlarmRuleId,omitempty" xml:"AlarmRuleId,omitempty"`
	AlarmStackInfoJsonStr *string                                                     `json:"AlarmStackInfoJsonStr,omitempty" xml:"AlarmStackInfoJsonStr,omitempty"`
	AlarmStatus           *int64                                                      `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	AlarmSubscribers      []*QueryRMSUnifiedAlarmNotifyHistoryRequestAlarmSubscribers `json:"AlarmSubscribers,omitempty" xml:"AlarmSubscribers,omitempty" type:"Repeated"`
	Channel               *string                                                     `json:"Channel,omitempty" xml:"Channel,omitempty"`
	CurrentPage           *int64                                                      `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime               *int64                                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EventId               *string                                                     `json:"EventId,omitempty" xml:"EventId,omitempty"`
	Keyword               *string                                                     `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageSize              *int64                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime             *int64                                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                *int64                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	Subscriber            *string                                                     `json:"Subscriber,omitempty" xml:"Subscriber,omitempty"`
	WorkspaceName         *string                                                     `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryRMSUnifiedAlarmNotifyHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRMSUnifiedAlarmNotifyHistoryRequest) GoString() string {
	return s.String()
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryRequest) SetAlarmRuleId(v int64) *QueryRMSUnifiedAlarmNotifyHistoryRequest {
	s.AlarmRuleId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryRequest) SetAlarmStackInfoJsonStr(v string) *QueryRMSUnifiedAlarmNotifyHistoryRequest {
	s.AlarmStackInfoJsonStr = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryRequest) SetAlarmStatus(v int64) *QueryRMSUnifiedAlarmNotifyHistoryRequest {
	s.AlarmStatus = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryRequest) SetAlarmSubscribers(v []*QueryRMSUnifiedAlarmNotifyHistoryRequestAlarmSubscribers) *QueryRMSUnifiedAlarmNotifyHistoryRequest {
	s.AlarmSubscribers = v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryRequest) SetChannel(v string) *QueryRMSUnifiedAlarmNotifyHistoryRequest {
	s.Channel = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryRequest) SetCurrentPage(v int64) *QueryRMSUnifiedAlarmNotifyHistoryRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryRequest) SetEndTime(v int64) *QueryRMSUnifiedAlarmNotifyHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryRequest) SetEventId(v string) *QueryRMSUnifiedAlarmNotifyHistoryRequest {
	s.EventId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryRequest) SetKeyword(v string) *QueryRMSUnifiedAlarmNotifyHistoryRequest {
	s.Keyword = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryRequest) SetPageSize(v int64) *QueryRMSUnifiedAlarmNotifyHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryRequest) SetStartTime(v int64) *QueryRMSUnifiedAlarmNotifyHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryRequest) SetStatus(v int64) *QueryRMSUnifiedAlarmNotifyHistoryRequest {
	s.Status = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryRequest) SetSubscriber(v string) *QueryRMSUnifiedAlarmNotifyHistoryRequest {
	s.Subscriber = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryRequest) SetWorkspaceName(v string) *QueryRMSUnifiedAlarmNotifyHistoryRequest {
	s.WorkspaceName = &v
	return s
}

type QueryRMSUnifiedAlarmNotifyHistoryRequestAlarmSubscribers struct {
	Subscriber       *string `json:"Subscriber,omitempty" xml:"Subscriber,omitempty"`
	SubscriberName   *string `json:"SubscriberName,omitempty" xml:"SubscriberName,omitempty"`
	SubscriberSource *string `json:"SubscriberSource,omitempty" xml:"SubscriberSource,omitempty"`
	SubscriberType   *string `json:"SubscriberType,omitempty" xml:"SubscriberType,omitempty"`
	SubscriberUuid   *string `json:"SubscriberUuid,omitempty" xml:"SubscriberUuid,omitempty"`
}

func (s QueryRMSUnifiedAlarmNotifyHistoryRequestAlarmSubscribers) String() string {
	return tea.Prettify(s)
}

func (s QueryRMSUnifiedAlarmNotifyHistoryRequestAlarmSubscribers) GoString() string {
	return s.String()
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryRequestAlarmSubscribers) SetSubscriber(v string) *QueryRMSUnifiedAlarmNotifyHistoryRequestAlarmSubscribers {
	s.Subscriber = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryRequestAlarmSubscribers) SetSubscriberName(v string) *QueryRMSUnifiedAlarmNotifyHistoryRequestAlarmSubscribers {
	s.SubscriberName = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryRequestAlarmSubscribers) SetSubscriberSource(v string) *QueryRMSUnifiedAlarmNotifyHistoryRequestAlarmSubscribers {
	s.SubscriberSource = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryRequestAlarmSubscribers) SetSubscriberType(v string) *QueryRMSUnifiedAlarmNotifyHistoryRequestAlarmSubscribers {
	s.SubscriberType = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryRequestAlarmSubscribers) SetSubscriberUuid(v string) *QueryRMSUnifiedAlarmNotifyHistoryRequestAlarmSubscribers {
	s.SubscriberUuid = &v
	return s
}

type QueryRMSUnifiedAlarmNotifyHistoryResponseBody struct {
	AlarmNotifyHistories []*QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories `json:"AlarmNotifyHistories,omitempty" xml:"AlarmNotifyHistories,omitempty" type:"Repeated"`
	AlarmNotifySubs      []*QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifySubs      `json:"AlarmNotifySubs,omitempty" xml:"AlarmNotifySubs,omitempty" type:"Repeated"`
	CurrentPage          *int64                                                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize             *int64                                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId            *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode           *string                                                              `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage        *string                                                              `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	TotalCount           *int64                                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryRMSUnifiedAlarmNotifyHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRMSUnifiedAlarmNotifyHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBody) SetAlarmNotifyHistories(v []*QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) *QueryRMSUnifiedAlarmNotifyHistoryResponseBody {
	s.AlarmNotifyHistories = v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBody) SetAlarmNotifySubs(v []*QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifySubs) *QueryRMSUnifiedAlarmNotifyHistoryResponseBody {
	s.AlarmNotifySubs = v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBody) SetCurrentPage(v int64) *QueryRMSUnifiedAlarmNotifyHistoryResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBody) SetPageSize(v int64) *QueryRMSUnifiedAlarmNotifyHistoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBody) SetRequestId(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBody) SetResultCode(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBody) SetResultMessage(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBody) SetTotalCount(v int64) *QueryRMSUnifiedAlarmNotifyHistoryResponseBody {
	s.TotalCount = &v
	return s
}

type QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories struct {
	AggCount           *int64  `json:"AggCount,omitempty" xml:"AggCount,omitempty"`
	AlarmEventId       *string `json:"AlarmEventId,omitempty" xml:"AlarmEventId,omitempty"`
	AlarmHistoryId     *int64  `json:"AlarmHistoryId,omitempty" xml:"AlarmHistoryId,omitempty"`
	AlarmLevel         *int64  `json:"AlarmLevel,omitempty" xml:"AlarmLevel,omitempty"`
	AlarmRuleId        *int64  `json:"AlarmRuleId,omitempty" xml:"AlarmRuleId,omitempty"`
	AlarmRuleName      *string `json:"AlarmRuleName,omitempty" xml:"AlarmRuleName,omitempty"`
	AlarmSilenceTime   *int64  `json:"AlarmSilenceTime,omitempty" xml:"AlarmSilenceTime,omitempty"`
	AlarmStatus        *int64  `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	AlarmTargetDeleted *bool   `json:"AlarmTargetDeleted,omitempty" xml:"AlarmTargetDeleted,omitempty"`
	AlarmTime          *int64  `json:"AlarmTime,omitempty" xml:"AlarmTime,omitempty"`
	AlarmType          *string `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	AlarmUrl           *string `json:"AlarmUrl,omitempty" xml:"AlarmUrl,omitempty"`
	AlarmUrlWithDomain *string `json:"AlarmUrlWithDomain,omitempty" xml:"AlarmUrlWithDomain,omitempty"`
	Channel            *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	DataSourceName     *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	ErrorMessage       *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified        *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	NotifyContent      *string `json:"NotifyContent,omitempty" xml:"NotifyContent,omitempty"`
	Status             *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	Subscriber         *string `json:"Subscriber,omitempty" xml:"Subscriber,omitempty"`
	SubscriberName     *string `json:"SubscriberName,omitempty" xml:"SubscriberName,omitempty"`
	SubscriberSource   *string `json:"SubscriberSource,omitempty" xml:"SubscriberSource,omitempty"`
	SubscriberType     *string `json:"SubscriberType,omitempty" xml:"SubscriberType,omitempty"`
	TenantId           *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TraceId            *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	WorkspaceId        *int64  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) String() string {
	return tea.Prettify(s)
}

func (s QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) GoString() string {
	return s.String()
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetAggCount(v int64) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.AggCount = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetAlarmEventId(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.AlarmEventId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetAlarmHistoryId(v int64) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.AlarmHistoryId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetAlarmLevel(v int64) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.AlarmLevel = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetAlarmRuleId(v int64) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.AlarmRuleId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetAlarmRuleName(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.AlarmRuleName = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetAlarmSilenceTime(v int64) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.AlarmSilenceTime = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetAlarmStatus(v int64) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.AlarmStatus = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetAlarmTargetDeleted(v bool) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.AlarmTargetDeleted = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetAlarmTime(v int64) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.AlarmTime = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetAlarmType(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.AlarmType = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetAlarmUrl(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.AlarmUrl = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetAlarmUrlWithDomain(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.AlarmUrlWithDomain = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetChannel(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.Channel = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetDataSourceName(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.DataSourceName = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetErrorMessage(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.ErrorMessage = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetGmtCreate(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.GmtCreate = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetGmtModified(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.GmtModified = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetId(v int64) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.Id = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetNotifyContent(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.NotifyContent = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetStatus(v int64) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.Status = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetSubscriber(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.Subscriber = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetSubscriberName(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.SubscriberName = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetSubscriberSource(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.SubscriberSource = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetSubscriberType(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.SubscriberType = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetTenantId(v int64) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.TenantId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetTraceId(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.TraceId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories) SetWorkspaceId(v int64) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifyHistories {
	s.WorkspaceId = &v
	return s
}

type QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifySubs struct {
	Subscriber       *string `json:"Subscriber,omitempty" xml:"Subscriber,omitempty"`
	SubscriberName   *string `json:"SubscriberName,omitempty" xml:"SubscriberName,omitempty"`
	SubscriberSource *string `json:"SubscriberSource,omitempty" xml:"SubscriberSource,omitempty"`
	SubscriberType   *string `json:"SubscriberType,omitempty" xml:"SubscriberType,omitempty"`
	SubscriberUuid   *string `json:"SubscriberUuid,omitempty" xml:"SubscriberUuid,omitempty"`
}

func (s QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifySubs) String() string {
	return tea.Prettify(s)
}

func (s QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifySubs) GoString() string {
	return s.String()
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifySubs) SetSubscriber(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifySubs {
	s.Subscriber = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifySubs) SetSubscriberName(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifySubs {
	s.SubscriberName = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifySubs) SetSubscriberSource(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifySubs {
	s.SubscriberSource = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifySubs) SetSubscriberType(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifySubs {
	s.SubscriberType = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifySubs) SetSubscriberUuid(v string) *QueryRMSUnifiedAlarmNotifyHistoryResponseBodyAlarmNotifySubs {
	s.SubscriberUuid = &v
	return s
}

type QueryRMSUnifiedAlarmNotifyHistoryResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryRMSUnifiedAlarmNotifyHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRMSUnifiedAlarmNotifyHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRMSUnifiedAlarmNotifyHistoryResponse) GoString() string {
	return s.String()
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponse) SetHeaders(v map[string]*string) *QueryRMSUnifiedAlarmNotifyHistoryResponse {
	s.Headers = v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponse) SetStatusCode(v int32) *QueryRMSUnifiedAlarmNotifyHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRMSUnifiedAlarmNotifyHistoryResponse) SetBody(v *QueryRMSUnifiedAlarmNotifyHistoryResponseBody) *QueryRMSUnifiedAlarmNotifyHistoryResponse {
	s.Body = v
	return s
}

type QueryRMSUnifiedAlarmRuleRequest struct {
	AlarmStatus        *int64  `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	AlarmTargetJsonStr *string `json:"AlarmTargetJsonStr,omitempty" xml:"AlarmTargetJsonStr,omitempty"`
	Category           *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CurrentPage        *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Keyword            *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Level              *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	OnlyMe             *bool   `json:"OnlyMe,omitempty" xml:"OnlyMe,omitempty"`
	PageSize           *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RefDatasourceType  *string `json:"RefDatasourceType,omitempty" xml:"RefDatasourceType,omitempty"`
	RuleId             *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleStatus         *int64  `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	RuleUniqueIdentity *string `json:"RuleUniqueIdentity,omitempty" xml:"RuleUniqueIdentity,omitempty"`
	WorkspaceName      *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryRMSUnifiedAlarmRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRMSUnifiedAlarmRuleRequest) GoString() string {
	return s.String()
}

func (s *QueryRMSUnifiedAlarmRuleRequest) SetAlarmStatus(v int64) *QueryRMSUnifiedAlarmRuleRequest {
	s.AlarmStatus = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleRequest) SetAlarmTargetJsonStr(v string) *QueryRMSUnifiedAlarmRuleRequest {
	s.AlarmTargetJsonStr = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleRequest) SetCategory(v string) *QueryRMSUnifiedAlarmRuleRequest {
	s.Category = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleRequest) SetCurrentPage(v int64) *QueryRMSUnifiedAlarmRuleRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleRequest) SetKeyword(v string) *QueryRMSUnifiedAlarmRuleRequest {
	s.Keyword = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleRequest) SetLevel(v int64) *QueryRMSUnifiedAlarmRuleRequest {
	s.Level = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleRequest) SetOnlyMe(v bool) *QueryRMSUnifiedAlarmRuleRequest {
	s.OnlyMe = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleRequest) SetPageSize(v int64) *QueryRMSUnifiedAlarmRuleRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleRequest) SetRefDatasourceType(v string) *QueryRMSUnifiedAlarmRuleRequest {
	s.RefDatasourceType = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleRequest) SetRuleId(v int64) *QueryRMSUnifiedAlarmRuleRequest {
	s.RuleId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleRequest) SetRuleStatus(v int64) *QueryRMSUnifiedAlarmRuleRequest {
	s.RuleStatus = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleRequest) SetRuleUniqueIdentity(v string) *QueryRMSUnifiedAlarmRuleRequest {
	s.RuleUniqueIdentity = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleRequest) SetWorkspaceName(v string) *QueryRMSUnifiedAlarmRuleRequest {
	s.WorkspaceName = &v
	return s
}

type QueryRMSUnifiedAlarmRuleResponseBody struct {
	AlarmRules    []*QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules `json:"AlarmRules,omitempty" xml:"AlarmRules,omitempty" type:"Repeated"`
	CurrentPage   *int64                                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize      *int64                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                           `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                           `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	TotalCount    *int64                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryRMSUnifiedAlarmRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRMSUnifiedAlarmRuleResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRMSUnifiedAlarmRuleResponseBody) SetAlarmRules(v []*QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) *QueryRMSUnifiedAlarmRuleResponseBody {
	s.AlarmRules = v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBody) SetCurrentPage(v int64) *QueryRMSUnifiedAlarmRuleResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBody) SetPageSize(v int64) *QueryRMSUnifiedAlarmRuleResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBody) SetRequestId(v string) *QueryRMSUnifiedAlarmRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBody) SetResultCode(v string) *QueryRMSUnifiedAlarmRuleResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBody) SetResultMessage(v string) *QueryRMSUnifiedAlarmRuleResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBody) SetTotalCount(v int64) *QueryRMSUnifiedAlarmRuleResponseBody {
	s.TotalCount = &v
	return s
}

type QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules struct {
	AlarmEmpty         *int64    `json:"AlarmEmpty,omitempty" xml:"AlarmEmpty,omitempty"`
	AlarmStatus        *int64    `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	Category           *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	Channels           *string   `json:"Channels,omitempty" xml:"Channels,omitempty"`
	ConditionsDes      []*string `json:"ConditionsDes,omitempty" xml:"ConditionsDes,omitempty" type:"Repeated"`
	Creator            *string   `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Deleted            *int64    `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Emergency          *string   `json:"Emergency,omitempty" xml:"Emergency,omitempty"`
	EmergencyUrl       *string   `json:"EmergencyUrl,omitempty" xml:"EmergencyUrl,omitempty"`
	GmtCreate          *string   `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified        *string   `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id                 *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	Level              *int64    `json:"Level,omitempty" xml:"Level,omitempty"`
	Modifier           *string   `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	Name               *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NotifyEmpty        *int64    `json:"NotifyEmpty,omitempty" xml:"NotifyEmpty,omitempty"`
	NotifyFiring       *int64    `json:"NotifyFiring,omitempty" xml:"NotifyFiring,omitempty"`
	NotifyRecovered    *int64    `json:"NotifyRecovered,omitempty" xml:"NotifyRecovered,omitempty"`
	NotifyTimeFilter   *string   `json:"NotifyTimeFilter,omitempty" xml:"NotifyTimeFilter,omitempty"`
	NotifyTplId        *int64    `json:"NotifyTplId,omitempty" xml:"NotifyTplId,omitempty"`
	PendingHit         *int64    `json:"PendingHit,omitempty" xml:"PendingHit,omitempty"`
	RecoveredHit       *int64    `json:"RecoveredHit,omitempty" xml:"RecoveredHit,omitempty"`
	RuleConfig         *string   `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	SilenceTime        *int64    `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	SourceId           *string   `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	Status             *int64    `json:"Status,omitempty" xml:"Status,omitempty"`
	Step               *int64    `json:"Step,omitempty" xml:"Step,omitempty"`
	SuspendedEndTime   *int64    `json:"SuspendedEndTime,omitempty" xml:"SuspendedEndTime,omitempty"`
	SuspendedReason    *string   `json:"SuspendedReason,omitempty" xml:"SuspendedReason,omitempty"`
	SuspendedStartTime *int64    `json:"SuspendedStartTime,omitempty" xml:"SuspendedStartTime,omitempty"`
	TenantId           *int64    `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TimeZone           *string   `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	UniqueIdentity     *string   `json:"UniqueIdentity,omitempty" xml:"UniqueIdentity,omitempty"`
	WorkspaceId        *int64    `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) String() string {
	return tea.Prettify(s)
}

func (s QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) GoString() string {
	return s.String()
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetAlarmEmpty(v int64) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.AlarmEmpty = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetAlarmStatus(v int64) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.AlarmStatus = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetCategory(v string) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.Category = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetChannels(v string) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.Channels = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetConditionsDes(v []*string) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.ConditionsDes = v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetCreator(v string) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.Creator = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetDeleted(v int64) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.Deleted = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetEmergency(v string) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.Emergency = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetEmergencyUrl(v string) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.EmergencyUrl = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetGmtCreate(v string) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.GmtCreate = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetGmtModified(v string) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.GmtModified = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetId(v int64) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.Id = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetLevel(v int64) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.Level = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetModifier(v string) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.Modifier = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetName(v string) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.Name = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetNotifyEmpty(v int64) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.NotifyEmpty = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetNotifyFiring(v int64) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.NotifyFiring = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetNotifyRecovered(v int64) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.NotifyRecovered = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetNotifyTimeFilter(v string) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.NotifyTimeFilter = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetNotifyTplId(v int64) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.NotifyTplId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetPendingHit(v int64) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.PendingHit = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetRecoveredHit(v int64) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.RecoveredHit = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetRuleConfig(v string) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.RuleConfig = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetSilenceTime(v int64) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.SilenceTime = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetSourceId(v string) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.SourceId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetStatus(v int64) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.Status = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetStep(v int64) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.Step = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetSuspendedEndTime(v int64) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.SuspendedEndTime = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetSuspendedReason(v string) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.SuspendedReason = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetSuspendedStartTime(v int64) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.SuspendedStartTime = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetTenantId(v int64) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.TenantId = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetTimeZone(v string) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.TimeZone = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetUniqueIdentity(v string) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.UniqueIdentity = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules) SetWorkspaceId(v int64) *QueryRMSUnifiedAlarmRuleResponseBodyAlarmRules {
	s.WorkspaceId = &v
	return s
}

type QueryRMSUnifiedAlarmRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryRMSUnifiedAlarmRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRMSUnifiedAlarmRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRMSUnifiedAlarmRuleResponse) GoString() string {
	return s.String()
}

func (s *QueryRMSUnifiedAlarmRuleResponse) SetHeaders(v map[string]*string) *QueryRMSUnifiedAlarmRuleResponse {
	s.Headers = v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponse) SetStatusCode(v int32) *QueryRMSUnifiedAlarmRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRMSUnifiedAlarmRuleResponse) SetBody(v *QueryRMSUnifiedAlarmRuleResponseBody) *QueryRMSUnifiedAlarmRuleResponse {
	s.Body = v
	return s
}

type ResendMqSofamqDLQMessageBatchRequest struct {
	Cell       *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgIds     *string `json:"MsgIds,omitempty" xml:"MsgIds,omitempty"`
}

func (s ResendMqSofamqDLQMessageBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s ResendMqSofamqDLQMessageBatchRequest) GoString() string {
	return s.String()
}

func (s *ResendMqSofamqDLQMessageBatchRequest) SetCell(v string) *ResendMqSofamqDLQMessageBatchRequest {
	s.Cell = &v
	return s
}

func (s *ResendMqSofamqDLQMessageBatchRequest) SetGroupId(v string) *ResendMqSofamqDLQMessageBatchRequest {
	s.GroupId = &v
	return s
}

func (s *ResendMqSofamqDLQMessageBatchRequest) SetInstanceId(v string) *ResendMqSofamqDLQMessageBatchRequest {
	s.InstanceId = &v
	return s
}

func (s *ResendMqSofamqDLQMessageBatchRequest) SetMsgIds(v string) *ResendMqSofamqDLQMessageBatchRequest {
	s.MsgIds = &v
	return s
}

type ResendMqSofamqDLQMessageBatchResponseBody struct {
	Data          *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ResendMqSofamqDLQMessageBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResendMqSofamqDLQMessageBatchResponseBody) GoString() string {
	return s.String()
}

func (s *ResendMqSofamqDLQMessageBatchResponseBody) SetData(v string) *ResendMqSofamqDLQMessageBatchResponseBody {
	s.Data = &v
	return s
}

func (s *ResendMqSofamqDLQMessageBatchResponseBody) SetRequestId(v string) *ResendMqSofamqDLQMessageBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResendMqSofamqDLQMessageBatchResponseBody) SetResultCode(v string) *ResendMqSofamqDLQMessageBatchResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ResendMqSofamqDLQMessageBatchResponseBody) SetResultMessage(v string) *ResendMqSofamqDLQMessageBatchResponseBody {
	s.ResultMessage = &v
	return s
}

type ResendMqSofamqDLQMessageBatchResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResendMqSofamqDLQMessageBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResendMqSofamqDLQMessageBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s ResendMqSofamqDLQMessageBatchResponse) GoString() string {
	return s.String()
}

func (s *ResendMqSofamqDLQMessageBatchResponse) SetHeaders(v map[string]*string) *ResendMqSofamqDLQMessageBatchResponse {
	s.Headers = v
	return s
}

func (s *ResendMqSofamqDLQMessageBatchResponse) SetStatusCode(v int32) *ResendMqSofamqDLQMessageBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *ResendMqSofamqDLQMessageBatchResponse) SetBody(v *ResendMqSofamqDLQMessageBatchResponseBody) *ResendMqSofamqDLQMessageBatchResponse {
	s.Body = v
	return s
}

type ResendMqSofamqDLQMessageByIdRequest struct {
	Cell       *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s ResendMqSofamqDLQMessageByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ResendMqSofamqDLQMessageByIdRequest) GoString() string {
	return s.String()
}

func (s *ResendMqSofamqDLQMessageByIdRequest) SetCell(v string) *ResendMqSofamqDLQMessageByIdRequest {
	s.Cell = &v
	return s
}

func (s *ResendMqSofamqDLQMessageByIdRequest) SetGroupId(v string) *ResendMqSofamqDLQMessageByIdRequest {
	s.GroupId = &v
	return s
}

func (s *ResendMqSofamqDLQMessageByIdRequest) SetInstanceId(v string) *ResendMqSofamqDLQMessageByIdRequest {
	s.InstanceId = &v
	return s
}

func (s *ResendMqSofamqDLQMessageByIdRequest) SetMsgId(v string) *ResendMqSofamqDLQMessageByIdRequest {
	s.MsgId = &v
	return s
}

type ResendMqSofamqDLQMessageByIdResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ResendMqSofamqDLQMessageByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResendMqSofamqDLQMessageByIdResponseBody) GoString() string {
	return s.String()
}

func (s *ResendMqSofamqDLQMessageByIdResponseBody) SetRequestId(v string) *ResendMqSofamqDLQMessageByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResendMqSofamqDLQMessageByIdResponseBody) SetResultCode(v string) *ResendMqSofamqDLQMessageByIdResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ResendMqSofamqDLQMessageByIdResponseBody) SetResultMessage(v string) *ResendMqSofamqDLQMessageByIdResponseBody {
	s.ResultMessage = &v
	return s
}

type ResendMqSofamqDLQMessageByIdResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResendMqSofamqDLQMessageByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResendMqSofamqDLQMessageByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ResendMqSofamqDLQMessageByIdResponse) GoString() string {
	return s.String()
}

func (s *ResendMqSofamqDLQMessageByIdResponse) SetHeaders(v map[string]*string) *ResendMqSofamqDLQMessageByIdResponse {
	s.Headers = v
	return s
}

func (s *ResendMqSofamqDLQMessageByIdResponse) SetStatusCode(v int32) *ResendMqSofamqDLQMessageByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ResendMqSofamqDLQMessageByIdResponse) SetBody(v *ResendMqSofamqDLQMessageByIdResponseBody) *ResendMqSofamqDLQMessageByIdResponse {
	s.Body = v
	return s
}

type ResetMqSofamqConsumerOffsetRequest struct {
	Cell           *string `json:"Cell,omitempty" xml:"Cell,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResetTimestamp *int64  `json:"ResetTimestamp,omitempty" xml:"ResetTimestamp,omitempty"`
	Topic          *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Type           *int64  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ResetMqSofamqConsumerOffsetRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetMqSofamqConsumerOffsetRequest) GoString() string {
	return s.String()
}

func (s *ResetMqSofamqConsumerOffsetRequest) SetCell(v string) *ResetMqSofamqConsumerOffsetRequest {
	s.Cell = &v
	return s
}

func (s *ResetMqSofamqConsumerOffsetRequest) SetGroupId(v string) *ResetMqSofamqConsumerOffsetRequest {
	s.GroupId = &v
	return s
}

func (s *ResetMqSofamqConsumerOffsetRequest) SetInstanceId(v string) *ResetMqSofamqConsumerOffsetRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetMqSofamqConsumerOffsetRequest) SetResetTimestamp(v int64) *ResetMqSofamqConsumerOffsetRequest {
	s.ResetTimestamp = &v
	return s
}

func (s *ResetMqSofamqConsumerOffsetRequest) SetTopic(v string) *ResetMqSofamqConsumerOffsetRequest {
	s.Topic = &v
	return s
}

func (s *ResetMqSofamqConsumerOffsetRequest) SetType(v int64) *ResetMqSofamqConsumerOffsetRequest {
	s.Type = &v
	return s
}

type ResetMqSofamqConsumerOffsetResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ResetMqSofamqConsumerOffsetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetMqSofamqConsumerOffsetResponseBody) GoString() string {
	return s.String()
}

func (s *ResetMqSofamqConsumerOffsetResponseBody) SetRequestId(v string) *ResetMqSofamqConsumerOffsetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetMqSofamqConsumerOffsetResponseBody) SetResultCode(v string) *ResetMqSofamqConsumerOffsetResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ResetMqSofamqConsumerOffsetResponseBody) SetResultMessage(v string) *ResetMqSofamqConsumerOffsetResponseBody {
	s.ResultMessage = &v
	return s
}

type ResetMqSofamqConsumerOffsetResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetMqSofamqConsumerOffsetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetMqSofamqConsumerOffsetResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetMqSofamqConsumerOffsetResponse) GoString() string {
	return s.String()
}

func (s *ResetMqSofamqConsumerOffsetResponse) SetHeaders(v map[string]*string) *ResetMqSofamqConsumerOffsetResponse {
	s.Headers = v
	return s
}

func (s *ResetMqSofamqConsumerOffsetResponse) SetStatusCode(v int32) *ResetMqSofamqConsumerOffsetResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetMqSofamqConsumerOffsetResponse) SetBody(v *ResetMqSofamqConsumerOffsetResponseBody) *ResetMqSofamqConsumerOffsetResponse {
	s.Body = v
	return s
}

type UpdateMqSofamqGroupRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ReadEnable *bool   `json:"ReadEnable,omitempty" xml:"ReadEnable,omitempty"`
}

func (s UpdateMqSofamqGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMqSofamqGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateMqSofamqGroupRequest) SetGroupId(v string) *UpdateMqSofamqGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateMqSofamqGroupRequest) SetInstanceId(v string) *UpdateMqSofamqGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateMqSofamqGroupRequest) SetReadEnable(v bool) *UpdateMqSofamqGroupRequest {
	s.ReadEnable = &v
	return s
}

type UpdateMqSofamqGroupResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s UpdateMqSofamqGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMqSofamqGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMqSofamqGroupResponseBody) SetRequestId(v string) *UpdateMqSofamqGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMqSofamqGroupResponseBody) SetResultCode(v string) *UpdateMqSofamqGroupResponseBody {
	s.ResultCode = &v
	return s
}

func (s *UpdateMqSofamqGroupResponseBody) SetResultMessage(v string) *UpdateMqSofamqGroupResponseBody {
	s.ResultMessage = &v
	return s
}

type UpdateMqSofamqGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateMqSofamqGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMqSofamqGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMqSofamqGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateMqSofamqGroupResponse) SetHeaders(v map[string]*string) *UpdateMqSofamqGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateMqSofamqGroupResponse) SetStatusCode(v int32) *UpdateMqSofamqGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMqSofamqGroupResponse) SetBody(v *UpdateMqSofamqGroupResponseBody) *UpdateMqSofamqGroupResponse {
	s.Body = v
	return s
}

type UpdateMqSofamqGroupRemarkRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateMqSofamqGroupRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMqSofamqGroupRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMqSofamqGroupRemarkRequest) SetGroupId(v string) *UpdateMqSofamqGroupRemarkRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateMqSofamqGroupRemarkRequest) SetInstanceId(v string) *UpdateMqSofamqGroupRemarkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateMqSofamqGroupRemarkRequest) SetRemark(v string) *UpdateMqSofamqGroupRemarkRequest {
	s.Remark = &v
	return s
}

type UpdateMqSofamqGroupRemarkResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s UpdateMqSofamqGroupRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMqSofamqGroupRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMqSofamqGroupRemarkResponseBody) SetRequestId(v string) *UpdateMqSofamqGroupRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMqSofamqGroupRemarkResponseBody) SetResultCode(v string) *UpdateMqSofamqGroupRemarkResponseBody {
	s.ResultCode = &v
	return s
}

func (s *UpdateMqSofamqGroupRemarkResponseBody) SetResultMessage(v string) *UpdateMqSofamqGroupRemarkResponseBody {
	s.ResultMessage = &v
	return s
}

type UpdateMqSofamqGroupRemarkResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateMqSofamqGroupRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMqSofamqGroupRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMqSofamqGroupRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateMqSofamqGroupRemarkResponse) SetHeaders(v map[string]*string) *UpdateMqSofamqGroupRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateMqSofamqGroupRemarkResponse) SetStatusCode(v int32) *UpdateMqSofamqGroupRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMqSofamqGroupRemarkResponse) SetBody(v *UpdateMqSofamqGroupRemarkResponseBody) *UpdateMqSofamqGroupRemarkResponse {
	s.Body = v
	return s
}

type UpdateMqSofamqTopicRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Perm       *int64  `json:"Perm,omitempty" xml:"Perm,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s UpdateMqSofamqTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMqSofamqTopicRequest) GoString() string {
	return s.String()
}

func (s *UpdateMqSofamqTopicRequest) SetInstanceId(v string) *UpdateMqSofamqTopicRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateMqSofamqTopicRequest) SetPerm(v int64) *UpdateMqSofamqTopicRequest {
	s.Perm = &v
	return s
}

func (s *UpdateMqSofamqTopicRequest) SetTopic(v string) *UpdateMqSofamqTopicRequest {
	s.Topic = &v
	return s
}

type UpdateMqSofamqTopicResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s UpdateMqSofamqTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMqSofamqTopicResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMqSofamqTopicResponseBody) SetRequestId(v string) *UpdateMqSofamqTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMqSofamqTopicResponseBody) SetResultCode(v string) *UpdateMqSofamqTopicResponseBody {
	s.ResultCode = &v
	return s
}

func (s *UpdateMqSofamqTopicResponseBody) SetResultMessage(v string) *UpdateMqSofamqTopicResponseBody {
	s.ResultMessage = &v
	return s
}

type UpdateMqSofamqTopicResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateMqSofamqTopicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMqSofamqTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMqSofamqTopicResponse) GoString() string {
	return s.String()
}

func (s *UpdateMqSofamqTopicResponse) SetHeaders(v map[string]*string) *UpdateMqSofamqTopicResponse {
	s.Headers = v
	return s
}

func (s *UpdateMqSofamqTopicResponse) SetStatusCode(v int32) *UpdateMqSofamqTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMqSofamqTopicResponse) SetBody(v *UpdateMqSofamqTopicResponseBody) *UpdateMqSofamqTopicResponse {
	s.Body = v
	return s
}

type UpdateMqSofamqTopicRemarkRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s UpdateMqSofamqTopicRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMqSofamqTopicRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMqSofamqTopicRemarkRequest) SetInstanceId(v string) *UpdateMqSofamqTopicRemarkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateMqSofamqTopicRemarkRequest) SetRemark(v string) *UpdateMqSofamqTopicRemarkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateMqSofamqTopicRemarkRequest) SetTopic(v string) *UpdateMqSofamqTopicRemarkRequest {
	s.Topic = &v
	return s
}

type UpdateMqSofamqTopicRemarkResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s UpdateMqSofamqTopicRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMqSofamqTopicRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMqSofamqTopicRemarkResponseBody) SetRequestId(v string) *UpdateMqSofamqTopicRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMqSofamqTopicRemarkResponseBody) SetResultCode(v string) *UpdateMqSofamqTopicRemarkResponseBody {
	s.ResultCode = &v
	return s
}

func (s *UpdateMqSofamqTopicRemarkResponseBody) SetResultMessage(v string) *UpdateMqSofamqTopicRemarkResponseBody {
	s.ResultMessage = &v
	return s
}

type UpdateMqSofamqTopicRemarkResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateMqSofamqTopicRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMqSofamqTopicRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMqSofamqTopicRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateMqSofamqTopicRemarkResponse) SetHeaders(v map[string]*string) *UpdateMqSofamqTopicRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateMqSofamqTopicRemarkResponse) SetStatusCode(v int32) *UpdateMqSofamqTopicRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMqSofamqTopicRemarkResponse) SetBody(v *UpdateMqSofamqTopicRemarkResponseBody) *UpdateMqSofamqTopicRemarkResponse {
	s.Body = v
	return s
}

type UpdateMqSofamqWarnRequest struct {
	AlertTime  *string `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	Contacts   *string `json:"Contacts,omitempty" xml:"Contacts,omitempty"`
	DelayTime  *int64  `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	Frequency  *int64  `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Threshold  *int64  `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	WarnId     *int64  `json:"WarnId,omitempty" xml:"WarnId,omitempty"`
}

func (s UpdateMqSofamqWarnRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMqSofamqWarnRequest) GoString() string {
	return s.String()
}

func (s *UpdateMqSofamqWarnRequest) SetAlertTime(v string) *UpdateMqSofamqWarnRequest {
	s.AlertTime = &v
	return s
}

func (s *UpdateMqSofamqWarnRequest) SetContacts(v string) *UpdateMqSofamqWarnRequest {
	s.Contacts = &v
	return s
}

func (s *UpdateMqSofamqWarnRequest) SetDelayTime(v int64) *UpdateMqSofamqWarnRequest {
	s.DelayTime = &v
	return s
}

func (s *UpdateMqSofamqWarnRequest) SetFrequency(v int64) *UpdateMqSofamqWarnRequest {
	s.Frequency = &v
	return s
}

func (s *UpdateMqSofamqWarnRequest) SetInstanceId(v string) *UpdateMqSofamqWarnRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateMqSofamqWarnRequest) SetThreshold(v int64) *UpdateMqSofamqWarnRequest {
	s.Threshold = &v
	return s
}

func (s *UpdateMqSofamqWarnRequest) SetWarnId(v int64) *UpdateMqSofamqWarnRequest {
	s.WarnId = &v
	return s
}

type UpdateMqSofamqWarnResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s UpdateMqSofamqWarnResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMqSofamqWarnResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMqSofamqWarnResponseBody) SetRequestId(v string) *UpdateMqSofamqWarnResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMqSofamqWarnResponseBody) SetResultCode(v string) *UpdateMqSofamqWarnResponseBody {
	s.ResultCode = &v
	return s
}

func (s *UpdateMqSofamqWarnResponseBody) SetResultMessage(v string) *UpdateMqSofamqWarnResponseBody {
	s.ResultMessage = &v
	return s
}

type UpdateMqSofamqWarnResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateMqSofamqWarnResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMqSofamqWarnResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMqSofamqWarnResponse) GoString() string {
	return s.String()
}

func (s *UpdateMqSofamqWarnResponse) SetHeaders(v map[string]*string) *UpdateMqSofamqWarnResponse {
	s.Headers = v
	return s
}

func (s *UpdateMqSofamqWarnResponse) SetStatusCode(v int32) *UpdateMqSofamqWarnResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMqSofamqWarnResponse) SetBody(v *UpdateMqSofamqWarnResponseBody) *UpdateMqSofamqWarnResponse {
	s.Body = v
	return s
}

type UpdateMsConfigAttributesRequest struct {
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	Desc          *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateMsConfigAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMsConfigAttributesRequest) GoString() string {
	return s.String()
}

func (s *UpdateMsConfigAttributesRequest) SetAttributeName(v string) *UpdateMsConfigAttributesRequest {
	s.AttributeName = &v
	return s
}

func (s *UpdateMsConfigAttributesRequest) SetDesc(v string) *UpdateMsConfigAttributesRequest {
	s.Desc = &v
	return s
}

func (s *UpdateMsConfigAttributesRequest) SetId(v int64) *UpdateMsConfigAttributesRequest {
	s.Id = &v
	return s
}

func (s *UpdateMsConfigAttributesRequest) SetInstanceId(v string) *UpdateMsConfigAttributesRequest {
	s.InstanceId = &v
	return s
}

type UpdateMsConfigAttributesResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s UpdateMsConfigAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMsConfigAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMsConfigAttributesResponseBody) SetRequestId(v string) *UpdateMsConfigAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMsConfigAttributesResponseBody) SetResultCode(v string) *UpdateMsConfigAttributesResponseBody {
	s.ResultCode = &v
	return s
}

func (s *UpdateMsConfigAttributesResponseBody) SetResultMessage(v string) *UpdateMsConfigAttributesResponseBody {
	s.ResultMessage = &v
	return s
}

type UpdateMsConfigAttributesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateMsConfigAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMsConfigAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMsConfigAttributesResponse) GoString() string {
	return s.String()
}

func (s *UpdateMsConfigAttributesResponse) SetHeaders(v map[string]*string) *UpdateMsConfigAttributesResponse {
	s.Headers = v
	return s
}

func (s *UpdateMsConfigAttributesResponse) SetStatusCode(v int32) *UpdateMsConfigAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMsConfigAttributesResponse) SetBody(v *UpdateMsConfigAttributesResponseBody) *UpdateMsConfigAttributesResponse {
	s.Body = v
	return s
}

type UpdateMsConfigResourcesRequest struct {
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Attributes      *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	Desc            *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	UpdateAttribute *bool   `json:"UpdateAttribute,omitempty" xml:"UpdateAttribute,omitempty"`
}

func (s UpdateMsConfigResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMsConfigResourcesRequest) GoString() string {
	return s.String()
}

func (s *UpdateMsConfigResourcesRequest) SetAppName(v string) *UpdateMsConfigResourcesRequest {
	s.AppName = &v
	return s
}

func (s *UpdateMsConfigResourcesRequest) SetAttributes(v string) *UpdateMsConfigResourcesRequest {
	s.Attributes = &v
	return s
}

func (s *UpdateMsConfigResourcesRequest) SetDesc(v string) *UpdateMsConfigResourcesRequest {
	s.Desc = &v
	return s
}

func (s *UpdateMsConfigResourcesRequest) SetId(v int64) *UpdateMsConfigResourcesRequest {
	s.Id = &v
	return s
}

func (s *UpdateMsConfigResourcesRequest) SetInstanceId(v string) *UpdateMsConfigResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateMsConfigResourcesRequest) SetRegion(v string) *UpdateMsConfigResourcesRequest {
	s.Region = &v
	return s
}

func (s *UpdateMsConfigResourcesRequest) SetResourceId(v string) *UpdateMsConfigResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *UpdateMsConfigResourcesRequest) SetUpdateAttribute(v bool) *UpdateMsConfigResourcesRequest {
	s.UpdateAttribute = &v
	return s
}

type UpdateMsConfigResourcesResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s UpdateMsConfigResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMsConfigResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMsConfigResourcesResponseBody) SetRequestId(v string) *UpdateMsConfigResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMsConfigResourcesResponseBody) SetResultCode(v string) *UpdateMsConfigResourcesResponseBody {
	s.ResultCode = &v
	return s
}

func (s *UpdateMsConfigResourcesResponseBody) SetResultMessage(v string) *UpdateMsConfigResourcesResponseBody {
	s.ResultMessage = &v
	return s
}

type UpdateMsConfigResourcesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateMsConfigResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMsConfigResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMsConfigResourcesResponse) GoString() string {
	return s.String()
}

func (s *UpdateMsConfigResourcesResponse) SetHeaders(v map[string]*string) *UpdateMsConfigResourcesResponse {
	s.Headers = v
	return s
}

func (s *UpdateMsConfigResourcesResponse) SetStatusCode(v int32) *UpdateMsConfigResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMsConfigResourcesResponse) SetBody(v *UpdateMsConfigResourcesResponseBody) *UpdateMsConfigResourcesResponse {
	s.Body = v
	return s
}

type UpdateRMSUnifiedAlarmRuleRequest struct {
	AlarmNodata             *int64                                          `json:"AlarmNodata,omitempty" xml:"AlarmNodata,omitempty"`
	Category                *string                                         `json:"Category,omitempty" xml:"Category,omitempty"`
	ChannelsRepeatList      []*string                                       `json:"ChannelsRepeatList,omitempty" xml:"ChannelsRepeatList,omitempty" type:"Repeated"`
	Emergency               *string                                         `json:"Emergency,omitempty" xml:"Emergency,omitempty"`
	EmergencyUrl            *string                                         `json:"EmergencyUrl,omitempty" xml:"EmergencyUrl,omitempty"`
	Id                      *int64                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	Level                   *int64                                          `json:"Level,omitempty" xml:"Level,omitempty"`
	Name                    *string                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	NotifyFiring            *int64                                          `json:"NotifyFiring,omitempty" xml:"NotifyFiring,omitempty"`
	NotifyNodata            *int64                                          `json:"NotifyNodata,omitempty" xml:"NotifyNodata,omitempty"`
	NotifyRecovered         *int64                                          `json:"NotifyRecovered,omitempty" xml:"NotifyRecovered,omitempty"`
	NotifyTarget            []*UpdateRMSUnifiedAlarmRuleRequestNotifyTarget `json:"NotifyTarget,omitempty" xml:"NotifyTarget,omitempty" type:"Repeated"`
	NotifyTimeFilterJsonStr *string                                         `json:"NotifyTimeFilterJsonStr,omitempty" xml:"NotifyTimeFilterJsonStr,omitempty"`
	PendingHit              *int64                                          `json:"PendingHit,omitempty" xml:"PendingHit,omitempty"`
	RecoveredHit            *int64                                          `json:"RecoveredHit,omitempty" xml:"RecoveredHit,omitempty"`
	RuleConfig              *string                                         `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	SilenceTime             *int64                                          `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	Status                  *string                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Step                    *int64                                          `json:"Step,omitempty" xml:"Step,omitempty"`
	SuspendedEndTime        *int64                                          `json:"SuspendedEndTime,omitempty" xml:"SuspendedEndTime,omitempty"`
	SuspendedReason         *string                                         `json:"SuspendedReason,omitempty" xml:"SuspendedReason,omitempty"`
	SuspendedStartTime      *int64                                          `json:"SuspendedStartTime,omitempty" xml:"SuspendedStartTime,omitempty"`
	WorkspaceName           *string                                         `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s UpdateRMSUnifiedAlarmRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRMSUnifiedAlarmRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetAlarmNodata(v int64) *UpdateRMSUnifiedAlarmRuleRequest {
	s.AlarmNodata = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetCategory(v string) *UpdateRMSUnifiedAlarmRuleRequest {
	s.Category = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetChannelsRepeatList(v []*string) *UpdateRMSUnifiedAlarmRuleRequest {
	s.ChannelsRepeatList = v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetEmergency(v string) *UpdateRMSUnifiedAlarmRuleRequest {
	s.Emergency = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetEmergencyUrl(v string) *UpdateRMSUnifiedAlarmRuleRequest {
	s.EmergencyUrl = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetId(v int64) *UpdateRMSUnifiedAlarmRuleRequest {
	s.Id = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetLevel(v int64) *UpdateRMSUnifiedAlarmRuleRequest {
	s.Level = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetName(v string) *UpdateRMSUnifiedAlarmRuleRequest {
	s.Name = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetNotifyFiring(v int64) *UpdateRMSUnifiedAlarmRuleRequest {
	s.NotifyFiring = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetNotifyNodata(v int64) *UpdateRMSUnifiedAlarmRuleRequest {
	s.NotifyNodata = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetNotifyRecovered(v int64) *UpdateRMSUnifiedAlarmRuleRequest {
	s.NotifyRecovered = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetNotifyTarget(v []*UpdateRMSUnifiedAlarmRuleRequestNotifyTarget) *UpdateRMSUnifiedAlarmRuleRequest {
	s.NotifyTarget = v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetNotifyTimeFilterJsonStr(v string) *UpdateRMSUnifiedAlarmRuleRequest {
	s.NotifyTimeFilterJsonStr = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetPendingHit(v int64) *UpdateRMSUnifiedAlarmRuleRequest {
	s.PendingHit = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetRecoveredHit(v int64) *UpdateRMSUnifiedAlarmRuleRequest {
	s.RecoveredHit = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetRuleConfig(v string) *UpdateRMSUnifiedAlarmRuleRequest {
	s.RuleConfig = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetSilenceTime(v int64) *UpdateRMSUnifiedAlarmRuleRequest {
	s.SilenceTime = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetStatus(v string) *UpdateRMSUnifiedAlarmRuleRequest {
	s.Status = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetStep(v int64) *UpdateRMSUnifiedAlarmRuleRequest {
	s.Step = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetSuspendedEndTime(v int64) *UpdateRMSUnifiedAlarmRuleRequest {
	s.SuspendedEndTime = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetSuspendedReason(v string) *UpdateRMSUnifiedAlarmRuleRequest {
	s.SuspendedReason = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetSuspendedStartTime(v int64) *UpdateRMSUnifiedAlarmRuleRequest {
	s.SuspendedStartTime = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequest) SetWorkspaceName(v string) *UpdateRMSUnifiedAlarmRuleRequest {
	s.WorkspaceName = &v
	return s
}

type UpdateRMSUnifiedAlarmRuleRequestNotifyTarget struct {
	Subscriber       *string `json:"Subscriber,omitempty" xml:"Subscriber,omitempty"`
	SubscriberName   *string `json:"SubscriberName,omitempty" xml:"SubscriberName,omitempty"`
	SubscriberSource *string `json:"SubscriberSource,omitempty" xml:"SubscriberSource,omitempty"`
	SubscriberType   *string `json:"SubscriberType,omitempty" xml:"SubscriberType,omitempty"`
	SubscriberUuid   *string `json:"SubscriberUuid,omitempty" xml:"SubscriberUuid,omitempty"`
}

func (s UpdateRMSUnifiedAlarmRuleRequestNotifyTarget) String() string {
	return tea.Prettify(s)
}

func (s UpdateRMSUnifiedAlarmRuleRequestNotifyTarget) GoString() string {
	return s.String()
}

func (s *UpdateRMSUnifiedAlarmRuleRequestNotifyTarget) SetSubscriber(v string) *UpdateRMSUnifiedAlarmRuleRequestNotifyTarget {
	s.Subscriber = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequestNotifyTarget) SetSubscriberName(v string) *UpdateRMSUnifiedAlarmRuleRequestNotifyTarget {
	s.SubscriberName = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequestNotifyTarget) SetSubscriberSource(v string) *UpdateRMSUnifiedAlarmRuleRequestNotifyTarget {
	s.SubscriberSource = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequestNotifyTarget) SetSubscriberType(v string) *UpdateRMSUnifiedAlarmRuleRequestNotifyTarget {
	s.SubscriberType = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleRequestNotifyTarget) SetSubscriberUuid(v string) *UpdateRMSUnifiedAlarmRuleRequestNotifyTarget {
	s.SubscriberUuid = &v
	return s
}

type UpdateRMSUnifiedAlarmRuleResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s UpdateRMSUnifiedAlarmRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRMSUnifiedAlarmRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRMSUnifiedAlarmRuleResponseBody) SetRequestId(v string) *UpdateRMSUnifiedAlarmRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleResponseBody) SetResultCode(v string) *UpdateRMSUnifiedAlarmRuleResponseBody {
	s.ResultCode = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleResponseBody) SetResultMessage(v string) *UpdateRMSUnifiedAlarmRuleResponseBody {
	s.ResultMessage = &v
	return s
}

type UpdateRMSUnifiedAlarmRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRMSUnifiedAlarmRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRMSUnifiedAlarmRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRMSUnifiedAlarmRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRMSUnifiedAlarmRuleResponse) SetHeaders(v map[string]*string) *UpdateRMSUnifiedAlarmRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleResponse) SetStatusCode(v int32) *UpdateRMSUnifiedAlarmRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRMSUnifiedAlarmRuleResponse) SetBody(v *UpdateRMSUnifiedAlarmRuleResponseBody) *UpdateRMSUnifiedAlarmRuleResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-hangzhou-finance": tea.String("sofa.cn-shanghai.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("sofa"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddMsConfigAttributesWithOptions(request *AddMsConfigAttributesRequest, runtime *util.RuntimeOptions) (_result *AddMsConfigAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttributeName)) {
		body["AttributeName"] = request.AttributeName
	}

	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		body["Desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddMsConfigAttributes"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddMsConfigAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddMsConfigAttributes(request *AddMsConfigAttributesRequest) (_result *AddMsConfigAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddMsConfigAttributesResponse{}
	_body, _err := client.AddMsConfigAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddMsConfigResourcesWithOptions(request *AddMsConfigResourcesRequest, runtime *util.RuntimeOptions) (_result *AddMsConfigResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Attributes)) {
		body["Attributes"] = request.Attributes
	}

	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		body["Desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddMsConfigResources"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddMsConfigResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddMsConfigResources(request *AddMsConfigResourcesRequest) (_result *AddMsConfigResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddMsConfigResourcesResponse{}
	_body, _err := client.AddMsConfigResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMqSofamqGroupWithOptions(request *CreateMqSofamqGroupRequest, runtime *util.RuntimeOptions) (_result *CreateMqSofamqGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		body["GroupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMqSofamqGroup"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMqSofamqGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMqSofamqGroup(request *CreateMqSofamqGroupRequest) (_result *CreateMqSofamqGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMqSofamqGroupResponse{}
	_body, _err := client.CreateMqSofamqGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMqSofamqTopicWithOptions(request *CreateMqSofamqTopicRequest, runtime *util.RuntimeOptions) (_result *CreateMqSofamqTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		body["MessageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMqSofamqTopic"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMqSofamqTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMqSofamqTopic(request *CreateMqSofamqTopicRequest) (_result *CreateMqSofamqTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMqSofamqTopicResponse{}
	_body, _err := client.CreateMqSofamqTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRMSUnifiedAlarmRuleWithOptions(request *CreateRMSUnifiedAlarmRuleRequest, runtime *util.RuntimeOptions) (_result *CreateRMSUnifiedAlarmRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmNodata)) {
		body["AlarmNodata"] = request.AlarmNodata
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelsRepeatList)) {
		body["ChannelsRepeatList"] = request.ChannelsRepeatList
	}

	if !tea.BoolValue(util.IsUnset(request.Emergency)) {
		body["Emergency"] = request.Emergency
	}

	if !tea.BoolValue(util.IsUnset(request.EmergencyUrl)) {
		body["EmergencyUrl"] = request.EmergencyUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyFiring)) {
		body["NotifyFiring"] = request.NotifyFiring
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyNodata)) {
		body["NotifyNodata"] = request.NotifyNodata
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyRecovered)) {
		body["NotifyRecovered"] = request.NotifyRecovered
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTarget)) {
		body["NotifyTarget"] = request.NotifyTarget
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTimeFilterJsonStr)) {
		body["NotifyTimeFilterJsonStr"] = request.NotifyTimeFilterJsonStr
	}

	if !tea.BoolValue(util.IsUnset(request.PendingHit)) {
		body["PendingHit"] = request.PendingHit
	}

	if !tea.BoolValue(util.IsUnset(request.RecoveredHit)) {
		body["RecoveredHit"] = request.RecoveredHit
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConfig)) {
		body["RuleConfig"] = request.RuleConfig
	}

	if !tea.BoolValue(util.IsUnset(request.SilenceTime)) {
		body["SilenceTime"] = request.SilenceTime
	}

	if !tea.BoolValue(util.IsUnset(request.Step)) {
		body["Step"] = request.Step
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceName)) {
		body["WorkspaceName"] = request.WorkspaceName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRMSUnifiedAlarmRule"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRMSUnifiedAlarmRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRMSUnifiedAlarmRule(request *CreateRMSUnifiedAlarmRuleRequest) (_result *CreateRMSUnifiedAlarmRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRMSUnifiedAlarmRuleResponse{}
	_body, _err := client.CreateRMSUnifiedAlarmRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMqSofamqGroupWithOptions(request *DeleteMqSofamqGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteMqSofamqGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMqSofamqGroup"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMqSofamqGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMqSofamqGroup(request *DeleteMqSofamqGroupRequest) (_result *DeleteMqSofamqGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMqSofamqGroupResponse{}
	_body, _err := client.DeleteMqSofamqGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMqSofamqTopicWithOptions(request *DeleteMqSofamqTopicRequest, runtime *util.RuntimeOptions) (_result *DeleteMqSofamqTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMqSofamqTopic"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMqSofamqTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMqSofamqTopic(request *DeleteMqSofamqTopicRequest) (_result *DeleteMqSofamqTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMqSofamqTopicResponse{}
	_body, _err := client.DeleteMqSofamqTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMqSofamqTraceWithOptions(request *DeleteMqSofamqTraceRequest, runtime *util.RuntimeOptions) (_result *DeleteMqSofamqTraceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryId)) {
		body["QueryId"] = request.QueryId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMqSofamqTrace"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMqSofamqTraceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMqSofamqTrace(request *DeleteMqSofamqTraceRequest) (_result *DeleteMqSofamqTraceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMqSofamqTraceResponse{}
	_body, _err := client.DeleteMqSofamqTraceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMqSofamqWarnWithOptions(request *DeleteMqSofamqWarnRequest, runtime *util.RuntimeOptions) (_result *DeleteMqSofamqWarnResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.WarnId)) {
		body["WarnId"] = request.WarnId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMqSofamqWarn"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMqSofamqWarnResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMqSofamqWarn(request *DeleteMqSofamqWarnRequest) (_result *DeleteMqSofamqWarnResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMqSofamqWarnResponse{}
	_body, _err := client.DeleteMqSofamqWarnWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMsConfigAttributesWithOptions(request *DeleteMsConfigAttributesRequest, runtime *util.RuntimeOptions) (_result *DeleteMsConfigAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMsConfigAttributes"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMsConfigAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMsConfigAttributes(request *DeleteMsConfigAttributesRequest) (_result *DeleteMsConfigAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMsConfigAttributesResponse{}
	_body, _err := client.DeleteMsConfigAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMsConfigResourcesWithOptions(request *DeleteMsConfigResourcesRequest, runtime *util.RuntimeOptions) (_result *DeleteMsConfigResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMsConfigResources"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMsConfigResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMsConfigResources(request *DeleteMsConfigResourcesRequest) (_result *DeleteMsConfigResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMsConfigResourcesResponse{}
	_body, _err := client.DeleteMsConfigResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRMSUnifiedAlarmRuleWithOptions(request *DeleteRMSUnifiedAlarmRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteRMSUnifiedAlarmRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceName)) {
		body["WorkspaceName"] = request.WorkspaceName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRMSUnifiedAlarmRule"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRMSUnifiedAlarmRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRMSUnifiedAlarmRule(request *DeleteRMSUnifiedAlarmRuleRequest) (_result *DeleteRMSUnifiedAlarmRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRMSUnifiedAlarmRuleResponse{}
	_body, _err := client.DeleteRMSUnifiedAlarmRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCasComputersWithOptions(request *DescribeCasComputersRequest, runtime *util.RuntimeOptions) (_result *DescribeCasComputersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppServiceIdsRepeatList)) {
		body["AppServiceIdsRepeatList"] = request.AppServiceIdsRepeatList
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		body["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCasComputers"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCasComputersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCasComputers(request *DescribeCasComputersRequest) (_result *DescribeCasComputersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCasComputersResponse{}
	_body, _err := client.DescribeCasComputersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableMqSofamqWarnWithOptions(request *DisableMqSofamqWarnRequest, runtime *util.RuntimeOptions) (_result *DisableMqSofamqWarnResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.WarnId)) {
		body["WarnId"] = request.WarnId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableMqSofamqWarn"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableMqSofamqWarnResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableMqSofamqWarn(request *DisableMqSofamqWarnRequest) (_result *DisableMqSofamqWarnResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableMqSofamqWarnResponse{}
	_body, _err := client.DisableMqSofamqWarnWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableMqSofamqWarnWithOptions(request *EnableMqSofamqWarnRequest, runtime *util.RuntimeOptions) (_result *EnableMqSofamqWarnResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.WarnId)) {
		body["WarnId"] = request.WarnId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableMqSofamqWarn"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableMqSofamqWarnResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableMqSofamqWarn(request *EnableMqSofamqWarnRequest) (_result *EnableMqSofamqWarnResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableMqSofamqWarnResponse{}
	_body, _err := client.EnableMqSofamqWarnWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMqSofamqConsumerJStackWithOptions(request *GetMqSofamqConsumerJStackRequest, runtime *util.RuntimeOptions) (_result *GetMqSofamqConsumerJStackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cell)) {
		body["Cell"] = request.Cell
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMqSofamqConsumerJStack"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMqSofamqConsumerJStackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMqSofamqConsumerJStack(request *GetMqSofamqConsumerJStackRequest) (_result *GetMqSofamqConsumerJStackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMqSofamqConsumerJStackResponse{}
	_body, _err := client.GetMqSofamqConsumerJStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMqSofamqConsumerStatusWithOptions(request *GetMqSofamqConsumerStatusRequest, runtime *util.RuntimeOptions) (_result *GetMqSofamqConsumerStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cell)) {
		body["Cell"] = request.Cell
	}

	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		body["Detail"] = request.Detail
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NeedJstack)) {
		body["NeedJstack"] = request.NeedJstack
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMqSofamqConsumerStatus"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMqSofamqConsumerStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMqSofamqConsumerStatus(request *GetMqSofamqConsumerStatusRequest) (_result *GetMqSofamqConsumerStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMqSofamqConsumerStatusResponse{}
	_body, _err := client.GetMqSofamqConsumerStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMqSofamqDLQMessageByIdWithOptions(request *GetMqSofamqDLQMessageByIdRequest, runtime *util.RuntimeOptions) (_result *GetMqSofamqDLQMessageByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cell)) {
		body["Cell"] = request.Cell
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgId)) {
		body["MsgId"] = request.MsgId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMqSofamqDLQMessageById"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMqSofamqDLQMessageByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMqSofamqDLQMessageById(request *GetMqSofamqDLQMessageByIdRequest) (_result *GetMqSofamqDLQMessageByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMqSofamqDLQMessageByIdResponse{}
	_body, _err := client.GetMqSofamqDLQMessageByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMqSofamqMessageByIdWithOptions(request *GetMqSofamqMessageByIdRequest, runtime *util.RuntimeOptions) (_result *GetMqSofamqMessageByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cell)) {
		body["Cell"] = request.Cell
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgId)) {
		body["MsgId"] = request.MsgId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMqSofamqMessageById"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMqSofamqMessageByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMqSofamqMessageById(request *GetMqSofamqMessageByIdRequest) (_result *GetMqSofamqMessageByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMqSofamqMessageByIdResponse{}
	_body, _err := client.GetMqSofamqMessageByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMqSofamqTraceByMsgIdWithOptions(request *GetMqSofamqTraceByMsgIdRequest, runtime *util.RuntimeOptions) (_result *GetMqSofamqTraceByMsgIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		body["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.Cell)) {
		body["Cell"] = request.Cell
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgId)) {
		body["MsgId"] = request.MsgId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMqSofamqTraceByMsgId"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMqSofamqTraceByMsgIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMqSofamqTraceByMsgId(request *GetMqSofamqTraceByMsgIdRequest) (_result *GetMqSofamqTraceByMsgIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMqSofamqTraceByMsgIdResponse{}
	_body, _err := client.GetMqSofamqTraceByMsgIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMqSofamqTraceResultWithOptions(request *GetMqSofamqTraceResultRequest, runtime *util.RuntimeOptions) (_result *GetMqSofamqTraceResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryId)) {
		body["QueryId"] = request.QueryId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMqSofamqTraceResult"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMqSofamqTraceResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMqSofamqTraceResult(request *GetMqSofamqTraceResultRequest) (_result *GetMqSofamqTraceResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMqSofamqTraceResultResponse{}
	_body, _err := client.GetMqSofamqTraceResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMsConfigAttributesWithOptions(request *GetMsConfigAttributesRequest, runtime *util.RuntimeOptions) (_result *GetMsConfigAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMsConfigAttributes"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMsConfigAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMsConfigAttributes(request *GetMsConfigAttributesRequest) (_result *GetMsConfigAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMsConfigAttributesResponse{}
	_body, _err := client.GetMsConfigAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMsConfigResourcesWithOptions(request *GetMsConfigResourcesRequest, runtime *util.RuntimeOptions) (_result *GetMsConfigResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMsConfigResources"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMsConfigResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMsConfigResources(request *GetMsConfigResourcesRequest) (_result *GetMsConfigResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMsConfigResourcesResponse{}
	_body, _err := client.GetMsConfigResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrayPushMsConfigDataWithOptions(request *GrayPushMsConfigDataRequest, runtime *util.RuntimeOptions) (_result *GrayPushMsConfigDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttributeId)) {
		body["AttributeId"] = request.AttributeId
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Hosts)) {
		body["Hosts"] = request.Hosts
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["Operator"] = request.Operator
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GrayPushMsConfigData"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GrayPushMsConfigDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrayPushMsConfigData(request *GrayPushMsConfigDataRequest) (_result *GrayPushMsConfigDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrayPushMsConfigDataResponse{}
	_body, _err := client.GrayPushMsConfigDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMqSofamqGroupWithOptions(request *ListMqSofamqGroupRequest, runtime *util.RuntimeOptions) (_result *ListMqSofamqGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		body["GroupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMqSofamqGroup"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMqSofamqGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMqSofamqGroup(request *ListMqSofamqGroupRequest) (_result *ListMqSofamqGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMqSofamqGroupResponse{}
	_body, _err := client.ListMqSofamqGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMqSofamqTopicWithOptions(request *ListMqSofamqTopicRequest, runtime *util.RuntimeOptions) (_result *ListMqSofamqTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMqSofamqTopic"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMqSofamqTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMqSofamqTopic(request *ListMqSofamqTopicRequest) (_result *ListMqSofamqTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMqSofamqTopicResponse{}
	_body, _err := client.ListMqSofamqTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMqSofamqTraceWithOptions(request *ListMqSofamqTraceRequest, runtime *util.RuntimeOptions) (_result *ListMqSofamqTraceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryItem)) {
		body["QueryItem"] = request.QueryItem
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMqSofamqTrace"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMqSofamqTraceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMqSofamqTrace(request *ListMqSofamqTraceRequest) (_result *ListMqSofamqTraceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMqSofamqTraceResponse{}
	_body, _err := client.ListMqSofamqTraceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMqSofamqWarnWithOptions(request *ListMqSofamqWarnRequest, runtime *util.RuntimeOptions) (_result *ListMqSofamqWarnResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMqSofamqWarn"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMqSofamqWarnResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMqSofamqWarn(request *ListMqSofamqWarnRequest) (_result *ListMqSofamqWarnResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMqSofamqWarnResponse{}
	_body, _err := client.ListMqSofamqWarnWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMqSofamqWarnHistoryWithOptions(request *ListMqSofamqWarnHistoryRequest, runtime *util.RuntimeOptions) (_result *ListMqSofamqWarnHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cell)) {
		body["Cell"] = request.Cell
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.WarnId)) {
		body["WarnId"] = request.WarnId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMqSofamqWarnHistory"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMqSofamqWarnHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMqSofamqWarnHistory(request *ListMqSofamqWarnHistoryRequest) (_result *ListMqSofamqWarnHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMqSofamqWarnHistoryResponse{}
	_body, _err := client.ListMqSofamqWarnHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LogoutMsRegistryServiceWithOptions(request *LogoutMsRegistryServiceRequest, runtime *util.RuntimeOptions) (_result *LogoutMsRegistryServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerIpsRepeatList)) {
		body["ServerIpsRepeatList"] = request.ServerIpsRepeatList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("LogoutMsRegistryService"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LogoutMsRegistryServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LogoutMsRegistryService(request *LogoutMsRegistryServiceRequest) (_result *LogoutMsRegistryServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LogoutMsRegistryServiceResponse{}
	_body, _err := client.LogoutMsRegistryServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushMsConfigDataWithOptions(request *PushMsConfigDataRequest, runtime *util.RuntimeOptions) (_result *PushMsConfigDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttributeId)) {
		body["AttributeId"] = request.AttributeId
	}

	if !tea.BoolValue(util.IsUnset(request.Cells)) {
		body["Cells"] = request.Cells
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["Operator"] = request.Operator
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PushMsConfigData"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PushMsConfigDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushMsConfigData(request *PushMsConfigDataRequest) (_result *PushMsConfigDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PushMsConfigDataResponse{}
	_body, _err := client.PushMsConfigDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMqSofamqConsumerAccumulateWithOptions(request *QueryMqSofamqConsumerAccumulateRequest, runtime *util.RuntimeOptions) (_result *QueryMqSofamqConsumerAccumulateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cell)) {
		body["Cell"] = request.Cell
	}

	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		body["Detail"] = request.Detail
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMqSofamqConsumerAccumulate"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMqSofamqConsumerAccumulateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMqSofamqConsumerAccumulate(request *QueryMqSofamqConsumerAccumulateRequest) (_result *QueryMqSofamqConsumerAccumulateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMqSofamqConsumerAccumulateResponse{}
	_body, _err := client.QueryMqSofamqConsumerAccumulateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMqSofamqConsumerConnectionWithOptions(request *QueryMqSofamqConsumerConnectionRequest, runtime *util.RuntimeOptions) (_result *QueryMqSofamqConsumerConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cell)) {
		body["Cell"] = request.Cell
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMqSofamqConsumerConnection"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMqSofamqConsumerConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMqSofamqConsumerConnection(request *QueryMqSofamqConsumerConnectionRequest) (_result *QueryMqSofamqConsumerConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMqSofamqConsumerConnectionResponse{}
	_body, _err := client.QueryMqSofamqConsumerConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMqSofamqConsumerTimespanWithOptions(request *QueryMqSofamqConsumerTimespanRequest, runtime *util.RuntimeOptions) (_result *QueryMqSofamqConsumerTimespanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cell)) {
		body["Cell"] = request.Cell
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMqSofamqConsumerTimespan"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMqSofamqConsumerTimespanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMqSofamqConsumerTimespan(request *QueryMqSofamqConsumerTimespanRequest) (_result *QueryMqSofamqConsumerTimespanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMqSofamqConsumerTimespanResponse{}
	_body, _err := client.QueryMqSofamqConsumerTimespanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMqSofamqDLQMessageByGroupIdWithOptions(request *QueryMqSofamqDLQMessageByGroupIdRequest, runtime *util.RuntimeOptions) (_result *QueryMqSofamqDLQMessageByGroupIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		body["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.Cell)) {
		body["Cell"] = request.Cell
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMqSofamqDLQMessageByGroupId"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMqSofamqDLQMessageByGroupIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMqSofamqDLQMessageByGroupId(request *QueryMqSofamqDLQMessageByGroupIdRequest) (_result *QueryMqSofamqDLQMessageByGroupIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMqSofamqDLQMessageByGroupIdResponse{}
	_body, _err := client.QueryMqSofamqDLQMessageByGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMqSofamqGroupSubDetailWithOptions(request *QueryMqSofamqGroupSubDetailRequest, runtime *util.RuntimeOptions) (_result *QueryMqSofamqGroupSubDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cell)) {
		body["Cell"] = request.Cell
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMqSofamqGroupSubDetail"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMqSofamqGroupSubDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMqSofamqGroupSubDetail(request *QueryMqSofamqGroupSubDetailRequest) (_result *QueryMqSofamqGroupSubDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMqSofamqGroupSubDetailResponse{}
	_body, _err := client.QueryMqSofamqGroupSubDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMqSofamqMessageByKeyWithOptions(request *QueryMqSofamqMessageByKeyRequest, runtime *util.RuntimeOptions) (_result *QueryMqSofamqMessageByKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cell)) {
		body["Cell"] = request.Cell
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		body["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMqSofamqMessageByKey"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMqSofamqMessageByKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMqSofamqMessageByKey(request *QueryMqSofamqMessageByKeyRequest) (_result *QueryMqSofamqMessageByKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMqSofamqMessageByKeyResponse{}
	_body, _err := client.QueryMqSofamqMessageByKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMqSofamqMessageByTopicWithOptions(request *QueryMqSofamqMessageByTopicRequest, runtime *util.RuntimeOptions) (_result *QueryMqSofamqMessageByTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		body["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.Cell)) {
		body["Cell"] = request.Cell
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMqSofamqMessageByTopic"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMqSofamqMessageByTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMqSofamqMessageByTopic(request *QueryMqSofamqMessageByTopicRequest) (_result *QueryMqSofamqMessageByTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMqSofamqMessageByTopicResponse{}
	_body, _err := client.QueryMqSofamqMessageByTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMqSofamqTraceByMsgKeyWithOptions(request *QueryMqSofamqTraceByMsgKeyRequest, runtime *util.RuntimeOptions) (_result *QueryMqSofamqTraceByMsgKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		body["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.Cell)) {
		body["Cell"] = request.Cell
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgKey)) {
		body["MsgKey"] = request.MsgKey
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMqSofamqTraceByMsgKey"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMqSofamqTraceByMsgKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMqSofamqTraceByMsgKey(request *QueryMqSofamqTraceByMsgKeyRequest) (_result *QueryMqSofamqTraceByMsgKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMqSofamqTraceByMsgKeyResponse{}
	_body, _err := client.QueryMqSofamqTraceByMsgKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMqSofamqTraceByTopicWithOptions(request *QueryMqSofamqTraceByTopicRequest, runtime *util.RuntimeOptions) (_result *QueryMqSofamqTraceByTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		body["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.Cell)) {
		body["Cell"] = request.Cell
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMqSofamqTraceByTopic"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMqSofamqTraceByTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMqSofamqTraceByTopic(request *QueryMqSofamqTraceByTopicRequest) (_result *QueryMqSofamqTraceByTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMqSofamqTraceByTopicResponse{}
	_body, _err := client.QueryMqSofamqTraceByTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMsConfigAttributesWithOptions(request *QueryMsConfigAttributesRequest, runtime *util.RuntimeOptions) (_result *QueryMsConfigAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AttributeName)) {
		body["AttributeName"] = request.AttributeName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMsConfigAttributes"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMsConfigAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMsConfigAttributes(request *QueryMsConfigAttributesRequest) (_result *QueryMsConfigAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMsConfigAttributesResponse{}
	_body, _err := client.QueryMsConfigAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMsConfigClientValuesWithOptions(request *QueryMsConfigClientValuesRequest, runtime *util.RuntimeOptions) (_result *QueryMsConfigClientValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttributeId)) {
		body["AttributeId"] = request.AttributeId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ips)) {
		body["Ips"] = request.Ips
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMsConfigClientValues"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMsConfigClientValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMsConfigClientValues(request *QueryMsConfigClientValuesRequest) (_result *QueryMsConfigClientValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMsConfigClientValuesResponse{}
	_body, _err := client.QueryMsConfigClientValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMsConfigClientsWithOptions(request *QueryMsConfigClientsRequest, runtime *util.RuntimeOptions) (_result *QueryMsConfigClientsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttributeId)) {
		body["AttributeId"] = request.AttributeId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMsConfigClients"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMsConfigClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMsConfigClients(request *QueryMsConfigClientsRequest) (_result *QueryMsConfigClientsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMsConfigClientsResponse{}
	_body, _err := client.QueryMsConfigClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMsConfigDataWithOptions(request *QueryMsConfigDataRequest, runtime *util.RuntimeOptions) (_result *QueryMsConfigDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttributeId)) {
		body["AttributeId"] = request.AttributeId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMsConfigData"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMsConfigDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMsConfigData(request *QueryMsConfigDataRequest) (_result *QueryMsConfigDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMsConfigDataResponse{}
	_body, _err := client.QueryMsConfigDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMsConfigResourcesWithOptions(request *QueryMsConfigResourcesRequest, runtime *util.RuntimeOptions) (_result *QueryMsConfigResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMsConfigResources"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMsConfigResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMsConfigResources(request *QueryMsConfigResourcesRequest) (_result *QueryMsConfigResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMsConfigResourcesResponse{}
	_body, _err := client.QueryMsConfigResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRMSMetricsWithOptions(request *QueryRMSMetricsRequest, runtime *util.RuntimeOptions) (_result *QueryRMSMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContentType)) {
		body["ContentType"] = request.ContentType
	}

	if !tea.BoolValue(util.IsUnset(request.DsId)) {
		body["DsId"] = request.DsId
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		body["End"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.FieldsRepeatList)) {
		body["FieldsRepeatList"] = request.FieldsRepeatList
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodType)) {
		body["PeriodType"] = request.PeriodType
	}

	if !tea.BoolValue(util.IsUnset(request.Plugin)) {
		body["Plugin"] = request.Plugin
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		body["Start"] = request.Start
	}

	if !tea.BoolValue(util.IsUnset(request.Where)) {
		body["Where"] = request.Where
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceName)) {
		body["WorkspaceName"] = request.WorkspaceName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRMSMetrics"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRMSMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRMSMetrics(request *QueryRMSMetricsRequest) (_result *QueryRMSMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRMSMetricsResponse{}
	_body, _err := client.QueryRMSMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRMSUnifiedAlarmEventWithOptions(request *QueryRMSUnifiedAlarmEventRequest, runtime *util.RuntimeOptions) (_result *QueryRMSUnifiedAlarmEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmLevel)) {
		body["AlarmLevel"] = request.AlarmLevel
	}

	if !tea.BoolValue(util.IsUnset(request.AlarmRuleId)) {
		body["AlarmRuleId"] = request.AlarmRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.AlarmRuleUuid)) {
		body["AlarmRuleUuid"] = request.AlarmRuleUuid
	}

	if !tea.BoolValue(util.IsUnset(request.AlarmStackInfoJsonStr)) {
		body["AlarmStackInfoJsonStr"] = request.AlarmStackInfoJsonStr
	}

	if !tea.BoolValue(util.IsUnset(request.AlarmStatusRepeatList)) {
		body["AlarmStatusRepeatList"] = request.AlarmStatusRepeatList
	}

	if !tea.BoolValue(util.IsUnset(request.AlarmTargetKeyword)) {
		body["AlarmTargetKeyword"] = request.AlarmTargetKeyword
	}

	if !tea.BoolValue(util.IsUnset(request.AlarmTargetType)) {
		body["AlarmTargetType"] = request.AlarmTargetType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceName)) {
		body["WorkspaceName"] = request.WorkspaceName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRMSUnifiedAlarmEvent"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRMSUnifiedAlarmEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRMSUnifiedAlarmEvent(request *QueryRMSUnifiedAlarmEventRequest) (_result *QueryRMSUnifiedAlarmEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRMSUnifiedAlarmEventResponse{}
	_body, _err := client.QueryRMSUnifiedAlarmEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRMSUnifiedAlarmNotifyHistoryWithOptions(request *QueryRMSUnifiedAlarmNotifyHistoryRequest, runtime *util.RuntimeOptions) (_result *QueryRMSUnifiedAlarmNotifyHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmRuleId)) {
		body["AlarmRuleId"] = request.AlarmRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.AlarmStackInfoJsonStr)) {
		body["AlarmStackInfoJsonStr"] = request.AlarmStackInfoJsonStr
	}

	if !tea.BoolValue(util.IsUnset(request.AlarmStatus)) {
		body["AlarmStatus"] = request.AlarmStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AlarmSubscribers)) {
		body["AlarmSubscribers"] = request.AlarmSubscribers
	}

	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["Channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EventId)) {
		body["EventId"] = request.EventId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Subscriber)) {
		body["Subscriber"] = request.Subscriber
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceName)) {
		body["WorkspaceName"] = request.WorkspaceName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRMSUnifiedAlarmNotifyHistory"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRMSUnifiedAlarmNotifyHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRMSUnifiedAlarmNotifyHistory(request *QueryRMSUnifiedAlarmNotifyHistoryRequest) (_result *QueryRMSUnifiedAlarmNotifyHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRMSUnifiedAlarmNotifyHistoryResponse{}
	_body, _err := client.QueryRMSUnifiedAlarmNotifyHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRMSUnifiedAlarmRuleWithOptions(request *QueryRMSUnifiedAlarmRuleRequest, runtime *util.RuntimeOptions) (_result *QueryRMSUnifiedAlarmRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmStatus)) {
		body["AlarmStatus"] = request.AlarmStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AlarmTargetJsonStr)) {
		body["AlarmTargetJsonStr"] = request.AlarmTargetJsonStr
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.OnlyMe)) {
		body["OnlyMe"] = request.OnlyMe
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RefDatasourceType)) {
		body["RefDatasourceType"] = request.RefDatasourceType
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		body["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleStatus)) {
		body["RuleStatus"] = request.RuleStatus
	}

	if !tea.BoolValue(util.IsUnset(request.RuleUniqueIdentity)) {
		body["RuleUniqueIdentity"] = request.RuleUniqueIdentity
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceName)) {
		body["WorkspaceName"] = request.WorkspaceName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRMSUnifiedAlarmRule"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRMSUnifiedAlarmRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRMSUnifiedAlarmRule(request *QueryRMSUnifiedAlarmRuleRequest) (_result *QueryRMSUnifiedAlarmRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRMSUnifiedAlarmRuleResponse{}
	_body, _err := client.QueryRMSUnifiedAlarmRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResendMqSofamqDLQMessageBatchWithOptions(request *ResendMqSofamqDLQMessageBatchRequest, runtime *util.RuntimeOptions) (_result *ResendMqSofamqDLQMessageBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cell)) {
		body["Cell"] = request.Cell
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgIds)) {
		body["MsgIds"] = request.MsgIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResendMqSofamqDLQMessageBatch"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResendMqSofamqDLQMessageBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResendMqSofamqDLQMessageBatch(request *ResendMqSofamqDLQMessageBatchRequest) (_result *ResendMqSofamqDLQMessageBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResendMqSofamqDLQMessageBatchResponse{}
	_body, _err := client.ResendMqSofamqDLQMessageBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResendMqSofamqDLQMessageByIdWithOptions(request *ResendMqSofamqDLQMessageByIdRequest, runtime *util.RuntimeOptions) (_result *ResendMqSofamqDLQMessageByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cell)) {
		body["Cell"] = request.Cell
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgId)) {
		body["MsgId"] = request.MsgId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResendMqSofamqDLQMessageById"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResendMqSofamqDLQMessageByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResendMqSofamqDLQMessageById(request *ResendMqSofamqDLQMessageByIdRequest) (_result *ResendMqSofamqDLQMessageByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResendMqSofamqDLQMessageByIdResponse{}
	_body, _err := client.ResendMqSofamqDLQMessageByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetMqSofamqConsumerOffsetWithOptions(request *ResetMqSofamqConsumerOffsetRequest, runtime *util.RuntimeOptions) (_result *ResetMqSofamqConsumerOffsetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cell)) {
		body["Cell"] = request.Cell
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResetTimestamp)) {
		body["ResetTimestamp"] = request.ResetTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetMqSofamqConsumerOffset"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetMqSofamqConsumerOffsetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetMqSofamqConsumerOffset(request *ResetMqSofamqConsumerOffsetRequest) (_result *ResetMqSofamqConsumerOffsetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetMqSofamqConsumerOffsetResponse{}
	_body, _err := client.ResetMqSofamqConsumerOffsetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMqSofamqGroupWithOptions(request *UpdateMqSofamqGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateMqSofamqGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ReadEnable)) {
		body["ReadEnable"] = request.ReadEnable
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMqSofamqGroup"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMqSofamqGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMqSofamqGroup(request *UpdateMqSofamqGroupRequest) (_result *UpdateMqSofamqGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMqSofamqGroupResponse{}
	_body, _err := client.UpdateMqSofamqGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMqSofamqGroupRemarkWithOptions(request *UpdateMqSofamqGroupRemarkRequest, runtime *util.RuntimeOptions) (_result *UpdateMqSofamqGroupRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMqSofamqGroupRemark"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMqSofamqGroupRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMqSofamqGroupRemark(request *UpdateMqSofamqGroupRemarkRequest) (_result *UpdateMqSofamqGroupRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMqSofamqGroupRemarkResponse{}
	_body, _err := client.UpdateMqSofamqGroupRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMqSofamqTopicWithOptions(request *UpdateMqSofamqTopicRequest, runtime *util.RuntimeOptions) (_result *UpdateMqSofamqTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Perm)) {
		body["Perm"] = request.Perm
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMqSofamqTopic"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMqSofamqTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMqSofamqTopic(request *UpdateMqSofamqTopicRequest) (_result *UpdateMqSofamqTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMqSofamqTopicResponse{}
	_body, _err := client.UpdateMqSofamqTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMqSofamqTopicRemarkWithOptions(request *UpdateMqSofamqTopicRemarkRequest, runtime *util.RuntimeOptions) (_result *UpdateMqSofamqTopicRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMqSofamqTopicRemark"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMqSofamqTopicRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMqSofamqTopicRemark(request *UpdateMqSofamqTopicRemarkRequest) (_result *UpdateMqSofamqTopicRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMqSofamqTopicRemarkResponse{}
	_body, _err := client.UpdateMqSofamqTopicRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMqSofamqWarnWithOptions(request *UpdateMqSofamqWarnRequest, runtime *util.RuntimeOptions) (_result *UpdateMqSofamqWarnResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertTime)) {
		body["AlertTime"] = request.AlertTime
	}

	if !tea.BoolValue(util.IsUnset(request.Contacts)) {
		body["Contacts"] = request.Contacts
	}

	if !tea.BoolValue(util.IsUnset(request.DelayTime)) {
		body["DelayTime"] = request.DelayTime
	}

	if !tea.BoolValue(util.IsUnset(request.Frequency)) {
		body["Frequency"] = request.Frequency
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		body["Threshold"] = request.Threshold
	}

	if !tea.BoolValue(util.IsUnset(request.WarnId)) {
		body["WarnId"] = request.WarnId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMqSofamqWarn"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMqSofamqWarnResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMqSofamqWarn(request *UpdateMqSofamqWarnRequest) (_result *UpdateMqSofamqWarnResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMqSofamqWarnResponse{}
	_body, _err := client.UpdateMqSofamqWarnWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMsConfigAttributesWithOptions(request *UpdateMsConfigAttributesRequest, runtime *util.RuntimeOptions) (_result *UpdateMsConfigAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttributeName)) {
		body["AttributeName"] = request.AttributeName
	}

	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		body["Desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMsConfigAttributes"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMsConfigAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMsConfigAttributes(request *UpdateMsConfigAttributesRequest) (_result *UpdateMsConfigAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMsConfigAttributesResponse{}
	_body, _err := client.UpdateMsConfigAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMsConfigResourcesWithOptions(request *UpdateMsConfigResourcesRequest, runtime *util.RuntimeOptions) (_result *UpdateMsConfigResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Attributes)) {
		body["Attributes"] = request.Attributes
	}

	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		body["Desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateAttribute)) {
		body["UpdateAttribute"] = request.UpdateAttribute
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMsConfigResources"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMsConfigResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMsConfigResources(request *UpdateMsConfigResourcesRequest) (_result *UpdateMsConfigResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMsConfigResourcesResponse{}
	_body, _err := client.UpdateMsConfigResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRMSUnifiedAlarmRuleWithOptions(request *UpdateRMSUnifiedAlarmRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateRMSUnifiedAlarmRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmNodata)) {
		body["AlarmNodata"] = request.AlarmNodata
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelsRepeatList)) {
		body["ChannelsRepeatList"] = request.ChannelsRepeatList
	}

	if !tea.BoolValue(util.IsUnset(request.Emergency)) {
		body["Emergency"] = request.Emergency
	}

	if !tea.BoolValue(util.IsUnset(request.EmergencyUrl)) {
		body["EmergencyUrl"] = request.EmergencyUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyFiring)) {
		body["NotifyFiring"] = request.NotifyFiring
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyNodata)) {
		body["NotifyNodata"] = request.NotifyNodata
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyRecovered)) {
		body["NotifyRecovered"] = request.NotifyRecovered
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTarget)) {
		body["NotifyTarget"] = request.NotifyTarget
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTimeFilterJsonStr)) {
		body["NotifyTimeFilterJsonStr"] = request.NotifyTimeFilterJsonStr
	}

	if !tea.BoolValue(util.IsUnset(request.PendingHit)) {
		body["PendingHit"] = request.PendingHit
	}

	if !tea.BoolValue(util.IsUnset(request.RecoveredHit)) {
		body["RecoveredHit"] = request.RecoveredHit
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConfig)) {
		body["RuleConfig"] = request.RuleConfig
	}

	if !tea.BoolValue(util.IsUnset(request.SilenceTime)) {
		body["SilenceTime"] = request.SilenceTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Step)) {
		body["Step"] = request.Step
	}

	if !tea.BoolValue(util.IsUnset(request.SuspendedEndTime)) {
		body["SuspendedEndTime"] = request.SuspendedEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.SuspendedReason)) {
		body["SuspendedReason"] = request.SuspendedReason
	}

	if !tea.BoolValue(util.IsUnset(request.SuspendedStartTime)) {
		body["SuspendedStartTime"] = request.SuspendedStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceName)) {
		body["WorkspaceName"] = request.WorkspaceName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRMSUnifiedAlarmRule"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRMSUnifiedAlarmRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRMSUnifiedAlarmRule(request *UpdateRMSUnifiedAlarmRuleRequest) (_result *UpdateRMSUnifiedAlarmRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRMSUnifiedAlarmRuleResponse{}
	_body, _err := client.UpdateRMSUnifiedAlarmRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
