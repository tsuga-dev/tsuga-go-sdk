# \PromqlAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryPromql**](PromqlAPI.md#QueryPromql) | **Post** /v1/promql | 
[**QueryPromqlV2**](PromqlAPI.md#QueryPromqlV2) | **Post** /v2/promql | 



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
	queryPromqlRequest := *openapiclient.NewQueryPromqlRequest("Query_example", "Step_example", *openapiclient.NewInputTimeRange(int64(123), int64(123))) // QueryPromqlRequest | PromQL range-query request. Send clusterId as a query parameter, not in this JSON body.
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


## QueryPromqlV2

> QueryPromqlV2Response QueryPromqlV2(ctx).QueryPromqlV2Request(queryPromqlV2Request).ClusterId(clusterId).Execute()





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
	queryPromqlV2Request := *openapiclient.NewQueryPromqlV2Request("Query_example", "Step_example", *openapiclient.NewInputV2TimeRange(int64(123), int64(123))) // QueryPromqlV2Request | PromQL range-query request with the time range in milliseconds. Send clusterId as a query parameter, not in this JSON body.
	clusterId := "clusterId_example" // string | Telemetry cluster to query. Set by the client as the clusterId query parameter, not in the JSON body. Required when the organization has multiple clusters; omit only for single-cluster organizations. Unknown cluster IDs return 404. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromqlAPI.QueryPromqlV2(context.Background()).QueryPromqlV2Request(queryPromqlV2Request).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromqlAPI.QueryPromqlV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryPromqlV2`: QueryPromqlV2Response
	fmt.Fprintf(os.Stdout, "Response from `PromqlAPI.QueryPromqlV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryPromqlV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryPromqlV2Request** | [**QueryPromqlV2Request**](QueryPromqlV2Request.md) | PromQL range-query request with the time range in milliseconds. Send clusterId as a query parameter, not in this JSON body. | 
 **clusterId** | **string** | Telemetry cluster to query. Set by the client as the clusterId query parameter, not in the JSON body. Required when the organization has multiple clusters; omit only for single-cluster organizations. Unknown cluster IDs return 404. | 

### Return type

[**QueryPromqlV2Response**](QueryPromqlV2Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

