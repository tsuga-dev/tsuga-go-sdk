# \MonitorFoldersAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMonitorFolder**](MonitorFoldersAPI.md#CreateMonitorFolder) | **Post** /v1/monitor-folders | 
[**DeleteMonitorFolder**](MonitorFoldersAPI.md#DeleteMonitorFolder) | **Delete** /v1/monitor-folders/{id} | 
[**GetMonitorFolder**](MonitorFoldersAPI.md#GetMonitorFolder) | **Get** /v1/monitor-folders/{id} | 
[**QueryMonitorFolders**](MonitorFoldersAPI.md#QueryMonitorFolders) | **Post** /v1/monitor-folders/query | 
[**UpdateMonitorFolder**](MonitorFoldersAPI.md#UpdateMonitorFolder) | **Put** /v1/monitor-folders/{id} | 



## CreateMonitorFolder

> CreateMonitorFolderResponse CreateMonitorFolder(ctx).UpdateMonitorFolderRequest(updateMonitorFolderRequest).Execute()





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
	updateMonitorFolderRequest := *openapiclient.NewUpdateMonitorFolderRequest("Name_example", "Owner_example") // UpdateMonitorFolderRequest | Monitor folder create or update request. Provide a name, owner team, optional parent folder, and optional tags.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitorFoldersAPI.CreateMonitorFolder(context.Background()).UpdateMonitorFolderRequest(updateMonitorFolderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorFoldersAPI.CreateMonitorFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMonitorFolder`: CreateMonitorFolderResponse
	fmt.Fprintf(os.Stdout, "Response from `MonitorFoldersAPI.CreateMonitorFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMonitorFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateMonitorFolderRequest** | [**UpdateMonitorFolderRequest**](UpdateMonitorFolderRequest.md) | Monitor folder create or update request. Provide a name, owner team, optional parent folder, and optional tags. | 

### Return type

[**CreateMonitorFolderResponse**](CreateMonitorFolderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMonitorFolder

> DeleteMonitorFolderResponse DeleteMonitorFolder(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the monitor folder to delete. Use the `id` returned by folder list, get, create, or update responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitorFoldersAPI.DeleteMonitorFolder(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorFoldersAPI.DeleteMonitorFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMonitorFolder`: DeleteMonitorFolderResponse
	fmt.Fprintf(os.Stdout, "Response from `MonitorFoldersAPI.DeleteMonitorFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the monitor folder to delete. Use the &#x60;id&#x60; returned by folder list, get, create, or update responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMonitorFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteMonitorFolderResponse**](DeleteMonitorFolderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorFolder

> GetMonitorFolderResponse GetMonitorFolder(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the monitor folder to retrieve. Use the `id` returned by folder list, get, create, or update responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitorFoldersAPI.GetMonitorFolder(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorFoldersAPI.GetMonitorFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMonitorFolder`: GetMonitorFolderResponse
	fmt.Fprintf(os.Stdout, "Response from `MonitorFoldersAPI.GetMonitorFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the monitor folder to retrieve. Use the &#x60;id&#x60; returned by folder list, get, create, or update responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMonitorFolderResponse**](GetMonitorFolderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryMonitorFolders

> QueryMonitorFoldersResponse QueryMonitorFolders(ctx).QueryMonitorFoldersRequest(queryMonitorFoldersRequest).Execute()





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
	queryMonitorFoldersRequest := *openapiclient.NewQueryMonitorFoldersRequest() // QueryMonitorFoldersRequest | Monitor folder query request. Use filters, sorting, limit, and offset to page through visible monitor folders.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitorFoldersAPI.QueryMonitorFolders(context.Background()).QueryMonitorFoldersRequest(queryMonitorFoldersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorFoldersAPI.QueryMonitorFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryMonitorFolders`: QueryMonitorFoldersResponse
	fmt.Fprintf(os.Stdout, "Response from `MonitorFoldersAPI.QueryMonitorFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryMonitorFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryMonitorFoldersRequest** | [**QueryMonitorFoldersRequest**](QueryMonitorFoldersRequest.md) | Monitor folder query request. Use filters, sorting, limit, and offset to page through visible monitor folders. | 

### Return type

[**QueryMonitorFoldersResponse**](QueryMonitorFoldersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMonitorFolder

> UpdateMonitorFolderResponse UpdateMonitorFolder(ctx, id).UpdateMonitorFolderRequest(updateMonitorFolderRequest).Execute()





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
	id := "id_example" // string | Identifier of the monitor folder to update. Use the `id` returned by folder list, get, create, or update responses.
	updateMonitorFolderRequest := *openapiclient.NewUpdateMonitorFolderRequest("Name_example", "Owner_example") // UpdateMonitorFolderRequest | Monitor folder create or update request. Provide a name, owner team, optional parent folder, and optional tags.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitorFoldersAPI.UpdateMonitorFolder(context.Background(), id).UpdateMonitorFolderRequest(updateMonitorFolderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorFoldersAPI.UpdateMonitorFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMonitorFolder`: UpdateMonitorFolderResponse
	fmt.Fprintf(os.Stdout, "Response from `MonitorFoldersAPI.UpdateMonitorFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the monitor folder to update. Use the &#x60;id&#x60; returned by folder list, get, create, or update responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMonitorFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMonitorFolderRequest** | [**UpdateMonitorFolderRequest**](UpdateMonitorFolderRequest.md) | Monitor folder create or update request. Provide a name, owner team, optional parent folder, and optional tags. | 

### Return type

[**UpdateMonitorFolderResponse**](UpdateMonitorFolderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

