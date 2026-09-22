# UpdateSloRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the SLO. | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 
**Configuration** | [**UpdateSloRequestConfiguration**](UpdateSloRequestConfiguration.md) |  | 
**Target** | **float32** | Target percentage (0 &lt; target &lt; 100, e.g. 99.9) | 
**TimeframeDays** | **int32** | Rolling SLO evaluation window in days. Set by the caller on create or update and returned on read and list responses. Allowed values are 7, 30, and 90. | 
**Owner** | **string** | Team ID that owns the SLO | 
**Permissions** | **string** | &#x60;all&#x60; allows the resource to query all permitted telemetry, &#x60;owning-team-and-public&#x60; limits it to the owning team plus public data, and &#x60;owning-team-only&#x60; limits it to the owning team. | 
**ClusterIds** | Pointer to **[]string** | Cluster IDs this SLO is evaluated on. Omit or send an empty array on input to run on all clusters; responses use an empty array to mean all clusters. Non-empty input values must reference existing clusters. | [optional] 
**Alerts** | [**[]UpdateSloRequestAlertsInner**](UpdateSloRequestAlertsInner.md) | Alerts attached to this SLO. The SLO owns its alerts: the set is reconciled to exactly this list — existing alerts are matched by id, alerts without an id are created, and existing alerts absent from the list are deleted. Send an empty array to clear all alerts. | 

## Methods

### NewUpdateSloRequest

`func NewUpdateSloRequest(name string, configuration UpdateSloRequestConfiguration, target float32, timeframeDays int32, owner string, permissions string, alerts []UpdateSloRequestAlertsInner, ) *UpdateSloRequest`

NewUpdateSloRequest instantiates a new UpdateSloRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSloRequestWithDefaults

`func NewUpdateSloRequestWithDefaults() *UpdateSloRequest`

NewUpdateSloRequestWithDefaults instantiates a new UpdateSloRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateSloRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSloRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSloRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateSloRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSloRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSloRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSloRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateSloRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateSloRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTags

`func (o *UpdateSloRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateSloRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateSloRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateSloRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetConfiguration

`func (o *UpdateSloRequest) GetConfiguration() UpdateSloRequestConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *UpdateSloRequest) GetConfigurationOk() (*UpdateSloRequestConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *UpdateSloRequest) SetConfiguration(v UpdateSloRequestConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetTarget

`func (o *UpdateSloRequest) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *UpdateSloRequest) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *UpdateSloRequest) SetTarget(v float32)`

SetTarget sets Target field to given value.


### GetTimeframeDays

`func (o *UpdateSloRequest) GetTimeframeDays() int32`

GetTimeframeDays returns the TimeframeDays field if non-nil, zero value otherwise.

### GetTimeframeDaysOk

`func (o *UpdateSloRequest) GetTimeframeDaysOk() (*int32, bool)`

GetTimeframeDaysOk returns a tuple with the TimeframeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframeDays

`func (o *UpdateSloRequest) SetTimeframeDays(v int32)`

SetTimeframeDays sets TimeframeDays field to given value.


### GetOwner

`func (o *UpdateSloRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateSloRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateSloRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetPermissions

`func (o *UpdateSloRequest) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateSloRequest) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateSloRequest) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.


### GetClusterIds

`func (o *UpdateSloRequest) GetClusterIds() []*string`

GetClusterIds returns the ClusterIds field if non-nil, zero value otherwise.

### GetClusterIdsOk

`func (o *UpdateSloRequest) GetClusterIdsOk() (*[]*string, bool)`

GetClusterIdsOk returns a tuple with the ClusterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIds

`func (o *UpdateSloRequest) SetClusterIds(v []*string)`

SetClusterIds sets ClusterIds field to given value.

### HasClusterIds

`func (o *UpdateSloRequest) HasClusterIds() bool`

HasClusterIds returns a boolean if a field has been set.

### GetAlerts

`func (o *UpdateSloRequest) GetAlerts() []UpdateSloRequestAlertsInner`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *UpdateSloRequest) GetAlertsOk() (*[]UpdateSloRequestAlertsInner, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *UpdateSloRequest) SetAlerts(v []UpdateSloRequestAlertsInner)`

SetAlerts sets Alerts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


