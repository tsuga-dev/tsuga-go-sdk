# \TeamMembershipsAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserToTeam**](TeamMembershipsAPI.md#AddUserToTeam) | **Post** /v1/team-memberships | 
[**GetTeamMemberships**](TeamMembershipsAPI.md#GetTeamMemberships) | **Get** /v1/team-memberships | 
[**RemoveUserFromTeam**](TeamMembershipsAPI.md#RemoveUserFromTeam) | **Delete** /v1/team-memberships | 
[**UpdateTeamMembership**](TeamMembershipsAPI.md#UpdateTeamMembership) | **Put** /v1/team-memberships | 



## AddUserToTeam

> UpdateTeamMembership200Response AddUserToTeam(ctx).UpdateTeamMembershipRequest(updateTeamMembershipRequest).Execute()





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
	updateTeamMembershipRequest := *openapiclient.NewUpdateTeamMembershipRequest("UserId_example", "TeamId_example", "RoleKey_example") // UpdateTeamMembershipRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamMembershipsAPI.AddUserToTeam(context.Background()).UpdateTeamMembershipRequest(updateTeamMembershipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamMembershipsAPI.AddUserToTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUserToTeam`: UpdateTeamMembership200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamMembershipsAPI.AddUserToTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUserToTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTeamMembershipRequest** | [**UpdateTeamMembershipRequest**](UpdateTeamMembershipRequest.md) |  | 

### Return type

[**UpdateTeamMembership200Response**](UpdateTeamMembership200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamMemberships

> GetTeamMemberships200Response GetTeamMemberships(ctx).UserId(userId).TeamId(teamId).Execute()





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
	userId := "userId_example" // string |  (optional)
	teamId := "teamId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamMembershipsAPI.GetTeamMemberships(context.Background()).UserId(userId).TeamId(teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamMembershipsAPI.GetTeamMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamMemberships`: GetTeamMemberships200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamMembershipsAPI.GetTeamMemberships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** |  | 
 **teamId** | **string** |  | 

### Return type

[**GetTeamMemberships200Response**](GetTeamMemberships200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserFromTeam

> DeleteIngestionApiKey200Response RemoveUserFromTeam(ctx).RemoveUserFromTeamRequest(removeUserFromTeamRequest).Execute()





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
	removeUserFromTeamRequest := *openapiclient.NewRemoveUserFromTeamRequest("UserId_example", "TeamId_example") // RemoveUserFromTeamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamMembershipsAPI.RemoveUserFromTeam(context.Background()).RemoveUserFromTeamRequest(removeUserFromTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamMembershipsAPI.RemoveUserFromTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveUserFromTeam`: DeleteIngestionApiKey200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamMembershipsAPI.RemoveUserFromTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserFromTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeUserFromTeamRequest** | [**RemoveUserFromTeamRequest**](RemoveUserFromTeamRequest.md) |  | 

### Return type

[**DeleteIngestionApiKey200Response**](DeleteIngestionApiKey200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeamMembership

> UpdateTeamMembership200Response UpdateTeamMembership(ctx).UpdateTeamMembershipRequest(updateTeamMembershipRequest).Execute()





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
	updateTeamMembershipRequest := *openapiclient.NewUpdateTeamMembershipRequest("UserId_example", "TeamId_example", "RoleKey_example") // UpdateTeamMembershipRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamMembershipsAPI.UpdateTeamMembership(context.Background()).UpdateTeamMembershipRequest(updateTeamMembershipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamMembershipsAPI.UpdateTeamMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTeamMembership`: UpdateTeamMembership200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamMembershipsAPI.UpdateTeamMembership`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTeamMembershipRequest** | [**UpdateTeamMembershipRequest**](UpdateTeamMembershipRequest.md) |  | 

### Return type

[**UpdateTeamMembership200Response**](UpdateTeamMembership200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

