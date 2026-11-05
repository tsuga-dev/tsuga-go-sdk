# DashboardSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Tsuga-generated dashboard ID assigned when the dashboard is created. | 
**Name** | **string** | Display name of the dashboard. Set by the caller on create or update. | 
**Owner** | **string** | ID of the team that owns the dashboard. Set by the caller on create or update and used for dashboard access control. | 
**Filters** | Pointer to [**[]DashboardSummaryFiltersInner**](DashboardSummaryFiltersInner.md) | Dashboard-wide filters applied to every widget on the dashboard. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | [optional] 
**TimePreset** | Pointer to **string** | Relative time preset used when opening the dashboard. | [optional] 
**FolderId** | Pointer to **string** | Folder that contains the dashboard. Set from the dashboard&#39;s saved folder relationship. Returned when the dashboard is in a folder; omitted when it is not. | [optional] 
**WidgetCount** | **float32** | Number of widgets on the dashboard. | 

## Methods

### NewDashboardSummary

`func NewDashboardSummary(id string, name string, owner string, widgetCount float32, ) *DashboardSummary`

NewDashboardSummary instantiates a new DashboardSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardSummaryWithDefaults

`func NewDashboardSummaryWithDefaults() *DashboardSummary`

NewDashboardSummaryWithDefaults instantiates a new DashboardSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DashboardSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardSummary) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DashboardSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardSummary) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *DashboardSummary) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DashboardSummary) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DashboardSummary) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetFilters

`func (o *DashboardSummary) GetFilters() []DashboardSummaryFiltersInner`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DashboardSummary) GetFiltersOk() (*[]DashboardSummaryFiltersInner, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DashboardSummary) SetFilters(v []DashboardSummaryFiltersInner)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *DashboardSummary) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetTags

`func (o *DashboardSummary) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DashboardSummary) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DashboardSummary) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DashboardSummary) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimePreset

`func (o *DashboardSummary) GetTimePreset() string`

GetTimePreset returns the TimePreset field if non-nil, zero value otherwise.

### GetTimePresetOk

`func (o *DashboardSummary) GetTimePresetOk() (*string, bool)`

GetTimePresetOk returns a tuple with the TimePreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePreset

`func (o *DashboardSummary) SetTimePreset(v string)`

SetTimePreset sets TimePreset field to given value.

### HasTimePreset

`func (o *DashboardSummary) HasTimePreset() bool`

HasTimePreset returns a boolean if a field has been set.

### GetFolderId

`func (o *DashboardSummary) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *DashboardSummary) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *DashboardSummary) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *DashboardSummary) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetWidgetCount

`func (o *DashboardSummary) GetWidgetCount() float32`

GetWidgetCount returns the WidgetCount field if non-nil, zero value otherwise.

### GetWidgetCountOk

`func (o *DashboardSummary) GetWidgetCountOk() (*float32, bool)`

GetWidgetCountOk returns a tuple with the WidgetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetCount

`func (o *DashboardSummary) SetWidgetCount(v float32)`

SetWidgetCount sets WidgetCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


