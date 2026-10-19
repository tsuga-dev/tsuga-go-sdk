# QueryInventoryResourcesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**[]QueryInventoryResourcesResponseDataInner**](QueryInventoryResourcesResponseDataInner.md) |  | 
**Metadata** | Pointer to [**ResponseMetadata**](ResponseMetadata.md) |  | [optional] 

## Methods

### NewQueryInventoryResourcesResponse

`func NewQueryInventoryResourcesResponse(requestId string, data []QueryInventoryResourcesResponseDataInner, ) *QueryInventoryResourcesResponse`

NewQueryInventoryResourcesResponse instantiates a new QueryInventoryResourcesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryInventoryResourcesResponseWithDefaults

`func NewQueryInventoryResourcesResponseWithDefaults() *QueryInventoryResourcesResponse`

NewQueryInventoryResourcesResponseWithDefaults instantiates a new QueryInventoryResourcesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *QueryInventoryResourcesResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *QueryInventoryResourcesResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *QueryInventoryResourcesResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *QueryInventoryResourcesResponse) GetData() []QueryInventoryResourcesResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QueryInventoryResourcesResponse) GetDataOk() (*[]QueryInventoryResourcesResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QueryInventoryResourcesResponse) SetData(v []QueryInventoryResourcesResponseDataInner)`

SetData sets Data field to given value.


### GetMetadata

`func (o *QueryInventoryResourcesResponse) GetMetadata() ResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *QueryInventoryResourcesResponse) GetMetadataOk() (*ResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *QueryInventoryResourcesResponse) SetMetadata(v ResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *QueryInventoryResourcesResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


