# CreateCloudAccountRequestConnectionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator identifying AWS connection settings. | 
**AccountId** | **string** | AWS account ID that Tsuga scans. | 
**ExternalId** | **string** | External ID Tsuga passes when assuming the cross-account IAM role. | 
**RoleArn** | **string** | ARN of the IAM role Tsuga assumes to read inventory in the account. | 
**ProjectId** | **string** | GCP project ID that Tsuga scans. | 
**ServiceAccountId** | **string** | Service account Tsuga impersonates to read inventory. | 
**WorkloadIdentityProvider** | **string** | Full resource name of the workload identity provider used for federated authentication, e.g. //iam.googleapis.com/projects/123456789/locations/global/workloadIdentityPools/my-pool/providers/my-provider. | 
**ClientId** | **string** | Client ID of the Azure application Tsuga authenticates as (GUID). | 
**SubscriptionId** | **string** | Azure subscription ID that Tsuga scans (GUID). | 
**TenantId** | **string** | Azure tenant ID that owns the subscription (GUID). | 

## Methods

### NewCreateCloudAccountRequestConnectionSettings

`func NewCreateCloudAccountRequestConnectionSettings(type_ string, accountId string, externalId string, roleArn string, projectId string, serviceAccountId string, workloadIdentityProvider string, clientId string, subscriptionId string, tenantId string, ) *CreateCloudAccountRequestConnectionSettings`

NewCreateCloudAccountRequestConnectionSettings instantiates a new CreateCloudAccountRequestConnectionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCloudAccountRequestConnectionSettingsWithDefaults

`func NewCreateCloudAccountRequestConnectionSettingsWithDefaults() *CreateCloudAccountRequestConnectionSettings`

NewCreateCloudAccountRequestConnectionSettingsWithDefaults instantiates a new CreateCloudAccountRequestConnectionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateCloudAccountRequestConnectionSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateCloudAccountRequestConnectionSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateCloudAccountRequestConnectionSettings) SetType(v string)`

SetType sets Type field to given value.


### GetAccountId

`func (o *CreateCloudAccountRequestConnectionSettings) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateCloudAccountRequestConnectionSettings) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateCloudAccountRequestConnectionSettings) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetExternalId

`func (o *CreateCloudAccountRequestConnectionSettings) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateCloudAccountRequestConnectionSettings) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateCloudAccountRequestConnectionSettings) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetRoleArn

`func (o *CreateCloudAccountRequestConnectionSettings) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *CreateCloudAccountRequestConnectionSettings) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *CreateCloudAccountRequestConnectionSettings) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetProjectId

`func (o *CreateCloudAccountRequestConnectionSettings) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateCloudAccountRequestConnectionSettings) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateCloudAccountRequestConnectionSettings) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetServiceAccountId

`func (o *CreateCloudAccountRequestConnectionSettings) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *CreateCloudAccountRequestConnectionSettings) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *CreateCloudAccountRequestConnectionSettings) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.


### GetWorkloadIdentityProvider

`func (o *CreateCloudAccountRequestConnectionSettings) GetWorkloadIdentityProvider() string`

GetWorkloadIdentityProvider returns the WorkloadIdentityProvider field if non-nil, zero value otherwise.

### GetWorkloadIdentityProviderOk

`func (o *CreateCloudAccountRequestConnectionSettings) GetWorkloadIdentityProviderOk() (*string, bool)`

GetWorkloadIdentityProviderOk returns a tuple with the WorkloadIdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadIdentityProvider

`func (o *CreateCloudAccountRequestConnectionSettings) SetWorkloadIdentityProvider(v string)`

SetWorkloadIdentityProvider sets WorkloadIdentityProvider field to given value.


### GetClientId

`func (o *CreateCloudAccountRequestConnectionSettings) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateCloudAccountRequestConnectionSettings) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateCloudAccountRequestConnectionSettings) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetSubscriptionId

`func (o *CreateCloudAccountRequestConnectionSettings) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateCloudAccountRequestConnectionSettings) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateCloudAccountRequestConnectionSettings) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTenantId

`func (o *CreateCloudAccountRequestConnectionSettings) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateCloudAccountRequestConnectionSettings) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateCloudAccountRequestConnectionSettings) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


