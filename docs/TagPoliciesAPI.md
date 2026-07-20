# \TagPoliciesAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTagPolicy**](TagPoliciesAPI.md#CreateTagPolicy) | **Post** /v1/tag-policies | 
[**DeleteTagPolicy**](TagPoliciesAPI.md#DeleteTagPolicy) | **Delete** /v1/tag-policies/{id} | 
[**GetTagPolicy**](TagPoliciesAPI.md#GetTagPolicy) | **Get** /v1/tag-policies/{id} | 
[**ListTagPolicies**](TagPoliciesAPI.md#ListTagPolicies) | **Get** /v1/tag-policies | 
[**UpdateTagPolicy**](TagPoliciesAPI.md#UpdateTagPolicy) | **Put** /v1/tag-policies/{id} | 



## CreateTagPolicy

> CreateTagPolicyResponse CreateTagPolicy(ctx).CreateTagPolicyRequest(createTagPolicyRequest).Execute()





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
	createTagPolicyRequest := *openapiclient.NewCreateTagPolicyRequest("Name_example", false, "TagKey_example", []*string{nil}, false, openapiclient.createTagPolicy_request_configuration{TelemetryTagPolicy: openapiclient.NewTelemetryTagPolicy("Type_example", []string{"AssetTypes_example"}, false)}, "Owner_example") // CreateTagPolicyRequest | Tag policy create or update request. Provide policy identity, owner, enforced tag key, allowed values, team scope, active state, and asset or telemetry configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagPoliciesAPI.CreateTagPolicy(context.Background()).CreateTagPolicyRequest(createTagPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagPoliciesAPI.CreateTagPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTagPolicy`: CreateTagPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `TagPoliciesAPI.CreateTagPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTagPolicyRequest** | [**CreateTagPolicyRequest**](CreateTagPolicyRequest.md) | Tag policy create or update request. Provide policy identity, owner, enforced tag key, allowed values, team scope, active state, and asset or telemetry configuration. | 

### Return type

[**CreateTagPolicyResponse**](CreateTagPolicyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTagPolicy

> DeleteTagPolicyResponse DeleteTagPolicy(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the tag policy to delete. Use the `id` returned by tag policy list, create, or update responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagPoliciesAPI.DeleteTagPolicy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagPoliciesAPI.DeleteTagPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTagPolicy`: DeleteTagPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `TagPoliciesAPI.DeleteTagPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the tag policy to delete. Use the &#x60;id&#x60; returned by tag policy list, create, or update responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteTagPolicyResponse**](DeleteTagPolicyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagPolicy

> GetTagPolicyResponse GetTagPolicy(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the tag policy to retrieve. Use the `id` returned by tag policy list or create responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagPoliciesAPI.GetTagPolicy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagPoliciesAPI.GetTagPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagPolicy`: GetTagPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `TagPoliciesAPI.GetTagPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the tag policy to retrieve. Use the &#x60;id&#x60; returned by tag policy list or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTagPolicyResponse**](GetTagPolicyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTagPolicies

> ListTagPoliciesResponse ListTagPolicies(ctx).Execute()





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
	resp, r, err := apiClient.TagPoliciesAPI.ListTagPolicies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagPoliciesAPI.ListTagPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTagPolicies`: ListTagPoliciesResponse
	fmt.Fprintf(os.Stdout, "Response from `TagPoliciesAPI.ListTagPolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTagPoliciesRequest struct via the builder pattern


### Return type

[**ListTagPoliciesResponse**](ListTagPoliciesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTagPolicy

> UpdateTagPolicyResponse UpdateTagPolicy(ctx, id).UpdateTagPolicyRequest(updateTagPolicyRequest).Execute()





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
	id := "id_example" // string | Identifier of the tag policy to update. Use the `id` returned by tag policy list, get, or create responses.
	updateTagPolicyRequest := *openapiclient.NewUpdateTagPolicyRequest("Name_example", false, "TagKey_example", []string{"AllowedTagValues_example"}, false, openapiclient.createTagPolicy_request_configuration{TelemetryTagPolicy: openapiclient.NewTelemetryTagPolicy("Type_example", []string{"AssetTypes_example"}, false)}, "Owner_example") // UpdateTagPolicyRequest | Tag policy create or update request. Provide policy identity, owner, enforced tag key, allowed values, team scope, active state, and asset or telemetry configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagPoliciesAPI.UpdateTagPolicy(context.Background(), id).UpdateTagPolicyRequest(updateTagPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagPoliciesAPI.UpdateTagPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTagPolicy`: UpdateTagPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `TagPoliciesAPI.UpdateTagPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the tag policy to update. Use the &#x60;id&#x60; returned by tag policy list, get, or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTagPolicyRequest** | [**UpdateTagPolicyRequest**](UpdateTagPolicyRequest.md) | Tag policy create or update request. Provide policy identity, owner, enforced tag key, allowed values, team scope, active state, and asset or telemetry configuration. | 

### Return type

[**UpdateTagPolicyResponse**](UpdateTagPolicyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

