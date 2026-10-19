# \InterestingFieldsAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInterestingFields**](InterestingFieldsAPI.md#GetInterestingFields) | **Post** /v1/interesting-fields | 



## GetInterestingFields

> GetInterestingFieldsResponse GetInterestingFields(ctx).GetInterestingFieldsRequest(getInterestingFieldsRequest).Execute()





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
	getInterestingFieldsRequest := *openapiclient.NewGetInterestingFieldsRequest([]openapiclient.GetInterestingFieldsRequestScopeInner{*openapiclient.NewGetInterestingFieldsRequestScopeInner("Key_example")}) // GetInterestingFieldsRequest | The scope to rank fields for, as attribute key/value pairs such as a service name and an environment. Entries are matched softly by default: a widget contributes in proportion to how many entries it matches, so a broader scope returns a broader ranking. Set `valueMustMatch` on an entry to make it a hard filter instead: a widget that does not match it exactly contributes nothing, even when other entries match.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InterestingFieldsAPI.GetInterestingFields(context.Background()).GetInterestingFieldsRequest(getInterestingFieldsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterestingFieldsAPI.GetInterestingFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInterestingFields`: GetInterestingFieldsResponse
	fmt.Fprintf(os.Stdout, "Response from `InterestingFieldsAPI.GetInterestingFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInterestingFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getInterestingFieldsRequest** | [**GetInterestingFieldsRequest**](GetInterestingFieldsRequest.md) | The scope to rank fields for, as attribute key/value pairs such as a service name and an environment. Entries are matched softly by default: a widget contributes in proportion to how many entries it matches, so a broader scope returns a broader ranking. Set &#x60;valueMustMatch&#x60; on an entry to make it a hard filter instead: a widget that does not match it exactly contributes nothing, even when other entries match. | 

### Return type

[**GetInterestingFieldsResponse**](GetInterestingFieldsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

