# UpdateNetworkRequestNetwork

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

### NewUpdateNetworkRequestNetwork

`func NewUpdateNetworkRequestNetwork() *UpdateNetworkRequestNetwork`

NewUpdateNetworkRequestNetwork instantiates a new UpdateNetworkRequestNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkRequestNetworkWithDefaults

`func NewUpdateNetworkRequestNetworkWithDefaults() *UpdateNetworkRequestNetwork`

NewUpdateNetworkRequestNetworkWithDefaults instantiates a new UpdateNetworkRequestNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *UpdateNetworkRequestNetwork) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateNetworkRequestNetwork) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateNetworkRequestNetwork) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateNetworkRequestNetwork) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateNetworkRequestNetwork) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateNetworkRequestNetwork) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateNetworkRequestNetwork) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateNetworkRequestNetwork) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateNetworkRequestNetwork) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateNetworkRequestNetwork) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateNetworkRequestNetwork) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateNetworkRequestNetwork) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCidr

`func (o *UpdateNetworkRequestNetwork) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *UpdateNetworkRequestNetwork) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *UpdateNetworkRequestNetwork) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *UpdateNetworkRequestNetwork) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetGateway

`func (o *UpdateNetworkRequestNetwork) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UpdateNetworkRequestNetwork) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UpdateNetworkRequestNetwork) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UpdateNetworkRequestNetwork) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetDnsPrimary

`func (o *UpdateNetworkRequestNetwork) GetDnsPrimary() string`

GetDnsPrimary returns the DnsPrimary field if non-nil, zero value otherwise.

### GetDnsPrimaryOk

`func (o *UpdateNetworkRequestNetwork) GetDnsPrimaryOk() (*string, bool)`

GetDnsPrimaryOk returns a tuple with the DnsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrimary

`func (o *UpdateNetworkRequestNetwork) SetDnsPrimary(v string)`

SetDnsPrimary sets DnsPrimary field to given value.

### HasDnsPrimary

`func (o *UpdateNetworkRequestNetwork) HasDnsPrimary() bool`

HasDnsPrimary returns a boolean if a field has been set.

### GetDnsSecondary

`func (o *UpdateNetworkRequestNetwork) GetDnsSecondary() string`

GetDnsSecondary returns the DnsSecondary field if non-nil, zero value otherwise.

### GetDnsSecondaryOk

`func (o *UpdateNetworkRequestNetwork) GetDnsSecondaryOk() (*string, bool)`

GetDnsSecondaryOk returns a tuple with the DnsSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSecondary

`func (o *UpdateNetworkRequestNetwork) SetDnsSecondary(v string)`

SetDnsSecondary sets DnsSecondary field to given value.

### HasDnsSecondary

`func (o *UpdateNetworkRequestNetwork) HasDnsSecondary() bool`

HasDnsSecondary returns a boolean if a field has been set.

### GetVlanId

`func (o *UpdateNetworkRequestNetwork) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *UpdateNetworkRequestNetwork) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *UpdateNetworkRequestNetwork) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *UpdateNetworkRequestNetwork) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetPool

`func (o *UpdateNetworkRequestNetwork) GetPool() int64`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *UpdateNetworkRequestNetwork) GetPoolOk() (*int64, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *UpdateNetworkRequestNetwork) SetPool(v int64)`

SetPool sets Pool field to given value.

### HasPool

`func (o *UpdateNetworkRequestNetwork) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetAllowStaticOverride

`func (o *UpdateNetworkRequestNetwork) GetAllowStaticOverride() bool`

GetAllowStaticOverride returns the AllowStaticOverride field if non-nil, zero value otherwise.

### GetAllowStaticOverrideOk

`func (o *UpdateNetworkRequestNetwork) GetAllowStaticOverrideOk() (*bool, bool)`

GetAllowStaticOverrideOk returns a tuple with the AllowStaticOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowStaticOverride

`func (o *UpdateNetworkRequestNetwork) SetAllowStaticOverride(v bool)`

SetAllowStaticOverride sets AllowStaticOverride field to given value.

### HasAllowStaticOverride

`func (o *UpdateNetworkRequestNetwork) HasAllowStaticOverride() bool`

HasAllowStaticOverride returns a boolean if a field has been set.

### GetAssignPublicIp

`func (o *UpdateNetworkRequestNetwork) GetAssignPublicIp() bool`

GetAssignPublicIp returns the AssignPublicIp field if non-nil, zero value otherwise.

### GetAssignPublicIpOk

`func (o *UpdateNetworkRequestNetwork) GetAssignPublicIpOk() (*bool, bool)`

GetAssignPublicIpOk returns a tuple with the AssignPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignPublicIp

`func (o *UpdateNetworkRequestNetwork) SetAssignPublicIp(v bool)`

SetAssignPublicIp sets AssignPublicIp field to given value.

### HasAssignPublicIp

`func (o *UpdateNetworkRequestNetwork) HasAssignPublicIp() bool`

HasAssignPublicIp returns a boolean if a field has been set.

### GetActive

`func (o *UpdateNetworkRequestNetwork) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateNetworkRequestNetwork) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateNetworkRequestNetwork) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateNetworkRequestNetwork) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDhcpServer

`func (o *UpdateNetworkRequestNetwork) GetDhcpServer() bool`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *UpdateNetworkRequestNetwork) GetDhcpServerOk() (*bool, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *UpdateNetworkRequestNetwork) SetDhcpServer(v bool)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *UpdateNetworkRequestNetwork) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *UpdateNetworkRequestNetwork) GetNetworkDomain() CreateNetworksRequestNetworkNetworkDomain`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *UpdateNetworkRequestNetwork) GetNetworkDomainOk() (*CreateNetworksRequestNetworkNetworkDomain, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *UpdateNetworkRequestNetwork) SetNetworkDomain(v CreateNetworksRequestNetworkNetworkDomain)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *UpdateNetworkRequestNetwork) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetSearchDomains

`func (o *UpdateNetworkRequestNetwork) GetSearchDomains() string`

GetSearchDomains returns the SearchDomains field if non-nil, zero value otherwise.

### GetSearchDomainsOk

`func (o *UpdateNetworkRequestNetwork) GetSearchDomainsOk() (*string, bool)`

GetSearchDomainsOk returns a tuple with the SearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomains

`func (o *UpdateNetworkRequestNetwork) SetSearchDomains(v string)`

SetSearchDomains sets SearchDomains field to given value.

### HasSearchDomains

`func (o *UpdateNetworkRequestNetwork) HasSearchDomains() bool`

HasSearchDomains returns a boolean if a field has been set.

### GetNetworkProxy

`func (o *UpdateNetworkRequestNetwork) GetNetworkProxy() CreateNetworksRequestNetworkNetworkProxy`

GetNetworkProxy returns the NetworkProxy field if non-nil, zero value otherwise.

### GetNetworkProxyOk

`func (o *UpdateNetworkRequestNetwork) GetNetworkProxyOk() (*CreateNetworksRequestNetworkNetworkProxy, bool)`

GetNetworkProxyOk returns a tuple with the NetworkProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProxy

`func (o *UpdateNetworkRequestNetwork) SetNetworkProxy(v CreateNetworksRequestNetworkNetworkProxy)`

SetNetworkProxy sets NetworkProxy field to given value.

### HasNetworkProxy

`func (o *UpdateNetworkRequestNetwork) HasNetworkProxy() bool`

HasNetworkProxy returns a boolean if a field has been set.

### GetApplianceUrlProxyBypass

`func (o *UpdateNetworkRequestNetwork) GetApplianceUrlProxyBypass() bool`

GetApplianceUrlProxyBypass returns the ApplianceUrlProxyBypass field if non-nil, zero value otherwise.

### GetApplianceUrlProxyBypassOk

`func (o *UpdateNetworkRequestNetwork) GetApplianceUrlProxyBypassOk() (*bool, bool)`

GetApplianceUrlProxyBypassOk returns a tuple with the ApplianceUrlProxyBypass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrlProxyBypass

`func (o *UpdateNetworkRequestNetwork) SetApplianceUrlProxyBypass(v bool)`

SetApplianceUrlProxyBypass sets ApplianceUrlProxyBypass field to given value.

### HasApplianceUrlProxyBypass

`func (o *UpdateNetworkRequestNetwork) HasApplianceUrlProxyBypass() bool`

HasApplianceUrlProxyBypass returns a boolean if a field has been set.

### GetNoProxy

`func (o *UpdateNetworkRequestNetwork) GetNoProxy() string`

GetNoProxy returns the NoProxy field if non-nil, zero value otherwise.

### GetNoProxyOk

`func (o *UpdateNetworkRequestNetwork) GetNoProxyOk() (*string, bool)`

GetNoProxyOk returns a tuple with the NoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxy

`func (o *UpdateNetworkRequestNetwork) SetNoProxy(v string)`

SetNoProxy sets NoProxy field to given value.

### HasNoProxy

`func (o *UpdateNetworkRequestNetwork) HasNoProxy() bool`

HasNoProxy returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateNetworkRequestNetwork) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateNetworkRequestNetwork) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateNetworkRequestNetwork) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateNetworkRequestNetwork) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateNetworkRequestNetwork) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateNetworkRequestNetwork) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateNetworkRequestNetwork) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateNetworkRequestNetwork) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTenants

`func (o *UpdateNetworkRequestNetwork) GetTenants() []GetAlerts200ResponseAllOfChecksInnerAccount`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateNetworkRequestNetwork) GetTenantsOk() (*[]GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateNetworkRequestNetwork) SetTenants(v []GetAlerts200ResponseAllOfChecksInnerAccount)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateNetworkRequestNetwork) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *UpdateNetworkRequestNetwork) GetResourcePermissions() UpdateNetworkRequestNetworkResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *UpdateNetworkRequestNetwork) GetResourcePermissionsOk() (*UpdateNetworkRequestNetworkResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *UpdateNetworkRequestNetwork) SetResourcePermissions(v UpdateNetworkRequestNetworkResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *UpdateNetworkRequestNetwork) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


