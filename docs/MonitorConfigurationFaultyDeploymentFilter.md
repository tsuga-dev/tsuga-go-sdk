# MonitorConfigurationFaultyDeploymentFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | Pointer to **[]string** | Service names to watch. Omit to watch every service. | [optional] 
**Envs** | Pointer to **[]string** | Environments to watch. Omit to watch every environment. | [optional] 

## Methods

### NewMonitorConfigurationFaultyDeploymentFilter

`func NewMonitorConfigurationFaultyDeploymentFilter() *MonitorConfigurationFaultyDeploymentFilter`

NewMonitorConfigurationFaultyDeploymentFilter instantiates a new MonitorConfigurationFaultyDeploymentFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorConfigurationFaultyDeploymentFilterWithDefaults

`func NewMonitorConfigurationFaultyDeploymentFilterWithDefaults() *MonitorConfigurationFaultyDeploymentFilter`

NewMonitorConfigurationFaultyDeploymentFilterWithDefaults instantiates a new MonitorConfigurationFaultyDeploymentFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *MonitorConfigurationFaultyDeploymentFilter) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *MonitorConfigurationFaultyDeploymentFilter) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *MonitorConfigurationFaultyDeploymentFilter) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *MonitorConfigurationFaultyDeploymentFilter) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetEnvs

`func (o *MonitorConfigurationFaultyDeploymentFilter) GetEnvs() []string`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *MonitorConfigurationFaultyDeploymentFilter) GetEnvsOk() (*[]string, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *MonitorConfigurationFaultyDeploymentFilter) SetEnvs(v []string)`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *MonitorConfigurationFaultyDeploymentFilter) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


