# \CustomRolesAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomRole**](CustomRolesAPI.md#CreateCustomRole) | **Post** /v1/custom-roles | 
[**DeleteCustomRole**](CustomRolesAPI.md#DeleteCustomRole) | **Delete** /v1/custom-roles/{id} | 
[**GetCustomRole**](CustomRolesAPI.md#GetCustomRole) | **Get** /v1/custom-roles/{id} | 
[**ListCustomRoles**](CustomRolesAPI.md#ListCustomRoles) | **Get** /v1/custom-roles | 
[**UpdateCustomRole**](CustomRolesAPI.md#UpdateCustomRole) | **Put** /v1/custom-roles/{id} | 



## CreateCustomRole

> CreateCustomRoleResponse CreateCustomRole(ctx).CreateCustomRoleRequest(createCustomRoleRequest).Execute()





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
	createCustomRoleRequest := *openapiclient.NewCreateCustomRoleRequest("Name_example", openapiclient.createCustomRole_request_teamScope{CreateCustomRoleRequestTeamScopeOneOf: openapiclient.NewCreateCustomRoleRequestTeamScopeOneOf("Type_example")}, *openapiclient.NewCreateCustomRoleRequestAssets("Team_example"), []string{"Telemetry_example"}) // CreateCustomRoleRequest | Custom role create or update request. Provide the role name, optional description, team scope, resource type grants, readable telemetry signals, and optional telemetry restriction.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomRolesAPI.CreateCustomRole(context.Background()).CreateCustomRoleRequest(createCustomRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomRolesAPI.CreateCustomRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomRole`: CreateCustomRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomRolesAPI.CreateCustomRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomRoleRequest** | [**CreateCustomRoleRequest**](CreateCustomRoleRequest.md) | Custom role create or update request. Provide the role name, optional description, team scope, resource type grants, readable telemetry signals, and optional telemetry restriction. | 

### Return type

[**CreateCustomRoleResponse**](CreateCustomRoleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomRole

> DeleteCustomRoleResponse DeleteCustomRole(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the custom role to delete. Use the `id` returned by custom role list, get, create, or update responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomRolesAPI.DeleteCustomRole(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomRolesAPI.DeleteCustomRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCustomRole`: DeleteCustomRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomRolesAPI.DeleteCustomRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the custom role to delete. Use the &#x60;id&#x60; returned by custom role list, get, create, or update responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteCustomRoleResponse**](DeleteCustomRoleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomRole

> GetCustomRoleResponse GetCustomRole(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the custom role to retrieve. Use the `id` returned by custom role list, create, or update responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomRolesAPI.GetCustomRole(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomRolesAPI.GetCustomRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomRole`: GetCustomRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomRolesAPI.GetCustomRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the custom role to retrieve. Use the &#x60;id&#x60; returned by custom role list, create, or update responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCustomRoleResponse**](GetCustomRoleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomRoles

> ListCustomRolesResponse ListCustomRoles(ctx).Execute()





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
	resp, r, err := apiClient.CustomRolesAPI.ListCustomRoles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomRolesAPI.ListCustomRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomRoles`: ListCustomRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomRolesAPI.ListCustomRoles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomRolesRequest struct via the builder pattern


### Return type

[**ListCustomRolesResponse**](ListCustomRolesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomRole

> UpdateCustomRoleResponse UpdateCustomRole(ctx, id).UpdateCustomRoleRequest(updateCustomRoleRequest).Execute()





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
	id := "id_example" // string | Identifier of the custom role to update. Use the `id` returned by custom role list, get, or create responses.
	updateCustomRoleRequest := *openapiclient.NewUpdateCustomRoleRequest("Name_example", openapiclient.updateCustomRole_request_teamScope{CreateCustomRoleRequestTeamScopeOneOf: openapiclient.NewCreateCustomRoleRequestTeamScopeOneOf("Type_example")}, *openapiclient.NewCreateCustomRoleRequestAssets("Team_example"), []string{"Telemetry_example"}) // UpdateCustomRoleRequest | Custom role create or update request. Provide the role name, optional description, team scope, resource type grants, readable telemetry signals, and optional telemetry restriction.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomRolesAPI.UpdateCustomRole(context.Background(), id).UpdateCustomRoleRequest(updateCustomRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomRolesAPI.UpdateCustomRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomRole`: UpdateCustomRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomRolesAPI.UpdateCustomRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the custom role to update. Use the &#x60;id&#x60; returned by custom role list, get, or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomRoleRequest** | [**UpdateCustomRoleRequest**](UpdateCustomRoleRequest.md) | Custom role create or update request. Provide the role name, optional description, team scope, resource type grants, readable telemetry signals, and optional telemetry restriction. | 

### Return type

[**UpdateCustomRoleResponse**](UpdateCustomRoleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

