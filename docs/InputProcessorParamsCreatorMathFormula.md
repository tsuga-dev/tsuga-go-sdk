# InputProcessorParamsCreatorMathFormula

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAttribute** | **string** | Attribute that will receive the computed value | 
**Formula** | **string** | Mathematical formula evaluated to populate the target attribute | 
**OverrideTarget** | Pointer to **bool** | Set to true to overwrite an existing target attribute value (defaults to true) | [optional] 
**Subtype** | **string** | Evaluates a mathematical expression and writes the numeric result. | 
**ReplaceMissingBy0** | Pointer to **bool** | Set to true for missing referenced attributes to evaluate as zero. Defaults to true. | [optional] 

## Methods

### NewInputProcessorParamsCreatorMathFormula

`func NewInputProcessorParamsCreatorMathFormula(targetAttribute string, formula string, subtype string, ) *InputProcessorParamsCreatorMathFormula`

NewInputProcessorParamsCreatorMathFormula instantiates a new InputProcessorParamsCreatorMathFormula object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputProcessorParamsCreatorMathFormulaWithDefaults

`func NewInputProcessorParamsCreatorMathFormulaWithDefaults() *InputProcessorParamsCreatorMathFormula`

NewInputProcessorParamsCreatorMathFormulaWithDefaults instantiates a new InputProcessorParamsCreatorMathFormula object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAttribute

`func (o *InputProcessorParamsCreatorMathFormula) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *InputProcessorParamsCreatorMathFormula) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *InputProcessorParamsCreatorMathFormula) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetFormula

`func (o *InputProcessorParamsCreatorMathFormula) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *InputProcessorParamsCreatorMathFormula) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *InputProcessorParamsCreatorMathFormula) SetFormula(v string)`

SetFormula sets Formula field to given value.


### GetOverrideTarget

`func (o *InputProcessorParamsCreatorMathFormula) GetOverrideTarget() bool`

GetOverrideTarget returns the OverrideTarget field if non-nil, zero value otherwise.

### GetOverrideTargetOk

`func (o *InputProcessorParamsCreatorMathFormula) GetOverrideTargetOk() (*bool, bool)`

GetOverrideTargetOk returns a tuple with the OverrideTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideTarget

`func (o *InputProcessorParamsCreatorMathFormula) SetOverrideTarget(v bool)`

SetOverrideTarget sets OverrideTarget field to given value.

### HasOverrideTarget

`func (o *InputProcessorParamsCreatorMathFormula) HasOverrideTarget() bool`

HasOverrideTarget returns a boolean if a field has been set.

### GetSubtype

`func (o *InputProcessorParamsCreatorMathFormula) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *InputProcessorParamsCreatorMathFormula) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *InputProcessorParamsCreatorMathFormula) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetReplaceMissingBy0

`func (o *InputProcessorParamsCreatorMathFormula) GetReplaceMissingBy0() bool`

GetReplaceMissingBy0 returns the ReplaceMissingBy0 field if non-nil, zero value otherwise.

### GetReplaceMissingBy0Ok

`func (o *InputProcessorParamsCreatorMathFormula) GetReplaceMissingBy0Ok() (*bool, bool)`

GetReplaceMissingBy0Ok returns a tuple with the ReplaceMissingBy0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceMissingBy0

`func (o *InputProcessorParamsCreatorMathFormula) SetReplaceMissingBy0(v bool)`

SetReplaceMissingBy0 sets ReplaceMissingBy0 field to given value.

### HasReplaceMissingBy0

`func (o *InputProcessorParamsCreatorMathFormula) HasReplaceMissingBy0() bool`

HasReplaceMissingBy0 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


