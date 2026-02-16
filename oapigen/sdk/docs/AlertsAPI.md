# \AlertsAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAlerts**](AlertsAPI.md#DeleteAlerts) | **Delete** /api/monitoring/alerts/{id} | Delete a Specific Alert
[**GetAlerts**](AlertsAPI.md#GetAlerts) | **Get** /api/monitoring/alerts/{id} | Get a Specific Alert
[**UpdateAlerts**](AlertsAPI.md#UpdateAlerts) | **Put** /api/monitoring/alerts/{id} | Update Alert



## DeleteAlerts

> DeleteAlerts200Response DeleteAlerts(ctx, id).Execute()

Delete a Specific Alert



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
	resp, r, err := apiClient.AlertsAPI.DeleteAlerts(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.DeleteAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAlerts`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.DeleteAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertsRequest struct via the builder pattern


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


## GetAlerts

> GetAlerts200Response GetAlerts(ctx, id).Execute()

Get a Specific Alert



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
	resp, r, err := apiClient.AlertsAPI.GetAlerts(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlerts`: GetAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAlerts200Response**](GetAlerts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlerts

> UpdateAlerts200Response UpdateAlerts(ctx, id).UpdateAlertsRequest(updateAlertsRequest).Execute()

Update Alert



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
	updateAlertsRequest := *openapiclient.NewUpdateAlertsRequest(*openapiclient.NewUpdateAlertsRequestAlert()) // UpdateAlertsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.UpdateAlerts(context.Background(), id).UpdateAlertsRequest(updateAlertsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.UpdateAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAlerts`: UpdateAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.UpdateAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAlertsRequest** | [**UpdateAlertsRequest**](UpdateAlertsRequest.md) |  | 

### Return type

[**UpdateAlerts200Response**](UpdateAlerts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

