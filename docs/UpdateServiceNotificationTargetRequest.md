# UpdateServiceNotificationTargetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**CreateNotificationRuleRequestTargetsInnerConfig**](CreateNotificationRuleRequestTargetsInnerConfig.md) |  | 
**RateLimit** | [**NullableCreateServiceNotificationTargetRequestRateLimit**](CreateServiceNotificationTargetRequestRateLimit.md) |  | 

## Methods

### NewUpdateServiceNotificationTargetRequest

`func NewUpdateServiceNotificationTargetRequest(config CreateNotificationRuleRequestTargetsInnerConfig, rateLimit NullableCreateServiceNotificationTargetRequestRateLimit, ) *UpdateServiceNotificationTargetRequest`

NewUpdateServiceNotificationTargetRequest instantiates a new UpdateServiceNotificationTargetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServiceNotificationTargetRequestWithDefaults

`func NewUpdateServiceNotificationTargetRequestWithDefaults() *UpdateServiceNotificationTargetRequest`

NewUpdateServiceNotificationTargetRequestWithDefaults instantiates a new UpdateServiceNotificationTargetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *UpdateServiceNotificationTargetRequest) GetConfig() CreateNotificationRuleRequestTargetsInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateServiceNotificationTargetRequest) GetConfigOk() (*CreateNotificationRuleRequestTargetsInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateServiceNotificationTargetRequest) SetConfig(v CreateNotificationRuleRequestTargetsInnerConfig)`

SetConfig sets Config field to given value.


### GetRateLimit

`func (o *UpdateServiceNotificationTargetRequest) GetRateLimit() CreateServiceNotificationTargetRequestRateLimit`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *UpdateServiceNotificationTargetRequest) GetRateLimitOk() (*CreateServiceNotificationTargetRequestRateLimit, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *UpdateServiceNotificationTargetRequest) SetRateLimit(v CreateServiceNotificationTargetRequestRateLimit)`

SetRateLimit sets RateLimit field to given value.


### SetRateLimitNil

`func (o *UpdateServiceNotificationTargetRequest) SetRateLimitNil(b bool)`

 SetRateLimitNil sets the value for RateLimit to be an explicit nil

### UnsetRateLimit
`func (o *UpdateServiceNotificationTargetRequest) UnsetRateLimit()`

UnsetRateLimit ensures that no value is present for RateLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


