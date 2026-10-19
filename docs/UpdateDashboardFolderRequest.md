# UpdateDashboardFolderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the dashboard folder. Required and non-empty. | 
**Owner** | **string** | ID of the team that owns the dashboard folder. Required and used for folder access control. | 
**ParentFolderId** | Pointer to **NullableString** | Folder filter for dashboard query results. Optional; pass a folder ID to return dashboards in that folder, or null to return only dashboards that are not in any folder. | [optional] 
**Tags** | Pointer to [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 

## Methods

### NewUpdateDashboardFolderRequest

`func NewUpdateDashboardFolderRequest(name string, owner string, ) *UpdateDashboardFolderRequest`

NewUpdateDashboardFolderRequest instantiates a new UpdateDashboardFolderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDashboardFolderRequestWithDefaults

`func NewUpdateDashboardFolderRequestWithDefaults() *UpdateDashboardFolderRequest`

NewUpdateDashboardFolderRequestWithDefaults instantiates a new UpdateDashboardFolderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateDashboardFolderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDashboardFolderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDashboardFolderRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *UpdateDashboardFolderRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateDashboardFolderRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateDashboardFolderRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetParentFolderId

`func (o *UpdateDashboardFolderRequest) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *UpdateDashboardFolderRequest) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *UpdateDashboardFolderRequest) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *UpdateDashboardFolderRequest) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *UpdateDashboardFolderRequest) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *UpdateDashboardFolderRequest) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetTags

`func (o *UpdateDashboardFolderRequest) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateDashboardFolderRequest) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateDashboardFolderRequest) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateDashboardFolderRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


