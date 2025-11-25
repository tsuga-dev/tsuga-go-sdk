# ProcessorParamsCreatorMathFormula

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAttribute** | **string** | Attribute that will receive the computed value | 
**Formula** | **string** | Mathematical formula evaluated to populate the target attribute | 
**OverrideTarget** | Pointer to **bool** | Set to true to overwrite an existing target attribute value (defaults to true) | [optional] 
**Subtype** | **string** |  | 
**ReplaceMissingBy0** | Pointer to **bool** |  | [optional] 

## Methods

### NewProcessorParamsCreatorMathFormula

`func NewProcessorParamsCreatorMathFormula(targetAttribute string, formula string, subtype string, ) *ProcessorParamsCreatorMathFormula`

NewProcessorParamsCreatorMathFormula instantiates a new ProcessorParamsCreatorMathFormula object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorParamsCreatorMathFormulaWithDefaults

`func NewProcessorParamsCreatorMathFormulaWithDefaults() *ProcessorParamsCreatorMathFormula`

NewProcessorParamsCreatorMathFormulaWithDefaults instantiates a new ProcessorParamsCreatorMathFormula object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAttribute

`func (o *ProcessorParamsCreatorMathFormula) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *ProcessorParamsCreatorMathFormula) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *ProcessorParamsCreatorMathFormula) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetFormula

`func (o *ProcessorParamsCreatorMathFormula) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *ProcessorParamsCreatorMathFormula) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *ProcessorParamsCreatorMathFormula) SetFormula(v string)`

SetFormula sets Formula field to given value.


### GetOverrideTarget

`func (o *ProcessorParamsCreatorMathFormula) GetOverrideTarget() bool`

GetOverrideTarget returns the OverrideTarget field if non-nil, zero value otherwise.

### GetOverrideTargetOk

`func (o *ProcessorParamsCreatorMathFormula) GetOverrideTargetOk() (*bool, bool)`

GetOverrideTargetOk returns a tuple with the OverrideTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideTarget

`func (o *ProcessorParamsCreatorMathFormula) SetOverrideTarget(v bool)`

SetOverrideTarget sets OverrideTarget field to given value.

### HasOverrideTarget

`func (o *ProcessorParamsCreatorMathFormula) HasOverrideTarget() bool`

HasOverrideTarget returns a boolean if a field has been set.

### GetSubtype

`func (o *ProcessorParamsCreatorMathFormula) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *ProcessorParamsCreatorMathFormula) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *ProcessorParamsCreatorMathFormula) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetReplaceMissingBy0

`func (o *ProcessorParamsCreatorMathFormula) GetReplaceMissingBy0() bool`

GetReplaceMissingBy0 returns the ReplaceMissingBy0 field if non-nil, zero value otherwise.

### GetReplaceMissingBy0Ok

`func (o *ProcessorParamsCreatorMathFormula) GetReplaceMissingBy0Ok() (*bool, bool)`

GetReplaceMissingBy0Ok returns a tuple with the ReplaceMissingBy0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceMissingBy0

`func (o *ProcessorParamsCreatorMathFormula) SetReplaceMissingBy0(v bool)`

SetReplaceMissingBy0 sets ReplaceMissingBy0 field to given value.

### HasReplaceMissingBy0

`func (o *ProcessorParamsCreatorMathFormula) HasReplaceMissingBy0() bool`

HasReplaceMissingBy0 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


