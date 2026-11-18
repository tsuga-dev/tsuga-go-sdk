# OutputSpansProcessorAnyOfParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**[]OutputSpansProcessorAnyOfParamsAttributesInner**](OutputSpansProcessorAnyOfParamsAttributesInner.md) | Mappings that map individual attributes to new targets | 
**Subtype** | **string** | Copies or moves one or more source attributes to target attributes. | 

## Methods

### NewOutputSpansProcessorAnyOfParams

`func NewOutputSpansProcessorAnyOfParams(attributes []OutputSpansProcessorAnyOfParamsAttributesInner, subtype string, ) *OutputSpansProcessorAnyOfParams`

NewOutputSpansProcessorAnyOfParams instantiates a new OutputSpansProcessorAnyOfParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputSpansProcessorAnyOfParamsWithDefaults

`func NewOutputSpansProcessorAnyOfParamsWithDefaults() *OutputSpansProcessorAnyOfParams`

NewOutputSpansProcessorAnyOfParamsWithDefaults instantiates a new OutputSpansProcessorAnyOfParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *OutputSpansProcessorAnyOfParams) GetAttributes() []OutputSpansProcessorAnyOfParamsAttributesInner`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OutputSpansProcessorAnyOfParams) GetAttributesOk() (*[]OutputSpansProcessorAnyOfParamsAttributesInner, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OutputSpansProcessorAnyOfParams) SetAttributes(v []OutputSpansProcessorAnyOfParamsAttributesInner)`

SetAttributes sets Attributes field to given value.


### GetSubtype

`func (o *OutputSpansProcessorAnyOfParams) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *OutputSpansProcessorAnyOfParams) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *OutputSpansProcessorAnyOfParams) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


