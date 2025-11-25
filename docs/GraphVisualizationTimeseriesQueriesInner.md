# GraphVisualizationTimeseriesQueriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregate** | [**GraphVisualizationTimeseriesQueriesInnerAggregate**](GraphVisualizationTimeseriesQueriesInnerAggregate.md) |  | 
**Filter** | Pointer to **string** |  | [optional] 
**Functions** | Pointer to [**[]GraphVisualizationTimeseriesQueriesInnerFunctionsInner**](GraphVisualizationTimeseriesQueriesInnerFunctionsInner.md) | Post-processing functions applied to aggregation results | [optional] 

## Methods

### NewGraphVisualizationTimeseriesQueriesInner

`func NewGraphVisualizationTimeseriesQueriesInner(aggregate GraphVisualizationTimeseriesQueriesInnerAggregate, ) *GraphVisualizationTimeseriesQueriesInner`

NewGraphVisualizationTimeseriesQueriesInner instantiates a new GraphVisualizationTimeseriesQueriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationTimeseriesQueriesInnerWithDefaults

`func NewGraphVisualizationTimeseriesQueriesInnerWithDefaults() *GraphVisualizationTimeseriesQueriesInner`

NewGraphVisualizationTimeseriesQueriesInnerWithDefaults instantiates a new GraphVisualizationTimeseriesQueriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregate

`func (o *GraphVisualizationTimeseriesQueriesInner) GetAggregate() GraphVisualizationTimeseriesQueriesInnerAggregate`

GetAggregate returns the Aggregate field if non-nil, zero value otherwise.

### GetAggregateOk

`func (o *GraphVisualizationTimeseriesQueriesInner) GetAggregateOk() (*GraphVisualizationTimeseriesQueriesInnerAggregate, bool)`

GetAggregateOk returns a tuple with the Aggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregate

`func (o *GraphVisualizationTimeseriesQueriesInner) SetAggregate(v GraphVisualizationTimeseriesQueriesInnerAggregate)`

SetAggregate sets Aggregate field to given value.


### GetFilter

`func (o *GraphVisualizationTimeseriesQueriesInner) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *GraphVisualizationTimeseriesQueriesInner) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *GraphVisualizationTimeseriesQueriesInner) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *GraphVisualizationTimeseriesQueriesInner) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFunctions

`func (o *GraphVisualizationTimeseriesQueriesInner) GetFunctions() []GraphVisualizationTimeseriesQueriesInnerFunctionsInner`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *GraphVisualizationTimeseriesQueriesInner) GetFunctionsOk() (*[]GraphVisualizationTimeseriesQueriesInnerFunctionsInner, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *GraphVisualizationTimeseriesQueriesInner) SetFunctions(v []GraphVisualizationTimeseriesQueriesInnerFunctionsInner)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *GraphVisualizationTimeseriesQueriesInner) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


