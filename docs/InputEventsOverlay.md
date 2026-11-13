# InputEventsOverlay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Draws the deployments contained in the time range over the chart | 
**Teams** | Pointer to **[]string** | IDs of the teams whose services the app offers in its Services picker. It does not select what is drawn: the overlay draws only the services listed in &#x60;services&#x60;. | [optional] 
**Services** | Pointer to **[]string** | OpenTelemetry service names whose deployments are drawn. Nothing is drawn until at least one service is set. | [optional] 
**Env** | Pointer to **string** | Environment whose deployments are drawn. Deployments from every environment are drawn when omitted. | [optional] 
**Query** | Pointer to **string** | Case-insensitive substring matched against the alert, the same way the Ongoing alerts page matches its search box. An empty string matches every alert. Nothing is drawn unless &#x60;query&#x60; or at least one filter is set. | [optional] 
**AlertTypes** | Pointer to **[]string** | Kinds of alert drawn over the chart. Both monitor and SLO alerts are drawn when omitted. On its own, it does not make the overlay draw: set &#x60;query&#x60; or a filter too. | [optional] 
**Filters** | Pointer to [**InputWidgetOngoingAlertsOverlayFilters**](InputWidgetOngoingAlertsOverlayFilters.md) |  | [optional] 

## Methods

### NewInputEventsOverlay

`func NewInputEventsOverlay(type_ string, ) *InputEventsOverlay`

NewInputEventsOverlay instantiates a new InputEventsOverlay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputEventsOverlayWithDefaults

`func NewInputEventsOverlayWithDefaults() *InputEventsOverlay`

NewInputEventsOverlayWithDefaults instantiates a new InputEventsOverlay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputEventsOverlay) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputEventsOverlay) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputEventsOverlay) SetType(v string)`

SetType sets Type field to given value.


### GetTeams

`func (o *InputEventsOverlay) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *InputEventsOverlay) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *InputEventsOverlay) SetTeams(v []string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *InputEventsOverlay) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetServices

`func (o *InputEventsOverlay) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *InputEventsOverlay) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *InputEventsOverlay) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *InputEventsOverlay) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetEnv

`func (o *InputEventsOverlay) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *InputEventsOverlay) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *InputEventsOverlay) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *InputEventsOverlay) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetQuery

`func (o *InputEventsOverlay) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *InputEventsOverlay) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *InputEventsOverlay) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *InputEventsOverlay) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetAlertTypes

`func (o *InputEventsOverlay) GetAlertTypes() []string`

GetAlertTypes returns the AlertTypes field if non-nil, zero value otherwise.

### GetAlertTypesOk

`func (o *InputEventsOverlay) GetAlertTypesOk() (*[]string, bool)`

GetAlertTypesOk returns a tuple with the AlertTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTypes

`func (o *InputEventsOverlay) SetAlertTypes(v []string)`

SetAlertTypes sets AlertTypes field to given value.

### HasAlertTypes

`func (o *InputEventsOverlay) HasAlertTypes() bool`

HasAlertTypes returns a boolean if a field has been set.

### GetFilters

`func (o *InputEventsOverlay) GetFilters() InputWidgetOngoingAlertsOverlayFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *InputEventsOverlay) GetFiltersOk() (*InputWidgetOngoingAlertsOverlayFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *InputEventsOverlay) SetFilters(v InputWidgetOngoingAlertsOverlayFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *InputEventsOverlay) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


