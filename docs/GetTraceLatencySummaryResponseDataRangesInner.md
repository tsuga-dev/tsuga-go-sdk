# GetTraceLatencySummaryResponseDataRangesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromNs** | **string** | Range start time in nanoseconds since Unix epoch. Tsuga sets it from the trace timeline split where active leaf spans change. Transmitted as a string to avoid losing precision. | 
**ToNs** | **string** | Range end time in nanoseconds since Unix epoch. Tsuga sets it from the trace timeline split where active leaf spans change. Transmitted as a string to avoid losing precision. | 
**DurationNs** | **string** | Range duration in nanoseconds, computed as &#x60;toNs - fromNs&#x60;. Transmitted as a string to avoid losing precision. | 
**Contributions** | [**[]GetTraceLatencySummaryResponseDataRangesInnerContributionsInner**](GetTraceLatencySummaryResponseDataRangesInnerContributionsInner.md) | Services contributing active leaf-span time to this range, with weighted shares for the range duration. | 
**SpanIds** | **[]string** | Leaf span IDs that produced this range. Tsuga returns IDs from the leaf-only attribution set, including merged sub-ranges after short-range collapsing. | 

## Methods

### NewGetTraceLatencySummaryResponseDataRangesInner

`func NewGetTraceLatencySummaryResponseDataRangesInner(fromNs string, toNs string, durationNs string, contributions []GetTraceLatencySummaryResponseDataRangesInnerContributionsInner, spanIds []string, ) *GetTraceLatencySummaryResponseDataRangesInner`

NewGetTraceLatencySummaryResponseDataRangesInner instantiates a new GetTraceLatencySummaryResponseDataRangesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTraceLatencySummaryResponseDataRangesInnerWithDefaults

`func NewGetTraceLatencySummaryResponseDataRangesInnerWithDefaults() *GetTraceLatencySummaryResponseDataRangesInner`

NewGetTraceLatencySummaryResponseDataRangesInnerWithDefaults instantiates a new GetTraceLatencySummaryResponseDataRangesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromNs

`func (o *GetTraceLatencySummaryResponseDataRangesInner) GetFromNs() string`

GetFromNs returns the FromNs field if non-nil, zero value otherwise.

### GetFromNsOk

`func (o *GetTraceLatencySummaryResponseDataRangesInner) GetFromNsOk() (*string, bool)`

GetFromNsOk returns a tuple with the FromNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromNs

`func (o *GetTraceLatencySummaryResponseDataRangesInner) SetFromNs(v string)`

SetFromNs sets FromNs field to given value.


### GetToNs

`func (o *GetTraceLatencySummaryResponseDataRangesInner) GetToNs() string`

GetToNs returns the ToNs field if non-nil, zero value otherwise.

### GetToNsOk

`func (o *GetTraceLatencySummaryResponseDataRangesInner) GetToNsOk() (*string, bool)`

GetToNsOk returns a tuple with the ToNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToNs

`func (o *GetTraceLatencySummaryResponseDataRangesInner) SetToNs(v string)`

SetToNs sets ToNs field to given value.


### GetDurationNs

`func (o *GetTraceLatencySummaryResponseDataRangesInner) GetDurationNs() string`

GetDurationNs returns the DurationNs field if non-nil, zero value otherwise.

### GetDurationNsOk

`func (o *GetTraceLatencySummaryResponseDataRangesInner) GetDurationNsOk() (*string, bool)`

GetDurationNsOk returns a tuple with the DurationNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationNs

`func (o *GetTraceLatencySummaryResponseDataRangesInner) SetDurationNs(v string)`

SetDurationNs sets DurationNs field to given value.


### GetContributions

`func (o *GetTraceLatencySummaryResponseDataRangesInner) GetContributions() []GetTraceLatencySummaryResponseDataRangesInnerContributionsInner`

GetContributions returns the Contributions field if non-nil, zero value otherwise.

### GetContributionsOk

`func (o *GetTraceLatencySummaryResponseDataRangesInner) GetContributionsOk() (*[]GetTraceLatencySummaryResponseDataRangesInnerContributionsInner, bool)`

GetContributionsOk returns a tuple with the Contributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributions

`func (o *GetTraceLatencySummaryResponseDataRangesInner) SetContributions(v []GetTraceLatencySummaryResponseDataRangesInnerContributionsInner)`

SetContributions sets Contributions field to given value.


### GetSpanIds

`func (o *GetTraceLatencySummaryResponseDataRangesInner) GetSpanIds() []string`

GetSpanIds returns the SpanIds field if non-nil, zero value otherwise.

### GetSpanIdsOk

`func (o *GetTraceLatencySummaryResponseDataRangesInner) GetSpanIdsOk() (*[]string, bool)`

GetSpanIdsOk returns a tuple with the SpanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanIds

`func (o *GetTraceLatencySummaryResponseDataRangesInner) SetSpanIds(v []string)`

SetSpanIds sets SpanIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


