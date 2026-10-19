# DocSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | &#x60;doc&#x60; for documentation pages (guides, references, release notes); &#x60;api&#x60; for API-reference endpoints indexed from the OpenAPI spec. Use this to decide how to present or follow up on the result. | 
**Path** | **string** | Path identifying the result. Documentation pages use a slug path (e.g. &#x60;data-collection/guides/how-to-add-resource-attributes&#x60;). API-reference results use &#x60;api/&lt;operationId&gt;&#x60; (e.g. &#x60;api/aggregateScalar&#x60;). | 
**Title** | **string** | Title of the result (e.g. \&quot;How to Add Resource Attributes\&quot; or \&quot;POST /v1/aggregation/multi-query/scalar\&quot;) | 
**Subsection** | Pointer to **string** | For documentation pages: the heading path inside the page that matched (e.g. \&quot;Language setup\&quot;). For API results: the matched field or parameter breadcrumb (e.g. \&quot;Body › queries\&quot;). Absent when the match is the page or endpoint intro. | [optional] 
**Tab** | Pointer to **string** | Tab label inside the page that the result matched. Set when the matching content is inside a &lt;Tabs&gt;&lt;Tab label&#x3D;\&quot;…\&quot;&gt; block. Always absent for API results. | [optional] 
**Anchor** | Pointer to **string** | For documentation pages: the slugified heading ID used as the URL fragment (e.g. &#x60;language-setup&#x60;). For API results: the Scalar deep-link fragment (&#x60;tag/&lt;tag&gt;/&lt;METHOD&gt;&lt;path&gt;&#x60;). The &#x60;link&#x60; field already has this baked in. | [optional] 
**Score** | **float32** | Relevance score between 0 and 1 (higher is more relevant) | 
**Excerpt** | **string** | Short snippet of matching content | 
**Link** | **string** | URL to the result in the Tsuga app. Documentation pages include the section anchor and tab query parameter when applicable. API results deep-link into the API reference at &#x60;/documentation/api#tag/…&#x60;. | 

## Methods

### NewDocSearchResult

`func NewDocSearchResult(kind string, path string, title string, score float32, excerpt string, link string, ) *DocSearchResult`

NewDocSearchResult instantiates a new DocSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocSearchResultWithDefaults

`func NewDocSearchResultWithDefaults() *DocSearchResult`

NewDocSearchResultWithDefaults instantiates a new DocSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *DocSearchResult) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DocSearchResult) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DocSearchResult) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPath

`func (o *DocSearchResult) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DocSearchResult) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DocSearchResult) SetPath(v string)`

SetPath sets Path field to given value.


### GetTitle

`func (o *DocSearchResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DocSearchResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DocSearchResult) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSubsection

`func (o *DocSearchResult) GetSubsection() string`

GetSubsection returns the Subsection field if non-nil, zero value otherwise.

### GetSubsectionOk

`func (o *DocSearchResult) GetSubsectionOk() (*string, bool)`

GetSubsectionOk returns a tuple with the Subsection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsection

`func (o *DocSearchResult) SetSubsection(v string)`

SetSubsection sets Subsection field to given value.

### HasSubsection

`func (o *DocSearchResult) HasSubsection() bool`

HasSubsection returns a boolean if a field has been set.

### GetTab

`func (o *DocSearchResult) GetTab() string`

GetTab returns the Tab field if non-nil, zero value otherwise.

### GetTabOk

`func (o *DocSearchResult) GetTabOk() (*string, bool)`

GetTabOk returns a tuple with the Tab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTab

`func (o *DocSearchResult) SetTab(v string)`

SetTab sets Tab field to given value.

### HasTab

`func (o *DocSearchResult) HasTab() bool`

HasTab returns a boolean if a field has been set.

### GetAnchor

`func (o *DocSearchResult) GetAnchor() string`

GetAnchor returns the Anchor field if non-nil, zero value otherwise.

### GetAnchorOk

`func (o *DocSearchResult) GetAnchorOk() (*string, bool)`

GetAnchorOk returns a tuple with the Anchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchor

`func (o *DocSearchResult) SetAnchor(v string)`

SetAnchor sets Anchor field to given value.

### HasAnchor

`func (o *DocSearchResult) HasAnchor() bool`

HasAnchor returns a boolean if a field has been set.

### GetScore

`func (o *DocSearchResult) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *DocSearchResult) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *DocSearchResult) SetScore(v float32)`

SetScore sets Score field to given value.


### GetExcerpt

`func (o *DocSearchResult) GetExcerpt() string`

GetExcerpt returns the Excerpt field if non-nil, zero value otherwise.

### GetExcerptOk

`func (o *DocSearchResult) GetExcerptOk() (*string, bool)`

GetExcerptOk returns a tuple with the Excerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcerpt

`func (o *DocSearchResult) SetExcerpt(v string)`

SetExcerpt sets Excerpt field to given value.


### GetLink

`func (o *DocSearchResult) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *DocSearchResult) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *DocSearchResult) SetLink(v string)`

SetLink sets Link field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


