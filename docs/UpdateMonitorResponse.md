# UpdateMonitorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**Monitor**](Monitor.md) |  | 

## Methods

### NewUpdateMonitorResponse

`func NewUpdateMonitorResponse(requestId string, data Monitor, ) *UpdateMonitorResponse`

NewUpdateMonitorResponse instantiates a new UpdateMonitorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMonitorResponseWithDefaults

`func NewUpdateMonitorResponseWithDefaults() *UpdateMonitorResponse`

NewUpdateMonitorResponseWithDefaults instantiates a new UpdateMonitorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *UpdateMonitorResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *UpdateMonitorResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *UpdateMonitorResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *UpdateMonitorResponse) GetData() Monitor`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateMonitorResponse) GetDataOk() (*Monitor, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateMonitorResponse) SetData(v Monitor)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


