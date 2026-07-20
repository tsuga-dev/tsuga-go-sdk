# GaugeColorThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **float32** | Lower bound of the Gauge color threshold; runs up to the next threshold or the max | 
**Color** | **string** | Color applied to the band starting at this value | 

## Methods

### NewGaugeColorThreshold

`func NewGaugeColorThreshold(from float32, color string, ) *GaugeColorThreshold`

NewGaugeColorThreshold instantiates a new GaugeColorThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGaugeColorThresholdWithDefaults

`func NewGaugeColorThresholdWithDefaults() *GaugeColorThreshold`

NewGaugeColorThresholdWithDefaults instantiates a new GaugeColorThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *GaugeColorThreshold) GetFrom() float32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GaugeColorThreshold) GetFromOk() (*float32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GaugeColorThreshold) SetFrom(v float32)`

SetFrom sets From field to given value.


### GetColor

`func (o *GaugeColorThreshold) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *GaugeColorThreshold) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *GaugeColorThreshold) SetColor(v string)`

SetColor sets Color field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


