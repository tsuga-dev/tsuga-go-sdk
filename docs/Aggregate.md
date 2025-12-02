# Aggregate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Counts the total number of records | 
**Field** | **string** | Attribute containing the values to aggregate | 
**Percentile** | **float32** |  | 

## Methods

### NewAggregate

`func NewAggregate(type_ string, field string, percentile float32, ) *Aggregate`

NewAggregate instantiates a new Aggregate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateWithDefaults

`func NewAggregateWithDefaults() *Aggregate`

NewAggregateWithDefaults instantiates a new Aggregate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Aggregate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Aggregate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Aggregate) SetType(v string)`

SetType sets Type field to given value.


### GetField

`func (o *Aggregate) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *Aggregate) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *Aggregate) SetField(v string)`

SetField sets Field field to given value.


### GetPercentile

`func (o *Aggregate) GetPercentile() float32`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *Aggregate) GetPercentileOk() (*float32, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *Aggregate) SetPercentile(v float32)`

SetPercentile sets Percentile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


