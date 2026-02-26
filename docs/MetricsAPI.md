# \MetricsAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetric**](MetricsAPI.md#GetMetric) | **Get** /v1/metrics/{name} | 
[**ListMetrics**](MetricsAPI.md#ListMetrics) | **Get** /v1/metrics | 



## GetMetric

> GetMetricResponse GetMetric(ctx, name).From(from).To(to).Execute()





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
	from := int64(1704067200) // int64 | Start timestamp in seconds
	to := int64(1704067200) // int64 | End timestamp in seconds
	name := "name_example" // string | The metric name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GetMetric(context.Background(), name).From(from).To(to).Execute()
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
**name** | **string** | The metric name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | Start timestamp in seconds | 
 **to** | **int64** | End timestamp in seconds | 


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

> ListMetricsResponse ListMetrics(ctx).From(from).To(to).Execute()





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
	from := int64(1704067200) // int64 | Start timestamp in seconds
	to := int64(1704067200) // int64 | End timestamp in seconds

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.ListMetrics(context.Background()).From(from).To(to).Execute()
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
 **from** | **int64** | Start timestamp in seconds | 
 **to** | **int64** | End timestamp in seconds | 

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

