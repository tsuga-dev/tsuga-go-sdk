# CreateMetricRouteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**MetricRoute**](MetricRoute.md) |  | 

## Methods

### NewCreateMetricRouteResponse

`func NewCreateMetricRouteResponse(requestId string, data MetricRoute, ) *CreateMetricRouteResponse`

NewCreateMetricRouteResponse instantiates a new CreateMetricRouteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMetricRouteResponseWithDefaults

`func NewCreateMetricRouteResponseWithDefaults() *CreateMetricRouteResponse`

NewCreateMetricRouteResponseWithDefaults instantiates a new CreateMetricRouteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateMetricRouteResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateMetricRouteResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateMetricRouteResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *CreateMetricRouteResponse) GetData() MetricRoute`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateMetricRouteResponse) GetDataOk() (*MetricRoute, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateMetricRouteResponse) SetData(v MetricRoute)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


