# SearchSpans200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**SearchSpans200ResponseData**](SearchSpans200ResponseData.md) |  | 

## Methods

### NewSearchSpans200Response

`func NewSearchSpans200Response(requestId string, data SearchSpans200ResponseData, ) *SearchSpans200Response`

NewSearchSpans200Response instantiates a new SearchSpans200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSpans200ResponseWithDefaults

`func NewSearchSpans200ResponseWithDefaults() *SearchSpans200Response`

NewSearchSpans200ResponseWithDefaults instantiates a new SearchSpans200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *SearchSpans200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SearchSpans200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SearchSpans200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *SearchSpans200Response) GetData() SearchSpans200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SearchSpans200Response) GetDataOk() (*SearchSpans200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SearchSpans200Response) SetData(v SearchSpans200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


