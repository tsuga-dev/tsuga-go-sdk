# ThresholdAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Alert configuration kind. Set to &#x60;threshold&#x60; for alerts that fire when the SLO indicator falls below a percentage threshold. | 
**Threshold** | **float32** | SLI percentage threshold for this alert. Required for threshold alerts; valid input is greater than 0 and less than 100. The alert passes when SLI is at or above this value and fires when SLI is strictly below it. | 

## Methods

### NewThresholdAlert

`func NewThresholdAlert(type_ string, threshold float32, ) *ThresholdAlert`

NewThresholdAlert instantiates a new ThresholdAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdAlertWithDefaults

`func NewThresholdAlertWithDefaults() *ThresholdAlert`

NewThresholdAlertWithDefaults instantiates a new ThresholdAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ThresholdAlert) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThresholdAlert) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThresholdAlert) SetType(v string)`

SetType sets Type field to given value.


### GetThreshold

`func (o *ThresholdAlert) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ThresholdAlert) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ThresholdAlert) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


