# BulkAddUsersToTeamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamId** | **string** | Team ID receiving all requested memberships. | 
**Members** | [**[]BulkAddUsersToTeamRequestMembersInner**](BulkAddUsersToTeamRequestMembersInner.md) | Users to add to the team. Each item must include either &#x60;userId&#x60; or &#x60;userEmail&#x60;, not both. Duplicate identifiers (including a &#x60;userId&#x60; and &#x60;userEmail&#x60; that resolve to the same user) are rejected. Items are processed independently: an unresolved user does not stop other items from being added. | 

## Methods

### NewBulkAddUsersToTeamRequest

`func NewBulkAddUsersToTeamRequest(teamId string, members []BulkAddUsersToTeamRequestMembersInner, ) *BulkAddUsersToTeamRequest`

NewBulkAddUsersToTeamRequest instantiates a new BulkAddUsersToTeamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAddUsersToTeamRequestWithDefaults

`func NewBulkAddUsersToTeamRequestWithDefaults() *BulkAddUsersToTeamRequest`

NewBulkAddUsersToTeamRequestWithDefaults instantiates a new BulkAddUsersToTeamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamId

`func (o *BulkAddUsersToTeamRequest) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *BulkAddUsersToTeamRequest) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *BulkAddUsersToTeamRequest) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetMembers

`func (o *BulkAddUsersToTeamRequest) GetMembers() []BulkAddUsersToTeamRequestMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *BulkAddUsersToTeamRequest) GetMembersOk() (*[]BulkAddUsersToTeamRequestMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *BulkAddUsersToTeamRequest) SetMembers(v []BulkAddUsersToTeamRequestMembersInner)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


