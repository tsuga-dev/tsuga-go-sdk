# UpdateSloRequestAlertsInnerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Alert configuration kind. Set to &#x60;burn-rate&#x60; for alerts that fire when the SLO consumes error budget faster than the configured multiplier. | 
**BurnRate** | **float32** | Burn-rate multiplier that triggers this alert. Required for burn-rate alerts; valid input is 1 through 100. Non-template values are accepted; short and long evaluation windows are derived from the predefined template with the closest burn rate. | 
**Threshold** | **float32** | SLI percentage threshold for this alert. Required for threshold alerts; valid input is greater than 0 and less than 100. The alert passes when SLI is at or above this value and fires when SLI is strictly below it. | 

## Methods

### NewUpdateSloRequestAlertsInnerConfiguration

`func NewUpdateSloRequestAlertsInnerConfiguration(type_ string, burnRate float32, threshold float32, ) *UpdateSloRequestAlertsInnerConfiguration`

NewUpdateSloRequestAlertsInnerConfiguration instantiates a new UpdateSloRequestAlertsInnerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSloRequestAlertsInnerConfigurationWithDefaults

`func NewUpdateSloRequestAlertsInnerConfigurationWithDefaults() *UpdateSloRequestAlertsInnerConfiguration`

NewUpdateSloRequestAlertsInnerConfigurationWithDefaults instantiates a new UpdateSloRequestAlertsInnerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateSloRequestAlertsInnerConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateSloRequestAlertsInnerConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateSloRequestAlertsInnerConfiguration) SetType(v string)`

SetType sets Type field to given value.


### GetBurnRate

`func (o *UpdateSloRequestAlertsInnerConfiguration) GetBurnRate() float32`

GetBurnRate returns the BurnRate field if non-nil, zero value otherwise.

### GetBurnRateOk

`func (o *UpdateSloRequestAlertsInnerConfiguration) GetBurnRateOk() (*float32, bool)`

GetBurnRateOk returns a tuple with the BurnRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnRate

`func (o *UpdateSloRequestAlertsInnerConfiguration) SetBurnRate(v float32)`

SetBurnRate sets BurnRate field to given value.


### GetThreshold

`func (o *UpdateSloRequestAlertsInnerConfiguration) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *UpdateSloRequestAlertsInnerConfiguration) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *UpdateSloRequestAlertsInnerConfiguration) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


