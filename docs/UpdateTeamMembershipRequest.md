# UpdateTeamMembershipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** |  | 
**TeamId** | **string** |  | 
**RoleKey** | **string** |  | 

## Methods

### NewUpdateTeamMembershipRequest

`func NewUpdateTeamMembershipRequest(userId string, teamId string, roleKey string, ) *UpdateTeamMembershipRequest`

NewUpdateTeamMembershipRequest instantiates a new UpdateTeamMembershipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTeamMembershipRequestWithDefaults

`func NewUpdateTeamMembershipRequestWithDefaults() *UpdateTeamMembershipRequest`

NewUpdateTeamMembershipRequestWithDefaults instantiates a new UpdateTeamMembershipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UpdateTeamMembershipRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UpdateTeamMembershipRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UpdateTeamMembershipRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetTeamId

`func (o *UpdateTeamMembershipRequest) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *UpdateTeamMembershipRequest) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *UpdateTeamMembershipRequest) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetRoleKey

`func (o *UpdateTeamMembershipRequest) GetRoleKey() string`

GetRoleKey returns the RoleKey field if non-nil, zero value otherwise.

### GetRoleKeyOk

`func (o *UpdateTeamMembershipRequest) GetRoleKeyOk() (*string, bool)`

GetRoleKeyOk returns a tuple with the RoleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleKey

`func (o *UpdateTeamMembershipRequest) SetRoleKey(v string)`

SetRoleKey sets RoleKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


