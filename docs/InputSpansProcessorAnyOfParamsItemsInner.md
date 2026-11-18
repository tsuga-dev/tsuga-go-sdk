# InputSpansProcessorAnyOfParamsItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | Query that determines whether spans enter this branch | 
**Processors** | [**[]InputSpansProcessor**](InputSpansProcessor.md) | Processors executed when the branch query matches | 

## Methods

### NewInputSpansProcessorAnyOfParamsItemsInner

`func NewInputSpansProcessorAnyOfParamsItemsInner(query string, processors []InputSpansProcessor, ) *InputSpansProcessorAnyOfParamsItemsInner`

NewInputSpansProcessorAnyOfParamsItemsInner instantiates a new InputSpansProcessorAnyOfParamsItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputSpansProcessorAnyOfParamsItemsInnerWithDefaults

`func NewInputSpansProcessorAnyOfParamsItemsInnerWithDefaults() *InputSpansProcessorAnyOfParamsItemsInner`

NewInputSpansProcessorAnyOfParamsItemsInnerWithDefaults instantiates a new InputSpansProcessorAnyOfParamsItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *InputSpansProcessorAnyOfParamsItemsInner) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *InputSpansProcessorAnyOfParamsItemsInner) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *InputSpansProcessorAnyOfParamsItemsInner) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetProcessors

`func (o *InputSpansProcessorAnyOfParamsItemsInner) GetProcessors() []InputSpansProcessor`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *InputSpansProcessorAnyOfParamsItemsInner) GetProcessorsOk() (*[]InputSpansProcessor, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *InputSpansProcessorAnyOfParamsItemsInner) SetProcessors(v []InputSpansProcessor)`

SetProcessors sets Processors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


