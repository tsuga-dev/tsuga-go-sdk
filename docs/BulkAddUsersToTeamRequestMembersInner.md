# BulkAddUsersToTeamRequestMembersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | Existing user ID to add to the team. | [optional] 
**UserEmail** | Pointer to **string** | Existing user email to resolve and add to the team. | [optional] 
**RoleKey** | **string** | Role granted to the user inside this team. &#x60;admin&#x60; gives full team control, including managing assets, members, and settings. &#x60;editor&#x60; can view and update most team assets (dashboards, monitors, etc.) but cannot manage members or settings. &#x60;viewer&#x60; can view team data but cannot make changes. | 

## Methods

### NewBulkAddUsersToTeamRequestMembersInner

`func NewBulkAddUsersToTeamRequestMembersInner(roleKey string, ) *BulkAddUsersToTeamRequestMembersInner`

NewBulkAddUsersToTeamRequestMembersInner instantiates a new BulkAddUsersToTeamRequestMembersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAddUsersToTeamRequestMembersInnerWithDefaults

`func NewBulkAddUsersToTeamRequestMembersInnerWithDefaults() *BulkAddUsersToTeamRequestMembersInner`

NewBulkAddUsersToTeamRequestMembersInnerWithDefaults instantiates a new BulkAddUsersToTeamRequestMembersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *BulkAddUsersToTeamRequestMembersInner) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BulkAddUsersToTeamRequestMembersInner) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BulkAddUsersToTeamRequestMembersInner) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BulkAddUsersToTeamRequestMembersInner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserEmail

`func (o *BulkAddUsersToTeamRequestMembersInner) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *BulkAddUsersToTeamRequestMembersInner) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *BulkAddUsersToTeamRequestMembersInner) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *BulkAddUsersToTeamRequestMembersInner) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetRoleKey

`func (o *BulkAddUsersToTeamRequestMembersInner) GetRoleKey() string`

GetRoleKey returns the RoleKey field if non-nil, zero value otherwise.

### GetRoleKeyOk

`func (o *BulkAddUsersToTeamRequestMembersInner) GetRoleKeyOk() (*string, bool)`

GetRoleKeyOk returns a tuple with the RoleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleKey

`func (o *BulkAddUsersToTeamRequestMembersInner) SetRoleKey(v string)`

SetRoleKey sets RoleKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


