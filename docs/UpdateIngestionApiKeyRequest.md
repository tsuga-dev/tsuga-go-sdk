# UpdateIngestionApiKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Kebab-case name for the ingestion API key. Maximum length is 100 characters. | 
**Owner** | **string** | Team ID that will own and manage the ingestion API key. | 
**Tags** | [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | 
**TeamOverrideFields** | **[]string** | Array form. Up to 3 fields are allowed. | 

## Methods

### NewUpdateIngestionApiKeyRequest

`func NewUpdateIngestionApiKeyRequest(name string, owner string, tags []InputTag, teamOverrideFields []string, ) *UpdateIngestionApiKeyRequest`

NewUpdateIngestionApiKeyRequest instantiates a new UpdateIngestionApiKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIngestionApiKeyRequestWithDefaults

`func NewUpdateIngestionApiKeyRequestWithDefaults() *UpdateIngestionApiKeyRequest`

NewUpdateIngestionApiKeyRequestWithDefaults instantiates a new UpdateIngestionApiKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateIngestionApiKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateIngestionApiKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateIngestionApiKeyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *UpdateIngestionApiKeyRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateIngestionApiKeyRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateIngestionApiKeyRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetTags

`func (o *UpdateIngestionApiKeyRequest) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateIngestionApiKeyRequest) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateIngestionApiKeyRequest) SetTags(v []InputTag)`

SetTags sets Tags field to given value.


### GetTeamOverrideFields

`func (o *UpdateIngestionApiKeyRequest) GetTeamOverrideFields() []string`

GetTeamOverrideFields returns the TeamOverrideFields field if non-nil, zero value otherwise.

### GetTeamOverrideFieldsOk

`func (o *UpdateIngestionApiKeyRequest) GetTeamOverrideFieldsOk() (*[]string, bool)`

GetTeamOverrideFieldsOk returns a tuple with the TeamOverrideFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamOverrideFields

`func (o *UpdateIngestionApiKeyRequest) SetTeamOverrideFields(v []string)`

SetTeamOverrideFields sets TeamOverrideFields field to given value.


### SetTeamOverrideFieldsNil

`func (o *UpdateIngestionApiKeyRequest) SetTeamOverrideFieldsNil(b bool)`

 SetTeamOverrideFieldsNil sets the value for TeamOverrideFields to be an explicit nil

### UnsetTeamOverrideFields
`func (o *UpdateIngestionApiKeyRequest) UnsetTeamOverrideFields()`

UnsetTeamOverrideFields ensures that no value is present for TeamOverrideFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


