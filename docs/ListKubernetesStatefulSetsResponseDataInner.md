# ListKubernetesStatefulSetsResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatefulSet** | **string** | Kubernetes StatefulSet name from workload telemetry. | 
**Namespace** | Pointer to **string** | Kubernetes namespace that contains the StatefulSet. Present when namespace telemetry is observed. | [optional] 
**Cluster** | Pointer to **string** | Kubernetes cluster name that contains the StatefulSet. Present when cluster telemetry is observed. | [optional] 
**Team** | Pointer to **string** | Team metadata associated with the StatefulSet. Present when a team label or mapping is observed. | [optional] 
**Env** | Pointer to **string** | Environment metadata associated with the StatefulSet. Present when an environment label or mapping is observed. | [optional] 
**Pods** | Pointer to **int32** | Number of observed pods associated with the StatefulSet. Present when pod telemetry and workload attributes are observed. | [optional] 
**CpuUsageMilli** | Pointer to **float32** | Recent aggregate CPU usage for observed StatefulSet pods, in millicores. Present when CPU usage telemetry is observed. | [optional] 
**MemoryUsageBytes** | Pointer to **float32** | Recent aggregate memory usage for observed StatefulSet pods, in bytes. Present when memory usage telemetry is observed. | [optional] 

## Methods

### NewListKubernetesStatefulSetsResponseDataInner

`func NewListKubernetesStatefulSetsResponseDataInner(statefulSet string, ) *ListKubernetesStatefulSetsResponseDataInner`

NewListKubernetesStatefulSetsResponseDataInner instantiates a new ListKubernetesStatefulSetsResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKubernetesStatefulSetsResponseDataInnerWithDefaults

`func NewListKubernetesStatefulSetsResponseDataInnerWithDefaults() *ListKubernetesStatefulSetsResponseDataInner`

NewListKubernetesStatefulSetsResponseDataInnerWithDefaults instantiates a new ListKubernetesStatefulSetsResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatefulSet

`func (o *ListKubernetesStatefulSetsResponseDataInner) GetStatefulSet() string`

GetStatefulSet returns the StatefulSet field if non-nil, zero value otherwise.

### GetStatefulSetOk

`func (o *ListKubernetesStatefulSetsResponseDataInner) GetStatefulSetOk() (*string, bool)`

GetStatefulSetOk returns a tuple with the StatefulSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatefulSet

`func (o *ListKubernetesStatefulSetsResponseDataInner) SetStatefulSet(v string)`

SetStatefulSet sets StatefulSet field to given value.


### GetNamespace

`func (o *ListKubernetesStatefulSetsResponseDataInner) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ListKubernetesStatefulSetsResponseDataInner) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ListKubernetesStatefulSetsResponseDataInner) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ListKubernetesStatefulSetsResponseDataInner) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetCluster

`func (o *ListKubernetesStatefulSetsResponseDataInner) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ListKubernetesStatefulSetsResponseDataInner) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ListKubernetesStatefulSetsResponseDataInner) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ListKubernetesStatefulSetsResponseDataInner) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetTeam

`func (o *ListKubernetesStatefulSetsResponseDataInner) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ListKubernetesStatefulSetsResponseDataInner) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ListKubernetesStatefulSetsResponseDataInner) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ListKubernetesStatefulSetsResponseDataInner) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetEnv

`func (o *ListKubernetesStatefulSetsResponseDataInner) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ListKubernetesStatefulSetsResponseDataInner) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ListKubernetesStatefulSetsResponseDataInner) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ListKubernetesStatefulSetsResponseDataInner) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetPods

`func (o *ListKubernetesStatefulSetsResponseDataInner) GetPods() int32`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *ListKubernetesStatefulSetsResponseDataInner) GetPodsOk() (*int32, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *ListKubernetesStatefulSetsResponseDataInner) SetPods(v int32)`

SetPods sets Pods field to given value.

### HasPods

`func (o *ListKubernetesStatefulSetsResponseDataInner) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetCpuUsageMilli

`func (o *ListKubernetesStatefulSetsResponseDataInner) GetCpuUsageMilli() float32`

GetCpuUsageMilli returns the CpuUsageMilli field if non-nil, zero value otherwise.

### GetCpuUsageMilliOk

`func (o *ListKubernetesStatefulSetsResponseDataInner) GetCpuUsageMilliOk() (*float32, bool)`

GetCpuUsageMilliOk returns a tuple with the CpuUsageMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageMilli

`func (o *ListKubernetesStatefulSetsResponseDataInner) SetCpuUsageMilli(v float32)`

SetCpuUsageMilli sets CpuUsageMilli field to given value.

### HasCpuUsageMilli

`func (o *ListKubernetesStatefulSetsResponseDataInner) HasCpuUsageMilli() bool`

HasCpuUsageMilli returns a boolean if a field has been set.

### GetMemoryUsageBytes

`func (o *ListKubernetesStatefulSetsResponseDataInner) GetMemoryUsageBytes() float32`

GetMemoryUsageBytes returns the MemoryUsageBytes field if non-nil, zero value otherwise.

### GetMemoryUsageBytesOk

`func (o *ListKubernetesStatefulSetsResponseDataInner) GetMemoryUsageBytesOk() (*float32, bool)`

GetMemoryUsageBytesOk returns a tuple with the MemoryUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsageBytes

`func (o *ListKubernetesStatefulSetsResponseDataInner) SetMemoryUsageBytes(v float32)`

SetMemoryUsageBytes sets MemoryUsageBytes field to given value.

### HasMemoryUsageBytes

`func (o *ListKubernetesStatefulSetsResponseDataInner) HasMemoryUsageBytes() bool`

HasMemoryUsageBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


