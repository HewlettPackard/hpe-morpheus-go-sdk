# \CloudsAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCloudResourcePool**](CloudsAPI.md#AddCloudResourcePool) | **Post** /api/zones/{zoneId}/resource-pools | Creates a Specified Resource Pool for Specified Cloud
[**AddClouds**](CloudsAPI.md#AddClouds) | **Post** /api/zones | Creates a Cloud
[**GetCloudDatastores**](CloudsAPI.md#GetCloudDatastores) | **Get** /api/zones/{zoneId}/data-stores/{id} | Retrieves a Datastore for Specified Cloud
[**GetCloudFolders**](CloudsAPI.md#GetCloudFolders) | **Get** /api/zones/{zoneId}/folders/{id} | Retrieves a Resource Folder for Specified Cloud
[**GetCloudResourcePools**](CloudsAPI.md#GetCloudResourcePools) | **Get** /api/zones/{zoneId}/resource-pools/{id} | Retrieves a Resource Pool for Specified Cloud
[**GetCloudTypes**](CloudsAPI.md#GetCloudTypes) | **Get** /api/zone-types/{id} | Retrieves a Specific Cloud Type
[**GetClouds**](CloudsAPI.md#GetClouds) | **Get** /api/zones/{id} | Retrieves a Specific Cloud
[**ListCloudDatastores**](CloudsAPI.md#ListCloudDatastores) | **Get** /api/zones/{zoneId}/data-stores | Retrieves all Datastores for Specified Cloud
[**ListCloudFolders**](CloudsAPI.md#ListCloudFolders) | **Get** /api/zones/{zoneId}/folders | Retrieves all resource folders for Specified Cloud
[**ListCloudResourcePools**](CloudsAPI.md#ListCloudResourcePools) | **Get** /api/zones/{zoneId}/resource-pools | Retrieves all Resource Pools for Specified Cloud
[**ListCloudSecurityGroups**](CloudsAPI.md#ListCloudSecurityGroups) | **Get** /api/zones/{id}/security-groups | Retrieves all Security Groups for a Cloud
[**ListCloudTypes**](CloudsAPI.md#ListCloudTypes) | **Get** /api/zone-types | Retrieves all Cloud Types
[**ListClouds**](CloudsAPI.md#ListClouds) | **Get** /api/zones | Retrieves all Clouds
[**RefreshClouds**](CloudsAPI.md#RefreshClouds) | **Post** /api/zones/{id}/refresh | Refreshes a Cloud
[**RemoveCloudResourcePools**](CloudsAPI.md#RemoveCloudResourcePools) | **Delete** /api/zones/{zoneId}/resource-pools/{id} | Deletes a Resource Pool for Specified Cloud
[**RemoveClouds**](CloudsAPI.md#RemoveClouds) | **Delete** /api/zones/{id} | Deletes a Cloud
[**UpdateCloudDatastores**](CloudsAPI.md#UpdateCloudDatastores) | **Put** /api/zones/{zoneId}/data-stores/{id} | Updates a Specified Datastore for Specified Cloud
[**UpdateCloudFolders**](CloudsAPI.md#UpdateCloudFolders) | **Put** /api/zones/{zoneId}/folders/{id} | Updates a Resource Folder for Specified Cloud
[**UpdateCloudLogo**](CloudsAPI.md#UpdateCloudLogo) | **Post** /api/zones/{id}/update-logo | Update Logo For Cloud
[**UpdateCloudResourcePool**](CloudsAPI.md#UpdateCloudResourcePool) | **Put** /api/zones/{zoneId}/resource-pools/{id} | Updates a Specified Resource Pool for Specified Cloud
[**UpdateCloudSecurityGroups**](CloudsAPI.md#UpdateCloudSecurityGroups) | **Post** /api/zones/{id}/security-groups | Sets Security Groups for a Cloud
[**UpdateClouds**](CloudsAPI.md#UpdateClouds) | **Put** /api/zones/{id} | Updates a Cloud



## AddCloudResourcePool

> AddCloudResourcePool200Response AddCloudResourcePool(ctx, zoneId).AddCloudResourcePoolRequest(addCloudResourcePoolRequest).Execute()

Creates a Specified Resource Pool for Specified Cloud



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	zoneId := float32(7) // float32 | The ID of the cloud
	addCloudResourcePoolRequest := *openapiclient.NewAddCloudResourcePoolRequest(*openapiclient.NewAddCloudResourcePoolRequestResourcePool("Name_example", *openapiclient.NewAddCloudResourcePoolRequestResourcePoolConfig())) // AddCloudResourcePoolRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.AddCloudResourcePool(context.Background(), zoneId).AddCloudResourcePoolRequest(addCloudResourcePoolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.AddCloudResourcePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddCloudResourcePool`: AddCloudResourcePool200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.AddCloudResourcePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCloudResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addCloudResourcePoolRequest** | [**AddCloudResourcePoolRequest**](AddCloudResourcePoolRequest.md) |  | 

### Return type

[**AddCloudResourcePool200Response**](AddCloudResourcePool200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddClouds

> AddClouds200Response AddClouds(ctx).AddCloudsRequest(addCloudsRequest).Execute()

Creates a Cloud



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	addCloudsRequest := *openapiclient.NewAddCloudsRequest(*openapiclient.NewAddCloudsRequestZone("My Cloud", *openapiclient.NewAddCloudsRequestZoneZoneType(), int64(3))) // AddCloudsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.AddClouds(context.Background()).AddCloudsRequest(addCloudsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.AddClouds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddClouds`: AddClouds200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.AddClouds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCloudsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addCloudsRequest** | [**AddCloudsRequest**](AddCloudsRequest.md) |  | 

### Return type

[**AddClouds200Response**](AddClouds200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudDatastores

> GetCloudDatastores200Response GetCloudDatastores(ctx, zoneId, id).Execute()

Retrieves a Datastore for Specified Cloud



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	zoneId := float32(7) // float32 | The ID of the cloud
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.GetCloudDatastores(context.Background(), zoneId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.GetCloudDatastores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudDatastores`: GetCloudDatastores200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.GetCloudDatastores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudDatastoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetCloudDatastores200Response**](GetCloudDatastores200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudFolders

> GetCloudFolders200Response GetCloudFolders(ctx, zoneId, id).Execute()

Retrieves a Resource Folder for Specified Cloud



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	zoneId := float32(7) // float32 | The ID of the cloud
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.GetCloudFolders(context.Background(), zoneId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.GetCloudFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudFolders`: GetCloudFolders200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.GetCloudFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetCloudFolders200Response**](GetCloudFolders200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudResourcePools

> GetCloudResourcePools200Response GetCloudResourcePools(ctx, zoneId, id).Execute()

Retrieves a Resource Pool for Specified Cloud



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	zoneId := float32(7) // float32 | The ID of the cloud
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.GetCloudResourcePools(context.Background(), zoneId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.GetCloudResourcePools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudResourcePools`: GetCloudResourcePools200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.GetCloudResourcePools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudResourcePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetCloudResourcePools200Response**](GetCloudResourcePools200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudTypes

> GetCloudTypes200Response GetCloudTypes(ctx, id).Execute()

Retrieves a Specific Cloud Type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.GetCloudTypes(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.GetCloudTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudTypes`: GetCloudTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.GetCloudTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCloudTypes200Response**](GetCloudTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClouds

> GetClouds200Response GetClouds(ctx, id).Execute()

Retrieves a Specific Cloud



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.GetClouds(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.GetClouds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClouds`: GetClouds200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.GetClouds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetClouds200Response**](GetClouds200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudDatastores

> ListCloudDatastores200Response ListCloudDatastores(ctx, zoneId).Name(name).Phrase(phrase).Max(max).Sort(sort).Direction(direction).Execute()

Retrieves all Datastores for Specified Cloud



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	zoneId := float32(7) // float32 | The ID of the cloud
	name := "example" // string | Filter by name (optional)
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
	direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.ListCloudDatastores(context.Background(), zoneId).Name(name).Phrase(phrase).Max(max).Sort(sort).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.ListCloudDatastores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCloudDatastores`: ListCloudDatastores200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.ListCloudDatastores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCloudDatastoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Filter by name | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]

### Return type

[**ListCloudDatastores200Response**](ListCloudDatastores200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudFolders

> ListCloudFolders200Response ListCloudFolders(ctx, zoneId).Name(name).Phrase(phrase).Max(max).Execute()

Retrieves all resource folders for Specified Cloud



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	zoneId := float32(7) // float32 | The ID of the cloud
	name := "example" // string | Filter by name (optional)
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.ListCloudFolders(context.Background(), zoneId).Name(name).Phrase(phrase).Max(max).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.ListCloudFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCloudFolders`: ListCloudFolders200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.ListCloudFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCloudFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Filter by name | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **max** | **int64** | Maximum number of records to return | [default to 25]

### Return type

[**ListCloudFolders200Response**](ListCloudFolders200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudResourcePools

> ListCloudResourcePools200Response ListCloudResourcePools(ctx, zoneId).Name(name).Phrase(phrase).Max(max).Execute()

Retrieves all Resource Pools for Specified Cloud



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	zoneId := float32(7) // float32 | The ID of the cloud
	name := "example" // string | Filter by name (optional)
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.ListCloudResourcePools(context.Background(), zoneId).Name(name).Phrase(phrase).Max(max).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.ListCloudResourcePools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCloudResourcePools`: ListCloudResourcePools200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.ListCloudResourcePools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCloudResourcePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Filter by name | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **max** | **int64** | Maximum number of records to return | [default to 25]

### Return type

[**ListCloudResourcePools200Response**](ListCloudResourcePools200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudSecurityGroups

> ListCloudSecurityGroups200Response ListCloudSecurityGroups(ctx, id).Execute()

Retrieves all Security Groups for a Cloud



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.ListCloudSecurityGroups(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.ListCloudSecurityGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCloudSecurityGroups`: ListCloudSecurityGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.ListCloudSecurityGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCloudSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListCloudSecurityGroups200Response**](ListCloudSecurityGroups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudTypes

> ListCloudTypes200Response ListCloudTypes(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Name(name).Code(code).Phrase(phrase).ProvisionType(provisionType).Execute()

Retrieves all Cloud Types



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
	sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
	direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
	name := "example" // string | Filter by name (optional)
	code := "azr" // string | If specified will return an exact match on code (optional)
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	provisionType := "provisionType_example" // string | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.ListCloudTypes(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Name(name).Code(code).Phrase(phrase).ProvisionType(provisionType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.ListCloudTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCloudTypes`: ListCloudTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.ListCloudTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCloudTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **name** | **string** | Filter by name | 
 **code** | **string** | If specified will return an exact match on code | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **provisionType** | **string** | Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | 

### Return type

[**ListCloudTypes200Response**](ListCloudTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClouds

> ListClouds200Response ListClouds(ctx).LastUpdated(lastUpdated).Type_(type_).GroupId(groupId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Labels(labels).Execute()

Retrieves all Clouds



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
	type_ := "type__example" // string | If specified will return all zones by cloud type code. Refer to `Zone Types` API for up to date listings.  (optional)
	groupId := int64(13) // int64 | If specified will return all zones assigned to a server group by id. Accepts multiple values. (optional)
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
	sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
	direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	name := "example" // string | Filter by name (optional)
	labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.ListClouds(context.Background()).LastUpdated(lastUpdated).Type_(type_).GroupId(groupId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Labels(labels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.ListClouds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClouds`: ListClouds200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.ListClouds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCloudsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 
 **type_** | **string** | If specified will return all zones by cloud type code. Refer to &#x60;Zone Types&#x60; API for up to date listings.  | 
 **groupId** | **int64** | If specified will return all zones assigned to a server group by id. Accepts multiple values. | 
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 

### Return type

[**ListClouds200Response**](ListClouds200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshClouds

> DeleteAlerts200Response RefreshClouds(ctx, id).RefreshCloudsRequest(refreshCloudsRequest).Execute()

Refreshes a Cloud



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	refreshCloudsRequest := *openapiclient.NewRefreshCloudsRequest() // RefreshCloudsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.RefreshClouds(context.Background(), id).RefreshCloudsRequest(refreshCloudsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.RefreshClouds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshClouds`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.RefreshClouds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshCloudsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refreshCloudsRequest** | [**RefreshCloudsRequest**](RefreshCloudsRequest.md) |  | 

### Return type

[**DeleteAlerts200Response**](DeleteAlerts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCloudResourcePools

> DeleteAlerts200Response RemoveCloudResourcePools(ctx, zoneId, id).Execute()

Deletes a Resource Pool for Specified Cloud



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	zoneId := float32(7) // float32 | The ID of the cloud
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.RemoveCloudResourcePools(context.Background(), zoneId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.RemoveCloudResourcePools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveCloudResourcePools`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.RemoveCloudResourcePools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCloudResourcePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteAlerts200Response**](DeleteAlerts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveClouds

> DeleteAlerts200Response RemoveClouds(ctx, id).RemoveResources(removeResources).Execute()

Deletes a Cloud



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	removeResources := true // bool | Removing associated resources will delete the instances and the associated resources underneath.  This includes Virtual Machines and other forms of compute. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.RemoveClouds(context.Background(), id).RemoveResources(removeResources).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.RemoveClouds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveClouds`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.RemoveClouds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCloudsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeResources** | **bool** | Removing associated resources will delete the instances and the associated resources underneath.  This includes Virtual Machines and other forms of compute. | [default to false]

### Return type

[**DeleteAlerts200Response**](DeleteAlerts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudDatastores

> UpdateCloudDatastores200Response UpdateCloudDatastores(ctx, zoneId, id).UpdateCloudDatastoresRequest(updateCloudDatastoresRequest).Execute()

Updates a Specified Datastore for Specified Cloud



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	zoneId := float32(7) // float32 | The ID of the cloud
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateCloudDatastoresRequest := *openapiclient.NewUpdateCloudDatastoresRequest(*openapiclient.NewUpdateCloudDatastoresRequestDatastore()) // UpdateCloudDatastoresRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.UpdateCloudDatastores(context.Background(), zoneId, id).UpdateCloudDatastoresRequest(updateCloudDatastoresRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.UpdateCloudDatastores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCloudDatastores`: UpdateCloudDatastores200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.UpdateCloudDatastores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudDatastoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCloudDatastoresRequest** | [**UpdateCloudDatastoresRequest**](UpdateCloudDatastoresRequest.md) |  | 

### Return type

[**UpdateCloudDatastores200Response**](UpdateCloudDatastores200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudFolders

> UpdateCloudFolders200Response UpdateCloudFolders(ctx, zoneId, id).UpdateCloudFoldersRequest(updateCloudFoldersRequest).Execute()

Updates a Resource Folder for Specified Cloud



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	zoneId := float32(7) // float32 | The ID of the cloud
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateCloudFoldersRequest := *openapiclient.NewUpdateCloudFoldersRequest(*openapiclient.NewUpdateCloudFoldersRequestFolder()) // UpdateCloudFoldersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.UpdateCloudFolders(context.Background(), zoneId, id).UpdateCloudFoldersRequest(updateCloudFoldersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.UpdateCloudFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCloudFolders`: UpdateCloudFolders200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.UpdateCloudFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCloudFoldersRequest** | [**UpdateCloudFoldersRequest**](UpdateCloudFoldersRequest.md) |  | 

### Return type

[**UpdateCloudFolders200Response**](UpdateCloudFolders200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudLogo

> DeleteAlerts200Response UpdateCloudLogo(ctx, id).Logo(logo).DarkLogo(darkLogo).Execute()

Update Logo For Cloud



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	logo := os.NewFile(1234, "some_file") // *os.File | Logo File png,jpg,svg (optional)
	darkLogo := os.NewFile(1234, "some_file") // *os.File | Dark Logo File png,jpg,svg (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.UpdateCloudLogo(context.Background(), id).Logo(logo).DarkLogo(darkLogo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.UpdateCloudLogo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCloudLogo`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.UpdateCloudLogo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logo** | ***os.File** | Logo File png,jpg,svg | 
 **darkLogo** | ***os.File** | Dark Logo File png,jpg,svg | 

### Return type

[**DeleteAlerts200Response**](DeleteAlerts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudResourcePool

> AddCloudResourcePool200Response UpdateCloudResourcePool(ctx, zoneId, id).UpdateCloudResourcePoolRequest(updateCloudResourcePoolRequest).Execute()

Updates a Specified Resource Pool for Specified Cloud



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	zoneId := float32(7) // float32 | The ID of the cloud
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateCloudResourcePoolRequest := *openapiclient.NewUpdateCloudResourcePoolRequest(*openapiclient.NewUpdateCloudResourcePoolRequestResourcePool()) // UpdateCloudResourcePoolRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.UpdateCloudResourcePool(context.Background(), zoneId, id).UpdateCloudResourcePoolRequest(updateCloudResourcePoolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.UpdateCloudResourcePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCloudResourcePool`: AddCloudResourcePool200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.UpdateCloudResourcePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **float32** | The ID of the cloud | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCloudResourcePoolRequest** | [**UpdateCloudResourcePoolRequest**](UpdateCloudResourcePoolRequest.md) |  | 

### Return type

[**AddCloudResourcePool200Response**](AddCloudResourcePool200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudSecurityGroups

> UpdateCloudSecurityGroups200Response UpdateCloudSecurityGroups(ctx, id).UpdateCloudSecurityGroupsRequest(updateCloudSecurityGroupsRequest).Execute()

Sets Security Groups for a Cloud



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateCloudSecurityGroupsRequest := *openapiclient.NewUpdateCloudSecurityGroupsRequest([]int64{int64(123)}) // UpdateCloudSecurityGroupsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.UpdateCloudSecurityGroups(context.Background(), id).UpdateCloudSecurityGroupsRequest(updateCloudSecurityGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.UpdateCloudSecurityGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCloudSecurityGroups`: UpdateCloudSecurityGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.UpdateCloudSecurityGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCloudSecurityGroupsRequest** | [**UpdateCloudSecurityGroupsRequest**](UpdateCloudSecurityGroupsRequest.md) |  | 

### Return type

[**UpdateCloudSecurityGroups200Response**](UpdateCloudSecurityGroups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClouds

> AddClouds200Response UpdateClouds(ctx, id).UpdateCloudsRequest(updateCloudsRequest).Execute()

Updates a Cloud



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateCloudsRequest := *openapiclient.NewUpdateCloudsRequest(*openapiclient.NewUpdateCloudsRequestZone("My Cloud", "ZoneType_example", int64(3), map[string]interface{}({"id":1}))) // UpdateCloudsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.UpdateClouds(context.Background(), id).UpdateCloudsRequest(updateCloudsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.UpdateClouds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClouds`: AddClouds200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.UpdateClouds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCloudsRequest** | [**UpdateCloudsRequest**](UpdateCloudsRequest.md) |  | 

### Return type

[**AddClouds200Response**](AddClouds200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

