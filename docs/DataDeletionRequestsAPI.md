# \DataDeletionRequestsAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelDataDeletionRequest**](DataDeletionRequestsAPI.md#CancelDataDeletionRequest) | **Post** /v1/data-deletion-requests/{id}/cancel | 
[**CreateDataDeletionRequest**](DataDeletionRequestsAPI.md#CreateDataDeletionRequest) | **Post** /v1/data-deletion-requests | 
[**ListDataDeletionRequests**](DataDeletionRequestsAPI.md#ListDataDeletionRequests) | **Get** /v1/data-deletion-requests | 



## CancelDataDeletionRequest

> CancelDataDeletionRequestResponse CancelDataDeletionRequest(ctx, id).Execute()





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
	id := "id_example" // string | ID of the deletion request to cancel.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataDeletionRequestsAPI.CancelDataDeletionRequest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataDeletionRequestsAPI.CancelDataDeletionRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelDataDeletionRequest`: CancelDataDeletionRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `DataDeletionRequestsAPI.CancelDataDeletionRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the deletion request to cancel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelDataDeletionRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CancelDataDeletionRequestResponse**](CancelDataDeletionRequestResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDataDeletionRequest

> CreateDataDeletionRequestResponse CreateDataDeletionRequest(ctx).CreateDataDeletionRequestRequest(createDataDeletionRequestRequest).Execute()





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
	createDataDeletionRequestRequest := *openapiclient.NewCreateDataDeletionRequestRequest("ClusterId_example", "DataSource_example", "Query_example", int64(123), int64(123), "Reason_example", false, "Owner_example") // CreateDataDeletionRequestRequest | Data deletion request to create: the cluster and telemetry signal to delete from, the query and time range selecting the events, the reason kept for audit, and the owning team.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataDeletionRequestsAPI.CreateDataDeletionRequest(context.Background()).CreateDataDeletionRequestRequest(createDataDeletionRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataDeletionRequestsAPI.CreateDataDeletionRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDataDeletionRequest`: CreateDataDeletionRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `DataDeletionRequestsAPI.CreateDataDeletionRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataDeletionRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDataDeletionRequestRequest** | [**CreateDataDeletionRequestRequest**](CreateDataDeletionRequestRequest.md) | Data deletion request to create: the cluster and telemetry signal to delete from, the query and time range selecting the events, the reason kept for audit, and the owning team. | 

### Return type

[**CreateDataDeletionRequestResponse**](CreateDataDeletionRequestResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataDeletionRequests

> ListDataDeletionRequestsResponse ListDataDeletionRequests(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).DataSource(dataSource).SortBy(sortBy).SortDirection(sortDirection).Execute()





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
	limit := int32(56) // int32 | Maximum number of items to return in this page. Omit to use the public API default of 100. (optional) (default to 100)
	offset := int32(56) // int32 | Zero-based index of the first matching item to return. Increase it with `limit` to request later pages. If `limit` is provided without `offset`, the offset defaults to 0. (optional)
	clusterId := "clusterId_example" // string | Set to only return the requests of this cluster. (optional)
	dataSource := "dataSource_example" // string | Set to only return the requests deleting from this telemetry signal. (optional)
	sortBy := "sortBy_example" // string | Field to sort the requests by. Omit to sort by creation time, newest first. (optional)
	sortDirection := "sortDirection_example" // string | Sort direction: ascending or descending. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataDeletionRequestsAPI.ListDataDeletionRequests(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).DataSource(dataSource).SortBy(sortBy).SortDirection(sortDirection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataDeletionRequestsAPI.ListDataDeletionRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDataDeletionRequests`: ListDataDeletionRequestsResponse
	fmt.Fprintf(os.Stdout, "Response from `DataDeletionRequestsAPI.ListDataDeletionRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDataDeletionRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [default to 100]
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 
 **clusterId** | **string** | Set to only return the requests of this cluster. | 
 **dataSource** | **string** | Set to only return the requests deleting from this telemetry signal. | 
 **sortBy** | **string** | Field to sort the requests by. Omit to sort by creation time, newest first. | 
 **sortDirection** | **string** | Sort direction: ascending or descending. | 

### Return type

[**ListDataDeletionRequestsResponse**](ListDataDeletionRequestsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

