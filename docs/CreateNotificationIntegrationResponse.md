# CreateNotificationIntegrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**NotificationIntegration**](NotificationIntegration.md) |  | 

## Methods

### NewCreateNotificationIntegrationResponse

`func NewCreateNotificationIntegrationResponse(requestId string, data NotificationIntegration, ) *CreateNotificationIntegrationResponse`

NewCreateNotificationIntegrationResponse instantiates a new CreateNotificationIntegrationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationIntegrationResponseWithDefaults

`func NewCreateNotificationIntegrationResponseWithDefaults() *CreateNotificationIntegrationResponse`

NewCreateNotificationIntegrationResponseWithDefaults instantiates a new CreateNotificationIntegrationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateNotificationIntegrationResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateNotificationIntegrationResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateNotificationIntegrationResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *CreateNotificationIntegrationResponse) GetData() NotificationIntegration`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateNotificationIntegrationResponse) GetDataOk() (*NotificationIntegration, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateNotificationIntegrationResponse) SetData(v NotificationIntegration)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


