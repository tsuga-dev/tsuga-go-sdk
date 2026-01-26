# RuleTargetInputGoogleChat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**IntegrationId** | **string** | Identifier of the Google Chat integration to use | 
**HideTransition** | Pointer to **bool** | When true, the transition info (e.g., \&quot;from ok to alert\&quot;) is hidden from the Google Chat message | [optional] 
**HideTime** | Pointer to **bool** | When true, the timestamp is hidden from the Google Chat message | [optional] 

## Methods

### NewRuleTargetInputGoogleChat

`func NewRuleTargetInputGoogleChat(type_ string, integrationId string, ) *RuleTargetInputGoogleChat`

NewRuleTargetInputGoogleChat instantiates a new RuleTargetInputGoogleChat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTargetInputGoogleChatWithDefaults

`func NewRuleTargetInputGoogleChatWithDefaults() *RuleTargetInputGoogleChat`

NewRuleTargetInputGoogleChatWithDefaults instantiates a new RuleTargetInputGoogleChat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuleTargetInputGoogleChat) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleTargetInputGoogleChat) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleTargetInputGoogleChat) SetType(v string)`

SetType sets Type field to given value.


### GetIntegrationId

`func (o *RuleTargetInputGoogleChat) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *RuleTargetInputGoogleChat) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *RuleTargetInputGoogleChat) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetHideTransition

`func (o *RuleTargetInputGoogleChat) GetHideTransition() bool`

GetHideTransition returns the HideTransition field if non-nil, zero value otherwise.

### GetHideTransitionOk

`func (o *RuleTargetInputGoogleChat) GetHideTransitionOk() (*bool, bool)`

GetHideTransitionOk returns a tuple with the HideTransition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideTransition

`func (o *RuleTargetInputGoogleChat) SetHideTransition(v bool)`

SetHideTransition sets HideTransition field to given value.

### HasHideTransition

`func (o *RuleTargetInputGoogleChat) HasHideTransition() bool`

HasHideTransition returns a boolean if a field has been set.

### GetHideTime

`func (o *RuleTargetInputGoogleChat) GetHideTime() bool`

GetHideTime returns the HideTime field if non-nil, zero value otherwise.

### GetHideTimeOk

`func (o *RuleTargetInputGoogleChat) GetHideTimeOk() (*bool, bool)`

GetHideTimeOk returns a tuple with the HideTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideTime

`func (o *RuleTargetInputGoogleChat) SetHideTime(v bool)`

SetHideTime sets HideTime field to given value.

### HasHideTime

`func (o *RuleTargetInputGoogleChat) HasHideTime() bool`

HasHideTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


