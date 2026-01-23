# CreateNotificationRuleRequestTargetsInnerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Channel** | **string** | Slack channel ID that receives the notification | 
**IntegrationId** | **string** | Identifier of the Google Chat integration to use | 
**HideTransition** | Pointer to **bool** | When true, the transition info (e.g., \&quot;from ok to alert\&quot;) is hidden from the Google Chat message | [optional] 
**HideTime** | Pointer to **bool** | When true, the timestamp is hidden from the Google Chat message | [optional] 
**Addresses** | **[]string** | Email addresses that will receive the alert | 

## Methods

### NewCreateNotificationRuleRequestTargetsInnerConfig

`func NewCreateNotificationRuleRequestTargetsInnerConfig(type_ string, channel string, integrationId string, addresses []string, ) *CreateNotificationRuleRequestTargetsInnerConfig`

NewCreateNotificationRuleRequestTargetsInnerConfig instantiates a new CreateNotificationRuleRequestTargetsInnerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationRuleRequestTargetsInnerConfigWithDefaults

`func NewCreateNotificationRuleRequestTargetsInnerConfigWithDefaults() *CreateNotificationRuleRequestTargetsInnerConfig`

NewCreateNotificationRuleRequestTargetsInnerConfigWithDefaults instantiates a new CreateNotificationRuleRequestTargetsInnerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) SetType(v string)`

SetType sets Type field to given value.


### GetChannel

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetIntegrationId

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetHideTransition

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetHideTransition() bool`

GetHideTransition returns the HideTransition field if non-nil, zero value otherwise.

### GetHideTransitionOk

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetHideTransitionOk() (*bool, bool)`

GetHideTransitionOk returns a tuple with the HideTransition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideTransition

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) SetHideTransition(v bool)`

SetHideTransition sets HideTransition field to given value.

### HasHideTransition

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) HasHideTransition() bool`

HasHideTransition returns a boolean if a field has been set.

### GetHideTime

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetHideTime() bool`

GetHideTime returns the HideTime field if non-nil, zero value otherwise.

### GetHideTimeOk

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetHideTimeOk() (*bool, bool)`

GetHideTimeOk returns a tuple with the HideTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideTime

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) SetHideTime(v bool)`

SetHideTime sets HideTime field to given value.

### HasHideTime

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) HasHideTime() bool`

HasHideTime returns a boolean if a field has been set.

### GetAddresses

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


