# ListKubernetesDaemonSetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**[]ListKubernetesDaemonSetsResponseDataInner**](ListKubernetesDaemonSetsResponseDataInner.md) |  | 
**Metadata** | Pointer to [**ResponseMetadata**](ResponseMetadata.md) |  | [optional] 

## Methods

### NewListKubernetesDaemonSetsResponse

`func NewListKubernetesDaemonSetsResponse(requestId string, data []ListKubernetesDaemonSetsResponseDataInner, ) *ListKubernetesDaemonSetsResponse`

NewListKubernetesDaemonSetsResponse instantiates a new ListKubernetesDaemonSetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKubernetesDaemonSetsResponseWithDefaults

`func NewListKubernetesDaemonSetsResponseWithDefaults() *ListKubernetesDaemonSetsResponse`

NewListKubernetesDaemonSetsResponseWithDefaults instantiates a new ListKubernetesDaemonSetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ListKubernetesDaemonSetsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListKubernetesDaemonSetsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListKubernetesDaemonSetsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *ListKubernetesDaemonSetsResponse) GetData() []ListKubernetesDaemonSetsResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListKubernetesDaemonSetsResponse) GetDataOk() (*[]ListKubernetesDaemonSetsResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListKubernetesDaemonSetsResponse) SetData(v []ListKubernetesDaemonSetsResponseDataInner)`

SetData sets Data field to given value.


### GetMetadata

`func (o *ListKubernetesDaemonSetsResponse) GetMetadata() ResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListKubernetesDaemonSetsResponse) GetMetadataOk() (*ResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListKubernetesDaemonSetsResponse) SetMetadata(v ResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ListKubernetesDaemonSetsResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


