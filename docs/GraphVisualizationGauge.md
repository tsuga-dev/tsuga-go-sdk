# GraphVisualizationGauge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the aggregation as a gauge | 
**Source** | **string** | Telemetry source queried by this aggregation: &#x60;logs&#x60;, &#x60;metrics&#x60;, &#x60;traces&#x60;, or &#x60;rum&#x60;. | 
**Queries** | [**[]AggregationQuery**](AggregationQuery.md) | Aggregations that may be combined together in the same query. Each item is referenced from &#x60;formula&#x60; as q1, q2, and so on, in submission order. | 
**Formula** | Pointer to **string** | Formula referencing submitted query outputs, such as &#x60;q1 + q2&#x60;. References must be within &#x60;q1&#x60; through &#x60;qN&#x60; for the submitted queries. | [optional] 
**Aliases** | Pointer to [**GraphVisualizationTimeseriesPromqlAliases**](GraphVisualizationTimeseriesPromqlAliases.md) |  | [optional] 
**VisibleSeries** | Pointer to **[]bool** | Flags indicating whether each query or formula series is visible | [optional] 
**Max** | Pointer to **float32** | Gauge maximum value | [optional] 
**ColorThresholds** | Pointer to [**[]GaugeColorThreshold**](GaugeColorThreshold.md) | Color thresholds inside the gauge range | [optional] 
**Precision** | Pointer to **float32** | Number of decimal places to display in the value | [optional] 
**Normalizer** | Pointer to [**Normalizer**](Normalizer.md) |  | [optional] 

## Methods

### NewGraphVisualizationGauge

`func NewGraphVisualizationGauge(type_ string, source string, queries []AggregationQuery, ) *GraphVisualizationGauge`

NewGraphVisualizationGauge instantiates a new GraphVisualizationGauge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationGaugeWithDefaults

`func NewGraphVisualizationGaugeWithDefaults() *GraphVisualizationGauge`

NewGraphVisualizationGaugeWithDefaults instantiates a new GraphVisualizationGauge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GraphVisualizationGauge) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphVisualizationGauge) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphVisualizationGauge) SetType(v string)`

SetType sets Type field to given value.


### GetSource

`func (o *GraphVisualizationGauge) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GraphVisualizationGauge) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GraphVisualizationGauge) SetSource(v string)`

SetSource sets Source field to given value.


### GetQueries

`func (o *GraphVisualizationGauge) GetQueries() []AggregationQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *GraphVisualizationGauge) GetQueriesOk() (*[]AggregationQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *GraphVisualizationGauge) SetQueries(v []AggregationQuery)`

SetQueries sets Queries field to given value.


### GetFormula

`func (o *GraphVisualizationGauge) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *GraphVisualizationGauge) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *GraphVisualizationGauge) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *GraphVisualizationGauge) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetAliases

`func (o *GraphVisualizationGauge) GetAliases() GraphVisualizationTimeseriesPromqlAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *GraphVisualizationGauge) GetAliasesOk() (*GraphVisualizationTimeseriesPromqlAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *GraphVisualizationGauge) SetAliases(v GraphVisualizationTimeseriesPromqlAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *GraphVisualizationGauge) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetVisibleSeries

`func (o *GraphVisualizationGauge) GetVisibleSeries() []bool`

GetVisibleSeries returns the VisibleSeries field if non-nil, zero value otherwise.

### GetVisibleSeriesOk

`func (o *GraphVisualizationGauge) GetVisibleSeriesOk() (*[]bool, bool)`

GetVisibleSeriesOk returns a tuple with the VisibleSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleSeries

`func (o *GraphVisualizationGauge) SetVisibleSeries(v []bool)`

SetVisibleSeries sets VisibleSeries field to given value.

### HasVisibleSeries

`func (o *GraphVisualizationGauge) HasVisibleSeries() bool`

HasVisibleSeries returns a boolean if a field has been set.

### GetMax

`func (o *GraphVisualizationGauge) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GraphVisualizationGauge) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GraphVisualizationGauge) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *GraphVisualizationGauge) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetColorThresholds

`func (o *GraphVisualizationGauge) GetColorThresholds() []GaugeColorThreshold`

GetColorThresholds returns the ColorThresholds field if non-nil, zero value otherwise.

### GetColorThresholdsOk

`func (o *GraphVisualizationGauge) GetColorThresholdsOk() (*[]GaugeColorThreshold, bool)`

GetColorThresholdsOk returns a tuple with the ColorThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorThresholds

`func (o *GraphVisualizationGauge) SetColorThresholds(v []GaugeColorThreshold)`

SetColorThresholds sets ColorThresholds field to given value.

### HasColorThresholds

`func (o *GraphVisualizationGauge) HasColorThresholds() bool`

HasColorThresholds returns a boolean if a field has been set.

### GetPrecision

`func (o *GraphVisualizationGauge) GetPrecision() float32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *GraphVisualizationGauge) GetPrecisionOk() (*float32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *GraphVisualizationGauge) SetPrecision(v float32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *GraphVisualizationGauge) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetNormalizer

`func (o *GraphVisualizationGauge) GetNormalizer() Normalizer`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *GraphVisualizationGauge) GetNormalizerOk() (*Normalizer, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *GraphVisualizationGauge) SetNormalizer(v Normalizer)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *GraphVisualizationGauge) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


