# \PromqlAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryPromql**](PromqlAPI.md#QueryPromql) | **Post** /v1/promql | 



## QueryPromql

> QueryPromqlResponse QueryPromql(ctx).QueryPromqlRequest(queryPromqlRequest).ClusterId(clusterId).Execute()





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
	queryPromqlRequest := *openapiclient.NewQueryPromqlRequest("Query_example", "Step_example", *openapiclient.NewInputTimeRange(int64(1704067200), int64(1704067200))) // QueryPromqlRequest | PromQL range-query request. Send clusterId as a query parameter, not in this JSON body.
	clusterId := "clusterId_example" // string | Telemetry cluster to query. Set by the client as the clusterId query parameter, not in the JSON body. Required when the organization has multiple clusters; omit only for single-cluster organizations. Unknown cluster IDs return 404. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromqlAPI.QueryPromql(context.Background()).QueryPromqlRequest(queryPromqlRequest).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromqlAPI.QueryPromql``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryPromql`: QueryPromqlResponse
	fmt.Fprintf(os.Stdout, "Response from `PromqlAPI.QueryPromql`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryPromqlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryPromqlRequest** | [**QueryPromqlRequest**](QueryPromqlRequest.md) | PromQL range-query request. Send clusterId as a query parameter, not in this JSON body. | 
 **clusterId** | **string** | Telemetry cluster to query. Set by the client as the clusterId query parameter, not in the JSON body. Required when the organization has multiple clusters; omit only for single-cluster organizations. Unknown cluster IDs return 404. | 

### Return type

[**QueryPromqlResponse**](QueryPromqlResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

