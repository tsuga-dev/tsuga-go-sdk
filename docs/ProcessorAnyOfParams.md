# ProcessorAnyOfParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**[]ProcessorParamsMapperMapAttributesAttributesInner**](ProcessorParamsMapperMapAttributesAttributesInner.md) | Mappings that map individual attributes to new targets | 
**Subtype** | **string** |  | 
**AttributeName** | **string** | Attribute whose value will determine the log timestamp | 

## Methods

### NewProcessorAnyOfParams

`func NewProcessorAnyOfParams(attributes []ProcessorParamsMapperMapAttributesAttributesInner, subtype string, attributeName string, ) *ProcessorAnyOfParams`

NewProcessorAnyOfParams instantiates a new ProcessorAnyOfParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorAnyOfParamsWithDefaults

`func NewProcessorAnyOfParamsWithDefaults() *ProcessorAnyOfParams`

NewProcessorAnyOfParamsWithDefaults instantiates a new ProcessorAnyOfParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *ProcessorAnyOfParams) GetAttributes() []ProcessorParamsMapperMapAttributesAttributesInner`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ProcessorAnyOfParams) GetAttributesOk() (*[]ProcessorParamsMapperMapAttributesAttributesInner, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ProcessorAnyOfParams) SetAttributes(v []ProcessorParamsMapperMapAttributesAttributesInner)`

SetAttributes sets Attributes field to given value.


### GetSubtype

`func (o *ProcessorAnyOfParams) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *ProcessorAnyOfParams) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *ProcessorAnyOfParams) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetAttributeName

`func (o *ProcessorAnyOfParams) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *ProcessorAnyOfParams) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *ProcessorAnyOfParams) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


