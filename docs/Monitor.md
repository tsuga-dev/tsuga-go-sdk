# Monitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the monitor | 
**Name** | **string** | Display name of the monitor | 
**Message** | Pointer to **string** | Message to be displayed if a notification is triggered | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List of key/value tags applied to the resource | [optional] 
**Configuration** | [**CreateMonitorRequestConfiguration**](CreateMonitorRequestConfiguration.md) |  | 
**Priority** | **float32** | Priority of the monitor | 
**Owner** | **string** | Team ID that owns and manages the monitor | 
**DashboardId** | Pointer to **string** | Identifier of a dashboard related to the monitor | [optional] 
**Permissions** | **string** | This controls which data the monitor can see. | 

## Methods

### NewMonitor

`func NewMonitor(id string, name string, configuration CreateMonitorRequestConfiguration, priority float32, owner string, permissions string, ) *Monitor`

NewMonitor instantiates a new Monitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorWithDefaults

`func NewMonitorWithDefaults() *Monitor`

NewMonitorWithDefaults instantiates a new Monitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Monitor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Monitor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Monitor) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Monitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Monitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Monitor) SetName(v string)`

SetName sets Name field to given value.


### GetMessage

`func (o *Monitor) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Monitor) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Monitor) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Monitor) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTags

`func (o *Monitor) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Monitor) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Monitor) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Monitor) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetConfiguration

`func (o *Monitor) GetConfiguration() CreateMonitorRequestConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Monitor) GetConfigurationOk() (*CreateMonitorRequestConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Monitor) SetConfiguration(v CreateMonitorRequestConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetPriority

`func (o *Monitor) GetPriority() float32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Monitor) GetPriorityOk() (*float32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Monitor) SetPriority(v float32)`

SetPriority sets Priority field to given value.


### GetOwner

`func (o *Monitor) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Monitor) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Monitor) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetDashboardId

`func (o *Monitor) GetDashboardId() string`

GetDashboardId returns the DashboardId field if non-nil, zero value otherwise.

### GetDashboardIdOk

`func (o *Monitor) GetDashboardIdOk() (*string, bool)`

GetDashboardIdOk returns a tuple with the DashboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardId

`func (o *Monitor) SetDashboardId(v string)`

SetDashboardId sets DashboardId field to given value.

### HasDashboardId

`func (o *Monitor) HasDashboardId() bool`

HasDashboardId returns a boolean if a field has been set.

### GetPermissions

`func (o *Monitor) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Monitor) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Monitor) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


