# QueryPromqlV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**TimeseriesAggregationResponseV2**](TimeseriesAggregationResponseV2.md) |  | 

## Methods

### NewQueryPromqlV2Response

`func NewQueryPromqlV2Response(requestId string, data TimeseriesAggregationResponseV2, ) *QueryPromqlV2Response`

NewQueryPromqlV2Response instantiates a new QueryPromqlV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryPromqlV2ResponseWithDefaults

`func NewQueryPromqlV2ResponseWithDefaults() *QueryPromqlV2Response`

NewQueryPromqlV2ResponseWithDefaults instantiates a new QueryPromqlV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *QueryPromqlV2Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *QueryPromqlV2Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *QueryPromqlV2Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *QueryPromqlV2Response) GetData() TimeseriesAggregationResponseV2`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QueryPromqlV2Response) GetDataOk() (*TimeseriesAggregationResponseV2, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QueryPromqlV2Response) SetData(v TimeseriesAggregationResponseV2)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


