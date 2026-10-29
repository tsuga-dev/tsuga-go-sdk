# OutputProcessorAnyOf1ParamsOneOf4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAttribute** | **string** | Attribute containing the key/value string segment to parse | 
**TargetAttribute** | **string** | Target map attribute where extracted key/value pairs will be written | 
**KeyValueSplitter** | **string** | Delimiter separating keys from values in the source string | 
**PairsSplitter** | **string** | Delimiter separating each key/value pair in the source string. | 
**AcceptStandaloneKey** | **bool** | Whether keys without explicit values are accepted by the parser. | 
**Subtype** | **string** | Parses delimited key/value text into attributes. | 

## Methods

### NewOutputProcessorAnyOf1ParamsOneOf4

`func NewOutputProcessorAnyOf1ParamsOneOf4(sourceAttribute string, targetAttribute string, keyValueSplitter string, pairsSplitter string, acceptStandaloneKey bool, subtype string, ) *OutputProcessorAnyOf1ParamsOneOf4`

NewOutputProcessorAnyOf1ParamsOneOf4 instantiates a new OutputProcessorAnyOf1ParamsOneOf4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputProcessorAnyOf1ParamsOneOf4WithDefaults

`func NewOutputProcessorAnyOf1ParamsOneOf4WithDefaults() *OutputProcessorAnyOf1ParamsOneOf4`

NewOutputProcessorAnyOf1ParamsOneOf4WithDefaults instantiates a new OutputProcessorAnyOf1ParamsOneOf4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAttribute

`func (o *OutputProcessorAnyOf1ParamsOneOf4) GetSourceAttribute() string`

GetSourceAttribute returns the SourceAttribute field if non-nil, zero value otherwise.

### GetSourceAttributeOk

`func (o *OutputProcessorAnyOf1ParamsOneOf4) GetSourceAttributeOk() (*string, bool)`

GetSourceAttributeOk returns a tuple with the SourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttribute

`func (o *OutputProcessorAnyOf1ParamsOneOf4) SetSourceAttribute(v string)`

SetSourceAttribute sets SourceAttribute field to given value.


### GetTargetAttribute

`func (o *OutputProcessorAnyOf1ParamsOneOf4) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *OutputProcessorAnyOf1ParamsOneOf4) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *OutputProcessorAnyOf1ParamsOneOf4) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetKeyValueSplitter

`func (o *OutputProcessorAnyOf1ParamsOneOf4) GetKeyValueSplitter() string`

GetKeyValueSplitter returns the KeyValueSplitter field if non-nil, zero value otherwise.

### GetKeyValueSplitterOk

`func (o *OutputProcessorAnyOf1ParamsOneOf4) GetKeyValueSplitterOk() (*string, bool)`

GetKeyValueSplitterOk returns a tuple with the KeyValueSplitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyValueSplitter

`func (o *OutputProcessorAnyOf1ParamsOneOf4) SetKeyValueSplitter(v string)`

SetKeyValueSplitter sets KeyValueSplitter field to given value.


### GetPairsSplitter

`func (o *OutputProcessorAnyOf1ParamsOneOf4) GetPairsSplitter() string`

GetPairsSplitter returns the PairsSplitter field if non-nil, zero value otherwise.

### GetPairsSplitterOk

`func (o *OutputProcessorAnyOf1ParamsOneOf4) GetPairsSplitterOk() (*string, bool)`

GetPairsSplitterOk returns a tuple with the PairsSplitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairsSplitter

`func (o *OutputProcessorAnyOf1ParamsOneOf4) SetPairsSplitter(v string)`

SetPairsSplitter sets PairsSplitter field to given value.


### GetAcceptStandaloneKey

`func (o *OutputProcessorAnyOf1ParamsOneOf4) GetAcceptStandaloneKey() bool`

GetAcceptStandaloneKey returns the AcceptStandaloneKey field if non-nil, zero value otherwise.

### GetAcceptStandaloneKeyOk

`func (o *OutputProcessorAnyOf1ParamsOneOf4) GetAcceptStandaloneKeyOk() (*bool, bool)`

GetAcceptStandaloneKeyOk returns a tuple with the AcceptStandaloneKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptStandaloneKey

`func (o *OutputProcessorAnyOf1ParamsOneOf4) SetAcceptStandaloneKey(v bool)`

SetAcceptStandaloneKey sets AcceptStandaloneKey field to given value.


### GetSubtype

`func (o *OutputProcessorAnyOf1ParamsOneOf4) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *OutputProcessorAnyOf1ParamsOneOf4) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *OutputProcessorAnyOf1ParamsOneOf4) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


