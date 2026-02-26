# SpanEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**SearchSpansResponseDataSpansInnerContext**](SearchSpansResponseDataSpansInnerContext.md) |  | [optional] 
**ParentSpanId** | Pointer to **string** |  | [optional] 
**ResourceAttributes** | **map[string]interface{}** |  | 
**SpanAttributes** | **map[string]interface{}** |  | 
**Duration** | **float32** | In milliseconds. | 
**StartTime** | **string** | In nanoseconds. Transmitted as a string to avoid losing precision, as BigInt is not JSON-compatible. | 
**EndTime** | **string** | In nanoseconds. Transmitted as a string to avoid losing precision, as BigInt is not JSON-compatible. | 
**Error** | Pointer to [**SearchSpansResponseDataSpansInnerError**](SearchSpansResponseDataSpansInnerError.md) |  | [optional] 
**Events** | Pointer to [**[]SearchSpansResponseDataSpansInnerEventsInner**](SearchSpansResponseDataSpansInnerEventsInner.md) |  | [optional] 
**SpanId** | **string** |  | 
**Span** | [**SearchSpansResponseDataSpansInnerSpan**](SearchSpansResponseDataSpansInnerSpan.md) |  | 
**StatusCode** | Pointer to **string** |  | [optional] 
**TraceId** | **string** |  | 

## Methods

### NewSpanEvent

`func NewSpanEvent(resourceAttributes map[string]interface{}, spanAttributes map[string]interface{}, duration float32, startTime string, endTime string, spanId string, span SearchSpansResponseDataSpansInnerSpan, traceId string, ) *SpanEvent`

NewSpanEvent instantiates a new SpanEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanEventWithDefaults

`func NewSpanEventWithDefaults() *SpanEvent`

NewSpanEventWithDefaults instantiates a new SpanEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *SpanEvent) GetContext() SearchSpansResponseDataSpansInnerContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SpanEvent) GetContextOk() (*SearchSpansResponseDataSpansInnerContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SpanEvent) SetContext(v SearchSpansResponseDataSpansInnerContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *SpanEvent) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetParentSpanId

`func (o *SpanEvent) GetParentSpanId() string`

GetParentSpanId returns the ParentSpanId field if non-nil, zero value otherwise.

### GetParentSpanIdOk

`func (o *SpanEvent) GetParentSpanIdOk() (*string, bool)`

GetParentSpanIdOk returns a tuple with the ParentSpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpanId

`func (o *SpanEvent) SetParentSpanId(v string)`

SetParentSpanId sets ParentSpanId field to given value.

### HasParentSpanId

`func (o *SpanEvent) HasParentSpanId() bool`

HasParentSpanId returns a boolean if a field has been set.

### GetResourceAttributes

`func (o *SpanEvent) GetResourceAttributes() map[string]interface{}`

GetResourceAttributes returns the ResourceAttributes field if non-nil, zero value otherwise.

### GetResourceAttributesOk

`func (o *SpanEvent) GetResourceAttributesOk() (*map[string]interface{}, bool)`

GetResourceAttributesOk returns a tuple with the ResourceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttributes

`func (o *SpanEvent) SetResourceAttributes(v map[string]interface{})`

SetResourceAttributes sets ResourceAttributes field to given value.


### GetSpanAttributes

`func (o *SpanEvent) GetSpanAttributes() map[string]interface{}`

GetSpanAttributes returns the SpanAttributes field if non-nil, zero value otherwise.

### GetSpanAttributesOk

`func (o *SpanEvent) GetSpanAttributesOk() (*map[string]interface{}, bool)`

GetSpanAttributesOk returns a tuple with the SpanAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanAttributes

`func (o *SpanEvent) SetSpanAttributes(v map[string]interface{})`

SetSpanAttributes sets SpanAttributes field to given value.


### GetDuration

`func (o *SpanEvent) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SpanEvent) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SpanEvent) SetDuration(v float32)`

SetDuration sets Duration field to given value.


### GetStartTime

`func (o *SpanEvent) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SpanEvent) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SpanEvent) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *SpanEvent) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SpanEvent) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SpanEvent) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### GetError

`func (o *SpanEvent) GetError() SearchSpansResponseDataSpansInnerError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SpanEvent) GetErrorOk() (*SearchSpansResponseDataSpansInnerError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SpanEvent) SetError(v SearchSpansResponseDataSpansInnerError)`

SetError sets Error field to given value.

### HasError

`func (o *SpanEvent) HasError() bool`

HasError returns a boolean if a field has been set.

### GetEvents

`func (o *SpanEvent) GetEvents() []SearchSpansResponseDataSpansInnerEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SpanEvent) GetEventsOk() (*[]SearchSpansResponseDataSpansInnerEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SpanEvent) SetEvents(v []SearchSpansResponseDataSpansInnerEventsInner)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *SpanEvent) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetSpanId

`func (o *SpanEvent) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *SpanEvent) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *SpanEvent) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### GetSpan

`func (o *SpanEvent) GetSpan() SearchSpansResponseDataSpansInnerSpan`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *SpanEvent) GetSpanOk() (*SearchSpansResponseDataSpansInnerSpan, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *SpanEvent) SetSpan(v SearchSpansResponseDataSpansInnerSpan)`

SetSpan sets Span field to given value.


### GetStatusCode

`func (o *SpanEvent) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *SpanEvent) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *SpanEvent) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *SpanEvent) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetTraceId

`func (o *SpanEvent) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *SpanEvent) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *SpanEvent) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


