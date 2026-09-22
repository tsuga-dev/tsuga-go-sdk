# Slo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Tsuga-generated SLO ID assigned when the SLO is created. | 
**Name** | **string** | Display name of the SLO, set by the caller. | 
**Description** | Pointer to **string** | Free-form, caller-provided description of what the SLO measures. Absent when not set. | [optional] 
**Tags** | Pointer to [**[]Tag1**](Tag1.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | [optional] 
**Configuration** | [**SloConfiguration**](SloConfiguration.md) |  | 
**Target** | **float32** | Target percentage (0 &lt; target &lt; 100, e.g. 99.9) | 
**TimeframeDays** | **int32** | Rolling SLO evaluation window in days. Set by the caller on create or update and returned on read and list responses. Allowed values are 7, 30, and 90. | 
**Owner** | **string** | Team ID that owns the SLO | 
**Permissions** | **string** | &#x60;all&#x60; allows the resource to query all permitted telemetry, &#x60;owning-team-and-public&#x60; limits it to the owning team plus public data, and &#x60;owning-team-only&#x60; limits it to the owning team. | 
**ClusterIds** | **[]string** | Cluster IDs this SLO is evaluated on. Omit or send an empty array on input to run on all clusters; responses use an empty array to mean all clusters. Non-empty input values must reference existing clusters. | 
**Alerts** | [**[]SloAlertsInner**](SloAlertsInner.md) | Alerts attached to this SLO | 

## Methods

### NewSlo

`func NewSlo(id string, name string, configuration SloConfiguration, target float32, timeframeDays int32, owner string, permissions string, clusterIds []string, alerts []SloAlertsInner, ) *Slo`

NewSlo instantiates a new Slo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloWithDefaults

`func NewSloWithDefaults() *Slo`

NewSloWithDefaults instantiates a new Slo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Slo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Slo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Slo) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Slo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Slo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Slo) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Slo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Slo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Slo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Slo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *Slo) GetTags() []Tag1`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Slo) GetTagsOk() (*[]Tag1, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Slo) SetTags(v []Tag1)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Slo) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetConfiguration

`func (o *Slo) GetConfiguration() SloConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Slo) GetConfigurationOk() (*SloConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Slo) SetConfiguration(v SloConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetTarget

`func (o *Slo) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Slo) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Slo) SetTarget(v float32)`

SetTarget sets Target field to given value.


### GetTimeframeDays

`func (o *Slo) GetTimeframeDays() int32`

GetTimeframeDays returns the TimeframeDays field if non-nil, zero value otherwise.

### GetTimeframeDaysOk

`func (o *Slo) GetTimeframeDaysOk() (*int32, bool)`

GetTimeframeDaysOk returns a tuple with the TimeframeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframeDays

`func (o *Slo) SetTimeframeDays(v int32)`

SetTimeframeDays sets TimeframeDays field to given value.


### GetOwner

`func (o *Slo) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Slo) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Slo) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetPermissions

`func (o *Slo) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Slo) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Slo) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.


### GetClusterIds

`func (o *Slo) GetClusterIds() []string`

GetClusterIds returns the ClusterIds field if non-nil, zero value otherwise.

### GetClusterIdsOk

`func (o *Slo) GetClusterIdsOk() (*[]string, bool)`

GetClusterIdsOk returns a tuple with the ClusterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIds

`func (o *Slo) SetClusterIds(v []string)`

SetClusterIds sets ClusterIds field to given value.


### GetAlerts

`func (o *Slo) GetAlerts() []SloAlertsInner`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *Slo) GetAlertsOk() (*[]SloAlertsInner, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *Slo) SetAlerts(v []SloAlertsInner)`

SetAlerts sets Alerts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


