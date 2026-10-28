# DataDeletionRequestExecutionProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchedRows** | **int32** | Events matching the request. A dry run counts them without deleting them. | 
**ScannedRows** | **int32** | Events the cluster scanned for matching ones. | 

## Methods

### NewDataDeletionRequestExecutionProgress

`func NewDataDeletionRequestExecutionProgress(matchedRows int32, scannedRows int32, ) *DataDeletionRequestExecutionProgress`

NewDataDeletionRequestExecutionProgress instantiates a new DataDeletionRequestExecutionProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataDeletionRequestExecutionProgressWithDefaults

`func NewDataDeletionRequestExecutionProgressWithDefaults() *DataDeletionRequestExecutionProgress`

NewDataDeletionRequestExecutionProgressWithDefaults instantiates a new DataDeletionRequestExecutionProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchedRows

`func (o *DataDeletionRequestExecutionProgress) GetMatchedRows() int32`

GetMatchedRows returns the MatchedRows field if non-nil, zero value otherwise.

### GetMatchedRowsOk

`func (o *DataDeletionRequestExecutionProgress) GetMatchedRowsOk() (*int32, bool)`

GetMatchedRowsOk returns a tuple with the MatchedRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedRows

`func (o *DataDeletionRequestExecutionProgress) SetMatchedRows(v int32)`

SetMatchedRows sets MatchedRows field to given value.


### GetScannedRows

`func (o *DataDeletionRequestExecutionProgress) GetScannedRows() int32`

GetScannedRows returns the ScannedRows field if non-nil, zero value otherwise.

### GetScannedRowsOk

`func (o *DataDeletionRequestExecutionProgress) GetScannedRowsOk() (*int32, bool)`

GetScannedRowsOk returns a tuple with the ScannedRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScannedRows

`func (o *DataDeletionRequestExecutionProgress) SetScannedRows(v int32)`

SetScannedRows sets ScannedRows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


