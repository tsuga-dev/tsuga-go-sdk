# RetentionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Env** | Pointer to **string** |  | [optional] 
**TeamId** | Pointer to **string** |  | [optional] 
**DataSource** | **string** |  | 
**DurationDays** | **string** |  | 
**IsEnabled** | **bool** |  | 

## Methods

### NewRetentionPolicy

`func NewRetentionPolicy(id string, dataSource string, durationDays string, isEnabled bool, ) *RetentionPolicy`

NewRetentionPolicy instantiates a new RetentionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetentionPolicyWithDefaults

`func NewRetentionPolicyWithDefaults() *RetentionPolicy`

NewRetentionPolicyWithDefaults instantiates a new RetentionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RetentionPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RetentionPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RetentionPolicy) SetId(v string)`

SetId sets Id field to given value.


### GetEnv

`func (o *RetentionPolicy) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *RetentionPolicy) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *RetentionPolicy) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *RetentionPolicy) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetTeamId

`func (o *RetentionPolicy) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *RetentionPolicy) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *RetentionPolicy) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *RetentionPolicy) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetDataSource

`func (o *RetentionPolicy) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *RetentionPolicy) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *RetentionPolicy) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.


### GetDurationDays

`func (o *RetentionPolicy) GetDurationDays() string`

GetDurationDays returns the DurationDays field if non-nil, zero value otherwise.

### GetDurationDaysOk

`func (o *RetentionPolicy) GetDurationDaysOk() (*string, bool)`

GetDurationDaysOk returns a tuple with the DurationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationDays

`func (o *RetentionPolicy) SetDurationDays(v string)`

SetDurationDays sets DurationDays field to given value.


### GetIsEnabled

`func (o *RetentionPolicy) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *RetentionPolicy) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *RetentionPolicy) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


