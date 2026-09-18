# GraphVisualizationListConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays database rows as a tabular list | 
**ConnectionId** | **string** | The ID of the connection to use to query the datastore. | 
**Query** | **string** | The read-only SQL query to execute against the connection. | 
**ListColumns** | Pointer to [**[]WidgetListColumn**](WidgetListColumn.md) | Custom columns to display for each database row | [optional] 
**ListColumnsSize** | Pointer to **map[string]float32** | List column widths in pixels, keyed by the &#x60;attribute&#x60; of the matching &#x60;listColumns&#x60; entry. Columns without an entry keep their default width. | [optional] 
**IsCellWrapped** | Pointer to **bool** | Whether list widget cell text wraps instead of truncating. Applies to connection list widgets. | [optional] 
**DefaultSorting** | Pointer to [**[]ListDefaultSorting**](ListDefaultSorting.md) | Default sorting applied to a list widget. Applies to log, span, or connection list widgets. Users can still change sorting by selecting columns in the rendered list. | [optional] 

## Methods

### NewGraphVisualizationListConnection

`func NewGraphVisualizationListConnection(type_ string, connectionId string, query string, ) *GraphVisualizationListConnection`

NewGraphVisualizationListConnection instantiates a new GraphVisualizationListConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationListConnectionWithDefaults

`func NewGraphVisualizationListConnectionWithDefaults() *GraphVisualizationListConnection`

NewGraphVisualizationListConnectionWithDefaults instantiates a new GraphVisualizationListConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GraphVisualizationListConnection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphVisualizationListConnection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphVisualizationListConnection) SetType(v string)`

SetType sets Type field to given value.


### GetConnectionId

`func (o *GraphVisualizationListConnection) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *GraphVisualizationListConnection) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *GraphVisualizationListConnection) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetQuery

`func (o *GraphVisualizationListConnection) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GraphVisualizationListConnection) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GraphVisualizationListConnection) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetListColumns

`func (o *GraphVisualizationListConnection) GetListColumns() []WidgetListColumn`

GetListColumns returns the ListColumns field if non-nil, zero value otherwise.

### GetListColumnsOk

`func (o *GraphVisualizationListConnection) GetListColumnsOk() (*[]WidgetListColumn, bool)`

GetListColumnsOk returns a tuple with the ListColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListColumns

`func (o *GraphVisualizationListConnection) SetListColumns(v []WidgetListColumn)`

SetListColumns sets ListColumns field to given value.

### HasListColumns

`func (o *GraphVisualizationListConnection) HasListColumns() bool`

HasListColumns returns a boolean if a field has been set.

### GetListColumnsSize

`func (o *GraphVisualizationListConnection) GetListColumnsSize() map[string]float32`

GetListColumnsSize returns the ListColumnsSize field if non-nil, zero value otherwise.

### GetListColumnsSizeOk

`func (o *GraphVisualizationListConnection) GetListColumnsSizeOk() (*map[string]float32, bool)`

GetListColumnsSizeOk returns a tuple with the ListColumnsSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListColumnsSize

`func (o *GraphVisualizationListConnection) SetListColumnsSize(v map[string]float32)`

SetListColumnsSize sets ListColumnsSize field to given value.

### HasListColumnsSize

`func (o *GraphVisualizationListConnection) HasListColumnsSize() bool`

HasListColumnsSize returns a boolean if a field has been set.

### GetIsCellWrapped

`func (o *GraphVisualizationListConnection) GetIsCellWrapped() bool`

GetIsCellWrapped returns the IsCellWrapped field if non-nil, zero value otherwise.

### GetIsCellWrappedOk

`func (o *GraphVisualizationListConnection) GetIsCellWrappedOk() (*bool, bool)`

GetIsCellWrappedOk returns a tuple with the IsCellWrapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCellWrapped

`func (o *GraphVisualizationListConnection) SetIsCellWrapped(v bool)`

SetIsCellWrapped sets IsCellWrapped field to given value.

### HasIsCellWrapped

`func (o *GraphVisualizationListConnection) HasIsCellWrapped() bool`

HasIsCellWrapped returns a boolean if a field has been set.

### GetDefaultSorting

`func (o *GraphVisualizationListConnection) GetDefaultSorting() []ListDefaultSorting`

GetDefaultSorting returns the DefaultSorting field if non-nil, zero value otherwise.

### GetDefaultSortingOk

`func (o *GraphVisualizationListConnection) GetDefaultSortingOk() (*[]ListDefaultSorting, bool)`

GetDefaultSortingOk returns a tuple with the DefaultSorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSorting

`func (o *GraphVisualizationListConnection) SetDefaultSorting(v []ListDefaultSorting)`

SetDefaultSorting sets DefaultSorting field to given value.

### HasDefaultSorting

`func (o *GraphVisualizationListConnection) HasDefaultSorting() bool`

HasDefaultSorting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


