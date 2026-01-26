# ProcessorAnyOf2Params

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAttribute** | **string** | Attribute that will receive the computed value | 
**FormatString** | **string** | Template string used to build the target attribute value | 
**OverrideTarget** | Pointer to **bool** | Set to true to overwrite an existing target attribute value (defaults to true) | [optional] 
**Subtype** | **string** |  | 
**ReplaceMissingByEmpty** | Pointer to **bool** |  | [optional] 
**Formula** | **string** | Mathematical formula evaluated to populate the target attribute | 
**ReplaceMissingBy0** | Pointer to **bool** |  | [optional] 

## Methods

### NewProcessorAnyOf2Params

`func NewProcessorAnyOf2Params(targetAttribute string, formatString string, subtype string, formula string, ) *ProcessorAnyOf2Params`

NewProcessorAnyOf2Params instantiates a new ProcessorAnyOf2Params object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorAnyOf2ParamsWithDefaults

`func NewProcessorAnyOf2ParamsWithDefaults() *ProcessorAnyOf2Params`

NewProcessorAnyOf2ParamsWithDefaults instantiates a new ProcessorAnyOf2Params object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAttribute

`func (o *ProcessorAnyOf2Params) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *ProcessorAnyOf2Params) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *ProcessorAnyOf2Params) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetFormatString

`func (o *ProcessorAnyOf2Params) GetFormatString() string`

GetFormatString returns the FormatString field if non-nil, zero value otherwise.

### GetFormatStringOk

`func (o *ProcessorAnyOf2Params) GetFormatStringOk() (*string, bool)`

GetFormatStringOk returns a tuple with the FormatString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatString

`func (o *ProcessorAnyOf2Params) SetFormatString(v string)`

SetFormatString sets FormatString field to given value.


### GetOverrideTarget

`func (o *ProcessorAnyOf2Params) GetOverrideTarget() bool`

GetOverrideTarget returns the OverrideTarget field if non-nil, zero value otherwise.

### GetOverrideTargetOk

`func (o *ProcessorAnyOf2Params) GetOverrideTargetOk() (*bool, bool)`

GetOverrideTargetOk returns a tuple with the OverrideTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideTarget

`func (o *ProcessorAnyOf2Params) SetOverrideTarget(v bool)`

SetOverrideTarget sets OverrideTarget field to given value.

### HasOverrideTarget

`func (o *ProcessorAnyOf2Params) HasOverrideTarget() bool`

HasOverrideTarget returns a boolean if a field has been set.

### GetSubtype

`func (o *ProcessorAnyOf2Params) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *ProcessorAnyOf2Params) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *ProcessorAnyOf2Params) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetReplaceMissingByEmpty

`func (o *ProcessorAnyOf2Params) GetReplaceMissingByEmpty() bool`

GetReplaceMissingByEmpty returns the ReplaceMissingByEmpty field if non-nil, zero value otherwise.

### GetReplaceMissingByEmptyOk

`func (o *ProcessorAnyOf2Params) GetReplaceMissingByEmptyOk() (*bool, bool)`

GetReplaceMissingByEmptyOk returns a tuple with the ReplaceMissingByEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceMissingByEmpty

`func (o *ProcessorAnyOf2Params) SetReplaceMissingByEmpty(v bool)`

SetReplaceMissingByEmpty sets ReplaceMissingByEmpty field to given value.

### HasReplaceMissingByEmpty

`func (o *ProcessorAnyOf2Params) HasReplaceMissingByEmpty() bool`

HasReplaceMissingByEmpty returns a boolean if a field has been set.

### GetFormula

`func (o *ProcessorAnyOf2Params) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *ProcessorAnyOf2Params) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *ProcessorAnyOf2Params) SetFormula(v string)`

SetFormula sets Formula field to given value.


### GetReplaceMissingBy0

`func (o *ProcessorAnyOf2Params) GetReplaceMissingBy0() bool`

GetReplaceMissingBy0 returns the ReplaceMissingBy0 field if non-nil, zero value otherwise.

### GetReplaceMissingBy0Ok

`func (o *ProcessorAnyOf2Params) GetReplaceMissingBy0Ok() (*bool, bool)`

GetReplaceMissingBy0Ok returns a tuple with the ReplaceMissingBy0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceMissingBy0

`func (o *ProcessorAnyOf2Params) SetReplaceMissingBy0(v bool)`

SetReplaceMissingBy0 sets ReplaceMissingBy0 field to given value.

### HasReplaceMissingBy0

`func (o *ProcessorAnyOf2Params) HasReplaceMissingBy0() bool`

HasReplaceMissingBy0 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


