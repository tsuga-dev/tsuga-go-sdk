# \DocsAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDocPage**](DocsAPI.md#GetDocPage) | **Get** /v1/docs | 
[**SearchDocs**](DocsAPI.md#SearchDocs) | **Get** /v1/docs/search | 



## GetDocPage

> GetDocPageResponse GetDocPage(ctx).Path(path).Execute()





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
	path := "path_example" // string | Documentation path to fetch. Use page slug paths for docs or `api/<operationId>` for API-reference pages.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocsAPI.GetDocPage(context.Background()).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocsAPI.GetDocPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDocPage`: GetDocPageResponse
	fmt.Fprintf(os.Stdout, "Response from `DocsAPI.GetDocPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDocPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Documentation path to fetch. Use page slug paths for docs or &#x60;api/&lt;operationId&gt;&#x60; for API-reference pages. | 

### Return type

[**GetDocPageResponse**](GetDocPageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDocs

> SearchDocsResponse SearchDocs(ctx).Query(query).Execute()





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
	query := "query_example" // string | Search text to match against documentation and API-reference content.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocsAPI.SearchDocs(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocsAPI.SearchDocs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDocs`: SearchDocsResponse
	fmt.Fprintf(os.Stdout, "Response from `DocsAPI.SearchDocs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchDocsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Search text to match against documentation and API-reference content. | 

### Return type

[**SearchDocsResponse**](SearchDocsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

