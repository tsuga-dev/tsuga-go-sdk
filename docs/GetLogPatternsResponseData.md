# GetLogPatternsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Patterns** | [**[]GetLogPatternsResponseDataPatternsInner**](GetLogPatternsResponseDataPatternsInner.md) |  | 
**SampleSize** | **float32** |  | 

## Methods

### NewGetLogPatternsResponseData

`func NewGetLogPatternsResponseData(patterns []GetLogPatternsResponseDataPatternsInner, sampleSize float32, ) *GetLogPatternsResponseData`

NewGetLogPatternsResponseData instantiates a new GetLogPatternsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLogPatternsResponseDataWithDefaults

`func NewGetLogPatternsResponseDataWithDefaults() *GetLogPatternsResponseData`

NewGetLogPatternsResponseDataWithDefaults instantiates a new GetLogPatternsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatterns

`func (o *GetLogPatternsResponseData) GetPatterns() []GetLogPatternsResponseDataPatternsInner`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *GetLogPatternsResponseData) GetPatternsOk() (*[]GetLogPatternsResponseDataPatternsInner, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *GetLogPatternsResponseData) SetPatterns(v []GetLogPatternsResponseDataPatternsInner)`

SetPatterns sets Patterns field to given value.


### GetSampleSize

`func (o *GetLogPatternsResponseData) GetSampleSize() float32`

GetSampleSize returns the SampleSize field if non-nil, zero value otherwise.

### GetSampleSizeOk

`func (o *GetLogPatternsResponseData) GetSampleSizeOk() (*float32, bool)`

GetSampleSizeOk returns a tuple with the SampleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleSize

`func (o *GetLogPatternsResponseData) SetSampleSize(v float32)`

SetSampleSize sets SampleSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


