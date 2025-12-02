# Route1ProcessorsInnerAnyOf1ParamsOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeName** | **string** | Attribute whose value will be parsed with Grok rules | 
**Rules** | **[]string** | Ordered Grok rules evaluated until one matches | 
**Samples** | Pointer to **[]string** | Example log lines for validation | [optional] 
**Subtype** | **string** |  | 

## Methods

### NewRoute1ProcessorsInnerAnyOf1ParamsOneOf

`func NewRoute1ProcessorsInnerAnyOf1ParamsOneOf(attributeName string, rules []string, subtype string, ) *Route1ProcessorsInnerAnyOf1ParamsOneOf`

NewRoute1ProcessorsInnerAnyOf1ParamsOneOf instantiates a new Route1ProcessorsInnerAnyOf1ParamsOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoute1ProcessorsInnerAnyOf1ParamsOneOfWithDefaults

`func NewRoute1ProcessorsInnerAnyOf1ParamsOneOfWithDefaults() *Route1ProcessorsInnerAnyOf1ParamsOneOf`

NewRoute1ProcessorsInnerAnyOf1ParamsOneOfWithDefaults instantiates a new Route1ProcessorsInnerAnyOf1ParamsOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeName

`func (o *Route1ProcessorsInnerAnyOf1ParamsOneOf) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *Route1ProcessorsInnerAnyOf1ParamsOneOf) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *Route1ProcessorsInnerAnyOf1ParamsOneOf) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.


### GetRules

`func (o *Route1ProcessorsInnerAnyOf1ParamsOneOf) GetRules() []string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Route1ProcessorsInnerAnyOf1ParamsOneOf) GetRulesOk() (*[]string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Route1ProcessorsInnerAnyOf1ParamsOneOf) SetRules(v []string)`

SetRules sets Rules field to given value.


### GetSamples

`func (o *Route1ProcessorsInnerAnyOf1ParamsOneOf) GetSamples() []string`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *Route1ProcessorsInnerAnyOf1ParamsOneOf) GetSamplesOk() (*[]string, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *Route1ProcessorsInnerAnyOf1ParamsOneOf) SetSamples(v []string)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *Route1ProcessorsInnerAnyOf1ParamsOneOf) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetSubtype

`func (o *Route1ProcessorsInnerAnyOf1ParamsOneOf) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *Route1ProcessorsInnerAnyOf1ParamsOneOf) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *Route1ProcessorsInnerAnyOf1ParamsOneOf) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


