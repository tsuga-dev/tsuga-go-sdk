# GraphVisualizationTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the aggregation as a table with multi-level grouping | 
**Columns** | [**[]TableColumn**](TableColumn.md) | Each column defines an independent aggregation displayed as a table column | 
**GroupBy** | Pointer to [**[]AggregationGroupBy**](AggregationGroupBy.md) | Fields used to group the results | [optional] 
**DefaultSorting** | Pointer to [**[]TableDefaultSorting**](TableDefaultSorting.md) |  | [optional] 

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

`func (o *GraphVisualizationTable) GetGroupBy() []AggregationGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *GraphVisualizationTable) GetGroupByOk() (*[]AggregationGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *GraphVisualizationTable) SetGroupBy(v []AggregationGroupBy)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *GraphVisualizationTable) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


