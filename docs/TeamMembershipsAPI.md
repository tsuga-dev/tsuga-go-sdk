# \TeamMembershipsAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserToTeam**](TeamMembershipsAPI.md#AddUserToTeam) | **Post** /v1/team-memberships | 
[**BulkAddUsersToTeam**](TeamMembershipsAPI.md#BulkAddUsersToTeam) | **Post** /v1/team-memberships/bulk-add | 
[**BulkRemoveUsersFromTeam**](TeamMembershipsAPI.md#BulkRemoveUsersFromTeam) | **Delete** /v1/team-memberships/bulk-remove | 
[**GetTeamMemberships**](TeamMembershipsAPI.md#GetTeamMemberships) | **Get** /v1/team-memberships | 
[**RemoveUserFromTeam**](TeamMembershipsAPI.md#RemoveUserFromTeam) | **Delete** /v1/team-memberships | 
[**UpdateTeamMembership**](TeamMembershipsAPI.md#UpdateTeamMembership) | **Put** /v1/team-memberships | 



## AddUserToTeam

> AddUserToTeamResponse AddUserToTeam(ctx).UpdateTeamMembershipRequest(updateTeamMembershipRequest).Execute()





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
	updateTeamMembershipRequest := *openapiclient.NewUpdateTeamMembershipRequest("UserId_example", "TeamId_example", "RoleKey_example") // UpdateTeamMembershipRequest | Team membership request. Provide the user, team, and role assignment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamMembershipsAPI.AddUserToTeam(context.Background()).UpdateTeamMembershipRequest(updateTeamMembershipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamMembershipsAPI.AddUserToTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUserToTeam`: AddUserToTeamResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamMembershipsAPI.AddUserToTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUserToTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTeamMembershipRequest** | [**UpdateTeamMembershipRequest**](UpdateTeamMembershipRequest.md) | Team membership request. Provide the user, team, and role assignment. | 

### Return type

[**AddUserToTeamResponse**](AddUserToTeamResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkAddUsersToTeam

> BulkAddUsersToTeamResponse BulkAddUsersToTeam(ctx).BulkAddUsersToTeamRequest(bulkAddUsersToTeamRequest).Execute()





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
	bulkAddUsersToTeamRequest := *openapiclient.NewBulkAddUsersToTeamRequest("TeamId_example", []openapiclient.BulkAddUsersToTeamRequestMembersInner{nil}) // BulkAddUsersToTeamRequest | Bulk team-membership creation request. Provide one target team and the users to add by ID or email, with a role for each user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamMembershipsAPI.BulkAddUsersToTeam(context.Background()).BulkAddUsersToTeamRequest(bulkAddUsersToTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamMembershipsAPI.BulkAddUsersToTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkAddUsersToTeam`: BulkAddUsersToTeamResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamMembershipsAPI.BulkAddUsersToTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkAddUsersToTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkAddUsersToTeamRequest** | [**BulkAddUsersToTeamRequest**](BulkAddUsersToTeamRequest.md) | Bulk team-membership creation request. Provide one target team and the users to add by ID or email, with a role for each user. | 

### Return type

[**BulkAddUsersToTeamResponse**](BulkAddUsersToTeamResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkRemoveUsersFromTeam

> BulkRemoveUsersFromTeamResponse BulkRemoveUsersFromTeam(ctx).BulkRemoveUsersFromTeamRequest(bulkRemoveUsersFromTeamRequest).Execute()





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
	bulkRemoveUsersFromTeamRequest := *openapiclient.NewBulkRemoveUsersFromTeamRequest("TeamId_example") // BulkRemoveUsersFromTeamRequest | Bulk team-membership removal request. Provide one target team and either `userIds` or `userEmails`, not both, with no duplicates.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamMembershipsAPI.BulkRemoveUsersFromTeam(context.Background()).BulkRemoveUsersFromTeamRequest(bulkRemoveUsersFromTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamMembershipsAPI.BulkRemoveUsersFromTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkRemoveUsersFromTeam`: BulkRemoveUsersFromTeamResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamMembershipsAPI.BulkRemoveUsersFromTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkRemoveUsersFromTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkRemoveUsersFromTeamRequest** | [**BulkRemoveUsersFromTeamRequest**](BulkRemoveUsersFromTeamRequest.md) | Bulk team-membership removal request. Provide one target team and either &#x60;userIds&#x60; or &#x60;userEmails&#x60;, not both, with no duplicates. | 

### Return type

[**BulkRemoveUsersFromTeamResponse**](BulkRemoveUsersFromTeamResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamMemberships

> GetTeamMembershipsResponse GetTeamMemberships(ctx).UserId(userId).TeamId(teamId).Execute()





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
	userId := "userId_example" // string | Restrict results to memberships for the user with this ID. Omit to include memberships for any user. (optional)
	teamId := "teamId_example" // string | Restrict results to memberships for the team with this ID. Omit to include memberships for any team. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamMembershipsAPI.GetTeamMemberships(context.Background()).UserId(userId).TeamId(teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamMembershipsAPI.GetTeamMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamMemberships`: GetTeamMembershipsResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamMembershipsAPI.GetTeamMemberships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | Restrict results to memberships for the user with this ID. Omit to include memberships for any user. | 
 **teamId** | **string** | Restrict results to memberships for the team with this ID. Omit to include memberships for any team. | 

### Return type

[**GetTeamMembershipsResponse**](GetTeamMembershipsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserFromTeam

> RemoveUserFromTeamResponse RemoveUserFromTeam(ctx).RemoveUserFromTeamRequest(removeUserFromTeamRequest).Execute()





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
	removeUserFromTeamRequest := *openapiclient.NewRemoveUserFromTeamRequest("UserId_example", "TeamId_example") // RemoveUserFromTeamRequest | Team membership removal request. Provide the user and team to disconnect.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamMembershipsAPI.RemoveUserFromTeam(context.Background()).RemoveUserFromTeamRequest(removeUserFromTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamMembershipsAPI.RemoveUserFromTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveUserFromTeam`: RemoveUserFromTeamResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamMembershipsAPI.RemoveUserFromTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserFromTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeUserFromTeamRequest** | [**RemoveUserFromTeamRequest**](RemoveUserFromTeamRequest.md) | Team membership removal request. Provide the user and team to disconnect. | 

### Return type

[**RemoveUserFromTeamResponse**](RemoveUserFromTeamResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeamMembership

> UpdateTeamMembershipResponse UpdateTeamMembership(ctx).UpdateTeamMembershipRequest(updateTeamMembershipRequest).Execute()





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
	updateTeamMembershipRequest := *openapiclient.NewUpdateTeamMembershipRequest("UserId_example", "TeamId_example", "RoleKey_example") // UpdateTeamMembershipRequest | Team membership request. Provide the user, team, and role assignment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamMembershipsAPI.UpdateTeamMembership(context.Background()).UpdateTeamMembershipRequest(updateTeamMembershipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamMembershipsAPI.UpdateTeamMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTeamMembership`: UpdateTeamMembershipResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamMembershipsAPI.UpdateTeamMembership`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTeamMembershipRequest** | [**UpdateTeamMembershipRequest**](UpdateTeamMembershipRequest.md) | Team membership request. Provide the user, team, and role assignment. | 

### Return type

[**UpdateTeamMembershipResponse**](UpdateTeamMembershipResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

