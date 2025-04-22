# NetworkUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Display Name | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**Cidr** | Pointer to **string** | CIDR Network | [optional] 
**Gateway** | Pointer to **string** | Network Gateway | [optional] 
**DnsPrimary** | Pointer to **string** | Primary DNS Server | [optional] 
**DnsSecondary** | Pointer to **string** | Secondary DNS Server | [optional] 
**VlanId** | Pointer to **int64** |  | [optional] 
**Pool** | Pointer to **int64** | Network Pool ID | [optional] 
**AllowStaticOverride** | Pointer to **bool** | Allow IP Override | [optional] 
**AssignPublicIp** | Pointer to **bool** | Assign Public IP | [optional] 
**Active** | Pointer to **bool** | Activate (true) or disable (false) the network | [optional] 
**DhcpServer** | Pointer to **bool** | DHCP Server enabled network | [optional] 
**NetworkDomain** | Pointer to [**CreateNetworksRequestNetworkNetworkDomain**](CreateNetworksRequestNetworkNetworkDomain.md) |  | [optional] 
**SearchDomains** | Pointer to **string** | Search Domains | [optional] 
**NetworkProxy** | Pointer to [**CreateNetworksRequestNetworkNetworkProxy**](CreateNetworksRequestNetworkNetworkProxy.md) |  | [optional] 
**ApplianceUrlProxyBypass** | Pointer to **bool** | Bypass Proxy for Appliance URL | [optional] 
**NoProxy** | Pointer to **string** | Comma-separated list of ip addresses or name servers to exclude proxy traversal for. Typically locally routable servers are excluded. | [optional] 
**Visibility** | Pointer to **string** | Visibility, private or public. | [optional] [default to "private"]
**Config** | Pointer to **map[string]interface{}** | Configuration object. Settings vary by type. | [optional] 
**Tenants** | Pointer to [**[]GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) | Array of tenant account ids that are allowed access | [optional] 
**ResourcePermissions** | Pointer to [**UpdateNetworkRequestNetworkResourcePermissions**](UpdateNetworkRequestNetworkResourcePermissions.md) |  | [optional] 

## Methods

### NewNetworkUpdate

`func NewNetworkUpdate() *NetworkUpdate`

NewNetworkUpdate instantiates a new NetworkUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkUpdateWithDefaults

`func NewNetworkUpdateWithDefaults() *NetworkUpdate`

NewNetworkUpdateWithDefaults instantiates a new NetworkUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLabels

`func (o *NetworkUpdate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NetworkUpdate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NetworkUpdate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *NetworkUpdate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCidr

`func (o *NetworkUpdate) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *NetworkUpdate) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *NetworkUpdate) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *NetworkUpdate) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetGateway

`func (o *NetworkUpdate) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *NetworkUpdate) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *NetworkUpdate) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *NetworkUpdate) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetDnsPrimary

`func (o *NetworkUpdate) GetDnsPrimary() string`

GetDnsPrimary returns the DnsPrimary field if non-nil, zero value otherwise.

### GetDnsPrimaryOk

`func (o *NetworkUpdate) GetDnsPrimaryOk() (*string, bool)`

GetDnsPrimaryOk returns a tuple with the DnsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrimary

`func (o *NetworkUpdate) SetDnsPrimary(v string)`

SetDnsPrimary sets DnsPrimary field to given value.

### HasDnsPrimary

`func (o *NetworkUpdate) HasDnsPrimary() bool`

HasDnsPrimary returns a boolean if a field has been set.

### GetDnsSecondary

`func (o *NetworkUpdate) GetDnsSecondary() string`

GetDnsSecondary returns the DnsSecondary field if non-nil, zero value otherwise.

### GetDnsSecondaryOk

`func (o *NetworkUpdate) GetDnsSecondaryOk() (*string, bool)`

GetDnsSecondaryOk returns a tuple with the DnsSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSecondary

`func (o *NetworkUpdate) SetDnsSecondary(v string)`

SetDnsSecondary sets DnsSecondary field to given value.

### HasDnsSecondary

`func (o *NetworkUpdate) HasDnsSecondary() bool`

HasDnsSecondary returns a boolean if a field has been set.

### GetVlanId

`func (o *NetworkUpdate) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *NetworkUpdate) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *NetworkUpdate) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *NetworkUpdate) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetPool

`func (o *NetworkUpdate) GetPool() int64`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *NetworkUpdate) GetPoolOk() (*int64, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *NetworkUpdate) SetPool(v int64)`

SetPool sets Pool field to given value.

### HasPool

`func (o *NetworkUpdate) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetAllowStaticOverride

`func (o *NetworkUpdate) GetAllowStaticOverride() bool`

GetAllowStaticOverride returns the AllowStaticOverride field if non-nil, zero value otherwise.

### GetAllowStaticOverrideOk

`func (o *NetworkUpdate) GetAllowStaticOverrideOk() (*bool, bool)`

GetAllowStaticOverrideOk returns a tuple with the AllowStaticOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowStaticOverride

`func (o *NetworkUpdate) SetAllowStaticOverride(v bool)`

SetAllowStaticOverride sets AllowStaticOverride field to given value.

### HasAllowStaticOverride

`func (o *NetworkUpdate) HasAllowStaticOverride() bool`

HasAllowStaticOverride returns a boolean if a field has been set.

### GetAssignPublicIp

`func (o *NetworkUpdate) GetAssignPublicIp() bool`

GetAssignPublicIp returns the AssignPublicIp field if non-nil, zero value otherwise.

### GetAssignPublicIpOk

`func (o *NetworkUpdate) GetAssignPublicIpOk() (*bool, bool)`

GetAssignPublicIpOk returns a tuple with the AssignPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignPublicIp

`func (o *NetworkUpdate) SetAssignPublicIp(v bool)`

SetAssignPublicIp sets AssignPublicIp field to given value.

### HasAssignPublicIp

`func (o *NetworkUpdate) HasAssignPublicIp() bool`

HasAssignPublicIp returns a boolean if a field has been set.

### GetActive

`func (o *NetworkUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *NetworkUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *NetworkUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *NetworkUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDhcpServer

`func (o *NetworkUpdate) GetDhcpServer() bool`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *NetworkUpdate) GetDhcpServerOk() (*bool, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *NetworkUpdate) SetDhcpServer(v bool)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *NetworkUpdate) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *NetworkUpdate) GetNetworkDomain() CreateNetworksRequestNetworkNetworkDomain`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *NetworkUpdate) GetNetworkDomainOk() (*CreateNetworksRequestNetworkNetworkDomain, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *NetworkUpdate) SetNetworkDomain(v CreateNetworksRequestNetworkNetworkDomain)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *NetworkUpdate) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetSearchDomains

`func (o *NetworkUpdate) GetSearchDomains() string`

GetSearchDomains returns the SearchDomains field if non-nil, zero value otherwise.

### GetSearchDomainsOk

`func (o *NetworkUpdate) GetSearchDomainsOk() (*string, bool)`

GetSearchDomainsOk returns a tuple with the SearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomains

`func (o *NetworkUpdate) SetSearchDomains(v string)`

SetSearchDomains sets SearchDomains field to given value.

### HasSearchDomains

`func (o *NetworkUpdate) HasSearchDomains() bool`

HasSearchDomains returns a boolean if a field has been set.

### GetNetworkProxy

`func (o *NetworkUpdate) GetNetworkProxy() CreateNetworksRequestNetworkNetworkProxy`

GetNetworkProxy returns the NetworkProxy field if non-nil, zero value otherwise.

### GetNetworkProxyOk

`func (o *NetworkUpdate) GetNetworkProxyOk() (*CreateNetworksRequestNetworkNetworkProxy, bool)`

GetNetworkProxyOk returns a tuple with the NetworkProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProxy

`func (o *NetworkUpdate) SetNetworkProxy(v CreateNetworksRequestNetworkNetworkProxy)`

SetNetworkProxy sets NetworkProxy field to given value.

### HasNetworkProxy

`func (o *NetworkUpdate) HasNetworkProxy() bool`

HasNetworkProxy returns a boolean if a field has been set.

### GetApplianceUrlProxyBypass

`func (o *NetworkUpdate) GetApplianceUrlProxyBypass() bool`

GetApplianceUrlProxyBypass returns the ApplianceUrlProxyBypass field if non-nil, zero value otherwise.

### GetApplianceUrlProxyBypassOk

`func (o *NetworkUpdate) GetApplianceUrlProxyBypassOk() (*bool, bool)`

GetApplianceUrlProxyBypassOk returns a tuple with the ApplianceUrlProxyBypass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrlProxyBypass

`func (o *NetworkUpdate) SetApplianceUrlProxyBypass(v bool)`

SetApplianceUrlProxyBypass sets ApplianceUrlProxyBypass field to given value.

### HasApplianceUrlProxyBypass

`func (o *NetworkUpdate) HasApplianceUrlProxyBypass() bool`

HasApplianceUrlProxyBypass returns a boolean if a field has been set.

### GetNoProxy

`func (o *NetworkUpdate) GetNoProxy() string`

GetNoProxy returns the NoProxy field if non-nil, zero value otherwise.

### GetNoProxyOk

`func (o *NetworkUpdate) GetNoProxyOk() (*string, bool)`

GetNoProxyOk returns a tuple with the NoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxy

`func (o *NetworkUpdate) SetNoProxy(v string)`

SetNoProxy sets NoProxy field to given value.

### HasNoProxy

`func (o *NetworkUpdate) HasNoProxy() bool`

HasNoProxy returns a boolean if a field has been set.

### GetVisibility

`func (o *NetworkUpdate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *NetworkUpdate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *NetworkUpdate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *NetworkUpdate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkUpdate) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkUpdate) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkUpdate) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkUpdate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTenants

`func (o *NetworkUpdate) GetTenants() []GetAlerts200ResponseAllOfChecksInnerAccount`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *NetworkUpdate) GetTenantsOk() (*[]GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *NetworkUpdate) SetTenants(v []GetAlerts200ResponseAllOfChecksInnerAccount)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *NetworkUpdate) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *NetworkUpdate) GetResourcePermissions() UpdateNetworkRequestNetworkResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *NetworkUpdate) GetResourcePermissionsOk() (*UpdateNetworkRequestNetworkResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *NetworkUpdate) SetResourcePermissions(v UpdateNetworkRequestNetworkResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *NetworkUpdate) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


