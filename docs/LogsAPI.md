# \LogsAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetErrorPatternIncrease**](LogsAPI.md#GetErrorPatternIncrease) | **Get** /v1/logs/patterns/increase | 
[**GetLogAttributes**](LogsAPI.md#GetLogAttributes) | **Get** /v1/logs/attributes | 
[**GetLogPatterns**](LogsAPI.md#GetLogPatterns) | **Get** /v1/logs/patterns | 
[**GetNewErrorPatterns**](LogsAPI.md#GetNewErrorPatterns) | **Get** /v1/logs/patterns/new | 
[**SearchLogs**](LogsAPI.md#SearchLogs) | **Get** /v1/logs/search | 



## GetErrorPatternIncrease

> GetErrorPatternIncreaseResponse GetErrorPatternIncrease(ctx).From(from).To(to).Team(team).Env(env).ClusterId(clusterId).Execute()





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
	team := "team_example" // string | Team name to analyze for pattern increases. Required; match it to the `context.team` value on the logs you want Tsuga to inspect.
	env := "env_example" // string | Environment to restrict the analysis to. Optional; match it to the `context.env` value on the logs, or omit it to analyze all environments for the team. (optional)
	clusterId := "clusterId_example" // string | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.GetErrorPatternIncrease(context.Background()).From(from).To(to).Team(team).Env(env).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetErrorPatternIncrease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetErrorPatternIncrease`: GetErrorPatternIncreaseResponse
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.GetErrorPatternIncrease`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetErrorPatternIncreaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | Start of the query window as a Unix timestamp in seconds | 
 **to** | **int64** | End of the query window as a Unix timestamp in seconds | 
 **team** | **string** | Team name to analyze for pattern increases. Required; match it to the &#x60;context.team&#x60; value on the logs you want Tsuga to inspect. | 
 **env** | **string** | Environment to restrict the analysis to. Optional; match it to the &#x60;context.env&#x60; value on the logs, or omit it to analyze all environments for the team. | 
 **clusterId** | **string** | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. | 

### Return type

[**GetErrorPatternIncreaseResponse**](GetErrorPatternIncreaseResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogAttributes

> GetLogAttributesResponse GetLogAttributes(ctx).From(from).To(to).ClusterId(clusterId).Search(search).Execute()





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
	clusterId := "clusterId_example" // string | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. (optional)
	search := "search_example" // string | Return only attribute dot-paths containing this text, case-insensitively, e.g. `k8s` or `http.status`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.GetLogAttributes(context.Background()).From(from).To(to).ClusterId(clusterId).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetLogAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogAttributes`: GetLogAttributesResponse
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.GetLogAttributes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | Start of the query window as a Unix timestamp in seconds | 
 **to** | **int64** | End of the query window as a Unix timestamp in seconds | 
 **clusterId** | **string** | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. | 
 **search** | **string** | Return only attribute dot-paths containing this text, case-insensitively, e.g. &#x60;k8s&#x60; or &#x60;http.status&#x60;. | 

### Return type

[**GetLogAttributesResponse**](GetLogAttributesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogPatterns

> GetLogPatternsResponse GetLogPatterns(ctx).From(from).To(to).Query(query).ClusterId(clusterId).Execute()





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
	query := "query_example" // string | Tsuga query selecting logs to cluster into patterns.
	clusterId := "clusterId_example" // string | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.GetLogPatterns(context.Background()).From(from).To(to).Query(query).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetLogPatterns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogPatterns`: GetLogPatternsResponse
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.GetLogPatterns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogPatternsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | Start of the query window as a Unix timestamp in seconds | 
 **to** | **int64** | End of the query window as a Unix timestamp in seconds | 
 **query** | **string** | Tsuga query selecting logs to cluster into patterns. | 
 **clusterId** | **string** | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. | 

### Return type

[**GetLogPatternsResponse**](GetLogPatternsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNewErrorPatterns

> GetNewErrorPatternsResponse GetNewErrorPatterns(ctx).From(from).To(to).Team(team).Env(env).Service(service).ClusterId(clusterId).Execute()





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
	team := "team_example" // string | Team name to restrict pattern detection to. (optional)
	env := "env_example" // string | Environment to restrict pattern detection to. (optional)
	service := "service_example" // string | Service name to restrict pattern detection to. (optional)
	clusterId := "clusterId_example" // string | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.GetNewErrorPatterns(context.Background()).From(from).To(to).Team(team).Env(env).Service(service).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetNewErrorPatterns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNewErrorPatterns`: GetNewErrorPatternsResponse
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.GetNewErrorPatterns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNewErrorPatternsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | Start of the query window as a Unix timestamp in seconds | 
 **to** | **int64** | End of the query window as a Unix timestamp in seconds | 
 **team** | **string** | Team name to restrict pattern detection to. | 
 **env** | **string** | Environment to restrict pattern detection to. | 
 **service** | **string** | Service name to restrict pattern detection to. | 
 **clusterId** | **string** | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. | 

### Return type

[**GetNewErrorPatternsResponse**](GetNewErrorPatternsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchLogs

> SearchLogsResponse SearchLogs(ctx).From(from).To(to).Query(query).MaxResults(maxResults).ClusterId(clusterId).Execute()





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
	query := "query_example" // string | Tsuga query selecting logs to return. Omit to search all logs with `*`. (optional) (default to "*")
	maxResults := int32(56) // int32 | Maximum number of logs to return. Defaults to 100. Valid values are 1 through 1000. (optional) (default to 100)
	clusterId := "clusterId_example" // string | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.SearchLogs(context.Background()).From(from).To(to).Query(query).MaxResults(maxResults).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.SearchLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchLogs`: SearchLogsResponse
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.SearchLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | Start of the query window as a Unix timestamp in seconds | 
 **to** | **int64** | End of the query window as a Unix timestamp in seconds | 
 **query** | **string** | Tsuga query selecting logs to return. Omit to search all logs with &#x60;*&#x60;. | [default to &quot;*&quot;]
 **maxResults** | **int32** | Maximum number of logs to return. Defaults to 100. Valid values are 1 through 1000. | [default to 100]
 **clusterId** | **string** | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. | 

### Return type

[**SearchLogsResponse**](SearchLogsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

