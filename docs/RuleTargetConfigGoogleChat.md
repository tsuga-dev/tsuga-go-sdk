# RuleTargetConfigGoogleChat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**IntegrationId** | **string** | Identifier of the Google Chat integration to use | 
**IntegrationName** | **string** | Human readable name of the Google Chat integration | 
**HideTransition** | Pointer to **bool** | When true, the transition info (e.g., \&quot;from ok to alert\&quot;) is hidden from the Google Chat message | [optional] 
**HideTime** | Pointer to **bool** | When true, the timestamp is hidden from the Google Chat message | [optional] 

## Methods

### NewRuleTargetConfigGoogleChat

`func NewRuleTargetConfigGoogleChat(type_ string, integrationId string, integrationName string, ) *RuleTargetConfigGoogleChat`

NewRuleTargetConfigGoogleChat instantiates a new RuleTargetConfigGoogleChat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTargetConfigGoogleChatWithDefaults

`func NewRuleTargetConfigGoogleChatWithDefaults() *RuleTargetConfigGoogleChat`

NewRuleTargetConfigGoogleChatWithDefaults instantiates a new RuleTargetConfigGoogleChat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuleTargetConfigGoogleChat) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleTargetConfigGoogleChat) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleTargetConfigGoogleChat) SetType(v string)`

SetType sets Type field to given value.


### GetIntegrationId

`func (o *RuleTargetConfigGoogleChat) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *RuleTargetConfigGoogleChat) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *RuleTargetConfigGoogleChat) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetIntegrationName

`func (o *RuleTargetConfigGoogleChat) GetIntegrationName() string`

GetIntegrationName returns the IntegrationName field if non-nil, zero value otherwise.

### GetIntegrationNameOk

`func (o *RuleTargetConfigGoogleChat) GetIntegrationNameOk() (*string, bool)`

GetIntegrationNameOk returns a tuple with the IntegrationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationName

`func (o *RuleTargetConfigGoogleChat) SetIntegrationName(v string)`

SetIntegrationName sets IntegrationName field to given value.


### GetHideTransition

`func (o *RuleTargetConfigGoogleChat) GetHideTransition() bool`

GetHideTransition returns the HideTransition field if non-nil, zero value otherwise.

### GetHideTransitionOk

`func (o *RuleTargetConfigGoogleChat) GetHideTransitionOk() (*bool, bool)`

GetHideTransitionOk returns a tuple with the HideTransition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideTransition

`func (o *RuleTargetConfigGoogleChat) SetHideTransition(v bool)`

SetHideTransition sets HideTransition field to given value.

### HasHideTransition

`func (o *RuleTargetConfigGoogleChat) HasHideTransition() bool`

HasHideTransition returns a boolean if a field has been set.

### GetHideTime

`func (o *RuleTargetConfigGoogleChat) GetHideTime() bool`

GetHideTime returns the HideTime field if non-nil, zero value otherwise.

### GetHideTimeOk

`func (o *RuleTargetConfigGoogleChat) GetHideTimeOk() (*bool, bool)`

GetHideTimeOk returns a tuple with the HideTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideTime

`func (o *RuleTargetConfigGoogleChat) SetHideTime(v bool)`

SetHideTime sets HideTime field to given value.

### HasHideTime

`func (o *RuleTargetConfigGoogleChat) HasHideTime() bool`

HasHideTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


