# InputWidgetListColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | **string** | Attribute (log field or database row column) displayed as a column in the list | 
**Normalizer** | Pointer to [**InputNormalizer**](InputNormalizer.md) |  | [optional] 

## Methods

### NewInputWidgetListColumn

`func NewInputWidgetListColumn(attribute string, ) *InputWidgetListColumn`

NewInputWidgetListColumn instantiates a new InputWidgetListColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputWidgetListColumnWithDefaults

`func NewInputWidgetListColumnWithDefaults() *InputWidgetListColumn`

NewInputWidgetListColumnWithDefaults instantiates a new InputWidgetListColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *InputWidgetListColumn) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *InputWidgetListColumn) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *InputWidgetListColumn) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetNormalizer

`func (o *InputWidgetListColumn) GetNormalizer() InputNormalizer`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *InputWidgetListColumn) GetNormalizerOk() (*InputNormalizer, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *InputWidgetListColumn) SetNormalizer(v InputNormalizer)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *InputWidgetListColumn) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


