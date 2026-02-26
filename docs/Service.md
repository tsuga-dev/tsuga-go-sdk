# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ServiceName** | **string** |  | 
**ServiceNamespace** | Pointer to **string** |  | [optional] 
**Env** | Pointer to **string** |  | [optional] 
**Teams** | [**[]ServiceTeamsInner**](ServiceTeamsInner.md) |  | 
**Versions** | [**[]ServiceVersionsInner**](ServiceVersionsInner.md) |  | 
**Languages** | [**[]ServiceLanguagesInner**](ServiceLanguagesInner.md) |  | 
**LastSeenAt** | **time.Time** |  | 
**FirstSeenAt** | **time.Time** |  | 
**Sources** | **[]string** |  | 
**LogsCount24h** | **float32** |  | 
**ErrorLogsCount24h** | **float32** |  | 
**TracesCount24h** | **float32** |  | 
**ErrorTracesCount24h** | **float32** |  | 

## Methods

### NewService

`func NewService(id string, serviceName string, teams []ServiceTeamsInner, versions []ServiceVersionsInner, languages []ServiceLanguagesInner, lastSeenAt time.Time, firstSeenAt time.Time, sources []string, logsCount24h float32, errorLogsCount24h float32, tracesCount24h float32, errorTracesCount24h float32, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Service) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v string)`

SetId sets Id field to given value.


### GetServiceName

`func (o *Service) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *Service) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *Service) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetServiceNamespace

`func (o *Service) GetServiceNamespace() string`

GetServiceNamespace returns the ServiceNamespace field if non-nil, zero value otherwise.

### GetServiceNamespaceOk

`func (o *Service) GetServiceNamespaceOk() (*string, bool)`

GetServiceNamespaceOk returns a tuple with the ServiceNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNamespace

`func (o *Service) SetServiceNamespace(v string)`

SetServiceNamespace sets ServiceNamespace field to given value.

### HasServiceNamespace

`func (o *Service) HasServiceNamespace() bool`

HasServiceNamespace returns a boolean if a field has been set.

### GetEnv

`func (o *Service) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *Service) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *Service) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *Service) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetTeams

`func (o *Service) GetTeams() []ServiceTeamsInner`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *Service) GetTeamsOk() (*[]ServiceTeamsInner, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *Service) SetTeams(v []ServiceTeamsInner)`

SetTeams sets Teams field to given value.


### GetVersions

`func (o *Service) GetVersions() []ServiceVersionsInner`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *Service) GetVersionsOk() (*[]ServiceVersionsInner, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *Service) SetVersions(v []ServiceVersionsInner)`

SetVersions sets Versions field to given value.


### GetLanguages

`func (o *Service) GetLanguages() []ServiceLanguagesInner`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *Service) GetLanguagesOk() (*[]ServiceLanguagesInner, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *Service) SetLanguages(v []ServiceLanguagesInner)`

SetLanguages sets Languages field to given value.


### GetLastSeenAt

`func (o *Service) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *Service) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *Service) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.


### GetFirstSeenAt

`func (o *Service) GetFirstSeenAt() time.Time`

GetFirstSeenAt returns the FirstSeenAt field if non-nil, zero value otherwise.

### GetFirstSeenAtOk

`func (o *Service) GetFirstSeenAtOk() (*time.Time, bool)`

GetFirstSeenAtOk returns a tuple with the FirstSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenAt

`func (o *Service) SetFirstSeenAt(v time.Time)`

SetFirstSeenAt sets FirstSeenAt field to given value.


### GetSources

`func (o *Service) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *Service) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *Service) SetSources(v []string)`

SetSources sets Sources field to given value.


### GetLogsCount24h

`func (o *Service) GetLogsCount24h() float32`

GetLogsCount24h returns the LogsCount24h field if non-nil, zero value otherwise.

### GetLogsCount24hOk

`func (o *Service) GetLogsCount24hOk() (*float32, bool)`

GetLogsCount24hOk returns a tuple with the LogsCount24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsCount24h

`func (o *Service) SetLogsCount24h(v float32)`

SetLogsCount24h sets LogsCount24h field to given value.


### GetErrorLogsCount24h

`func (o *Service) GetErrorLogsCount24h() float32`

GetErrorLogsCount24h returns the ErrorLogsCount24h field if non-nil, zero value otherwise.

### GetErrorLogsCount24hOk

`func (o *Service) GetErrorLogsCount24hOk() (*float32, bool)`

GetErrorLogsCount24hOk returns a tuple with the ErrorLogsCount24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorLogsCount24h

`func (o *Service) SetErrorLogsCount24h(v float32)`

SetErrorLogsCount24h sets ErrorLogsCount24h field to given value.


### GetTracesCount24h

`func (o *Service) GetTracesCount24h() float32`

GetTracesCount24h returns the TracesCount24h field if non-nil, zero value otherwise.

### GetTracesCount24hOk

`func (o *Service) GetTracesCount24hOk() (*float32, bool)`

GetTracesCount24hOk returns a tuple with the TracesCount24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracesCount24h

`func (o *Service) SetTracesCount24h(v float32)`

SetTracesCount24h sets TracesCount24h field to given value.


### GetErrorTracesCount24h

`func (o *Service) GetErrorTracesCount24h() float32`

GetErrorTracesCount24h returns the ErrorTracesCount24h field if non-nil, zero value otherwise.

### GetErrorTracesCount24hOk

`func (o *Service) GetErrorTracesCount24hOk() (*float32, bool)`

GetErrorTracesCount24hOk returns a tuple with the ErrorTracesCount24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorTracesCount24h

`func (o *Service) SetErrorTracesCount24h(v float32)`

SetErrorTracesCount24h sets ErrorTracesCount24h field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


