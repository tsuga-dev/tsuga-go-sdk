# InputGraphVisualizationTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the aggregation as a table with multi-level grouping | 
**Columns** | [**[]InputTableColumn**](InputTableColumn.md) | Each column defines an independent aggregation displayed as a table column | 
**GroupBy** | Pointer to [**[]InputAggregationGroupBy**](InputAggregationGroupBy.md) | Nested grouping levels applied to aggregation results, outermost first (e.g. group by service, then by level within each service). Each level splits results further, so the response contains one result per unique combination of group values instead of one aggregated total. Defaults to an empty array (one ungrouped result) when omitted. | [optional] 
**GroupByMode** | Pointer to **string** | &#x60;absolute&#x60; keeps each group at its own value; &#x60;relative&#x60; shows it as a percentage of the ungrouped total (defaults to absolute) | [optional] 
**DefaultSorting** | Pointer to [**[]InputTableDefaultSorting**](InputTableDefaultSorting.md) | Default sorting applied to a table widget. Column IDs are &#x60;label&#x60; for the grouping column and &#x60;col-&lt;index&gt;&#x60; for each entry in &#x60;columns&#x60;. Users can still change sorting by selecting columns in the rendered table. | [optional] 
**ColumnSizes** | Pointer to **map[string]float32** | Table column widths in pixels, keyed by column id: &#x60;label&#x60; for the grouping column and &#x60;col-&lt;index&gt;&#x60; for each entry in &#x60;columns&#x60;. Columns without an entry keep their default width. | [optional] 
**CustomLink** | Pointer to [**InputCustomLink**](InputCustomLink.md) |  | [optional] 

## Methods

### NewInputGraphVisualizationTable

`func NewInputGraphVisualizationTable(type_ string, columns []InputTableColumn, ) *InputGraphVisualizationTable`

NewInputGraphVisualizationTable instantiates a new InputGraphVisualizationTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationTableWithDefaults

`func NewInputGraphVisualizationTableWithDefaults() *InputGraphVisualizationTable`

NewInputGraphVisualizationTableWithDefaults instantiates a new InputGraphVisualizationTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputGraphVisualizationTable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputGraphVisualizationTable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputGraphVisualizationTable) SetType(v string)`

SetType sets Type field to given value.


### GetColumns

`func (o *InputGraphVisualizationTable) GetColumns() []InputTableColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *InputGraphVisualizationTable) GetColumnsOk() (*[]InputTableColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *InputGraphVisualizationTable) SetColumns(v []InputTableColumn)`

SetColumns sets Columns field to given value.


### GetGroupBy

`func (o *InputGraphVisualizationTable) GetGroupBy() []InputAggregationGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *InputGraphVisualizationTable) GetGroupByOk() (*[]InputAggregationGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *InputGraphVisualizationTable) SetGroupBy(v []InputAggregationGroupBy)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *InputGraphVisualizationTable) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetGroupByMode

`func (o *InputGraphVisualizationTable) GetGroupByMode() string`

GetGroupByMode returns the GroupByMode field if non-nil, zero value otherwise.

### GetGroupByModeOk

`func (o *InputGraphVisualizationTable) GetGroupByModeOk() (*string, bool)`

GetGroupByModeOk returns a tuple with the GroupByMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByMode

`func (o *InputGraphVisualizationTable) SetGroupByMode(v string)`

SetGroupByMode sets GroupByMode field to given value.

### HasGroupByMode

`func (o *InputGraphVisualizationTable) HasGroupByMode() bool`

HasGroupByMode returns a boolean if a field has been set.

### GetDefaultSorting

`func (o *InputGraphVisualizationTable) GetDefaultSorting() []InputTableDefaultSorting`

GetDefaultSorting returns the DefaultSorting field if non-nil, zero value otherwise.

### GetDefaultSortingOk

`func (o *InputGraphVisualizationTable) GetDefaultSortingOk() (*[]InputTableDefaultSorting, bool)`

GetDefaultSortingOk returns a tuple with the DefaultSorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSorting

`func (o *InputGraphVisualizationTable) SetDefaultSorting(v []InputTableDefaultSorting)`

SetDefaultSorting sets DefaultSorting field to given value.

### HasDefaultSorting

`func (o *InputGraphVisualizationTable) HasDefaultSorting() bool`

HasDefaultSorting returns a boolean if a field has been set.

### GetColumnSizes

`func (o *InputGraphVisualizationTable) GetColumnSizes() map[string]float32`

GetColumnSizes returns the ColumnSizes field if non-nil, zero value otherwise.

### GetColumnSizesOk

`func (o *InputGraphVisualizationTable) GetColumnSizesOk() (*map[string]float32, bool)`

GetColumnSizesOk returns a tuple with the ColumnSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnSizes

`func (o *InputGraphVisualizationTable) SetColumnSizes(v map[string]float32)`

SetColumnSizes sets ColumnSizes field to given value.

### HasColumnSizes

`func (o *InputGraphVisualizationTable) HasColumnSizes() bool`

HasColumnSizes returns a boolean if a field has been set.

### GetCustomLink

`func (o *InputGraphVisualizationTable) GetCustomLink() InputCustomLink`

GetCustomLink returns the CustomLink field if non-nil, zero value otherwise.

### GetCustomLinkOk

`func (o *InputGraphVisualizationTable) GetCustomLinkOk() (*InputCustomLink, bool)`

GetCustomLinkOk returns a tuple with the CustomLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLink

`func (o *InputGraphVisualizationTable) SetCustomLink(v InputCustomLink)`

SetCustomLink sets CustomLink field to given value.

### HasCustomLink

`func (o *InputGraphVisualizationTable) HasCustomLink() bool`

HasCustomLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


