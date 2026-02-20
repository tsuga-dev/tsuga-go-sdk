# SearchSpans200ResponseDataSpansInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**SearchSpans200ResponseDataSpansInnerContext**](SearchSpans200ResponseDataSpansInnerContext.md) |  | [optional] 
**ParentSpanId** | Pointer to **string** |  | [optional] 
**ResourceAttributes** | **map[string]interface{}** |  | 
**SpanAttributes** | **map[string]interface{}** |  | 
**Duration** | **float32** | In milliseconds. | 
**StartTime** | **string** | In nanoseconds. Transmitted as a string to avoid losing precision, as BigInt is not JSON-compatible. | 
**EndTime** | **string** | In nanoseconds. Transmitted as a string to avoid losing precision, as BigInt is not JSON-compatible. | 
**Error** | Pointer to [**SearchSpans200ResponseDataSpansInnerError**](SearchSpans200ResponseDataSpansInnerError.md) |  | [optional] 
**Events** | Pointer to [**[]SearchSpans200ResponseDataSpansInnerEventsInner**](SearchSpans200ResponseDataSpansInnerEventsInner.md) |  | [optional] 
**SpanId** | **string** |  | 
**Span** | [**SearchSpans200ResponseDataSpansInnerSpan**](SearchSpans200ResponseDataSpansInnerSpan.md) |  | 
**StatusCode** | Pointer to **string** |  | [optional] 
**TraceId** | **string** |  | 

## Methods

### NewSearchSpans200ResponseDataSpansInner

`func NewSearchSpans200ResponseDataSpansInner(resourceAttributes map[string]interface{}, spanAttributes map[string]interface{}, duration float32, startTime string, endTime string, spanId string, span SearchSpans200ResponseDataSpansInnerSpan, traceId string, ) *SearchSpans200ResponseDataSpansInner`

NewSearchSpans200ResponseDataSpansInner instantiates a new SearchSpans200ResponseDataSpansInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSpans200ResponseDataSpansInnerWithDefaults

`func NewSearchSpans200ResponseDataSpansInnerWithDefaults() *SearchSpans200ResponseDataSpansInner`

NewSearchSpans200ResponseDataSpansInnerWithDefaults instantiates a new SearchSpans200ResponseDataSpansInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *SearchSpans200ResponseDataSpansInner) GetContext() SearchSpans200ResponseDataSpansInnerContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SearchSpans200ResponseDataSpansInner) GetContextOk() (*SearchSpans200ResponseDataSpansInnerContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SearchSpans200ResponseDataSpansInner) SetContext(v SearchSpans200ResponseDataSpansInnerContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *SearchSpans200ResponseDataSpansInner) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetParentSpanId

`func (o *SearchSpans200ResponseDataSpansInner) GetParentSpanId() string`

GetParentSpanId returns the ParentSpanId field if non-nil, zero value otherwise.

### GetParentSpanIdOk

`func (o *SearchSpans200ResponseDataSpansInner) GetParentSpanIdOk() (*string, bool)`

GetParentSpanIdOk returns a tuple with the ParentSpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpanId

`func (o *SearchSpans200ResponseDataSpansInner) SetParentSpanId(v string)`

SetParentSpanId sets ParentSpanId field to given value.

### HasParentSpanId

`func (o *SearchSpans200ResponseDataSpansInner) HasParentSpanId() bool`

HasParentSpanId returns a boolean if a field has been set.

### GetResourceAttributes

`func (o *SearchSpans200ResponseDataSpansInner) GetResourceAttributes() map[string]interface{}`

GetResourceAttributes returns the ResourceAttributes field if non-nil, zero value otherwise.

### GetResourceAttributesOk

`func (o *SearchSpans200ResponseDataSpansInner) GetResourceAttributesOk() (*map[string]interface{}, bool)`

GetResourceAttributesOk returns a tuple with the ResourceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttributes

`func (o *SearchSpans200ResponseDataSpansInner) SetResourceAttributes(v map[string]interface{})`

SetResourceAttributes sets ResourceAttributes field to given value.


### GetSpanAttributes

`func (o *SearchSpans200ResponseDataSpansInner) GetSpanAttributes() map[string]interface{}`

GetSpanAttributes returns the SpanAttributes field if non-nil, zero value otherwise.

### GetSpanAttributesOk

`func (o *SearchSpans200ResponseDataSpansInner) GetSpanAttributesOk() (*map[string]interface{}, bool)`

GetSpanAttributesOk returns a tuple with the SpanAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanAttributes

`func (o *SearchSpans200ResponseDataSpansInner) SetSpanAttributes(v map[string]interface{})`

SetSpanAttributes sets SpanAttributes field to given value.


### GetDuration

`func (o *SearchSpans200ResponseDataSpansInner) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SearchSpans200ResponseDataSpansInner) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SearchSpans200ResponseDataSpansInner) SetDuration(v float32)`

SetDuration sets Duration field to given value.


### GetStartTime

`func (o *SearchSpans200ResponseDataSpansInner) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SearchSpans200ResponseDataSpansInner) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SearchSpans200ResponseDataSpansInner) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *SearchSpans200ResponseDataSpansInner) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SearchSpans200ResponseDataSpansInner) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SearchSpans200ResponseDataSpansInner) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### GetError

`func (o *SearchSpans200ResponseDataSpansInner) GetError() SearchSpans200ResponseDataSpansInnerError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SearchSpans200ResponseDataSpansInner) GetErrorOk() (*SearchSpans200ResponseDataSpansInnerError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SearchSpans200ResponseDataSpansInner) SetError(v SearchSpans200ResponseDataSpansInnerError)`

SetError sets Error field to given value.

### HasError

`func (o *SearchSpans200ResponseDataSpansInner) HasError() bool`

HasError returns a boolean if a field has been set.

### GetEvents

`func (o *SearchSpans200ResponseDataSpansInner) GetEvents() []SearchSpans200ResponseDataSpansInnerEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SearchSpans200ResponseDataSpansInner) GetEventsOk() (*[]SearchSpans200ResponseDataSpansInnerEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SearchSpans200ResponseDataSpansInner) SetEvents(v []SearchSpans200ResponseDataSpansInnerEventsInner)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *SearchSpans200ResponseDataSpansInner) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetSpanId

`func (o *SearchSpans200ResponseDataSpansInner) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *SearchSpans200ResponseDataSpansInner) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *SearchSpans200ResponseDataSpansInner) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### GetSpan

`func (o *SearchSpans200ResponseDataSpansInner) GetSpan() SearchSpans200ResponseDataSpansInnerSpan`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *SearchSpans200ResponseDataSpansInner) GetSpanOk() (*SearchSpans200ResponseDataSpansInnerSpan, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *SearchSpans200ResponseDataSpansInner) SetSpan(v SearchSpans200ResponseDataSpansInnerSpan)`

SetSpan sets Span field to given value.


### GetStatusCode

`func (o *SearchSpans200ResponseDataSpansInner) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *SearchSpans200ResponseDataSpansInner) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *SearchSpans200ResponseDataSpansInner) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *SearchSpans200ResponseDataSpansInner) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetTraceId

`func (o *SearchSpans200ResponseDataSpansInner) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *SearchSpans200ResponseDataSpansInner) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *SearchSpans200ResponseDataSpansInner) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


