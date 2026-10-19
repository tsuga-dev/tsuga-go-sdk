# ListKubernetesDaemonSetsResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaemonSet** | **string** | Kubernetes DaemonSet name from workload telemetry. | 
**Namespace** | Pointer to **string** | Kubernetes namespace that contains the DaemonSet. Present when namespace telemetry is observed. | [optional] 
**Cluster** | Pointer to **string** | Kubernetes cluster name that contains the DaemonSet. Present when cluster telemetry is observed. | [optional] 
**Team** | Pointer to **string** | Team metadata associated with the DaemonSet. Present when a team label or mapping is observed. | [optional] 
**Env** | Pointer to **string** | Environment metadata associated with the DaemonSet. Present when an environment label or mapping is observed. | [optional] 
**Pods** | Pointer to **int32** | Number of observed pods associated with the DaemonSet. Present when pod telemetry and workload attributes are observed. | [optional] 
**ReadyNodes** | Pointer to **int32** | Number of eligible nodes with a ready DaemonSet pod. Present when DaemonSet status telemetry is observed. | [optional] 
**DesiredNodes** | Pointer to **int32** | Number of eligible nodes where the DaemonSet should run a pod. Present when DaemonSet status telemetry is observed. | [optional] 
**CpuUsageMilli** | Pointer to **float32** | Recent aggregate CPU usage for observed DaemonSet pods, in millicores. Present when CPU usage telemetry is observed. | [optional] 
**MemoryUsageBytes** | Pointer to **float32** | Recent aggregate memory usage for observed DaemonSet pods, in bytes. Present when memory usage telemetry is observed. | [optional] 

## Methods

### NewListKubernetesDaemonSetsResponseDataInner

`func NewListKubernetesDaemonSetsResponseDataInner(daemonSet string, ) *ListKubernetesDaemonSetsResponseDataInner`

NewListKubernetesDaemonSetsResponseDataInner instantiates a new ListKubernetesDaemonSetsResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKubernetesDaemonSetsResponseDataInnerWithDefaults

`func NewListKubernetesDaemonSetsResponseDataInnerWithDefaults() *ListKubernetesDaemonSetsResponseDataInner`

NewListKubernetesDaemonSetsResponseDataInnerWithDefaults instantiates a new ListKubernetesDaemonSetsResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaemonSet

`func (o *ListKubernetesDaemonSetsResponseDataInner) GetDaemonSet() string`

GetDaemonSet returns the DaemonSet field if non-nil, zero value otherwise.

### GetDaemonSetOk

`func (o *ListKubernetesDaemonSetsResponseDataInner) GetDaemonSetOk() (*string, bool)`

GetDaemonSetOk returns a tuple with the DaemonSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonSet

`func (o *ListKubernetesDaemonSetsResponseDataInner) SetDaemonSet(v string)`

SetDaemonSet sets DaemonSet field to given value.


### GetNamespace

`func (o *ListKubernetesDaemonSetsResponseDataInner) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ListKubernetesDaemonSetsResponseDataInner) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ListKubernetesDaemonSetsResponseDataInner) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ListKubernetesDaemonSetsResponseDataInner) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetCluster

`func (o *ListKubernetesDaemonSetsResponseDataInner) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ListKubernetesDaemonSetsResponseDataInner) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ListKubernetesDaemonSetsResponseDataInner) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ListKubernetesDaemonSetsResponseDataInner) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetTeam

`func (o *ListKubernetesDaemonSetsResponseDataInner) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ListKubernetesDaemonSetsResponseDataInner) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ListKubernetesDaemonSetsResponseDataInner) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ListKubernetesDaemonSetsResponseDataInner) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetEnv

`func (o *ListKubernetesDaemonSetsResponseDataInner) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ListKubernetesDaemonSetsResponseDataInner) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ListKubernetesDaemonSetsResponseDataInner) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ListKubernetesDaemonSetsResponseDataInner) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetPods

`func (o *ListKubernetesDaemonSetsResponseDataInner) GetPods() int32`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *ListKubernetesDaemonSetsResponseDataInner) GetPodsOk() (*int32, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *ListKubernetesDaemonSetsResponseDataInner) SetPods(v int32)`

SetPods sets Pods field to given value.

### HasPods

`func (o *ListKubernetesDaemonSetsResponseDataInner) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetReadyNodes

`func (o *ListKubernetesDaemonSetsResponseDataInner) GetReadyNodes() int32`

GetReadyNodes returns the ReadyNodes field if non-nil, zero value otherwise.

### GetReadyNodesOk

`func (o *ListKubernetesDaemonSetsResponseDataInner) GetReadyNodesOk() (*int32, bool)`

GetReadyNodesOk returns a tuple with the ReadyNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyNodes

`func (o *ListKubernetesDaemonSetsResponseDataInner) SetReadyNodes(v int32)`

SetReadyNodes sets ReadyNodes field to given value.

### HasReadyNodes

`func (o *ListKubernetesDaemonSetsResponseDataInner) HasReadyNodes() bool`

HasReadyNodes returns a boolean if a field has been set.

### GetDesiredNodes

`func (o *ListKubernetesDaemonSetsResponseDataInner) GetDesiredNodes() int32`

GetDesiredNodes returns the DesiredNodes field if non-nil, zero value otherwise.

### GetDesiredNodesOk

`func (o *ListKubernetesDaemonSetsResponseDataInner) GetDesiredNodesOk() (*int32, bool)`

GetDesiredNodesOk returns a tuple with the DesiredNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredNodes

`func (o *ListKubernetesDaemonSetsResponseDataInner) SetDesiredNodes(v int32)`

SetDesiredNodes sets DesiredNodes field to given value.

### HasDesiredNodes

`func (o *ListKubernetesDaemonSetsResponseDataInner) HasDesiredNodes() bool`

HasDesiredNodes returns a boolean if a field has been set.

### GetCpuUsageMilli

`func (o *ListKubernetesDaemonSetsResponseDataInner) GetCpuUsageMilli() float32`

GetCpuUsageMilli returns the CpuUsageMilli field if non-nil, zero value otherwise.

### GetCpuUsageMilliOk

`func (o *ListKubernetesDaemonSetsResponseDataInner) GetCpuUsageMilliOk() (*float32, bool)`

GetCpuUsageMilliOk returns a tuple with the CpuUsageMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageMilli

`func (o *ListKubernetesDaemonSetsResponseDataInner) SetCpuUsageMilli(v float32)`

SetCpuUsageMilli sets CpuUsageMilli field to given value.

### HasCpuUsageMilli

`func (o *ListKubernetesDaemonSetsResponseDataInner) HasCpuUsageMilli() bool`

HasCpuUsageMilli returns a boolean if a field has been set.

### GetMemoryUsageBytes

`func (o *ListKubernetesDaemonSetsResponseDataInner) GetMemoryUsageBytes() float32`

GetMemoryUsageBytes returns the MemoryUsageBytes field if non-nil, zero value otherwise.

### GetMemoryUsageBytesOk

`func (o *ListKubernetesDaemonSetsResponseDataInner) GetMemoryUsageBytesOk() (*float32, bool)`

GetMemoryUsageBytesOk returns a tuple with the MemoryUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsageBytes

`func (o *ListKubernetesDaemonSetsResponseDataInner) SetMemoryUsageBytes(v float32)`

SetMemoryUsageBytes sets MemoryUsageBytes field to given value.

### HasMemoryUsageBytes

`func (o *ListKubernetesDaemonSetsResponseDataInner) HasMemoryUsageBytes() bool`

HasMemoryUsageBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


