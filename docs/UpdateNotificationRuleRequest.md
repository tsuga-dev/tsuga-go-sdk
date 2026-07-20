# UpdateNotificationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the notification rule. | 
**QueryString** | Pointer to **string** | Optional query that narrows which alert transitions trigger the rule. Matches on the monitor transition group key and the monitor tags, e.g. &#x60;env:prod service:api&#x60;. Omit or leave empty to match regardless of tags. | [optional] 
**TeamsFilter** | [**CreateNotificationRuleRequestTeamsFilter**](CreateNotificationRuleRequestTeamsFilter.md) |  | 
**PrioritiesFilter** | **[]float32** | Monitor priorities that must match for this rule to fire. An empty array matches every priority. | 
**TransitionTypesFilter** | **[]string** | Alert state transitions that must match for this rule to fire. An empty array matches every transition type. | 
**ClusterIdsFilter** | Pointer to **[]string** | Cluster IDs that must match for this rule to fire. Omit it, or leave it empty, to match every cluster. | [optional] 
**Owner** | **string** | Team ID that owns and manages the rule | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key/value tags to apply to the resource. Up to 50 tags are accepted and tag policies may require specific keys or values. | [optional] 
**IsActive** | **bool** | Set to true for the rule to send notifications when its filters match. | 
**Targets** | [**[]CreateNotificationRuleRequestTargetsInner**](CreateNotificationRuleRequestTargetsInner.md) | Destinations that receive a notification whenever this rule matches an alert transition. At least one target is required. This list replaces the existing targets on update. | 

## Methods

### NewUpdateNotificationRuleRequest

`func NewUpdateNotificationRuleRequest(name string, teamsFilter CreateNotificationRuleRequestTeamsFilter, prioritiesFilter []float32, transitionTypesFilter []string, owner string, isActive bool, targets []CreateNotificationRuleRequestTargetsInner, ) *UpdateNotificationRuleRequest`

NewUpdateNotificationRuleRequest instantiates a new UpdateNotificationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNotificationRuleRequestWithDefaults

`func NewUpdateNotificationRuleRequestWithDefaults() *UpdateNotificationRuleRequest`

NewUpdateNotificationRuleRequestWithDefaults instantiates a new UpdateNotificationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNotificationRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNotificationRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNotificationRuleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetQueryString

`func (o *UpdateNotificationRuleRequest) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *UpdateNotificationRuleRequest) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *UpdateNotificationRuleRequest) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *UpdateNotificationRuleRequest) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetTeamsFilter

`func (o *UpdateNotificationRuleRequest) GetTeamsFilter() CreateNotificationRuleRequestTeamsFilter`

GetTeamsFilter returns the TeamsFilter field if non-nil, zero value otherwise.

### GetTeamsFilterOk

`func (o *UpdateNotificationRuleRequest) GetTeamsFilterOk() (*CreateNotificationRuleRequestTeamsFilter, bool)`

GetTeamsFilterOk returns a tuple with the TeamsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsFilter

`func (o *UpdateNotificationRuleRequest) SetTeamsFilter(v CreateNotificationRuleRequestTeamsFilter)`

SetTeamsFilter sets TeamsFilter field to given value.


### GetPrioritiesFilter

`func (o *UpdateNotificationRuleRequest) GetPrioritiesFilter() []float32`

GetPrioritiesFilter returns the PrioritiesFilter field if non-nil, zero value otherwise.

### GetPrioritiesFilterOk

`func (o *UpdateNotificationRuleRequest) GetPrioritiesFilterOk() (*[]float32, bool)`

GetPrioritiesFilterOk returns a tuple with the PrioritiesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioritiesFilter

`func (o *UpdateNotificationRuleRequest) SetPrioritiesFilter(v []float32)`

SetPrioritiesFilter sets PrioritiesFilter field to given value.


### GetTransitionTypesFilter

`func (o *UpdateNotificationRuleRequest) GetTransitionTypesFilter() []string`

GetTransitionTypesFilter returns the TransitionTypesFilter field if non-nil, zero value otherwise.

### GetTransitionTypesFilterOk

`func (o *UpdateNotificationRuleRequest) GetTransitionTypesFilterOk() (*[]string, bool)`

GetTransitionTypesFilterOk returns a tuple with the TransitionTypesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionTypesFilter

`func (o *UpdateNotificationRuleRequest) SetTransitionTypesFilter(v []string)`

SetTransitionTypesFilter sets TransitionTypesFilter field to given value.


### GetClusterIdsFilter

`func (o *UpdateNotificationRuleRequest) GetClusterIdsFilter() []string`

GetClusterIdsFilter returns the ClusterIdsFilter field if non-nil, zero value otherwise.

### GetClusterIdsFilterOk

`func (o *UpdateNotificationRuleRequest) GetClusterIdsFilterOk() (*[]string, bool)`

GetClusterIdsFilterOk returns a tuple with the ClusterIdsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdsFilter

`func (o *UpdateNotificationRuleRequest) SetClusterIdsFilter(v []string)`

SetClusterIdsFilter sets ClusterIdsFilter field to given value.

### HasClusterIdsFilter

`func (o *UpdateNotificationRuleRequest) HasClusterIdsFilter() bool`

HasClusterIdsFilter returns a boolean if a field has been set.

### GetOwner

`func (o *UpdateNotificationRuleRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateNotificationRuleRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateNotificationRuleRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetTags

`func (o *UpdateNotificationRuleRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateNotificationRuleRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateNotificationRuleRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateNotificationRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsActive

`func (o *UpdateNotificationRuleRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateNotificationRuleRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateNotificationRuleRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetTargets

`func (o *UpdateNotificationRuleRequest) GetTargets() []CreateNotificationRuleRequestTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *UpdateNotificationRuleRequest) GetTargetsOk() (*[]CreateNotificationRuleRequestTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *UpdateNotificationRuleRequest) SetTargets(v []CreateNotificationRuleRequestTargetsInner)`

SetTargets sets Targets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


