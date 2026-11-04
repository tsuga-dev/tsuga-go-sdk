# InviteUsersRequestInnerAnyOf1UserRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Built-in organization role: &#x60;admin&#x60;, &#x60;public_editor&#x60;, &#x60;public_viewer&#x60;, or &#x60;teams_only&#x60;. | 
**CustomRoleId** | **string** | Identifier of the custom role to assign. Use the &#x60;id&#x60; returned by the custom role endpoints. | 

## Methods

### NewInviteUsersRequestInnerAnyOf1UserRole

`func NewInviteUsersRequestInnerAnyOf1UserRole(type_ string, customRoleId string, ) *InviteUsersRequestInnerAnyOf1UserRole`

NewInviteUsersRequestInnerAnyOf1UserRole instantiates a new InviteUsersRequestInnerAnyOf1UserRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteUsersRequestInnerAnyOf1UserRoleWithDefaults

`func NewInviteUsersRequestInnerAnyOf1UserRoleWithDefaults() *InviteUsersRequestInnerAnyOf1UserRole`

NewInviteUsersRequestInnerAnyOf1UserRoleWithDefaults instantiates a new InviteUsersRequestInnerAnyOf1UserRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InviteUsersRequestInnerAnyOf1UserRole) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InviteUsersRequestInnerAnyOf1UserRole) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InviteUsersRequestInnerAnyOf1UserRole) SetType(v string)`

SetType sets Type field to given value.


### GetCustomRoleId

`func (o *InviteUsersRequestInnerAnyOf1UserRole) GetCustomRoleId() string`

GetCustomRoleId returns the CustomRoleId field if non-nil, zero value otherwise.

### GetCustomRoleIdOk

`func (o *InviteUsersRequestInnerAnyOf1UserRole) GetCustomRoleIdOk() (*string, bool)`

GetCustomRoleIdOk returns a tuple with the CustomRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoleId

`func (o *InviteUsersRequestInnerAnyOf1UserRole) SetCustomRoleId(v string)`

SetCustomRoleId sets CustomRoleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


