# UserRoleBuiltIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Built-in organization role: &#x60;admin&#x60;, &#x60;public_editor&#x60;, &#x60;public_viewer&#x60;, or &#x60;teams_only&#x60;. | 

## Methods

### NewUserRoleBuiltIn

`func NewUserRoleBuiltIn(type_ string, ) *UserRoleBuiltIn`

NewUserRoleBuiltIn instantiates a new UserRoleBuiltIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRoleBuiltInWithDefaults

`func NewUserRoleBuiltInWithDefaults() *UserRoleBuiltIn`

NewUserRoleBuiltInWithDefaults instantiates a new UserRoleBuiltIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserRoleBuiltIn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserRoleBuiltIn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserRoleBuiltIn) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


