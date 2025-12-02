# RuleTargetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the notification target | 
**RateLimit** | Pointer to [**CreateNotificationRuleRequestTargetsInnerRateLimit**](CreateNotificationRuleRequestTargetsInnerRateLimit.md) |  | [optional] 
**RenotifyConfig** | Pointer to [**CreateNotificationRuleRequestTargetsInnerRenotifyConfig**](CreateNotificationRuleRequestTargetsInnerRenotifyConfig.md) |  | [optional] 
**Config** | [**RuleTargetsInnerConfig**](RuleTargetsInnerConfig.md) |  | 

## Methods

### NewRuleTargetsInner

`func NewRuleTargetsInner(id string, config RuleTargetsInnerConfig, ) *RuleTargetsInner`

NewRuleTargetsInner instantiates a new RuleTargetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTargetsInnerWithDefaults

`func NewRuleTargetsInnerWithDefaults() *RuleTargetsInner`

NewRuleTargetsInnerWithDefaults instantiates a new RuleTargetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RuleTargetsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleTargetsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleTargetsInner) SetId(v string)`

SetId sets Id field to given value.


### GetRateLimit

`func (o *RuleTargetsInner) GetRateLimit() CreateNotificationRuleRequestTargetsInnerRateLimit`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *RuleTargetsInner) GetRateLimitOk() (*CreateNotificationRuleRequestTargetsInnerRateLimit, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *RuleTargetsInner) SetRateLimit(v CreateNotificationRuleRequestTargetsInnerRateLimit)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *RuleTargetsInner) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetRenotifyConfig

`func (o *RuleTargetsInner) GetRenotifyConfig() CreateNotificationRuleRequestTargetsInnerRenotifyConfig`

GetRenotifyConfig returns the RenotifyConfig field if non-nil, zero value otherwise.

### GetRenotifyConfigOk

`func (o *RuleTargetsInner) GetRenotifyConfigOk() (*CreateNotificationRuleRequestTargetsInnerRenotifyConfig, bool)`

GetRenotifyConfigOk returns a tuple with the RenotifyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenotifyConfig

`func (o *RuleTargetsInner) SetRenotifyConfig(v CreateNotificationRuleRequestTargetsInnerRenotifyConfig)`

SetRenotifyConfig sets RenotifyConfig field to given value.

### HasRenotifyConfig

`func (o *RuleTargetsInner) HasRenotifyConfig() bool`

HasRenotifyConfig returns a boolean if a field has been set.

### GetConfig

`func (o *RuleTargetsInner) GetConfig() RuleTargetsInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *RuleTargetsInner) GetConfigOk() (*RuleTargetsInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *RuleTargetsInner) SetConfig(v RuleTargetsInnerConfig)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


