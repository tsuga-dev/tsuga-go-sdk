# BurnRateAlert1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Alert configuration kind. Set to &#x60;burn-rate&#x60; for alerts that fire when the SLO consumes error budget faster than the configured multiplier. | 
**BurnRate** | **float32** | Burn-rate multiplier that triggers this alert. The short and long evaluation windows are derived from the predefined template with the closest burn rate. | 

## Methods

### NewBurnRateAlert1

`func NewBurnRateAlert1(type_ string, burnRate float32, ) *BurnRateAlert1`

NewBurnRateAlert1 instantiates a new BurnRateAlert1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBurnRateAlert1WithDefaults

`func NewBurnRateAlert1WithDefaults() *BurnRateAlert1`

NewBurnRateAlert1WithDefaults instantiates a new BurnRateAlert1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BurnRateAlert1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BurnRateAlert1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BurnRateAlert1) SetType(v string)`

SetType sets Type field to given value.


### GetBurnRate

`func (o *BurnRateAlert1) GetBurnRate() float32`

GetBurnRate returns the BurnRate field if non-nil, zero value otherwise.

### GetBurnRateOk

`func (o *BurnRateAlert1) GetBurnRateOk() (*float32, bool)`

GetBurnRateOk returns a tuple with the BurnRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnRate

`func (o *BurnRateAlert1) SetBurnRate(v float32)`

SetBurnRate sets BurnRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


