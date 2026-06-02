# RuleTargetsInnerRenotifyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** |  | 
**RenotificationStates** | **[]string** | Alert states that will trigger a renotification | 
**RenotifyIntervalMinutes** | **int32** | Minimum number of minutes to wait before renotifying | 

## Methods

### NewRuleTargetsInnerRenotifyConfig

`func NewRuleTargetsInnerRenotifyConfig(mode string, renotificationStates []string, renotifyIntervalMinutes int32, ) *RuleTargetsInnerRenotifyConfig`

NewRuleTargetsInnerRenotifyConfig instantiates a new RuleTargetsInnerRenotifyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTargetsInnerRenotifyConfigWithDefaults

`func NewRuleTargetsInnerRenotifyConfigWithDefaults() *RuleTargetsInnerRenotifyConfig`

NewRuleTargetsInnerRenotifyConfigWithDefaults instantiates a new RuleTargetsInnerRenotifyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *RuleTargetsInnerRenotifyConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RuleTargetsInnerRenotifyConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RuleTargetsInnerRenotifyConfig) SetMode(v string)`

SetMode sets Mode field to given value.


### GetRenotificationStates

`func (o *RuleTargetsInnerRenotifyConfig) GetRenotificationStates() []string`

GetRenotificationStates returns the RenotificationStates field if non-nil, zero value otherwise.

### GetRenotificationStatesOk

`func (o *RuleTargetsInnerRenotifyConfig) GetRenotificationStatesOk() (*[]string, bool)`

GetRenotificationStatesOk returns a tuple with the RenotificationStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenotificationStates

`func (o *RuleTargetsInnerRenotifyConfig) SetRenotificationStates(v []string)`

SetRenotificationStates sets RenotificationStates field to given value.


### GetRenotifyIntervalMinutes

`func (o *RuleTargetsInnerRenotifyConfig) GetRenotifyIntervalMinutes() int32`

GetRenotifyIntervalMinutes returns the RenotifyIntervalMinutes field if non-nil, zero value otherwise.

### GetRenotifyIntervalMinutesOk

`func (o *RuleTargetsInnerRenotifyConfig) GetRenotifyIntervalMinutesOk() (*int32, bool)`

GetRenotifyIntervalMinutesOk returns a tuple with the RenotifyIntervalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenotifyIntervalMinutes

`func (o *RuleTargetsInnerRenotifyConfig) SetRenotifyIntervalMinutes(v int32)`

SetRenotifyIntervalMinutes sets RenotifyIntervalMinutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


