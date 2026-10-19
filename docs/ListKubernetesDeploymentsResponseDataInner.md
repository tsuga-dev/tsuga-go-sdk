# ListKubernetesDeploymentsResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployment** | **string** | Kubernetes Deployment name from workload telemetry. | 
**Namespace** | Pointer to **string** | Kubernetes namespace that contains the Deployment. Present when namespace telemetry is observed. | [optional] 
**Cluster** | Pointer to **string** | Kubernetes cluster name that contains the Deployment. Present when cluster telemetry is observed. | [optional] 
**Team** | Pointer to **string** | Team metadata associated with the Deployment. Present when a team label or mapping is observed. | [optional] 
**Env** | Pointer to **string** | Environment metadata associated with the Deployment. Present when an environment label or mapping is observed. | [optional] 
**Pods** | Pointer to **int32** | Number of observed pods associated with the Deployment. Present when pod telemetry and workload attributes are observed. | [optional] 
**AvailableReplicas** | Pointer to **int32** | Number of available Deployment replicas from Deployment status telemetry. Present when deployment status is observed. | [optional] 
**DesiredReplicas** | Pointer to **int32** | Desired replica count configured for the Deployment. Present when Deployment spec telemetry is observed. | [optional] 
**CpuUsageMilli** | Pointer to **float32** | Recent aggregate CPU usage for observed Deployment pods, in millicores. Present when CPU usage telemetry is observed. | [optional] 
**MemoryUsageBytes** | Pointer to **float32** | Recent aggregate memory usage for observed Deployment pods, in bytes. Present when memory usage telemetry is observed. | [optional] 

## Methods

### NewListKubernetesDeploymentsResponseDataInner

`func NewListKubernetesDeploymentsResponseDataInner(deployment string, ) *ListKubernetesDeploymentsResponseDataInner`

NewListKubernetesDeploymentsResponseDataInner instantiates a new ListKubernetesDeploymentsResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKubernetesDeploymentsResponseDataInnerWithDefaults

`func NewListKubernetesDeploymentsResponseDataInnerWithDefaults() *ListKubernetesDeploymentsResponseDataInner`

NewListKubernetesDeploymentsResponseDataInnerWithDefaults instantiates a new ListKubernetesDeploymentsResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployment

`func (o *ListKubernetesDeploymentsResponseDataInner) GetDeployment() string`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *ListKubernetesDeploymentsResponseDataInner) GetDeploymentOk() (*string, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *ListKubernetesDeploymentsResponseDataInner) SetDeployment(v string)`

SetDeployment sets Deployment field to given value.


### GetNamespace

`func (o *ListKubernetesDeploymentsResponseDataInner) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ListKubernetesDeploymentsResponseDataInner) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ListKubernetesDeploymentsResponseDataInner) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ListKubernetesDeploymentsResponseDataInner) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetCluster

`func (o *ListKubernetesDeploymentsResponseDataInner) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ListKubernetesDeploymentsResponseDataInner) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ListKubernetesDeploymentsResponseDataInner) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ListKubernetesDeploymentsResponseDataInner) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetTeam

`func (o *ListKubernetesDeploymentsResponseDataInner) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ListKubernetesDeploymentsResponseDataInner) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ListKubernetesDeploymentsResponseDataInner) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ListKubernetesDeploymentsResponseDataInner) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetEnv

`func (o *ListKubernetesDeploymentsResponseDataInner) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ListKubernetesDeploymentsResponseDataInner) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ListKubernetesDeploymentsResponseDataInner) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ListKubernetesDeploymentsResponseDataInner) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetPods

`func (o *ListKubernetesDeploymentsResponseDataInner) GetPods() int32`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *ListKubernetesDeploymentsResponseDataInner) GetPodsOk() (*int32, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *ListKubernetesDeploymentsResponseDataInner) SetPods(v int32)`

SetPods sets Pods field to given value.

### HasPods

`func (o *ListKubernetesDeploymentsResponseDataInner) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetAvailableReplicas

`func (o *ListKubernetesDeploymentsResponseDataInner) GetAvailableReplicas() int32`

GetAvailableReplicas returns the AvailableReplicas field if non-nil, zero value otherwise.

### GetAvailableReplicasOk

`func (o *ListKubernetesDeploymentsResponseDataInner) GetAvailableReplicasOk() (*int32, bool)`

GetAvailableReplicasOk returns a tuple with the AvailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableReplicas

`func (o *ListKubernetesDeploymentsResponseDataInner) SetAvailableReplicas(v int32)`

SetAvailableReplicas sets AvailableReplicas field to given value.

### HasAvailableReplicas

`func (o *ListKubernetesDeploymentsResponseDataInner) HasAvailableReplicas() bool`

HasAvailableReplicas returns a boolean if a field has been set.

### GetDesiredReplicas

`func (o *ListKubernetesDeploymentsResponseDataInner) GetDesiredReplicas() int32`

GetDesiredReplicas returns the DesiredReplicas field if non-nil, zero value otherwise.

### GetDesiredReplicasOk

`func (o *ListKubernetesDeploymentsResponseDataInner) GetDesiredReplicasOk() (*int32, bool)`

GetDesiredReplicasOk returns a tuple with the DesiredReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredReplicas

`func (o *ListKubernetesDeploymentsResponseDataInner) SetDesiredReplicas(v int32)`

SetDesiredReplicas sets DesiredReplicas field to given value.

### HasDesiredReplicas

`func (o *ListKubernetesDeploymentsResponseDataInner) HasDesiredReplicas() bool`

HasDesiredReplicas returns a boolean if a field has been set.

### GetCpuUsageMilli

`func (o *ListKubernetesDeploymentsResponseDataInner) GetCpuUsageMilli() float32`

GetCpuUsageMilli returns the CpuUsageMilli field if non-nil, zero value otherwise.

### GetCpuUsageMilliOk

`func (o *ListKubernetesDeploymentsResponseDataInner) GetCpuUsageMilliOk() (*float32, bool)`

GetCpuUsageMilliOk returns a tuple with the CpuUsageMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageMilli

`func (o *ListKubernetesDeploymentsResponseDataInner) SetCpuUsageMilli(v float32)`

SetCpuUsageMilli sets CpuUsageMilli field to given value.

### HasCpuUsageMilli

`func (o *ListKubernetesDeploymentsResponseDataInner) HasCpuUsageMilli() bool`

HasCpuUsageMilli returns a boolean if a field has been set.

### GetMemoryUsageBytes

`func (o *ListKubernetesDeploymentsResponseDataInner) GetMemoryUsageBytes() float32`

GetMemoryUsageBytes returns the MemoryUsageBytes field if non-nil, zero value otherwise.

### GetMemoryUsageBytesOk

`func (o *ListKubernetesDeploymentsResponseDataInner) GetMemoryUsageBytesOk() (*float32, bool)`

GetMemoryUsageBytesOk returns a tuple with the MemoryUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsageBytes

`func (o *ListKubernetesDeploymentsResponseDataInner) SetMemoryUsageBytes(v float32)`

SetMemoryUsageBytes sets MemoryUsageBytes field to given value.

### HasMemoryUsageBytes

`func (o *ListKubernetesDeploymentsResponseDataInner) HasMemoryUsageBytes() bool`

HasMemoryUsageBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


