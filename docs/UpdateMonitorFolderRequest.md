# UpdateMonitorFolderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the monitor folder. | 
**Owner** | **string** | ID of the team that owns the monitor folder and controls folder access. | 
**ParentFolderId** | Pointer to **NullableString** | ID of the top-level folder to nest this folder under. Omit on create or pass null to create a top-level folder; omit on update to keep the current parent; pass null on update to move the folder to top level. Non-null values must name a readable top-level folder. Folders can be nested at most two levels deep. | [optional] 
**Tags** | Pointer to [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 

## Methods

### NewUpdateMonitorFolderRequest

`func NewUpdateMonitorFolderRequest(name string, owner string, ) *UpdateMonitorFolderRequest`

NewUpdateMonitorFolderRequest instantiates a new UpdateMonitorFolderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMonitorFolderRequestWithDefaults

`func NewUpdateMonitorFolderRequestWithDefaults() *UpdateMonitorFolderRequest`

NewUpdateMonitorFolderRequestWithDefaults instantiates a new UpdateMonitorFolderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateMonitorFolderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMonitorFolderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMonitorFolderRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *UpdateMonitorFolderRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateMonitorFolderRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateMonitorFolderRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetParentFolderId

`func (o *UpdateMonitorFolderRequest) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *UpdateMonitorFolderRequest) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *UpdateMonitorFolderRequest) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *UpdateMonitorFolderRequest) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *UpdateMonitorFolderRequest) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *UpdateMonitorFolderRequest) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetTags

`func (o *UpdateMonitorFolderRequest) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateMonitorFolderRequest) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateMonitorFolderRequest) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateMonitorFolderRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


