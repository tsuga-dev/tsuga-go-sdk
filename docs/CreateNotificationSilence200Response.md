# CreateNotificationSilence200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**NotificationSilence**](NotificationSilence.md) |  | 

## Methods

### NewCreateNotificationSilence200Response

`func NewCreateNotificationSilence200Response(requestId string, data NotificationSilence, ) *CreateNotificationSilence200Response`

NewCreateNotificationSilence200Response instantiates a new CreateNotificationSilence200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationSilence200ResponseWithDefaults

`func NewCreateNotificationSilence200ResponseWithDefaults() *CreateNotificationSilence200Response`

NewCreateNotificationSilence200ResponseWithDefaults instantiates a new CreateNotificationSilence200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateNotificationSilence200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateNotificationSilence200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateNotificationSilence200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *CreateNotificationSilence200Response) GetData() NotificationSilence`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateNotificationSilence200Response) GetDataOk() (*NotificationSilence, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateNotificationSilence200Response) SetData(v NotificationSilence)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


