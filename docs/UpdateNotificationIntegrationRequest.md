# UpdateNotificationIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Display name of the graph widget. Pass null to remove it. | [optional] 
**Owner** | **string** | ID of the team that will own the integration, determining which team members can view and manage it. Only an org admin or an admin of this team may set it. | 
**Tags** | Pointer to [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 
**Setting** | Pointer to [**UpdateNotificationIntegrationRequestSetting**](UpdateNotificationIntegrationRequestSetting.md) |  | [optional] 

## Methods

### NewUpdateNotificationIntegrationRequest

`func NewUpdateNotificationIntegrationRequest(owner string, ) *UpdateNotificationIntegrationRequest`

NewUpdateNotificationIntegrationRequest instantiates a new UpdateNotificationIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNotificationIntegrationRequestWithDefaults

`func NewUpdateNotificationIntegrationRequestWithDefaults() *UpdateNotificationIntegrationRequest`

NewUpdateNotificationIntegrationRequestWithDefaults instantiates a new UpdateNotificationIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNotificationIntegrationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNotificationIntegrationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNotificationIntegrationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNotificationIntegrationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateNotificationIntegrationRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateNotificationIntegrationRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwner

`func (o *UpdateNotificationIntegrationRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateNotificationIntegrationRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateNotificationIntegrationRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetTags

`func (o *UpdateNotificationIntegrationRequest) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateNotificationIntegrationRequest) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateNotificationIntegrationRequest) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateNotificationIntegrationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSetting

`func (o *UpdateNotificationIntegrationRequest) GetSetting() UpdateNotificationIntegrationRequestSetting`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *UpdateNotificationIntegrationRequest) GetSettingOk() (*UpdateNotificationIntegrationRequestSetting, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *UpdateNotificationIntegrationRequest) SetSetting(v UpdateNotificationIntegrationRequestSetting)`

SetSetting sets Setting field to given value.

### HasSetting

`func (o *UpdateNotificationIntegrationRequest) HasSetting() bool`

HasSetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


