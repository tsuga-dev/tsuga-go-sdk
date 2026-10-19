# GetInterestingFieldsRequestScopeInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Telemetry attribute key to scope the ranking by: a service, environment, team, route, or other positive equality term in the population you want fields for. Use dot-path keys such as &#x60;context.service.name&#x60;. | 
**Value** | Pointer to **string** | Attribute value the scoped telemetry should match, when the scope is a specific service, environment, team, route, or other known value. When omitted, artifacts that reference the key match this entry at full weight regardless of value. | [optional] 
**ValueMustMatch** | Pointer to **bool** | When true, this entry is a hard filter rather than a weight: any widget or monitor that does not match it exactly (key &#x3D; value) contributes no fields at all, instead of contributing at a reduced weight. Other widgets of the same dashboard still contribute if they match. Requires &#x60;value&#x60;. | [optional] 

## Methods

### NewGetInterestingFieldsRequestScopeInner

`func NewGetInterestingFieldsRequestScopeInner(key string, ) *GetInterestingFieldsRequestScopeInner`

NewGetInterestingFieldsRequestScopeInner instantiates a new GetInterestingFieldsRequestScopeInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInterestingFieldsRequestScopeInnerWithDefaults

`func NewGetInterestingFieldsRequestScopeInnerWithDefaults() *GetInterestingFieldsRequestScopeInner`

NewGetInterestingFieldsRequestScopeInnerWithDefaults instantiates a new GetInterestingFieldsRequestScopeInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *GetInterestingFieldsRequestScopeInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetInterestingFieldsRequestScopeInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetInterestingFieldsRequestScopeInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *GetInterestingFieldsRequestScopeInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetInterestingFieldsRequestScopeInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetInterestingFieldsRequestScopeInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetInterestingFieldsRequestScopeInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueMustMatch

`func (o *GetInterestingFieldsRequestScopeInner) GetValueMustMatch() bool`

GetValueMustMatch returns the ValueMustMatch field if non-nil, zero value otherwise.

### GetValueMustMatchOk

`func (o *GetInterestingFieldsRequestScopeInner) GetValueMustMatchOk() (*bool, bool)`

GetValueMustMatchOk returns a tuple with the ValueMustMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMustMatch

`func (o *GetInterestingFieldsRequestScopeInner) SetValueMustMatch(v bool)`

SetValueMustMatch sets ValueMustMatch field to given value.

### HasValueMustMatch

`func (o *GetInterestingFieldsRequestScopeInner) HasValueMustMatch() bool`

HasValueMustMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


