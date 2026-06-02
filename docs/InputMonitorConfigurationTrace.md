# InputMonitorConfigurationTrace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Conditions** | [**[]InputMonitorConfigurationMetricConditionsInner**](InputMonitorConfigurationMetricConditionsInner.md) |  | 
**NoDataBehavior** | **string** |  | 
**Timeframe** | **float32** | Timeframe of the monitor in minutes | 
**GroupByFields** | [**[]InputMonitorConfigurationMetricGroupByFieldsInner**](InputMonitorConfigurationMetricGroupByFieldsInner.md) | Monitor group by configuration. Warning! Note that the limit setting is currently ignored. | 
**AggregationAlertLogic** | Pointer to **string** |  | [optional] 
**ProportionAlertThreshold** | Pointer to **int32** |  | [optional] 
**Queries** | [**[]MonitorAggregationQuery1**](MonitorAggregationQuery1.md) | Aggregations that may be combined together in the same query | 

## Methods

### NewInputMonitorConfigurationTrace

`func NewInputMonitorConfigurationTrace(type_ string, conditions []InputMonitorConfigurationMetricConditionsInner, noDataBehavior string, timeframe float32, groupByFields []InputMonitorConfigurationMetricGroupByFieldsInner, queries []MonitorAggregationQuery1, ) *InputMonitorConfigurationTrace`

NewInputMonitorConfigurationTrace instantiates a new InputMonitorConfigurationTrace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputMonitorConfigurationTraceWithDefaults

`func NewInputMonitorConfigurationTraceWithDefaults() *InputMonitorConfigurationTrace`

NewInputMonitorConfigurationTraceWithDefaults instantiates a new InputMonitorConfigurationTrace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputMonitorConfigurationTrace) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputMonitorConfigurationTrace) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputMonitorConfigurationTrace) SetType(v string)`

SetType sets Type field to given value.


### GetConditions

`func (o *InputMonitorConfigurationTrace) GetConditions() []InputMonitorConfigurationMetricConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *InputMonitorConfigurationTrace) GetConditionsOk() (*[]InputMonitorConfigurationMetricConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *InputMonitorConfigurationTrace) SetConditions(v []InputMonitorConfigurationMetricConditionsInner)`

SetConditions sets Conditions field to given value.


### GetNoDataBehavior

`func (o *InputMonitorConfigurationTrace) GetNoDataBehavior() string`

GetNoDataBehavior returns the NoDataBehavior field if non-nil, zero value otherwise.

### GetNoDataBehaviorOk

`func (o *InputMonitorConfigurationTrace) GetNoDataBehaviorOk() (*string, bool)`

GetNoDataBehaviorOk returns a tuple with the NoDataBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataBehavior

`func (o *InputMonitorConfigurationTrace) SetNoDataBehavior(v string)`

SetNoDataBehavior sets NoDataBehavior field to given value.


### GetTimeframe

`func (o *InputMonitorConfigurationTrace) GetTimeframe() float32`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *InputMonitorConfigurationTrace) GetTimeframeOk() (*float32, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *InputMonitorConfigurationTrace) SetTimeframe(v float32)`

SetTimeframe sets Timeframe field to given value.


### GetGroupByFields

`func (o *InputMonitorConfigurationTrace) GetGroupByFields() []InputMonitorConfigurationMetricGroupByFieldsInner`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *InputMonitorConfigurationTrace) GetGroupByFieldsOk() (*[]InputMonitorConfigurationMetricGroupByFieldsInner, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *InputMonitorConfigurationTrace) SetGroupByFields(v []InputMonitorConfigurationMetricGroupByFieldsInner)`

SetGroupByFields sets GroupByFields field to given value.


### GetAggregationAlertLogic

`func (o *InputMonitorConfigurationTrace) GetAggregationAlertLogic() string`

GetAggregationAlertLogic returns the AggregationAlertLogic field if non-nil, zero value otherwise.

### GetAggregationAlertLogicOk

`func (o *InputMonitorConfigurationTrace) GetAggregationAlertLogicOk() (*string, bool)`

GetAggregationAlertLogicOk returns a tuple with the AggregationAlertLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationAlertLogic

`func (o *InputMonitorConfigurationTrace) SetAggregationAlertLogic(v string)`

SetAggregationAlertLogic sets AggregationAlertLogic field to given value.

### HasAggregationAlertLogic

`func (o *InputMonitorConfigurationTrace) HasAggregationAlertLogic() bool`

HasAggregationAlertLogic returns a boolean if a field has been set.

### GetProportionAlertThreshold

`func (o *InputMonitorConfigurationTrace) GetProportionAlertThreshold() int32`

GetProportionAlertThreshold returns the ProportionAlertThreshold field if non-nil, zero value otherwise.

### GetProportionAlertThresholdOk

`func (o *InputMonitorConfigurationTrace) GetProportionAlertThresholdOk() (*int32, bool)`

GetProportionAlertThresholdOk returns a tuple with the ProportionAlertThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProportionAlertThreshold

`func (o *InputMonitorConfigurationTrace) SetProportionAlertThreshold(v int32)`

SetProportionAlertThreshold sets ProportionAlertThreshold field to given value.

### HasProportionAlertThreshold

`func (o *InputMonitorConfigurationTrace) HasProportionAlertThreshold() bool`

HasProportionAlertThreshold returns a boolean if a field has been set.

### GetQueries

`func (o *InputMonitorConfigurationTrace) GetQueries() []MonitorAggregationQuery1`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *InputMonitorConfigurationTrace) GetQueriesOk() (*[]MonitorAggregationQuery1, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *InputMonitorConfigurationTrace) SetQueries(v []MonitorAggregationQuery1)`

SetQueries sets Queries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


