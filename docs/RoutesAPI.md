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

> CreateRoute200Response CreateRoute(ctx).CreateRouteRequest(createRouteRequest).Execute()





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
	createRouteRequest := *openapiclient.NewCreateRouteRequest("Name_example", false, "Query_example", "Owner_example", []openapiclient.CreateRouteRequestProcessorsInner{*openapiclient.NewCreateRouteRequestProcessorsInner("Id_example", "Name_example", "Type_example", *openapiclient.NewCreateRouteRequestProcessorsInnerAnyOf3Params([]openapiclient.CreateRouteRequestProcessorsInnerAnyOf3ParamsItemsInner{*openapiclient.NewCreateRouteRequestProcessorsInnerAnyOf3ParamsItemsInner("Query_example", []openapiclient.Processor{*openapiclient.NewProcessor("Id_example", "Name_example", "Type_example", *openapiclient.NewCreateRouteRequestProcessorsInnerAnyOf3Params([]openapiclient.CreateRouteRequestProcessorsInnerAnyOf3ParamsItemsInner{*openapiclient.NewCreateRouteRequestProcessorsInnerAnyOf3ParamsItemsInner("Query_example", []openapiclient.Processor{*openapiclient.NewProcessor("Id_example", "Name_example", "Type_example", )})}))})}))}) // CreateRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.CreateRoute(context.Background()).CreateRouteRequest(createRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.CreateRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoute`: CreateRoute200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.CreateRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRouteRequest** | [**CreateRouteRequest**](CreateRouteRequest.md) |  | 

### Return type

[**CreateRoute200Response**](CreateRoute200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoute

> DeleteDashboard200Response DeleteRoute(ctx, id).Execute()





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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.DeleteRoute(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.DeleteRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRoute`: DeleteDashboard200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.DeleteRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteDashboard200Response**](DeleteDashboard200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoute

> CreateRoute200Response GetRoute(ctx, id).Execute()





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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.GetRoute(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.GetRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoute`: CreateRoute200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.GetRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateRoute200Response**](CreateRoute200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoutes

> ListRoutes200Response ListRoutes(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.ListRoutes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.ListRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoutes`: ListRoutes200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.ListRoutes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRoutesRequest struct via the builder pattern


### Return type

[**ListRoutes200Response**](ListRoutes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoute

> CreateRoute200Response UpdateRoute(ctx, id).CreateRouteRequest(createRouteRequest).Execute()





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
	id := "id_example" // string | 
	createRouteRequest := *openapiclient.NewCreateRouteRequest("Name_example", false, "Query_example", "Owner_example", []openapiclient.CreateRouteRequestProcessorsInner{*openapiclient.NewCreateRouteRequestProcessorsInner("Id_example", "Name_example", "Type_example", *openapiclient.NewCreateRouteRequestProcessorsInnerAnyOf3Params([]openapiclient.CreateRouteRequestProcessorsInnerAnyOf3ParamsItemsInner{*openapiclient.NewCreateRouteRequestProcessorsInnerAnyOf3ParamsItemsInner("Query_example", []openapiclient.Processor{*openapiclient.NewProcessor("Id_example", "Name_example", "Type_example", *openapiclient.NewCreateRouteRequestProcessorsInnerAnyOf3Params([]openapiclient.CreateRouteRequestProcessorsInnerAnyOf3ParamsItemsInner{*openapiclient.NewCreateRouteRequestProcessorsInnerAnyOf3ParamsItemsInner("Query_example", []openapiclient.Processor{*openapiclient.NewProcessor("Id_example", "Name_example", "Type_example", )})}))})}))}) // CreateRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.UpdateRoute(context.Background(), id).CreateRouteRequest(createRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.UpdateRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoute`: CreateRoute200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.UpdateRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRouteRequest** | [**CreateRouteRequest**](CreateRouteRequest.md) |  | 

### Return type

[**CreateRoute200Response**](CreateRoute200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

