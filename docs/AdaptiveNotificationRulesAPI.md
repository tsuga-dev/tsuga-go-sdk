# \AdaptiveNotificationRulesAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAdaptiveNotificationRule**](AdaptiveNotificationRulesAPI.md#CreateAdaptiveNotificationRule) | **Post** /v1/adaptive-notification-rules | 
[**DeleteAdaptiveNotificationRule**](AdaptiveNotificationRulesAPI.md#DeleteAdaptiveNotificationRule) | **Delete** /v1/adaptive-notification-rules/{id} | 
[**GetAdaptiveNotificationRule**](AdaptiveNotificationRulesAPI.md#GetAdaptiveNotificationRule) | **Get** /v1/adaptive-notification-rules/{id} | 
[**ListAdaptiveNotificationRules**](AdaptiveNotificationRulesAPI.md#ListAdaptiveNotificationRules) | **Get** /v1/adaptive-notification-rules | 
[**UpdateAdaptiveNotificationRule**](AdaptiveNotificationRulesAPI.md#UpdateAdaptiveNotificationRule) | **Put** /v1/adaptive-notification-rules/{id} | 



## CreateAdaptiveNotificationRule

> CreateAdaptiveNotificationRuleResponse CreateAdaptiveNotificationRule(ctx).CreateAdaptiveNotificationRuleRequest(createAdaptiveNotificationRuleRequest).Execute()





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
	createAdaptiveNotificationRuleRequest := *openapiclient.NewCreateAdaptiveNotificationRuleRequest("Name_example", "Owner_example", false, *openapiclient.NewCreateNotificationRuleRequestTargetsInner("Id_example", openapiclient.createNotificationRule_request_targets_inner_config{RuleTargetInputEmail: openapiclient.NewRuleTargetInputEmail("Type_example", []string{"Addresses_example"})})) // CreateAdaptiveNotificationRuleRequest | Adaptive notification rule create or update request. Provide the monitors or SLOs the rule reacts to, the owner, the active state, and the fallback target.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdaptiveNotificationRulesAPI.CreateAdaptiveNotificationRule(context.Background()).CreateAdaptiveNotificationRuleRequest(createAdaptiveNotificationRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptiveNotificationRulesAPI.CreateAdaptiveNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAdaptiveNotificationRule`: CreateAdaptiveNotificationRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `AdaptiveNotificationRulesAPI.CreateAdaptiveNotificationRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdaptiveNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAdaptiveNotificationRuleRequest** | [**CreateAdaptiveNotificationRuleRequest**](CreateAdaptiveNotificationRuleRequest.md) | Adaptive notification rule create or update request. Provide the monitors or SLOs the rule reacts to, the owner, the active state, and the fallback target. | 

### Return type

[**CreateAdaptiveNotificationRuleResponse**](CreateAdaptiveNotificationRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAdaptiveNotificationRule

> DeleteAdaptiveNotificationRuleResponse DeleteAdaptiveNotificationRule(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the adaptive notification rule to delete. Use the `id` returned by adaptive notification rule list, get, create, or update responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdaptiveNotificationRulesAPI.DeleteAdaptiveNotificationRule(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptiveNotificationRulesAPI.DeleteAdaptiveNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAdaptiveNotificationRule`: DeleteAdaptiveNotificationRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `AdaptiveNotificationRulesAPI.DeleteAdaptiveNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the adaptive notification rule to delete. Use the &#x60;id&#x60; returned by adaptive notification rule list, get, create, or update responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdaptiveNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteAdaptiveNotificationRuleResponse**](DeleteAdaptiveNotificationRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdaptiveNotificationRule

> GetAdaptiveNotificationRuleResponse GetAdaptiveNotificationRule(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the adaptive notification rule to retrieve. Use the `id` returned by adaptive notification rule list or create responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdaptiveNotificationRulesAPI.GetAdaptiveNotificationRule(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptiveNotificationRulesAPI.GetAdaptiveNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdaptiveNotificationRule`: GetAdaptiveNotificationRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `AdaptiveNotificationRulesAPI.GetAdaptiveNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the adaptive notification rule to retrieve. Use the &#x60;id&#x60; returned by adaptive notification rule list or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdaptiveNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAdaptiveNotificationRuleResponse**](GetAdaptiveNotificationRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAdaptiveNotificationRules

> ListAdaptiveNotificationRulesResponse ListAdaptiveNotificationRules(ctx).Limit(limit).Offset(offset).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdaptiveNotificationRulesAPI.ListAdaptiveNotificationRules(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptiveNotificationRulesAPI.ListAdaptiveNotificationRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAdaptiveNotificationRules`: ListAdaptiveNotificationRulesResponse
	fmt.Fprintf(os.Stdout, "Response from `AdaptiveNotificationRulesAPI.ListAdaptiveNotificationRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAdaptiveNotificationRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [default to 100]
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 

### Return type

[**ListAdaptiveNotificationRulesResponse**](ListAdaptiveNotificationRulesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdaptiveNotificationRule

> UpdateAdaptiveNotificationRuleResponse UpdateAdaptiveNotificationRule(ctx, id).UpdateAdaptiveNotificationRuleRequest(updateAdaptiveNotificationRuleRequest).Execute()





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
	id := "id_example" // string | Identifier of the adaptive notification rule to update. Use the `id` returned by adaptive notification rule list, get, or create responses.
	updateAdaptiveNotificationRuleRequest := *openapiclient.NewUpdateAdaptiveNotificationRuleRequest("Name_example", "Owner_example", false, *openapiclient.NewCreateNotificationRuleRequestTargetsInner("Id_example", openapiclient.createNotificationRule_request_targets_inner_config{RuleTargetInputEmail: openapiclient.NewRuleTargetInputEmail("Type_example", []string{"Addresses_example"})})) // UpdateAdaptiveNotificationRuleRequest | Adaptive notification rule create or update request. Provide the monitors or SLOs the rule reacts to, the owner, the active state, and the fallback target.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdaptiveNotificationRulesAPI.UpdateAdaptiveNotificationRule(context.Background(), id).UpdateAdaptiveNotificationRuleRequest(updateAdaptiveNotificationRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptiveNotificationRulesAPI.UpdateAdaptiveNotificationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAdaptiveNotificationRule`: UpdateAdaptiveNotificationRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `AdaptiveNotificationRulesAPI.UpdateAdaptiveNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the adaptive notification rule to update. Use the &#x60;id&#x60; returned by adaptive notification rule list, get, or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdaptiveNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAdaptiveNotificationRuleRequest** | [**UpdateAdaptiveNotificationRuleRequest**](UpdateAdaptiveNotificationRuleRequest.md) | Adaptive notification rule create or update request. Provide the monitors or SLOs the rule reacts to, the owner, the active state, and the fallback target. | 

### Return type

[**UpdateAdaptiveNotificationRuleResponse**](UpdateAdaptiveNotificationRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

