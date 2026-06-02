# GraphVisualizationTopListConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the database rows-based aggregation as a ranked top list | 
**ConnectionId** | **string** | The ID of the connection to use to query the datastore. | 
**Queries** | **[]string** |  | 

## Methods

### NewGraphVisualizationTopListConnection

`func NewGraphVisualizationTopListConnection(type_ string, connectionId string, queries []string, ) *GraphVisualizationTopListConnection`

NewGraphVisualizationTopListConnection instantiates a new GraphVisualizationTopListConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationTopListConnectionWithDefaults

`func NewGraphVisualizationTopListConnectionWithDefaults() *GraphVisualizationTopListConnection`

NewGraphVisualizationTopListConnectionWithDefaults instantiates a new GraphVisualizationTopListConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GraphVisualizationTopListConnection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphVisualizationTopListConnection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphVisualizationTopListConnection) SetType(v string)`

SetType sets Type field to given value.


### GetConnectionId

`func (o *GraphVisualizationTopListConnection) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *GraphVisualizationTopListConnection) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *GraphVisualizationTopListConnection) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetQueries

`func (o *GraphVisualizationTopListConnection) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *GraphVisualizationTopListConnection) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *GraphVisualizationTopListConnection) SetQueries(v []string)`

SetQueries sets Queries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


