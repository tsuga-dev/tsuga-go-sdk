# InputMonitorConfigurationLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Threshold monitor over log aggregations. | 
**Conditions** | [**[]InputMonitorConfigurationMetricConditionsInner**](InputMonitorConfigurationMetricConditionsInner.md) | Threshold conditions evaluated against query and formula results. All listed conditions must match for a non-grouped threshold monitor to alert. | 
**NoDataBehavior** | **string** | How the monitor updates state when no data is returned. &#x60;alert&#x60; enters no-data alert state, &#x60;resolve&#x60; resolves, &#x60;keep_last_status&#x60; preserves the previous state, and &#x60;consider_zero&#x60; evaluates missing values as zero. | 
**Timeframe** | **int32** | Lookback window, in minutes, that each query-backed threshold monitor evaluation aggregates over. | 
**GroupByFields** | [**[]InputMonitorConfigurationMetricGroupByFieldsInner**](InputMonitorConfigurationMetricGroupByFieldsInner.md) | Monitor group by configuration. The &#x60;limit&#x60; setting is currently ignored; evaluation applies a fixed limit of 100 groups per field. | 
**AggregationAlertLogic** | Pointer to **string** | How grouped results are combined into alert state. Use &#x60;no_aggregation&#x60; only when &#x60;groupByFields&#x60; is empty; use &#x60;all&#x60;, &#x60;any&#x60;, &#x60;each&#x60;, or &#x60;proportion&#x60; with non-empty &#x60;groupByFields&#x60;. &#x60;proportion&#x60; also requires &#x60;proportionAlertThreshold&#x60;. With &#x60;groupByFields&#x60;, send &#x60;each&#x60; so every group opens and resolves its own alert, as the app does by default; &#x60;all&#x60;, &#x60;any&#x60;, and &#x60;proportion&#x60; raise one alert for the whole monitor. On create, omitted query-monitor values default to &#x60;no_aggregation&#x60;; on update, omitted values keep the existing aggregation logic. | [optional] 
**ProportionAlertThreshold** | Pointer to **int32** | Percentage threshold used when &#x60;aggregationAlertLogic&#x60; is &#x60;proportion&#x60;. | [optional] 
**ConsecutiveWindows** | Pointer to **int32** | Number of consecutive &#x60;timeframe&#x60;-minute windows that must all meet the conditions before the monitor alerts. 1 or omitted is the standard threshold behavior: one window ending at evaluation time. Above 1, each evaluation checks the last &#x60;consecutiveWindows&#x60; windows and alerts only when every one meets the conditions, per group for grouped monitors. A window that does not meet the conditions breaks the run. A window with no data, such as one where a query with &#x60;fill.mode.type&#x60; &#x60;null&#x60; returns nothing, breaks the run with &#x60;alert&#x60; or &#x60;resolve&#x60;, takes the result of the window before it with &#x60;keep_last_status&#x60;, and is compared as 0 with &#x60;consider_zero&#x60;. A query without &#x60;fill&#x60; counts an empty window as 0. When every window has no data (with &#x60;keep_last_status&#x60;, when the oldest window has none and no window breaks the run), the group follows &#x60;noDataBehavior&#x60; as a single window does. At most 10, and above 1, &#x60;consecutiveWindows&#x60; × &#x60;timeframe&#x60; must not exceed 4320 minutes (3 days). | [optional] 
**Queries** | [**[]InputMonitorAggregationQuery**](InputMonitorAggregationQuery.md) | Aggregation queries used by alerting. Each query is referenced from formulas as q1, q2, and so on. | 

## Methods

### NewInputMonitorConfigurationLog

`func NewInputMonitorConfigurationLog(type_ string, conditions []InputMonitorConfigurationMetricConditionsInner, noDataBehavior string, timeframe int32, groupByFields []InputMonitorConfigurationMetricGroupByFieldsInner, queries []InputMonitorAggregationQuery, ) *InputMonitorConfigurationLog`

NewInputMonitorConfigurationLog instantiates a new InputMonitorConfigurationLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputMonitorConfigurationLogWithDefaults

`func NewInputMonitorConfigurationLogWithDefaults() *InputMonitorConfigurationLog`

NewInputMonitorConfigurationLogWithDefaults instantiates a new InputMonitorConfigurationLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputMonitorConfigurationLog) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputMonitorConfigurationLog) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputMonitorConfigurationLog) SetType(v string)`

SetType sets Type field to given value.


### GetConditions

`func (o *InputMonitorConfigurationLog) GetConditions() []InputMonitorConfigurationMetricConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *InputMonitorConfigurationLog) GetConditionsOk() (*[]InputMonitorConfigurationMetricConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *InputMonitorConfigurationLog) SetConditions(v []InputMonitorConfigurationMetricConditionsInner)`

SetConditions sets Conditions field to given value.


### GetNoDataBehavior

`func (o *InputMonitorConfigurationLog) GetNoDataBehavior() string`

GetNoDataBehavior returns the NoDataBehavior field if non-nil, zero value otherwise.

### GetNoDataBehaviorOk

`func (o *InputMonitorConfigurationLog) GetNoDataBehaviorOk() (*string, bool)`

GetNoDataBehaviorOk returns a tuple with the NoDataBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataBehavior

`func (o *InputMonitorConfigurationLog) SetNoDataBehavior(v string)`

SetNoDataBehavior sets NoDataBehavior field to given value.


### GetTimeframe

`func (o *InputMonitorConfigurationLog) GetTimeframe() int32`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *InputMonitorConfigurationLog) GetTimeframeOk() (*int32, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *InputMonitorConfigurationLog) SetTimeframe(v int32)`

SetTimeframe sets Timeframe field to given value.


### GetGroupByFields

`func (o *InputMonitorConfigurationLog) GetGroupByFields() []InputMonitorConfigurationMetricGroupByFieldsInner`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *InputMonitorConfigurationLog) GetGroupByFieldsOk() (*[]InputMonitorConfigurationMetricGroupByFieldsInner, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *InputMonitorConfigurationLog) SetGroupByFields(v []InputMonitorConfigurationMetricGroupByFieldsInner)`

SetGroupByFields sets GroupByFields field to given value.


### GetAggregationAlertLogic

`func (o *InputMonitorConfigurationLog) GetAggregationAlertLogic() string`

GetAggregationAlertLogic returns the AggregationAlertLogic field if non-nil, zero value otherwise.

### GetAggregationAlertLogicOk

`func (o *InputMonitorConfigurationLog) GetAggregationAlertLogicOk() (*string, bool)`

GetAggregationAlertLogicOk returns a tuple with the AggregationAlertLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationAlertLogic

`func (o *InputMonitorConfigurationLog) SetAggregationAlertLogic(v string)`

SetAggregationAlertLogic sets AggregationAlertLogic field to given value.

### HasAggregationAlertLogic

`func (o *InputMonitorConfigurationLog) HasAggregationAlertLogic() bool`

HasAggregationAlertLogic returns a boolean if a field has been set.

### GetProportionAlertThreshold

`func (o *InputMonitorConfigurationLog) GetProportionAlertThreshold() int32`

GetProportionAlertThreshold returns the ProportionAlertThreshold field if non-nil, zero value otherwise.

### GetProportionAlertThresholdOk

`func (o *InputMonitorConfigurationLog) GetProportionAlertThresholdOk() (*int32, bool)`

GetProportionAlertThresholdOk returns a tuple with the ProportionAlertThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProportionAlertThreshold

`func (o *InputMonitorConfigurationLog) SetProportionAlertThreshold(v int32)`

SetProportionAlertThreshold sets ProportionAlertThreshold field to given value.

### HasProportionAlertThreshold

`func (o *InputMonitorConfigurationLog) HasProportionAlertThreshold() bool`

HasProportionAlertThreshold returns a boolean if a field has been set.

### GetConsecutiveWindows

`func (o *InputMonitorConfigurationLog) GetConsecutiveWindows() int32`

GetConsecutiveWindows returns the ConsecutiveWindows field if non-nil, zero value otherwise.

### GetConsecutiveWindowsOk

`func (o *InputMonitorConfigurationLog) GetConsecutiveWindowsOk() (*int32, bool)`

GetConsecutiveWindowsOk returns a tuple with the ConsecutiveWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveWindows

`func (o *InputMonitorConfigurationLog) SetConsecutiveWindows(v int32)`

SetConsecutiveWindows sets ConsecutiveWindows field to given value.

### HasConsecutiveWindows

`func (o *InputMonitorConfigurationLog) HasConsecutiveWindows() bool`

HasConsecutiveWindows returns a boolean if a field has been set.

### GetQueries

`func (o *InputMonitorConfigurationLog) GetQueries() []InputMonitorAggregationQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *InputMonitorConfigurationLog) GetQueriesOk() (*[]InputMonitorAggregationQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *InputMonitorConfigurationLog) SetQueries(v []InputMonitorAggregationQuery)`

SetQueries sets Queries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


