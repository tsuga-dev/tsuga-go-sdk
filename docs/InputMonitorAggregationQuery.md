# InputMonitorAggregationQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregate** | [**InputAggregate**](InputAggregate.md) |  | 
**Functions** | Pointer to [**[]InputFunction**](InputFunction.md) | Post-processing functions applied to aggregation results. Defaults to an empty array when omitted. | [optional] 
**Fill** | Pointer to [**MonitorAggregationQueryFill**](MonitorAggregationQueryFill.md) |  | [optional] 
**TimeAggregate** | Pointer to **string** | Per-series rollup applied within each time bucket before the cross-series aggregate. Use it on metric queries when &#x60;aggregate.type&#x60; is &#x60;sum&#x60; and no &#x60;rate&#x60;, &#x60;increase&#x60;, &#x60;last&#x60;, or &#x60;rolling&#x60; function is present. When omitted, Tsuga derives the rollup from the metric type. | [optional] 
**MetricKind** | Pointer to **string** | Variant of the metric in &#x60;aggregate.field&#x60; to query, for metric names reported under several kinds (for instance both cumulative and delta). When omitted, Tsuga picks one variant by type and temporality precedence. | [optional] 
**Filter** | **string** | Tsuga query filter applied to this alerting or SLO aggregation query. Required by the API; use the query language for the selected data source. | 

## Methods

### NewInputMonitorAggregationQuery

`func NewInputMonitorAggregationQuery(aggregate InputAggregate, filter string, ) *InputMonitorAggregationQuery`

NewInputMonitorAggregationQuery instantiates a new InputMonitorAggregationQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputMonitorAggregationQueryWithDefaults

`func NewInputMonitorAggregationQueryWithDefaults() *InputMonitorAggregationQuery`

NewInputMonitorAggregationQueryWithDefaults instantiates a new InputMonitorAggregationQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregate

`func (o *InputMonitorAggregationQuery) GetAggregate() InputAggregate`

GetAggregate returns the Aggregate field if non-nil, zero value otherwise.

### GetAggregateOk

`func (o *InputMonitorAggregationQuery) GetAggregateOk() (*InputAggregate, bool)`

GetAggregateOk returns a tuple with the Aggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregate

`func (o *InputMonitorAggregationQuery) SetAggregate(v InputAggregate)`

SetAggregate sets Aggregate field to given value.


### GetFunctions

`func (o *InputMonitorAggregationQuery) GetFunctions() []InputFunction`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *InputMonitorAggregationQuery) GetFunctionsOk() (*[]InputFunction, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *InputMonitorAggregationQuery) SetFunctions(v []InputFunction)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *InputMonitorAggregationQuery) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### GetFill

`func (o *InputMonitorAggregationQuery) GetFill() MonitorAggregationQueryFill`

GetFill returns the Fill field if non-nil, zero value otherwise.

### GetFillOk

`func (o *InputMonitorAggregationQuery) GetFillOk() (*MonitorAggregationQueryFill, bool)`

GetFillOk returns a tuple with the Fill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFill

`func (o *InputMonitorAggregationQuery) SetFill(v MonitorAggregationQueryFill)`

SetFill sets Fill field to given value.

### HasFill

`func (o *InputMonitorAggregationQuery) HasFill() bool`

HasFill returns a boolean if a field has been set.

### GetTimeAggregate

`func (o *InputMonitorAggregationQuery) GetTimeAggregate() string`

GetTimeAggregate returns the TimeAggregate field if non-nil, zero value otherwise.

### GetTimeAggregateOk

`func (o *InputMonitorAggregationQuery) GetTimeAggregateOk() (*string, bool)`

GetTimeAggregateOk returns a tuple with the TimeAggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAggregate

`func (o *InputMonitorAggregationQuery) SetTimeAggregate(v string)`

SetTimeAggregate sets TimeAggregate field to given value.

### HasTimeAggregate

`func (o *InputMonitorAggregationQuery) HasTimeAggregate() bool`

HasTimeAggregate returns a boolean if a field has been set.

### GetMetricKind

`func (o *InputMonitorAggregationQuery) GetMetricKind() string`

GetMetricKind returns the MetricKind field if non-nil, zero value otherwise.

### GetMetricKindOk

`func (o *InputMonitorAggregationQuery) GetMetricKindOk() (*string, bool)`

GetMetricKindOk returns a tuple with the MetricKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKind

`func (o *InputMonitorAggregationQuery) SetMetricKind(v string)`

SetMetricKind sets MetricKind field to given value.

### HasMetricKind

`func (o *InputMonitorAggregationQuery) HasMetricKind() bool`

HasMetricKind returns a boolean if a field has been set.

### GetFilter

`func (o *InputMonitorAggregationQuery) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *InputMonitorAggregationQuery) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *InputMonitorAggregationQuery) SetFilter(v string)`

SetFilter sets Filter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


