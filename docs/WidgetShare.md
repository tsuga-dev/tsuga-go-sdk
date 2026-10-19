# WidgetShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Code of the share link, the last path segment of &#x60;https://api.tsuga.com/v0/widget-shares/&lt;orgId&gt;/&lt;id&gt;&#x60;. | 
**Title** | Pointer to **string** | Name of the widget the link was shared from. Omitted when the widget had no name. | [optional] 
**Search** | **string** | Analytics search parameters captured when the link was created, without the leading &#x60;?&#x60;. Values use the app URL encoding: object and array values are JSON-stringified, then URL-encoded. &#x60;timeZone&#x60; is the time zone of the user who created the link, which the shared image uses. | 
**DashboardId** | Pointer to **string** | ID of the dashboard holding the shared widget. Omitted for Analytics shares. | [optional] 
**GraphId** | Pointer to **string** | ID of the shared widget on its dashboard. Omitted for Analytics shares. | [optional] 
**Url** | **string** | Absolute app URL a signed-in browser opening the share link is redirected to: the widget on its dashboard, or Analytics. It carries no &#x60;timeZone&#x60;, so the app shows the result in the time zone of the viewer. | 
**ExpiresAt** | **string** | Timestamp after which the share link stops resolving. | 

## Methods

### NewWidgetShare

`func NewWidgetShare(id string, search string, url string, expiresAt string, ) *WidgetShare`

NewWidgetShare instantiates a new WidgetShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetShareWithDefaults

`func NewWidgetShareWithDefaults() *WidgetShare`

NewWidgetShareWithDefaults instantiates a new WidgetShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WidgetShare) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WidgetShare) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WidgetShare) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *WidgetShare) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WidgetShare) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WidgetShare) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WidgetShare) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSearch

`func (o *WidgetShare) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *WidgetShare) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *WidgetShare) SetSearch(v string)`

SetSearch sets Search field to given value.


### GetDashboardId

`func (o *WidgetShare) GetDashboardId() string`

GetDashboardId returns the DashboardId field if non-nil, zero value otherwise.

### GetDashboardIdOk

`func (o *WidgetShare) GetDashboardIdOk() (*string, bool)`

GetDashboardIdOk returns a tuple with the DashboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardId

`func (o *WidgetShare) SetDashboardId(v string)`

SetDashboardId sets DashboardId field to given value.

### HasDashboardId

`func (o *WidgetShare) HasDashboardId() bool`

HasDashboardId returns a boolean if a field has been set.

### GetGraphId

`func (o *WidgetShare) GetGraphId() string`

GetGraphId returns the GraphId field if non-nil, zero value otherwise.

### GetGraphIdOk

`func (o *WidgetShare) GetGraphIdOk() (*string, bool)`

GetGraphIdOk returns a tuple with the GraphId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphId

`func (o *WidgetShare) SetGraphId(v string)`

SetGraphId sets GraphId field to given value.

### HasGraphId

`func (o *WidgetShare) HasGraphId() bool`

HasGraphId returns a boolean if a field has been set.

### GetUrl

`func (o *WidgetShare) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WidgetShare) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WidgetShare) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetExpiresAt

`func (o *WidgetShare) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *WidgetShare) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *WidgetShare) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


