# IngestionApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**KeyLastCharacters** | **string** |  | 
**Owner** | **string** |  | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List of key/value tags applied to the resource | [optional] 
**TeamOverrideFields** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIngestionApiKey

`func NewIngestionApiKey(id string, name string, keyLastCharacters string, owner string, ) *IngestionApiKey`

NewIngestionApiKey instantiates a new IngestionApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionApiKeyWithDefaults

`func NewIngestionApiKeyWithDefaults() *IngestionApiKey`

NewIngestionApiKeyWithDefaults instantiates a new IngestionApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IngestionApiKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IngestionApiKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IngestionApiKey) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *IngestionApiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IngestionApiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IngestionApiKey) SetName(v string)`

SetName sets Name field to given value.


### GetKeyLastCharacters

`func (o *IngestionApiKey) GetKeyLastCharacters() string`

GetKeyLastCharacters returns the KeyLastCharacters field if non-nil, zero value otherwise.

### GetKeyLastCharactersOk

`func (o *IngestionApiKey) GetKeyLastCharactersOk() (*string, bool)`

GetKeyLastCharactersOk returns a tuple with the KeyLastCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLastCharacters

`func (o *IngestionApiKey) SetKeyLastCharacters(v string)`

SetKeyLastCharacters sets KeyLastCharacters field to given value.


### GetOwner

`func (o *IngestionApiKey) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IngestionApiKey) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IngestionApiKey) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetTags

`func (o *IngestionApiKey) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IngestionApiKey) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IngestionApiKey) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IngestionApiKey) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTeamOverrideFields

`func (o *IngestionApiKey) GetTeamOverrideFields() []string`

GetTeamOverrideFields returns the TeamOverrideFields field if non-nil, zero value otherwise.

### GetTeamOverrideFieldsOk

`func (o *IngestionApiKey) GetTeamOverrideFieldsOk() (*[]string, bool)`

GetTeamOverrideFieldsOk returns a tuple with the TeamOverrideFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamOverrideFields

`func (o *IngestionApiKey) SetTeamOverrideFields(v []string)`

SetTeamOverrideFields sets TeamOverrideFields field to given value.

### HasTeamOverrideFields

`func (o *IngestionApiKey) HasTeamOverrideFields() bool`

HasTeamOverrideFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


