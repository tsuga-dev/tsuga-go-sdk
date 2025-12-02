# ConditionalFormatting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** | Comparator used to evaluate the metric value | 
**Value** | **float32** | Threshold value used when evaluating the operator | 
**Color** | **string** | Color applied when the condition is met | 

## Methods

### NewConditionalFormatting

`func NewConditionalFormatting(operator string, value float32, color string, ) *ConditionalFormatting`

NewConditionalFormatting instantiates a new ConditionalFormatting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionalFormattingWithDefaults

`func NewConditionalFormattingWithDefaults() *ConditionalFormatting`

NewConditionalFormattingWithDefaults instantiates a new ConditionalFormatting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ConditionalFormatting) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ConditionalFormatting) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ConditionalFormatting) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *ConditionalFormatting) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConditionalFormatting) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConditionalFormatting) SetValue(v float32)`

SetValue sets Value field to given value.


### GetColor

`func (o *ConditionalFormatting) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ConditionalFormatting) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ConditionalFormatting) SetColor(v string)`

SetColor sets Color field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


