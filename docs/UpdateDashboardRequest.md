# UpdateDashboardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Display name of the dashboard | [optional] 
**Owner** | Pointer to **string** | Team ID that owns and manages the dashboard | [optional] 
**Graphs** | Pointer to [**[]ListDashboards200ResponseDataInnerGraphsInner**](ListDashboards200ResponseDataInnerGraphsInner.md) | Ordered widgets that compose the dashboard | [optional] 
**Filters** | Pointer to **[]string** | Filters applied to every widget on the dashboard | [optional] 
**Tags** | Pointer to [**[]ListDashboards200ResponseDataInnerTagsInner**](ListDashboards200ResponseDataInnerTagsInner.md) | List of key/value tags applied to the resource | [optional] 

## Methods

### NewUpdateDashboardRequest

`func NewUpdateDashboardRequest() *UpdateDashboardRequest`

NewUpdateDashboardRequest instantiates a new UpdateDashboardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDashboardRequestWithDefaults

`func NewUpdateDashboardRequestWithDefaults() *UpdateDashboardRequest`

NewUpdateDashboardRequestWithDefaults instantiates a new UpdateDashboardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateDashboardRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDashboardRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDashboardRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDashboardRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *UpdateDashboardRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateDashboardRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateDashboardRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdateDashboardRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetGraphs

`func (o *UpdateDashboardRequest) GetGraphs() []ListDashboards200ResponseDataInnerGraphsInner`

GetGraphs returns the Graphs field if non-nil, zero value otherwise.

### GetGraphsOk

`func (o *UpdateDashboardRequest) GetGraphsOk() (*[]ListDashboards200ResponseDataInnerGraphsInner, bool)`

GetGraphsOk returns a tuple with the Graphs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphs

`func (o *UpdateDashboardRequest) SetGraphs(v []ListDashboards200ResponseDataInnerGraphsInner)`

SetGraphs sets Graphs field to given value.

### HasGraphs

`func (o *UpdateDashboardRequest) HasGraphs() bool`

HasGraphs returns a boolean if a field has been set.

### GetFilters

`func (o *UpdateDashboardRequest) GetFilters() []string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *UpdateDashboardRequest) GetFiltersOk() (*[]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *UpdateDashboardRequest) SetFilters(v []string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *UpdateDashboardRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetTags

`func (o *UpdateDashboardRequest) GetTags() []ListDashboards200ResponseDataInnerTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateDashboardRequest) GetTagsOk() (*[]ListDashboards200ResponseDataInnerTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateDashboardRequest) SetTags(v []ListDashboards200ResponseDataInnerTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateDashboardRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


