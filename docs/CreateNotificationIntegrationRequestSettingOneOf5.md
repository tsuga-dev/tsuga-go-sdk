# CreateNotificationIntegrationRequestSettingOneOf5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**InstanceName** | **string** | Instance name of the ServiceNow instance. This corresponds to the subdomain part of the instance URL, like \&quot;my-instance\&quot; in \&quot;my-instance.service-now.com\&quot; | 
**Username** | **NullableString** | Folder filter for dashboard query results. Optional; pass a folder ID to return dashboards in that folder, or null to return only dashboards that are not in any folder. | 
**Password** | **NullableString** | Folder filter for dashboard query results. Optional; pass a folder ID to return dashboards in that folder, or null to return only dashboards that are not in any folder. | 

## Methods

### NewCreateNotificationIntegrationRequestSettingOneOf5

`func NewCreateNotificationIntegrationRequestSettingOneOf5(type_ string, instanceName string, username NullableString, password NullableString, ) *CreateNotificationIntegrationRequestSettingOneOf5`

NewCreateNotificationIntegrationRequestSettingOneOf5 instantiates a new CreateNotificationIntegrationRequestSettingOneOf5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationIntegrationRequestSettingOneOf5WithDefaults

`func NewCreateNotificationIntegrationRequestSettingOneOf5WithDefaults() *CreateNotificationIntegrationRequestSettingOneOf5`

NewCreateNotificationIntegrationRequestSettingOneOf5WithDefaults instantiates a new CreateNotificationIntegrationRequestSettingOneOf5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateNotificationIntegrationRequestSettingOneOf5) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNotificationIntegrationRequestSettingOneOf5) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNotificationIntegrationRequestSettingOneOf5) SetType(v string)`

SetType sets Type field to given value.


### GetInstanceName

`func (o *CreateNotificationIntegrationRequestSettingOneOf5) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *CreateNotificationIntegrationRequestSettingOneOf5) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *CreateNotificationIntegrationRequestSettingOneOf5) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### GetUsername

`func (o *CreateNotificationIntegrationRequestSettingOneOf5) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateNotificationIntegrationRequestSettingOneOf5) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateNotificationIntegrationRequestSettingOneOf5) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *CreateNotificationIntegrationRequestSettingOneOf5) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *CreateNotificationIntegrationRequestSettingOneOf5) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *CreateNotificationIntegrationRequestSettingOneOf5) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateNotificationIntegrationRequestSettingOneOf5) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateNotificationIntegrationRequestSettingOneOf5) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *CreateNotificationIntegrationRequestSettingOneOf5) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CreateNotificationIntegrationRequestSettingOneOf5) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


