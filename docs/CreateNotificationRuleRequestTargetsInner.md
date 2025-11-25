# CreateNotificationRuleRequestTargetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the notification target | 
**RateLimit** | Pointer to [**ListNotificationRules200ResponseDataInnerTargetsInnerRateLimit**](ListNotificationRules200ResponseDataInnerTargetsInnerRateLimit.md) |  | [optional] 
**Config** | [**CreateNotificationRuleRequestTargetsInnerConfig**](CreateNotificationRuleRequestTargetsInnerConfig.md) |  | 
**RenotifyConfig** | Pointer to [**ListNotificationRules200ResponseDataInnerTargetsInnerRenotifyConfig**](ListNotificationRules200ResponseDataInnerTargetsInnerRenotifyConfig.md) |  | [optional] 

## Methods

### NewCreateNotificationRuleRequestTargetsInner

`func NewCreateNotificationRuleRequestTargetsInner(id string, config CreateNotificationRuleRequestTargetsInnerConfig, ) *CreateNotificationRuleRequestTargetsInner`

NewCreateNotificationRuleRequestTargetsInner instantiates a new CreateNotificationRuleRequestTargetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationRuleRequestTargetsInnerWithDefaults

`func NewCreateNotificationRuleRequestTargetsInnerWithDefaults() *CreateNotificationRuleRequestTargetsInner`

NewCreateNotificationRuleRequestTargetsInnerWithDefaults instantiates a new CreateNotificationRuleRequestTargetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateNotificationRuleRequestTargetsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNotificationRuleRequestTargetsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNotificationRuleRequestTargetsInner) SetId(v string)`

SetId sets Id field to given value.


### GetRateLimit

`func (o *CreateNotificationRuleRequestTargetsInner) GetRateLimit() ListNotificationRules200ResponseDataInnerTargetsInnerRateLimit`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *CreateNotificationRuleRequestTargetsInner) GetRateLimitOk() (*ListNotificationRules200ResponseDataInnerTargetsInnerRateLimit, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *CreateNotificationRuleRequestTargetsInner) SetRateLimit(v ListNotificationRules200ResponseDataInnerTargetsInnerRateLimit)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *CreateNotificationRuleRequestTargetsInner) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetConfig

`func (o *CreateNotificationRuleRequestTargetsInner) GetConfig() CreateNotificationRuleRequestTargetsInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateNotificationRuleRequestTargetsInner) GetConfigOk() (*CreateNotificationRuleRequestTargetsInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateNotificationRuleRequestTargetsInner) SetConfig(v CreateNotificationRuleRequestTargetsInnerConfig)`

SetConfig sets Config field to given value.


### GetRenotifyConfig

`func (o *CreateNotificationRuleRequestTargetsInner) GetRenotifyConfig() ListNotificationRules200ResponseDataInnerTargetsInnerRenotifyConfig`

GetRenotifyConfig returns the RenotifyConfig field if non-nil, zero value otherwise.

### GetRenotifyConfigOk

`func (o *CreateNotificationRuleRequestTargetsInner) GetRenotifyConfigOk() (*ListNotificationRules200ResponseDataInnerTargetsInnerRenotifyConfig, bool)`

GetRenotifyConfigOk returns a tuple with the RenotifyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenotifyConfig

`func (o *CreateNotificationRuleRequestTargetsInner) SetRenotifyConfig(v ListNotificationRules200ResponseDataInnerTargetsInnerRenotifyConfig)`

SetRenotifyConfig sets RenotifyConfig field to given value.

### HasRenotifyConfig

`func (o *CreateNotificationRuleRequestTargetsInner) HasRenotifyConfig() bool`

HasRenotifyConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


