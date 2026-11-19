# UpdateRumRouteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**RumRoute**](RumRoute.md) |  | 

## Methods

### NewUpdateRumRouteResponse

`func NewUpdateRumRouteResponse(requestId string, data RumRoute, ) *UpdateRumRouteResponse`

NewUpdateRumRouteResponse instantiates a new UpdateRumRouteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRumRouteResponseWithDefaults

`func NewUpdateRumRouteResponseWithDefaults() *UpdateRumRouteResponse`

NewUpdateRumRouteResponseWithDefaults instantiates a new UpdateRumRouteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *UpdateRumRouteResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *UpdateRumRouteResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *UpdateRumRouteResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *UpdateRumRouteResponse) GetData() RumRoute`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateRumRouteResponse) GetDataOk() (*RumRoute, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateRumRouteResponse) SetData(v RumRoute)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


