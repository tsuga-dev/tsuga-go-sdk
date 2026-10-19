# SummarizeTraceResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spans** | [**[]SearchSpansResponseDataSpansInner**](SearchSpansResponseDataSpansInner.md) | Summarized trace spans, sorted by start time unless &#x60;sortSpans&#x60; is false. Aggregated groups are replaced by synthetic summary spans marked with &#x60;spanAttributes.aggregation.is_summary&#x60;. | 

## Methods

### NewSummarizeTraceResponseData

`func NewSummarizeTraceResponseData(spans []SearchSpansResponseDataSpansInner, ) *SummarizeTraceResponseData`

NewSummarizeTraceResponseData instantiates a new SummarizeTraceResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummarizeTraceResponseDataWithDefaults

`func NewSummarizeTraceResponseDataWithDefaults() *SummarizeTraceResponseData`

NewSummarizeTraceResponseDataWithDefaults instantiates a new SummarizeTraceResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpans

`func (o *SummarizeTraceResponseData) GetSpans() []SearchSpansResponseDataSpansInner`

GetSpans returns the Spans field if non-nil, zero value otherwise.

### GetSpansOk

`func (o *SummarizeTraceResponseData) GetSpansOk() (*[]SearchSpansResponseDataSpansInner, bool)`

GetSpansOk returns a tuple with the Spans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpans

`func (o *SummarizeTraceResponseData) SetSpans(v []SearchSpansResponseDataSpansInner)`

SetSpans sets Spans field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


