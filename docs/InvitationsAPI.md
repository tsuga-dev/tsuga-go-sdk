# \InvitationsAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InviteUsers**](InvitationsAPI.md#InviteUsers) | **Post** /v1/invitations | 
[**ListInvitations**](InvitationsAPI.md#ListInvitations) | **Get** /v1/invitations | 



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
	inviteUsersRequestInner := []openapiclient.InviteUsersRequestInner{*openapiclient.NewInviteUsersRequestInner("Email_example", "TeamId_example", "TeamRole_example", "UserRole_example")} // []InviteUsersRequestInner | 

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
 **inviteUsersRequestInner** | [**[]InviteUsersRequestInner**](InviteUsersRequestInner.md) |  | 

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


## ListInvitations

> ListInvitationsResponse ListInvitations(ctx).Execute()





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
	resp, r, err := apiClient.InvitationsAPI.ListInvitations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsAPI.ListInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInvitations`: ListInvitationsResponse
	fmt.Fprintf(os.Stdout, "Response from `InvitationsAPI.ListInvitations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListInvitationsRequest struct via the builder pattern


### Return type

[**ListInvitationsResponse**](ListInvitationsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

