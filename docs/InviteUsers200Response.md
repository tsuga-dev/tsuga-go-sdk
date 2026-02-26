# InviteUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**InviteUsers200ResponseData**](InviteUsers200ResponseData.md) |  | 

## Methods

### NewInviteUsers200Response

`func NewInviteUsers200Response(requestId string, data InviteUsers200ResponseData, ) *InviteUsers200Response`

NewInviteUsers200Response instantiates a new InviteUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteUsers200ResponseWithDefaults

`func NewInviteUsers200ResponseWithDefaults() *InviteUsers200Response`

NewInviteUsers200ResponseWithDefaults instantiates a new InviteUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *InviteUsers200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *InviteUsers200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *InviteUsers200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *InviteUsers200Response) GetData() InviteUsers200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InviteUsers200Response) GetDataOk() (*InviteUsers200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InviteUsers200Response) SetData(v InviteUsers200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


