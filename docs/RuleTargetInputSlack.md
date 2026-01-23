# RuleTargetInputSlack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Channel** | **string** | Slack channel ID that receives the notification | 
**IntegrationId** | **string** | Slack workspace ID that receives the notification | 
**HideTransition** | Pointer to **bool** | When true, the transition info (e.g., \&quot;from ok to alert\&quot;) is hidden from the Slack message | [optional] 
**HideTime** | Pointer to **bool** | When true, the timestamp is hidden from the Slack message | [optional] 

## Methods

### NewRuleTargetInputSlack

`func NewRuleTargetInputSlack(type_ string, channel string, integrationId string, ) *RuleTargetInputSlack`

NewRuleTargetInputSlack instantiates a new RuleTargetInputSlack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTargetInputSlackWithDefaults

`func NewRuleTargetInputSlackWithDefaults() *RuleTargetInputSlack`

NewRuleTargetInputSlackWithDefaults instantiates a new RuleTargetInputSlack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuleTargetInputSlack) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleTargetInputSlack) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleTargetInputSlack) SetType(v string)`

SetType sets Type field to given value.


### GetChannel

`func (o *RuleTargetInputSlack) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *RuleTargetInputSlack) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *RuleTargetInputSlack) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetIntegrationId

`func (o *RuleTargetInputSlack) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *RuleTargetInputSlack) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *RuleTargetInputSlack) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetHideTransition

`func (o *RuleTargetInputSlack) GetHideTransition() bool`

GetHideTransition returns the HideTransition field if non-nil, zero value otherwise.

### GetHideTransitionOk

`func (o *RuleTargetInputSlack) GetHideTransitionOk() (*bool, bool)`

GetHideTransitionOk returns a tuple with the HideTransition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideTransition

`func (o *RuleTargetInputSlack) SetHideTransition(v bool)`

SetHideTransition sets HideTransition field to given value.

### HasHideTransition

`func (o *RuleTargetInputSlack) HasHideTransition() bool`

HasHideTransition returns a boolean if a field has been set.

### GetHideTime

`func (o *RuleTargetInputSlack) GetHideTime() bool`

GetHideTime returns the HideTime field if non-nil, zero value otherwise.

### GetHideTimeOk

`func (o *RuleTargetInputSlack) GetHideTimeOk() (*bool, bool)`

GetHideTimeOk returns a tuple with the HideTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideTime

`func (o *RuleTargetInputSlack) SetHideTime(v bool)`

SetHideTime sets HideTime field to given value.

### HasHideTime

`func (o *RuleTargetInputSlack) HasHideTime() bool`

HasHideTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


