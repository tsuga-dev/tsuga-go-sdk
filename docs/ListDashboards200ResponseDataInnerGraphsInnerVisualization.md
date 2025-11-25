# ListDashboards200ResponseDataInnerGraphsInnerVisualization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the aggregation as a time series chart | 
**Source** | **string** | Indicates that the widget queries logs | 
**Queries** | [**[]GraphVisualizationTimeseriesQueriesInner**](GraphVisualizationTimeseriesQueriesInner.md) | Aggregations that may be combined together in the same query | 
**Formula** | Pointer to **string** | Formula referencing query outputs (e.g. q1+q2) to compute derived series | [optional] 
**VisibleSeries** | Pointer to **[]bool** | Flags indicating whether each query or formula series is visible | [optional] 
**GroupBy** | Pointer to [**[]GraphVisualizationTimeseriesGroupByInner**](GraphVisualizationTimeseriesGroupByInner.md) | Fields used to group the results | [optional] 
**Normalizer** | Pointer to [**GraphVisualizationTimeseriesNormalizer**](GraphVisualizationTimeseriesNormalizer.md) |  | [optional] 
**BackgroundMode** | Pointer to **string** | Controls whether the widget uses a solid or transparent background | [optional] 
**Conditions** | Pointer to [**[]GraphVisualizationQueryValueConditionsInner**](GraphVisualizationQueryValueConditionsInner.md) | Conditional formatting rules applied to the displayed value | [optional] 
**TimeBucket** | Pointer to [**GraphVisualizationBarTimeBucket**](GraphVisualizationBarTimeBucket.md) |  | [optional] 
**Query** | **string** | Log query that selects records for the list | 
**ListColumns** | Pointer to [**[]GraphVisualizationListListColumnsInner**](GraphVisualizationListListColumnsInner.md) | Custom columns to display for each log entry | [optional] 
**Note** | Pointer to **string** | Markdown-compatible text shown in the note | [optional] 
**NoteColor** | Pointer to **string** | Background color used to render the note | [optional] 
**NoteAlign** | Pointer to **string** | Flex alignment keyword used for widget layout | [optional] 
**NoteJustifyContent** | Pointer to **string** | Flex alignment keyword used for widget layout | [optional] 

## Methods

### NewListDashboards200ResponseDataInnerGraphsInnerVisualization

`func NewListDashboards200ResponseDataInnerGraphsInnerVisualization(type_ string, source string, queries []GraphVisualizationTimeseriesQueriesInner, query string, ) *ListDashboards200ResponseDataInnerGraphsInnerVisualization`

NewListDashboards200ResponseDataInnerGraphsInnerVisualization instantiates a new ListDashboards200ResponseDataInnerGraphsInnerVisualization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDashboards200ResponseDataInnerGraphsInnerVisualizationWithDefaults

`func NewListDashboards200ResponseDataInnerGraphsInnerVisualizationWithDefaults() *ListDashboards200ResponseDataInnerGraphsInnerVisualization`

NewListDashboards200ResponseDataInnerGraphsInnerVisualizationWithDefaults instantiates a new ListDashboards200ResponseDataInnerGraphsInnerVisualization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) SetType(v string)`

SetType sets Type field to given value.


### GetSource

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) SetSource(v string)`

SetSource sets Source field to given value.


### GetQueries

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetQueries() []GraphVisualizationTimeseriesQueriesInner`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetQueriesOk() (*[]GraphVisualizationTimeseriesQueriesInner, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) SetQueries(v []GraphVisualizationTimeseriesQueriesInner)`

SetQueries sets Queries field to given value.


### GetFormula

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetVisibleSeries

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetVisibleSeries() []bool`

GetVisibleSeries returns the VisibleSeries field if non-nil, zero value otherwise.

### GetVisibleSeriesOk

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetVisibleSeriesOk() (*[]bool, bool)`

GetVisibleSeriesOk returns a tuple with the VisibleSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleSeries

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) SetVisibleSeries(v []bool)`

SetVisibleSeries sets VisibleSeries field to given value.

### HasVisibleSeries

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) HasVisibleSeries() bool`

HasVisibleSeries returns a boolean if a field has been set.

### GetGroupBy

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetGroupBy() []GraphVisualizationTimeseriesGroupByInner`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetGroupByOk() (*[]GraphVisualizationTimeseriesGroupByInner, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) SetGroupBy(v []GraphVisualizationTimeseriesGroupByInner)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetNormalizer

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetNormalizer() GraphVisualizationTimeseriesNormalizer`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetNormalizerOk() (*GraphVisualizationTimeseriesNormalizer, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) SetNormalizer(v GraphVisualizationTimeseriesNormalizer)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetBackgroundMode

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetBackgroundMode() string`

GetBackgroundMode returns the BackgroundMode field if non-nil, zero value otherwise.

### GetBackgroundModeOk

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetBackgroundModeOk() (*string, bool)`

GetBackgroundModeOk returns a tuple with the BackgroundMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundMode

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) SetBackgroundMode(v string)`

SetBackgroundMode sets BackgroundMode field to given value.

### HasBackgroundMode

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) HasBackgroundMode() bool`

HasBackgroundMode returns a boolean if a field has been set.

### GetConditions

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetConditions() []GraphVisualizationQueryValueConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetConditionsOk() (*[]GraphVisualizationQueryValueConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) SetConditions(v []GraphVisualizationQueryValueConditionsInner)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetTimeBucket

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetTimeBucket() GraphVisualizationBarTimeBucket`

GetTimeBucket returns the TimeBucket field if non-nil, zero value otherwise.

### GetTimeBucketOk

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetTimeBucketOk() (*GraphVisualizationBarTimeBucket, bool)`

GetTimeBucketOk returns a tuple with the TimeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBucket

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) SetTimeBucket(v GraphVisualizationBarTimeBucket)`

SetTimeBucket sets TimeBucket field to given value.

### HasTimeBucket

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) HasTimeBucket() bool`

HasTimeBucket returns a boolean if a field has been set.

### GetQuery

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetListColumns

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetListColumns() []GraphVisualizationListListColumnsInner`

GetListColumns returns the ListColumns field if non-nil, zero value otherwise.

### GetListColumnsOk

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetListColumnsOk() (*[]GraphVisualizationListListColumnsInner, bool)`

GetListColumnsOk returns a tuple with the ListColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListColumns

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) SetListColumns(v []GraphVisualizationListListColumnsInner)`

SetListColumns sets ListColumns field to given value.

### HasListColumns

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) HasListColumns() bool`

HasListColumns returns a boolean if a field has been set.

### GetNote

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetNoteColor

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetNoteColor() string`

GetNoteColor returns the NoteColor field if non-nil, zero value otherwise.

### GetNoteColorOk

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetNoteColorOk() (*string, bool)`

GetNoteColorOk returns a tuple with the NoteColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteColor

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) SetNoteColor(v string)`

SetNoteColor sets NoteColor field to given value.

### HasNoteColor

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) HasNoteColor() bool`

HasNoteColor returns a boolean if a field has been set.

### GetNoteAlign

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetNoteAlign() string`

GetNoteAlign returns the NoteAlign field if non-nil, zero value otherwise.

### GetNoteAlignOk

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetNoteAlignOk() (*string, bool)`

GetNoteAlignOk returns a tuple with the NoteAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteAlign

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) SetNoteAlign(v string)`

SetNoteAlign sets NoteAlign field to given value.

### HasNoteAlign

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) HasNoteAlign() bool`

HasNoteAlign returns a boolean if a field has been set.

### GetNoteJustifyContent

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetNoteJustifyContent() string`

GetNoteJustifyContent returns the NoteJustifyContent field if non-nil, zero value otherwise.

### GetNoteJustifyContentOk

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) GetNoteJustifyContentOk() (*string, bool)`

GetNoteJustifyContentOk returns a tuple with the NoteJustifyContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteJustifyContent

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) SetNoteJustifyContent(v string)`

SetNoteJustifyContent sets NoteJustifyContent field to given value.

### HasNoteJustifyContent

`func (o *ListDashboards200ResponseDataInnerGraphsInnerVisualization) HasNoteJustifyContent() bool`

HasNoteJustifyContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


