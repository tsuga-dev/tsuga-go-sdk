# AggregateCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Counts the total number of records | 
**Field** | Pointer to **string** | Attribute containing the values to aggregate | [optional] 

## Methods

### NewAggregateCount

`func NewAggregateCount(type_ string, ) *AggregateCount`

NewAggregateCount instantiates a new AggregateCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateCountWithDefaults

`func NewAggregateCountWithDefaults() *AggregateCount`

NewAggregateCountWithDefaults instantiates a new AggregateCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AggregateCount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AggregateCount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AggregateCount) SetType(v string)`

SetType sets Type field to given value.


### GetField

`func (o *AggregateCount) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AggregateCount) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AggregateCount) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *AggregateCount) HasField() bool`

HasField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


