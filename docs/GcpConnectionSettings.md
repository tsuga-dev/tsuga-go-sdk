# GcpConnectionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator identifying GCP connection settings. | 
**ProjectId** | **string** | GCP project ID that Tsuga scans. | 
**ServiceAccountId** | **string** | Service account Tsuga impersonates to read inventory. | 
**WorkloadIdentityProvider** | **string** | Full resource name of the workload identity provider used for federated authentication, e.g. //iam.googleapis.com/projects/123456789/locations/global/workloadIdentityPools/my-pool/providers/my-provider. | 

## Methods

### NewGcpConnectionSettings

`func NewGcpConnectionSettings(type_ string, projectId string, serviceAccountId string, workloadIdentityProvider string, ) *GcpConnectionSettings`

NewGcpConnectionSettings instantiates a new GcpConnectionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpConnectionSettingsWithDefaults

`func NewGcpConnectionSettingsWithDefaults() *GcpConnectionSettings`

NewGcpConnectionSettingsWithDefaults instantiates a new GcpConnectionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GcpConnectionSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GcpConnectionSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GcpConnectionSettings) SetType(v string)`

SetType sets Type field to given value.


### GetProjectId

`func (o *GcpConnectionSettings) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GcpConnectionSettings) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GcpConnectionSettings) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetServiceAccountId

`func (o *GcpConnectionSettings) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *GcpConnectionSettings) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *GcpConnectionSettings) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.


### GetWorkloadIdentityProvider

`func (o *GcpConnectionSettings) GetWorkloadIdentityProvider() string`

GetWorkloadIdentityProvider returns the WorkloadIdentityProvider field if non-nil, zero value otherwise.

### GetWorkloadIdentityProviderOk

`func (o *GcpConnectionSettings) GetWorkloadIdentityProviderOk() (*string, bool)`

GetWorkloadIdentityProviderOk returns a tuple with the WorkloadIdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadIdentityProvider

`func (o *GcpConnectionSettings) SetWorkloadIdentityProvider(v string)`

SetWorkloadIdentityProvider sets WorkloadIdentityProvider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


