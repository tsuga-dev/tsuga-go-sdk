# InputGraphVisualizationTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the aggregation as a table with multi-level grouping | 
**Columns** | [**[]TableColumn1**](TableColumn1.md) | Each column defines an independent aggregation displayed as a table column | 
**GroupBy** | Pointer to [**[]AggregationGroupBy1**](AggregationGroupBy1.md) | Fields used to group the results | [optional] 
**DefaultSorting** | Pointer to [**[]TableDefaultSorting1**](TableDefaultSorting1.md) |  | [optional] 

## Methods

### NewInputGraphVisualizationTable

`func NewInputGraphVisualizationTable(type_ string, columns []TableColumn1, ) *InputGraphVisualizationTable`

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

`func (o *InputGraphVisualizationTable) GetColumns() []TableColumn1`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *InputGraphVisualizationTable) GetColumnsOk() (*[]TableColumn1, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *InputGraphVisualizationTable) SetColumns(v []TableColumn1)`

SetColumns sets Columns field to given value.


### GetGroupBy

`func (o *InputGraphVisualizationTable) GetGroupBy() []AggregationGroupBy1`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *InputGraphVisualizationTable) GetGroupByOk() (*[]AggregationGroupBy1, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *InputGraphVisualizationTable) SetGroupBy(v []AggregationGroupBy1)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *InputGraphVisualizationTable) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetDefaultSorting

`func (o *InputGraphVisualizationTable) GetDefaultSorting() []TableDefaultSorting1`

GetDefaultSorting returns the DefaultSorting field if non-nil, zero value otherwise.

### GetDefaultSortingOk

`func (o *InputGraphVisualizationTable) GetDefaultSortingOk() (*[]TableDefaultSorting1, bool)`

GetDefaultSortingOk returns a tuple with the DefaultSorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSorting

`func (o *InputGraphVisualizationTable) SetDefaultSorting(v []TableDefaultSorting1)`

SetDefaultSorting sets DefaultSorting field to given value.

### HasDefaultSorting

`func (o *InputGraphVisualizationTable) HasDefaultSorting() bool`

HasDefaultSorting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


