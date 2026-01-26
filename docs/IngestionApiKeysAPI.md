# \IngestionApiKeysAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIngestionApiKey**](IngestionApiKeysAPI.md#CreateIngestionApiKey) | **Post** /v1/ingestion-api-keys | 
[**DeleteIngestionApiKey**](IngestionApiKeysAPI.md#DeleteIngestionApiKey) | **Delete** /v1/ingestion-api-keys/{id} | 
[**ListIngestionApiKeys**](IngestionApiKeysAPI.md#ListIngestionApiKeys) | **Get** /v1/ingestion-api-keys | 
[**UpdateIngestionApiKey**](IngestionApiKeysAPI.md#UpdateIngestionApiKey) | **Put** /v1/ingestion-api-keys/{id} | 



## CreateIngestionApiKey

> CreateIngestionApiKey200Response CreateIngestionApiKey(ctx).CreateIngestionApiKeyRequest(createIngestionApiKeyRequest).Execute()





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
	createIngestionApiKeyRequest := *openapiclient.NewCreateIngestionApiKeyRequest("Name_example", "Owner_example") // CreateIngestionApiKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionApiKeysAPI.CreateIngestionApiKey(context.Background()).CreateIngestionApiKeyRequest(createIngestionApiKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionApiKeysAPI.CreateIngestionApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIngestionApiKey`: CreateIngestionApiKey200Response
	fmt.Fprintf(os.Stdout, "Response from `IngestionApiKeysAPI.CreateIngestionApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIngestionApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIngestionApiKeyRequest** | [**CreateIngestionApiKeyRequest**](CreateIngestionApiKeyRequest.md) |  | 

### Return type

[**CreateIngestionApiKey200Response**](CreateIngestionApiKey200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIngestionApiKey

> DeleteIngestionApiKey200Response DeleteIngestionApiKey(ctx, id).Execute()





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
	resp, r, err := apiClient.IngestionApiKeysAPI.DeleteIngestionApiKey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionApiKeysAPI.DeleteIngestionApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIngestionApiKey`: DeleteIngestionApiKey200Response
	fmt.Fprintf(os.Stdout, "Response from `IngestionApiKeysAPI.DeleteIngestionApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIngestionApiKeyRequest struct via the builder pattern


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


## ListIngestionApiKeys

> ListIngestionApiKeys200Response ListIngestionApiKeys(ctx).Execute()





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
	resp, r, err := apiClient.IngestionApiKeysAPI.ListIngestionApiKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionApiKeysAPI.ListIngestionApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIngestionApiKeys`: ListIngestionApiKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `IngestionApiKeysAPI.ListIngestionApiKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIngestionApiKeysRequest struct via the builder pattern


### Return type

[**ListIngestionApiKeys200Response**](ListIngestionApiKeys200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIngestionApiKey

> UpdateIngestionApiKey200Response UpdateIngestionApiKey(ctx, id).CreateIngestionApiKeyRequest(createIngestionApiKeyRequest).Execute()





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
	createIngestionApiKeyRequest := *openapiclient.NewCreateIngestionApiKeyRequest("Name_example", "Owner_example") // CreateIngestionApiKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngestionApiKeysAPI.UpdateIngestionApiKey(context.Background(), id).CreateIngestionApiKeyRequest(createIngestionApiKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngestionApiKeysAPI.UpdateIngestionApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIngestionApiKey`: UpdateIngestionApiKey200Response
	fmt.Fprintf(os.Stdout, "Response from `IngestionApiKeysAPI.UpdateIngestionApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIngestionApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createIngestionApiKeyRequest** | [**CreateIngestionApiKeyRequest**](CreateIngestionApiKeyRequest.md) |  | 

### Return type

[**UpdateIngestionApiKey200Response**](UpdateIngestionApiKey200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

