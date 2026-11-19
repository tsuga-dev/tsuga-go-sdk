# MetricsProcessorFormatString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAttribute** | **string** | Attribute that will receive the formatted value | 
**FormatString** | **string** | Template string used to build the target attribute value | 
**OverrideTarget** | Pointer to **bool** | Whether to overwrite the target attribute when it already exists. | [optional] 
**Subtype** | **string** |  | 
**ReplaceMissingByEmpty** | Pointer to **bool** | Whether missing referenced attributes render as empty strings. | [optional] 

## Methods

### NewMetricsProcessorFormatString

`func NewMetricsProcessorFormatString(targetAttribute string, formatString string, subtype string, ) *MetricsProcessorFormatString`

NewMetricsProcessorFormatString instantiates a new MetricsProcessorFormatString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsProcessorFormatStringWithDefaults

`func NewMetricsProcessorFormatStringWithDefaults() *MetricsProcessorFormatString`

NewMetricsProcessorFormatStringWithDefaults instantiates a new MetricsProcessorFormatString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAttribute

`func (o *MetricsProcessorFormatString) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *MetricsProcessorFormatString) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *MetricsProcessorFormatString) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetFormatString

`func (o *MetricsProcessorFormatString) GetFormatString() string`

GetFormatString returns the FormatString field if non-nil, zero value otherwise.

### GetFormatStringOk

`func (o *MetricsProcessorFormatString) GetFormatStringOk() (*string, bool)`

GetFormatStringOk returns a tuple with the FormatString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatString

`func (o *MetricsProcessorFormatString) SetFormatString(v string)`

SetFormatString sets FormatString field to given value.


### GetOverrideTarget

`func (o *MetricsProcessorFormatString) GetOverrideTarget() bool`

GetOverrideTarget returns the OverrideTarget field if non-nil, zero value otherwise.

### GetOverrideTargetOk

`func (o *MetricsProcessorFormatString) GetOverrideTargetOk() (*bool, bool)`

GetOverrideTargetOk returns a tuple with the OverrideTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideTarget

`func (o *MetricsProcessorFormatString) SetOverrideTarget(v bool)`

SetOverrideTarget sets OverrideTarget field to given value.

### HasOverrideTarget

`func (o *MetricsProcessorFormatString) HasOverrideTarget() bool`

HasOverrideTarget returns a boolean if a field has been set.

### GetSubtype

`func (o *MetricsProcessorFormatString) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *MetricsProcessorFormatString) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *MetricsProcessorFormatString) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetReplaceMissingByEmpty

`func (o *MetricsProcessorFormatString) GetReplaceMissingByEmpty() bool`

GetReplaceMissingByEmpty returns the ReplaceMissingByEmpty field if non-nil, zero value otherwise.

### GetReplaceMissingByEmptyOk

`func (o *MetricsProcessorFormatString) GetReplaceMissingByEmptyOk() (*bool, bool)`

GetReplaceMissingByEmptyOk returns a tuple with the ReplaceMissingByEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceMissingByEmpty

`func (o *MetricsProcessorFormatString) SetReplaceMissingByEmpty(v bool)`

SetReplaceMissingByEmpty sets ReplaceMissingByEmpty field to given value.

### HasReplaceMissingByEmpty

`func (o *MetricsProcessorFormatString) HasReplaceMissingByEmpty() bool`

HasReplaceMissingByEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


