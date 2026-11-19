# CreateRumRouteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**RumRoute**](RumRoute.md) |  | 

## Methods

### NewCreateRumRouteResponse

`func NewCreateRumRouteResponse(requestId string, data RumRoute, ) *CreateRumRouteResponse`

NewCreateRumRouteResponse instantiates a new CreateRumRouteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRumRouteResponseWithDefaults

`func NewCreateRumRouteResponseWithDefaults() *CreateRumRouteResponse`

NewCreateRumRouteResponseWithDefaults instantiates a new CreateRumRouteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateRumRouteResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateRumRouteResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateRumRouteResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *CreateRumRouteResponse) GetData() RumRoute`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateRumRouteResponse) GetDataOk() (*RumRoute, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateRumRouteResponse) SetData(v RumRoute)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


