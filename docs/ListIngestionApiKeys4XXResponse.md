# ListIngestionApiKeys4XXResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** |  | 
**Error** | [**ErrorResponse**](ErrorResponse.md) |  | 

## Methods

### NewListIngestionApiKeys4XXResponse

`func NewListIngestionApiKeys4XXResponse(requestId string, error_ ErrorResponse, ) *ListIngestionApiKeys4XXResponse`

NewListIngestionApiKeys4XXResponse instantiates a new ListIngestionApiKeys4XXResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIngestionApiKeys4XXResponseWithDefaults

`func NewListIngestionApiKeys4XXResponseWithDefaults() *ListIngestionApiKeys4XXResponse`

NewListIngestionApiKeys4XXResponseWithDefaults instantiates a new ListIngestionApiKeys4XXResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ListIngestionApiKeys4XXResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListIngestionApiKeys4XXResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListIngestionApiKeys4XXResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetError

`func (o *ListIngestionApiKeys4XXResponse) GetError() ErrorResponse`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListIngestionApiKeys4XXResponse) GetErrorOk() (*ErrorResponse, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListIngestionApiKeys4XXResponse) SetError(v ErrorResponse)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


