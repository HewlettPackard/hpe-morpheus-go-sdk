# \RolesAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRoles**](RolesAPI.md#AddRoles) | **Post** /api/roles | Create role
[**DeleteRole**](RolesAPI.md#DeleteRole) | **Delete** /api/roles/{id} | Delete role
[**GetRole**](RolesAPI.md#GetRole) | **Get** /api/roles/{id} | Get role
[**ListRoles**](RolesAPI.md#ListRoles) | **Get** /api/roles | List roles
[**UpdateRole**](RolesAPI.md#UpdateRole) | **Put** /api/roles/{id} | Update role
[**UpdateRoleBlueprintAccess**](RolesAPI.md#UpdateRoleBlueprintAccess) | **Put** /api/roles/{id}/update-blueprint | Customizing Blueprint Access
[**UpdateRoleCatalogItemTypeAccess**](RolesAPI.md#UpdateRoleCatalogItemTypeAccess) | **Put** /api/roles/{id}/update-catalog-item-type | Customizing Catalog Item Type Access
[**UpdateRoleCloudAccess**](RolesAPI.md#UpdateRoleCloudAccess) | **Put** /api/roles/{id}/update-cloud | Customizing Cloud Access
[**UpdateRoleGroupAccess**](RolesAPI.md#UpdateRoleGroupAccess) | **Put** /api/roles/{id}/update-group | Customizing Group Access
[**UpdateRoleInstanceTypeAccess**](RolesAPI.md#UpdateRoleInstanceTypeAccess) | **Put** /api/roles/{id}/update-instance-type | Customizing Instance Type Access
[**UpdateRolePermission**](RolesAPI.md#UpdateRolePermission) | **Put** /api/roles/{id}/update-permission | Updating Role Permissions
[**UpdateRolePersonaAccess**](RolesAPI.md#UpdateRolePersonaAccess) | **Put** /api/roles/{id}/update-persona | Customizing Persona Access
[**UpdateRoleReportTypeAccess**](RolesAPI.md#UpdateRoleReportTypeAccess) | **Put** /api/roles/{id}/update-report-type | Customizing Report Type Access
[**UpdateRoleTaskAccess**](RolesAPI.md#UpdateRoleTaskAccess) | **Put** /api/roles/{id}/update-task | Customizing Task Access
[**UpdateRoleVDIPoolAccess**](RolesAPI.md#UpdateRoleVDIPoolAccess) | **Put** /api/roles/{id}/update-vdi-pool | Customizing VDI Pool Access
[**UpdateRoleWorkflowAccess**](RolesAPI.md#UpdateRoleWorkflowAccess) | **Put** /api/roles/{id}/update-task-set | Customizing Workflow Access



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
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/sdk"
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

> DeleteAlerts200Response DeleteRole(ctx, id).Execute()

Delete role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/sdk"
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
	// response from `DeleteRole`: DeleteAlerts200Response
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

[**DeleteAlerts200Response**](DeleteAlerts200Response.md)

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
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/sdk"
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
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/sdk"
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

> AddRoles200Response UpdateRole(ctx, id).IncludeDefaultAccess(includeDefaultAccess).UpdateRoleRequest(updateRoleRequest).Execute()

Update role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/sdk"
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
	// response from `UpdateRole`: AddRoles200Response
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

[**AddRoles200Response**](AddRoles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleBlueprintAccess

> UpdateRoleBlueprintAccess200Response UpdateRoleBlueprintAccess(ctx, id).UpdateRoleBlueprintAccessRequest(updateRoleBlueprintAccessRequest).Execute()

Customizing Blueprint Access



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateRoleBlueprintAccessRequest := openapiclient.updateRoleBlueprintAccess_request{UpdateRoleBlueprintAccessRequestOneOf: openapiclient.NewUpdateRoleBlueprintAccessRequestOneOf(int32(123), "Access_example")} // UpdateRoleBlueprintAccessRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.UpdateRoleBlueprintAccess(context.Background(), id).UpdateRoleBlueprintAccessRequest(updateRoleBlueprintAccessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.UpdateRoleBlueprintAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoleBlueprintAccess`: UpdateRoleBlueprintAccess200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.UpdateRoleBlueprintAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleBlueprintAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRoleBlueprintAccessRequest** | [**UpdateRoleBlueprintAccessRequest**](UpdateRoleBlueprintAccessRequest.md) |  | 

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleCatalogItemTypeAccess

> UpdateRoleBlueprintAccess200Response UpdateRoleCatalogItemTypeAccess(ctx, id).UpdateRoleCatalogItemTypeAccessRequest(updateRoleCatalogItemTypeAccessRequest).Execute()

Customizing Catalog Item Type Access



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateRoleCatalogItemTypeAccessRequest := openapiclient.updateRoleCatalogItemTypeAccess_request{UpdateRoleCatalogItemTypeAccessRequestOneOf: openapiclient.NewUpdateRoleCatalogItemTypeAccessRequestOneOf(int32(123), "Access_example")} // UpdateRoleCatalogItemTypeAccessRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.UpdateRoleCatalogItemTypeAccess(context.Background(), id).UpdateRoleCatalogItemTypeAccessRequest(updateRoleCatalogItemTypeAccessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.UpdateRoleCatalogItemTypeAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoleCatalogItemTypeAccess`: UpdateRoleBlueprintAccess200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.UpdateRoleCatalogItemTypeAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleCatalogItemTypeAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRoleCatalogItemTypeAccessRequest** | [**UpdateRoleCatalogItemTypeAccessRequest**](UpdateRoleCatalogItemTypeAccessRequest.md) |  | 

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleCloudAccess

> UpdateRoleBlueprintAccess200Response UpdateRoleCloudAccess(ctx, id).UpdateRoleCloudAccessRequest(updateRoleCloudAccessRequest).Execute()

Customizing Cloud Access



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateRoleCloudAccessRequest := openapiclient.updateRoleCloudAccess_request{UpdateRoleCloudAccessRequestOneOf: openapiclient.NewUpdateRoleCloudAccessRequestOneOf(int32(123), "Access_example")} // UpdateRoleCloudAccessRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.UpdateRoleCloudAccess(context.Background(), id).UpdateRoleCloudAccessRequest(updateRoleCloudAccessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.UpdateRoleCloudAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoleCloudAccess`: UpdateRoleBlueprintAccess200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.UpdateRoleCloudAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleCloudAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRoleCloudAccessRequest** | [**UpdateRoleCloudAccessRequest**](UpdateRoleCloudAccessRequest.md) |  | 

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleGroupAccess

> UpdateRoleBlueprintAccess200Response UpdateRoleGroupAccess(ctx, id).UpdateRoleGroupAccessRequest(updateRoleGroupAccessRequest).Execute()

Customizing Group Access



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateRoleGroupAccessRequest := openapiclient.updateRoleGroupAccess_request{UpdateRoleGroupAccessRequestOneOf: openapiclient.NewUpdateRoleGroupAccessRequestOneOf(int32(123), "Access_example")} // UpdateRoleGroupAccessRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.UpdateRoleGroupAccess(context.Background(), id).UpdateRoleGroupAccessRequest(updateRoleGroupAccessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.UpdateRoleGroupAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoleGroupAccess`: UpdateRoleBlueprintAccess200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.UpdateRoleGroupAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleGroupAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRoleGroupAccessRequest** | [**UpdateRoleGroupAccessRequest**](UpdateRoleGroupAccessRequest.md) |  | 

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleInstanceTypeAccess

> UpdateRoleBlueprintAccess200Response UpdateRoleInstanceTypeAccess(ctx, id).UpdateRoleInstanceTypeAccessRequest(updateRoleInstanceTypeAccessRequest).Execute()

Customizing Instance Type Access



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateRoleInstanceTypeAccessRequest := openapiclient.updateRoleInstanceTypeAccess_request{UpdateRoleInstanceTypeAccessRequestOneOf: openapiclient.NewUpdateRoleInstanceTypeAccessRequestOneOf(int32(123), "Access_example")} // UpdateRoleInstanceTypeAccessRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.UpdateRoleInstanceTypeAccess(context.Background(), id).UpdateRoleInstanceTypeAccessRequest(updateRoleInstanceTypeAccessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.UpdateRoleInstanceTypeAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoleInstanceTypeAccess`: UpdateRoleBlueprintAccess200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.UpdateRoleInstanceTypeAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleInstanceTypeAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRoleInstanceTypeAccessRequest** | [**UpdateRoleInstanceTypeAccessRequest**](UpdateRoleInstanceTypeAccessRequest.md) |  | 

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRolePermission

> UpdateRoleBlueprintAccess200Response UpdateRolePermission(ctx, id).UpdateRolePermissionRequest(updateRolePermissionRequest).Execute()

Updating Role Permissions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateRolePermissionRequest := openapiclient.updateRolePermission_request{DefaultBlueprintPermission: openapiclient.NewDefaultBlueprintPermission("PermissionCode_example", "Access_example")} // UpdateRolePermissionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.UpdateRolePermission(context.Background(), id).UpdateRolePermissionRequest(updateRolePermissionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.UpdateRolePermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRolePermission`: UpdateRoleBlueprintAccess200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.UpdateRolePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRolePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRolePermissionRequest** | [**UpdateRolePermissionRequest**](UpdateRolePermissionRequest.md) |  | 

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRolePersonaAccess

> UpdateRoleBlueprintAccess200Response UpdateRolePersonaAccess(ctx, id).UpdateRolePersonaAccessRequest(updateRolePersonaAccessRequest).Execute()

Customizing Persona Access



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateRolePersonaAccessRequest := openapiclient.updateRolePersonaAccess_request{UpdateRolePersonaAccessRequestOneOf: openapiclient.NewUpdateRolePersonaAccessRequestOneOf("PersonaCode_example", "Access_example")} // UpdateRolePersonaAccessRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.UpdateRolePersonaAccess(context.Background(), id).UpdateRolePersonaAccessRequest(updateRolePersonaAccessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.UpdateRolePersonaAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRolePersonaAccess`: UpdateRoleBlueprintAccess200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.UpdateRolePersonaAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRolePersonaAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRolePersonaAccessRequest** | [**UpdateRolePersonaAccessRequest**](UpdateRolePersonaAccessRequest.md) |  | 

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleReportTypeAccess

> UpdateRoleBlueprintAccess200Response UpdateRoleReportTypeAccess(ctx, id).UpdateRoleReportTypeAccessRequest(updateRoleReportTypeAccessRequest).Execute()

Customizing Report Type Access



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateRoleReportTypeAccessRequest := openapiclient.updateRoleReportTypeAccess_request{UpdateRoleReportTypeAccessRequestOneOf: openapiclient.NewUpdateRoleReportTypeAccessRequestOneOf(int32(123), "Access_example")} // UpdateRoleReportTypeAccessRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.UpdateRoleReportTypeAccess(context.Background(), id).UpdateRoleReportTypeAccessRequest(updateRoleReportTypeAccessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.UpdateRoleReportTypeAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoleReportTypeAccess`: UpdateRoleBlueprintAccess200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.UpdateRoleReportTypeAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleReportTypeAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRoleReportTypeAccessRequest** | [**UpdateRoleReportTypeAccessRequest**](UpdateRoleReportTypeAccessRequest.md) |  | 

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleTaskAccess

> UpdateRoleBlueprintAccess200Response UpdateRoleTaskAccess(ctx, id).UpdateRoleTaskAccessRequest(updateRoleTaskAccessRequest).Execute()

Customizing Task Access



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateRoleTaskAccessRequest := openapiclient.updateRoleTaskAccess_request{UpdateRoleTaskAccessRequestOneOf: openapiclient.NewUpdateRoleTaskAccessRequestOneOf(int32(123), "Access_example")} // UpdateRoleTaskAccessRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.UpdateRoleTaskAccess(context.Background(), id).UpdateRoleTaskAccessRequest(updateRoleTaskAccessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.UpdateRoleTaskAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoleTaskAccess`: UpdateRoleBlueprintAccess200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.UpdateRoleTaskAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleTaskAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRoleTaskAccessRequest** | [**UpdateRoleTaskAccessRequest**](UpdateRoleTaskAccessRequest.md) |  | 

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleVDIPoolAccess

> UpdateRoleBlueprintAccess200Response UpdateRoleVDIPoolAccess(ctx, id).UpdateRoleVDIPoolAccessRequest(updateRoleVDIPoolAccessRequest).Execute()

Customizing VDI Pool Access



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateRoleVDIPoolAccessRequest := *openapiclient.NewUpdateRoleVDIPoolAccessRequest(int32(123), "Access_example") // UpdateRoleVDIPoolAccessRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.UpdateRoleVDIPoolAccess(context.Background(), id).UpdateRoleVDIPoolAccessRequest(updateRoleVDIPoolAccessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.UpdateRoleVDIPoolAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoleVDIPoolAccess`: UpdateRoleBlueprintAccess200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.UpdateRoleVDIPoolAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleVDIPoolAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRoleVDIPoolAccessRequest** | [**UpdateRoleVDIPoolAccessRequest**](UpdateRoleVDIPoolAccessRequest.md) |  | 

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleWorkflowAccess

> UpdateRoleBlueprintAccess200Response UpdateRoleWorkflowAccess(ctx, id).UpdateRoleWorkflowAccessRequest(updateRoleWorkflowAccessRequest).Execute()

Customizing Workflow Access



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client/sdk"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateRoleWorkflowAccessRequest := openapiclient.updateRoleWorkflowAccess_request{UpdateRoleWorkflowAccessRequestOneOf: openapiclient.NewUpdateRoleWorkflowAccessRequestOneOf(int32(123), "Access_example")} // UpdateRoleWorkflowAccessRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.UpdateRoleWorkflowAccess(context.Background(), id).UpdateRoleWorkflowAccessRequest(updateRoleWorkflowAccessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.UpdateRoleWorkflowAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoleWorkflowAccess`: UpdateRoleBlueprintAccess200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.UpdateRoleWorkflowAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleWorkflowAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRoleWorkflowAccessRequest** | [**UpdateRoleWorkflowAccessRequest**](UpdateRoleWorkflowAccessRequest.md) |  | 

### Return type

[**UpdateRoleBlueprintAccess200Response**](UpdateRoleBlueprintAccess200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

