# RuleTeamsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Teams** | **[]string** | Team IDs that narrow down the teams that can receive notifications from this rule | 

## Methods

### NewRuleTeamsFilter

`func NewRuleTeamsFilter(type_ string, teams []string, ) *RuleTeamsFilter`

NewRuleTeamsFilter instantiates a new RuleTeamsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTeamsFilterWithDefaults

`func NewRuleTeamsFilterWithDefaults() *RuleTeamsFilter`

NewRuleTeamsFilterWithDefaults instantiates a new RuleTeamsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuleTeamsFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleTeamsFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleTeamsFilter) SetType(v string)`

SetType sets Type field to given value.


### GetTeams

`func (o *RuleTeamsFilter) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *RuleTeamsFilter) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *RuleTeamsFilter) SetTeams(v []string)`

SetTeams sets Teams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


