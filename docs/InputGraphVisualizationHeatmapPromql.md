# InputGraphVisualizationHeatmapPromql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays a PromQL histogram query grouped by &#x60;le&#x60; as a heatmap chart, e.g. &#x60;sum by (le) (rate(http_duration_bucket[5m]))&#x60; | 
**Queries** | **[]string** | PromQL expressions configured for this widget. Only the first expression is rendered; later entries are stored and ignored. | 
**Precision** | Pointer to [**GraphVisualizationQueryValueConnectionPrecision**](GraphVisualizationQueryValueConnectionPrecision.md) |  | [optional] 
**Normalizer** | Pointer to [**InputNormalizer**](InputNormalizer.md) |  | [optional] 
**Palette** | Pointer to **string** | Color palette used to render the heatmap intensity gradient | [optional] 
**YAxisSettings** | Pointer to [**InputGraphVisualizationTimeseriesConnectionYAxisSettings**](InputGraphVisualizationTimeseriesConnectionYAxisSettings.md) |  | [optional] 

## Methods

### NewInputGraphVisualizationHeatmapPromql

`func NewInputGraphVisualizationHeatmapPromql(type_ string, queries []string, ) *InputGraphVisualizationHeatmapPromql`

NewInputGraphVisualizationHeatmapPromql instantiates a new InputGraphVisualizationHeatmapPromql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationHeatmapPromqlWithDefaults

`func NewInputGraphVisualizationHeatmapPromqlWithDefaults() *InputGraphVisualizationHeatmapPromql`

NewInputGraphVisualizationHeatmapPromqlWithDefaults instantiates a new InputGraphVisualizationHeatmapPromql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputGraphVisualizationHeatmapPromql) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputGraphVisualizationHeatmapPromql) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputGraphVisualizationHeatmapPromql) SetType(v string)`

SetType sets Type field to given value.


### GetQueries

`func (o *InputGraphVisualizationHeatmapPromql) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *InputGraphVisualizationHeatmapPromql) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *InputGraphVisualizationHeatmapPromql) SetQueries(v []string)`

SetQueries sets Queries field to given value.


### GetPrecision

`func (o *InputGraphVisualizationHeatmapPromql) GetPrecision() GraphVisualizationQueryValueConnectionPrecision`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *InputGraphVisualizationHeatmapPromql) GetPrecisionOk() (*GraphVisualizationQueryValueConnectionPrecision, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *InputGraphVisualizationHeatmapPromql) SetPrecision(v GraphVisualizationQueryValueConnectionPrecision)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *InputGraphVisualizationHeatmapPromql) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetNormalizer

`func (o *InputGraphVisualizationHeatmapPromql) GetNormalizer() InputNormalizer`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *InputGraphVisualizationHeatmapPromql) GetNormalizerOk() (*InputNormalizer, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *InputGraphVisualizationHeatmapPromql) SetNormalizer(v InputNormalizer)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *InputGraphVisualizationHeatmapPromql) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetPalette

`func (o *InputGraphVisualizationHeatmapPromql) GetPalette() string`

GetPalette returns the Palette field if non-nil, zero value otherwise.

### GetPaletteOk

`func (o *InputGraphVisualizationHeatmapPromql) GetPaletteOk() (*string, bool)`

GetPaletteOk returns a tuple with the Palette field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPalette

`func (o *InputGraphVisualizationHeatmapPromql) SetPalette(v string)`

SetPalette sets Palette field to given value.

### HasPalette

`func (o *InputGraphVisualizationHeatmapPromql) HasPalette() bool`

HasPalette returns a boolean if a field has been set.

### GetYAxisSettings

`func (o *InputGraphVisualizationHeatmapPromql) GetYAxisSettings() InputGraphVisualizationTimeseriesConnectionYAxisSettings`

GetYAxisSettings returns the YAxisSettings field if non-nil, zero value otherwise.

### GetYAxisSettingsOk

`func (o *InputGraphVisualizationHeatmapPromql) GetYAxisSettingsOk() (*InputGraphVisualizationTimeseriesConnectionYAxisSettings, bool)`

GetYAxisSettingsOk returns a tuple with the YAxisSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisSettings

`func (o *InputGraphVisualizationHeatmapPromql) SetYAxisSettings(v InputGraphVisualizationTimeseriesConnectionYAxisSettings)`

SetYAxisSettings sets YAxisSettings field to given value.

### HasYAxisSettings

`func (o *InputGraphVisualizationHeatmapPromql) HasYAxisSettings() bool`

HasYAxisSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


