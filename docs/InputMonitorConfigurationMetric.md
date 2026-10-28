# InputMonitorConfigurationMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Threshold monitor over metric aggregations. | 
**Conditions** | [**[]InputMonitorConfigurationMetricConditionsInner**](InputMonitorConfigurationMetricConditionsInner.md) | Threshold conditions evaluated against query and formula results. All listed conditions must match for a non-grouped threshold monitor to alert. | 
**NoDataBehavior** | **string** | How the monitor updates state when no data is returned. &#x60;alert&#x60; enters no-data alert state, &#x60;resolve&#x60; resolves, &#x60;keep_last_status&#x60; preserves the previous state, and &#x60;consider_zero&#x60; evaluates missing values as zero. | 
**Timeframe** | **int32** | Lookback window, in minutes, that each query-backed threshold monitor evaluation aggregates over. | 
**GroupByFields** | [**[]InputMonitorConfigurationMetricGroupByFieldsInner**](InputMonitorConfigurationMetricGroupByFieldsInner.md) | Monitor group by configuration. The &#x60;limit&#x60; setting is currently ignored; evaluation applies a fixed limit of 100 groups per field. | 
**AggregationAlertLogic** | Pointer to **string** | How grouped results are combined into alert state. Use &#x60;no_aggregation&#x60; only when &#x60;groupByFields&#x60; is empty; use &#x60;all&#x60;, &#x60;any&#x60;, &#x60;each&#x60;, or &#x60;proportion&#x60; with non-empty &#x60;groupByFields&#x60;. &#x60;proportion&#x60; also requires &#x60;proportionAlertThreshold&#x60;. With &#x60;groupByFields&#x60;, send &#x60;each&#x60; so every group opens and resolves its own alert, as the app does by default; &#x60;all&#x60;, &#x60;any&#x60;, and &#x60;proportion&#x60; raise one alert for the whole monitor. On create, omitted query-monitor values default to &#x60;no_aggregation&#x60;; on update, omitted values keep the existing aggregation logic. | [optional] 
**ProportionAlertThreshold** | Pointer to **int32** | Percentage threshold used when &#x60;aggregationAlertLogic&#x60; is &#x60;proportion&#x60;. | [optional] 
**ConsecutiveWindows** | Pointer to **int32** | Number of consecutive &#x60;timeframe&#x60;-minute windows that must all meet the conditions before the monitor alerts. 1 or omitted is the standard threshold behavior: one window ending at evaluation time. Above 1, each evaluation checks the last &#x60;consecutiveWindows&#x60; windows and alerts only when every one meets the conditions, per group for grouped monitors. A window that does not meet the conditions breaks the run, and so does a window with no data unless &#x60;noDataBehavior&#x60; is &#x60;consider_zero&#x60;, which evaluates it as 0. At most 10, and above 1, &#x60;consecutiveWindows&#x60; × &#x60;timeframe&#x60; must not exceed 4320 minutes (3 days). | [optional] 
**Queries** | [**[]InputMonitorAggregationQuery**](InputMonitorAggregationQuery.md) | Aggregation queries used by alerting. Each query is referenced from formulas as q1, q2, and so on. | 

## Methods

### NewInputMonitorConfigurationMetric

`func NewInputMonitorConfigurationMetric(type_ string, conditions []InputMonitorConfigurationMetricConditionsInner, noDataBehavior string, timeframe int32, groupByFields []InputMonitorConfigurationMetricGroupByFieldsInner, queries []InputMonitorAggregationQuery, ) *InputMonitorConfigurationMetric`

NewInputMonitorConfigurationMetric instantiates a new InputMonitorConfigurationMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputMonitorConfigurationMetricWithDefaults

`func NewInputMonitorConfigurationMetricWithDefaults() *InputMonitorConfigurationMetric`

NewInputMonitorConfigurationMetricWithDefaults instantiates a new InputMonitorConfigurationMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputMonitorConfigurationMetric) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputMonitorConfigurationMetric) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputMonitorConfigurationMetric) SetType(v string)`

SetType sets Type field to given value.


### GetConditions

`func (o *InputMonitorConfigurationMetric) GetConditions() []InputMonitorConfigurationMetricConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *InputMonitorConfigurationMetric) GetConditionsOk() (*[]InputMonitorConfigurationMetricConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *InputMonitorConfigurationMetric) SetConditions(v []InputMonitorConfigurationMetricConditionsInner)`

SetConditions sets Conditions field to given value.


### GetNoDataBehavior

`func (o *InputMonitorConfigurationMetric) GetNoDataBehavior() string`

GetNoDataBehavior returns the NoDataBehavior field if non-nil, zero value otherwise.

### GetNoDataBehaviorOk

`func (o *InputMonitorConfigurationMetric) GetNoDataBehaviorOk() (*string, bool)`

GetNoDataBehaviorOk returns a tuple with the NoDataBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataBehavior

`func (o *InputMonitorConfigurationMetric) SetNoDataBehavior(v string)`

SetNoDataBehavior sets NoDataBehavior field to given value.


### GetTimeframe

`func (o *InputMonitorConfigurationMetric) GetTimeframe() int32`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *InputMonitorConfigurationMetric) GetTimeframeOk() (*int32, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *InputMonitorConfigurationMetric) SetTimeframe(v int32)`

SetTimeframe sets Timeframe field to given value.


### GetGroupByFields

`func (o *InputMonitorConfigurationMetric) GetGroupByFields() []InputMonitorConfigurationMetricGroupByFieldsInner`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *InputMonitorConfigurationMetric) GetGroupByFieldsOk() (*[]InputMonitorConfigurationMetricGroupByFieldsInner, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *InputMonitorConfigurationMetric) SetGroupByFields(v []InputMonitorConfigurationMetricGroupByFieldsInner)`

SetGroupByFields sets GroupByFields field to given value.


### GetAggregationAlertLogic

`func (o *InputMonitorConfigurationMetric) GetAggregationAlertLogic() string`

GetAggregationAlertLogic returns the AggregationAlertLogic field if non-nil, zero value otherwise.

### GetAggregationAlertLogicOk

`func (o *InputMonitorConfigurationMetric) GetAggregationAlertLogicOk() (*string, bool)`

GetAggregationAlertLogicOk returns a tuple with the AggregationAlertLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationAlertLogic

`func (o *InputMonitorConfigurationMetric) SetAggregationAlertLogic(v string)`

SetAggregationAlertLogic sets AggregationAlertLogic field to given value.

### HasAggregationAlertLogic

`func (o *InputMonitorConfigurationMetric) HasAggregationAlertLogic() bool`

HasAggregationAlertLogic returns a boolean if a field has been set.

### GetProportionAlertThreshold

`func (o *InputMonitorConfigurationMetric) GetProportionAlertThreshold() int32`

GetProportionAlertThreshold returns the ProportionAlertThreshold field if non-nil, zero value otherwise.

### GetProportionAlertThresholdOk

`func (o *InputMonitorConfigurationMetric) GetProportionAlertThresholdOk() (*int32, bool)`

GetProportionAlertThresholdOk returns a tuple with the ProportionAlertThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProportionAlertThreshold

`func (o *InputMonitorConfigurationMetric) SetProportionAlertThreshold(v int32)`

SetProportionAlertThreshold sets ProportionAlertThreshold field to given value.

### HasProportionAlertThreshold

`func (o *InputMonitorConfigurationMetric) HasProportionAlertThreshold() bool`

HasProportionAlertThreshold returns a boolean if a field has been set.

### GetConsecutiveWindows

`func (o *InputMonitorConfigurationMetric) GetConsecutiveWindows() int32`

GetConsecutiveWindows returns the ConsecutiveWindows field if non-nil, zero value otherwise.

### GetConsecutiveWindowsOk

`func (o *InputMonitorConfigurationMetric) GetConsecutiveWindowsOk() (*int32, bool)`

GetConsecutiveWindowsOk returns a tuple with the ConsecutiveWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveWindows

`func (o *InputMonitorConfigurationMetric) SetConsecutiveWindows(v int32)`

SetConsecutiveWindows sets ConsecutiveWindows field to given value.

### HasConsecutiveWindows

`func (o *InputMonitorConfigurationMetric) HasConsecutiveWindows() bool`

HasConsecutiveWindows returns a boolean if a field has been set.

### GetQueries

`func (o *InputMonitorConfigurationMetric) GetQueries() []InputMonitorAggregationQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *InputMonitorConfigurationMetric) GetQueriesOk() (*[]InputMonitorAggregationQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *InputMonitorConfigurationMetric) SetQueries(v []InputMonitorAggregationQuery)`

SetQueries sets Queries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


