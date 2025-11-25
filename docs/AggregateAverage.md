# AggregateAverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Calculates the average value of the selected field | 
**Field** | **string** | Attribute containing the values to aggregate | 

## Methods

### NewAggregateAverage

`func NewAggregateAverage(type_ string, field string, ) *AggregateAverage`

NewAggregateAverage instantiates a new AggregateAverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateAverageWithDefaults

`func NewAggregateAverageWithDefaults() *AggregateAverage`

NewAggregateAverageWithDefaults instantiates a new AggregateAverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AggregateAverage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AggregateAverage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AggregateAverage) SetType(v string)`

SetType sets Type field to given value.


### GetField

`func (o *AggregateAverage) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AggregateAverage) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AggregateAverage) SetField(v string)`

SetField sets Field field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


