# QueryServicesRequestFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchQuery** | Pointer to [**QueryServicesRequestFiltersSearchQuery**](QueryServicesRequestFiltersSearchQuery.md) |  | [optional] 
**Teams** | Pointer to [**QueryServicesRequestFiltersTeams**](QueryServicesRequestFiltersTeams.md) |  | [optional] 
**Envs** | Pointer to [**QueryServicesRequestFiltersEnvs**](QueryServicesRequestFiltersEnvs.md) |  | [optional] 
**Versions** | Pointer to [**QueryServicesRequestFiltersVersions**](QueryServicesRequestFiltersVersions.md) |  | [optional] 
**Languages** | Pointer to [**QueryServicesRequestFiltersLanguages**](QueryServicesRequestFiltersLanguages.md) |  | [optional] 
**HasRecentData** | Pointer to **bool** | When true, only return services whose request rate over the queried window is non-zero | [optional] 

## Methods

### NewQueryServicesRequestFilters

`func NewQueryServicesRequestFilters() *QueryServicesRequestFilters`

NewQueryServicesRequestFilters instantiates a new QueryServicesRequestFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryServicesRequestFiltersWithDefaults

`func NewQueryServicesRequestFiltersWithDefaults() *QueryServicesRequestFilters`

NewQueryServicesRequestFiltersWithDefaults instantiates a new QueryServicesRequestFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchQuery

`func (o *QueryServicesRequestFilters) GetSearchQuery() QueryServicesRequestFiltersSearchQuery`

GetSearchQuery returns the SearchQuery field if non-nil, zero value otherwise.

### GetSearchQueryOk

`func (o *QueryServicesRequestFilters) GetSearchQueryOk() (*QueryServicesRequestFiltersSearchQuery, bool)`

GetSearchQueryOk returns a tuple with the SearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchQuery

`func (o *QueryServicesRequestFilters) SetSearchQuery(v QueryServicesRequestFiltersSearchQuery)`

SetSearchQuery sets SearchQuery field to given value.

### HasSearchQuery

`func (o *QueryServicesRequestFilters) HasSearchQuery() bool`

HasSearchQuery returns a boolean if a field has been set.

### GetTeams

`func (o *QueryServicesRequestFilters) GetTeams() QueryServicesRequestFiltersTeams`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *QueryServicesRequestFilters) GetTeamsOk() (*QueryServicesRequestFiltersTeams, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *QueryServicesRequestFilters) SetTeams(v QueryServicesRequestFiltersTeams)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *QueryServicesRequestFilters) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetEnvs

`func (o *QueryServicesRequestFilters) GetEnvs() QueryServicesRequestFiltersEnvs`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *QueryServicesRequestFilters) GetEnvsOk() (*QueryServicesRequestFiltersEnvs, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *QueryServicesRequestFilters) SetEnvs(v QueryServicesRequestFiltersEnvs)`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *QueryServicesRequestFilters) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.

### GetVersions

`func (o *QueryServicesRequestFilters) GetVersions() QueryServicesRequestFiltersVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *QueryServicesRequestFilters) GetVersionsOk() (*QueryServicesRequestFiltersVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *QueryServicesRequestFilters) SetVersions(v QueryServicesRequestFiltersVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *QueryServicesRequestFilters) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetLanguages

`func (o *QueryServicesRequestFilters) GetLanguages() QueryServicesRequestFiltersLanguages`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *QueryServicesRequestFilters) GetLanguagesOk() (*QueryServicesRequestFiltersLanguages, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *QueryServicesRequestFilters) SetLanguages(v QueryServicesRequestFiltersLanguages)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *QueryServicesRequestFilters) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetHasRecentData

`func (o *QueryServicesRequestFilters) GetHasRecentData() bool`

GetHasRecentData returns the HasRecentData field if non-nil, zero value otherwise.

### GetHasRecentDataOk

`func (o *QueryServicesRequestFilters) GetHasRecentDataOk() (*bool, bool)`

GetHasRecentDataOk returns a tuple with the HasRecentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRecentData

`func (o *QueryServicesRequestFilters) SetHasRecentData(v bool)`

SetHasRecentData sets HasRecentData field to given value.

### HasHasRecentData

`func (o *QueryServicesRequestFilters) HasHasRecentData() bool`

HasHasRecentData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


