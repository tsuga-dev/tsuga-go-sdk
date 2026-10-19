# QueryOngoingAlertsResponseDataAlertsPageInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator identifying an alert raised by a monitor. | 
**Monitor** | [**OngoingMonitorAlertMonitor**](OngoingMonitorAlertMonitor.md) |  | 
**TransitionId** | **string** | ID of the state transition that opened this alert. It stays the same for as long as the alert remains open, so it identifies the alert across successive polls. | 
**ClusterId** | Pointer to **string** | Cluster the alert was raised on. Omitted for alerts that are not scoped to a cluster. | [optional] 
**CurrentState** | **string** | Current alert state. &#x60;alert&#x60; means the condition is breached; &#x60;alert_no_data&#x60; means evaluation stopped receiving data and the monitor is configured to alert on that. | 
**CurrentStateSince** | **string** | Time the alert entered &#x60;currentState&#x60;. Use it as the anchor when building the time range of a query that investigates the alert. | 
**TriggeredBy** | [**[]map[string]OngoingMonitorAlertTriggeredByInnerValue**](map[string]OngoingMonitorAlertTriggeredByInnerValue.md) | Group-by dimension values that triggered the alert, one entry per firing group. Empty when the monitor or SLO has no group-by. Combine these with the asset configuration to scope an investigation query to the firing group. | 
**Slo** | [**OngoingSloAlertSlo**](OngoingSloAlertSlo.md) |  | 
**SloAlert** | [**OngoingSloAlertSloAlert**](OngoingSloAlertSloAlert.md) |  | 

## Methods

### NewQueryOngoingAlertsResponseDataAlertsPageInner

`func NewQueryOngoingAlertsResponseDataAlertsPageInner(type_ string, monitor OngoingMonitorAlertMonitor, transitionId string, currentState string, currentStateSince string, triggeredBy []map[string]OngoingMonitorAlertTriggeredByInnerValue, slo OngoingSloAlertSlo, sloAlert OngoingSloAlertSloAlert, ) *QueryOngoingAlertsResponseDataAlertsPageInner`

NewQueryOngoingAlertsResponseDataAlertsPageInner instantiates a new QueryOngoingAlertsResponseDataAlertsPageInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryOngoingAlertsResponseDataAlertsPageInnerWithDefaults

`func NewQueryOngoingAlertsResponseDataAlertsPageInnerWithDefaults() *QueryOngoingAlertsResponseDataAlertsPageInner`

NewQueryOngoingAlertsResponseDataAlertsPageInnerWithDefaults instantiates a new QueryOngoingAlertsResponseDataAlertsPageInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) SetType(v string)`

SetType sets Type field to given value.


### GetMonitor

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) GetMonitor() OngoingMonitorAlertMonitor`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) GetMonitorOk() (*OngoingMonitorAlertMonitor, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) SetMonitor(v OngoingMonitorAlertMonitor)`

SetMonitor sets Monitor field to given value.


### GetTransitionId

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) GetTransitionId() string`

GetTransitionId returns the TransitionId field if non-nil, zero value otherwise.

### GetTransitionIdOk

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) GetTransitionIdOk() (*string, bool)`

GetTransitionIdOk returns a tuple with the TransitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionId

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) SetTransitionId(v string)`

SetTransitionId sets TransitionId field to given value.


### GetClusterId

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCurrentState

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) GetCurrentState() string`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) GetCurrentStateOk() (*string, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) SetCurrentState(v string)`

SetCurrentState sets CurrentState field to given value.


### GetCurrentStateSince

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) GetCurrentStateSince() string`

GetCurrentStateSince returns the CurrentStateSince field if non-nil, zero value otherwise.

### GetCurrentStateSinceOk

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) GetCurrentStateSinceOk() (*string, bool)`

GetCurrentStateSinceOk returns a tuple with the CurrentStateSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStateSince

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) SetCurrentStateSince(v string)`

SetCurrentStateSince sets CurrentStateSince field to given value.


### GetTriggeredBy

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) GetTriggeredBy() []map[string]OngoingMonitorAlertTriggeredByInnerValue`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) GetTriggeredByOk() (*[]map[string]OngoingMonitorAlertTriggeredByInnerValue, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) SetTriggeredBy(v []map[string]OngoingMonitorAlertTriggeredByInnerValue)`

SetTriggeredBy sets TriggeredBy field to given value.


### GetSlo

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) GetSlo() OngoingSloAlertSlo`

GetSlo returns the Slo field if non-nil, zero value otherwise.

### GetSloOk

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) GetSloOk() (*OngoingSloAlertSlo, bool)`

GetSloOk returns a tuple with the Slo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlo

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) SetSlo(v OngoingSloAlertSlo)`

SetSlo sets Slo field to given value.


### GetSloAlert

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) GetSloAlert() OngoingSloAlertSloAlert`

GetSloAlert returns the SloAlert field if non-nil, zero value otherwise.

### GetSloAlertOk

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) GetSloAlertOk() (*OngoingSloAlertSloAlert, bool)`

GetSloAlertOk returns a tuple with the SloAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloAlert

`func (o *QueryOngoingAlertsResponseDataAlertsPageInner) SetSloAlert(v OngoingSloAlertSloAlert)`

SetSloAlert sets SloAlert field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


