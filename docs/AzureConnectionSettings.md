# AzureConnectionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator identifying Azure connection settings. | 
**ClientId** | **string** | Client ID of the Azure application Tsuga authenticates as (GUID). | 
**SubscriptionId** | **string** | Azure subscription ID that Tsuga scans (GUID). | 
**TenantId** | **string** | Azure tenant ID that owns the subscription (GUID). | 

## Methods

### NewAzureConnectionSettings

`func NewAzureConnectionSettings(type_ string, clientId string, subscriptionId string, tenantId string, ) *AzureConnectionSettings`

NewAzureConnectionSettings instantiates a new AzureConnectionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureConnectionSettingsWithDefaults

`func NewAzureConnectionSettingsWithDefaults() *AzureConnectionSettings`

NewAzureConnectionSettingsWithDefaults instantiates a new AzureConnectionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AzureConnectionSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AzureConnectionSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AzureConnectionSettings) SetType(v string)`

SetType sets Type field to given value.


### GetClientId

`func (o *AzureConnectionSettings) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AzureConnectionSettings) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AzureConnectionSettings) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetSubscriptionId

`func (o *AzureConnectionSettings) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureConnectionSettings) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureConnectionSettings) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTenantId

`func (o *AzureConnectionSettings) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureConnectionSettings) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureConnectionSettings) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


