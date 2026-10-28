# GetKubernetesDeploymentDetailsResponseDataRevisionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplicaSetName** | **string** | ReplicaSet name for a Deployment revision from workload telemetry. | 
**Revision** | **int32** | Deployment revision number from ReplicaSet annotations. Integer greater than or equal to zero. | 
**FirstSeen** | Pointer to **string** | ReplicaSet metadata.creationTimestamp from the latest observed ReplicaSet object. Present when creation time is observed. | [optional] 
**FaultyReasons** | Pointer to **[]string** | Container status reasons the Kubernetes faulty deployment monitor reported for this revision within the revisions lookback. Present only when the revision was flagged as faulty. | [optional] 
**LogsSearchFilter** | **string** | TQL log search filter scoped to this ReplicaSet revision. Built from cluster, namespace, and ReplicaSet identifiers. | 

## Methods

### NewGetKubernetesDeploymentDetailsResponseDataRevisionsInner

`func NewGetKubernetesDeploymentDetailsResponseDataRevisionsInner(replicaSetName string, revision int32, logsSearchFilter string, ) *GetKubernetesDeploymentDetailsResponseDataRevisionsInner`

NewGetKubernetesDeploymentDetailsResponseDataRevisionsInner instantiates a new GetKubernetesDeploymentDetailsResponseDataRevisionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKubernetesDeploymentDetailsResponseDataRevisionsInnerWithDefaults

`func NewGetKubernetesDeploymentDetailsResponseDataRevisionsInnerWithDefaults() *GetKubernetesDeploymentDetailsResponseDataRevisionsInner`

NewGetKubernetesDeploymentDetailsResponseDataRevisionsInnerWithDefaults instantiates a new GetKubernetesDeploymentDetailsResponseDataRevisionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicaSetName

`func (o *GetKubernetesDeploymentDetailsResponseDataRevisionsInner) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *GetKubernetesDeploymentDetailsResponseDataRevisionsInner) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *GetKubernetesDeploymentDetailsResponseDataRevisionsInner) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.


### GetRevision

`func (o *GetKubernetesDeploymentDetailsResponseDataRevisionsInner) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *GetKubernetesDeploymentDetailsResponseDataRevisionsInner) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *GetKubernetesDeploymentDetailsResponseDataRevisionsInner) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetFirstSeen

`func (o *GetKubernetesDeploymentDetailsResponseDataRevisionsInner) GetFirstSeen() string`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *GetKubernetesDeploymentDetailsResponseDataRevisionsInner) GetFirstSeenOk() (*string, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *GetKubernetesDeploymentDetailsResponseDataRevisionsInner) SetFirstSeen(v string)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *GetKubernetesDeploymentDetailsResponseDataRevisionsInner) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetFaultyReasons

`func (o *GetKubernetesDeploymentDetailsResponseDataRevisionsInner) GetFaultyReasons() []string`

GetFaultyReasons returns the FaultyReasons field if non-nil, zero value otherwise.

### GetFaultyReasonsOk

`func (o *GetKubernetesDeploymentDetailsResponseDataRevisionsInner) GetFaultyReasonsOk() (*[]string, bool)`

GetFaultyReasonsOk returns a tuple with the FaultyReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultyReasons

`func (o *GetKubernetesDeploymentDetailsResponseDataRevisionsInner) SetFaultyReasons(v []string)`

SetFaultyReasons sets FaultyReasons field to given value.

### HasFaultyReasons

`func (o *GetKubernetesDeploymentDetailsResponseDataRevisionsInner) HasFaultyReasons() bool`

HasFaultyReasons returns a boolean if a field has been set.

### GetLogsSearchFilter

`func (o *GetKubernetesDeploymentDetailsResponseDataRevisionsInner) GetLogsSearchFilter() string`

GetLogsSearchFilter returns the LogsSearchFilter field if non-nil, zero value otherwise.

### GetLogsSearchFilterOk

`func (o *GetKubernetesDeploymentDetailsResponseDataRevisionsInner) GetLogsSearchFilterOk() (*string, bool)`

GetLogsSearchFilterOk returns a tuple with the LogsSearchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsSearchFilter

`func (o *GetKubernetesDeploymentDetailsResponseDataRevisionsInner) SetLogsSearchFilter(v string)`

SetLogsSearchFilter sets LogsSearchFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


