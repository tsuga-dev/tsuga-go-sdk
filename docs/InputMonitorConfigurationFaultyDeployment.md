# InputMonitorConfigurationFaultyDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Monitor that alerts when Tsuga faulty deployment detection flags a newly deployed service version as faulty. | 
**AggregationAlertLogic** | **string** | Fixed aggregation logic for faulty deployment monitors. Each faulty service version in an environment is evaluated as its own alert group. | 
**Filter** | [**InputMonitorConfigurationFaultyDeploymentFilter**](InputMonitorConfigurationFaultyDeploymentFilter.md) |  | 

## Methods

### NewInputMonitorConfigurationFaultyDeployment

`func NewInputMonitorConfigurationFaultyDeployment(type_ string, aggregationAlertLogic string, filter InputMonitorConfigurationFaultyDeploymentFilter, ) *InputMonitorConfigurationFaultyDeployment`

NewInputMonitorConfigurationFaultyDeployment instantiates a new InputMonitorConfigurationFaultyDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputMonitorConfigurationFaultyDeploymentWithDefaults

`func NewInputMonitorConfigurationFaultyDeploymentWithDefaults() *InputMonitorConfigurationFaultyDeployment`

NewInputMonitorConfigurationFaultyDeploymentWithDefaults instantiates a new InputMonitorConfigurationFaultyDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputMonitorConfigurationFaultyDeployment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputMonitorConfigurationFaultyDeployment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputMonitorConfigurationFaultyDeployment) SetType(v string)`

SetType sets Type field to given value.


### GetAggregationAlertLogic

`func (o *InputMonitorConfigurationFaultyDeployment) GetAggregationAlertLogic() string`

GetAggregationAlertLogic returns the AggregationAlertLogic field if non-nil, zero value otherwise.

### GetAggregationAlertLogicOk

`func (o *InputMonitorConfigurationFaultyDeployment) GetAggregationAlertLogicOk() (*string, bool)`

GetAggregationAlertLogicOk returns a tuple with the AggregationAlertLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationAlertLogic

`func (o *InputMonitorConfigurationFaultyDeployment) SetAggregationAlertLogic(v string)`

SetAggregationAlertLogic sets AggregationAlertLogic field to given value.


### GetFilter

`func (o *InputMonitorConfigurationFaultyDeployment) GetFilter() InputMonitorConfigurationFaultyDeploymentFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *InputMonitorConfigurationFaultyDeployment) GetFilterOk() (*InputMonitorConfigurationFaultyDeploymentFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *InputMonitorConfigurationFaultyDeployment) SetFilter(v InputMonitorConfigurationFaultyDeploymentFilter)`

SetFilter sets Filter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


