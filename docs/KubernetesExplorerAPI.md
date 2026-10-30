# \KubernetesExplorerAPI

All URIs are relative to *https://api.tsuga.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKubernetesDeploymentDetails**](KubernetesExplorerAPI.md#GetKubernetesDeploymentDetails) | **Get** /v1/kubernetes-explorer/deployment-details | 
[**GetKubernetesPodDetails**](KubernetesExplorerAPI.md#GetKubernetesPodDetails) | **Get** /v1/kubernetes-explorer/pod-details | 
[**GetKubernetesPodManifest**](KubernetesExplorerAPI.md#GetKubernetesPodManifest) | **Get** /v1/kubernetes-explorer/pod-manifest | 
[**ListKubernetesClusters**](KubernetesExplorerAPI.md#ListKubernetesClusters) | **Get** /v1/kubernetes-explorer/clusters | 
[**ListKubernetesDaemonSets**](KubernetesExplorerAPI.md#ListKubernetesDaemonSets) | **Get** /v1/kubernetes-explorer/daemonsets | 
[**ListKubernetesDeployments**](KubernetesExplorerAPI.md#ListKubernetesDeployments) | **Get** /v1/kubernetes-explorer/deployments | 
[**ListKubernetesEvents**](KubernetesExplorerAPI.md#ListKubernetesEvents) | **Get** /v1/kubernetes-explorer/events | 
[**ListKubernetesNamespaces**](KubernetesExplorerAPI.md#ListKubernetesNamespaces) | **Get** /v1/kubernetes-explorer/namespaces | 
[**ListKubernetesNodes**](KubernetesExplorerAPI.md#ListKubernetesNodes) | **Get** /v1/kubernetes-explorer/nodes | 
[**ListKubernetesPods**](KubernetesExplorerAPI.md#ListKubernetesPods) | **Get** /v1/kubernetes-explorer/pods | 
[**ListKubernetesStatefulSets**](KubernetesExplorerAPI.md#ListKubernetesStatefulSets) | **Get** /v1/kubernetes-explorer/statefulsets | 



## GetKubernetesDeploymentDetails

> GetKubernetesDeploymentDetailsResponse GetKubernetesDeploymentDetails(ctx).K8sCluster(k8sCluster).K8sNamespace(k8sNamespace).DeploymentName(deploymentName).ClusterId(clusterId).Execute()





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
	k8sCluster := "k8sCluster_example" // string | Kubernetes cluster name that contains the Deployment. From cluster telemetry. Required to identify the Deployment because names are only unique within a namespace and cluster.
	k8sNamespace := "k8sNamespace_example" // string | Kubernetes namespace that contains the Deployment. From workload metadata. Required to identify the Deployment because names are only unique within a namespace and cluster.
	deploymentName := "deploymentName_example" // string | Kubernetes Deployment name to retrieve. From workload telemetry. Required to fetch deployment details for the selected cluster and namespace.
	clusterId := "clusterId_example" // string | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesExplorerAPI.GetKubernetesDeploymentDetails(context.Background()).K8sCluster(k8sCluster).K8sNamespace(k8sNamespace).DeploymentName(deploymentName).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesExplorerAPI.GetKubernetesDeploymentDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesDeploymentDetails`: GetKubernetesDeploymentDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesExplorerAPI.GetKubernetesDeploymentDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesDeploymentDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **k8sCluster** | **string** | Kubernetes cluster name that contains the Deployment. From cluster telemetry. Required to identify the Deployment because names are only unique within a namespace and cluster. | 
 **k8sNamespace** | **string** | Kubernetes namespace that contains the Deployment. From workload metadata. Required to identify the Deployment because names are only unique within a namespace and cluster. | 
 **deploymentName** | **string** | Kubernetes Deployment name to retrieve. From workload telemetry. Required to fetch deployment details for the selected cluster and namespace. | 
 **clusterId** | **string** | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. | 

### Return type

[**GetKubernetesDeploymentDetailsResponse**](GetKubernetesDeploymentDetailsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesPodDetails

> GetKubernetesPodDetailsResponse GetKubernetesPodDetails(ctx).K8sCluster(k8sCluster).K8sNamespace(k8sNamespace).PodName(podName).ClusterId(clusterId).Execute()





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
	k8sCluster := "k8sCluster_example" // string | Kubernetes cluster name that contains the pod. From cluster telemetry. Required to identify the pod because pod names are only unique within a namespace and cluster.
	k8sNamespace := "k8sNamespace_example" // string | Kubernetes namespace that contains the pod. From pod metadata. Required to identify the pod because pod names are only unique within a namespace and cluster.
	podName := "podName_example" // string | Kubernetes pod name to retrieve. Set when the pod is created. Required to fetch pod details for the selected cluster and namespace.
	clusterId := "clusterId_example" // string | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesExplorerAPI.GetKubernetesPodDetails(context.Background()).K8sCluster(k8sCluster).K8sNamespace(k8sNamespace).PodName(podName).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesExplorerAPI.GetKubernetesPodDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesPodDetails`: GetKubernetesPodDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesExplorerAPI.GetKubernetesPodDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesPodDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **k8sCluster** | **string** | Kubernetes cluster name that contains the pod. From cluster telemetry. Required to identify the pod because pod names are only unique within a namespace and cluster. | 
 **k8sNamespace** | **string** | Kubernetes namespace that contains the pod. From pod metadata. Required to identify the pod because pod names are only unique within a namespace and cluster. | 
 **podName** | **string** | Kubernetes pod name to retrieve. Set when the pod is created. Required to fetch pod details for the selected cluster and namespace. | 
 **clusterId** | **string** | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. | 

### Return type

[**GetKubernetesPodDetailsResponse**](GetKubernetesPodDetailsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesPodManifest

> GetKubernetesPodManifestResponse GetKubernetesPodManifest(ctx).K8sCluster(k8sCluster).K8sNamespace(k8sNamespace).PodName(podName).ClusterId(clusterId).Execute()





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
	k8sCluster := "k8sCluster_example" // string | Kubernetes cluster name that contains the pod. From cluster telemetry. Required to identify the pod because pod names are only unique within a namespace and cluster.
	k8sNamespace := "k8sNamespace_example" // string | Kubernetes namespace that contains the pod. From pod metadata. Required to identify the pod because pod names are only unique within a namespace and cluster.
	podName := "podName_example" // string | Kubernetes pod name to retrieve. Set when the pod is created. Required to fetch pod details for the selected cluster and namespace.
	clusterId := "clusterId_example" // string | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesExplorerAPI.GetKubernetesPodManifest(context.Background()).K8sCluster(k8sCluster).K8sNamespace(k8sNamespace).PodName(podName).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesExplorerAPI.GetKubernetesPodManifest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesPodManifest`: GetKubernetesPodManifestResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesExplorerAPI.GetKubernetesPodManifest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesPodManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **k8sCluster** | **string** | Kubernetes cluster name that contains the pod. From cluster telemetry. Required to identify the pod because pod names are only unique within a namespace and cluster. | 
 **k8sNamespace** | **string** | Kubernetes namespace that contains the pod. From pod metadata. Required to identify the pod because pod names are only unique within a namespace and cluster. | 
 **podName** | **string** | Kubernetes pod name to retrieve. Set when the pod is created. Required to fetch pod details for the selected cluster and namespace. | 
 **clusterId** | **string** | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. | 

### Return type

[**GetKubernetesPodManifestResponse**](GetKubernetesPodManifestResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKubernetesClusters

> ListKubernetesClustersResponse ListKubernetesClusters(ctx).ClusterId(clusterId).Search(search).Execute()





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
	clusterId := "clusterId_example" // string | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. (optional)
	search := "search_example" // string | Substring matched against the Kubernetes cluster name. Unlike Kubernetes resource names this is a collector-assigned identifier, so uppercase letters, digits, `.`, `-` and `_` are accepted. Wildcard and other query-syntax characters are treated literally. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesExplorerAPI.ListKubernetesClusters(context.Background()).ClusterId(clusterId).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesExplorerAPI.ListKubernetesClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKubernetesClusters`: ListKubernetesClustersResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesExplorerAPI.ListKubernetesClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKubernetesClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. | 
 **search** | **string** | Substring matched against the Kubernetes cluster name. Unlike Kubernetes resource names this is a collector-assigned identifier, so uppercase letters, digits, &#x60;.&#x60;, &#x60;-&#x60; and &#x60;_&#x60; are accepted. Wildcard and other query-syntax characters are treated literally. | 

### Return type

[**ListKubernetesClustersResponse**](ListKubernetesClustersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKubernetesDaemonSets

> ListKubernetesDaemonSetsResponse ListKubernetesDaemonSets(ctx).ClusterId(clusterId).Search(search).SortBy(sortBy).SortDirection(sortDirection).Limit(limit).Offset(offset).Execute()





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
	clusterId := "clusterId_example" // string | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. (optional)
	search := "search_example" // string | Substring matched against the resource name. Kubernetes names are lowercase RFC 1123, so only lowercase letters, digits, `.` and `-` are accepted; uppercase and other characters are rejected. Wildcard and other query-syntax characters are treated literally. (optional)
	sortBy := "sortBy_example" // string | Field to sort results by. (optional)
	sortDirection := "sortDirection_example" // string | Sort direction: ascending or descending. (optional)
	limit := int32(56) // int32 | Maximum number of items to return in this page. Omit to use the public API default of 100. (optional) (default to 100)
	offset := int32(56) // int32 | Zero-based index of the first matching item to return. Increase it with `limit` to request later pages. If `limit` is provided without `offset`, the offset defaults to 0. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesExplorerAPI.ListKubernetesDaemonSets(context.Background()).ClusterId(clusterId).Search(search).SortBy(sortBy).SortDirection(sortDirection).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesExplorerAPI.ListKubernetesDaemonSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKubernetesDaemonSets`: ListKubernetesDaemonSetsResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesExplorerAPI.ListKubernetesDaemonSets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKubernetesDaemonSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. | 
 **search** | **string** | Substring matched against the resource name. Kubernetes names are lowercase RFC 1123, so only lowercase letters, digits, &#x60;.&#x60; and &#x60;-&#x60; are accepted; uppercase and other characters are rejected. Wildcard and other query-syntax characters are treated literally. | 
 **sortBy** | **string** | Field to sort results by. | 
 **sortDirection** | **string** | Sort direction: ascending or descending. | 
 **limit** | **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [default to 100]
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 

### Return type

[**ListKubernetesDaemonSetsResponse**](ListKubernetesDaemonSetsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKubernetesDeployments

> ListKubernetesDeploymentsResponse ListKubernetesDeployments(ctx).ClusterId(clusterId).Search(search).SortBy(sortBy).SortDirection(sortDirection).Limit(limit).Offset(offset).Execute()





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
	clusterId := "clusterId_example" // string | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. (optional)
	search := "search_example" // string | Substring matched against the resource name. Kubernetes names are lowercase RFC 1123, so only lowercase letters, digits, `.` and `-` are accepted; uppercase and other characters are rejected. Wildcard and other query-syntax characters are treated literally. (optional)
	sortBy := "sortBy_example" // string | Field to sort results by. (optional)
	sortDirection := "sortDirection_example" // string | Sort direction: ascending or descending. (optional)
	limit := int32(56) // int32 | Maximum number of items to return in this page. Omit to use the public API default of 100. (optional) (default to 100)
	offset := int32(56) // int32 | Zero-based index of the first matching item to return. Increase it with `limit` to request later pages. If `limit` is provided without `offset`, the offset defaults to 0. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesExplorerAPI.ListKubernetesDeployments(context.Background()).ClusterId(clusterId).Search(search).SortBy(sortBy).SortDirection(sortDirection).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesExplorerAPI.ListKubernetesDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKubernetesDeployments`: ListKubernetesDeploymentsResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesExplorerAPI.ListKubernetesDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKubernetesDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. | 
 **search** | **string** | Substring matched against the resource name. Kubernetes names are lowercase RFC 1123, so only lowercase letters, digits, &#x60;.&#x60; and &#x60;-&#x60; are accepted; uppercase and other characters are rejected. Wildcard and other query-syntax characters are treated literally. | 
 **sortBy** | **string** | Field to sort results by. | 
 **sortDirection** | **string** | Sort direction: ascending or descending. | 
 **limit** | **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [default to 100]
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 

### Return type

[**ListKubernetesDeploymentsResponse**](ListKubernetesDeploymentsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKubernetesEvents

> ListKubernetesEventsResponse ListKubernetesEvents(ctx).ClusterId(clusterId).Search(search).IncludeInternalNamespaces(includeInternalNamespaces).Clusters(clusters).Namespaces(namespaces).ResourceKinds(resourceKinds).Types(types).Limit(limit).Execute()





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
	clusterId := "clusterId_example" // string | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. (optional)
	search := "search_example" // string | Substring matched against the resource name. Kubernetes names are lowercase RFC 1123, so only lowercase letters, digits, `.` and `-` are accepted; uppercase and other characters are rejected. Wildcard and other query-syntax characters are treated literally. (optional)
	includeInternalNamespaces := true // bool | When false or omitted, resources in internal namespaces (kube-system, kube-public, kube-node-lease, local-path-storage, gke-managed-*) are excluded, even when `namespaces` names them. (optional)
	clusters := []*string{"Inner_example"} // []*string | Restrict results to these Kubernetes cluster names. (optional)
	namespaces := []string{"Inner_example"} // []string | Restrict results to these namespaces. (optional)
	resourceKinds := []string{"Inner_example"} // []string | Restrict results to events about these resource kinds (Pod, Node, Deployment...). (optional)
	types := []string{"Types_example"} // []string | Restrict results to events of these types. (optional)
	limit := int32(56) // int32 | Maximum number of events to return, newest first. Defaults to 100. Valid values are 1 through 1000. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesExplorerAPI.ListKubernetesEvents(context.Background()).ClusterId(clusterId).Search(search).IncludeInternalNamespaces(includeInternalNamespaces).Clusters(clusters).Namespaces(namespaces).ResourceKinds(resourceKinds).Types(types).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesExplorerAPI.ListKubernetesEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKubernetesEvents`: ListKubernetesEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesExplorerAPI.ListKubernetesEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKubernetesEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. | 
 **search** | **string** | Substring matched against the resource name. Kubernetes names are lowercase RFC 1123, so only lowercase letters, digits, &#x60;.&#x60; and &#x60;-&#x60; are accepted; uppercase and other characters are rejected. Wildcard and other query-syntax characters are treated literally. | 
 **includeInternalNamespaces** | **bool** | When false or omitted, resources in internal namespaces (kube-system, kube-public, kube-node-lease, local-path-storage, gke-managed-*) are excluded, even when &#x60;namespaces&#x60; names them. | 
 **clusters** | **[]string** | Restrict results to these Kubernetes cluster names. | 
 **namespaces** | **[]string** | Restrict results to these namespaces. | 
 **resourceKinds** | **[]string** | Restrict results to events about these resource kinds (Pod, Node, Deployment...). | 
 **types** | **[]string** | Restrict results to events of these types. | 
 **limit** | **int32** | Maximum number of events to return, newest first. Defaults to 100. Valid values are 1 through 1000. | [default to 100]

### Return type

[**ListKubernetesEventsResponse**](ListKubernetesEventsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKubernetesNamespaces

> ListKubernetesNamespacesResponse ListKubernetesNamespaces(ctx).ClusterId(clusterId).Search(search).SortBy(sortBy).SortDirection(sortDirection).Limit(limit).Offset(offset).Execute()





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
	clusterId := "clusterId_example" // string | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. (optional)
	search := "search_example" // string | Substring matched against the resource name. Kubernetes names are lowercase RFC 1123, so only lowercase letters, digits, `.` and `-` are accepted; uppercase and other characters are rejected. Wildcard and other query-syntax characters are treated literally. (optional)
	sortBy := "sortBy_example" // string | Field to sort results by. (optional)
	sortDirection := "sortDirection_example" // string | Sort direction: ascending or descending. (optional)
	limit := int32(56) // int32 | Maximum number of items to return in this page. Omit to use the public API default of 100. (optional) (default to 100)
	offset := int32(56) // int32 | Zero-based index of the first matching item to return. Increase it with `limit` to request later pages. If `limit` is provided without `offset`, the offset defaults to 0. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesExplorerAPI.ListKubernetesNamespaces(context.Background()).ClusterId(clusterId).Search(search).SortBy(sortBy).SortDirection(sortDirection).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesExplorerAPI.ListKubernetesNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKubernetesNamespaces`: ListKubernetesNamespacesResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesExplorerAPI.ListKubernetesNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKubernetesNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. | 
 **search** | **string** | Substring matched against the resource name. Kubernetes names are lowercase RFC 1123, so only lowercase letters, digits, &#x60;.&#x60; and &#x60;-&#x60; are accepted; uppercase and other characters are rejected. Wildcard and other query-syntax characters are treated literally. | 
 **sortBy** | **string** | Field to sort results by. | 
 **sortDirection** | **string** | Sort direction: ascending or descending. | 
 **limit** | **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [default to 100]
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 

### Return type

[**ListKubernetesNamespacesResponse**](ListKubernetesNamespacesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKubernetesNodes

> ListKubernetesNodesResponse ListKubernetesNodes(ctx).ClusterId(clusterId).Search(search).SortBy(sortBy).SortDirection(sortDirection).Limit(limit).Offset(offset).Execute()





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
	clusterId := "clusterId_example" // string | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. (optional)
	search := "search_example" // string | Substring matched against the resource name. Kubernetes names are lowercase RFC 1123, so only lowercase letters, digits, `.` and `-` are accepted; uppercase and other characters are rejected. Wildcard and other query-syntax characters are treated literally. (optional)
	sortBy := "sortBy_example" // string | Field to sort results by. (optional)
	sortDirection := "sortDirection_example" // string | Sort direction: ascending or descending. (optional)
	limit := int32(56) // int32 | Maximum number of items to return in this page. Omit to use the public API default of 100. (optional) (default to 100)
	offset := int32(56) // int32 | Zero-based index of the first matching item to return. Increase it with `limit` to request later pages. If `limit` is provided without `offset`, the offset defaults to 0. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesExplorerAPI.ListKubernetesNodes(context.Background()).ClusterId(clusterId).Search(search).SortBy(sortBy).SortDirection(sortDirection).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesExplorerAPI.ListKubernetesNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKubernetesNodes`: ListKubernetesNodesResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesExplorerAPI.ListKubernetesNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKubernetesNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. | 
 **search** | **string** | Substring matched against the resource name. Kubernetes names are lowercase RFC 1123, so only lowercase letters, digits, &#x60;.&#x60; and &#x60;-&#x60; are accepted; uppercase and other characters are rejected. Wildcard and other query-syntax characters are treated literally. | 
 **sortBy** | **string** | Field to sort results by. | 
 **sortDirection** | **string** | Sort direction: ascending or descending. | 
 **limit** | **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [default to 100]
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 

### Return type

[**ListKubernetesNodesResponse**](ListKubernetesNodesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKubernetesPods

> ListKubernetesPodsResponse ListKubernetesPods(ctx).ClusterId(clusterId).Search(search).SortBy(sortBy).SortDirection(sortDirection).Limit(limit).Offset(offset).Execute()





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
	clusterId := "clusterId_example" // string | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. (optional)
	search := "search_example" // string | Substring matched against the resource name. Kubernetes names are lowercase RFC 1123, so only lowercase letters, digits, `.` and `-` are accepted; uppercase and other characters are rejected. Wildcard and other query-syntax characters are treated literally. (optional)
	sortBy := "sortBy_example" // string | Field to sort results by. (optional)
	sortDirection := "sortDirection_example" // string | Sort direction: ascending or descending. (optional)
	limit := int32(56) // int32 | Maximum number of items to return in this page. Omit to use the public API default of 100. (optional) (default to 100)
	offset := int32(56) // int32 | Zero-based index of the first matching item to return. Increase it with `limit` to request later pages. If `limit` is provided without `offset`, the offset defaults to 0. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesExplorerAPI.ListKubernetesPods(context.Background()).ClusterId(clusterId).Search(search).SortBy(sortBy).SortDirection(sortDirection).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesExplorerAPI.ListKubernetesPods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKubernetesPods`: ListKubernetesPodsResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesExplorerAPI.ListKubernetesPods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKubernetesPodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. | 
 **search** | **string** | Substring matched against the resource name. Kubernetes names are lowercase RFC 1123, so only lowercase letters, digits, &#x60;.&#x60; and &#x60;-&#x60; are accepted; uppercase and other characters are rejected. Wildcard and other query-syntax characters are treated literally. | 
 **sortBy** | **string** | Field to sort results by. | 
 **sortDirection** | **string** | Sort direction: ascending or descending. | 
 **limit** | **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [default to 100]
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 

### Return type

[**ListKubernetesPodsResponse**](ListKubernetesPodsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKubernetesStatefulSets

> ListKubernetesStatefulSetsResponse ListKubernetesStatefulSets(ctx).ClusterId(clusterId).Search(search).SortBy(sortBy).SortDirection(sortDirection).Limit(limit).Offset(offset).Execute()





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
	clusterId := "clusterId_example" // string | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. (optional)
	search := "search_example" // string | Substring matched against the resource name. Kubernetes names are lowercase RFC 1123, so only lowercase letters, digits, `.` and `-` are accepted; uppercase and other characters are rejected. Wildcard and other query-syntax characters are treated literally. (optional)
	sortBy := "sortBy_example" // string | Field to sort results by. (optional)
	sortDirection := "sortDirection_example" // string | Sort direction: ascending or descending. (optional)
	limit := int32(56) // int32 | Maximum number of items to return in this page. Omit to use the public API default of 100. (optional) (default to 100)
	offset := int32(56) // int32 | Zero-based index of the first matching item to return. Increase it with `limit` to request later pages. If `limit` is provided without `offset`, the offset defaults to 0. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesExplorerAPI.ListKubernetesStatefulSets(context.Background()).ClusterId(clusterId).Search(search).SortBy(sortBy).SortDirection(sortDirection).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesExplorerAPI.ListKubernetesStatefulSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKubernetesStatefulSets`: ListKubernetesStatefulSetsResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesExplorerAPI.ListKubernetesStatefulSets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKubernetesStatefulSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Tsuga cluster ID to target. Tsuga assigns it when a cluster is connected to the organization. Optional for single-cluster organizations and required for organizations with multiple clusters to scope the request to one cluster. | 
 **search** | **string** | Substring matched against the resource name. Kubernetes names are lowercase RFC 1123, so only lowercase letters, digits, &#x60;.&#x60; and &#x60;-&#x60; are accepted; uppercase and other characters are rejected. Wildcard and other query-syntax characters are treated literally. | 
 **sortBy** | **string** | Field to sort results by. | 
 **sortDirection** | **string** | Sort direction: ascending or descending. | 
 **limit** | **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [default to 100]
 **offset** | **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | 

### Return type

[**ListKubernetesStatefulSetsResponse**](ListKubernetesStatefulSetsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

