# InputMonitorConfigurationK8sFaultyDeploymentFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterNames** | Pointer to **[]string** | Kubernetes cluster names to watch. Omit to watch every cluster. | [optional] 
**Namespaces** | Pointer to **[]string** | Kubernetes namespaces to watch. Omit to watch every namespace. | [optional] 
**Envs** | Pointer to **[]string** | Environments to watch. Omit to watch every environment. | [optional] 
**WorkloadTypes** | Pointer to **[]string** | Kubernetes workload types to watch. Omit to watch every workload type. | [optional] 

## Methods

### NewInputMonitorConfigurationK8sFaultyDeploymentFilter

`func NewInputMonitorConfigurationK8sFaultyDeploymentFilter() *InputMonitorConfigurationK8sFaultyDeploymentFilter`

NewInputMonitorConfigurationK8sFaultyDeploymentFilter instantiates a new InputMonitorConfigurationK8sFaultyDeploymentFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputMonitorConfigurationK8sFaultyDeploymentFilterWithDefaults

`func NewInputMonitorConfigurationK8sFaultyDeploymentFilterWithDefaults() *InputMonitorConfigurationK8sFaultyDeploymentFilter`

NewInputMonitorConfigurationK8sFaultyDeploymentFilterWithDefaults instantiates a new InputMonitorConfigurationK8sFaultyDeploymentFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterNames

`func (o *InputMonitorConfigurationK8sFaultyDeploymentFilter) GetClusterNames() []string`

GetClusterNames returns the ClusterNames field if non-nil, zero value otherwise.

### GetClusterNamesOk

`func (o *InputMonitorConfigurationK8sFaultyDeploymentFilter) GetClusterNamesOk() (*[]string, bool)`

GetClusterNamesOk returns a tuple with the ClusterNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNames

`func (o *InputMonitorConfigurationK8sFaultyDeploymentFilter) SetClusterNames(v []string)`

SetClusterNames sets ClusterNames field to given value.

### HasClusterNames

`func (o *InputMonitorConfigurationK8sFaultyDeploymentFilter) HasClusterNames() bool`

HasClusterNames returns a boolean if a field has been set.

### GetNamespaces

`func (o *InputMonitorConfigurationK8sFaultyDeploymentFilter) GetNamespaces() []string`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *InputMonitorConfigurationK8sFaultyDeploymentFilter) GetNamespacesOk() (*[]string, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *InputMonitorConfigurationK8sFaultyDeploymentFilter) SetNamespaces(v []string)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *InputMonitorConfigurationK8sFaultyDeploymentFilter) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetEnvs

`func (o *InputMonitorConfigurationK8sFaultyDeploymentFilter) GetEnvs() []string`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *InputMonitorConfigurationK8sFaultyDeploymentFilter) GetEnvsOk() (*[]string, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *InputMonitorConfigurationK8sFaultyDeploymentFilter) SetEnvs(v []string)`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *InputMonitorConfigurationK8sFaultyDeploymentFilter) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.

### GetWorkloadTypes

`func (o *InputMonitorConfigurationK8sFaultyDeploymentFilter) GetWorkloadTypes() []string`

GetWorkloadTypes returns the WorkloadTypes field if non-nil, zero value otherwise.

### GetWorkloadTypesOk

`func (o *InputMonitorConfigurationK8sFaultyDeploymentFilter) GetWorkloadTypesOk() (*[]string, bool)`

GetWorkloadTypesOk returns a tuple with the WorkloadTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadTypes

`func (o *InputMonitorConfigurationK8sFaultyDeploymentFilter) SetWorkloadTypes(v []string)`

SetWorkloadTypes sets WorkloadTypes field to given value.

### HasWorkloadTypes

`func (o *InputMonitorConfigurationK8sFaultyDeploymentFilter) HasWorkloadTypes() bool`

HasWorkloadTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


