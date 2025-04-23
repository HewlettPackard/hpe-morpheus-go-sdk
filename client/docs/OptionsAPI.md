# \OptionsAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOptionSourceData**](OptionsAPI.md#GetOptionSourceData) | **Get** /api/options/{optionSource} | Get Option Source Data
[**ListCodeRepositories**](OptionsAPI.md#ListCodeRepositories) | **Get** /api/options/codeRepositories | Retrieves a list of Code/GIT Repositories
[**ListOptionNetworkOptions**](OptionsAPI.md#ListOptionNetworkOptions) | **Get** /api/options/zoneNetworkOptions | Retrieves network options by zone/cloud
[**ListOptionValues**](OptionsAPI.md#ListOptionValues) | **Get** /api/options/list | Retrieves input option values



## GetOptionSourceData

> GetOptionSourceData200Response GetOptionSourceData(ctx, optionSource).Execute()

Get Option Source Data



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
	optionSource := "keyPairs" // string | `optionSource` to be listed

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetOptionSourceData(context.Background(), optionSource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetOptionSourceData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOptionSourceData`: GetOptionSourceData200Response
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetOptionSourceData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optionSource** | **string** | &#x60;optionSource&#x60; to be listed | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOptionSourceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOptionSourceData200Response**](GetOptionSourceData200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCodeRepositories

> ListCodeRepositories200Response ListCodeRepositories(ctx).IntegrationId(integrationId).Execute()

Retrieves a list of Code/GIT Repositories



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
	integrationId := int64(33) // int64 | Filter by an integration Id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.ListCodeRepositories(context.Background()).IntegrationId(integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.ListCodeRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCodeRepositories`: ListCodeRepositories200Response
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.ListCodeRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCodeRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationId** | **int64** | Filter by an integration Id. | 

### Return type

[**ListCodeRepositories200Response**](ListCodeRepositories200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOptionNetworkOptions

> ListOptionNetworkOptions200Response ListOptionNetworkOptions(ctx).ZoneId(zoneId).ProvisionTypeId(provisionTypeId).Execute()

Retrieves network options by zone/cloud



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
	zoneId := int64(3) // int64 | The Zone ID for Filtering (optional)
	provisionTypeId := int64(22) // int64 | Provision type filter, restricts query to only load service plans of specified provision type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.ListOptionNetworkOptions(context.Background()).ZoneId(zoneId).ProvisionTypeId(provisionTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.ListOptionNetworkOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOptionNetworkOptions`: ListOptionNetworkOptions200Response
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.ListOptionNetworkOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOptionNetworkOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zoneId** | **int64** | The Zone ID for Filtering | 
 **provisionTypeId** | **int64** | Provision type filter, restricts query to only load service plans of specified provision type | 

### Return type

[**ListOptionNetworkOptions200Response**](ListOptionNetworkOptions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOptionValues

> ListOptionValues200Response ListOptionValues(ctx).OptionTypeId(optionTypeId).Config(config).Execute()

Retrieves input option values



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
	optionTypeId := int64(789) // int64 | Input or Option Type ID
	config := map[string]interface{}{ ... } // map[string]interface{} | Input parameters are required if the input is dependent on them.  Fields must be prefixed with `config.` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.ListOptionValues(context.Background()).OptionTypeId(optionTypeId).Config(config).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.ListOptionValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOptionValues`: ListOptionValues200Response
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.ListOptionValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOptionValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optionTypeId** | **int64** | Input or Option Type ID | 
 **config** | [**map[string]interface{}**](map[string]interface{}.md) | Input parameters are required if the input is dependent on them.  Fields must be prefixed with &#x60;config.&#x60; | 

### Return type

[**ListOptionValues200Response**](ListOptionValues200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

