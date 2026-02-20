# CreateTagPolicyRequestConfigurationOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**AssetTypes** | **[]string** |  | 
**ShouldInsertWarning** | **bool** |  | 
**DropSample** | Pointer to **float32** |  | [optional] 

## Methods

### NewCreateTagPolicyRequestConfigurationOneOf

`func NewCreateTagPolicyRequestConfigurationOneOf(type_ string, assetTypes []string, shouldInsertWarning bool, ) *CreateTagPolicyRequestConfigurationOneOf`

NewCreateTagPolicyRequestConfigurationOneOf instantiates a new CreateTagPolicyRequestConfigurationOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTagPolicyRequestConfigurationOneOfWithDefaults

`func NewCreateTagPolicyRequestConfigurationOneOfWithDefaults() *CreateTagPolicyRequestConfigurationOneOf`

NewCreateTagPolicyRequestConfigurationOneOfWithDefaults instantiates a new CreateTagPolicyRequestConfigurationOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateTagPolicyRequestConfigurationOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTagPolicyRequestConfigurationOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTagPolicyRequestConfigurationOneOf) SetType(v string)`

SetType sets Type field to given value.


### GetAssetTypes

`func (o *CreateTagPolicyRequestConfigurationOneOf) GetAssetTypes() []string`

GetAssetTypes returns the AssetTypes field if non-nil, zero value otherwise.

### GetAssetTypesOk

`func (o *CreateTagPolicyRequestConfigurationOneOf) GetAssetTypesOk() (*[]string, bool)`

GetAssetTypesOk returns a tuple with the AssetTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTypes

`func (o *CreateTagPolicyRequestConfigurationOneOf) SetAssetTypes(v []string)`

SetAssetTypes sets AssetTypes field to given value.


### GetShouldInsertWarning

`func (o *CreateTagPolicyRequestConfigurationOneOf) GetShouldInsertWarning() bool`

GetShouldInsertWarning returns the ShouldInsertWarning field if non-nil, zero value otherwise.

### GetShouldInsertWarningOk

`func (o *CreateTagPolicyRequestConfigurationOneOf) GetShouldInsertWarningOk() (*bool, bool)`

GetShouldInsertWarningOk returns a tuple with the ShouldInsertWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldInsertWarning

`func (o *CreateTagPolicyRequestConfigurationOneOf) SetShouldInsertWarning(v bool)`

SetShouldInsertWarning sets ShouldInsertWarning field to given value.


### GetDropSample

`func (o *CreateTagPolicyRequestConfigurationOneOf) GetDropSample() float32`

GetDropSample returns the DropSample field if non-nil, zero value otherwise.

### GetDropSampleOk

`func (o *CreateTagPolicyRequestConfigurationOneOf) GetDropSampleOk() (*float32, bool)`

GetDropSampleOk returns a tuple with the DropSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropSample

`func (o *CreateTagPolicyRequestConfigurationOneOf) SetDropSample(v float32)`

SetDropSample sets DropSample field to given value.

### HasDropSample

`func (o *CreateTagPolicyRequestConfigurationOneOf) HasDropSample() bool`

HasDropSample returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


