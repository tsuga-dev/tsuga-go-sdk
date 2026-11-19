# InputSpansProcessorAnyOfParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]InputSpansProcessorAnyOfParamsItemsInner**](InputSpansProcessorAnyOfParamsItemsInner.md) | Branches evaluated in order; only the first matching branch runs. A final branch with an empty query matches all remaining events. If no branch matches, the event is unchanged. | 

## Methods

### NewInputSpansProcessorAnyOfParams

`func NewInputSpansProcessorAnyOfParams(items []InputSpansProcessorAnyOfParamsItemsInner, ) *InputSpansProcessorAnyOfParams`

NewInputSpansProcessorAnyOfParams instantiates a new InputSpansProcessorAnyOfParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputSpansProcessorAnyOfParamsWithDefaults

`func NewInputSpansProcessorAnyOfParamsWithDefaults() *InputSpansProcessorAnyOfParams`

NewInputSpansProcessorAnyOfParamsWithDefaults instantiates a new InputSpansProcessorAnyOfParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InputSpansProcessorAnyOfParams) GetItems() []InputSpansProcessorAnyOfParamsItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InputSpansProcessorAnyOfParams) GetItemsOk() (*[]InputSpansProcessorAnyOfParamsItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InputSpansProcessorAnyOfParams) SetItems(v []InputSpansProcessorAnyOfParamsItemsInner)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


