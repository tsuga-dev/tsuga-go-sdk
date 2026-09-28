# CustomLink1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Text of the entry added to the widget context menu. Defaults to \&quot;Open custom link\&quot; when omitted. | [optional] 
**Url** | **string** | URL opened in a new tab by the context menu entry. Every &#x60;{{group.&lt;field&gt;}}&#x60; placeholder is replaced by the URL-encoded value the clicked series has for that group-by field; without placeholders, every series opens the same URL. &#x60;{&#x60; and &#x60;}&#x60; are only allowed as part of a placeholder. Tsuga omits the menu entry when the clicked series has no value for one of the referenced fields. | 

## Methods

### NewCustomLink1

`func NewCustomLink1(url string, ) *CustomLink1`

NewCustomLink1 instantiates a new CustomLink1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomLink1WithDefaults

`func NewCustomLink1WithDefaults() *CustomLink1`

NewCustomLink1WithDefaults instantiates a new CustomLink1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CustomLink1) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomLink1) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomLink1) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CustomLink1) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetUrl

`func (o *CustomLink1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CustomLink1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CustomLink1) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


