# CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAttribute** | **string** | Attribute that will receive the formatted value | 
**FormatString** | **string** | Template string used to build the target attribute value | 
**OverrideTarget** | Pointer to **bool** | Set to true to overwrite an existing target attribute value (defaults to true) | [optional] 
**Subtype** | **string** |  | 
**ReplaceMissingByEmpty** | Pointer to **bool** | Set to true for missing referenced attributes to render as empty strings. Defaults to true. | [optional] 

## Methods

### NewCreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf

`func NewCreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf(targetAttribute string, formatString string, subtype string, ) *CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf`

NewCreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf instantiates a new CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOfWithDefaults

`func NewCreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOfWithDefaults() *CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf`

NewCreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOfWithDefaults instantiates a new CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAttribute

`func (o *CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetFormatString

`func (o *CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf) GetFormatString() string`

GetFormatString returns the FormatString field if non-nil, zero value otherwise.

### GetFormatStringOk

`func (o *CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf) GetFormatStringOk() (*string, bool)`

GetFormatStringOk returns a tuple with the FormatString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatString

`func (o *CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf) SetFormatString(v string)`

SetFormatString sets FormatString field to given value.


### GetOverrideTarget

`func (o *CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf) GetOverrideTarget() bool`

GetOverrideTarget returns the OverrideTarget field if non-nil, zero value otherwise.

### GetOverrideTargetOk

`func (o *CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf) GetOverrideTargetOk() (*bool, bool)`

GetOverrideTargetOk returns a tuple with the OverrideTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideTarget

`func (o *CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf) SetOverrideTarget(v bool)`

SetOverrideTarget sets OverrideTarget field to given value.

### HasOverrideTarget

`func (o *CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf) HasOverrideTarget() bool`

HasOverrideTarget returns a boolean if a field has been set.

### GetSubtype

`func (o *CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetReplaceMissingByEmpty

`func (o *CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf) GetReplaceMissingByEmpty() bool`

GetReplaceMissingByEmpty returns the ReplaceMissingByEmpty field if non-nil, zero value otherwise.

### GetReplaceMissingByEmptyOk

`func (o *CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf) GetReplaceMissingByEmptyOk() (*bool, bool)`

GetReplaceMissingByEmptyOk returns a tuple with the ReplaceMissingByEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceMissingByEmpty

`func (o *CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf) SetReplaceMissingByEmpty(v bool)`

SetReplaceMissingByEmpty sets ReplaceMissingByEmpty field to given value.

### HasReplaceMissingByEmpty

`func (o *CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf) HasReplaceMissingByEmpty() bool`

HasReplaceMissingByEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


