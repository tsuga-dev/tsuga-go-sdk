# ContrastSetShift

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuantileLift** | **float32** | How far apart the two groups are, as a ratio: 1 means no change, 3 means about three times apart. It does not tell which group is larger. Read &#x60;histogram&#x60; for the direction. | 
**Ks** | **float32** | Kolmogorov-Smirnov statistic, from 0 to 1: the largest gap between the two cumulative distributions. | 
**PValue** | **float32** | Probability that a difference this large would occur if the two groups had the same distribution. Smaller is stronger. | 
**TargetSampleCount** | **float32** | Number of target spans that carry the attribute. | 
**BaselineSampleCount** | **float32** | Number of baseline spans that carry the attribute. | 
**Histogram** | [**ContrastSetHistogram**](ContrastSetHistogram.md) |  | 

## Methods

### NewContrastSetShift

`func NewContrastSetShift(quantileLift float32, ks float32, pValue float32, targetSampleCount float32, baselineSampleCount float32, histogram ContrastSetHistogram, ) *ContrastSetShift`

NewContrastSetShift instantiates a new ContrastSetShift object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContrastSetShiftWithDefaults

`func NewContrastSetShiftWithDefaults() *ContrastSetShift`

NewContrastSetShiftWithDefaults instantiates a new ContrastSetShift object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantileLift

`func (o *ContrastSetShift) GetQuantileLift() float32`

GetQuantileLift returns the QuantileLift field if non-nil, zero value otherwise.

### GetQuantileLiftOk

`func (o *ContrastSetShift) GetQuantileLiftOk() (*float32, bool)`

GetQuantileLiftOk returns a tuple with the QuantileLift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantileLift

`func (o *ContrastSetShift) SetQuantileLift(v float32)`

SetQuantileLift sets QuantileLift field to given value.


### GetKs

`func (o *ContrastSetShift) GetKs() float32`

GetKs returns the Ks field if non-nil, zero value otherwise.

### GetKsOk

`func (o *ContrastSetShift) GetKsOk() (*float32, bool)`

GetKsOk returns a tuple with the Ks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKs

`func (o *ContrastSetShift) SetKs(v float32)`

SetKs sets Ks field to given value.


### GetPValue

`func (o *ContrastSetShift) GetPValue() float32`

GetPValue returns the PValue field if non-nil, zero value otherwise.

### GetPValueOk

`func (o *ContrastSetShift) GetPValueOk() (*float32, bool)`

GetPValueOk returns a tuple with the PValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPValue

`func (o *ContrastSetShift) SetPValue(v float32)`

SetPValue sets PValue field to given value.


### GetTargetSampleCount

`func (o *ContrastSetShift) GetTargetSampleCount() float32`

GetTargetSampleCount returns the TargetSampleCount field if non-nil, zero value otherwise.

### GetTargetSampleCountOk

`func (o *ContrastSetShift) GetTargetSampleCountOk() (*float32, bool)`

GetTargetSampleCountOk returns a tuple with the TargetSampleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSampleCount

`func (o *ContrastSetShift) SetTargetSampleCount(v float32)`

SetTargetSampleCount sets TargetSampleCount field to given value.


### GetBaselineSampleCount

`func (o *ContrastSetShift) GetBaselineSampleCount() float32`

GetBaselineSampleCount returns the BaselineSampleCount field if non-nil, zero value otherwise.

### GetBaselineSampleCountOk

`func (o *ContrastSetShift) GetBaselineSampleCountOk() (*float32, bool)`

GetBaselineSampleCountOk returns a tuple with the BaselineSampleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineSampleCount

`func (o *ContrastSetShift) SetBaselineSampleCount(v float32)`

SetBaselineSampleCount sets BaselineSampleCount field to given value.


### GetHistogram

`func (o *ContrastSetShift) GetHistogram() ContrastSetHistogram`

GetHistogram returns the Histogram field if non-nil, zero value otherwise.

### GetHistogramOk

`func (o *ContrastSetShift) GetHistogramOk() (*ContrastSetHistogram, bool)`

GetHistogramOk returns a tuple with the Histogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogram

`func (o *ContrastSetShift) SetHistogram(v ContrastSetHistogram)`

SetHistogram sets Histogram field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


