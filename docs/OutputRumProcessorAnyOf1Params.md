# OutputRumProcessorAnyOf1Params

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]OutputRumProcessorAnyOf1ParamsItemsInner**](OutputRumProcessorAnyOf1ParamsItemsInner.md) | Branches evaluated in order; only the first matching branch runs. A final branch with an empty query matches all remaining events. If no branch matches, the event is unchanged. | 

## Methods

### NewOutputRumProcessorAnyOf1Params

`func NewOutputRumProcessorAnyOf1Params(items []OutputRumProcessorAnyOf1ParamsItemsInner, ) *OutputRumProcessorAnyOf1Params`

NewOutputRumProcessorAnyOf1Params instantiates a new OutputRumProcessorAnyOf1Params object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputRumProcessorAnyOf1ParamsWithDefaults

`func NewOutputRumProcessorAnyOf1ParamsWithDefaults() *OutputRumProcessorAnyOf1Params`

NewOutputRumProcessorAnyOf1ParamsWithDefaults instantiates a new OutputRumProcessorAnyOf1Params object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *OutputRumProcessorAnyOf1Params) GetItems() []OutputRumProcessorAnyOf1ParamsItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OutputRumProcessorAnyOf1Params) GetItemsOk() (*[]OutputRumProcessorAnyOf1ParamsItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OutputRumProcessorAnyOf1Params) SetItems(v []OutputRumProcessorAnyOf1ParamsItemsInner)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


