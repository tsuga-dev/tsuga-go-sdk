# OutputProcessorAnyOf2Params

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAttribute** | **string** | Attribute that will receive the category value | 
**FormatString** | **string** | Template string used to build the target attribute value | 
**OverrideTarget** | Pointer to **bool** | Whether to overwrite the target attribute when it already exists. If omitted, route processing treats this as true. | [optional] 
**Subtype** | **string** | Assigns the first matching category value from ordered query clauses, with an optional default. | 
**ReplaceMissingByEmpty** | Pointer to **bool** | Whether missing referenced attributes render as empty strings. If omitted, route processing treats this as true. | [optional] 
**Formula** | **string** | Mathematical formula evaluated to populate the target attribute | 
**ReplaceMissingBy0** | Pointer to **bool** | Whether missing referenced attributes evaluate as zero. If omitted, route processing treats this as true. | [optional] 
**Clauses** | [**[]ProcessorParamsCreatorCategoryClausesInner**](ProcessorParamsCreatorCategoryClausesInner.md) | Conditions evaluated in order to determine the category value. The first matching clause wins. | 
**DefaultValue** | Pointer to **string** | String category value used when no clause matches. If omitted and no clause matches, no category value is written. | [optional] 

## Methods

### NewOutputProcessorAnyOf2Params

`func NewOutputProcessorAnyOf2Params(targetAttribute string, formatString string, subtype string, formula string, clauses []ProcessorParamsCreatorCategoryClausesInner, ) *OutputProcessorAnyOf2Params`

NewOutputProcessorAnyOf2Params instantiates a new OutputProcessorAnyOf2Params object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputProcessorAnyOf2ParamsWithDefaults

`func NewOutputProcessorAnyOf2ParamsWithDefaults() *OutputProcessorAnyOf2Params`

NewOutputProcessorAnyOf2ParamsWithDefaults instantiates a new OutputProcessorAnyOf2Params object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAttribute

`func (o *OutputProcessorAnyOf2Params) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *OutputProcessorAnyOf2Params) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *OutputProcessorAnyOf2Params) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetFormatString

`func (o *OutputProcessorAnyOf2Params) GetFormatString() string`

GetFormatString returns the FormatString field if non-nil, zero value otherwise.

### GetFormatStringOk

`func (o *OutputProcessorAnyOf2Params) GetFormatStringOk() (*string, bool)`

GetFormatStringOk returns a tuple with the FormatString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatString

`func (o *OutputProcessorAnyOf2Params) SetFormatString(v string)`

SetFormatString sets FormatString field to given value.


### GetOverrideTarget

`func (o *OutputProcessorAnyOf2Params) GetOverrideTarget() bool`

GetOverrideTarget returns the OverrideTarget field if non-nil, zero value otherwise.

### GetOverrideTargetOk

`func (o *OutputProcessorAnyOf2Params) GetOverrideTargetOk() (*bool, bool)`

GetOverrideTargetOk returns a tuple with the OverrideTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideTarget

`func (o *OutputProcessorAnyOf2Params) SetOverrideTarget(v bool)`

SetOverrideTarget sets OverrideTarget field to given value.

### HasOverrideTarget

`func (o *OutputProcessorAnyOf2Params) HasOverrideTarget() bool`

HasOverrideTarget returns a boolean if a field has been set.

### GetSubtype

`func (o *OutputProcessorAnyOf2Params) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *OutputProcessorAnyOf2Params) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *OutputProcessorAnyOf2Params) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetReplaceMissingByEmpty

`func (o *OutputProcessorAnyOf2Params) GetReplaceMissingByEmpty() bool`

GetReplaceMissingByEmpty returns the ReplaceMissingByEmpty field if non-nil, zero value otherwise.

### GetReplaceMissingByEmptyOk

`func (o *OutputProcessorAnyOf2Params) GetReplaceMissingByEmptyOk() (*bool, bool)`

GetReplaceMissingByEmptyOk returns a tuple with the ReplaceMissingByEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceMissingByEmpty

`func (o *OutputProcessorAnyOf2Params) SetReplaceMissingByEmpty(v bool)`

SetReplaceMissingByEmpty sets ReplaceMissingByEmpty field to given value.

### HasReplaceMissingByEmpty

`func (o *OutputProcessorAnyOf2Params) HasReplaceMissingByEmpty() bool`

HasReplaceMissingByEmpty returns a boolean if a field has been set.

### GetFormula

`func (o *OutputProcessorAnyOf2Params) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *OutputProcessorAnyOf2Params) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *OutputProcessorAnyOf2Params) SetFormula(v string)`

SetFormula sets Formula field to given value.


### GetReplaceMissingBy0

`func (o *OutputProcessorAnyOf2Params) GetReplaceMissingBy0() bool`

GetReplaceMissingBy0 returns the ReplaceMissingBy0 field if non-nil, zero value otherwise.

### GetReplaceMissingBy0Ok

`func (o *OutputProcessorAnyOf2Params) GetReplaceMissingBy0Ok() (*bool, bool)`

GetReplaceMissingBy0Ok returns a tuple with the ReplaceMissingBy0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceMissingBy0

`func (o *OutputProcessorAnyOf2Params) SetReplaceMissingBy0(v bool)`

SetReplaceMissingBy0 sets ReplaceMissingBy0 field to given value.

### HasReplaceMissingBy0

`func (o *OutputProcessorAnyOf2Params) HasReplaceMissingBy0() bool`

HasReplaceMissingBy0 returns a boolean if a field has been set.

### GetClauses

`func (o *OutputProcessorAnyOf2Params) GetClauses() []ProcessorParamsCreatorCategoryClausesInner`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *OutputProcessorAnyOf2Params) GetClausesOk() (*[]ProcessorParamsCreatorCategoryClausesInner, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *OutputProcessorAnyOf2Params) SetClauses(v []ProcessorParamsCreatorCategoryClausesInner)`

SetClauses sets Clauses field to given value.


### GetDefaultValue

`func (o *OutputProcessorAnyOf2Params) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *OutputProcessorAnyOf2Params) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *OutputProcessorAnyOf2Params) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *OutputProcessorAnyOf2Params) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


