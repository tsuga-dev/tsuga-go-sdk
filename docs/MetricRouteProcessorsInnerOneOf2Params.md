# MetricRouteProcessorsInnerOneOf2Params

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAttribute** | **string** | Attribute that will receive the category value | 
**FormatString** | **string** | Template string used to build the target attribute value | 
**OverrideTarget** | Pointer to **bool** | Whether to overwrite the target attribute when it already exists. | [optional] 
**Subtype** | **string** |  | 
**ReplaceMissingByEmpty** | Pointer to **bool** | Whether missing referenced attributes render as empty strings. | [optional] 
**Clauses** | [**[]MetricRouteProcessorsInnerOneOf2ParamsOneOf1ClausesInner**](MetricRouteProcessorsInnerOneOf2ParamsOneOf1ClausesInner.md) | Conditions evaluated in order to determine the category value. | 
**DefaultValue** | Pointer to **string** | String category value used when no clause matches. | [optional] 

## Methods

### NewMetricRouteProcessorsInnerOneOf2Params

`func NewMetricRouteProcessorsInnerOneOf2Params(targetAttribute string, formatString string, subtype string, clauses []MetricRouteProcessorsInnerOneOf2ParamsOneOf1ClausesInner, ) *MetricRouteProcessorsInnerOneOf2Params`

NewMetricRouteProcessorsInnerOneOf2Params instantiates a new MetricRouteProcessorsInnerOneOf2Params object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricRouteProcessorsInnerOneOf2ParamsWithDefaults

`func NewMetricRouteProcessorsInnerOneOf2ParamsWithDefaults() *MetricRouteProcessorsInnerOneOf2Params`

NewMetricRouteProcessorsInnerOneOf2ParamsWithDefaults instantiates a new MetricRouteProcessorsInnerOneOf2Params object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAttribute

`func (o *MetricRouteProcessorsInnerOneOf2Params) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *MetricRouteProcessorsInnerOneOf2Params) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *MetricRouteProcessorsInnerOneOf2Params) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetFormatString

`func (o *MetricRouteProcessorsInnerOneOf2Params) GetFormatString() string`

GetFormatString returns the FormatString field if non-nil, zero value otherwise.

### GetFormatStringOk

`func (o *MetricRouteProcessorsInnerOneOf2Params) GetFormatStringOk() (*string, bool)`

GetFormatStringOk returns a tuple with the FormatString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatString

`func (o *MetricRouteProcessorsInnerOneOf2Params) SetFormatString(v string)`

SetFormatString sets FormatString field to given value.


### GetOverrideTarget

`func (o *MetricRouteProcessorsInnerOneOf2Params) GetOverrideTarget() bool`

GetOverrideTarget returns the OverrideTarget field if non-nil, zero value otherwise.

### GetOverrideTargetOk

`func (o *MetricRouteProcessorsInnerOneOf2Params) GetOverrideTargetOk() (*bool, bool)`

GetOverrideTargetOk returns a tuple with the OverrideTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideTarget

`func (o *MetricRouteProcessorsInnerOneOf2Params) SetOverrideTarget(v bool)`

SetOverrideTarget sets OverrideTarget field to given value.

### HasOverrideTarget

`func (o *MetricRouteProcessorsInnerOneOf2Params) HasOverrideTarget() bool`

HasOverrideTarget returns a boolean if a field has been set.

### GetSubtype

`func (o *MetricRouteProcessorsInnerOneOf2Params) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *MetricRouteProcessorsInnerOneOf2Params) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *MetricRouteProcessorsInnerOneOf2Params) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetReplaceMissingByEmpty

`func (o *MetricRouteProcessorsInnerOneOf2Params) GetReplaceMissingByEmpty() bool`

GetReplaceMissingByEmpty returns the ReplaceMissingByEmpty field if non-nil, zero value otherwise.

### GetReplaceMissingByEmptyOk

`func (o *MetricRouteProcessorsInnerOneOf2Params) GetReplaceMissingByEmptyOk() (*bool, bool)`

GetReplaceMissingByEmptyOk returns a tuple with the ReplaceMissingByEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceMissingByEmpty

`func (o *MetricRouteProcessorsInnerOneOf2Params) SetReplaceMissingByEmpty(v bool)`

SetReplaceMissingByEmpty sets ReplaceMissingByEmpty field to given value.

### HasReplaceMissingByEmpty

`func (o *MetricRouteProcessorsInnerOneOf2Params) HasReplaceMissingByEmpty() bool`

HasReplaceMissingByEmpty returns a boolean if a field has been set.

### GetClauses

`func (o *MetricRouteProcessorsInnerOneOf2Params) GetClauses() []MetricRouteProcessorsInnerOneOf2ParamsOneOf1ClausesInner`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *MetricRouteProcessorsInnerOneOf2Params) GetClausesOk() (*[]MetricRouteProcessorsInnerOneOf2ParamsOneOf1ClausesInner, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *MetricRouteProcessorsInnerOneOf2Params) SetClauses(v []MetricRouteProcessorsInnerOneOf2ParamsOneOf1ClausesInner)`

SetClauses sets Clauses field to given value.


### GetDefaultValue

`func (o *MetricRouteProcessorsInnerOneOf2Params) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *MetricRouteProcessorsInnerOneOf2Params) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *MetricRouteProcessorsInnerOneOf2Params) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *MetricRouteProcessorsInnerOneOf2Params) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


