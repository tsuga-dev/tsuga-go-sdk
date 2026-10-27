# InputWidgetOngoingAlertsOverlayFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Team** | Pointer to **[]string** | Names of the teams owning the monitor, or the SLO for an SLO alert. Names rather than IDs, matching how the Ongoing alerts page filters by owner. | [optional] 
**TriggeredBy** | Pointer to **[]string** | Group-by dimensions that raised the alert, each formatted as &#x60;key:value&#x60;. An alert matches when any of its triggering dimensions is listed. | [optional] 
**Priority** | Pointer to **[]int32** | Priorities the alert was raised at. A monitor alert takes its priority from the monitor, an SLO alert from the SLO alert that fired. | [optional] 
**Cluster** | Pointer to **[]string** | IDs of the clusters the alert was raised on. Alerts raised on no cluster are always drawn. | [optional] 

## Methods

### NewInputWidgetOngoingAlertsOverlayFilters

`func NewInputWidgetOngoingAlertsOverlayFilters() *InputWidgetOngoingAlertsOverlayFilters`

NewInputWidgetOngoingAlertsOverlayFilters instantiates a new InputWidgetOngoingAlertsOverlayFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputWidgetOngoingAlertsOverlayFiltersWithDefaults

`func NewInputWidgetOngoingAlertsOverlayFiltersWithDefaults() *InputWidgetOngoingAlertsOverlayFilters`

NewInputWidgetOngoingAlertsOverlayFiltersWithDefaults instantiates a new InputWidgetOngoingAlertsOverlayFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeam

`func (o *InputWidgetOngoingAlertsOverlayFilters) GetTeam() []string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *InputWidgetOngoingAlertsOverlayFilters) GetTeamOk() (*[]string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *InputWidgetOngoingAlertsOverlayFilters) SetTeam(v []string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *InputWidgetOngoingAlertsOverlayFilters) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetTriggeredBy

`func (o *InputWidgetOngoingAlertsOverlayFilters) GetTriggeredBy() []string`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *InputWidgetOngoingAlertsOverlayFilters) GetTriggeredByOk() (*[]string, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *InputWidgetOngoingAlertsOverlayFilters) SetTriggeredBy(v []string)`

SetTriggeredBy sets TriggeredBy field to given value.

### HasTriggeredBy

`func (o *InputWidgetOngoingAlertsOverlayFilters) HasTriggeredBy() bool`

HasTriggeredBy returns a boolean if a field has been set.

### GetPriority

`func (o *InputWidgetOngoingAlertsOverlayFilters) GetPriority() []int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InputWidgetOngoingAlertsOverlayFilters) GetPriorityOk() (*[]int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InputWidgetOngoingAlertsOverlayFilters) SetPriority(v []int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InputWidgetOngoingAlertsOverlayFilters) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetCluster

`func (o *InputWidgetOngoingAlertsOverlayFilters) GetCluster() []string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *InputWidgetOngoingAlertsOverlayFilters) GetClusterOk() (*[]string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *InputWidgetOngoingAlertsOverlayFilters) SetCluster(v []string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *InputWidgetOngoingAlertsOverlayFilters) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


