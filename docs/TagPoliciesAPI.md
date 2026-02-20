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

> CreateTagPolicy200Response CreateTagPolicy(ctx).CreateTagPolicyRequest(createTagPolicyRequest).Execute()





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
	createTagPolicyRequest := *openapiclient.NewCreateTagPolicyRequest("Name_example", false, "TagKey_example", []string{"AllowedTagValues_example"}, false, openapiclient.createTagPolicy_request_configuration{CreateTagPolicyRequestConfigurationOneOf: openapiclient.NewCreateTagPolicyRequestConfigurationOneOf("Type_example", []string{"AssetTypes_example"}, false)}, "Owner_example") // CreateTagPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagPoliciesAPI.CreateTagPolicy(context.Background()).CreateTagPolicyRequest(createTagPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagPoliciesAPI.CreateTagPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTagPolicy`: CreateTagPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `TagPoliciesAPI.CreateTagPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTagPolicyRequest** | [**CreateTagPolicyRequest**](CreateTagPolicyRequest.md) |  | 

### Return type

[**CreateTagPolicy200Response**](CreateTagPolicy200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTagPolicy

> DeleteIngestionApiKey200Response DeleteTagPolicy(ctx, id).Execute()





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
	resp, r, err := apiClient.TagPoliciesAPI.DeleteTagPolicy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagPoliciesAPI.DeleteTagPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTagPolicy`: DeleteIngestionApiKey200Response
	fmt.Fprintf(os.Stdout, "Response from `TagPoliciesAPI.DeleteTagPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagPolicyRequest struct via the builder pattern


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


## GetTagPolicy

> CreateTagPolicy200Response GetTagPolicy(ctx, id).Execute()





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
	resp, r, err := apiClient.TagPoliciesAPI.GetTagPolicy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagPoliciesAPI.GetTagPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagPolicy`: CreateTagPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `TagPoliciesAPI.GetTagPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateTagPolicy200Response**](CreateTagPolicy200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTagPolicies

> ListTagPolicies200Response ListTagPolicies(ctx).Execute()





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
	// response from `ListTagPolicies`: ListTagPolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `TagPoliciesAPI.ListTagPolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTagPoliciesRequest struct via the builder pattern


### Return type

[**ListTagPolicies200Response**](ListTagPolicies200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTagPolicy

> CreateTagPolicy200Response UpdateTagPolicy(ctx, id).CreateTagPolicyRequest(createTagPolicyRequest).Execute()





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
	createTagPolicyRequest := *openapiclient.NewCreateTagPolicyRequest("Name_example", false, "TagKey_example", []string{"AllowedTagValues_example"}, false, openapiclient.createTagPolicy_request_configuration{CreateTagPolicyRequestConfigurationOneOf: openapiclient.NewCreateTagPolicyRequestConfigurationOneOf("Type_example", []string{"AssetTypes_example"}, false)}, "Owner_example") // CreateTagPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagPoliciesAPI.UpdateTagPolicy(context.Background(), id).CreateTagPolicyRequest(createTagPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagPoliciesAPI.UpdateTagPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTagPolicy`: CreateTagPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `TagPoliciesAPI.UpdateTagPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTagPolicyRequest** | [**CreateTagPolicyRequest**](CreateTagPolicyRequest.md) |  | 

### Return type

[**CreateTagPolicy200Response**](CreateTagPolicy200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

