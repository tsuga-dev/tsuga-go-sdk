# QueryMonitorsRequestFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owners** | Pointer to [**QueryMonitorsRequestFiltersOwners**](QueryMonitorsRequestFiltersOwners.md) |  | [optional] 
**SearchQuery** | Pointer to [**QueryMonitorsRequestFiltersSearchQuery**](QueryMonitorsRequestFiltersSearchQuery.md) |  | [optional] 
**Tags** | Pointer to [**QueryMonitorsRequestFiltersTags**](QueryMonitorsRequestFiltersTags.md) |  | [optional] 
**Priorities** | Pointer to [**QueryMonitorsRequestFiltersPriorities**](QueryMonitorsRequestFiltersPriorities.md) |  | [optional] 
**Types** | Pointer to [**QueryMonitorsRequestFiltersTypes**](QueryMonitorsRequestFiltersTypes.md) |  | [optional] 
**Activity** | Pointer to **string** | Restrict results by activity status: active monitors or currently snoozed monitors | [optional] 
**ClusterIds** | Pointer to [**QueryMonitorsRequestFiltersClusterIds**](QueryMonitorsRequestFiltersClusterIds.md) |  | [optional] 

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

`func (o *QueryMonitorsRequestFilters) GetOwners() QueryMonitorsRequestFiltersOwners`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *QueryMonitorsRequestFilters) GetOwnersOk() (*QueryMonitorsRequestFiltersOwners, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *QueryMonitorsRequestFilters) SetOwners(v QueryMonitorsRequestFiltersOwners)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *QueryMonitorsRequestFilters) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSearchQuery

`func (o *QueryMonitorsRequestFilters) GetSearchQuery() QueryMonitorsRequestFiltersSearchQuery`

GetSearchQuery returns the SearchQuery field if non-nil, zero value otherwise.

### GetSearchQueryOk

`func (o *QueryMonitorsRequestFilters) GetSearchQueryOk() (*QueryMonitorsRequestFiltersSearchQuery, bool)`

GetSearchQueryOk returns a tuple with the SearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchQuery

`func (o *QueryMonitorsRequestFilters) SetSearchQuery(v QueryMonitorsRequestFiltersSearchQuery)`

SetSearchQuery sets SearchQuery field to given value.

### HasSearchQuery

`func (o *QueryMonitorsRequestFilters) HasSearchQuery() bool`

HasSearchQuery returns a boolean if a field has been set.

### GetTags

`func (o *QueryMonitorsRequestFilters) GetTags() QueryMonitorsRequestFiltersTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *QueryMonitorsRequestFilters) GetTagsOk() (*QueryMonitorsRequestFiltersTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *QueryMonitorsRequestFilters) SetTags(v QueryMonitorsRequestFiltersTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *QueryMonitorsRequestFilters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPriorities

`func (o *QueryMonitorsRequestFilters) GetPriorities() QueryMonitorsRequestFiltersPriorities`

GetPriorities returns the Priorities field if non-nil, zero value otherwise.

### GetPrioritiesOk

`func (o *QueryMonitorsRequestFilters) GetPrioritiesOk() (*QueryMonitorsRequestFiltersPriorities, bool)`

GetPrioritiesOk returns a tuple with the Priorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorities

`func (o *QueryMonitorsRequestFilters) SetPriorities(v QueryMonitorsRequestFiltersPriorities)`

SetPriorities sets Priorities field to given value.

### HasPriorities

`func (o *QueryMonitorsRequestFilters) HasPriorities() bool`

HasPriorities returns a boolean if a field has been set.

### GetTypes

`func (o *QueryMonitorsRequestFilters) GetTypes() QueryMonitorsRequestFiltersTypes`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *QueryMonitorsRequestFilters) GetTypesOk() (*QueryMonitorsRequestFiltersTypes, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *QueryMonitorsRequestFilters) SetTypes(v QueryMonitorsRequestFiltersTypes)`

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

`func (o *QueryMonitorsRequestFilters) GetClusterIds() QueryMonitorsRequestFiltersClusterIds`

GetClusterIds returns the ClusterIds field if non-nil, zero value otherwise.

### GetClusterIdsOk

`func (o *QueryMonitorsRequestFilters) GetClusterIdsOk() (*QueryMonitorsRequestFiltersClusterIds, bool)`

GetClusterIdsOk returns a tuple with the ClusterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIds

`func (o *QueryMonitorsRequestFilters) SetClusterIds(v QueryMonitorsRequestFiltersClusterIds)`

SetClusterIds sets ClusterIds field to given value.

### HasClusterIds

`func (o *QueryMonitorsRequestFilters) HasClusterIds() bool`

HasClusterIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


