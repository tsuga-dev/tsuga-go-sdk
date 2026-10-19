# InputProcessorAnyOf3ParamsItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | Query that determines whether logs enter this branch | 
**Processors** | [**[]InputProcessor**](InputProcessor.md) | Processors executed when the branch query matches | 

## Methods

### NewInputProcessorAnyOf3ParamsItemsInner

`func NewInputProcessorAnyOf3ParamsItemsInner(query string, processors []InputProcessor, ) *InputProcessorAnyOf3ParamsItemsInner`

NewInputProcessorAnyOf3ParamsItemsInner instantiates a new InputProcessorAnyOf3ParamsItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputProcessorAnyOf3ParamsItemsInnerWithDefaults

`func NewInputProcessorAnyOf3ParamsItemsInnerWithDefaults() *InputProcessorAnyOf3ParamsItemsInner`

NewInputProcessorAnyOf3ParamsItemsInnerWithDefaults instantiates a new InputProcessorAnyOf3ParamsItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *InputProcessorAnyOf3ParamsItemsInner) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *InputProcessorAnyOf3ParamsItemsInner) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *InputProcessorAnyOf3ParamsItemsInner) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetProcessors

`func (o *InputProcessorAnyOf3ParamsItemsInner) GetProcessors() []InputProcessor`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *InputProcessorAnyOf3ParamsItemsInner) GetProcessorsOk() (*[]InputProcessor, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *InputProcessorAnyOf3ParamsItemsInner) SetProcessors(v []InputProcessor)`

SetProcessors sets Processors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


