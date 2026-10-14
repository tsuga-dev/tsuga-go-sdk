# UpdateMonitorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the monitor and alert source. | 
**Message** | Pointer to **NullableString** | Message included in notifications triggered by this monitor. Optional on create. | [optional] 
**Tags** | Pointer to [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 
**Configuration** | [**UpdateMonitorRequestConfiguration**](UpdateMonitorRequestConfiguration.md) |  | 
**Priority** | **float32** | Monitor priority from 1 through 5, where 1 is highest priority. | 
**Owner** | **string** | Team ID that owns and manages the monitor. | 
**DashboardId** | Pointer to **NullableString** | Folder filter for dashboard query results. Optional; pass a folder ID to return dashboards in that folder, or null to return only dashboards that are not in any folder. | [optional] 
**FolderId** | Pointer to **NullableString** | Folder filter for dashboard query results. Optional; pass a folder ID to return dashboards in that folder, or null to return only dashboards that are not in any folder. | [optional] 
**PauseUntil** | Pointer to **NullableTime** | ISO timestamp until which the monitor is snoozed. Must be in the future and at most one month from now. On create: omit or send &#x60;null&#x60; to leave the monitor active. On update: omit to keep the current snooze, send &#x60;null&#x60; to clear it. | [optional] 
**Permissions** | **string** | &#x60;all&#x60; allows the resource to query all permitted telemetry, &#x60;owning-team-and-public&#x60; limits it to the owning team plus public data, and &#x60;owning-team-only&#x60; limits it to the owning team. | 
**ClusterIds** | Pointer to **[]string** | Cluster IDs where the monitor should be deployed. Empty or omitted means all eligible clusters for cluster-deployed monitor types. | [optional] 

## Methods

### NewUpdateMonitorRequest

`func NewUpdateMonitorRequest(name string, configuration UpdateMonitorRequestConfiguration, priority float32, owner string, permissions string, ) *UpdateMonitorRequest`

NewUpdateMonitorRequest instantiates a new UpdateMonitorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMonitorRequestWithDefaults

`func NewUpdateMonitorRequestWithDefaults() *UpdateMonitorRequest`

NewUpdateMonitorRequestWithDefaults instantiates a new UpdateMonitorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateMonitorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMonitorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMonitorRequest) SetName(v string)`

SetName sets Name field to given value.


### GetMessage

`func (o *UpdateMonitorRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateMonitorRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateMonitorRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdateMonitorRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *UpdateMonitorRequest) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *UpdateMonitorRequest) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetTags

`func (o *UpdateMonitorRequest) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateMonitorRequest) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateMonitorRequest) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateMonitorRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetConfiguration

`func (o *UpdateMonitorRequest) GetConfiguration() UpdateMonitorRequestConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *UpdateMonitorRequest) GetConfigurationOk() (*UpdateMonitorRequestConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *UpdateMonitorRequest) SetConfiguration(v UpdateMonitorRequestConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetPriority

`func (o *UpdateMonitorRequest) GetPriority() float32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpdateMonitorRequest) GetPriorityOk() (*float32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpdateMonitorRequest) SetPriority(v float32)`

SetPriority sets Priority field to given value.


### GetOwner

`func (o *UpdateMonitorRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateMonitorRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateMonitorRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetDashboardId

`func (o *UpdateMonitorRequest) GetDashboardId() string`

GetDashboardId returns the DashboardId field if non-nil, zero value otherwise.

### GetDashboardIdOk

`func (o *UpdateMonitorRequest) GetDashboardIdOk() (*string, bool)`

GetDashboardIdOk returns a tuple with the DashboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardId

`func (o *UpdateMonitorRequest) SetDashboardId(v string)`

SetDashboardId sets DashboardId field to given value.

### HasDashboardId

`func (o *UpdateMonitorRequest) HasDashboardId() bool`

HasDashboardId returns a boolean if a field has been set.

### SetDashboardIdNil

`func (o *UpdateMonitorRequest) SetDashboardIdNil(b bool)`

 SetDashboardIdNil sets the value for DashboardId to be an explicit nil

### UnsetDashboardId
`func (o *UpdateMonitorRequest) UnsetDashboardId()`

UnsetDashboardId ensures that no value is present for DashboardId, not even an explicit nil
### GetFolderId

`func (o *UpdateMonitorRequest) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *UpdateMonitorRequest) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *UpdateMonitorRequest) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *UpdateMonitorRequest) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *UpdateMonitorRequest) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *UpdateMonitorRequest) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetPauseUntil

`func (o *UpdateMonitorRequest) GetPauseUntil() time.Time`

GetPauseUntil returns the PauseUntil field if non-nil, zero value otherwise.

### GetPauseUntilOk

`func (o *UpdateMonitorRequest) GetPauseUntilOk() (*time.Time, bool)`

GetPauseUntilOk returns a tuple with the PauseUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseUntil

`func (o *UpdateMonitorRequest) SetPauseUntil(v time.Time)`

SetPauseUntil sets PauseUntil field to given value.

### HasPauseUntil

`func (o *UpdateMonitorRequest) HasPauseUntil() bool`

HasPauseUntil returns a boolean if a field has been set.

### SetPauseUntilNil

`func (o *UpdateMonitorRequest) SetPauseUntilNil(b bool)`

 SetPauseUntilNil sets the value for PauseUntil to be an explicit nil

### UnsetPauseUntil
`func (o *UpdateMonitorRequest) UnsetPauseUntil()`

UnsetPauseUntil ensures that no value is present for PauseUntil, not even an explicit nil
### GetPermissions

`func (o *UpdateMonitorRequest) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateMonitorRequest) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateMonitorRequest) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.


### GetClusterIds

`func (o *UpdateMonitorRequest) GetClusterIds() []*string`

GetClusterIds returns the ClusterIds field if non-nil, zero value otherwise.

### GetClusterIdsOk

`func (o *UpdateMonitorRequest) GetClusterIdsOk() (*[]*string, bool)`

GetClusterIdsOk returns a tuple with the ClusterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIds

`func (o *UpdateMonitorRequest) SetClusterIds(v []*string)`

SetClusterIds sets ClusterIds field to given value.

### HasClusterIds

`func (o *UpdateMonitorRequest) HasClusterIds() bool`

HasClusterIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


