# InputGraphVisualizationSloFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owners** | Pointer to [**QuerySlosRequestFiltersOwners**](QuerySlosRequestFiltersOwners.md) |  | [optional] 
**Tags** | Pointer to [**QuerySlosRequestFiltersTags**](QuerySlosRequestFiltersTags.md) |  | [optional] 
**Statuses** | Pointer to [**InputGraphVisualizationSloFiltersStatuses**](InputGraphVisualizationSloFiltersStatuses.md) |  | [optional] 
**Keys** | Pointer to [**[]InputGraphVisualizationSloFiltersKeysInner**](InputGraphVisualizationSloFiltersKeysInner.md) | Keep only the results matching every key filter. A result&#39;s values for a key are its group-by value when its SLO groups by that field, and its SLO&#39;s tag values for that key. A result matches when one of its values is in &#x60;values&#x60;, or, with &#x60;exclude: true&#x60;, when none is; a result with no value for the key only matches an exclusion. | [optional] 

## Methods

### NewInputGraphVisualizationSloFilters

`func NewInputGraphVisualizationSloFilters() *InputGraphVisualizationSloFilters`

NewInputGraphVisualizationSloFilters instantiates a new InputGraphVisualizationSloFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationSloFiltersWithDefaults

`func NewInputGraphVisualizationSloFiltersWithDefaults() *InputGraphVisualizationSloFilters`

NewInputGraphVisualizationSloFiltersWithDefaults instantiates a new InputGraphVisualizationSloFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwners

`func (o *InputGraphVisualizationSloFilters) GetOwners() QuerySlosRequestFiltersOwners`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *InputGraphVisualizationSloFilters) GetOwnersOk() (*QuerySlosRequestFiltersOwners, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *InputGraphVisualizationSloFilters) SetOwners(v QuerySlosRequestFiltersOwners)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *InputGraphVisualizationSloFilters) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetTags

`func (o *InputGraphVisualizationSloFilters) GetTags() QuerySlosRequestFiltersTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InputGraphVisualizationSloFilters) GetTagsOk() (*QuerySlosRequestFiltersTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InputGraphVisualizationSloFilters) SetTags(v QuerySlosRequestFiltersTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InputGraphVisualizationSloFilters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetStatuses

`func (o *InputGraphVisualizationSloFilters) GetStatuses() InputGraphVisualizationSloFiltersStatuses`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *InputGraphVisualizationSloFilters) GetStatusesOk() (*InputGraphVisualizationSloFiltersStatuses, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *InputGraphVisualizationSloFilters) SetStatuses(v InputGraphVisualizationSloFiltersStatuses)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *InputGraphVisualizationSloFilters) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetKeys

`func (o *InputGraphVisualizationSloFilters) GetKeys() []InputGraphVisualizationSloFiltersKeysInner`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *InputGraphVisualizationSloFilters) GetKeysOk() (*[]InputGraphVisualizationSloFiltersKeysInner, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *InputGraphVisualizationSloFilters) SetKeys(v []InputGraphVisualizationSloFiltersKeysInner)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *InputGraphVisualizationSloFilters) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


