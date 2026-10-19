# ServiceNotificationTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Tsuga-generated identifier of this service notification target. Tsuga returns it when the target is created or listed, and clients use it to update or delete the target. It is unique across service targets. | 
**ServiceId** | **string** | Identifier of the service catalog entry this target is attached to. Tsuga sets it from the service path parameter and returns it on read and write responses. | 
**Config** | [**RuleTargetsInnerConfig**](RuleTargetsInnerConfig.md) |  | 
**RateLimit** | Pointer to [**RuleTargetsInnerRateLimit**](RuleTargetsInnerRateLimit.md) |  | [optional] 

## Methods

### NewServiceNotificationTarget

`func NewServiceNotificationTarget(id string, serviceId string, config RuleTargetsInnerConfig, ) *ServiceNotificationTarget`

NewServiceNotificationTarget instantiates a new ServiceNotificationTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceNotificationTargetWithDefaults

`func NewServiceNotificationTargetWithDefaults() *ServiceNotificationTarget`

NewServiceNotificationTargetWithDefaults instantiates a new ServiceNotificationTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceNotificationTarget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceNotificationTarget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceNotificationTarget) SetId(v string)`

SetId sets Id field to given value.


### GetServiceId

`func (o *ServiceNotificationTarget) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServiceNotificationTarget) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServiceNotificationTarget) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetConfig

`func (o *ServiceNotificationTarget) GetConfig() RuleTargetsInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ServiceNotificationTarget) GetConfigOk() (*RuleTargetsInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ServiceNotificationTarget) SetConfig(v RuleTargetsInnerConfig)`

SetConfig sets Config field to given value.


### GetRateLimit

`func (o *ServiceNotificationTarget) GetRateLimit() RuleTargetsInnerRateLimit`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *ServiceNotificationTarget) GetRateLimitOk() (*RuleTargetsInnerRateLimit, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *ServiceNotificationTarget) SetRateLimit(v RuleTargetsInnerRateLimit)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *ServiceNotificationTarget) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


