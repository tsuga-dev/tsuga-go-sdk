# \DashboardFoldersAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboardFolder**](DashboardFoldersAPI.md#CreateDashboardFolder) | **Post** /v1/dashboard-folders | 
[**DeleteDashboardFolder**](DashboardFoldersAPI.md#DeleteDashboardFolder) | **Delete** /v1/dashboard-folders/{id} | 
[**GetDashboardFolder**](DashboardFoldersAPI.md#GetDashboardFolder) | **Get** /v1/dashboard-folders/{id} | 
[**QueryDashboardFolders**](DashboardFoldersAPI.md#QueryDashboardFolders) | **Post** /v1/dashboard-folders/query | 
[**UpdateDashboardFolder**](DashboardFoldersAPI.md#UpdateDashboardFolder) | **Put** /v1/dashboard-folders/{id} | 



## CreateDashboardFolder

> CreateDashboardFolderResponse CreateDashboardFolder(ctx).CreateDashboardFolderRequest(createDashboardFolderRequest).Execute()





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
	createDashboardFolderRequest := *openapiclient.NewCreateDashboardFolderRequest("Name_example", "Owner_example") // CreateDashboardFolderRequest | Dashboard folder create or update request. Provide a name, owner team, optional parent folder, and optional tags.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardFoldersAPI.CreateDashboardFolder(context.Background()).CreateDashboardFolderRequest(createDashboardFolderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardFoldersAPI.CreateDashboardFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDashboardFolder`: CreateDashboardFolderResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardFoldersAPI.CreateDashboardFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDashboardFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDashboardFolderRequest** | [**CreateDashboardFolderRequest**](CreateDashboardFolderRequest.md) | Dashboard folder create or update request. Provide a name, owner team, optional parent folder, and optional tags. | 

### Return type

[**CreateDashboardFolderResponse**](CreateDashboardFolderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboardFolder

> DeleteDashboardFolderResponse DeleteDashboardFolder(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the dashboard folder to delete. Use the `id` returned by folder list, get, create, or update responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardFoldersAPI.DeleteDashboardFolder(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardFoldersAPI.DeleteDashboardFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDashboardFolder`: DeleteDashboardFolderResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardFoldersAPI.DeleteDashboardFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the dashboard folder to delete. Use the &#x60;id&#x60; returned by folder list, get, create, or update responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteDashboardFolderResponse**](DeleteDashboardFolderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboardFolder

> GetDashboardFolderResponse GetDashboardFolder(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the dashboard folder to retrieve. Use the `id` returned by folder list, get, create, or update responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardFoldersAPI.GetDashboardFolder(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardFoldersAPI.GetDashboardFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboardFolder`: GetDashboardFolderResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardFoldersAPI.GetDashboardFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the dashboard folder to retrieve. Use the &#x60;id&#x60; returned by folder list, get, create, or update responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDashboardFolderResponse**](GetDashboardFolderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryDashboardFolders

> QueryDashboardFoldersResponse QueryDashboardFolders(ctx).QueryDashboardFoldersRequest(queryDashboardFoldersRequest).Execute()





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
	queryDashboardFoldersRequest := *openapiclient.NewQueryDashboardFoldersRequest() // QueryDashboardFoldersRequest | Dashboard folder query request. Use filters, sorting, limit, and offset to page through visible dashboard folders.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardFoldersAPI.QueryDashboardFolders(context.Background()).QueryDashboardFoldersRequest(queryDashboardFoldersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardFoldersAPI.QueryDashboardFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryDashboardFolders`: QueryDashboardFoldersResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardFoldersAPI.QueryDashboardFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryDashboardFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryDashboardFoldersRequest** | [**QueryDashboardFoldersRequest**](QueryDashboardFoldersRequest.md) | Dashboard folder query request. Use filters, sorting, limit, and offset to page through visible dashboard folders. | 

### Return type

[**QueryDashboardFoldersResponse**](QueryDashboardFoldersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboardFolder

> UpdateDashboardFolderResponse UpdateDashboardFolder(ctx, id).UpdateDashboardFolderRequest(updateDashboardFolderRequest).Execute()





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
	id := "id_example" // string | Identifier of the dashboard folder to update. Use the `id` returned by folder list, get, create, or update responses.
	updateDashboardFolderRequest := *openapiclient.NewUpdateDashboardFolderRequest("Name_example", "Owner_example") // UpdateDashboardFolderRequest | Dashboard folder create or update request. Provide a name, owner team, optional parent folder, and optional tags.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardFoldersAPI.UpdateDashboardFolder(context.Background(), id).UpdateDashboardFolderRequest(updateDashboardFolderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardFoldersAPI.UpdateDashboardFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDashboardFolder`: UpdateDashboardFolderResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardFoldersAPI.UpdateDashboardFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the dashboard folder to update. Use the &#x60;id&#x60; returned by folder list, get, create, or update responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDashboardFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDashboardFolderRequest** | [**UpdateDashboardFolderRequest**](UpdateDashboardFolderRequest.md) | Dashboard folder create or update request. Provide a name, owner team, optional parent folder, and optional tags. | 

### Return type

[**UpdateDashboardFolderResponse**](UpdateDashboardFolderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

