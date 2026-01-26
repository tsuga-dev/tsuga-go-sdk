# \RetentionPoliciesAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRetentionPolicy**](RetentionPoliciesAPI.md#CreateRetentionPolicy) | **Post** /v1/retention-policies | 
[**DeleteRetentionPolicy**](RetentionPoliciesAPI.md#DeleteRetentionPolicy) | **Delete** /v1/retention-policies/{id} | 
[**GetRetentionPolicy**](RetentionPoliciesAPI.md#GetRetentionPolicy) | **Get** /v1/retention-policies/{id} | 
[**ListRetentionPolicies**](RetentionPoliciesAPI.md#ListRetentionPolicies) | **Get** /v1/retention-policies | 
[**UpdateRetentionPolicy**](RetentionPoliciesAPI.md#UpdateRetentionPolicy) | **Put** /v1/retention-policies/{id} | 



## CreateRetentionPolicy

> CreateRetentionPolicy200Response CreateRetentionPolicy(ctx).CreateRetentionPolicyRequest(createRetentionPolicyRequest).Execute()





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
	createRetentionPolicyRequest := *openapiclient.NewCreateRetentionPolicyRequest("DataSource_example", "DurationDays_example", false) // CreateRetentionPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetentionPoliciesAPI.CreateRetentionPolicy(context.Background()).CreateRetentionPolicyRequest(createRetentionPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetentionPoliciesAPI.CreateRetentionPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRetentionPolicy`: CreateRetentionPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `RetentionPoliciesAPI.CreateRetentionPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRetentionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRetentionPolicyRequest** | [**CreateRetentionPolicyRequest**](CreateRetentionPolicyRequest.md) |  | 

### Return type

[**CreateRetentionPolicy200Response**](CreateRetentionPolicy200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRetentionPolicy

> DeleteIngestionApiKey200Response DeleteRetentionPolicy(ctx, id).Execute()





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
	resp, r, err := apiClient.RetentionPoliciesAPI.DeleteRetentionPolicy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetentionPoliciesAPI.DeleteRetentionPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRetentionPolicy`: DeleteIngestionApiKey200Response
	fmt.Fprintf(os.Stdout, "Response from `RetentionPoliciesAPI.DeleteRetentionPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRetentionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteIngestionApiKey200Response**](DeleteIngestionApiKey200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRetentionPolicy

> CreateRetentionPolicy200Response GetRetentionPolicy(ctx, id).Execute()





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
	resp, r, err := apiClient.RetentionPoliciesAPI.GetRetentionPolicy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetentionPoliciesAPI.GetRetentionPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRetentionPolicy`: CreateRetentionPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `RetentionPoliciesAPI.GetRetentionPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRetentionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateRetentionPolicy200Response**](CreateRetentionPolicy200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRetentionPolicies

> ListRetentionPolicies200Response ListRetentionPolicies(ctx).Execute()





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
	resp, r, err := apiClient.RetentionPoliciesAPI.ListRetentionPolicies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetentionPoliciesAPI.ListRetentionPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRetentionPolicies`: ListRetentionPolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `RetentionPoliciesAPI.ListRetentionPolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRetentionPoliciesRequest struct via the builder pattern


### Return type

[**ListRetentionPolicies200Response**](ListRetentionPolicies200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRetentionPolicy

> CreateRetentionPolicy200Response UpdateRetentionPolicy(ctx, id).CreateRetentionPolicyRequest(createRetentionPolicyRequest).Execute()





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
	createRetentionPolicyRequest := *openapiclient.NewCreateRetentionPolicyRequest("DataSource_example", "DurationDays_example", false) // CreateRetentionPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetentionPoliciesAPI.UpdateRetentionPolicy(context.Background(), id).CreateRetentionPolicyRequest(createRetentionPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetentionPoliciesAPI.UpdateRetentionPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRetentionPolicy`: CreateRetentionPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `RetentionPoliciesAPI.UpdateRetentionPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRetentionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRetentionPolicyRequest** | [**CreateRetentionPolicyRequest**](CreateRetentionPolicyRequest.md) |  | 

### Return type

[**CreateRetentionPolicy200Response**](CreateRetentionPolicy200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

