# TsugaAssetTagPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Policy applies to Tsuga-managed resources such as dashboards, routes, and keys. | 
**AssetTypes** | **[]string** | Tsuga-managed asset types covered by the policy. An empty array means all supported asset types. | 

## Methods

### NewTsugaAssetTagPolicy

`func NewTsugaAssetTagPolicy(type_ string, assetTypes []string, ) *TsugaAssetTagPolicy`

NewTsugaAssetTagPolicy instantiates a new TsugaAssetTagPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTsugaAssetTagPolicyWithDefaults

`func NewTsugaAssetTagPolicyWithDefaults() *TsugaAssetTagPolicy`

NewTsugaAssetTagPolicyWithDefaults instantiates a new TsugaAssetTagPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TsugaAssetTagPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TsugaAssetTagPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TsugaAssetTagPolicy) SetType(v string)`

SetType sets Type field to given value.


### GetAssetTypes

`func (o *TsugaAssetTagPolicy) GetAssetTypes() []string`

GetAssetTypes returns the AssetTypes field if non-nil, zero value otherwise.

### GetAssetTypesOk

`func (o *TsugaAssetTagPolicy) GetAssetTypesOk() (*[]string, bool)`

GetAssetTypesOk returns a tuple with the AssetTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTypes

`func (o *TsugaAssetTagPolicy) SetAssetTypes(v []string)`

SetAssetTypes sets AssetTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


