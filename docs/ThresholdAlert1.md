# ThresholdAlert1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Alert configuration kind. Set to &#x60;threshold&#x60; for alerts that fire when the SLO indicator falls below a percentage threshold. | 
**Threshold** | **float32** | SLI percentage threshold for this alert. The alert passes when SLI is at or above this value and fires when SLI is strictly below it. | 

## Methods

### NewThresholdAlert1

`func NewThresholdAlert1(type_ string, threshold float32, ) *ThresholdAlert1`

NewThresholdAlert1 instantiates a new ThresholdAlert1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdAlert1WithDefaults

`func NewThresholdAlert1WithDefaults() *ThresholdAlert1`

NewThresholdAlert1WithDefaults instantiates a new ThresholdAlert1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ThresholdAlert1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThresholdAlert1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThresholdAlert1) SetType(v string)`

SetType sets Type field to given value.


### GetThreshold

`func (o *ThresholdAlert1) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ThresholdAlert1) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ThresholdAlert1) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


