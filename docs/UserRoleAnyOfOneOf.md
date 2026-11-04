# UserRoleAnyOfOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The user holds a custom role. | 
**CustomRoleId** | **string** | Identifier of the custom role. Use it with the custom role endpoints to read what the role grants. | 

## Methods

### NewUserRoleAnyOfOneOf

`func NewUserRoleAnyOfOneOf(type_ string, customRoleId string, ) *UserRoleAnyOfOneOf`

NewUserRoleAnyOfOneOf instantiates a new UserRoleAnyOfOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRoleAnyOfOneOfWithDefaults

`func NewUserRoleAnyOfOneOfWithDefaults() *UserRoleAnyOfOneOf`

NewUserRoleAnyOfOneOfWithDefaults instantiates a new UserRoleAnyOfOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserRoleAnyOfOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserRoleAnyOfOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserRoleAnyOfOneOf) SetType(v string)`

SetType sets Type field to given value.


### GetCustomRoleId

`func (o *UserRoleAnyOfOneOf) GetCustomRoleId() string`

GetCustomRoleId returns the CustomRoleId field if non-nil, zero value otherwise.

### GetCustomRoleIdOk

`func (o *UserRoleAnyOfOneOf) GetCustomRoleIdOk() (*string, bool)`

GetCustomRoleIdOk returns a tuple with the CustomRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoleId

`func (o *UserRoleAnyOfOneOf) SetCustomRoleId(v string)`

SetCustomRoleId sets CustomRoleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


