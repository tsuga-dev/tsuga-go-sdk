# CreateCustomUsageTagResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**CustomUsageTag**](CustomUsageTag.md) |  | 

## Methods

### NewCreateCustomUsageTagResponse

`func NewCreateCustomUsageTagResponse(requestId string, data CustomUsageTag, ) *CreateCustomUsageTagResponse`

NewCreateCustomUsageTagResponse instantiates a new CreateCustomUsageTagResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomUsageTagResponseWithDefaults

`func NewCreateCustomUsageTagResponseWithDefaults() *CreateCustomUsageTagResponse`

NewCreateCustomUsageTagResponseWithDefaults instantiates a new CreateCustomUsageTagResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateCustomUsageTagResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateCustomUsageTagResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateCustomUsageTagResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *CreateCustomUsageTagResponse) GetData() CustomUsageTag`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateCustomUsageTagResponse) GetDataOk() (*CustomUsageTag, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateCustomUsageTagResponse) SetData(v CustomUsageTag)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


