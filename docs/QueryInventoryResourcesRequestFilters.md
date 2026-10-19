# QueryInventoryResourcesRequestFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudPlatforms** | Pointer to [**QueryInventoryResourcesRequestFiltersCloudPlatforms**](QueryInventoryResourcesRequestFiltersCloudPlatforms.md) |  | [optional] 
**CloudAccounts** | Pointer to [**QueryInventoryResourcesRequestFiltersCloudAccounts**](QueryInventoryResourcesRequestFiltersCloudAccounts.md) |  | [optional] 
**CloudRegions** | Pointer to [**QueryInventoryResourcesRequestFiltersCloudRegions**](QueryInventoryResourcesRequestFiltersCloudRegions.md) |  | [optional] 
**NativeResourceTypes** | Pointer to [**QueryInventoryResourcesRequestFiltersNativeResourceTypes**](QueryInventoryResourcesRequestFiltersNativeResourceTypes.md) |  | [optional] 
**ResourceTypes** | Pointer to [**QueryInventoryResourcesRequestFiltersResourceTypes**](QueryInventoryResourcesRequestFiltersResourceTypes.md) |  | [optional] 
**Tags** | Pointer to [**QueryInventoryResourcesRequestFiltersTags**](QueryInventoryResourcesRequestFiltersTags.md) |  | [optional] 
**IsIacManaged** | Pointer to **bool** | Restrict results by infrastructure-as-code coverage. When true, only resources covered by the configured IaC identifiers are returned; when false, only uncovered resources are returned. | [optional] 
**SearchQuery** | Pointer to [**QueryInventoryResourcesRequestFiltersSearchQuery**](QueryInventoryResourcesRequestFiltersSearchQuery.md) |  | [optional] 

## Methods

### NewQueryInventoryResourcesRequestFilters

`func NewQueryInventoryResourcesRequestFilters() *QueryInventoryResourcesRequestFilters`

NewQueryInventoryResourcesRequestFilters instantiates a new QueryInventoryResourcesRequestFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryInventoryResourcesRequestFiltersWithDefaults

`func NewQueryInventoryResourcesRequestFiltersWithDefaults() *QueryInventoryResourcesRequestFilters`

NewQueryInventoryResourcesRequestFiltersWithDefaults instantiates a new QueryInventoryResourcesRequestFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudPlatforms

`func (o *QueryInventoryResourcesRequestFilters) GetCloudPlatforms() QueryInventoryResourcesRequestFiltersCloudPlatforms`

GetCloudPlatforms returns the CloudPlatforms field if non-nil, zero value otherwise.

### GetCloudPlatformsOk

`func (o *QueryInventoryResourcesRequestFilters) GetCloudPlatformsOk() (*QueryInventoryResourcesRequestFiltersCloudPlatforms, bool)`

GetCloudPlatformsOk returns a tuple with the CloudPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPlatforms

`func (o *QueryInventoryResourcesRequestFilters) SetCloudPlatforms(v QueryInventoryResourcesRequestFiltersCloudPlatforms)`

SetCloudPlatforms sets CloudPlatforms field to given value.

### HasCloudPlatforms

`func (o *QueryInventoryResourcesRequestFilters) HasCloudPlatforms() bool`

HasCloudPlatforms returns a boolean if a field has been set.

### GetCloudAccounts

`func (o *QueryInventoryResourcesRequestFilters) GetCloudAccounts() QueryInventoryResourcesRequestFiltersCloudAccounts`

GetCloudAccounts returns the CloudAccounts field if non-nil, zero value otherwise.

### GetCloudAccountsOk

`func (o *QueryInventoryResourcesRequestFilters) GetCloudAccountsOk() (*QueryInventoryResourcesRequestFiltersCloudAccounts, bool)`

GetCloudAccountsOk returns a tuple with the CloudAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccounts

`func (o *QueryInventoryResourcesRequestFilters) SetCloudAccounts(v QueryInventoryResourcesRequestFiltersCloudAccounts)`

SetCloudAccounts sets CloudAccounts field to given value.

### HasCloudAccounts

`func (o *QueryInventoryResourcesRequestFilters) HasCloudAccounts() bool`

HasCloudAccounts returns a boolean if a field has been set.

### GetCloudRegions

`func (o *QueryInventoryResourcesRequestFilters) GetCloudRegions() QueryInventoryResourcesRequestFiltersCloudRegions`

GetCloudRegions returns the CloudRegions field if non-nil, zero value otherwise.

### GetCloudRegionsOk

`func (o *QueryInventoryResourcesRequestFilters) GetCloudRegionsOk() (*QueryInventoryResourcesRequestFiltersCloudRegions, bool)`

GetCloudRegionsOk returns a tuple with the CloudRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudRegions

`func (o *QueryInventoryResourcesRequestFilters) SetCloudRegions(v QueryInventoryResourcesRequestFiltersCloudRegions)`

SetCloudRegions sets CloudRegions field to given value.

### HasCloudRegions

`func (o *QueryInventoryResourcesRequestFilters) HasCloudRegions() bool`

HasCloudRegions returns a boolean if a field has been set.

### GetNativeResourceTypes

`func (o *QueryInventoryResourcesRequestFilters) GetNativeResourceTypes() QueryInventoryResourcesRequestFiltersNativeResourceTypes`

GetNativeResourceTypes returns the NativeResourceTypes field if non-nil, zero value otherwise.

### GetNativeResourceTypesOk

`func (o *QueryInventoryResourcesRequestFilters) GetNativeResourceTypesOk() (*QueryInventoryResourcesRequestFiltersNativeResourceTypes, bool)`

GetNativeResourceTypesOk returns a tuple with the NativeResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeResourceTypes

`func (o *QueryInventoryResourcesRequestFilters) SetNativeResourceTypes(v QueryInventoryResourcesRequestFiltersNativeResourceTypes)`

SetNativeResourceTypes sets NativeResourceTypes field to given value.

### HasNativeResourceTypes

`func (o *QueryInventoryResourcesRequestFilters) HasNativeResourceTypes() bool`

HasNativeResourceTypes returns a boolean if a field has been set.

### GetResourceTypes

`func (o *QueryInventoryResourcesRequestFilters) GetResourceTypes() QueryInventoryResourcesRequestFiltersResourceTypes`

GetResourceTypes returns the ResourceTypes field if non-nil, zero value otherwise.

### GetResourceTypesOk

`func (o *QueryInventoryResourcesRequestFilters) GetResourceTypesOk() (*QueryInventoryResourcesRequestFiltersResourceTypes, bool)`

GetResourceTypesOk returns a tuple with the ResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTypes

`func (o *QueryInventoryResourcesRequestFilters) SetResourceTypes(v QueryInventoryResourcesRequestFiltersResourceTypes)`

SetResourceTypes sets ResourceTypes field to given value.

### HasResourceTypes

`func (o *QueryInventoryResourcesRequestFilters) HasResourceTypes() bool`

HasResourceTypes returns a boolean if a field has been set.

### GetTags

`func (o *QueryInventoryResourcesRequestFilters) GetTags() QueryInventoryResourcesRequestFiltersTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *QueryInventoryResourcesRequestFilters) GetTagsOk() (*QueryInventoryResourcesRequestFiltersTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *QueryInventoryResourcesRequestFilters) SetTags(v QueryInventoryResourcesRequestFiltersTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *QueryInventoryResourcesRequestFilters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsIacManaged

`func (o *QueryInventoryResourcesRequestFilters) GetIsIacManaged() bool`

GetIsIacManaged returns the IsIacManaged field if non-nil, zero value otherwise.

### GetIsIacManagedOk

`func (o *QueryInventoryResourcesRequestFilters) GetIsIacManagedOk() (*bool, bool)`

GetIsIacManagedOk returns a tuple with the IsIacManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIacManaged

`func (o *QueryInventoryResourcesRequestFilters) SetIsIacManaged(v bool)`

SetIsIacManaged sets IsIacManaged field to given value.

### HasIsIacManaged

`func (o *QueryInventoryResourcesRequestFilters) HasIsIacManaged() bool`

HasIsIacManaged returns a boolean if a field has been set.

### GetSearchQuery

`func (o *QueryInventoryResourcesRequestFilters) GetSearchQuery() QueryInventoryResourcesRequestFiltersSearchQuery`

GetSearchQuery returns the SearchQuery field if non-nil, zero value otherwise.

### GetSearchQueryOk

`func (o *QueryInventoryResourcesRequestFilters) GetSearchQueryOk() (*QueryInventoryResourcesRequestFiltersSearchQuery, bool)`

GetSearchQueryOk returns a tuple with the SearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchQuery

`func (o *QueryInventoryResourcesRequestFilters) SetSearchQuery(v QueryInventoryResourcesRequestFiltersSearchQuery)`

SetSearchQuery sets SearchQuery field to given value.

### HasSearchQuery

`func (o *QueryInventoryResourcesRequestFilters) HasSearchQuery() bool`

HasSearchQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


