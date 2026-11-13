# QueryPromqlV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | PromQL expression to evaluate. | 
**Step** | **string** | Evaluation step interval: a positive whole number followed by one unit, &#x60;s&#x60;, &#x60;m&#x60;, &#x60;h&#x60;, &#x60;d&#x60;, or &#x60;w&#x60;, such as &#x60;60s&#x60;, &#x60;1m&#x60;, or &#x60;1h&#x60;. A number without a unit, such as &#x60;60&#x60;, is not accepted. | 
**TimeRange** | [**InputV2TimeRange**](InputV2TimeRange.md) |  | 

## Methods

### NewQueryPromqlV2Request

`func NewQueryPromqlV2Request(query string, step string, timeRange InputV2TimeRange, ) *QueryPromqlV2Request`

NewQueryPromqlV2Request instantiates a new QueryPromqlV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryPromqlV2RequestWithDefaults

`func NewQueryPromqlV2RequestWithDefaults() *QueryPromqlV2Request`

NewQueryPromqlV2RequestWithDefaults instantiates a new QueryPromqlV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *QueryPromqlV2Request) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QueryPromqlV2Request) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QueryPromqlV2Request) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetStep

`func (o *QueryPromqlV2Request) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *QueryPromqlV2Request) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *QueryPromqlV2Request) SetStep(v string)`

SetStep sets Step field to given value.


### GetTimeRange

`func (o *QueryPromqlV2Request) GetTimeRange() InputV2TimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *QueryPromqlV2Request) GetTimeRangeOk() (*InputV2TimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *QueryPromqlV2Request) SetTimeRange(v InputV2TimeRange)`

SetTimeRange sets TimeRange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


