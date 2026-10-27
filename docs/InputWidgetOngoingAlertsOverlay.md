# InputWidgetOngoingAlertsOverlay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Draws monitor and SLO alerts over the chart | 
**AlertTypes** | Pointer to **[]string** | Kinds of alert drawn over the chart. Both monitor and SLO alerts are drawn when omitted. | [optional] 
**Query** | Pointer to **string** | Case-insensitive substring matched against the alert, the same way the Ongoing alerts page matches its search box. Every alert is drawn when omitted. | [optional] 
**Filters** | Pointer to [**InputWidgetOngoingAlertsOverlayFilters**](InputWidgetOngoingAlertsOverlayFilters.md) |  | [optional] 

## Methods

### NewInputWidgetOngoingAlertsOverlay

`func NewInputWidgetOngoingAlertsOverlay(type_ string, ) *InputWidgetOngoingAlertsOverlay`

NewInputWidgetOngoingAlertsOverlay instantiates a new InputWidgetOngoingAlertsOverlay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputWidgetOngoingAlertsOverlayWithDefaults

`func NewInputWidgetOngoingAlertsOverlayWithDefaults() *InputWidgetOngoingAlertsOverlay`

NewInputWidgetOngoingAlertsOverlayWithDefaults instantiates a new InputWidgetOngoingAlertsOverlay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputWidgetOngoingAlertsOverlay) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputWidgetOngoingAlertsOverlay) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputWidgetOngoingAlertsOverlay) SetType(v string)`

SetType sets Type field to given value.


### GetAlertTypes

`func (o *InputWidgetOngoingAlertsOverlay) GetAlertTypes() []string`

GetAlertTypes returns the AlertTypes field if non-nil, zero value otherwise.

### GetAlertTypesOk

`func (o *InputWidgetOngoingAlertsOverlay) GetAlertTypesOk() (*[]string, bool)`

GetAlertTypesOk returns a tuple with the AlertTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTypes

`func (o *InputWidgetOngoingAlertsOverlay) SetAlertTypes(v []string)`

SetAlertTypes sets AlertTypes field to given value.

### HasAlertTypes

`func (o *InputWidgetOngoingAlertsOverlay) HasAlertTypes() bool`

HasAlertTypes returns a boolean if a field has been set.

### GetQuery

`func (o *InputWidgetOngoingAlertsOverlay) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *InputWidgetOngoingAlertsOverlay) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *InputWidgetOngoingAlertsOverlay) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *InputWidgetOngoingAlertsOverlay) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetFilters

`func (o *InputWidgetOngoingAlertsOverlay) GetFilters() InputWidgetOngoingAlertsOverlayFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *InputWidgetOngoingAlertsOverlay) GetFiltersOk() (*InputWidgetOngoingAlertsOverlayFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *InputWidgetOngoingAlertsOverlay) SetFilters(v InputWidgetOngoingAlertsOverlayFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *InputWidgetOngoingAlertsOverlay) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


