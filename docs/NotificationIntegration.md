# NotificationIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Owner** | Pointer to **string** | ID of the team that owns the integration. Absent on integrations created before ownership was introduced. | [optional] 
**Setting** | [**NotificationIntegrationSetting**](NotificationIntegrationSetting.md) |  | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | [optional] 

## Methods

### NewNotificationIntegration

`func NewNotificationIntegration(id string, name string, setting NotificationIntegrationSetting, ) *NotificationIntegration`

NewNotificationIntegration instantiates a new NotificationIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationIntegrationWithDefaults

`func NewNotificationIntegrationWithDefaults() *NotificationIntegration`

NewNotificationIntegrationWithDefaults instantiates a new NotificationIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationIntegration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationIntegration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationIntegration) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *NotificationIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationIntegration) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *NotificationIntegration) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NotificationIntegration) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NotificationIntegration) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NotificationIntegration) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSetting

`func (o *NotificationIntegration) GetSetting() NotificationIntegrationSetting`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *NotificationIntegration) GetSettingOk() (*NotificationIntegrationSetting, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *NotificationIntegration) SetSetting(v NotificationIntegrationSetting)`

SetSetting sets Setting field to given value.


### GetTags

`func (o *NotificationIntegration) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NotificationIntegration) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NotificationIntegration) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NotificationIntegration) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


