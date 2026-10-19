# GetKubernetesDeploymentDetailsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overview** | [**GetKubernetesDeploymentDetailsResponseDataOverview**](GetKubernetesDeploymentDetailsResponseDataOverview.md) |  | 
**PodReplicas** | [**[]GetKubernetesDeploymentDetailsResponseDataPodReplicasInner**](GetKubernetesDeploymentDetailsResponseDataPodReplicasInner.md) | Pod instances with recent CPU or memory telemetry for the Deployment. Tsuga builds this list from pod usage telemetry and workload attributes. Empty when no matching pod usage telemetry is observed. | 
**Revisions** | [**[]GetKubernetesDeploymentDetailsResponseDataRevisionsInner**](GetKubernetesDeploymentDetailsResponseDataRevisionsInner.md) | Deployment rollout revisions from recent ReplicaSet object telemetry. Includes the latest observed event per ReplicaSet, excludes ReplicaSets whose latest observed spec.replicas is 0, and returns up to 8 revisions sorted by revision number descending. | 

## Methods

### NewGetKubernetesDeploymentDetailsResponseData

`func NewGetKubernetesDeploymentDetailsResponseData(overview GetKubernetesDeploymentDetailsResponseDataOverview, podReplicas []GetKubernetesDeploymentDetailsResponseDataPodReplicasInner, revisions []GetKubernetesDeploymentDetailsResponseDataRevisionsInner, ) *GetKubernetesDeploymentDetailsResponseData`

NewGetKubernetesDeploymentDetailsResponseData instantiates a new GetKubernetesDeploymentDetailsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKubernetesDeploymentDetailsResponseDataWithDefaults

`func NewGetKubernetesDeploymentDetailsResponseDataWithDefaults() *GetKubernetesDeploymentDetailsResponseData`

NewGetKubernetesDeploymentDetailsResponseDataWithDefaults instantiates a new GetKubernetesDeploymentDetailsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverview

`func (o *GetKubernetesDeploymentDetailsResponseData) GetOverview() GetKubernetesDeploymentDetailsResponseDataOverview`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *GetKubernetesDeploymentDetailsResponseData) GetOverviewOk() (*GetKubernetesDeploymentDetailsResponseDataOverview, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *GetKubernetesDeploymentDetailsResponseData) SetOverview(v GetKubernetesDeploymentDetailsResponseDataOverview)`

SetOverview sets Overview field to given value.


### GetPodReplicas

`func (o *GetKubernetesDeploymentDetailsResponseData) GetPodReplicas() []GetKubernetesDeploymentDetailsResponseDataPodReplicasInner`

GetPodReplicas returns the PodReplicas field if non-nil, zero value otherwise.

### GetPodReplicasOk

`func (o *GetKubernetesDeploymentDetailsResponseData) GetPodReplicasOk() (*[]GetKubernetesDeploymentDetailsResponseDataPodReplicasInner, bool)`

GetPodReplicasOk returns a tuple with the PodReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodReplicas

`func (o *GetKubernetesDeploymentDetailsResponseData) SetPodReplicas(v []GetKubernetesDeploymentDetailsResponseDataPodReplicasInner)`

SetPodReplicas sets PodReplicas field to given value.


### GetRevisions

`func (o *GetKubernetesDeploymentDetailsResponseData) GetRevisions() []GetKubernetesDeploymentDetailsResponseDataRevisionsInner`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *GetKubernetesDeploymentDetailsResponseData) GetRevisionsOk() (*[]GetKubernetesDeploymentDetailsResponseDataRevisionsInner, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *GetKubernetesDeploymentDetailsResponseData) SetRevisions(v []GetKubernetesDeploymentDetailsResponseDataRevisionsInner)`

SetRevisions sets Revisions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


