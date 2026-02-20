# ServerErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Machine readable error code describing what went wrong | [optional] 
**Message** | **string** | Human readable explanation of the error | 
**StatusCode** | **float32** | HTTP status code that was returned | 

## Methods

### NewServerErrorResponse

`func NewServerErrorResponse(message string, statusCode float32, ) *ServerErrorResponse`

NewServerErrorResponse instantiates a new ServerErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerErrorResponseWithDefaults

`func NewServerErrorResponseWithDefaults() *ServerErrorResponse`

NewServerErrorResponseWithDefaults instantiates a new ServerErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ServerErrorResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ServerErrorResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ServerErrorResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ServerErrorResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ServerErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServerErrorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServerErrorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStatusCode

`func (o *ServerErrorResponse) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ServerErrorResponse) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ServerErrorResponse) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


