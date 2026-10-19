# ParseGrokResponseDataResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sample** | **string** | Echo of the submitted log line, so results can be matched to the samples in the request. | 
**Matched** | **bool** | Whether any submitted rule matched this sample. &#x60;true&#x60; means &#x60;extracted&#x60; contains the captured fields; &#x60;false&#x60; means no rule matched. | 
**Extracted** | **map[string]interface{}** | Fields extracted from the first matching Grok rule when &#x60;matched&#x60; is &#x60;true&#x60;; dotted capture names are expanded into nested objects. &#x60;null&#x60; when no rule matched. | 

## Methods

### NewParseGrokResponseDataResultsInner

`func NewParseGrokResponseDataResultsInner(sample string, matched bool, extracted map[string]*interface{}, ) *ParseGrokResponseDataResultsInner`

NewParseGrokResponseDataResultsInner instantiates a new ParseGrokResponseDataResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParseGrokResponseDataResultsInnerWithDefaults

`func NewParseGrokResponseDataResultsInnerWithDefaults() *ParseGrokResponseDataResultsInner`

NewParseGrokResponseDataResultsInnerWithDefaults instantiates a new ParseGrokResponseDataResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSample

`func (o *ParseGrokResponseDataResultsInner) GetSample() string`

GetSample returns the Sample field if non-nil, zero value otherwise.

### GetSampleOk

`func (o *ParseGrokResponseDataResultsInner) GetSampleOk() (*string, bool)`

GetSampleOk returns a tuple with the Sample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSample

`func (o *ParseGrokResponseDataResultsInner) SetSample(v string)`

SetSample sets Sample field to given value.


### GetMatched

`func (o *ParseGrokResponseDataResultsInner) GetMatched() bool`

GetMatched returns the Matched field if non-nil, zero value otherwise.

### GetMatchedOk

`func (o *ParseGrokResponseDataResultsInner) GetMatchedOk() (*bool, bool)`

GetMatchedOk returns a tuple with the Matched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatched

`func (o *ParseGrokResponseDataResultsInner) SetMatched(v bool)`

SetMatched sets Matched field to given value.


### GetExtracted

`func (o *ParseGrokResponseDataResultsInner) GetExtracted() map[string]*interface{}`

GetExtracted returns the Extracted field if non-nil, zero value otherwise.

### GetExtractedOk

`func (o *ParseGrokResponseDataResultsInner) GetExtractedOk() (*map[string]*interface{}, bool)`

GetExtractedOk returns a tuple with the Extracted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtracted

`func (o *ParseGrokResponseDataResultsInner) SetExtracted(v map[string]*interface{})`

SetExtracted sets Extracted field to given value.


### SetExtractedNil

`func (o *ParseGrokResponseDataResultsInner) SetExtractedNil(b bool)`

 SetExtractedNil sets the value for Extracted to be an explicit nil

### UnsetExtracted
`func (o *ParseGrokResponseDataResultsInner) UnsetExtracted()`

UnsetExtracted ensures that no value is present for Extracted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


