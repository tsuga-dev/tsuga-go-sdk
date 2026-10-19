# DocPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | Path identifying the page. Documentation pages use a slug path (e.g. &#x60;getting-started/send-your-first-logs&#x60;). API-reference pages use &#x60;api/&lt;operationId&gt;&#x60; (e.g. &#x60;api/aggregateScalar&#x60;). | 
**Title** | **string** | Title of the page (e.g. \&quot;Send your first logs\&quot; or \&quot;POST /v1/aggregation/multi-query/scalar\&quot;) | 
**Content** | **string** | Raw markdown content of the page | 

## Methods

### NewDocPage

`func NewDocPage(path string, title string, content string, ) *DocPage`

NewDocPage instantiates a new DocPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocPageWithDefaults

`func NewDocPageWithDefaults() *DocPage`

NewDocPageWithDefaults instantiates a new DocPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *DocPage) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DocPage) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DocPage) SetPath(v string)`

SetPath sets Path field to given value.


### GetTitle

`func (o *DocPage) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DocPage) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DocPage) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetContent

`func (o *DocPage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DocPage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DocPage) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


