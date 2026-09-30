# MonitorConfigurationLogErrorPatternIncreaseFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamIds** | Pointer to **[]string** | Team IDs whose error logs are watched for volume increases. Tsuga resolves these team IDs to team names when exporting monitor assets. Omit to watch every team, in which case a service scope is required. | [optional] 
**Env** | **string** | Environment whose error logs are watched for volume increases. | 
**Services** | Pointer to **[]string** | Service names whose error logs are watched for volume increases. Omit to watch every service matching the team and environment scope. | [optional] 

## Methods

### NewMonitorConfigurationLogErrorPatternIncreaseFilter

`func NewMonitorConfigurationLogErrorPatternIncreaseFilter(env string, ) *MonitorConfigurationLogErrorPatternIncreaseFilter`

NewMonitorConfigurationLogErrorPatternIncreaseFilter instantiates a new MonitorConfigurationLogErrorPatternIncreaseFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorConfigurationLogErrorPatternIncreaseFilterWithDefaults

`func NewMonitorConfigurationLogErrorPatternIncreaseFilterWithDefaults() *MonitorConfigurationLogErrorPatternIncreaseFilter`

NewMonitorConfigurationLogErrorPatternIncreaseFilterWithDefaults instantiates a new MonitorConfigurationLogErrorPatternIncreaseFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamIds

`func (o *MonitorConfigurationLogErrorPatternIncreaseFilter) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *MonitorConfigurationLogErrorPatternIncreaseFilter) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *MonitorConfigurationLogErrorPatternIncreaseFilter) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *MonitorConfigurationLogErrorPatternIncreaseFilter) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

### GetEnv

`func (o *MonitorConfigurationLogErrorPatternIncreaseFilter) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *MonitorConfigurationLogErrorPatternIncreaseFilter) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *MonitorConfigurationLogErrorPatternIncreaseFilter) SetEnv(v string)`

SetEnv sets Env field to given value.


### GetServices

`func (o *MonitorConfigurationLogErrorPatternIncreaseFilter) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *MonitorConfigurationLogErrorPatternIncreaseFilter) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *MonitorConfigurationLogErrorPatternIncreaseFilter) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *MonitorConfigurationLogErrorPatternIncreaseFilter) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


