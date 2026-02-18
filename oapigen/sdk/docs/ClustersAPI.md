# \ClustersAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCluster**](ClustersAPI.md#AddCluster) | **Post** /api/clusters | Create a Cluster
[**DeleteClusterDatastore**](ClustersAPI.md#DeleteClusterDatastore) | **Delete** /api/clusters/{clusterId}/datastores/{id} | Delete a Cluster Datastore
[**GetClusterDatastore**](ClustersAPI.md#GetClusterDatastore) | **Get** /api/clusters/{clusterId}/datastores/{id} | Get a Specific Cluster Datastore
[**ListClusterDatastores**](ClustersAPI.md#ListClusterDatastores) | **Get** /api/clusters/{clusterId}/datastores | Get Cluster Datastores
[**ListClusters**](ClustersAPI.md#ListClusters) | **Get** /api/clusters | Get All Clusters
[**SaveClusterDatastore**](ClustersAPI.md#SaveClusterDatastore) | **Post** /api/clusters/{clusterId}/datastores | Create a Cluster Datastore
[**UpdateClusterDatastore**](ClustersAPI.md#UpdateClusterDatastore) | **Put** /api/clusters/{clusterId}/datastores/{id} | Update Cluster Datastore



## AddCluster

> AddCluster200Response AddCluster(ctx).AddClusterRequest(addClusterRequest).Execute()

Create a Cluster



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
	addClusterRequest := *openapiclient.NewAddClusterRequest() // AddClusterRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.AddCluster(context.Background()).AddClusterRequest(addClusterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.AddCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddCluster`: AddCluster200Response
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.AddCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addClusterRequest** | [**AddClusterRequest**](AddClusterRequest.md) |  | 

### Return type

[**AddCluster200Response**](AddCluster200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClusterDatastore

> DeleteClusterDatastore200Response DeleteClusterDatastore(ctx, clusterId, id).Execute()

Delete a Cluster Datastore



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
	clusterId := int64(5) // int64 | The ID of the cluster
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.DeleteClusterDatastore(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.DeleteClusterDatastore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteClusterDatastore`: DeleteClusterDatastore200Response
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.DeleteClusterDatastore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int64** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterDatastoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteClusterDatastore200Response**](DeleteClusterDatastore200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterDatastore

> GetClusterDatastore200Response GetClusterDatastore(ctx, clusterId, id).Execute()

Get a Specific Cluster Datastore



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
	clusterId := int64(5) // int64 | The ID of the cluster
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetClusterDatastore(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClusterDatastore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterDatastore`: GetClusterDatastore200Response
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClusterDatastore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int64** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterDatastoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetClusterDatastore200Response**](GetClusterDatastore200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterDatastores

> ListClusterDatastores200Response ListClusterDatastores(ctx, clusterId).Max(max).Offset(offset).Sort(sort).Order(order).Phrase(phrase).Name(name).Code(code).HideInactive(hideInactive).Execute()

Get Cluster Datastores



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
	clusterId := int64(5) // int64 | The ID of the cluster
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
	sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
	order := "order_example" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	name := "example" // string | Filter by name (optional)
	code := "azr" // string | If specified will return an exact match on code (optional)
	hideInactive := true // bool | If true restricts query to only load active datastores (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.ListClusterDatastores(context.Background(), clusterId).Max(max).Offset(offset).Sort(sort).Order(order).Phrase(phrase).Name(name).Code(code).HideInactive(hideInactive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.ListClusterDatastores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusterDatastores`: ListClusterDatastores200Response
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.ListClusterDatastores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int64** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterDatastoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **order** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name | 
 **code** | **string** | If specified will return an exact match on code | 
 **hideInactive** | **bool** | If true restricts query to only load active datastores | [default to false]

### Return type

[**ListClusterDatastores200Response**](ListClusterDatastores200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusters

> ListClusters200Response ListClusters(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).ZoneId(zoneId).TypeId(typeId).Labels(labels).AllLabels(allLabels).Execute()

Get All Clusters



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
	name := "example" // string | Filter by name (optional)
	zoneId := int64(3) // int64 | The Cloud ID (Zone ID) for Filtering (optional)
	typeId := int64(3) // int64 | Type filter, restricts query to only load clusters of a specified cluster type (optional)
	labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
	allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.ListClusters(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).ZoneId(zoneId).TypeId(typeId).Labels(labels).AllLabels(allLabels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.ListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusters`: ListClusters200Response
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.ListClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name | 
 **zoneId** | **int64** | The Cloud ID (Zone ID) for Filtering | 
 **typeId** | **int64** | Type filter, restricts query to only load clusters of a specified cluster type | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 

### Return type

[**ListClusters200Response**](ListClusters200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveClusterDatastore

> SaveClusterDatastore200Response SaveClusterDatastore(ctx, clusterId).SaveClusterDatastoreRequest(saveClusterDatastoreRequest).Execute()

Create a Cluster Datastore



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
	clusterId := int64(5) // int64 | The ID of the cluster
	saveClusterDatastoreRequest := *openapiclient.NewSaveClusterDatastoreRequest() // SaveClusterDatastoreRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.SaveClusterDatastore(context.Background(), clusterId).SaveClusterDatastoreRequest(saveClusterDatastoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.SaveClusterDatastore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveClusterDatastore`: SaveClusterDatastore200Response
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.SaveClusterDatastore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int64** | The ID of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveClusterDatastoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **saveClusterDatastoreRequest** | [**SaveClusterDatastoreRequest**](SaveClusterDatastoreRequest.md) |  | 

### Return type

[**SaveClusterDatastore200Response**](SaveClusterDatastore200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClusterDatastore

> UpdateClusterDatastore200Response UpdateClusterDatastore(ctx, clusterId, id).UpdateClusterDatastoreRequest(updateClusterDatastoreRequest).Execute()

Update Cluster Datastore



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
	clusterId := int64(5) // int64 | The ID of the cluster
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateClusterDatastoreRequest := *openapiclient.NewUpdateClusterDatastoreRequest() // UpdateClusterDatastoreRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.UpdateClusterDatastore(context.Background(), clusterId, id).UpdateClusterDatastoreRequest(updateClusterDatastoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.UpdateClusterDatastore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClusterDatastore`: UpdateClusterDatastore200Response
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.UpdateClusterDatastore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int64** | The ID of the cluster | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterDatastoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateClusterDatastoreRequest** | [**UpdateClusterDatastoreRequest**](UpdateClusterDatastoreRequest.md) |  | 

### Return type

[**UpdateClusterDatastore200Response**](UpdateClusterDatastore200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

