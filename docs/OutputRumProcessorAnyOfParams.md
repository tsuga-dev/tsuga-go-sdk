# OutputRumProcessorAnyOfParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**[]OutputRumProcessorAnyOfParamsAttributesInner**](OutputRumProcessorAnyOfParamsAttributesInner.md) | Mappings that map individual attributes to new targets | 
**Subtype** | **string** | Copies or moves one or more source attributes to target attributes. | 

## Methods

### NewOutputRumProcessorAnyOfParams

`func NewOutputRumProcessorAnyOfParams(attributes []OutputRumProcessorAnyOfParamsAttributesInner, subtype string, ) *OutputRumProcessorAnyOfParams`

NewOutputRumProcessorAnyOfParams instantiates a new OutputRumProcessorAnyOfParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputRumProcessorAnyOfParamsWithDefaults

`func NewOutputRumProcessorAnyOfParamsWithDefaults() *OutputRumProcessorAnyOfParams`

NewOutputRumProcessorAnyOfParamsWithDefaults instantiates a new OutputRumProcessorAnyOfParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *OutputRumProcessorAnyOfParams) GetAttributes() []OutputRumProcessorAnyOfParamsAttributesInner`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OutputRumProcessorAnyOfParams) GetAttributesOk() (*[]OutputRumProcessorAnyOfParamsAttributesInner, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OutputRumProcessorAnyOfParams) SetAttributes(v []OutputRumProcessorAnyOfParamsAttributesInner)`

SetAttributes sets Attributes field to given value.


### GetSubtype

`func (o *OutputRumProcessorAnyOfParams) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *OutputRumProcessorAnyOfParams) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *OutputRumProcessorAnyOfParams) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


