# \LibraryAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInstanceType**](LibraryAPI.md#AddInstanceType) | **Post** /api/library/instance-types | Create an Instance Type
[**AddLayout**](LibraryAPI.md#AddLayout) | **Post** /api/library/instance-types/{instanceTypeId}/layouts | Create a Layout
[**AddVirtualImage**](LibraryAPI.md#AddVirtualImage) | **Post** /api/virtual-images | Create a Virtual Image
[**DeleteInstanceType**](LibraryAPI.md#DeleteInstanceType) | **Delete** /api/library/instance-types/{instanceTypeId} | Delete an Instance Type
[**DeleteLayout**](LibraryAPI.md#DeleteLayout) | **Delete** /api/library/layouts/{id} | Delete a Layout
[**GetInstanceType**](LibraryAPI.md#GetInstanceType) | **Get** /api/library/instance-types/{instanceTypeId} | Get a Specific Instance Type
[**GetLayout**](LibraryAPI.md#GetLayout) | **Get** /api/library/layouts/{id} | Get a Specific Layout
[**GetVirtualImage**](LibraryAPI.md#GetVirtualImage) | **Get** /api/virtual-images/{virtualImageId} | Get a Specific Virtual Image
[**ListInstanceTypes**](LibraryAPI.md#ListInstanceTypes) | **Get** /api/library/instance-types | Get All Instance Types
[**ListLayouts**](LibraryAPI.md#ListLayouts) | **Get** /api/library/layouts | Get All Layouts
[**ListLayoutsForInstanceType**](LibraryAPI.md#ListLayoutsForInstanceType) | **Get** /api/library/instance-types/{instanceTypeId}/layouts | Get All Layouts For an Instance Type
[**ListVirtualImages**](LibraryAPI.md#ListVirtualImages) | **Get** /api/virtual-images | Get List of Virtual Images
[**RemoveVirtualImage**](LibraryAPI.md#RemoveVirtualImage) | **Delete** /api/virtual-images/{virtualImageId} | Delete a Virtual Image
[**UpdateInstanceType**](LibraryAPI.md#UpdateInstanceType) | **Put** /api/library/instance-types/{instanceTypeId} | Update an Instance Type
[**UpdateLayout**](LibraryAPI.md#UpdateLayout) | **Put** /api/library/layouts/{id} | Update a Layout
[**UpdateVirtualImage**](LibraryAPI.md#UpdateVirtualImage) | **Put** /api/virtual-images/{virtualImageId} | Update a Virtual Image



## AddInstanceType

> AddInstanceType200Response AddInstanceType(ctx).AddInstanceTypeRequest(addInstanceTypeRequest).Execute()

Create an Instance Type



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
	addInstanceTypeRequest := *openapiclient.NewAddInstanceTypeRequest() // AddInstanceTypeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.AddInstanceType(context.Background()).AddInstanceTypeRequest(addInstanceTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.AddInstanceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddInstanceType`: AddInstanceType200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.AddInstanceType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addInstanceTypeRequest** | [**AddInstanceTypeRequest**](AddInstanceTypeRequest.md) |  | 

### Return type

[**AddInstanceType200Response**](AddInstanceType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddLayout

> AddLayout200Response AddLayout(ctx, instanceTypeId).AddLayoutRequest(addLayoutRequest).Execute()

Create a Layout



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
	instanceTypeId := int64(2) // int64 | The ID of the instance type
	addLayoutRequest := *openapiclient.NewAddLayoutRequest() // AddLayoutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.AddLayout(context.Background(), instanceTypeId).AddLayoutRequest(addLayoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.AddLayout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddLayout`: AddLayout200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.AddLayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceTypeId** | **int64** | The ID of the instance type | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addLayoutRequest** | [**AddLayoutRequest**](AddLayoutRequest.md) |  | 

### Return type

[**AddLayout200Response**](AddLayout200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddVirtualImage

> AddVirtualImage200Response AddVirtualImage(ctx).AddVirtualImageRequest(addVirtualImageRequest).Execute()

Create a Virtual Image



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
	addVirtualImageRequest := *openapiclient.NewAddVirtualImageRequest(*openapiclient.NewAddVirtualImageRequestVirtualImage()) // AddVirtualImageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.AddVirtualImage(context.Background()).AddVirtualImageRequest(addVirtualImageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.AddVirtualImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddVirtualImage`: AddVirtualImage200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.AddVirtualImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddVirtualImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addVirtualImageRequest** | [**AddVirtualImageRequest**](AddVirtualImageRequest.md) |  | 

### Return type

[**AddVirtualImage200Response**](AddVirtualImage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstanceType

> DeleteInstanceType200Response DeleteInstanceType(ctx, instanceTypeId).Execute()

Delete an Instance Type



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
	instanceTypeId := int64(2) // int64 | The ID of the instance type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.DeleteInstanceType(context.Background(), instanceTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.DeleteInstanceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInstanceType`: DeleteInstanceType200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.DeleteInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceTypeId** | **int64** | The ID of the instance type | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteInstanceType200Response**](DeleteInstanceType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLayout

> DeleteLayout200Response DeleteLayout(ctx, id).Execute()

Delete a Layout



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
	resp, r, err := apiClient.LibraryAPI.DeleteLayout(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.DeleteLayout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLayout`: DeleteLayout200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.DeleteLayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteLayout200Response**](DeleteLayout200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceType

> GetInstanceType200Response GetInstanceType(ctx, instanceTypeId).Execute()

Get a Specific Instance Type



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
	instanceTypeId := int64(2) // int64 | The ID of the instance type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetInstanceType(context.Background(), instanceTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetInstanceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceType`: GetInstanceType200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceTypeId** | **int64** | The ID of the instance type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInstanceType200Response**](GetInstanceType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLayout

> GetLayout200Response GetLayout(ctx, id).Execute()

Get a Specific Layout



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
	resp, r, err := apiClient.LibraryAPI.GetLayout(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetLayout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLayout`: GetLayout200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetLayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLayout200Response**](GetLayout200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualImage

> GetVirtualImage200Response GetVirtualImage(ctx, virtualImageId).Execute()

Get a Specific Virtual Image



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
	virtualImageId := int64(4) // int64 | Virtual Image ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetVirtualImage(context.Background(), virtualImageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetVirtualImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualImage`: GetVirtualImage200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetVirtualImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualImageId** | **int64** | Virtual Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetVirtualImage200Response**](GetVirtualImage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstanceTypes

> ListInstanceTypes200Response ListInstanceTypes(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).Featured(featured).Labels(labels).AllLabels(allLabels).Details(details).Execute()

Get All Instance Types



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
	code := "azr" // string | If specified will return an exact match on code (optional)
	featured := false // bool | Filter by featured (optional)
	labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
	allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)
	details := true // bool | Load full details including optionTypes, environmentVariables, etc. These properties are excluded by default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.ListInstanceTypes(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).Featured(featured).Labels(labels).AllLabels(allLabels).Details(details).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.ListInstanceTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstanceTypes`: ListInstanceTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.ListInstanceTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstanceTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name | 
 **code** | **string** | If specified will return an exact match on code | 
 **featured** | **bool** | Filter by featured | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 
 **details** | **bool** | Load full details including optionTypes, environmentVariables, etc. These properties are excluded by default. | 

### Return type

[**ListInstanceTypes200Response**](ListInstanceTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLayouts

> ListLayouts200Response ListLayouts(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).ProvisionType(provisionType).Labels(labels).AllLabels(allLabels).Execute()

Get All Layouts



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
	code := "azr" // string | If specified will return an exact match on code (optional)
	provisionType := "provisionType_example" // string | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.  (optional)
	labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
	allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.ListLayouts(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).ProvisionType(provisionType).Labels(labels).AllLabels(allLabels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.ListLayouts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLayouts`: ListLayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.ListLayouts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLayoutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name | 
 **code** | **string** | If specified will return an exact match on code | 
 **provisionType** | **string** | Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 

### Return type

[**ListLayouts200Response**](ListLayouts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLayoutsForInstanceType

> ListLayoutsForInstanceType200Response ListLayoutsForInstanceType(ctx, instanceTypeId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).ProvisionType(provisionType).Labels(labels).AllLabels(allLabels).Execute()

Get All Layouts For an Instance Type



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
	instanceTypeId := int64(2) // int64 | The ID of the instance type
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
	sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
	direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	name := "example" // string | Filter by name (optional)
	code := "azr" // string | If specified will return an exact match on code (optional)
	provisionType := "provisionType_example" // string | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.  (optional)
	labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
	allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.ListLayoutsForInstanceType(context.Background(), instanceTypeId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).ProvisionType(provisionType).Labels(labels).AllLabels(allLabels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.ListLayoutsForInstanceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLayoutsForInstanceType`: ListLayoutsForInstanceType200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.ListLayoutsForInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceTypeId** | **int64** | The ID of the instance type | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLayoutsForInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name | 
 **code** | **string** | If specified will return an exact match on code | 
 **provisionType** | **string** | Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 

### Return type

[**ListLayoutsForInstanceType200Response**](ListLayoutsForInstanceType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualImages

> ListVirtualImages200Response ListVirtualImages(ctx).Max(max).Offset(offset).Name(name).Phrase(phrase).LastUpdated(lastUpdated).FilterType(filterType).ImageType(imageType).TagsName(tagsName).Labels(labels).AllLabels(allLabels).Execute()

Get List of Virtual Images



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
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
	name := "example" // string | Filter by name (optional)
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
	filterType := "System" // string | Filter by type, \"User\", \"System\", \"Synced\", or \"All\" (optional) (default to "User")
	imageType := "vmware" // string | Filter by image type code, \"vmware\", \"ami\", etc (optional)
	tagsName := "value" // string | Filter by tags (metadata). This allows filtering by a tag name and value(s)  (optional)
	labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
	allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.ListVirtualImages(context.Background()).Max(max).Offset(offset).Name(name).Phrase(phrase).LastUpdated(lastUpdated).FilterType(filterType).ImageType(imageType).TagsName(tagsName).Labels(labels).AllLabels(allLabels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.ListVirtualImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVirtualImages`: ListVirtualImages200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.ListVirtualImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **name** | **string** | Filter by name | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 
 **filterType** | **string** | Filter by type, \&quot;User\&quot;, \&quot;System\&quot;, \&quot;Synced\&quot;, or \&quot;All\&quot; | [default to &quot;User&quot;]
 **imageType** | **string** | Filter by image type code, \&quot;vmware\&quot;, \&quot;ami\&quot;, etc | 
 **tagsName** | **string** | Filter by tags (metadata). This allows filtering by a tag name and value(s)  | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 

### Return type

[**ListVirtualImages200Response**](ListVirtualImages200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveVirtualImage

> RemoveGroups200Response RemoveVirtualImage(ctx, virtualImageId).RemoveFromCloud(removeFromCloud).Execute()

Delete a Virtual Image



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
	virtualImageId := int64(4) // int64 | Virtual Image ID
	removeFromCloud := true // bool | Include removeFromCloud=true to delete the image location files from all clouds. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.RemoveVirtualImage(context.Background(), virtualImageId).RemoveFromCloud(removeFromCloud).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.RemoveVirtualImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveVirtualImage`: RemoveGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.RemoveVirtualImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualImageId** | **int64** | Virtual Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveVirtualImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeFromCloud** | **bool** | Include removeFromCloud&#x3D;true to delete the image location files from all clouds. | 

### Return type

[**RemoveGroups200Response**](RemoveGroups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInstanceType

> UpdateInstanceType200Response UpdateInstanceType(ctx, instanceTypeId).UpdateInstanceTypeRequest(updateInstanceTypeRequest).Execute()

Update an Instance Type



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
	instanceTypeId := int64(2) // int64 | The ID of the instance type
	updateInstanceTypeRequest := *openapiclient.NewUpdateInstanceTypeRequest() // UpdateInstanceTypeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.UpdateInstanceType(context.Background(), instanceTypeId).UpdateInstanceTypeRequest(updateInstanceTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.UpdateInstanceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInstanceType`: UpdateInstanceType200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.UpdateInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceTypeId** | **int64** | The ID of the instance type | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInstanceTypeRequest** | [**UpdateInstanceTypeRequest**](UpdateInstanceTypeRequest.md) |  | 

### Return type

[**UpdateInstanceType200Response**](UpdateInstanceType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLayout

> UpdateLayout200Response UpdateLayout(ctx, id).UpdateLayoutRequest(updateLayoutRequest).Execute()

Update a Layout



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
	updateLayoutRequest := *openapiclient.NewUpdateLayoutRequest() // UpdateLayoutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.UpdateLayout(context.Background(), id).UpdateLayoutRequest(updateLayoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.UpdateLayout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLayout`: UpdateLayout200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.UpdateLayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateLayoutRequest** | [**UpdateLayoutRequest**](UpdateLayoutRequest.md) |  | 

### Return type

[**UpdateLayout200Response**](UpdateLayout200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualImage

> UpdateVirtualImage200Response UpdateVirtualImage(ctx, virtualImageId).UpdateVirtualImageRequest(updateVirtualImageRequest).Execute()

Update a Virtual Image



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
	virtualImageId := int64(4) // int64 | Virtual Image ID
	updateVirtualImageRequest := *openapiclient.NewUpdateVirtualImageRequest(*openapiclient.NewUpdateVirtualImageRequestVirtualImage()) // UpdateVirtualImageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.UpdateVirtualImage(context.Background(), virtualImageId).UpdateVirtualImageRequest(updateVirtualImageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.UpdateVirtualImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVirtualImage`: UpdateVirtualImage200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.UpdateVirtualImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualImageId** | **int64** | Virtual Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateVirtualImageRequest** | [**UpdateVirtualImageRequest**](UpdateVirtualImageRequest.md) |  | 

### Return type

[**UpdateVirtualImage200Response**](UpdateVirtualImage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

