# AggregateSum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Sums the values of the selected field | 
**Field** | **string** | Attribute containing the values to aggregate | 

## Methods

### NewAggregateSum

`func NewAggregateSum(type_ string, field string, ) *AggregateSum`

NewAggregateSum instantiates a new AggregateSum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateSumWithDefaults

`func NewAggregateSumWithDefaults() *AggregateSum`

NewAggregateSumWithDefaults instantiates a new AggregateSum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AggregateSum) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AggregateSum) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AggregateSum) SetType(v string)`

SetType sets Type field to given value.


### GetField

`func (o *AggregateSum) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AggregateSum) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AggregateSum) SetField(v string)`

SetField sets Field field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


