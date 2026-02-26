# \MonitorsAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMonitor**](MonitorsAPI.md#CreateMonitor) | **Post** /v1/monitors | 
[**DeleteMonitor**](MonitorsAPI.md#DeleteMonitor) | **Delete** /v1/monitors/{id} | 
[**GetMonitor**](MonitorsAPI.md#GetMonitor) | **Get** /v1/monitors/{id} | 
[**ListMonitors**](MonitorsAPI.md#ListMonitors) | **Get** /v1/monitors | 
[**UpdateMonitor**](MonitorsAPI.md#UpdateMonitor) | **Put** /v1/monitors/{id} | 



## CreateMonitor

> CreateMonitorResponse CreateMonitor(ctx).CreateMonitorRequest(createMonitorRequest).Execute()





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
	createMonitorRequest := *openapiclient.NewCreateMonitorRequest("Name_example", openapiclient.createMonitor_request_configuration{MonitorConfigurationAnomalyLog: openapiclient.NewMonitorConfigurationAnomalyLog("Type_example", *openapiclient.NewMonitorConfigurationAnomalyLogCondition("Formula_example", "ConditionType_example"), "NoDataBehavior_example", float32(123), []openapiclient.MonitorConfigurationMetricGroupByFieldsInner{*openapiclient.NewMonitorConfigurationMetricGroupByFieldsInner([]string{"Fields_example"}, float32(123))}, []openapiclient.MonitorAggregationQuery{*openapiclient.NewMonitorAggregationQuery(openapiclient.Aggregate{AggregateAverage: openapiclient.NewAggregateAverage("Type_example", "Field_example")}, "Filter_example")})}, float32(123), "Owner_example", "Permissions_example") // CreateMonitorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitorsAPI.CreateMonitor(context.Background()).CreateMonitorRequest(createMonitorRequest).Execute()
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
 **createMonitorRequest** | [**CreateMonitorRequest**](CreateMonitorRequest.md) |  | 

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
	id := "id_example" // string | 

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
**id** | **string** |  | 

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


## ListMonitors

> ListMonitorsResponse ListMonitors(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitorsAPI.ListMonitors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsAPI.ListMonitors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMonitors`: ListMonitorsResponse
	fmt.Fprintf(os.Stdout, "Response from `MonitorsAPI.ListMonitors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMonitorsRequest struct via the builder pattern


### Return type

[**ListMonitorsResponse**](ListMonitorsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMonitor

> UpdateMonitorResponse UpdateMonitor(ctx, id).CreateMonitorRequest(createMonitorRequest).Execute()





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
	createMonitorRequest := *openapiclient.NewCreateMonitorRequest("Name_example", openapiclient.createMonitor_request_configuration{MonitorConfigurationAnomalyLog: openapiclient.NewMonitorConfigurationAnomalyLog("Type_example", *openapiclient.NewMonitorConfigurationAnomalyLogCondition("Formula_example", "ConditionType_example"), "NoDataBehavior_example", float32(123), []openapiclient.MonitorConfigurationMetricGroupByFieldsInner{*openapiclient.NewMonitorConfigurationMetricGroupByFieldsInner([]string{"Fields_example"}, float32(123))}, []openapiclient.MonitorAggregationQuery{*openapiclient.NewMonitorAggregationQuery(openapiclient.Aggregate{AggregateAverage: openapiclient.NewAggregateAverage("Type_example", "Field_example")}, "Filter_example")})}, float32(123), "Owner_example", "Permissions_example") // CreateMonitorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitorsAPI.UpdateMonitor(context.Background(), id).CreateMonitorRequest(createMonitorRequest).Execute()
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

 **createMonitorRequest** | [**CreateMonitorRequest**](CreateMonitorRequest.md) |  | 

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

