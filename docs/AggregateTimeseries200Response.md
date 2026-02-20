# AggregateTimeseries200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**TimeseriesAggregationResponse**](TimeseriesAggregationResponse.md) |  | 

## Methods

### NewAggregateTimeseries200Response

`func NewAggregateTimeseries200Response(requestId string, data TimeseriesAggregationResponse, ) *AggregateTimeseries200Response`

NewAggregateTimeseries200Response instantiates a new AggregateTimeseries200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateTimeseries200ResponseWithDefaults

`func NewAggregateTimeseries200ResponseWithDefaults() *AggregateTimeseries200Response`

NewAggregateTimeseries200ResponseWithDefaults instantiates a new AggregateTimeseries200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *AggregateTimeseries200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AggregateTimeseries200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AggregateTimeseries200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *AggregateTimeseries200Response) GetData() TimeseriesAggregationResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AggregateTimeseries200Response) GetDataOk() (*TimeseriesAggregationResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AggregateTimeseries200Response) SetData(v TimeseriesAggregationResponse)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


