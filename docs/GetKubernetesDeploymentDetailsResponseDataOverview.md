# GetKubernetesDeploymentDetailsResponseDataOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentName** | **string** | Kubernetes Deployment name for this detail response. | 
**Namespace** | Pointer to **string** | Kubernetes namespace for this detail response. Present when any deployment signal is found. | [optional] 
**Cluster** | Pointer to **string** | Kubernetes cluster name for this detail response. Present when any deployment signal is found. | [optional] 
**Team** | Pointer to **string** | Team metadata from the Deployment available-replica metric. Present when the metric includes team metadata. | [optional] 
**Env** | Pointer to **string** | Environment metadata from the Deployment available-replica metric. Present when the metric includes environment metadata. | [optional] 
**AvailableReplicas** | Pointer to **int32** | Number of available Deployment replicas from Deployment status telemetry. Present when deployment status is observed. | [optional] 
**DesiredReplicas** | Pointer to **int32** | Desired replica count configured for the Deployment. Present when Deployment spec telemetry is observed. | [optional] 
**LatestRevisionAt** | Pointer to **string** | Latest ReplicaSet metadata.creationTimestamp observed for the Deployment. Present when ReplicaSet object telemetry is observed. | [optional] 

## Methods

### NewGetKubernetesDeploymentDetailsResponseDataOverview

`func NewGetKubernetesDeploymentDetailsResponseDataOverview(deploymentName string, ) *GetKubernetesDeploymentDetailsResponseDataOverview`

NewGetKubernetesDeploymentDetailsResponseDataOverview instantiates a new GetKubernetesDeploymentDetailsResponseDataOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKubernetesDeploymentDetailsResponseDataOverviewWithDefaults

`func NewGetKubernetesDeploymentDetailsResponseDataOverviewWithDefaults() *GetKubernetesDeploymentDetailsResponseDataOverview`

NewGetKubernetesDeploymentDetailsResponseDataOverviewWithDefaults instantiates a new GetKubernetesDeploymentDetailsResponseDataOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentName

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) GetDeploymentName() string`

GetDeploymentName returns the DeploymentName field if non-nil, zero value otherwise.

### GetDeploymentNameOk

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) GetDeploymentNameOk() (*string, bool)`

GetDeploymentNameOk returns a tuple with the DeploymentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentName

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) SetDeploymentName(v string)`

SetDeploymentName sets DeploymentName field to given value.


### GetNamespace

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetCluster

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetTeam

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetEnv

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetAvailableReplicas

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) GetAvailableReplicas() int32`

GetAvailableReplicas returns the AvailableReplicas field if non-nil, zero value otherwise.

### GetAvailableReplicasOk

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) GetAvailableReplicasOk() (*int32, bool)`

GetAvailableReplicasOk returns a tuple with the AvailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableReplicas

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) SetAvailableReplicas(v int32)`

SetAvailableReplicas sets AvailableReplicas field to given value.

### HasAvailableReplicas

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) HasAvailableReplicas() bool`

HasAvailableReplicas returns a boolean if a field has been set.

### GetDesiredReplicas

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) GetDesiredReplicas() int32`

GetDesiredReplicas returns the DesiredReplicas field if non-nil, zero value otherwise.

### GetDesiredReplicasOk

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) GetDesiredReplicasOk() (*int32, bool)`

GetDesiredReplicasOk returns a tuple with the DesiredReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredReplicas

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) SetDesiredReplicas(v int32)`

SetDesiredReplicas sets DesiredReplicas field to given value.

### HasDesiredReplicas

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) HasDesiredReplicas() bool`

HasDesiredReplicas returns a boolean if a field has been set.

### GetLatestRevisionAt

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) GetLatestRevisionAt() string`

GetLatestRevisionAt returns the LatestRevisionAt field if non-nil, zero value otherwise.

### GetLatestRevisionAtOk

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) GetLatestRevisionAtOk() (*string, bool)`

GetLatestRevisionAtOk returns a tuple with the LatestRevisionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevisionAt

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) SetLatestRevisionAt(v string)`

SetLatestRevisionAt sets LatestRevisionAt field to given value.

### HasLatestRevisionAt

`func (o *GetKubernetesDeploymentDetailsResponseDataOverview) HasLatestRevisionAt() bool`

HasLatestRevisionAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


