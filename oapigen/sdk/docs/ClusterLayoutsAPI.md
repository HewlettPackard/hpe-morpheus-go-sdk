# \ClusterLayoutsAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddClusterLayouts**](ClusterLayoutsAPI.md#AddClusterLayouts) | **Post** /api/library/cluster-layouts | Create a Cluster Layout
[**ListClusterLayouts**](ClusterLayoutsAPI.md#ListClusterLayouts) | **Get** /api/library/cluster-layouts | Get All Cluster Layouts



## AddClusterLayouts

> AddClusterLayouts200Response AddClusterLayouts(ctx).AddClusterLayoutsRequest(addClusterLayoutsRequest).Execute()

Create a Cluster Layout



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	addClusterLayoutsRequest := *openapiclient.NewAddClusterLayoutsRequest() // AddClusterLayoutsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterLayoutsAPI.AddClusterLayouts(context.Background()).AddClusterLayoutsRequest(addClusterLayoutsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterLayoutsAPI.AddClusterLayouts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddClusterLayouts`: AddClusterLayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterLayoutsAPI.AddClusterLayouts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddClusterLayoutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addClusterLayoutsRequest** | [**AddClusterLayoutsRequest**](AddClusterLayoutsRequest.md) |  | 

### Return type

[**AddClusterLayouts200Response**](AddClusterLayouts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterLayouts

> ListClusterLayouts200Response ListClusterLayouts(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).ProvisionType(provisionType).Labels(labels).AllLabels(allLabels).Execute()

Get All Cluster Layouts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
	sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
	direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	provisionType := "provisionType_example" // string | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.  (optional)
	labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
	allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterLayoutsAPI.ListClusterLayouts(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).ProvisionType(provisionType).Labels(labels).AllLabels(allLabels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterLayoutsAPI.ListClusterLayouts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusterLayouts`: ListClusterLayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterLayoutsAPI.ListClusterLayouts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClusterLayoutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **provisionType** | **string** | Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 

### Return type

[**ListClusterLayouts200Response**](ListClusterLayouts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

