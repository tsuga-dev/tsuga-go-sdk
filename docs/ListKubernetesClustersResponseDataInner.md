# ListKubernetesClustersResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **string** | Kubernetes cluster name from cluster telemetry. | 
**Teams** | Pointer to **[]string** | Team metadata values observed on pods in non-default Kubernetes namespaces on the cluster. Present when team labels or mappings are observed. | [optional] 
**Envs** | Pointer to **[]string** | Environment metadata values observed on pods in non-default Kubernetes namespaces on the cluster. Present when environment labels or mappings are observed. | [optional] 
**CpuUsageMilli** | Pointer to **float32** | Aggregate recent CPU usage for observed pods in the cluster, excluding default Kubernetes namespaces, in millicores. Present when CPU usage telemetry is observed. | [optional] 
**MemoryUsageBytes** | Pointer to **float32** | Aggregate recent memory usage for observed pods in the cluster, excluding default Kubernetes namespaces, in bytes. Present when memory usage telemetry is observed. | [optional] 
**ReadyNodes** | Pointer to **int32** | Number of Kubernetes nodes in the cluster that are ready. Present when node readiness telemetry is observed. | [optional] 
**PodPhaseCounts** | Pointer to [**ListKubernetesClustersResponseDataInnerPodPhaseCounts**](ListKubernetesClustersResponseDataInnerPodPhaseCounts.md) |  | [optional] 

## Methods

### NewListKubernetesClustersResponseDataInner

`func NewListKubernetesClustersResponseDataInner(cluster string, ) *ListKubernetesClustersResponseDataInner`

NewListKubernetesClustersResponseDataInner instantiates a new ListKubernetesClustersResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKubernetesClustersResponseDataInnerWithDefaults

`func NewListKubernetesClustersResponseDataInnerWithDefaults() *ListKubernetesClustersResponseDataInner`

NewListKubernetesClustersResponseDataInnerWithDefaults instantiates a new ListKubernetesClustersResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *ListKubernetesClustersResponseDataInner) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ListKubernetesClustersResponseDataInner) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ListKubernetesClustersResponseDataInner) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetTeams

`func (o *ListKubernetesClustersResponseDataInner) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *ListKubernetesClustersResponseDataInner) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *ListKubernetesClustersResponseDataInner) SetTeams(v []string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *ListKubernetesClustersResponseDataInner) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetEnvs

`func (o *ListKubernetesClustersResponseDataInner) GetEnvs() []string`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *ListKubernetesClustersResponseDataInner) GetEnvsOk() (*[]string, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *ListKubernetesClustersResponseDataInner) SetEnvs(v []string)`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *ListKubernetesClustersResponseDataInner) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.

### GetCpuUsageMilli

`func (o *ListKubernetesClustersResponseDataInner) GetCpuUsageMilli() float32`

GetCpuUsageMilli returns the CpuUsageMilli field if non-nil, zero value otherwise.

### GetCpuUsageMilliOk

`func (o *ListKubernetesClustersResponseDataInner) GetCpuUsageMilliOk() (*float32, bool)`

GetCpuUsageMilliOk returns a tuple with the CpuUsageMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageMilli

`func (o *ListKubernetesClustersResponseDataInner) SetCpuUsageMilli(v float32)`

SetCpuUsageMilli sets CpuUsageMilli field to given value.

### HasCpuUsageMilli

`func (o *ListKubernetesClustersResponseDataInner) HasCpuUsageMilli() bool`

HasCpuUsageMilli returns a boolean if a field has been set.

### GetMemoryUsageBytes

`func (o *ListKubernetesClustersResponseDataInner) GetMemoryUsageBytes() float32`

GetMemoryUsageBytes returns the MemoryUsageBytes field if non-nil, zero value otherwise.

### GetMemoryUsageBytesOk

`func (o *ListKubernetesClustersResponseDataInner) GetMemoryUsageBytesOk() (*float32, bool)`

GetMemoryUsageBytesOk returns a tuple with the MemoryUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsageBytes

`func (o *ListKubernetesClustersResponseDataInner) SetMemoryUsageBytes(v float32)`

SetMemoryUsageBytes sets MemoryUsageBytes field to given value.

### HasMemoryUsageBytes

`func (o *ListKubernetesClustersResponseDataInner) HasMemoryUsageBytes() bool`

HasMemoryUsageBytes returns a boolean if a field has been set.

### GetReadyNodes

`func (o *ListKubernetesClustersResponseDataInner) GetReadyNodes() int32`

GetReadyNodes returns the ReadyNodes field if non-nil, zero value otherwise.

### GetReadyNodesOk

`func (o *ListKubernetesClustersResponseDataInner) GetReadyNodesOk() (*int32, bool)`

GetReadyNodesOk returns a tuple with the ReadyNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyNodes

`func (o *ListKubernetesClustersResponseDataInner) SetReadyNodes(v int32)`

SetReadyNodes sets ReadyNodes field to given value.

### HasReadyNodes

`func (o *ListKubernetesClustersResponseDataInner) HasReadyNodes() bool`

HasReadyNodes returns a boolean if a field has been set.

### GetPodPhaseCounts

`func (o *ListKubernetesClustersResponseDataInner) GetPodPhaseCounts() ListKubernetesClustersResponseDataInnerPodPhaseCounts`

GetPodPhaseCounts returns the PodPhaseCounts field if non-nil, zero value otherwise.

### GetPodPhaseCountsOk

`func (o *ListKubernetesClustersResponseDataInner) GetPodPhaseCountsOk() (*ListKubernetesClustersResponseDataInnerPodPhaseCounts, bool)`

GetPodPhaseCountsOk returns a tuple with the PodPhaseCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPhaseCounts

`func (o *ListKubernetesClustersResponseDataInner) SetPodPhaseCounts(v ListKubernetesClustersResponseDataInnerPodPhaseCounts)`

SetPodPhaseCounts sets PodPhaseCounts field to given value.

### HasPodPhaseCounts

`func (o *ListKubernetesClustersResponseDataInner) HasPodPhaseCounts() bool`

HasPodPhaseCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


