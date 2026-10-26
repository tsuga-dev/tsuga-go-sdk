# ClientErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Machine readable error code describing what went wrong | [optional] 
**Message** | **string** | Human readable explanation of the error | 
**StatusCode** | **float32** | HTTP status code that was returned | 
**Details** | Pointer to [**ClientErrorResponseDetails**](ClientErrorResponseDetails.md) |  | [optional] 
**Hint** | Pointer to **string** | One next step that fixes the error, such as the field to change or how long to wait. Omitted when the code covers too many cases for one step, and the message then says what to change. | [optional] 
**DocUrl** | Pointer to **string** | Absolute URL of the public documentation section that explains this error and how to fix it. | [optional] 

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


### GetDetails

`func (o *ClientErrorResponse) GetDetails() ClientErrorResponseDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ClientErrorResponse) GetDetailsOk() (*ClientErrorResponseDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ClientErrorResponse) SetDetails(v ClientErrorResponseDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ClientErrorResponse) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetHint

`func (o *ClientErrorResponse) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *ClientErrorResponse) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *ClientErrorResponse) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *ClientErrorResponse) HasHint() bool`

HasHint returns a boolean if a field has been set.

### GetDocUrl

`func (o *ClientErrorResponse) GetDocUrl() string`

GetDocUrl returns the DocUrl field if non-nil, zero value otherwise.

### GetDocUrlOk

`func (o *ClientErrorResponse) GetDocUrlOk() (*string, bool)`

GetDocUrlOk returns a tuple with the DocUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocUrl

`func (o *ClientErrorResponse) SetDocUrl(v string)`

SetDocUrl sets DocUrl field to given value.

### HasDocUrl

`func (o *ClientErrorResponse) HasDocUrl() bool`

HasDocUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


