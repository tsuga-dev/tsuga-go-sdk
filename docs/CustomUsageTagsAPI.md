# \CustomUsageTagsAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomUsageTag**](CustomUsageTagsAPI.md#CreateCustomUsageTag) | **Post** /v1/custom-usage-tags | 
[**DeleteCustomUsageTag**](CustomUsageTagsAPI.md#DeleteCustomUsageTag) | **Delete** /v1/custom-usage-tags/{id} | 
[**GetCustomUsageTag**](CustomUsageTagsAPI.md#GetCustomUsageTag) | **Get** /v1/custom-usage-tags/{id} | 
[**ListCustomUsageTags**](CustomUsageTagsAPI.md#ListCustomUsageTags) | **Get** /v1/custom-usage-tags | 



## CreateCustomUsageTag

> CreateCustomUsageTagResponse CreateCustomUsageTag(ctx).CreateCustomUsageTagRequest(createCustomUsageTagRequest).Execute()





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
	createCustomUsageTagRequest := *openapiclient.NewCreateCustomUsageTagRequest("TagKey_example") // CreateCustomUsageTagRequest | Custom usage tag creation request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomUsageTagsAPI.CreateCustomUsageTag(context.Background()).CreateCustomUsageTagRequest(createCustomUsageTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomUsageTagsAPI.CreateCustomUsageTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomUsageTag`: CreateCustomUsageTagResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomUsageTagsAPI.CreateCustomUsageTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomUsageTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomUsageTagRequest** | [**CreateCustomUsageTagRequest**](CreateCustomUsageTagRequest.md) | Custom usage tag creation request. | 

### Return type

[**CreateCustomUsageTagResponse**](CreateCustomUsageTagResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomUsageTag

> DeleteCustomUsageTagResponse DeleteCustomUsageTag(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the custom usage tag configuration to delete. Use the `id` returned by custom usage tag list, get, or create responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomUsageTagsAPI.DeleteCustomUsageTag(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomUsageTagsAPI.DeleteCustomUsageTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCustomUsageTag`: DeleteCustomUsageTagResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomUsageTagsAPI.DeleteCustomUsageTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the custom usage tag configuration to delete. Use the &#x60;id&#x60; returned by custom usage tag list, get, or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomUsageTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteCustomUsageTagResponse**](DeleteCustomUsageTagResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomUsageTag

> GetCustomUsageTagResponse GetCustomUsageTag(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the custom usage tag configuration to retrieve. Use the `id` returned by custom usage tag list or create responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomUsageTagsAPI.GetCustomUsageTag(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomUsageTagsAPI.GetCustomUsageTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomUsageTag`: GetCustomUsageTagResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomUsageTagsAPI.GetCustomUsageTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the custom usage tag configuration to retrieve. Use the &#x60;id&#x60; returned by custom usage tag list or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomUsageTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCustomUsageTagResponse**](GetCustomUsageTagResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomUsageTags

> ListCustomUsageTagsResponse ListCustomUsageTags(ctx).Execute()





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
	resp, r, err := apiClient.CustomUsageTagsAPI.ListCustomUsageTags(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomUsageTagsAPI.ListCustomUsageTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomUsageTags`: ListCustomUsageTagsResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomUsageTagsAPI.ListCustomUsageTags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomUsageTagsRequest struct via the builder pattern


### Return type

[**ListCustomUsageTagsResponse**](ListCustomUsageTagsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

