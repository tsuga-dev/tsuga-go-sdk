# InputGraphVisualizationListRum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays matching RUM events in a tabular list | 
**Query** | **string** | Query that selects RUM events for the list. Prefix every field with the stream it belongs to (&#x60;measurements.&#x60;, &#x60;events.&#x60;, or &#x60;exceptions.&#x60;) and filter on a single stream per query. A query without stream-prefixed fields searches the &#x60;events&#x60; stream. | 
**ListColumns** | Pointer to [**[]InputWidgetListColumn**](InputWidgetListColumn.md) | Custom columns to display for each RUM event. Set attributes without their stream prefix, e.g. &#x60;view.name&#x60; rather than &#x60;events.view.name&#x60;. | [optional] 
**ListColumnsSize** | Pointer to **map[string]float32** | List column widths in pixels, keyed by the &#x60;attribute&#x60; of the matching &#x60;listColumns&#x60; entry. Columns without an entry keep their default width. | [optional] 
**IsCellWrapped** | Pointer to **bool** | Whether list widget cell text wraps instead of truncating. Applies to RUM event list widgets. Optional; omit or set false to use truncated cells. | [optional] 
**DefaultSorting** | Pointer to [**[]InputListDefaultSorting**](InputListDefaultSorting.md) | Default sorting applied to a list widget. Optional on create or update for log, span, RUM event, or connection list widgets. Users can still change sorting by selecting columns in the rendered list. | [optional] 

## Methods

### NewInputGraphVisualizationListRum

`func NewInputGraphVisualizationListRum(type_ string, query string, ) *InputGraphVisualizationListRum`

NewInputGraphVisualizationListRum instantiates a new InputGraphVisualizationListRum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationListRumWithDefaults

`func NewInputGraphVisualizationListRumWithDefaults() *InputGraphVisualizationListRum`

NewInputGraphVisualizationListRumWithDefaults instantiates a new InputGraphVisualizationListRum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputGraphVisualizationListRum) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputGraphVisualizationListRum) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputGraphVisualizationListRum) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *InputGraphVisualizationListRum) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *InputGraphVisualizationListRum) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *InputGraphVisualizationListRum) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetListColumns

`func (o *InputGraphVisualizationListRum) GetListColumns() []InputWidgetListColumn`

GetListColumns returns the ListColumns field if non-nil, zero value otherwise.

### GetListColumnsOk

`func (o *InputGraphVisualizationListRum) GetListColumnsOk() (*[]InputWidgetListColumn, bool)`

GetListColumnsOk returns a tuple with the ListColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListColumns

`func (o *InputGraphVisualizationListRum) SetListColumns(v []InputWidgetListColumn)`

SetListColumns sets ListColumns field to given value.

### HasListColumns

`func (o *InputGraphVisualizationListRum) HasListColumns() bool`

HasListColumns returns a boolean if a field has been set.

### GetListColumnsSize

`func (o *InputGraphVisualizationListRum) GetListColumnsSize() map[string]float32`

GetListColumnsSize returns the ListColumnsSize field if non-nil, zero value otherwise.

### GetListColumnsSizeOk

`func (o *InputGraphVisualizationListRum) GetListColumnsSizeOk() (*map[string]float32, bool)`

GetListColumnsSizeOk returns a tuple with the ListColumnsSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListColumnsSize

`func (o *InputGraphVisualizationListRum) SetListColumnsSize(v map[string]float32)`

SetListColumnsSize sets ListColumnsSize field to given value.

### HasListColumnsSize

`func (o *InputGraphVisualizationListRum) HasListColumnsSize() bool`

HasListColumnsSize returns a boolean if a field has been set.

### GetIsCellWrapped

`func (o *InputGraphVisualizationListRum) GetIsCellWrapped() bool`

GetIsCellWrapped returns the IsCellWrapped field if non-nil, zero value otherwise.

### GetIsCellWrappedOk

`func (o *InputGraphVisualizationListRum) GetIsCellWrappedOk() (*bool, bool)`

GetIsCellWrappedOk returns a tuple with the IsCellWrapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCellWrapped

`func (o *InputGraphVisualizationListRum) SetIsCellWrapped(v bool)`

SetIsCellWrapped sets IsCellWrapped field to given value.

### HasIsCellWrapped

`func (o *InputGraphVisualizationListRum) HasIsCellWrapped() bool`

HasIsCellWrapped returns a boolean if a field has been set.

### GetDefaultSorting

`func (o *InputGraphVisualizationListRum) GetDefaultSorting() []InputListDefaultSorting`

GetDefaultSorting returns the DefaultSorting field if non-nil, zero value otherwise.

### GetDefaultSortingOk

`func (o *InputGraphVisualizationListRum) GetDefaultSortingOk() (*[]InputListDefaultSorting, bool)`

GetDefaultSortingOk returns a tuple with the DefaultSorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSorting

`func (o *InputGraphVisualizationListRum) SetDefaultSorting(v []InputListDefaultSorting)`

SetDefaultSorting sets DefaultSorting field to given value.

### HasDefaultSorting

`func (o *InputGraphVisualizationListRum) HasDefaultSorting() bool`

HasDefaultSorting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


