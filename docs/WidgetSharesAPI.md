# \WidgetSharesAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWidgetShare**](WidgetSharesAPI.md#GetWidgetShare) | **Get** /v1/widget-shares/{id} | 



## GetWidgetShare

> GetWidgetShareResponse GetWidgetShare(ctx, id).Execute()





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
	id := "id_example" // string | Code of the share link, the last path segment of `https://api.tsuga.com/v0/widget-shares/<orgId>/<id>`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WidgetSharesAPI.GetWidgetShare(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WidgetSharesAPI.GetWidgetShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWidgetShare`: GetWidgetShareResponse
	fmt.Fprintf(os.Stdout, "Response from `WidgetSharesAPI.GetWidgetShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Code of the share link, the last path segment of &#x60;https://api.tsuga.com/v0/widget-shares/&lt;orgId&gt;/&lt;id&gt;&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWidgetShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWidgetShareResponse**](GetWidgetShareResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

