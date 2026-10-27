# WidgetLogsOverlay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Draws log events over the chart | 
**Query** | Pointer to **string** | Tsuga query selecting the log events drawn over the chart. All log events in the time range are drawn when omitted. | [optional] 

## Methods

### NewWidgetLogsOverlay

`func NewWidgetLogsOverlay(type_ string, ) *WidgetLogsOverlay`

NewWidgetLogsOverlay instantiates a new WidgetLogsOverlay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetLogsOverlayWithDefaults

`func NewWidgetLogsOverlayWithDefaults() *WidgetLogsOverlay`

NewWidgetLogsOverlayWithDefaults instantiates a new WidgetLogsOverlay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WidgetLogsOverlay) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WidgetLogsOverlay) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WidgetLogsOverlay) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *WidgetLogsOverlay) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *WidgetLogsOverlay) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *WidgetLogsOverlay) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *WidgetLogsOverlay) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


