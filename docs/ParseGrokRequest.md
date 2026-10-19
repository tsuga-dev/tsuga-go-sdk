# ParseGrokRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | **[]string** | Ordered Grok parsing rules to preview. For each sample, Tsuga tries the rules in order and returns the first successful match, matching route Grok processor semantics. | 
**Samples** | **[]string** | Sample log lines to run against the rules. Results are returned in the same order as this array. | 

## Methods

### NewParseGrokRequest

`func NewParseGrokRequest(rules []string, samples []string, ) *ParseGrokRequest`

NewParseGrokRequest instantiates a new ParseGrokRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParseGrokRequestWithDefaults

`func NewParseGrokRequestWithDefaults() *ParseGrokRequest`

NewParseGrokRequestWithDefaults instantiates a new ParseGrokRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *ParseGrokRequest) GetRules() []string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ParseGrokRequest) GetRulesOk() (*[]string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ParseGrokRequest) SetRules(v []string)`

SetRules sets Rules field to given value.


### GetSamples

`func (o *ParseGrokRequest) GetSamples() []string`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *ParseGrokRequest) GetSamplesOk() (*[]string, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *ParseGrokRequest) SetSamples(v []string)`

SetSamples sets Samples field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


