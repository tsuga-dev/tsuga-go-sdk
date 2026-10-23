# \CliAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordCliRun**](CliAPI.md#RecordCliRun) | **Post** /v1/cli-runs | 



## RecordCliRun

> RecordCliRunResponse RecordCliRun(ctx).RecordCliRunRequest(recordCliRunRequest).Execute()





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
	recordCliRunRequest := *openapiclient.NewRecordCliRunRequest("CliVersion_example", "Command_example", int32(123), int32(123), "OsType_example", "NodeVersion_example", "TraceId_example", "SpanId_example") // RecordCliRunRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CliAPI.RecordCliRun(context.Background()).RecordCliRunRequest(recordCliRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CliAPI.RecordCliRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordCliRun`: RecordCliRunResponse
	fmt.Fprintf(os.Stdout, "Response from `CliAPI.RecordCliRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecordCliRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recordCliRunRequest** | [**RecordCliRunRequest**](RecordCliRunRequest.md) |  | 

### Return type

[**RecordCliRunResponse**](RecordCliRunResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

