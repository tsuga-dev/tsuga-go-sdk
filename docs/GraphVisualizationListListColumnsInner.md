# GraphVisualizationListListColumnsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | **string** | Attribute displayed as a column in the log list | 
**Normalizer** | Pointer to [**GraphVisualizationTimeseriesNormalizer**](GraphVisualizationTimeseriesNormalizer.md) |  | [optional] 

## Methods

### NewGraphVisualizationListListColumnsInner

`func NewGraphVisualizationListListColumnsInner(attribute string, ) *GraphVisualizationListListColumnsInner`

NewGraphVisualizationListListColumnsInner instantiates a new GraphVisualizationListListColumnsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationListListColumnsInnerWithDefaults

`func NewGraphVisualizationListListColumnsInnerWithDefaults() *GraphVisualizationListListColumnsInner`

NewGraphVisualizationListListColumnsInnerWithDefaults instantiates a new GraphVisualizationListListColumnsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *GraphVisualizationListListColumnsInner) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *GraphVisualizationListListColumnsInner) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *GraphVisualizationListListColumnsInner) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetNormalizer

`func (o *GraphVisualizationListListColumnsInner) GetNormalizer() GraphVisualizationTimeseriesNormalizer`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *GraphVisualizationListListColumnsInner) GetNormalizerOk() (*GraphVisualizationTimeseriesNormalizer, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *GraphVisualizationListListColumnsInner) SetNormalizer(v GraphVisualizationTimeseriesNormalizer)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *GraphVisualizationListListColumnsInner) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


