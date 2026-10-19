# ProcessorParamsCreatorCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAttribute** | **string** | Attribute that will receive the category value | 
**Clauses** | [**[]ProcessorParamsCreatorCategoryClausesInner**](ProcessorParamsCreatorCategoryClausesInner.md) | Conditions evaluated in order to determine the category value. The first matching clause wins. | 
**DefaultValue** | Pointer to **string** | String category value used when no clause matches. If omitted and no clause matches, no category value is written. | [optional] 
**Subtype** | **string** | Assigns the first matching category value from ordered query clauses, with an optional default. | 

## Methods

### NewProcessorParamsCreatorCategory

`func NewProcessorParamsCreatorCategory(targetAttribute string, clauses []ProcessorParamsCreatorCategoryClausesInner, subtype string, ) *ProcessorParamsCreatorCategory`

NewProcessorParamsCreatorCategory instantiates a new ProcessorParamsCreatorCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorParamsCreatorCategoryWithDefaults

`func NewProcessorParamsCreatorCategoryWithDefaults() *ProcessorParamsCreatorCategory`

NewProcessorParamsCreatorCategoryWithDefaults instantiates a new ProcessorParamsCreatorCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAttribute

`func (o *ProcessorParamsCreatorCategory) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *ProcessorParamsCreatorCategory) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *ProcessorParamsCreatorCategory) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetClauses

`func (o *ProcessorParamsCreatorCategory) GetClauses() []ProcessorParamsCreatorCategoryClausesInner`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *ProcessorParamsCreatorCategory) GetClausesOk() (*[]ProcessorParamsCreatorCategoryClausesInner, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *ProcessorParamsCreatorCategory) SetClauses(v []ProcessorParamsCreatorCategoryClausesInner)`

SetClauses sets Clauses field to given value.


### GetDefaultValue

`func (o *ProcessorParamsCreatorCategory) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ProcessorParamsCreatorCategory) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ProcessorParamsCreatorCategory) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ProcessorParamsCreatorCategory) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetSubtype

`func (o *ProcessorParamsCreatorCategory) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *ProcessorParamsCreatorCategory) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *ProcessorParamsCreatorCategory) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


