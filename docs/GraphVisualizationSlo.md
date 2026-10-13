# GraphVisualizationSlo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the latest status of SLOs as summary cards with the result count per status, a table of results, or both. Status is not recomputed over the dashboard time range. | 
**SloId** | Pointer to **string** | Id of the SLO to display. Omit it to cover every SLO that matches &#x60;filters&#x60;. | [optional] 
**Filters** | Pointer to [**GraphVisualizationSloFilters**](GraphVisualizationSloFilters.md) |  | [optional] 
**OrganizeBy** | Pointer to **string** | Groups the results by owning team or by cluster. Omit it to list results ungrouped. | [optional] 
**Display** | **[]string** | Sections to show: &#x60;summary&#x60; for the result count per status, &#x60;table&#x60; for the list of results. | 

## Methods

### NewGraphVisualizationSlo

`func NewGraphVisualizationSlo(type_ string, display []string, ) *GraphVisualizationSlo`

NewGraphVisualizationSlo instantiates a new GraphVisualizationSlo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationSloWithDefaults

`func NewGraphVisualizationSloWithDefaults() *GraphVisualizationSlo`

NewGraphVisualizationSloWithDefaults instantiates a new GraphVisualizationSlo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GraphVisualizationSlo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphVisualizationSlo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphVisualizationSlo) SetType(v string)`

SetType sets Type field to given value.


### GetSloId

`func (o *GraphVisualizationSlo) GetSloId() string`

GetSloId returns the SloId field if non-nil, zero value otherwise.

### GetSloIdOk

`func (o *GraphVisualizationSlo) GetSloIdOk() (*string, bool)`

GetSloIdOk returns a tuple with the SloId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloId

`func (o *GraphVisualizationSlo) SetSloId(v string)`

SetSloId sets SloId field to given value.

### HasSloId

`func (o *GraphVisualizationSlo) HasSloId() bool`

HasSloId returns a boolean if a field has been set.

### GetFilters

`func (o *GraphVisualizationSlo) GetFilters() GraphVisualizationSloFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GraphVisualizationSlo) GetFiltersOk() (*GraphVisualizationSloFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GraphVisualizationSlo) SetFilters(v GraphVisualizationSloFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GraphVisualizationSlo) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetOrganizeBy

`func (o *GraphVisualizationSlo) GetOrganizeBy() string`

GetOrganizeBy returns the OrganizeBy field if non-nil, zero value otherwise.

### GetOrganizeByOk

`func (o *GraphVisualizationSlo) GetOrganizeByOk() (*string, bool)`

GetOrganizeByOk returns a tuple with the OrganizeBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizeBy

`func (o *GraphVisualizationSlo) SetOrganizeBy(v string)`

SetOrganizeBy sets OrganizeBy field to given value.

### HasOrganizeBy

`func (o *GraphVisualizationSlo) HasOrganizeBy() bool`

HasOrganizeBy returns a boolean if a field has been set.

### GetDisplay

`func (o *GraphVisualizationSlo) GetDisplay() []string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *GraphVisualizationSlo) GetDisplayOk() (*[]string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *GraphVisualizationSlo) SetDisplay(v []string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


