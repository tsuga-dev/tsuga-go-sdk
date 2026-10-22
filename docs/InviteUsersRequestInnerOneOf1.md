# InviteUsersRequestInnerOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Lowercase email address to invite into the organization. | 
**TeamId** | **string** | Team ID to add the invited user to after acceptance. | 
**TeamRole** | **string** | Role the invited user receives in the target team once accepted: &#x60;admin&#x60;, &#x60;editor&#x60;, or &#x60;viewer&#x60;. | 
**UserRole** | **string** | Assigns the custom role identified by &#x60;customRoleId&#x60; after acceptance. | 
**CustomRoleId** | **string** | Identifier of the custom role to assign. Use the &#x60;id&#x60; returned by the custom role endpoints. | 

## Methods

### NewInviteUsersRequestInnerOneOf1

`func NewInviteUsersRequestInnerOneOf1(email string, teamId string, teamRole string, userRole string, customRoleId string, ) *InviteUsersRequestInnerOneOf1`

NewInviteUsersRequestInnerOneOf1 instantiates a new InviteUsersRequestInnerOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteUsersRequestInnerOneOf1WithDefaults

`func NewInviteUsersRequestInnerOneOf1WithDefaults() *InviteUsersRequestInnerOneOf1`

NewInviteUsersRequestInnerOneOf1WithDefaults instantiates a new InviteUsersRequestInnerOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InviteUsersRequestInnerOneOf1) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteUsersRequestInnerOneOf1) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteUsersRequestInnerOneOf1) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetTeamId

`func (o *InviteUsersRequestInnerOneOf1) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *InviteUsersRequestInnerOneOf1) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *InviteUsersRequestInnerOneOf1) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetTeamRole

`func (o *InviteUsersRequestInnerOneOf1) GetTeamRole() string`

GetTeamRole returns the TeamRole field if non-nil, zero value otherwise.

### GetTeamRoleOk

`func (o *InviteUsersRequestInnerOneOf1) GetTeamRoleOk() (*string, bool)`

GetTeamRoleOk returns a tuple with the TeamRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamRole

`func (o *InviteUsersRequestInnerOneOf1) SetTeamRole(v string)`

SetTeamRole sets TeamRole field to given value.


### GetUserRole

`func (o *InviteUsersRequestInnerOneOf1) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *InviteUsersRequestInnerOneOf1) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *InviteUsersRequestInnerOneOf1) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.


### GetCustomRoleId

`func (o *InviteUsersRequestInnerOneOf1) GetCustomRoleId() string`

GetCustomRoleId returns the CustomRoleId field if non-nil, zero value otherwise.

### GetCustomRoleIdOk

`func (o *InviteUsersRequestInnerOneOf1) GetCustomRoleIdOk() (*string, bool)`

GetCustomRoleIdOk returns a tuple with the CustomRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoleId

`func (o *InviteUsersRequestInnerOneOf1) SetCustomRoleId(v string)`

SetCustomRoleId sets CustomRoleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


