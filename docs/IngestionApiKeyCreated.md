# IngestionApiKeyCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Tsuga-generated ingestion API key ID assigned when the key is created. | 
**Name** | **string** | Human-readable key name. Names must be kebab-case when creating or updating keys. | 
**KeyLastCharacters** | **string** | Last visible characters of the secret key. Use this to identify a key without exposing the full secret. | 
**Owner** | **string** | Team ID that owns and manages the ingestion API key. | 
**Tags** | [**[]Tag**](Tag.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | 
**TeamOverrideFields** | Pointer to **[]string** | Telemetry attribute names used to override ownership for data sent with this key. | [optional] 
**Key** | **string** | Full Tsuga-generated ingestion API key secret. Returned only in the create response; list, get, and update responses expose only &#x60;keyLastCharacters&#x60;. | 

## Methods

### NewIngestionApiKeyCreated

`func NewIngestionApiKeyCreated(id string, name string, keyLastCharacters string, owner string, tags []Tag, key string, ) *IngestionApiKeyCreated`

NewIngestionApiKeyCreated instantiates a new IngestionApiKeyCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionApiKeyCreatedWithDefaults

`func NewIngestionApiKeyCreatedWithDefaults() *IngestionApiKeyCreated`

NewIngestionApiKeyCreatedWithDefaults instantiates a new IngestionApiKeyCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IngestionApiKeyCreated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IngestionApiKeyCreated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IngestionApiKeyCreated) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *IngestionApiKeyCreated) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IngestionApiKeyCreated) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IngestionApiKeyCreated) SetName(v string)`

SetName sets Name field to given value.


### GetKeyLastCharacters

`func (o *IngestionApiKeyCreated) GetKeyLastCharacters() string`

GetKeyLastCharacters returns the KeyLastCharacters field if non-nil, zero value otherwise.

### GetKeyLastCharactersOk

`func (o *IngestionApiKeyCreated) GetKeyLastCharactersOk() (*string, bool)`

GetKeyLastCharactersOk returns a tuple with the KeyLastCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLastCharacters

`func (o *IngestionApiKeyCreated) SetKeyLastCharacters(v string)`

SetKeyLastCharacters sets KeyLastCharacters field to given value.


### GetOwner

`func (o *IngestionApiKeyCreated) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IngestionApiKeyCreated) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IngestionApiKeyCreated) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetTags

`func (o *IngestionApiKeyCreated) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IngestionApiKeyCreated) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IngestionApiKeyCreated) SetTags(v []Tag)`

SetTags sets Tags field to given value.


### GetTeamOverrideFields

`func (o *IngestionApiKeyCreated) GetTeamOverrideFields() []string`

GetTeamOverrideFields returns the TeamOverrideFields field if non-nil, zero value otherwise.

### GetTeamOverrideFieldsOk

`func (o *IngestionApiKeyCreated) GetTeamOverrideFieldsOk() (*[]string, bool)`

GetTeamOverrideFieldsOk returns a tuple with the TeamOverrideFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamOverrideFields

`func (o *IngestionApiKeyCreated) SetTeamOverrideFields(v []string)`

SetTeamOverrideFields sets TeamOverrideFields field to given value.

### HasTeamOverrideFields

`func (o *IngestionApiKeyCreated) HasTeamOverrideFields() bool`

HasTeamOverrideFields returns a boolean if a field has been set.

### GetKey

`func (o *IngestionApiKeyCreated) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IngestionApiKeyCreated) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IngestionApiKeyCreated) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


