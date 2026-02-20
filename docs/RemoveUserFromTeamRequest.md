# RemoveUserFromTeamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** |  | 
**TeamId** | **string** |  | 

## Methods

### NewRemoveUserFromTeamRequest

`func NewRemoveUserFromTeamRequest(userId string, teamId string, ) *RemoveUserFromTeamRequest`

NewRemoveUserFromTeamRequest instantiates a new RemoveUserFromTeamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveUserFromTeamRequestWithDefaults

`func NewRemoveUserFromTeamRequestWithDefaults() *RemoveUserFromTeamRequest`

NewRemoveUserFromTeamRequestWithDefaults instantiates a new RemoveUserFromTeamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *RemoveUserFromTeamRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RemoveUserFromTeamRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RemoveUserFromTeamRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetTeamId

`func (o *RemoveUserFromTeamRequest) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *RemoveUserFromTeamRequest) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *RemoveUserFromTeamRequest) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


