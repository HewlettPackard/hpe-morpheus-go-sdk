# \LogSettingsAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLogSettingsSyslogRules**](LogSettingsAPI.md#AddLogSettingsSyslogRules) | **Post** /api/log-settings/syslog-rules | Create a New Syslog Rule
[**DeleteLogSettingsSyslogRules**](LogSettingsAPI.md#DeleteLogSettingsSyslogRules) | **Delete** /api/log-settings/syslog-rules/{id} | Delete a Specific Syslog Rule
[**ListLogSettings**](LogSettingsAPI.md#ListLogSettings) | **Get** /api/log-settings | List All Log Settings
[**UpdateLogSettings**](LogSettingsAPI.md#UpdateLogSettings) | **Put** /api/log-settings | Update Log Settings



## AddLogSettingsSyslogRules

> DeleteAlerts200Response AddLogSettingsSyslogRules(ctx).AddLogSettingsSyslogRulesRequest(addLogSettingsSyslogRulesRequest).Execute()

Create a New Syslog Rule



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
	addLogSettingsSyslogRulesRequest := *openapiclient.NewAddLogSettingsSyslogRulesRequest(*openapiclient.NewAddLogSettingsSyslogRulesRequestSyslogRule("foo", "*.* @@bar.com:8080")) // AddLogSettingsSyslogRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogSettingsAPI.AddLogSettingsSyslogRules(context.Background()).AddLogSettingsSyslogRulesRequest(addLogSettingsSyslogRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogSettingsAPI.AddLogSettingsSyslogRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddLogSettingsSyslogRules`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `LogSettingsAPI.AddLogSettingsSyslogRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLogSettingsSyslogRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addLogSettingsSyslogRulesRequest** | [**AddLogSettingsSyslogRulesRequest**](AddLogSettingsSyslogRulesRequest.md) |  | 

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


## DeleteLogSettingsSyslogRules

> DeleteAlerts200Response DeleteLogSettingsSyslogRules(ctx, id).Execute()

Delete a Specific Syslog Rule



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
	resp, r, err := apiClient.LogSettingsAPI.DeleteLogSettingsSyslogRules(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogSettingsAPI.DeleteLogSettingsSyslogRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLogSettingsSyslogRules`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `LogSettingsAPI.DeleteLogSettingsSyslogRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogSettingsSyslogRulesRequest struct via the builder pattern


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


## ListLogSettings

> ListLogSettings200Response ListLogSettings(ctx).Execute()

List All Log Settings



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
	resp, r, err := apiClient.LogSettingsAPI.ListLogSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogSettingsAPI.ListLogSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLogSettings`: ListLogSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `LogSettingsAPI.ListLogSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLogSettingsRequest struct via the builder pattern


### Return type

[**ListLogSettings200Response**](ListLogSettings200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogSettings

> UpdateGuidanceSettings200Response UpdateLogSettings(ctx).ListLogSettings200Response(listLogSettings200Response).Execute()

Update Log Settings



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
	listLogSettings200Response := *openapiclient.NewListLogSettings200Response() // ListLogSettings200Response |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogSettingsAPI.UpdateLogSettings(context.Background()).ListLogSettings200Response(listLogSettings200Response).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogSettingsAPI.UpdateLogSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLogSettings`: UpdateGuidanceSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `LogSettingsAPI.UpdateLogSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listLogSettings200Response** | [**ListLogSettings200Response**](ListLogSettings200Response.md) |  | 

### Return type

[**UpdateGuidanceSettings200Response**](UpdateGuidanceSettings200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

