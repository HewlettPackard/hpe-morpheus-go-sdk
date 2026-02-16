# \AutomationAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTasks**](AutomationAPI.md#AddTasks) | **Post** /api/tasks | Creates a Task
[**GetTasks**](AutomationAPI.md#GetTasks) | **Get** /api/tasks/{id} | Retrieves a Specific Task
[**ListTasks**](AutomationAPI.md#ListTasks) | **Get** /api/tasks | Retrieves all Tasks
[**RemoveTasks**](AutomationAPI.md#RemoveTasks) | **Delete** /api/tasks/{id} | Deletes a Task
[**UpdateTasks**](AutomationAPI.md#UpdateTasks) | **Put** /api/tasks/{id} | Updates a Task



## AddTasks

> AddTasks200Response AddTasks(ctx).AddTasksRequest(addTasksRequest).Execute()

Creates a Task



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
	addTasksRequest := *openapiclient.NewAddTasksRequest(*openapiclient.NewAddTasksRequestTask("Sample Task", *openapiclient.NewAddTasksRequestTaskTaskType("Code_example"), "ExecuteTarget_example")) // AddTasksRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationAPI.AddTasks(context.Background()).AddTasksRequest(addTasksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationAPI.AddTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTasks`: AddTasks200Response
	fmt.Fprintf(os.Stdout, "Response from `AutomationAPI.AddTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addTasksRequest** | [**AddTasksRequest**](AddTasksRequest.md) |  | 

### Return type

[**AddTasks200Response**](AddTasks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasks

> GetTasks200Response GetTasks(ctx, id).Execute()

Retrieves a Specific Task



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
	resp, r, err := apiClient.AutomationAPI.GetTasks(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationAPI.GetTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTasks`: GetTasks200Response
	fmt.Fprintf(os.Stdout, "Response from `AutomationAPI.GetTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTasks200Response**](GetTasks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTasks

> ListTasks200Response ListTasks(ctx).Type_(type_).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Labels(labels).AllLabels(allLabels).Execute()

Retrieves all Tasks



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
	type_ := "type__example" // string | If specified will return all tasks by `task type` code. Refer to `Task Types` API for up to date listings.  (optional)
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
	sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
	direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	name := "example" // string | Filter by name (optional)
	labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
	allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationAPI.ListTasks(context.Background()).Type_(type_).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Labels(labels).AllLabels(allLabels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationAPI.ListTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTasks`: ListTasks200Response
	fmt.Fprintf(os.Stdout, "Response from `AutomationAPI.ListTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | If specified will return all tasks by &#x60;task type&#x60; code. Refer to &#x60;Task Types&#x60; API for up to date listings.  | 
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 

### Return type

[**ListTasks200Response**](ListTasks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTasks

> RemoveGroups200Response RemoveTasks(ctx, id).Execute()

Deletes a Task



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
	resp, r, err := apiClient.AutomationAPI.RemoveTasks(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationAPI.RemoveTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveTasks`: RemoveGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `AutomationAPI.RemoveTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemoveGroups200Response**](RemoveGroups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTasks

> UpdateTasks200Response UpdateTasks(ctx, id).UpdateTasksRequest(updateTasksRequest).Execute()

Updates a Task



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
	updateTasksRequest := *openapiclient.NewUpdateTasksRequest(*openapiclient.NewUpdateTasksRequestTask()) // UpdateTasksRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationAPI.UpdateTasks(context.Background(), id).UpdateTasksRequest(updateTasksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationAPI.UpdateTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTasks`: UpdateTasks200Response
	fmt.Fprintf(os.Stdout, "Response from `AutomationAPI.UpdateTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTasksRequest** | [**UpdateTasksRequest**](UpdateTasksRequest.md) |  | 

### Return type

[**UpdateTasks200Response**](UpdateTasks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

