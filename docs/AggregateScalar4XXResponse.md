# AggregateScalar4XXResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Error** | [**ClientErrorResponse**](ClientErrorResponse.md) |  | 

## Methods

### NewAggregateScalar4XXResponse

`func NewAggregateScalar4XXResponse(requestId string, error_ ClientErrorResponse, ) *AggregateScalar4XXResponse`

NewAggregateScalar4XXResponse instantiates a new AggregateScalar4XXResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateScalar4XXResponseWithDefaults

`func NewAggregateScalar4XXResponseWithDefaults() *AggregateScalar4XXResponse`

NewAggregateScalar4XXResponseWithDefaults instantiates a new AggregateScalar4XXResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *AggregateScalar4XXResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AggregateScalar4XXResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AggregateScalar4XXResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetError

`func (o *AggregateScalar4XXResponse) GetError() ClientErrorResponse`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AggregateScalar4XXResponse) GetErrorOk() (*ClientErrorResponse, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AggregateScalar4XXResponse) SetError(v ClientErrorResponse)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


