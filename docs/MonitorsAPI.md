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

> CreateMonitor200Response CreateMonitor(ctx).CreateMonitorRequest(createMonitorRequest).Execute()





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
	// response from `CreateMonitor`: CreateMonitor200Response
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

[**CreateMonitor200Response**](CreateMonitor200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMonitor

> DeleteIngestionApiKey200Response DeleteMonitor(ctx, id).Execute()





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
	// response from `DeleteMonitor`: DeleteIngestionApiKey200Response
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

[**DeleteIngestionApiKey200Response**](DeleteIngestionApiKey200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitor

> CreateMonitor200Response GetMonitor(ctx, id).Execute()





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
	// response from `GetMonitor`: CreateMonitor200Response
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

[**CreateMonitor200Response**](CreateMonitor200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMonitors

> ListMonitors200Response ListMonitors(ctx).Execute()





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
	// response from `ListMonitors`: ListMonitors200Response
	fmt.Fprintf(os.Stdout, "Response from `MonitorsAPI.ListMonitors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMonitorsRequest struct via the builder pattern


### Return type

[**ListMonitors200Response**](ListMonitors200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMonitor

> CreateMonitor200Response UpdateMonitor(ctx, id).CreateMonitorRequest(createMonitorRequest).Execute()





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
	// response from `UpdateMonitor`: CreateMonitor200Response
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

[**CreateMonitor200Response**](CreateMonitor200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

