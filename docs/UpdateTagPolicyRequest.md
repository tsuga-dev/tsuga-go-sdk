# UpdateTagPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable tag policy name. | 
**Description** | Pointer to **string** | Optional policy description. | [optional] 
**IsActive** | **bool** | Set to true for Tsuga to evaluate this policy. Reserved policies (for example the built-in &#x60;env&#x60; policy on ingestion API keys) reject a false value. | 
**TagKey** | **string** | Tag key enforced by this policy. Tsuga trims surrounding whitespace before storing the policy. | 
**AllowedTagValues** | **[]string** | Allowed values for &#x60;tagKey&#x60;. Leave empty to allow any value when the tag exists. Tsuga trims each submitted value before storing the policy. | 
**IsRequired** | **bool** | Set to true to require the tag. If false, allowed values still apply when the tag exists. Reserved policies reject a false value. | 
**TeamScope** | Pointer to [**CreateTagPolicyRequestTeamScope**](CreateTagPolicyRequestTeamScope.md) |  | [optional] 
**Configuration** | [**CreateTagPolicyRequestConfiguration**](CreateTagPolicyRequestConfiguration.md) |  | 
**Owner** | **string** | Team ID that will own and manage the policy. | 

## Methods

### NewUpdateTagPolicyRequest

`func NewUpdateTagPolicyRequest(name string, isActive bool, tagKey string, allowedTagValues []string, isRequired bool, configuration CreateTagPolicyRequestConfiguration, owner string, ) *UpdateTagPolicyRequest`

NewUpdateTagPolicyRequest instantiates a new UpdateTagPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTagPolicyRequestWithDefaults

`func NewUpdateTagPolicyRequestWithDefaults() *UpdateTagPolicyRequest`

NewUpdateTagPolicyRequestWithDefaults instantiates a new UpdateTagPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTagPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTagPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTagPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateTagPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateTagPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateTagPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateTagPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsActive

`func (o *UpdateTagPolicyRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateTagPolicyRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateTagPolicyRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetTagKey

`func (o *UpdateTagPolicyRequest) GetTagKey() string`

GetTagKey returns the TagKey field if non-nil, zero value otherwise.

### GetTagKeyOk

`func (o *UpdateTagPolicyRequest) GetTagKeyOk() (*string, bool)`

GetTagKeyOk returns a tuple with the TagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKey

`func (o *UpdateTagPolicyRequest) SetTagKey(v string)`

SetTagKey sets TagKey field to given value.


### GetAllowedTagValues

`func (o *UpdateTagPolicyRequest) GetAllowedTagValues() []string`

GetAllowedTagValues returns the AllowedTagValues field if non-nil, zero value otherwise.

### GetAllowedTagValuesOk

`func (o *UpdateTagPolicyRequest) GetAllowedTagValuesOk() (*[]string, bool)`

GetAllowedTagValuesOk returns a tuple with the AllowedTagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTagValues

`func (o *UpdateTagPolicyRequest) SetAllowedTagValues(v []string)`

SetAllowedTagValues sets AllowedTagValues field to given value.


### GetIsRequired

`func (o *UpdateTagPolicyRequest) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *UpdateTagPolicyRequest) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *UpdateTagPolicyRequest) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.


### GetTeamScope

`func (o *UpdateTagPolicyRequest) GetTeamScope() CreateTagPolicyRequestTeamScope`

GetTeamScope returns the TeamScope field if non-nil, zero value otherwise.

### GetTeamScopeOk

`func (o *UpdateTagPolicyRequest) GetTeamScopeOk() (*CreateTagPolicyRequestTeamScope, bool)`

GetTeamScopeOk returns a tuple with the TeamScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamScope

`func (o *UpdateTagPolicyRequest) SetTeamScope(v CreateTagPolicyRequestTeamScope)`

SetTeamScope sets TeamScope field to given value.

### HasTeamScope

`func (o *UpdateTagPolicyRequest) HasTeamScope() bool`

HasTeamScope returns a boolean if a field has been set.

### GetConfiguration

`func (o *UpdateTagPolicyRequest) GetConfiguration() CreateTagPolicyRequestConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *UpdateTagPolicyRequest) GetConfigurationOk() (*CreateTagPolicyRequestConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *UpdateTagPolicyRequest) SetConfiguration(v CreateTagPolicyRequestConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetOwner

`func (o *UpdateTagPolicyRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateTagPolicyRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateTagPolicyRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


