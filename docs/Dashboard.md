# Dashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Tsuga-generated dashboard ID assigned when the dashboard is created. | 
**Name** | **string** | Display name of the dashboard. Set by the caller on create or update. | 
**Owner** | **string** | ID of the team that owns the dashboard. Set by the caller on create or update and used for dashboard access control. | 
**Graphs** | [**[]Graph1**](Graph1.md) | Ordered widgets that compose the dashboard | 
**Filters** | Pointer to [**[]DashboardFiltersInner**](DashboardFiltersInner.md) | Dashboard-wide filters applied to every widget on the dashboard. | [optional] 
**Tags** | Pointer to [**[]Tag1**](Tag1.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | [optional] 
**TimePreset** | Pointer to **string** | Relative time preset used when opening the dashboard. | [optional] 
**FolderId** | Pointer to **string** | Folder that contains the dashboard. Set from the dashboard’s saved folder relationship. Returned when the dashboard is in a folder; omitted when it is not. | [optional] 

## Methods

### NewDashboard

`func NewDashboard(id string, name string, owner string, graphs []Graph1, ) *Dashboard`

NewDashboard instantiates a new Dashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardWithDefaults

`func NewDashboardWithDefaults() *Dashboard`

NewDashboardWithDefaults instantiates a new Dashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Dashboard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dashboard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Dashboard) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Dashboard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dashboard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dashboard) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *Dashboard) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Dashboard) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Dashboard) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetGraphs

`func (o *Dashboard) GetGraphs() []Graph1`

GetGraphs returns the Graphs field if non-nil, zero value otherwise.

### GetGraphsOk

`func (o *Dashboard) GetGraphsOk() (*[]Graph1, bool)`

GetGraphsOk returns a tuple with the Graphs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphs

`func (o *Dashboard) SetGraphs(v []Graph1)`

SetGraphs sets Graphs field to given value.


### GetFilters

`func (o *Dashboard) GetFilters() []DashboardFiltersInner`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Dashboard) GetFiltersOk() (*[]DashboardFiltersInner, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Dashboard) SetFilters(v []DashboardFiltersInner)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *Dashboard) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetTags

`func (o *Dashboard) GetTags() []Tag1`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Dashboard) GetTagsOk() (*[]Tag1, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Dashboard) SetTags(v []Tag1)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Dashboard) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimePreset

`func (o *Dashboard) GetTimePreset() string`

GetTimePreset returns the TimePreset field if non-nil, zero value otherwise.

### GetTimePresetOk

`func (o *Dashboard) GetTimePresetOk() (*string, bool)`

GetTimePresetOk returns a tuple with the TimePreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePreset

`func (o *Dashboard) SetTimePreset(v string)`

SetTimePreset sets TimePreset field to given value.

### HasTimePreset

`func (o *Dashboard) HasTimePreset() bool`

HasTimePreset returns a boolean if a field has been set.

### GetFolderId

`func (o *Dashboard) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *Dashboard) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *Dashboard) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *Dashboard) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


