# GraphVisualizationTimeseriesConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the database rows-based aggregation as a time series chart | 
**ConnectionId** | **string** | The ID of the connection to use to query the datastore. | 
**Queries** | **[]string** |  | 
**LegendMode** | Pointer to **string** | Controls whether and how the widget displays legend or series details (e.g. table, legend-only, or no legend) | [optional] 
**Thresholds** | Pointer to [**[]ThresholdMarker**](ThresholdMarker.md) | Threshold markers displayed on the chart | [optional] 
**YAxisSettings** | Pointer to [**GraphVisualizationTimeseriesConnectionYAxisSettings**](GraphVisualizationTimeseriesConnectionYAxisSettings.md) |  | [optional] 
**LineStyleOptions** | Pointer to [**map[string]GraphVisualizationTimeseriesConnectionLineStyleOptionsValue**](GraphVisualizationTimeseriesConnectionLineStyleOptionsValue.md) | Line style of each series, keyed by 1-based query index. The last index is the formula when there is one. Each entry is &#x60;{groupId, lineStyle}&#x60;, where &#x60;groupId&#x60; is &#x60;_q&lt;n&gt;&#x60; for query n or &#x60;_value&#x60; for the formula. For widgets with a single query, only the &#x60;1&#x60; entry is read and it applies to every series. Defaults to regular. | [optional] 
**Normalizer** | Pointer to [**Normalizer**](Normalizer.md) |  | [optional] 
**EventsOverlay** | Pointer to [**EventsOverlay**](EventsOverlay.md) |  | [optional] 

## Methods

### NewGraphVisualizationTimeseriesConnection

`func NewGraphVisualizationTimeseriesConnection(type_ string, connectionId string, queries []string, ) *GraphVisualizationTimeseriesConnection`

NewGraphVisualizationTimeseriesConnection instantiates a new GraphVisualizationTimeseriesConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationTimeseriesConnectionWithDefaults

`func NewGraphVisualizationTimeseriesConnectionWithDefaults() *GraphVisualizationTimeseriesConnection`

NewGraphVisualizationTimeseriesConnectionWithDefaults instantiates a new GraphVisualizationTimeseriesConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GraphVisualizationTimeseriesConnection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphVisualizationTimeseriesConnection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphVisualizationTimeseriesConnection) SetType(v string)`

SetType sets Type field to given value.


### GetConnectionId

`func (o *GraphVisualizationTimeseriesConnection) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *GraphVisualizationTimeseriesConnection) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *GraphVisualizationTimeseriesConnection) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetQueries

`func (o *GraphVisualizationTimeseriesConnection) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *GraphVisualizationTimeseriesConnection) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *GraphVisualizationTimeseriesConnection) SetQueries(v []string)`

SetQueries sets Queries field to given value.


### GetLegendMode

`func (o *GraphVisualizationTimeseriesConnection) GetLegendMode() string`

GetLegendMode returns the LegendMode field if non-nil, zero value otherwise.

### GetLegendModeOk

`func (o *GraphVisualizationTimeseriesConnection) GetLegendModeOk() (*string, bool)`

GetLegendModeOk returns a tuple with the LegendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendMode

`func (o *GraphVisualizationTimeseriesConnection) SetLegendMode(v string)`

SetLegendMode sets LegendMode field to given value.

### HasLegendMode

`func (o *GraphVisualizationTimeseriesConnection) HasLegendMode() bool`

HasLegendMode returns a boolean if a field has been set.

### GetThresholds

`func (o *GraphVisualizationTimeseriesConnection) GetThresholds() []ThresholdMarker`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *GraphVisualizationTimeseriesConnection) GetThresholdsOk() (*[]ThresholdMarker, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *GraphVisualizationTimeseriesConnection) SetThresholds(v []ThresholdMarker)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *GraphVisualizationTimeseriesConnection) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### GetYAxisSettings

`func (o *GraphVisualizationTimeseriesConnection) GetYAxisSettings() GraphVisualizationTimeseriesConnectionYAxisSettings`

GetYAxisSettings returns the YAxisSettings field if non-nil, zero value otherwise.

### GetYAxisSettingsOk

`func (o *GraphVisualizationTimeseriesConnection) GetYAxisSettingsOk() (*GraphVisualizationTimeseriesConnectionYAxisSettings, bool)`

GetYAxisSettingsOk returns a tuple with the YAxisSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisSettings

`func (o *GraphVisualizationTimeseriesConnection) SetYAxisSettings(v GraphVisualizationTimeseriesConnectionYAxisSettings)`

SetYAxisSettings sets YAxisSettings field to given value.

### HasYAxisSettings

`func (o *GraphVisualizationTimeseriesConnection) HasYAxisSettings() bool`

HasYAxisSettings returns a boolean if a field has been set.

### GetLineStyleOptions

`func (o *GraphVisualizationTimeseriesConnection) GetLineStyleOptions() map[string]GraphVisualizationTimeseriesConnectionLineStyleOptionsValue`

GetLineStyleOptions returns the LineStyleOptions field if non-nil, zero value otherwise.

### GetLineStyleOptionsOk

`func (o *GraphVisualizationTimeseriesConnection) GetLineStyleOptionsOk() (*map[string]GraphVisualizationTimeseriesConnectionLineStyleOptionsValue, bool)`

GetLineStyleOptionsOk returns a tuple with the LineStyleOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineStyleOptions

`func (o *GraphVisualizationTimeseriesConnection) SetLineStyleOptions(v map[string]GraphVisualizationTimeseriesConnectionLineStyleOptionsValue)`

SetLineStyleOptions sets LineStyleOptions field to given value.

### HasLineStyleOptions

`func (o *GraphVisualizationTimeseriesConnection) HasLineStyleOptions() bool`

HasLineStyleOptions returns a boolean if a field has been set.

### GetNormalizer

`func (o *GraphVisualizationTimeseriesConnection) GetNormalizer() Normalizer`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *GraphVisualizationTimeseriesConnection) GetNormalizerOk() (*Normalizer, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *GraphVisualizationTimeseriesConnection) SetNormalizer(v Normalizer)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *GraphVisualizationTimeseriesConnection) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetEventsOverlay

`func (o *GraphVisualizationTimeseriesConnection) GetEventsOverlay() EventsOverlay`

GetEventsOverlay returns the EventsOverlay field if non-nil, zero value otherwise.

### GetEventsOverlayOk

`func (o *GraphVisualizationTimeseriesConnection) GetEventsOverlayOk() (*EventsOverlay, bool)`

GetEventsOverlayOk returns a tuple with the EventsOverlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsOverlay

`func (o *GraphVisualizationTimeseriesConnection) SetEventsOverlay(v EventsOverlay)`

SetEventsOverlay sets EventsOverlay field to given value.

### HasEventsOverlay

`func (o *GraphVisualizationTimeseriesConnection) HasEventsOverlay() bool`

HasEventsOverlay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


