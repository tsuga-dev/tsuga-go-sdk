# UpdatePublicTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Kebab-case name for the RUM public token. Maximum length is 100 characters. | 
**Owner** | **string** | Team ID that will own and manage the RUM public token. | 
**ApplicationType** | **string** | Kind of application the token is issued for. Only &#x60;web&#x60; telemetry is checked against &#x60;allowedUrlPrefixes&#x60;, since mobile applications do not send an &#x60;Origin&#x60; header. Fixed at creation: an update must resend the token&#39;s current value. | 
**AllowedUrlPrefixes** | **[]string** | URL prefixes which RUM telemetry data can be sent from. Supports wildcard matching. At least one prefix is required for &#x60;web&#x60; applications, and the array must be empty for &#x60;ios&#x60; and &#x60;android&#x60; applications. | 
**Tags** | [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | 
**TeamOverrideFields** | **[]string** | Array form. Up to 3 fields are allowed. | 

## Methods

### NewUpdatePublicTokenRequest

`func NewUpdatePublicTokenRequest(name string, owner string, applicationType string, allowedUrlPrefixes []string, tags []InputTag, teamOverrideFields []string, ) *UpdatePublicTokenRequest`

NewUpdatePublicTokenRequest instantiates a new UpdatePublicTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePublicTokenRequestWithDefaults

`func NewUpdatePublicTokenRequestWithDefaults() *UpdatePublicTokenRequest`

NewUpdatePublicTokenRequestWithDefaults instantiates a new UpdatePublicTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdatePublicTokenRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePublicTokenRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePublicTokenRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *UpdatePublicTokenRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdatePublicTokenRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdatePublicTokenRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetApplicationType

`func (o *UpdatePublicTokenRequest) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *UpdatePublicTokenRequest) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *UpdatePublicTokenRequest) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.


### GetAllowedUrlPrefixes

`func (o *UpdatePublicTokenRequest) GetAllowedUrlPrefixes() []string`

GetAllowedUrlPrefixes returns the AllowedUrlPrefixes field if non-nil, zero value otherwise.

### GetAllowedUrlPrefixesOk

`func (o *UpdatePublicTokenRequest) GetAllowedUrlPrefixesOk() (*[]string, bool)`

GetAllowedUrlPrefixesOk returns a tuple with the AllowedUrlPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUrlPrefixes

`func (o *UpdatePublicTokenRequest) SetAllowedUrlPrefixes(v []string)`

SetAllowedUrlPrefixes sets AllowedUrlPrefixes field to given value.


### GetTags

`func (o *UpdatePublicTokenRequest) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdatePublicTokenRequest) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdatePublicTokenRequest) SetTags(v []InputTag)`

SetTags sets Tags field to given value.


### GetTeamOverrideFields

`func (o *UpdatePublicTokenRequest) GetTeamOverrideFields() []string`

GetTeamOverrideFields returns the TeamOverrideFields field if non-nil, zero value otherwise.

### GetTeamOverrideFieldsOk

`func (o *UpdatePublicTokenRequest) GetTeamOverrideFieldsOk() (*[]string, bool)`

GetTeamOverrideFieldsOk returns a tuple with the TeamOverrideFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamOverrideFields

`func (o *UpdatePublicTokenRequest) SetTeamOverrideFields(v []string)`

SetTeamOverrideFields sets TeamOverrideFields field to given value.


### SetTeamOverrideFieldsNil

`func (o *UpdatePublicTokenRequest) SetTeamOverrideFieldsNil(b bool)`

 SetTeamOverrideFieldsNil sets the value for TeamOverrideFields to be an explicit nil

### UnsetTeamOverrideFields
`func (o *UpdatePublicTokenRequest) UnsetTeamOverrideFields()`

UnsetTeamOverrideFields ensures that no value is present for TeamOverrideFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


