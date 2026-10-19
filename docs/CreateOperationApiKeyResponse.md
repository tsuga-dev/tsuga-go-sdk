# CreateOperationApiKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**OperationApiKeyCreated**](OperationApiKeyCreated.md) |  | 

## Methods

### NewCreateOperationApiKeyResponse

`func NewCreateOperationApiKeyResponse(requestId string, data OperationApiKeyCreated, ) *CreateOperationApiKeyResponse`

NewCreateOperationApiKeyResponse instantiates a new CreateOperationApiKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOperationApiKeyResponseWithDefaults

`func NewCreateOperationApiKeyResponseWithDefaults() *CreateOperationApiKeyResponse`

NewCreateOperationApiKeyResponseWithDefaults instantiates a new CreateOperationApiKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateOperationApiKeyResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateOperationApiKeyResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateOperationApiKeyResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *CreateOperationApiKeyResponse) GetData() OperationApiKeyCreated`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateOperationApiKeyResponse) GetDataOk() (*OperationApiKeyCreated, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateOperationApiKeyResponse) SetData(v OperationApiKeyCreated)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


