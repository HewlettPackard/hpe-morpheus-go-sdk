# \NetworksAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkDhcpRelay**](NetworksAPI.md#CreateNetworkDhcpRelay) | **Post** /api/networks/servers/{serverId}/dhcp-relays | Create a Network DHCP Relay
[**CreateNetworkDhcpServer**](NetworksAPI.md#CreateNetworkDhcpServer) | **Post** /api/networks/servers/{serverId}/dhcp-servers | Create a Network DHCP Server
[**CreateNetworkDomain**](NetworksAPI.md#CreateNetworkDomain) | **Post** /api/networks/domains | Create a Network Domain
[**CreateNetworkFirewallRule**](NetworksAPI.md#CreateNetworkFirewallRule) | **Post** /api/networks/servers/{serverId}/firewall-rules | Create a Network Firewall Rule
[**CreateNetworkFirewallRuleGroup**](NetworksAPI.md#CreateNetworkFirewallRuleGroup) | **Post** /api/networks/servers/{serverId}/firewall-rule-groups | Create a Network Firewall Rule Group
[**CreateNetworkGroup**](NetworksAPI.md#CreateNetworkGroup) | **Post** /api/networks/groups | Create a Network Group
[**CreateNetworkPool**](NetworksAPI.md#CreateNetworkPool) | **Post** /api/networks/pools | Create a Network Pool
[**CreateNetworkPoolIp**](NetworksAPI.md#CreateNetworkPoolIp) | **Post** /api/networks/pools/{id}/ips | Create a Network Pool IP Address
[**CreateNetworkPoolServer**](NetworksAPI.md#CreateNetworkPoolServer) | **Post** /api/networks/pool-servers | Create a Network Pool Server
[**CreateNetworkProxy**](NetworksAPI.md#CreateNetworkProxy) | **Post** /api/networks/proxies | Create a Network Proxy
[**CreateNetworkRouter**](NetworksAPI.md#CreateNetworkRouter) | **Post** /api/networks/routers | Create a Network Router
[**CreateNetworkRouterBgpNeighbor**](NetworksAPI.md#CreateNetworkRouterBgpNeighbor) | **Post** /api/networks/routers/{routerId}/bgp-neighbors | Create a Network Router BGP Neighbor
[**CreateNetworkRouterFirewallRule**](NetworksAPI.md#CreateNetworkRouterFirewallRule) | **Post** /api/networks/routers/{routerId}/firewall-rules | Create a Network Router Firewall Rule
[**CreateNetworkRouterFirewallRuleGroup**](NetworksAPI.md#CreateNetworkRouterFirewallRuleGroup) | **Post** /api/networks/routers/{routerId}/firewall-rule-groups | Create a Network Router Firewall Rule Group
[**CreateNetworkRouterNat**](NetworksAPI.md#CreateNetworkRouterNat) | **Post** /api/networks/routers/{routerId}/nats | Create a Network Router NAT
[**CreateNetworkRouterRoute**](NetworksAPI.md#CreateNetworkRouterRoute) | **Post** /api/networks/routers/{routerId}/routes | Create a Network Router Route
[**CreateNetworkServer**](NetworksAPI.md#CreateNetworkServer) | **Post** /api/networks/servers | Create a Network Server
[**CreateNetworkServerGroup**](NetworksAPI.md#CreateNetworkServerGroup) | **Post** /api/networks/servers/{serverId}/groups | Create a Network Server Group
[**CreateNetworkTransportZone**](NetworksAPI.md#CreateNetworkTransportZone) | **Post** /api/networks/servers/{serverId}/scopes | Create a Network Transport Zone
[**CreateNetworks**](NetworksAPI.md#CreateNetworks) | **Post** /api/networks | Create a Network
[**CreateStaticRoute**](NetworksAPI.md#CreateStaticRoute) | **Put** /api/networks/{id}/routes | Create a Network Static Route
[**CreateSubnet**](NetworksAPI.md#CreateSubnet) | **Post** /api/subnets | Create a Subnet
[**DeleteNetwork**](NetworksAPI.md#DeleteNetwork) | **Delete** /api/networks/{id} | Delete a Network
[**DeleteNetworkDhcpRelay**](NetworksAPI.md#DeleteNetworkDhcpRelay) | **Delete** /api/networks/servers/{serverId}/dhcp-relays/{id} | Delete a Network DHCP Relay
[**DeleteNetworkDhcpServer**](NetworksAPI.md#DeleteNetworkDhcpServer) | **Delete** /api/networks/servers/{serverId}/dhcp-servers/{id} | Delete a Network DHCP Server
[**DeleteNetworkDomain**](NetworksAPI.md#DeleteNetworkDomain) | **Delete** /api/networks/domains/{id} | Delete a Network Domain
[**DeleteNetworkFirewallRule**](NetworksAPI.md#DeleteNetworkFirewallRule) | **Delete** /api/networks/servers/{serverId}/firewall-rules/{id} | Delete a Network Firewall Rule
[**DeleteNetworkFirewallRuleGroup**](NetworksAPI.md#DeleteNetworkFirewallRuleGroup) | **Delete** /api/networks/servers/{serverId}/firewall-rule-groups/{id} | Delete a Network firewall rule group
[**DeleteNetworkGroup**](NetworksAPI.md#DeleteNetworkGroup) | **Delete** /api/networks/groups/{id} | Delete a Network Group
[**DeleteNetworkPool**](NetworksAPI.md#DeleteNetworkPool) | **Delete** /api/networks/pools/{id} | Delete a Network Pool
[**DeleteNetworkPoolIp**](NetworksAPI.md#DeleteNetworkPoolIp) | **Delete** /api/networks/pools/{networkPoolId}/ips/{id} | Delete a host record associated with an IP Address for a Specific Network Pool
[**DeleteNetworkPoolServer**](NetworksAPI.md#DeleteNetworkPoolServer) | **Delete** /api/networks/pool-servers/{id} | Delete a Network Pool Server
[**DeleteNetworkProxy**](NetworksAPI.md#DeleteNetworkProxy) | **Delete** /api/networks/proxies/{id} | Delete a Network Proxy
[**DeleteNetworkRouter**](NetworksAPI.md#DeleteNetworkRouter) | **Delete** /api/networks/routers/{id} | Delete a Network Router
[**DeleteNetworkRouterBgpNeighbor**](NetworksAPI.md#DeleteNetworkRouterBgpNeighbor) | **Delete** /api/networks/routers/{routerId}/bgp-neighbors/{id} | Delete a Network Router BGP Neighbor
[**DeleteNetworkRouterFirewallRule**](NetworksAPI.md#DeleteNetworkRouterFirewallRule) | **Delete** /api/networks/routers/{routerId}/firewall-rules/{id} | Delete a Network Router Firewall Rule
[**DeleteNetworkRouterFirewallRuleGroup**](NetworksAPI.md#DeleteNetworkRouterFirewallRuleGroup) | **Delete** /api/networks/routers/{routerId}/firewall-rule-groups/{id} | Delete a Network Router firewall rule group
[**DeleteNetworkRouterNat**](NetworksAPI.md#DeleteNetworkRouterNat) | **Delete** /api/networks/routers/{routerId}/nats/{id} | Delete a Network Router NAT
[**DeleteNetworkRouterRoute**](NetworksAPI.md#DeleteNetworkRouterRoute) | **Delete** /api/networks/routers/{routerId}/routes/{id} | Delete a Network Router Route
[**DeleteNetworkServer**](NetworksAPI.md#DeleteNetworkServer) | **Delete** /api/networks/servers/{id} | Delete a Network Server
[**DeleteNetworkServerGroup**](NetworksAPI.md#DeleteNetworkServerGroup) | **Delete** /api/networks/servers/{serverId}/groups/{id} | Delete a Network Server Group
[**DeleteNetworkTransportZone**](NetworksAPI.md#DeleteNetworkTransportZone) | **Delete** /api/networks/servers/{serverId}/scopes/{id} | Delete a Network Transport Zone
[**DeleteStaticRoute**](NetworksAPI.md#DeleteStaticRoute) | **Delete** /api/networks/{id}/routes/{routeId} | Delete a Network Static Route
[**DeleteSubnet**](NetworksAPI.md#DeleteSubnet) | **Delete** /api/subnets/{id} | Delete a Subnet
[**GetAllNetworkFloatingIps**](NetworksAPI.md#GetAllNetworkFloatingIps) | **Get** /api/networks/floating-ips | Get All Floating IPs
[**GetNetwork**](NetworksAPI.md#GetNetwork) | **Get** /api/networks/{id} | Get a Specific Network
[**GetNetworkDhcpRelay**](NetworksAPI.md#GetNetworkDhcpRelay) | **Get** /api/networks/servers/{serverId}/dhcp-relays/{id} | Get a Specific Network DHCP Relay
[**GetNetworkDhcpRelays**](NetworksAPI.md#GetNetworkDhcpRelays) | **Get** /api/networks/servers/{serverId}/dhcp-relays | Get all Network DHCP Relays for Network Relay
[**GetNetworkDhcpServer**](NetworksAPI.md#GetNetworkDhcpServer) | **Get** /api/networks/servers/{serverId}/dhcp-servers/{id} | Get a Specific Network DHCP Server
[**GetNetworkDhcpServers**](NetworksAPI.md#GetNetworkDhcpServers) | **Get** /api/networks/servers/{serverId}/dhcp-servers | Get all Network DHCP Servers for Network Server
[**GetNetworkDomain**](NetworksAPI.md#GetNetworkDomain) | **Get** /api/networks/domains/{id} | Get a Specific Network Domain
[**GetNetworkDomains**](NetworksAPI.md#GetNetworkDomains) | **Get** /api/networks/domains | Get all Network Domains
[**GetNetworkEdgeCluster**](NetworksAPI.md#GetNetworkEdgeCluster) | **Get** /api/networks/servers/{serverId}/edge-clusters/{id} | Get a Specific Network Edge Cluster
[**GetNetworkEdgeClusters**](NetworksAPI.md#GetNetworkEdgeClusters) | **Get** /api/networks/servers/{serverId}/edge-clusters | Get all Network Edge Clusters for Network Server
[**GetNetworkFirewallRule**](NetworksAPI.md#GetNetworkFirewallRule) | **Get** /api/networks/servers/{serverId}/firewall-rules/{id} | Get a Specific Network Firewall Rule
[**GetNetworkFirewallRuleGroup**](NetworksAPI.md#GetNetworkFirewallRuleGroup) | **Get** /api/networks/servers/{serverId}/firewall-rule-groups/{id} | Get a Specific Network Firewall Rule Group
[**GetNetworkFirewallRuleGroups**](NetworksAPI.md#GetNetworkFirewallRuleGroups) | **Get** /api/networks/servers/{serverId}/firewall-rule-groups | Get all Network Firewall Rule Groups for Network Server
[**GetNetworkFirewallRules**](NetworksAPI.md#GetNetworkFirewallRules) | **Get** /api/networks/servers/{serverId}/firewall-rules | Get all Network Firewall Rules for Network Server
[**GetNetworkFloatingIp**](NetworksAPI.md#GetNetworkFloatingIp) | **Get** /api/networks/floating-ips/{id} | Get a Specific Floating IP
[**GetNetworkGroup**](NetworksAPI.md#GetNetworkGroup) | **Get** /api/networks/groups/{id} | Get a Specific Network Group
[**GetNetworkGroups**](NetworksAPI.md#GetNetworkGroups) | **Get** /api/networks/groups | Get all Network Groups
[**GetNetworkPool**](NetworksAPI.md#GetNetworkPool) | **Get** /api/networks/pools/{id} | Get a Specific Network Pool
[**GetNetworkPoolIp**](NetworksAPI.md#GetNetworkPoolIp) | **Get** /api/networks/pools/{networkPoolId}/ips/{id} | Get a Specific IP Address for a Specific Network Pool
[**GetNetworkPoolIps**](NetworksAPI.md#GetNetworkPoolIps) | **Get** /api/networks/pools/{id}/ips | Get all IP Addresses for a Specific Network Pool
[**GetNetworkPoolServer**](NetworksAPI.md#GetNetworkPoolServer) | **Get** /api/networks/pool-servers/{id} | Get a Specific Network Pool Server
[**GetNetworkPoolServerType**](NetworksAPI.md#GetNetworkPoolServerType) | **Get** /api/networks/pool-server-types/{id} | Retrieves a Specific Network Pool Server Type
[**GetNetworkPools**](NetworksAPI.md#GetNetworkPools) | **Get** /api/networks/pools | Get all Network Pools
[**GetNetworkProxies**](NetworksAPI.md#GetNetworkProxies) | **Get** /api/networks/proxies | Get all Network Proxies
[**GetNetworkProxy**](NetworksAPI.md#GetNetworkProxy) | **Get** /api/networks/proxies/{id} | Get a Specific Network Proxy
[**GetNetworkRouter**](NetworksAPI.md#GetNetworkRouter) | **Get** /api/networks/routers/{id} | Get a Specific Network Router
[**GetNetworkRouterBgpNeighbor**](NetworksAPI.md#GetNetworkRouterBgpNeighbor) | **Get** /api/networks/routers/{routerId}/bgp-neighbors/{id} | Get a Network Router BGP Neighbor
[**GetNetworkRouterFirewallRule**](NetworksAPI.md#GetNetworkRouterFirewallRule) | **Get** /api/networks/routers/{routerId}/firewall-rules/{id} | Get a Firewall Rule for Network Router
[**GetNetworkRouterFirewallRuleGroup**](NetworksAPI.md#GetNetworkRouterFirewallRuleGroup) | **Get** /api/networks/routers/{routerId}/firewall-rule-groups/{id} | Get a Firewall Rule Group for Network Router
[**GetNetworkRouterFirewallRuleGroups**](NetworksAPI.md#GetNetworkRouterFirewallRuleGroups) | **Get** /api/networks/routers/{routerId}/firewall-rule-groups | Get all Network Firewall Rule Groups for Network Router
[**GetNetworkRouterNat**](NetworksAPI.md#GetNetworkRouterNat) | **Get** /api/networks/routers/{routerId}/nats/{id} | Get a Network Router NAT
[**GetNetworkRouterRoute**](NetworksAPI.md#GetNetworkRouterRoute) | **Get** /api/networks/routers/{routerId}/routes/{id} | Get a Route for Network Router
[**GetNetworkRouterType**](NetworksAPI.md#GetNetworkRouterType) | **Get** /api/network-router-types/{id} | Get a Specific Network Router Type
[**GetNetworkRouters**](NetworksAPI.md#GetNetworkRouters) | **Get** /api/networks/routers | Get all Network Routers
[**GetNetworkRoutersBgpNeighbors**](NetworksAPI.md#GetNetworkRoutersBgpNeighbors) | **Get** /api/networks/routers/{routerId}/bgp-neighbors | Get all BGP Neighbors for Network Router
[**GetNetworkRoutersFirewallRules**](NetworksAPI.md#GetNetworkRoutersFirewallRules) | **Get** /api/networks/routers/{routerId}/firewall-rules | Get all Firewall Rules for Network Router
[**GetNetworkRoutersNats**](NetworksAPI.md#GetNetworkRoutersNats) | **Get** /api/networks/routers/{routerId}/nats | Get all Network Router NATs for Network Router
[**GetNetworkRoutersRoutes**](NetworksAPI.md#GetNetworkRoutersRoutes) | **Get** /api/networks/routers/{routerId}/routes | Get all Routes for Network Router
[**GetNetworkServer**](NetworksAPI.md#GetNetworkServer) | **Get** /api/networks/servers/{id} | Get a Specific Network Server
[**GetNetworkServerGroup**](NetworksAPI.md#GetNetworkServerGroup) | **Get** /api/networks/servers/{serverId}/groups/{id} | Get a specific Network Server Group
[**GetNetworkSubnets**](NetworksAPI.md#GetNetworkSubnets) | **Get** /api/networks/{id}/subnets | Get Subnets for a Network
[**GetNetworkTransportZone**](NetworksAPI.md#GetNetworkTransportZone) | **Get** /api/networks/servers/{serverId}/scopes/{id} | Get a Specific Network Transport Zone
[**GetNetworkTransportZones**](NetworksAPI.md#GetNetworkTransportZones) | **Get** /api/networks/servers/{serverId}/scopes | Get all Network Transport Zones for Network Server
[**GetNetworkType**](NetworksAPI.md#GetNetworkType) | **Get** /api/network-types/{id} | Get a Specific Network Type
[**GetStaticRoute**](NetworksAPI.md#GetStaticRoute) | **Get** /api/networks/{id}/routes/{routeId} | Get a Static Route for a Network
[**GetStaticRoutes**](NetworksAPI.md#GetStaticRoutes) | **Get** /api/networks/{id}/routes | Get all Static Routes for a Network
[**GetSubnet**](NetworksAPI.md#GetSubnet) | **Get** /api/subnets/{id} | Get a Specific Subnet
[**ListNetworkPoolServerTypes**](NetworksAPI.md#ListNetworkPoolServerTypes) | **Get** /api/networks/pool-server-types | Get All Network Pool Server Types
[**ListNetworkPoolServers**](NetworksAPI.md#ListNetworkPoolServers) | **Get** /api/networks/pool-servers | Get All Network Pool Servers
[**ListNetworkRouterTypes**](NetworksAPI.md#ListNetworkRouterTypes) | **Get** /api/network-router-types | Get all Network Router Types
[**ListNetworkServerGroups**](NetworksAPI.md#ListNetworkServerGroups) | **Get** /api/networks/servers/{serverId}/groups | Get all Network Server Groups for Network Server
[**ListNetworkServers**](NetworksAPI.md#ListNetworkServers) | **Get** /api/networks/servers | Get All Network Servers
[**ListNetworkServices**](NetworksAPI.md#ListNetworkServices) | **Get** /api/networks/services | Get All Network Services
[**ListNetworkTypes**](NetworksAPI.md#ListNetworkTypes) | **Get** /api/network-types | Network Types
[**ListNetworks**](NetworksAPI.md#ListNetworks) | **Get** /api/networks | Get All Networks
[**ListSubnetTypes**](NetworksAPI.md#ListSubnetTypes) | **Get** /api/subnet-types | Get All Subnet Types
[**ListSubnets**](NetworksAPI.md#ListSubnets) | **Get** /api/subnets | Get All Subnets
[**RefreshNetworkServer**](NetworksAPI.md#RefreshNetworkServer) | **Post** /api/networks/servers/{id}/refresh | Refresh a Network Server/Integration
[**ReleaseNetworkFloatingIp**](NetworksAPI.md#ReleaseNetworkFloatingIp) | **Put** /api/networks/floating-ips/{id}/release | Release a Floating IP
[**UpdateNetwork**](NetworksAPI.md#UpdateNetwork) | **Put** /api/networks/{id} | Update a Network
[**UpdateNetworkDhcpRelay**](NetworksAPI.md#UpdateNetworkDhcpRelay) | **Put** /api/networks/servers/{serverId}/dhcp-relays/{id} | Update a Network DHCP Relay
[**UpdateNetworkDhcpServer**](NetworksAPI.md#UpdateNetworkDhcpServer) | **Put** /api/networks/servers/{serverId}/dhcp-servers/{id} | Update a Network DHCP Server
[**UpdateNetworkDomain**](NetworksAPI.md#UpdateNetworkDomain) | **Put** /api/networks/domains/{id} | Update a Network Domain
[**UpdateNetworkEdgeCluster**](NetworksAPI.md#UpdateNetworkEdgeCluster) | **Put** /api/networks/servers/{serverId}/edge-clusters/{id} | Update a Network Edge Cluster
[**UpdateNetworkFirewallRule**](NetworksAPI.md#UpdateNetworkFirewallRule) | **Put** /api/networks/servers/{serverId}/firewall-rules/{id} | Update a Network Firewall Rule
[**UpdateNetworkFirewallRuleGroup**](NetworksAPI.md#UpdateNetworkFirewallRuleGroup) | **Put** /api/networks/servers/{serverId}/firewall-rule-groups/{id} | Update a Network Firewall Rule Group
[**UpdateNetworkGroup**](NetworksAPI.md#UpdateNetworkGroup) | **Put** /api/networks/groups/{id} | Update a Network Group
[**UpdateNetworkPool**](NetworksAPI.md#UpdateNetworkPool) | **Put** /api/networks/pools/{id} | Update a Network Pool
[**UpdateNetworkPoolServer**](NetworksAPI.md#UpdateNetworkPoolServer) | **Put** /api/networks/pool-servers/{id} | Update a Network Pool Server
[**UpdateNetworkProxy**](NetworksAPI.md#UpdateNetworkProxy) | **Put** /api/networks/proxies/{id} | Update a Network Proxy
[**UpdateNetworkRouter**](NetworksAPI.md#UpdateNetworkRouter) | **Put** /api/networks/routers/{id} | Update a Network Router
[**UpdateNetworkRouterBgpNeighbor**](NetworksAPI.md#UpdateNetworkRouterBgpNeighbor) | **Put** /api/networks/routers/{routerId}/bgp-neighbors/{id} | Update Network Router BGP Neighbor
[**UpdateNetworkRouterFirewallRule**](NetworksAPI.md#UpdateNetworkRouterFirewallRule) | **Put** /api/networks/routers/{routerId}/firewall-rules/{id} | Update a Network Router Firewall Rule
[**UpdateNetworkRouterFirewallRuleGroup**](NetworksAPI.md#UpdateNetworkRouterFirewallRuleGroup) | **Put** /api/networks/routers/{routerId}/firewall-rule-groups/{id} | Update a Network Router Firewall Rule Group
[**UpdateNetworkRouterNat**](NetworksAPI.md#UpdateNetworkRouterNat) | **Put** /api/networks/routers/{routerId}/nats/{id} | Update Network Router NAT
[**UpdateNetworkRouterPermissions**](NetworksAPI.md#UpdateNetworkRouterPermissions) | **Put** /api/networks/routers/{routerId}/permissions | Update Network Router Permissions
[**UpdateNetworkServer**](NetworksAPI.md#UpdateNetworkServer) | **Put** /api/networks/servers/{id} | Update a Network Server
[**UpdateNetworkServerGroup**](NetworksAPI.md#UpdateNetworkServerGroup) | **Put** /api/networks/servers/{serverId}/groups/{id} | Update a Network Server Group
[**UpdateNetworkTransportZone**](NetworksAPI.md#UpdateNetworkTransportZone) | **Put** /api/networks/servers/{serverId}/scopes/{id} | Update a Network Transport Zone
[**UpdateStaticRoute**](NetworksAPI.md#UpdateStaticRoute) | **Put** /api/networks/{id}/routes/{routeId} | Update a Network Static Route
[**UpdateSubnet**](NetworksAPI.md#UpdateSubnet) | **Put** /api/subnets/{id} | Update a Subnet



## CreateNetworkDhcpRelay

> AddClusterLayouts200Response CreateNetworkDhcpRelay(ctx, serverId).CreateNetworkDhcpRelayRequest(createNetworkDhcpRelayRequest).Execute()

Create a Network DHCP Relay



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	serverId := float32(4) // float32 | Server ID
	createNetworkDhcpRelayRequest := *openapiclient.NewCreateNetworkDhcpRelayRequest() // CreateNetworkDhcpRelayRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkDhcpRelay(context.Background(), serverId).CreateNetworkDhcpRelayRequest(createNetworkDhcpRelayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkDhcpRelay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkDhcpRelay`: AddClusterLayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkDhcpRelay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkDhcpRelayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkDhcpRelayRequest** | [**CreateNetworkDhcpRelayRequest**](CreateNetworkDhcpRelayRequest.md) |  | 

### Return type

[**AddClusterLayouts200Response**](AddClusterLayouts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkDhcpServer

> AddClusterLayouts200Response CreateNetworkDhcpServer(ctx, serverId).CreateNetworkDhcpServerRequest(createNetworkDhcpServerRequest).Execute()

Create a Network DHCP Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	serverId := float32(4) // float32 | Server ID
	createNetworkDhcpServerRequest := *openapiclient.NewCreateNetworkDhcpServerRequest() // CreateNetworkDhcpServerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkDhcpServer(context.Background(), serverId).CreateNetworkDhcpServerRequest(createNetworkDhcpServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkDhcpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkDhcpServer`: AddClusterLayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkDhcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkDhcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkDhcpServerRequest** | [**CreateNetworkDhcpServerRequest**](CreateNetworkDhcpServerRequest.md) |  | 

### Return type

[**AddClusterLayouts200Response**](AddClusterLayouts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkDomain

> CreateNetworkDomain200Response CreateNetworkDomain(ctx).CreateNetworkDomainRequest(createNetworkDomainRequest).Execute()

Create a Network Domain



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	createNetworkDomainRequest := *openapiclient.NewCreateNetworkDomainRequest() // CreateNetworkDomainRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkDomain(context.Background()).CreateNetworkDomainRequest(createNetworkDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkDomain`: CreateNetworkDomain200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkDomainRequest** | [**CreateNetworkDomainRequest**](CreateNetworkDomainRequest.md) |  | 

### Return type

[**CreateNetworkDomain200Response**](CreateNetworkDomain200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkFirewallRule

> AddClusterLayouts200Response CreateNetworkFirewallRule(ctx, serverId).CreateNetworkFirewallRuleRequest(createNetworkFirewallRuleRequest).Execute()

Create a Network Firewall Rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	serverId := float32(4) // float32 | Server ID
	createNetworkFirewallRuleRequest := *openapiclient.NewCreateNetworkFirewallRuleRequest() // CreateNetworkFirewallRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkFirewallRule(context.Background(), serverId).CreateNetworkFirewallRuleRequest(createNetworkFirewallRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkFirewallRule`: AddClusterLayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkFirewallRuleRequest** | [**CreateNetworkFirewallRuleRequest**](CreateNetworkFirewallRuleRequest.md) |  | 

### Return type

[**AddClusterLayouts200Response**](AddClusterLayouts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkFirewallRuleGroup

> AddClusterLayouts200Response CreateNetworkFirewallRuleGroup(ctx, serverId).CreateNetworkRouterFirewallRuleGroupRequest(createNetworkRouterFirewallRuleGroupRequest).Execute()

Create a Network Firewall Rule Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	serverId := float32(4) // float32 | Server ID
	createNetworkRouterFirewallRuleGroupRequest := *openapiclient.NewCreateNetworkRouterFirewallRuleGroupRequest() // CreateNetworkRouterFirewallRuleGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkFirewallRuleGroup(context.Background(), serverId).CreateNetworkRouterFirewallRuleGroupRequest(createNetworkRouterFirewallRuleGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkFirewallRuleGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkFirewallRuleGroup`: AddClusterLayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkFirewallRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFirewallRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkRouterFirewallRuleGroupRequest** | [**CreateNetworkRouterFirewallRuleGroupRequest**](CreateNetworkRouterFirewallRuleGroupRequest.md) |  | 

### Return type

[**AddClusterLayouts200Response**](AddClusterLayouts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkGroup

> AddClusterLayouts200Response CreateNetworkGroup(ctx).CreateNetworkGroupRequest(createNetworkGroupRequest).Execute()

Create a Network Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	createNetworkGroupRequest := *openapiclient.NewCreateNetworkGroupRequest() // CreateNetworkGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkGroup(context.Background()).CreateNetworkGroupRequest(createNetworkGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkGroup`: AddClusterLayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkGroupRequest** | [**CreateNetworkGroupRequest**](CreateNetworkGroupRequest.md) |  | 

### Return type

[**AddClusterLayouts200Response**](AddClusterLayouts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkPool

> CreateNetworkPool200Response CreateNetworkPool(ctx).CreateNetworkPoolRequest(createNetworkPoolRequest).Execute()

Create a Network Pool



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	createNetworkPoolRequest := *openapiclient.NewCreateNetworkPoolRequest() // CreateNetworkPoolRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkPool(context.Background()).CreateNetworkPoolRequest(createNetworkPoolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkPool`: CreateNetworkPool200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkPool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkPoolRequest** | [**CreateNetworkPoolRequest**](CreateNetworkPoolRequest.md) |  | 

### Return type

[**CreateNetworkPool200Response**](CreateNetworkPool200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkPoolIp

> CreateNetworkPoolIp200Response CreateNetworkPoolIp(ctx, id).CreateNetworkPoolIpRequest(createNetworkPoolIpRequest).Execute()

Create a Network Pool IP Address



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	createNetworkPoolIpRequest := *openapiclient.NewCreateNetworkPoolIpRequest() // CreateNetworkPoolIpRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkPoolIp(context.Background(), id).CreateNetworkPoolIpRequest(createNetworkPoolIpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkPoolIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkPoolIp`: CreateNetworkPoolIp200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkPoolIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkPoolIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkPoolIpRequest** | [**CreateNetworkPoolIpRequest**](CreateNetworkPoolIpRequest.md) |  | 

### Return type

[**CreateNetworkPoolIp200Response**](CreateNetworkPoolIp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkPoolServer

> CreateNetworkPoolServer200Response CreateNetworkPoolServer(ctx).CreateNetworkPoolServerRequest(createNetworkPoolServerRequest).Execute()

Create a Network Pool Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	createNetworkPoolServerRequest := *openapiclient.NewCreateNetworkPoolServerRequest() // CreateNetworkPoolServerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkPoolServer(context.Background()).CreateNetworkPoolServerRequest(createNetworkPoolServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkPoolServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkPoolServer`: CreateNetworkPoolServer200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkPoolServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkPoolServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkPoolServerRequest** | [**CreateNetworkPoolServerRequest**](CreateNetworkPoolServerRequest.md) |  | 

### Return type

[**CreateNetworkPoolServer200Response**](CreateNetworkPoolServer200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkProxy

> CreateNetworkProxy200Response CreateNetworkProxy(ctx).CreateNetworkProxyRequest(createNetworkProxyRequest).Execute()

Create a Network Proxy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	createNetworkProxyRequest := *openapiclient.NewCreateNetworkProxyRequest() // CreateNetworkProxyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkProxy(context.Background()).CreateNetworkProxyRequest(createNetworkProxyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkProxy`: CreateNetworkProxy200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkProxy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkProxyRequest** | [**CreateNetworkProxyRequest**](CreateNetworkProxyRequest.md) |  | 

### Return type

[**CreateNetworkProxy200Response**](CreateNetworkProxy200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkRouter

> AddClusterLayouts200Response CreateNetworkRouter(ctx).CreateNetworkRouterRequest(createNetworkRouterRequest).Execute()

Create a Network Router



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	createNetworkRouterRequest := *openapiclient.NewCreateNetworkRouterRequest() // CreateNetworkRouterRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkRouter(context.Background()).CreateNetworkRouterRequest(createNetworkRouterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkRouter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkRouter`: AddClusterLayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkRouter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkRouterRequest** | [**CreateNetworkRouterRequest**](CreateNetworkRouterRequest.md) |  | 

### Return type

[**AddClusterLayouts200Response**](AddClusterLayouts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkRouterBgpNeighbor

> AddClusterLayouts200Response CreateNetworkRouterBgpNeighbor(ctx, routerId).CreateNetworkRouterBgpNeighborRequest(createNetworkRouterBgpNeighborRequest).Execute()

Create a Network Router BGP Neighbor



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	routerId := float32(4) // float32 | Router ID
	createNetworkRouterBgpNeighborRequest := *openapiclient.NewCreateNetworkRouterBgpNeighborRequest() // CreateNetworkRouterBgpNeighborRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkRouterBgpNeighbor(context.Background(), routerId).CreateNetworkRouterBgpNeighborRequest(createNetworkRouterBgpNeighborRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkRouterBgpNeighbor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkRouterBgpNeighbor`: AddClusterLayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkRouterBgpNeighbor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkRouterBgpNeighborRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkRouterBgpNeighborRequest** | [**CreateNetworkRouterBgpNeighborRequest**](CreateNetworkRouterBgpNeighborRequest.md) |  | 

### Return type

[**AddClusterLayouts200Response**](AddClusterLayouts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkRouterFirewallRule

> AddClusterLayouts200Response CreateNetworkRouterFirewallRule(ctx, routerId).CreateNetworkRouterFirewallRuleRequest(createNetworkRouterFirewallRuleRequest).Execute()

Create a Network Router Firewall Rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	routerId := float32(4) // float32 | Router ID
	createNetworkRouterFirewallRuleRequest := *openapiclient.NewCreateNetworkRouterFirewallRuleRequest() // CreateNetworkRouterFirewallRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkRouterFirewallRule(context.Background(), routerId).CreateNetworkRouterFirewallRuleRequest(createNetworkRouterFirewallRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkRouterFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkRouterFirewallRule`: AddClusterLayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkRouterFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkRouterFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkRouterFirewallRuleRequest** | [**CreateNetworkRouterFirewallRuleRequest**](CreateNetworkRouterFirewallRuleRequest.md) |  | 

### Return type

[**AddClusterLayouts200Response**](AddClusterLayouts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkRouterFirewallRuleGroup

> AddClusterLayouts200Response CreateNetworkRouterFirewallRuleGroup(ctx, routerId).CreateNetworkRouterFirewallRuleGroupRequest(createNetworkRouterFirewallRuleGroupRequest).Execute()

Create a Network Router Firewall Rule Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	routerId := float32(4) // float32 | Router ID
	createNetworkRouterFirewallRuleGroupRequest := *openapiclient.NewCreateNetworkRouterFirewallRuleGroupRequest() // CreateNetworkRouterFirewallRuleGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkRouterFirewallRuleGroup(context.Background(), routerId).CreateNetworkRouterFirewallRuleGroupRequest(createNetworkRouterFirewallRuleGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkRouterFirewallRuleGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkRouterFirewallRuleGroup`: AddClusterLayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkRouterFirewallRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkRouterFirewallRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkRouterFirewallRuleGroupRequest** | [**CreateNetworkRouterFirewallRuleGroupRequest**](CreateNetworkRouterFirewallRuleGroupRequest.md) |  | 

### Return type

[**AddClusterLayouts200Response**](AddClusterLayouts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkRouterNat

> AddClusterLayouts200Response CreateNetworkRouterNat(ctx, routerId).CreateNetworkRouterNatRequest(createNetworkRouterNatRequest).Execute()

Create a Network Router NAT



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	routerId := float32(4) // float32 | Router ID
	createNetworkRouterNatRequest := *openapiclient.NewCreateNetworkRouterNatRequest() // CreateNetworkRouterNatRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkRouterNat(context.Background(), routerId).CreateNetworkRouterNatRequest(createNetworkRouterNatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkRouterNat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkRouterNat`: AddClusterLayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkRouterNat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkRouterNatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkRouterNatRequest** | [**CreateNetworkRouterNatRequest**](CreateNetworkRouterNatRequest.md) |  | 

### Return type

[**AddClusterLayouts200Response**](AddClusterLayouts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkRouterRoute

> AddClusterLayouts200Response CreateNetworkRouterRoute(ctx, routerId).CreateNetworkRouterRouteRequest(createNetworkRouterRouteRequest).Execute()

Create a Network Router Route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	routerId := float32(4) // float32 | Router ID
	createNetworkRouterRouteRequest := *openapiclient.NewCreateNetworkRouterRouteRequest() // CreateNetworkRouterRouteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkRouterRoute(context.Background(), routerId).CreateNetworkRouterRouteRequest(createNetworkRouterRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkRouterRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkRouterRoute`: AddClusterLayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkRouterRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkRouterRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkRouterRouteRequest** | [**CreateNetworkRouterRouteRequest**](CreateNetworkRouterRouteRequest.md) |  | 

### Return type

[**AddClusterLayouts200Response**](AddClusterLayouts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkServer

> CreateNetworkServer200Response CreateNetworkServer(ctx).CreateNetworkServerRequest(createNetworkServerRequest).Execute()

Create a Network Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	createNetworkServerRequest := *openapiclient.NewCreateNetworkServerRequest() // CreateNetworkServerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkServer(context.Background()).CreateNetworkServerRequest(createNetworkServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkServer`: CreateNetworkServer200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkServerRequest** | [**CreateNetworkServerRequest**](CreateNetworkServerRequest.md) |  | 

### Return type

[**CreateNetworkServer200Response**](CreateNetworkServer200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkServerGroup

> AddClusterLayouts200Response CreateNetworkServerGroup(ctx, serverId).CreateNetworkServerGroupRequest(createNetworkServerGroupRequest).Execute()

Create a Network Server Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	serverId := float32(4) // float32 | Server ID
	createNetworkServerGroupRequest := *openapiclient.NewCreateNetworkServerGroupRequest() // CreateNetworkServerGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkServerGroup(context.Background(), serverId).CreateNetworkServerGroupRequest(createNetworkServerGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkServerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkServerGroup`: AddClusterLayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkServerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkServerGroupRequest** | [**CreateNetworkServerGroupRequest**](CreateNetworkServerGroupRequest.md) |  | 

### Return type

[**AddClusterLayouts200Response**](AddClusterLayouts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkTransportZone

> AddClusterLayouts200Response CreateNetworkTransportZone(ctx, serverId).CreateNetworkTransportZoneRequest(createNetworkTransportZoneRequest).Execute()

Create a Network Transport Zone



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	serverId := float32(4) // float32 | Server ID
	createNetworkTransportZoneRequest := *openapiclient.NewCreateNetworkTransportZoneRequest() // CreateNetworkTransportZoneRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkTransportZone(context.Background(), serverId).CreateNetworkTransportZoneRequest(createNetworkTransportZoneRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkTransportZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkTransportZone`: AddClusterLayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkTransportZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkTransportZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkTransportZoneRequest** | [**CreateNetworkTransportZoneRequest**](CreateNetworkTransportZoneRequest.md) |  | 

### Return type

[**AddClusterLayouts200Response**](AddClusterLayouts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworks

> CreateNetworks200Response CreateNetworks(ctx).CreateNetworksRequest(createNetworksRequest).Execute()

Create a Network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	createNetworksRequest := *openapiclient.NewCreateNetworksRequest() // CreateNetworksRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworks(context.Background()).CreateNetworksRequest(createNetworksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworks`: CreateNetworks200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworksRequest** | [**CreateNetworksRequest**](CreateNetworksRequest.md) |  | 

### Return type

[**CreateNetworks200Response**](CreateNetworks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStaticRoute

> AddClusterLayouts200Response CreateStaticRoute(ctx, id).CreateStaticRouteRequest(createStaticRouteRequest).Execute()

Create a Network Static Route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	createStaticRouteRequest := *openapiclient.NewCreateStaticRouteRequest() // CreateStaticRouteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateStaticRoute(context.Background(), id).CreateStaticRouteRequest(createStaticRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStaticRoute`: AddClusterLayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createStaticRouteRequest** | [**CreateStaticRouteRequest**](CreateStaticRouteRequest.md) |  | 

### Return type

[**AddClusterLayouts200Response**](AddClusterLayouts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubnet

> CreateSubnet200Response CreateSubnet(ctx).CreateSubnetRequest(createSubnetRequest).Execute()

Create a Subnet



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	createSubnetRequest := *openapiclient.NewCreateSubnetRequest() // CreateSubnetRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateSubnet(context.Background()).CreateSubnetRequest(createSubnetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubnet`: CreateSubnet200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateSubnet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSubnetRequest** | [**CreateSubnetRequest**](CreateSubnetRequest.md) |  | 

### Return type

[**CreateSubnet200Response**](CreateSubnet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetwork

> DeleteAlerts200Response DeleteNetwork(ctx, id).Execute()

Delete a Network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetwork(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetwork`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRequest struct via the builder pattern


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


## DeleteNetworkDhcpRelay

> DeleteAlerts200Response DeleteNetworkDhcpRelay(ctx, id, serverId).Execute()

Delete a Network DHCP Relay



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	serverId := float32(4) // float32 | Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworkDhcpRelay(context.Background(), id, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkDhcpRelay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkDhcpRelay`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworkDhcpRelay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkDhcpRelayRequest struct via the builder pattern


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


## DeleteNetworkDhcpServer

> DeleteAlerts200Response DeleteNetworkDhcpServer(ctx, id, serverId).Execute()

Delete a Network DHCP Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	serverId := float32(4) // float32 | Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworkDhcpServer(context.Background(), id, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkDhcpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkDhcpServer`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworkDhcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkDhcpServerRequest struct via the builder pattern


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


## DeleteNetworkDomain

> DeleteAlerts200Response DeleteNetworkDomain(ctx, id).Execute()

Delete a Network Domain



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworkDomain(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkDomain`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworkDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkDomainRequest struct via the builder pattern


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


## DeleteNetworkFirewallRule

> DeleteAlerts200Response DeleteNetworkFirewallRule(ctx, id, serverId).Execute()

Delete a Network Firewall Rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	serverId := float32(4) // float32 | Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworkFirewallRule(context.Background(), id, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkFirewallRule`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworkFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkFirewallRuleRequest struct via the builder pattern


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


## DeleteNetworkFirewallRuleGroup

> DeleteAlerts200Response DeleteNetworkFirewallRuleGroup(ctx, id, serverId).Execute()

Delete a Network firewall rule group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	serverId := float32(4) // float32 | Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworkFirewallRuleGroup(context.Background(), id, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkFirewallRuleGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkFirewallRuleGroup`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworkFirewallRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkFirewallRuleGroupRequest struct via the builder pattern


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


## DeleteNetworkGroup

> DeleteAlerts200Response DeleteNetworkGroup(ctx, id).Execute()

Delete a Network Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworkGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkGroup`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworkGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkGroupRequest struct via the builder pattern


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


## DeleteNetworkPool

> DeleteAlerts200Response DeleteNetworkPool(ctx, id).Execute()

Delete a Network Pool



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworkPool(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkPool`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworkPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkPoolRequest struct via the builder pattern


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


## DeleteNetworkPoolIp

> DeleteAlerts200Response DeleteNetworkPoolIp(ctx, networkPoolId, id).Execute()

Delete a host record associated with an IP Address for a Specific Network Pool



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	networkPoolId := int64(1) // int64 | Morpheus Network Pool ID of the Object being referenced
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworkPoolIp(context.Background(), networkPoolId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkPoolIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkPoolIp`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworkPoolIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkPoolId** | **int64** | Morpheus Network Pool ID of the Object being referenced | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkPoolIpRequest struct via the builder pattern


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


## DeleteNetworkPoolServer

> DeleteAlerts200Response DeleteNetworkPoolServer(ctx, id).Execute()

Delete a Network Pool Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworkPoolServer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkPoolServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkPoolServer`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworkPoolServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkPoolServerRequest struct via the builder pattern


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


## DeleteNetworkProxy

> DeleteAlerts200Response DeleteNetworkProxy(ctx, id).Execute()

Delete a Network Proxy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworkProxy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkProxy`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworkProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkProxyRequest struct via the builder pattern


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


## DeleteNetworkRouter

> DeleteAlerts200Response DeleteNetworkRouter(ctx, id).Execute()

Delete a Network Router



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworkRouter(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkRouter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkRouter`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworkRouter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRouterRequest struct via the builder pattern


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


## DeleteNetworkRouterBgpNeighbor

> DeleteAlerts200Response DeleteNetworkRouterBgpNeighbor(ctx, id, routerId).Execute()

Delete a Network Router BGP Neighbor



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	routerId := float32(4) // float32 | Router ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworkRouterBgpNeighbor(context.Background(), id, routerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkRouterBgpNeighbor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkRouterBgpNeighbor`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworkRouterBgpNeighbor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRouterBgpNeighborRequest struct via the builder pattern


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


## DeleteNetworkRouterFirewallRule

> DeleteAlerts200Response DeleteNetworkRouterFirewallRule(ctx, id, routerId).Execute()

Delete a Network Router Firewall Rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	routerId := float32(4) // float32 | Router ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworkRouterFirewallRule(context.Background(), id, routerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkRouterFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkRouterFirewallRule`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworkRouterFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRouterFirewallRuleRequest struct via the builder pattern


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


## DeleteNetworkRouterFirewallRuleGroup

> DeleteAlerts200Response DeleteNetworkRouterFirewallRuleGroup(ctx, id, routerId).Execute()

Delete a Network Router firewall rule group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	routerId := float32(4) // float32 | Router ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworkRouterFirewallRuleGroup(context.Background(), id, routerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkRouterFirewallRuleGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkRouterFirewallRuleGroup`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworkRouterFirewallRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRouterFirewallRuleGroupRequest struct via the builder pattern


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


## DeleteNetworkRouterNat

> DeleteAlerts200Response DeleteNetworkRouterNat(ctx, id, routerId).Execute()

Delete a Network Router NAT



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	routerId := float32(4) // float32 | Router ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworkRouterNat(context.Background(), id, routerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkRouterNat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkRouterNat`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworkRouterNat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRouterNatRequest struct via the builder pattern


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


## DeleteNetworkRouterRoute

> DeleteAlerts200Response DeleteNetworkRouterRoute(ctx, id, routerId).Execute()

Delete a Network Router Route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	routerId := float32(4) // float32 | Router ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworkRouterRoute(context.Background(), id, routerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkRouterRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkRouterRoute`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworkRouterRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRouterRouteRequest struct via the builder pattern


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


## DeleteNetworkServer

> DeleteAlerts200Response DeleteNetworkServer(ctx, id).Execute()

Delete a Network Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworkServer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkServer`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworkServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkServerRequest struct via the builder pattern


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


## DeleteNetworkServerGroup

> DeleteAlerts200Response DeleteNetworkServerGroup(ctx, id, serverId).Execute()

Delete a Network Server Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	serverId := float32(4) // float32 | Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworkServerGroup(context.Background(), id, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkServerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkServerGroup`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworkServerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkServerGroupRequest struct via the builder pattern


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


## DeleteNetworkTransportZone

> DeleteAlerts200Response DeleteNetworkTransportZone(ctx, id, serverId).Execute()

Delete a Network Transport Zone



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	serverId := float32(4) // float32 | Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworkTransportZone(context.Background(), id, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkTransportZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkTransportZone`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworkTransportZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkTransportZoneRequest struct via the builder pattern


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


## DeleteStaticRoute

> DeleteAlerts200Response DeleteStaticRoute(ctx, id, routeId).Execute()

Delete a Network Static Route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	routeId := float32(4) // float32 | The ID of the route

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteStaticRoute(context.Background(), id, routeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStaticRoute`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routeId** | **float32** | The ID of the route | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStaticRouteRequest struct via the builder pattern


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


## DeleteSubnet

> DeleteAlerts200Response DeleteSubnet(ctx, id).Execute()

Delete a Subnet



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteSubnet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSubnet`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubnetRequest struct via the builder pattern


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


## GetAllNetworkFloatingIps

> GetAllNetworkFloatingIps200Response GetAllNetworkFloatingIps(ctx).Phrase(phrase).IpAddress(ipAddress).IpStatus(ipStatus).ZoneId(zoneId).ServerId(serverId).Execute()

Get All Floating IPs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	ipAddress := "10.32.23.188" // string | Filter by IP Address (optional)
	ipStatus := "ipStatus_example" // string | Filter by IP Status (optional)
	zoneId := int64(789) // int64 | Filter by Cloud ID (optional)
	serverId := int64(789) // int64 | Filter by Server ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetAllNetworkFloatingIps(context.Background()).Phrase(phrase).IpAddress(ipAddress).IpStatus(ipStatus).ZoneId(zoneId).ServerId(serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetAllNetworkFloatingIps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllNetworkFloatingIps`: GetAllNetworkFloatingIps200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetAllNetworkFloatingIps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllNetworkFloatingIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **ipAddress** | **string** | Filter by IP Address | 
 **ipStatus** | **string** | Filter by IP Status | 
 **zoneId** | **int64** | Filter by Cloud ID | 
 **serverId** | **int64** | Filter by Server ID | 

### Return type

[**GetAllNetworkFloatingIps200Response**](GetAllNetworkFloatingIps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetwork

> GetNetwork200Response GetNetwork(ctx, id).Execute()

Get a Specific Network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetwork(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetwork`: GetNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetwork200Response**](GetNetwork200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDhcpRelay

> GetNetworkDhcpRelay200Response GetNetworkDhcpRelay(ctx, id, serverId).Execute()

Get a Specific Network DHCP Relay



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	serverId := float32(4) // float32 | Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkDhcpRelay(context.Background(), id, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkDhcpRelay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDhcpRelay`: GetNetworkDhcpRelay200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkDhcpRelay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDhcpRelayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkDhcpRelay200Response**](GetNetworkDhcpRelay200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDhcpRelays

> GetNetworkDhcpRelays200Response GetNetworkDhcpRelays(ctx, serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()

Get all Network DHCP Relays for Network Relay



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	serverId := float32(4) // float32 | Server ID
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
	sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
	direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkDhcpRelays(context.Background(), serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkDhcpRelays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDhcpRelays`: GetNetworkDhcpRelays200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkDhcpRelays`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDhcpRelaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

[**GetNetworkDhcpRelays200Response**](GetNetworkDhcpRelays200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDhcpServer

> GetNetworkDhcpServer200Response GetNetworkDhcpServer(ctx, id, serverId).Execute()

Get a Specific Network DHCP Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	serverId := float32(4) // float32 | Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkDhcpServer(context.Background(), id, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkDhcpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDhcpServer`: GetNetworkDhcpServer200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkDhcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDhcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkDhcpServer200Response**](GetNetworkDhcpServer200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDhcpServers

> GetNetworkDhcpServers200Response GetNetworkDhcpServers(ctx, serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()

Get all Network DHCP Servers for Network Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	serverId := float32(4) // float32 | Server ID
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
	sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
	direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkDhcpServers(context.Background(), serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkDhcpServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDhcpServers`: GetNetworkDhcpServers200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkDhcpServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDhcpServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

[**GetNetworkDhcpServers200Response**](GetNetworkDhcpServers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDomain

> CreateNetworkDomain200Response GetNetworkDomain(ctx, id).Execute()

Get a Specific Network Domain



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkDomain(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDomain`: CreateNetworkDomain200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateNetworkDomain200Response**](CreateNetworkDomain200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDomains

> GetNetworkDomains200Response GetNetworkDomains(ctx).Name(name).Phrase(phrase).Execute()

Get all Network Domains



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	name := "example" // string | Filter by name (optional)
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkDomains(context.Background()).Name(name).Phrase(phrase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDomains`: GetNetworkDomains200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

[**GetNetworkDomains200Response**](GetNetworkDomains200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkEdgeCluster

> GetNetworkEdgeCluster200Response GetNetworkEdgeCluster(ctx, id, serverId).Execute()

Get a Specific Network Edge Cluster



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	serverId := float32(4) // float32 | Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkEdgeCluster(context.Background(), id, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkEdgeCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEdgeCluster`: GetNetworkEdgeCluster200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkEdgeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEdgeClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkEdgeCluster200Response**](GetNetworkEdgeCluster200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkEdgeClusters

> GetNetworkEdgeClusters200Response GetNetworkEdgeClusters(ctx, serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()

Get all Network Edge Clusters for Network Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	serverId := float32(4) // float32 | Server ID
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
	sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
	direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkEdgeClusters(context.Background(), serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkEdgeClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEdgeClusters`: GetNetworkEdgeClusters200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkEdgeClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEdgeClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

[**GetNetworkEdgeClusters200Response**](GetNetworkEdgeClusters200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirewallRule

> GetNetworkFirewallRule200Response GetNetworkFirewallRule(ctx, id, serverId).Execute()

Get a Specific Network Firewall Rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	serverId := float32(4) // float32 | Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkFirewallRule(context.Background(), id, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFirewallRule`: GetNetworkFirewallRule200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkFirewallRule200Response**](GetNetworkFirewallRule200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirewallRuleGroup

> GetNetworkFirewallRuleGroup200Response GetNetworkFirewallRuleGroup(ctx, id, serverId).Execute()

Get a Specific Network Firewall Rule Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	serverId := float32(4) // float32 | Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkFirewallRuleGroup(context.Background(), id, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkFirewallRuleGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFirewallRuleGroup`: GetNetworkFirewallRuleGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkFirewallRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirewallRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkFirewallRuleGroup200Response**](GetNetworkFirewallRuleGroup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirewallRuleGroups

> GetNetworkFirewallRuleGroups200Response GetNetworkFirewallRuleGroups(ctx, serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()

Get all Network Firewall Rule Groups for Network Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	serverId := float32(4) // float32 | Server ID
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
	sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
	direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkFirewallRuleGroups(context.Background(), serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkFirewallRuleGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFirewallRuleGroups`: GetNetworkFirewallRuleGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkFirewallRuleGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirewallRuleGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

[**GetNetworkFirewallRuleGroups200Response**](GetNetworkFirewallRuleGroups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirewallRules

> GetNetworkFirewallRules200Response GetNetworkFirewallRules(ctx, serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()

Get all Network Firewall Rules for Network Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	serverId := float32(4) // float32 | Server ID
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
	sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
	direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkFirewallRules(context.Background(), serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFirewallRules`: GetNetworkFirewallRules200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

[**GetNetworkFirewallRules200Response**](GetNetworkFirewallRules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFloatingIp

> GetNetworkFloatingIp200Response GetNetworkFloatingIp(ctx, id).Execute()

Get a Specific Floating IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkFloatingIp(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkFloatingIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFloatingIp`: GetNetworkFloatingIp200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkFloatingIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFloatingIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkFloatingIp200Response**](GetNetworkFloatingIp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkGroup

> GetNetworkGroup200Response GetNetworkGroup(ctx, id).Execute()

Get a Specific Network Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkGroup`: GetNetworkGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkGroup200Response**](GetNetworkGroup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkGroups

> GetNetworkGroups200Response GetNetworkGroups(ctx).Execute()

Get all Network Groups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkGroups`: GetNetworkGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkGroupsRequest struct via the builder pattern


### Return type

[**GetNetworkGroups200Response**](GetNetworkGroups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPool

> CreateNetworkPool200Response GetNetworkPool(ctx, id).Execute()

Get a Specific Network Pool



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkPool(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPool`: CreateNetworkPool200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateNetworkPool200Response**](CreateNetworkPool200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPoolIp

> GetNetworkPoolIps200Response GetNetworkPoolIp(ctx, networkPoolId, id).Execute()

Get a Specific IP Address for a Specific Network Pool



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	networkPoolId := int64(1) // int64 | Morpheus Network Pool ID of the Object being referenced
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkPoolIp(context.Background(), networkPoolId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkPoolIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPoolIp`: GetNetworkPoolIps200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkPoolIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkPoolId** | **int64** | Morpheus Network Pool ID of the Object being referenced | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPoolIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkPoolIps200Response**](GetNetworkPoolIps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPoolIps

> GetNetworkPoolIps200Response GetNetworkPoolIps(ctx, id).Phrase(phrase).IpAddress(ipAddress).Hostname(hostname).Execute()

Get all IP Addresses for a Specific Network Pool



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	phrase := "phrase_example" // string | Search phrase for partial matches on ipAddress or hostname (optional)
	ipAddress := "ipAddress_example" // string | If specified will return an exact match on ipAddress (optional)
	hostname := "hostname_example" // string | If specified will return an exact match on hostname (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkPoolIps(context.Background(), id).Phrase(phrase).IpAddress(ipAddress).Hostname(hostname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkPoolIps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPoolIps`: GetNetworkPoolIps200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkPoolIps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPoolIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phrase** | **string** | Search phrase for partial matches on ipAddress or hostname | 
 **ipAddress** | **string** | If specified will return an exact match on ipAddress | 
 **hostname** | **string** | If specified will return an exact match on hostname | 

### Return type

[**GetNetworkPoolIps200Response**](GetNetworkPoolIps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPoolServer

> GetNetworkPoolServer200Response GetNetworkPoolServer(ctx, id).Execute()

Get a Specific Network Pool Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkPoolServer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkPoolServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPoolServer`: GetNetworkPoolServer200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkPoolServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPoolServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkPoolServer200Response**](GetNetworkPoolServer200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPoolServerType

> GetNetworkPoolServerType200Response GetNetworkPoolServerType(ctx, id).Execute()

Retrieves a Specific Network Pool Server Type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkPoolServerType(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkPoolServerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPoolServerType`: GetNetworkPoolServerType200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkPoolServerType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPoolServerTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkPoolServerType200Response**](GetNetworkPoolServerType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPools

> GetNetworkPools200Response GetNetworkPools(ctx).Name(name).Phrase(phrase).Execute()

Get all Network Pools



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	name := "example" // string | Filter by name (optional)
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkPools(context.Background()).Name(name).Phrase(phrase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPools`: GetNetworkPools200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkPools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

[**GetNetworkPools200Response**](GetNetworkPools200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkProxies

> GetNetworkProxies200Response GetNetworkProxies(ctx).Name(name).Phrase(phrase).Execute()

Get all Network Proxies



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	name := "example" // string | Filter by name (optional)
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkProxies(context.Background()).Name(name).Phrase(phrase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkProxies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkProxies`: GetNetworkProxies200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkProxies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkProxiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

[**GetNetworkProxies200Response**](GetNetworkProxies200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkProxy

> CreateNetworkProxy200Response GetNetworkProxy(ctx, id).Execute()

Get a Specific Network Proxy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkProxy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkProxy`: CreateNetworkProxy200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateNetworkProxy200Response**](CreateNetworkProxy200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRouter

> GetNetworkRouter200Response GetNetworkRouter(ctx, id).Execute()

Get a Specific Network Router



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkRouter(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkRouter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkRouter`: GetNetworkRouter200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkRouter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkRouter200Response**](GetNetworkRouter200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRouterBgpNeighbor

> GetNetworkRouterBgpNeighbor200Response GetNetworkRouterBgpNeighbor(ctx, id, routerId).Execute()

Get a Network Router BGP Neighbor



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	routerId := float32(4) // float32 | Router ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkRouterBgpNeighbor(context.Background(), id, routerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkRouterBgpNeighbor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkRouterBgpNeighbor`: GetNetworkRouterBgpNeighbor200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkRouterBgpNeighbor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRouterBgpNeighborRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkRouterBgpNeighbor200Response**](GetNetworkRouterBgpNeighbor200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRouterFirewallRule

> GetNetworkRouterFirewallRule200Response GetNetworkRouterFirewallRule(ctx, id, routerId).Execute()

Get a Firewall Rule for Network Router



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	routerId := float32(4) // float32 | Router ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkRouterFirewallRule(context.Background(), id, routerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkRouterFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkRouterFirewallRule`: GetNetworkRouterFirewallRule200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkRouterFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRouterFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkRouterFirewallRule200Response**](GetNetworkRouterFirewallRule200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRouterFirewallRuleGroup

> GetNetworkRouterFirewallRuleGroup200Response GetNetworkRouterFirewallRuleGroup(ctx, id, routerId).Execute()

Get a Firewall Rule Group for Network Router



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	routerId := float32(4) // float32 | Router ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkRouterFirewallRuleGroup(context.Background(), id, routerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkRouterFirewallRuleGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkRouterFirewallRuleGroup`: GetNetworkRouterFirewallRuleGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkRouterFirewallRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRouterFirewallRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkRouterFirewallRuleGroup200Response**](GetNetworkRouterFirewallRuleGroup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRouterFirewallRuleGroups

> GetNetworkRouterFirewallRuleGroups200Response GetNetworkRouterFirewallRuleGroups(ctx, routerId).Execute()

Get all Network Firewall Rule Groups for Network Router



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	routerId := float32(4) // float32 | Router ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkRouterFirewallRuleGroups(context.Background(), routerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkRouterFirewallRuleGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkRouterFirewallRuleGroups`: GetNetworkRouterFirewallRuleGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkRouterFirewallRuleGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRouterFirewallRuleGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkRouterFirewallRuleGroups200Response**](GetNetworkRouterFirewallRuleGroups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRouterNat

> GetNetworkRouterNat200Response GetNetworkRouterNat(ctx, id, routerId).Execute()

Get a Network Router NAT



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	routerId := float32(4) // float32 | Router ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkRouterNat(context.Background(), id, routerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkRouterNat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkRouterNat`: GetNetworkRouterNat200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkRouterNat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRouterNatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkRouterNat200Response**](GetNetworkRouterNat200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRouterRoute

> GetNetworkRouterRoute200Response GetNetworkRouterRoute(ctx, id, routerId).Execute()

Get a Route for Network Router



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	routerId := float32(4) // float32 | Router ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkRouterRoute(context.Background(), id, routerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkRouterRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkRouterRoute`: GetNetworkRouterRoute200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkRouterRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRouterRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkRouterRoute200Response**](GetNetworkRouterRoute200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRouterType

> GetNetworkRouterType200Response GetNetworkRouterType(ctx, id).Execute()

Get a Specific Network Router Type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkRouterType(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkRouterType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkRouterType`: GetNetworkRouterType200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkRouterType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRouterTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkRouterType200Response**](GetNetworkRouterType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRouters

> GetNetworkRouters200Response GetNetworkRouters(ctx).Execute()

Get all Network Routers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkRouters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkRouters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkRouters`: GetNetworkRouters200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkRouters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRoutersRequest struct via the builder pattern


### Return type

[**GetNetworkRouters200Response**](GetNetworkRouters200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRoutersBgpNeighbors

> GetNetworkRoutersBgpNeighbors200Response GetNetworkRoutersBgpNeighbors(ctx, routerId).Execute()

Get all BGP Neighbors for Network Router



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	routerId := float32(4) // float32 | Router ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkRoutersBgpNeighbors(context.Background(), routerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkRoutersBgpNeighbors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkRoutersBgpNeighbors`: GetNetworkRoutersBgpNeighbors200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkRoutersBgpNeighbors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRoutersBgpNeighborsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkRoutersBgpNeighbors200Response**](GetNetworkRoutersBgpNeighbors200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRoutersFirewallRules

> GetNetworkRoutersFirewallRules200Response GetNetworkRoutersFirewallRules(ctx, routerId).Execute()

Get all Firewall Rules for Network Router



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	routerId := float32(4) // float32 | Router ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkRoutersFirewallRules(context.Background(), routerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkRoutersFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkRoutersFirewallRules`: GetNetworkRoutersFirewallRules200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkRoutersFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRoutersFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkRoutersFirewallRules200Response**](GetNetworkRoutersFirewallRules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRoutersNats

> GetNetworkRoutersNats200Response GetNetworkRoutersNats(ctx, routerId).Execute()

Get all Network Router NATs for Network Router



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	routerId := float32(4) // float32 | Router ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkRoutersNats(context.Background(), routerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkRoutersNats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkRoutersNats`: GetNetworkRoutersNats200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkRoutersNats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRoutersNatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkRoutersNats200Response**](GetNetworkRoutersNats200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRoutersRoutes

> GetNetworkRoutersRoutes200Response GetNetworkRoutersRoutes(ctx, routerId).Execute()

Get all Routes for Network Router



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	routerId := float32(4) // float32 | Router ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkRoutersRoutes(context.Background(), routerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkRoutersRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkRoutersRoutes`: GetNetworkRoutersRoutes200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkRoutersRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRoutersRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkRoutersRoutes200Response**](GetNetworkRoutersRoutes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkServer

> GetNetworkServer200Response GetNetworkServer(ctx, id).Execute()

Get a Specific Network Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkServer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkServer`: GetNetworkServer200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkServer200Response**](GetNetworkServer200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkServerGroup

> GetNetworkServerGroup200Response GetNetworkServerGroup(ctx, serverId, id).Execute()

Get a specific Network Server Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	serverId := float32(4) // float32 | Server ID
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkServerGroup(context.Background(), serverId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkServerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkServerGroup`: GetNetworkServerGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkServerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkServerGroup200Response**](GetNetworkServerGroup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSubnets

> GetNetworkSubnets200Response GetNetworkSubnets(ctx, id).Execute()

Get Subnets for a Network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkSubnets(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkSubnets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSubnets`: GetNetworkSubnets200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkSubnets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSubnetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSubnets200Response**](GetNetworkSubnets200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkTransportZone

> GetNetworkTransportZone200Response GetNetworkTransportZone(ctx, id, serverId).Execute()

Get a Specific Network Transport Zone



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	serverId := float32(4) // float32 | Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkTransportZone(context.Background(), id, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkTransportZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkTransportZone`: GetNetworkTransportZone200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkTransportZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTransportZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkTransportZone200Response**](GetNetworkTransportZone200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkTransportZones

> GetNetworkTransportZones200Response GetNetworkTransportZones(ctx, serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()

Get all Network Transport Zones for Network Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	serverId := float32(4) // float32 | Server ID
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
	sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
	direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkTransportZones(context.Background(), serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkTransportZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkTransportZones`: GetNetworkTransportZones200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkTransportZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTransportZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

[**GetNetworkTransportZones200Response**](GetNetworkTransportZones200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkType

> GetNetworkType200Response GetNetworkType(ctx, id).Execute()

Get a Specific Network Type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkType(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkType`: GetNetworkType200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkType200Response**](GetNetworkType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStaticRoute

> GetStaticRoute200Response GetStaticRoute(ctx, id, routeId).Execute()

Get a Static Route for a Network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	routeId := float32(4) // float32 | The ID of the route

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetStaticRoute(context.Background(), id, routeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStaticRoute`: GetStaticRoute200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routeId** | **float32** | The ID of the route | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetStaticRoute200Response**](GetStaticRoute200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStaticRoutes

> GetStaticRoutes200Response GetStaticRoutes(ctx, id).Execute()

Get all Static Routes for a Network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetStaticRoutes(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetStaticRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStaticRoutes`: GetStaticRoutes200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetStaticRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStaticRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetStaticRoutes200Response**](GetStaticRoutes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubnet

> CreateSubnet200Response GetSubnet(ctx, id).Execute()

Get a Specific Subnet



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetSubnet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubnet`: CreateSubnet200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateSubnet200Response**](CreateSubnet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkPoolServerTypes

> ListNetworkPoolServerTypes200Response ListNetworkPoolServerTypes(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).Execute()

Get All Network Pool Server Types



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
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
	resp, r, err := apiClient.NetworksAPI.ListNetworkPoolServerTypes(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.ListNetworkPoolServerTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkPoolServerTypes`: ListNetworkPoolServerTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.ListNetworkPoolServerTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkPoolServerTypesRequest struct via the builder pattern


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

[**ListNetworkPoolServerTypes200Response**](ListNetworkPoolServerTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkPoolServers

> ListNetworkPoolServers200Response ListNetworkPoolServers(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()

Get All Network Pool Servers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
	sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
	direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	name := "example" // string | Filter by name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.ListNetworkPoolServers(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.ListNetworkPoolServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkPoolServers`: ListNetworkPoolServers200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.ListNetworkPoolServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkPoolServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name | 

### Return type

[**ListNetworkPoolServers200Response**](ListNetworkPoolServers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkRouterTypes

> ListNetworkRouterTypes200Response ListNetworkRouterTypes(ctx).Execute()

Get all Network Router Types



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.ListNetworkRouterTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.ListNetworkRouterTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkRouterTypes`: ListNetworkRouterTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.ListNetworkRouterTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkRouterTypesRequest struct via the builder pattern


### Return type

[**ListNetworkRouterTypes200Response**](ListNetworkRouterTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkServerGroups

> ListNetworkServerGroups200Response ListNetworkServerGroups(ctx, serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()

Get all Network Server Groups for Network Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	serverId := float32(4) // float32 | Server ID
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
	sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
	direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.ListNetworkServerGroups(context.Background(), serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.ListNetworkServerGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkServerGroups`: ListNetworkServerGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.ListNetworkServerGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkServerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

[**ListNetworkServerGroups200Response**](ListNetworkServerGroups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkServers

> ListNetworkServers200Response ListNetworkServers(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()

Get All Network Servers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	max := int64(789) // int64 | Maximum number of records to return (optional) (default to 25)
	offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
	sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
	direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	name := "example" // string | Filter by name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.ListNetworkServers(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.ListNetworkServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkServers`: ListNetworkServers200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.ListNetworkServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name | 

### Return type

[**ListNetworkServers200Response**](ListNetworkServers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkServices

> ListNetworkServices200Response ListNetworkServices(ctx).Name(name).Phrase(phrase).Execute()

Get All Network Services



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	name := "example" // string | Filter by name (optional)
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.ListNetworkServices(context.Background()).Name(name).Phrase(phrase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.ListNetworkServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkServices`: ListNetworkServices200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.ListNetworkServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

[**ListNetworkServices200Response**](ListNetworkServices200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkTypes

> ListNetworkTypes200Response ListNetworkTypes(ctx).Name(name).Code(code).Phrase(phrase).Execute()

Network Types



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	name := "example" // string | Filter by name (optional)
	code := "azr" // string | If specified will return an exact match on code (optional)
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.ListNetworkTypes(context.Background()).Name(name).Code(code).Phrase(phrase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.ListNetworkTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkTypes`: ListNetworkTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.ListNetworkTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name | 
 **code** | **string** | If specified will return an exact match on code | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

[**ListNetworkTypes200Response**](ListNetworkTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworks

> ListNetworks200Response ListNetworks(ctx).Name(name).Phrase(phrase).Labels(labels).AllLabels(allLabels).ZoneId(zoneId).Cidr(cidr).Execute()

Get All Networks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	name := "example" // string | Filter by name (optional)
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
	allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)
	zoneId := int64(3) // int64 | The Zone ID for Filtering (optional)
	cidr := "cidr_example" // string | Filter by cidr, matches beginning of value. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.ListNetworks(context.Background()).Name(name).Phrase(phrase).Labels(labels).AllLabels(allLabels).ZoneId(zoneId).Cidr(cidr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.ListNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworks`: ListNetworks200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.ListNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 
 **zoneId** | **int64** | The Zone ID for Filtering | 
 **cidr** | **string** | Filter by cidr, matches beginning of value. | 

### Return type

[**ListNetworks200Response**](ListNetworks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubnetTypes

> ListSubnetTypes200Response ListSubnetTypes(ctx).Name(name).Phrase(phrase).Execute()

Get All Subnet Types



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	name := "example" // string | Filter by name (optional)
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.ListSubnetTypes(context.Background()).Name(name).Phrase(phrase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.ListSubnetTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubnetTypes`: ListSubnetTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.ListSubnetTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubnetTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

[**ListSubnetTypes200Response**](ListSubnetTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubnets

> GetNetworkSubnets200Response ListSubnets(ctx).Name(name).Phrase(phrase).Labels(labels).AllLabels(allLabels).Execute()

Get All Subnets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	name := "example" // string | Filter by name (optional)
	phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
	labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
	allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.ListSubnets(context.Background()).Name(name).Phrase(phrase).Labels(labels).AllLabels(allLabels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.ListSubnets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubnets`: GetNetworkSubnets200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.ListSubnets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubnetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 

### Return type

[**GetNetworkSubnets200Response**](GetNetworkSubnets200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshNetworkServer

> AddClusterLayouts200Response RefreshNetworkServer(ctx, id).Execute()

Refresh a Network Server/Integration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.RefreshNetworkServer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.RefreshNetworkServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshNetworkServer`: AddClusterLayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.RefreshNetworkServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshNetworkServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddClusterLayouts200Response**](AddClusterLayouts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseNetworkFloatingIp

> DeleteAlerts200Response ReleaseNetworkFloatingIp(ctx, id).Execute()

Release a Floating IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.ReleaseNetworkFloatingIp(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.ReleaseNetworkFloatingIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReleaseNetworkFloatingIp`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.ReleaseNetworkFloatingIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseNetworkFloatingIpRequest struct via the builder pattern


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


## UpdateNetwork

> CreateNetworks200Response UpdateNetwork(ctx, id).UpdateNetworkRequest(updateNetworkRequest).Execute()

Update a Network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateNetworkRequest := *openapiclient.NewUpdateNetworkRequest() // UpdateNetworkRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetwork(context.Background(), id).UpdateNetworkRequest(updateNetworkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetwork`: CreateNetworks200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkRequest** | [**UpdateNetworkRequest**](UpdateNetworkRequest.md) |  | 

### Return type

[**CreateNetworks200Response**](CreateNetworks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkDhcpRelay

> DeleteAlerts200Response UpdateNetworkDhcpRelay(ctx, id, serverId).UpdateNetworkDhcpRelayRequest(updateNetworkDhcpRelayRequest).Execute()

Update a Network DHCP Relay



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	serverId := float32(4) // float32 | Server ID
	updateNetworkDhcpRelayRequest := *openapiclient.NewUpdateNetworkDhcpRelayRequest() // UpdateNetworkDhcpRelayRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkDhcpRelay(context.Background(), id, serverId).UpdateNetworkDhcpRelayRequest(updateNetworkDhcpRelayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkDhcpRelay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkDhcpRelay`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkDhcpRelay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkDhcpRelayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkDhcpRelayRequest** | [**UpdateNetworkDhcpRelayRequest**](UpdateNetworkDhcpRelayRequest.md) |  | 

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


## UpdateNetworkDhcpServer

> DeleteAlerts200Response UpdateNetworkDhcpServer(ctx, id, serverId).UpdateNetworkDhcpServerRequest(updateNetworkDhcpServerRequest).Execute()

Update a Network DHCP Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	serverId := float32(4) // float32 | Server ID
	updateNetworkDhcpServerRequest := *openapiclient.NewUpdateNetworkDhcpServerRequest() // UpdateNetworkDhcpServerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkDhcpServer(context.Background(), id, serverId).UpdateNetworkDhcpServerRequest(updateNetworkDhcpServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkDhcpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkDhcpServer`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkDhcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkDhcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkDhcpServerRequest** | [**UpdateNetworkDhcpServerRequest**](UpdateNetworkDhcpServerRequest.md) |  | 

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


## UpdateNetworkDomain

> CreateNetworkDomain200Response UpdateNetworkDomain(ctx, id).CreateNetworkDomainRequest(createNetworkDomainRequest).Execute()

Update a Network Domain



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	createNetworkDomainRequest := *openapiclient.NewCreateNetworkDomainRequest() // CreateNetworkDomainRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkDomain(context.Background(), id).CreateNetworkDomainRequest(createNetworkDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkDomain`: CreateNetworkDomain200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkDomainRequest** | [**CreateNetworkDomainRequest**](CreateNetworkDomainRequest.md) |  | 

### Return type

[**CreateNetworkDomain200Response**](CreateNetworkDomain200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkEdgeCluster

> DeleteAlerts200Response UpdateNetworkEdgeCluster(ctx, id, serverId).UpdateNetworkEdgeClusterRequest(updateNetworkEdgeClusterRequest).Execute()

Update a Network Edge Cluster



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	serverId := float32(4) // float32 | Server ID
	updateNetworkEdgeClusterRequest := *openapiclient.NewUpdateNetworkEdgeClusterRequest() // UpdateNetworkEdgeClusterRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkEdgeCluster(context.Background(), id, serverId).UpdateNetworkEdgeClusterRequest(updateNetworkEdgeClusterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkEdgeCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkEdgeCluster`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkEdgeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkEdgeClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkEdgeClusterRequest** | [**UpdateNetworkEdgeClusterRequest**](UpdateNetworkEdgeClusterRequest.md) |  | 

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


## UpdateNetworkFirewallRule

> DeleteAlerts200Response UpdateNetworkFirewallRule(ctx, id, serverId).UpdateNetworkFirewallRuleRequest(updateNetworkFirewallRuleRequest).Execute()

Update a Network Firewall Rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	serverId := float32(4) // float32 | Server ID
	updateNetworkFirewallRuleRequest := *openapiclient.NewUpdateNetworkFirewallRuleRequest() // UpdateNetworkFirewallRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkFirewallRule(context.Background(), id, serverId).UpdateNetworkFirewallRuleRequest(updateNetworkFirewallRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkFirewallRule`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkFirewallRuleRequest** | [**UpdateNetworkFirewallRuleRequest**](UpdateNetworkFirewallRuleRequest.md) |  | 

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


## UpdateNetworkFirewallRuleGroup

> DeleteAlerts200Response UpdateNetworkFirewallRuleGroup(ctx, id, serverId).UpdateNetworkFirewallRuleGroupRequest(updateNetworkFirewallRuleGroupRequest).Execute()

Update a Network Firewall Rule Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	serverId := float32(4) // float32 | Server ID
	updateNetworkFirewallRuleGroupRequest := *openapiclient.NewUpdateNetworkFirewallRuleGroupRequest() // UpdateNetworkFirewallRuleGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkFirewallRuleGroup(context.Background(), id, serverId).UpdateNetworkFirewallRuleGroupRequest(updateNetworkFirewallRuleGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkFirewallRuleGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkFirewallRuleGroup`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkFirewallRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFirewallRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkFirewallRuleGroupRequest** | [**UpdateNetworkFirewallRuleGroupRequest**](UpdateNetworkFirewallRuleGroupRequest.md) |  | 

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


## UpdateNetworkGroup

> DeleteAlerts200Response UpdateNetworkGroup(ctx, id).CreateNetworkGroupRequest(createNetworkGroupRequest).Execute()

Update a Network Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	createNetworkGroupRequest := *openapiclient.NewCreateNetworkGroupRequest() // CreateNetworkGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkGroup(context.Background(), id).CreateNetworkGroupRequest(createNetworkGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkGroup`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkGroupRequest** | [**CreateNetworkGroupRequest**](CreateNetworkGroupRequest.md) |  | 

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


## UpdateNetworkPool

> CreateNetworkPool200Response UpdateNetworkPool(ctx, id).CreateNetworkPoolRequest(createNetworkPoolRequest).Execute()

Update a Network Pool



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	createNetworkPoolRequest := *openapiclient.NewCreateNetworkPoolRequest() // CreateNetworkPoolRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkPool(context.Background(), id).CreateNetworkPoolRequest(createNetworkPoolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkPool`: CreateNetworkPool200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkPoolRequest** | [**CreateNetworkPoolRequest**](CreateNetworkPoolRequest.md) |  | 

### Return type

[**CreateNetworkPool200Response**](CreateNetworkPool200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkPoolServer

> CreateNetworkPoolServer200Response UpdateNetworkPoolServer(ctx, id).UpdateNetworkPoolServerRequest(updateNetworkPoolServerRequest).Execute()

Update a Network Pool Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateNetworkPoolServerRequest := *openapiclient.NewUpdateNetworkPoolServerRequest() // UpdateNetworkPoolServerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkPoolServer(context.Background(), id).UpdateNetworkPoolServerRequest(updateNetworkPoolServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkPoolServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkPoolServer`: CreateNetworkPoolServer200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkPoolServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkPoolServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkPoolServerRequest** | [**UpdateNetworkPoolServerRequest**](UpdateNetworkPoolServerRequest.md) |  | 

### Return type

[**CreateNetworkPoolServer200Response**](CreateNetworkPoolServer200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkProxy

> DeleteAlerts200Response UpdateNetworkProxy(ctx, id).CreateNetworkProxyRequest(createNetworkProxyRequest).Execute()

Update a Network Proxy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	createNetworkProxyRequest := *openapiclient.NewCreateNetworkProxyRequest() // CreateNetworkProxyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkProxy(context.Background(), id).CreateNetworkProxyRequest(createNetworkProxyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkProxy`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkProxyRequest** | [**CreateNetworkProxyRequest**](CreateNetworkProxyRequest.md) |  | 

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


## UpdateNetworkRouter

> DeleteAlerts200Response UpdateNetworkRouter(ctx, id).UpdateNetworkRouterRequest(updateNetworkRouterRequest).Execute()

Update a Network Router



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateNetworkRouterRequest := *openapiclient.NewUpdateNetworkRouterRequest() // UpdateNetworkRouterRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkRouter(context.Background(), id).UpdateNetworkRouterRequest(updateNetworkRouterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkRouter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkRouter`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkRouter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkRouterRequest** | [**UpdateNetworkRouterRequest**](UpdateNetworkRouterRequest.md) |  | 

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


## UpdateNetworkRouterBgpNeighbor

> DeleteAlerts200Response UpdateNetworkRouterBgpNeighbor(ctx, id, routerId).UpdateNetworkRouterBgpNeighborRequest(updateNetworkRouterBgpNeighborRequest).Execute()

Update Network Router BGP Neighbor



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	routerId := float32(4) // float32 | Router ID
	updateNetworkRouterBgpNeighborRequest := *openapiclient.NewUpdateNetworkRouterBgpNeighborRequest() // UpdateNetworkRouterBgpNeighborRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkRouterBgpNeighbor(context.Background(), id, routerId).UpdateNetworkRouterBgpNeighborRequest(updateNetworkRouterBgpNeighborRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkRouterBgpNeighbor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkRouterBgpNeighbor`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkRouterBgpNeighbor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRouterBgpNeighborRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkRouterBgpNeighborRequest** | [**UpdateNetworkRouterBgpNeighborRequest**](UpdateNetworkRouterBgpNeighborRequest.md) |  | 

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


## UpdateNetworkRouterFirewallRule

> DeleteAlerts200Response UpdateNetworkRouterFirewallRule(ctx, id, routerId).UpdateNetworkRouterFirewallRuleRequest(updateNetworkRouterFirewallRuleRequest).Execute()

Update a Network Router Firewall Rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	routerId := float32(4) // float32 | Router ID
	updateNetworkRouterFirewallRuleRequest := *openapiclient.NewUpdateNetworkRouterFirewallRuleRequest() // UpdateNetworkRouterFirewallRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkRouterFirewallRule(context.Background(), id, routerId).UpdateNetworkRouterFirewallRuleRequest(updateNetworkRouterFirewallRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkRouterFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkRouterFirewallRule`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkRouterFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRouterFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkRouterFirewallRuleRequest** | [**UpdateNetworkRouterFirewallRuleRequest**](UpdateNetworkRouterFirewallRuleRequest.md) |  | 

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


## UpdateNetworkRouterFirewallRuleGroup

> DeleteAlerts200Response UpdateNetworkRouterFirewallRuleGroup(ctx, id, routerId).UpdateNetworkRouterFirewallRuleGroupRequest(updateNetworkRouterFirewallRuleGroupRequest).Execute()

Update a Network Router Firewall Rule Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	routerId := float32(4) // float32 | Router ID
	updateNetworkRouterFirewallRuleGroupRequest := *openapiclient.NewUpdateNetworkRouterFirewallRuleGroupRequest() // UpdateNetworkRouterFirewallRuleGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkRouterFirewallRuleGroup(context.Background(), id, routerId).UpdateNetworkRouterFirewallRuleGroupRequest(updateNetworkRouterFirewallRuleGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkRouterFirewallRuleGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkRouterFirewallRuleGroup`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkRouterFirewallRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRouterFirewallRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkRouterFirewallRuleGroupRequest** | [**UpdateNetworkRouterFirewallRuleGroupRequest**](UpdateNetworkRouterFirewallRuleGroupRequest.md) |  | 

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


## UpdateNetworkRouterNat

> DeleteAlerts200Response UpdateNetworkRouterNat(ctx, id, routerId).UpdateNetworkRouterNatRequest(updateNetworkRouterNatRequest).Execute()

Update Network Router NAT



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	routerId := float32(4) // float32 | Router ID
	updateNetworkRouterNatRequest := *openapiclient.NewUpdateNetworkRouterNatRequest() // UpdateNetworkRouterNatRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkRouterNat(context.Background(), id, routerId).UpdateNetworkRouterNatRequest(updateNetworkRouterNatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkRouterNat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkRouterNat`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkRouterNat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRouterNatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkRouterNatRequest** | [**UpdateNetworkRouterNatRequest**](UpdateNetworkRouterNatRequest.md) |  | 

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


## UpdateNetworkRouterPermissions

> AddClusterLayouts200Response UpdateNetworkRouterPermissions(ctx, routerId).UpdateNetworkRouterPermissionsRequest(updateNetworkRouterPermissionsRequest).Execute()

Update Network Router Permissions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	routerId := float32(4) // float32 | Router ID
	updateNetworkRouterPermissionsRequest := *openapiclient.NewUpdateNetworkRouterPermissionsRequest() // UpdateNetworkRouterPermissionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkRouterPermissions(context.Background(), routerId).UpdateNetworkRouterPermissionsRequest(updateNetworkRouterPermissionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkRouterPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkRouterPermissions`: AddClusterLayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkRouterPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRouterPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkRouterPermissionsRequest** | [**UpdateNetworkRouterPermissionsRequest**](UpdateNetworkRouterPermissionsRequest.md) |  | 

### Return type

[**AddClusterLayouts200Response**](AddClusterLayouts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkServer

> CreateNetworkServer200Response UpdateNetworkServer(ctx, id).UpdateNetworkServerRequest(updateNetworkServerRequest).Execute()

Update a Network Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	updateNetworkServerRequest := *openapiclient.NewUpdateNetworkServerRequest() // UpdateNetworkServerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkServer(context.Background(), id).UpdateNetworkServerRequest(updateNetworkServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkServer`: CreateNetworkServer200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkServerRequest** | [**UpdateNetworkServerRequest**](UpdateNetworkServerRequest.md) |  | 

### Return type

[**CreateNetworkServer200Response**](CreateNetworkServer200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkServerGroup

> DeleteAlerts200Response UpdateNetworkServerGroup(ctx, id, serverId).UpdateNetworkServerGroupRequest(updateNetworkServerGroupRequest).Execute()

Update a Network Server Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	serverId := float32(4) // float32 | Server ID
	updateNetworkServerGroupRequest := *openapiclient.NewUpdateNetworkServerGroupRequest() // UpdateNetworkServerGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkServerGroup(context.Background(), id, serverId).UpdateNetworkServerGroupRequest(updateNetworkServerGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkServerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkServerGroup`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkServerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkServerGroupRequest** | [**UpdateNetworkServerGroupRequest**](UpdateNetworkServerGroupRequest.md) |  | 

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


## UpdateNetworkTransportZone

> DeleteAlerts200Response UpdateNetworkTransportZone(ctx, id, serverId).UpdateNetworkTransportZoneRequest(updateNetworkTransportZoneRequest).Execute()

Update a Network Transport Zone



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	serverId := float32(4) // float32 | Server ID
	updateNetworkTransportZoneRequest := *openapiclient.NewUpdateNetworkTransportZoneRequest() // UpdateNetworkTransportZoneRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkTransportZone(context.Background(), id, serverId).UpdateNetworkTransportZoneRequest(updateNetworkTransportZoneRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkTransportZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkTransportZone`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkTransportZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkTransportZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkTransportZoneRequest** | [**UpdateNetworkTransportZoneRequest**](UpdateNetworkTransportZoneRequest.md) |  | 

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


## UpdateStaticRoute

> AddClusterLayouts200Response UpdateStaticRoute(ctx, id, routeId).CreateStaticRouteRequest(createStaticRouteRequest).Execute()

Update a Network Static Route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	routeId := float32(4) // float32 | The ID of the route
	createStaticRouteRequest := *openapiclient.NewCreateStaticRouteRequest() // CreateStaticRouteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateStaticRoute(context.Background(), id, routeId).CreateStaticRouteRequest(createStaticRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStaticRoute`: AddClusterLayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routeId** | **float32** | The ID of the route | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createStaticRouteRequest** | [**CreateStaticRouteRequest**](CreateStaticRouteRequest.md) |  | 

### Return type

[**AddClusterLayouts200Response**](AddClusterLayouts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubnet

> CreateSubnet200Response UpdateSubnet(ctx, id).CreateSubnetRequest(createSubnetRequest).Execute()

Update a Subnet



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/HewlettPackard/hpe-morpheus-client"
)

func main() {
	id := int64(1) // int64 | Morpheus ID of the Object being referenced
	createSubnetRequest := *openapiclient.NewCreateSubnetRequest() // CreateSubnetRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateSubnet(context.Background(), id).CreateSubnetRequest(createSubnetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubnet`: CreateSubnet200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSubnetRequest** | [**CreateSubnetRequest**](CreateSubnetRequest.md) |  | 

### Return type

[**CreateSubnet200Response**](CreateSubnet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

