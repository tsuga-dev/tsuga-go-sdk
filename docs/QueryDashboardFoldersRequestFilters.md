# QueryDashboardFoldersRequestFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owners** | Pointer to [**QueryDashboardFoldersRequestFiltersOwners**](QueryDashboardFoldersRequestFiltersOwners.md) |  | [optional] 
**SearchQuery** | Pointer to [**QueryDashboardFoldersRequestFiltersSearchQuery**](QueryDashboardFoldersRequestFiltersSearchQuery.md) |  | [optional] 
**Tags** | Pointer to [**QueryDashboardFoldersRequestFiltersTags**](QueryDashboardFoldersRequestFiltersTags.md) |  | [optional] 
**ParentFolderId** | Pointer to **NullableString** | Folder filter for dashboard query results. Optional; pass a folder ID to return dashboards in that folder, or null to return only dashboards that are not in any folder. | [optional] 

## Methods

### NewQueryDashboardFoldersRequestFilters

`func NewQueryDashboardFoldersRequestFilters() *QueryDashboardFoldersRequestFilters`

NewQueryDashboardFoldersRequestFilters instantiates a new QueryDashboardFoldersRequestFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryDashboardFoldersRequestFiltersWithDefaults

`func NewQueryDashboardFoldersRequestFiltersWithDefaults() *QueryDashboardFoldersRequestFilters`

NewQueryDashboardFoldersRequestFiltersWithDefaults instantiates a new QueryDashboardFoldersRequestFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwners

`func (o *QueryDashboardFoldersRequestFilters) GetOwners() QueryDashboardFoldersRequestFiltersOwners`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *QueryDashboardFoldersRequestFilters) GetOwnersOk() (*QueryDashboardFoldersRequestFiltersOwners, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *QueryDashboardFoldersRequestFilters) SetOwners(v QueryDashboardFoldersRequestFiltersOwners)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *QueryDashboardFoldersRequestFilters) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSearchQuery

`func (o *QueryDashboardFoldersRequestFilters) GetSearchQuery() QueryDashboardFoldersRequestFiltersSearchQuery`

GetSearchQuery returns the SearchQuery field if non-nil, zero value otherwise.

### GetSearchQueryOk

`func (o *QueryDashboardFoldersRequestFilters) GetSearchQueryOk() (*QueryDashboardFoldersRequestFiltersSearchQuery, bool)`

GetSearchQueryOk returns a tuple with the SearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchQuery

`func (o *QueryDashboardFoldersRequestFilters) SetSearchQuery(v QueryDashboardFoldersRequestFiltersSearchQuery)`

SetSearchQuery sets SearchQuery field to given value.

### HasSearchQuery

`func (o *QueryDashboardFoldersRequestFilters) HasSearchQuery() bool`

HasSearchQuery returns a boolean if a field has been set.

### GetTags

`func (o *QueryDashboardFoldersRequestFilters) GetTags() QueryDashboardFoldersRequestFiltersTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *QueryDashboardFoldersRequestFilters) GetTagsOk() (*QueryDashboardFoldersRequestFiltersTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *QueryDashboardFoldersRequestFilters) SetTags(v QueryDashboardFoldersRequestFiltersTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *QueryDashboardFoldersRequestFilters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetParentFolderId

`func (o *QueryDashboardFoldersRequestFilters) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *QueryDashboardFoldersRequestFilters) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *QueryDashboardFoldersRequestFilters) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *QueryDashboardFoldersRequestFilters) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *QueryDashboardFoldersRequestFilters) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *QueryDashboardFoldersRequestFilters) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


