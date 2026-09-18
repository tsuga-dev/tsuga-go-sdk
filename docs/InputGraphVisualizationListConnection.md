# InputGraphVisualizationListConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays database rows as a tabular list | 
**ConnectionId** | **string** | The ID of the connection to use to query the datastore. | 
**Query** | **string** | The read-only SQL query to execute against the connection. | 
**ListColumns** | Pointer to [**[]WidgetListColumn1**](WidgetListColumn1.md) | Custom columns to display for each database row | [optional] 
**ListColumnsSize** | Pointer to **map[string]float32** | List column widths in pixels, keyed by the &#x60;attribute&#x60; of the matching &#x60;listColumns&#x60; entry. Columns without an entry keep their default width. | [optional] 
**IsCellWrapped** | Pointer to **bool** | Whether list widget cell text wraps instead of truncating. Applies to connection list widgets. Optional; omit or set false to use truncated cells. | [optional] 
**DefaultSorting** | Pointer to [**[]ListDefaultSorting1**](ListDefaultSorting1.md) | Default sorting applied to a list widget. Optional on create or update for log, span, or connection list widgets. Users can still change sorting by selecting columns in the rendered list. | [optional] 

## Methods

### NewInputGraphVisualizationListConnection

`func NewInputGraphVisualizationListConnection(type_ string, connectionId string, query string, ) *InputGraphVisualizationListConnection`

NewInputGraphVisualizationListConnection instantiates a new InputGraphVisualizationListConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationListConnectionWithDefaults

`func NewInputGraphVisualizationListConnectionWithDefaults() *InputGraphVisualizationListConnection`

NewInputGraphVisualizationListConnectionWithDefaults instantiates a new InputGraphVisualizationListConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputGraphVisualizationListConnection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputGraphVisualizationListConnection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputGraphVisualizationListConnection) SetType(v string)`

SetType sets Type field to given value.


### GetConnectionId

`func (o *InputGraphVisualizationListConnection) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *InputGraphVisualizationListConnection) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *InputGraphVisualizationListConnection) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetQuery

`func (o *InputGraphVisualizationListConnection) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *InputGraphVisualizationListConnection) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *InputGraphVisualizationListConnection) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetListColumns

`func (o *InputGraphVisualizationListConnection) GetListColumns() []WidgetListColumn1`

GetListColumns returns the ListColumns field if non-nil, zero value otherwise.

### GetListColumnsOk

`func (o *InputGraphVisualizationListConnection) GetListColumnsOk() (*[]WidgetListColumn1, bool)`

GetListColumnsOk returns a tuple with the ListColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListColumns

`func (o *InputGraphVisualizationListConnection) SetListColumns(v []WidgetListColumn1)`

SetListColumns sets ListColumns field to given value.

### HasListColumns

`func (o *InputGraphVisualizationListConnection) HasListColumns() bool`

HasListColumns returns a boolean if a field has been set.

### GetListColumnsSize

`func (o *InputGraphVisualizationListConnection) GetListColumnsSize() map[string]float32`

GetListColumnsSize returns the ListColumnsSize field if non-nil, zero value otherwise.

### GetListColumnsSizeOk

`func (o *InputGraphVisualizationListConnection) GetListColumnsSizeOk() (*map[string]float32, bool)`

GetListColumnsSizeOk returns a tuple with the ListColumnsSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListColumnsSize

`func (o *InputGraphVisualizationListConnection) SetListColumnsSize(v map[string]float32)`

SetListColumnsSize sets ListColumnsSize field to given value.

### HasListColumnsSize

`func (o *InputGraphVisualizationListConnection) HasListColumnsSize() bool`

HasListColumnsSize returns a boolean if a field has been set.

### GetIsCellWrapped

`func (o *InputGraphVisualizationListConnection) GetIsCellWrapped() bool`

GetIsCellWrapped returns the IsCellWrapped field if non-nil, zero value otherwise.

### GetIsCellWrappedOk

`func (o *InputGraphVisualizationListConnection) GetIsCellWrappedOk() (*bool, bool)`

GetIsCellWrappedOk returns a tuple with the IsCellWrapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCellWrapped

`func (o *InputGraphVisualizationListConnection) SetIsCellWrapped(v bool)`

SetIsCellWrapped sets IsCellWrapped field to given value.

### HasIsCellWrapped

`func (o *InputGraphVisualizationListConnection) HasIsCellWrapped() bool`

HasIsCellWrapped returns a boolean if a field has been set.

### GetDefaultSorting

`func (o *InputGraphVisualizationListConnection) GetDefaultSorting() []ListDefaultSorting1`

GetDefaultSorting returns the DefaultSorting field if non-nil, zero value otherwise.

### GetDefaultSortingOk

`func (o *InputGraphVisualizationListConnection) GetDefaultSortingOk() (*[]ListDefaultSorting1, bool)`

GetDefaultSortingOk returns a tuple with the DefaultSorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSorting

`func (o *InputGraphVisualizationListConnection) SetDefaultSorting(v []ListDefaultSorting1)`

SetDefaultSorting sets DefaultSorting field to given value.

### HasDefaultSorting

`func (o *InputGraphVisualizationListConnection) HasDefaultSorting() bool`

HasDefaultSorting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


