# CreateTagPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**IsActive** | **bool** |  | 
**TagKey** | **string** |  | 
**AllowedTagValues** | **[]string** |  | 
**IsRequired** | **bool** |  | 
**TeamScope** | Pointer to [**CreateTagPolicyRequestTeamScope**](CreateTagPolicyRequestTeamScope.md) |  | [optional] 
**Configuration** | [**CreateTagPolicyRequestConfiguration**](CreateTagPolicyRequestConfiguration.md) |  | 
**Owner** | **string** |  | 

## Methods

### NewCreateTagPolicyRequest

`func NewCreateTagPolicyRequest(name string, isActive bool, tagKey string, allowedTagValues []string, isRequired bool, configuration CreateTagPolicyRequestConfiguration, owner string, ) *CreateTagPolicyRequest`

NewCreateTagPolicyRequest instantiates a new CreateTagPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTagPolicyRequestWithDefaults

`func NewCreateTagPolicyRequestWithDefaults() *CreateTagPolicyRequest`

NewCreateTagPolicyRequestWithDefaults instantiates a new CreateTagPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTagPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTagPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTagPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateTagPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTagPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTagPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTagPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsActive

`func (o *CreateTagPolicyRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateTagPolicyRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateTagPolicyRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetTagKey

`func (o *CreateTagPolicyRequest) GetTagKey() string`

GetTagKey returns the TagKey field if non-nil, zero value otherwise.

### GetTagKeyOk

`func (o *CreateTagPolicyRequest) GetTagKeyOk() (*string, bool)`

GetTagKeyOk returns a tuple with the TagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKey

`func (o *CreateTagPolicyRequest) SetTagKey(v string)`

SetTagKey sets TagKey field to given value.


### GetAllowedTagValues

`func (o *CreateTagPolicyRequest) GetAllowedTagValues() []string`

GetAllowedTagValues returns the AllowedTagValues field if non-nil, zero value otherwise.

### GetAllowedTagValuesOk

`func (o *CreateTagPolicyRequest) GetAllowedTagValuesOk() (*[]string, bool)`

GetAllowedTagValuesOk returns a tuple with the AllowedTagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTagValues

`func (o *CreateTagPolicyRequest) SetAllowedTagValues(v []string)`

SetAllowedTagValues sets AllowedTagValues field to given value.


### GetIsRequired

`func (o *CreateTagPolicyRequest) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *CreateTagPolicyRequest) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *CreateTagPolicyRequest) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.


### GetTeamScope

`func (o *CreateTagPolicyRequest) GetTeamScope() CreateTagPolicyRequestTeamScope`

GetTeamScope returns the TeamScope field if non-nil, zero value otherwise.

### GetTeamScopeOk

`func (o *CreateTagPolicyRequest) GetTeamScopeOk() (*CreateTagPolicyRequestTeamScope, bool)`

GetTeamScopeOk returns a tuple with the TeamScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamScope

`func (o *CreateTagPolicyRequest) SetTeamScope(v CreateTagPolicyRequestTeamScope)`

SetTeamScope sets TeamScope field to given value.

### HasTeamScope

`func (o *CreateTagPolicyRequest) HasTeamScope() bool`

HasTeamScope returns a boolean if a field has been set.

### GetConfiguration

`func (o *CreateTagPolicyRequest) GetConfiguration() CreateTagPolicyRequestConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreateTagPolicyRequest) GetConfigurationOk() (*CreateTagPolicyRequestConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreateTagPolicyRequest) SetConfiguration(v CreateTagPolicyRequestConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetOwner

`func (o *CreateTagPolicyRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateTagPolicyRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateTagPolicyRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


