# \AlertsAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryOngoingAlerts**](AlertsAPI.md#QueryOngoingAlerts) | **Post** /v1/alerts/ongoing/query | 



## QueryOngoingAlerts

> QueryOngoingAlertsResponse QueryOngoingAlerts(ctx).QueryOngoingAlertsRequest(queryOngoingAlertsRequest).Execute()





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
	queryOngoingAlertsRequest := *openapiclient.NewQueryOngoingAlertsRequest() // QueryOngoingAlertsRequest | Ongoing alert query request. Use filters, sorting, limit, and offset to page through the alerts currently open.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.QueryOngoingAlerts(context.Background()).QueryOngoingAlertsRequest(queryOngoingAlertsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.QueryOngoingAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryOngoingAlerts`: QueryOngoingAlertsResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.QueryOngoingAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryOngoingAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryOngoingAlertsRequest** | [**QueryOngoingAlertsRequest**](QueryOngoingAlertsRequest.md) | Ongoing alert query request. Use filters, sorting, limit, and offset to page through the alerts currently open. | 

### Return type

[**QueryOngoingAlertsResponse**](QueryOngoingAlertsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

