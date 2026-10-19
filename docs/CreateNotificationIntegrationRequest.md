# CreateNotificationIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Display name of the graph widget. Pass null to remove it. | 
**Owner** | **string** | ID of the team that will own the integration, determining which team members can view and manage it. Only an org admin or an admin of this team may set it. | 
**Setting** | [**CreateNotificationIntegrationRequestSetting**](CreateNotificationIntegrationRequestSetting.md) |  | 
**Tags** | Pointer to [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 

## Methods

### NewCreateNotificationIntegrationRequest

`func NewCreateNotificationIntegrationRequest(name NullableString, owner string, setting CreateNotificationIntegrationRequestSetting, ) *CreateNotificationIntegrationRequest`

NewCreateNotificationIntegrationRequest instantiates a new CreateNotificationIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationIntegrationRequestWithDefaults

`func NewCreateNotificationIntegrationRequestWithDefaults() *CreateNotificationIntegrationRequest`

NewCreateNotificationIntegrationRequestWithDefaults instantiates a new CreateNotificationIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNotificationIntegrationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNotificationIntegrationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNotificationIntegrationRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateNotificationIntegrationRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateNotificationIntegrationRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwner

`func (o *CreateNotificationIntegrationRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateNotificationIntegrationRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateNotificationIntegrationRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetSetting

`func (o *CreateNotificationIntegrationRequest) GetSetting() CreateNotificationIntegrationRequestSetting`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *CreateNotificationIntegrationRequest) GetSettingOk() (*CreateNotificationIntegrationRequestSetting, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *CreateNotificationIntegrationRequest) SetSetting(v CreateNotificationIntegrationRequestSetting)`

SetSetting sets Setting field to given value.


### GetTags

`func (o *CreateNotificationIntegrationRequest) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateNotificationIntegrationRequest) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateNotificationIntegrationRequest) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateNotificationIntegrationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


