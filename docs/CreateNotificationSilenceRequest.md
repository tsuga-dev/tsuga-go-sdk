# CreateNotificationSilenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the silence | 
**Reason** | **string** | Reason for the silence | 
**Owner** | **string** | Team ID that owns and manages the silence | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List of key/value tags applied to the resource | [optional] 
**IsActive** | **bool** | Whether the silence is currently enabled | 
**Schedule** | [**CreateNotificationSilenceRequestSchedule**](CreateNotificationSilenceRequestSchedule.md) |  | 
**NotificationRuleIds** | Pointer to **[]string** | Notification rule IDs this silence applies to | [optional] 
**QueryString** | Pointer to **string** | Query string to filter which alerts this silence applies to | [optional] 
**TeamsFilter** | [**CreateNotificationRuleRequestTeamsFilter**](CreateNotificationRuleRequestTeamsFilter.md) |  | 
**PrioritiesFilter** | **[]float32** | Monitor priorities to filter which alerts this silence applies to | 
**TransitionTypesFilter** | **[]string** | Transition types to filter which alerts this silence applies to | 

## Methods

### NewCreateNotificationSilenceRequest

`func NewCreateNotificationSilenceRequest(name string, reason string, owner string, isActive bool, schedule CreateNotificationSilenceRequestSchedule, teamsFilter CreateNotificationRuleRequestTeamsFilter, prioritiesFilter []float32, transitionTypesFilter []string, ) *CreateNotificationSilenceRequest`

NewCreateNotificationSilenceRequest instantiates a new CreateNotificationSilenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationSilenceRequestWithDefaults

`func NewCreateNotificationSilenceRequestWithDefaults() *CreateNotificationSilenceRequest`

NewCreateNotificationSilenceRequestWithDefaults instantiates a new CreateNotificationSilenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNotificationSilenceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNotificationSilenceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNotificationSilenceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetReason

`func (o *CreateNotificationSilenceRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateNotificationSilenceRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateNotificationSilenceRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetOwner

`func (o *CreateNotificationSilenceRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateNotificationSilenceRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateNotificationSilenceRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetTags

`func (o *CreateNotificationSilenceRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateNotificationSilenceRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateNotificationSilenceRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateNotificationSilenceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsActive

`func (o *CreateNotificationSilenceRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateNotificationSilenceRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateNotificationSilenceRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetSchedule

`func (o *CreateNotificationSilenceRequest) GetSchedule() CreateNotificationSilenceRequestSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CreateNotificationSilenceRequest) GetScheduleOk() (*CreateNotificationSilenceRequestSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CreateNotificationSilenceRequest) SetSchedule(v CreateNotificationSilenceRequestSchedule)`

SetSchedule sets Schedule field to given value.


### GetNotificationRuleIds

`func (o *CreateNotificationSilenceRequest) GetNotificationRuleIds() []string`

GetNotificationRuleIds returns the NotificationRuleIds field if non-nil, zero value otherwise.

### GetNotificationRuleIdsOk

`func (o *CreateNotificationSilenceRequest) GetNotificationRuleIdsOk() (*[]string, bool)`

GetNotificationRuleIdsOk returns a tuple with the NotificationRuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationRuleIds

`func (o *CreateNotificationSilenceRequest) SetNotificationRuleIds(v []string)`

SetNotificationRuleIds sets NotificationRuleIds field to given value.

### HasNotificationRuleIds

`func (o *CreateNotificationSilenceRequest) HasNotificationRuleIds() bool`

HasNotificationRuleIds returns a boolean if a field has been set.

### GetQueryString

`func (o *CreateNotificationSilenceRequest) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *CreateNotificationSilenceRequest) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *CreateNotificationSilenceRequest) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *CreateNotificationSilenceRequest) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetTeamsFilter

`func (o *CreateNotificationSilenceRequest) GetTeamsFilter() CreateNotificationRuleRequestTeamsFilter`

GetTeamsFilter returns the TeamsFilter field if non-nil, zero value otherwise.

### GetTeamsFilterOk

`func (o *CreateNotificationSilenceRequest) GetTeamsFilterOk() (*CreateNotificationRuleRequestTeamsFilter, bool)`

GetTeamsFilterOk returns a tuple with the TeamsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsFilter

`func (o *CreateNotificationSilenceRequest) SetTeamsFilter(v CreateNotificationRuleRequestTeamsFilter)`

SetTeamsFilter sets TeamsFilter field to given value.


### GetPrioritiesFilter

`func (o *CreateNotificationSilenceRequest) GetPrioritiesFilter() []float32`

GetPrioritiesFilter returns the PrioritiesFilter field if non-nil, zero value otherwise.

### GetPrioritiesFilterOk

`func (o *CreateNotificationSilenceRequest) GetPrioritiesFilterOk() (*[]float32, bool)`

GetPrioritiesFilterOk returns a tuple with the PrioritiesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioritiesFilter

`func (o *CreateNotificationSilenceRequest) SetPrioritiesFilter(v []float32)`

SetPrioritiesFilter sets PrioritiesFilter field to given value.


### GetTransitionTypesFilter

`func (o *CreateNotificationSilenceRequest) GetTransitionTypesFilter() []string`

GetTransitionTypesFilter returns the TransitionTypesFilter field if non-nil, zero value otherwise.

### GetTransitionTypesFilterOk

`func (o *CreateNotificationSilenceRequest) GetTransitionTypesFilterOk() (*[]string, bool)`

GetTransitionTypesFilterOk returns a tuple with the TransitionTypesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionTypesFilter

`func (o *CreateNotificationSilenceRequest) SetTransitionTypesFilter(v []string)`

SetTransitionTypesFilter sets TransitionTypesFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


