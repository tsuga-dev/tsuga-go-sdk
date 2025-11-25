# FunctionRolling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Computes a rolling window aggregate | 
**Window** | **string** | Duration of the rolling window | 

## Methods

### NewFunctionRolling

`func NewFunctionRolling(type_ string, window string, ) *FunctionRolling`

NewFunctionRolling instantiates a new FunctionRolling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionRollingWithDefaults

`func NewFunctionRollingWithDefaults() *FunctionRolling`

NewFunctionRollingWithDefaults instantiates a new FunctionRolling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FunctionRolling) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionRolling) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionRolling) SetType(v string)`

SetType sets Type field to given value.


### GetWindow

`func (o *FunctionRolling) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *FunctionRolling) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *FunctionRolling) SetWindow(v string)`

SetWindow sets Window field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


