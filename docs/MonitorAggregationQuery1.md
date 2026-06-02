# MonitorAggregationQuery1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregate** | [**InputAggregate**](InputAggregate.md) |  | 
**Functions** | Pointer to [**[]InputFunction**](InputFunction.md) | Post-processing functions applied to aggregation results | [optional] 
**Fill** | Pointer to [**MonitorAggregationQueryFill**](MonitorAggregationQueryFill.md) |  | [optional] 
**Filter** | **string** |  | 

## Methods

### NewMonitorAggregationQuery1

`func NewMonitorAggregationQuery1(aggregate InputAggregate, filter string, ) *MonitorAggregationQuery1`

NewMonitorAggregationQuery1 instantiates a new MonitorAggregationQuery1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorAggregationQuery1WithDefaults

`func NewMonitorAggregationQuery1WithDefaults() *MonitorAggregationQuery1`

NewMonitorAggregationQuery1WithDefaults instantiates a new MonitorAggregationQuery1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregate

`func (o *MonitorAggregationQuery1) GetAggregate() InputAggregate`

GetAggregate returns the Aggregate field if non-nil, zero value otherwise.

### GetAggregateOk

`func (o *MonitorAggregationQuery1) GetAggregateOk() (*InputAggregate, bool)`

GetAggregateOk returns a tuple with the Aggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregate

`func (o *MonitorAggregationQuery1) SetAggregate(v InputAggregate)`

SetAggregate sets Aggregate field to given value.


### GetFunctions

`func (o *MonitorAggregationQuery1) GetFunctions() []InputFunction`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *MonitorAggregationQuery1) GetFunctionsOk() (*[]InputFunction, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *MonitorAggregationQuery1) SetFunctions(v []InputFunction)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *MonitorAggregationQuery1) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### GetFill

`func (o *MonitorAggregationQuery1) GetFill() MonitorAggregationQueryFill`

GetFill returns the Fill field if non-nil, zero value otherwise.

### GetFillOk

`func (o *MonitorAggregationQuery1) GetFillOk() (*MonitorAggregationQueryFill, bool)`

GetFillOk returns a tuple with the Fill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFill

`func (o *MonitorAggregationQuery1) SetFill(v MonitorAggregationQueryFill)`

SetFill sets Fill field to given value.

### HasFill

`func (o *MonitorAggregationQuery1) HasFill() bool`

HasFill returns a boolean if a field has been set.

### GetFilter

`func (o *MonitorAggregationQuery1) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MonitorAggregationQuery1) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MonitorAggregationQuery1) SetFilter(v string)`

SetFilter sets Filter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


