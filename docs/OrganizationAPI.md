# \OrganizationAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllowedDomains**](OrganizationAPI.md#GetAllowedDomains) | **Get** /v1/allowed-domains | 
[**UpdateAllowedDomains**](OrganizationAPI.md#UpdateAllowedDomains) | **Put** /v1/allowed-domains | 



## GetAllowedDomains

> GetAllowedDomainsResponse GetAllowedDomains(ctx).Execute()





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
	resp, r, err := apiClient.OrganizationAPI.GetAllowedDomains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetAllowedDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllowedDomains`: GetAllowedDomainsResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetAllowedDomains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllowedDomainsRequest struct via the builder pattern


### Return type

[**GetAllowedDomainsResponse**](GetAllowedDomainsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAllowedDomains

> UpdateAllowedDomainsResponse UpdateAllowedDomains(ctx).UpdateAllowedDomainsRequest(updateAllowedDomainsRequest).Execute()





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
	updateAllowedDomainsRequest := *openapiclient.NewUpdateAllowedDomainsRequest([]string{"AllowedDomains_example"}) // UpdateAllowedDomainsRequest | Allowed-domain update request. Send the complete replacement list of domains that should remain allowed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.UpdateAllowedDomains(context.Background()).UpdateAllowedDomainsRequest(updateAllowedDomainsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpdateAllowedDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAllowedDomains`: UpdateAllowedDomainsResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.UpdateAllowedDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAllowedDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAllowedDomainsRequest** | [**UpdateAllowedDomainsRequest**](UpdateAllowedDomainsRequest.md) | Allowed-domain update request. Send the complete replacement list of domains that should remain allowed. | 

### Return type

[**UpdateAllowedDomainsResponse**](UpdateAllowedDomainsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

