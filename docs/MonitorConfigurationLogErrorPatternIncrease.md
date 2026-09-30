# MonitorConfigurationLogErrorPatternIncrease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Monitor that alerts when the volume of an existing log error pattern increases sharply for the configured teams and environment. | 
**AggregationAlertLogic** | **string** | Fixed aggregation logic for log error pattern increase monitors. Each pattern whose volume increased is evaluated as its own alert group. | 
**NoDataBehavior** | **string** | Fixed no-data behavior for log error pattern increase monitors. A pattern group resolves once its volume stops increasing. | 
**Filter** | [**MonitorConfigurationLogErrorPatternIncreaseFilter**](MonitorConfigurationLogErrorPatternIncreaseFilter.md) |  | 

## Methods

### NewMonitorConfigurationLogErrorPatternIncrease

`func NewMonitorConfigurationLogErrorPatternIncrease(type_ string, aggregationAlertLogic string, noDataBehavior string, filter MonitorConfigurationLogErrorPatternIncreaseFilter, ) *MonitorConfigurationLogErrorPatternIncrease`

NewMonitorConfigurationLogErrorPatternIncrease instantiates a new MonitorConfigurationLogErrorPatternIncrease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorConfigurationLogErrorPatternIncreaseWithDefaults

`func NewMonitorConfigurationLogErrorPatternIncreaseWithDefaults() *MonitorConfigurationLogErrorPatternIncrease`

NewMonitorConfigurationLogErrorPatternIncreaseWithDefaults instantiates a new MonitorConfigurationLogErrorPatternIncrease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MonitorConfigurationLogErrorPatternIncrease) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorConfigurationLogErrorPatternIncrease) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorConfigurationLogErrorPatternIncrease) SetType(v string)`

SetType sets Type field to given value.


### GetAggregationAlertLogic

`func (o *MonitorConfigurationLogErrorPatternIncrease) GetAggregationAlertLogic() string`

GetAggregationAlertLogic returns the AggregationAlertLogic field if non-nil, zero value otherwise.

### GetAggregationAlertLogicOk

`func (o *MonitorConfigurationLogErrorPatternIncrease) GetAggregationAlertLogicOk() (*string, bool)`

GetAggregationAlertLogicOk returns a tuple with the AggregationAlertLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationAlertLogic

`func (o *MonitorConfigurationLogErrorPatternIncrease) SetAggregationAlertLogic(v string)`

SetAggregationAlertLogic sets AggregationAlertLogic field to given value.


### GetNoDataBehavior

`func (o *MonitorConfigurationLogErrorPatternIncrease) GetNoDataBehavior() string`

GetNoDataBehavior returns the NoDataBehavior field if non-nil, zero value otherwise.

### GetNoDataBehaviorOk

`func (o *MonitorConfigurationLogErrorPatternIncrease) GetNoDataBehaviorOk() (*string, bool)`

GetNoDataBehaviorOk returns a tuple with the NoDataBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataBehavior

`func (o *MonitorConfigurationLogErrorPatternIncrease) SetNoDataBehavior(v string)`

SetNoDataBehavior sets NoDataBehavior field to given value.


### GetFilter

`func (o *MonitorConfigurationLogErrorPatternIncrease) GetFilter() MonitorConfigurationLogErrorPatternIncreaseFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MonitorConfigurationLogErrorPatternIncrease) GetFilterOk() (*MonitorConfigurationLogErrorPatternIncreaseFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MonitorConfigurationLogErrorPatternIncrease) SetFilter(v MonitorConfigurationLogErrorPatternIncreaseFilter)`

SetFilter sets Filter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


