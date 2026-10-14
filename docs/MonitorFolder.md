# MonitorFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Tsuga-generated monitor folder ID assigned when the folder is created. Use it as monitor &#x60;folderId&#x60; or child-folder &#x60;parentFolderId&#x60;. | 
**Name** | **string** | Display name of the monitor folder. Set by the caller on create or update. | 
**Owner** | **string** | ID of the team that owns the monitor folder. Set by the caller on create or update and used for folder access control. | 
**ParentFolderId** | Pointer to **string** | ID of the top-level parent folder when this folder is nested. Set by the caller on create or update. Returned only for child folders; omitted for top-level folders. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | [optional] 

## Methods

### NewMonitorFolder

`func NewMonitorFolder(id string, name string, owner string, ) *MonitorFolder`

NewMonitorFolder instantiates a new MonitorFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorFolderWithDefaults

`func NewMonitorFolderWithDefaults() *MonitorFolder`

NewMonitorFolderWithDefaults instantiates a new MonitorFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MonitorFolder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MonitorFolder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MonitorFolder) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MonitorFolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitorFolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitorFolder) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *MonitorFolder) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MonitorFolder) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MonitorFolder) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetParentFolderId

`func (o *MonitorFolder) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *MonitorFolder) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *MonitorFolder) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *MonitorFolder) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### GetTags

`func (o *MonitorFolder) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MonitorFolder) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MonitorFolder) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MonitorFolder) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


