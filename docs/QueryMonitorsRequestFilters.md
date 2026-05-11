# QueryMonitorsRequestFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owners** | Pointer to **[]string** | Restrict results to monitors owned by any of these team IDs | [optional] 
**SearchQuery** | Pointer to **string** | Substring matched against monitor name and configuration | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Restrict results to monitors that have at least one of the listed key/value tags | [optional] 
**Priorities** | Pointer to **[]float32** | Restrict results to monitors with one of these priorities | [optional] 
**Types** | Pointer to **[]string** | Restrict results to monitors with these types | [optional] 
**Activity** | Pointer to **string** | Restrict results by activity status: active monitors or currently snoozed monitors | [optional] 
**ClusterIds** | Pointer to **[]string** | Restrict results to monitors deployed to any of these cluster IDs (monitors with no cluster restriction are always included) | [optional] 

## Methods

### NewQueryMonitorsRequestFilters

`func NewQueryMonitorsRequestFilters() *QueryMonitorsRequestFilters`

NewQueryMonitorsRequestFilters instantiates a new QueryMonitorsRequestFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMonitorsRequestFiltersWithDefaults

`func NewQueryMonitorsRequestFiltersWithDefaults() *QueryMonitorsRequestFilters`

NewQueryMonitorsRequestFiltersWithDefaults instantiates a new QueryMonitorsRequestFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwners

`func (o *QueryMonitorsRequestFilters) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *QueryMonitorsRequestFilters) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *QueryMonitorsRequestFilters) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *QueryMonitorsRequestFilters) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSearchQuery

`func (o *QueryMonitorsRequestFilters) GetSearchQuery() string`

GetSearchQuery returns the SearchQuery field if non-nil, zero value otherwise.

### GetSearchQueryOk

`func (o *QueryMonitorsRequestFilters) GetSearchQueryOk() (*string, bool)`

GetSearchQueryOk returns a tuple with the SearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchQuery

`func (o *QueryMonitorsRequestFilters) SetSearchQuery(v string)`

SetSearchQuery sets SearchQuery field to given value.

### HasSearchQuery

`func (o *QueryMonitorsRequestFilters) HasSearchQuery() bool`

HasSearchQuery returns a boolean if a field has been set.

### GetTags

`func (o *QueryMonitorsRequestFilters) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *QueryMonitorsRequestFilters) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *QueryMonitorsRequestFilters) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *QueryMonitorsRequestFilters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPriorities

`func (o *QueryMonitorsRequestFilters) GetPriorities() []float32`

GetPriorities returns the Priorities field if non-nil, zero value otherwise.

### GetPrioritiesOk

`func (o *QueryMonitorsRequestFilters) GetPrioritiesOk() (*[]float32, bool)`

GetPrioritiesOk returns a tuple with the Priorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorities

`func (o *QueryMonitorsRequestFilters) SetPriorities(v []float32)`

SetPriorities sets Priorities field to given value.

### HasPriorities

`func (o *QueryMonitorsRequestFilters) HasPriorities() bool`

HasPriorities returns a boolean if a field has been set.

### GetTypes

`func (o *QueryMonitorsRequestFilters) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *QueryMonitorsRequestFilters) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *QueryMonitorsRequestFilters) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *QueryMonitorsRequestFilters) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetActivity

`func (o *QueryMonitorsRequestFilters) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *QueryMonitorsRequestFilters) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *QueryMonitorsRequestFilters) SetActivity(v string)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *QueryMonitorsRequestFilters) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetClusterIds

`func (o *QueryMonitorsRequestFilters) GetClusterIds() []string`

GetClusterIds returns the ClusterIds field if non-nil, zero value otherwise.

### GetClusterIdsOk

`func (o *QueryMonitorsRequestFilters) GetClusterIdsOk() (*[]string, bool)`

GetClusterIdsOk returns a tuple with the ClusterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIds

`func (o *QueryMonitorsRequestFilters) SetClusterIds(v []string)`

SetClusterIds sets ClusterIds field to given value.

### HasClusterIds

`func (o *QueryMonitorsRequestFilters) HasClusterIds() bool`

HasClusterIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


