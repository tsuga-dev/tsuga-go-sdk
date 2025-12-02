# CreateNotificationRuleRequestTargetsInnerRateLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxMessages** | **int32** | Maximum number of messages allowed during the rate-limiting window | 
**Minutes** | **int32** | Length of the rate-limiting window in minutes | 

## Methods

### NewCreateNotificationRuleRequestTargetsInnerRateLimit

`func NewCreateNotificationRuleRequestTargetsInnerRateLimit(maxMessages int32, minutes int32, ) *CreateNotificationRuleRequestTargetsInnerRateLimit`

NewCreateNotificationRuleRequestTargetsInnerRateLimit instantiates a new CreateNotificationRuleRequestTargetsInnerRateLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationRuleRequestTargetsInnerRateLimitWithDefaults

`func NewCreateNotificationRuleRequestTargetsInnerRateLimitWithDefaults() *CreateNotificationRuleRequestTargetsInnerRateLimit`

NewCreateNotificationRuleRequestTargetsInnerRateLimitWithDefaults instantiates a new CreateNotificationRuleRequestTargetsInnerRateLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxMessages

`func (o *CreateNotificationRuleRequestTargetsInnerRateLimit) GetMaxMessages() int32`

GetMaxMessages returns the MaxMessages field if non-nil, zero value otherwise.

### GetMaxMessagesOk

`func (o *CreateNotificationRuleRequestTargetsInnerRateLimit) GetMaxMessagesOk() (*int32, bool)`

GetMaxMessagesOk returns a tuple with the MaxMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMessages

`func (o *CreateNotificationRuleRequestTargetsInnerRateLimit) SetMaxMessages(v int32)`

SetMaxMessages sets MaxMessages field to given value.


### GetMinutes

`func (o *CreateNotificationRuleRequestTargetsInnerRateLimit) GetMinutes() int32`

GetMinutes returns the Minutes field if non-nil, zero value otherwise.

### GetMinutesOk

`func (o *CreateNotificationRuleRequestTargetsInnerRateLimit) GetMinutesOk() (*int32, bool)`

GetMinutesOk returns a tuple with the Minutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutes

`func (o *CreateNotificationRuleRequestTargetsInnerRateLimit) SetMinutes(v int32)`

SetMinutes sets Minutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


