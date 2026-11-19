# OutputSpansProcessorAnyOfParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]OutputSpansProcessorAnyOfParamsItemsInner**](OutputSpansProcessorAnyOfParamsItemsInner.md) | Branches evaluated in order; only the first matching branch runs. A final branch with an empty query matches all remaining events. If no branch matches, the event is unchanged. | 

## Methods

### NewOutputSpansProcessorAnyOfParams

`func NewOutputSpansProcessorAnyOfParams(items []OutputSpansProcessorAnyOfParamsItemsInner, ) *OutputSpansProcessorAnyOfParams`

NewOutputSpansProcessorAnyOfParams instantiates a new OutputSpansProcessorAnyOfParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputSpansProcessorAnyOfParamsWithDefaults

`func NewOutputSpansProcessorAnyOfParamsWithDefaults() *OutputSpansProcessorAnyOfParams`

NewOutputSpansProcessorAnyOfParamsWithDefaults instantiates a new OutputSpansProcessorAnyOfParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *OutputSpansProcessorAnyOfParams) GetItems() []OutputSpansProcessorAnyOfParamsItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OutputSpansProcessorAnyOfParams) GetItemsOk() (*[]OutputSpansProcessorAnyOfParamsItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OutputSpansProcessorAnyOfParams) SetItems(v []OutputSpansProcessorAnyOfParamsItemsInner)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


