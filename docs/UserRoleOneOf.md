# UserRoleOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The user holds a custom role. | 
**CustomRoleId** | **string** | Identifier of the custom role. Use it with the custom role endpoints to read what the role grants. | 

## Methods

### NewUserRoleOneOf

`func NewUserRoleOneOf(type_ string, customRoleId string, ) *UserRoleOneOf`

NewUserRoleOneOf instantiates a new UserRoleOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRoleOneOfWithDefaults

`func NewUserRoleOneOfWithDefaults() *UserRoleOneOf`

NewUserRoleOneOfWithDefaults instantiates a new UserRoleOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserRoleOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserRoleOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserRoleOneOf) SetType(v string)`

SetType sets Type field to given value.


### GetCustomRoleId

`func (o *UserRoleOneOf) GetCustomRoleId() string`

GetCustomRoleId returns the CustomRoleId field if non-nil, zero value otherwise.

### GetCustomRoleIdOk

`func (o *UserRoleOneOf) GetCustomRoleIdOk() (*string, bool)`

GetCustomRoleIdOk returns a tuple with the CustomRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoleId

`func (o *UserRoleOneOf) SetCustomRoleId(v string)`

SetCustomRoleId sets CustomRoleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


