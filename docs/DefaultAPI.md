# \DefaultAPI

All URIs are relative to *https://api.tsuga-staging.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboard**](DefaultAPI.md#CreateDashboard) | **Post** /v1/dashboards | 
[**CreateNotificationRule**](DefaultAPI.md#CreateNotificationRule) | **Post** /v1/notification-rules | 
[**CreateRoute**](DefaultAPI.md#CreateRoute) | **Post** /v1/routes | 
[**CreateTeam**](DefaultAPI.md#CreateTeam) | **Post** /v1/teams | 
[**DeleteDashboard**](DefaultAPI.md#DeleteDashboard) | **Delete** /v1/dashboards/{id} | 
[**DeleteNotificationRule**](DefaultAPI.md#DeleteNotificationRule) | **Delete** /v1/notification-rules/{id} | 
[**DeleteRoute**](DefaultAPI.md#DeleteRoute) | **Delete** /v1/routes/{id} | 
[**DeleteTeam**](DefaultAPI.md#DeleteTeam) | **Delete** /v1/teams/{id} | 
[**GetDashboard**](DefaultAPI.md#GetDashboard) | **Get** /v1/dashboards/{id} | 
[**GetNotificationRule**](DefaultAPI.md#GetNotificationRule) | **Get** /v1/notification-rules/{id} | 
[**GetRoute**](DefaultAPI.md#GetRoute) | **Get** /v1/routes/{id} | 
[**GetTeam**](DefaultAPI.md#GetTeam) | **Get** /v1/teams/{id} | 
[**ListDashboards**](DefaultAPI.md#ListDashboards) | **Get** /v1/dashboards | 
[**ListNotificationRules**](DefaultAPI.md#ListNotificationRules) | **Get** /v1/notification-rules | 
[**ListRoutes**](DefaultAPI.md#ListRoutes) | **Get** /v1/routes | 
[**ListTeams**](DefaultAPI.md#ListTeams) | **Get** /v1/teams | 
[**UpdateDashboard**](DefaultAPI.md#UpdateDashboard) | **Put** /v1/dashboards/{id} | 
[**UpdateNotificationRule**](DefaultAPI.md#UpdateNotificationRule) | **Put** /v1/notification-rules/{id} | 
[**UpdateRoute**](DefaultAPI.md#UpdateRoute) | **Put** /v1/routes/{id} | 
[**UpdateTeam**](DefaultAPI.md#UpdateTeam) | **Put** /v1/teams/{id} | 



## CreateDashboard

> CreateDashboard200Response CreateDashboard(ctx).CreateDashboardRequest(createDashboardRequest).Execute()





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
	createDashboardRequest := *openapiclient.NewCreateDashboardRequest("Name_example", "Owner_example", []openapiclient.ListDashboards200ResponseDataInnerGraphsInner{*openapiclient.NewListDashboards200ResponseDataInnerGraphsInner("Id_example", openapiclient.listDashboards_200_response_data_inner_graphs_inner_visualization{GraphVisualizationBar: openapiclient.NewGraphVisualizationBar("Type_example", "Source_example", []openapiclient.GraphVisualizationTimeseriesQueriesInner{*openapiclient.NewGraphVisualizationTimeseriesQueriesInner(openapiclient.GraphVisualizationTimeseries_queries_inner_aggregate{AggregateAverage: openapiclient.NewAggregateAverage("Type_example", "Field_example")})})})}) // CreateDashboardRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateDashboard(context.Background()).CreateDashboardRequest(createDashboardRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDashboard`: CreateDashboard200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDashboardRequest** | [**CreateDashboardRequest**](CreateDashboardRequest.md) |  | 

### Return type

[**CreateDashboard200Response**](CreateDashboard200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNotificationRule

> CreateNotificationRule200Response CreateNotificationRule(ctx).CreateNotificationRuleRequest(createNotificationRuleRequest).Execute()





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
	createNotificationRuleRequest := *openapiclient.NewCreateNotificationRuleRequest("Name_example", []string{"TeamsFilter_example"}, []int32{int32(123)}, []string{"TransitionTypesFilter_example"}, "Owner_example", false, []openapiclient.CreateNotificationRuleRequestTargetsInner{*openapiclient.NewCreateNotificationRuleRequestTargetsInner("Id_example", openapiclient.createNotificationRule_request_targets_inner_config{RuleTargetInputEmail: openapiclient.NewRuleTargetInputEmail("Type_example", []string{"Addresses_example"})})}) // CreateNotificationRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateNotificationRule(context.Background()).CreateNotificationRuleRequest(createNotificationRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNotificationRule`: CreateNotificationRule200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateNotificationRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNotificationRuleRequest** | [**CreateNotificationRuleRequest**](CreateNotificationRuleRequest.md) |  | 

### Return type

[**CreateNotificationRule200Response**](CreateNotificationRule200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoute

> CreateRoute200Response CreateRoute(ctx).CreateRouteRequest(createRouteRequest).Execute()





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
	createRouteRequest := *openapiclient.NewCreateRouteRequest("Name_example", false, "Query_example", "Owner_example", []openapiclient.ListRoutes200ResponseDataInnerProcessorsInner{*openapiclient.NewListRoutes200ResponseDataInnerProcessorsInner("Id_example", "Name_example", "Type_example", *openapiclient.NewListRoutes200ResponseDataInnerProcessorsInnerAnyOf4Params([]openapiclient.ListRoutes200ResponseDataInnerProcessorsInnerAnyOf4ParamsItemsInner{*openapiclient.NewListRoutes200ResponseDataInnerProcessorsInnerAnyOf4ParamsItemsInner("Query_example", []openapiclient.Processor{*openapiclient.NewProcessor("Id_example", "Name_example", "Type_example", *openapiclient.NewListRoutes200ResponseDataInnerProcessorsInnerAnyOf4Params([]openapiclient.ListRoutes200ResponseDataInnerProcessorsInnerAnyOf4ParamsItemsInner{*openapiclient.NewListRoutes200ResponseDataInnerProcessorsInnerAnyOf4ParamsItemsInner("Query_example", []openapiclient.Processor{*openapiclient.NewProcessor("Id_example", "Name_example", "Type_example", )})}))})}))}) // CreateRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateRoute(context.Background()).CreateRouteRequest(createRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoute`: CreateRoute200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRouteRequest** | [**CreateRouteRequest**](CreateRouteRequest.md) |  | 

### Return type

[**CreateRoute200Response**](CreateRoute200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTeam

> CreateTeam200Response CreateTeam(ctx).CreateTeamRequest(createTeamRequest).Execute()





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
	createTeamRequest := *openapiclient.NewCreateTeamRequest("Name_example", "Visibility_example") // CreateTeamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateTeam(context.Background()).CreateTeamRequest(createTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTeam`: CreateTeam200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTeamRequest** | [**CreateTeamRequest**](CreateTeamRequest.md) |  | 

### Return type

[**CreateTeam200Response**](CreateTeam200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboard

> DeleteDashboard200Response DeleteDashboard(ctx, id).Execute()





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
	resp, r, err := apiClient.DefaultAPI.DeleteDashboard(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDashboard`: DeleteDashboard200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteDashboard`: %v\n", resp)
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

[**DeleteDashboard200Response**](DeleteDashboard200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationRule

> DeleteDashboard200Response DeleteNotificationRule(ctx, id).Execute()





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
	resp, r, err := apiClient.DefaultAPI.DeleteNotificationRule(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNotificationRule`: DeleteDashboard200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteDashboard200Response**](DeleteDashboard200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoute

> DeleteDashboard200Response DeleteRoute(ctx, id).Execute()





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
	resp, r, err := apiClient.DefaultAPI.DeleteRoute(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRoute`: DeleteDashboard200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteDashboard200Response**](DeleteDashboard200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeam

> DeleteDashboard200Response DeleteTeam(ctx, id).Execute()





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
	resp, r, err := apiClient.DefaultAPI.DeleteTeam(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTeam`: DeleteDashboard200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteDashboard200Response**](DeleteDashboard200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboard

> CreateDashboard200Response GetDashboard(ctx, id).Execute()





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
	resp, r, err := apiClient.DefaultAPI.GetDashboard(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboard`: CreateDashboard200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetDashboard`: %v\n", resp)
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

[**CreateDashboard200Response**](CreateDashboard200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationRule

> CreateNotificationRule200Response GetNotificationRule(ctx, id).Execute()





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
	resp, r, err := apiClient.DefaultAPI.GetNotificationRule(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationRule`: CreateNotificationRule200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateNotificationRule200Response**](CreateNotificationRule200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoute

> CreateRoute200Response GetRoute(ctx, id).Execute()





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
	resp, r, err := apiClient.DefaultAPI.GetRoute(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoute`: CreateRoute200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateRoute200Response**](CreateRoute200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeam

> CreateTeam200Response GetTeam(ctx, id).Execute()





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
	resp, r, err := apiClient.DefaultAPI.GetTeam(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeam`: CreateTeam200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateTeam200Response**](CreateTeam200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDashboards

> ListDashboards200Response ListDashboards(ctx).Execute()





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
	resp, r, err := apiClient.DefaultAPI.ListDashboards(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListDashboards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDashboards`: ListDashboards200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListDashboards`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDashboardsRequest struct via the builder pattern


### Return type

[**ListDashboards200Response**](ListDashboards200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationRules

> ListNotificationRules200Response ListNotificationRules(ctx).Execute()





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
	resp, r, err := apiClient.DefaultAPI.ListNotificationRules(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListNotificationRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNotificationRules`: ListNotificationRules200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListNotificationRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationRulesRequest struct via the builder pattern


### Return type

[**ListNotificationRules200Response**](ListNotificationRules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoutes

> ListRoutes200Response ListRoutes(ctx).Execute()





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
	resp, r, err := apiClient.DefaultAPI.ListRoutes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoutes`: ListRoutes200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListRoutes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRoutesRequest struct via the builder pattern


### Return type

[**ListRoutes200Response**](ListRoutes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeams

> ListTeams200Response ListTeams(ctx).Execute()





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
	resp, r, err := apiClient.DefaultAPI.ListTeams(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTeams`: ListTeams200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListTeams`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTeamsRequest struct via the builder pattern


### Return type

[**ListTeams200Response**](ListTeams200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboard

> CreateDashboard200Response UpdateDashboard(ctx, id).UpdateDashboardRequest(updateDashboardRequest).Execute()





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
	updateDashboardRequest := *openapiclient.NewUpdateDashboardRequest() // UpdateDashboardRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateDashboard(context.Background(), id).UpdateDashboardRequest(updateDashboardRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDashboard`: CreateDashboard200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateDashboard`: %v\n", resp)
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

[**CreateDashboard200Response**](CreateDashboard200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationRule

> CreateNotificationRule200Response UpdateNotificationRule(ctx, id).CreateNotificationRuleRequest(createNotificationRuleRequest).Execute()





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
	createNotificationRuleRequest := *openapiclient.NewCreateNotificationRuleRequest("Name_example", []string{"TeamsFilter_example"}, []int32{int32(123)}, []string{"TransitionTypesFilter_example"}, "Owner_example", false, []openapiclient.CreateNotificationRuleRequestTargetsInner{*openapiclient.NewCreateNotificationRuleRequestTargetsInner("Id_example", openapiclient.createNotificationRule_request_targets_inner_config{RuleTargetInputEmail: openapiclient.NewRuleTargetInputEmail("Type_example", []string{"Addresses_example"})})}) // CreateNotificationRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateNotificationRule(context.Background(), id).CreateNotificationRuleRequest(createNotificationRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNotificationRule`: CreateNotificationRule200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNotificationRuleRequest** | [**CreateNotificationRuleRequest**](CreateNotificationRuleRequest.md) |  | 

### Return type

[**CreateNotificationRule200Response**](CreateNotificationRule200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoute

> CreateRoute200Response UpdateRoute(ctx, id).CreateRouteRequest(createRouteRequest).Execute()





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
	createRouteRequest := *openapiclient.NewCreateRouteRequest("Name_example", false, "Query_example", "Owner_example", []openapiclient.ListRoutes200ResponseDataInnerProcessorsInner{*openapiclient.NewListRoutes200ResponseDataInnerProcessorsInner("Id_example", "Name_example", "Type_example", *openapiclient.NewListRoutes200ResponseDataInnerProcessorsInnerAnyOf4Params([]openapiclient.ListRoutes200ResponseDataInnerProcessorsInnerAnyOf4ParamsItemsInner{*openapiclient.NewListRoutes200ResponseDataInnerProcessorsInnerAnyOf4ParamsItemsInner("Query_example", []openapiclient.Processor{*openapiclient.NewProcessor("Id_example", "Name_example", "Type_example", *openapiclient.NewListRoutes200ResponseDataInnerProcessorsInnerAnyOf4Params([]openapiclient.ListRoutes200ResponseDataInnerProcessorsInnerAnyOf4ParamsItemsInner{*openapiclient.NewListRoutes200ResponseDataInnerProcessorsInnerAnyOf4ParamsItemsInner("Query_example", []openapiclient.Processor{*openapiclient.NewProcessor("Id_example", "Name_example", "Type_example", )})}))})}))}) // CreateRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateRoute(context.Background(), id).CreateRouteRequest(createRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoute`: CreateRoute200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRouteRequest** | [**CreateRouteRequest**](CreateRouteRequest.md) |  | 

### Return type

[**CreateRoute200Response**](CreateRoute200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeam

> CreateTeam200Response UpdateTeam(ctx, id).CreateTeamRequest(createTeamRequest).Execute()





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
	createTeamRequest := *openapiclient.NewCreateTeamRequest("Name_example", "Visibility_example") // CreateTeamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateTeam(context.Background(), id).CreateTeamRequest(createTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTeam`: CreateTeam200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTeamRequest** | [**CreateTeamRequest**](CreateTeamRequest.md) |  | 

### Return type

[**CreateTeam200Response**](CreateTeam200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

