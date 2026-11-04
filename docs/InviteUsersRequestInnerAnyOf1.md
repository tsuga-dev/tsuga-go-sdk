# InviteUsersRequestInnerAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Lowercase email address to invite into the organization. | 
**TeamId** | **string** | Team ID to add the invited user to after acceptance. | 
**TeamRole** | **string** | Role the invited user receives in the target team once accepted: &#x60;admin&#x60;, &#x60;editor&#x60;, or &#x60;viewer&#x60;. | 
**UserRole** | [**InviteUsersRequestInnerAnyOf1UserRole**](InviteUsersRequestInnerAnyOf1UserRole.md) |  | 

## Methods

### NewInviteUsersRequestInnerAnyOf1

`func NewInviteUsersRequestInnerAnyOf1(email string, teamId string, teamRole string, userRole InviteUsersRequestInnerAnyOf1UserRole, ) *InviteUsersRequestInnerAnyOf1`

NewInviteUsersRequestInnerAnyOf1 instantiates a new InviteUsersRequestInnerAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteUsersRequestInnerAnyOf1WithDefaults

`func NewInviteUsersRequestInnerAnyOf1WithDefaults() *InviteUsersRequestInnerAnyOf1`

NewInviteUsersRequestInnerAnyOf1WithDefaults instantiates a new InviteUsersRequestInnerAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InviteUsersRequestInnerAnyOf1) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteUsersRequestInnerAnyOf1) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteUsersRequestInnerAnyOf1) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetTeamId

`func (o *InviteUsersRequestInnerAnyOf1) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *InviteUsersRequestInnerAnyOf1) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *InviteUsersRequestInnerAnyOf1) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetTeamRole

`func (o *InviteUsersRequestInnerAnyOf1) GetTeamRole() string`

GetTeamRole returns the TeamRole field if non-nil, zero value otherwise.

### GetTeamRoleOk

`func (o *InviteUsersRequestInnerAnyOf1) GetTeamRoleOk() (*string, bool)`

GetTeamRoleOk returns a tuple with the TeamRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamRole

`func (o *InviteUsersRequestInnerAnyOf1) SetTeamRole(v string)`

SetTeamRole sets TeamRole field to given value.


### GetUserRole

`func (o *InviteUsersRequestInnerAnyOf1) GetUserRole() InviteUsersRequestInnerAnyOf1UserRole`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *InviteUsersRequestInnerAnyOf1) GetUserRoleOk() (*InviteUsersRequestInnerAnyOf1UserRole, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *InviteUsersRequestInnerAnyOf1) SetUserRole(v InviteUsersRequestInnerAnyOf1UserRole)`

SetUserRole sets UserRole field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


