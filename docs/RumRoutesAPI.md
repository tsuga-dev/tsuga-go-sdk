# \RumRoutesAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRumRoute**](RumRoutesAPI.md#CreateRumRoute) | **Post** /v1/rum/routes | 
[**DeleteRumRoute**](RumRoutesAPI.md#DeleteRumRoute) | **Delete** /v1/rum/routes/{id} | 
[**EnableRumRoute**](RumRoutesAPI.md#EnableRumRoute) | **Put** /v1/rum/routes/{id}/enable | 
[**GetRumRoute**](RumRoutesAPI.md#GetRumRoute) | **Get** /v1/rum/routes/{id} | 
[**ListRumRoutes**](RumRoutesAPI.md#ListRumRoutes) | **Get** /v1/rum/routes | 
[**ModifyRumRouteRanking**](RumRoutesAPI.md#ModifyRumRouteRanking) | **Put** /v1/rum/routes/{id}/modify-ranking | 
[**UpdateRumRoute**](RumRoutesAPI.md#UpdateRumRoute) | **Put** /v1/rum/routes/{id} | 



## CreateRumRoute

> CreateRumRouteResponse CreateRumRoute(ctx).CreateRumRouteRequest(createRumRouteRequest).Execute()





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
	createRumRouteRequest := *openapiclient.NewCreateRumRouteRequest("Name_example", false, "Query_example", "Owner_example", []openapiclient.InputRumProcessor{*openapiclient.NewInputRumProcessor("Id_example", "Type_example", *openapiclient.NewInputRumProcessorAnyOfParams([]openapiclient.InputRumProcessorAnyOfParamsItemsInner{*openapiclient.NewInputRumProcessorAnyOfParamsItemsInner("Query_example", []openapiclient.InputRumProcessor{*openapiclient.NewInputRumProcessor("Id_example", "Type_example", *openapiclient.NewInputRumProcessorAnyOfParams([]openapiclient.InputRumProcessorAnyOfParamsItemsInner{*openapiclient.NewInputRumProcessorAnyOfParamsItemsInner("Query_example", []openapiclient.InputRumProcessor{})}))})}))}) // CreateRumRouteRequest | RUM route create or update request. Provide the route metadata, source query, owner team, enabled state, tags, and ordered processors.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RumRoutesAPI.CreateRumRoute(context.Background()).CreateRumRouteRequest(createRumRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumRoutesAPI.CreateRumRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRumRoute`: CreateRumRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `RumRoutesAPI.CreateRumRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRumRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRumRouteRequest** | [**CreateRumRouteRequest**](CreateRumRouteRequest.md) | RUM route create or update request. Provide the route metadata, source query, owner team, enabled state, tags, and ordered processors. | 

### Return type

[**CreateRumRouteResponse**](CreateRumRouteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRumRoute

> DeleteRumRouteResponse DeleteRumRoute(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the RUM route to delete. Use the `id` returned by RUM route list, create, or update responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RumRoutesAPI.DeleteRumRoute(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumRoutesAPI.DeleteRumRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRumRoute`: DeleteRumRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `RumRoutesAPI.DeleteRumRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the RUM route to delete. Use the &#x60;id&#x60; returned by RUM route list, create, or update responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRumRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteRumRouteResponse**](DeleteRumRouteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableRumRoute

> EnableRumRouteResponse EnableRumRoute(ctx, id).EnableRumRouteRequest(enableRumRouteRequest).Execute()





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
	id := "id_example" // string | Identifier of the RUM route to enable or disable.
	enableRumRouteRequest := *openapiclient.NewEnableRumRouteRequest(false) // EnableRumRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RumRoutesAPI.EnableRumRoute(context.Background(), id).EnableRumRouteRequest(enableRumRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumRoutesAPI.EnableRumRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableRumRoute`: EnableRumRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `RumRoutesAPI.EnableRumRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the RUM route to enable or disable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableRumRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableRumRouteRequest** | [**EnableRumRouteRequest**](EnableRumRouteRequest.md) |  | 

### Return type

[**EnableRumRouteResponse**](EnableRumRouteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRumRoute

> GetRumRouteResponse GetRumRoute(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the RUM route to retrieve. Use the `id` returned by RUM route list or create responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RumRoutesAPI.GetRumRoute(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumRoutesAPI.GetRumRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRumRoute`: GetRumRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `RumRoutesAPI.GetRumRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the RUM route to retrieve. Use the &#x60;id&#x60; returned by RUM route list or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRumRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRumRouteResponse**](GetRumRouteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRumRoutes

> ListRumRoutesResponse ListRumRoutes(ctx).Limit(limit).Offset(offset).Execute()





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
	resp, r, err := apiClient.RumRoutesAPI.ListRumRoutes(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumRoutesAPI.ListRumRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRumRoutes`: ListRumRoutesResponse
	fmt.Fprintf(os.Stdout, "Response from `RumRoutesAPI.ListRumRoutes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRumRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [default to 100]
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 

### Return type

[**ListRumRoutesResponse**](ListRumRoutesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyRumRouteRanking

> ModifyRumRouteRankingResponse ModifyRumRouteRanking(ctx, id).ModifyRumRouteRankingRequest(modifyRumRouteRankingRequest).Execute()





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
	id := "id_example" // string | Identifier of the RUM route to move.
	modifyRumRouteRankingRequest := *openapiclient.NewModifyRumRouteRankingRequest(int32(123)) // ModifyRumRouteRankingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RumRoutesAPI.ModifyRumRouteRanking(context.Background(), id).ModifyRumRouteRankingRequest(modifyRumRouteRankingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumRoutesAPI.ModifyRumRouteRanking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRumRouteRanking`: ModifyRumRouteRankingResponse
	fmt.Fprintf(os.Stdout, "Response from `RumRoutesAPI.ModifyRumRouteRanking`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the RUM route to move. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRumRouteRankingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyRumRouteRankingRequest** | [**ModifyRumRouteRankingRequest**](ModifyRumRouteRankingRequest.md) |  | 

### Return type

[**ModifyRumRouteRankingResponse**](ModifyRumRouteRankingResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRumRoute

> UpdateRumRouteResponse UpdateRumRoute(ctx, id).UpdateRumRouteRequest(updateRumRouteRequest).Execute()





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
	id := "id_example" // string | Identifier of the RUM route to update. Use the `id` returned by RUM route list, get, or create responses.
	updateRumRouteRequest := *openapiclient.NewUpdateRumRouteRequest("Name_example", false, "Query_example", "Owner_example", []openapiclient.InputRumProcessor{*openapiclient.NewInputRumProcessor("Id_example", "Type_example", *openapiclient.NewInputRumProcessorAnyOfParams([]openapiclient.InputRumProcessorAnyOfParamsItemsInner{*openapiclient.NewInputRumProcessorAnyOfParamsItemsInner("Query_example", []openapiclient.InputRumProcessor{*openapiclient.NewInputRumProcessor("Id_example", "Type_example", *openapiclient.NewInputRumProcessorAnyOfParams([]openapiclient.InputRumProcessorAnyOfParamsItemsInner{*openapiclient.NewInputRumProcessorAnyOfParamsItemsInner("Query_example", []openapiclient.InputRumProcessor{})}))})}))}) // UpdateRumRouteRequest | RUM route create or update request. Provide the route metadata, source query, owner team, enabled state, tags, and ordered processors.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RumRoutesAPI.UpdateRumRoute(context.Background(), id).UpdateRumRouteRequest(updateRumRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumRoutesAPI.UpdateRumRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRumRoute`: UpdateRumRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `RumRoutesAPI.UpdateRumRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the RUM route to update. Use the &#x60;id&#x60; returned by RUM route list, get, or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRumRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRumRouteRequest** | [**UpdateRumRouteRequest**](UpdateRumRouteRequest.md) | RUM route create or update request. Provide the route metadata, source query, owner team, enabled state, tags, and ordered processors. | 

### Return type

[**UpdateRumRouteResponse**](UpdateRumRouteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

