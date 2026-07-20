# QueryMonitorsRequestSort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**By** | **string** | Field to sort the returned monitors by. For &#x60;priority&#x60;, &#x60;desc&#x60; returns highest-priority monitors first because priority 1 is highest. | 
**Direction** | **string** | Sort direction. For &#x60;priority&#x60;, &#x60;desc&#x60; returns priority 1 first and &#x60;asc&#x60; returns priority 5 first. | 

## Methods

### NewQueryMonitorsRequestSort

`func NewQueryMonitorsRequestSort(by string, direction string, ) *QueryMonitorsRequestSort`

NewQueryMonitorsRequestSort instantiates a new QueryMonitorsRequestSort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMonitorsRequestSortWithDefaults

`func NewQueryMonitorsRequestSortWithDefaults() *QueryMonitorsRequestSort`

NewQueryMonitorsRequestSortWithDefaults instantiates a new QueryMonitorsRequestSort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBy

`func (o *QueryMonitorsRequestSort) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *QueryMonitorsRequestSort) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *QueryMonitorsRequestSort) SetBy(v string)`

SetBy sets By field to given value.


### GetDirection

`func (o *QueryMonitorsRequestSort) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *QueryMonitorsRequestSort) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *QueryMonitorsRequestSort) SetDirection(v string)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


