# GrokInvalidFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator identifying a rule whose filter is a known filter used incorrectly. | 
**Error** | [**GrokInvalidFilterError**](GrokInvalidFilterError.md) |  | 

## Methods

### NewGrokInvalidFilter

`func NewGrokInvalidFilter(type_ string, error_ GrokInvalidFilterError, ) *GrokInvalidFilter`

NewGrokInvalidFilter instantiates a new GrokInvalidFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrokInvalidFilterWithDefaults

`func NewGrokInvalidFilterWithDefaults() *GrokInvalidFilter`

NewGrokInvalidFilterWithDefaults instantiates a new GrokInvalidFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GrokInvalidFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GrokInvalidFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GrokInvalidFilter) SetType(v string)`

SetType sets Type field to given value.


### GetError

`func (o *GrokInvalidFilter) GetError() GrokInvalidFilterError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GrokInvalidFilter) GetErrorOk() (*GrokInvalidFilterError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GrokInvalidFilter) SetError(v GrokInvalidFilterError)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


