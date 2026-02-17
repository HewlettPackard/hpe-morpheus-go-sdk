# \TenantsAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserTenant**](TenantsAPI.md#AddUserTenant) | **Post** /api/accounts/{accountId}/users | Create a User For a Tenant



## AddUserTenant

> AddUserTenant200Response AddUserTenant(ctx, accountId).AddUserTenantRequest(addUserTenantRequest).Execute()

Create a User For a Tenant



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
	accountId := int64(7) // int64 | The ID of the subtenant account
	addUserTenantRequest := *openapiclient.NewAddUserTenantRequest(*openapiclient.NewAddUserTenantRequestUser("jsmith", "jsmith@email.com", "Passw0rd!", []openapiclient.AddUserTenantRequestUserRolesInner{*openapiclient.NewAddUserTenantRequestUserRolesInner()})) // AddUserTenantRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantsAPI.AddUserTenant(context.Background(), accountId).AddUserTenantRequest(addUserTenantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.AddUserTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUserTenant`: AddUserTenant200Response
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.AddUserTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** | The ID of the subtenant account | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addUserTenantRequest** | [**AddUserTenantRequest**](AddUserTenantRequest.md) |  | 

### Return type

[**AddUserTenant200Response**](AddUserTenant200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

