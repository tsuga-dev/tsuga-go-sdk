# \NotificationIntegrationsAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationIntegration**](NotificationIntegrationsAPI.md#CreateNotificationIntegration) | **Post** /v1/notification-integrations | 
[**DeleteNotificationIntegration**](NotificationIntegrationsAPI.md#DeleteNotificationIntegration) | **Delete** /v1/notification-integrations/{id} | 
[**GetNotificationIntegration**](NotificationIntegrationsAPI.md#GetNotificationIntegration) | **Get** /v1/notification-integrations/{id} | 
[**ListNotificationIntegrations**](NotificationIntegrationsAPI.md#ListNotificationIntegrations) | **Get** /v1/notification-integrations | 
[**UpdateNotificationIntegration**](NotificationIntegrationsAPI.md#UpdateNotificationIntegration) | **Put** /v1/notification-integrations/{id} | 



## CreateNotificationIntegration

> CreateNotificationIntegrationResponse CreateNotificationIntegration(ctx).CreateNotificationIntegrationRequest(createNotificationIntegrationRequest).Execute()





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
	createNotificationIntegrationRequest := *openapiclient.NewCreateNotificationIntegrationRequest("Name_example", "Owner_example", openapiclient.createNotificationIntegration_request_setting{CreateNotificationIntegrationRequestSettingOneOf: openapiclient.NewCreateNotificationIntegrationRequestSettingOneOf("Type_example", "AlertSourceConfigId_example", "Token_example")}) // CreateNotificationIntegrationRequest | Notification integration create request containing the display name, owning team, destination setting, and tags. `setting.type` cannot be `slack`; add Slack integrations from the app instead.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationIntegrationsAPI.CreateNotificationIntegration(context.Background()).CreateNotificationIntegrationRequest(createNotificationIntegrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationIntegrationsAPI.CreateNotificationIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNotificationIntegration`: CreateNotificationIntegrationResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationIntegrationsAPI.CreateNotificationIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNotificationIntegrationRequest** | [**CreateNotificationIntegrationRequest**](CreateNotificationIntegrationRequest.md) | Notification integration create request containing the display name, owning team, destination setting, and tags. &#x60;setting.type&#x60; cannot be &#x60;slack&#x60;; add Slack integrations from the app instead. | 

### Return type

[**CreateNotificationIntegrationResponse**](CreateNotificationIntegrationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationIntegration

> DeleteNotificationIntegrationResponse DeleteNotificationIntegration(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the notification integration to delete. Use the `id` returned by notification integration list, get, or create responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationIntegrationsAPI.DeleteNotificationIntegration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationIntegrationsAPI.DeleteNotificationIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNotificationIntegration`: DeleteNotificationIntegrationResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationIntegrationsAPI.DeleteNotificationIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the notification integration to delete. Use the &#x60;id&#x60; returned by notification integration list, get, or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteNotificationIntegrationResponse**](DeleteNotificationIntegrationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationIntegration

> GetNotificationIntegrationResponse GetNotificationIntegration(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the notification integration to retrieve. Use the `id` returned by notification integration list or create responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationIntegrationsAPI.GetNotificationIntegration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationIntegrationsAPI.GetNotificationIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationIntegration`: GetNotificationIntegrationResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationIntegrationsAPI.GetNotificationIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the notification integration to retrieve. Use the &#x60;id&#x60; returned by notification integration list or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNotificationIntegrationResponse**](GetNotificationIntegrationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationIntegrations

> ListNotificationIntegrationsResponse ListNotificationIntegrations(ctx).Limit(limit).Offset(offset).Execute()





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
	resp, r, err := apiClient.NotificationIntegrationsAPI.ListNotificationIntegrations(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationIntegrationsAPI.ListNotificationIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNotificationIntegrations`: ListNotificationIntegrationsResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationIntegrationsAPI.ListNotificationIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [default to 100]
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 

### Return type

[**ListNotificationIntegrationsResponse**](ListNotificationIntegrationsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationIntegration

> UpdateNotificationIntegrationResponse UpdateNotificationIntegration(ctx, id).UpdateNotificationIntegrationRequest(updateNotificationIntegrationRequest).Execute()





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
	id := "id_example" // string | Identifier of the notification integration to update. Use the `id` returned by notification integration list, get, or create responses.
	updateNotificationIntegrationRequest := *openapiclient.NewUpdateNotificationIntegrationRequest("Owner_example") // UpdateNotificationIntegrationRequest | Notification integration update request. `owner` has to be sent on every update; the other omitted fields keep their existing values.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationIntegrationsAPI.UpdateNotificationIntegration(context.Background(), id).UpdateNotificationIntegrationRequest(updateNotificationIntegrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationIntegrationsAPI.UpdateNotificationIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNotificationIntegration`: UpdateNotificationIntegrationResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationIntegrationsAPI.UpdateNotificationIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the notification integration to update. Use the &#x60;id&#x60; returned by notification integration list, get, or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNotificationIntegrationRequest** | [**UpdateNotificationIntegrationRequest**](UpdateNotificationIntegrationRequest.md) | Notification integration update request. &#x60;owner&#x60; has to be sent on every update; the other omitted fields keep their existing values. | 

### Return type

[**UpdateNotificationIntegrationResponse**](UpdateNotificationIntegrationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

