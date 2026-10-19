# \PublicTokensAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePublicToken**](PublicTokensAPI.md#CreatePublicToken) | **Post** /v1/public-tokens | 
[**DeletePublicToken**](PublicTokensAPI.md#DeletePublicToken) | **Delete** /v1/public-tokens/{id} | 
[**GetPublicToken**](PublicTokensAPI.md#GetPublicToken) | **Get** /v1/public-tokens/{id} | 
[**ListPublicTokens**](PublicTokensAPI.md#ListPublicTokens) | **Get** /v1/public-tokens | 
[**UpdatePublicToken**](PublicTokensAPI.md#UpdatePublicToken) | **Put** /v1/public-tokens/{id} | 



## CreatePublicToken

> CreatePublicTokenResponse CreatePublicToken(ctx).CreatePublicTokenRequest(createPublicTokenRequest).Execute()





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
	createPublicTokenRequest := *openapiclient.NewCreatePublicTokenRequest("Name_example", "Owner_example", "ApplicationType_example", []string{"AllowedUrlPrefixes_example"}, []openapiclient.InputTag{*openapiclient.NewInputTag("Key_example", "Value_example")}, []string{"TeamOverrideFields_example"}) // CreatePublicTokenRequest | RUM public token creation request. Provide a kebab-case name, owner team, application type, allowed URL prefixes (required for `web`, empty for `ios` and `android`), tags including an `env` tag, and team override fields as an array of fields or null.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicTokensAPI.CreatePublicToken(context.Background()).CreatePublicTokenRequest(createPublicTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicTokensAPI.CreatePublicToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePublicToken`: CreatePublicTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicTokensAPI.CreatePublicToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePublicTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPublicTokenRequest** | [**CreatePublicTokenRequest**](CreatePublicTokenRequest.md) | RUM public token creation request. Provide a kebab-case name, owner team, application type, allowed URL prefixes (required for &#x60;web&#x60;, empty for &#x60;ios&#x60; and &#x60;android&#x60;), tags including an &#x60;env&#x60; tag, and team override fields as an array of fields or null. | 

### Return type

[**CreatePublicTokenResponse**](CreatePublicTokenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePublicToken

> DeletePublicTokenResponse DeletePublicToken(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the public token to delete. Use the `id` returned by public token list, get, or create responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicTokensAPI.DeletePublicToken(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicTokensAPI.DeletePublicToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePublicToken`: DeletePublicTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicTokensAPI.DeletePublicToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the public token to delete. Use the &#x60;id&#x60; returned by public token list, get, or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePublicTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeletePublicTokenResponse**](DeletePublicTokenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicToken

> GetPublicTokenResponse GetPublicToken(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the public token to retrieve. Use the `id` returned by public token list or create responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicTokensAPI.GetPublicToken(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicTokensAPI.GetPublicToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicToken`: GetPublicTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicTokensAPI.GetPublicToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the public token to retrieve. Use the &#x60;id&#x60; returned by public token list or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPublicTokenResponse**](GetPublicTokenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPublicTokens

> ListPublicTokensResponse ListPublicTokens(ctx).Limit(limit).Offset(offset).Execute()





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
	resp, r, err := apiClient.PublicTokensAPI.ListPublicTokens(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicTokensAPI.ListPublicTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPublicTokens`: ListPublicTokensResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicTokensAPI.ListPublicTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPublicTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [default to 100]
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 

### Return type

[**ListPublicTokensResponse**](ListPublicTokensResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePublicToken

> UpdatePublicTokenResponse UpdatePublicToken(ctx, id).UpdatePublicTokenRequest(updatePublicTokenRequest).Execute()





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
	id := "id_example" // string | Identifier of the public token to update. Use the `id` returned by public token list, get, or create responses.
	updatePublicTokenRequest := *openapiclient.NewUpdatePublicTokenRequest("Name_example", "Owner_example", "ApplicationType_example", []string{"AllowedUrlPrefixes_example"}, []openapiclient.InputTag{*openapiclient.NewInputTag("Key_example", "Value_example")}, []string{"TeamOverrideFields_example"}) // UpdatePublicTokenRequest | RUM public token update request. Provide the full definition, including name, owner team, allowed URL prefixes, tags including an `env` tag, and team override fields. The application type must match the existing token and the token value itself never changes.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicTokensAPI.UpdatePublicToken(context.Background(), id).UpdatePublicTokenRequest(updatePublicTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicTokensAPI.UpdatePublicToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePublicToken`: UpdatePublicTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicTokensAPI.UpdatePublicToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the public token to update. Use the &#x60;id&#x60; returned by public token list, get, or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePublicTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePublicTokenRequest** | [**UpdatePublicTokenRequest**](UpdatePublicTokenRequest.md) | RUM public token update request. Provide the full definition, including name, owner team, allowed URL prefixes, tags including an &#x60;env&#x60; tag, and team override fields. The application type must match the existing token and the token value itself never changes. | 

### Return type

[**UpdatePublicTokenResponse**](UpdatePublicTokenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

