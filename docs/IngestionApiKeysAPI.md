# \IngestionApiKeysAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIngestionApiKey**](IngestionApiKeysAPI.md#CreateIngestionApiKey) | **Post** /v1/ingestion-api-keys | 
[**DeleteIngestionApiKey**](IngestionApiKeysAPI.md#DeleteIngestionApiKey) | **Delete** /v1/ingestion-api-keys/{id} | 
[**GetIngestionApiKey**](IngestionApiKeysAPI.md#GetIngestionApiKey) | **Get** /v1/ingestion-api-keys/{id} | 
[**ListIngestionApiKeys**](IngestionApiKeysAPI.md#ListIngestionApiKeys) | **Get** /v1/ingestion-api-keys | 
[**UpdateIngestionApiKey**](IngestionApiKeysAPI.md#UpdateIngestionApiKey) | **Put** /v1/ingestion-api-keys/{id} | 



## CreateIngestionApiKey

> CreateIngestionApiKeyResponse CreateIngestionApiKey(ctx).CreateIngestionApiKeyRequest(createIngestionApiKeyRequest).Execute()





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
	createIngestionApiKeyRequest := *openapiclient.NewCreateIngestionApiKeyRequest("Name_example", "Owner_example", []openapiclient.InputTag{*openapiclient.NewInputTag("Key_example", "Value_example")}, []*string{nil}) // CreateIngestionApiKeyRequest | Ingestion API key creation request. Provide a kebab-case name, owner team, required tags, and required team override fields as an array of fields or null.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionApiKeysAPI.CreateIngestionApiKey(context.Background()).CreateIngestionApiKeyRequest(createIngestionApiKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionApiKeysAPI.CreateIngestionApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIngestionApiKey`: CreateIngestionApiKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `IngestionApiKeysAPI.CreateIngestionApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIngestionApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIngestionApiKeyRequest** | [**CreateIngestionApiKeyRequest**](CreateIngestionApiKeyRequest.md) | Ingestion API key creation request. Provide a kebab-case name, owner team, required tags, and required team override fields as an array of fields or null. | 

### Return type

[**CreateIngestionApiKeyResponse**](CreateIngestionApiKeyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIngestionApiKey

> DeleteIngestionApiKeyResponse DeleteIngestionApiKey(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the ingestion API key to delete. Use the `id` returned by ingestion API key list, create, or update responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionApiKeysAPI.DeleteIngestionApiKey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionApiKeysAPI.DeleteIngestionApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIngestionApiKey`: DeleteIngestionApiKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `IngestionApiKeysAPI.DeleteIngestionApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the ingestion API key to delete. Use the &#x60;id&#x60; returned by ingestion API key list, create, or update responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIngestionApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteIngestionApiKeyResponse**](DeleteIngestionApiKeyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIngestionApiKey

> GetIngestionApiKeyResponse GetIngestionApiKey(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the ingestion API key to retrieve. Use the `id` returned by ingestion API key list or create responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionApiKeysAPI.GetIngestionApiKey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionApiKeysAPI.GetIngestionApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIngestionApiKey`: GetIngestionApiKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `IngestionApiKeysAPI.GetIngestionApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the ingestion API key to retrieve. Use the &#x60;id&#x60; returned by ingestion API key list or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIngestionApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetIngestionApiKeyResponse**](GetIngestionApiKeyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIngestionApiKeys

> ListIngestionApiKeysResponse ListIngestionApiKeys(ctx).Limit(limit).Offset(offset).Execute()





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
	resp, r, err := apiClient.IngestionApiKeysAPI.ListIngestionApiKeys(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionApiKeysAPI.ListIngestionApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIngestionApiKeys`: ListIngestionApiKeysResponse
	fmt.Fprintf(os.Stdout, "Response from `IngestionApiKeysAPI.ListIngestionApiKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIngestionApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [default to 100]
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 

### Return type

[**ListIngestionApiKeysResponse**](ListIngestionApiKeysResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIngestionApiKey

> UpdateIngestionApiKeyResponse UpdateIngestionApiKey(ctx, id).UpdateIngestionApiKeyRequest(updateIngestionApiKeyRequest).Execute()





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
	id := "id_example" // string | Identifier of the ingestion API key to update. Use the `id` returned by ingestion API key list, get, or create responses.
	updateIngestionApiKeyRequest := *openapiclient.NewUpdateIngestionApiKeyRequest("Name_example", "Owner_example", []openapiclient.InputTag{*openapiclient.NewInputTag("Key_example", "Value_example")}, []string{"TeamOverrideFields_example"}) // UpdateIngestionApiKeyRequest | Ingestion API key update request. Provide the full metadata definition, including name, owner team, tags, and team override fields.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionApiKeysAPI.UpdateIngestionApiKey(context.Background(), id).UpdateIngestionApiKeyRequest(updateIngestionApiKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionApiKeysAPI.UpdateIngestionApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIngestionApiKey`: UpdateIngestionApiKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `IngestionApiKeysAPI.UpdateIngestionApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the ingestion API key to update. Use the &#x60;id&#x60; returned by ingestion API key list, get, or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIngestionApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIngestionApiKeyRequest** | [**UpdateIngestionApiKeyRequest**](UpdateIngestionApiKeyRequest.md) | Ingestion API key update request. Provide the full metadata definition, including name, owner team, tags, and team override fields. | 

### Return type

[**UpdateIngestionApiKeyResponse**](UpdateIngestionApiKeyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

