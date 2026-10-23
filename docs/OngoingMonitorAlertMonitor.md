# OngoingMonitorAlertMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Tsuga-generated monitor ID assigned when the monitor is created. | 
**Name** | **string** | Display name of the monitor and alert source. | 
**Owner** | **string** | Team ID that owns and manages the monitor. | 
**Message** | Pointer to **string** | Caller-supplied message included in notifications triggered by this monitor. Omitted when the monitor was created or last updated without one. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | [optional] 
**Configuration** | [**OngoingMonitorAlertMonitorConfiguration**](OngoingMonitorAlertMonitorConfiguration.md) |  | 
**Priority** | **int32** | Monitor priority from 1 through 5, where 1 is highest priority. | 
**Permissions** | **string** | &#x60;all&#x60; allows the resource to query all permitted telemetry, &#x60;owning-team-and-public&#x60; limits it to the owning team plus public data, and &#x60;owning-team-only&#x60; limits it to the owning team. | 
**DashboardId** | Pointer to **string** | ID of the dashboard linked to this monitor for context when investigating alerts. Caller-supplied. Omitted when no dashboard is linked. | [optional] 
**FolderId** | Pointer to **string** | ID of the monitor folder containing this monitor. Omitted for unfiled monitors. | [optional] 
**PauseUntil** | Pointer to **string** | Timestamp until which the monitor is snoozed. Omitted when the monitor is not snoozed. | [optional] 
**ClusterIds** | **[]string** | Cluster IDs where the monitor is deployed. Empty means all eligible clusters. | 

## Methods

### NewOngoingMonitorAlertMonitor

`func NewOngoingMonitorAlertMonitor(id string, name string, owner string, configuration OngoingMonitorAlertMonitorConfiguration, priority int32, permissions string, clusterIds []string, ) *OngoingMonitorAlertMonitor`

NewOngoingMonitorAlertMonitor instantiates a new OngoingMonitorAlertMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOngoingMonitorAlertMonitorWithDefaults

`func NewOngoingMonitorAlertMonitorWithDefaults() *OngoingMonitorAlertMonitor`

NewOngoingMonitorAlertMonitorWithDefaults instantiates a new OngoingMonitorAlertMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OngoingMonitorAlertMonitor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OngoingMonitorAlertMonitor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OngoingMonitorAlertMonitor) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *OngoingMonitorAlertMonitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OngoingMonitorAlertMonitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OngoingMonitorAlertMonitor) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *OngoingMonitorAlertMonitor) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *OngoingMonitorAlertMonitor) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *OngoingMonitorAlertMonitor) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetMessage

`func (o *OngoingMonitorAlertMonitor) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OngoingMonitorAlertMonitor) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OngoingMonitorAlertMonitor) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OngoingMonitorAlertMonitor) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTags

`func (o *OngoingMonitorAlertMonitor) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OngoingMonitorAlertMonitor) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OngoingMonitorAlertMonitor) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OngoingMonitorAlertMonitor) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetConfiguration

`func (o *OngoingMonitorAlertMonitor) GetConfiguration() OngoingMonitorAlertMonitorConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *OngoingMonitorAlertMonitor) GetConfigurationOk() (*OngoingMonitorAlertMonitorConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *OngoingMonitorAlertMonitor) SetConfiguration(v OngoingMonitorAlertMonitorConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetPriority

`func (o *OngoingMonitorAlertMonitor) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *OngoingMonitorAlertMonitor) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *OngoingMonitorAlertMonitor) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetPermissions

`func (o *OngoingMonitorAlertMonitor) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *OngoingMonitorAlertMonitor) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *OngoingMonitorAlertMonitor) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.


### GetDashboardId

`func (o *OngoingMonitorAlertMonitor) GetDashboardId() string`

GetDashboardId returns the DashboardId field if non-nil, zero value otherwise.

### GetDashboardIdOk

`func (o *OngoingMonitorAlertMonitor) GetDashboardIdOk() (*string, bool)`

GetDashboardIdOk returns a tuple with the DashboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardId

`func (o *OngoingMonitorAlertMonitor) SetDashboardId(v string)`

SetDashboardId sets DashboardId field to given value.

### HasDashboardId

`func (o *OngoingMonitorAlertMonitor) HasDashboardId() bool`

HasDashboardId returns a boolean if a field has been set.

### GetFolderId

`func (o *OngoingMonitorAlertMonitor) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *OngoingMonitorAlertMonitor) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *OngoingMonitorAlertMonitor) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *OngoingMonitorAlertMonitor) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetPauseUntil

`func (o *OngoingMonitorAlertMonitor) GetPauseUntil() string`

GetPauseUntil returns the PauseUntil field if non-nil, zero value otherwise.

### GetPauseUntilOk

`func (o *OngoingMonitorAlertMonitor) GetPauseUntilOk() (*string, bool)`

GetPauseUntilOk returns a tuple with the PauseUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseUntil

`func (o *OngoingMonitorAlertMonitor) SetPauseUntil(v string)`

SetPauseUntil sets PauseUntil field to given value.

### HasPauseUntil

`func (o *OngoingMonitorAlertMonitor) HasPauseUntil() bool`

HasPauseUntil returns a boolean if a field has been set.

### GetClusterIds

`func (o *OngoingMonitorAlertMonitor) GetClusterIds() []string`

GetClusterIds returns the ClusterIds field if non-nil, zero value otherwise.

### GetClusterIdsOk

`func (o *OngoingMonitorAlertMonitor) GetClusterIdsOk() (*[]string, bool)`

GetClusterIdsOk returns a tuple with the ClusterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIds

`func (o *OngoingMonitorAlertMonitor) SetClusterIds(v []string)`

SetClusterIds sets ClusterIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


