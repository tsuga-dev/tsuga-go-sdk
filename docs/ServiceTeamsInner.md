# ServiceTeamsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Team** | **string** | Team name observed on telemetry for this service. The emitting service sets it through telemetry attributes. Returned when Tsuga has observed a team value for the service; metadata refreshes weekly or when a new service version appears. For a service the service catalog has not recorded yet, it comes from the trace span metrics of the queried window. | 

## Methods

### NewServiceTeamsInner

`func NewServiceTeamsInner(team string, ) *ServiceTeamsInner`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


