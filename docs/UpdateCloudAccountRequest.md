# UpdateCloudAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountFriendlyName** | **string** | Human-readable name for the account shown in Tsuga. | 
**Owner** | Pointer to **NullableString** | Folder filter for dashboard query results. Optional; pass a folder ID to return dashboards in that folder, or null to return only dashboards that are not in any folder. | [optional] 

## Methods

### NewUpdateCloudAccountRequest

`func NewUpdateCloudAccountRequest(accountFriendlyName string, ) *UpdateCloudAccountRequest`

NewUpdateCloudAccountRequest instantiates a new UpdateCloudAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCloudAccountRequestWithDefaults

`func NewUpdateCloudAccountRequestWithDefaults() *UpdateCloudAccountRequest`

NewUpdateCloudAccountRequestWithDefaults instantiates a new UpdateCloudAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountFriendlyName

`func (o *UpdateCloudAccountRequest) GetAccountFriendlyName() string`

GetAccountFriendlyName returns the AccountFriendlyName field if non-nil, zero value otherwise.

### GetAccountFriendlyNameOk

`func (o *UpdateCloudAccountRequest) GetAccountFriendlyNameOk() (*string, bool)`

GetAccountFriendlyNameOk returns a tuple with the AccountFriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFriendlyName

`func (o *UpdateCloudAccountRequest) SetAccountFriendlyName(v string)`

SetAccountFriendlyName sets AccountFriendlyName field to given value.


### GetOwner

`func (o *UpdateCloudAccountRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateCloudAccountRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateCloudAccountRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdateCloudAccountRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *UpdateCloudAccountRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *UpdateCloudAccountRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


