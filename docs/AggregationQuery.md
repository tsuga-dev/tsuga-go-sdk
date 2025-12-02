# AggregationQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregate** | [**Aggregate**](Aggregate.md) |  | 
**Filter** | Pointer to **string** |  | [optional] 
**Functions** | Pointer to [**[]Function**](Function.md) | Post-processing functions applied to aggregation results | [optional] 

## Methods

### NewAggregationQuery

`func NewAggregationQuery(aggregate Aggregate, ) *AggregationQuery`

NewAggregationQuery instantiates a new AggregationQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationQueryWithDefaults

`func NewAggregationQueryWithDefaults() *AggregationQuery`

NewAggregationQueryWithDefaults instantiates a new AggregationQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregate

`func (o *AggregationQuery) GetAggregate() Aggregate`

GetAggregate returns the Aggregate field if non-nil, zero value otherwise.

### GetAggregateOk

`func (o *AggregationQuery) GetAggregateOk() (*Aggregate, bool)`

GetAggregateOk returns a tuple with the Aggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregate

`func (o *AggregationQuery) SetAggregate(v Aggregate)`

SetAggregate sets Aggregate field to given value.


### GetFilter

`func (o *AggregationQuery) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AggregationQuery) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AggregationQuery) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AggregationQuery) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFunctions

`func (o *AggregationQuery) GetFunctions() []Function`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *AggregationQuery) GetFunctionsOk() (*[]Function, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *AggregationQuery) SetFunctions(v []Function)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *AggregationQuery) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


