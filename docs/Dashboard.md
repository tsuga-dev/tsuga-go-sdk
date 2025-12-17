# Dashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the dashboard | 
**Name** | **string** | Display name of the dashboard | 
**Owner** | **string** | Team ID that owns and manages the dashboard | 
**Graphs** | [**[]Graph**](Graph.md) | Ordered widgets that compose the dashboard | 
**Filters** | Pointer to **[]string** | Filters applied to every widget on the dashboard | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List of key/value tags applied to the resource | [optional] 
**TimePreset** | Pointer to **string** |  | [optional] 

## Methods

### NewDashboard

`func NewDashboard(id string, name string, owner string, graphs []Graph, ) *Dashboard`

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

`func (o *Dashboard) GetGraphs() []Graph`

GetGraphs returns the Graphs field if non-nil, zero value otherwise.

### GetGraphsOk

`func (o *Dashboard) GetGraphsOk() (*[]Graph, bool)`

GetGraphsOk returns a tuple with the Graphs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphs

`func (o *Dashboard) SetGraphs(v []Graph)`

SetGraphs sets Graphs field to given value.


### GetFilters

`func (o *Dashboard) GetFilters() []string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Dashboard) GetFiltersOk() (*[]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Dashboard) SetFilters(v []string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *Dashboard) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetTags

`func (o *Dashboard) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Dashboard) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Dashboard) SetTags(v []Tag)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


