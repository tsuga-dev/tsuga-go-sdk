# \InventoryAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCloudAccount**](InventoryAPI.md#CreateCloudAccount) | **Post** /v1/inventory/cloud-accounts | 
[**DeleteCloudAccount**](InventoryAPI.md#DeleteCloudAccount) | **Delete** /v1/inventory/cloud-accounts/{id} | 
[**GetCloudAccount**](InventoryAPI.md#GetCloudAccount) | **Get** /v1/inventory/cloud-accounts/{id} | 
[**ListCloudAccounts**](InventoryAPI.md#ListCloudAccounts) | **Get** /v1/inventory/cloud-accounts | 
[**QueryInventoryResources**](InventoryAPI.md#QueryInventoryResources) | **Post** /v1/inventory/resources/query | 
[**UpdateCloudAccount**](InventoryAPI.md#UpdateCloudAccount) | **Put** /v1/inventory/cloud-accounts/{id} | 



## CreateCloudAccount

> CreateCloudAccountResponse CreateCloudAccount(ctx).CreateCloudAccountRequest(createCloudAccountRequest).Execute()





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
	createCloudAccountRequest := *openapiclient.NewCreateCloudAccountRequest("CloudType_example", "CloudAccountId_example", openapiclient.createCloudAccount_request_connectionSettings{AwsConnectionSettings: openapiclient.NewAwsConnectionSettings("Type_example", "AccountId_example", "ExternalId_example", "RoleArn_example")}) // CreateCloudAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.CreateCloudAccount(context.Background()).CreateCloudAccountRequest(createCloudAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.CreateCloudAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCloudAccount`: CreateCloudAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.CreateCloudAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloudAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCloudAccountRequest** | [**CreateCloudAccountRequest**](CreateCloudAccountRequest.md) |  | 

### Return type

[**CreateCloudAccountResponse**](CreateCloudAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCloudAccount

> DeleteCloudAccountResponse DeleteCloudAccount(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the cloud account to delete. Use the `id` returned by the cloud account list or create operations.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.DeleteCloudAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.DeleteCloudAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCloudAccount`: DeleteCloudAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.DeleteCloudAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the cloud account to delete. Use the &#x60;id&#x60; returned by the cloud account list or create operations. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteCloudAccountResponse**](DeleteCloudAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudAccount

> GetCloudAccountResponse GetCloudAccount(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the cloud account to retrieve. Use the `id` returned by the cloud account list or create operations.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.GetCloudAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.GetCloudAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudAccount`: GetCloudAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.GetCloudAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the cloud account to retrieve. Use the &#x60;id&#x60; returned by the cloud account list or create operations. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCloudAccountResponse**](GetCloudAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudAccounts

> ListCloudAccountsResponse ListCloudAccounts(ctx).Execute()





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
	resp, r, err := apiClient.InventoryAPI.ListCloudAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.ListCloudAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCloudAccounts`: ListCloudAccountsResponse
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.ListCloudAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCloudAccountsRequest struct via the builder pattern


### Return type

[**ListCloudAccountsResponse**](ListCloudAccountsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryInventoryResources

> QueryInventoryResourcesResponse QueryInventoryResources(ctx).QueryInventoryResourcesRequest(queryInventoryResourcesRequest).Execute()





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
	queryInventoryResourcesRequest := *openapiclient.NewQueryInventoryResourcesRequest() // QueryInventoryResourcesRequest | Inventory resource query request. Use filters, sorting, limit, and offset to page through discovered resources.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.QueryInventoryResources(context.Background()).QueryInventoryResourcesRequest(queryInventoryResourcesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.QueryInventoryResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryInventoryResources`: QueryInventoryResourcesResponse
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.QueryInventoryResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryInventoryResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryInventoryResourcesRequest** | [**QueryInventoryResourcesRequest**](QueryInventoryResourcesRequest.md) | Inventory resource query request. Use filters, sorting, limit, and offset to page through discovered resources. | 

### Return type

[**QueryInventoryResourcesResponse**](QueryInventoryResourcesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudAccount

> UpdateCloudAccountResponse UpdateCloudAccount(ctx, id).UpdateCloudAccountRequest(updateCloudAccountRequest).Execute()





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
	id := "id_example" // string | Identifier of the cloud account to update. Use the `id` returned by the cloud account list or create operations.
	updateCloudAccountRequest := *openapiclient.NewUpdateCloudAccountRequest("AccountFriendlyName_example") // UpdateCloudAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.UpdateCloudAccount(context.Background(), id).UpdateCloudAccountRequest(updateCloudAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.UpdateCloudAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCloudAccount`: UpdateCloudAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.UpdateCloudAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the cloud account to update. Use the &#x60;id&#x60; returned by the cloud account list or create operations. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCloudAccountRequest** | [**UpdateCloudAccountRequest**](UpdateCloudAccountRequest.md) |  | 

### Return type

[**UpdateCloudAccountResponse**](UpdateCloudAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

