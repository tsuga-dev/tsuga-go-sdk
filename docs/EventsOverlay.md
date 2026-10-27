# EventsOverlay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Draws the deployments contained in the time range over the chart | 
**Teams** | Pointer to **[]string** | IDs of the teams whose services are drawn. The set is resolved when the widget renders, so a service joining one of these teams later is picked up automatically. Every service is drawn when omitted. | [optional] 
**Services** | Pointer to **[]string** | OpenTelemetry service names to draw, narrowing &#x60;teams&#x60; down to an explicit list. Every service matching &#x60;teams&#x60; is drawn when omitted. | [optional] 
**Env** | Pointer to **string** | Environment whose deployments are drawn. Deployments from every environment are drawn when omitted. | [optional] 
**Query** | Pointer to **string** | Case-insensitive substring matched against the alert, the same way the Ongoing alerts page matches its search box. Every alert is drawn when omitted. | [optional] 
**AlertTypes** | Pointer to **[]string** | Kinds of alert drawn over the chart. Both monitor and SLO alerts are drawn when omitted. | [optional] 
**Filters** | Pointer to [**WidgetOngoingAlertsOverlayFilters**](WidgetOngoingAlertsOverlayFilters.md) |  | [optional] 

## Methods

### NewEventsOverlay

`func NewEventsOverlay(type_ string, ) *EventsOverlay`

NewEventsOverlay instantiates a new EventsOverlay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsOverlayWithDefaults

`func NewEventsOverlayWithDefaults() *EventsOverlay`

NewEventsOverlayWithDefaults instantiates a new EventsOverlay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventsOverlay) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventsOverlay) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventsOverlay) SetType(v string)`

SetType sets Type field to given value.


### GetTeams

`func (o *EventsOverlay) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *EventsOverlay) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *EventsOverlay) SetTeams(v []string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *EventsOverlay) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetServices

`func (o *EventsOverlay) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *EventsOverlay) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *EventsOverlay) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *EventsOverlay) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetEnv

`func (o *EventsOverlay) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *EventsOverlay) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *EventsOverlay) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *EventsOverlay) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetQuery

`func (o *EventsOverlay) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EventsOverlay) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EventsOverlay) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *EventsOverlay) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetAlertTypes

`func (o *EventsOverlay) GetAlertTypes() []string`

GetAlertTypes returns the AlertTypes field if non-nil, zero value otherwise.

### GetAlertTypesOk

`func (o *EventsOverlay) GetAlertTypesOk() (*[]string, bool)`

GetAlertTypesOk returns a tuple with the AlertTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTypes

`func (o *EventsOverlay) SetAlertTypes(v []string)`

SetAlertTypes sets AlertTypes field to given value.

### HasAlertTypes

`func (o *EventsOverlay) HasAlertTypes() bool`

HasAlertTypes returns a boolean if a field has been set.

### GetFilters

`func (o *EventsOverlay) GetFilters() WidgetOngoingAlertsOverlayFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *EventsOverlay) GetFiltersOk() (*WidgetOngoingAlertsOverlayFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *EventsOverlay) SetFilters(v WidgetOngoingAlertsOverlayFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *EventsOverlay) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


