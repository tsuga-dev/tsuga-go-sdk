# TeamsFilterSpecificTeams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Match alert transitions associated with the listed teams. The caller must be allowed to access every listed team. | 
**Teams** | **[]string** | Team IDs selected by this filter. | 

## Methods

### NewTeamsFilterSpecificTeams

`func NewTeamsFilterSpecificTeams(type_ string, teams []string, ) *TeamsFilterSpecificTeams`

NewTeamsFilterSpecificTeams instantiates a new TeamsFilterSpecificTeams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsFilterSpecificTeamsWithDefaults

`func NewTeamsFilterSpecificTeamsWithDefaults() *TeamsFilterSpecificTeams`

NewTeamsFilterSpecificTeamsWithDefaults instantiates a new TeamsFilterSpecificTeams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TeamsFilterSpecificTeams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TeamsFilterSpecificTeams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TeamsFilterSpecificTeams) SetType(v string)`

SetType sets Type field to given value.


### GetTeams

`func (o *TeamsFilterSpecificTeams) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *TeamsFilterSpecificTeams) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *TeamsFilterSpecificTeams) SetTeams(v []string)`

SetTeams sets Teams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


