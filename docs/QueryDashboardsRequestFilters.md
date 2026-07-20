# QueryDashboardsRequestFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owners** | Pointer to [**QueryDashboardsRequestFiltersOwners**](QueryDashboardsRequestFiltersOwners.md) |  | [optional] 
**SearchQuery** | Pointer to [**QueryDashboardsRequestFiltersSearchQuery**](QueryDashboardsRequestFiltersSearchQuery.md) |  | [optional] 
**Tags** | Pointer to [**QueryDashboardsRequestFiltersTags**](QueryDashboardsRequestFiltersTags.md) |  | [optional] 
**FolderId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewQueryDashboardsRequestFilters

`func NewQueryDashboardsRequestFilters() *QueryDashboardsRequestFilters`

NewQueryDashboardsRequestFilters instantiates a new QueryDashboardsRequestFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryDashboardsRequestFiltersWithDefaults

`func NewQueryDashboardsRequestFiltersWithDefaults() *QueryDashboardsRequestFilters`

NewQueryDashboardsRequestFiltersWithDefaults instantiates a new QueryDashboardsRequestFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwners

`func (o *QueryDashboardsRequestFilters) GetOwners() QueryDashboardsRequestFiltersOwners`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *QueryDashboardsRequestFilters) GetOwnersOk() (*QueryDashboardsRequestFiltersOwners, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *QueryDashboardsRequestFilters) SetOwners(v QueryDashboardsRequestFiltersOwners)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *QueryDashboardsRequestFilters) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSearchQuery

`func (o *QueryDashboardsRequestFilters) GetSearchQuery() QueryDashboardsRequestFiltersSearchQuery`

GetSearchQuery returns the SearchQuery field if non-nil, zero value otherwise.

### GetSearchQueryOk

`func (o *QueryDashboardsRequestFilters) GetSearchQueryOk() (*QueryDashboardsRequestFiltersSearchQuery, bool)`

GetSearchQueryOk returns a tuple with the SearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchQuery

`func (o *QueryDashboardsRequestFilters) SetSearchQuery(v QueryDashboardsRequestFiltersSearchQuery)`

SetSearchQuery sets SearchQuery field to given value.

### HasSearchQuery

`func (o *QueryDashboardsRequestFilters) HasSearchQuery() bool`

HasSearchQuery returns a boolean if a field has been set.

### GetTags

`func (o *QueryDashboardsRequestFilters) GetTags() QueryDashboardsRequestFiltersTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *QueryDashboardsRequestFilters) GetTagsOk() (*QueryDashboardsRequestFiltersTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *QueryDashboardsRequestFilters) SetTags(v QueryDashboardsRequestFiltersTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *QueryDashboardsRequestFilters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFolderId

`func (o *QueryDashboardsRequestFilters) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *QueryDashboardsRequestFilters) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *QueryDashboardsRequestFilters) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *QueryDashboardsRequestFilters) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *QueryDashboardsRequestFilters) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *QueryDashboardsRequestFilters) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


