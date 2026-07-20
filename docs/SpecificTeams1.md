# SpecificTeams1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Match alert transitions associated with the listed teams. The caller must be allowed to access every listed team. | 
**Teams** | **[]string** | Team IDs selected by this filter. | 

## Methods

### NewSpecificTeams1

`func NewSpecificTeams1(type_ string, teams []string, ) *SpecificTeams1`

NewSpecificTeams1 instantiates a new SpecificTeams1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecificTeams1WithDefaults

`func NewSpecificTeams1WithDefaults() *SpecificTeams1`

NewSpecificTeams1WithDefaults instantiates a new SpecificTeams1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SpecificTeams1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpecificTeams1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpecificTeams1) SetType(v string)`

SetType sets Type field to given value.


### GetTeams

`func (o *SpecificTeams1) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *SpecificTeams1) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *SpecificTeams1) SetTeams(v []string)`

SetTeams sets Teams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


