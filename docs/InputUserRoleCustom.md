# InputUserRoleCustom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Assigns the custom role &#x60;customRoleId&#x60;. | 
**CustomRoleId** | **string** | Identifier of the custom role to assign. Use the &#x60;id&#x60; returned by the custom role endpoints. | 

## Methods

### NewInputUserRoleCustom

`func NewInputUserRoleCustom(type_ string, customRoleId string, ) *InputUserRoleCustom`

NewInputUserRoleCustom instantiates a new InputUserRoleCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputUserRoleCustomWithDefaults

`func NewInputUserRoleCustomWithDefaults() *InputUserRoleCustom`

NewInputUserRoleCustomWithDefaults instantiates a new InputUserRoleCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputUserRoleCustom) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputUserRoleCustom) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputUserRoleCustom) SetType(v string)`

SetType sets Type field to given value.


### GetCustomRoleId

`func (o *InputUserRoleCustom) GetCustomRoleId() string`

GetCustomRoleId returns the CustomRoleId field if non-nil, zero value otherwise.

### GetCustomRoleIdOk

`func (o *InputUserRoleCustom) GetCustomRoleIdOk() (*string, bool)`

GetCustomRoleIdOk returns a tuple with the CustomRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoleId

`func (o *InputUserRoleCustom) SetCustomRoleId(v string)`

SetCustomRoleId sets CustomRoleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


