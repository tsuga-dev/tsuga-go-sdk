# ServiceCloudProvidersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | Observed cloud provider, read from the &#x60;cloud.provider&#x60; resource attribute, such as &#x60;aws&#x60;, &#x60;gcp&#x60;, or &#x60;azure&#x60;. The emitting service or collector sets it. | 
**LastSeenAt** | **string** | Most recent time this provider was observed. Tsuga sets it from telemetry timestamps during service metadata refreshes. | 

## Methods

### NewServiceCloudProvidersInner

`func NewServiceCloudProvidersInner(provider string, lastSeenAt string, ) *ServiceCloudProvidersInner`

NewServiceCloudProvidersInner instantiates a new ServiceCloudProvidersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCloudProvidersInnerWithDefaults

`func NewServiceCloudProvidersInnerWithDefaults() *ServiceCloudProvidersInner`

NewServiceCloudProvidersInnerWithDefaults instantiates a new ServiceCloudProvidersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *ServiceCloudProvidersInner) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ServiceCloudProvidersInner) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ServiceCloudProvidersInner) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetLastSeenAt

`func (o *ServiceCloudProvidersInner) GetLastSeenAt() string`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *ServiceCloudProvidersInner) GetLastSeenAtOk() (*string, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *ServiceCloudProvidersInner) SetLastSeenAt(v string)`

SetLastSeenAt sets LastSeenAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


