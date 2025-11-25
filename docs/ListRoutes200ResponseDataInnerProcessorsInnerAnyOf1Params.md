# ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeName** | **string** | Attribute whose value will be parsed with Grok rules | 
**Rules** | **[]string** | Ordered Grok rules evaluated until one matches | 
**Samples** | Pointer to **[]string** | Example log lines for validation | [optional] 
**Subtype** | **string** |  | 
**SourceAttribute** | **string** | Attribute containing the key/value string segment to parse | 
**TargetAttribute** | **string** | Attribute prefix where extracted key/value pairs will be written | 
**KeyValueSplitter** | **string** | Delimiter separating keys from values in the source string | 
**PairsSplitter** | **string** | Delimiter separating each key/value pair | 
**AcceptStandaloneKey** | **bool** |  | 

## Methods

### NewListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params

`func NewListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params(attributeName string, rules []string, subtype string, sourceAttribute string, targetAttribute string, keyValueSplitter string, pairsSplitter string, acceptStandaloneKey bool, ) *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params`

NewListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params instantiates a new ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRoutes200ResponseDataInnerProcessorsInnerAnyOf1ParamsWithDefaults

`func NewListRoutes200ResponseDataInnerProcessorsInnerAnyOf1ParamsWithDefaults() *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params`

NewListRoutes200ResponseDataInnerProcessorsInnerAnyOf1ParamsWithDefaults instantiates a new ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeName

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.


### GetRules

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) GetRules() []string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) GetRulesOk() (*[]string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) SetRules(v []string)`

SetRules sets Rules field to given value.


### GetSamples

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) GetSamples() []string`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) GetSamplesOk() (*[]string, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) SetSamples(v []string)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetSubtype

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetSourceAttribute

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) GetSourceAttribute() string`

GetSourceAttribute returns the SourceAttribute field if non-nil, zero value otherwise.

### GetSourceAttributeOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) GetSourceAttributeOk() (*string, bool)`

GetSourceAttributeOk returns a tuple with the SourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttribute

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) SetSourceAttribute(v string)`

SetSourceAttribute sets SourceAttribute field to given value.


### GetTargetAttribute

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetKeyValueSplitter

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) GetKeyValueSplitter() string`

GetKeyValueSplitter returns the KeyValueSplitter field if non-nil, zero value otherwise.

### GetKeyValueSplitterOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) GetKeyValueSplitterOk() (*string, bool)`

GetKeyValueSplitterOk returns a tuple with the KeyValueSplitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyValueSplitter

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) SetKeyValueSplitter(v string)`

SetKeyValueSplitter sets KeyValueSplitter field to given value.


### GetPairsSplitter

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) GetPairsSplitter() string`

GetPairsSplitter returns the PairsSplitter field if non-nil, zero value otherwise.

### GetPairsSplitterOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) GetPairsSplitterOk() (*string, bool)`

GetPairsSplitterOk returns a tuple with the PairsSplitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairsSplitter

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) SetPairsSplitter(v string)`

SetPairsSplitter sets PairsSplitter field to given value.


### GetAcceptStandaloneKey

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) GetAcceptStandaloneKey() bool`

GetAcceptStandaloneKey returns the AcceptStandaloneKey field if non-nil, zero value otherwise.

### GetAcceptStandaloneKeyOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) GetAcceptStandaloneKeyOk() (*bool, bool)`

GetAcceptStandaloneKeyOk returns a tuple with the AcceptStandaloneKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptStandaloneKey

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf1Params) SetAcceptStandaloneKey(v bool)`

SetAcceptStandaloneKey sets AcceptStandaloneKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


