# ServiceFrameworksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Framework** | **string** | Observed application framework, such as &#x60;express&#x60;, &#x60;fastify&#x60;, &#x60;django&#x60;, or &#x60;rails&#x60;. Tsuga detects it from OpenTelemetry instrumentation scope. Returned when framework metadata is available for the service. | 
**LastSeenAt** | **string** | Most recent time this framework was observed. Tsuga sets it from telemetry timestamps during service metadata refreshes. | 

## Methods

### NewServiceFrameworksInner

`func NewServiceFrameworksInner(framework string, lastSeenAt string, ) *ServiceFrameworksInner`

NewServiceFrameworksInner instantiates a new ServiceFrameworksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceFrameworksInnerWithDefaults

`func NewServiceFrameworksInnerWithDefaults() *ServiceFrameworksInner`

NewServiceFrameworksInnerWithDefaults instantiates a new ServiceFrameworksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFramework

`func (o *ServiceFrameworksInner) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *ServiceFrameworksInner) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *ServiceFrameworksInner) SetFramework(v string)`

SetFramework sets Framework field to given value.


### GetLastSeenAt

`func (o *ServiceFrameworksInner) GetLastSeenAt() string`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *ServiceFrameworksInner) GetLastSeenAtOk() (*string, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *ServiceFrameworksInner) SetLastSeenAt(v string)`

SetLastSeenAt sets LastSeenAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


