# CreateDashboardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the dashboard | 
**Owner** | **string** | Team ID that owns and manages the dashboard | 
**Graphs** | [**[]Graph**](Graph.md) | Ordered widgets that compose the dashboard | 
**Filters** | Pointer to [**[]CreateDashboardRequestFiltersInner**](CreateDashboardRequestFiltersInner.md) | Filters applied to every widget on the dashboard | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List of key/value tags applied to the resource | [optional] 
**TimePreset** | Pointer to **string** |  | [optional] 

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

`func (o *CreateDashboardRequest) GetFilters() []CreateDashboardRequestFiltersInner`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CreateDashboardRequest) GetFiltersOk() (*[]CreateDashboardRequestFiltersInner, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CreateDashboardRequest) SetFilters(v []CreateDashboardRequestFiltersInner)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


