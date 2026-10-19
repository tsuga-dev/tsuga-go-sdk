# AwsConnectionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator identifying AWS connection settings. | 
**AccountId** | **string** | AWS account ID that Tsuga scans. | 
**ExternalId** | **string** | External ID Tsuga passes when assuming the cross-account IAM role. | 
**RoleArn** | **string** | ARN of the IAM role Tsuga assumes to read inventory in the account. | 

## Methods

### NewAwsConnectionSettings

`func NewAwsConnectionSettings(type_ string, accountId string, externalId string, roleArn string, ) *AwsConnectionSettings`

NewAwsConnectionSettings instantiates a new AwsConnectionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsConnectionSettingsWithDefaults

`func NewAwsConnectionSettingsWithDefaults() *AwsConnectionSettings`

NewAwsConnectionSettingsWithDefaults instantiates a new AwsConnectionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AwsConnectionSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwsConnectionSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwsConnectionSettings) SetType(v string)`

SetType sets Type field to given value.


### GetAccountId

`func (o *AwsConnectionSettings) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AwsConnectionSettings) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AwsConnectionSettings) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetExternalId

`func (o *AwsConnectionSettings) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AwsConnectionSettings) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AwsConnectionSettings) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetRoleArn

`func (o *AwsConnectionSettings) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AwsConnectionSettings) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AwsConnectionSettings) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


