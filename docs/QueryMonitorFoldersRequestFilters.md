# QueryMonitorFoldersRequestFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owners** | Pointer to [**QueryMonitorFoldersRequestFiltersOwners**](QueryMonitorFoldersRequestFiltersOwners.md) |  | [optional] 
**SearchQuery** | Pointer to [**QueryMonitorFoldersRequestFiltersSearchQuery**](QueryMonitorFoldersRequestFiltersSearchQuery.md) |  | [optional] 
**Tags** | Pointer to [**QueryMonitorFoldersRequestFiltersTags**](QueryMonitorFoldersRequestFiltersTags.md) |  | [optional] 
**ParentFolderId** | Pointer to **NullableString** | Folder filter for dashboard query results. Optional; pass a folder ID to return dashboards in that folder, or null to return only dashboards that are not in any folder. | [optional] 

## Methods

### NewQueryMonitorFoldersRequestFilters

`func NewQueryMonitorFoldersRequestFilters() *QueryMonitorFoldersRequestFilters`

NewQueryMonitorFoldersRequestFilters instantiates a new QueryMonitorFoldersRequestFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMonitorFoldersRequestFiltersWithDefaults

`func NewQueryMonitorFoldersRequestFiltersWithDefaults() *QueryMonitorFoldersRequestFilters`

NewQueryMonitorFoldersRequestFiltersWithDefaults instantiates a new QueryMonitorFoldersRequestFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwners

`func (o *QueryMonitorFoldersRequestFilters) GetOwners() QueryMonitorFoldersRequestFiltersOwners`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *QueryMonitorFoldersRequestFilters) GetOwnersOk() (*QueryMonitorFoldersRequestFiltersOwners, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *QueryMonitorFoldersRequestFilters) SetOwners(v QueryMonitorFoldersRequestFiltersOwners)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *QueryMonitorFoldersRequestFilters) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSearchQuery

`func (o *QueryMonitorFoldersRequestFilters) GetSearchQuery() QueryMonitorFoldersRequestFiltersSearchQuery`

GetSearchQuery returns the SearchQuery field if non-nil, zero value otherwise.

### GetSearchQueryOk

`func (o *QueryMonitorFoldersRequestFilters) GetSearchQueryOk() (*QueryMonitorFoldersRequestFiltersSearchQuery, bool)`

GetSearchQueryOk returns a tuple with the SearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchQuery

`func (o *QueryMonitorFoldersRequestFilters) SetSearchQuery(v QueryMonitorFoldersRequestFiltersSearchQuery)`

SetSearchQuery sets SearchQuery field to given value.

### HasSearchQuery

`func (o *QueryMonitorFoldersRequestFilters) HasSearchQuery() bool`

HasSearchQuery returns a boolean if a field has been set.

### GetTags

`func (o *QueryMonitorFoldersRequestFilters) GetTags() QueryMonitorFoldersRequestFiltersTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *QueryMonitorFoldersRequestFilters) GetTagsOk() (*QueryMonitorFoldersRequestFiltersTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *QueryMonitorFoldersRequestFilters) SetTags(v QueryMonitorFoldersRequestFiltersTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *QueryMonitorFoldersRequestFilters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetParentFolderId

`func (o *QueryMonitorFoldersRequestFilters) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *QueryMonitorFoldersRequestFilters) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *QueryMonitorFoldersRequestFilters) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *QueryMonitorFoldersRequestFilters) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *QueryMonitorFoldersRequestFilters) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *QueryMonitorFoldersRequestFilters) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


