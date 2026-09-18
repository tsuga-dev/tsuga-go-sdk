# UpdateTagPolicyRequestTeamScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamIds** | **[]string** | Team IDs used by this policy scope. | 
**Mode** | **string** | &#x60;include&#x60; applies the policy only to listed teams. &#x60;exclude&#x60; applies it to all teams except the listed teams. | 

## Methods

### NewUpdateTagPolicyRequestTeamScope

`func NewUpdateTagPolicyRequestTeamScope(teamIds []string, mode string, ) *UpdateTagPolicyRequestTeamScope`

NewUpdateTagPolicyRequestTeamScope instantiates a new UpdateTagPolicyRequestTeamScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTagPolicyRequestTeamScopeWithDefaults

`func NewUpdateTagPolicyRequestTeamScopeWithDefaults() *UpdateTagPolicyRequestTeamScope`

NewUpdateTagPolicyRequestTeamScopeWithDefaults instantiates a new UpdateTagPolicyRequestTeamScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamIds

`func (o *UpdateTagPolicyRequestTeamScope) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *UpdateTagPolicyRequestTeamScope) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *UpdateTagPolicyRequestTeamScope) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.


### GetMode

`func (o *UpdateTagPolicyRequestTeamScope) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *UpdateTagPolicyRequestTeamScope) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *UpdateTagPolicyRequestTeamScope) SetMode(v string)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


