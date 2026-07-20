# GraphVisualizationTimeseriesPromqlTimeBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **float32** | Numeric count of &#x60;metric&#x60; units in each bucket. Required when &#x60;timeBucket&#x60; is present. | 
**Metric** | **string** | Unit used to measure the bucket width | 

## Methods

### NewGraphVisualizationTimeseriesPromqlTimeBucket

`func NewGraphVisualizationTimeseriesPromqlTimeBucket(time float32, metric string, ) *GraphVisualizationTimeseriesPromqlTimeBucket`

NewGraphVisualizationTimeseriesPromqlTimeBucket instantiates a new GraphVisualizationTimeseriesPromqlTimeBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationTimeseriesPromqlTimeBucketWithDefaults

`func NewGraphVisualizationTimeseriesPromqlTimeBucketWithDefaults() *GraphVisualizationTimeseriesPromqlTimeBucket`

NewGraphVisualizationTimeseriesPromqlTimeBucketWithDefaults instantiates a new GraphVisualizationTimeseriesPromqlTimeBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *GraphVisualizationTimeseriesPromqlTimeBucket) GetTime() float32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GraphVisualizationTimeseriesPromqlTimeBucket) GetTimeOk() (*float32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GraphVisualizationTimeseriesPromqlTimeBucket) SetTime(v float32)`

SetTime sets Time field to given value.


### GetMetric

`func (o *GraphVisualizationTimeseriesPromqlTimeBucket) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *GraphVisualizationTimeseriesPromqlTimeBucket) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *GraphVisualizationTimeseriesPromqlTimeBucket) SetMetric(v string)`

SetMetric sets Metric field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


