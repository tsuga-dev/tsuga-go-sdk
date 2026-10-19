# GetKubernetesDeploymentDetailsResponseDataPodReplicasInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PodName** | **string** | Name of a pod instance associated with the workload. Resolved from recent pod usage telemetry and workload attributes. | 
**PodUid** | Pointer to **string** | Kubernetes UID for the pod replica. Present when UID telemetry is observed. | [optional] 
**CpuUsageMilli** | Pointer to **float32** | Recent CPU usage for the pod replica, in millicores. Present when CPU usage telemetry is observed. | [optional] 
**MemoryUsageBytes** | Pointer to **float32** | Recent memory usage for the pod replica, in bytes. Present when memory usage telemetry is observed. | [optional] 

## Methods

### NewGetKubernetesDeploymentDetailsResponseDataPodReplicasInner

`func NewGetKubernetesDeploymentDetailsResponseDataPodReplicasInner(podName string, ) *GetKubernetesDeploymentDetailsResponseDataPodReplicasInner`

NewGetKubernetesDeploymentDetailsResponseDataPodReplicasInner instantiates a new GetKubernetesDeploymentDetailsResponseDataPodReplicasInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKubernetesDeploymentDetailsResponseDataPodReplicasInnerWithDefaults

`func NewGetKubernetesDeploymentDetailsResponseDataPodReplicasInnerWithDefaults() *GetKubernetesDeploymentDetailsResponseDataPodReplicasInner`

NewGetKubernetesDeploymentDetailsResponseDataPodReplicasInnerWithDefaults instantiates a new GetKubernetesDeploymentDetailsResponseDataPodReplicasInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPodName

`func (o *GetKubernetesDeploymentDetailsResponseDataPodReplicasInner) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *GetKubernetesDeploymentDetailsResponseDataPodReplicasInner) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *GetKubernetesDeploymentDetailsResponseDataPodReplicasInner) SetPodName(v string)`

SetPodName sets PodName field to given value.


### GetPodUid

`func (o *GetKubernetesDeploymentDetailsResponseDataPodReplicasInner) GetPodUid() string`

GetPodUid returns the PodUid field if non-nil, zero value otherwise.

### GetPodUidOk

`func (o *GetKubernetesDeploymentDetailsResponseDataPodReplicasInner) GetPodUidOk() (*string, bool)`

GetPodUidOk returns a tuple with the PodUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodUid

`func (o *GetKubernetesDeploymentDetailsResponseDataPodReplicasInner) SetPodUid(v string)`

SetPodUid sets PodUid field to given value.

### HasPodUid

`func (o *GetKubernetesDeploymentDetailsResponseDataPodReplicasInner) HasPodUid() bool`

HasPodUid returns a boolean if a field has been set.

### GetCpuUsageMilli

`func (o *GetKubernetesDeploymentDetailsResponseDataPodReplicasInner) GetCpuUsageMilli() float32`

GetCpuUsageMilli returns the CpuUsageMilli field if non-nil, zero value otherwise.

### GetCpuUsageMilliOk

`func (o *GetKubernetesDeploymentDetailsResponseDataPodReplicasInner) GetCpuUsageMilliOk() (*float32, bool)`

GetCpuUsageMilliOk returns a tuple with the CpuUsageMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageMilli

`func (o *GetKubernetesDeploymentDetailsResponseDataPodReplicasInner) SetCpuUsageMilli(v float32)`

SetCpuUsageMilli sets CpuUsageMilli field to given value.

### HasCpuUsageMilli

`func (o *GetKubernetesDeploymentDetailsResponseDataPodReplicasInner) HasCpuUsageMilli() bool`

HasCpuUsageMilli returns a boolean if a field has been set.

### GetMemoryUsageBytes

`func (o *GetKubernetesDeploymentDetailsResponseDataPodReplicasInner) GetMemoryUsageBytes() float32`

GetMemoryUsageBytes returns the MemoryUsageBytes field if non-nil, zero value otherwise.

### GetMemoryUsageBytesOk

`func (o *GetKubernetesDeploymentDetailsResponseDataPodReplicasInner) GetMemoryUsageBytesOk() (*float32, bool)`

GetMemoryUsageBytesOk returns a tuple with the MemoryUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsageBytes

`func (o *GetKubernetesDeploymentDetailsResponseDataPodReplicasInner) SetMemoryUsageBytes(v float32)`

SetMemoryUsageBytes sets MemoryUsageBytes field to given value.

### HasMemoryUsageBytes

`func (o *GetKubernetesDeploymentDetailsResponseDataPodReplicasInner) HasMemoryUsageBytes() bool`

HasMemoryUsageBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


