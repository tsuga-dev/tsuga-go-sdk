# SloAlertsInnerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Alert configuration kind. Set to &#x60;burn-rate&#x60; for alerts that fire when the SLO consumes error budget faster than the configured multiplier. | 
**BurnRate** | **float32** | Burn-rate multiplier that triggers this alert. The short and long evaluation windows are derived from the predefined template with the closest burn rate. | 
**Threshold** | **float32** | SLI percentage threshold for this alert. The alert passes when SLI is at or above this value and fires when SLI is strictly below it. | 

## Methods

### NewSloAlertsInnerConfiguration

`func NewSloAlertsInnerConfiguration(type_ string, burnRate float32, threshold float32, ) *SloAlertsInnerConfiguration`

NewSloAlertsInnerConfiguration instantiates a new SloAlertsInnerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloAlertsInnerConfigurationWithDefaults

`func NewSloAlertsInnerConfigurationWithDefaults() *SloAlertsInnerConfiguration`

NewSloAlertsInnerConfigurationWithDefaults instantiates a new SloAlertsInnerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SloAlertsInnerConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SloAlertsInnerConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SloAlertsInnerConfiguration) SetType(v string)`

SetType sets Type field to given value.


### GetBurnRate

`func (o *SloAlertsInnerConfiguration) GetBurnRate() float32`

GetBurnRate returns the BurnRate field if non-nil, zero value otherwise.

### GetBurnRateOk

`func (o *SloAlertsInnerConfiguration) GetBurnRateOk() (*float32, bool)`

GetBurnRateOk returns a tuple with the BurnRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnRate

`func (o *SloAlertsInnerConfiguration) SetBurnRate(v float32)`

SetBurnRate sets BurnRate field to given value.


### GetThreshold

`func (o *SloAlertsInnerConfiguration) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *SloAlertsInnerConfiguration) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *SloAlertsInnerConfiguration) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


