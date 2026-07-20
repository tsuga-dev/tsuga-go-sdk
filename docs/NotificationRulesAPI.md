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

> CreateNotificationRuleResponse CreateNotificationRule(ctx).CreateNotificationRuleRequest(createNotificationRuleRequest).Execute()





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
	createNotificationRuleRequest := *openapiclient.NewCreateNotificationRuleRequest("Name_example", openapiclient.createNotificationRule_request_teamsFilter{AllPublicTeams: openapiclient.NewAllPublicTeams("Type_example")}, []float32{float32(123)}, []string{"TransitionTypesFilter_example"}, "Owner_example", false, []openapiclient.CreateNotificationRuleRequestTargetsInner{*openapiclient.NewCreateNotificationRuleRequestTargetsInner("Id_example", openapiclient.createNotificationRule_request_targets_inner_config{RuleTargetInputEmail: openapiclient.NewRuleTargetInputEmail("Type_example", []string{"Addresses_example"})})}) // CreateNotificationRuleRequest | Notification rule create or update request. Provide matching filters, owner, tags, active state, and delivery targets.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationRulesAPI.CreateNotificationRule(context.Background()).CreateNotificationRuleRequest(createNotificationRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRulesAPI.CreateNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNotificationRule`: CreateNotificationRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationRulesAPI.CreateNotificationRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNotificationRuleRequest** | [**CreateNotificationRuleRequest**](CreateNotificationRuleRequest.md) | Notification rule create or update request. Provide matching filters, owner, tags, active state, and delivery targets. | 

### Return type

[**CreateNotificationRuleResponse**](CreateNotificationRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationRule

> DeleteNotificationRuleResponse DeleteNotificationRule(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the notification rule to delete. Use the `id` returned by notification rule list, get, create, or update responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationRulesAPI.DeleteNotificationRule(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRulesAPI.DeleteNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNotificationRule`: DeleteNotificationRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationRulesAPI.DeleteNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the notification rule to delete. Use the &#x60;id&#x60; returned by notification rule list, get, create, or update responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteNotificationRuleResponse**](DeleteNotificationRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationRule

> GetNotificationRuleResponse GetNotificationRule(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the notification rule to retrieve. Use the `id` returned by notification rule list or create responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationRulesAPI.GetNotificationRule(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRulesAPI.GetNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationRule`: GetNotificationRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationRulesAPI.GetNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the notification rule to retrieve. Use the &#x60;id&#x60; returned by notification rule list or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNotificationRuleResponse**](GetNotificationRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationRules

> ListNotificationRulesResponse ListNotificationRules(ctx).Limit(limit).Offset(offset).Execute()





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
	limit := int32(56) // int32 | Maximum number of items to return in this page. Valid values are 1 through 1000. (optional)
	offset := int32(56) // int32 | Zero-based index of the first matching item to return. Increase it with `limit` to request later pages. If `limit` is provided without `offset`, the offset defaults to 0. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationRulesAPI.ListNotificationRules(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRulesAPI.ListNotificationRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNotificationRules`: ListNotificationRulesResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationRulesAPI.ListNotificationRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return in this page. Valid values are 1 through 1000. | 
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 

### Return type

[**ListNotificationRulesResponse**](ListNotificationRulesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationRule

> UpdateNotificationRuleResponse UpdateNotificationRule(ctx, id).UpdateNotificationRuleRequest(updateNotificationRuleRequest).Execute()





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
	id := "id_example" // string | Identifier of the notification rule to update. Use the `id` returned by notification rule list, get, or create responses.
	updateNotificationRuleRequest := *openapiclient.NewUpdateNotificationRuleRequest("Name_example", openapiclient.createNotificationRule_request_teamsFilter{AllPublicTeams: openapiclient.NewAllPublicTeams("Type_example")}, []float32{float32(123)}, []string{"TransitionTypesFilter_example"}, "Owner_example", false, []openapiclient.CreateNotificationRuleRequestTargetsInner{*openapiclient.NewCreateNotificationRuleRequestTargetsInner("Id_example", openapiclient.createNotificationRule_request_targets_inner_config{RuleTargetInputEmail: openapiclient.NewRuleTargetInputEmail("Type_example", []string{"Addresses_example"})})}) // UpdateNotificationRuleRequest | Notification rule create or update request. Provide matching filters, owner, tags, active state, and delivery targets.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationRulesAPI.UpdateNotificationRule(context.Background(), id).UpdateNotificationRuleRequest(updateNotificationRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRulesAPI.UpdateNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNotificationRule`: UpdateNotificationRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationRulesAPI.UpdateNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the notification rule to update. Use the &#x60;id&#x60; returned by notification rule list, get, or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNotificationRuleRequest** | [**UpdateNotificationRuleRequest**](UpdateNotificationRuleRequest.md) | Notification rule create or update request. Provide matching filters, owner, tags, active state, and delivery targets. | 

### Return type

[**UpdateNotificationRuleResponse**](UpdateNotificationRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

