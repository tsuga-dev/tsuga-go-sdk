# CreateServiceNotificationTargetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**CreateNotificationRuleRequestTargetsInnerConfig**](CreateNotificationRuleRequestTargetsInnerConfig.md) |  | 
**RateLimit** | Pointer to [**NullableCreateServiceNotificationTargetRequestRateLimit**](CreateServiceNotificationTargetRequestRateLimit.md) |  | [optional] 

## Methods

### NewCreateServiceNotificationTargetRequest

`func NewCreateServiceNotificationTargetRequest(config CreateNotificationRuleRequestTargetsInnerConfig, ) *CreateServiceNotificationTargetRequest`

NewCreateServiceNotificationTargetRequest instantiates a new CreateServiceNotificationTargetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceNotificationTargetRequestWithDefaults

`func NewCreateServiceNotificationTargetRequestWithDefaults() *CreateServiceNotificationTargetRequest`

NewCreateServiceNotificationTargetRequestWithDefaults instantiates a new CreateServiceNotificationTargetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *CreateServiceNotificationTargetRequest) GetConfig() CreateNotificationRuleRequestTargetsInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateServiceNotificationTargetRequest) GetConfigOk() (*CreateNotificationRuleRequestTargetsInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateServiceNotificationTargetRequest) SetConfig(v CreateNotificationRuleRequestTargetsInnerConfig)`

SetConfig sets Config field to given value.


### GetRateLimit

`func (o *CreateServiceNotificationTargetRequest) GetRateLimit() CreateServiceNotificationTargetRequestRateLimit`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *CreateServiceNotificationTargetRequest) GetRateLimitOk() (*CreateServiceNotificationTargetRequestRateLimit, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *CreateServiceNotificationTargetRequest) SetRateLimit(v CreateServiceNotificationTargetRequestRateLimit)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *CreateServiceNotificationTargetRequest) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### SetRateLimitNil

`func (o *CreateServiceNotificationTargetRequest) SetRateLimitNil(b bool)`

 SetRateLimitNil sets the value for RateLimit to be an explicit nil

### UnsetRateLimit
`func (o *CreateServiceNotificationTargetRequest) UnsetRateLimit()`

UnsetRateLimit ensures that no value is present for RateLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


