# \QualityReportsAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLatestQualityReport**](QualityReportsAPI.md#GetLatestQualityReport) | **Get** /v1/quality-reports/latest | 
[**GetQualityReportExamples**](QualityReportsAPI.md#GetQualityReportExamples) | **Get** /v1/quality-reports/{id}/examples | 



## GetLatestQualityReport

> GetLatestQualityReportResponse GetLatestQualityReport(ctx).ClusterId(clusterId).Team(team).Limit(limit).Offset(offset).Execute()





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
	clusterId := "clusterId_example" // string | Identifier of the cluster whose latest quality-report rows are returned. Use the id from the clusters API. Required when the organization has multiple clusters; omit only in single-cluster organizations. Unknown cluster IDs return RESOURCE_NOT_FOUND. (optional)
	team := "team_example" // string | Team name used to scope the report. Matched against Tsuga team names. Optional; when set, returns only rows for that team and omits global rows. Unknown team names return RESOURCE_NOT_FOUND; teams the caller cannot access return an empty result. (optional)
	limit := int32(56) // int32 | Maximum number of items to return in this page. Omit to use the public API default of 100. (optional) (default to 100)
	offset := int32(56) // int32 | Zero-based index of the first matching item to return. Increase it with `limit` to request later pages. If `limit` is provided without `offset`, the offset defaults to 0. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QualityReportsAPI.GetLatestQualityReport(context.Background()).ClusterId(clusterId).Team(team).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityReportsAPI.GetLatestQualityReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestQualityReport`: GetLatestQualityReportResponse
	fmt.Fprintf(os.Stdout, "Response from `QualityReportsAPI.GetLatestQualityReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestQualityReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Identifier of the cluster whose latest quality-report rows are returned. Use the id from the clusters API. Required when the organization has multiple clusters; omit only in single-cluster organizations. Unknown cluster IDs return RESOURCE_NOT_FOUND. | 
 **team** | **string** | Team name used to scope the report. Matched against Tsuga team names. Optional; when set, returns only rows for that team and omits global rows. Unknown team names return RESOURCE_NOT_FOUND; teams the caller cannot access return an empty result. | 
 **limit** | **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [default to 100]
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 

### Return type

[**GetLatestQualityReportResponse**](GetLatestQualityReportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQualityReportExamples

> GetQualityReportExamplesResponse GetQualityReportExamples(ctx, id).Limit(limit).Offset(offset).Execute()





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
	id := "id_example" // string | Identifier of the quality-report row whose examples are returned. Use the id from the latest-report endpoint. Unknown ids, rows the caller cannot access, and cluster-wide rows requested by a non-admin all return RESOURCE_NOT_FOUND.
	limit := int32(56) // int32 | Maximum number of items to return in this page. Omit to use the public API default of 100. (optional) (default to 100)
	offset := int32(56) // int32 | Zero-based index of the first matching item to return. Increase it with `limit` to request later pages. If `limit` is provided without `offset`, the offset defaults to 0. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QualityReportsAPI.GetQualityReportExamples(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityReportsAPI.GetQualityReportExamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQualityReportExamples`: GetQualityReportExamplesResponse
	fmt.Fprintf(os.Stdout, "Response from `QualityReportsAPI.GetQualityReportExamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the quality-report row whose examples are returned. Use the id from the latest-report endpoint. Unknown ids, rows the caller cannot access, and cluster-wide rows requested by a non-admin all return RESOURCE_NOT_FOUND. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQualityReportExamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [default to 100]
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 

### Return type

[**GetQualityReportExamplesResponse**](GetQualityReportExamplesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

