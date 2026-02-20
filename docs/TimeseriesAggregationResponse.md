# TimeseriesAggregationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Series** | [**[]Series**](Series.md) | Time series data per query and group | 

## Methods

### NewTimeseriesAggregationResponse

`func NewTimeseriesAggregationResponse(series []Series, ) *TimeseriesAggregationResponse`

NewTimeseriesAggregationResponse instantiates a new TimeseriesAggregationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeseriesAggregationResponseWithDefaults

`func NewTimeseriesAggregationResponseWithDefaults() *TimeseriesAggregationResponse`

NewTimeseriesAggregationResponseWithDefaults instantiates a new TimeseriesAggregationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeries

`func (o *TimeseriesAggregationResponse) GetSeries() []Series`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *TimeseriesAggregationResponse) GetSeriesOk() (*[]Series, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *TimeseriesAggregationResponse) SetSeries(v []Series)`

SetSeries sets Series field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


