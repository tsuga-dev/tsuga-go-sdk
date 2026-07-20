# MonitorConfigurationAnomalyLogCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Formula** | **string** | Formula result analyzed for anomalous behavior, usually &#x60;q1&#x60; or a formula alias. | 
**ConditionType** | **string** | Anomaly model type Tsuga uses to classify expected behavior: &#x60;rate&#x60; flags downward anomalies (e.g. count rates), &#x60;error&#x60; and &#x60;cpu&#x60; flag upward anomalies (e.g. error counts/rates, CPU usage), &#x60;general&#x60; flags both directions, and &#x60;to_be_set&#x60; is a placeholder Tsuga replaces by inferring the model type from the monitor query when the monitor is created or updated. | 

## Methods

### NewMonitorConfigurationAnomalyLogCondition

`func NewMonitorConfigurationAnomalyLogCondition(formula string, conditionType string, ) *MonitorConfigurationAnomalyLogCondition`

NewMonitorConfigurationAnomalyLogCondition instantiates a new MonitorConfigurationAnomalyLogCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorConfigurationAnomalyLogConditionWithDefaults

`func NewMonitorConfigurationAnomalyLogConditionWithDefaults() *MonitorConfigurationAnomalyLogCondition`

NewMonitorConfigurationAnomalyLogConditionWithDefaults instantiates a new MonitorConfigurationAnomalyLogCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormula

`func (o *MonitorConfigurationAnomalyLogCondition) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *MonitorConfigurationAnomalyLogCondition) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *MonitorConfigurationAnomalyLogCondition) SetFormula(v string)`

SetFormula sets Formula field to given value.


### GetConditionType

`func (o *MonitorConfigurationAnomalyLogCondition) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *MonitorConfigurationAnomalyLogCondition) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *MonitorConfigurationAnomalyLogCondition) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


