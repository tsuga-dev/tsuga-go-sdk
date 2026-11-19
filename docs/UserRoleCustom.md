# UserRoleCustom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The user holds a custom role. | 
**CustomRoleId** | **string** | Identifier of the custom role. Use it with the custom role endpoints to read what the role grants. | 

## Methods

### NewUserRoleCustom

`func NewUserRoleCustom(type_ string, customRoleId string, ) *UserRoleCustom`

NewUserRoleCustom instantiates a new UserRoleCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRoleCustomWithDefaults

`func NewUserRoleCustomWithDefaults() *UserRoleCustom`

NewUserRoleCustomWithDefaults instantiates a new UserRoleCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserRoleCustom) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserRoleCustom) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserRoleCustom) SetType(v string)`

SetType sets Type field to given value.


### GetCustomRoleId

`func (o *UserRoleCustom) GetCustomRoleId() string`

GetCustomRoleId returns the CustomRoleId field if non-nil, zero value otherwise.

### GetCustomRoleIdOk

`func (o *UserRoleCustom) GetCustomRoleIdOk() (*string, bool)`

GetCustomRoleIdOk returns a tuple with the CustomRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoleId

`func (o *UserRoleCustom) SetCustomRoleId(v string)`

SetCustomRoleId sets CustomRoleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


