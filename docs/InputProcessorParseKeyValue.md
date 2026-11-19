# InputProcessorParseKeyValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAttribute** | **string** | Attribute containing the key/value string segment to parse | 
**TargetAttribute** | **string** | Target map attribute where extracted key/value pairs will be written | 
**KeyValueSplitter** | **string** | Delimiter separating keys from values in the source string | 
**PairsSplitter** | **string** | Delimiter separating each key/value pair in the source string. | 
**AcceptStandaloneKey** | **bool** | Set to true to accept keys without explicit values. | 
**Subtype** | **string** | Parses delimited key/value text into attributes. | 

## Methods

### NewInputProcessorParseKeyValue

`func NewInputProcessorParseKeyValue(sourceAttribute string, targetAttribute string, keyValueSplitter string, pairsSplitter string, acceptStandaloneKey bool, subtype string, ) *InputProcessorParseKeyValue`

NewInputProcessorParseKeyValue instantiates a new InputProcessorParseKeyValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputProcessorParseKeyValueWithDefaults

`func NewInputProcessorParseKeyValueWithDefaults() *InputProcessorParseKeyValue`

NewInputProcessorParseKeyValueWithDefaults instantiates a new InputProcessorParseKeyValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAttribute

`func (o *InputProcessorParseKeyValue) GetSourceAttribute() string`

GetSourceAttribute returns the SourceAttribute field if non-nil, zero value otherwise.

### GetSourceAttributeOk

`func (o *InputProcessorParseKeyValue) GetSourceAttributeOk() (*string, bool)`

GetSourceAttributeOk returns a tuple with the SourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttribute

`func (o *InputProcessorParseKeyValue) SetSourceAttribute(v string)`

SetSourceAttribute sets SourceAttribute field to given value.


### GetTargetAttribute

`func (o *InputProcessorParseKeyValue) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *InputProcessorParseKeyValue) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *InputProcessorParseKeyValue) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetKeyValueSplitter

`func (o *InputProcessorParseKeyValue) GetKeyValueSplitter() string`

GetKeyValueSplitter returns the KeyValueSplitter field if non-nil, zero value otherwise.

### GetKeyValueSplitterOk

`func (o *InputProcessorParseKeyValue) GetKeyValueSplitterOk() (*string, bool)`

GetKeyValueSplitterOk returns a tuple with the KeyValueSplitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyValueSplitter

`func (o *InputProcessorParseKeyValue) SetKeyValueSplitter(v string)`

SetKeyValueSplitter sets KeyValueSplitter field to given value.


### GetPairsSplitter

`func (o *InputProcessorParseKeyValue) GetPairsSplitter() string`

GetPairsSplitter returns the PairsSplitter field if non-nil, zero value otherwise.

### GetPairsSplitterOk

`func (o *InputProcessorParseKeyValue) GetPairsSplitterOk() (*string, bool)`

GetPairsSplitterOk returns a tuple with the PairsSplitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairsSplitter

`func (o *InputProcessorParseKeyValue) SetPairsSplitter(v string)`

SetPairsSplitter sets PairsSplitter field to given value.


### GetAcceptStandaloneKey

`func (o *InputProcessorParseKeyValue) GetAcceptStandaloneKey() bool`

GetAcceptStandaloneKey returns the AcceptStandaloneKey field if non-nil, zero value otherwise.

### GetAcceptStandaloneKeyOk

`func (o *InputProcessorParseKeyValue) GetAcceptStandaloneKeyOk() (*bool, bool)`

GetAcceptStandaloneKeyOk returns a tuple with the AcceptStandaloneKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptStandaloneKey

`func (o *InputProcessorParseKeyValue) SetAcceptStandaloneKey(v bool)`

SetAcceptStandaloneKey sets AcceptStandaloneKey field to given value.


### GetSubtype

`func (o *InputProcessorParseKeyValue) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *InputProcessorParseKeyValue) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *InputProcessorParseKeyValue) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


