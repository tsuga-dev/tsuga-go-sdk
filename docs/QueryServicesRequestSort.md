# QueryServicesRequestSort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**By** | **string** | Field to sort the returned services by | 
**Direction** | **string** | Sort direction: ascending or descending. | 

## Methods

### NewQueryServicesRequestSort

`func NewQueryServicesRequestSort(by string, direction string, ) *QueryServicesRequestSort`

NewQueryServicesRequestSort instantiates a new QueryServicesRequestSort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryServicesRequestSortWithDefaults

`func NewQueryServicesRequestSortWithDefaults() *QueryServicesRequestSort`

NewQueryServicesRequestSortWithDefaults instantiates a new QueryServicesRequestSort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBy

`func (o *QueryServicesRequestSort) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *QueryServicesRequestSort) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *QueryServicesRequestSort) SetBy(v string)`

SetBy sets By field to given value.


### GetDirection

`func (o *QueryServicesRequestSort) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *QueryServicesRequestSort) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *QueryServicesRequestSort) SetDirection(v string)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


