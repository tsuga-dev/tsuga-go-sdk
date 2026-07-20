# \TeamsAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTeam**](TeamsAPI.md#CreateTeam) | **Post** /v1/teams | 
[**DeleteTeam**](TeamsAPI.md#DeleteTeam) | **Delete** /v1/teams/{id} | 
[**GetTeam**](TeamsAPI.md#GetTeam) | **Get** /v1/teams/{id} | 
[**ListTeams**](TeamsAPI.md#ListTeams) | **Get** /v1/teams | 
[**UpdateTeam**](TeamsAPI.md#UpdateTeam) | **Put** /v1/teams/{id} | 



## CreateTeam

> CreateTeamResponse CreateTeam(ctx).CreateTeamRequest(createTeamRequest).Execute()





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
	createTeamRequest := *openapiclient.NewCreateTeamRequest("Name_example", "Visibility_example") // CreateTeamRequest | Team create or update request. Provide the team name, optional description, visibility, and tags.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.CreateTeam(context.Background()).CreateTeamRequest(createTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.CreateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTeam`: CreateTeamResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.CreateTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTeamRequest** | [**CreateTeamRequest**](CreateTeamRequest.md) | Team create or update request. Provide the team name, optional description, visibility, and tags. | 

### Return type

[**CreateTeamResponse**](CreateTeamResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeam

> DeleteTeamResponse DeleteTeam(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the team to delete. Use the `id` returned by team list, create, or update responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.DeleteTeam(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.DeleteTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTeam`: DeleteTeamResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.DeleteTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the team to delete. Use the &#x60;id&#x60; returned by team list, create, or update responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteTeamResponse**](DeleteTeamResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeam

> GetTeamResponse GetTeam(ctx, id).Execute()





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
	id := "id_example" // string | Identifier of the team to retrieve. Use the `id` returned by team list or create responses.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.GetTeam(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GetTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeam`: GetTeamResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GetTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the team to retrieve. Use the &#x60;id&#x60; returned by team list or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTeamResponse**](GetTeamResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeams

> ListTeamsResponse ListTeams(ctx).Limit(limit).Offset(offset).Execute()





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
	limit := int32(56) // int32 | Maximum number of items to return in this page. Valid values are 1 through 1000. (optional)
	offset := int32(56) // int32 | Zero-based index of the first matching item to return. Increase it with `limit` to request later pages. If `limit` is provided without `offset`, the offset defaults to 0. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.ListTeams(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.ListTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTeams`: ListTeamsResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.ListTeams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return in this page. Valid values are 1 through 1000. | 
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 

### Return type

[**ListTeamsResponse**](ListTeamsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeam

> UpdateTeamResponse UpdateTeam(ctx, id).CreateTeamRequest(createTeamRequest).Execute()





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
	id := "id_example" // string | Identifier of the team to update. Use the `id` returned by team list, get, or create responses.
	createTeamRequest := *openapiclient.NewCreateTeamRequest("Name_example", "Visibility_example") // CreateTeamRequest | Team create or update request. Provide the team name, optional description, visibility, and tags.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.UpdateTeam(context.Background(), id).CreateTeamRequest(createTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.UpdateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTeam`: UpdateTeamResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.UpdateTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the team to update. Use the &#x60;id&#x60; returned by team list, get, or create responses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTeamRequest** | [**CreateTeamRequest**](CreateTeamRequest.md) | Team create or update request. Provide the team name, optional description, visibility, and tags. | 

### Return type

[**UpdateTeamResponse**](UpdateTeamResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

