# InputMonitorConfigurationLogErrorPattern

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Monitor that alerts on newly detected log error patterns for the configured team, environment, and optional service. Recovery notifications are not sent for this monitor type. | 
**AggregationAlertLogic** | **string** | Fixed aggregation logic for log error pattern monitors. Each new pattern is evaluated as its own alert group. | 
**NoDataBehavior** | **string** | Fixed no-data behavior for log error pattern monitors. Pattern state is kept until inactive pattern groups expire. | 
**Filter** | [**InputMonitorConfigurationLogErrorPatternFilter**](InputMonitorConfigurationLogErrorPatternFilter.md) |  | 

## Methods

### NewInputMonitorConfigurationLogErrorPattern

`func NewInputMonitorConfigurationLogErrorPattern(type_ string, aggregationAlertLogic string, noDataBehavior string, filter InputMonitorConfigurationLogErrorPatternFilter, ) *InputMonitorConfigurationLogErrorPattern`

NewInputMonitorConfigurationLogErrorPattern instantiates a new InputMonitorConfigurationLogErrorPattern object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputMonitorConfigurationLogErrorPatternWithDefaults

`func NewInputMonitorConfigurationLogErrorPatternWithDefaults() *InputMonitorConfigurationLogErrorPattern`

NewInputMonitorConfigurationLogErrorPatternWithDefaults instantiates a new InputMonitorConfigurationLogErrorPattern object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputMonitorConfigurationLogErrorPattern) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputMonitorConfigurationLogErrorPattern) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputMonitorConfigurationLogErrorPattern) SetType(v string)`

SetType sets Type field to given value.


### GetAggregationAlertLogic

`func (o *InputMonitorConfigurationLogErrorPattern) GetAggregationAlertLogic() string`

GetAggregationAlertLogic returns the AggregationAlertLogic field if non-nil, zero value otherwise.

### GetAggregationAlertLogicOk

`func (o *InputMonitorConfigurationLogErrorPattern) GetAggregationAlertLogicOk() (*string, bool)`

GetAggregationAlertLogicOk returns a tuple with the AggregationAlertLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationAlertLogic

`func (o *InputMonitorConfigurationLogErrorPattern) SetAggregationAlertLogic(v string)`

SetAggregationAlertLogic sets AggregationAlertLogic field to given value.


### GetNoDataBehavior

`func (o *InputMonitorConfigurationLogErrorPattern) GetNoDataBehavior() string`

GetNoDataBehavior returns the NoDataBehavior field if non-nil, zero value otherwise.

### GetNoDataBehaviorOk

`func (o *InputMonitorConfigurationLogErrorPattern) GetNoDataBehaviorOk() (*string, bool)`

GetNoDataBehaviorOk returns a tuple with the NoDataBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataBehavior

`func (o *InputMonitorConfigurationLogErrorPattern) SetNoDataBehavior(v string)`

SetNoDataBehavior sets NoDataBehavior field to given value.


### GetFilter

`func (o *InputMonitorConfigurationLogErrorPattern) GetFilter() InputMonitorConfigurationLogErrorPatternFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *InputMonitorConfigurationLogErrorPattern) GetFilterOk() (*InputMonitorConfigurationLogErrorPatternFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *InputMonitorConfigurationLogErrorPattern) SetFilter(v InputMonitorConfigurationLogErrorPatternFilter)`

SetFilter sets Filter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


