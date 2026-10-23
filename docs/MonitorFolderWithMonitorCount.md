# MonitorFolderWithMonitorCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Tsuga-generated monitor folder ID assigned when the folder is created. Use it as monitor &#x60;folderId&#x60; or child-folder &#x60;parentFolderId&#x60;. | 
**Name** | **string** | Display name of the monitor folder. Set by the caller on create or update. | 
**Owner** | **string** | ID of the team that owns the monitor folder. Set by the caller on create or update and used for folder access control. | 
**ParentFolderId** | Pointer to **string** | ID of the top-level parent folder when this folder is nested. Set by the caller on create or update. Returned only for child folders; omitted for top-level folders. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | [optional] 
**MonitorCount** | **float32** | Number of monitors in this folder and its readable subfolders that the caller can read. Returned only by folder list responses; callers without monitor read access see 0. | 

## Methods

### NewMonitorFolderWithMonitorCount

`func NewMonitorFolderWithMonitorCount(id string, name string, owner string, monitorCount float32, ) *MonitorFolderWithMonitorCount`

NewMonitorFolderWithMonitorCount instantiates a new MonitorFolderWithMonitorCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorFolderWithMonitorCountWithDefaults

`func NewMonitorFolderWithMonitorCountWithDefaults() *MonitorFolderWithMonitorCount`

NewMonitorFolderWithMonitorCountWithDefaults instantiates a new MonitorFolderWithMonitorCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MonitorFolderWithMonitorCount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MonitorFolderWithMonitorCount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MonitorFolderWithMonitorCount) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MonitorFolderWithMonitorCount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitorFolderWithMonitorCount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitorFolderWithMonitorCount) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *MonitorFolderWithMonitorCount) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MonitorFolderWithMonitorCount) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MonitorFolderWithMonitorCount) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetParentFolderId

`func (o *MonitorFolderWithMonitorCount) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *MonitorFolderWithMonitorCount) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *MonitorFolderWithMonitorCount) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *MonitorFolderWithMonitorCount) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### GetTags

`func (o *MonitorFolderWithMonitorCount) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MonitorFolderWithMonitorCount) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MonitorFolderWithMonitorCount) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MonitorFolderWithMonitorCount) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMonitorCount

`func (o *MonitorFolderWithMonitorCount) GetMonitorCount() float32`

GetMonitorCount returns the MonitorCount field if non-nil, zero value otherwise.

### GetMonitorCountOk

`func (o *MonitorFolderWithMonitorCount) GetMonitorCountOk() (*float32, bool)`

GetMonitorCountOk returns a tuple with the MonitorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorCount

`func (o *MonitorFolderWithMonitorCount) SetMonitorCount(v float32)`

SetMonitorCount sets MonitorCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


