# QueryInvitationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**[]UserInvitation**](UserInvitation.md) |  | 
**Metadata** | Pointer to [**ResponseMetadata**](ResponseMetadata.md) |  | [optional] 

## Methods

### NewQueryInvitationsResponse

`func NewQueryInvitationsResponse(requestId string, data []UserInvitation, ) *QueryInvitationsResponse`

NewQueryInvitationsResponse instantiates a new QueryInvitationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryInvitationsResponseWithDefaults

`func NewQueryInvitationsResponseWithDefaults() *QueryInvitationsResponse`

NewQueryInvitationsResponseWithDefaults instantiates a new QueryInvitationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *QueryInvitationsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *QueryInvitationsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *QueryInvitationsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *QueryInvitationsResponse) GetData() []UserInvitation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QueryInvitationsResponse) GetDataOk() (*[]UserInvitation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QueryInvitationsResponse) SetData(v []UserInvitation)`

SetData sets Data field to given value.


### GetMetadata

`func (o *QueryInvitationsResponse) GetMetadata() ResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *QueryInvitationsResponse) GetMetadataOk() (*ResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *QueryInvitationsResponse) SetMetadata(v ResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *QueryInvitationsResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


