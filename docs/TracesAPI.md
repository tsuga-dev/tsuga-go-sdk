# \TracesAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchSpans**](TracesAPI.md#SearchSpans) | **Get** /v1/traces/search | 



## SearchSpans

> SearchSpansResponse SearchSpans(ctx).From(from).To(to).Query(query).MaxResults(maxResults).Execute()





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
	from := int64(1704067200) // int64 | Start timestamp in seconds
	to := int64(1704067200) // int64 | End timestamp in seconds
	query := "query_example" // string |  (optional)
	maxResults := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracesAPI.SearchSpans(context.Background()).From(from).To(to).Query(query).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.SearchSpans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSpans`: SearchSpansResponse
	fmt.Fprintf(os.Stdout, "Response from `TracesAPI.SearchSpans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSpansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | Start timestamp in seconds | 
 **to** | **int64** | End timestamp in seconds | 
 **query** | **string** |  | 
 **maxResults** | **int32** |  | 

### Return type

[**SearchSpansResponse**](SearchSpansResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

