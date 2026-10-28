# CreateDataDeletionRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** | ID of the cluster to delete the events from. | 
**DataSource** | **string** | Telemetry signal to delete events from. | 
**Query** | **string** | Tsuga query selecting the events to delete. Queries matching every event, full-text searches and queries that only exclude events (&#x60;NOT user.id:123&#x60;) are rejected: select events by attribute. | 
**From** | **int64** | Start of the query window as a Unix timestamp in seconds | 
**To** | **int64** | End of the query window as a Unix timestamp in seconds | 
**Reason** | **string** | Why the deletion is requested, for example the reference of a GDPR request. Kept for audit. | 
**DryRun** | **bool** | Set to true to only count the events the request matches, without deleting them. | 
**Owner** | **string** | ID of the team that owns the request. | 

## Methods

### NewCreateDataDeletionRequestRequest

`func NewCreateDataDeletionRequestRequest(clusterId string, dataSource string, query string, from int64, to int64, reason string, dryRun bool, owner string, ) *CreateDataDeletionRequestRequest`

NewCreateDataDeletionRequestRequest instantiates a new CreateDataDeletionRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDataDeletionRequestRequestWithDefaults

`func NewCreateDataDeletionRequestRequestWithDefaults() *CreateDataDeletionRequestRequest`

NewCreateDataDeletionRequestRequestWithDefaults instantiates a new CreateDataDeletionRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *CreateDataDeletionRequestRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *CreateDataDeletionRequestRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *CreateDataDeletionRequestRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetDataSource

`func (o *CreateDataDeletionRequestRequest) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *CreateDataDeletionRequestRequest) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *CreateDataDeletionRequestRequest) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.


### GetQuery

`func (o *CreateDataDeletionRequestRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CreateDataDeletionRequestRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CreateDataDeletionRequestRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetFrom

`func (o *CreateDataDeletionRequestRequest) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CreateDataDeletionRequestRequest) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CreateDataDeletionRequestRequest) SetFrom(v int64)`

SetFrom sets From field to given value.


### GetTo

`func (o *CreateDataDeletionRequestRequest) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CreateDataDeletionRequestRequest) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CreateDataDeletionRequestRequest) SetTo(v int64)`

SetTo sets To field to given value.


### GetReason

`func (o *CreateDataDeletionRequestRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateDataDeletionRequestRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateDataDeletionRequestRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetDryRun

`func (o *CreateDataDeletionRequestRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateDataDeletionRequestRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateDataDeletionRequestRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.


### GetOwner

`func (o *CreateDataDeletionRequestRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateDataDeletionRequestRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateDataDeletionRequestRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


