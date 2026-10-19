# GetLogAttributesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**GetLogAttributesResponseData**](GetLogAttributesResponseData.md) |  | 

## Methods

### NewGetLogAttributesResponse

`func NewGetLogAttributesResponse(requestId string, data GetLogAttributesResponseData, ) *GetLogAttributesResponse`

NewGetLogAttributesResponse instantiates a new GetLogAttributesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLogAttributesResponseWithDefaults

`func NewGetLogAttributesResponseWithDefaults() *GetLogAttributesResponse`

NewGetLogAttributesResponseWithDefaults instantiates a new GetLogAttributesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetLogAttributesResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetLogAttributesResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetLogAttributesResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *GetLogAttributesResponse) GetData() GetLogAttributesResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetLogAttributesResponse) GetDataOk() (*GetLogAttributesResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetLogAttributesResponse) SetData(v GetLogAttributesResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


