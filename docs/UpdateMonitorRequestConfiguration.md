# UpdateMonitorRequestConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Threshold monitor over metric aggregations. | 
**Conditions** | [**[]InputMonitorConfigurationMetricConditionsInner**](InputMonitorConfigurationMetricConditionsInner.md) | Threshold conditions evaluated against query and formula results. All listed conditions must match for a non-grouped threshold monitor to alert. | 
**NoDataBehavior** | **string** | Certificate expiry monitors resolve when matching certificate data disappears. | 
**Timeframe** | **int32** | Lookback window, in minutes, that each anomaly monitor evaluation aggregates over. | 
**GroupByFields** | [**[]InputMonitorConfigurationMetricGroupByFieldsInner**](InputMonitorConfigurationMetricGroupByFieldsInner.md) | Monitor group by configuration. The &#x60;limit&#x60; setting is currently ignored; evaluation applies a fixed limit of 100 groups per field. | 
**AggregationAlertLogic** | **string** | Fixed aggregation logic for faulty deployment monitors. Each faulty service version in an environment is evaluated as its own alert group. | 
**ProportionAlertThreshold** | Pointer to **int32** | Percentage threshold used when &#x60;aggregationAlertLogic&#x60; is &#x60;proportion&#x60;. | [optional] 
**ConsecutiveWindows** | Pointer to **int32** | Number of consecutive &#x60;timeframe&#x60;-minute windows that must all meet the conditions before the monitor alerts. 1 or omitted is the standard threshold behavior: one window ending at evaluation time. Above 1, each evaluation checks the last &#x60;consecutiveWindows&#x60; windows and alerts only when every one meets the conditions, per group for grouped monitors. A window that does not meet the conditions breaks the run. A window with no data, such as one where a query with &#x60;fill.mode.type&#x60; &#x60;null&#x60; returns nothing, breaks the run with &#x60;alert&#x60; or &#x60;resolve&#x60;, takes the result of the window before it with &#x60;keep_last_status&#x60;, and is compared as 0 with &#x60;consider_zero&#x60;. A query without &#x60;fill&#x60; counts an empty window as 0. When every window has no data (with &#x60;keep_last_status&#x60;, when the oldest window has none and no window breaks the run), the group follows &#x60;noDataBehavior&#x60; as a single window does. At most 10, and above 1, &#x60;consecutiveWindows&#x60; × &#x60;timeframe&#x60; must not exceed 4320 minutes (3 days). | [optional] 
**Queries** | [**[]InputMonitorAggregationQuery**](InputMonitorAggregationQuery.md) | Aggregation queries used by alerting. Each query is referenced from formulas as q1, q2, and so on. | 
**Condition** | [**InputMonitorConfigurationAnomalyLogCondition**](InputMonitorConfigurationAnomalyLogCondition.md) |  | 
**Filter** | [**InputMonitorConfigurationFaultyDeploymentFilter**](InputMonitorConfigurationFaultyDeploymentFilter.md) |  | 
**AlertOnFirstOccurrence** | Pointer to **bool** | When &#x60;true&#x60;, every new error pattern alerts on the next evaluation, even if it appeared in a single log. When &#x60;false&#x60; or omitted, a new pattern alerts only once it has lasted at least 5 minutes and matched at least 20 logs, which filters out one-off errors. | [optional] 
**WarnBeforeInDays** | **int32** | Number of days before certificate expiry when the monitor should warn. Valid values are 1 through 365. | 
**CloudAccounts** | Pointer to **[]string** | Cloud account IDs whose certificates are checked. Omit to check certificates from all cloud accounts. | [optional] 

## Methods

### NewUpdateMonitorRequestConfiguration

`func NewUpdateMonitorRequestConfiguration(type_ string, conditions []InputMonitorConfigurationMetricConditionsInner, noDataBehavior string, timeframe int32, groupByFields []InputMonitorConfigurationMetricGroupByFieldsInner, aggregationAlertLogic string, queries []InputMonitorAggregationQuery, condition InputMonitorConfigurationAnomalyLogCondition, filter InputMonitorConfigurationFaultyDeploymentFilter, warnBeforeInDays int32, ) *UpdateMonitorRequestConfiguration`

NewUpdateMonitorRequestConfiguration instantiates a new UpdateMonitorRequestConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMonitorRequestConfigurationWithDefaults

`func NewUpdateMonitorRequestConfigurationWithDefaults() *UpdateMonitorRequestConfiguration`

NewUpdateMonitorRequestConfigurationWithDefaults instantiates a new UpdateMonitorRequestConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateMonitorRequestConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateMonitorRequestConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateMonitorRequestConfiguration) SetType(v string)`

SetType sets Type field to given value.


### GetConditions

`func (o *UpdateMonitorRequestConfiguration) GetConditions() []InputMonitorConfigurationMetricConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *UpdateMonitorRequestConfiguration) GetConditionsOk() (*[]InputMonitorConfigurationMetricConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *UpdateMonitorRequestConfiguration) SetConditions(v []InputMonitorConfigurationMetricConditionsInner)`

SetConditions sets Conditions field to given value.


### GetNoDataBehavior

`func (o *UpdateMonitorRequestConfiguration) GetNoDataBehavior() string`

GetNoDataBehavior returns the NoDataBehavior field if non-nil, zero value otherwise.

### GetNoDataBehaviorOk

`func (o *UpdateMonitorRequestConfiguration) GetNoDataBehaviorOk() (*string, bool)`

GetNoDataBehaviorOk returns a tuple with the NoDataBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataBehavior

`func (o *UpdateMonitorRequestConfiguration) SetNoDataBehavior(v string)`

SetNoDataBehavior sets NoDataBehavior field to given value.


### GetTimeframe

`func (o *UpdateMonitorRequestConfiguration) GetTimeframe() int32`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *UpdateMonitorRequestConfiguration) GetTimeframeOk() (*int32, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *UpdateMonitorRequestConfiguration) SetTimeframe(v int32)`

SetTimeframe sets Timeframe field to given value.


### GetGroupByFields

`func (o *UpdateMonitorRequestConfiguration) GetGroupByFields() []InputMonitorConfigurationMetricGroupByFieldsInner`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *UpdateMonitorRequestConfiguration) GetGroupByFieldsOk() (*[]InputMonitorConfigurationMetricGroupByFieldsInner, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *UpdateMonitorRequestConfiguration) SetGroupByFields(v []InputMonitorConfigurationMetricGroupByFieldsInner)`

SetGroupByFields sets GroupByFields field to given value.


### GetAggregationAlertLogic

`func (o *UpdateMonitorRequestConfiguration) GetAggregationAlertLogic() string`

GetAggregationAlertLogic returns the AggregationAlertLogic field if non-nil, zero value otherwise.

### GetAggregationAlertLogicOk

`func (o *UpdateMonitorRequestConfiguration) GetAggregationAlertLogicOk() (*string, bool)`

GetAggregationAlertLogicOk returns a tuple with the AggregationAlertLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationAlertLogic

`func (o *UpdateMonitorRequestConfiguration) SetAggregationAlertLogic(v string)`

SetAggregationAlertLogic sets AggregationAlertLogic field to given value.


### GetProportionAlertThreshold

`func (o *UpdateMonitorRequestConfiguration) GetProportionAlertThreshold() int32`

GetProportionAlertThreshold returns the ProportionAlertThreshold field if non-nil, zero value otherwise.

### GetProportionAlertThresholdOk

`func (o *UpdateMonitorRequestConfiguration) GetProportionAlertThresholdOk() (*int32, bool)`

GetProportionAlertThresholdOk returns a tuple with the ProportionAlertThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProportionAlertThreshold

`func (o *UpdateMonitorRequestConfiguration) SetProportionAlertThreshold(v int32)`

SetProportionAlertThreshold sets ProportionAlertThreshold field to given value.

### HasProportionAlertThreshold

`func (o *UpdateMonitorRequestConfiguration) HasProportionAlertThreshold() bool`

HasProportionAlertThreshold returns a boolean if a field has been set.

### GetConsecutiveWindows

`func (o *UpdateMonitorRequestConfiguration) GetConsecutiveWindows() int32`

GetConsecutiveWindows returns the ConsecutiveWindows field if non-nil, zero value otherwise.

### GetConsecutiveWindowsOk

`func (o *UpdateMonitorRequestConfiguration) GetConsecutiveWindowsOk() (*int32, bool)`

GetConsecutiveWindowsOk returns a tuple with the ConsecutiveWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveWindows

`func (o *UpdateMonitorRequestConfiguration) SetConsecutiveWindows(v int32)`

SetConsecutiveWindows sets ConsecutiveWindows field to given value.

### HasConsecutiveWindows

`func (o *UpdateMonitorRequestConfiguration) HasConsecutiveWindows() bool`

HasConsecutiveWindows returns a boolean if a field has been set.

### GetQueries

`func (o *UpdateMonitorRequestConfiguration) GetQueries() []InputMonitorAggregationQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *UpdateMonitorRequestConfiguration) GetQueriesOk() (*[]InputMonitorAggregationQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *UpdateMonitorRequestConfiguration) SetQueries(v []InputMonitorAggregationQuery)`

SetQueries sets Queries field to given value.


### GetCondition

`func (o *UpdateMonitorRequestConfiguration) GetCondition() InputMonitorConfigurationAnomalyLogCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *UpdateMonitorRequestConfiguration) GetConditionOk() (*InputMonitorConfigurationAnomalyLogCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *UpdateMonitorRequestConfiguration) SetCondition(v InputMonitorConfigurationAnomalyLogCondition)`

SetCondition sets Condition field to given value.


### GetFilter

`func (o *UpdateMonitorRequestConfiguration) GetFilter() InputMonitorConfigurationFaultyDeploymentFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *UpdateMonitorRequestConfiguration) GetFilterOk() (*InputMonitorConfigurationFaultyDeploymentFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *UpdateMonitorRequestConfiguration) SetFilter(v InputMonitorConfigurationFaultyDeploymentFilter)`

SetFilter sets Filter field to given value.


### GetAlertOnFirstOccurrence

`func (o *UpdateMonitorRequestConfiguration) GetAlertOnFirstOccurrence() bool`

GetAlertOnFirstOccurrence returns the AlertOnFirstOccurrence field if non-nil, zero value otherwise.

### GetAlertOnFirstOccurrenceOk

`func (o *UpdateMonitorRequestConfiguration) GetAlertOnFirstOccurrenceOk() (*bool, bool)`

GetAlertOnFirstOccurrenceOk returns a tuple with the AlertOnFirstOccurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFirstOccurrence

`func (o *UpdateMonitorRequestConfiguration) SetAlertOnFirstOccurrence(v bool)`

SetAlertOnFirstOccurrence sets AlertOnFirstOccurrence field to given value.

### HasAlertOnFirstOccurrence

`func (o *UpdateMonitorRequestConfiguration) HasAlertOnFirstOccurrence() bool`

HasAlertOnFirstOccurrence returns a boolean if a field has been set.

### GetWarnBeforeInDays

`func (o *UpdateMonitorRequestConfiguration) GetWarnBeforeInDays() int32`

GetWarnBeforeInDays returns the WarnBeforeInDays field if non-nil, zero value otherwise.

### GetWarnBeforeInDaysOk

`func (o *UpdateMonitorRequestConfiguration) GetWarnBeforeInDaysOk() (*int32, bool)`

GetWarnBeforeInDaysOk returns a tuple with the WarnBeforeInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnBeforeInDays

`func (o *UpdateMonitorRequestConfiguration) SetWarnBeforeInDays(v int32)`

SetWarnBeforeInDays sets WarnBeforeInDays field to given value.


### GetCloudAccounts

`func (o *UpdateMonitorRequestConfiguration) GetCloudAccounts() []string`

GetCloudAccounts returns the CloudAccounts field if non-nil, zero value otherwise.

### GetCloudAccountsOk

`func (o *UpdateMonitorRequestConfiguration) GetCloudAccountsOk() (*[]string, bool)`

GetCloudAccountsOk returns a tuple with the CloudAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccounts

`func (o *UpdateMonitorRequestConfiguration) SetCloudAccounts(v []string)`

SetCloudAccounts sets CloudAccounts field to given value.

### HasCloudAccounts

`func (o *UpdateMonitorRequestConfiguration) HasCloudAccounts() bool`

HasCloudAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


