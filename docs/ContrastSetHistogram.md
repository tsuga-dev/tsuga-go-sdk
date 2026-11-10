# ContrastSetHistogram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bounds** | **[]float32** | Bin edges, in the unit of the attribute. One more item than each count list. The bins go from the 2nd to the 98th percentile of both groups, so they do not hold all samples. | 
**TargetCounts** | **[]float32** | Number of target samples in each bin. | 
**BaselineCounts** | **[]float32** | Number of baseline samples in each bin. | 

## Methods

### NewContrastSetHistogram

`func NewContrastSetHistogram(bounds []float32, targetCounts []float32, baselineCounts []float32, ) *ContrastSetHistogram`

NewContrastSetHistogram instantiates a new ContrastSetHistogram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContrastSetHistogramWithDefaults

`func NewContrastSetHistogramWithDefaults() *ContrastSetHistogram`

NewContrastSetHistogramWithDefaults instantiates a new ContrastSetHistogram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBounds

`func (o *ContrastSetHistogram) GetBounds() []float32`

GetBounds returns the Bounds field if non-nil, zero value otherwise.

### GetBoundsOk

`func (o *ContrastSetHistogram) GetBoundsOk() (*[]float32, bool)`

GetBoundsOk returns a tuple with the Bounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounds

`func (o *ContrastSetHistogram) SetBounds(v []float32)`

SetBounds sets Bounds field to given value.


### GetTargetCounts

`func (o *ContrastSetHistogram) GetTargetCounts() []float32`

GetTargetCounts returns the TargetCounts field if non-nil, zero value otherwise.

### GetTargetCountsOk

`func (o *ContrastSetHistogram) GetTargetCountsOk() (*[]float32, bool)`

GetTargetCountsOk returns a tuple with the TargetCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCounts

`func (o *ContrastSetHistogram) SetTargetCounts(v []float32)`

SetTargetCounts sets TargetCounts field to given value.


### GetBaselineCounts

`func (o *ContrastSetHistogram) GetBaselineCounts() []float32`

GetBaselineCounts returns the BaselineCounts field if non-nil, zero value otherwise.

### GetBaselineCountsOk

`func (o *ContrastSetHistogram) GetBaselineCountsOk() (*[]float32, bool)`

GetBaselineCountsOk returns a tuple with the BaselineCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineCounts

`func (o *ContrastSetHistogram) SetBaselineCounts(v []float32)`

SetBaselineCounts sets BaselineCounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


