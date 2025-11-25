# RuleTargetConfigSlack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Channel** | **string** | Slack channel ID that receives the notification | 
**IntegrationId** | **string** | Slack workspace ID that receives the notification | 
**IntegrationName** | **string** | Human readable name of the Slack integration | 

## Methods

### NewRuleTargetConfigSlack

`func NewRuleTargetConfigSlack(type_ string, channel string, integrationId string, integrationName string, ) *RuleTargetConfigSlack`

NewRuleTargetConfigSlack instantiates a new RuleTargetConfigSlack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTargetConfigSlackWithDefaults

`func NewRuleTargetConfigSlackWithDefaults() *RuleTargetConfigSlack`

NewRuleTargetConfigSlackWithDefaults instantiates a new RuleTargetConfigSlack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuleTargetConfigSlack) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleTargetConfigSlack) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleTargetConfigSlack) SetType(v string)`

SetType sets Type field to given value.


### GetChannel

`func (o *RuleTargetConfigSlack) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *RuleTargetConfigSlack) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *RuleTargetConfigSlack) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetIntegrationId

`func (o *RuleTargetConfigSlack) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *RuleTargetConfigSlack) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *RuleTargetConfigSlack) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetIntegrationName

`func (o *RuleTargetConfigSlack) GetIntegrationName() string`

GetIntegrationName returns the IntegrationName field if non-nil, zero value otherwise.

### GetIntegrationNameOk

`func (o *RuleTargetConfigSlack) GetIntegrationNameOk() (*string, bool)`

GetIntegrationNameOk returns a tuple with the IntegrationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationName

`func (o *RuleTargetConfigSlack) SetIntegrationName(v string)`

SetIntegrationName sets IntegrationName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


