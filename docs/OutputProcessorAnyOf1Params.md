# OutputProcessorAnyOf1Params

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeName** | **string** | Attribute whose value will be parsed with Grok rules | 
**Rules** | **[]string** | Ordered Grok rules evaluated until one matches. | 
**Samples** | Pointer to **[]string** | Example log lines for validating the Grok rules. Samples are not sent to the ingestion engine. | [optional] 
**Subtype** | **string** | Parses delimited key/value text into attributes. | 
**SourceAttribute** | **string** | Attribute containing the key/value string segment to parse | 
**TargetAttribute** | **string** | Target map attribute where extracted key/value pairs will be written | 
**RedactSource** | **bool** | Whether to replace a public IP address in the source attribute with &#x60;[REDACTED_IP]&#x60;. Private IP addresses and non-IP values are left unchanged. | 
**KeyValueSplitter** | **string** | Delimiter separating keys from values in the source string | 
**PairsSplitter** | **string** | Delimiter separating each key/value pair in the source string. | 
**AcceptStandaloneKey** | **bool** | Whether keys without explicit values are accepted by the parser. | 

## Methods

### NewOutputProcessorAnyOf1Params

`func NewOutputProcessorAnyOf1Params(attributeName string, rules []string, subtype string, sourceAttribute string, targetAttribute string, redactSource bool, keyValueSplitter string, pairsSplitter string, acceptStandaloneKey bool, ) *OutputProcessorAnyOf1Params`

NewOutputProcessorAnyOf1Params instantiates a new OutputProcessorAnyOf1Params object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputProcessorAnyOf1ParamsWithDefaults

`func NewOutputProcessorAnyOf1ParamsWithDefaults() *OutputProcessorAnyOf1Params`

NewOutputProcessorAnyOf1ParamsWithDefaults instantiates a new OutputProcessorAnyOf1Params object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeName

`func (o *OutputProcessorAnyOf1Params) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *OutputProcessorAnyOf1Params) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *OutputProcessorAnyOf1Params) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.


### GetRules

`func (o *OutputProcessorAnyOf1Params) GetRules() []string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *OutputProcessorAnyOf1Params) GetRulesOk() (*[]string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *OutputProcessorAnyOf1Params) SetRules(v []string)`

SetRules sets Rules field to given value.


### GetSamples

`func (o *OutputProcessorAnyOf1Params) GetSamples() []string`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *OutputProcessorAnyOf1Params) GetSamplesOk() (*[]string, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *OutputProcessorAnyOf1Params) SetSamples(v []string)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *OutputProcessorAnyOf1Params) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetSubtype

`func (o *OutputProcessorAnyOf1Params) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *OutputProcessorAnyOf1Params) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *OutputProcessorAnyOf1Params) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetSourceAttribute

`func (o *OutputProcessorAnyOf1Params) GetSourceAttribute() string`

GetSourceAttribute returns the SourceAttribute field if non-nil, zero value otherwise.

### GetSourceAttributeOk

`func (o *OutputProcessorAnyOf1Params) GetSourceAttributeOk() (*string, bool)`

GetSourceAttributeOk returns a tuple with the SourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttribute

`func (o *OutputProcessorAnyOf1Params) SetSourceAttribute(v string)`

SetSourceAttribute sets SourceAttribute field to given value.


### GetTargetAttribute

`func (o *OutputProcessorAnyOf1Params) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *OutputProcessorAnyOf1Params) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *OutputProcessorAnyOf1Params) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetRedactSource

`func (o *OutputProcessorAnyOf1Params) GetRedactSource() bool`

GetRedactSource returns the RedactSource field if non-nil, zero value otherwise.

### GetRedactSourceOk

`func (o *OutputProcessorAnyOf1Params) GetRedactSourceOk() (*bool, bool)`

GetRedactSourceOk returns a tuple with the RedactSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactSource

`func (o *OutputProcessorAnyOf1Params) SetRedactSource(v bool)`

SetRedactSource sets RedactSource field to given value.


### GetKeyValueSplitter

`func (o *OutputProcessorAnyOf1Params) GetKeyValueSplitter() string`

GetKeyValueSplitter returns the KeyValueSplitter field if non-nil, zero value otherwise.

### GetKeyValueSplitterOk

`func (o *OutputProcessorAnyOf1Params) GetKeyValueSplitterOk() (*string, bool)`

GetKeyValueSplitterOk returns a tuple with the KeyValueSplitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyValueSplitter

`func (o *OutputProcessorAnyOf1Params) SetKeyValueSplitter(v string)`

SetKeyValueSplitter sets KeyValueSplitter field to given value.


### GetPairsSplitter

`func (o *OutputProcessorAnyOf1Params) GetPairsSplitter() string`

GetPairsSplitter returns the PairsSplitter field if non-nil, zero value otherwise.

### GetPairsSplitterOk

`func (o *OutputProcessorAnyOf1Params) GetPairsSplitterOk() (*string, bool)`

GetPairsSplitterOk returns a tuple with the PairsSplitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairsSplitter

`func (o *OutputProcessorAnyOf1Params) SetPairsSplitter(v string)`

SetPairsSplitter sets PairsSplitter field to given value.


### GetAcceptStandaloneKey

`func (o *OutputProcessorAnyOf1Params) GetAcceptStandaloneKey() bool`

GetAcceptStandaloneKey returns the AcceptStandaloneKey field if non-nil, zero value otherwise.

### GetAcceptStandaloneKeyOk

`func (o *OutputProcessorAnyOf1Params) GetAcceptStandaloneKeyOk() (*bool, bool)`

GetAcceptStandaloneKeyOk returns a tuple with the AcceptStandaloneKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptStandaloneKey

`func (o *OutputProcessorAnyOf1Params) SetAcceptStandaloneKey(v bool)`

SetAcceptStandaloneKey sets AcceptStandaloneKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


