# ProcessorAnyOf3ParamsItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | Query that determines whether logs enter this branch | 
**Processors** | [**[]Processor**](Processor.md) | Processors executed when the branch query matches | 

## Methods

### NewProcessorAnyOf3ParamsItemsInner

`func NewProcessorAnyOf3ParamsItemsInner(query string, processors []Processor, ) *ProcessorAnyOf3ParamsItemsInner`

NewProcessorAnyOf3ParamsItemsInner instantiates a new ProcessorAnyOf3ParamsItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorAnyOf3ParamsItemsInnerWithDefaults

`func NewProcessorAnyOf3ParamsItemsInnerWithDefaults() *ProcessorAnyOf3ParamsItemsInner`

NewProcessorAnyOf3ParamsItemsInnerWithDefaults instantiates a new ProcessorAnyOf3ParamsItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *ProcessorAnyOf3ParamsItemsInner) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ProcessorAnyOf3ParamsItemsInner) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ProcessorAnyOf3ParamsItemsInner) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetProcessors

`func (o *ProcessorAnyOf3ParamsItemsInner) GetProcessors() []Processor`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *ProcessorAnyOf3ParamsItemsInner) GetProcessorsOk() (*[]Processor, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *ProcessorAnyOf3ParamsItemsInner) SetProcessors(v []Processor)`

SetProcessors sets Processors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


