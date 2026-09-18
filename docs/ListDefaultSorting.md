# ListDefaultSorting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Column attribute used for the default list sort. It matches an attribute in &#x60;listColumns&#x60; or a built-in list column for the selected list source. | 
**Desc** | **bool** | Sort direction for the default list sort. &#x60;true&#x60; sorts descending; &#x60;false&#x60; sorts ascending. | 

## Methods

### NewListDefaultSorting

`func NewListDefaultSorting(id string, desc bool, ) *ListDefaultSorting`

NewListDefaultSorting instantiates a new ListDefaultSorting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDefaultSortingWithDefaults

`func NewListDefaultSortingWithDefaults() *ListDefaultSorting`

NewListDefaultSortingWithDefaults instantiates a new ListDefaultSorting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListDefaultSorting) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListDefaultSorting) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListDefaultSorting) SetId(v string)`

SetId sets Id field to given value.


### GetDesc

`func (o *ListDefaultSorting) GetDesc() bool`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *ListDefaultSorting) GetDescOk() (*bool, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *ListDefaultSorting) SetDesc(v bool)`

SetDesc sets Desc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


