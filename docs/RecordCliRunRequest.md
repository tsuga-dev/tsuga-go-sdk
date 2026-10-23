# RecordCliRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CliVersion** | **string** | Version of the CLI that ran. | 
**Command** | **string** | Command path without its arguments, e.g. &#x60;monitors list&#x60;. | 
**DurationMs** | **int32** | Run duration. | 
**ExitCode** | **int32** | Process exit code. | 
**ErrorType** | Pointer to **string** | Kind of failure, e.g. &#x60;TransportError&#x60;. | [optional] 
**OsType** | **string** | e.g. &#x60;darwin&#x60;. | 
**NodeVersion** | **string** | e.g. &#x60;24.1.0&#x60;. | 
**TraceId** | **string** | W3C trace ID the CLI sent with its API calls during the run. | 
**SpanId** | **string** | W3C span ID the CLI named as the parent of its API calls. | 

## Methods

### NewRecordCliRunRequest

`func NewRecordCliRunRequest(cliVersion string, command string, durationMs int32, exitCode int32, osType string, nodeVersion string, traceId string, spanId string, ) *RecordCliRunRequest`

NewRecordCliRunRequest instantiates a new RecordCliRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordCliRunRequestWithDefaults

`func NewRecordCliRunRequestWithDefaults() *RecordCliRunRequest`

NewRecordCliRunRequestWithDefaults instantiates a new RecordCliRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCliVersion

`func (o *RecordCliRunRequest) GetCliVersion() string`

GetCliVersion returns the CliVersion field if non-nil, zero value otherwise.

### GetCliVersionOk

`func (o *RecordCliRunRequest) GetCliVersionOk() (*string, bool)`

GetCliVersionOk returns a tuple with the CliVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliVersion

`func (o *RecordCliRunRequest) SetCliVersion(v string)`

SetCliVersion sets CliVersion field to given value.


### GetCommand

`func (o *RecordCliRunRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *RecordCliRunRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *RecordCliRunRequest) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetDurationMs

`func (o *RecordCliRunRequest) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *RecordCliRunRequest) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *RecordCliRunRequest) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.


### GetExitCode

`func (o *RecordCliRunRequest) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *RecordCliRunRequest) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *RecordCliRunRequest) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.


### GetErrorType

`func (o *RecordCliRunRequest) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *RecordCliRunRequest) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *RecordCliRunRequest) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *RecordCliRunRequest) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetOsType

`func (o *RecordCliRunRequest) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *RecordCliRunRequest) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *RecordCliRunRequest) SetOsType(v string)`

SetOsType sets OsType field to given value.


### GetNodeVersion

`func (o *RecordCliRunRequest) GetNodeVersion() string`

GetNodeVersion returns the NodeVersion field if non-nil, zero value otherwise.

### GetNodeVersionOk

`func (o *RecordCliRunRequest) GetNodeVersionOk() (*string, bool)`

GetNodeVersionOk returns a tuple with the NodeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeVersion

`func (o *RecordCliRunRequest) SetNodeVersion(v string)`

SetNodeVersion sets NodeVersion field to given value.


### GetTraceId

`func (o *RecordCliRunRequest) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *RecordCliRunRequest) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *RecordCliRunRequest) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetSpanId

`func (o *RecordCliRunRequest) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *RecordCliRunRequest) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *RecordCliRunRequest) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


