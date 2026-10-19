# \MetricRoutesAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMetricRoute**](MetricRoutesAPI.md#CreateMetricRoute) | **Post** /v1/metrics/routes | 
[**DeleteMetricRoute**](MetricRoutesAPI.md#DeleteMetricRoute) | **Delete** /v1/metrics/routes/{id} | 
[**GetMetricRoute**](MetricRoutesAPI.md#GetMetricRoute) | **Get** /v1/metrics/routes/{id} | 
[**ListMetricRoutes**](MetricRoutesAPI.md#ListMetricRoutes) | **Get** /v1/metrics/routes | 
[**UpdateMetricRoute**](MetricRoutesAPI.md#UpdateMetricRoute) | **Put** /v1/metrics/routes/{id} | 



## CreateMetricRoute

> CreateMetricRouteResponse CreateMetricRoute(ctx).CreateMetricRouteRequest(createMetricRouteRequest).Execute()





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
	createMetricRouteRequest := *openapiclient.NewCreateMetricRouteRequest("Name_example", false, "MetricNameRegex_example", "Owner_example", []openapiclient.CreateMetricRouteRequestProcessorsInner{openapiclient.createMetricRoute_request_processors_inner{CreateMetricRouteRequestProcessorsInnerOneOf: openapiclient.NewCreateMetricRouteRequestProcessorsInnerOneOf("Id_example", "Type_example", *openapiclient.NewCreateMetricRouteRequestProcessorsInnerOneOfParams([]openapiclient.CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner{*openapiclient.NewCreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner("OriginAttribute_example", "TargetAttribute_example")}, "Subtype_example"))}}) // CreateMetricRouteRequest | Metrics route create or update request. Provide the route metadata, metric name regex, owner team, enabled state, tags, and ordered processors.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricRoutesAPI.CreateMetricRoute(context.Background()).CreateMetricRouteRequest(createMetricRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricRoutesAPI.CreateMetricRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMetricRoute`: CreateMetricRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricRoutesAPI.CreateMetricRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMetricRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMetricRouteRequest** | [**CreateMetricRouteRequest**](CreateMetricRouteRequest.md) | Metrics route create or update request. Provide the route metadata, metric name regex, owner team, enabled state, tags, and ordered processors. | 

### Return type

[**CreateMetricRouteResponse**](CreateMetricRouteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMetricRoute

> DeleteMetricRouteResponse DeleteMetricRoute(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the metric route to delete. Use the `id` returned by metric route list, create, or update responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricRoutesAPI.DeleteMetricRoute(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricRoutesAPI.DeleteMetricRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMetricRoute`: DeleteMetricRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricRoutesAPI.DeleteMetricRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the metric route to delete. Use the &#x60;id&#x60; returned by metric route list, create, or update responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMetricRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteMetricRouteResponse**](DeleteMetricRouteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricRoute

> GetMetricRouteResponse GetMetricRoute(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the metric route to retrieve. Use the `id` returned by metric route list or create responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricRoutesAPI.GetMetricRoute(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricRoutesAPI.GetMetricRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricRoute`: GetMetricRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricRoutesAPI.GetMetricRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the metric route to retrieve. Use the &#x60;id&#x60; returned by metric route list or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMetricRouteResponse**](GetMetricRouteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetricRoutes

> ListMetricRoutesResponse ListMetricRoutes(ctx).Limit(limit).Offset(offset).Execute()





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
	limit := int32(56) // int32 | Maximum number of items to return in this page. Omit to use the public API default of 100. (optional) (default to 100)
	offset := int32(56) // int32 | Zero-based index of the first matching item to return. Increase it with `limit` to request later pages. If `limit` is provided without `offset`, the offset defaults to 0. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricRoutesAPI.ListMetricRoutes(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricRoutesAPI.ListMetricRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMetricRoutes`: ListMetricRoutesResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricRoutesAPI.ListMetricRoutes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetricRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [default to 100]
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 

### Return type

[**ListMetricRoutesResponse**](ListMetricRoutesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMetricRoute

> UpdateMetricRouteResponse UpdateMetricRoute(ctx, id).UpdateMetricRouteRequest(updateMetricRouteRequest).Execute()





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
	id := "id_example" // string | Identifier of the metric route to update. Use the `id` returned by metric route list, get, or create responses.
	updateMetricRouteRequest := *openapiclient.NewUpdateMetricRouteRequest("Name_example", false, "MetricNameRegex_example", "Owner_example", []openapiclient.UpdateMetricRouteRequestProcessorsInner{openapiclient.updateMetricRoute_request_processors_inner{CreateMetricRouteRequestProcessorsInnerOneOf: openapiclient.NewCreateMetricRouteRequestProcessorsInnerOneOf("Id_example", "Type_example", *openapiclient.NewCreateMetricRouteRequestProcessorsInnerOneOfParams([]openapiclient.CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner{*openapiclient.NewCreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner("OriginAttribute_example", "TargetAttribute_example")}, "Subtype_example"))}}) // UpdateMetricRouteRequest | Metrics route create or update request. Provide the route metadata, metric name regex, owner team, enabled state, tags, and ordered processors.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricRoutesAPI.UpdateMetricRoute(context.Background(), id).UpdateMetricRouteRequest(updateMetricRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricRoutesAPI.UpdateMetricRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMetricRoute`: UpdateMetricRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricRoutesAPI.UpdateMetricRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the metric route to update. Use the &#x60;id&#x60; returned by metric route list, get, or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetricRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMetricRouteRequest** | [**UpdateMetricRouteRequest**](UpdateMetricRouteRequest.md) | Metrics route create or update request. Provide the route metadata, metric name regex, owner team, enabled state, tags, and ordered processors. | 

### Return type

[**UpdateMetricRouteResponse**](UpdateMetricRouteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

