# ProcessorParseUserAgent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAttribute** | **string** | Attribute containing the user agent string to parse | 
**Subtype** | **string** | Parses a user-agent string and writes browser, operating system, and device details. | 

## Methods

### NewProcessorParseUserAgent

`func NewProcessorParseUserAgent(sourceAttribute string, subtype string, ) *ProcessorParseUserAgent`

NewProcessorParseUserAgent instantiates a new ProcessorParseUserAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorParseUserAgentWithDefaults

`func NewProcessorParseUserAgentWithDefaults() *ProcessorParseUserAgent`

NewProcessorParseUserAgentWithDefaults instantiates a new ProcessorParseUserAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAttribute

`func (o *ProcessorParseUserAgent) GetSourceAttribute() string`

GetSourceAttribute returns the SourceAttribute field if non-nil, zero value otherwise.

### GetSourceAttributeOk

`func (o *ProcessorParseUserAgent) GetSourceAttributeOk() (*string, bool)`

GetSourceAttributeOk returns a tuple with the SourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttribute

`func (o *ProcessorParseUserAgent) SetSourceAttribute(v string)`

SetSourceAttribute sets SourceAttribute field to given value.


### GetSubtype

`func (o *ProcessorParseUserAgent) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *ProcessorParseUserAgent) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *ProcessorParseUserAgent) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


