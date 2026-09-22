# \SlosAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSlo**](SlosAPI.md#CreateSlo) | **Post** /v1/slos | 
[**DeleteSlo**](SlosAPI.md#DeleteSlo) | **Delete** /v1/slos/{id} | 
[**GetSlo**](SlosAPI.md#GetSlo) | **Get** /v1/slos/{id} | 
[**QuerySlos**](SlosAPI.md#QuerySlos) | **Post** /v1/slos/query | 
[**UpdateSlo**](SlosAPI.md#UpdateSlo) | **Put** /v1/slos/{id} | 



## CreateSlo

> CreateSloResponse CreateSlo(ctx).CreateSloRequest(createSloRequest).Execute()





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
	createSloRequest := *openapiclient.NewCreateSloRequest("Name_example", openapiclient.updateSlo_request_configuration{EventSloConfiguration: openapiclient.NewEventSloConfiguration("Type_example", "DataSource_example", *openapiclient.NewSloQueryFormula([]openapiclient.SloAggregationQuery{*openapiclient.NewSloAggregationQuery(openapiclient.InputAggregate{InputAggregateAverage: openapiclient.NewInputAggregateAverage("Type_example", "Field_example")}, "Filter_example")}, "Formula_example"), *openapiclient.NewSloQueryFormula([]openapiclient.SloAggregationQuery{*openapiclient.NewSloAggregationQuery(openapiclient.InputAggregate{InputAggregateAverage: openapiclient.NewInputAggregateAverage("Type_example", "Field_example")}, "Filter_example")}, "Formula_example"), "NoDataBehavior_example")}, float32(123), int32(123), "Owner_example", "Permissions_example", []openapiclient.CreateSloRequestAlertsInner{*openapiclient.NewCreateSloRequestAlertsInner(float32(123), openapiclient.updateSlo_request_alerts_inner_configuration{BurnRateAlert: openapiclient.NewBurnRateAlert("Type_example", float32(123))})}) // CreateSloRequest | SLO creation request. Provide ownership, data access, SLI configuration, target, timeframe, optional cluster scope, and the alerts to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlosAPI.CreateSlo(context.Background()).CreateSloRequest(createSloRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlosAPI.CreateSlo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSlo`: CreateSloResponse
	fmt.Fprintf(os.Stdout, "Response from `SlosAPI.CreateSlo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSloRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSloRequest** | [**CreateSloRequest**](CreateSloRequest.md) | SLO creation request. Provide ownership, data access, SLI configuration, target, timeframe, optional cluster scope, and the alerts to create. | 

### Return type

[**CreateSloResponse**](CreateSloResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSlo

> DeleteSloResponse DeleteSlo(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the SLO to delete. Use the `id` returned by SLO query or create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlosAPI.DeleteSlo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlosAPI.DeleteSlo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSlo`: DeleteSloResponse
	fmt.Fprintf(os.Stdout, "Response from `SlosAPI.DeleteSlo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the SLO to delete. Use the &#x60;id&#x60; returned by SLO query or create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSloRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteSloResponse**](DeleteSloResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlo

> GetSloResponse GetSlo(ctx, id).Execute()





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
	id := "id_example" // string | SLO ID returned by SLO query, create, or update operations.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlosAPI.GetSlo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlosAPI.GetSlo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSlo`: GetSloResponse
	fmt.Fprintf(os.Stdout, "Response from `SlosAPI.GetSlo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | SLO ID returned by SLO query, create, or update operations. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSloRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSloResponse**](GetSloResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuerySlos

> QuerySlosResponse QuerySlos(ctx).QuerySlosRequest(querySlosRequest).Execute()





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
	querySlosRequest := *openapiclient.NewQuerySlosRequest() // QuerySlosRequest | SLO query request. Use filters, sorting, limit, and offset to page through visible SLOs.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlosAPI.QuerySlos(context.Background()).QuerySlosRequest(querySlosRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlosAPI.QuerySlos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuerySlos`: QuerySlosResponse
	fmt.Fprintf(os.Stdout, "Response from `SlosAPI.QuerySlos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuerySlosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **querySlosRequest** | [**QuerySlosRequest**](QuerySlosRequest.md) | SLO query request. Use filters, sorting, limit, and offset to page through visible SLOs. | 

### Return type

[**QuerySlosResponse**](QuerySlosResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSlo

> UpdateSloResponse UpdateSlo(ctx, id).UpdateSloRequest(updateSloRequest).Execute()





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
	id := "id_example" // string | Identifier of the SLO to update. Use the `id` returned by SLO query or create.
	updateSloRequest := *openapiclient.NewUpdateSloRequest("Name_example", openapiclient.updateSlo_request_configuration{EventSloConfiguration: openapiclient.NewEventSloConfiguration("Type_example", "DataSource_example", *openapiclient.NewSloQueryFormula([]openapiclient.SloAggregationQuery{*openapiclient.NewSloAggregationQuery(openapiclient.InputAggregate{InputAggregateAverage: openapiclient.NewInputAggregateAverage("Type_example", "Field_example")}, "Filter_example")}, "Formula_example"), *openapiclient.NewSloQueryFormula([]openapiclient.SloAggregationQuery{*openapiclient.NewSloAggregationQuery(openapiclient.InputAggregate{InputAggregateAverage: openapiclient.NewInputAggregateAverage("Type_example", "Field_example")}, "Filter_example")}, "Formula_example"), "NoDataBehavior_example")}, float32(123), int32(123), "Owner_example", "Permissions_example", []openapiclient.UpdateSloRequestAlertsInner{*openapiclient.NewUpdateSloRequestAlertsInner(float32(123), openapiclient.updateSlo_request_alerts_inner_configuration{BurnRateAlert: openapiclient.NewBurnRateAlert("Type_example", float32(123))})}) // UpdateSloRequest | SLO update request. Send all required SLO fields and the full alert list. Alerts are reconciled to exactly the provided list: alerts with a known id are updated, alerts without an id are created, and existing alerts absent from the list are deleted. Omit `clusterIds` to preserve the current cluster scope; send an empty array to run on all clusters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlosAPI.UpdateSlo(context.Background(), id).UpdateSloRequest(updateSloRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlosAPI.UpdateSlo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSlo`: UpdateSloResponse
	fmt.Fprintf(os.Stdout, "Response from `SlosAPI.UpdateSlo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the SLO to update. Use the &#x60;id&#x60; returned by SLO query or create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSloRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSloRequest** | [**UpdateSloRequest**](UpdateSloRequest.md) | SLO update request. Send all required SLO fields and the full alert list. Alerts are reconciled to exactly the provided list: alerts with a known id are updated, alerts without an id are created, and existing alerts absent from the list are deleted. Omit &#x60;clusterIds&#x60; to preserve the current cluster scope; send an empty array to run on all clusters. | 

### Return type

[**UpdateSloResponse**](UpdateSloResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

