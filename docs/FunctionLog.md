# FunctionLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Applies a logarithm transformation to the values | 
**Base** | **int32** | The base of the logarithm | 

## Methods

### NewFunctionLog

`func NewFunctionLog(type_ string, base int32, ) *FunctionLog`

NewFunctionLog instantiates a new FunctionLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionLogWithDefaults

`func NewFunctionLogWithDefaults() *FunctionLog`

NewFunctionLogWithDefaults instantiates a new FunctionLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FunctionLog) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionLog) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionLog) SetType(v string)`

SetType sets Type field to given value.


### GetBase

`func (o *FunctionLog) GetBase() int32`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *FunctionLog) GetBaseOk() (*int32, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *FunctionLog) SetBase(v int32)`

SetBase sets Base field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


