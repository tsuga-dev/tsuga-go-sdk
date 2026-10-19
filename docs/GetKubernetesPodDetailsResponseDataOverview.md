# GetKubernetesPodDetailsResponseDataOverview

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

## Methods

### NewGetKubernetesPodDetailsResponseDataOverview

`func NewGetKubernetesPodDetailsResponseDataOverview(podName string, phase string, readyContainers int32, totalContainers int32, restarts int32, ) *GetKubernetesPodDetailsResponseDataOverview`

NewGetKubernetesPodDetailsResponseDataOverview instantiates a new GetKubernetesPodDetailsResponseDataOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKubernetesPodDetailsResponseDataOverviewWithDefaults

`func NewGetKubernetesPodDetailsResponseDataOverviewWithDefaults() *GetKubernetesPodDetailsResponseDataOverview`

NewGetKubernetesPodDetailsResponseDataOverviewWithDefaults instantiates a new GetKubernetesPodDetailsResponseDataOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPodName

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *GetKubernetesPodDetailsResponseDataOverview) SetPodName(v string)`

SetPodName sets PodName field to given value.


### GetNamespace

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *GetKubernetesPodDetailsResponseDataOverview) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *GetKubernetesPodDetailsResponseDataOverview) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetCluster

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *GetKubernetesPodDetailsResponseDataOverview) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *GetKubernetesPodDetailsResponseDataOverview) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetTeam

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *GetKubernetesPodDetailsResponseDataOverview) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *GetKubernetesPodDetailsResponseDataOverview) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetEnv

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *GetKubernetesPodDetailsResponseDataOverview) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *GetKubernetesPodDetailsResponseDataOverview) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetNode

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *GetKubernetesPodDetailsResponseDataOverview) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *GetKubernetesPodDetailsResponseDataOverview) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetPhase

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *GetKubernetesPodDetailsResponseDataOverview) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetReadyContainers

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetReadyContainers() int32`

GetReadyContainers returns the ReadyContainers field if non-nil, zero value otherwise.

### GetReadyContainersOk

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetReadyContainersOk() (*int32, bool)`

GetReadyContainersOk returns a tuple with the ReadyContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyContainers

`func (o *GetKubernetesPodDetailsResponseDataOverview) SetReadyContainers(v int32)`

SetReadyContainers sets ReadyContainers field to given value.


### GetTotalContainers

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetTotalContainers() int32`

GetTotalContainers returns the TotalContainers field if non-nil, zero value otherwise.

### GetTotalContainersOk

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetTotalContainersOk() (*int32, bool)`

GetTotalContainersOk returns a tuple with the TotalContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalContainers

`func (o *GetKubernetesPodDetailsResponseDataOverview) SetTotalContainers(v int32)`

SetTotalContainers sets TotalContainers field to given value.


### GetRestarts

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetRestarts() int32`

GetRestarts returns the Restarts field if non-nil, zero value otherwise.

### GetRestartsOk

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetRestartsOk() (*int32, bool)`

GetRestartsOk returns a tuple with the Restarts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestarts

`func (o *GetKubernetesPodDetailsResponseDataOverview) SetRestarts(v int32)`

SetRestarts sets Restarts field to given value.


### GetCreatedAt

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetKubernetesPodDetailsResponseDataOverview) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetKubernetesPodDetailsResponseDataOverview) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetWorkload

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetWorkload() ListKubernetesPodsResponseDataInnerWorkload`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *GetKubernetesPodDetailsResponseDataOverview) GetWorkloadOk() (*ListKubernetesPodsResponseDataInnerWorkload, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *GetKubernetesPodDetailsResponseDataOverview) SetWorkload(v ListKubernetesPodsResponseDataInnerWorkload)`

SetWorkload sets Workload field to given value.

### HasWorkload

`func (o *GetKubernetesPodDetailsResponseDataOverview) HasWorkload() bool`

HasWorkload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


