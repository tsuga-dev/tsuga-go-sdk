# InputMonitorConfigurationAnomalyMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Anomaly monitor over metric aggregations. | 
**Condition** | [**InputMonitorConfigurationAnomalyLogCondition**](InputMonitorConfigurationAnomalyLogCondition.md) |  | 
**NoDataBehavior** | **string** | How the anomaly monitor updates state when no data is returned. &#x60;alert&#x60; enters no-data alert state, &#x60;resolve&#x60; resolves, and &#x60;keep_last_status&#x60; preserves the previous state. | 
**Timeframe** | **float32** | Lookback window, in minutes, that each anomaly monitor evaluation aggregates over. Valid input is 5 through 1440 minutes. | 
**GroupByFields** | [**[]InputMonitorConfigurationMetricGroupByFieldsInner**](InputMonitorConfigurationMetricGroupByFieldsInner.md) | Monitor group by configuration. Warning! Note that the limit setting is currently ignored. | 
**AggregationAlertLogic** | Pointer to **string** | How grouped results are combined into alert state. Use &#x60;no_aggregation&#x60; only when &#x60;groupByFields&#x60; is empty; use &#x60;all&#x60;, &#x60;any&#x60;, &#x60;each&#x60;, or &#x60;proportion&#x60; with non-empty &#x60;groupByFields&#x60;. &#x60;proportion&#x60; also requires &#x60;proportionAlertThreshold&#x60;. On create, omitted query-monitor values default to &#x60;no_aggregation&#x60;; on update, omitted values keep the existing aggregation logic. | [optional] 
**ProportionAlertThreshold** | Pointer to **int32** | Percentage threshold used when &#x60;aggregationAlertLogic&#x60; is &#x60;proportion&#x60;. Valid values are 1 through 99. | [optional] 
**Queries** | [**[]MonitorAggregationQuery1**](MonitorAggregationQuery1.md) | Aggregation queries used by alerting and SLO evaluation. Each query is referenced from formulas as q1, q2, and so on. | 

## Methods

### NewInputMonitorConfigurationAnomalyMetric

`func NewInputMonitorConfigurationAnomalyMetric(type_ string, condition InputMonitorConfigurationAnomalyLogCondition, noDataBehavior string, timeframe float32, groupByFields []InputMonitorConfigurationMetricGroupByFieldsInner, queries []MonitorAggregationQuery1, ) *InputMonitorConfigurationAnomalyMetric`

NewInputMonitorConfigurationAnomalyMetric instantiates a new InputMonitorConfigurationAnomalyMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputMonitorConfigurationAnomalyMetricWithDefaults

`func NewInputMonitorConfigurationAnomalyMetricWithDefaults() *InputMonitorConfigurationAnomalyMetric`

NewInputMonitorConfigurationAnomalyMetricWithDefaults instantiates a new InputMonitorConfigurationAnomalyMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputMonitorConfigurationAnomalyMetric) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputMonitorConfigurationAnomalyMetric) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputMonitorConfigurationAnomalyMetric) SetType(v string)`

SetType sets Type field to given value.


### GetCondition

`func (o *InputMonitorConfigurationAnomalyMetric) GetCondition() InputMonitorConfigurationAnomalyLogCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *InputMonitorConfigurationAnomalyMetric) GetConditionOk() (*InputMonitorConfigurationAnomalyLogCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *InputMonitorConfigurationAnomalyMetric) SetCondition(v InputMonitorConfigurationAnomalyLogCondition)`

SetCondition sets Condition field to given value.


### GetNoDataBehavior

`func (o *InputMonitorConfigurationAnomalyMetric) GetNoDataBehavior() string`

GetNoDataBehavior returns the NoDataBehavior field if non-nil, zero value otherwise.

### GetNoDataBehaviorOk

`func (o *InputMonitorConfigurationAnomalyMetric) GetNoDataBehaviorOk() (*string, bool)`

GetNoDataBehaviorOk returns a tuple with the NoDataBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataBehavior

`func (o *InputMonitorConfigurationAnomalyMetric) SetNoDataBehavior(v string)`

SetNoDataBehavior sets NoDataBehavior field to given value.


### GetTimeframe

`func (o *InputMonitorConfigurationAnomalyMetric) GetTimeframe() float32`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *InputMonitorConfigurationAnomalyMetric) GetTimeframeOk() (*float32, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *InputMonitorConfigurationAnomalyMetric) SetTimeframe(v float32)`

SetTimeframe sets Timeframe field to given value.


### GetGroupByFields

`func (o *InputMonitorConfigurationAnomalyMetric) GetGroupByFields() []InputMonitorConfigurationMetricGroupByFieldsInner`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *InputMonitorConfigurationAnomalyMetric) GetGroupByFieldsOk() (*[]InputMonitorConfigurationMetricGroupByFieldsInner, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *InputMonitorConfigurationAnomalyMetric) SetGroupByFields(v []InputMonitorConfigurationMetricGroupByFieldsInner)`

SetGroupByFields sets GroupByFields field to given value.


### GetAggregationAlertLogic

`func (o *InputMonitorConfigurationAnomalyMetric) GetAggregationAlertLogic() string`

GetAggregationAlertLogic returns the AggregationAlertLogic field if non-nil, zero value otherwise.

### GetAggregationAlertLogicOk

`func (o *InputMonitorConfigurationAnomalyMetric) GetAggregationAlertLogicOk() (*string, bool)`

GetAggregationAlertLogicOk returns a tuple with the AggregationAlertLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationAlertLogic

`func (o *InputMonitorConfigurationAnomalyMetric) SetAggregationAlertLogic(v string)`

SetAggregationAlertLogic sets AggregationAlertLogic field to given value.

### HasAggregationAlertLogic

`func (o *InputMonitorConfigurationAnomalyMetric) HasAggregationAlertLogic() bool`

HasAggregationAlertLogic returns a boolean if a field has been set.

### GetProportionAlertThreshold

`func (o *InputMonitorConfigurationAnomalyMetric) GetProportionAlertThreshold() int32`

GetProportionAlertThreshold returns the ProportionAlertThreshold field if non-nil, zero value otherwise.

### GetProportionAlertThresholdOk

`func (o *InputMonitorConfigurationAnomalyMetric) GetProportionAlertThresholdOk() (*int32, bool)`

GetProportionAlertThresholdOk returns a tuple with the ProportionAlertThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProportionAlertThreshold

`func (o *InputMonitorConfigurationAnomalyMetric) SetProportionAlertThreshold(v int32)`

SetProportionAlertThreshold sets ProportionAlertThreshold field to given value.

### HasProportionAlertThreshold

`func (o *InputMonitorConfigurationAnomalyMetric) HasProportionAlertThreshold() bool`

HasProportionAlertThreshold returns a boolean if a field has been set.

### GetQueries

`func (o *InputMonitorConfigurationAnomalyMetric) GetQueries() []MonitorAggregationQuery1`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *InputMonitorConfigurationAnomalyMetric) GetQueriesOk() (*[]MonitorAggregationQuery1, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *InputMonitorConfigurationAnomalyMetric) SetQueries(v []MonitorAggregationQuery1)`

SetQueries sets Queries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


