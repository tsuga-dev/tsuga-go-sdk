# InputRumProcessorAnyOfParamsItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | Query that determines whether RUM events enter this branch | 
**Processors** | [**[]InputRumProcessor**](InputRumProcessor.md) | Processors executed when the branch query matches | 

## Methods

### NewInputRumProcessorAnyOfParamsItemsInner

`func NewInputRumProcessorAnyOfParamsItemsInner(query string, processors []InputRumProcessor, ) *InputRumProcessorAnyOfParamsItemsInner`

NewInputRumProcessorAnyOfParamsItemsInner instantiates a new InputRumProcessorAnyOfParamsItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputRumProcessorAnyOfParamsItemsInnerWithDefaults

`func NewInputRumProcessorAnyOfParamsItemsInnerWithDefaults() *InputRumProcessorAnyOfParamsItemsInner`

NewInputRumProcessorAnyOfParamsItemsInnerWithDefaults instantiates a new InputRumProcessorAnyOfParamsItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *InputRumProcessorAnyOfParamsItemsInner) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *InputRumProcessorAnyOfParamsItemsInner) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *InputRumProcessorAnyOfParamsItemsInner) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetProcessors

`func (o *InputRumProcessorAnyOfParamsItemsInner) GetProcessors() []InputRumProcessor`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *InputRumProcessorAnyOfParamsItemsInner) GetProcessorsOk() (*[]InputRumProcessor, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *InputRumProcessorAnyOfParamsItemsInner) SetProcessors(v []InputRumProcessor)`

SetProcessors sets Processors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


