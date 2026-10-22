# MonitorConfigurationK8sFaultyDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Monitor that alerts when a Kubernetes workload (Deployment, StatefulSet, or DaemonSet) keeps failing right after a rollout. | 
**AggregationAlertLogic** | **string** | Fixed aggregation logic for Kubernetes faulty deployment monitors. Each faulty workload revision is evaluated as its own alert group. | 
**Filter** | [**MonitorConfigurationK8sFaultyDeploymentFilter**](MonitorConfigurationK8sFaultyDeploymentFilter.md) |  | 

## Methods

### NewMonitorConfigurationK8sFaultyDeployment

`func NewMonitorConfigurationK8sFaultyDeployment(type_ string, aggregationAlertLogic string, filter MonitorConfigurationK8sFaultyDeploymentFilter, ) *MonitorConfigurationK8sFaultyDeployment`

NewMonitorConfigurationK8sFaultyDeployment instantiates a new MonitorConfigurationK8sFaultyDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorConfigurationK8sFaultyDeploymentWithDefaults

`func NewMonitorConfigurationK8sFaultyDeploymentWithDefaults() *MonitorConfigurationK8sFaultyDeployment`

NewMonitorConfigurationK8sFaultyDeploymentWithDefaults instantiates a new MonitorConfigurationK8sFaultyDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MonitorConfigurationK8sFaultyDeployment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorConfigurationK8sFaultyDeployment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorConfigurationK8sFaultyDeployment) SetType(v string)`

SetType sets Type field to given value.


### GetAggregationAlertLogic

`func (o *MonitorConfigurationK8sFaultyDeployment) GetAggregationAlertLogic() string`

GetAggregationAlertLogic returns the AggregationAlertLogic field if non-nil, zero value otherwise.

### GetAggregationAlertLogicOk

`func (o *MonitorConfigurationK8sFaultyDeployment) GetAggregationAlertLogicOk() (*string, bool)`

GetAggregationAlertLogicOk returns a tuple with the AggregationAlertLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationAlertLogic

`func (o *MonitorConfigurationK8sFaultyDeployment) SetAggregationAlertLogic(v string)`

SetAggregationAlertLogic sets AggregationAlertLogic field to given value.


### GetFilter

`func (o *MonitorConfigurationK8sFaultyDeployment) GetFilter() MonitorConfigurationK8sFaultyDeploymentFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MonitorConfigurationK8sFaultyDeployment) GetFilterOk() (*MonitorConfigurationK8sFaultyDeploymentFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MonitorConfigurationK8sFaultyDeployment) SetFilter(v MonitorConfigurationK8sFaultyDeploymentFilter)`

SetFilter sets Filter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


