# UpdateAdaptiveNotificationRuleRequest

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

### NewUpdateAdaptiveNotificationRuleRequest

`func NewUpdateAdaptiveNotificationRuleRequest(name string, owner string, isActive bool, fallbackTarget CreateNotificationRuleRequestTargetsInner, ) *UpdateAdaptiveNotificationRuleRequest`

NewUpdateAdaptiveNotificationRuleRequest instantiates a new UpdateAdaptiveNotificationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAdaptiveNotificationRuleRequestWithDefaults

`func NewUpdateAdaptiveNotificationRuleRequestWithDefaults() *UpdateAdaptiveNotificationRuleRequest`

NewUpdateAdaptiveNotificationRuleRequestWithDefaults instantiates a new UpdateAdaptiveNotificationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAdaptiveNotificationRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAdaptiveNotificationRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAdaptiveNotificationRuleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *UpdateAdaptiveNotificationRuleRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateAdaptiveNotificationRuleRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateAdaptiveNotificationRuleRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetTags

`func (o *UpdateAdaptiveNotificationRuleRequest) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateAdaptiveNotificationRuleRequest) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateAdaptiveNotificationRuleRequest) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateAdaptiveNotificationRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsActive

`func (o *UpdateAdaptiveNotificationRuleRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateAdaptiveNotificationRuleRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateAdaptiveNotificationRuleRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetShouldCreateInvestigation

`func (o *UpdateAdaptiveNotificationRuleRequest) GetShouldCreateInvestigation() bool`

GetShouldCreateInvestigation returns the ShouldCreateInvestigation field if non-nil, zero value otherwise.

### GetShouldCreateInvestigationOk

`func (o *UpdateAdaptiveNotificationRuleRequest) GetShouldCreateInvestigationOk() (*bool, bool)`

GetShouldCreateInvestigationOk returns a tuple with the ShouldCreateInvestigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCreateInvestigation

`func (o *UpdateAdaptiveNotificationRuleRequest) SetShouldCreateInvestigation(v bool)`

SetShouldCreateInvestigation sets ShouldCreateInvestigation field to given value.

### HasShouldCreateInvestigation

`func (o *UpdateAdaptiveNotificationRuleRequest) HasShouldCreateInvestigation() bool`

HasShouldCreateInvestigation returns a boolean if a field has been set.

### GetMonitorIds

`func (o *UpdateAdaptiveNotificationRuleRequest) GetMonitorIds() []string`

GetMonitorIds returns the MonitorIds field if non-nil, zero value otherwise.

### GetMonitorIdsOk

`func (o *UpdateAdaptiveNotificationRuleRequest) GetMonitorIdsOk() (*[]string, bool)`

GetMonitorIdsOk returns a tuple with the MonitorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorIds

`func (o *UpdateAdaptiveNotificationRuleRequest) SetMonitorIds(v []string)`

SetMonitorIds sets MonitorIds field to given value.

### HasMonitorIds

`func (o *UpdateAdaptiveNotificationRuleRequest) HasMonitorIds() bool`

HasMonitorIds returns a boolean if a field has been set.

### GetSloIds

`func (o *UpdateAdaptiveNotificationRuleRequest) GetSloIds() []string`

GetSloIds returns the SloIds field if non-nil, zero value otherwise.

### GetSloIdsOk

`func (o *UpdateAdaptiveNotificationRuleRequest) GetSloIdsOk() (*[]string, bool)`

GetSloIdsOk returns a tuple with the SloIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloIds

`func (o *UpdateAdaptiveNotificationRuleRequest) SetSloIds(v []string)`

SetSloIds sets SloIds field to given value.

### HasSloIds

`func (o *UpdateAdaptiveNotificationRuleRequest) HasSloIds() bool`

HasSloIds returns a boolean if a field has been set.

### GetFallbackTarget

`func (o *UpdateAdaptiveNotificationRuleRequest) GetFallbackTarget() CreateNotificationRuleRequestTargetsInner`

GetFallbackTarget returns the FallbackTarget field if non-nil, zero value otherwise.

### GetFallbackTargetOk

`func (o *UpdateAdaptiveNotificationRuleRequest) GetFallbackTargetOk() (*CreateNotificationRuleRequestTargetsInner, bool)`

GetFallbackTargetOk returns a tuple with the FallbackTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackTarget

`func (o *UpdateAdaptiveNotificationRuleRequest) SetFallbackTarget(v CreateNotificationRuleRequestTargetsInner)`

SetFallbackTarget sets FallbackTarget field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


