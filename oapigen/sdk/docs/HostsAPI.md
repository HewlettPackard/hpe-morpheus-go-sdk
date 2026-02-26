# \HostsAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBaremetalHost**](HostsAPI.md#AddBaremetalHost) | **Post** /api/servers | Add a Baremetal Host
[**GetHost**](HostsAPI.md#GetHost) | **Get** /api/servers/{id} | Get a Specific Host
[**ListHosts**](HostsAPI.md#ListHosts) | **Get** /api/servers | Get All Hosts
[**RemoveHost**](HostsAPI.md#RemoveHost) | **Delete** /api/servers/{id} | Delete a Host
[**StopHost**](HostsAPI.md#StopHost) | **Put** /api/servers/{id}/stop | Stop a Host
[**UpdateHost**](HostsAPI.md#UpdateHost) | **Put** /api/servers/{id} | Updating a Host



## AddBaremetalHost

> AddBaremetalHost200Response AddBaremetalHost(ctx).AddBaremetalHostRequest(addBaremetalHostRequest).Execute()

Add a Baremetal Host



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
	addBaremetalHostRequest := *openapiclient.NewAddBaremetalHostRequest() // AddBaremetalHostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.AddBaremetalHost(context.Background()).AddBaremetalHostRequest(addBaremetalHostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.AddBaremetalHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddBaremetalHost`: AddBaremetalHost200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.AddBaremetalHost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddBaremetalHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addBaremetalHostRequest** | [**AddBaremetalHostRequest**](AddBaremetalHostRequest.md) |  | 

### Return type

[**AddBaremetalHost200Response**](AddBaremetalHost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHost

> GetHost200Response GetHost(ctx, id).Execute()

Get a Specific Host



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
	id := openapiclient.getHost_id_parameter{Int64: new(int64)} // GetHostIdParameter | The `id` or `uuid` to identify the server record.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.GetHost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.GetHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHost`: GetHost200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.GetHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**GetHostIdParameter**](.md) | The &#x60;id&#x60; or &#x60;uuid&#x60; to identify the server record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetHost200Response**](GetHost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHosts

> ListHosts200Response ListHosts(ctx).Name(name).Phrase(phrase).ZoneId(zoneId).SiteId(siteId).ClusterId(clusterId).ParentServerId(parentServerId).Managed(managed).ServerType(serverType).PowerState(powerState).Ip(ip).Vm(vm).VmHypervisor(vmHypervisor).BareMetalHost(bareMetalHost).Status(status).AgentInstalled(agentInstalled).GuestAgentStatus(guestAgentStatus).Max(max).Offset(offset).LastUpdated(lastUpdated).CreatedBy(createdBy).TagsName(tagsName).Metadata(metadata).Uuid(uuid).ExternalId(externalId).InternalId(internalId).ExternalUniquelId(externalUniquelId).Stats(stats).Execute()

Get All Hosts



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
	name := "example" // string | Filter by name (optional)
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	zoneId := int64(3) // int64 | The Cloud ID (Zone ID) for Filtering (optional)
	siteId := int64(7) // int64 | The Group ID (Site ID) for Filtering (optional)
	clusterId := int64(135) // int64 | The Cluster ID(s) for filtering. Accepts multiple values. (optional)
	parentServerId := int64(36) // int64 | The Parent Server (Hypervisor) ID for Filtering (optional)
	managed := false // bool | Filter by managed (true) or unmanaged (false) (optional)
	serverType := "vmwareHypervisor" // string | Filter by server type code (optional)
	powerState := "off" // string | Filter by power status (optional)
	ip := "192.168.1.45" // string | Filter by IP address (optional)
	vm := false // bool | Filter to show only Virtual Machines (true) (optional)
	vmHypervisor := false // bool | Filter to show only VM Hypervisors (true) (optional)
	bareMetalHost := false // bool | Filter to show only Baremetal Servers (optional)
	status := "status_example" // string | Filter by status (optional)
	agentInstalled := true // bool | Filter by agent installed (true) (optional)
	guestAgentStatus := "connected" // string | Filter by Guest Agent Status. (optional)
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
	lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
	createdBy := int64(63) // int64 | The User ID for Filtering (optional)
	tagsName := "value" // string | Filter by tags (metadata). This allows filtering by a tag name and value(s)  (optional)
	metadata := "tags.env=qa&tags.env=test" // string | Alias for tags (optional)
	uuid := "uuid_example" // string | Filter by UUID, accepts multiple values (optional)
	externalId := "externalId_example" // string | Filter by External ID (optional)
	internalId := "internalId_example" // string | Filter by Internal ID (optional)
	externalUniquelId := "externalUniquelId_example" // string | Filter by External Unique ID (optional)
	stats := "stats_example" // string | This can be used to exclude the `stats` object in the response by passing `false` which can increase performance when returning a large number of servers. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.ListHosts(context.Background()).Name(name).Phrase(phrase).ZoneId(zoneId).SiteId(siteId).ClusterId(clusterId).ParentServerId(parentServerId).Managed(managed).ServerType(serverType).PowerState(powerState).Ip(ip).Vm(vm).VmHypervisor(vmHypervisor).BareMetalHost(bareMetalHost).Status(status).AgentInstalled(agentInstalled).GuestAgentStatus(guestAgentStatus).Max(max).Offset(offset).LastUpdated(lastUpdated).CreatedBy(createdBy).TagsName(tagsName).Metadata(metadata).Uuid(uuid).ExternalId(externalId).InternalId(internalId).ExternalUniquelId(externalUniquelId).Stats(stats).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.ListHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHosts`: ListHosts200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.ListHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **zoneId** | **int64** | The Cloud ID (Zone ID) for Filtering | 
 **siteId** | **int64** | The Group ID (Site ID) for Filtering | 
 **clusterId** | **int64** | The Cluster ID(s) for filtering. Accepts multiple values. | 
 **parentServerId** | **int64** | The Parent Server (Hypervisor) ID for Filtering | 
 **managed** | **bool** | Filter by managed (true) or unmanaged (false) | 
 **serverType** | **string** | Filter by server type code | 
 **powerState** | **string** | Filter by power status | 
 **ip** | **string** | Filter by IP address | 
 **vm** | **bool** | Filter to show only Virtual Machines (true) | 
 **vmHypervisor** | **bool** | Filter to show only VM Hypervisors (true) | 
 **bareMetalHost** | **bool** | Filter to show only Baremetal Servers | 
 **status** | **string** | Filter by status | 
 **agentInstalled** | **bool** | Filter by agent installed (true) | 
 **guestAgentStatus** | **string** | Filter by Guest Agent Status. | 
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 
 **createdBy** | **int64** | The User ID for Filtering | 
 **tagsName** | **string** | Filter by tags (metadata). This allows filtering by a tag name and value(s)  | 
 **metadata** | **string** | Alias for tags | 
 **uuid** | **string** | Filter by UUID, accepts multiple values | 
 **externalId** | **string** | Filter by External ID | 
 **internalId** | **string** | Filter by Internal ID | 
 **externalUniquelId** | **string** | Filter by External Unique ID | 
 **stats** | **string** | This can be used to exclude the &#x60;stats&#x60; object in the response by passing &#x60;false&#x60; which can increase performance when returning a large number of servers. | 

### Return type

[**ListHosts200Response**](ListHosts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveHost

> RemoveGroups200Response RemoveHost(ctx, id).RemoveResources(removeResources).RemoveInstances(removeInstances).PreserveVolumes(preserveVolumes).ReleaseFloatingIps(releaseFloatingIps).ReleaseEIPs(releaseEIPs).Force(force).Execute()

Delete a Host



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
	id := openapiclient.updateHost_id_parameter{Int64: new(int64)} // UpdateHostIdParameter | The `id` or `uuid` to identify the server record.
	removeResources := "off" // string | Remove Resources (optional) (default to "on")
	removeInstances := "on" // string | Remove Instances (optional) (default to "off")
	preserveVolumes := "on" // string | Preserve Volumes (optional) (default to "off")
	releaseFloatingIps := "off" // string | Release Floating IPs (optional) (default to "on")
	releaseEIPs := "off" // string | Alias for releaseFloatingIps (optional) (default to "on")
	force := "on" // string | Force Delete (optional) (default to "off")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.RemoveHost(context.Background(), id).RemoveResources(removeResources).RemoveInstances(removeInstances).PreserveVolumes(preserveVolumes).ReleaseFloatingIps(releaseFloatingIps).ReleaseEIPs(releaseEIPs).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.RemoveHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveHost`: RemoveGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.RemoveHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**UpdateHostIdParameter**](.md) | The &#x60;id&#x60; or &#x60;uuid&#x60; to identify the server record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeResources** | **string** | Remove Resources | [default to &quot;on&quot;]
 **removeInstances** | **string** | Remove Instances | [default to &quot;off&quot;]
 **preserveVolumes** | **string** | Preserve Volumes | [default to &quot;off&quot;]
 **releaseFloatingIps** | **string** | Release Floating IPs | [default to &quot;on&quot;]
 **releaseEIPs** | **string** | Alias for releaseFloatingIps | [default to &quot;on&quot;]
 **force** | **string** | Force Delete | [default to &quot;off&quot;]

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


## StopHost

> StopHost200Response StopHost(ctx, id).Execute()

Stop a Host



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
	id := openapiclient.updateHost_id_parameter{Int64: new(int64)} // UpdateHostIdParameter | The `id` or `uuid` to identify the server record.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.StopHost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.StopHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopHost`: StopHost200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.StopHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**UpdateHostIdParameter**](.md) | The &#x60;id&#x60; or &#x60;uuid&#x60; to identify the server record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StopHost200Response**](StopHost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHost

> GetHost200Response UpdateHost(ctx, id).UpdateHostRequest(updateHostRequest).Execute()

Updating a Host



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
	id := openapiclient.updateHost_id_parameter{Int64: new(int64)} // UpdateHostIdParameter | The `id` or `uuid` to identify the server record.
	updateHostRequest := *openapiclient.NewUpdateHostRequest() // UpdateHostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.UpdateHost(context.Background(), id).UpdateHostRequest(updateHostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.UpdateHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHost`: GetHost200Response
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.UpdateHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**UpdateHostIdParameter**](.md) | The &#x60;id&#x60; or &#x60;uuid&#x60; to identify the server record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateHostRequest** | [**UpdateHostRequest**](UpdateHostRequest.md) |  | 

### Return type

[**GetHost200Response**](GetHost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

