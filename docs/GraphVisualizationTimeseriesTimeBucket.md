# GraphVisualizationTimeseriesTimeBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **float32** | Number of units that make up each bar bucket | 
**Metric** | **string** | Unit used to measure the bucket width | 

## Methods

### NewGraphVisualizationTimeseriesTimeBucket

`func NewGraphVisualizationTimeseriesTimeBucket(time float32, metric string, ) *GraphVisualizationTimeseriesTimeBucket`

NewGraphVisualizationTimeseriesTimeBucket instantiates a new GraphVisualizationTimeseriesTimeBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationTimeseriesTimeBucketWithDefaults

`func NewGraphVisualizationTimeseriesTimeBucketWithDefaults() *GraphVisualizationTimeseriesTimeBucket`

NewGraphVisualizationTimeseriesTimeBucketWithDefaults instantiates a new GraphVisualizationTimeseriesTimeBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *GraphVisualizationTimeseriesTimeBucket) GetTime() float32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GraphVisualizationTimeseriesTimeBucket) GetTimeOk() (*float32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GraphVisualizationTimeseriesTimeBucket) SetTime(v float32)`

SetTime sets Time field to given value.


### GetMetric

`func (o *GraphVisualizationTimeseriesTimeBucket) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *GraphVisualizationTimeseriesTimeBucket) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *GraphVisualizationTimeseriesTimeBucket) SetMetric(v string)`

SetMetric sets Metric field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


