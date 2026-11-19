# MetricsProcessorCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAttribute** | **string** | Attribute that will receive the category value | 
**Clauses** | [**[]MetricsProcessorCategoryClausesInner**](MetricsProcessorCategoryClausesInner.md) | Conditions evaluated in order to determine the category value. | 
**DefaultValue** | Pointer to **string** | String category value used when no clause matches. | [optional] 
**Subtype** | **string** |  | 

## Methods

### NewMetricsProcessorCategory

`func NewMetricsProcessorCategory(targetAttribute string, clauses []MetricsProcessorCategoryClausesInner, subtype string, ) *MetricsProcessorCategory`

NewMetricsProcessorCategory instantiates a new MetricsProcessorCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsProcessorCategoryWithDefaults

`func NewMetricsProcessorCategoryWithDefaults() *MetricsProcessorCategory`

NewMetricsProcessorCategoryWithDefaults instantiates a new MetricsProcessorCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAttribute

`func (o *MetricsProcessorCategory) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *MetricsProcessorCategory) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *MetricsProcessorCategory) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetClauses

`func (o *MetricsProcessorCategory) GetClauses() []MetricsProcessorCategoryClausesInner`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *MetricsProcessorCategory) GetClausesOk() (*[]MetricsProcessorCategoryClausesInner, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *MetricsProcessorCategory) SetClauses(v []MetricsProcessorCategoryClausesInner)`

SetClauses sets Clauses field to given value.


### GetDefaultValue

`func (o *MetricsProcessorCategory) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *MetricsProcessorCategory) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *MetricsProcessorCategory) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *MetricsProcessorCategory) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetSubtype

`func (o *MetricsProcessorCategory) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *MetricsProcessorCategory) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *MetricsProcessorCategory) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


