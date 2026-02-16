# \RolesAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRoles**](RolesAPI.md#AddRoles) | **Post** /api/roles | Create role
[**DeleteRole**](RolesAPI.md#DeleteRole) | **Delete** /api/roles/{id} | Delete role
[**GetRole**](RolesAPI.md#GetRole) | **Get** /api/roles/{id} | Get role
[**ListRoles**](RolesAPI.md#ListRoles) | **Get** /api/roles | List roles
[**UpdateRole**](RolesAPI.md#UpdateRole) | **Put** /api/roles/{id} | Update role



## AddRoles

> AddRoles200Response AddRoles(ctx).IncludeDefaultAccess(includeDefaultAccess).AddRolesRequest(addRolesRequest).Execute()

Create role



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
	includeDefaultAccess := true // bool | Pass true to include all resource permissions in the response including those with access set to `default`. Only resources with access specific levels are returned by default. eg. `full`, `read` or `none`  (optional)
	addRolesRequest := *openapiclient.NewAddRolesRequest(*openapiclient.NewAddRolesRequestRole("Authority_example")) // AddRolesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.AddRoles(context.Background()).IncludeDefaultAccess(includeDefaultAccess).AddRolesRequest(addRolesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AddRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddRoles`: AddRoles200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.AddRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeDefaultAccess** | **bool** | Pass true to include all resource permissions in the response including those with access set to &#x60;default&#x60;. Only resources with access specific levels are returned by default. eg. &#x60;full&#x60;, &#x60;read&#x60; or &#x60;none&#x60;  | 
 **addRolesRequest** | [**AddRolesRequest**](AddRolesRequest.md) |  | 

### Return type

[**AddRoles200Response**](AddRoles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> RemoveClouds200Response DeleteRole(ctx, id).Execute()

Delete role



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
	resp, r, err := apiClient.RolesAPI.DeleteRole(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.DeleteRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRole`: RemoveClouds200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.DeleteRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetRole

> GetRole200Response GetRole(ctx, id).IncludeDefaultAccess(includeDefaultAccess).Execute()

Get role



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
	includeDefaultAccess := true // bool | Pass true to include all resource permissions in the response including those with access set to `default`. Only resources with access specific levels are returned by default. eg. `full`, `read` or `none`  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.GetRole(context.Background(), id).IncludeDefaultAccess(includeDefaultAccess).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRole`: GetRole200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDefaultAccess** | **bool** | Pass true to include all resource permissions in the response including those with access set to &#x60;default&#x60;. Only resources with access specific levels are returned by default. eg. &#x60;full&#x60;, &#x60;read&#x60; or &#x60;none&#x60;  | 

### Return type

[**GetRole200Response**](GetRole200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoles

> ListRoles200Response ListRoles(ctx).Phrase(phrase).Max(max).Offset(offset).Sort(sort).Direction(direction).Authority(authority).RoleType(roleType).Execute()

List roles



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
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
	sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
	direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
	authority := "authority_example" // string | Filter by authority (optional)
	roleType := "roleType_example" // string | Filter by roleType (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.ListRoles(context.Background()).Phrase(phrase).Max(max).Offset(offset).Sort(sort).Direction(direction).Authority(authority).RoleType(roleType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.ListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoles`: ListRoles200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.ListRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **authority** | **string** | Filter by authority | 
 **roleType** | **string** | Filter by roleType | 

### Return type

[**ListRoles200Response**](ListRoles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> UpdateRole200Response UpdateRole(ctx, id).IncludeDefaultAccess(includeDefaultAccess).UpdateRoleRequest(updateRoleRequest).Execute()

Update role



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
	includeDefaultAccess := true // bool | Pass true to include all resource permissions in the response including those with access set to `default`. Only resources with access specific levels are returned by default. eg. `full`, `read` or `none`  (optional)
	updateRoleRequest := *openapiclient.NewUpdateRoleRequest(*openapiclient.NewUpdateRoleRequestRole()) // UpdateRoleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.UpdateRole(context.Background(), id).IncludeDefaultAccess(includeDefaultAccess).UpdateRoleRequest(updateRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.UpdateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRole`: UpdateRole200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.UpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDefaultAccess** | **bool** | Pass true to include all resource permissions in the response including those with access set to &#x60;default&#x60;. Only resources with access specific levels are returned by default. eg. &#x60;full&#x60;, &#x60;read&#x60; or &#x60;none&#x60;  | 
 **updateRoleRequest** | [**UpdateRoleRequest**](UpdateRoleRequest.md) |  | 

### Return type

[**UpdateRole200Response**](UpdateRole200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

