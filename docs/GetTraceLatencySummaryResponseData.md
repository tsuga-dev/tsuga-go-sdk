# GetTraceLatencySummaryResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TraceId** | **string** | Trace ID summarized by the response. Echoes the required &#x60;traceId&#x60; query parameter. | 
**StartTimeNs** | **string** | Earliest &#x60;start_time&#x60; across fetched spans, in nanoseconds since Unix epoch. If the request window does not cover the full trace, this reflects the fetched subset, not necessarily the trace&#39;s true start. Transmitted as a string to avoid losing precision. | 
**EndTimeNs** | **string** | Latest &#x60;end_time&#x60; across fetched spans, in nanoseconds since Unix epoch. If the request window does not cover the full trace, this reflects the fetched subset, not necessarily the trace&#39;s true end. Transmitted as a string to avoid losing precision. | 
**TotalDurationNs** | **string** | Total fetched trace duration in nanoseconds, computed as &#x60;endTimeNs - startTimeNs&#x60;. Transmitted as a string to avoid losing precision. | 
**Truncated** | **bool** | Whether the underlying span fetch hit its result cap. When true, the latency summary may be incomplete because not every span in the trace was available. | 
**Services** | [**[]GetTraceLatencySummaryResponseDataServicesInner**](GetTraceLatencySummaryResponseDataServicesInner.md) | Unique services that contributed leaf-span time to the trace, enriched with catalog IDs when Tsuga can resolve them. Referenced by &#x60;serviceKey&#x60; from &#x60;ranges&#x60; and &#x60;serviceTotals&#x60;. | 
**Ranges** | [**[]GetTraceLatencySummaryResponseDataRangesInner**](GetTraceLatencySummaryResponseDataRangesInner.md) | Sequential trace timeline split into ranges wherever the active leaf-span set changes. Empty when the request sets &#x60;includeRanges&#x60; to false. | 
**ServiceTotals** | [**[]GetTraceLatencySummaryResponseDataServiceTotalsInner**](GetTraceLatencySummaryResponseDataServiceTotalsInner.md) | Aggregated wall-clock contribution per service across the fetched trace. Use this for service-level latency ownership without processing the per-range timeline. | 

## Methods

### NewGetTraceLatencySummaryResponseData

`func NewGetTraceLatencySummaryResponseData(traceId string, startTimeNs string, endTimeNs string, totalDurationNs string, truncated bool, services []GetTraceLatencySummaryResponseDataServicesInner, ranges []GetTraceLatencySummaryResponseDataRangesInner, serviceTotals []GetTraceLatencySummaryResponseDataServiceTotalsInner, ) *GetTraceLatencySummaryResponseData`

NewGetTraceLatencySummaryResponseData instantiates a new GetTraceLatencySummaryResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTraceLatencySummaryResponseDataWithDefaults

`func NewGetTraceLatencySummaryResponseDataWithDefaults() *GetTraceLatencySummaryResponseData`

NewGetTraceLatencySummaryResponseDataWithDefaults instantiates a new GetTraceLatencySummaryResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceId

`func (o *GetTraceLatencySummaryResponseData) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *GetTraceLatencySummaryResponseData) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *GetTraceLatencySummaryResponseData) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetStartTimeNs

`func (o *GetTraceLatencySummaryResponseData) GetStartTimeNs() string`

GetStartTimeNs returns the StartTimeNs field if non-nil, zero value otherwise.

### GetStartTimeNsOk

`func (o *GetTraceLatencySummaryResponseData) GetStartTimeNsOk() (*string, bool)`

GetStartTimeNsOk returns a tuple with the StartTimeNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeNs

`func (o *GetTraceLatencySummaryResponseData) SetStartTimeNs(v string)`

SetStartTimeNs sets StartTimeNs field to given value.


### GetEndTimeNs

`func (o *GetTraceLatencySummaryResponseData) GetEndTimeNs() string`

GetEndTimeNs returns the EndTimeNs field if non-nil, zero value otherwise.

### GetEndTimeNsOk

`func (o *GetTraceLatencySummaryResponseData) GetEndTimeNsOk() (*string, bool)`

GetEndTimeNsOk returns a tuple with the EndTimeNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeNs

`func (o *GetTraceLatencySummaryResponseData) SetEndTimeNs(v string)`

SetEndTimeNs sets EndTimeNs field to given value.


### GetTotalDurationNs

`func (o *GetTraceLatencySummaryResponseData) GetTotalDurationNs() string`

GetTotalDurationNs returns the TotalDurationNs field if non-nil, zero value otherwise.

### GetTotalDurationNsOk

`func (o *GetTraceLatencySummaryResponseData) GetTotalDurationNsOk() (*string, bool)`

GetTotalDurationNsOk returns a tuple with the TotalDurationNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationNs

`func (o *GetTraceLatencySummaryResponseData) SetTotalDurationNs(v string)`

SetTotalDurationNs sets TotalDurationNs field to given value.


### GetTruncated

`func (o *GetTraceLatencySummaryResponseData) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *GetTraceLatencySummaryResponseData) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *GetTraceLatencySummaryResponseData) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.


### GetServices

`func (o *GetTraceLatencySummaryResponseData) GetServices() []GetTraceLatencySummaryResponseDataServicesInner`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *GetTraceLatencySummaryResponseData) GetServicesOk() (*[]GetTraceLatencySummaryResponseDataServicesInner, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *GetTraceLatencySummaryResponseData) SetServices(v []GetTraceLatencySummaryResponseDataServicesInner)`

SetServices sets Services field to given value.


### GetRanges

`func (o *GetTraceLatencySummaryResponseData) GetRanges() []GetTraceLatencySummaryResponseDataRangesInner`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *GetTraceLatencySummaryResponseData) GetRangesOk() (*[]GetTraceLatencySummaryResponseDataRangesInner, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *GetTraceLatencySummaryResponseData) SetRanges(v []GetTraceLatencySummaryResponseDataRangesInner)`

SetRanges sets Ranges field to given value.


### GetServiceTotals

`func (o *GetTraceLatencySummaryResponseData) GetServiceTotals() []GetTraceLatencySummaryResponseDataServiceTotalsInner`

GetServiceTotals returns the ServiceTotals field if non-nil, zero value otherwise.

### GetServiceTotalsOk

`func (o *GetTraceLatencySummaryResponseData) GetServiceTotalsOk() (*[]GetTraceLatencySummaryResponseDataServiceTotalsInner, bool)`

GetServiceTotalsOk returns a tuple with the ServiceTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTotals

`func (o *GetTraceLatencySummaryResponseData) SetServiceTotals(v []GetTraceLatencySummaryResponseDataServiceTotalsInner)`

SetServiceTotals sets ServiceTotals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


