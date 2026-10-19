# DeletePublicTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**DeleteResponse**](DeleteResponse.md) |  | 

## Methods

### NewDeletePublicTokenResponse

`func NewDeletePublicTokenResponse(requestId string, data DeleteResponse, ) *DeletePublicTokenResponse`

NewDeletePublicTokenResponse instantiates a new DeletePublicTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletePublicTokenResponseWithDefaults

`func NewDeletePublicTokenResponseWithDefaults() *DeletePublicTokenResponse`

NewDeletePublicTokenResponseWithDefaults instantiates a new DeletePublicTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *DeletePublicTokenResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DeletePublicTokenResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DeletePublicTokenResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *DeletePublicTokenResponse) GetData() DeleteResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeletePublicTokenResponse) GetDataOk() (*DeleteResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeletePublicTokenResponse) SetData(v DeleteResponse)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


