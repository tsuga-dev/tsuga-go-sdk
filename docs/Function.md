# Function

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Normalizes the metric to events per second | 
**Window** | **string** | Duration of the rolling window | 
**Base** | **float32** | The base of the logarithm | 
**Exponent** | **float32** | The exponent to raise values to | 

## Methods

### NewFunction

`func NewFunction(type_ string, window string, base float32, exponent float32, ) *Function`

NewFunction instantiates a new Function object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionWithDefaults

`func NewFunctionWithDefaults() *Function`

NewFunctionWithDefaults instantiates a new Function object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Function) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Function) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Function) SetType(v string)`

SetType sets Type field to given value.


### GetWindow

`func (o *Function) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *Function) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *Function) SetWindow(v string)`

SetWindow sets Window field to given value.


### GetBase

`func (o *Function) GetBase() float32`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *Function) GetBaseOk() (*float32, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *Function) SetBase(v float32)`

SetBase sets Base field to given value.


### GetExponent

`func (o *Function) GetExponent() float32`

GetExponent returns the Exponent field if non-nil, zero value otherwise.

### GetExponentOk

`func (o *Function) GetExponentOk() (*float32, bool)`

GetExponentOk returns a tuple with the Exponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExponent

`func (o *Function) SetExponent(v float32)`

SetExponent sets Exponent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


