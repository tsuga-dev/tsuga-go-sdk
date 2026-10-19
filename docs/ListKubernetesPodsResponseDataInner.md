# ListKubernetesPodsResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PodName** | **string** | Name of the Kubernetes pod. Set when the pod is created. | 
**Namespace** | Pointer to **string** | Kubernetes namespace that contains the pod. Present when namespace telemetry is observed. | [optional] 
**Cluster** | Pointer to **string** | Kubernetes cluster name that reported the pod. Present when cluster telemetry is observed. | [optional] 
**Team** | Pointer to **string** | Team metadata associated with the pod. Present when a team label or mapping is observed. | [optional] 
**Env** | Pointer to **string** | Environment metadata associated with the pod, such as production or staging. Present when an environment label or mapping is observed. | [optional] 
**Node** | Pointer to **string** | Kubernetes node currently or most recently running the pod. Present when node telemetry is observed. | [optional] 
**Phase** | **string** | Current or most recently observed pod status. From pod status and container state telemetry. Values include the Kubernetes phases Pending, Running, Succeeded, Failed, and Unknown, plus Tsuga-derived CrashLoopBackOff and Error values. | 
**ReadyContainers** | **int32** | Number of containers in the pod that are currently ready. From container readiness telemetry. Zero or greater. | 
**TotalContainers** | **int32** | Number of observed containers for the pod based on container readiness telemetry. Zero or greater. | 
**Restarts** | **int32** | Total restart count across containers in the pod. From container restart telemetry. Zero or greater. | 
**CreatedAt** | Pointer to **string** | Timestamp from pod start time telemetry for the selected pod instance. Present when pod start time telemetry is observed. | [optional] 
**Workload** | Pointer to [**ListKubernetesPodsResponseDataInnerWorkload**](ListKubernetesPodsResponseDataInnerWorkload.md) |  | [optional] 
**CpuUsageMilli** | Pointer to **float32** | Recent pod CPU usage in millicores. Present when CPU usage telemetry is observed. | [optional] 
**CpuLimitMilli** | Pointer to **float32** | Configured pod CPU limit in millicores, summed across observed containers. Present when a CPU limit is configured and observed. | [optional] 
**MemoryUsageBytes** | Pointer to **float32** | Recent pod memory usage in bytes. Present when memory usage telemetry is observed. | [optional] 
**MemoryLimitBytes** | Pointer to **float32** | Configured pod memory limit in bytes, summed across observed containers. Present when a memory limit is configured and observed. | [optional] 
**CpuUtilization** | Pointer to **float32** | Recent CPU usage divided by the configured CPU limit. Present when both usage and limit are observed. | [optional] 
**MemoryUtilization** | Pointer to **float32** | Recent memory usage divided by the configured memory limit. Present when both usage and limit are observed. | [optional] 

## Methods

### NewListKubernetesPodsResponseDataInner

`func NewListKubernetesPodsResponseDataInner(podName string, phase string, readyContainers int32, totalContainers int32, restarts int32, ) *ListKubernetesPodsResponseDataInner`

NewListKubernetesPodsResponseDataInner instantiates a new ListKubernetesPodsResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKubernetesPodsResponseDataInnerWithDefaults

`func NewListKubernetesPodsResponseDataInnerWithDefaults() *ListKubernetesPodsResponseDataInner`

NewListKubernetesPodsResponseDataInnerWithDefaults instantiates a new ListKubernetesPodsResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPodName

`func (o *ListKubernetesPodsResponseDataInner) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *ListKubernetesPodsResponseDataInner) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *ListKubernetesPodsResponseDataInner) SetPodName(v string)`

SetPodName sets PodName field to given value.


### GetNamespace

`func (o *ListKubernetesPodsResponseDataInner) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ListKubernetesPodsResponseDataInner) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ListKubernetesPodsResponseDataInner) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ListKubernetesPodsResponseDataInner) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetCluster

`func (o *ListKubernetesPodsResponseDataInner) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ListKubernetesPodsResponseDataInner) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ListKubernetesPodsResponseDataInner) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ListKubernetesPodsResponseDataInner) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetTeam

`func (o *ListKubernetesPodsResponseDataInner) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ListKubernetesPodsResponseDataInner) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ListKubernetesPodsResponseDataInner) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ListKubernetesPodsResponseDataInner) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetEnv

`func (o *ListKubernetesPodsResponseDataInner) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ListKubernetesPodsResponseDataInner) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ListKubernetesPodsResponseDataInner) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ListKubernetesPodsResponseDataInner) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetNode

`func (o *ListKubernetesPodsResponseDataInner) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *ListKubernetesPodsResponseDataInner) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *ListKubernetesPodsResponseDataInner) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *ListKubernetesPodsResponseDataInner) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetPhase

`func (o *ListKubernetesPodsResponseDataInner) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ListKubernetesPodsResponseDataInner) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ListKubernetesPodsResponseDataInner) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetReadyContainers

`func (o *ListKubernetesPodsResponseDataInner) GetReadyContainers() int32`

GetReadyContainers returns the ReadyContainers field if non-nil, zero value otherwise.

### GetReadyContainersOk

`func (o *ListKubernetesPodsResponseDataInner) GetReadyContainersOk() (*int32, bool)`

GetReadyContainersOk returns a tuple with the ReadyContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyContainers

`func (o *ListKubernetesPodsResponseDataInner) SetReadyContainers(v int32)`

SetReadyContainers sets ReadyContainers field to given value.


### GetTotalContainers

`func (o *ListKubernetesPodsResponseDataInner) GetTotalContainers() int32`

GetTotalContainers returns the TotalContainers field if non-nil, zero value otherwise.

### GetTotalContainersOk

`func (o *ListKubernetesPodsResponseDataInner) GetTotalContainersOk() (*int32, bool)`

GetTotalContainersOk returns a tuple with the TotalContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalContainers

`func (o *ListKubernetesPodsResponseDataInner) SetTotalContainers(v int32)`

SetTotalContainers sets TotalContainers field to given value.


### GetRestarts

`func (o *ListKubernetesPodsResponseDataInner) GetRestarts() int32`

GetRestarts returns the Restarts field if non-nil, zero value otherwise.

### GetRestartsOk

`func (o *ListKubernetesPodsResponseDataInner) GetRestartsOk() (*int32, bool)`

GetRestartsOk returns a tuple with the Restarts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestarts

`func (o *ListKubernetesPodsResponseDataInner) SetRestarts(v int32)`

SetRestarts sets Restarts field to given value.


### GetCreatedAt

`func (o *ListKubernetesPodsResponseDataInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListKubernetesPodsResponseDataInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListKubernetesPodsResponseDataInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListKubernetesPodsResponseDataInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetWorkload

`func (o *ListKubernetesPodsResponseDataInner) GetWorkload() ListKubernetesPodsResponseDataInnerWorkload`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *ListKubernetesPodsResponseDataInner) GetWorkloadOk() (*ListKubernetesPodsResponseDataInnerWorkload, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *ListKubernetesPodsResponseDataInner) SetWorkload(v ListKubernetesPodsResponseDataInnerWorkload)`

SetWorkload sets Workload field to given value.

### HasWorkload

`func (o *ListKubernetesPodsResponseDataInner) HasWorkload() bool`

HasWorkload returns a boolean if a field has been set.

### GetCpuUsageMilli

`func (o *ListKubernetesPodsResponseDataInner) GetCpuUsageMilli() float32`

GetCpuUsageMilli returns the CpuUsageMilli field if non-nil, zero value otherwise.

### GetCpuUsageMilliOk

`func (o *ListKubernetesPodsResponseDataInner) GetCpuUsageMilliOk() (*float32, bool)`

GetCpuUsageMilliOk returns a tuple with the CpuUsageMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageMilli

`func (o *ListKubernetesPodsResponseDataInner) SetCpuUsageMilli(v float32)`

SetCpuUsageMilli sets CpuUsageMilli field to given value.

### HasCpuUsageMilli

`func (o *ListKubernetesPodsResponseDataInner) HasCpuUsageMilli() bool`

HasCpuUsageMilli returns a boolean if a field has been set.

### GetCpuLimitMilli

`func (o *ListKubernetesPodsResponseDataInner) GetCpuLimitMilli() float32`

GetCpuLimitMilli returns the CpuLimitMilli field if non-nil, zero value otherwise.

### GetCpuLimitMilliOk

`func (o *ListKubernetesPodsResponseDataInner) GetCpuLimitMilliOk() (*float32, bool)`

GetCpuLimitMilliOk returns a tuple with the CpuLimitMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimitMilli

`func (o *ListKubernetesPodsResponseDataInner) SetCpuLimitMilli(v float32)`

SetCpuLimitMilli sets CpuLimitMilli field to given value.

### HasCpuLimitMilli

`func (o *ListKubernetesPodsResponseDataInner) HasCpuLimitMilli() bool`

HasCpuLimitMilli returns a boolean if a field has been set.

### GetMemoryUsageBytes

`func (o *ListKubernetesPodsResponseDataInner) GetMemoryUsageBytes() float32`

GetMemoryUsageBytes returns the MemoryUsageBytes field if non-nil, zero value otherwise.

### GetMemoryUsageBytesOk

`func (o *ListKubernetesPodsResponseDataInner) GetMemoryUsageBytesOk() (*float32, bool)`

GetMemoryUsageBytesOk returns a tuple with the MemoryUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsageBytes

`func (o *ListKubernetesPodsResponseDataInner) SetMemoryUsageBytes(v float32)`

SetMemoryUsageBytes sets MemoryUsageBytes field to given value.

### HasMemoryUsageBytes

`func (o *ListKubernetesPodsResponseDataInner) HasMemoryUsageBytes() bool`

HasMemoryUsageBytes returns a boolean if a field has been set.

### GetMemoryLimitBytes

`func (o *ListKubernetesPodsResponseDataInner) GetMemoryLimitBytes() float32`

GetMemoryLimitBytes returns the MemoryLimitBytes field if non-nil, zero value otherwise.

### GetMemoryLimitBytesOk

`func (o *ListKubernetesPodsResponseDataInner) GetMemoryLimitBytesOk() (*float32, bool)`

GetMemoryLimitBytesOk returns a tuple with the MemoryLimitBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimitBytes

`func (o *ListKubernetesPodsResponseDataInner) SetMemoryLimitBytes(v float32)`

SetMemoryLimitBytes sets MemoryLimitBytes field to given value.

### HasMemoryLimitBytes

`func (o *ListKubernetesPodsResponseDataInner) HasMemoryLimitBytes() bool`

HasMemoryLimitBytes returns a boolean if a field has been set.

### GetCpuUtilization

`func (o *ListKubernetesPodsResponseDataInner) GetCpuUtilization() float32`

GetCpuUtilization returns the CpuUtilization field if non-nil, zero value otherwise.

### GetCpuUtilizationOk

`func (o *ListKubernetesPodsResponseDataInner) GetCpuUtilizationOk() (*float32, bool)`

GetCpuUtilizationOk returns a tuple with the CpuUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtilization

`func (o *ListKubernetesPodsResponseDataInner) SetCpuUtilization(v float32)`

SetCpuUtilization sets CpuUtilization field to given value.

### HasCpuUtilization

`func (o *ListKubernetesPodsResponseDataInner) HasCpuUtilization() bool`

HasCpuUtilization returns a boolean if a field has been set.

### GetMemoryUtilization

`func (o *ListKubernetesPodsResponseDataInner) GetMemoryUtilization() float32`

GetMemoryUtilization returns the MemoryUtilization field if non-nil, zero value otherwise.

### GetMemoryUtilizationOk

`func (o *ListKubernetesPodsResponseDataInner) GetMemoryUtilizationOk() (*float32, bool)`

GetMemoryUtilizationOk returns a tuple with the MemoryUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUtilization

`func (o *ListKubernetesPodsResponseDataInner) SetMemoryUtilization(v float32)`

SetMemoryUtilization sets MemoryUtilization field to given value.

### HasMemoryUtilization

`func (o *ListKubernetesPodsResponseDataInner) HasMemoryUtilization() bool`

HasMemoryUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


