# MetricRouteProcessorsInnerOneOf2ParamsOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAttribute** | **string** | Attribute that will receive the category value | 
**Clauses** | [**[]MetricRouteProcessorsInnerOneOf2ParamsOneOf1ClausesInner**](MetricRouteProcessorsInnerOneOf2ParamsOneOf1ClausesInner.md) | Conditions evaluated in order to determine the category value. | 
**DefaultValue** | Pointer to **string** | String category value used when no clause matches. | [optional] 
**Subtype** | **string** |  | 

## Methods

### NewMetricRouteProcessorsInnerOneOf2ParamsOneOf1

`func NewMetricRouteProcessorsInnerOneOf2ParamsOneOf1(targetAttribute string, clauses []MetricRouteProcessorsInnerOneOf2ParamsOneOf1ClausesInner, subtype string, ) *MetricRouteProcessorsInnerOneOf2ParamsOneOf1`

NewMetricRouteProcessorsInnerOneOf2ParamsOneOf1 instantiates a new MetricRouteProcessorsInnerOneOf2ParamsOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricRouteProcessorsInnerOneOf2ParamsOneOf1WithDefaults

`func NewMetricRouteProcessorsInnerOneOf2ParamsOneOf1WithDefaults() *MetricRouteProcessorsInnerOneOf2ParamsOneOf1`

NewMetricRouteProcessorsInnerOneOf2ParamsOneOf1WithDefaults instantiates a new MetricRouteProcessorsInnerOneOf2ParamsOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAttribute

`func (o *MetricRouteProcessorsInnerOneOf2ParamsOneOf1) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *MetricRouteProcessorsInnerOneOf2ParamsOneOf1) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *MetricRouteProcessorsInnerOneOf2ParamsOneOf1) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetClauses

`func (o *MetricRouteProcessorsInnerOneOf2ParamsOneOf1) GetClauses() []MetricRouteProcessorsInnerOneOf2ParamsOneOf1ClausesInner`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *MetricRouteProcessorsInnerOneOf2ParamsOneOf1) GetClausesOk() (*[]MetricRouteProcessorsInnerOneOf2ParamsOneOf1ClausesInner, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *MetricRouteProcessorsInnerOneOf2ParamsOneOf1) SetClauses(v []MetricRouteProcessorsInnerOneOf2ParamsOneOf1ClausesInner)`

SetClauses sets Clauses field to given value.


### GetDefaultValue

`func (o *MetricRouteProcessorsInnerOneOf2ParamsOneOf1) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *MetricRouteProcessorsInnerOneOf2ParamsOneOf1) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *MetricRouteProcessorsInnerOneOf2ParamsOneOf1) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *MetricRouteProcessorsInnerOneOf2ParamsOneOf1) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetSubtype

`func (o *MetricRouteProcessorsInnerOneOf2ParamsOneOf1) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *MetricRouteProcessorsInnerOneOf2ParamsOneOf1) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *MetricRouteProcessorsInnerOneOf2ParamsOneOf1) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


