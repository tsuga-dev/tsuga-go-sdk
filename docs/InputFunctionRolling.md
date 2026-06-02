# InputFunctionRolling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Computes a rolling window aggregate | 
**Window** | **string** | Duration of the rolling window | 

## Methods

### NewInputFunctionRolling

`func NewInputFunctionRolling(type_ string, window string, ) *InputFunctionRolling`

NewInputFunctionRolling instantiates a new InputFunctionRolling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputFunctionRollingWithDefaults

`func NewInputFunctionRollingWithDefaults() *InputFunctionRolling`

NewInputFunctionRollingWithDefaults instantiates a new InputFunctionRolling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputFunctionRolling) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputFunctionRolling) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputFunctionRolling) SetType(v string)`

SetType sets Type field to given value.


### GetWindow

`func (o *InputFunctionRolling) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *InputFunctionRolling) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *InputFunctionRolling) SetWindow(v string)`

SetWindow sets Window field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


