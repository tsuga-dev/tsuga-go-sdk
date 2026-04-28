# ResponseMetadataPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** | The total number of items matching the query | 
**PageCount** | **int32** | The total number of pages available | 

## Methods

### NewResponseMetadataPagination

`func NewResponseMetadataPagination(totalCount int32, pageCount int32, ) *ResponseMetadataPagination`

NewResponseMetadataPagination instantiates a new ResponseMetadataPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseMetadataPaginationWithDefaults

`func NewResponseMetadataPaginationWithDefaults() *ResponseMetadataPagination`

NewResponseMetadataPaginationWithDefaults instantiates a new ResponseMetadataPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ResponseMetadataPagination) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ResponseMetadataPagination) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ResponseMetadataPagination) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetPageCount

`func (o *ResponseMetadataPagination) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *ResponseMetadataPagination) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *ResponseMetadataPagination) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


