# RuleTargetConfigSlackRenotifyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | Renotification mode. &#x60;each&#x60; sends repeat notifications for each matching monitor alert that remains in a configured state. | 
**RenotificationStates** | **[]string** | Monitor alert states that trigger repeat notifications while the target remains configured for renotification. | 
**RenotifyIntervalMinutes** | **int32** | Minimum number of minutes to wait after the previous notification. | 

## Methods

### NewRuleTargetConfigSlackRenotifyConfig

`func NewRuleTargetConfigSlackRenotifyConfig(mode string, renotificationStates []string, renotifyIntervalMinutes int32, ) *RuleTargetConfigSlackRenotifyConfig`

NewRuleTargetConfigSlackRenotifyConfig instantiates a new RuleTargetConfigSlackRenotifyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTargetConfigSlackRenotifyConfigWithDefaults

`func NewRuleTargetConfigSlackRenotifyConfigWithDefaults() *RuleTargetConfigSlackRenotifyConfig`

NewRuleTargetConfigSlackRenotifyConfigWithDefaults instantiates a new RuleTargetConfigSlackRenotifyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *RuleTargetConfigSlackRenotifyConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RuleTargetConfigSlackRenotifyConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RuleTargetConfigSlackRenotifyConfig) SetMode(v string)`

SetMode sets Mode field to given value.


### GetRenotificationStates

`func (o *RuleTargetConfigSlackRenotifyConfig) GetRenotificationStates() []string`

GetRenotificationStates returns the RenotificationStates field if non-nil, zero value otherwise.

### GetRenotificationStatesOk

`func (o *RuleTargetConfigSlackRenotifyConfig) GetRenotificationStatesOk() (*[]string, bool)`

GetRenotificationStatesOk returns a tuple with the RenotificationStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenotificationStates

`func (o *RuleTargetConfigSlackRenotifyConfig) SetRenotificationStates(v []string)`

SetRenotificationStates sets RenotificationStates field to given value.


### GetRenotifyIntervalMinutes

`func (o *RuleTargetConfigSlackRenotifyConfig) GetRenotifyIntervalMinutes() int32`

GetRenotifyIntervalMinutes returns the RenotifyIntervalMinutes field if non-nil, zero value otherwise.

### GetRenotifyIntervalMinutesOk

`func (o *RuleTargetConfigSlackRenotifyConfig) GetRenotifyIntervalMinutesOk() (*int32, bool)`

GetRenotifyIntervalMinutesOk returns a tuple with the RenotifyIntervalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenotifyIntervalMinutes

`func (o *RuleTargetConfigSlackRenotifyConfig) SetRenotifyIntervalMinutes(v int32)`

SetRenotifyIntervalMinutes sets RenotifyIntervalMinutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


