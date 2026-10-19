# \QueuesAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQueue**](QueuesAPI.md#GetQueue) | **Get** /v1/queues/by-name | 
[**ListQueues**](QueuesAPI.md#ListQueues) | **Get** /v1/queues | 



## GetQueue

> GetQueueResponse GetQueue(ctx).Name(name).MessagingSystem(messagingSystem).Env(env).ClusterId(clusterId).Execute()





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
	name := "name_example" // string | Queue name (`messaging.destination.name`).
	messagingSystem := "messagingSystem_example" // string | Messaging system (`messaging.system`), for example `kafka` or `aws_sqs`.
	env := "env_example" // string | Environment the queue was observed in.
	clusterId := "clusterId_example" // string | Telemetry cluster to target. Omit to use the first cluster in the organization; this does not query all clusters. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueuesAPI.GetQueue(context.Background()).Name(name).MessagingSystem(messagingSystem).Env(env).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueuesAPI.GetQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueue`: GetQueueResponse
	fmt.Fprintf(os.Stdout, "Response from `QueuesAPI.GetQueue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Queue name (&#x60;messaging.destination.name&#x60;). | 
 **messagingSystem** | **string** | Messaging system (&#x60;messaging.system&#x60;), for example &#x60;kafka&#x60; or &#x60;aws_sqs&#x60;. | 
 **env** | **string** | Environment the queue was observed in. | 
 **clusterId** | **string** | Telemetry cluster to target. Omit to use the first cluster in the organization; this does not query all clusters. | 

### Return type

[**GetQueueResponse**](GetQueueResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQueues

> ListQueuesResponse ListQueues(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).Env(env).StatsWindow(statsWindow).Execute()





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
	clusterId := "clusterId_example" // string | Telemetry cluster to target. Omit to use the first cluster in the organization; this does not query all clusters. (optional)
	env := "env_example" // string | Environment used to restrict the queues. (optional)
	statsWindow := "statsWindow_example" // string | Lookback over messaging spans used to list the queues and compute their rates. Defaults to `1h`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueuesAPI.ListQueues(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).Env(env).StatsWindow(statsWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueuesAPI.ListQueues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQueues`: ListQueuesResponse
	fmt.Fprintf(os.Stdout, "Response from `QueuesAPI.ListQueues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListQueuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [default to 100]
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 
 **clusterId** | **string** | Telemetry cluster to target. Omit to use the first cluster in the organization; this does not query all clusters. | 
 **env** | **string** | Environment used to restrict the queues. | 
 **statsWindow** | **string** | Lookback over messaging spans used to list the queues and compute their rates. Defaults to &#x60;1h&#x60;. | 

### Return type

[**ListQueuesResponse**](ListQueuesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

