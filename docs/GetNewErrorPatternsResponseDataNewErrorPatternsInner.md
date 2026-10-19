# GetNewErrorPatternsResponseDataNewErrorPatternsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExampleLog** | Pointer to [**SearchLogsResponseDataLogsInner**](SearchLogsResponseDataLogsInner.md) |  | [optional] 
**Team** | **string** | Team associated with the new error pattern. | 
**Env** | **string** | Environment associated with the new error pattern. | 
**Service** | **string** | Service associated with the new error pattern. | 
**FirstSeen** | **float32** | Unix timestamp in seconds when this pattern was first seen. | 
**LastSeen** | **float32** | Unix timestamp in seconds when this pattern was last seen. | 

## Methods

### NewGetNewErrorPatternsResponseDataNewErrorPatternsInner

`func NewGetNewErrorPatternsResponseDataNewErrorPatternsInner(team string, env string, service string, firstSeen float32, lastSeen float32, ) *GetNewErrorPatternsResponseDataNewErrorPatternsInner`

NewGetNewErrorPatternsResponseDataNewErrorPatternsInner instantiates a new GetNewErrorPatternsResponseDataNewErrorPatternsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNewErrorPatternsResponseDataNewErrorPatternsInnerWithDefaults

`func NewGetNewErrorPatternsResponseDataNewErrorPatternsInnerWithDefaults() *GetNewErrorPatternsResponseDataNewErrorPatternsInner`

NewGetNewErrorPatternsResponseDataNewErrorPatternsInnerWithDefaults instantiates a new GetNewErrorPatternsResponseDataNewErrorPatternsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExampleLog

`func (o *GetNewErrorPatternsResponseDataNewErrorPatternsInner) GetExampleLog() SearchLogsResponseDataLogsInner`

GetExampleLog returns the ExampleLog field if non-nil, zero value otherwise.

### GetExampleLogOk

`func (o *GetNewErrorPatternsResponseDataNewErrorPatternsInner) GetExampleLogOk() (*SearchLogsResponseDataLogsInner, bool)`

GetExampleLogOk returns a tuple with the ExampleLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExampleLog

`func (o *GetNewErrorPatternsResponseDataNewErrorPatternsInner) SetExampleLog(v SearchLogsResponseDataLogsInner)`

SetExampleLog sets ExampleLog field to given value.

### HasExampleLog

`func (o *GetNewErrorPatternsResponseDataNewErrorPatternsInner) HasExampleLog() bool`

HasExampleLog returns a boolean if a field has been set.

### GetTeam

`func (o *GetNewErrorPatternsResponseDataNewErrorPatternsInner) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *GetNewErrorPatternsResponseDataNewErrorPatternsInner) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *GetNewErrorPatternsResponseDataNewErrorPatternsInner) SetTeam(v string)`

SetTeam sets Team field to given value.


### GetEnv

`func (o *GetNewErrorPatternsResponseDataNewErrorPatternsInner) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *GetNewErrorPatternsResponseDataNewErrorPatternsInner) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *GetNewErrorPatternsResponseDataNewErrorPatternsInner) SetEnv(v string)`

SetEnv sets Env field to given value.


### GetService

`func (o *GetNewErrorPatternsResponseDataNewErrorPatternsInner) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GetNewErrorPatternsResponseDataNewErrorPatternsInner) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GetNewErrorPatternsResponseDataNewErrorPatternsInner) SetService(v string)`

SetService sets Service field to given value.


### GetFirstSeen

`func (o *GetNewErrorPatternsResponseDataNewErrorPatternsInner) GetFirstSeen() float32`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *GetNewErrorPatternsResponseDataNewErrorPatternsInner) GetFirstSeenOk() (*float32, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *GetNewErrorPatternsResponseDataNewErrorPatternsInner) SetFirstSeen(v float32)`

SetFirstSeen sets FirstSeen field to given value.


### GetLastSeen

`func (o *GetNewErrorPatternsResponseDataNewErrorPatternsInner) GetLastSeen() float32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *GetNewErrorPatternsResponseDataNewErrorPatternsInner) GetLastSeenOk() (*float32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *GetNewErrorPatternsResponseDataNewErrorPatternsInner) SetLastSeen(v float32)`

SetLastSeen sets LastSeen field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


