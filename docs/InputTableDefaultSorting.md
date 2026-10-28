# InputTableDefaultSorting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the column to sort by: &#x60;label&#x60; for the grouping column, or &#x60;col-&lt;index&gt;&#x60; with the zero-based position of the entry in &#x60;columns&#x60;. | 
**Desc** | **bool** | Whether the sort is descending | 

## Methods

### NewInputTableDefaultSorting

`func NewInputTableDefaultSorting(id string, desc bool, ) *InputTableDefaultSorting`

NewInputTableDefaultSorting instantiates a new InputTableDefaultSorting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputTableDefaultSortingWithDefaults

`func NewInputTableDefaultSortingWithDefaults() *InputTableDefaultSorting`

NewInputTableDefaultSortingWithDefaults instantiates a new InputTableDefaultSorting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InputTableDefaultSorting) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputTableDefaultSorting) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputTableDefaultSorting) SetId(v string)`

SetId sets Id field to given value.


### GetDesc

`func (o *InputTableDefaultSorting) GetDesc() bool`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *InputTableDefaultSorting) GetDescOk() (*bool, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *InputTableDefaultSorting) SetDesc(v bool)`

SetDesc sets Desc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


