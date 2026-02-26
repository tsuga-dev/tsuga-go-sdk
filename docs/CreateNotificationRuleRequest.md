# CreateNotificationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the notification rule | 
**TeamsFilter** | [**CreateNotificationRuleRequestTeamsFilter**](CreateNotificationRuleRequestTeamsFilter.md) |  | 
**PrioritiesFilter** | **[]float32** | Priorities that narrow down the alerts that can trigger a notification | 
**TransitionTypesFilter** | **[]string** | Alert state transitions that can trigger a notification | 
**ClusterIdsFilter** | Pointer to **[]string** | Cluster IDs that can trigger a notification | [optional] 
**Owner** | **string** | Team ID that owns and manages the rule | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List of key/value tags applied to the resource | [optional] 
**IsActive** | **bool** |  | 
**Targets** | [**[]CreateNotificationRuleRequestTargetsInner**](CreateNotificationRuleRequestTargetsInner.md) | Notification targets that can receive notifications when the rule matches | 

## Methods

### NewCreateNotificationRuleRequest

`func NewCreateNotificationRuleRequest(name string, teamsFilter CreateNotificationRuleRequestTeamsFilter, prioritiesFilter []float32, transitionTypesFilter []string, owner string, isActive bool, targets []CreateNotificationRuleRequestTargetsInner, ) *CreateNotificationRuleRequest`

NewCreateNotificationRuleRequest instantiates a new CreateNotificationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationRuleRequestWithDefaults

`func NewCreateNotificationRuleRequestWithDefaults() *CreateNotificationRuleRequest`

NewCreateNotificationRuleRequestWithDefaults instantiates a new CreateNotificationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNotificationRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNotificationRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNotificationRuleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTeamsFilter

`func (o *CreateNotificationRuleRequest) GetTeamsFilter() CreateNotificationRuleRequestTeamsFilter`

GetTeamsFilter returns the TeamsFilter field if non-nil, zero value otherwise.

### GetTeamsFilterOk

`func (o *CreateNotificationRuleRequest) GetTeamsFilterOk() (*CreateNotificationRuleRequestTeamsFilter, bool)`

GetTeamsFilterOk returns a tuple with the TeamsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsFilter

`func (o *CreateNotificationRuleRequest) SetTeamsFilter(v CreateNotificationRuleRequestTeamsFilter)`

SetTeamsFilter sets TeamsFilter field to given value.


### GetPrioritiesFilter

`func (o *CreateNotificationRuleRequest) GetPrioritiesFilter() []float32`

GetPrioritiesFilter returns the PrioritiesFilter field if non-nil, zero value otherwise.

### GetPrioritiesFilterOk

`func (o *CreateNotificationRuleRequest) GetPrioritiesFilterOk() (*[]float32, bool)`

GetPrioritiesFilterOk returns a tuple with the PrioritiesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioritiesFilter

`func (o *CreateNotificationRuleRequest) SetPrioritiesFilter(v []float32)`

SetPrioritiesFilter sets PrioritiesFilter field to given value.


### GetTransitionTypesFilter

`func (o *CreateNotificationRuleRequest) GetTransitionTypesFilter() []string`

GetTransitionTypesFilter returns the TransitionTypesFilter field if non-nil, zero value otherwise.

### GetTransitionTypesFilterOk

`func (o *CreateNotificationRuleRequest) GetTransitionTypesFilterOk() (*[]string, bool)`

GetTransitionTypesFilterOk returns a tuple with the TransitionTypesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionTypesFilter

`func (o *CreateNotificationRuleRequest) SetTransitionTypesFilter(v []string)`

SetTransitionTypesFilter sets TransitionTypesFilter field to given value.


### GetClusterIdsFilter

`func (o *CreateNotificationRuleRequest) GetClusterIdsFilter() []string`

GetClusterIdsFilter returns the ClusterIdsFilter field if non-nil, zero value otherwise.

### GetClusterIdsFilterOk

`func (o *CreateNotificationRuleRequest) GetClusterIdsFilterOk() (*[]string, bool)`

GetClusterIdsFilterOk returns a tuple with the ClusterIdsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdsFilter

`func (o *CreateNotificationRuleRequest) SetClusterIdsFilter(v []string)`

SetClusterIdsFilter sets ClusterIdsFilter field to given value.

### HasClusterIdsFilter

`func (o *CreateNotificationRuleRequest) HasClusterIdsFilter() bool`

HasClusterIdsFilter returns a boolean if a field has been set.

### GetOwner

`func (o *CreateNotificationRuleRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateNotificationRuleRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateNotificationRuleRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetTags

`func (o *CreateNotificationRuleRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateNotificationRuleRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateNotificationRuleRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateNotificationRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsActive

`func (o *CreateNotificationRuleRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateNotificationRuleRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateNotificationRuleRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetTargets

`func (o *CreateNotificationRuleRequest) GetTargets() []CreateNotificationRuleRequestTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *CreateNotificationRuleRequest) GetTargetsOk() (*[]CreateNotificationRuleRequestTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *CreateNotificationRuleRequest) SetTargets(v []CreateNotificationRuleRequestTargetsInner)`

SetTargets sets Targets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


