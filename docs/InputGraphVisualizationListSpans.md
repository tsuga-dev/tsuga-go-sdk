# InputGraphVisualizationListSpans

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays matching trace spans in a tabular list | 
**Query** | **string** | Query that selects trace spans for the list | 
**ListColumns** | Pointer to [**[]WidgetListColumn1**](WidgetListColumn1.md) | Custom columns to display for each span | [optional] 
**ListColumnsSize** | Pointer to **map[string]float32** | List column widths in pixels, keyed by the &#x60;attribute&#x60; of the matching &#x60;listColumns&#x60; entry. Columns without an entry keep their default width. | [optional] 
**IsCellWrapped** | Pointer to **bool** | Whether list widget cell text wraps instead of truncating. Applies to span list widgets. Optional; omit or set false to use truncated cells. | [optional] 
**DefaultSorting** | Pointer to [**[]ListDefaultSorting1**](ListDefaultSorting1.md) | Default sorting applied to a list widget. Optional on create or update for log, span, or connection list widgets. Users can still change sorting by selecting columns in the rendered list. | [optional] 

## Methods

### NewInputGraphVisualizationListSpans

`func NewInputGraphVisualizationListSpans(type_ string, query string, ) *InputGraphVisualizationListSpans`

NewInputGraphVisualizationListSpans instantiates a new InputGraphVisualizationListSpans object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationListSpansWithDefaults

`func NewInputGraphVisualizationListSpansWithDefaults() *InputGraphVisualizationListSpans`

NewInputGraphVisualizationListSpansWithDefaults instantiates a new InputGraphVisualizationListSpans object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputGraphVisualizationListSpans) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputGraphVisualizationListSpans) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputGraphVisualizationListSpans) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *InputGraphVisualizationListSpans) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *InputGraphVisualizationListSpans) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *InputGraphVisualizationListSpans) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetListColumns

`func (o *InputGraphVisualizationListSpans) GetListColumns() []WidgetListColumn1`

GetListColumns returns the ListColumns field if non-nil, zero value otherwise.

### GetListColumnsOk

`func (o *InputGraphVisualizationListSpans) GetListColumnsOk() (*[]WidgetListColumn1, bool)`

GetListColumnsOk returns a tuple with the ListColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListColumns

`func (o *InputGraphVisualizationListSpans) SetListColumns(v []WidgetListColumn1)`

SetListColumns sets ListColumns field to given value.

### HasListColumns

`func (o *InputGraphVisualizationListSpans) HasListColumns() bool`

HasListColumns returns a boolean if a field has been set.

### GetListColumnsSize

`func (o *InputGraphVisualizationListSpans) GetListColumnsSize() map[string]float32`

GetListColumnsSize returns the ListColumnsSize field if non-nil, zero value otherwise.

### GetListColumnsSizeOk

`func (o *InputGraphVisualizationListSpans) GetListColumnsSizeOk() (*map[string]float32, bool)`

GetListColumnsSizeOk returns a tuple with the ListColumnsSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListColumnsSize

`func (o *InputGraphVisualizationListSpans) SetListColumnsSize(v map[string]float32)`

SetListColumnsSize sets ListColumnsSize field to given value.

### HasListColumnsSize

`func (o *InputGraphVisualizationListSpans) HasListColumnsSize() bool`

HasListColumnsSize returns a boolean if a field has been set.

### GetIsCellWrapped

`func (o *InputGraphVisualizationListSpans) GetIsCellWrapped() bool`

GetIsCellWrapped returns the IsCellWrapped field if non-nil, zero value otherwise.

### GetIsCellWrappedOk

`func (o *InputGraphVisualizationListSpans) GetIsCellWrappedOk() (*bool, bool)`

GetIsCellWrappedOk returns a tuple with the IsCellWrapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCellWrapped

`func (o *InputGraphVisualizationListSpans) SetIsCellWrapped(v bool)`

SetIsCellWrapped sets IsCellWrapped field to given value.

### HasIsCellWrapped

`func (o *InputGraphVisualizationListSpans) HasIsCellWrapped() bool`

HasIsCellWrapped returns a boolean if a field has been set.

### GetDefaultSorting

`func (o *InputGraphVisualizationListSpans) GetDefaultSorting() []ListDefaultSorting1`

GetDefaultSorting returns the DefaultSorting field if non-nil, zero value otherwise.

### GetDefaultSortingOk

`func (o *InputGraphVisualizationListSpans) GetDefaultSortingOk() (*[]ListDefaultSorting1, bool)`

GetDefaultSortingOk returns a tuple with the DefaultSorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSorting

`func (o *InputGraphVisualizationListSpans) SetDefaultSorting(v []ListDefaultSorting1)`

SetDefaultSorting sets DefaultSorting field to given value.

### HasDefaultSorting

`func (o *InputGraphVisualizationListSpans) HasDefaultSorting() bool`

HasDefaultSorting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


