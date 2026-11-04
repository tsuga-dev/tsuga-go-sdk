# InviteUsersRequestInnerAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Lowercase email address to invite into the organization. | 
**TeamId** | **string** | Team ID to add the invited user to after acceptance. | 
**TeamRole** | **string** | Role the invited user receives in the target team once accepted: &#x60;admin&#x60;, &#x60;editor&#x60;, or &#x60;viewer&#x60;. | 
**UserRole** | **string** | Assigns the custom role identified by &#x60;customRoleId&#x60; after acceptance. | 
**CustomRoleId** | **string** | Identifier of the custom role to assign. Use the &#x60;id&#x60; returned by the custom role endpoints. | 

## Methods

### NewInviteUsersRequestInnerAnyOf

`func NewInviteUsersRequestInnerAnyOf(email string, teamId string, teamRole string, userRole string, customRoleId string, ) *InviteUsersRequestInnerAnyOf`

NewInviteUsersRequestInnerAnyOf instantiates a new InviteUsersRequestInnerAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteUsersRequestInnerAnyOfWithDefaults

`func NewInviteUsersRequestInnerAnyOfWithDefaults() *InviteUsersRequestInnerAnyOf`

NewInviteUsersRequestInnerAnyOfWithDefaults instantiates a new InviteUsersRequestInnerAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InviteUsersRequestInnerAnyOf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteUsersRequestInnerAnyOf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteUsersRequestInnerAnyOf) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetTeamId

`func (o *InviteUsersRequestInnerAnyOf) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *InviteUsersRequestInnerAnyOf) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *InviteUsersRequestInnerAnyOf) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetTeamRole

`func (o *InviteUsersRequestInnerAnyOf) GetTeamRole() string`

GetTeamRole returns the TeamRole field if non-nil, zero value otherwise.

### GetTeamRoleOk

`func (o *InviteUsersRequestInnerAnyOf) GetTeamRoleOk() (*string, bool)`

GetTeamRoleOk returns a tuple with the TeamRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamRole

`func (o *InviteUsersRequestInnerAnyOf) SetTeamRole(v string)`

SetTeamRole sets TeamRole field to given value.


### GetUserRole

`func (o *InviteUsersRequestInnerAnyOf) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *InviteUsersRequestInnerAnyOf) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *InviteUsersRequestInnerAnyOf) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.


### GetCustomRoleId

`func (o *InviteUsersRequestInnerAnyOf) GetCustomRoleId() string`

GetCustomRoleId returns the CustomRoleId field if non-nil, zero value otherwise.

### GetCustomRoleIdOk

`func (o *InviteUsersRequestInnerAnyOf) GetCustomRoleIdOk() (*string, bool)`

GetCustomRoleIdOk returns a tuple with the CustomRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoleId

`func (o *InviteUsersRequestInnerAnyOf) SetCustomRoleId(v string)`

SetCustomRoleId sets CustomRoleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


