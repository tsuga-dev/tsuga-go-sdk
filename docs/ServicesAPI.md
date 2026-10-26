# \ServicesAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceNotificationTarget**](ServicesAPI.md#CreateServiceNotificationTarget) | **Post** /v1/services/{serviceId}/notification-targets | 
[**DeleteServiceNotificationTarget**](ServicesAPI.md#DeleteServiceNotificationTarget) | **Delete** /v1/services/{serviceId}/notification-targets/{targetId} | 
[**GetService**](ServicesAPI.md#GetService) | **Get** /v1/services/{serviceId} | 
[**GetServiceGraph**](ServicesAPI.md#GetServiceGraph) | **Get** /v1/services/{serviceId}/graph | 
[**ListServiceNotificationTargets**](ServicesAPI.md#ListServiceNotificationTargets) | **Get** /v1/services/{serviceId}/notification-targets | 
[**QueryServices**](ServicesAPI.md#QueryServices) | **Post** /v1/services/query | 
[**UpdateServiceNotificationTarget**](ServicesAPI.md#UpdateServiceNotificationTarget) | **Put** /v1/services/{serviceId}/notification-targets/{targetId} | 



## CreateServiceNotificationTarget

> CreateServiceNotificationTargetResponse CreateServiceNotificationTarget(ctx, serviceId).CreateServiceNotificationTargetRequest(createServiceNotificationTargetRequest).Execute()





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
	serviceId := "serviceId_example" // string | Identifier of the service the notification target belongs to. Use the `id` returned by service list or get responses.
	createServiceNotificationTargetRequest := *openapiclient.NewCreateServiceNotificationTargetRequest(openapiclient.createNotificationRule_request_targets_inner_config{RuleTargetInputEmail: openapiclient.NewRuleTargetInputEmail("Type_example", []string{"Addresses_example"})}) // CreateServiceNotificationTargetRequest | Service notification target create request. Provide the delivery destination and an optional per-target rate limit.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.CreateServiceNotificationTarget(context.Background(), serviceId).CreateServiceNotificationTargetRequest(createServiceNotificationTargetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.CreateServiceNotificationTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceNotificationTarget`: CreateServiceNotificationTargetResponse
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.CreateServiceNotificationTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Identifier of the service the notification target belongs to. Use the &#x60;id&#x60; returned by service list or get responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceNotificationTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createServiceNotificationTargetRequest** | [**CreateServiceNotificationTargetRequest**](CreateServiceNotificationTargetRequest.md) | Service notification target create request. Provide the delivery destination and an optional per-target rate limit. | 

### Return type

[**CreateServiceNotificationTargetResponse**](CreateServiceNotificationTargetResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceNotificationTarget

> DeleteServiceNotificationTargetResponse DeleteServiceNotificationTarget(ctx, serviceId, targetId).Execute()





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
	serviceId := "serviceId_example" // string | Identifier of the service the notification target belongs to. Use the `id` returned by service list or get responses.
	targetId := "targetId_example" // string | Identifier of the notification target to delete. Use the `id` returned by service notification target list or create responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.DeleteServiceNotificationTarget(context.Background(), serviceId, targetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.DeleteServiceNotificationTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteServiceNotificationTarget`: DeleteServiceNotificationTargetResponse
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.DeleteServiceNotificationTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Identifier of the service the notification target belongs to. Use the &#x60;id&#x60; returned by service list or get responses. | 
**targetId** | **string** | Identifier of the notification target to delete. Use the &#x60;id&#x60; returned by service notification target list or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceNotificationTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteServiceNotificationTargetResponse**](DeleteServiceNotificationTargetResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetService

> GetServiceResponse GetService(ctx, serviceId).Execute()





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
	serviceId := "serviceId_example" // string | Identifier of the service to retrieve. Use the `id` returned by service query responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.GetService(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.GetService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetService`: GetServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.GetService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Identifier of the service to retrieve. Use the &#x60;id&#x60; returned by service query responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetServiceResponse**](GetServiceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceGraph

> GetServiceGraphResponse GetServiceGraph(ctx, serviceId).From(from).To(to).Query(query).Execute()





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
	serviceId := "serviceId_example" // string | Identifier of the service to use as the graph root. Use the `id` returned by service list or get responses.
	from := int64(1704067200) // int64 | Start of the service graph query window as a Unix timestamp in seconds. Provide both `from` and `to` for a deterministic graph window; otherwise Tsuga uses backend graph defaults. (optional)
	to := int64(1704067200) // int64 | End of the service graph query window as a Unix timestamp in seconds. Provide both `from` and `to` for a deterministic graph window; otherwise Tsuga uses backend graph defaults. (optional)
	query := "query_example" // string | Tsuga query used to filter spans considered for the graph. Defaults to `*`, which applies no additional filter beyond the service identity. (optional) (default to "*")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.GetServiceGraph(context.Background(), serviceId).From(from).To(to).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.GetServiceGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceGraph`: GetServiceGraphResponse
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.GetServiceGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Identifier of the service to use as the graph root. Use the &#x60;id&#x60; returned by service list or get responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int64** | Start of the service graph query window as a Unix timestamp in seconds. Provide both &#x60;from&#x60; and &#x60;to&#x60; for a deterministic graph window; otherwise Tsuga uses backend graph defaults. | 
 **to** | **int64** | End of the service graph query window as a Unix timestamp in seconds. Provide both &#x60;from&#x60; and &#x60;to&#x60; for a deterministic graph window; otherwise Tsuga uses backend graph defaults. | 
 **query** | **string** | Tsuga query used to filter spans considered for the graph. Defaults to &#x60;*&#x60;, which applies no additional filter beyond the service identity. | [default to &quot;*&quot;]

### Return type

[**GetServiceGraphResponse**](GetServiceGraphResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceNotificationTargets

> ListServiceNotificationTargetsResponse ListServiceNotificationTargets(ctx, serviceId).Execute()





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
	serviceId := "serviceId_example" // string | Identifier of the service the notification target belongs to. Use the `id` returned by service list or get responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ListServiceNotificationTargets(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ListServiceNotificationTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServiceNotificationTargets`: ListServiceNotificationTargetsResponse
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ListServiceNotificationTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Identifier of the service the notification target belongs to. Use the &#x60;id&#x60; returned by service list or get responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceNotificationTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListServiceNotificationTargetsResponse**](ListServiceNotificationTargetsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryServices

> QueryServicesResponse QueryServices(ctx).QueryServicesRequest(queryServicesRequest).Execute()





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
	queryServicesRequest := *openapiclient.NewQueryServicesRequest() // QueryServicesRequest | Service query request. Use filters, sorting, limit, and offset to page through services. Each service returns one row per environment it was observed in.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.QueryServices(context.Background()).QueryServicesRequest(queryServicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.QueryServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryServices`: QueryServicesResponse
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.QueryServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryServicesRequest** | [**QueryServicesRequest**](QueryServicesRequest.md) | Service query request. Use filters, sorting, limit, and offset to page through services. Each service returns one row per environment it was observed in. | 

### Return type

[**QueryServicesResponse**](QueryServicesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceNotificationTarget

> UpdateServiceNotificationTargetResponse UpdateServiceNotificationTarget(ctx, serviceId, targetId).UpdateServiceNotificationTargetRequest(updateServiceNotificationTargetRequest).Execute()





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
	serviceId := "serviceId_example" // string | Identifier of the service the notification target belongs to. Use the `id` returned by service list or get responses.
	targetId := "targetId_example" // string | Identifier of the notification target to update. Use the `id` returned by service notification target list or create responses.
	updateServiceNotificationTargetRequest := *openapiclient.NewUpdateServiceNotificationTargetRequest(openapiclient.createNotificationRule_request_targets_inner_config{RuleTargetInputEmail: openapiclient.NewRuleTargetInputEmail("Type_example", []string{"Addresses_example"})}, "TODO") // UpdateServiceNotificationTargetRequest | Service notification target update request. The submitted destination and rate limit overwrite the stored ones.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.UpdateServiceNotificationTarget(context.Background(), serviceId, targetId).UpdateServiceNotificationTargetRequest(updateServiceNotificationTargetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.UpdateServiceNotificationTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServiceNotificationTarget`: UpdateServiceNotificationTargetResponse
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.UpdateServiceNotificationTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Identifier of the service the notification target belongs to. Use the &#x60;id&#x60; returned by service list or get responses. | 
**targetId** | **string** | Identifier of the notification target to update. Use the &#x60;id&#x60; returned by service notification target list or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceNotificationTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateServiceNotificationTargetRequest** | [**UpdateServiceNotificationTargetRequest**](UpdateServiceNotificationTargetRequest.md) | Service notification target update request. The submitted destination and rate limit overwrite the stored ones. | 

### Return type

[**UpdateServiceNotificationTargetResponse**](UpdateServiceNotificationTargetResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

