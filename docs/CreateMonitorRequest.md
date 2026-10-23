# CreateMonitorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the monitor and alert source. | 
**Message** | Pointer to **NullableString** | Message included in notifications triggered by this monitor. Optional on create. | [optional] 
**Tags** | Pointer to [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 
**Configuration** | [**UpdateMonitorRequestConfiguration**](UpdateMonitorRequestConfiguration.md) |  | 
**Priority** | **int32** | Monitor priority from 1 through 5, where 1 is highest priority. | 
**Owner** | **string** | Team ID that owns and manages the monitor. | 
**DashboardId** | Pointer to **NullableString** | ID of the dashboard to link to this monitor for context when investigating alerts. On create: omit or send &#x60;null&#x60; to leave the monitor unlinked. On update: omit to skip updating this field, send &#x60;null&#x60; to clear an existing link. | [optional] 
**FolderId** | Pointer to **NullableString** | Folder to organize this monitor in. Omit on create or send null to leave it unfiled; omit on update to keep its folder, or send null to remove it. A non-null ID must name a readable monitor folder. Assignment does not change ownership, data access, or evaluation state. | [optional] 
**PauseUntil** | Pointer to **NullableTime** | ISO timestamp until which the monitor is snoozed. Must be in the future and at most one month from now. On create: omit or send &#x60;null&#x60; to leave the monitor active. On update: omit to keep the current snooze, send &#x60;null&#x60; to clear it. | [optional] 
**Permissions** | **string** | &#x60;all&#x60; allows the resource to query all permitted telemetry, &#x60;owning-team-and-public&#x60; limits it to the owning team plus public data, and &#x60;owning-team-only&#x60; limits it to the owning team. | 
**ClusterIds** | Pointer to **[]string** | Cluster IDs where the monitor should be deployed. Empty or omitted means all eligible clusters for cluster-deployed monitor types. | [optional] 

## Methods

### NewCreateMonitorRequest

`func NewCreateMonitorRequest(name string, configuration UpdateMonitorRequestConfiguration, priority int32, owner string, permissions string, ) *CreateMonitorRequest`

NewCreateMonitorRequest instantiates a new CreateMonitorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMonitorRequestWithDefaults

`func NewCreateMonitorRequestWithDefaults() *CreateMonitorRequest`

NewCreateMonitorRequestWithDefaults instantiates a new CreateMonitorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateMonitorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMonitorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMonitorRequest) SetName(v string)`

SetName sets Name field to given value.


### GetMessage

`func (o *CreateMonitorRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateMonitorRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateMonitorRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateMonitorRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CreateMonitorRequest) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CreateMonitorRequest) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetTags

`func (o *CreateMonitorRequest) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateMonitorRequest) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateMonitorRequest) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateMonitorRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetConfiguration

`func (o *CreateMonitorRequest) GetConfiguration() UpdateMonitorRequestConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreateMonitorRequest) GetConfigurationOk() (*UpdateMonitorRequestConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreateMonitorRequest) SetConfiguration(v UpdateMonitorRequestConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetPriority

`func (o *CreateMonitorRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateMonitorRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateMonitorRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetOwner

`func (o *CreateMonitorRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateMonitorRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateMonitorRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetDashboardId

`func (o *CreateMonitorRequest) GetDashboardId() string`

GetDashboardId returns the DashboardId field if non-nil, zero value otherwise.

### GetDashboardIdOk

`func (o *CreateMonitorRequest) GetDashboardIdOk() (*string, bool)`

GetDashboardIdOk returns a tuple with the DashboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardId

`func (o *CreateMonitorRequest) SetDashboardId(v string)`

SetDashboardId sets DashboardId field to given value.

### HasDashboardId

`func (o *CreateMonitorRequest) HasDashboardId() bool`

HasDashboardId returns a boolean if a field has been set.

### SetDashboardIdNil

`func (o *CreateMonitorRequest) SetDashboardIdNil(b bool)`

 SetDashboardIdNil sets the value for DashboardId to be an explicit nil

### UnsetDashboardId
`func (o *CreateMonitorRequest) UnsetDashboardId()`

UnsetDashboardId ensures that no value is present for DashboardId, not even an explicit nil
### GetFolderId

`func (o *CreateMonitorRequest) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *CreateMonitorRequest) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *CreateMonitorRequest) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *CreateMonitorRequest) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *CreateMonitorRequest) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *CreateMonitorRequest) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetPauseUntil

`func (o *CreateMonitorRequest) GetPauseUntil() time.Time`

GetPauseUntil returns the PauseUntil field if non-nil, zero value otherwise.

### GetPauseUntilOk

`func (o *CreateMonitorRequest) GetPauseUntilOk() (*time.Time, bool)`

GetPauseUntilOk returns a tuple with the PauseUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseUntil

`func (o *CreateMonitorRequest) SetPauseUntil(v time.Time)`

SetPauseUntil sets PauseUntil field to given value.

### HasPauseUntil

`func (o *CreateMonitorRequest) HasPauseUntil() bool`

HasPauseUntil returns a boolean if a field has been set.

### SetPauseUntilNil

`func (o *CreateMonitorRequest) SetPauseUntilNil(b bool)`

 SetPauseUntilNil sets the value for PauseUntil to be an explicit nil

### UnsetPauseUntil
`func (o *CreateMonitorRequest) UnsetPauseUntil()`

UnsetPauseUntil ensures that no value is present for PauseUntil, not even an explicit nil
### GetPermissions

`func (o *CreateMonitorRequest) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateMonitorRequest) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateMonitorRequest) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.


### GetClusterIds

`func (o *CreateMonitorRequest) GetClusterIds() []string`

GetClusterIds returns the ClusterIds field if non-nil, zero value otherwise.

### GetClusterIdsOk

`func (o *CreateMonitorRequest) GetClusterIdsOk() (*[]string, bool)`

GetClusterIdsOk returns a tuple with the ClusterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIds

`func (o *CreateMonitorRequest) SetClusterIds(v []string)`

SetClusterIds sets ClusterIds field to given value.

### HasClusterIds

`func (o *CreateMonitorRequest) HasClusterIds() bool`

HasClusterIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


