# ListKubernetesClustersResponseDataInnerPodPhaseCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pending** | **int32** | Number of observed pods currently in the Pending phase. | 
**Running** | **int32** | Number of observed pods currently in the Running phase. | 
**Succeeded** | **int32** | Number of observed pods currently in the Succeeded phase. | 
**Failed** | **int32** | Number of observed pods currently in the Failed phase. | 
**Unknown** | **int32** | Number of observed pods currently in the Unknown phase. | 

## Methods

### NewListKubernetesClustersResponseDataInnerPodPhaseCounts

`func NewListKubernetesClustersResponseDataInnerPodPhaseCounts(pending int32, running int32, succeeded int32, failed int32, unknown int32, ) *ListKubernetesClustersResponseDataInnerPodPhaseCounts`

NewListKubernetesClustersResponseDataInnerPodPhaseCounts instantiates a new ListKubernetesClustersResponseDataInnerPodPhaseCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKubernetesClustersResponseDataInnerPodPhaseCountsWithDefaults

`func NewListKubernetesClustersResponseDataInnerPodPhaseCountsWithDefaults() *ListKubernetesClustersResponseDataInnerPodPhaseCounts`

NewListKubernetesClustersResponseDataInnerPodPhaseCountsWithDefaults instantiates a new ListKubernetesClustersResponseDataInnerPodPhaseCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPending

`func (o *ListKubernetesClustersResponseDataInnerPodPhaseCounts) GetPending() int32`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *ListKubernetesClustersResponseDataInnerPodPhaseCounts) GetPendingOk() (*int32, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *ListKubernetesClustersResponseDataInnerPodPhaseCounts) SetPending(v int32)`

SetPending sets Pending field to given value.


### GetRunning

`func (o *ListKubernetesClustersResponseDataInnerPodPhaseCounts) GetRunning() int32`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *ListKubernetesClustersResponseDataInnerPodPhaseCounts) GetRunningOk() (*int32, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *ListKubernetesClustersResponseDataInnerPodPhaseCounts) SetRunning(v int32)`

SetRunning sets Running field to given value.


### GetSucceeded

`func (o *ListKubernetesClustersResponseDataInnerPodPhaseCounts) GetSucceeded() int32`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *ListKubernetesClustersResponseDataInnerPodPhaseCounts) GetSucceededOk() (*int32, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *ListKubernetesClustersResponseDataInnerPodPhaseCounts) SetSucceeded(v int32)`

SetSucceeded sets Succeeded field to given value.


### GetFailed

`func (o *ListKubernetesClustersResponseDataInnerPodPhaseCounts) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ListKubernetesClustersResponseDataInnerPodPhaseCounts) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ListKubernetesClustersResponseDataInnerPodPhaseCounts) SetFailed(v int32)`

SetFailed sets Failed field to given value.


### GetUnknown

`func (o *ListKubernetesClustersResponseDataInnerPodPhaseCounts) GetUnknown() int32`

GetUnknown returns the Unknown field if non-nil, zero value otherwise.

### GetUnknownOk

`func (o *ListKubernetesClustersResponseDataInnerPodPhaseCounts) GetUnknownOk() (*int32, bool)`

GetUnknownOk returns a tuple with the Unknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknown

`func (o *ListKubernetesClustersResponseDataInnerPodPhaseCounts) SetUnknown(v int32)`

SetUnknown sets Unknown field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


