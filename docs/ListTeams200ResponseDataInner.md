# ListTeams200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the team | 
**Name** | **string** | Human readable team name displayed throughout the app | 
**Description** | Pointer to **string** |  | [optional] 
**Visibility** | **string** | Controls whether the resources of the team are discoverable by users | 
**Tags** | Pointer to [**[]ListDashboards200ResponseDataInnerTagsInner**](ListDashboards200ResponseDataInnerTagsInner.md) | List of key/value tags applied to the resource | [optional] 

## Methods

### NewListTeams200ResponseDataInner

`func NewListTeams200ResponseDataInner(id string, name string, visibility string, ) *ListTeams200ResponseDataInner`

NewListTeams200ResponseDataInner instantiates a new ListTeams200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTeams200ResponseDataInnerWithDefaults

`func NewListTeams200ResponseDataInnerWithDefaults() *ListTeams200ResponseDataInner`

NewListTeams200ResponseDataInnerWithDefaults instantiates a new ListTeams200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListTeams200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListTeams200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListTeams200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListTeams200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListTeams200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListTeams200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ListTeams200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListTeams200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListTeams200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListTeams200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *ListTeams200ResponseDataInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListTeams200ResponseDataInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListTeams200ResponseDataInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetTags

`func (o *ListTeams200ResponseDataInner) GetTags() []ListDashboards200ResponseDataInnerTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListTeams200ResponseDataInner) GetTagsOk() (*[]ListDashboards200ResponseDataInnerTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListTeams200ResponseDataInner) SetTags(v []ListDashboards200ResponseDataInnerTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListTeams200ResponseDataInner) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


