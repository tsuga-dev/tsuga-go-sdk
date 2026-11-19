# InputMetricsProcessorFormatString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAttribute** | **string** | Attribute that will receive the formatted value | 
**FormatString** | **string** | Template string used to build the target attribute value | 
**OverrideTarget** | Pointer to **bool** | Set to true to overwrite an existing target attribute value (defaults to true) | [optional] 
**Subtype** | **string** |  | 
**ReplaceMissingByEmpty** | Pointer to **bool** | Set to true for missing referenced attributes to render as empty strings. Defaults to true. | [optional] 

## Methods

### NewInputMetricsProcessorFormatString

`func NewInputMetricsProcessorFormatString(targetAttribute string, formatString string, subtype string, ) *InputMetricsProcessorFormatString`

NewInputMetricsProcessorFormatString instantiates a new InputMetricsProcessorFormatString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputMetricsProcessorFormatStringWithDefaults

`func NewInputMetricsProcessorFormatStringWithDefaults() *InputMetricsProcessorFormatString`

NewInputMetricsProcessorFormatStringWithDefaults instantiates a new InputMetricsProcessorFormatString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAttribute

`func (o *InputMetricsProcessorFormatString) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *InputMetricsProcessorFormatString) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *InputMetricsProcessorFormatString) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetFormatString

`func (o *InputMetricsProcessorFormatString) GetFormatString() string`

GetFormatString returns the FormatString field if non-nil, zero value otherwise.

### GetFormatStringOk

`func (o *InputMetricsProcessorFormatString) GetFormatStringOk() (*string, bool)`

GetFormatStringOk returns a tuple with the FormatString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatString

`func (o *InputMetricsProcessorFormatString) SetFormatString(v string)`

SetFormatString sets FormatString field to given value.


### GetOverrideTarget

`func (o *InputMetricsProcessorFormatString) GetOverrideTarget() bool`

GetOverrideTarget returns the OverrideTarget field if non-nil, zero value otherwise.

### GetOverrideTargetOk

`func (o *InputMetricsProcessorFormatString) GetOverrideTargetOk() (*bool, bool)`

GetOverrideTargetOk returns a tuple with the OverrideTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideTarget

`func (o *InputMetricsProcessorFormatString) SetOverrideTarget(v bool)`

SetOverrideTarget sets OverrideTarget field to given value.

### HasOverrideTarget

`func (o *InputMetricsProcessorFormatString) HasOverrideTarget() bool`

HasOverrideTarget returns a boolean if a field has been set.

### GetSubtype

`func (o *InputMetricsProcessorFormatString) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *InputMetricsProcessorFormatString) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *InputMetricsProcessorFormatString) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetReplaceMissingByEmpty

`func (o *InputMetricsProcessorFormatString) GetReplaceMissingByEmpty() bool`

GetReplaceMissingByEmpty returns the ReplaceMissingByEmpty field if non-nil, zero value otherwise.

### GetReplaceMissingByEmptyOk

`func (o *InputMetricsProcessorFormatString) GetReplaceMissingByEmptyOk() (*bool, bool)`

GetReplaceMissingByEmptyOk returns a tuple with the ReplaceMissingByEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceMissingByEmpty

`func (o *InputMetricsProcessorFormatString) SetReplaceMissingByEmpty(v bool)`

SetReplaceMissingByEmpty sets ReplaceMissingByEmpty field to given value.

### HasReplaceMissingByEmpty

`func (o *InputMetricsProcessorFormatString) HasReplaceMissingByEmpty() bool`

HasReplaceMissingByEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


