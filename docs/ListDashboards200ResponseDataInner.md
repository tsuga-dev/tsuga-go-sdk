# ListDashboards200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the dashboard | 
**Name** | **string** | Display name of the dashboard | 
**Owner** | **string** | Team ID that owns and manages the dashboard | 
**Graphs** | [**[]ListDashboards200ResponseDataInnerGraphsInner**](ListDashboards200ResponseDataInnerGraphsInner.md) | Ordered widgets that compose the dashboard | 
**Filters** | Pointer to **[]string** | Filters applied to every widget on the dashboard | [optional] 
**Tags** | Pointer to [**[]ListDashboards200ResponseDataInnerTagsInner**](ListDashboards200ResponseDataInnerTagsInner.md) | List of key/value tags applied to the resource | [optional] 

## Methods

### NewListDashboards200ResponseDataInner

`func NewListDashboards200ResponseDataInner(id string, name string, owner string, graphs []ListDashboards200ResponseDataInnerGraphsInner, ) *ListDashboards200ResponseDataInner`

NewListDashboards200ResponseDataInner instantiates a new ListDashboards200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDashboards200ResponseDataInnerWithDefaults

`func NewListDashboards200ResponseDataInnerWithDefaults() *ListDashboards200ResponseDataInner`

NewListDashboards200ResponseDataInnerWithDefaults instantiates a new ListDashboards200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListDashboards200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListDashboards200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListDashboards200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListDashboards200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListDashboards200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListDashboards200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *ListDashboards200ResponseDataInner) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListDashboards200ResponseDataInner) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListDashboards200ResponseDataInner) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetGraphs

`func (o *ListDashboards200ResponseDataInner) GetGraphs() []ListDashboards200ResponseDataInnerGraphsInner`

GetGraphs returns the Graphs field if non-nil, zero value otherwise.

### GetGraphsOk

`func (o *ListDashboards200ResponseDataInner) GetGraphsOk() (*[]ListDashboards200ResponseDataInnerGraphsInner, bool)`

GetGraphsOk returns a tuple with the Graphs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphs

`func (o *ListDashboards200ResponseDataInner) SetGraphs(v []ListDashboards200ResponseDataInnerGraphsInner)`

SetGraphs sets Graphs field to given value.


### GetFilters

`func (o *ListDashboards200ResponseDataInner) GetFilters() []string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ListDashboards200ResponseDataInner) GetFiltersOk() (*[]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ListDashboards200ResponseDataInner) SetFilters(v []string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ListDashboards200ResponseDataInner) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetTags

`func (o *ListDashboards200ResponseDataInner) GetTags() []ListDashboards200ResponseDataInnerTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListDashboards200ResponseDataInner) GetTagsOk() (*[]ListDashboards200ResponseDataInnerTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListDashboards200ResponseDataInner) SetTags(v []ListDashboards200ResponseDataInnerTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListDashboards200ResponseDataInner) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


