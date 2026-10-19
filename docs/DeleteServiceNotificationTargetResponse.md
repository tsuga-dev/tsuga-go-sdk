# DeleteServiceNotificationTargetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**DeleteResponse**](DeleteResponse.md) |  | 

## Methods

### NewDeleteServiceNotificationTargetResponse

`func NewDeleteServiceNotificationTargetResponse(requestId string, data DeleteResponse, ) *DeleteServiceNotificationTargetResponse`

NewDeleteServiceNotificationTargetResponse instantiates a new DeleteServiceNotificationTargetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteServiceNotificationTargetResponseWithDefaults

`func NewDeleteServiceNotificationTargetResponseWithDefaults() *DeleteServiceNotificationTargetResponse`

NewDeleteServiceNotificationTargetResponseWithDefaults instantiates a new DeleteServiceNotificationTargetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *DeleteServiceNotificationTargetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DeleteServiceNotificationTargetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DeleteServiceNotificationTargetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *DeleteServiceNotificationTargetResponse) GetData() DeleteResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteServiceNotificationTargetResponse) GetDataOk() (*DeleteResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteServiceNotificationTargetResponse) SetData(v DeleteResponse)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


