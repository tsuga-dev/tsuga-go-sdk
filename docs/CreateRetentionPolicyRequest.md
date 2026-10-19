# CreateRetentionPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | Pointer to **NullableString** | Display name of the graph widget. Pass null to remove it. | [optional] 
**TeamId** | Pointer to **NullableString** | Display name of the graph widget. Pass null to remove it. | [optional] 
**DataSource** | **string** | Telemetry signal retained by this policy: logs, metrics, traces or rum. | 
**DurationDays** | **int32** | Number of days Tsuga should retain matching telemetry. Valid values are 1 through 3650. | 
**IsEnabled** | **bool** | Set to true for the policy to participate in retention resolution. | 

## Methods

### NewCreateRetentionPolicyRequest

`func NewCreateRetentionPolicyRequest(dataSource string, durationDays int32, isEnabled bool, ) *CreateRetentionPolicyRequest`

NewCreateRetentionPolicyRequest instantiates a new CreateRetentionPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRetentionPolicyRequestWithDefaults

`func NewCreateRetentionPolicyRequestWithDefaults() *CreateRetentionPolicyRequest`

NewCreateRetentionPolicyRequestWithDefaults instantiates a new CreateRetentionPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnv

`func (o *CreateRetentionPolicyRequest) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *CreateRetentionPolicyRequest) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *CreateRetentionPolicyRequest) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *CreateRetentionPolicyRequest) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### SetEnvNil

`func (o *CreateRetentionPolicyRequest) SetEnvNil(b bool)`

 SetEnvNil sets the value for Env to be an explicit nil

### UnsetEnv
`func (o *CreateRetentionPolicyRequest) UnsetEnv()`

UnsetEnv ensures that no value is present for Env, not even an explicit nil
### GetTeamId

`func (o *CreateRetentionPolicyRequest) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *CreateRetentionPolicyRequest) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *CreateRetentionPolicyRequest) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *CreateRetentionPolicyRequest) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### SetTeamIdNil

`func (o *CreateRetentionPolicyRequest) SetTeamIdNil(b bool)`

 SetTeamIdNil sets the value for TeamId to be an explicit nil

### UnsetTeamId
`func (o *CreateRetentionPolicyRequest) UnsetTeamId()`

UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil
### GetDataSource

`func (o *CreateRetentionPolicyRequest) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *CreateRetentionPolicyRequest) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *CreateRetentionPolicyRequest) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.


### GetDurationDays

`func (o *CreateRetentionPolicyRequest) GetDurationDays() int32`

GetDurationDays returns the DurationDays field if non-nil, zero value otherwise.

### GetDurationDaysOk

`func (o *CreateRetentionPolicyRequest) GetDurationDaysOk() (*int32, bool)`

GetDurationDaysOk returns a tuple with the DurationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationDays

`func (o *CreateRetentionPolicyRequest) SetDurationDays(v int32)`

SetDurationDays sets DurationDays field to given value.


### GetIsEnabled

`func (o *CreateRetentionPolicyRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CreateRetentionPolicyRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CreateRetentionPolicyRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


