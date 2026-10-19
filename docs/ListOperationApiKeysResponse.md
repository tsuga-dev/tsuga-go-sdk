# ListOperationApiKeysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**[]OperationApiKey**](OperationApiKey.md) |  | 
**Metadata** | Pointer to [**ResponseMetadata**](ResponseMetadata.md) |  | [optional] 

## Methods

### NewListOperationApiKeysResponse

`func NewListOperationApiKeysResponse(requestId string, data []OperationApiKey, ) *ListOperationApiKeysResponse`

NewListOperationApiKeysResponse instantiates a new ListOperationApiKeysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOperationApiKeysResponseWithDefaults

`func NewListOperationApiKeysResponseWithDefaults() *ListOperationApiKeysResponse`

NewListOperationApiKeysResponseWithDefaults instantiates a new ListOperationApiKeysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ListOperationApiKeysResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListOperationApiKeysResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListOperationApiKeysResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *ListOperationApiKeysResponse) GetData() []OperationApiKey`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListOperationApiKeysResponse) GetDataOk() (*[]OperationApiKey, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListOperationApiKeysResponse) SetData(v []OperationApiKey)`

SetData sets Data field to given value.


### GetMetadata

`func (o *ListOperationApiKeysResponse) GetMetadata() ResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListOperationApiKeysResponse) GetMetadataOk() (*ResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListOperationApiKeysResponse) SetMetadata(v ResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ListOperationApiKeysResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


