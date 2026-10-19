# BulkRemoveUsersFromTeamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamId** | **string** | Team ID to remove the requested users from. | 
**UserIds** | Pointer to **[]string** | User IDs to remove from the team. Items are processed independently: a user who is not a member or cannot be resolved does not stop the other items from being removed. | [optional] 
**UserEmails** | Pointer to **[]string** | User emails to resolve and remove from the team. | [optional] 

## Methods

### NewBulkRemoveUsersFromTeamRequest

`func NewBulkRemoveUsersFromTeamRequest(teamId string, ) *BulkRemoveUsersFromTeamRequest`

NewBulkRemoveUsersFromTeamRequest instantiates a new BulkRemoveUsersFromTeamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkRemoveUsersFromTeamRequestWithDefaults

`func NewBulkRemoveUsersFromTeamRequestWithDefaults() *BulkRemoveUsersFromTeamRequest`

NewBulkRemoveUsersFromTeamRequestWithDefaults instantiates a new BulkRemoveUsersFromTeamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamId

`func (o *BulkRemoveUsersFromTeamRequest) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *BulkRemoveUsersFromTeamRequest) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *BulkRemoveUsersFromTeamRequest) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetUserIds

`func (o *BulkRemoveUsersFromTeamRequest) GetUserIds() []*string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *BulkRemoveUsersFromTeamRequest) GetUserIdsOk() (*[]*string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *BulkRemoveUsersFromTeamRequest) SetUserIds(v []*string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *BulkRemoveUsersFromTeamRequest) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### GetUserEmails

`func (o *BulkRemoveUsersFromTeamRequest) GetUserEmails() []string`

GetUserEmails returns the UserEmails field if non-nil, zero value otherwise.

### GetUserEmailsOk

`func (o *BulkRemoveUsersFromTeamRequest) GetUserEmailsOk() (*[]string, bool)`

GetUserEmailsOk returns a tuple with the UserEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmails

`func (o *BulkRemoveUsersFromTeamRequest) SetUserEmails(v []string)`

SetUserEmails sets UserEmails field to given value.

### HasUserEmails

`func (o *BulkRemoveUsersFromTeamRequest) HasUserEmails() bool`

HasUserEmails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


