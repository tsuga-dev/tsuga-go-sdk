# DataDeletionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the deletion request. | 
**ClusterId** | **string** | ID of the cluster the request deletes events from. | 
**DataSource** | **string** | Telemetry signal the request deletes events from. | 
**Query** | **string** | Tsuga query selecting the events the request deletes. | 
**From** | **string** | Inclusive start of the event time range. | 
**To** | **string** | Exclusive end of the event time range. | 
**Reason** | **string** | Why the deletion was requested, kept for audit. | 
**DryRun** | **bool** | Whether the cluster only counts the events the request matches, without deleting them. | 
**Status** | **string** | Status of the request: &#x60;cancelled&#x60; once cancelled, which is final. Otherwise the status the cluster executes it with: &#x60;pending&#x60; until the cluster starts it, then &#x60;running&#x60;, then &#x60;done&#x60;, or &#x60;failed&#x60; if it cannot be executed. A &#x60;done&#x60; request goes back to &#x60;running&#x60; when late events arrive. Also &#x60;pending&#x60; when the cluster could not be reached. | 
**Execution** | Pointer to [**DataDeletionRequestExecution**](DataDeletionRequestExecution.md) |  | [optional] 
**MatchedEventCount** | **int32** | Number of events the request matched at creation. The cluster deletes the events matching when it executes the request, which can differ. | 
**DueAt** | **string** | Time from which the cluster executes the request. Cancelling before then prevents any deletion; cancelling later stops the remaining work, and events already deleted stay deleted. | 
**Owner** | **string** | ID of the team that owns the request. | 

## Methods

### NewDataDeletionRequest

`func NewDataDeletionRequest(id string, clusterId string, dataSource string, query string, from string, to string, reason string, dryRun bool, status string, matchedEventCount int32, dueAt string, owner string, ) *DataDeletionRequest`

NewDataDeletionRequest instantiates a new DataDeletionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataDeletionRequestWithDefaults

`func NewDataDeletionRequestWithDefaults() *DataDeletionRequest`

NewDataDeletionRequestWithDefaults instantiates a new DataDeletionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataDeletionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataDeletionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataDeletionRequest) SetId(v string)`

SetId sets Id field to given value.


### GetClusterId

`func (o *DataDeletionRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *DataDeletionRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *DataDeletionRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetDataSource

`func (o *DataDeletionRequest) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *DataDeletionRequest) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *DataDeletionRequest) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.


### GetQuery

`func (o *DataDeletionRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *DataDeletionRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *DataDeletionRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetFrom

`func (o *DataDeletionRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *DataDeletionRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *DataDeletionRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *DataDeletionRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *DataDeletionRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *DataDeletionRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetReason

`func (o *DataDeletionRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DataDeletionRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DataDeletionRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetDryRun

`func (o *DataDeletionRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DataDeletionRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DataDeletionRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.


### GetStatus

`func (o *DataDeletionRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DataDeletionRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DataDeletionRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetExecution

`func (o *DataDeletionRequest) GetExecution() DataDeletionRequestExecution`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *DataDeletionRequest) GetExecutionOk() (*DataDeletionRequestExecution, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *DataDeletionRequest) SetExecution(v DataDeletionRequestExecution)`

SetExecution sets Execution field to given value.

### HasExecution

`func (o *DataDeletionRequest) HasExecution() bool`

HasExecution returns a boolean if a field has been set.

### GetMatchedEventCount

`func (o *DataDeletionRequest) GetMatchedEventCount() int32`

GetMatchedEventCount returns the MatchedEventCount field if non-nil, zero value otherwise.

### GetMatchedEventCountOk

`func (o *DataDeletionRequest) GetMatchedEventCountOk() (*int32, bool)`

GetMatchedEventCountOk returns a tuple with the MatchedEventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedEventCount

`func (o *DataDeletionRequest) SetMatchedEventCount(v int32)`

SetMatchedEventCount sets MatchedEventCount field to given value.


### GetDueAt

`func (o *DataDeletionRequest) GetDueAt() string`

GetDueAt returns the DueAt field if non-nil, zero value otherwise.

### GetDueAtOk

`func (o *DataDeletionRequest) GetDueAtOk() (*string, bool)`

GetDueAtOk returns a tuple with the DueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAt

`func (o *DataDeletionRequest) SetDueAt(v string)`

SetDueAt sets DueAt field to given value.


### GetOwner

`func (o *DataDeletionRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DataDeletionRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DataDeletionRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


