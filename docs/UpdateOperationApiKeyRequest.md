# UpdateOperationApiKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Kebab-case name for the operation API key. | 
**Owner** | **string** | ID of the team that owns the operation API key. When the caller is an operation API key, the team must be within the caller&#39;s scope. | 
**Tags** | Pointer to [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 
**Scope** | **string** | Team-access boundary for the operation API key. Set when the key is created or updated. &#x60;all&#x60; can reach all teams, &#x60;owning-team-and-public&#x60; can reach the owning team and public teams, and &#x60;owning-team-only&#x60; can reach only the owning team. | 
**Permissions** | [**CreateOperationApiKeyRequestPermissions**](CreateOperationApiKeyRequestPermissions.md) |  | 
**ExpiresAt** | **time.Time** | Date and time after which the key stops working. When the caller is an operation API key, set a future expiration, no later than the caller&#39;s own when the caller expires. | 

## Methods

### NewUpdateOperationApiKeyRequest

`func NewUpdateOperationApiKeyRequest(name string, owner string, scope string, permissions CreateOperationApiKeyRequestPermissions, expiresAt time.Time, ) *UpdateOperationApiKeyRequest`

NewUpdateOperationApiKeyRequest instantiates a new UpdateOperationApiKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOperationApiKeyRequestWithDefaults

`func NewUpdateOperationApiKeyRequestWithDefaults() *UpdateOperationApiKeyRequest`

NewUpdateOperationApiKeyRequestWithDefaults instantiates a new UpdateOperationApiKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateOperationApiKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOperationApiKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOperationApiKeyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *UpdateOperationApiKeyRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateOperationApiKeyRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateOperationApiKeyRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetTags

`func (o *UpdateOperationApiKeyRequest) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateOperationApiKeyRequest) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateOperationApiKeyRequest) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateOperationApiKeyRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetScope

`func (o *UpdateOperationApiKeyRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UpdateOperationApiKeyRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UpdateOperationApiKeyRequest) SetScope(v string)`

SetScope sets Scope field to given value.


### GetPermissions

`func (o *UpdateOperationApiKeyRequest) GetPermissions() CreateOperationApiKeyRequestPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateOperationApiKeyRequest) GetPermissionsOk() (*CreateOperationApiKeyRequestPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateOperationApiKeyRequest) SetPermissions(v CreateOperationApiKeyRequestPermissions)`

SetPermissions sets Permissions field to given value.


### GetExpiresAt

`func (o *UpdateOperationApiKeyRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *UpdateOperationApiKeyRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *UpdateOperationApiKeyRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


