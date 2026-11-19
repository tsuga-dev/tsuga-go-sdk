# InputMetricsCreateProcessorParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAttribute** | **string** | Attribute that will receive the category value | 
**FormatString** | **string** | Template string used to build the target attribute value | 
**OverrideTarget** | Pointer to **bool** | Set to true to overwrite an existing target attribute value (defaults to true) | [optional] 
**Subtype** | **string** |  | 
**ReplaceMissingByEmpty** | Pointer to **bool** | Set to true for missing referenced attributes to render as empty strings. Defaults to true. | [optional] 
**Clauses** | [**[]InputProcessorParamsCreatorCategoryClausesInner**](InputProcessorParamsCreatorCategoryClausesInner.md) | Conditions evaluated in order to determine the category value. The first matching clause wins. | 
**DefaultValue** | Pointer to **string** | String category value used when no clause matches. | [optional] 

## Methods

### NewInputMetricsCreateProcessorParams

`func NewInputMetricsCreateProcessorParams(targetAttribute string, formatString string, subtype string, clauses []InputProcessorParamsCreatorCategoryClausesInner, ) *InputMetricsCreateProcessorParams`

NewInputMetricsCreateProcessorParams instantiates a new InputMetricsCreateProcessorParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputMetricsCreateProcessorParamsWithDefaults

`func NewInputMetricsCreateProcessorParamsWithDefaults() *InputMetricsCreateProcessorParams`

NewInputMetricsCreateProcessorParamsWithDefaults instantiates a new InputMetricsCreateProcessorParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAttribute

`func (o *InputMetricsCreateProcessorParams) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *InputMetricsCreateProcessorParams) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *InputMetricsCreateProcessorParams) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetFormatString

`func (o *InputMetricsCreateProcessorParams) GetFormatString() string`

GetFormatString returns the FormatString field if non-nil, zero value otherwise.

### GetFormatStringOk

`func (o *InputMetricsCreateProcessorParams) GetFormatStringOk() (*string, bool)`

GetFormatStringOk returns a tuple with the FormatString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatString

`func (o *InputMetricsCreateProcessorParams) SetFormatString(v string)`

SetFormatString sets FormatString field to given value.


### GetOverrideTarget

`func (o *InputMetricsCreateProcessorParams) GetOverrideTarget() bool`

GetOverrideTarget returns the OverrideTarget field if non-nil, zero value otherwise.

### GetOverrideTargetOk

`func (o *InputMetricsCreateProcessorParams) GetOverrideTargetOk() (*bool, bool)`

GetOverrideTargetOk returns a tuple with the OverrideTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideTarget

`func (o *InputMetricsCreateProcessorParams) SetOverrideTarget(v bool)`

SetOverrideTarget sets OverrideTarget field to given value.

### HasOverrideTarget

`func (o *InputMetricsCreateProcessorParams) HasOverrideTarget() bool`

HasOverrideTarget returns a boolean if a field has been set.

### GetSubtype

`func (o *InputMetricsCreateProcessorParams) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *InputMetricsCreateProcessorParams) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *InputMetricsCreateProcessorParams) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetReplaceMissingByEmpty

`func (o *InputMetricsCreateProcessorParams) GetReplaceMissingByEmpty() bool`

GetReplaceMissingByEmpty returns the ReplaceMissingByEmpty field if non-nil, zero value otherwise.

### GetReplaceMissingByEmptyOk

`func (o *InputMetricsCreateProcessorParams) GetReplaceMissingByEmptyOk() (*bool, bool)`

GetReplaceMissingByEmptyOk returns a tuple with the ReplaceMissingByEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceMissingByEmpty

`func (o *InputMetricsCreateProcessorParams) SetReplaceMissingByEmpty(v bool)`

SetReplaceMissingByEmpty sets ReplaceMissingByEmpty field to given value.

### HasReplaceMissingByEmpty

`func (o *InputMetricsCreateProcessorParams) HasReplaceMissingByEmpty() bool`

HasReplaceMissingByEmpty returns a boolean if a field has been set.

### GetClauses

`func (o *InputMetricsCreateProcessorParams) GetClauses() []InputProcessorParamsCreatorCategoryClausesInner`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *InputMetricsCreateProcessorParams) GetClausesOk() (*[]InputProcessorParamsCreatorCategoryClausesInner, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *InputMetricsCreateProcessorParams) SetClauses(v []InputProcessorParamsCreatorCategoryClausesInner)`

SetClauses sets Clauses field to given value.


### GetDefaultValue

`func (o *InputMetricsCreateProcessorParams) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *InputMetricsCreateProcessorParams) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *InputMetricsCreateProcessorParams) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *InputMetricsCreateProcessorParams) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


