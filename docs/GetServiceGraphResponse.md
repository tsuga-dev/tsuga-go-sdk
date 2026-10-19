# GetServiceGraphResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**GetServiceGraphResponseData**](GetServiceGraphResponseData.md) |  | 

## Methods

### NewGetServiceGraphResponse

`func NewGetServiceGraphResponse(requestId string, data GetServiceGraphResponseData, ) *GetServiceGraphResponse`

NewGetServiceGraphResponse instantiates a new GetServiceGraphResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServiceGraphResponseWithDefaults

`func NewGetServiceGraphResponseWithDefaults() *GetServiceGraphResponse`

NewGetServiceGraphResponseWithDefaults instantiates a new GetServiceGraphResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetServiceGraphResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetServiceGraphResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetServiceGraphResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *GetServiceGraphResponse) GetData() GetServiceGraphResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetServiceGraphResponse) GetDataOk() (*GetServiceGraphResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetServiceGraphResponse) SetData(v GetServiceGraphResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


