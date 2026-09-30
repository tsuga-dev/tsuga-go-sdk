# InputMonitorConfigurationLogErrorPatternIncrease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Monitor that alerts when the volume of an existing log error pattern increases sharply for the configured teams and environment. | 
**AggregationAlertLogic** | **string** | Fixed aggregation logic for log error pattern increase monitors. Each pattern whose volume increased is evaluated as its own alert group. | 
**NoDataBehavior** | **string** | Fixed no-data behavior for log error pattern increase monitors. A pattern group resolves once its volume stops increasing. | 
**Filter** | [**InputMonitorConfigurationLogErrorPatternIncreaseFilter**](InputMonitorConfigurationLogErrorPatternIncreaseFilter.md) |  | 

## Methods

### NewInputMonitorConfigurationLogErrorPatternIncrease

`func NewInputMonitorConfigurationLogErrorPatternIncrease(type_ string, aggregationAlertLogic string, noDataBehavior string, filter InputMonitorConfigurationLogErrorPatternIncreaseFilter, ) *InputMonitorConfigurationLogErrorPatternIncrease`

NewInputMonitorConfigurationLogErrorPatternIncrease instantiates a new InputMonitorConfigurationLogErrorPatternIncrease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputMonitorConfigurationLogErrorPatternIncreaseWithDefaults

`func NewInputMonitorConfigurationLogErrorPatternIncreaseWithDefaults() *InputMonitorConfigurationLogErrorPatternIncrease`

NewInputMonitorConfigurationLogErrorPatternIncreaseWithDefaults instantiates a new InputMonitorConfigurationLogErrorPatternIncrease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputMonitorConfigurationLogErrorPatternIncrease) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputMonitorConfigurationLogErrorPatternIncrease) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputMonitorConfigurationLogErrorPatternIncrease) SetType(v string)`

SetType sets Type field to given value.


### GetAggregationAlertLogic

`func (o *InputMonitorConfigurationLogErrorPatternIncrease) GetAggregationAlertLogic() string`

GetAggregationAlertLogic returns the AggregationAlertLogic field if non-nil, zero value otherwise.

### GetAggregationAlertLogicOk

`func (o *InputMonitorConfigurationLogErrorPatternIncrease) GetAggregationAlertLogicOk() (*string, bool)`

GetAggregationAlertLogicOk returns a tuple with the AggregationAlertLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationAlertLogic

`func (o *InputMonitorConfigurationLogErrorPatternIncrease) SetAggregationAlertLogic(v string)`

SetAggregationAlertLogic sets AggregationAlertLogic field to given value.


### GetNoDataBehavior

`func (o *InputMonitorConfigurationLogErrorPatternIncrease) GetNoDataBehavior() string`

GetNoDataBehavior returns the NoDataBehavior field if non-nil, zero value otherwise.

### GetNoDataBehaviorOk

`func (o *InputMonitorConfigurationLogErrorPatternIncrease) GetNoDataBehaviorOk() (*string, bool)`

GetNoDataBehaviorOk returns a tuple with the NoDataBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataBehavior

`func (o *InputMonitorConfigurationLogErrorPatternIncrease) SetNoDataBehavior(v string)`

SetNoDataBehavior sets NoDataBehavior field to given value.


### GetFilter

`func (o *InputMonitorConfigurationLogErrorPatternIncrease) GetFilter() InputMonitorConfigurationLogErrorPatternIncreaseFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *InputMonitorConfigurationLogErrorPatternIncrease) GetFilterOk() (*InputMonitorConfigurationLogErrorPatternIncreaseFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *InputMonitorConfigurationLogErrorPatternIncrease) SetFilter(v InputMonitorConfigurationLogErrorPatternIncreaseFilter)`

SetFilter sets Filter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


