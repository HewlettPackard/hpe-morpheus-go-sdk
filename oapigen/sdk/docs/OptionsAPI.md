# \OptionsAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOptionSourceData**](OptionsAPI.md#GetOptionSourceData) | **Get** /api/options/{optionSource} | Get Option Source Data



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
	openapiclient "github.com/HewlettPackard/hpe-morpheus-go-sdk/sdk"
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

