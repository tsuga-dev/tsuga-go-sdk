# WidgetListColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | **string** | Attribute displayed as a column in the log list | 
**Normalizer** | Pointer to [**Normalizer**](Normalizer.md) |  | [optional] 

## Methods

### NewWidgetListColumn

`func NewWidgetListColumn(attribute string, ) *WidgetListColumn`

NewWidgetListColumn instantiates a new WidgetListColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetListColumnWithDefaults

`func NewWidgetListColumnWithDefaults() *WidgetListColumn`

NewWidgetListColumnWithDefaults instantiates a new WidgetListColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *WidgetListColumn) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *WidgetListColumn) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *WidgetListColumn) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetNormalizer

`func (o *WidgetListColumn) GetNormalizer() Normalizer`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *WidgetListColumn) GetNormalizerOk() (*Normalizer, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *WidgetListColumn) SetNormalizer(v Normalizer)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *WidgetListColumn) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


