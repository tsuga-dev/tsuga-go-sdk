# UpdateIngestionApiKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**IngestionApiKey**](IngestionApiKey.md) |  | 

## Methods

### NewUpdateIngestionApiKeyResponse

`func NewUpdateIngestionApiKeyResponse(requestId string, data IngestionApiKey, ) *UpdateIngestionApiKeyResponse`

NewUpdateIngestionApiKeyResponse instantiates a new UpdateIngestionApiKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIngestionApiKeyResponseWithDefaults

`func NewUpdateIngestionApiKeyResponseWithDefaults() *UpdateIngestionApiKeyResponse`

NewUpdateIngestionApiKeyResponseWithDefaults instantiates a new UpdateIngestionApiKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *UpdateIngestionApiKeyResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *UpdateIngestionApiKeyResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *UpdateIngestionApiKeyResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *UpdateIngestionApiKeyResponse) GetData() IngestionApiKey`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateIngestionApiKeyResponse) GetDataOk() (*IngestionApiKey, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateIngestionApiKeyResponse) SetData(v IngestionApiKey)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


