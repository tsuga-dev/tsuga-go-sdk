# InputProcessorParamsCreatorCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAttribute** | **string** | Attribute that will receive the category value | 
**Clauses** | [**[]CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf1ClausesInner**](CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf1ClausesInner.md) | Conditions evaluated in order to determine the category value. The first matching clause wins. | 
**DefaultValue** | Pointer to **string** | String category value used when no clause matches. If omitted and no clause matches, no category value is written. | [optional] 
**Subtype** | **string** | Assigns the first matching category value from ordered query clauses, with an optional default. | 

## Methods

### NewInputProcessorParamsCreatorCategory

`func NewInputProcessorParamsCreatorCategory(targetAttribute string, clauses []CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf1ClausesInner, subtype string, ) *InputProcessorParamsCreatorCategory`

NewInputProcessorParamsCreatorCategory instantiates a new InputProcessorParamsCreatorCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputProcessorParamsCreatorCategoryWithDefaults

`func NewInputProcessorParamsCreatorCategoryWithDefaults() *InputProcessorParamsCreatorCategory`

NewInputProcessorParamsCreatorCategoryWithDefaults instantiates a new InputProcessorParamsCreatorCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAttribute

`func (o *InputProcessorParamsCreatorCategory) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *InputProcessorParamsCreatorCategory) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *InputProcessorParamsCreatorCategory) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetClauses

`func (o *InputProcessorParamsCreatorCategory) GetClauses() []CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf1ClausesInner`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *InputProcessorParamsCreatorCategory) GetClausesOk() (*[]CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf1ClausesInner, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *InputProcessorParamsCreatorCategory) SetClauses(v []CreateMetricRouteRequestProcessorsInnerOneOf2ParamsOneOf1ClausesInner)`

SetClauses sets Clauses field to given value.


### GetDefaultValue

`func (o *InputProcessorParamsCreatorCategory) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *InputProcessorParamsCreatorCategory) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *InputProcessorParamsCreatorCategory) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *InputProcessorParamsCreatorCategory) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetSubtype

`func (o *InputProcessorParamsCreatorCategory) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *InputProcessorParamsCreatorCategory) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *InputProcessorParamsCreatorCategory) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


