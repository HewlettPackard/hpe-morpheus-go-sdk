# \ProvisioningAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListProvisionTypes**](ProvisioningAPI.md#ListProvisionTypes) | **Get** /api/provision-types | Retrieves all Provision Types



## ListProvisionTypes

> ListProvisionTypes200Response ListProvisionTypes(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).Execute()

Retrieves all Provision Types



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
	code := "azr" // string | If specified will return an exact match on code (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.ListProvisionTypes(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.ListProvisionTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProvisionTypes`: ListProvisionTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.ListProvisionTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProvisionTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name | 
 **code** | **string** | If specified will return an exact match on code | 

### Return type

[**ListProvisionTypes200Response**](ListProvisionTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

