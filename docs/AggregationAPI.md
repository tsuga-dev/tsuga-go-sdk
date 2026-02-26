# \AggregationAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AggregateScalar**](AggregationAPI.md#AggregateScalar) | **Post** /v1/aggregation/multi-query/scalar | 
[**AggregateTimeseries**](AggregationAPI.md#AggregateTimeseries) | **Post** /v1/aggregation/multi-query/timeseries | 



## AggregateScalar

> AggregateScalarResponse AggregateScalar(ctx).AggregateScalarRequest(aggregateScalarRequest).Execute()





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
	aggregateScalarRequest := *openapiclient.NewAggregateScalarRequest(*openapiclient.NewTimeRange(int64(1704067200), int64(1704067200)), []openapiclient.AggregationQuery{*openapiclient.NewAggregationQuery(openapiclient.Aggregate{AggregateAverage: openapiclient.NewAggregateAverage("Type_example", "Field_example")})}, "DataSource_example") // AggregateScalarRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AggregationAPI.AggregateScalar(context.Background()).AggregateScalarRequest(aggregateScalarRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AggregationAPI.AggregateScalar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AggregateScalar`: AggregateScalarResponse
	fmt.Fprintf(os.Stdout, "Response from `AggregationAPI.AggregateScalar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAggregateScalarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aggregateScalarRequest** | [**AggregateScalarRequest**](AggregateScalarRequest.md) |  | 

### Return type

[**AggregateScalarResponse**](AggregateScalarResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AggregateTimeseries

> AggregateTimeseriesResponse AggregateTimeseries(ctx).AggregateTimeseriesRequest(aggregateTimeseriesRequest).Execute()





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
	aggregateTimeseriesRequest := *openapiclient.NewAggregateTimeseriesRequest(*openapiclient.NewTimeRange(int64(1704067200), int64(1704067200)), []openapiclient.AggregationQuery{*openapiclient.NewAggregationQuery(openapiclient.Aggregate{AggregateAverage: openapiclient.NewAggregateAverage("Type_example", "Field_example")})}, "DataSource_example", "AggregationWindow_example") // AggregateTimeseriesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AggregationAPI.AggregateTimeseries(context.Background()).AggregateTimeseriesRequest(aggregateTimeseriesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AggregationAPI.AggregateTimeseries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AggregateTimeseries`: AggregateTimeseriesResponse
	fmt.Fprintf(os.Stdout, "Response from `AggregationAPI.AggregateTimeseries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAggregateTimeseriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aggregateTimeseriesRequest** | [**AggregateTimeseriesRequest**](AggregateTimeseriesRequest.md) |  | 

### Return type

[**AggregateTimeseriesResponse**](AggregateTimeseriesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

