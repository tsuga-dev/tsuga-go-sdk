# CreateDashboardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the dashboard. Set by the caller on create or update. | 
**Owner** | **string** | ID of the team that owns the dashboard. Set by the caller on create or update. Required and used for dashboard access control. | 
**Graphs** | [**[]Graph**](Graph.md) | Ordered widgets that compose the dashboard | 
**Filters** | Pointer to [**[]UpdateDashboardRequestFiltersInner**](UpdateDashboardRequestFiltersInner.md) | Dashboard-wide filters applied to every widget on the dashboard. Up to 10 filters are allowed. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key/value tags to apply to the resource. Up to 50 tags are accepted and tag policies may require specific keys or values. | [optional] 
**TimePreset** | Pointer to **string** | Relative time preset used when opening the dashboard. Set by the caller or by Tsuga’s default. Optional; create defaults to &#x60;past-30-minutes&#x60; when omitted, and update omission preserves the current value. | [optional] 
**FolderId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateDashboardRequest

`func NewCreateDashboardRequest(name string, owner string, graphs []Graph, ) *CreateDashboardRequest`

NewCreateDashboardRequest instantiates a new CreateDashboardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDashboardRequestWithDefaults

`func NewCreateDashboardRequestWithDefaults() *CreateDashboardRequest`

NewCreateDashboardRequestWithDefaults instantiates a new CreateDashboardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateDashboardRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDashboardRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDashboardRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *CreateDashboardRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateDashboardRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateDashboardRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetGraphs

`func (o *CreateDashboardRequest) GetGraphs() []Graph`

GetGraphs returns the Graphs field if non-nil, zero value otherwise.

### GetGraphsOk

`func (o *CreateDashboardRequest) GetGraphsOk() (*[]Graph, bool)`

GetGraphsOk returns a tuple with the Graphs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphs

`func (o *CreateDashboardRequest) SetGraphs(v []Graph)`

SetGraphs sets Graphs field to given value.


### GetFilters

`func (o *CreateDashboardRequest) GetFilters() []UpdateDashboardRequestFiltersInner`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CreateDashboardRequest) GetFiltersOk() (*[]UpdateDashboardRequestFiltersInner, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CreateDashboardRequest) SetFilters(v []UpdateDashboardRequestFiltersInner)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *CreateDashboardRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetTags

`func (o *CreateDashboardRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateDashboardRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateDashboardRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateDashboardRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimePreset

`func (o *CreateDashboardRequest) GetTimePreset() string`

GetTimePreset returns the TimePreset field if non-nil, zero value otherwise.

### GetTimePresetOk

`func (o *CreateDashboardRequest) GetTimePresetOk() (*string, bool)`

GetTimePresetOk returns a tuple with the TimePreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePreset

`func (o *CreateDashboardRequest) SetTimePreset(v string)`

SetTimePreset sets TimePreset field to given value.

### HasTimePreset

`func (o *CreateDashboardRequest) HasTimePreset() bool`

HasTimePreset returns a boolean if a field has been set.

### GetFolderId

`func (o *CreateDashboardRequest) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *CreateDashboardRequest) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *CreateDashboardRequest) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *CreateDashboardRequest) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *CreateDashboardRequest) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *CreateDashboardRequest) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


