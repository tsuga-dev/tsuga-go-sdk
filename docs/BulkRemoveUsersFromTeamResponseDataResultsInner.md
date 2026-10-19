# BulkRemoveUsersFromTeamResponseDataResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | Resolved user ID for this bulk remove result. For &#x60;user_not_found&#x60;, this is the submitted user ID or email that Tsuga could not resolve. | 
**Status** | **string** | &#x60;removed&#x60; means the membership was deleted, &#x60;not_member&#x60; means no membership existed, and &#x60;user_not_found&#x60; means Tsuga could not resolve the requested user. | 

## Methods

### NewBulkRemoveUsersFromTeamResponseDataResultsInner

`func NewBulkRemoveUsersFromTeamResponseDataResultsInner(userId string, status string, ) *BulkRemoveUsersFromTeamResponseDataResultsInner`

NewBulkRemoveUsersFromTeamResponseDataResultsInner instantiates a new BulkRemoveUsersFromTeamResponseDataResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkRemoveUsersFromTeamResponseDataResultsInnerWithDefaults

`func NewBulkRemoveUsersFromTeamResponseDataResultsInnerWithDefaults() *BulkRemoveUsersFromTeamResponseDataResultsInner`

NewBulkRemoveUsersFromTeamResponseDataResultsInnerWithDefaults instantiates a new BulkRemoveUsersFromTeamResponseDataResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *BulkRemoveUsersFromTeamResponseDataResultsInner) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BulkRemoveUsersFromTeamResponseDataResultsInner) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BulkRemoveUsersFromTeamResponseDataResultsInner) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetStatus

`func (o *BulkRemoveUsersFromTeamResponseDataResultsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkRemoveUsersFromTeamResponseDataResultsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkRemoveUsersFromTeamResponseDataResultsInner) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


