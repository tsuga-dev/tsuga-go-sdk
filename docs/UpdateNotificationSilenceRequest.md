# UpdateNotificationSilenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the silence | 
**Reason** | **string** | Free-text explanation of why notifications are being silenced, shown to other operators reviewing active silences. | 
**Owner** | **string** | Team ID that owns and manages the silence | 
**Tags** | Pointer to [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 
**IsActive** | **bool** | Set to true for the silence to suppress matching notifications while its schedule is active. Set to false to keep the silence without suppressing anything. | 
**Schedule** | [**CreateNotificationSilenceRequestSchedule**](CreateNotificationSilenceRequestSchedule.md) |  | 
**NotificationRuleIds** | Pointer to **[]string** | Notification rule IDs this silence applies to directly, matching any notification produced by one of these rules regardless of priority or transition type. Mutually exclusive with &#x60;prioritiesFilter&#x60;, &#x60;transitionTypesFilter&#x60;, and &#x60;queryString&#x60;: a request combining this field with any of those is rejected. | [optional] 
**QueryString** | Pointer to **NullableString** | Query on the alerting monitor tags and group-by dimensions, e.g. &#x60;env:prod AND service:bridge&#x60;. The monitor itself can be targeted with &#x60;monitor_id:&lt;monitor id&gt;&#x60;. | [optional] 
**TeamsFilter** | [**CreateNotificationRuleRequestTeamsFilter**](CreateNotificationRuleRequestTeamsFilter.md) |  | 
**PrioritiesFilter** | **[]float32** | Monitor priorities that must match for this silence to apply. An empty array matches every priority. | 
**TransitionTypesFilter** | **[]string** | Alert state transitions that must match for this silence to apply. An empty array matches every transition type. | 

## Methods

### NewUpdateNotificationSilenceRequest

`func NewUpdateNotificationSilenceRequest(name string, reason string, owner string, isActive bool, schedule CreateNotificationSilenceRequestSchedule, teamsFilter CreateNotificationRuleRequestTeamsFilter, prioritiesFilter []float32, transitionTypesFilter []string, ) *UpdateNotificationSilenceRequest`

NewUpdateNotificationSilenceRequest instantiates a new UpdateNotificationSilenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNotificationSilenceRequestWithDefaults

`func NewUpdateNotificationSilenceRequestWithDefaults() *UpdateNotificationSilenceRequest`

NewUpdateNotificationSilenceRequestWithDefaults instantiates a new UpdateNotificationSilenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNotificationSilenceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNotificationSilenceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNotificationSilenceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetReason

`func (o *UpdateNotificationSilenceRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UpdateNotificationSilenceRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UpdateNotificationSilenceRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetOwner

`func (o *UpdateNotificationSilenceRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateNotificationSilenceRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateNotificationSilenceRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetTags

`func (o *UpdateNotificationSilenceRequest) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateNotificationSilenceRequest) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateNotificationSilenceRequest) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateNotificationSilenceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsActive

`func (o *UpdateNotificationSilenceRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateNotificationSilenceRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateNotificationSilenceRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetSchedule

`func (o *UpdateNotificationSilenceRequest) GetSchedule() CreateNotificationSilenceRequestSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *UpdateNotificationSilenceRequest) GetScheduleOk() (*CreateNotificationSilenceRequestSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *UpdateNotificationSilenceRequest) SetSchedule(v CreateNotificationSilenceRequestSchedule)`

SetSchedule sets Schedule field to given value.


### GetNotificationRuleIds

`func (o *UpdateNotificationSilenceRequest) GetNotificationRuleIds() []string`

GetNotificationRuleIds returns the NotificationRuleIds field if non-nil, zero value otherwise.

### GetNotificationRuleIdsOk

`func (o *UpdateNotificationSilenceRequest) GetNotificationRuleIdsOk() (*[]string, bool)`

GetNotificationRuleIdsOk returns a tuple with the NotificationRuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationRuleIds

`func (o *UpdateNotificationSilenceRequest) SetNotificationRuleIds(v []string)`

SetNotificationRuleIds sets NotificationRuleIds field to given value.

### HasNotificationRuleIds

`func (o *UpdateNotificationSilenceRequest) HasNotificationRuleIds() bool`

HasNotificationRuleIds returns a boolean if a field has been set.

### SetNotificationRuleIdsNil

`func (o *UpdateNotificationSilenceRequest) SetNotificationRuleIdsNil(b bool)`

 SetNotificationRuleIdsNil sets the value for NotificationRuleIds to be an explicit nil

### UnsetNotificationRuleIds
`func (o *UpdateNotificationSilenceRequest) UnsetNotificationRuleIds()`

UnsetNotificationRuleIds ensures that no value is present for NotificationRuleIds, not even an explicit nil
### GetQueryString

`func (o *UpdateNotificationSilenceRequest) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *UpdateNotificationSilenceRequest) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *UpdateNotificationSilenceRequest) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *UpdateNotificationSilenceRequest) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### SetQueryStringNil

`func (o *UpdateNotificationSilenceRequest) SetQueryStringNil(b bool)`

 SetQueryStringNil sets the value for QueryString to be an explicit nil

### UnsetQueryString
`func (o *UpdateNotificationSilenceRequest) UnsetQueryString()`

UnsetQueryString ensures that no value is present for QueryString, not even an explicit nil
### GetTeamsFilter

`func (o *UpdateNotificationSilenceRequest) GetTeamsFilter() CreateNotificationRuleRequestTeamsFilter`

GetTeamsFilter returns the TeamsFilter field if non-nil, zero value otherwise.

### GetTeamsFilterOk

`func (o *UpdateNotificationSilenceRequest) GetTeamsFilterOk() (*CreateNotificationRuleRequestTeamsFilter, bool)`

GetTeamsFilterOk returns a tuple with the TeamsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsFilter

`func (o *UpdateNotificationSilenceRequest) SetTeamsFilter(v CreateNotificationRuleRequestTeamsFilter)`

SetTeamsFilter sets TeamsFilter field to given value.


### GetPrioritiesFilter

`func (o *UpdateNotificationSilenceRequest) GetPrioritiesFilter() []float32`

GetPrioritiesFilter returns the PrioritiesFilter field if non-nil, zero value otherwise.

### GetPrioritiesFilterOk

`func (o *UpdateNotificationSilenceRequest) GetPrioritiesFilterOk() (*[]float32, bool)`

GetPrioritiesFilterOk returns a tuple with the PrioritiesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioritiesFilter

`func (o *UpdateNotificationSilenceRequest) SetPrioritiesFilter(v []float32)`

SetPrioritiesFilter sets PrioritiesFilter field to given value.


### GetTransitionTypesFilter

`func (o *UpdateNotificationSilenceRequest) GetTransitionTypesFilter() []string`

GetTransitionTypesFilter returns the TransitionTypesFilter field if non-nil, zero value otherwise.

### GetTransitionTypesFilterOk

`func (o *UpdateNotificationSilenceRequest) GetTransitionTypesFilterOk() (*[]string, bool)`

GetTransitionTypesFilterOk returns a tuple with the TransitionTypesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionTypesFilter

`func (o *UpdateNotificationSilenceRequest) SetTransitionTypesFilter(v []string)`

SetTransitionTypesFilter sets TransitionTypesFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


