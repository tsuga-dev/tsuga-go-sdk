# SearchLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**SearchLogs200ResponseData**](SearchLogs200ResponseData.md) |  | 

## Methods

### NewSearchLogs200Response

`func NewSearchLogs200Response(requestId string, data SearchLogs200ResponseData, ) *SearchLogs200Response`

NewSearchLogs200Response instantiates a new SearchLogs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchLogs200ResponseWithDefaults

`func NewSearchLogs200ResponseWithDefaults() *SearchLogs200Response`

NewSearchLogs200ResponseWithDefaults instantiates a new SearchLogs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *SearchLogs200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SearchLogs200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SearchLogs200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *SearchLogs200Response) GetData() SearchLogs200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SearchLogs200Response) GetDataOk() (*SearchLogs200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SearchLogs200Response) SetData(v SearchLogs200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


