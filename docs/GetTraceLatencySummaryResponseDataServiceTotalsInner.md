# GetTraceLatencySummaryResponseDataServiceTotalsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceKey** | **string** | Contributing service key for this aggregate total. References &#x60;services[].key&#x60; in the same response. | 
**DurationNs** | **string** | Total wall-clock contribution of this service across the trace, in nanoseconds. Sum of &#x60;weight * durationNs&#x60; over every range the service contributes to; authoritative when it disagrees with rounded per-range values. Transmitted as a string to avoid losing precision. | 
**Share** | **float32** | Share of total trace duration attributed to this service, from &#x60;0&#x60; through &#x60;1&#x60;. Computed as &#x60;durationNs / totalDurationNs&#x60;; multiply by 100 for a percentage. Shares sum to less than &#x60;1&#x60; for two reasons: time where no span was running is credited to no service, and merging adjacent ranges truncates to the microsecond grid at each step, so the loss compounds with the number of ranges merged. A shortfall on a trace with no gap is that truncation; it is usually a few parts in 10,000, but a long run of very short adjacent ranges can make it large. | 

## Methods

### NewGetTraceLatencySummaryResponseDataServiceTotalsInner

`func NewGetTraceLatencySummaryResponseDataServiceTotalsInner(serviceKey string, durationNs string, share float32, ) *GetTraceLatencySummaryResponseDataServiceTotalsInner`

NewGetTraceLatencySummaryResponseDataServiceTotalsInner instantiates a new GetTraceLatencySummaryResponseDataServiceTotalsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTraceLatencySummaryResponseDataServiceTotalsInnerWithDefaults

`func NewGetTraceLatencySummaryResponseDataServiceTotalsInnerWithDefaults() *GetTraceLatencySummaryResponseDataServiceTotalsInner`

NewGetTraceLatencySummaryResponseDataServiceTotalsInnerWithDefaults instantiates a new GetTraceLatencySummaryResponseDataServiceTotalsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceKey

`func (o *GetTraceLatencySummaryResponseDataServiceTotalsInner) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *GetTraceLatencySummaryResponseDataServiceTotalsInner) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *GetTraceLatencySummaryResponseDataServiceTotalsInner) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.


### GetDurationNs

`func (o *GetTraceLatencySummaryResponseDataServiceTotalsInner) GetDurationNs() string`

GetDurationNs returns the DurationNs field if non-nil, zero value otherwise.

### GetDurationNsOk

`func (o *GetTraceLatencySummaryResponseDataServiceTotalsInner) GetDurationNsOk() (*string, bool)`

GetDurationNsOk returns a tuple with the DurationNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationNs

`func (o *GetTraceLatencySummaryResponseDataServiceTotalsInner) SetDurationNs(v string)`

SetDurationNs sets DurationNs field to given value.


### GetShare

`func (o *GetTraceLatencySummaryResponseDataServiceTotalsInner) GetShare() float32`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *GetTraceLatencySummaryResponseDataServiceTotalsInner) GetShareOk() (*float32, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *GetTraceLatencySummaryResponseDataServiceTotalsInner) SetShare(v float32)`

SetShare sets Share field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


