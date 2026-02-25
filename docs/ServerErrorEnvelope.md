# ServerErrorEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Error** | [**ServerErrorResponse**](ServerErrorResponse.md) |  | 

## Methods

### NewServerErrorEnvelope

`func NewServerErrorEnvelope(requestId string, error_ ServerErrorResponse, ) *ServerErrorEnvelope`

NewServerErrorEnvelope instantiates a new ServerErrorEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerErrorEnvelopeWithDefaults

`func NewServerErrorEnvelopeWithDefaults() *ServerErrorEnvelope`

NewServerErrorEnvelopeWithDefaults instantiates a new ServerErrorEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ServerErrorEnvelope) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ServerErrorEnvelope) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ServerErrorEnvelope) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetError

`func (o *ServerErrorEnvelope) GetError() ServerErrorResponse`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ServerErrorEnvelope) GetErrorOk() (*ServerErrorResponse, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ServerErrorEnvelope) SetError(v ServerErrorResponse)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


