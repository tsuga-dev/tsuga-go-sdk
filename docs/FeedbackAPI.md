# \FeedbackAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFeedback**](FeedbackAPI.md#CreateFeedback) | **Post** /v1/feedback | 



## CreateFeedback

> CreateFeedbackResponse CreateFeedback(ctx).CreateFeedbackRequest(createFeedbackRequest).Execute()





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
	createFeedbackRequest := *openapiclient.NewCreateFeedbackRequest("Content_example") // CreateFeedbackRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedbackAPI.CreateFeedback(context.Background()).CreateFeedbackRequest(createFeedbackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedbackAPI.CreateFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFeedback`: CreateFeedbackResponse
	fmt.Fprintf(os.Stdout, "Response from `FeedbackAPI.CreateFeedback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFeedbackRequest** | [**CreateFeedbackRequest**](CreateFeedbackRequest.md) |  | 

### Return type

[**CreateFeedbackResponse**](CreateFeedbackResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

