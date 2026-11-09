# SummarizeTraceResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | **map[string]map[string]interface{}** | Distinct resource contexts of the trace, keyed by the &#x60;resource_id&#x60; each span carries. Emitted once per resource instead of on every span. | 
**Spans** | [**[]SearchSpansResponseDataSpansInner**](SearchSpansResponseDataSpansInner.md) | Summarized trace spans, sorted by start time unless &#x60;sortSpans&#x60; is false. Aggregated groups are replaced by synthetic summary spans marked with &#x60;span_attributes.aggregation.is_summary&#x60;. | 

## Methods

### NewSummarizeTraceResponseData

`func NewSummarizeTraceResponseData(resources map[string]map[string]interface{}, spans []SearchSpansResponseDataSpansInner, ) *SummarizeTraceResponseData`

NewSummarizeTraceResponseData instantiates a new SummarizeTraceResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummarizeTraceResponseDataWithDefaults

`func NewSummarizeTraceResponseDataWithDefaults() *SummarizeTraceResponseData`

NewSummarizeTraceResponseDataWithDefaults instantiates a new SummarizeTraceResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *SummarizeTraceResponseData) GetResources() map[string]map[string]interface{}`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *SummarizeTraceResponseData) GetResourcesOk() (*map[string]map[string]interface{}, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *SummarizeTraceResponseData) SetResources(v map[string]map[string]interface{})`

SetResources sets Resources field to given value.


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


