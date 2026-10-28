# OngoingMonitorAlertMonitorConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Threshold monitor over metric aggregations. | 
**Conditions** | [**[]MonitorConfigurationMetricConditionsInner**](MonitorConfigurationMetricConditionsInner.md) | Threshold conditions evaluated against query and formula results. All listed conditions must match for a non-grouped threshold monitor to alert. | 
**NoDataBehavior** | **string** | Fixed no-data behavior for log error pattern increase monitors. A pattern group resolves once its volume stops increasing. | 
**Timeframe** | **float32** | Lookback window, in minutes, that each anomaly monitor evaluation aggregates over. Always a whole number from 5 through 1440. | 
**GroupByFields** | [**[]MonitorConfigurationMetricGroupByFieldsInner**](MonitorConfigurationMetricGroupByFieldsInner.md) | Monitor group by configuration. The &#x60;limit&#x60; setting is currently ignored; evaluation applies a fixed limit of 100 groups per field. | 
**AggregationAlertLogic** | **string** | Fixed aggregation logic for faulty deployment monitors. Each faulty service version in an environment is evaluated as its own alert group. | 
**ProportionAlertThreshold** | Pointer to **int32** | Percentage threshold used when &#x60;aggregationAlertLogic&#x60; is &#x60;proportion&#x60;. | [optional] 
**ConsecutiveWindows** | Pointer to **int32** | Number of consecutive &#x60;timeframe&#x60;-minute windows that must all meet the conditions before the monitor alerts. 1 or omitted is the standard threshold behavior: one window ending at evaluation time. Above 1, each evaluation checks the last &#x60;consecutiveWindows&#x60; windows and alerts only when every one meets the conditions, per group for grouped monitors. A window that does not meet the conditions breaks the run, and so does a window with no data unless &#x60;noDataBehavior&#x60; is &#x60;consider_zero&#x60;, which evaluates it as 0. At most 10, and above 1, &#x60;consecutiveWindows&#x60; × &#x60;timeframe&#x60; must not exceed 4320 minutes (3 days). | [optional] 
**Queries** | [**[]MonitorAggregationQuery**](MonitorAggregationQuery.md) | Aggregation queries used by alerting and SLO evaluation. Each query is referenced from formulas as q1, q2, and so on. | 
**Condition** | [**MonitorConfigurationAnomalyLogCondition**](MonitorConfigurationAnomalyLogCondition.md) |  | 
**WarnBeforeInDays** | **int32** | Number of days before certificate expiry when the monitor should warn. | 
**CloudAccounts** | Pointer to **[]string** | Cloud account IDs whose certificates are checked. Omitted means certificates from all cloud accounts are checked. | [optional] 
**Filter** | [**MonitorConfigurationFaultyDeploymentFilter**](MonitorConfigurationFaultyDeploymentFilter.md) |  | 

## Methods

### NewOngoingMonitorAlertMonitorConfiguration

`func NewOngoingMonitorAlertMonitorConfiguration(type_ string, conditions []MonitorConfigurationMetricConditionsInner, noDataBehavior string, timeframe float32, groupByFields []MonitorConfigurationMetricGroupByFieldsInner, aggregationAlertLogic string, queries []MonitorAggregationQuery, condition MonitorConfigurationAnomalyLogCondition, warnBeforeInDays int32, filter MonitorConfigurationFaultyDeploymentFilter, ) *OngoingMonitorAlertMonitorConfiguration`

NewOngoingMonitorAlertMonitorConfiguration instantiates a new OngoingMonitorAlertMonitorConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOngoingMonitorAlertMonitorConfigurationWithDefaults

`func NewOngoingMonitorAlertMonitorConfigurationWithDefaults() *OngoingMonitorAlertMonitorConfiguration`

NewOngoingMonitorAlertMonitorConfigurationWithDefaults instantiates a new OngoingMonitorAlertMonitorConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OngoingMonitorAlertMonitorConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OngoingMonitorAlertMonitorConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OngoingMonitorAlertMonitorConfiguration) SetType(v string)`

SetType sets Type field to given value.


### GetConditions

`func (o *OngoingMonitorAlertMonitorConfiguration) GetConditions() []MonitorConfigurationMetricConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *OngoingMonitorAlertMonitorConfiguration) GetConditionsOk() (*[]MonitorConfigurationMetricConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *OngoingMonitorAlertMonitorConfiguration) SetConditions(v []MonitorConfigurationMetricConditionsInner)`

SetConditions sets Conditions field to given value.


### GetNoDataBehavior

`func (o *OngoingMonitorAlertMonitorConfiguration) GetNoDataBehavior() string`

GetNoDataBehavior returns the NoDataBehavior field if non-nil, zero value otherwise.

### GetNoDataBehaviorOk

`func (o *OngoingMonitorAlertMonitorConfiguration) GetNoDataBehaviorOk() (*string, bool)`

GetNoDataBehaviorOk returns a tuple with the NoDataBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataBehavior

`func (o *OngoingMonitorAlertMonitorConfiguration) SetNoDataBehavior(v string)`

SetNoDataBehavior sets NoDataBehavior field to given value.


### GetTimeframe

`func (o *OngoingMonitorAlertMonitorConfiguration) GetTimeframe() float32`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *OngoingMonitorAlertMonitorConfiguration) GetTimeframeOk() (*float32, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *OngoingMonitorAlertMonitorConfiguration) SetTimeframe(v float32)`

SetTimeframe sets Timeframe field to given value.


### GetGroupByFields

`func (o *OngoingMonitorAlertMonitorConfiguration) GetGroupByFields() []MonitorConfigurationMetricGroupByFieldsInner`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *OngoingMonitorAlertMonitorConfiguration) GetGroupByFieldsOk() (*[]MonitorConfigurationMetricGroupByFieldsInner, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *OngoingMonitorAlertMonitorConfiguration) SetGroupByFields(v []MonitorConfigurationMetricGroupByFieldsInner)`

SetGroupByFields sets GroupByFields field to given value.


### GetAggregationAlertLogic

`func (o *OngoingMonitorAlertMonitorConfiguration) GetAggregationAlertLogic() string`

GetAggregationAlertLogic returns the AggregationAlertLogic field if non-nil, zero value otherwise.

### GetAggregationAlertLogicOk

`func (o *OngoingMonitorAlertMonitorConfiguration) GetAggregationAlertLogicOk() (*string, bool)`

GetAggregationAlertLogicOk returns a tuple with the AggregationAlertLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationAlertLogic

`func (o *OngoingMonitorAlertMonitorConfiguration) SetAggregationAlertLogic(v string)`

SetAggregationAlertLogic sets AggregationAlertLogic field to given value.


### GetProportionAlertThreshold

`func (o *OngoingMonitorAlertMonitorConfiguration) GetProportionAlertThreshold() int32`

GetProportionAlertThreshold returns the ProportionAlertThreshold field if non-nil, zero value otherwise.

### GetProportionAlertThresholdOk

`func (o *OngoingMonitorAlertMonitorConfiguration) GetProportionAlertThresholdOk() (*int32, bool)`

GetProportionAlertThresholdOk returns a tuple with the ProportionAlertThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProportionAlertThreshold

`func (o *OngoingMonitorAlertMonitorConfiguration) SetProportionAlertThreshold(v int32)`

SetProportionAlertThreshold sets ProportionAlertThreshold field to given value.

### HasProportionAlertThreshold

`func (o *OngoingMonitorAlertMonitorConfiguration) HasProportionAlertThreshold() bool`

HasProportionAlertThreshold returns a boolean if a field has been set.

### GetConsecutiveWindows

`func (o *OngoingMonitorAlertMonitorConfiguration) GetConsecutiveWindows() int32`

GetConsecutiveWindows returns the ConsecutiveWindows field if non-nil, zero value otherwise.

### GetConsecutiveWindowsOk

`func (o *OngoingMonitorAlertMonitorConfiguration) GetConsecutiveWindowsOk() (*int32, bool)`

GetConsecutiveWindowsOk returns a tuple with the ConsecutiveWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveWindows

`func (o *OngoingMonitorAlertMonitorConfiguration) SetConsecutiveWindows(v int32)`

SetConsecutiveWindows sets ConsecutiveWindows field to given value.

### HasConsecutiveWindows

`func (o *OngoingMonitorAlertMonitorConfiguration) HasConsecutiveWindows() bool`

HasConsecutiveWindows returns a boolean if a field has been set.

### GetQueries

`func (o *OngoingMonitorAlertMonitorConfiguration) GetQueries() []MonitorAggregationQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *OngoingMonitorAlertMonitorConfiguration) GetQueriesOk() (*[]MonitorAggregationQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *OngoingMonitorAlertMonitorConfiguration) SetQueries(v []MonitorAggregationQuery)`

SetQueries sets Queries field to given value.


### GetCondition

`func (o *OngoingMonitorAlertMonitorConfiguration) GetCondition() MonitorConfigurationAnomalyLogCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *OngoingMonitorAlertMonitorConfiguration) GetConditionOk() (*MonitorConfigurationAnomalyLogCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *OngoingMonitorAlertMonitorConfiguration) SetCondition(v MonitorConfigurationAnomalyLogCondition)`

SetCondition sets Condition field to given value.


### GetWarnBeforeInDays

`func (o *OngoingMonitorAlertMonitorConfiguration) GetWarnBeforeInDays() int32`

GetWarnBeforeInDays returns the WarnBeforeInDays field if non-nil, zero value otherwise.

### GetWarnBeforeInDaysOk

`func (o *OngoingMonitorAlertMonitorConfiguration) GetWarnBeforeInDaysOk() (*int32, bool)`

GetWarnBeforeInDaysOk returns a tuple with the WarnBeforeInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnBeforeInDays

`func (o *OngoingMonitorAlertMonitorConfiguration) SetWarnBeforeInDays(v int32)`

SetWarnBeforeInDays sets WarnBeforeInDays field to given value.


### GetCloudAccounts

`func (o *OngoingMonitorAlertMonitorConfiguration) GetCloudAccounts() []string`

GetCloudAccounts returns the CloudAccounts field if non-nil, zero value otherwise.

### GetCloudAccountsOk

`func (o *OngoingMonitorAlertMonitorConfiguration) GetCloudAccountsOk() (*[]string, bool)`

GetCloudAccountsOk returns a tuple with the CloudAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccounts

`func (o *OngoingMonitorAlertMonitorConfiguration) SetCloudAccounts(v []string)`

SetCloudAccounts sets CloudAccounts field to given value.

### HasCloudAccounts

`func (o *OngoingMonitorAlertMonitorConfiguration) HasCloudAccounts() bool`

HasCloudAccounts returns a boolean if a field has been set.

### GetFilter

`func (o *OngoingMonitorAlertMonitorConfiguration) GetFilter() MonitorConfigurationFaultyDeploymentFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *OngoingMonitorAlertMonitorConfiguration) GetFilterOk() (*MonitorConfigurationFaultyDeploymentFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *OngoingMonitorAlertMonitorConfiguration) SetFilter(v MonitorConfigurationFaultyDeploymentFilter)`

SetFilter sets Filter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


