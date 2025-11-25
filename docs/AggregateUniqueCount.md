# AggregateUniqueCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Counts the distinct values of the field | 
**Field** | **string** | Attribute containing the values to aggregate | 

## Methods

### NewAggregateUniqueCount

`func NewAggregateUniqueCount(type_ string, field string, ) *AggregateUniqueCount`

NewAggregateUniqueCount instantiates a new AggregateUniqueCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateUniqueCountWithDefaults

`func NewAggregateUniqueCountWithDefaults() *AggregateUniqueCount`

NewAggregateUniqueCountWithDefaults instantiates a new AggregateUniqueCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AggregateUniqueCount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AggregateUniqueCount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AggregateUniqueCount) SetType(v string)`

SetType sets Type field to given value.


### GetField

`func (o *AggregateUniqueCount) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AggregateUniqueCount) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AggregateUniqueCount) SetField(v string)`

SetField sets Field field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


