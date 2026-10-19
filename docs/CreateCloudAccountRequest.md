# CreateCloudAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudType** | **string** | Cloud provider of the account. | 
**CloudAccountId** | **string** | Cloud-native account identifier: an AWS account ID, GCP project ID, or Azure subscription ID. | 
**AccountFriendlyName** | Pointer to **string** | Human-readable name for the account shown in Tsuga. | [optional] 
**ConnectionSettings** | [**CreateCloudAccountRequestConnectionSettings**](CreateCloudAccountRequestConnectionSettings.md) |  | 

## Methods

### NewCreateCloudAccountRequest

`func NewCreateCloudAccountRequest(cloudType string, cloudAccountId string, connectionSettings CreateCloudAccountRequestConnectionSettings, ) *CreateCloudAccountRequest`

NewCreateCloudAccountRequest instantiates a new CreateCloudAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCloudAccountRequestWithDefaults

`func NewCreateCloudAccountRequestWithDefaults() *CreateCloudAccountRequest`

NewCreateCloudAccountRequestWithDefaults instantiates a new CreateCloudAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudType

`func (o *CreateCloudAccountRequest) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *CreateCloudAccountRequest) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *CreateCloudAccountRequest) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.


### GetCloudAccountId

`func (o *CreateCloudAccountRequest) GetCloudAccountId() string`

GetCloudAccountId returns the CloudAccountId field if non-nil, zero value otherwise.

### GetCloudAccountIdOk

`func (o *CreateCloudAccountRequest) GetCloudAccountIdOk() (*string, bool)`

GetCloudAccountIdOk returns a tuple with the CloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccountId

`func (o *CreateCloudAccountRequest) SetCloudAccountId(v string)`

SetCloudAccountId sets CloudAccountId field to given value.


### GetAccountFriendlyName

`func (o *CreateCloudAccountRequest) GetAccountFriendlyName() string`

GetAccountFriendlyName returns the AccountFriendlyName field if non-nil, zero value otherwise.

### GetAccountFriendlyNameOk

`func (o *CreateCloudAccountRequest) GetAccountFriendlyNameOk() (*string, bool)`

GetAccountFriendlyNameOk returns a tuple with the AccountFriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFriendlyName

`func (o *CreateCloudAccountRequest) SetAccountFriendlyName(v string)`

SetAccountFriendlyName sets AccountFriendlyName field to given value.

### HasAccountFriendlyName

`func (o *CreateCloudAccountRequest) HasAccountFriendlyName() bool`

HasAccountFriendlyName returns a boolean if a field has been set.

### GetConnectionSettings

`func (o *CreateCloudAccountRequest) GetConnectionSettings() CreateCloudAccountRequestConnectionSettings`

GetConnectionSettings returns the ConnectionSettings field if non-nil, zero value otherwise.

### GetConnectionSettingsOk

`func (o *CreateCloudAccountRequest) GetConnectionSettingsOk() (*CreateCloudAccountRequestConnectionSettings, bool)`

GetConnectionSettingsOk returns a tuple with the ConnectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSettings

`func (o *CreateCloudAccountRequest) SetConnectionSettings(v CreateCloudAccountRequestConnectionSettings)`

SetConnectionSettings sets ConnectionSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


