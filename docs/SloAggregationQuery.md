# SloAggregationQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregate** | [**InputAggregate**](InputAggregate.md) |  | 
**Functions** | Pointer to [**[]InputFunction**](InputFunction.md) | Post-processing functions applied to aggregation results. Defaults to an empty array when omitted. | [optional] 
**TimeAggregate** | Pointer to **string** | Per-series rollup applied within each time bucket before the cross-series aggregate. Use it on metric queries when &#x60;aggregate.type&#x60; is &#x60;sum&#x60; and no &#x60;rate&#x60;, &#x60;increase&#x60;, &#x60;last&#x60;, or &#x60;rolling&#x60; function is present. When omitted, Tsuga derives the rollup from the metric type. | [optional] 
**Filter** | **string** | Tsuga query filter applied to this alerting or SLO aggregation query. Required by the API; use the query language for the selected data source. | 

## Methods

### NewSloAggregationQuery

`func NewSloAggregationQuery(aggregate InputAggregate, filter string, ) *SloAggregationQuery`

NewSloAggregationQuery instantiates a new SloAggregationQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloAggregationQueryWithDefaults

`func NewSloAggregationQueryWithDefaults() *SloAggregationQuery`

NewSloAggregationQueryWithDefaults instantiates a new SloAggregationQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregate

`func (o *SloAggregationQuery) GetAggregate() InputAggregate`

GetAggregate returns the Aggregate field if non-nil, zero value otherwise.

### GetAggregateOk

`func (o *SloAggregationQuery) GetAggregateOk() (*InputAggregate, bool)`

GetAggregateOk returns a tuple with the Aggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregate

`func (o *SloAggregationQuery) SetAggregate(v InputAggregate)`

SetAggregate sets Aggregate field to given value.


### GetFunctions

`func (o *SloAggregationQuery) GetFunctions() []InputFunction`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *SloAggregationQuery) GetFunctionsOk() (*[]InputFunction, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *SloAggregationQuery) SetFunctions(v []InputFunction)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *SloAggregationQuery) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### GetTimeAggregate

`func (o *SloAggregationQuery) GetTimeAggregate() string`

GetTimeAggregate returns the TimeAggregate field if non-nil, zero value otherwise.

### GetTimeAggregateOk

`func (o *SloAggregationQuery) GetTimeAggregateOk() (*string, bool)`

GetTimeAggregateOk returns a tuple with the TimeAggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAggregate

`func (o *SloAggregationQuery) SetTimeAggregate(v string)`

SetTimeAggregate sets TimeAggregate field to given value.

### HasTimeAggregate

`func (o *SloAggregationQuery) HasTimeAggregate() bool`

HasTimeAggregate returns a boolean if a field has been set.

### GetFilter

`func (o *SloAggregationQuery) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SloAggregationQuery) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SloAggregationQuery) SetFilter(v string)`

SetFilter sets Filter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


