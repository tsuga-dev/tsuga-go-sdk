# AggregateTimeseriesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeRange** | [**TimeRange**](TimeRange.md) |  | 
**Queries** | [**[]AggregationQuery**](AggregationQuery.md) | Aggregations that may be combined together in the same query | 
**GroupBy** | Pointer to [**[]AggregationGroupBy**](AggregationGroupBy.md) | Fields used to group the results | [optional] 
**DataSource** | **string** | Data source being queried for this aggregation | 
**Formula** | Pointer to **string** | Formula referencing query outputs (e.g. q1+q2) to compute derived series | [optional] 
**AggregationWindow** | **string** | Time bucket size for aggregation (e.g. \&quot;10s\&quot;, \&quot;1m\&quot;, \&quot;1h\&quot;) | 

## Methods

### NewAggregateTimeseriesRequest

`func NewAggregateTimeseriesRequest(timeRange TimeRange, queries []AggregationQuery, dataSource string, aggregationWindow string, ) *AggregateTimeseriesRequest`

NewAggregateTimeseriesRequest instantiates a new AggregateTimeseriesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateTimeseriesRequestWithDefaults

`func NewAggregateTimeseriesRequestWithDefaults() *AggregateTimeseriesRequest`

NewAggregateTimeseriesRequestWithDefaults instantiates a new AggregateTimeseriesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeRange

`func (o *AggregateTimeseriesRequest) GetTimeRange() TimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *AggregateTimeseriesRequest) GetTimeRangeOk() (*TimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *AggregateTimeseriesRequest) SetTimeRange(v TimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetQueries

`func (o *AggregateTimeseriesRequest) GetQueries() []AggregationQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *AggregateTimeseriesRequest) GetQueriesOk() (*[]AggregationQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *AggregateTimeseriesRequest) SetQueries(v []AggregationQuery)`

SetQueries sets Queries field to given value.


### GetGroupBy

`func (o *AggregateTimeseriesRequest) GetGroupBy() []AggregationGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *AggregateTimeseriesRequest) GetGroupByOk() (*[]AggregationGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *AggregateTimeseriesRequest) SetGroupBy(v []AggregationGroupBy)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *AggregateTimeseriesRequest) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetDataSource

`func (o *AggregateTimeseriesRequest) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *AggregateTimeseriesRequest) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *AggregateTimeseriesRequest) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.


### GetFormula

`func (o *AggregateTimeseriesRequest) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *AggregateTimeseriesRequest) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *AggregateTimeseriesRequest) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *AggregateTimeseriesRequest) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetAggregationWindow

`func (o *AggregateTimeseriesRequest) GetAggregationWindow() string`

GetAggregationWindow returns the AggregationWindow field if non-nil, zero value otherwise.

### GetAggregationWindowOk

`func (o *AggregateTimeseriesRequest) GetAggregationWindowOk() (*string, bool)`

GetAggregationWindowOk returns a tuple with the AggregationWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationWindow

`func (o *AggregateTimeseriesRequest) SetAggregationWindow(v string)`

SetAggregationWindow sets AggregationWindow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


