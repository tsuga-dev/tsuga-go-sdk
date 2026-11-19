# ProcessorParseUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAttribute** | **string** | Attribute containing the URL to parse | 
**Subtype** | **string** | Parses a URL attribute and writes URL detail fields. | 

## Methods

### NewProcessorParseUrl

`func NewProcessorParseUrl(sourceAttribute string, subtype string, ) *ProcessorParseUrl`

NewProcessorParseUrl instantiates a new ProcessorParseUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorParseUrlWithDefaults

`func NewProcessorParseUrlWithDefaults() *ProcessorParseUrl`

NewProcessorParseUrlWithDefaults instantiates a new ProcessorParseUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAttribute

`func (o *ProcessorParseUrl) GetSourceAttribute() string`

GetSourceAttribute returns the SourceAttribute field if non-nil, zero value otherwise.

### GetSourceAttributeOk

`func (o *ProcessorParseUrl) GetSourceAttributeOk() (*string, bool)`

GetSourceAttributeOk returns a tuple with the SourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttribute

`func (o *ProcessorParseUrl) SetSourceAttribute(v string)`

SetSourceAttribute sets SourceAttribute field to given value.


### GetSubtype

`func (o *ProcessorParseUrl) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *ProcessorParseUrl) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *ProcessorParseUrl) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


