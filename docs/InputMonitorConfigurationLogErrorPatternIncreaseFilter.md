# InputMonitorConfigurationLogErrorPatternIncreaseFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamIds** | Pointer to **[]string** | Team IDs whose error logs are watched for volume increases. Tsuga resolves these team IDs to team names when exporting monitor assets. Omit to watch every team, in which case &#x60;services&#x60; is required. | [optional] 
**Env** | **string** | Environment whose error logs are watched for volume increases. | 
**Services** | Pointer to **[]string** | Service names whose error logs are watched for volume increases. Omit to watch every service matching the team and environment scope. | [optional] 

## Methods

### NewInputMonitorConfigurationLogErrorPatternIncreaseFilter

`func NewInputMonitorConfigurationLogErrorPatternIncreaseFilter(env string, ) *InputMonitorConfigurationLogErrorPatternIncreaseFilter`

NewInputMonitorConfigurationLogErrorPatternIncreaseFilter instantiates a new InputMonitorConfigurationLogErrorPatternIncreaseFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputMonitorConfigurationLogErrorPatternIncreaseFilterWithDefaults

`func NewInputMonitorConfigurationLogErrorPatternIncreaseFilterWithDefaults() *InputMonitorConfigurationLogErrorPatternIncreaseFilter`

NewInputMonitorConfigurationLogErrorPatternIncreaseFilterWithDefaults instantiates a new InputMonitorConfigurationLogErrorPatternIncreaseFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamIds

`func (o *InputMonitorConfigurationLogErrorPatternIncreaseFilter) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *InputMonitorConfigurationLogErrorPatternIncreaseFilter) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *InputMonitorConfigurationLogErrorPatternIncreaseFilter) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *InputMonitorConfigurationLogErrorPatternIncreaseFilter) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

### GetEnv

`func (o *InputMonitorConfigurationLogErrorPatternIncreaseFilter) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *InputMonitorConfigurationLogErrorPatternIncreaseFilter) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *InputMonitorConfigurationLogErrorPatternIncreaseFilter) SetEnv(v string)`

SetEnv sets Env field to given value.


### GetServices

`func (o *InputMonitorConfigurationLogErrorPatternIncreaseFilter) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *InputMonitorConfigurationLogErrorPatternIncreaseFilter) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *InputMonitorConfigurationLogErrorPatternIncreaseFilter) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *InputMonitorConfigurationLogErrorPatternIncreaseFilter) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


