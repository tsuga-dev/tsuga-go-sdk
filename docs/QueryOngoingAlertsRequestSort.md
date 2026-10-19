# QueryOngoingAlertsRequestSort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**By** | **string** | Field to sort the returned alerts by. For &#x60;priority&#x60;, &#x60;desc&#x60; returns highest-priority alerts first because priority 1 is highest. For &#x60;clusterId&#x60;, alerts raised on no cluster sort last in both directions. | 
**Direction** | **string** | Sort direction: ascending or descending. | 

## Methods

### NewQueryOngoingAlertsRequestSort

`func NewQueryOngoingAlertsRequestSort(by string, direction string, ) *QueryOngoingAlertsRequestSort`

NewQueryOngoingAlertsRequestSort instantiates a new QueryOngoingAlertsRequestSort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryOngoingAlertsRequestSortWithDefaults

`func NewQueryOngoingAlertsRequestSortWithDefaults() *QueryOngoingAlertsRequestSort`

NewQueryOngoingAlertsRequestSortWithDefaults instantiates a new QueryOngoingAlertsRequestSort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBy

`func (o *QueryOngoingAlertsRequestSort) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *QueryOngoingAlertsRequestSort) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *QueryOngoingAlertsRequestSort) SetBy(v string)`

SetBy sets By field to given value.


### GetDirection

`func (o *QueryOngoingAlertsRequestSort) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *QueryOngoingAlertsRequestSort) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *QueryOngoingAlertsRequestSort) SetDirection(v string)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


