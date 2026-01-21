# Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the notification rule | 
**Name** | **string** | Display name of the notification rule | 
**TeamsFilter** | [**RuleTeamsFilter**](RuleTeamsFilter.md) |  | 
**PrioritiesFilter** | **[]float32** | Priorities that narrow down the alerts that can trigger a notification | 
**TransitionTypesFilter** | **[]string** | Alert state transitions that can trigger a notification | 
**Owner** | **string** | Team ID that owns and manages the rule | 
**IsActive** | **bool** |  | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List of key/value tags applied to the resource | [optional] 
**Targets** | [**[]RuleTargetsInner**](RuleTargetsInner.md) | Notification targets that can receive notifications when the rule matches | 

## Methods

### NewRule

`func NewRule(id string, name string, teamsFilter RuleTeamsFilter, prioritiesFilter []float32, transitionTypesFilter []string, owner string, isActive bool, targets []RuleTargetsInner, ) *Rule`

NewRule instantiates a new Rule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleWithDefaults

`func NewRuleWithDefaults() *Rule`

NewRuleWithDefaults instantiates a new Rule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Rule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Rule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Rule) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Rule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Rule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Rule) SetName(v string)`

SetName sets Name field to given value.


### GetTeamsFilter

`func (o *Rule) GetTeamsFilter() RuleTeamsFilter`

GetTeamsFilter returns the TeamsFilter field if non-nil, zero value otherwise.

### GetTeamsFilterOk

`func (o *Rule) GetTeamsFilterOk() (*RuleTeamsFilter, bool)`

GetTeamsFilterOk returns a tuple with the TeamsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsFilter

`func (o *Rule) SetTeamsFilter(v RuleTeamsFilter)`

SetTeamsFilter sets TeamsFilter field to given value.


### GetPrioritiesFilter

`func (o *Rule) GetPrioritiesFilter() []float32`

GetPrioritiesFilter returns the PrioritiesFilter field if non-nil, zero value otherwise.

### GetPrioritiesFilterOk

`func (o *Rule) GetPrioritiesFilterOk() (*[]float32, bool)`

GetPrioritiesFilterOk returns a tuple with the PrioritiesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioritiesFilter

`func (o *Rule) SetPrioritiesFilter(v []float32)`

SetPrioritiesFilter sets PrioritiesFilter field to given value.


### GetTransitionTypesFilter

`func (o *Rule) GetTransitionTypesFilter() []string`

GetTransitionTypesFilter returns the TransitionTypesFilter field if non-nil, zero value otherwise.

### GetTransitionTypesFilterOk

`func (o *Rule) GetTransitionTypesFilterOk() (*[]string, bool)`

GetTransitionTypesFilterOk returns a tuple with the TransitionTypesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionTypesFilter

`func (o *Rule) SetTransitionTypesFilter(v []string)`

SetTransitionTypesFilter sets TransitionTypesFilter field to given value.


### GetOwner

`func (o *Rule) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Rule) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Rule) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetIsActive

`func (o *Rule) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Rule) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Rule) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetTags

`func (o *Rule) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Rule) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Rule) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Rule) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargets

`func (o *Rule) GetTargets() []RuleTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *Rule) GetTargetsOk() (*[]RuleTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *Rule) SetTargets(v []RuleTargetsInner)`

SetTargets sets Targets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


