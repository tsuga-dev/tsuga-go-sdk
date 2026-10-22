# UpdateCustomRoleRequestTeamScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**TeamIds** | **[]string** | Team IDs the scope applies to. | 

## Methods

### NewUpdateCustomRoleRequestTeamScope

`func NewUpdateCustomRoleRequestTeamScope(type_ string, teamIds []string, ) *UpdateCustomRoleRequestTeamScope`

NewUpdateCustomRoleRequestTeamScope instantiates a new UpdateCustomRoleRequestTeamScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomRoleRequestTeamScopeWithDefaults

`func NewUpdateCustomRoleRequestTeamScopeWithDefaults() *UpdateCustomRoleRequestTeamScope`

NewUpdateCustomRoleRequestTeamScopeWithDefaults instantiates a new UpdateCustomRoleRequestTeamScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateCustomRoleRequestTeamScope) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateCustomRoleRequestTeamScope) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateCustomRoleRequestTeamScope) SetType(v string)`

SetType sets Type field to given value.


### GetTeamIds

`func (o *UpdateCustomRoleRequestTeamScope) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *UpdateCustomRoleRequestTeamScope) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *UpdateCustomRoleRequestTeamScope) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


