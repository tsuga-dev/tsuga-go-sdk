# CreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeName** | **string** | Attribute whose value will be parsed with Grok rules | 
**Rules** | **[]string** | Ordered Grok rules evaluated until one matches | 
**Samples** | Pointer to **[]string** | Example log lines for validation | [optional] 
**Subtype** | **string** |  | 

## Methods

### NewCreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf

`func NewCreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf(attributeName string, rules []string, subtype string, ) *CreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf`

NewCreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf instantiates a new CreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRouteRequestProcessorsInnerAnyOf1ParamsOneOfWithDefaults

`func NewCreateRouteRequestProcessorsInnerAnyOf1ParamsOneOfWithDefaults() *CreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf`

NewCreateRouteRequestProcessorsInnerAnyOf1ParamsOneOfWithDefaults instantiates a new CreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeName

`func (o *CreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *CreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *CreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.


### GetRules

`func (o *CreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf) GetRules() []string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf) GetRulesOk() (*[]string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf) SetRules(v []string)`

SetRules sets Rules field to given value.


### GetSamples

`func (o *CreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf) GetSamples() []string`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *CreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf) GetSamplesOk() (*[]string, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *CreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf) SetSamples(v []string)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *CreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetSubtype

`func (o *CreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *CreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *CreateRouteRequestProcessorsInnerAnyOf1ParamsOneOf) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


