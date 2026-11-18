# UpdateSpanRouteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**SpanRoute**](SpanRoute.md) |  | 

## Methods

### NewUpdateSpanRouteResponse

`func NewUpdateSpanRouteResponse(requestId string, data SpanRoute, ) *UpdateSpanRouteResponse`

NewUpdateSpanRouteResponse instantiates a new UpdateSpanRouteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSpanRouteResponseWithDefaults

`func NewUpdateSpanRouteResponseWithDefaults() *UpdateSpanRouteResponse`

NewUpdateSpanRouteResponseWithDefaults instantiates a new UpdateSpanRouteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *UpdateSpanRouteResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *UpdateSpanRouteResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *UpdateSpanRouteResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *UpdateSpanRouteResponse) GetData() SpanRoute`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateSpanRouteResponse) GetDataOk() (*SpanRoute, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateSpanRouteResponse) SetData(v SpanRoute)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


