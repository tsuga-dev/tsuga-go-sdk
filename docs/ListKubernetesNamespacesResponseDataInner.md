# ListKubernetesNamespacesResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** | Kubernetes namespace name from namespace telemetry. Includes namespaces with no observed pods when namespace telemetry is observed. | 
**Cluster** | Pointer to **string** | Kubernetes cluster name that contains the namespace. Present when cluster telemetry is observed. | [optional] 
**Teams** | Pointer to **[]string** | Team metadata values observed in workloads in the namespace. Present when team labels or mappings are observed. | [optional] 
**Envs** | Pointer to **[]string** | Environment metadata values observed in workloads in the namespace. Present when environment labels or mappings are observed. | [optional] 
**Pods** | Pointer to **int32** | Number of observed pods in the namespace. Present when pod telemetry is observed. | [optional] 
**CpuUsageMilli** | Pointer to **float32** | Aggregate recent CPU usage for observed pods in the namespace, in millicores. Present when CPU usage telemetry is observed. | [optional] 
**MemoryUsageBytes** | Pointer to **float32** | Aggregate recent memory usage for observed pods in the namespace, in bytes. Present when memory usage telemetry is observed. | [optional] 

## Methods

### NewListKubernetesNamespacesResponseDataInner

`func NewListKubernetesNamespacesResponseDataInner(namespace string, ) *ListKubernetesNamespacesResponseDataInner`

NewListKubernetesNamespacesResponseDataInner instantiates a new ListKubernetesNamespacesResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKubernetesNamespacesResponseDataInnerWithDefaults

`func NewListKubernetesNamespacesResponseDataInnerWithDefaults() *ListKubernetesNamespacesResponseDataInner`

NewListKubernetesNamespacesResponseDataInnerWithDefaults instantiates a new ListKubernetesNamespacesResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *ListKubernetesNamespacesResponseDataInner) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ListKubernetesNamespacesResponseDataInner) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ListKubernetesNamespacesResponseDataInner) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetCluster

`func (o *ListKubernetesNamespacesResponseDataInner) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ListKubernetesNamespacesResponseDataInner) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ListKubernetesNamespacesResponseDataInner) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ListKubernetesNamespacesResponseDataInner) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetTeams

`func (o *ListKubernetesNamespacesResponseDataInner) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *ListKubernetesNamespacesResponseDataInner) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *ListKubernetesNamespacesResponseDataInner) SetTeams(v []string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *ListKubernetesNamespacesResponseDataInner) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetEnvs

`func (o *ListKubernetesNamespacesResponseDataInner) GetEnvs() []string`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *ListKubernetesNamespacesResponseDataInner) GetEnvsOk() (*[]string, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *ListKubernetesNamespacesResponseDataInner) SetEnvs(v []string)`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *ListKubernetesNamespacesResponseDataInner) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.

### GetPods

`func (o *ListKubernetesNamespacesResponseDataInner) GetPods() int32`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *ListKubernetesNamespacesResponseDataInner) GetPodsOk() (*int32, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *ListKubernetesNamespacesResponseDataInner) SetPods(v int32)`

SetPods sets Pods field to given value.

### HasPods

`func (o *ListKubernetesNamespacesResponseDataInner) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetCpuUsageMilli

`func (o *ListKubernetesNamespacesResponseDataInner) GetCpuUsageMilli() float32`

GetCpuUsageMilli returns the CpuUsageMilli field if non-nil, zero value otherwise.

### GetCpuUsageMilliOk

`func (o *ListKubernetesNamespacesResponseDataInner) GetCpuUsageMilliOk() (*float32, bool)`

GetCpuUsageMilliOk returns a tuple with the CpuUsageMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageMilli

`func (o *ListKubernetesNamespacesResponseDataInner) SetCpuUsageMilli(v float32)`

SetCpuUsageMilli sets CpuUsageMilli field to given value.

### HasCpuUsageMilli

`func (o *ListKubernetesNamespacesResponseDataInner) HasCpuUsageMilli() bool`

HasCpuUsageMilli returns a boolean if a field has been set.

### GetMemoryUsageBytes

`func (o *ListKubernetesNamespacesResponseDataInner) GetMemoryUsageBytes() float32`

GetMemoryUsageBytes returns the MemoryUsageBytes field if non-nil, zero value otherwise.

### GetMemoryUsageBytesOk

`func (o *ListKubernetesNamespacesResponseDataInner) GetMemoryUsageBytesOk() (*float32, bool)`

GetMemoryUsageBytesOk returns a tuple with the MemoryUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsageBytes

`func (o *ListKubernetesNamespacesResponseDataInner) SetMemoryUsageBytes(v float32)`

SetMemoryUsageBytes sets MemoryUsageBytes field to given value.

### HasMemoryUsageBytes

`func (o *ListKubernetesNamespacesResponseDataInner) HasMemoryUsageBytes() bool`

HasMemoryUsageBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


