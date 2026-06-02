# InputGraphVisualizationPieConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the database rows-based aggregation as a pie chart | 
**ConnectionId** | **string** | The ID of the connection to use to query the datastore. | 
**Queries** | **[]string** |  | 
**LegendMode** | Pointer to **string** | Controls whether and how the widget displays legend or series details (e.g. table, legend-only, or no legend) | [optional] 

## Methods

### NewInputGraphVisualizationPieConnection

`func NewInputGraphVisualizationPieConnection(type_ string, connectionId string, queries []string, ) *InputGraphVisualizationPieConnection`

NewInputGraphVisualizationPieConnection instantiates a new InputGraphVisualizationPieConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationPieConnectionWithDefaults

`func NewInputGraphVisualizationPieConnectionWithDefaults() *InputGraphVisualizationPieConnection`

NewInputGraphVisualizationPieConnectionWithDefaults instantiates a new InputGraphVisualizationPieConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputGraphVisualizationPieConnection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputGraphVisualizationPieConnection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputGraphVisualizationPieConnection) SetType(v string)`

SetType sets Type field to given value.


### GetConnectionId

`func (o *InputGraphVisualizationPieConnection) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *InputGraphVisualizationPieConnection) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *InputGraphVisualizationPieConnection) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetQueries

`func (o *InputGraphVisualizationPieConnection) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *InputGraphVisualizationPieConnection) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *InputGraphVisualizationPieConnection) SetQueries(v []string)`

SetQueries sets Queries field to given value.


### GetLegendMode

`func (o *InputGraphVisualizationPieConnection) GetLegendMode() string`

GetLegendMode returns the LegendMode field if non-nil, zero value otherwise.

### GetLegendModeOk

`func (o *InputGraphVisualizationPieConnection) GetLegendModeOk() (*string, bool)`

GetLegendModeOk returns a tuple with the LegendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendMode

`func (o *InputGraphVisualizationPieConnection) SetLegendMode(v string)`

SetLegendMode sets LegendMode field to given value.

### HasLegendMode

`func (o *InputGraphVisualizationPieConnection) HasLegendMode() bool`

HasLegendMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


