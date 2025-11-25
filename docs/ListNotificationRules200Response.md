# ListNotificationRules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**[]ListNotificationRules200ResponseDataInner**](ListNotificationRules200ResponseDataInner.md) |  | 

## Methods

### NewListNotificationRules200Response

`func NewListNotificationRules200Response(requestId string, data []ListNotificationRules200ResponseDataInner, ) *ListNotificationRules200Response`

NewListNotificationRules200Response instantiates a new ListNotificationRules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNotificationRules200ResponseWithDefaults

`func NewListNotificationRules200ResponseWithDefaults() *ListNotificationRules200Response`

NewListNotificationRules200ResponseWithDefaults instantiates a new ListNotificationRules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ListNotificationRules200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListNotificationRules200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListNotificationRules200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *ListNotificationRules200Response) GetData() []ListNotificationRules200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListNotificationRules200Response) GetDataOk() (*[]ListNotificationRules200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListNotificationRules200Response) SetData(v []ListNotificationRules200ResponseDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


