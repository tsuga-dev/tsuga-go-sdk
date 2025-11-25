# ListNotificationRules200ResponseDataInnerTargetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the notification target | 
**RateLimit** | Pointer to [**ListNotificationRules200ResponseDataInnerTargetsInnerRateLimit**](ListNotificationRules200ResponseDataInnerTargetsInnerRateLimit.md) |  | [optional] 
**RenotifyConfig** | Pointer to [**ListNotificationRules200ResponseDataInnerTargetsInnerRenotifyConfig**](ListNotificationRules200ResponseDataInnerTargetsInnerRenotifyConfig.md) |  | [optional] 
**Config** | [**ListNotificationRules200ResponseDataInnerTargetsInnerConfig**](ListNotificationRules200ResponseDataInnerTargetsInnerConfig.md) |  | 

## Methods

### NewListNotificationRules200ResponseDataInnerTargetsInner

`func NewListNotificationRules200ResponseDataInnerTargetsInner(id string, config ListNotificationRules200ResponseDataInnerTargetsInnerConfig, ) *ListNotificationRules200ResponseDataInnerTargetsInner`

NewListNotificationRules200ResponseDataInnerTargetsInner instantiates a new ListNotificationRules200ResponseDataInnerTargetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNotificationRules200ResponseDataInnerTargetsInnerWithDefaults

`func NewListNotificationRules200ResponseDataInnerTargetsInnerWithDefaults() *ListNotificationRules200ResponseDataInnerTargetsInner`

NewListNotificationRules200ResponseDataInnerTargetsInnerWithDefaults instantiates a new ListNotificationRules200ResponseDataInnerTargetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListNotificationRules200ResponseDataInnerTargetsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListNotificationRules200ResponseDataInnerTargetsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListNotificationRules200ResponseDataInnerTargetsInner) SetId(v string)`

SetId sets Id field to given value.


### GetRateLimit

`func (o *ListNotificationRules200ResponseDataInnerTargetsInner) GetRateLimit() ListNotificationRules200ResponseDataInnerTargetsInnerRateLimit`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *ListNotificationRules200ResponseDataInnerTargetsInner) GetRateLimitOk() (*ListNotificationRules200ResponseDataInnerTargetsInnerRateLimit, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *ListNotificationRules200ResponseDataInnerTargetsInner) SetRateLimit(v ListNotificationRules200ResponseDataInnerTargetsInnerRateLimit)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *ListNotificationRules200ResponseDataInnerTargetsInner) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetRenotifyConfig

`func (o *ListNotificationRules200ResponseDataInnerTargetsInner) GetRenotifyConfig() ListNotificationRules200ResponseDataInnerTargetsInnerRenotifyConfig`

GetRenotifyConfig returns the RenotifyConfig field if non-nil, zero value otherwise.

### GetRenotifyConfigOk

`func (o *ListNotificationRules200ResponseDataInnerTargetsInner) GetRenotifyConfigOk() (*ListNotificationRules200ResponseDataInnerTargetsInnerRenotifyConfig, bool)`

GetRenotifyConfigOk returns a tuple with the RenotifyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenotifyConfig

`func (o *ListNotificationRules200ResponseDataInnerTargetsInner) SetRenotifyConfig(v ListNotificationRules200ResponseDataInnerTargetsInnerRenotifyConfig)`

SetRenotifyConfig sets RenotifyConfig field to given value.

### HasRenotifyConfig

`func (o *ListNotificationRules200ResponseDataInnerTargetsInner) HasRenotifyConfig() bool`

HasRenotifyConfig returns a boolean if a field has been set.

### GetConfig

`func (o *ListNotificationRules200ResponseDataInnerTargetsInner) GetConfig() ListNotificationRules200ResponseDataInnerTargetsInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListNotificationRules200ResponseDataInnerTargetsInner) GetConfigOk() (*ListNotificationRules200ResponseDataInnerTargetsInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListNotificationRules200ResponseDataInnerTargetsInner) SetConfig(v ListNotificationRules200ResponseDataInnerTargetsInnerConfig)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


