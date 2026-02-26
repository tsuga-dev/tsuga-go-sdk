# \DashboardsAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboard**](DashboardsAPI.md#CreateDashboard) | **Post** /v1/dashboards | 
[**DeleteDashboard**](DashboardsAPI.md#DeleteDashboard) | **Delete** /v1/dashboards/{id} | 
[**GetDashboard**](DashboardsAPI.md#GetDashboard) | **Get** /v1/dashboards/{id} | 
[**ListDashboards**](DashboardsAPI.md#ListDashboards) | **Get** /v1/dashboards | 
[**UpdateDashboard**](DashboardsAPI.md#UpdateDashboard) | **Put** /v1/dashboards/{id} | 



## CreateDashboard

> CreateDashboardResponse CreateDashboard(ctx).CreateDashboardRequest(createDashboardRequest).Execute()





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
	createDashboardRequest := *openapiclient.NewCreateDashboardRequest("Name_example", "Owner_example", []openapiclient.Graph{*openapiclient.NewGraph("Id_example", openapiclient.Graph_visualization{GraphVisualizationBar: openapiclient.NewGraphVisualizationBar("Type_example", "Source_example", []openapiclient.AggregationQuery1{*openapiclient.NewAggregationQuery1(openapiclient.Aggregate{AggregateAverage: openapiclient.NewAggregateAverage("Type_example", "Field_example")})})})}) // CreateDashboardRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.CreateDashboard(context.Background()).CreateDashboardRequest(createDashboardRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.CreateDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDashboard`: CreateDashboardResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.CreateDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDashboardRequest** | [**CreateDashboardRequest**](CreateDashboardRequest.md) |  | 

### Return type

[**CreateDashboardResponse**](CreateDashboardResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboard

> DeleteDashboardResponse DeleteDashboard(ctx, id).Execute()





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
	resp, r, err := apiClient.DashboardsAPI.DeleteDashboard(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DeleteDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDashboard`: DeleteDashboardResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.DeleteDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteDashboardResponse**](DeleteDashboardResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboard

> GetDashboardResponse GetDashboard(ctx, id).Execute()





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
	resp, r, err := apiClient.DashboardsAPI.GetDashboard(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.GetDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboard`: GetDashboardResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.GetDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDashboardResponse**](GetDashboardResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDashboards

> ListDashboardsResponse ListDashboards(ctx).Owners(owners).Execute()





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
	owners := []string{"Inner_example"} // []string | Filter by owner team IDs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.ListDashboards(context.Background()).Owners(owners).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.ListDashboards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDashboards`: ListDashboardsResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.ListDashboards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDashboardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owners** | **[]string** | Filter by owner team IDs | 

### Return type

[**ListDashboardsResponse**](ListDashboardsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboard

> UpdateDashboardResponse UpdateDashboard(ctx, id).UpdateDashboardRequest(updateDashboardRequest).Execute()





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
	updateDashboardRequest := *openapiclient.NewUpdateDashboardRequest() // UpdateDashboardRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.UpdateDashboard(context.Background(), id).UpdateDashboardRequest(updateDashboardRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.UpdateDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDashboard`: UpdateDashboardResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.UpdateDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDashboardRequest** | [**UpdateDashboardRequest**](UpdateDashboardRequest.md) |  | 

### Return type

[**UpdateDashboardResponse**](UpdateDashboardResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

