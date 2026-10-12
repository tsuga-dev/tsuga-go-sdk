# InputAllTeams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Match alert transitions for every team in the organization. Public API callers need an API key with &#x60;all&#x60; scope to create or update this filter. | 

## Methods

### NewInputAllTeams

`func NewInputAllTeams(type_ string, ) *InputAllTeams`

NewInputAllTeams instantiates a new InputAllTeams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputAllTeamsWithDefaults

`func NewInputAllTeamsWithDefaults() *InputAllTeams`

NewInputAllTeamsWithDefaults instantiates a new InputAllTeams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputAllTeams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputAllTeams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputAllTeams) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


