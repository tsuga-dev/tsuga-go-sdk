# \InvitationsAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InviteUsers**](InvitationsAPI.md#InviteUsers) | **Post** /v1/invitations | 
[**QueryInvitations**](InvitationsAPI.md#QueryInvitations) | **Post** /v1/invitations/query | 



## InviteUsers

> InviteUsersResponse InviteUsers(ctx).InviteUsersRequestInner(inviteUsersRequestInner).Execute()





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
	inviteUsersRequestInner := []openapiclient.InviteUsersRequestInner{openapiclient.inviteUsers_request_inner{InviteUsersRequestInnerOneOf: openapiclient.NewInviteUsersRequestInnerOneOf("Email_example", "TeamId_example", "TeamRole_example", "UserRole_example")}} // []InviteUsersRequestInner | Invitations to create. Up to 10 invitations are accepted per request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitationsAPI.InviteUsers(context.Background()).InviteUsersRequestInner(inviteUsersRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsAPI.InviteUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InviteUsers`: InviteUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `InvitationsAPI.InviteUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInviteUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inviteUsersRequestInner** | [**[]InviteUsersRequestInner**](InviteUsersRequestInner.md) | Invitations to create. Up to 10 invitations are accepted per request. | 

### Return type

[**InviteUsersResponse**](InviteUsersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryInvitations

> QueryInvitationsResponse QueryInvitations(ctx).QueryInvitationsRequest(queryInvitationsRequest).Execute()





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
	queryInvitationsRequest := *openapiclient.NewQueryInvitationsRequest() // QueryInvitationsRequest | Invitation query request. Use filters, sorting, limit, and offset to page through the organization invitations.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitationsAPI.QueryInvitations(context.Background()).QueryInvitationsRequest(queryInvitationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsAPI.QueryInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryInvitations`: QueryInvitationsResponse
	fmt.Fprintf(os.Stdout, "Response from `InvitationsAPI.QueryInvitations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryInvitationsRequest** | [**QueryInvitationsRequest**](QueryInvitationsRequest.md) | Invitation query request. Use filters, sorting, limit, and offset to page through the organization invitations. | 

### Return type

[**QueryInvitationsResponse**](QueryInvitationsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

