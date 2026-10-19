# OutputProcessorAnyOfParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**[]ProcessorParamsMapperMapAttributesAttributesInner**](ProcessorParamsMapperMapAttributesAttributesInner.md) | Mappings that map individual attributes to new targets | 
**Subtype** | **string** | Uses an event-time attribute as the log timestamp. | 
**AttributeName** | **string** | Attribute whose value will determine the log timestamp | 

## Methods

### NewOutputProcessorAnyOfParams

`func NewOutputProcessorAnyOfParams(attributes []ProcessorParamsMapperMapAttributesAttributesInner, subtype string, attributeName string, ) *OutputProcessorAnyOfParams`

NewOutputProcessorAnyOfParams instantiates a new OutputProcessorAnyOfParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputProcessorAnyOfParamsWithDefaults

`func NewOutputProcessorAnyOfParamsWithDefaults() *OutputProcessorAnyOfParams`

NewOutputProcessorAnyOfParamsWithDefaults instantiates a new OutputProcessorAnyOfParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *OutputProcessorAnyOfParams) GetAttributes() []ProcessorParamsMapperMapAttributesAttributesInner`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OutputProcessorAnyOfParams) GetAttributesOk() (*[]ProcessorParamsMapperMapAttributesAttributesInner, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OutputProcessorAnyOfParams) SetAttributes(v []ProcessorParamsMapperMapAttributesAttributesInner)`

SetAttributes sets Attributes field to given value.


### GetSubtype

`func (o *OutputProcessorAnyOfParams) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *OutputProcessorAnyOfParams) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *OutputProcessorAnyOfParams) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetAttributeName

`func (o *OutputProcessorAnyOfParams) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *OutputProcessorAnyOfParams) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *OutputProcessorAnyOfParams) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


