# \InstancesAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInstance**](InstancesAPI.md#AddInstance) | **Post** /api/instances | Create an Instance
[**DeleteInstance**](InstancesAPI.md#DeleteInstance) | **Delete** /api/instances/{id} | Delete an instance
[**GetEnvVariables**](InstancesAPI.md#GetEnvVariables) | **Get** /api/instances/{id}/envs | Get Env Variables
[**GetInstance**](InstancesAPI.md#GetInstance) | **Get** /api/instances/{id} | Retrieves a Specific Instance
[**GetInstanceTypeProvisioning**](InstancesAPI.md#GetInstanceTypeProvisioning) | **Get** /api/instance-types/{id} | Get Specific Instance Type for Provisioning
[**ListInstances**](InstancesAPI.md#ListInstances) | **Get** /api/instances | Get All Instances
[**ResizeInstance**](InstancesAPI.md#ResizeInstance) | **Put** /api/instances/{id}/resize | Resize an Instance
[**UpdateInstance**](InstancesAPI.md#UpdateInstance) | **Put** /api/instances/{id} | Updating an Instance



## AddInstance

> AddInstance200Response AddInstance(ctx).AddInstanceRequest(addInstanceRequest).Execute()

Create an Instance



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
	addInstanceRequest := *openapiclient.NewAddInstanceRequest(*openapiclient.NewAddInstanceRequestInstance("myinstance", *openapiclient.NewAddInstanceRequestInstanceSite(int64(2)), *openapiclient.NewAddInstanceRequestInstanceInstanceType("Ubuntu"), *openapiclient.NewAddInstanceRequestInstanceLayout(int64(105)), *openapiclient.NewAddInstanceRequestInstancePlan(int64(75))), *openapiclient.NewAddInstanceRequestConfig()) // AddInstanceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.AddInstance(context.Background()).AddInstanceRequest(addInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.AddInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddInstance`: AddInstance200Response
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.AddInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addInstanceRequest** | [**AddInstanceRequest**](AddInstanceRequest.md) |  | 

### Return type

[**AddInstance200Response**](AddInstance200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstance

> DeleteInstance200Response DeleteInstance(ctx, id).PreserveVolumes(preserveVolumes).KeepBackups(keepBackups).ReleaseFloatingIps(releaseFloatingIps).ReleaseEIPs(releaseEIPs).Force(force).RemoveVolumes(removeVolumes).Execute()

Delete an instance



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
	preserveVolumes := "on" // string | Preserve Volumes (optional) (default to "off")
	keepBackups := "on" // string | Preserve copy of backups (optional) (default to "off")
	releaseFloatingIps := "off" // string | Release Floating IPs (optional) (default to "on")
	releaseEIPs := "off" // string | Alias for releaseFloatingIps (optional) (default to "on")
	force := "on" // string | Force Delete (optional) (default to "off")
	removeVolumes := "on" // string | Remove Volumes (optional) (default to "off")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.DeleteInstance(context.Background(), id).PreserveVolumes(preserveVolumes).KeepBackups(keepBackups).ReleaseFloatingIps(releaseFloatingIps).ReleaseEIPs(releaseEIPs).Force(force).RemoveVolumes(removeVolumes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.DeleteInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInstance`: DeleteInstance200Response
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.DeleteInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **preserveVolumes** | **string** | Preserve Volumes | [default to &quot;off&quot;]
 **keepBackups** | **string** | Preserve copy of backups | [default to &quot;off&quot;]
 **releaseFloatingIps** | **string** | Release Floating IPs | [default to &quot;on&quot;]
 **releaseEIPs** | **string** | Alias for releaseFloatingIps | [default to &quot;on&quot;]
 **force** | **string** | Force Delete | [default to &quot;off&quot;]
 **removeVolumes** | **string** | Remove Volumes | [default to &quot;off&quot;]

### Return type

[**DeleteInstance200Response**](DeleteInstance200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvVariables

> GetEnvVariables200Response GetEnvVariables(ctx, id).Execute()

Get Env Variables



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
	resp, r, err := apiClient.InstancesAPI.GetEnvVariables(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.GetEnvVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvVariables`: GetEnvVariables200Response
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.GetEnvVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEnvVariables200Response**](GetEnvVariables200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstance

> GetInstance200Response GetInstance(ctx, id).Details(details).Execute()

Retrieves a Specific Instance



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
	details := true // bool | Include details=false to exclude extra details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.GetInstance(context.Background(), id).Details(details).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstance`: GetInstance200Response
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **details** | **bool** | Include details&#x3D;false to exclude extra details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2. | [default to true]

### Return type

[**GetInstance200Response**](GetInstance200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceTypeProvisioning

> GetInstanceTypeProvisioning200Response GetInstanceTypeProvisioning(ctx, id).Execute()

Get Specific Instance Type for Provisioning



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
	resp, r, err := apiClient.InstancesAPI.GetInstanceTypeProvisioning(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.GetInstanceTypeProvisioning``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceTypeProvisioning`: GetInstanceTypeProvisioning200Response
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.GetInstanceTypeProvisioning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceTypeProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInstanceTypeProvisioning200Response**](GetInstanceTypeProvisioning200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstances

> ListInstances200Response ListInstances(ctx).Max(max).Offset(offset).Name(name).Phrase(phrase).ZoneId(zoneId).SiteId(siteId).PlanId(planId).InstanceType(instanceType).LastUpdated(lastUpdated).CreatedBy(createdBy).AgentInstalled(agentInstalled).Status(status).Environment(environment).ShowDeleted(showDeleted).Deleted(deleted).ExpireDate(expireDate).ExpireDateMin(expireDateMin).ExpireDays(expireDays).ExpireDaysMin(expireDaysMin).ShutdownDate(shutdownDate).ShutdownDateMin(shutdownDateMin).ShutdownDays(shutdownDays).ShutdownDaysMin(shutdownDaysMin).Labels(labels).AllLabels(allLabels).TagsName(tagsName).Details(details).Execute()

Get All Instances



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
	name := "example" // string | Filter by name (optional)
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	zoneId := int64(3) // int64 | The Cloud ID (Zone ID) for Filtering (optional)
	siteId := int64(7) // int64 | The Group ID (Site ID) for Filtering (optional)
	planId := int64(1) // int64 | The Plan ID for Filtering (optional)
	instanceType := "ubuntu" // string | The Instance Type Code for Filtering (optional)
	lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
	createdBy := int64(63) // int64 | The User ID for Filtering (optional)
	agentInstalled := true // bool | Filter instances by if agent is installed or not on the associated servers. (optional)
	status := "running" // string | The instance status for filtering. (optional)
	environment := "lab" // string | The environment for filtering. (optional)
	showDeleted := true // bool | If true, includes instances in pending removal status. (optional) (default to false)
	deleted := true // bool | If true, only deleted resources or instances in pending removal status are returned. (optional)
	expireDate := "2019-01-01" // string | Filter by expireDate less than or equal to specified date (optional)
	expireDateMin := "2019-01-01" // string | Filter expireDate greater than or equal to the specified date (optional)
	expireDays := "20" // string | Filter by expireDays less than or equal to the specified value (optional)
	expireDaysMin := "20" // string | Filter by expireDays greater than or equal to the specified value (optional)
	shutdownDate := "2019-01-01" // string | Filter by shutdownDate less than equal to the specified date (optional)
	shutdownDateMin := "2019-01-01" // string | Filter by shutdownDate greater than or equal to the specified date (optional)
	shutdownDays := "20" // string | Filter by shutdownDays less than or equal to the specified value (optional)
	shutdownDaysMin := "20" // string | Filter by shutdownDays greater than or equal to the specified value (optional)
	labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
	allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)
	tagsName := "value" // string | Filter by tags (metadata). This allows filtering by a tag name and value(s)  (optional)
	details := true // bool | Include details=true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.ListInstances(context.Background()).Max(max).Offset(offset).Name(name).Phrase(phrase).ZoneId(zoneId).SiteId(siteId).PlanId(planId).InstanceType(instanceType).LastUpdated(lastUpdated).CreatedBy(createdBy).AgentInstalled(agentInstalled).Status(status).Environment(environment).ShowDeleted(showDeleted).Deleted(deleted).ExpireDate(expireDate).ExpireDateMin(expireDateMin).ExpireDays(expireDays).ExpireDaysMin(expireDaysMin).ShutdownDate(shutdownDate).ShutdownDateMin(shutdownDateMin).ShutdownDays(shutdownDays).ShutdownDaysMin(shutdownDaysMin).Labels(labels).AllLabels(allLabels).TagsName(tagsName).Details(details).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.ListInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstances`: ListInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.ListInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **name** | **string** | Filter by name | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **zoneId** | **int64** | The Cloud ID (Zone ID) for Filtering | 
 **siteId** | **int64** | The Group ID (Site ID) for Filtering | 
 **planId** | **int64** | The Plan ID for Filtering | 
 **instanceType** | **string** | The Instance Type Code for Filtering | 
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 
 **createdBy** | **int64** | The User ID for Filtering | 
 **agentInstalled** | **bool** | Filter instances by if agent is installed or not on the associated servers. | 
 **status** | **string** | The instance status for filtering. | 
 **environment** | **string** | The environment for filtering. | 
 **showDeleted** | **bool** | If true, includes instances in pending removal status. | [default to false]
 **deleted** | **bool** | If true, only deleted resources or instances in pending removal status are returned. | 
 **expireDate** | **string** | Filter by expireDate less than or equal to specified date | 
 **expireDateMin** | **string** | Filter expireDate greater than or equal to the specified date | 
 **expireDays** | **string** | Filter by expireDays less than or equal to the specified value | 
 **expireDaysMin** | **string** | Filter by expireDays greater than or equal to the specified value | 
 **shutdownDate** | **string** | Filter by shutdownDate less than equal to the specified date | 
 **shutdownDateMin** | **string** | Filter by shutdownDate greater than or equal to the specified date | 
 **shutdownDays** | **string** | Filter by shutdownDays less than or equal to the specified value | 
 **shutdownDaysMin** | **string** | Filter by shutdownDays greater than or equal to the specified value | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 
 **tagsName** | **string** | Filter by tags (metadata). This allows filtering by a tag name and value(s)  | 
 **details** | **bool** | Include details&#x3D;true to return more details about the instance, ie. containerDetails. Available in api version 5.2.8/5.3.2. | [default to false]

### Return type

[**ListInstances200Response**](ListInstances200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResizeInstance

> ResizeInstance200Response ResizeInstance(ctx, id).ResizeInstanceRequest(resizeInstanceRequest).Execute()

Resize an Instance



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
	resizeInstanceRequest := *openapiclient.NewResizeInstanceRequest() // ResizeInstanceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.ResizeInstance(context.Background(), id).ResizeInstanceRequest(resizeInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.ResizeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResizeInstance`: ResizeInstance200Response
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.ResizeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiResizeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resizeInstanceRequest** | [**ResizeInstanceRequest**](ResizeInstanceRequest.md) |  | 

### Return type

[**ResizeInstance200Response**](ResizeInstance200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInstance

> UpdateInstance200Response UpdateInstance(ctx, id).UpdateInstanceRequest(updateInstanceRequest).Execute()

Updating an Instance



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
	updateInstanceRequest := *openapiclient.NewUpdateInstanceRequest() // UpdateInstanceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.UpdateInstance(context.Background(), id).UpdateInstanceRequest(updateInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.UpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInstance`: UpdateInstance200Response
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.UpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInstanceRequest** | [**UpdateInstanceRequest**](UpdateInstanceRequest.md) |  | 

### Return type

[**UpdateInstance200Response**](UpdateInstance200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

