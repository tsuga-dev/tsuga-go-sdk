# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of this service. Use it with service detail and graph endpoints. Absent until the hourly service catalog task records the service. | [optional] 
**ServiceName** | **string** | OpenTelemetry service name observed in telemetry. | 
**ServiceNamespace** | Pointer to **string** | OpenTelemetry service namespace observed in telemetry. | [optional] 
**Env** | **string** | Environment value observed for this service. | 
**Teams** | [**[]ServiceTeamsInner**](ServiceTeamsInner.md) | Team names observed on telemetry for this service. | 
**Versions** | [**[]ServiceVersionsInner**](ServiceVersionsInner.md) | Service version values observed in telemetry, including first and last seen timestamps. | 
**Languages** | [**[]ServiceLanguagesInner**](ServiceLanguagesInner.md) | Programming languages observed for this service. | 
**ServerProtocols** | [**[]ServiceServerProtocolsInner**](ServiceServerProtocolsInner.md) | API protocols this service exposes (from server-kind spans) | 
**Frameworks** | [**[]ServiceFrameworksInner**](ServiceFrameworksInner.md) | Application frameworks detected from OTel instrumentation scope | 
**CloudPlatforms** | Pointer to [**[]ServiceCloudPlatformsInner**](ServiceCloudPlatformsInner.md) | Cloud platforms this service is running on. Absent or empty when Tsuga has not observed platform metadata for the service; treat both the same. | [optional] 
**CloudProviders** | Pointer to [**[]ServiceCloudProvidersInner**](ServiceCloudProvidersInner.md) | Cloud providers this service is running on. Absent or empty when Tsuga has not observed provider metadata for the service; treat both the same. | [optional] 
**LastSeenAt** | Pointer to **string** | Most recent time Tsuga observed telemetry for this service. Absent until the hourly service catalog task records the service. | [optional] 
**FirstSeenAt** | Pointer to **string** | First time Tsuga observed telemetry for this service. Absent until the hourly service catalog task records the service. | [optional] 
**TraceRequestRate** | Pointer to **float32** | Requests per second observed for this service over the queried window, counting calls of its primary span kind. The primary span kind is the busiest of &#x60;SPAN_KIND_SERVER&#x60;, &#x60;SPAN_KIND_PRODUCER&#x60;, and &#x60;SPAN_KIND_CONSUMER&#x60; in the service’s busiest env, or &#x60;SPAN_KIND_INTERNAL&#x60; when the service has none of those; &#x60;SPAN_KIND_CLIENT&#x60; never counts. Every env of the service uses that same span kind. Omitted when the trace query failed, which is distinct from &#x60;0&#x60; (the service is quiet). | [optional] 
**TraceErrorRate** | Pointer to **float32** | Percentage of this service’s primary-span-kind calls (see &#x60;traceRequestRate&#x60;) over the queried window that ended with error status. Omitted when the trace query failed. | [optional] 

## Methods

### NewService

`func NewService(serviceName string, env string, teams []ServiceTeamsInner, versions []ServiceVersionsInner, languages []ServiceLanguagesInner, serverProtocols []ServiceServerProtocolsInner, frameworks []ServiceFrameworksInner, ) *Service`

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

### HasId

`func (o *Service) HasId() bool`

HasId returns a boolean if a field has been set.

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


### GetServerProtocols

`func (o *Service) GetServerProtocols() []ServiceServerProtocolsInner`

GetServerProtocols returns the ServerProtocols field if non-nil, zero value otherwise.

### GetServerProtocolsOk

`func (o *Service) GetServerProtocolsOk() (*[]ServiceServerProtocolsInner, bool)`

GetServerProtocolsOk returns a tuple with the ServerProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerProtocols

`func (o *Service) SetServerProtocols(v []ServiceServerProtocolsInner)`

SetServerProtocols sets ServerProtocols field to given value.


### GetFrameworks

`func (o *Service) GetFrameworks() []ServiceFrameworksInner`

GetFrameworks returns the Frameworks field if non-nil, zero value otherwise.

### GetFrameworksOk

`func (o *Service) GetFrameworksOk() (*[]ServiceFrameworksInner, bool)`

GetFrameworksOk returns a tuple with the Frameworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworks

`func (o *Service) SetFrameworks(v []ServiceFrameworksInner)`

SetFrameworks sets Frameworks field to given value.


### GetCloudPlatforms

`func (o *Service) GetCloudPlatforms() []ServiceCloudPlatformsInner`

GetCloudPlatforms returns the CloudPlatforms field if non-nil, zero value otherwise.

### GetCloudPlatformsOk

`func (o *Service) GetCloudPlatformsOk() (*[]ServiceCloudPlatformsInner, bool)`

GetCloudPlatformsOk returns a tuple with the CloudPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPlatforms

`func (o *Service) SetCloudPlatforms(v []ServiceCloudPlatformsInner)`

SetCloudPlatforms sets CloudPlatforms field to given value.

### HasCloudPlatforms

`func (o *Service) HasCloudPlatforms() bool`

HasCloudPlatforms returns a boolean if a field has been set.

### GetCloudProviders

`func (o *Service) GetCloudProviders() []ServiceCloudProvidersInner`

GetCloudProviders returns the CloudProviders field if non-nil, zero value otherwise.

### GetCloudProvidersOk

`func (o *Service) GetCloudProvidersOk() (*[]ServiceCloudProvidersInner, bool)`

GetCloudProvidersOk returns a tuple with the CloudProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviders

`func (o *Service) SetCloudProviders(v []ServiceCloudProvidersInner)`

SetCloudProviders sets CloudProviders field to given value.

### HasCloudProviders

`func (o *Service) HasCloudProviders() bool`

HasCloudProviders returns a boolean if a field has been set.

### GetLastSeenAt

`func (o *Service) GetLastSeenAt() string`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *Service) GetLastSeenAtOk() (*string, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *Service) SetLastSeenAt(v string)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *Service) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### GetFirstSeenAt

`func (o *Service) GetFirstSeenAt() string`

GetFirstSeenAt returns the FirstSeenAt field if non-nil, zero value otherwise.

### GetFirstSeenAtOk

`func (o *Service) GetFirstSeenAtOk() (*string, bool)`

GetFirstSeenAtOk returns a tuple with the FirstSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenAt

`func (o *Service) SetFirstSeenAt(v string)`

SetFirstSeenAt sets FirstSeenAt field to given value.

### HasFirstSeenAt

`func (o *Service) HasFirstSeenAt() bool`

HasFirstSeenAt returns a boolean if a field has been set.

### GetTraceRequestRate

`func (o *Service) GetTraceRequestRate() float32`

GetTraceRequestRate returns the TraceRequestRate field if non-nil, zero value otherwise.

### GetTraceRequestRateOk

`func (o *Service) GetTraceRequestRateOk() (*float32, bool)`

GetTraceRequestRateOk returns a tuple with the TraceRequestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRequestRate

`func (o *Service) SetTraceRequestRate(v float32)`

SetTraceRequestRate sets TraceRequestRate field to given value.

### HasTraceRequestRate

`func (o *Service) HasTraceRequestRate() bool`

HasTraceRequestRate returns a boolean if a field has been set.

### GetTraceErrorRate

`func (o *Service) GetTraceErrorRate() float32`

GetTraceErrorRate returns the TraceErrorRate field if non-nil, zero value otherwise.

### GetTraceErrorRateOk

`func (o *Service) GetTraceErrorRateOk() (*float32, bool)`

GetTraceErrorRateOk returns a tuple with the TraceErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceErrorRate

`func (o *Service) SetTraceErrorRate(v float32)`

SetTraceErrorRate sets TraceErrorRate field to given value.

### HasTraceErrorRate

`func (o *Service) HasTraceErrorRate() bool`

HasTraceErrorRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


