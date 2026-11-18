# OutputSpansProcessorAnyOf1ParamsItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | Query that determines whether spans enter this branch | 
**Processors** | [**[]OutputSpansProcessor**](OutputSpansProcessor.md) | Processors executed when the branch query matches | 

## Methods

### NewOutputSpansProcessorAnyOf1ParamsItemsInner

`func NewOutputSpansProcessorAnyOf1ParamsItemsInner(query string, processors []OutputSpansProcessor, ) *OutputSpansProcessorAnyOf1ParamsItemsInner`

NewOutputSpansProcessorAnyOf1ParamsItemsInner instantiates a new OutputSpansProcessorAnyOf1ParamsItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputSpansProcessorAnyOf1ParamsItemsInnerWithDefaults

`func NewOutputSpansProcessorAnyOf1ParamsItemsInnerWithDefaults() *OutputSpansProcessorAnyOf1ParamsItemsInner`

NewOutputSpansProcessorAnyOf1ParamsItemsInnerWithDefaults instantiates a new OutputSpansProcessorAnyOf1ParamsItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *OutputSpansProcessorAnyOf1ParamsItemsInner) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *OutputSpansProcessorAnyOf1ParamsItemsInner) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *OutputSpansProcessorAnyOf1ParamsItemsInner) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetProcessors

`func (o *OutputSpansProcessorAnyOf1ParamsItemsInner) GetProcessors() []OutputSpansProcessor`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *OutputSpansProcessorAnyOf1ParamsItemsInner) GetProcessorsOk() (*[]OutputSpansProcessor, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *OutputSpansProcessorAnyOf1ParamsItemsInner) SetProcessors(v []OutputSpansProcessor)`

SetProcessors sets Processors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


