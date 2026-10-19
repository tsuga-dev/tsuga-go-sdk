# TeamMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Tsuga-generated team membership ID assigned when the membership is created. | 
**UserId** | **string** | User ID that belongs to the team. | 
**TeamId** | **string** | Team ID that contains the user. | 
**RoleKey** | **string** | Role granted to the user inside this team. &#x60;admin&#x60; gives full team control, including managing assets, members, and settings. &#x60;editor&#x60; can view and update most team assets (dashboards, monitors, etc.) but cannot manage members or settings. &#x60;viewer&#x60; can view team data but cannot make changes. | 

## Methods

### NewTeamMembership

`func NewTeamMembership(id string, userId string, teamId string, roleKey string, ) *TeamMembership`

NewTeamMembership instantiates a new TeamMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamMembershipWithDefaults

`func NewTeamMembershipWithDefaults() *TeamMembership`

NewTeamMembershipWithDefaults instantiates a new TeamMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamMembership) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamMembership) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamMembership) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *TeamMembership) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TeamMembership) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TeamMembership) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetTeamId

`func (o *TeamMembership) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *TeamMembership) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *TeamMembership) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetRoleKey

`func (o *TeamMembership) GetRoleKey() string`

GetRoleKey returns the RoleKey field if non-nil, zero value otherwise.

### GetRoleKeyOk

`func (o *TeamMembership) GetRoleKeyOk() (*string, bool)`

GetRoleKeyOk returns a tuple with the RoleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleKey

`func (o *TeamMembership) SetRoleKey(v string)`

SetRoleKey sets RoleKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


