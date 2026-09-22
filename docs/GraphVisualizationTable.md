# GraphVisualizationTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the aggregation as a table with multi-level grouping | 
**Columns** | [**[]TableColumn**](TableColumn.md) | Each column defines an independent aggregation displayed as a table column | 
**GroupBy** | Pointer to [**[]AggregationGroupBy1**](AggregationGroupBy1.md) | Nested grouping levels applied to the results, outermost first (e.g. group by service, then by level within each service). Each level splits results further, so the response contains one result per unique combination of group values. | [optional] 
**GroupByMode** | Pointer to **string** | &#x60;absolute&#x60; keeps each group at its own value; &#x60;relative&#x60; shows it as a percentage of the ungrouped total (defaults to absolute) | [optional] 
**DefaultSorting** | Pointer to [**[]TableDefaultSorting**](TableDefaultSorting.md) | Default sorting applied to a table widget. Column IDs are &#x60;label&#x60; for the grouping column and &#x60;col-&lt;index&gt;&#x60; for each entry in &#x60;columns&#x60;. Users can still change sorting by selecting columns in the rendered table. | [optional] 
**ColumnSizes** | Pointer to **map[string]float32** | Table column widths in pixels, keyed by column id: &#x60;label&#x60; for the grouping column and &#x60;col-&lt;index&gt;&#x60; for each entry in &#x60;columns&#x60;. Columns without an entry keep their default width. | [optional] 

## Methods

### NewGraphVisualizationTable

`func NewGraphVisualizationTable(type_ string, columns []TableColumn, ) *GraphVisualizationTable`

NewGraphVisualizationTable instantiates a new GraphVisualizationTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationTableWithDefaults

`func NewGraphVisualizationTableWithDefaults() *GraphVisualizationTable`

NewGraphVisualizationTableWithDefaults instantiates a new GraphVisualizationTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GraphVisualizationTable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphVisualizationTable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphVisualizationTable) SetType(v string)`

SetType sets Type field to given value.


### GetColumns

`func (o *GraphVisualizationTable) GetColumns() []TableColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *GraphVisualizationTable) GetColumnsOk() (*[]TableColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *GraphVisualizationTable) SetColumns(v []TableColumn)`

SetColumns sets Columns field to given value.


### GetGroupBy

`func (o *GraphVisualizationTable) GetGroupBy() []AggregationGroupBy1`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *GraphVisualizationTable) GetGroupByOk() (*[]AggregationGroupBy1, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *GraphVisualizationTable) SetGroupBy(v []AggregationGroupBy1)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *GraphVisualizationTable) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetGroupByMode

`func (o *GraphVisualizationTable) GetGroupByMode() string`

GetGroupByMode returns the GroupByMode field if non-nil, zero value otherwise.

### GetGroupByModeOk

`func (o *GraphVisualizationTable) GetGroupByModeOk() (*string, bool)`

GetGroupByModeOk returns a tuple with the GroupByMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByMode

`func (o *GraphVisualizationTable) SetGroupByMode(v string)`

SetGroupByMode sets GroupByMode field to given value.

### HasGroupByMode

`func (o *GraphVisualizationTable) HasGroupByMode() bool`

HasGroupByMode returns a boolean if a field has been set.

### GetDefaultSorting

`func (o *GraphVisualizationTable) GetDefaultSorting() []TableDefaultSorting`

GetDefaultSorting returns the DefaultSorting field if non-nil, zero value otherwise.

### GetDefaultSortingOk

`func (o *GraphVisualizationTable) GetDefaultSortingOk() (*[]TableDefaultSorting, bool)`

GetDefaultSortingOk returns a tuple with the DefaultSorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSorting

`func (o *GraphVisualizationTable) SetDefaultSorting(v []TableDefaultSorting)`

SetDefaultSorting sets DefaultSorting field to given value.

### HasDefaultSorting

`func (o *GraphVisualizationTable) HasDefaultSorting() bool`

HasDefaultSorting returns a boolean if a field has been set.

### GetColumnSizes

`func (o *GraphVisualizationTable) GetColumnSizes() map[string]float32`

GetColumnSizes returns the ColumnSizes field if non-nil, zero value otherwise.

### GetColumnSizesOk

`func (o *GraphVisualizationTable) GetColumnSizesOk() (*map[string]float32, bool)`

GetColumnSizesOk returns a tuple with the ColumnSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnSizes

`func (o *GraphVisualizationTable) SetColumnSizes(v map[string]float32)`

SetColumnSizes sets ColumnSizes field to given value.

### HasColumnSizes

`func (o *GraphVisualizationTable) HasColumnSizes() bool`

HasColumnSizes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


