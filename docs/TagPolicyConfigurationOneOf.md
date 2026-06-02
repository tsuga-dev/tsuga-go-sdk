# TagPolicyConfigurationOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**AssetTypes** | **[]string** |  | 
**ShouldInsertWarning** | **bool** |  | 
**DropSample** | Pointer to **float32** |  | [optional] 

## Methods

### NewTagPolicyConfigurationOneOf

`func NewTagPolicyConfigurationOneOf(type_ string, assetTypes []string, shouldInsertWarning bool, ) *TagPolicyConfigurationOneOf`

NewTagPolicyConfigurationOneOf instantiates a new TagPolicyConfigurationOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagPolicyConfigurationOneOfWithDefaults

`func NewTagPolicyConfigurationOneOfWithDefaults() *TagPolicyConfigurationOneOf`

NewTagPolicyConfigurationOneOfWithDefaults instantiates a new TagPolicyConfigurationOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TagPolicyConfigurationOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TagPolicyConfigurationOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TagPolicyConfigurationOneOf) SetType(v string)`

SetType sets Type field to given value.


### GetAssetTypes

`func (o *TagPolicyConfigurationOneOf) GetAssetTypes() []string`

GetAssetTypes returns the AssetTypes field if non-nil, zero value otherwise.

### GetAssetTypesOk

`func (o *TagPolicyConfigurationOneOf) GetAssetTypesOk() (*[]string, bool)`

GetAssetTypesOk returns a tuple with the AssetTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTypes

`func (o *TagPolicyConfigurationOneOf) SetAssetTypes(v []string)`

SetAssetTypes sets AssetTypes field to given value.


### GetShouldInsertWarning

`func (o *TagPolicyConfigurationOneOf) GetShouldInsertWarning() bool`

GetShouldInsertWarning returns the ShouldInsertWarning field if non-nil, zero value otherwise.

### GetShouldInsertWarningOk

`func (o *TagPolicyConfigurationOneOf) GetShouldInsertWarningOk() (*bool, bool)`

GetShouldInsertWarningOk returns a tuple with the ShouldInsertWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldInsertWarning

`func (o *TagPolicyConfigurationOneOf) SetShouldInsertWarning(v bool)`

SetShouldInsertWarning sets ShouldInsertWarning field to given value.


### GetDropSample

`func (o *TagPolicyConfigurationOneOf) GetDropSample() float32`

GetDropSample returns the DropSample field if non-nil, zero value otherwise.

### GetDropSampleOk

`func (o *TagPolicyConfigurationOneOf) GetDropSampleOk() (*float32, bool)`

GetDropSampleOk returns a tuple with the DropSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropSample

`func (o *TagPolicyConfigurationOneOf) SetDropSample(v float32)`

SetDropSample sets DropSample field to given value.

### HasDropSample

`func (o *TagPolicyConfigurationOneOf) HasDropSample() bool`

HasDropSample returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


