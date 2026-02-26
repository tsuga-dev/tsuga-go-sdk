# ServiceTeamsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Team** | **string** |  | 
**LastSeenAt** | **time.Time** |  | 

## Methods

### NewServiceTeamsInner

`func NewServiceTeamsInner(team string, lastSeenAt time.Time, ) *ServiceTeamsInner`

NewServiceTeamsInner instantiates a new ServiceTeamsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTeamsInnerWithDefaults

`func NewServiceTeamsInnerWithDefaults() *ServiceTeamsInner`

NewServiceTeamsInnerWithDefaults instantiates a new ServiceTeamsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeam

`func (o *ServiceTeamsInner) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ServiceTeamsInner) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ServiceTeamsInner) SetTeam(v string)`

SetTeam sets Team field to given value.


### GetLastSeenAt

`func (o *ServiceTeamsInner) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *ServiceTeamsInner) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *ServiceTeamsInner) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


