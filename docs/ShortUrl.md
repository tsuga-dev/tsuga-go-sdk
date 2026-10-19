# ShortUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Code of the short link, the last path segment of &#x60;https://app.tsuga.com/l/&lt;id&gt;&#x60;. | 
**Url** | **string** | Relative app URL (path and search) the short link opens. Search parameters use the app URL encoding: object and array values are JSON-stringified, then URL-encoded. | 
**ExpiresAt** | **string** | Timestamp after which the short link stops resolving. | 

## Methods

### NewShortUrl

`func NewShortUrl(id string, url string, expiresAt string, ) *ShortUrl`

NewShortUrl instantiates a new ShortUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShortUrlWithDefaults

`func NewShortUrlWithDefaults() *ShortUrl`

NewShortUrlWithDefaults instantiates a new ShortUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShortUrl) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShortUrl) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShortUrl) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *ShortUrl) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ShortUrl) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ShortUrl) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetExpiresAt

`func (o *ShortUrl) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ShortUrl) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ShortUrl) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


