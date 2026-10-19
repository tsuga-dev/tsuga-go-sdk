# QueryPromqlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | PromQL expression to evaluate. | 
**Step** | **string** | Evaluation step interval (e.g. &#x60;1m&#x60;, &#x60;5m&#x60;, &#x60;1h&#x60;). Must be a valid Prometheus duration string. | 
**TimeRange** | [**InputTimeRange**](InputTimeRange.md) |  | 

## Methods

### NewQueryPromqlRequest

`func NewQueryPromqlRequest(query string, step string, timeRange InputTimeRange, ) *QueryPromqlRequest`

NewQueryPromqlRequest instantiates a new QueryPromqlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryPromqlRequestWithDefaults

`func NewQueryPromqlRequestWithDefaults() *QueryPromqlRequest`

NewQueryPromqlRequestWithDefaults instantiates a new QueryPromqlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *QueryPromqlRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QueryPromqlRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QueryPromqlRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetStep

`func (o *QueryPromqlRequest) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *QueryPromqlRequest) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *QueryPromqlRequest) SetStep(v string)`

SetStep sets Step field to given value.


### GetTimeRange

`func (o *QueryPromqlRequest) GetTimeRange() InputTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *QueryPromqlRequest) GetTimeRangeOk() (*InputTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *QueryPromqlRequest) SetTimeRange(v InputTimeRange)`

SetTimeRange sets TimeRange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


