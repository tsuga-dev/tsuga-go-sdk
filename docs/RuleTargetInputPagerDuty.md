# RuleTargetInputPagerDuty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | PagerDuty target backed by a configured PagerDuty integration. | 
**IntegrationId** | **string** | Identifier of the PagerDuty integration to use | 
**RenotifyConfig** | Pointer to [**RuleTargetInputSlackRenotifyConfig**](RuleTargetInputSlackRenotifyConfig.md) |  | [optional] 

## Methods

### NewRuleTargetInputPagerDuty

`func NewRuleTargetInputPagerDuty(type_ string, integrationId string, ) *RuleTargetInputPagerDuty`

NewRuleTargetInputPagerDuty instantiates a new RuleTargetInputPagerDuty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTargetInputPagerDutyWithDefaults

`func NewRuleTargetInputPagerDutyWithDefaults() *RuleTargetInputPagerDuty`

NewRuleTargetInputPagerDutyWithDefaults instantiates a new RuleTargetInputPagerDuty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuleTargetInputPagerDuty) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleTargetInputPagerDuty) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleTargetInputPagerDuty) SetType(v string)`

SetType sets Type field to given value.


### GetIntegrationId

`func (o *RuleTargetInputPagerDuty) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *RuleTargetInputPagerDuty) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *RuleTargetInputPagerDuty) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetRenotifyConfig

`func (o *RuleTargetInputPagerDuty) GetRenotifyConfig() RuleTargetInputSlackRenotifyConfig`

GetRenotifyConfig returns the RenotifyConfig field if non-nil, zero value otherwise.

### GetRenotifyConfigOk

`func (o *RuleTargetInputPagerDuty) GetRenotifyConfigOk() (*RuleTargetInputSlackRenotifyConfig, bool)`

GetRenotifyConfigOk returns a tuple with the RenotifyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenotifyConfig

`func (o *RuleTargetInputPagerDuty) SetRenotifyConfig(v RuleTargetInputSlackRenotifyConfig)`

SetRenotifyConfig sets RenotifyConfig field to given value.

### HasRenotifyConfig

`func (o *RuleTargetInputPagerDuty) HasRenotifyConfig() bool`

HasRenotifyConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


