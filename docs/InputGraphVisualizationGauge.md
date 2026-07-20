# InputGraphVisualizationGauge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the aggregation as a gauge | 
**Source** | **string** | Telemetry source queried by this aggregation: &#x60;logs&#x60;, &#x60;metrics&#x60;, &#x60;traces&#x60;, or &#x60;rum&#x60;. | 
**Queries** | [**[]AggregationQuery1**](AggregationQuery1.md) | Aggregations that may be combined together in the same query. Each item is referenced from &#x60;formula&#x60; as q1, q2, and so on, in submission order. Limited to 15 items. For dataSource \&quot;metrics\&quot;, each aggregate&#39;s &#x60;field&#x60; is the metric name, not an attribute; to count distinct values of an attribute use unique-count with field \&quot;&lt;metricName&gt;.context.&lt;attribute&gt;\&quot; (e.g. \&quot;system.cpu.utilization.context.host.name\&quot;). | 
**Formula** | Pointer to **string** | Formula referencing query outputs, such as &#x60;q1 + q2&#x60;, to compute derived results. Defaults to &#x60;q1&#x60;. Formulas may reference only submitted queries (&#x60;q1&#x60; through &#x60;qN&#x60;); undefined query references return 400. | [optional] 
**Aliases** | Pointer to [**InputGraphVisualizationTimeseriesPromqlAliases**](InputGraphVisualizationTimeseriesPromqlAliases.md) |  | [optional] 
**VisibleSeries** | Pointer to **[]bool** | Flags indicating whether each query or formula series is visible | [optional] 
**Max** | Pointer to **float32** | Gauge maximum value | [optional] 
**ColorThresholds** | Pointer to [**[]GaugeColorThreshold**](GaugeColorThreshold.md) | Color thresholds inside the gauge range | [optional] 
**Precision** | Pointer to **float32** | Number of decimal places to display in the value | [optional] 
**Normalizer** | Pointer to [**Normalizer1**](Normalizer1.md) |  | [optional] 

## Methods

### NewInputGraphVisualizationGauge

`func NewInputGraphVisualizationGauge(type_ string, source string, queries []AggregationQuery1, ) *InputGraphVisualizationGauge`

NewInputGraphVisualizationGauge instantiates a new InputGraphVisualizationGauge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationGaugeWithDefaults

`func NewInputGraphVisualizationGaugeWithDefaults() *InputGraphVisualizationGauge`

NewInputGraphVisualizationGaugeWithDefaults instantiates a new InputGraphVisualizationGauge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputGraphVisualizationGauge) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputGraphVisualizationGauge) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputGraphVisualizationGauge) SetType(v string)`

SetType sets Type field to given value.


### GetSource

`func (o *InputGraphVisualizationGauge) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InputGraphVisualizationGauge) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InputGraphVisualizationGauge) SetSource(v string)`

SetSource sets Source field to given value.


### GetQueries

`func (o *InputGraphVisualizationGauge) GetQueries() []AggregationQuery1`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *InputGraphVisualizationGauge) GetQueriesOk() (*[]AggregationQuery1, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *InputGraphVisualizationGauge) SetQueries(v []AggregationQuery1)`

SetQueries sets Queries field to given value.


### GetFormula

`func (o *InputGraphVisualizationGauge) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *InputGraphVisualizationGauge) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *InputGraphVisualizationGauge) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *InputGraphVisualizationGauge) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetAliases

`func (o *InputGraphVisualizationGauge) GetAliases() InputGraphVisualizationTimeseriesPromqlAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *InputGraphVisualizationGauge) GetAliasesOk() (*InputGraphVisualizationTimeseriesPromqlAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *InputGraphVisualizationGauge) SetAliases(v InputGraphVisualizationTimeseriesPromqlAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *InputGraphVisualizationGauge) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetVisibleSeries

`func (o *InputGraphVisualizationGauge) GetVisibleSeries() []bool`

GetVisibleSeries returns the VisibleSeries field if non-nil, zero value otherwise.

### GetVisibleSeriesOk

`func (o *InputGraphVisualizationGauge) GetVisibleSeriesOk() (*[]bool, bool)`

GetVisibleSeriesOk returns a tuple with the VisibleSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleSeries

`func (o *InputGraphVisualizationGauge) SetVisibleSeries(v []bool)`

SetVisibleSeries sets VisibleSeries field to given value.

### HasVisibleSeries

`func (o *InputGraphVisualizationGauge) HasVisibleSeries() bool`

HasVisibleSeries returns a boolean if a field has been set.

### GetMax

`func (o *InputGraphVisualizationGauge) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *InputGraphVisualizationGauge) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *InputGraphVisualizationGauge) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *InputGraphVisualizationGauge) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetColorThresholds

`func (o *InputGraphVisualizationGauge) GetColorThresholds() []GaugeColorThreshold`

GetColorThresholds returns the ColorThresholds field if non-nil, zero value otherwise.

### GetColorThresholdsOk

`func (o *InputGraphVisualizationGauge) GetColorThresholdsOk() (*[]GaugeColorThreshold, bool)`

GetColorThresholdsOk returns a tuple with the ColorThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorThresholds

`func (o *InputGraphVisualizationGauge) SetColorThresholds(v []GaugeColorThreshold)`

SetColorThresholds sets ColorThresholds field to given value.

### HasColorThresholds

`func (o *InputGraphVisualizationGauge) HasColorThresholds() bool`

HasColorThresholds returns a boolean if a field has been set.

### GetPrecision

`func (o *InputGraphVisualizationGauge) GetPrecision() float32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *InputGraphVisualizationGauge) GetPrecisionOk() (*float32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *InputGraphVisualizationGauge) SetPrecision(v float32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *InputGraphVisualizationGauge) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetNormalizer

`func (o *InputGraphVisualizationGauge) GetNormalizer() Normalizer1`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *InputGraphVisualizationGauge) GetNormalizerOk() (*Normalizer1, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *InputGraphVisualizationGauge) SetNormalizer(v Normalizer1)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *InputGraphVisualizationGauge) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


