# DashboardFolderWithDashboardCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Tsuga-generated dashboard folder ID assigned when the folder is created. Use it as dashboard &#x60;folderId&#x60; or child-folder &#x60;parentFolderId&#x60;. | 
**Name** | **string** | Display name of the dashboard folder. Set by the caller on create or update. | 
**Owner** | **string** | ID of the team that owns the dashboard folder. Set by the caller on create or update and used for folder access control. | 
**ParentFolderId** | Pointer to **string** | ID of the top-level parent folder when this folder is nested. Set by the caller on create or update. Returned only for child folders; omitted for top-level folders. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | [optional] 
**DashboardCount** | **float32** | Number of dashboards in this folder and its readable subfolders that the caller can read. Returned only by folder list responses; callers without dashboard read access see 0. | 

## Methods

### NewDashboardFolderWithDashboardCount

`func NewDashboardFolderWithDashboardCount(id string, name string, owner string, dashboardCount float32, ) *DashboardFolderWithDashboardCount`

NewDashboardFolderWithDashboardCount instantiates a new DashboardFolderWithDashboardCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardFolderWithDashboardCountWithDefaults

`func NewDashboardFolderWithDashboardCountWithDefaults() *DashboardFolderWithDashboardCount`

NewDashboardFolderWithDashboardCountWithDefaults instantiates a new DashboardFolderWithDashboardCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DashboardFolderWithDashboardCount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardFolderWithDashboardCount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardFolderWithDashboardCount) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DashboardFolderWithDashboardCount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardFolderWithDashboardCount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardFolderWithDashboardCount) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *DashboardFolderWithDashboardCount) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DashboardFolderWithDashboardCount) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DashboardFolderWithDashboardCount) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetParentFolderId

`func (o *DashboardFolderWithDashboardCount) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *DashboardFolderWithDashboardCount) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *DashboardFolderWithDashboardCount) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *DashboardFolderWithDashboardCount) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### GetTags

`func (o *DashboardFolderWithDashboardCount) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DashboardFolderWithDashboardCount) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DashboardFolderWithDashboardCount) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DashboardFolderWithDashboardCount) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDashboardCount

`func (o *DashboardFolderWithDashboardCount) GetDashboardCount() float32`

GetDashboardCount returns the DashboardCount field if non-nil, zero value otherwise.

### GetDashboardCountOk

`func (o *DashboardFolderWithDashboardCount) GetDashboardCountOk() (*float32, bool)`

GetDashboardCountOk returns a tuple with the DashboardCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardCount

`func (o *DashboardFolderWithDashboardCount) SetDashboardCount(v float32)`

SetDashboardCount sets DashboardCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


