# \TracesAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindContrastSets**](TracesAPI.md#FindContrastSets) | **Post** /v1/traces/contrast-sets | 
[**GetTraceLatencySummary**](TracesAPI.md#GetTraceLatencySummary) | **Get** /v1/traces/latency-summary | 
[**SearchSpans**](TracesAPI.md#SearchSpans) | **Get** /v1/traces/search | 
[**SummarizeTrace**](TracesAPI.md#SummarizeTrace) | **Get** /v1/traces/summarize | 



## FindContrastSets

> FindContrastSetsResponse FindContrastSets(ctx).FindContrastSetsRequest(findContrastSetsRequest).ClusterId(clusterId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tsuga-dev/tsuga-go-sdk"
)

func main() {
	findContrastSetsRequest := *openapiclient.NewFindContrastSetsRequest(*openapiclient.NewFindContrastSetsRequestTargetGroup("Filter_example", *openapiclient.NewInputTimeRange(int64(123), int64(123))), *openapiclient.NewFindContrastSetsRequestBaselineGroup("Filter_example", *openapiclient.NewInputTimeRange(int64(123), int64(123)))) // FindContrastSetsRequest | Contrast-set request. Provide the target group to explain and the baseline to compare it against, plus optional significance tuning. Send `clusterId` as a query parameter, not in the JSON body.
	clusterId := "clusterId_example" // string | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracesAPI.FindContrastSets(context.Background()).FindContrastSetsRequest(findContrastSetsRequest).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.FindContrastSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindContrastSets`: FindContrastSetsResponse
	fmt.Fprintf(os.Stdout, "Response from `TracesAPI.FindContrastSets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindContrastSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **findContrastSetsRequest** | [**FindContrastSetsRequest**](FindContrastSetsRequest.md) | Contrast-set request. Provide the target group to explain and the baseline to compare it against, plus optional significance tuning. Send &#x60;clusterId&#x60; as a query parameter, not in the JSON body. | 
 **clusterId** | **string** | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. | 

### Return type

[**FindContrastSetsResponse**](FindContrastSetsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraceLatencySummary

> GetTraceLatencySummaryResponse GetTraceLatencySummary(ctx).From(from).To(to).TraceId(traceId).ClusterId(clusterId).MinRangeDurationMs(minRangeDurationMs).IncludeRanges(includeRanges).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tsuga-dev/tsuga-go-sdk"
)

func main() {
	from := int64(1704067200) // int64 | Start of the query window as a Unix timestamp in seconds
	to := int64(1704067200) // int64 | End of the query window as a Unix timestamp in seconds
	traceId := "traceId_example" // string | Trace ID to summarize. Use the `traceId` returned by span search, the trace details page, or your tracing instrumentation. Required; the requested time range must include the trace.
	clusterId := "clusterId_example" // string | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. (optional)
	minRangeDurationMs := float32(8.14) // float32 | Minimum range duration in milliseconds. Shorter ranges are collapsed (absorbed into a neighbor or blended with adjacent shorts) to keep the timeline readable. When omitted, the default depends on the trace: max(1ms, 1% of the total trace duration), so e.g. a 10s trace collapses ranges shorter than 100ms. Pass 0 to disable collapsing. (optional)
	includeRanges := true // bool | When false, omits the per-range timeline (returns `ranges: []`) and returns only `services` and `serviceTotals`. Cuts payload size for callers that only need the headline. Defaults to true. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracesAPI.GetTraceLatencySummary(context.Background()).From(from).To(to).TraceId(traceId).ClusterId(clusterId).MinRangeDurationMs(minRangeDurationMs).IncludeRanges(includeRanges).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.GetTraceLatencySummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTraceLatencySummary`: GetTraceLatencySummaryResponse
	fmt.Fprintf(os.Stdout, "Response from `TracesAPI.GetTraceLatencySummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceLatencySummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | Start of the query window as a Unix timestamp in seconds | 
 **to** | **int64** | End of the query window as a Unix timestamp in seconds | 
 **traceId** | **string** | Trace ID to summarize. Use the &#x60;traceId&#x60; returned by span search, the trace details page, or your tracing instrumentation. Required; the requested time range must include the trace. | 
 **clusterId** | **string** | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. | 
 **minRangeDurationMs** | **float32** | Minimum range duration in milliseconds. Shorter ranges are collapsed (absorbed into a neighbor or blended with adjacent shorts) to keep the timeline readable. When omitted, the default depends on the trace: max(1ms, 1% of the total trace duration), so e.g. a 10s trace collapses ranges shorter than 100ms. Pass 0 to disable collapsing. | 
 **includeRanges** | **bool** | When false, omits the per-range timeline (returns &#x60;ranges: []&#x60;) and returns only &#x60;services&#x60; and &#x60;serviceTotals&#x60;. Cuts payload size for callers that only need the headline. Defaults to true. | [default to true]

### Return type

[**GetTraceLatencySummaryResponse**](GetTraceLatencySummaryResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSpans

> SearchSpansResponse SearchSpans(ctx).From(from).To(to).Query(query).MaxResults(maxResults).ClusterId(clusterId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tsuga-dev/tsuga-go-sdk"
)

func main() {
	from := int64(1704067200) // int64 | Start of the query window as a Unix timestamp in seconds
	to := int64(1704067200) // int64 | End of the query window as a Unix timestamp in seconds
	query := "query_example" // string | Tsuga query selecting spans to return. Omit to search all spans in the range with `*`. (optional) (default to "*")
	maxResults := int32(56) // int32 | Maximum number of spans to return. Defaults to 100. Valid values are 1 through 10000. (optional) (default to 100)
	clusterId := "clusterId_example" // string | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracesAPI.SearchSpans(context.Background()).From(from).To(to).Query(query).MaxResults(maxResults).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.SearchSpans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSpans`: SearchSpansResponse
	fmt.Fprintf(os.Stdout, "Response from `TracesAPI.SearchSpans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSpansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | Start of the query window as a Unix timestamp in seconds | 
 **to** | **int64** | End of the query window as a Unix timestamp in seconds | 
 **query** | **string** | Tsuga query selecting spans to return. Omit to search all spans in the range with &#x60;*&#x60;. | [default to &quot;*&quot;]
 **maxResults** | **int32** | Maximum number of spans to return. Defaults to 100. Valid values are 1 through 10000. | [default to 100]
 **clusterId** | **string** | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. | 

### Return type

[**SearchSpansResponse**](SearchSpansResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SummarizeTrace

> SummarizeTraceResponse SummarizeTrace(ctx).From(from).To(to).TraceId(traceId).ClusterId(clusterId).GroupByAttributes(groupByAttributes).MinLeafSpansToAggregate(minLeafSpansToAggregate).MaxParentDepth(maxParentDepth).AggregationHistogramBucketsSeconds(aggregationHistogramBucketsSeconds).SortSpans(sortSpans).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tsuga-dev/tsuga-go-sdk"
)

func main() {
	from := int64(1704067200) // int64 | Start of the query window as a Unix timestamp in seconds
	to := int64(1704067200) // int64 | End of the query window as a Unix timestamp in seconds
	traceId := "traceId_example" // string | Trace ID to summarize. Use the `traceId` returned by span search, the trace details page, or your tracing instrumentation. Required; the requested time range must include the trace.
	clusterId := "clusterId_example" // string | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. (optional)
	groupByAttributes := []string{"Inner_example"} // []string | Glob patterns of span attribute keys to use as additional grouping dimensions. Use this when spans should be summarized separately for matching attributes, such as `db.*`. Matches flattened dotted span attribute keys; omitted by default. (optional)
	minLeafSpansToAggregate := int32(56) // int32 | Minimum number of similar leaf spans in a group before Tsuga replaces them with one synthetic summary span. Use higher values for less aggressive summarization. Optional; defaults to 5 and must be at least 2. (optional) (default to 5)
	maxParentDepth := int32(56) // int32 | Maximum depth of recursive parent aggregation after leaf spans are summarized. Use this to limit parent-level aggregation. Use `0` to aggregate leaf spans only or `-1` for unlimited depth; defaults to `-1`. (optional) (default to -1)
	aggregationHistogramBucketsSeconds := []float32{float32(123)} // []float32 | Upper-bound bucket values, in seconds, for the latency histogram recorded on synthetic summary spans. Provide custom buckets when the default latency buckets are not suitable. Pass an empty array to omit histogram attributes; omitted values use Tsuga defaults from 5 ms through 10 s. (optional)
	sortSpans := true // bool | Whether Tsuga sorts returned spans by start time, using span ID as the tie-breaker for identical start times. Defaults to `true`; set to `false` only when reconstructing the trace tree yourself, in which case summary spans are appended after surviving spans. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracesAPI.SummarizeTrace(context.Background()).From(from).To(to).TraceId(traceId).ClusterId(clusterId).GroupByAttributes(groupByAttributes).MinLeafSpansToAggregate(minLeafSpansToAggregate).MaxParentDepth(maxParentDepth).AggregationHistogramBucketsSeconds(aggregationHistogramBucketsSeconds).SortSpans(sortSpans).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.SummarizeTrace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SummarizeTrace`: SummarizeTraceResponse
	fmt.Fprintf(os.Stdout, "Response from `TracesAPI.SummarizeTrace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSummarizeTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | Start of the query window as a Unix timestamp in seconds | 
 **to** | **int64** | End of the query window as a Unix timestamp in seconds | 
 **traceId** | **string** | Trace ID to summarize. Use the &#x60;traceId&#x60; returned by span search, the trace details page, or your tracing instrumentation. Required; the requested time range must include the trace. | 
 **clusterId** | **string** | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. | 
 **groupByAttributes** | **[]string** | Glob patterns of span attribute keys to use as additional grouping dimensions. Use this when spans should be summarized separately for matching attributes, such as &#x60;db.*&#x60;. Matches flattened dotted span attribute keys; omitted by default. | 
 **minLeafSpansToAggregate** | **int32** | Minimum number of similar leaf spans in a group before Tsuga replaces them with one synthetic summary span. Use higher values for less aggressive summarization. Optional; defaults to 5 and must be at least 2. | [default to 5]
 **maxParentDepth** | **int32** | Maximum depth of recursive parent aggregation after leaf spans are summarized. Use this to limit parent-level aggregation. Use &#x60;0&#x60; to aggregate leaf spans only or &#x60;-1&#x60; for unlimited depth; defaults to &#x60;-1&#x60;. | [default to -1]
 **aggregationHistogramBucketsSeconds** | **[]float32** | Upper-bound bucket values, in seconds, for the latency histogram recorded on synthetic summary spans. Provide custom buckets when the default latency buckets are not suitable. Pass an empty array to omit histogram attributes; omitted values use Tsuga defaults from 5 ms through 10 s. | 
 **sortSpans** | **bool** | Whether Tsuga sorts returned spans by start time, using span ID as the tie-breaker for identical start times. Defaults to &#x60;true&#x60;; set to &#x60;false&#x60; only when reconstructing the trace tree yourself, in which case summary spans are appended after surviving spans. | [default to true]

### Return type

[**SummarizeTraceResponse**](SummarizeTraceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

