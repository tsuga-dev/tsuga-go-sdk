# InviteUsersRequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**TeamId** | **string** |  | 
**TeamRole** | **string** |  | 
**UserRole** | **string** |  | 

## Methods

### NewInviteUsersRequestInner

`func NewInviteUsersRequestInner(email string, teamId string, teamRole string, userRole string, ) *InviteUsersRequestInner`

NewInviteUsersRequestInner instantiates a new InviteUsersRequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteUsersRequestInnerWithDefaults

`func NewInviteUsersRequestInnerWithDefaults() *InviteUsersRequestInner`

NewInviteUsersRequestInnerWithDefaults instantiates a new InviteUsersRequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InviteUsersRequestInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteUsersRequestInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteUsersRequestInner) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetTeamId

`func (o *InviteUsersRequestInner) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *InviteUsersRequestInner) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *InviteUsersRequestInner) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetTeamRole

`func (o *InviteUsersRequestInner) GetTeamRole() string`

GetTeamRole returns the TeamRole field if non-nil, zero value otherwise.

### GetTeamRoleOk

`func (o *InviteUsersRequestInner) GetTeamRoleOk() (*string, bool)`

GetTeamRoleOk returns a tuple with the TeamRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamRole

`func (o *InviteUsersRequestInner) SetTeamRole(v string)`

SetTeamRole sets TeamRole field to given value.


### GetUserRole

`func (o *InviteUsersRequestInner) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *InviteUsersRequestInner) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *InviteUsersRequestInner) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


