# GetTraceLatencySummaryResponseDataRangesInnerContributionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceKey** | **string** | Contributing service key for this range. References &#x60;services[].key&#x60; in the same response. | 
**Weight** | **float32** | Share of this range attributed to the service, from &#x60;0&#x60; through &#x60;1&#x60;, computed from active leaf spans in the range. Weights in one range sum to &#x60;1&#x60;, except on a range merged from shorter ones: each pairwise merge truncates to the microsecond grid, so the loss grows with the number of merges, a few parts in 10,000 on a typical trace but well below &#x60;1&#x60; after a long run of very short adjacent ranges. &#x60;serviceTotals&#x60; carries the same loss. | 
**DurationNs** | **string** | Estimated duration attributed to this service in the range, computed as &#x60;weight&#x60; times the range&#39;s own &#x60;durationNs&#x60; and rounded to whole nanoseconds. Use &#x60;serviceTotals&#x60; as authoritative when rounded range values do not add up exactly. | 

## Methods

### NewGetTraceLatencySummaryResponseDataRangesInnerContributionsInner

`func NewGetTraceLatencySummaryResponseDataRangesInnerContributionsInner(serviceKey string, weight float32, durationNs string, ) *GetTraceLatencySummaryResponseDataRangesInnerContributionsInner`

NewGetTraceLatencySummaryResponseDataRangesInnerContributionsInner instantiates a new GetTraceLatencySummaryResponseDataRangesInnerContributionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTraceLatencySummaryResponseDataRangesInnerContributionsInnerWithDefaults

`func NewGetTraceLatencySummaryResponseDataRangesInnerContributionsInnerWithDefaults() *GetTraceLatencySummaryResponseDataRangesInnerContributionsInner`

NewGetTraceLatencySummaryResponseDataRangesInnerContributionsInnerWithDefaults instantiates a new GetTraceLatencySummaryResponseDataRangesInnerContributionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceKey

`func (o *GetTraceLatencySummaryResponseDataRangesInnerContributionsInner) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *GetTraceLatencySummaryResponseDataRangesInnerContributionsInner) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *GetTraceLatencySummaryResponseDataRangesInnerContributionsInner) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.


### GetWeight

`func (o *GetTraceLatencySummaryResponseDataRangesInnerContributionsInner) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GetTraceLatencySummaryResponseDataRangesInnerContributionsInner) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GetTraceLatencySummaryResponseDataRangesInnerContributionsInner) SetWeight(v float32)`

SetWeight sets Weight field to given value.


### GetDurationNs

`func (o *GetTraceLatencySummaryResponseDataRangesInnerContributionsInner) GetDurationNs() string`

GetDurationNs returns the DurationNs field if non-nil, zero value otherwise.

### GetDurationNsOk

`func (o *GetTraceLatencySummaryResponseDataRangesInnerContributionsInner) GetDurationNsOk() (*string, bool)`

GetDurationNsOk returns a tuple with the DurationNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationNs

`func (o *GetTraceLatencySummaryResponseDataRangesInnerContributionsInner) SetDurationNs(v string)`

SetDurationNs sets DurationNs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


