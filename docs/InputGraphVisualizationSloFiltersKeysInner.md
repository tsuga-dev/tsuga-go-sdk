# InputGraphVisualizationSloFiltersKeysInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Group-by field or tag key to filter on, such as &#x60;context.team&#x60; | 
**Values** | **[]string** | Values compared with the result&#39;s group-by value and its SLO&#39;s tag values for &#x60;key&#x60; | 
**Exclude** | Pointer to **bool** | If true, keep the results with none of these values instead of those with at least one | [optional] 

## Methods

### NewInputGraphVisualizationSloFiltersKeysInner

`func NewInputGraphVisualizationSloFiltersKeysInner(key string, values []string, ) *InputGraphVisualizationSloFiltersKeysInner`

NewInputGraphVisualizationSloFiltersKeysInner instantiates a new InputGraphVisualizationSloFiltersKeysInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationSloFiltersKeysInnerWithDefaults

`func NewInputGraphVisualizationSloFiltersKeysInnerWithDefaults() *InputGraphVisualizationSloFiltersKeysInner`

NewInputGraphVisualizationSloFiltersKeysInnerWithDefaults instantiates a new InputGraphVisualizationSloFiltersKeysInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *InputGraphVisualizationSloFiltersKeysInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InputGraphVisualizationSloFiltersKeysInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InputGraphVisualizationSloFiltersKeysInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetValues

`func (o *InputGraphVisualizationSloFiltersKeysInner) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *InputGraphVisualizationSloFiltersKeysInner) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *InputGraphVisualizationSloFiltersKeysInner) SetValues(v []string)`

SetValues sets Values field to given value.


### GetExclude

`func (o *InputGraphVisualizationSloFiltersKeysInner) GetExclude() bool`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *InputGraphVisualizationSloFiltersKeysInner) GetExcludeOk() (*bool, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *InputGraphVisualizationSloFiltersKeysInner) SetExclude(v bool)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *InputGraphVisualizationSloFiltersKeysInner) HasExclude() bool`

HasExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


