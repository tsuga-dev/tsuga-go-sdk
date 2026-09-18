# GraphVisualizationListSpans

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays matching trace spans in a tabular list | 
**Query** | **string** | Query that selects trace spans for the list | 
**ListColumns** | Pointer to [**[]WidgetListColumn**](WidgetListColumn.md) | Custom columns to display for each span | [optional] 
**ListColumnsSize** | Pointer to **map[string]float32** | List column widths in pixels, keyed by the &#x60;attribute&#x60; of the matching &#x60;listColumns&#x60; entry. Columns without an entry keep their default width. | [optional] 
**IsCellWrapped** | Pointer to **bool** | Whether list widget cell text wraps instead of truncating. Applies to span list widgets. | [optional] 
**DefaultSorting** | Pointer to [**[]ListDefaultSorting**](ListDefaultSorting.md) | Default sorting applied to a list widget. Applies to log, span, or connection list widgets. Users can still change sorting by selecting columns in the rendered list. | [optional] 

## Methods

### NewGraphVisualizationListSpans

`func NewGraphVisualizationListSpans(type_ string, query string, ) *GraphVisualizationListSpans`

NewGraphVisualizationListSpans instantiates a new GraphVisualizationListSpans object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationListSpansWithDefaults

`func NewGraphVisualizationListSpansWithDefaults() *GraphVisualizationListSpans`

NewGraphVisualizationListSpansWithDefaults instantiates a new GraphVisualizationListSpans object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GraphVisualizationListSpans) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphVisualizationListSpans) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphVisualizationListSpans) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *GraphVisualizationListSpans) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GraphVisualizationListSpans) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GraphVisualizationListSpans) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetListColumns

`func (o *GraphVisualizationListSpans) GetListColumns() []WidgetListColumn`

GetListColumns returns the ListColumns field if non-nil, zero value otherwise.

### GetListColumnsOk

`func (o *GraphVisualizationListSpans) GetListColumnsOk() (*[]WidgetListColumn, bool)`

GetListColumnsOk returns a tuple with the ListColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListColumns

`func (o *GraphVisualizationListSpans) SetListColumns(v []WidgetListColumn)`

SetListColumns sets ListColumns field to given value.

### HasListColumns

`func (o *GraphVisualizationListSpans) HasListColumns() bool`

HasListColumns returns a boolean if a field has been set.

### GetListColumnsSize

`func (o *GraphVisualizationListSpans) GetListColumnsSize() map[string]float32`

GetListColumnsSize returns the ListColumnsSize field if non-nil, zero value otherwise.

### GetListColumnsSizeOk

`func (o *GraphVisualizationListSpans) GetListColumnsSizeOk() (*map[string]float32, bool)`

GetListColumnsSizeOk returns a tuple with the ListColumnsSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListColumnsSize

`func (o *GraphVisualizationListSpans) SetListColumnsSize(v map[string]float32)`

SetListColumnsSize sets ListColumnsSize field to given value.

### HasListColumnsSize

`func (o *GraphVisualizationListSpans) HasListColumnsSize() bool`

HasListColumnsSize returns a boolean if a field has been set.

### GetIsCellWrapped

`func (o *GraphVisualizationListSpans) GetIsCellWrapped() bool`

GetIsCellWrapped returns the IsCellWrapped field if non-nil, zero value otherwise.

### GetIsCellWrappedOk

`func (o *GraphVisualizationListSpans) GetIsCellWrappedOk() (*bool, bool)`

GetIsCellWrappedOk returns a tuple with the IsCellWrapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCellWrapped

`func (o *GraphVisualizationListSpans) SetIsCellWrapped(v bool)`

SetIsCellWrapped sets IsCellWrapped field to given value.

### HasIsCellWrapped

`func (o *GraphVisualizationListSpans) HasIsCellWrapped() bool`

HasIsCellWrapped returns a boolean if a field has been set.

### GetDefaultSorting

`func (o *GraphVisualizationListSpans) GetDefaultSorting() []ListDefaultSorting`

GetDefaultSorting returns the DefaultSorting field if non-nil, zero value otherwise.

### GetDefaultSortingOk

`func (o *GraphVisualizationListSpans) GetDefaultSortingOk() (*[]ListDefaultSorting, bool)`

GetDefaultSortingOk returns a tuple with the DefaultSorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSorting

`func (o *GraphVisualizationListSpans) SetDefaultSorting(v []ListDefaultSorting)`

SetDefaultSorting sets DefaultSorting field to given value.

### HasDefaultSorting

`func (o *GraphVisualizationListSpans) HasDefaultSorting() bool`

HasDefaultSorting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


