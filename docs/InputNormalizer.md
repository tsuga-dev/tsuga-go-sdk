# InputNormalizer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Normalizes values as durations using standard time units | 
**Unit** | **string** | Text label describing the custom unit. Must not be empty. | 

## Methods

### NewInputNormalizer

`func NewInputNormalizer(type_ string, unit string, ) *InputNormalizer`

NewInputNormalizer instantiates a new InputNormalizer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputNormalizerWithDefaults

`func NewInputNormalizerWithDefaults() *InputNormalizer`

NewInputNormalizerWithDefaults instantiates a new InputNormalizer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputNormalizer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputNormalizer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputNormalizer) SetType(v string)`

SetType sets Type field to given value.


### GetUnit

`func (o *InputNormalizer) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *InputNormalizer) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *InputNormalizer) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


