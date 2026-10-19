# ListPublicTokensResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**[]PublicToken**](PublicToken.md) |  | 
**Metadata** | Pointer to [**ResponseMetadata**](ResponseMetadata.md) |  | [optional] 

## Methods

### NewListPublicTokensResponse

`func NewListPublicTokensResponse(requestId string, data []PublicToken, ) *ListPublicTokensResponse`

NewListPublicTokensResponse instantiates a new ListPublicTokensResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPublicTokensResponseWithDefaults

`func NewListPublicTokensResponseWithDefaults() *ListPublicTokensResponse`

NewListPublicTokensResponseWithDefaults instantiates a new ListPublicTokensResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ListPublicTokensResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListPublicTokensResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListPublicTokensResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *ListPublicTokensResponse) GetData() []PublicToken`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPublicTokensResponse) GetDataOk() (*[]PublicToken, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPublicTokensResponse) SetData(v []PublicToken)`

SetData sets Data field to given value.


### GetMetadata

`func (o *ListPublicTokensResponse) GetMetadata() ResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListPublicTokensResponse) GetMetadataOk() (*ResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListPublicTokensResponse) SetMetadata(v ResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ListPublicTokensResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


