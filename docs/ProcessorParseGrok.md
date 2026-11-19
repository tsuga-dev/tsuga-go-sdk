# ProcessorParseGrok

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeName** | **string** | Attribute whose value will be parsed with Grok rules | 
**Rules** | **[]string** | Ordered Grok rules evaluated until one matches. | 
**Samples** | Pointer to **[]string** | Sample attribute values for validating the Grok rules. Samples are not sent to the ingestion engine. | [optional] 
**Subtype** | **string** | Extracts fields from unstructured text using ordered Grok rules. | 

## Methods

### NewProcessorParseGrok

`func NewProcessorParseGrok(attributeName string, rules []string, subtype string, ) *ProcessorParseGrok`

NewProcessorParseGrok instantiates a new ProcessorParseGrok object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorParseGrokWithDefaults

`func NewProcessorParseGrokWithDefaults() *ProcessorParseGrok`

NewProcessorParseGrokWithDefaults instantiates a new ProcessorParseGrok object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeName

`func (o *ProcessorParseGrok) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *ProcessorParseGrok) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *ProcessorParseGrok) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.


### GetRules

`func (o *ProcessorParseGrok) GetRules() []string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ProcessorParseGrok) GetRulesOk() (*[]string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ProcessorParseGrok) SetRules(v []string)`

SetRules sets Rules field to given value.


### GetSamples

`func (o *ProcessorParseGrok) GetSamples() []string`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *ProcessorParseGrok) GetSamplesOk() (*[]string, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *ProcessorParseGrok) SetSamples(v []string)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *ProcessorParseGrok) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetSubtype

`func (o *ProcessorParseGrok) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *ProcessorParseGrok) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *ProcessorParseGrok) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


