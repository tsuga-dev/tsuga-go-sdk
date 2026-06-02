# InputGraphVisualizationTimeseriesConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the database rows-based aggregation as a time series chart | 
**ConnectionId** | **string** | The ID of the connection to use to query the datastore. | 
**Queries** | **[]string** |  | 
**LegendMode** | Pointer to **string** | Controls whether and how the widget displays legend or series details (e.g. table, legend-only, or no legend) | [optional] 
**Thresholds** | Pointer to [**[]ThresholdMarker**](ThresholdMarker.md) | Threshold markers displayed on the chart | [optional] 
**YAxisSettings** | Pointer to [**InputGraphVisualizationTimeseriesConnectionYAxisSettings**](InputGraphVisualizationTimeseriesConnectionYAxisSettings.md) |  | [optional] 

## Methods

### NewInputGraphVisualizationTimeseriesConnection

`func NewInputGraphVisualizationTimeseriesConnection(type_ string, connectionId string, queries []string, ) *InputGraphVisualizationTimeseriesConnection`

NewInputGraphVisualizationTimeseriesConnection instantiates a new InputGraphVisualizationTimeseriesConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationTimeseriesConnectionWithDefaults

`func NewInputGraphVisualizationTimeseriesConnectionWithDefaults() *InputGraphVisualizationTimeseriesConnection`

NewInputGraphVisualizationTimeseriesConnectionWithDefaults instantiates a new InputGraphVisualizationTimeseriesConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputGraphVisualizationTimeseriesConnection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputGraphVisualizationTimeseriesConnection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputGraphVisualizationTimeseriesConnection) SetType(v string)`

SetType sets Type field to given value.


### GetConnectionId

`func (o *InputGraphVisualizationTimeseriesConnection) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *InputGraphVisualizationTimeseriesConnection) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *InputGraphVisualizationTimeseriesConnection) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetQueries

`func (o *InputGraphVisualizationTimeseriesConnection) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *InputGraphVisualizationTimeseriesConnection) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *InputGraphVisualizationTimeseriesConnection) SetQueries(v []string)`

SetQueries sets Queries field to given value.


### GetLegendMode

`func (o *InputGraphVisualizationTimeseriesConnection) GetLegendMode() string`

GetLegendMode returns the LegendMode field if non-nil, zero value otherwise.

### GetLegendModeOk

`func (o *InputGraphVisualizationTimeseriesConnection) GetLegendModeOk() (*string, bool)`

GetLegendModeOk returns a tuple with the LegendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendMode

`func (o *InputGraphVisualizationTimeseriesConnection) SetLegendMode(v string)`

SetLegendMode sets LegendMode field to given value.

### HasLegendMode

`func (o *InputGraphVisualizationTimeseriesConnection) HasLegendMode() bool`

HasLegendMode returns a boolean if a field has been set.

### GetThresholds

`func (o *InputGraphVisualizationTimeseriesConnection) GetThresholds() []ThresholdMarker`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *InputGraphVisualizationTimeseriesConnection) GetThresholdsOk() (*[]ThresholdMarker, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *InputGraphVisualizationTimeseriesConnection) SetThresholds(v []ThresholdMarker)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *InputGraphVisualizationTimeseriesConnection) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### GetYAxisSettings

`func (o *InputGraphVisualizationTimeseriesConnection) GetYAxisSettings() InputGraphVisualizationTimeseriesConnectionYAxisSettings`

GetYAxisSettings returns the YAxisSettings field if non-nil, zero value otherwise.

### GetYAxisSettingsOk

`func (o *InputGraphVisualizationTimeseriesConnection) GetYAxisSettingsOk() (*InputGraphVisualizationTimeseriesConnectionYAxisSettings, bool)`

GetYAxisSettingsOk returns a tuple with the YAxisSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisSettings

`func (o *InputGraphVisualizationTimeseriesConnection) SetYAxisSettings(v InputGraphVisualizationTimeseriesConnectionYAxisSettings)`

SetYAxisSettings sets YAxisSettings field to given value.

### HasYAxisSettings

`func (o *InputGraphVisualizationTimeseriesConnection) HasYAxisSettings() bool`

HasYAxisSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


