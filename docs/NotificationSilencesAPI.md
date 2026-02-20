# \NotificationSilencesAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationSilence**](NotificationSilencesAPI.md#CreateNotificationSilence) | **Post** /v1/notification-silences | 
[**DeleteNotificationSilence**](NotificationSilencesAPI.md#DeleteNotificationSilence) | **Delete** /v1/notification-silences/{id} | 
[**GetNotificationSilence**](NotificationSilencesAPI.md#GetNotificationSilence) | **Get** /v1/notification-silences/{id} | 
[**ListNotificationSilences**](NotificationSilencesAPI.md#ListNotificationSilences) | **Get** /v1/notification-silences | 
[**UpdateNotificationSilence**](NotificationSilencesAPI.md#UpdateNotificationSilence) | **Put** /v1/notification-silences/{id} | 



## CreateNotificationSilence

> CreateNotificationSilence200Response CreateNotificationSilence(ctx).CreateNotificationSilenceRequest(createNotificationSilenceRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/tsuga-dev/tsuga-go-sdk"
)

func main() {
	createNotificationSilenceRequest := *openapiclient.NewCreateNotificationSilenceRequest("Name_example", "Reason_example", "Owner_example", false, openapiclient.createNotificationSilence_request_schedule{CreateNotificationSilenceRequestScheduleOneOf: openapiclient.NewCreateNotificationSilenceRequestScheduleOneOf("Type_example", time.Now(), time.Now())}, openapiclient.createNotificationRule_request_teamsFilter{CreateNotificationRuleRequestTeamsFilterOneOf: openapiclient.NewCreateNotificationRuleRequestTeamsFilterOneOf("Type_example", []string{"Teams_example"})}, []float32{float32(123)}, []string{"TransitionTypesFilter_example"}) // CreateNotificationSilenceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationSilencesAPI.CreateNotificationSilence(context.Background()).CreateNotificationSilenceRequest(createNotificationSilenceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationSilencesAPI.CreateNotificationSilence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNotificationSilence`: CreateNotificationSilence200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationSilencesAPI.CreateNotificationSilence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationSilenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNotificationSilenceRequest** | [**CreateNotificationSilenceRequest**](CreateNotificationSilenceRequest.md) |  | 

### Return type

[**CreateNotificationSilence200Response**](CreateNotificationSilence200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationSilence

> DeleteIngestionApiKey200Response DeleteNotificationSilence(ctx, id).Execute()





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
	resp, r, err := apiClient.NotificationSilencesAPI.DeleteNotificationSilence(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationSilencesAPI.DeleteNotificationSilence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNotificationSilence`: DeleteIngestionApiKey200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationSilencesAPI.DeleteNotificationSilence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationSilenceRequest struct via the builder pattern


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


## GetNotificationSilence

> CreateNotificationSilence200Response GetNotificationSilence(ctx, id).Execute()





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
	resp, r, err := apiClient.NotificationSilencesAPI.GetNotificationSilence(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationSilencesAPI.GetNotificationSilence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationSilence`: CreateNotificationSilence200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationSilencesAPI.GetNotificationSilence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationSilenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateNotificationSilence200Response**](CreateNotificationSilence200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationSilences

> ListNotificationSilences200Response ListNotificationSilences(ctx).Execute()





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
	resp, r, err := apiClient.NotificationSilencesAPI.ListNotificationSilences(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationSilencesAPI.ListNotificationSilences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNotificationSilences`: ListNotificationSilences200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationSilencesAPI.ListNotificationSilences`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationSilencesRequest struct via the builder pattern


### Return type

[**ListNotificationSilences200Response**](ListNotificationSilences200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationSilence

> CreateNotificationSilence200Response UpdateNotificationSilence(ctx, id).CreateNotificationSilenceRequest(createNotificationSilenceRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/tsuga-dev/tsuga-go-sdk"
)

func main() {
	id := "id_example" // string | 
	createNotificationSilenceRequest := *openapiclient.NewCreateNotificationSilenceRequest("Name_example", "Reason_example", "Owner_example", false, openapiclient.createNotificationSilence_request_schedule{CreateNotificationSilenceRequestScheduleOneOf: openapiclient.NewCreateNotificationSilenceRequestScheduleOneOf("Type_example", time.Now(), time.Now())}, openapiclient.createNotificationRule_request_teamsFilter{CreateNotificationRuleRequestTeamsFilterOneOf: openapiclient.NewCreateNotificationRuleRequestTeamsFilterOneOf("Type_example", []string{"Teams_example"})}, []float32{float32(123)}, []string{"TransitionTypesFilter_example"}) // CreateNotificationSilenceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationSilencesAPI.UpdateNotificationSilence(context.Background(), id).CreateNotificationSilenceRequest(createNotificationSilenceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationSilencesAPI.UpdateNotificationSilence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNotificationSilence`: CreateNotificationSilence200Response
	fmt.Fprintf(os.Stdout, "Response from `NotificationSilencesAPI.UpdateNotificationSilence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationSilenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNotificationSilenceRequest** | [**CreateNotificationSilenceRequest**](CreateNotificationSilenceRequest.md) |  | 

### Return type

[**CreateNotificationSilence200Response**](CreateNotificationSilence200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

