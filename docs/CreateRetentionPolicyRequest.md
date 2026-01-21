# CreateRetentionPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | Pointer to **string** |  | [optional] 
**TeamId** | Pointer to **string** |  | [optional] 
**DataSource** | **string** |  | 
**DurationDays** | **string** |  | 
**IsEnabled** | **bool** |  | 

## Methods

### NewCreateRetentionPolicyRequest

`func NewCreateRetentionPolicyRequest(dataSource string, durationDays string, isEnabled bool, ) *CreateRetentionPolicyRequest`

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

`func (o *CreateRetentionPolicyRequest) GetDurationDays() string`

GetDurationDays returns the DurationDays field if non-nil, zero value otherwise.

### GetDurationDaysOk

`func (o *CreateRetentionPolicyRequest) GetDurationDaysOk() (*string, bool)`

GetDurationDaysOk returns a tuple with the DurationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationDays

`func (o *CreateRetentionPolicyRequest) SetDurationDays(v string)`

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


