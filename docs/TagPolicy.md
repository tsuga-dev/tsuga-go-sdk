# TagPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**IsActive** | **bool** |  | 
**TagKey** | **string** |  | 
**AllowedTagValues** | **[]string** |  | 
**IsRequired** | **bool** |  | 
**TeamScope** | Pointer to [**CreateTagPolicyRequestTeamScope**](CreateTagPolicyRequestTeamScope.md) |  | [optional] 
**Configuration** | [**CreateTagPolicyRequestConfiguration**](CreateTagPolicyRequestConfiguration.md) |  | 
**Owner** | Pointer to **string** |  | [optional] 

## Methods

### NewTagPolicy

`func NewTagPolicy(id string, name string, isActive bool, tagKey string, allowedTagValues []string, isRequired bool, configuration CreateTagPolicyRequestConfiguration, ) *TagPolicy`

NewTagPolicy instantiates a new TagPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagPolicyWithDefaults

`func NewTagPolicyWithDefaults() *TagPolicy`

NewTagPolicyWithDefaults instantiates a new TagPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TagPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagPolicy) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TagPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TagPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TagPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TagPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TagPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsActive

`func (o *TagPolicy) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *TagPolicy) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *TagPolicy) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetTagKey

`func (o *TagPolicy) GetTagKey() string`

GetTagKey returns the TagKey field if non-nil, zero value otherwise.

### GetTagKeyOk

`func (o *TagPolicy) GetTagKeyOk() (*string, bool)`

GetTagKeyOk returns a tuple with the TagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKey

`func (o *TagPolicy) SetTagKey(v string)`

SetTagKey sets TagKey field to given value.


### GetAllowedTagValues

`func (o *TagPolicy) GetAllowedTagValues() []string`

GetAllowedTagValues returns the AllowedTagValues field if non-nil, zero value otherwise.

### GetAllowedTagValuesOk

`func (o *TagPolicy) GetAllowedTagValuesOk() (*[]string, bool)`

GetAllowedTagValuesOk returns a tuple with the AllowedTagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTagValues

`func (o *TagPolicy) SetAllowedTagValues(v []string)`

SetAllowedTagValues sets AllowedTagValues field to given value.


### GetIsRequired

`func (o *TagPolicy) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *TagPolicy) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *TagPolicy) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.


### GetTeamScope

`func (o *TagPolicy) GetTeamScope() CreateTagPolicyRequestTeamScope`

GetTeamScope returns the TeamScope field if non-nil, zero value otherwise.

### GetTeamScopeOk

`func (o *TagPolicy) GetTeamScopeOk() (*CreateTagPolicyRequestTeamScope, bool)`

GetTeamScopeOk returns a tuple with the TeamScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamScope

`func (o *TagPolicy) SetTeamScope(v CreateTagPolicyRequestTeamScope)`

SetTeamScope sets TeamScope field to given value.

### HasTeamScope

`func (o *TagPolicy) HasTeamScope() bool`

HasTeamScope returns a boolean if a field has been set.

### GetConfiguration

`func (o *TagPolicy) GetConfiguration() CreateTagPolicyRequestConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *TagPolicy) GetConfigurationOk() (*CreateTagPolicyRequestConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *TagPolicy) SetConfiguration(v CreateTagPolicyRequestConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetOwner

`func (o *TagPolicy) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *TagPolicy) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *TagPolicy) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *TagPolicy) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


