# \RoutesAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoute**](RoutesAPI.md#CreateRoute) | **Post** /v1/routes | 
[**DeleteRoute**](RoutesAPI.md#DeleteRoute) | **Delete** /v1/routes/{id} | 
[**GetRoute**](RoutesAPI.md#GetRoute) | **Get** /v1/routes/{id} | 
[**ListRoutes**](RoutesAPI.md#ListRoutes) | **Get** /v1/routes | 
[**UpdateRoute**](RoutesAPI.md#UpdateRoute) | **Put** /v1/routes/{id} | 



## CreateRoute

> CreateRouteResponse CreateRoute(ctx).CreateRouteRequest(createRouteRequest).Execute()





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
	createRouteRequest := *openapiclient.NewCreateRouteRequest("Name_example", false, "Query_example", "Owner_example", []openapiclient.InputProcessor{*openapiclient.NewInputProcessor("Id_example", "Type_example", *openapiclient.NewInputProcessorAnyOf3Params([]openapiclient.InputProcessorAnyOf3ParamsItemsInner{*openapiclient.NewInputProcessorAnyOf3ParamsItemsInner("Query_example", []openapiclient.InputProcessor{*openapiclient.NewInputProcessor("Id_example", "Type_example", *openapiclient.NewInputProcessorAnyOf3Params([]openapiclient.InputProcessorAnyOf3ParamsItemsInner{*openapiclient.NewInputProcessorAnyOf3ParamsItemsInner("Query_example", []openapiclient.InputProcessor{})}))})}))}) // CreateRouteRequest | Log route create or update request. Provide the route metadata, source query, owner team, enabled state, tags, and ordered processors.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.CreateRoute(context.Background()).CreateRouteRequest(createRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.CreateRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoute`: CreateRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.CreateRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRouteRequest** | [**CreateRouteRequest**](CreateRouteRequest.md) | Log route create or update request. Provide the route metadata, source query, owner team, enabled state, tags, and ordered processors. | 

### Return type

[**CreateRouteResponse**](CreateRouteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoute

> DeleteRouteResponse DeleteRoute(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the log route to delete. Use the `id` returned by route list, create, or update responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.DeleteRoute(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.DeleteRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRoute`: DeleteRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.DeleteRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the log route to delete. Use the &#x60;id&#x60; returned by route list, create, or update responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteRouteResponse**](DeleteRouteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoute

> GetRouteResponse GetRoute(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the log route to retrieve. Use the `id` returned by route list or create responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.GetRoute(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.GetRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoute`: GetRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.GetRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the log route to retrieve. Use the &#x60;id&#x60; returned by route list or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRouteResponse**](GetRouteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoutes

> ListRoutesResponse ListRoutes(ctx).Limit(limit).Offset(offset).Execute()





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
	resp, r, err := apiClient.RoutesAPI.ListRoutes(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.ListRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoutes`: ListRoutesResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.ListRoutes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [default to 100]
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 

### Return type

[**ListRoutesResponse**](ListRoutesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoute

> UpdateRouteResponse UpdateRoute(ctx, id).UpdateRouteRequest(updateRouteRequest).Execute()





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
	id := "id_example" // string | Identifier of the log route to update. Use the `id` returned by route list, get, or create responses.
	updateRouteRequest := *openapiclient.NewUpdateRouteRequest("Name_example", false, "Query_example", "Owner_example", []openapiclient.InputProcessor{*openapiclient.NewInputProcessor("Id_example", "Type_example", *openapiclient.NewInputProcessorAnyOf3Params([]openapiclient.InputProcessorAnyOf3ParamsItemsInner{*openapiclient.NewInputProcessorAnyOf3ParamsItemsInner("Query_example", []openapiclient.InputProcessor{*openapiclient.NewInputProcessor("Id_example", "Type_example", *openapiclient.NewInputProcessorAnyOf3Params([]openapiclient.InputProcessorAnyOf3ParamsItemsInner{*openapiclient.NewInputProcessorAnyOf3ParamsItemsInner("Query_example", []openapiclient.InputProcessor{})}))})}))}) // UpdateRouteRequest | Log route create or update request. Provide the route metadata, source query, owner team, enabled state, tags, and ordered processors.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.UpdateRoute(context.Background(), id).UpdateRouteRequest(updateRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.UpdateRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoute`: UpdateRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.UpdateRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the log route to update. Use the &#x60;id&#x60; returned by route list, get, or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRouteRequest** | [**UpdateRouteRequest**](UpdateRouteRequest.md) | Log route create or update request. Provide the route metadata, source query, owner team, enabled state, tags, and ordered processors. | 

### Return type

[**UpdateRouteResponse**](UpdateRouteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

