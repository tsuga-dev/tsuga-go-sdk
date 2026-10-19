# OutputProcessorAnyOf3ParamsItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | Query that determines whether logs enter this branch | 
**Processors** | [**[]OutputProcessor**](OutputProcessor.md) | Processors executed when the branch query matches | 

## Methods

### NewOutputProcessorAnyOf3ParamsItemsInner

`func NewOutputProcessorAnyOf3ParamsItemsInner(query string, processors []OutputProcessor, ) *OutputProcessorAnyOf3ParamsItemsInner`

NewOutputProcessorAnyOf3ParamsItemsInner instantiates a new OutputProcessorAnyOf3ParamsItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputProcessorAnyOf3ParamsItemsInnerWithDefaults

`func NewOutputProcessorAnyOf3ParamsItemsInnerWithDefaults() *OutputProcessorAnyOf3ParamsItemsInner`

NewOutputProcessorAnyOf3ParamsItemsInnerWithDefaults instantiates a new OutputProcessorAnyOf3ParamsItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *OutputProcessorAnyOf3ParamsItemsInner) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *OutputProcessorAnyOf3ParamsItemsInner) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *OutputProcessorAnyOf3ParamsItemsInner) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetProcessors

`func (o *OutputProcessorAnyOf3ParamsItemsInner) GetProcessors() []OutputProcessor`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *OutputProcessorAnyOf3ParamsItemsInner) GetProcessorsOk() (*[]OutputProcessor, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *OutputProcessorAnyOf3ParamsItemsInner) SetProcessors(v []OutputProcessor)`

SetProcessors sets Processors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


