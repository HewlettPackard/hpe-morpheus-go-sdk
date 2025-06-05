# \UsersAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUser**](UsersAPI.md#AddUser) | **Post** /api/users | Create a New User
[**AddUserGroup**](UsersAPI.md#AddUserGroup) | **Post** /api/user-groups | Creates a User Group
[**DeleteUser**](UsersAPI.md#DeleteUser) | **Delete** /api/users/{id} | Delete a User
[**DeleteUserGroup**](UsersAPI.md#DeleteUserGroup) | **Delete** /api/user-groups/{id} | Delete User Group
[**DeleteUserSettingsAccessToken**](UsersAPI.md#DeleteUserSettingsAccessToken) | **Put** /api/user-settings/clear-access-token | Revoke API Access Token
[**DeleteUserSettingsAvatar**](UsersAPI.md#DeleteUserSettingsAvatar) | **Delete** /api/user-settings/avatar | Delete Avatar
[**DeleteUserSettingsDesktopBackground**](UsersAPI.md#DeleteUserSettingsDesktopBackground) | **Delete** /api/user-settings/desktop-background | Delete Desktop Background
[**GetUser**](UsersAPI.md#GetUser) | **Get** /api/users/{id} | Get a Specific User
[**GetUserGroup**](UsersAPI.md#GetUserGroup) | **Get** /api/user-groups/{id} | Get a Specific User Group
[**GetUserPermissions**](UsersAPI.md#GetUserPermissions) | **Get** /api/users/{id}/permissions | Get a Specific User Permissions
[**GetUserSettingsApiClients**](UsersAPI.md#GetUserSettingsApiClients) | **Get** /api/user-settings/api-clients | Get Available API Clients
[**ListUserGroups**](UsersAPI.md#ListUserGroups) | **Get** /api/user-groups | Retrieves all User Groups
[**ListUserSettings**](UsersAPI.md#ListUserSettings) | **Get** /api/user-settings | User Settings
[**ListUsers**](UsersAPI.md#ListUsers) | **Get** /api/users | List All Users
[**ListUsersAvailableRoles**](UsersAPI.md#ListUsersAvailableRoles) | **Get** /api/users/available-roles | List available roles for a user
[**UpdateUser**](UsersAPI.md#UpdateUser) | **Put** /api/users/{id} | Update user
[**UpdateUserGroup**](UsersAPI.md#UpdateUserGroup) | **Put** /api/user-groups/{id} | Update User Group
[**UpdateUserSettings**](UsersAPI.md#UpdateUserSettings) | **Put** /api/user-settings | Update User Settings
[**UpdateUserSettingsAccessToken**](UsersAPI.md#UpdateUserSettingsAccessToken) | **Put** /api/user-settings/regenerate-access-token | Regenerate API Access Token
[**UpdateUserSettingsAvatar**](UsersAPI.md#UpdateUserSettingsAvatar) | **Post** /api/user-settings/avatar | Update Avatar
[**UpdateUserSettingsDesktopBackground**](UsersAPI.md#UpdateUserSettingsDesktopBackground) | **Post** /api/user-settings/desktop-background | Update Desktop Background



## AddUser

> AddUserTenant200Response AddUser(ctx).AccountId(accountId).AddUserTenantRequest(addUserTenantRequest).Execute()

Create a New User



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
	accountId := int64(789) // int64 | Tenant ID, create user in a sub tenant account instead of your own. (optional)
	addUserTenantRequest := *openapiclient.NewAddUserTenantRequest(*openapiclient.NewAddUserTenantRequestUser("jsmith", "jsmith@email.com", "Passw0rd!", []openapiclient.GetAlerts200ResponseAllOfChecksInnerAccount{*openapiclient.NewGetAlerts200ResponseAllOfChecksInnerAccount()})) // AddUserTenantRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AddUser(context.Background()).AccountId(accountId).AddUserTenantRequest(addUserTenantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AddUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUser`: AddUserTenant200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AddUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | Tenant ID, create user in a sub tenant account instead of your own. | 
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


## AddUserGroup

> AddUserGroup200Response AddUserGroup(ctx).AddUserGroupRequest(addUserGroupRequest).Execute()

Creates a User Group



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
	addUserGroupRequest := *openapiclient.NewAddUserGroupRequest(*openapiclient.NewAddUserGroupRequestUserGroup()) // AddUserGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AddUserGroup(context.Background()).AddUserGroupRequest(addUserGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AddUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUserGroup`: AddUserGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AddUserGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addUserGroupRequest** | [**AddUserGroupRequest**](AddUserGroupRequest.md) |  | 

### Return type

[**AddUserGroup200Response**](AddUserGroup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteAlerts200Response DeleteUser(ctx, id).Execute()

Delete a User



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
	resp, r, err := apiClient.UsersAPI.DeleteUser(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUser`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## DeleteUserGroup

> DeleteAlerts200Response DeleteUserGroup(ctx, id).Execute()

Delete User Group



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
	resp, r, err := apiClient.UsersAPI.DeleteUserGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUserGroup`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.DeleteUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserGroupRequest struct via the builder pattern


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


## DeleteUserSettingsAccessToken

> DeleteAlerts200Response DeleteUserSettingsAccessToken(ctx).UserId(userId).ClientId(clientId).Execute()

Revoke API Access Token



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
	userId := int64(789) // int64 | ID of User (Only available to the master tenant.)  Default is the current user (optional)
	clientId := "morph-cli" // string | Client ID.  See `Get Available API Clients` for a list of valid `clientId` values. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.DeleteUserSettingsAccessToken(context.Background()).UserId(userId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUserSettingsAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUserSettingsAccessToken`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.DeleteUserSettingsAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserSettingsAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | ID of User (Only available to the master tenant.)  Default is the current user | 
 **clientId** | **string** | Client ID.  See &#x60;Get Available API Clients&#x60; for a list of valid &#x60;clientId&#x60; values. | 

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


## DeleteUserSettingsAvatar

> DeleteAlerts200Response DeleteUserSettingsAvatar(ctx).UserId(userId).Execute()

Delete Avatar



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
	userId := int64(789) // int64 | ID of User (Only available to the master tenant.)  Default is the current user (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.DeleteUserSettingsAvatar(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUserSettingsAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUserSettingsAvatar`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.DeleteUserSettingsAvatar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserSettingsAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | ID of User (Only available to the master tenant.)  Default is the current user | 

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


## DeleteUserSettingsDesktopBackground

> DeleteAlerts200Response DeleteUserSettingsDesktopBackground(ctx).UserId(userId).Execute()

Delete Desktop Background



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
	userId := int64(789) // int64 | ID of User (Only available to the master tenant.)  Default is the current user (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.DeleteUserSettingsDesktopBackground(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUserSettingsDesktopBackground``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUserSettingsDesktopBackground`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.DeleteUserSettingsDesktopBackground`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserSettingsDesktopBackgroundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | ID of User (Only available to the master tenant.)  Default is the current user | 

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


## GetUser

> GetUser200Response GetUser(ctx, id).IncludeAccess(includeAccess).Execute()

Get a Specific User



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
	includeAccess := true // bool | Include `access` information in the response. This is the permissions, clouds, instanceTypes, etc. that the user is authorized for based on their assigned Role(s). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUser(context.Background(), id).IncludeAccess(includeAccess).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: GetUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAccess** | **bool** | Include &#x60;access&#x60; information in the response. This is the permissions, clouds, instanceTypes, etc. that the user is authorized for based on their assigned Role(s). | 

### Return type

[**GetUser200Response**](GetUser200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroup

> GetUserGroup200Response GetUserGroup(ctx, id).Execute()

Get a Specific User Group



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
	resp, r, err := apiClient.UsersAPI.GetUserGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserGroup`: GetUserGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUserGroup200Response**](GetUserGroup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPermissions

> GetUserPermissions200Response GetUserPermissions(ctx, id).Execute()

Get a Specific User Permissions



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
	resp, r, err := apiClient.UsersAPI.GetUserPermissions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserPermissions`: GetUserPermissions200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUserPermissions200Response**](GetUserPermissions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSettingsApiClients

> GetUserSettingsApiClients200Response GetUserSettingsApiClients(ctx).UserId(userId).Execute()

Get Available API Clients



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
	userId := int64(789) // int64 | ID of User (Only available to the master tenant.)  Default is the current user (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUserSettingsApiClients(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserSettingsApiClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserSettingsApiClients`: GetUserSettingsApiClients200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserSettingsApiClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSettingsApiClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | ID of User (Only available to the master tenant.)  Default is the current user | 

### Return type

[**GetUserSettingsApiClients200Response**](GetUserSettingsApiClients200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserGroups

> ListUserGroups200Response ListUserGroups(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()

Retrieves all User Groups



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
	resp, r, err := apiClient.UsersAPI.ListUserGroups(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserGroups`: ListUserGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name | 

### Return type

[**ListUserGroups200Response**](ListUserGroups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserSettings

> ListUserSettings200Response ListUserSettings(ctx).UserId(userId).Execute()

User Settings



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
	userId := int64(789) // int64 | ID of User (Only available to the master tenant.)  Default is the current user (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ListUserSettings(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserSettings`: ListUserSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUserSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | ID of User (Only available to the master tenant.)  Default is the current user | 

### Return type

[**ListUserSettings200Response**](ListUserSettings200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> ListUsers200Response ListUsers(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Username(username).Email(email).RoleId(roleId).LastUpdated(lastUpdated).AccountId(accountId).Global(global).Execute()

List All Users



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
	sort := "sort_example" // string | Sort order, the name of the property to sort by. The default sort order is `username` ascending (optional)
	direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
	phrase := "phrase_example" // string | Search phrase for partial matches on username or email (optional)
	username := "administrator" // string | Username (optional)
	email := "mytest@email.com" // string | E-Mail Address (optional)
	roleId := int64(13) // int64 | Filter by Role ID (supports multiple values) (optional)
	lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
	accountId := int64(3) // int64 | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)
	global := true // bool | Global (All Tenants), load users from all tenants. The default is to only see your own tenant. This is only available to master tenant users with permission to manage tenants and users. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ListUsers(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Username(username).Email(email).RoleId(roleId).LastUpdated(lastUpdated).AccountId(accountId).Global(global).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsers`: ListUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by. The default sort order is &#x60;username&#x60; ascending | 
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on username or email | 
 **username** | **string** | Username | 
 **email** | **string** | E-Mail Address | 
 **roleId** | **int64** | Filter by Role ID (supports multiple values) | 
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 
 **accountId** | **int64** | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | 
 **global** | **bool** | Global (All Tenants), load users from all tenants. The default is to only see your own tenant. This is only available to master tenant users with permission to manage tenants and users. | [default to false]

### Return type

[**ListUsers200Response**](ListUsers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsersAvailableRoles

> ListUsersAvailableRoles200Response ListUsersAvailableRoles(ctx).AccountId(accountId).Execute()

List available roles for a user



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
	accountId := int64(789) // int64 | Tenant ID, find roles available to users of a sub tenant account. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ListUsersAvailableRoles(context.Background()).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUsersAvailableRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsersAvailableRoles`: ListUsersAvailableRoles200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUsersAvailableRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersAvailableRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int64** | Tenant ID, find roles available to users of a sub tenant account. | 

### Return type

[**ListUsersAvailableRoles200Response**](ListUsersAvailableRoles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> AddUserTenant200Response UpdateUser(ctx, id).UpdateUserRequest(updateUserRequest).Execute()

Update user



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
	updateUserRequest := *openapiclient.NewUpdateUserRequest(*openapiclient.NewUpdateUserRequestUser()) // UpdateUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUser(context.Background(), id).UpdateUserRequest(updateUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUser`: AddUserTenant200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserRequest** | [**UpdateUserRequest**](UpdateUserRequest.md) |  | 

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


## UpdateUserGroup

> AddUserGroup200Response UpdateUserGroup(ctx, id).AddUserGroupRequest(addUserGroupRequest).Execute()

Update User Group



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
	addUserGroupRequest := *openapiclient.NewAddUserGroupRequest(*openapiclient.NewAddUserGroupRequestUserGroup()) // AddUserGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserGroup(context.Background(), id).AddUserGroupRequest(addUserGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserGroup`: AddUserGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addUserGroupRequest** | [**AddUserGroupRequest**](AddUserGroupRequest.md) |  | 

### Return type

[**AddUserGroup200Response**](AddUserGroup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserSettings

> DeleteAlerts200Response UpdateUserSettings(ctx).UserId(userId).UpdateUserSettingsRequest(updateUserSettingsRequest).Execute()

Update User Settings



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
	userId := int64(789) // int64 | ID of User (Only available to the master tenant.)  Default is the current user (optional)
	updateUserSettingsRequest := *openapiclient.NewUpdateUserSettingsRequest() // UpdateUserSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserSettings(context.Background()).UserId(userId).UpdateUserSettingsRequest(updateUserSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserSettings`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | ID of User (Only available to the master tenant.)  Default is the current user | 
 **updateUserSettingsRequest** | [**UpdateUserSettingsRequest**](UpdateUserSettingsRequest.md) |  | 

### Return type

[**DeleteAlerts200Response**](DeleteAlerts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserSettingsAccessToken

> UpdateUserSettingsAccessToken200Response UpdateUserSettingsAccessToken(ctx).UserId(userId).ClientId(clientId).Execute()

Regenerate API Access Token



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
	userId := int64(789) // int64 | ID of User (Only available to the master tenant.)  Default is the current user (optional)
	clientId := "morph-cli" // string | Client ID.  See `Get Available API Clients` for a list of valid `clientId` values. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserSettingsAccessToken(context.Background()).UserId(userId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserSettingsAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserSettingsAccessToken`: UpdateUserSettingsAccessToken200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserSettingsAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserSettingsAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | ID of User (Only available to the master tenant.)  Default is the current user | 
 **clientId** | **string** | Client ID.  See &#x60;Get Available API Clients&#x60; for a list of valid &#x60;clientId&#x60; values. | 

### Return type

[**UpdateUserSettingsAccessToken200Response**](UpdateUserSettingsAccessToken200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserSettingsAvatar

> DeleteAlerts200Response UpdateUserSettingsAvatar(ctx).UserId(userId).UserAvatar(userAvatar).Execute()

Update Avatar



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
	userId := int64(789) // int64 | ID of User (Only available to the master tenant.)  Default is the current user (optional)
	userAvatar := "userAvatar_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserSettingsAvatar(context.Background()).UserId(userId).UserAvatar(userAvatar).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserSettingsAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserSettingsAvatar`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserSettingsAvatar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserSettingsAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | ID of User (Only available to the master tenant.)  Default is the current user | 
 **userAvatar** | **string** |  | 

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


## UpdateUserSettingsDesktopBackground

> DeleteAlerts200Response UpdateUserSettingsDesktopBackground(ctx).UserId(userId).UserDesktopBackground(userDesktopBackground).Execute()

Update Desktop Background



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
	userId := int64(789) // int64 | ID of User (Only available to the master tenant.)  Default is the current user (optional)
	userDesktopBackground := "userDesktopBackground_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserSettingsDesktopBackground(context.Background()).UserId(userId).UserDesktopBackground(userDesktopBackground).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserSettingsDesktopBackground``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserSettingsDesktopBackground`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserSettingsDesktopBackground`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserSettingsDesktopBackgroundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | ID of User (Only available to the master tenant.)  Default is the current user | 
 **userDesktopBackground** | **string** |  | 

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

