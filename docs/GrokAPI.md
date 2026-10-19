# \GrokAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ParseGrok**](GrokAPI.md#ParseGrok) | **Post** /v1/grok/parse | 



## ParseGrok

> ParseGrokResponse ParseGrok(ctx).ParseGrokRequest(parseGrokRequest).Execute()





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
	parseGrokRequest := *openapiclient.NewParseGrokRequest([]string{"Rules_example"}, []string{"Samples_example"}) // ParseGrokRequest | Grok rules and sample log lines to preview extraction before saving a route processor. The endpoint uses the same parser as the route builder preview and does not create or update routes.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GrokAPI.ParseGrok(context.Background()).ParseGrokRequest(parseGrokRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GrokAPI.ParseGrok``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParseGrok`: ParseGrokResponse
	fmt.Fprintf(os.Stdout, "Response from `GrokAPI.ParseGrok`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiParseGrokRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parseGrokRequest** | [**ParseGrokRequest**](ParseGrokRequest.md) | Grok rules and sample log lines to preview extraction before saving a route processor. The endpoint uses the same parser as the route builder preview and does not create or update routes. | 

### Return type

[**ParseGrokResponse**](ParseGrokResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

