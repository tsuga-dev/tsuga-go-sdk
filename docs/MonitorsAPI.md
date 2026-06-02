# \MonitorsAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMonitor**](MonitorsAPI.md#CreateMonitor) | **Post** /v1/monitors | 
[**DeleteMonitor**](MonitorsAPI.md#DeleteMonitor) | **Delete** /v1/monitors/{id} | 
[**GetMonitor**](MonitorsAPI.md#GetMonitor) | **Get** /v1/monitors/{id} | 
[**QueryMonitors**](MonitorsAPI.md#QueryMonitors) | **Post** /v1/monitors/query | 
[**UpdateMonitor**](MonitorsAPI.md#UpdateMonitor) | **Put** /v1/monitors/{id} | 



## CreateMonitor

> CreateMonitorResponse CreateMonitor(ctx).UpdateMonitorRequest(updateMonitorRequest).Execute()





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
	updateMonitorRequest := *openapiclient.NewUpdateMonitorRequest("Name_example", openapiclient.updateMonitor_request_configuration{InputMonitorConfigurationAnomalyLog: openapiclient.NewInputMonitorConfigurationAnomalyLog("Type_example", *openapiclient.NewInputMonitorConfigurationAnomalyLogCondition("Formula_example", "ConditionType_example"), "NoDataBehavior_example", float32(123), []openapiclient.InputMonitorConfigurationMetricGroupByFieldsInner{*openapiclient.NewInputMonitorConfigurationMetricGroupByFieldsInner([]string{"Fields_example"}, float32(123))}, []openapiclient.MonitorAggregationQuery1{*openapiclient.NewMonitorAggregationQuery1(openapiclient.InputAggregate{InputAggregateAverage: openapiclient.NewInputAggregateAverage("Type_example", "Field_example")}, "Filter_example")})}, float32(123), "Owner_example", "Permissions_example") // UpdateMonitorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitorsAPI.CreateMonitor(context.Background()).UpdateMonitorRequest(updateMonitorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsAPI.CreateMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMonitor`: CreateMonitorResponse
	fmt.Fprintf(os.Stdout, "Response from `MonitorsAPI.CreateMonitor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateMonitorRequest** | [**UpdateMonitorRequest**](UpdateMonitorRequest.md) |  | 

### Return type

[**CreateMonitorResponse**](CreateMonitorResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMonitor

> DeleteMonitorResponse DeleteMonitor(ctx, id).Execute()





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
	id := "id_example" // string | The monitor ID to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitorsAPI.DeleteMonitor(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsAPI.DeleteMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMonitor`: DeleteMonitorResponse
	fmt.Fprintf(os.Stdout, "Response from `MonitorsAPI.DeleteMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The monitor ID to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteMonitorResponse**](DeleteMonitorResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitor

> GetMonitorResponse GetMonitor(ctx, id).Execute()





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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitorsAPI.GetMonitor(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsAPI.GetMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMonitor`: GetMonitorResponse
	fmt.Fprintf(os.Stdout, "Response from `MonitorsAPI.GetMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMonitorResponse**](GetMonitorResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryMonitors

> QueryMonitorsResponse QueryMonitors(ctx).QueryMonitorsRequest(queryMonitorsRequest).Execute()





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
	queryMonitorsRequest := *openapiclient.NewQueryMonitorsRequest() // QueryMonitorsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitorsAPI.QueryMonitors(context.Background()).QueryMonitorsRequest(queryMonitorsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsAPI.QueryMonitors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryMonitors`: QueryMonitorsResponse
	fmt.Fprintf(os.Stdout, "Response from `MonitorsAPI.QueryMonitors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryMonitorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryMonitorsRequest** | [**QueryMonitorsRequest**](QueryMonitorsRequest.md) |  | 

### Return type

[**QueryMonitorsResponse**](QueryMonitorsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMonitor

> UpdateMonitorResponse UpdateMonitor(ctx, id).UpdateMonitorRequest(updateMonitorRequest).Execute()





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
	id := "id_example" // string | 
	updateMonitorRequest := *openapiclient.NewUpdateMonitorRequest("Name_example", openapiclient.updateMonitor_request_configuration{InputMonitorConfigurationAnomalyLog: openapiclient.NewInputMonitorConfigurationAnomalyLog("Type_example", *openapiclient.NewInputMonitorConfigurationAnomalyLogCondition("Formula_example", "ConditionType_example"), "NoDataBehavior_example", float32(123), []openapiclient.InputMonitorConfigurationMetricGroupByFieldsInner{*openapiclient.NewInputMonitorConfigurationMetricGroupByFieldsInner([]string{"Fields_example"}, float32(123))}, []openapiclient.MonitorAggregationQuery1{*openapiclient.NewMonitorAggregationQuery1(openapiclient.InputAggregate{InputAggregateAverage: openapiclient.NewInputAggregateAverage("Type_example", "Field_example")}, "Filter_example")})}, float32(123), "Owner_example", "Permissions_example") // UpdateMonitorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitorsAPI.UpdateMonitor(context.Background(), id).UpdateMonitorRequest(updateMonitorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsAPI.UpdateMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMonitor`: UpdateMonitorResponse
	fmt.Fprintf(os.Stdout, "Response from `MonitorsAPI.UpdateMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMonitorRequest** | [**UpdateMonitorRequest**](UpdateMonitorRequest.md) |  | 

### Return type

[**UpdateMonitorResponse**](UpdateMonitorResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

