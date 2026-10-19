# ListKubernetesNodesResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | **string** | Kubernetes node name from node telemetry. | 
**Cluster** | Pointer to **string** | Kubernetes cluster name that contains the node. Present when cluster telemetry is observed. | [optional] 
**Status** | Pointer to **string** | Readiness of the node, derived from its Ready condition. Present when node condition telemetry is observed. | [optional] 
**InstanceType** | Pointer to **string** | Cloud instance type backing the node (for example &#x60;r7g.2xlarge&#x60;). Present when host telemetry is observed. | [optional] 
**Pods** | Pointer to **int32** | Number of observed pods scheduled on the node, including pods in internal namespaces since they occupy node capacity. Present when pod telemetry is observed. | [optional] 
**AllocatablePods** | Pointer to **int32** | Maximum number of pods the node can run. Present when node allocatable telemetry is observed. | [optional] 
**CpuUsageMilli** | Pointer to **float32** | Recent CPU usage for the node, in millicores. Reported by the kubelet, so it is absent on nodes that only report cluster-level telemetry. | [optional] 
**CpuAllocatableMilli** | Pointer to **float32** | CPU available to pods on the node, in millicores. Present when node allocatable telemetry is observed. | [optional] 
**CpuUtilization** | Pointer to **float32** | Node CPU usage divided by allocatable CPU. Present when both usage and allocatable telemetry are observed. | [optional] 
**MemoryUsageBytes** | Pointer to **float32** | Recent memory usage for the node, in bytes. Reported by the kubelet, so it is absent on nodes that only report cluster-level telemetry. | [optional] 
**MemoryAllocatableBytes** | Pointer to **float32** | Memory available to pods on the node, in bytes. Present when node allocatable telemetry is observed. | [optional] 
**MemoryUtilization** | Pointer to **float32** | Node memory usage divided by allocatable memory. Present when both usage and allocatable telemetry are observed. | [optional] 

## Methods

### NewListKubernetesNodesResponseDataInner

`func NewListKubernetesNodesResponseDataInner(node string, ) *ListKubernetesNodesResponseDataInner`

NewListKubernetesNodesResponseDataInner instantiates a new ListKubernetesNodesResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKubernetesNodesResponseDataInnerWithDefaults

`func NewListKubernetesNodesResponseDataInnerWithDefaults() *ListKubernetesNodesResponseDataInner`

NewListKubernetesNodesResponseDataInnerWithDefaults instantiates a new ListKubernetesNodesResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNode

`func (o *ListKubernetesNodesResponseDataInner) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *ListKubernetesNodesResponseDataInner) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *ListKubernetesNodesResponseDataInner) SetNode(v string)`

SetNode sets Node field to given value.


### GetCluster

`func (o *ListKubernetesNodesResponseDataInner) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ListKubernetesNodesResponseDataInner) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ListKubernetesNodesResponseDataInner) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ListKubernetesNodesResponseDataInner) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetStatus

`func (o *ListKubernetesNodesResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListKubernetesNodesResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListKubernetesNodesResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListKubernetesNodesResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInstanceType

`func (o *ListKubernetesNodesResponseDataInner) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ListKubernetesNodesResponseDataInner) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ListKubernetesNodesResponseDataInner) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ListKubernetesNodesResponseDataInner) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetPods

`func (o *ListKubernetesNodesResponseDataInner) GetPods() int32`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *ListKubernetesNodesResponseDataInner) GetPodsOk() (*int32, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *ListKubernetesNodesResponseDataInner) SetPods(v int32)`

SetPods sets Pods field to given value.

### HasPods

`func (o *ListKubernetesNodesResponseDataInner) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetAllocatablePods

`func (o *ListKubernetesNodesResponseDataInner) GetAllocatablePods() int32`

GetAllocatablePods returns the AllocatablePods field if non-nil, zero value otherwise.

### GetAllocatablePodsOk

`func (o *ListKubernetesNodesResponseDataInner) GetAllocatablePodsOk() (*int32, bool)`

GetAllocatablePodsOk returns a tuple with the AllocatablePods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatablePods

`func (o *ListKubernetesNodesResponseDataInner) SetAllocatablePods(v int32)`

SetAllocatablePods sets AllocatablePods field to given value.

### HasAllocatablePods

`func (o *ListKubernetesNodesResponseDataInner) HasAllocatablePods() bool`

HasAllocatablePods returns a boolean if a field has been set.

### GetCpuUsageMilli

`func (o *ListKubernetesNodesResponseDataInner) GetCpuUsageMilli() float32`

GetCpuUsageMilli returns the CpuUsageMilli field if non-nil, zero value otherwise.

### GetCpuUsageMilliOk

`func (o *ListKubernetesNodesResponseDataInner) GetCpuUsageMilliOk() (*float32, bool)`

GetCpuUsageMilliOk returns a tuple with the CpuUsageMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageMilli

`func (o *ListKubernetesNodesResponseDataInner) SetCpuUsageMilli(v float32)`

SetCpuUsageMilli sets CpuUsageMilli field to given value.

### HasCpuUsageMilli

`func (o *ListKubernetesNodesResponseDataInner) HasCpuUsageMilli() bool`

HasCpuUsageMilli returns a boolean if a field has been set.

### GetCpuAllocatableMilli

`func (o *ListKubernetesNodesResponseDataInner) GetCpuAllocatableMilli() float32`

GetCpuAllocatableMilli returns the CpuAllocatableMilli field if non-nil, zero value otherwise.

### GetCpuAllocatableMilliOk

`func (o *ListKubernetesNodesResponseDataInner) GetCpuAllocatableMilliOk() (*float32, bool)`

GetCpuAllocatableMilliOk returns a tuple with the CpuAllocatableMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuAllocatableMilli

`func (o *ListKubernetesNodesResponseDataInner) SetCpuAllocatableMilli(v float32)`

SetCpuAllocatableMilli sets CpuAllocatableMilli field to given value.

### HasCpuAllocatableMilli

`func (o *ListKubernetesNodesResponseDataInner) HasCpuAllocatableMilli() bool`

HasCpuAllocatableMilli returns a boolean if a field has been set.

### GetCpuUtilization

`func (o *ListKubernetesNodesResponseDataInner) GetCpuUtilization() float32`

GetCpuUtilization returns the CpuUtilization field if non-nil, zero value otherwise.

### GetCpuUtilizationOk

`func (o *ListKubernetesNodesResponseDataInner) GetCpuUtilizationOk() (*float32, bool)`

GetCpuUtilizationOk returns a tuple with the CpuUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtilization

`func (o *ListKubernetesNodesResponseDataInner) SetCpuUtilization(v float32)`

SetCpuUtilization sets CpuUtilization field to given value.

### HasCpuUtilization

`func (o *ListKubernetesNodesResponseDataInner) HasCpuUtilization() bool`

HasCpuUtilization returns a boolean if a field has been set.

### GetMemoryUsageBytes

`func (o *ListKubernetesNodesResponseDataInner) GetMemoryUsageBytes() float32`

GetMemoryUsageBytes returns the MemoryUsageBytes field if non-nil, zero value otherwise.

### GetMemoryUsageBytesOk

`func (o *ListKubernetesNodesResponseDataInner) GetMemoryUsageBytesOk() (*float32, bool)`

GetMemoryUsageBytesOk returns a tuple with the MemoryUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsageBytes

`func (o *ListKubernetesNodesResponseDataInner) SetMemoryUsageBytes(v float32)`

SetMemoryUsageBytes sets MemoryUsageBytes field to given value.

### HasMemoryUsageBytes

`func (o *ListKubernetesNodesResponseDataInner) HasMemoryUsageBytes() bool`

HasMemoryUsageBytes returns a boolean if a field has been set.

### GetMemoryAllocatableBytes

`func (o *ListKubernetesNodesResponseDataInner) GetMemoryAllocatableBytes() float32`

GetMemoryAllocatableBytes returns the MemoryAllocatableBytes field if non-nil, zero value otherwise.

### GetMemoryAllocatableBytesOk

`func (o *ListKubernetesNodesResponseDataInner) GetMemoryAllocatableBytesOk() (*float32, bool)`

GetMemoryAllocatableBytesOk returns a tuple with the MemoryAllocatableBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryAllocatableBytes

`func (o *ListKubernetesNodesResponseDataInner) SetMemoryAllocatableBytes(v float32)`

SetMemoryAllocatableBytes sets MemoryAllocatableBytes field to given value.

### HasMemoryAllocatableBytes

`func (o *ListKubernetesNodesResponseDataInner) HasMemoryAllocatableBytes() bool`

HasMemoryAllocatableBytes returns a boolean if a field has been set.

### GetMemoryUtilization

`func (o *ListKubernetesNodesResponseDataInner) GetMemoryUtilization() float32`

GetMemoryUtilization returns the MemoryUtilization field if non-nil, zero value otherwise.

### GetMemoryUtilizationOk

`func (o *ListKubernetesNodesResponseDataInner) GetMemoryUtilizationOk() (*float32, bool)`

GetMemoryUtilizationOk returns a tuple with the MemoryUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUtilization

`func (o *ListKubernetesNodesResponseDataInner) SetMemoryUtilization(v float32)`

SetMemoryUtilization sets MemoryUtilization field to given value.

### HasMemoryUtilization

`func (o *ListKubernetesNodesResponseDataInner) HasMemoryUtilization() bool`

HasMemoryUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


