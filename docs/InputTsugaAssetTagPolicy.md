# InputTsugaAssetTagPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Policy applies to Tsuga-managed resources such as dashboards, routes, and keys. | 
**AssetTypes** | **[]string** | Tsuga-managed asset types covered by the policy. An empty array means all supported asset types. | 

## Methods

### NewInputTsugaAssetTagPolicy

`func NewInputTsugaAssetTagPolicy(type_ string, assetTypes []string, ) *InputTsugaAssetTagPolicy`

NewInputTsugaAssetTagPolicy instantiates a new InputTsugaAssetTagPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputTsugaAssetTagPolicyWithDefaults

`func NewInputTsugaAssetTagPolicyWithDefaults() *InputTsugaAssetTagPolicy`

NewInputTsugaAssetTagPolicyWithDefaults instantiates a new InputTsugaAssetTagPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputTsugaAssetTagPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputTsugaAssetTagPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputTsugaAssetTagPolicy) SetType(v string)`

SetType sets Type field to given value.


### GetAssetTypes

`func (o *InputTsugaAssetTagPolicy) GetAssetTypes() []string`

GetAssetTypes returns the AssetTypes field if non-nil, zero value otherwise.

### GetAssetTypesOk

`func (o *InputTsugaAssetTagPolicy) GetAssetTypesOk() (*[]string, bool)`

GetAssetTypesOk returns a tuple with the AssetTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTypes

`func (o *InputTsugaAssetTagPolicy) SetAssetTypes(v []string)`

SetAssetTypes sets AssetTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


