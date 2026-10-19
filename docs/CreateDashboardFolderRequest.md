# CreateDashboardFolderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the dashboard folder. Required and non-empty. | 
**Owner** | **string** | ID of the team that owns the dashboard folder. Required and used for folder access control. | 
**ParentFolderId** | Pointer to **NullableString** | ID of the top-level folder to nest this folder under. Omit on create or pass null to create a top-level folder; omit on update to keep the current parent; pass null on update to move the folder to top level. Non-null values must name a readable top-level folder. Folders can be nested at most two levels deep. | [optional] 
**Tags** | Pointer to [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 

## Methods

### NewCreateDashboardFolderRequest

`func NewCreateDashboardFolderRequest(name string, owner string, ) *CreateDashboardFolderRequest`

NewCreateDashboardFolderRequest instantiates a new CreateDashboardFolderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDashboardFolderRequestWithDefaults

`func NewCreateDashboardFolderRequestWithDefaults() *CreateDashboardFolderRequest`

NewCreateDashboardFolderRequestWithDefaults instantiates a new CreateDashboardFolderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateDashboardFolderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDashboardFolderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDashboardFolderRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *CreateDashboardFolderRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateDashboardFolderRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateDashboardFolderRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetParentFolderId

`func (o *CreateDashboardFolderRequest) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *CreateDashboardFolderRequest) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *CreateDashboardFolderRequest) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *CreateDashboardFolderRequest) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *CreateDashboardFolderRequest) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *CreateDashboardFolderRequest) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetTags

`func (o *CreateDashboardFolderRequest) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateDashboardFolderRequest) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateDashboardFolderRequest) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateDashboardFolderRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


