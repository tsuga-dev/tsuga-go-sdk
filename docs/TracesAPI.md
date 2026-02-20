# \TracesAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchSpans**](TracesAPI.md#SearchSpans) | **Get** /v1/traces/search | 



## SearchSpans

> SearchSpans200Response SearchSpans(ctx).From(from).To(to).Query(query).MaxResults(maxResults).Execute()





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
	from := int32(1704067200) // int32 | Start timestamp in seconds
	to := int32(1704067200) // int32 | End timestamp in seconds
	query := "query_example" // string |  (optional)
	maxResults := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracesAPI.SearchSpans(context.Background()).From(from).To(to).Query(query).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.SearchSpans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSpans`: SearchSpans200Response
	fmt.Fprintf(os.Stdout, "Response from `TracesAPI.SearchSpans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSpansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int32** | Start timestamp in seconds | 
 **to** | **int32** | End timestamp in seconds | 
 **query** | **string** |  | 
 **maxResults** | **int32** |  | 

### Return type

[**SearchSpans200Response**](SearchSpans200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

