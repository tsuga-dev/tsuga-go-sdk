# InputMonitorConfigurationLogErrorPatternFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamIds** | **[]string** | Team IDs whose logs are searched for new error patterns. Tsuga resolves these team IDs to team names when exporting monitor assets. | 
**Env** | **string** | Environment whose logs are searched for new error patterns. | 
**Service** | Pointer to **string** | Optional service name whose logs are searched for new error patterns. If omitted, the monitor searches all services matching the team and environment filter. | [optional] 

## Methods

### NewInputMonitorConfigurationLogErrorPatternFilter

`func NewInputMonitorConfigurationLogErrorPatternFilter(teamIds []string, env string, ) *InputMonitorConfigurationLogErrorPatternFilter`

NewInputMonitorConfigurationLogErrorPatternFilter instantiates a new InputMonitorConfigurationLogErrorPatternFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputMonitorConfigurationLogErrorPatternFilterWithDefaults

`func NewInputMonitorConfigurationLogErrorPatternFilterWithDefaults() *InputMonitorConfigurationLogErrorPatternFilter`

NewInputMonitorConfigurationLogErrorPatternFilterWithDefaults instantiates a new InputMonitorConfigurationLogErrorPatternFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamIds

`func (o *InputMonitorConfigurationLogErrorPatternFilter) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *InputMonitorConfigurationLogErrorPatternFilter) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *InputMonitorConfigurationLogErrorPatternFilter) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.


### GetEnv

`func (o *InputMonitorConfigurationLogErrorPatternFilter) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *InputMonitorConfigurationLogErrorPatternFilter) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *InputMonitorConfigurationLogErrorPatternFilter) SetEnv(v string)`

SetEnv sets Env field to given value.


### GetService

`func (o *InputMonitorConfigurationLogErrorPatternFilter) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *InputMonitorConfigurationLogErrorPatternFilter) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *InputMonitorConfigurationLogErrorPatternFilter) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *InputMonitorConfigurationLogErrorPatternFilter) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


