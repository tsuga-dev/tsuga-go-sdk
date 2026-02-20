# GetMetric200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**ListMetrics200ResponseDataInner**](ListMetrics200ResponseDataInner.md) |  | 

## Methods

### NewGetMetric200Response

`func NewGetMetric200Response(requestId string, data ListMetrics200ResponseDataInner, ) *GetMetric200Response`

NewGetMetric200Response instantiates a new GetMetric200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMetric200ResponseWithDefaults

`func NewGetMetric200ResponseWithDefaults() *GetMetric200Response`

NewGetMetric200ResponseWithDefaults instantiates a new GetMetric200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetMetric200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetMetric200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetMetric200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *GetMetric200Response) GetData() ListMetrics200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMetric200Response) GetDataOk() (*ListMetrics200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMetric200Response) SetData(v ListMetrics200ResponseDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


