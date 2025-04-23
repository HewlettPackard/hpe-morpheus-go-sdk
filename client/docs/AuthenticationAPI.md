# \AuthenticationAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ForgotPassword**](AuthenticationAPI.md#ForgotPassword) | **Post** /api/forgot/send-email | Request a reset password email
[**GetAccessToken**](AuthenticationAPI.md#GetAccessToken) | **Post** /oauth/token | Provides authentication via username and password
[**ResetPassword**](AuthenticationAPI.md#ResetPassword) | **Post** /api/forgot/reset-password | Reset user password
[**Whoami**](AuthenticationAPI.md#Whoami) | **Get** /api/whoami | Retrieves information about current user roles and permissions



## ForgotPassword

> ForgotPassword200Response ForgotPassword(ctx).ForgotPasswordRequest(forgotPasswordRequest).Execute()

Request a reset password email



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
	forgotPasswordRequest := *openapiclient.NewForgotPasswordRequest("Username_example") // ForgotPasswordRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.ForgotPassword(context.Background()).ForgotPasswordRequest(forgotPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.ForgotPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ForgotPassword`: ForgotPassword200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.ForgotPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiForgotPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forgotPasswordRequest** | [**ForgotPasswordRequest**](ForgotPasswordRequest.md) |  | 

### Return type

[**ForgotPassword200Response**](ForgotPassword200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessToken

> GetAccessToken200Response GetAccessToken(ctx).ClientId(clientId).GrantType(grantType).Scope(scope).Username(username).Password(password).RefreshToken(refreshToken).Execute()

Provides authentication via username and password



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
	clientId := "clientId_example" // string | Client ID, use morph-api. Users may only have one access token per Client ID. The CLI uses morph-cli.
	grantType := "grantType_example" // string | OAuth Grant Type, use password.
	scope := "scope_example" // string | OAuth token scope, use write.
	username := "username_example" // string | Specified as \\\"username\\\" or \\\"tenantId\\\\username\\\" with proper HTML encoding and used in conjunction with `password`. Not utilized with `refresh_token`. (optional)
	password := "password_example" // string | The Password for defined `username`. Must have proper HTML encoding (optional)
	refreshToken := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.GetAccessToken(context.Background()).ClientId(clientId).GrantType(grantType).Scope(scope).Username(username).Password(password).RefreshToken(refreshToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessToken`: GetAccessToken200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | Client ID, use morph-api. Users may only have one access token per Client ID. The CLI uses morph-cli. | 
 **grantType** | **string** | OAuth Grant Type, use password. | 
 **scope** | **string** | OAuth token scope, use write. | 
 **username** | **string** | Specified as \\\&quot;username\\\&quot; or \\\&quot;tenantId\\\\username\\\&quot; with proper HTML encoding and used in conjunction with &#x60;password&#x60;. Not utilized with &#x60;refresh_token&#x60;. | 
 **password** | **string** | The Password for defined &#x60;username&#x60;. Must have proper HTML encoding | 
 **refreshToken** | [**interface{}**](interface{}.md) |  | 

### Return type

[**GetAccessToken200Response**](GetAccessToken200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPassword

> ResetPassword200Response ResetPassword(ctx).ResetPasswordRequest(resetPasswordRequest).Execute()

Reset user password



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
	resetPasswordRequest := *openapiclient.NewResetPasswordRequest("Token_example", "Password_example") // ResetPasswordRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.ResetPassword(context.Background()).ResetPasswordRequest(resetPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.ResetPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetPassword`: ResetPassword200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.ResetPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordRequest** | [**ResetPasswordRequest**](ResetPasswordRequest.md) |  | 

### Return type

[**ResetPassword200Response**](ResetPassword200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Whoami

> Whoami200Response Whoami(ctx).Execute()

Retrieves information about current user roles and permissions



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.Whoami(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Whoami``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Whoami`: Whoami200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Whoami`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWhoamiRequest struct via the builder pattern


### Return type

[**Whoami200Response**](Whoami200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

