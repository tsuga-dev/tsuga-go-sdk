# MonitorConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Conditions** | [**[]MonitorConfigurationMetricConditionsInner**](MonitorConfigurationMetricConditionsInner.md) |  | 
**NoDataBehavior** | **string** |  | 
**Timeframe** | **float32** | Timeframe of the monitor in minutes | 
**GroupByFields** | [**[]MonitorConfigurationMetricGroupByFieldsInner**](MonitorConfigurationMetricGroupByFieldsInner.md) | Monitor group by configuration. Warning! Note that the limit setting is currently ignored. | 
**AggregationAlertLogic** | **string** |  | 
**ProportionAlertThreshold** | Pointer to **int32** |  | [optional] 
**Queries** | [**[]MonitorAggregationQuery**](MonitorAggregationQuery.md) |  | 
**Condition** | [**MonitorConfigurationAnomalyLogCondition**](MonitorConfigurationAnomalyLogCondition.md) |  | 
**Filter** | [**MonitorConfigurationLogErrorPatternFilter**](MonitorConfigurationLogErrorPatternFilter.md) |  | 
**WarnBeforeInDays** | **int32** |  | 
**CloudAccounts** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMonitorConfiguration

`func NewMonitorConfiguration(type_ string, conditions []MonitorConfigurationMetricConditionsInner, noDataBehavior string, timeframe float32, groupByFields []MonitorConfigurationMetricGroupByFieldsInner, aggregationAlertLogic string, queries []MonitorAggregationQuery, condition MonitorConfigurationAnomalyLogCondition, filter MonitorConfigurationLogErrorPatternFilter, warnBeforeInDays int32, ) *MonitorConfiguration`

NewMonitorConfiguration instantiates a new MonitorConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorConfigurationWithDefaults

`func NewMonitorConfigurationWithDefaults() *MonitorConfiguration`

NewMonitorConfigurationWithDefaults instantiates a new MonitorConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MonitorConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorConfiguration) SetType(v string)`

SetType sets Type field to given value.


### GetConditions

`func (o *MonitorConfiguration) GetConditions() []MonitorConfigurationMetricConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *MonitorConfiguration) GetConditionsOk() (*[]MonitorConfigurationMetricConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *MonitorConfiguration) SetConditions(v []MonitorConfigurationMetricConditionsInner)`

SetConditions sets Conditions field to given value.


### GetNoDataBehavior

`func (o *MonitorConfiguration) GetNoDataBehavior() string`

GetNoDataBehavior returns the NoDataBehavior field if non-nil, zero value otherwise.

### GetNoDataBehaviorOk

`func (o *MonitorConfiguration) GetNoDataBehaviorOk() (*string, bool)`

GetNoDataBehaviorOk returns a tuple with the NoDataBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataBehavior

`func (o *MonitorConfiguration) SetNoDataBehavior(v string)`

SetNoDataBehavior sets NoDataBehavior field to given value.


### GetTimeframe

`func (o *MonitorConfiguration) GetTimeframe() float32`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *MonitorConfiguration) GetTimeframeOk() (*float32, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *MonitorConfiguration) SetTimeframe(v float32)`

SetTimeframe sets Timeframe field to given value.


### GetGroupByFields

`func (o *MonitorConfiguration) GetGroupByFields() []MonitorConfigurationMetricGroupByFieldsInner`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *MonitorConfiguration) GetGroupByFieldsOk() (*[]MonitorConfigurationMetricGroupByFieldsInner, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *MonitorConfiguration) SetGroupByFields(v []MonitorConfigurationMetricGroupByFieldsInner)`

SetGroupByFields sets GroupByFields field to given value.


### GetAggregationAlertLogic

`func (o *MonitorConfiguration) GetAggregationAlertLogic() string`

GetAggregationAlertLogic returns the AggregationAlertLogic field if non-nil, zero value otherwise.

### GetAggregationAlertLogicOk

`func (o *MonitorConfiguration) GetAggregationAlertLogicOk() (*string, bool)`

GetAggregationAlertLogicOk returns a tuple with the AggregationAlertLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationAlertLogic

`func (o *MonitorConfiguration) SetAggregationAlertLogic(v string)`

SetAggregationAlertLogic sets AggregationAlertLogic field to given value.


### GetProportionAlertThreshold

`func (o *MonitorConfiguration) GetProportionAlertThreshold() int32`

GetProportionAlertThreshold returns the ProportionAlertThreshold field if non-nil, zero value otherwise.

### GetProportionAlertThresholdOk

`func (o *MonitorConfiguration) GetProportionAlertThresholdOk() (*int32, bool)`

GetProportionAlertThresholdOk returns a tuple with the ProportionAlertThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProportionAlertThreshold

`func (o *MonitorConfiguration) SetProportionAlertThreshold(v int32)`

SetProportionAlertThreshold sets ProportionAlertThreshold field to given value.

### HasProportionAlertThreshold

`func (o *MonitorConfiguration) HasProportionAlertThreshold() bool`

HasProportionAlertThreshold returns a boolean if a field has been set.

### GetQueries

`func (o *MonitorConfiguration) GetQueries() []MonitorAggregationQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *MonitorConfiguration) GetQueriesOk() (*[]MonitorAggregationQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *MonitorConfiguration) SetQueries(v []MonitorAggregationQuery)`

SetQueries sets Queries field to given value.


### GetCondition

`func (o *MonitorConfiguration) GetCondition() MonitorConfigurationAnomalyLogCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *MonitorConfiguration) GetConditionOk() (*MonitorConfigurationAnomalyLogCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *MonitorConfiguration) SetCondition(v MonitorConfigurationAnomalyLogCondition)`

SetCondition sets Condition field to given value.


### GetFilter

`func (o *MonitorConfiguration) GetFilter() MonitorConfigurationLogErrorPatternFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MonitorConfiguration) GetFilterOk() (*MonitorConfigurationLogErrorPatternFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MonitorConfiguration) SetFilter(v MonitorConfigurationLogErrorPatternFilter)`

SetFilter sets Filter field to given value.


### GetWarnBeforeInDays

`func (o *MonitorConfiguration) GetWarnBeforeInDays() int32`

GetWarnBeforeInDays returns the WarnBeforeInDays field if non-nil, zero value otherwise.

### GetWarnBeforeInDaysOk

`func (o *MonitorConfiguration) GetWarnBeforeInDaysOk() (*int32, bool)`

GetWarnBeforeInDaysOk returns a tuple with the WarnBeforeInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnBeforeInDays

`func (o *MonitorConfiguration) SetWarnBeforeInDays(v int32)`

SetWarnBeforeInDays sets WarnBeforeInDays field to given value.


### GetCloudAccounts

`func (o *MonitorConfiguration) GetCloudAccounts() []string`

GetCloudAccounts returns the CloudAccounts field if non-nil, zero value otherwise.

### GetCloudAccountsOk

`func (o *MonitorConfiguration) GetCloudAccountsOk() (*[]string, bool)`

GetCloudAccountsOk returns a tuple with the CloudAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccounts

`func (o *MonitorConfiguration) SetCloudAccounts(v []string)`

SetCloudAccounts sets CloudAccounts field to given value.

### HasCloudAccounts

`func (o *MonitorConfiguration) HasCloudAccounts() bool`

HasCloudAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


