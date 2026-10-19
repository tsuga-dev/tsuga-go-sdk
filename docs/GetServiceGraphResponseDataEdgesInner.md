# GetServiceGraphResponseDataEdgesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**GetServiceGraphResponseDataEdgesInnerFrom**](GetServiceGraphResponseDataEdgesInnerFrom.md) |  | 
**To** | [**GetServiceGraphResponseDataEdgesInnerTo**](GetServiceGraphResponseDataEdgesInnerTo.md) |  | 
**Count** | **float32** | Estimated number of cross-service transitions on this directed edge in the graph query window. The graph is built from a uniform sample of traces, so this is projected onto the full population rather than counted directly. | 
**ErrorRate** | Pointer to **float32** | Ratio of errored client spans to total client spans observed on this edge. Omitted when no client spans were observed. For edges to inferred services, the originating span is treated as the client span regardless of its kind. Computed from the raw sampled counts, so trace sampling does not affect it. | [optional] 
**ErrorCount** | **float32** | Estimated number of client spans with error status on this edge, projected from the sampled traces onto the full population. | 
**ClientSpanCount** | **float32** | Estimated number of client spans on this edge, projected from the sampled traces onto the full population. | 

## Methods

### NewGetServiceGraphResponseDataEdgesInner

`func NewGetServiceGraphResponseDataEdgesInner(from GetServiceGraphResponseDataEdgesInnerFrom, to GetServiceGraphResponseDataEdgesInnerTo, count float32, errorCount float32, clientSpanCount float32, ) *GetServiceGraphResponseDataEdgesInner`

NewGetServiceGraphResponseDataEdgesInner instantiates a new GetServiceGraphResponseDataEdgesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServiceGraphResponseDataEdgesInnerWithDefaults

`func NewGetServiceGraphResponseDataEdgesInnerWithDefaults() *GetServiceGraphResponseDataEdgesInner`

NewGetServiceGraphResponseDataEdgesInnerWithDefaults instantiates a new GetServiceGraphResponseDataEdgesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *GetServiceGraphResponseDataEdgesInner) GetFrom() GetServiceGraphResponseDataEdgesInnerFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetServiceGraphResponseDataEdgesInner) GetFromOk() (*GetServiceGraphResponseDataEdgesInnerFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetServiceGraphResponseDataEdgesInner) SetFrom(v GetServiceGraphResponseDataEdgesInnerFrom)`

SetFrom sets From field to given value.


### GetTo

`func (o *GetServiceGraphResponseDataEdgesInner) GetTo() GetServiceGraphResponseDataEdgesInnerTo`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GetServiceGraphResponseDataEdgesInner) GetToOk() (*GetServiceGraphResponseDataEdgesInnerTo, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GetServiceGraphResponseDataEdgesInner) SetTo(v GetServiceGraphResponseDataEdgesInnerTo)`

SetTo sets To field to given value.


### GetCount

`func (o *GetServiceGraphResponseDataEdgesInner) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetServiceGraphResponseDataEdgesInner) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetServiceGraphResponseDataEdgesInner) SetCount(v float32)`

SetCount sets Count field to given value.


### GetErrorRate

`func (o *GetServiceGraphResponseDataEdgesInner) GetErrorRate() float32`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *GetServiceGraphResponseDataEdgesInner) GetErrorRateOk() (*float32, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *GetServiceGraphResponseDataEdgesInner) SetErrorRate(v float32)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *GetServiceGraphResponseDataEdgesInner) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetErrorCount

`func (o *GetServiceGraphResponseDataEdgesInner) GetErrorCount() float32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *GetServiceGraphResponseDataEdgesInner) GetErrorCountOk() (*float32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *GetServiceGraphResponseDataEdgesInner) SetErrorCount(v float32)`

SetErrorCount sets ErrorCount field to given value.


### GetClientSpanCount

`func (o *GetServiceGraphResponseDataEdgesInner) GetClientSpanCount() float32`

GetClientSpanCount returns the ClientSpanCount field if non-nil, zero value otherwise.

### GetClientSpanCountOk

`func (o *GetServiceGraphResponseDataEdgesInner) GetClientSpanCountOk() (*float32, bool)`

GetClientSpanCountOk returns a tuple with the ClientSpanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSpanCount

`func (o *GetServiceGraphResponseDataEdgesInner) SetClientSpanCount(v float32)`

SetClientSpanCount sets ClientSpanCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


