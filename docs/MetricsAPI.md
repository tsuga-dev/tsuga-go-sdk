# \MetricsAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetric**](MetricsAPI.md#GetMetric) | **Get** /v1/metrics/{name} | 
[**ListMetrics**](MetricsAPI.md#ListMetrics) | **Get** /v1/metrics | 



## GetMetric

> GetMetricResponse GetMetric(ctx, name).From(from).To(to).ClusterId(clusterId).Execute()





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
	from := int64(1704067200) // int64 | Unix timestamp in seconds. Required, but does not scope the result.
	to := int64(1704067200) // int64 | Unix timestamp in seconds. Required, but does not scope the result.
	name := "name_example" // string | Metric name to retrieve, for example `http.server.request.duration`.
	clusterId := "clusterId_example" // string | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GetMetric(context.Background(), name).From(from).To(to).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GetMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetric`: GetMetricResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GetMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Metric name to retrieve, for example &#x60;http.server.request.duration&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | Unix timestamp in seconds. Required, but does not scope the result. | 
 **to** | **int64** | Unix timestamp in seconds. Required, but does not scope the result. | 

 **clusterId** | **string** | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. | 

### Return type

[**GetMetricResponse**](GetMetricResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetrics

> ListMetricsResponse ListMetrics(ctx).From(from).To(to).ClusterId(clusterId).Limit(limit).Offset(offset).Search(search).Execute()





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
	from := int64(1704067200) // int64 | Unix timestamp in seconds. Required, but does not scope the result.
	to := int64(1704067200) // int64 | Unix timestamp in seconds. Required, but does not scope the result.
	clusterId := "clusterId_example" // string | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. (optional)
	limit := int32(56) // int32 | Maximum number of items to return in this page. Omit to use the public API default of 100. (optional) (default to 100)
	offset := int32(56) // int32 | Zero-based index of the first matching item to return. Increase it with `limit` to request later pages. If `limit` is provided without `offset`, the offset defaults to 0. (optional)
	search := "search_example" // string | Keep only metric names containing this value, matched case-insensitively. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.ListMetrics(context.Background()).From(from).To(to).ClusterId(clusterId).Limit(limit).Offset(offset).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.ListMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMetrics`: ListMetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.ListMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | Unix timestamp in seconds. Required, but does not scope the result. | 
 **to** | **int64** | Unix timestamp in seconds. Required, but does not scope the result. | 
 **clusterId** | **string** | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. | 
 **limit** | **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [default to 100]
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 
 **search** | **string** | Keep only metric names containing this value, matched case-insensitively. | 

### Return type

[**ListMetricsResponse**](ListMetricsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

