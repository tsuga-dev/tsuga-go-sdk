# GetCustomUsageTagResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**CustomUsageTag**](CustomUsageTag.md) |  | 

## Methods

### NewGetCustomUsageTagResponse

`func NewGetCustomUsageTagResponse(requestId string, data CustomUsageTag, ) *GetCustomUsageTagResponse`

NewGetCustomUsageTagResponse instantiates a new GetCustomUsageTagResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomUsageTagResponseWithDefaults

`func NewGetCustomUsageTagResponseWithDefaults() *GetCustomUsageTagResponse`

NewGetCustomUsageTagResponseWithDefaults instantiates a new GetCustomUsageTagResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetCustomUsageTagResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetCustomUsageTagResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetCustomUsageTagResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *GetCustomUsageTagResponse) GetData() CustomUsageTag`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCustomUsageTagResponse) GetDataOk() (*CustomUsageTag, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCustomUsageTagResponse) SetData(v CustomUsageTag)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


