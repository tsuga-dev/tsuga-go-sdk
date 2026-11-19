# OutputSpansProcessorAnyOfParamsItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | Query that determines whether spans enter this branch | 
**Processors** | [**[]OutputSpansProcessor**](OutputSpansProcessor.md) | Processors executed when the branch query matches | 

## Methods

### NewOutputSpansProcessorAnyOfParamsItemsInner

`func NewOutputSpansProcessorAnyOfParamsItemsInner(query string, processors []OutputSpansProcessor, ) *OutputSpansProcessorAnyOfParamsItemsInner`

NewOutputSpansProcessorAnyOfParamsItemsInner instantiates a new OutputSpansProcessorAnyOfParamsItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputSpansProcessorAnyOfParamsItemsInnerWithDefaults

`func NewOutputSpansProcessorAnyOfParamsItemsInnerWithDefaults() *OutputSpansProcessorAnyOfParamsItemsInner`

NewOutputSpansProcessorAnyOfParamsItemsInnerWithDefaults instantiates a new OutputSpansProcessorAnyOfParamsItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *OutputSpansProcessorAnyOfParamsItemsInner) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *OutputSpansProcessorAnyOfParamsItemsInner) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *OutputSpansProcessorAnyOfParamsItemsInner) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetProcessors

`func (o *OutputSpansProcessorAnyOfParamsItemsInner) GetProcessors() []OutputSpansProcessor`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *OutputSpansProcessorAnyOfParamsItemsInner) GetProcessorsOk() (*[]OutputSpansProcessor, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *OutputSpansProcessorAnyOfParamsItemsInner) SetProcessors(v []OutputSpansProcessor)`

SetProcessors sets Processors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


