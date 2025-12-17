# MonitorConfigurationMetricCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Formula** | **string** |  | 
**Operator** | **string** |  | 
**Threshold** | **float32** |  | 

## Methods

### NewMonitorConfigurationMetricCondition

`func NewMonitorConfigurationMetricCondition(formula string, operator string, threshold float32, ) *MonitorConfigurationMetricCondition`

NewMonitorConfigurationMetricCondition instantiates a new MonitorConfigurationMetricCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorConfigurationMetricConditionWithDefaults

`func NewMonitorConfigurationMetricConditionWithDefaults() *MonitorConfigurationMetricCondition`

NewMonitorConfigurationMetricConditionWithDefaults instantiates a new MonitorConfigurationMetricCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormula

`func (o *MonitorConfigurationMetricCondition) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *MonitorConfigurationMetricCondition) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *MonitorConfigurationMetricCondition) SetFormula(v string)`

SetFormula sets Formula field to given value.


### GetOperator

`func (o *MonitorConfigurationMetricCondition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *MonitorConfigurationMetricCondition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *MonitorConfigurationMetricCondition) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetThreshold

`func (o *MonitorConfigurationMetricCondition) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *MonitorConfigurationMetricCondition) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *MonitorConfigurationMetricCondition) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


