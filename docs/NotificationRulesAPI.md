# \NotificationRulesAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationRule**](NotificationRulesAPI.md#CreateNotificationRule) | **Post** /v1/notification-rules | 
[**DeleteNotificationRule**](NotificationRulesAPI.md#DeleteNotificationRule) | **Delete** /v1/notification-rules/{id} | 
[**GetNotificationRule**](NotificationRulesAPI.md#GetNotificationRule) | **Get** /v1/notification-rules/{id} | 
[**ListNotificationRules**](NotificationRulesAPI.md#ListNotificationRules) | **Get** /v1/notification-rules | 
[**UpdateNotificationRule**](NotificationRulesAPI.md#UpdateNotificationRule) | **Put** /v1/notification-rules/{id} | 



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
	resp, r, err := apiClient.NotificationRulesAPI.CreateNotificationRule(context.Background()).CreateNotificationRuleRequest(createNotificationRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRulesAPI.CreateNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNotificationRule`: CreateNotificationRule200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationRulesAPI.CreateNotificationRule`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationRulesAPI.DeleteNotificationRule(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRulesAPI.DeleteNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNotificationRule`: DeleteDashboard200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationRulesAPI.DeleteNotificationRule`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationRulesAPI.GetNotificationRule(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRulesAPI.GetNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationRule`: CreateNotificationRule200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationRulesAPI.GetNotificationRule`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationRulesAPI.ListNotificationRules(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRulesAPI.ListNotificationRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNotificationRules`: ListNotificationRules200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationRulesAPI.ListNotificationRules`: %v\n", resp)
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
	resp, r, err := apiClient.NotificationRulesAPI.UpdateNotificationRule(context.Background(), id).CreateNotificationRuleRequest(createNotificationRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRulesAPI.UpdateNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNotificationRule`: CreateNotificationRule200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationRulesAPI.UpdateNotificationRule`: %v\n", resp)
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

