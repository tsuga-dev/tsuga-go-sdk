# ListNotificationRules200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the notification rule | 
**Name** | **string** | Display name of the notification rule | 
**TeamsFilter** | **[]string** | Team IDs that narrow down the teams that can receive notifications from this rule | 
**PrioritiesFilter** | **[]int32** | Priorities that narrow down the alerts that can trigger a notification | 
**TransitionTypesFilter** | **[]string** | Alert state transitions that can trigger a notification | 
**Owner** | **string** | Team ID that owns and manages the rule | 
**IsActive** | **bool** |  | 
**Tags** | Pointer to [**[]ListDashboards200ResponseDataInnerTagsInner**](ListDashboards200ResponseDataInnerTagsInner.md) | List of key/value tags applied to the resource | [optional] 
**Targets** | [**[]ListNotificationRules200ResponseDataInnerTargetsInner**](ListNotificationRules200ResponseDataInnerTargetsInner.md) | Notification targets that can receive notifications when the rule matches | 

## Methods

### NewListNotificationRules200ResponseDataInner

`func NewListNotificationRules200ResponseDataInner(id string, name string, teamsFilter []string, prioritiesFilter []int32, transitionTypesFilter []string, owner string, isActive bool, targets []ListNotificationRules200ResponseDataInnerTargetsInner, ) *ListNotificationRules200ResponseDataInner`

NewListNotificationRules200ResponseDataInner instantiates a new ListNotificationRules200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNotificationRules200ResponseDataInnerWithDefaults

`func NewListNotificationRules200ResponseDataInnerWithDefaults() *ListNotificationRules200ResponseDataInner`

NewListNotificationRules200ResponseDataInnerWithDefaults instantiates a new ListNotificationRules200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListNotificationRules200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListNotificationRules200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListNotificationRules200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListNotificationRules200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListNotificationRules200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListNotificationRules200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.


### GetTeamsFilter

`func (o *ListNotificationRules200ResponseDataInner) GetTeamsFilter() []string`

GetTeamsFilter returns the TeamsFilter field if non-nil, zero value otherwise.

### GetTeamsFilterOk

`func (o *ListNotificationRules200ResponseDataInner) GetTeamsFilterOk() (*[]string, bool)`

GetTeamsFilterOk returns a tuple with the TeamsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsFilter

`func (o *ListNotificationRules200ResponseDataInner) SetTeamsFilter(v []string)`

SetTeamsFilter sets TeamsFilter field to given value.


### GetPrioritiesFilter

`func (o *ListNotificationRules200ResponseDataInner) GetPrioritiesFilter() []int32`

GetPrioritiesFilter returns the PrioritiesFilter field if non-nil, zero value otherwise.

### GetPrioritiesFilterOk

`func (o *ListNotificationRules200ResponseDataInner) GetPrioritiesFilterOk() (*[]int32, bool)`

GetPrioritiesFilterOk returns a tuple with the PrioritiesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioritiesFilter

`func (o *ListNotificationRules200ResponseDataInner) SetPrioritiesFilter(v []int32)`

SetPrioritiesFilter sets PrioritiesFilter field to given value.


### GetTransitionTypesFilter

`func (o *ListNotificationRules200ResponseDataInner) GetTransitionTypesFilter() []string`

GetTransitionTypesFilter returns the TransitionTypesFilter field if non-nil, zero value otherwise.

### GetTransitionTypesFilterOk

`func (o *ListNotificationRules200ResponseDataInner) GetTransitionTypesFilterOk() (*[]string, bool)`

GetTransitionTypesFilterOk returns a tuple with the TransitionTypesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionTypesFilter

`func (o *ListNotificationRules200ResponseDataInner) SetTransitionTypesFilter(v []string)`

SetTransitionTypesFilter sets TransitionTypesFilter field to given value.


### GetOwner

`func (o *ListNotificationRules200ResponseDataInner) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListNotificationRules200ResponseDataInner) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListNotificationRules200ResponseDataInner) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetIsActive

`func (o *ListNotificationRules200ResponseDataInner) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ListNotificationRules200ResponseDataInner) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ListNotificationRules200ResponseDataInner) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetTags

`func (o *ListNotificationRules200ResponseDataInner) GetTags() []ListDashboards200ResponseDataInnerTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListNotificationRules200ResponseDataInner) GetTagsOk() (*[]ListDashboards200ResponseDataInnerTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListNotificationRules200ResponseDataInner) SetTags(v []ListDashboards200ResponseDataInnerTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListNotificationRules200ResponseDataInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargets

`func (o *ListNotificationRules200ResponseDataInner) GetTargets() []ListNotificationRules200ResponseDataInnerTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ListNotificationRules200ResponseDataInner) GetTargetsOk() (*[]ListNotificationRules200ResponseDataInnerTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ListNotificationRules200ResponseDataInner) SetTargets(v []ListNotificationRules200ResponseDataInnerTargetsInner)`

SetTargets sets Targets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


