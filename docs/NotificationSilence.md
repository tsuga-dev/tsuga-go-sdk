# NotificationSilence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Tsuga-generated notification silence ID assigned when the silence is created. | 
**Name** | **string** | Display name of the silence | 
**Reason** | **string** | Free-text explanation of why notifications are being silenced, shown to other operators reviewing active silences. | 
**Owner** | **string** | Team ID that owns and manages the silence | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | [optional] 
**IsActive** | **bool** | Whether the silence suppresses matching notifications while its schedule is active. When false, the silence is kept but never suppresses anything. | 
**Schedule** | [**NotificationSilenceSchedule**](NotificationSilenceSchedule.md) |  | 
**NotificationRuleIds** | Pointer to **[]string** | Notification rule IDs this silence applies to directly, matching any notification produced by one of these rules regardless of priority or transition type. Set only when &#x60;prioritiesFilter&#x60;, &#x60;transitionTypesFilter&#x60;, and &#x60;queryString&#x60; are not used to filter alerts. | [optional] 
**QueryString** | Pointer to **string** | Query on the alerting monitor tags and group-by dimensions, e.g. &#x60;env:prod AND service:bridge&#x60;. The monitor itself can be targeted with &#x60;monitor_id:&lt;monitor id&gt;&#x60;. | [optional] 
**TeamsFilter** | [**RuleTeamsFilter**](RuleTeamsFilter.md) |  | 
**PrioritiesFilter** | **[]float32** | Monitor priorities that must match for this silence to apply. An empty array matches every priority. | 
**TransitionTypesFilter** | **[]string** | Alert state transitions that must match for this silence to apply. An empty array matches every transition type. | 

## Methods

### NewNotificationSilence

`func NewNotificationSilence(id string, name string, reason string, owner string, isActive bool, schedule NotificationSilenceSchedule, teamsFilter RuleTeamsFilter, prioritiesFilter []float32, transitionTypesFilter []string, ) *NotificationSilence`

NewNotificationSilence instantiates a new NotificationSilence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSilenceWithDefaults

`func NewNotificationSilenceWithDefaults() *NotificationSilence`

NewNotificationSilenceWithDefaults instantiates a new NotificationSilence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationSilence) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationSilence) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationSilence) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *NotificationSilence) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationSilence) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationSilence) SetName(v string)`

SetName sets Name field to given value.


### GetReason

`func (o *NotificationSilence) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *NotificationSilence) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *NotificationSilence) SetReason(v string)`

SetReason sets Reason field to given value.


### GetOwner

`func (o *NotificationSilence) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NotificationSilence) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NotificationSilence) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetTags

`func (o *NotificationSilence) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NotificationSilence) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NotificationSilence) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NotificationSilence) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsActive

`func (o *NotificationSilence) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *NotificationSilence) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *NotificationSilence) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetSchedule

`func (o *NotificationSilence) GetSchedule() NotificationSilenceSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *NotificationSilence) GetScheduleOk() (*NotificationSilenceSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *NotificationSilence) SetSchedule(v NotificationSilenceSchedule)`

SetSchedule sets Schedule field to given value.


### GetNotificationRuleIds

`func (o *NotificationSilence) GetNotificationRuleIds() []string`

GetNotificationRuleIds returns the NotificationRuleIds field if non-nil, zero value otherwise.

### GetNotificationRuleIdsOk

`func (o *NotificationSilence) GetNotificationRuleIdsOk() (*[]string, bool)`

GetNotificationRuleIdsOk returns a tuple with the NotificationRuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationRuleIds

`func (o *NotificationSilence) SetNotificationRuleIds(v []string)`

SetNotificationRuleIds sets NotificationRuleIds field to given value.

### HasNotificationRuleIds

`func (o *NotificationSilence) HasNotificationRuleIds() bool`

HasNotificationRuleIds returns a boolean if a field has been set.

### GetQueryString

`func (o *NotificationSilence) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *NotificationSilence) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *NotificationSilence) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *NotificationSilence) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetTeamsFilter

`func (o *NotificationSilence) GetTeamsFilter() RuleTeamsFilter`

GetTeamsFilter returns the TeamsFilter field if non-nil, zero value otherwise.

### GetTeamsFilterOk

`func (o *NotificationSilence) GetTeamsFilterOk() (*RuleTeamsFilter, bool)`

GetTeamsFilterOk returns a tuple with the TeamsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsFilter

`func (o *NotificationSilence) SetTeamsFilter(v RuleTeamsFilter)`

SetTeamsFilter sets TeamsFilter field to given value.


### GetPrioritiesFilter

`func (o *NotificationSilence) GetPrioritiesFilter() []float32`

GetPrioritiesFilter returns the PrioritiesFilter field if non-nil, zero value otherwise.

### GetPrioritiesFilterOk

`func (o *NotificationSilence) GetPrioritiesFilterOk() (*[]float32, bool)`

GetPrioritiesFilterOk returns a tuple with the PrioritiesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioritiesFilter

`func (o *NotificationSilence) SetPrioritiesFilter(v []float32)`

SetPrioritiesFilter sets PrioritiesFilter field to given value.


### GetTransitionTypesFilter

`func (o *NotificationSilence) GetTransitionTypesFilter() []string`

GetTransitionTypesFilter returns the TransitionTypesFilter field if non-nil, zero value otherwise.

### GetTransitionTypesFilterOk

`func (o *NotificationSilence) GetTransitionTypesFilterOk() (*[]string, bool)`

GetTransitionTypesFilterOk returns a tuple with the TransitionTypesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionTypesFilter

`func (o *NotificationSilence) SetTransitionTypesFilter(v []string)`

SetTransitionTypesFilter sets TransitionTypesFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


