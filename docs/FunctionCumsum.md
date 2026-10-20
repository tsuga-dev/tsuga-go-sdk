# FunctionCumsum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Cumulative sum of the values over the time frame. Computed by the Tsuga web app only: the aggregation API ignores it and returns the raw values. | 

## Methods

### NewFunctionCumsum

`func NewFunctionCumsum(type_ string, ) *FunctionCumsum`

NewFunctionCumsum instantiates a new FunctionCumsum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionCumsumWithDefaults

`func NewFunctionCumsumWithDefaults() *FunctionCumsum`

NewFunctionCumsumWithDefaults instantiates a new FunctionCumsum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FunctionCumsum) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionCumsum) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionCumsum) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


