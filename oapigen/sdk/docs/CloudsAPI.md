# \CloudsAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddClouds**](CloudsAPI.md#AddClouds) | **Post** /api/zones | Creates a Cloud
[**GetCloudDatastores**](CloudsAPI.md#GetCloudDatastores) | **Get** /api/zones/{zoneId}/data-stores/{id} | Retrieves a Datastore for Specified Cloud
[**GetClouds**](CloudsAPI.md#GetClouds) | **Get** /api/zones/{id} | Retrieves a Specific Cloud
[**ListCloudDatastores**](CloudsAPI.md#ListCloudDatastores) | **Get** /api/zones/{zoneId}/data-stores | Retrieves all Datastores for Specified Cloud
[**ListClouds**](CloudsAPI.md#ListClouds) | **Get** /api/zones | Retrieves all Clouds
[**RemoveClouds**](CloudsAPI.md#RemoveClouds) | **Delete** /api/zones/{id} | Deletes a Cloud
[**SaveCloudDatastore**](CloudsAPI.md#SaveCloudDatastore) | **Post** /api/zones/{zoneId}/data-stores | Create a Datastore for Specified Cloud
[**UpdateCloudDatastores**](CloudsAPI.md#UpdateCloudDatastores) | **Put** /api/zones/{zoneId}/data-stores/{id} | Updates a Specified Datastore for Specified Cloud
[**UpdateClouds**](CloudsAPI.md#UpdateClouds) | **Put** /api/zones/{id} | Updates a Cloud



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
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	addCloudsRequest := *openapiclient.NewAddCloudsRequest(*openapiclient.NewAddCloudsRequestZone("My Cloud", int64(3), *openapiclient.NewAddCloudsRequestZoneZoneType(), *openapiclient.NewAddCloudsRequestZoneConfig("ec2.ca-central-1.amazonaws.com", "https://vcenter.morpheus.local/sdk", "ApiVersion_example", "Datacenter_example"))) // AddCloudsRequest |  (optional)

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
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	zoneId := int64(7) // int64 | The ID of the cloud
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
**zoneId** | **int64** | The ID of the cloud | 
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
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
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
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	zoneId := int64(7) // int64 | The ID of the cloud
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
**zoneId** | **int64** | The ID of the cloud | 

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
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
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


## RemoveClouds

> RemoveClouds200Response RemoveClouds(ctx, id).Force(force).RemoveResources(removeResources).Execute()

Deletes a Cloud



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
	force := true // bool | Force the deletion of the cloud. (optional) (default to false)
	removeResources := true // bool | Removing associated resources will delete the instances and the associated resources underneath.  This includes Virtual Machines and other forms of compute. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.RemoveClouds(context.Background(), id).Force(force).RemoveResources(removeResources).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.RemoveClouds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveClouds`: RemoveClouds200Response
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

 **force** | **bool** | Force the deletion of the cloud. | [default to false]
 **removeResources** | **bool** | Removing associated resources will delete the instances and the associated resources underneath.  This includes Virtual Machines and other forms of compute. | [default to false]

### Return type

[**RemoveClouds200Response**](RemoveClouds200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveCloudDatastore

> SaveCloudDatastore200Response SaveCloudDatastore(ctx, zoneId).SaveCloudDatastoreRequest(saveCloudDatastoreRequest).Execute()

Create a Datastore for Specified Cloud



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
	zoneId := int64(7) // int64 | The ID of the cloud
	saveCloudDatastoreRequest := *openapiclient.NewSaveCloudDatastoreRequest() // SaveCloudDatastoreRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.SaveCloudDatastore(context.Background(), zoneId).SaveCloudDatastoreRequest(saveCloudDatastoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.SaveCloudDatastore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveCloudDatastore`: SaveCloudDatastore200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudsAPI.SaveCloudDatastore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **int64** | The ID of the cloud | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveCloudDatastoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **saveCloudDatastoreRequest** | [**SaveCloudDatastoreRequest**](SaveCloudDatastoreRequest.md) |  | 

### Return type

[**SaveCloudDatastore200Response**](SaveCloudDatastore200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
)

func main() {
	zoneId := int64(7) // int64 | The ID of the cloud
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
**zoneId** | **int64** | The ID of the cloud | 
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


## UpdateClouds

> UpdateClouds200Response UpdateClouds(ctx, id).UpdateCloudsRequest(updateCloudsRequest).Execute()

Updates a Cloud



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
	updateCloudsRequest := *openapiclient.NewUpdateCloudsRequest(*openapiclient.NewUpdateCloudsRequestZone("My Cloud", "ZoneType_example", int64(3), map[string]interface{}({"id":1}))) // UpdateCloudsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudsAPI.UpdateClouds(context.Background(), id).UpdateCloudsRequest(updateCloudsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudsAPI.UpdateClouds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClouds`: UpdateClouds200Response
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

[**UpdateClouds200Response**](UpdateClouds200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

