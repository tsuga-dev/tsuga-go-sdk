# GraphVisualizationSloFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owners** | Pointer to [**GraphVisualizationSloFiltersOwners**](GraphVisualizationSloFiltersOwners.md) |  | [optional] 
**Tags** | Pointer to [**GraphVisualizationSloFiltersTags**](GraphVisualizationSloFiltersTags.md) |  | [optional] 
**Statuses** | Pointer to [**GraphVisualizationSloFiltersStatuses**](GraphVisualizationSloFiltersStatuses.md) |  | [optional] 
**Keys** | Pointer to [**[]GraphVisualizationSloFiltersKeysInner**](GraphVisualizationSloFiltersKeysInner.md) | Keep only the results matching every key filter. A result&#39;s values for a key are its group-by value when its SLO groups by that field, and its SLO&#39;s tag values for that key. A result matches when one of its values is in &#x60;values&#x60;, or, with &#x60;exclude: true&#x60;, when none is; a result with no value for the key only matches an exclusion. | [optional] 

## Methods

### NewGraphVisualizationSloFilters

`func NewGraphVisualizationSloFilters() *GraphVisualizationSloFilters`

NewGraphVisualizationSloFilters instantiates a new GraphVisualizationSloFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationSloFiltersWithDefaults

`func NewGraphVisualizationSloFiltersWithDefaults() *GraphVisualizationSloFilters`

NewGraphVisualizationSloFiltersWithDefaults instantiates a new GraphVisualizationSloFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwners

`func (o *GraphVisualizationSloFilters) GetOwners() GraphVisualizationSloFiltersOwners`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *GraphVisualizationSloFilters) GetOwnersOk() (*GraphVisualizationSloFiltersOwners, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *GraphVisualizationSloFilters) SetOwners(v GraphVisualizationSloFiltersOwners)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *GraphVisualizationSloFilters) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetTags

`func (o *GraphVisualizationSloFilters) GetTags() GraphVisualizationSloFiltersTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GraphVisualizationSloFilters) GetTagsOk() (*GraphVisualizationSloFiltersTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GraphVisualizationSloFilters) SetTags(v GraphVisualizationSloFiltersTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GraphVisualizationSloFilters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetStatuses

`func (o *GraphVisualizationSloFilters) GetStatuses() GraphVisualizationSloFiltersStatuses`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *GraphVisualizationSloFilters) GetStatusesOk() (*GraphVisualizationSloFiltersStatuses, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *GraphVisualizationSloFilters) SetStatuses(v GraphVisualizationSloFiltersStatuses)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *GraphVisualizationSloFilters) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetKeys

`func (o *GraphVisualizationSloFilters) GetKeys() []GraphVisualizationSloFiltersKeysInner`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *GraphVisualizationSloFilters) GetKeysOk() (*[]GraphVisualizationSloFiltersKeysInner, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *GraphVisualizationSloFilters) SetKeys(v []GraphVisualizationSloFiltersKeysInner)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *GraphVisualizationSloFilters) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


