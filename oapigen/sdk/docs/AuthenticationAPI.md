# \AuthenticationAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccessToken**](AuthenticationAPI.md#GetAccessToken) | **Post** /oauth/token | Provides authentication via username and password



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
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
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

