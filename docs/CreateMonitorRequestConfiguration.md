# CreateMonitorRequestConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Condition** | [**MonitorConfigurationAnomalyLogCondition**](MonitorConfigurationAnomalyLogCondition.md) |  | 
**NoDataBehavior** | **string** |  | 
**Timeframe** | **float32** | Timeframe of the monitor in minutes | 
**GroupByFields** | [**[]AggregationGroupBy**](AggregationGroupBy.md) |  | 
**AggregationAlertLogic** | Pointer to **string** |  | [optional] 
**ProportionAlertThreshold** | Pointer to **float32** |  | [optional] 
**Queries** | [**[]MonitorConfigurationMetricQueriesInner**](MonitorConfigurationMetricQueriesInner.md) |  | 

## Methods

### NewCreateMonitorRequestConfiguration

`func NewCreateMonitorRequestConfiguration(type_ string, condition MonitorConfigurationAnomalyLogCondition, noDataBehavior string, timeframe float32, groupByFields []AggregationGroupBy, queries []MonitorConfigurationMetricQueriesInner, ) *CreateMonitorRequestConfiguration`

NewCreateMonitorRequestConfiguration instantiates a new CreateMonitorRequestConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMonitorRequestConfigurationWithDefaults

`func NewCreateMonitorRequestConfigurationWithDefaults() *CreateMonitorRequestConfiguration`

NewCreateMonitorRequestConfigurationWithDefaults instantiates a new CreateMonitorRequestConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateMonitorRequestConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateMonitorRequestConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateMonitorRequestConfiguration) SetType(v string)`

SetType sets Type field to given value.


### GetCondition

`func (o *CreateMonitorRequestConfiguration) GetCondition() MonitorConfigurationAnomalyLogCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *CreateMonitorRequestConfiguration) GetConditionOk() (*MonitorConfigurationAnomalyLogCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *CreateMonitorRequestConfiguration) SetCondition(v MonitorConfigurationAnomalyLogCondition)`

SetCondition sets Condition field to given value.


### GetNoDataBehavior

`func (o *CreateMonitorRequestConfiguration) GetNoDataBehavior() string`

GetNoDataBehavior returns the NoDataBehavior field if non-nil, zero value otherwise.

### GetNoDataBehaviorOk

`func (o *CreateMonitorRequestConfiguration) GetNoDataBehaviorOk() (*string, bool)`

GetNoDataBehaviorOk returns a tuple with the NoDataBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataBehavior

`func (o *CreateMonitorRequestConfiguration) SetNoDataBehavior(v string)`

SetNoDataBehavior sets NoDataBehavior field to given value.


### GetTimeframe

`func (o *CreateMonitorRequestConfiguration) GetTimeframe() float32`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *CreateMonitorRequestConfiguration) GetTimeframeOk() (*float32, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *CreateMonitorRequestConfiguration) SetTimeframe(v float32)`

SetTimeframe sets Timeframe field to given value.


### GetGroupByFields

`func (o *CreateMonitorRequestConfiguration) GetGroupByFields() []AggregationGroupBy`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *CreateMonitorRequestConfiguration) GetGroupByFieldsOk() (*[]AggregationGroupBy, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *CreateMonitorRequestConfiguration) SetGroupByFields(v []AggregationGroupBy)`

SetGroupByFields sets GroupByFields field to given value.


### GetAggregationAlertLogic

`func (o *CreateMonitorRequestConfiguration) GetAggregationAlertLogic() string`

GetAggregationAlertLogic returns the AggregationAlertLogic field if non-nil, zero value otherwise.

### GetAggregationAlertLogicOk

`func (o *CreateMonitorRequestConfiguration) GetAggregationAlertLogicOk() (*string, bool)`

GetAggregationAlertLogicOk returns a tuple with the AggregationAlertLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationAlertLogic

`func (o *CreateMonitorRequestConfiguration) SetAggregationAlertLogic(v string)`

SetAggregationAlertLogic sets AggregationAlertLogic field to given value.

### HasAggregationAlertLogic

`func (o *CreateMonitorRequestConfiguration) HasAggregationAlertLogic() bool`

HasAggregationAlertLogic returns a boolean if a field has been set.

### GetProportionAlertThreshold

`func (o *CreateMonitorRequestConfiguration) GetProportionAlertThreshold() float32`

GetProportionAlertThreshold returns the ProportionAlertThreshold field if non-nil, zero value otherwise.

### GetProportionAlertThresholdOk

`func (o *CreateMonitorRequestConfiguration) GetProportionAlertThresholdOk() (*float32, bool)`

GetProportionAlertThresholdOk returns a tuple with the ProportionAlertThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProportionAlertThreshold

`func (o *CreateMonitorRequestConfiguration) SetProportionAlertThreshold(v float32)`

SetProportionAlertThreshold sets ProportionAlertThreshold field to given value.

### HasProportionAlertThreshold

`func (o *CreateMonitorRequestConfiguration) HasProportionAlertThreshold() bool`

HasProportionAlertThreshold returns a boolean if a field has been set.

### GetQueries

`func (o *CreateMonitorRequestConfiguration) GetQueries() []MonitorConfigurationMetricQueriesInner`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *CreateMonitorRequestConfiguration) GetQueriesOk() (*[]MonitorConfigurationMetricQueriesInner, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *CreateMonitorRequestConfiguration) SetQueries(v []MonitorConfigurationMetricQueriesInner)`

SetQueries sets Queries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


