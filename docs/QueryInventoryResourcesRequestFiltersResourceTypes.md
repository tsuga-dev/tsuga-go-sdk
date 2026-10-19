# QueryInventoryResourcesRequestFiltersResourceTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | **[]string** | Values compared by this filter. When &#x60;exclude&#x60; is omitted or false, rows matching any value are kept; when &#x60;exclude&#x60; is true, matching rows are removed. | 
**Exclude** | Pointer to **bool** | If true, exclude rows matching these values instead of including them | [optional] 

## Methods

### NewQueryInventoryResourcesRequestFiltersResourceTypes

`func NewQueryInventoryResourcesRequestFiltersResourceTypes(values []string, ) *QueryInventoryResourcesRequestFiltersResourceTypes`

NewQueryInventoryResourcesRequestFiltersResourceTypes instantiates a new QueryInventoryResourcesRequestFiltersResourceTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryInventoryResourcesRequestFiltersResourceTypesWithDefaults

`func NewQueryInventoryResourcesRequestFiltersResourceTypesWithDefaults() *QueryInventoryResourcesRequestFiltersResourceTypes`

NewQueryInventoryResourcesRequestFiltersResourceTypesWithDefaults instantiates a new QueryInventoryResourcesRequestFiltersResourceTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *QueryInventoryResourcesRequestFiltersResourceTypes) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *QueryInventoryResourcesRequestFiltersResourceTypes) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *QueryInventoryResourcesRequestFiltersResourceTypes) SetValues(v []string)`

SetValues sets Values field to given value.


### GetExclude

`func (o *QueryInventoryResourcesRequestFiltersResourceTypes) GetExclude() bool`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *QueryInventoryResourcesRequestFiltersResourceTypes) GetExcludeOk() (*bool, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *QueryInventoryResourcesRequestFiltersResourceTypes) SetExclude(v bool)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *QueryInventoryResourcesRequestFiltersResourceTypes) HasExclude() bool`

HasExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


