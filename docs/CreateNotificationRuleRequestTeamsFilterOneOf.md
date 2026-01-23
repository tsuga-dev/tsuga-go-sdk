# CreateNotificationRuleRequestTeamsFilterOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Teams** | **[]string** | Team IDs that narrow down the teams that can receive notifications from this rule | 

## Methods

### NewCreateNotificationRuleRequestTeamsFilterOneOf

`func NewCreateNotificationRuleRequestTeamsFilterOneOf(type_ string, teams []string, ) *CreateNotificationRuleRequestTeamsFilterOneOf`

NewCreateNotificationRuleRequestTeamsFilterOneOf instantiates a new CreateNotificationRuleRequestTeamsFilterOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationRuleRequestTeamsFilterOneOfWithDefaults

`func NewCreateNotificationRuleRequestTeamsFilterOneOfWithDefaults() *CreateNotificationRuleRequestTeamsFilterOneOf`

NewCreateNotificationRuleRequestTeamsFilterOneOfWithDefaults instantiates a new CreateNotificationRuleRequestTeamsFilterOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateNotificationRuleRequestTeamsFilterOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNotificationRuleRequestTeamsFilterOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNotificationRuleRequestTeamsFilterOneOf) SetType(v string)`

SetType sets Type field to given value.


### GetTeams

`func (o *CreateNotificationRuleRequestTeamsFilterOneOf) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *CreateNotificationRuleRequestTeamsFilterOneOf) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *CreateNotificationRuleRequestTeamsFilterOneOf) SetTeams(v []string)`

SetTeams sets Teams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


