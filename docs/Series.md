# Series

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Result identifier. Query results use &#x60;q1&#x60;, &#x60;q2&#x60;, and so on; formula results use &#x60;formula&#x60;. | 
**Group** | [**map[string]SeriesGroupValue**](SeriesGroupValue.md) | Group-by field values for this result. Empty when no group-by is requested. | 
**Points** | [**[]SeriesPointsInner**](SeriesPointsInner.md) | Ordered data points for this query or formula and group. | 

## Methods

### NewSeries

`func NewSeries(id string, group map[string]SeriesGroupValue, points []SeriesPointsInner, ) *Series`

NewSeries instantiates a new Series object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesWithDefaults

`func NewSeriesWithDefaults() *Series`

NewSeriesWithDefaults instantiates a new Series object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Series) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Series) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Series) SetId(v string)`

SetId sets Id field to given value.


### GetGroup

`func (o *Series) GetGroup() map[string]SeriesGroupValue`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Series) GetGroupOk() (*map[string]SeriesGroupValue, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Series) SetGroup(v map[string]SeriesGroupValue)`

SetGroup sets Group field to given value.


### GetPoints

`func (o *Series) GetPoints() []SeriesPointsInner`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *Series) GetPointsOk() (*[]SeriesPointsInner, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *Series) SetPoints(v []SeriesPointsInner)`

SetPoints sets Points field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


