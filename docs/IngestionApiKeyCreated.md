# IngestionApiKeyCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**KeyLastCharacters** | **string** |  | 
**Owner** | **string** |  | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List of key/value tags applied to the resource | [optional] 
**TeamOverrideFields** | Pointer to **[]string** |  | [optional] 
**Key** | **string** | The full API key value, returned only at creation time | 

## Methods

### NewIngestionApiKeyCreated

`func NewIngestionApiKeyCreated(id string, name string, keyLastCharacters string, owner string, key string, ) *IngestionApiKeyCreated`

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

### HasTags

`func (o *IngestionApiKeyCreated) HasTags() bool`

HasTags returns a boolean if a field has been set.

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


