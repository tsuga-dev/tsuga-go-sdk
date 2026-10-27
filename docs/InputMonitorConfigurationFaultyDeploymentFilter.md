# InputMonitorConfigurationFaultyDeploymentFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | Pointer to **[]string** | Service names to watch. Omit to watch every service. | [optional] 
**Envs** | Pointer to **[]string** | Environments to watch. Omit to watch every environment. | [optional] 

## Methods

### NewInputMonitorConfigurationFaultyDeploymentFilter

`func NewInputMonitorConfigurationFaultyDeploymentFilter() *InputMonitorConfigurationFaultyDeploymentFilter`

NewInputMonitorConfigurationFaultyDeploymentFilter instantiates a new InputMonitorConfigurationFaultyDeploymentFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputMonitorConfigurationFaultyDeploymentFilterWithDefaults

`func NewInputMonitorConfigurationFaultyDeploymentFilterWithDefaults() *InputMonitorConfigurationFaultyDeploymentFilter`

NewInputMonitorConfigurationFaultyDeploymentFilterWithDefaults instantiates a new InputMonitorConfigurationFaultyDeploymentFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *InputMonitorConfigurationFaultyDeploymentFilter) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *InputMonitorConfigurationFaultyDeploymentFilter) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *InputMonitorConfigurationFaultyDeploymentFilter) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *InputMonitorConfigurationFaultyDeploymentFilter) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetEnvs

`func (o *InputMonitorConfigurationFaultyDeploymentFilter) GetEnvs() []string`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *InputMonitorConfigurationFaultyDeploymentFilter) GetEnvsOk() (*[]string, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *InputMonitorConfigurationFaultyDeploymentFilter) SetEnvs(v []string)`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *InputMonitorConfigurationFaultyDeploymentFilter) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


