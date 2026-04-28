# MonitorFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Normalizes the metric to events per second | 
**Window** | **string** | Duration of the rolling window | 
**Base** | **int32** | The base of the logarithm | 
**Exponent** | **int32** | The exponent to raise values to | 

## Methods

### NewMonitorFunction

`func NewMonitorFunction(type_ string, window string, base int32, exponent int32, ) *MonitorFunction`

NewMonitorFunction instantiates a new MonitorFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorFunctionWithDefaults

`func NewMonitorFunctionWithDefaults() *MonitorFunction`

NewMonitorFunctionWithDefaults instantiates a new MonitorFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MonitorFunction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorFunction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorFunction) SetType(v string)`

SetType sets Type field to given value.


### GetWindow

`func (o *MonitorFunction) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *MonitorFunction) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *MonitorFunction) SetWindow(v string)`

SetWindow sets Window field to given value.


### GetBase

`func (o *MonitorFunction) GetBase() int32`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *MonitorFunction) GetBaseOk() (*int32, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *MonitorFunction) SetBase(v int32)`

SetBase sets Base field to given value.


### GetExponent

`func (o *MonitorFunction) GetExponent() int32`

GetExponent returns the Exponent field if non-nil, zero value otherwise.

### GetExponentOk

`func (o *MonitorFunction) GetExponentOk() (*int32, bool)`

GetExponentOk returns a tuple with the Exponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExponent

`func (o *MonitorFunction) SetExponent(v int32)`

SetExponent sets Exponent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


