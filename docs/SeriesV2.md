# SeriesV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**map[string]SeriesGroupValue**](SeriesGroupValue.md) | Group-by field values for this row, keyed by field name. Empty when no group-by is requested. | 
**Values** | **map[string][]*float32** | One value per entry of &#x60;timestamps&#x60;, &#x60;null&#x60; where the bucket has none. Keyed by result id: &#x60;q1&#x60;, &#x60;q2&#x60;, and so on for each query in submission order, plus &#x60;formula&#x60; when a computed formula is given, or &#x60;value&#x60; for a PromQL query. Every row has the same keys. | 

## Methods

### NewSeriesV2

`func NewSeriesV2(group map[string]SeriesGroupValue, values map[string][]*float32, ) *SeriesV2`

NewSeriesV2 instantiates a new SeriesV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesV2WithDefaults

`func NewSeriesV2WithDefaults() *SeriesV2`

NewSeriesV2WithDefaults instantiates a new SeriesV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *SeriesV2) GetGroup() map[string]SeriesGroupValue`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SeriesV2) GetGroupOk() (*map[string]SeriesGroupValue, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SeriesV2) SetGroup(v map[string]SeriesGroupValue)`

SetGroup sets Group field to given value.


### GetValues

`func (o *SeriesV2) GetValues() map[string][]*float32`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SeriesV2) GetValuesOk() (*map[string][]*float32, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SeriesV2) SetValues(v map[string][]*float32)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


