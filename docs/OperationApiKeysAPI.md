# \OperationApiKeysAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOperationApiKey**](OperationApiKeysAPI.md#CreateOperationApiKey) | **Post** /v1/operation-api-keys | 
[**DeleteOperationApiKey**](OperationApiKeysAPI.md#DeleteOperationApiKey) | **Delete** /v1/operation-api-keys/{id} | 
[**GetOperationApiKey**](OperationApiKeysAPI.md#GetOperationApiKey) | **Get** /v1/operation-api-keys/{id} | 
[**ListOperationApiKeys**](OperationApiKeysAPI.md#ListOperationApiKeys) | **Get** /v1/operation-api-keys | 
[**UpdateOperationApiKey**](OperationApiKeysAPI.md#UpdateOperationApiKey) | **Put** /v1/operation-api-keys/{id} | 



## CreateOperationApiKey

> CreateOperationApiKeyResponse CreateOperationApiKey(ctx).CreateOperationApiKeyRequest(createOperationApiKeyRequest).Execute()





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
	createOperationApiKeyRequest := *openapiclient.NewCreateOperationApiKeyRequest("Name_example", "Owner_example", "Scope_example", *openapiclient.NewCreateOperationApiKeyRequestPermissions()) // CreateOperationApiKeyRequest | Operation API key creation request. Provide a kebab-case name, owner team, scope, and permissions, and optionally tags and an expiration date.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationApiKeysAPI.CreateOperationApiKey(context.Background()).CreateOperationApiKeyRequest(createOperationApiKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationApiKeysAPI.CreateOperationApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOperationApiKey`: CreateOperationApiKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationApiKeysAPI.CreateOperationApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOperationApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOperationApiKeyRequest** | [**CreateOperationApiKeyRequest**](CreateOperationApiKeyRequest.md) | Operation API key creation request. Provide a kebab-case name, owner team, scope, and permissions, and optionally tags and an expiration date. | 

### Return type

[**CreateOperationApiKeyResponse**](CreateOperationApiKeyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOperationApiKey

> DeleteOperationApiKeyResponse DeleteOperationApiKey(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the operation API key to delete. Use the `id` returned by operation API key list, get, or create responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationApiKeysAPI.DeleteOperationApiKey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationApiKeysAPI.DeleteOperationApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOperationApiKey`: DeleteOperationApiKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationApiKeysAPI.DeleteOperationApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the operation API key to delete. Use the &#x60;id&#x60; returned by operation API key list, get, or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOperationApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteOperationApiKeyResponse**](DeleteOperationApiKeyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOperationApiKey

> GetOperationApiKeyResponse GetOperationApiKey(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the operation API key to retrieve. Use the `id` returned by operation API key list or create responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationApiKeysAPI.GetOperationApiKey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationApiKeysAPI.GetOperationApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOperationApiKey`: GetOperationApiKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationApiKeysAPI.GetOperationApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the operation API key to retrieve. Use the &#x60;id&#x60; returned by operation API key list or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOperationApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOperationApiKeyResponse**](GetOperationApiKeyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperationApiKeys

> ListOperationApiKeysResponse ListOperationApiKeys(ctx).Limit(limit).Offset(offset).Execute()





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
	resp, r, err := apiClient.OperationApiKeysAPI.ListOperationApiKeys(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationApiKeysAPI.ListOperationApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOperationApiKeys`: ListOperationApiKeysResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationApiKeysAPI.ListOperationApiKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOperationApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [default to 100]
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 

### Return type

[**ListOperationApiKeysResponse**](ListOperationApiKeysResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOperationApiKey

> UpdateOperationApiKeyResponse UpdateOperationApiKey(ctx, id).UpdateOperationApiKeyRequest(updateOperationApiKeyRequest).Execute()





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
	id := "id_example" // string | Identifier of the operation API key to update. Use the `id` returned by operation API key list, get, or create responses.
	updateOperationApiKeyRequest := *openapiclient.NewUpdateOperationApiKeyRequest("Name_example", "Owner_example", "Scope_example", *openapiclient.NewCreateOperationApiKeyRequestPermissions()) // UpdateOperationApiKeyRequest | Operation API key update request. Provide the full definition, including name, owner team, scope, and permissions.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationApiKeysAPI.UpdateOperationApiKey(context.Background(), id).UpdateOperationApiKeyRequest(updateOperationApiKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationApiKeysAPI.UpdateOperationApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOperationApiKey`: UpdateOperationApiKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `OperationApiKeysAPI.UpdateOperationApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the operation API key to update. Use the &#x60;id&#x60; returned by operation API key list, get, or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOperationApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOperationApiKeyRequest** | [**UpdateOperationApiKeyRequest**](UpdateOperationApiKeyRequest.md) | Operation API key update request. Provide the full definition, including name, owner team, scope, and permissions. | 

### Return type

[**UpdateOperationApiKeyResponse**](UpdateOperationApiKeyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

