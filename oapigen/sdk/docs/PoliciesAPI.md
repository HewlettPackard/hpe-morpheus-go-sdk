# \PoliciesAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPolicies**](PoliciesAPI.md#AddPolicies) | **Post** /api/policies | Creates a Policy
[**GetPolicies**](PoliciesAPI.md#GetPolicies) | **Get** /api/policies/{id} | Retrieves a Specific Policy
[**ListPolicies**](PoliciesAPI.md#ListPolicies) | **Get** /api/policies | Retrieves all Policies
[**RemovePolicies**](PoliciesAPI.md#RemovePolicies) | **Delete** /api/policies/{id} | Deletes a Policy
[**UpdatePolicies**](PoliciesAPI.md#UpdatePolicies) | **Put** /api/policies/{id} | Updates a Policy



## AddPolicies

> AddPolicies200Response AddPolicies(ctx).AddPoliciesRequest(addPoliciesRequest).Execute()

Creates a Policy



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
	addPoliciesRequest := *openapiclient.NewAddPoliciesRequest(*openapiclient.NewAddPoliciesRequestPolicy("Sample Policy", *openapiclient.NewAddPoliciesRequestPolicyPolicyType(), *openapiclient.NewAddPoliciesRequestPolicyConfig(*openapiclient.NewMaxMemoryPolicyTypeConfiguration1MaxMemory()))) // AddPoliciesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.AddPolicies(context.Background()).AddPoliciesRequest(addPoliciesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.AddPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPolicies`: AddPolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.AddPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addPoliciesRequest** | [**AddPoliciesRequest**](AddPoliciesRequest.md) |  | 

### Return type

[**AddPolicies200Response**](AddPolicies200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicies

> GetPolicies200Response GetPolicies(ctx, id).Execute()

Retrieves a Specific Policy



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
	resp, r, err := apiClient.PoliciesAPI.GetPolicies(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.GetPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicies`: GetPolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.GetPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPolicies200Response**](GetPolicies200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicies

> ListPolicies200Response ListPolicies(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()

Retrieves all Policies



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.ListPolicies(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.ListPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPolicies`: ListPolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.ListPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name | 

### Return type

[**ListPolicies200Response**](ListPolicies200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePolicies

> RemoveGroups200Response RemovePolicies(ctx, id).Execute()

Deletes a Policy



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
	resp, r, err := apiClient.PoliciesAPI.RemovePolicies(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.RemovePolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemovePolicies`: RemoveGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.RemovePolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UpdatePolicies

> UpdatePolicies200Response UpdatePolicies(ctx, id).UpdatePoliciesRequest(updatePoliciesRequest).Execute()

Updates a Policy



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
	updatePoliciesRequest := *openapiclient.NewUpdatePoliciesRequest(*openapiclient.NewUpdatePoliciesRequestPolicy()) // UpdatePoliciesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.UpdatePolicies(context.Background(), id).UpdatePoliciesRequest(updatePoliciesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.UpdatePolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePolicies`: UpdatePolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.UpdatePolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePoliciesRequest** | [**UpdatePoliciesRequest**](UpdatePoliciesRequest.md) |  | 

### Return type

[**UpdatePolicies200Response**](UpdatePolicies200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

