# DataDeletionRequestExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureReason** | Pointer to **string** | Why the cluster could not execute the request, when its status is &#x60;failed&#x60;. | [optional] 
**StartedAt** | Pointer to **string** | When the cluster started executing the request. | [optional] 
**CompletedAt** | Pointer to **string** | When the execution last reached &#x60;done&#x60;, &#x60;failed&#x60; or &#x60;cancelled&#x60;. | [optional] 
**Progress** | [**DataDeletionRequestExecutionProgress**](DataDeletionRequestExecutionProgress.md) |  | 

## Methods

### NewDataDeletionRequestExecution

`func NewDataDeletionRequestExecution(progress DataDeletionRequestExecutionProgress, ) *DataDeletionRequestExecution`

NewDataDeletionRequestExecution instantiates a new DataDeletionRequestExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataDeletionRequestExecutionWithDefaults

`func NewDataDeletionRequestExecutionWithDefaults() *DataDeletionRequestExecution`

NewDataDeletionRequestExecutionWithDefaults instantiates a new DataDeletionRequestExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureReason

`func (o *DataDeletionRequestExecution) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *DataDeletionRequestExecution) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *DataDeletionRequestExecution) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *DataDeletionRequestExecution) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetStartedAt

`func (o *DataDeletionRequestExecution) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DataDeletionRequestExecution) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DataDeletionRequestExecution) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DataDeletionRequestExecution) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *DataDeletionRequestExecution) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *DataDeletionRequestExecution) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *DataDeletionRequestExecution) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *DataDeletionRequestExecution) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetProgress

`func (o *DataDeletionRequestExecution) GetProgress() DataDeletionRequestExecutionProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *DataDeletionRequestExecution) GetProgressOk() (*DataDeletionRequestExecutionProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *DataDeletionRequestExecution) SetProgress(v DataDeletionRequestExecutionProgress)`

SetProgress sets Progress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


