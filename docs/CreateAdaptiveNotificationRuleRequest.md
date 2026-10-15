# CreateAdaptiveNotificationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the adaptive notification rule. | 
**Owner** | **string** | Team ID that owns and manages the rule | 
**Tags** | Pointer to [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 
**IsActive** | **bool** | Set to true for the rule to send notifications when one of its monitors fires. | 
**ShouldCreateInvestigation** | Pointer to **bool** | Set to true to record an investigation for each page, holding the sampled traces and the call graph the destination was resolved from, and to link it from the notification body. Defaults to false. | [optional] 
**MonitorIds** | Pointer to **[]string** | Monitors this rule reacts to. Only trace monitors and metric monitors over OTel span metrics are eligible. On update, omitting this field keeps the current monitors; send an empty array to remove them. Send at least one monitor or SLO. | [optional] 
**SloIds** | Pointer to **[]string** | SLOs this rule reacts to. Only SLOs over traces or OTel span metrics are eligible. On update, omitting this field keeps the current SLOs; send an empty array to remove them. Send at least one monitor or SLO. | [optional] 
**FallbackTarget** | [**CreateNotificationRuleRequestTargetsInner**](CreateNotificationRuleRequestTargetsInner.md) |  | 

## Methods

### NewCreateAdaptiveNotificationRuleRequest

`func NewCreateAdaptiveNotificationRuleRequest(name string, owner string, isActive bool, fallbackTarget CreateNotificationRuleRequestTargetsInner, ) *CreateAdaptiveNotificationRuleRequest`

NewCreateAdaptiveNotificationRuleRequest instantiates a new CreateAdaptiveNotificationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAdaptiveNotificationRuleRequestWithDefaults

`func NewCreateAdaptiveNotificationRuleRequestWithDefaults() *CreateAdaptiveNotificationRuleRequest`

NewCreateAdaptiveNotificationRuleRequestWithDefaults instantiates a new CreateAdaptiveNotificationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAdaptiveNotificationRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAdaptiveNotificationRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAdaptiveNotificationRuleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *CreateAdaptiveNotificationRuleRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateAdaptiveNotificationRuleRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateAdaptiveNotificationRuleRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetTags

`func (o *CreateAdaptiveNotificationRuleRequest) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateAdaptiveNotificationRuleRequest) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateAdaptiveNotificationRuleRequest) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateAdaptiveNotificationRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsActive

`func (o *CreateAdaptiveNotificationRuleRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateAdaptiveNotificationRuleRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateAdaptiveNotificationRuleRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetShouldCreateInvestigation

`func (o *CreateAdaptiveNotificationRuleRequest) GetShouldCreateInvestigation() bool`

GetShouldCreateInvestigation returns the ShouldCreateInvestigation field if non-nil, zero value otherwise.

### GetShouldCreateInvestigationOk

`func (o *CreateAdaptiveNotificationRuleRequest) GetShouldCreateInvestigationOk() (*bool, bool)`

GetShouldCreateInvestigationOk returns a tuple with the ShouldCreateInvestigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCreateInvestigation

`func (o *CreateAdaptiveNotificationRuleRequest) SetShouldCreateInvestigation(v bool)`

SetShouldCreateInvestigation sets ShouldCreateInvestigation field to given value.

### HasShouldCreateInvestigation

`func (o *CreateAdaptiveNotificationRuleRequest) HasShouldCreateInvestigation() bool`

HasShouldCreateInvestigation returns a boolean if a field has been set.

### GetMonitorIds

`func (o *CreateAdaptiveNotificationRuleRequest) GetMonitorIds() []*string`

GetMonitorIds returns the MonitorIds field if non-nil, zero value otherwise.

### GetMonitorIdsOk

`func (o *CreateAdaptiveNotificationRuleRequest) GetMonitorIdsOk() (*[]*string, bool)`

GetMonitorIdsOk returns a tuple with the MonitorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorIds

`func (o *CreateAdaptiveNotificationRuleRequest) SetMonitorIds(v []*string)`

SetMonitorIds sets MonitorIds field to given value.

### HasMonitorIds

`func (o *CreateAdaptiveNotificationRuleRequest) HasMonitorIds() bool`

HasMonitorIds returns a boolean if a field has been set.

### GetSloIds

`func (o *CreateAdaptiveNotificationRuleRequest) GetSloIds() []*string`

GetSloIds returns the SloIds field if non-nil, zero value otherwise.

### GetSloIdsOk

`func (o *CreateAdaptiveNotificationRuleRequest) GetSloIdsOk() (*[]*string, bool)`

GetSloIdsOk returns a tuple with the SloIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloIds

`func (o *CreateAdaptiveNotificationRuleRequest) SetSloIds(v []*string)`

SetSloIds sets SloIds field to given value.

### HasSloIds

`func (o *CreateAdaptiveNotificationRuleRequest) HasSloIds() bool`

HasSloIds returns a boolean if a field has been set.

### GetFallbackTarget

`func (o *CreateAdaptiveNotificationRuleRequest) GetFallbackTarget() CreateNotificationRuleRequestTargetsInner`

GetFallbackTarget returns the FallbackTarget field if non-nil, zero value otherwise.

### GetFallbackTargetOk

`func (o *CreateAdaptiveNotificationRuleRequest) GetFallbackTargetOk() (*CreateNotificationRuleRequestTargetsInner, bool)`

GetFallbackTargetOk returns a tuple with the FallbackTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackTarget

`func (o *CreateAdaptiveNotificationRuleRequest) SetFallbackTarget(v CreateNotificationRuleRequestTargetsInner)`

SetFallbackTarget sets FallbackTarget field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


