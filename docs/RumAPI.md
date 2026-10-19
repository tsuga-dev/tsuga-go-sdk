# \RumAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRumAttributes**](RumAPI.md#GetRumAttributes) | **Get** /v1/rum/attributes | 
[**SearchRumEvents**](RumAPI.md#SearchRumEvents) | **Get** /v1/rum/search | 



## GetRumAttributes

> GetRumAttributesResponse GetRumAttributes(ctx).ClusterId(clusterId).Execute()





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
	clusterId := "clusterId_example" // string | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RumAPI.GetRumAttributes(context.Background()).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumAPI.GetRumAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRumAttributes`: GetRumAttributesResponse
	fmt.Fprintf(os.Stdout, "Response from `RumAPI.GetRumAttributes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRumAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. | 

### Return type

[**GetRumAttributesResponse**](GetRumAttributesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRumEvents

> SearchRumEventsResponse SearchRumEvents(ctx).From(from).To(to).Query(query).MaxResults(maxResults).ClusterId(clusterId).Execute()





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
	from := int64(1704067200) // int64 | Start of the query window as a Unix timestamp in seconds
	to := int64(1704067200) // int64 | End of the query window as a Unix timestamp in seconds
	query := "query_example" // string | Tsuga query selecting RUM events to return. Prefix every field with the stream it belongs to (`measurements.`, `events.`, or `exceptions.`), e.g. `events.view.name:/checkout`, and filter on a single stream per query. A query without stream-prefixed fields, including the default `*`, searches the `events` stream. (optional) (default to "*")
	maxResults := int32(56) // int32 | Maximum number of RUM events to return. (optional) (default to 100)
	clusterId := "clusterId_example" // string | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RumAPI.SearchRumEvents(context.Background()).From(from).To(to).Query(query).MaxResults(maxResults).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumAPI.SearchRumEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchRumEvents`: SearchRumEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `RumAPI.SearchRumEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRumEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | Start of the query window as a Unix timestamp in seconds | 
 **to** | **int64** | End of the query window as a Unix timestamp in seconds | 
 **query** | **string** | Tsuga query selecting RUM events to return. Prefix every field with the stream it belongs to (&#x60;measurements.&#x60;, &#x60;events.&#x60;, or &#x60;exceptions.&#x60;), e.g. &#x60;events.view.name:/checkout&#x60;, and filter on a single stream per query. A query without stream-prefixed fields, including the default &#x60;*&#x60;, searches the &#x60;events&#x60; stream. | [default to &quot;*&quot;]
 **maxResults** | **int32** | Maximum number of RUM events to return. | [default to 100]
 **clusterId** | **string** | Telemetry cluster to target. Required when the organization has multiple clusters; omit only for single-cluster organizations. | 

### Return type

[**SearchRumEventsResponse**](SearchRumEventsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

