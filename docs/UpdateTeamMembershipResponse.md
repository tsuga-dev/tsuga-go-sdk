# UpdateTeamMembershipResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**TeamMembership**](TeamMembership.md) |  | 

## Methods

### NewUpdateTeamMembershipResponse

`func NewUpdateTeamMembershipResponse(requestId string, data TeamMembership, ) *UpdateTeamMembershipResponse`

NewUpdateTeamMembershipResponse instantiates a new UpdateTeamMembershipResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTeamMembershipResponseWithDefaults

`func NewUpdateTeamMembershipResponseWithDefaults() *UpdateTeamMembershipResponse`

NewUpdateTeamMembershipResponseWithDefaults instantiates a new UpdateTeamMembershipResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *UpdateTeamMembershipResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *UpdateTeamMembershipResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *UpdateTeamMembershipResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *UpdateTeamMembershipResponse) GetData() TeamMembership`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateTeamMembershipResponse) GetDataOk() (*TeamMembership, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateTeamMembershipResponse) SetData(v TeamMembership)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


