# LogAggregate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Counts the total number of records | 
**Field** | **string** | Attribute containing the values to aggregate | 
**Percentile** | **float32** |  | 

## Methods

### NewLogAggregate

`func NewLogAggregate(type_ string, field string, percentile float32, ) *LogAggregate`

NewLogAggregate instantiates a new LogAggregate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogAggregateWithDefaults

`func NewLogAggregateWithDefaults() *LogAggregate`

NewLogAggregateWithDefaults instantiates a new LogAggregate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LogAggregate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogAggregate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogAggregate) SetType(v string)`

SetType sets Type field to given value.


### GetField

`func (o *LogAggregate) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *LogAggregate) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *LogAggregate) SetField(v string)`

SetField sets Field field to given value.


### GetPercentile

`func (o *LogAggregate) GetPercentile() float32`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *LogAggregate) GetPercentileOk() (*float32, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *LogAggregate) SetPercentile(v float32)`

SetPercentile sets Percentile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


