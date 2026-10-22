# CreateCustomRoleRequestTeamScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**TeamIds** | **[]string** | Team IDs the scope applies to. | 

## Methods

### NewCreateCustomRoleRequestTeamScope

`func NewCreateCustomRoleRequestTeamScope(type_ string, teamIds []string, ) *CreateCustomRoleRequestTeamScope`

NewCreateCustomRoleRequestTeamScope instantiates a new CreateCustomRoleRequestTeamScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomRoleRequestTeamScopeWithDefaults

`func NewCreateCustomRoleRequestTeamScopeWithDefaults() *CreateCustomRoleRequestTeamScope`

NewCreateCustomRoleRequestTeamScopeWithDefaults instantiates a new CreateCustomRoleRequestTeamScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateCustomRoleRequestTeamScope) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateCustomRoleRequestTeamScope) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateCustomRoleRequestTeamScope) SetType(v string)`

SetType sets Type field to given value.


### GetTeamIds

`func (o *CreateCustomRoleRequestTeamScope) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *CreateCustomRoleRequestTeamScope) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *CreateCustomRoleRequestTeamScope) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


