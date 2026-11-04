# TimeseriesAggregationResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamps** | **[]int32** | Start of each time bucket as a Unix timestamp in milliseconds, ascending, spaced by &#x60;aggregationWindow&#x60; (&#x60;step&#x60; for PromQL) from the first bucket with data to the last. Every &#x60;values&#x60; array in &#x60;series&#x60; aligns with it index by index. | 
**Series** | [**[]SeriesV2**](SeriesV2.md) | One series per group, in group-by sort order for aggregation queries and in no particular order for PromQL. | 

## Methods

### NewTimeseriesAggregationResponseV2

`func NewTimeseriesAggregationResponseV2(timestamps []int32, series []SeriesV2, ) *TimeseriesAggregationResponseV2`

NewTimeseriesAggregationResponseV2 instantiates a new TimeseriesAggregationResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeseriesAggregationResponseV2WithDefaults

`func NewTimeseriesAggregationResponseV2WithDefaults() *TimeseriesAggregationResponseV2`

NewTimeseriesAggregationResponseV2WithDefaults instantiates a new TimeseriesAggregationResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamps

`func (o *TimeseriesAggregationResponseV2) GetTimestamps() []int32`

GetTimestamps returns the Timestamps field if non-nil, zero value otherwise.

### GetTimestampsOk

`func (o *TimeseriesAggregationResponseV2) GetTimestampsOk() (*[]int32, bool)`

GetTimestampsOk returns a tuple with the Timestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamps

`func (o *TimeseriesAggregationResponseV2) SetTimestamps(v []int32)`

SetTimestamps sets Timestamps field to given value.


### GetSeries

`func (o *TimeseriesAggregationResponseV2) GetSeries() []SeriesV2`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *TimeseriesAggregationResponseV2) GetSeriesOk() (*[]SeriesV2, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *TimeseriesAggregationResponseV2) SetSeries(v []SeriesV2)`

SetSeries sets Series field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


