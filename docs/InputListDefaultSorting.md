# InputListDefaultSorting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Column attribute used for the default list sort. Set this to an attribute in &#x60;listColumns&#x60; or a built-in list column for the selected list source. | 
**Desc** | **bool** | Sort direction for the default list sort. Set &#x60;true&#x60; for descending order or &#x60;false&#x60; for ascending order. | 

## Methods

### NewInputListDefaultSorting

`func NewInputListDefaultSorting(id string, desc bool, ) *InputListDefaultSorting`

NewInputListDefaultSorting instantiates a new InputListDefaultSorting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputListDefaultSortingWithDefaults

`func NewInputListDefaultSortingWithDefaults() *InputListDefaultSorting`

NewInputListDefaultSortingWithDefaults instantiates a new InputListDefaultSorting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InputListDefaultSorting) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputListDefaultSorting) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputListDefaultSorting) SetId(v string)`

SetId sets Id field to given value.


### GetDesc

`func (o *InputListDefaultSorting) GetDesc() bool`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *InputListDefaultSorting) GetDescOk() (*bool, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *InputListDefaultSorting) SetDesc(v bool)`

SetDesc sets Desc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


