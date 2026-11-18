# \SpanRoutesAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpanRoute**](SpanRoutesAPI.md#CreateSpanRoute) | **Post** /v1/spans/routes | 
[**DeleteSpanRoute**](SpanRoutesAPI.md#DeleteSpanRoute) | **Delete** /v1/spans/routes/{id} | 
[**GetSpanRoute**](SpanRoutesAPI.md#GetSpanRoute) | **Get** /v1/spans/routes/{id} | 
[**ListSpanRoutes**](SpanRoutesAPI.md#ListSpanRoutes) | **Get** /v1/spans/routes | 
[**UpdateSpanRoute**](SpanRoutesAPI.md#UpdateSpanRoute) | **Put** /v1/spans/routes/{id} | 



## CreateSpanRoute

> CreateSpanRouteResponse CreateSpanRoute(ctx).CreateSpanRouteRequest(createSpanRouteRequest).Execute()





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
	createSpanRouteRequest := *openapiclient.NewCreateSpanRouteRequest("Name_example", false, "Query_example", "Owner_example", []openapiclient.InputSpansProcessor{*openapiclient.NewInputSpansProcessor("Id_example", "Type_example", *openapiclient.NewInputSpansProcessorAnyOfParams([]openapiclient.InputSpansProcessorAnyOfParamsItemsInner{*openapiclient.NewInputSpansProcessorAnyOfParamsItemsInner("Query_example", []openapiclient.InputSpansProcessor{*openapiclient.NewInputSpansProcessor("Id_example", "Type_example", *openapiclient.NewInputSpansProcessorAnyOfParams([]openapiclient.InputSpansProcessorAnyOfParamsItemsInner{*openapiclient.NewInputSpansProcessorAnyOfParamsItemsInner("Query_example", []openapiclient.InputSpansProcessor{})}))})}))}) // CreateSpanRouteRequest | Span route create or update request. Provide the route metadata, source query, owner team, enabled state, tags, and ordered processors.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpanRoutesAPI.CreateSpanRoute(context.Background()).CreateSpanRouteRequest(createSpanRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpanRoutesAPI.CreateSpanRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSpanRoute`: CreateSpanRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `SpanRoutesAPI.CreateSpanRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpanRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSpanRouteRequest** | [**CreateSpanRouteRequest**](CreateSpanRouteRequest.md) | Span route create or update request. Provide the route metadata, source query, owner team, enabled state, tags, and ordered processors. | 

### Return type

[**CreateSpanRouteResponse**](CreateSpanRouteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSpanRoute

> DeleteSpanRouteResponse DeleteSpanRoute(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the span route to delete. Use the `id` returned by span route list, create, or update responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpanRoutesAPI.DeleteSpanRoute(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpanRoutesAPI.DeleteSpanRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSpanRoute`: DeleteSpanRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `SpanRoutesAPI.DeleteSpanRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the span route to delete. Use the &#x60;id&#x60; returned by span route list, create, or update responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpanRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteSpanRouteResponse**](DeleteSpanRouteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpanRoute

> GetSpanRouteResponse GetSpanRoute(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the span route to retrieve. Use the `id` returned by span route list or create responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpanRoutesAPI.GetSpanRoute(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpanRoutesAPI.GetSpanRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpanRoute`: GetSpanRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `SpanRoutesAPI.GetSpanRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the span route to retrieve. Use the &#x60;id&#x60; returned by span route list or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpanRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSpanRouteResponse**](GetSpanRouteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpanRoutes

> ListSpanRoutesResponse ListSpanRoutes(ctx).Limit(limit).Offset(offset).Execute()





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
	resp, r, err := apiClient.SpanRoutesAPI.ListSpanRoutes(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpanRoutesAPI.ListSpanRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSpanRoutes`: ListSpanRoutesResponse
	fmt.Fprintf(os.Stdout, "Response from `SpanRoutesAPI.ListSpanRoutes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSpanRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [default to 100]
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 

### Return type

[**ListSpanRoutesResponse**](ListSpanRoutesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpanRoute

> UpdateSpanRouteResponse UpdateSpanRoute(ctx, id).UpdateSpanRouteRequest(updateSpanRouteRequest).Execute()





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
	id := "id_example" // string | Identifier of the span route to update. Use the `id` returned by span route list, get, or create responses.
	updateSpanRouteRequest := *openapiclient.NewUpdateSpanRouteRequest("Name_example", false, "Query_example", "Owner_example", []openapiclient.InputSpansProcessor{*openapiclient.NewInputSpansProcessor("Id_example", "Type_example", *openapiclient.NewInputSpansProcessorAnyOfParams([]openapiclient.InputSpansProcessorAnyOfParamsItemsInner{*openapiclient.NewInputSpansProcessorAnyOfParamsItemsInner("Query_example", []openapiclient.InputSpansProcessor{*openapiclient.NewInputSpansProcessor("Id_example", "Type_example", *openapiclient.NewInputSpansProcessorAnyOfParams([]openapiclient.InputSpansProcessorAnyOfParamsItemsInner{*openapiclient.NewInputSpansProcessorAnyOfParamsItemsInner("Query_example", []openapiclient.InputSpansProcessor{})}))})}))}) // UpdateSpanRouteRequest | Span route create or update request. Provide the route metadata, source query, owner team, enabled state, tags, and ordered processors.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpanRoutesAPI.UpdateSpanRoute(context.Background(), id).UpdateSpanRouteRequest(updateSpanRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpanRoutesAPI.UpdateSpanRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSpanRoute`: UpdateSpanRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `SpanRoutesAPI.UpdateSpanRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the span route to update. Use the &#x60;id&#x60; returned by span route list, get, or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpanRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSpanRouteRequest** | [**UpdateSpanRouteRequest**](UpdateSpanRouteRequest.md) | Span route create or update request. Provide the route metadata, source query, owner team, enabled state, tags, and ordered processors. | 

### Return type

[**UpdateSpanRouteResponse**](UpdateSpanRouteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

