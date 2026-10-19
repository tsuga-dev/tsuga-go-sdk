# PublicToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Tsuga-generated RUM public token ID assigned when the token is created. | 
**Name** | **string** | Human-readable token name. Names must be kebab-case when creating or updating tokens. | 
**Token** | **string** | Full token value, prefixed with &#x60;tsuga_rum:v1:&#x60;, to configure in the RUM browser or mobile SDK. RUM tokens are public by design, so the value is not masked on subsequent reads and never changes after creation. | 
**Owner** | **string** | Team ID that owns and manages the RUM public token. | 
**ApplicationType** | **string** | Kind of application the token is issued for. Only &#x60;web&#x60; telemetry is checked against &#x60;allowedUrlPrefixes&#x60;, since mobile applications do not send an &#x60;Origin&#x60; header. | 
**AllowedUrlPrefixes** | **[]string** | URL prefixes which RUM telemetry data can be sent from. Supports wildcard matching. Always empty for &#x60;ios&#x60; and &#x60;android&#x60; applications. | 
**Tags** | [**[]Tag**](Tag.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | 
**TeamOverrideFields** | Pointer to **[]string** | Telemetry attribute names used to override ownership for data sent with this token. | [optional] 

## Methods

### NewPublicToken

`func NewPublicToken(id string, name string, token string, owner string, applicationType string, allowedUrlPrefixes []string, tags []Tag, ) *PublicToken`

NewPublicToken instantiates a new PublicToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicTokenWithDefaults

`func NewPublicTokenWithDefaults() *PublicToken`

NewPublicTokenWithDefaults instantiates a new PublicToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublicToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicToken) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PublicToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicToken) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *PublicToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PublicToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PublicToken) SetToken(v string)`

SetToken sets Token field to given value.


### GetOwner

`func (o *PublicToken) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PublicToken) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PublicToken) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetApplicationType

`func (o *PublicToken) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *PublicToken) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *PublicToken) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.


### GetAllowedUrlPrefixes

`func (o *PublicToken) GetAllowedUrlPrefixes() []string`

GetAllowedUrlPrefixes returns the AllowedUrlPrefixes field if non-nil, zero value otherwise.

### GetAllowedUrlPrefixesOk

`func (o *PublicToken) GetAllowedUrlPrefixesOk() (*[]string, bool)`

GetAllowedUrlPrefixesOk returns a tuple with the AllowedUrlPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUrlPrefixes

`func (o *PublicToken) SetAllowedUrlPrefixes(v []string)`

SetAllowedUrlPrefixes sets AllowedUrlPrefixes field to given value.


### GetTags

`func (o *PublicToken) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PublicToken) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PublicToken) SetTags(v []Tag)`

SetTags sets Tags field to given value.


### GetTeamOverrideFields

`func (o *PublicToken) GetTeamOverrideFields() []string`

GetTeamOverrideFields returns the TeamOverrideFields field if non-nil, zero value otherwise.

### GetTeamOverrideFieldsOk

`func (o *PublicToken) GetTeamOverrideFieldsOk() (*[]string, bool)`

GetTeamOverrideFieldsOk returns a tuple with the TeamOverrideFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamOverrideFields

`func (o *PublicToken) SetTeamOverrideFields(v []string)`

SetTeamOverrideFields sets TeamOverrideFields field to given value.

### HasTeamOverrideFields

`func (o *PublicToken) HasTeamOverrideFields() bool`

HasTeamOverrideFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


