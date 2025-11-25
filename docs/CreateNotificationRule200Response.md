# CreateNotificationRule200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**Rule**](Rule.md) |  | 

## Methods

### NewCreateNotificationRule200Response

`func NewCreateNotificationRule200Response(requestId string, data Rule, ) *CreateNotificationRule200Response`

NewCreateNotificationRule200Response instantiates a new CreateNotificationRule200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationRule200ResponseWithDefaults

`func NewCreateNotificationRule200ResponseWithDefaults() *CreateNotificationRule200Response`

NewCreateNotificationRule200ResponseWithDefaults instantiates a new CreateNotificationRule200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateNotificationRule200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateNotificationRule200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateNotificationRule200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *CreateNotificationRule200Response) GetData() Rule`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateNotificationRule200Response) GetDataOk() (*Rule, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateNotificationRule200Response) SetData(v Rule)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


