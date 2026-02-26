# \NetworksAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworks**](NetworksAPI.md#CreateNetworks) | **Post** /api/networks | Create a Network
[**DeleteNetwork**](NetworksAPI.md#DeleteNetwork) | **Delete** /api/networks/{id} | Delete a Network
[**GetNetwork**](NetworksAPI.md#GetNetwork) | **Get** /api/networks/{id} | Get a Specific Network
[**ListNetworks**](NetworksAPI.md#ListNetworks) | **Get** /api/networks | Get All Networks
[**UpdateNetwork**](NetworksAPI.md#UpdateNetwork) | **Put** /api/networks/{id} | Update a Network



## CreateNetworks

> CreateNetworks200Response CreateNetworks(ctx).CreateNetworksRequest(createNetworksRequest).Execute()

Create a Network



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
	createNetworksRequest := *openapiclient.NewCreateNetworksRequest() // CreateNetworksRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworks(context.Background()).CreateNetworksRequest(createNetworksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworks`: CreateNetworks200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworksRequest** | [**CreateNetworksRequest**](CreateNetworksRequest.md) |  | 

### Return type

[**CreateNetworks200Response**](CreateNetworks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetwork

> DeleteNetwork200Response DeleteNetwork(ctx, id).Execute()

Delete a Network



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
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetwork(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetwork`: DeleteNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteNetwork200Response**](DeleteNetwork200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetwork

> GetNetwork200Response GetNetwork(ctx, id).Execute()

Get a Specific Network



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
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetwork(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetwork`: GetNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetwork200Response**](GetNetwork200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworks

> ListNetworks200Response ListNetworks(ctx).Name(name).Phrase(phrase).Labels(labels).AllLabels(allLabels).ZoneId(zoneId).Cidr(cidr).ClusterId(clusterId).Execute()

Get All Networks



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
	name := "example" // string | Filter by name (optional)
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
	allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)
	zoneId := int64(3) // int64 | The Cloud ID (Zone ID) for Filtering (optional)
	cidr := "cidr_example" // string | Filter by cidr, matches beginning of value. (optional)
	clusterId := int64(135) // int64 | The Cluster ID for filtering. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.ListNetworks(context.Background()).Name(name).Phrase(phrase).Labels(labels).AllLabels(allLabels).ZoneId(zoneId).Cidr(cidr).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.ListNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworks`: ListNetworks200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.ListNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 
 **zoneId** | **int64** | The Cloud ID (Zone ID) for Filtering | 
 **cidr** | **string** | Filter by cidr, matches beginning of value. | 
 **clusterId** | **int64** | The Cluster ID for filtering. | 

### Return type

[**ListNetworks200Response**](ListNetworks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetwork

> UpdateNetwork200Response UpdateNetwork(ctx, id).UpdateNetworkRequest(updateNetworkRequest).Execute()

Update a Network



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
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateNetworkRequest := *openapiclient.NewUpdateNetworkRequest() // UpdateNetworkRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetwork(context.Background(), id).UpdateNetworkRequest(updateNetworkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetwork`: UpdateNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkRequest** | [**UpdateNetworkRequest**](UpdateNetworkRequest.md) |  | 

### Return type

[**UpdateNetwork200Response**](UpdateNetwork200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

