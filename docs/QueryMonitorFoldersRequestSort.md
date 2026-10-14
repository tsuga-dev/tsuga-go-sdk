# QueryMonitorFoldersRequestSort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**By** | **string** | Field to sort the returned monitor folders by: &#x60;owner&#x60; sorts by owning team name, &#x60;tags&#x60; sorts by tag key/value pairs, &#x60;monitorCount&#x60; sorts by the number of readable monitors directly in the folder, &#x60;updatedAt&#x60; sorts by last-updated timestamp, and &#x60;name&#x60; sorts alphabetically by folder name. | 
**Direction** | **string** | Sort direction: ascending or descending. | 

## Methods

### NewQueryMonitorFoldersRequestSort

`func NewQueryMonitorFoldersRequestSort(by string, direction string, ) *QueryMonitorFoldersRequestSort`

NewQueryMonitorFoldersRequestSort instantiates a new QueryMonitorFoldersRequestSort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMonitorFoldersRequestSortWithDefaults

`func NewQueryMonitorFoldersRequestSortWithDefaults() *QueryMonitorFoldersRequestSort`

NewQueryMonitorFoldersRequestSortWithDefaults instantiates a new QueryMonitorFoldersRequestSort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBy

`func (o *QueryMonitorFoldersRequestSort) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *QueryMonitorFoldersRequestSort) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *QueryMonitorFoldersRequestSort) SetBy(v string)`

SetBy sets By field to given value.


### GetDirection

`func (o *QueryMonitorFoldersRequestSort) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *QueryMonitorFoldersRequestSort) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *QueryMonitorFoldersRequestSort) SetDirection(v string)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


