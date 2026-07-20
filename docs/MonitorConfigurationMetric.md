# MonitorConfigurationMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Threshold monitor over metric aggregations. | 
**Conditions** | [**[]MonitorConfigurationMetricConditionsInner**](MonitorConfigurationMetricConditionsInner.md) | Threshold conditions evaluated against query and formula results. All listed conditions must match for a non-grouped threshold monitor to alert. | 
**NoDataBehavior** | **string** | How the monitor updates state when no data is returned. &#x60;alert&#x60; enters no-data alert state, &#x60;resolve&#x60; resolves, &#x60;keep_last_status&#x60; preserves the previous state, and &#x60;consider_zero&#x60; evaluates missing values as zero. | 
**Timeframe** | **float32** | Lookback window, in minutes, that each monitor evaluation aggregates over. | 
**GroupByFields** | [**[]MonitorConfigurationMetricGroupByFieldsInner**](MonitorConfigurationMetricGroupByFieldsInner.md) | Monitor group by configuration. Warning! Note that the limit setting is currently ignored. | 
**AggregationAlertLogic** | Pointer to **string** | How grouped results are combined into alert state. Use &#x60;no_aggregation&#x60; only when &#x60;groupByFields&#x60; is empty; use &#x60;all&#x60;, &#x60;any&#x60;, &#x60;each&#x60;, or &#x60;proportion&#x60; with non-empty &#x60;groupByFields&#x60;. &#x60;proportion&#x60; also requires &#x60;proportionAlertThreshold&#x60;. | [optional] 
**ProportionAlertThreshold** | Pointer to **int32** | Percentage threshold used when &#x60;aggregationAlertLogic&#x60; is &#x60;proportion&#x60;. | [optional] 
**Queries** | [**[]MonitorAggregationQuery**](MonitorAggregationQuery.md) | Aggregation queries used by alerting and SLO evaluation. Each query is referenced from formulas as q1, q2, and so on. | 

## Methods

### NewMonitorConfigurationMetric

`func NewMonitorConfigurationMetric(type_ string, conditions []MonitorConfigurationMetricConditionsInner, noDataBehavior string, timeframe float32, groupByFields []MonitorConfigurationMetricGroupByFieldsInner, queries []MonitorAggregationQuery, ) *MonitorConfigurationMetric`

NewMonitorConfigurationMetric instantiates a new MonitorConfigurationMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorConfigurationMetricWithDefaults

`func NewMonitorConfigurationMetricWithDefaults() *MonitorConfigurationMetric`

NewMonitorConfigurationMetricWithDefaults instantiates a new MonitorConfigurationMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MonitorConfigurationMetric) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorConfigurationMetric) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorConfigurationMetric) SetType(v string)`

SetType sets Type field to given value.


### GetConditions

`func (o *MonitorConfigurationMetric) GetConditions() []MonitorConfigurationMetricConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *MonitorConfigurationMetric) GetConditionsOk() (*[]MonitorConfigurationMetricConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *MonitorConfigurationMetric) SetConditions(v []MonitorConfigurationMetricConditionsInner)`

SetConditions sets Conditions field to given value.


### GetNoDataBehavior

`func (o *MonitorConfigurationMetric) GetNoDataBehavior() string`

GetNoDataBehavior returns the NoDataBehavior field if non-nil, zero value otherwise.

### GetNoDataBehaviorOk

`func (o *MonitorConfigurationMetric) GetNoDataBehaviorOk() (*string, bool)`

GetNoDataBehaviorOk returns a tuple with the NoDataBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataBehavior

`func (o *MonitorConfigurationMetric) SetNoDataBehavior(v string)`

SetNoDataBehavior sets NoDataBehavior field to given value.


### GetTimeframe

`func (o *MonitorConfigurationMetric) GetTimeframe() float32`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *MonitorConfigurationMetric) GetTimeframeOk() (*float32, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *MonitorConfigurationMetric) SetTimeframe(v float32)`

SetTimeframe sets Timeframe field to given value.


### GetGroupByFields

`func (o *MonitorConfigurationMetric) GetGroupByFields() []MonitorConfigurationMetricGroupByFieldsInner`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *MonitorConfigurationMetric) GetGroupByFieldsOk() (*[]MonitorConfigurationMetricGroupByFieldsInner, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *MonitorConfigurationMetric) SetGroupByFields(v []MonitorConfigurationMetricGroupByFieldsInner)`

SetGroupByFields sets GroupByFields field to given value.


### GetAggregationAlertLogic

`func (o *MonitorConfigurationMetric) GetAggregationAlertLogic() string`

GetAggregationAlertLogic returns the AggregationAlertLogic field if non-nil, zero value otherwise.

### GetAggregationAlertLogicOk

`func (o *MonitorConfigurationMetric) GetAggregationAlertLogicOk() (*string, bool)`

GetAggregationAlertLogicOk returns a tuple with the AggregationAlertLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationAlertLogic

`func (o *MonitorConfigurationMetric) SetAggregationAlertLogic(v string)`

SetAggregationAlertLogic sets AggregationAlertLogic field to given value.

### HasAggregationAlertLogic

`func (o *MonitorConfigurationMetric) HasAggregationAlertLogic() bool`

HasAggregationAlertLogic returns a boolean if a field has been set.

### GetProportionAlertThreshold

`func (o *MonitorConfigurationMetric) GetProportionAlertThreshold() int32`

GetProportionAlertThreshold returns the ProportionAlertThreshold field if non-nil, zero value otherwise.

### GetProportionAlertThresholdOk

`func (o *MonitorConfigurationMetric) GetProportionAlertThresholdOk() (*int32, bool)`

GetProportionAlertThresholdOk returns a tuple with the ProportionAlertThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProportionAlertThreshold

`func (o *MonitorConfigurationMetric) SetProportionAlertThreshold(v int32)`

SetProportionAlertThreshold sets ProportionAlertThreshold field to given value.

### HasProportionAlertThreshold

`func (o *MonitorConfigurationMetric) HasProportionAlertThreshold() bool`

HasProportionAlertThreshold returns a boolean if a field has been set.

### GetQueries

`func (o *MonitorConfigurationMetric) GetQueries() []MonitorAggregationQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *MonitorConfigurationMetric) GetQueriesOk() (*[]MonitorAggregationQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *MonitorConfigurationMetric) SetQueries(v []MonitorAggregationQuery)`

SetQueries sets Queries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


