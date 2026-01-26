# MonitorConfigurationLogErrorPattern

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**AggregationAlertLogic** | **string** |  | 
**NoDataBehavior** | **string** |  | 
**Filter** | [**MonitorConfigurationLogErrorPatternFilter**](MonitorConfigurationLogErrorPatternFilter.md) |  | 

## Methods

### NewMonitorConfigurationLogErrorPattern

`func NewMonitorConfigurationLogErrorPattern(type_ string, aggregationAlertLogic string, noDataBehavior string, filter MonitorConfigurationLogErrorPatternFilter, ) *MonitorConfigurationLogErrorPattern`

NewMonitorConfigurationLogErrorPattern instantiates a new MonitorConfigurationLogErrorPattern object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorConfigurationLogErrorPatternWithDefaults

`func NewMonitorConfigurationLogErrorPatternWithDefaults() *MonitorConfigurationLogErrorPattern`

NewMonitorConfigurationLogErrorPatternWithDefaults instantiates a new MonitorConfigurationLogErrorPattern object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MonitorConfigurationLogErrorPattern) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorConfigurationLogErrorPattern) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorConfigurationLogErrorPattern) SetType(v string)`

SetType sets Type field to given value.


### GetAggregationAlertLogic

`func (o *MonitorConfigurationLogErrorPattern) GetAggregationAlertLogic() string`

GetAggregationAlertLogic returns the AggregationAlertLogic field if non-nil, zero value otherwise.

### GetAggregationAlertLogicOk

`func (o *MonitorConfigurationLogErrorPattern) GetAggregationAlertLogicOk() (*string, bool)`

GetAggregationAlertLogicOk returns a tuple with the AggregationAlertLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationAlertLogic

`func (o *MonitorConfigurationLogErrorPattern) SetAggregationAlertLogic(v string)`

SetAggregationAlertLogic sets AggregationAlertLogic field to given value.


### GetNoDataBehavior

`func (o *MonitorConfigurationLogErrorPattern) GetNoDataBehavior() string`

GetNoDataBehavior returns the NoDataBehavior field if non-nil, zero value otherwise.

### GetNoDataBehaviorOk

`func (o *MonitorConfigurationLogErrorPattern) GetNoDataBehaviorOk() (*string, bool)`

GetNoDataBehaviorOk returns a tuple with the NoDataBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataBehavior

`func (o *MonitorConfigurationLogErrorPattern) SetNoDataBehavior(v string)`

SetNoDataBehavior sets NoDataBehavior field to given value.


### GetFilter

`func (o *MonitorConfigurationLogErrorPattern) GetFilter() MonitorConfigurationLogErrorPatternFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MonitorConfigurationLogErrorPattern) GetFilterOk() (*MonitorConfigurationLogErrorPatternFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MonitorConfigurationLogErrorPattern) SetFilter(v MonitorConfigurationLogErrorPatternFilter)`

SetFilter sets Filter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


