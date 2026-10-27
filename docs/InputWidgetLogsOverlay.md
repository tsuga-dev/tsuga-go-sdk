# InputWidgetLogsOverlay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Draws log events over the chart | 
**Query** | Pointer to **string** | Tsuga query selecting the log events drawn over the chart. All log events in the time range are drawn when omitted. | [optional] 

## Methods

### NewInputWidgetLogsOverlay

`func NewInputWidgetLogsOverlay(type_ string, ) *InputWidgetLogsOverlay`

NewInputWidgetLogsOverlay instantiates a new InputWidgetLogsOverlay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputWidgetLogsOverlayWithDefaults

`func NewInputWidgetLogsOverlayWithDefaults() *InputWidgetLogsOverlay`

NewInputWidgetLogsOverlayWithDefaults instantiates a new InputWidgetLogsOverlay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputWidgetLogsOverlay) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputWidgetLogsOverlay) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputWidgetLogsOverlay) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *InputWidgetLogsOverlay) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *InputWidgetLogsOverlay) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *InputWidgetLogsOverlay) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *InputWidgetLogsOverlay) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


