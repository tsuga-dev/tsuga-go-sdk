# ClientErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Machine readable error code describing what went wrong | [optional] 
**Message** | **string** | Human readable explanation of the error | 
**StatusCode** | **float32** | HTTP status code that was returned | 

## Methods

### NewClientErrorResponse

`func NewClientErrorResponse(message string, statusCode float32, ) *ClientErrorResponse`

NewClientErrorResponse instantiates a new ClientErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientErrorResponseWithDefaults

`func NewClientErrorResponseWithDefaults() *ClientErrorResponse`

NewClientErrorResponseWithDefaults instantiates a new ClientErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ClientErrorResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ClientErrorResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ClientErrorResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ClientErrorResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ClientErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClientErrorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClientErrorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStatusCode

`func (o *ClientErrorResponse) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ClientErrorResponse) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ClientErrorResponse) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


