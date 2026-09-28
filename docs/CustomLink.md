# CustomLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Text of the entry added to the widget context menu. Defaults to \&quot;Open custom link\&quot; when omitted. | [optional] 
**Url** | **string** | URL opened in a new tab by the context menu entry. Every &#x60;{{group.&lt;field&gt;}}&#x60; placeholder is replaced by the URL-encoded value the clicked series has for that group-by field; without placeholders, every series opens the same URL. &#x60;{&#x60; and &#x60;}&#x60; are only allowed as part of a placeholder. Tsuga omits the menu entry when the clicked series has no value for one of the referenced fields. | 

## Methods

### NewCustomLink

`func NewCustomLink(url string, ) *CustomLink`

NewCustomLink instantiates a new CustomLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomLinkWithDefaults

`func NewCustomLinkWithDefaults() *CustomLink`

NewCustomLinkWithDefaults instantiates a new CustomLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CustomLink) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomLink) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomLink) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CustomLink) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetUrl

`func (o *CustomLink) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CustomLink) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CustomLink) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


