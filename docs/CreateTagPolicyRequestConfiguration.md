# CreateTagPolicyRequestConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**AssetTypes** | **[]string** |  | 
**ShouldInsertWarning** | **bool** |  | 
**DropSample** | Pointer to **float32** |  | [optional] 

## Methods

### NewCreateTagPolicyRequestConfiguration

`func NewCreateTagPolicyRequestConfiguration(type_ string, assetTypes []string, shouldInsertWarning bool, ) *CreateTagPolicyRequestConfiguration`

NewCreateTagPolicyRequestConfiguration instantiates a new CreateTagPolicyRequestConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTagPolicyRequestConfigurationWithDefaults

`func NewCreateTagPolicyRequestConfigurationWithDefaults() *CreateTagPolicyRequestConfiguration`

NewCreateTagPolicyRequestConfigurationWithDefaults instantiates a new CreateTagPolicyRequestConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateTagPolicyRequestConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTagPolicyRequestConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTagPolicyRequestConfiguration) SetType(v string)`

SetType sets Type field to given value.


### GetAssetTypes

`func (o *CreateTagPolicyRequestConfiguration) GetAssetTypes() []string`

GetAssetTypes returns the AssetTypes field if non-nil, zero value otherwise.

### GetAssetTypesOk

`func (o *CreateTagPolicyRequestConfiguration) GetAssetTypesOk() (*[]string, bool)`

GetAssetTypesOk returns a tuple with the AssetTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTypes

`func (o *CreateTagPolicyRequestConfiguration) SetAssetTypes(v []string)`

SetAssetTypes sets AssetTypes field to given value.


### GetShouldInsertWarning

`func (o *CreateTagPolicyRequestConfiguration) GetShouldInsertWarning() bool`

GetShouldInsertWarning returns the ShouldInsertWarning field if non-nil, zero value otherwise.

### GetShouldInsertWarningOk

`func (o *CreateTagPolicyRequestConfiguration) GetShouldInsertWarningOk() (*bool, bool)`

GetShouldInsertWarningOk returns a tuple with the ShouldInsertWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldInsertWarning

`func (o *CreateTagPolicyRequestConfiguration) SetShouldInsertWarning(v bool)`

SetShouldInsertWarning sets ShouldInsertWarning field to given value.


### GetDropSample

`func (o *CreateTagPolicyRequestConfiguration) GetDropSample() float32`

GetDropSample returns the DropSample field if non-nil, zero value otherwise.

### GetDropSampleOk

`func (o *CreateTagPolicyRequestConfiguration) GetDropSampleOk() (*float32, bool)`

GetDropSampleOk returns a tuple with the DropSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropSample

`func (o *CreateTagPolicyRequestConfiguration) SetDropSample(v float32)`

SetDropSample sets DropSample field to given value.

### HasDropSample

`func (o *CreateTagPolicyRequestConfiguration) HasDropSample() bool`

HasDropSample returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


