# UpdateAllowedDomainsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedDomains** | **[]string** | The new full list of email domains whose users are allowed to join the account via invitation or SSO/SAML just-in-time provisioning. Replaces the existing list; omitted domains are no longer allowed. | 

## Methods

### NewUpdateAllowedDomainsRequest

`func NewUpdateAllowedDomainsRequest(allowedDomains []string, ) *UpdateAllowedDomainsRequest`

NewUpdateAllowedDomainsRequest instantiates a new UpdateAllowedDomainsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAllowedDomainsRequestWithDefaults

`func NewUpdateAllowedDomainsRequestWithDefaults() *UpdateAllowedDomainsRequest`

NewUpdateAllowedDomainsRequestWithDefaults instantiates a new UpdateAllowedDomainsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedDomains

`func (o *UpdateAllowedDomainsRequest) GetAllowedDomains() []string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *UpdateAllowedDomainsRequest) GetAllowedDomainsOk() (*[]string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *UpdateAllowedDomainsRequest) SetAllowedDomains(v []string)`

SetAllowedDomains sets AllowedDomains field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


