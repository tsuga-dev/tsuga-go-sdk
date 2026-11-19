# ListRumRoutesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**[]RumRoute**](RumRoute.md) |  | 
**Metadata** | Pointer to [**ResponseMetadata**](ResponseMetadata.md) |  | [optional] 

## Methods

### NewListRumRoutesResponse

`func NewListRumRoutesResponse(requestId string, data []RumRoute, ) *ListRumRoutesResponse`

NewListRumRoutesResponse instantiates a new ListRumRoutesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRumRoutesResponseWithDefaults

`func NewListRumRoutesResponseWithDefaults() *ListRumRoutesResponse`

NewListRumRoutesResponseWithDefaults instantiates a new ListRumRoutesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ListRumRoutesResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListRumRoutesResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListRumRoutesResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *ListRumRoutesResponse) GetData() []RumRoute`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListRumRoutesResponse) GetDataOk() (*[]RumRoute, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListRumRoutesResponse) SetData(v []RumRoute)`

SetData sets Data field to given value.


### GetMetadata

`func (o *ListRumRoutesResponse) GetMetadata() ResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListRumRoutesResponse) GetMetadataOk() (*ResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListRumRoutesResponse) SetMetadata(v ResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ListRumRoutesResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


