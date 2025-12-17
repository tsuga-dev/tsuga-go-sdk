# CreateMonitor200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**Monitor**](Monitor.md) |  | 

## Methods

### NewCreateMonitor200Response

`func NewCreateMonitor200Response(requestId string, data Monitor, ) *CreateMonitor200Response`

NewCreateMonitor200Response instantiates a new CreateMonitor200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMonitor200ResponseWithDefaults

`func NewCreateMonitor200ResponseWithDefaults() *CreateMonitor200Response`

NewCreateMonitor200ResponseWithDefaults instantiates a new CreateMonitor200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateMonitor200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateMonitor200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateMonitor200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *CreateMonitor200Response) GetData() Monitor`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateMonitor200Response) GetDataOk() (*Monitor, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateMonitor200Response) SetData(v Monitor)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


