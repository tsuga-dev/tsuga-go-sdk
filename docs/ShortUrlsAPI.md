# \ShortUrlsAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetShortUrl**](ShortUrlsAPI.md#GetShortUrl) | **Get** /v1/short-urls/{id} | 



## GetShortUrl

> GetShortUrlResponse GetShortUrl(ctx, id).Execute()





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
	id := "id_example" // string | Code of the short link, the last path segment of `https://app.tsuga.com/l/<id>`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShortUrlsAPI.GetShortUrl(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShortUrlsAPI.GetShortUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShortUrl`: GetShortUrlResponse
	fmt.Fprintf(os.Stdout, "Response from `ShortUrlsAPI.GetShortUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Code of the short link, the last path segment of &#x60;https://app.tsuga.com/l/&lt;id&gt;&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShortUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetShortUrlResponse**](GetShortUrlResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

