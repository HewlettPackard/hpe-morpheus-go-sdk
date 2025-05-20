# CreateNetworksRequestNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name | 
**DisplayName** | Pointer to **string** | Display Name | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**Site** | [**CreateNetworksRequestNetworkSite**](CreateNetworksRequestNetworkSite.md) |  | 
**Zone** | [**CreateNetworksRequestNetworkZone**](CreateNetworksRequestNetworkZone.md) |  | 
**Type** | Pointer to [**CreateNetworksRequestNetworkType**](CreateNetworksRequestNetworkType.md) |  | [optional] 
**Ipv4Enabled** | Pointer to **bool** |  | [optional] 
**Ipv6Enabled** | Pointer to **bool** |  | [optional] 
**Cidr** | Pointer to **string** | CIDR Network | [optional] 
**Gateway** | Pointer to **string** | Network Gateway | [optional] 
**DnsPrimary** | Pointer to **string** | Primary DNS Server | [optional] 
**DnsSecondary** | Pointer to **string** | Secondary DNS Server | [optional] 
**GatewayIPv6** | Pointer to **string** | IPv6 Network Gateway | [optional] 
**NetmaskIPv6** | Pointer to **string** |  | [optional] 
**DnsPrimaryIPv6** | Pointer to **string** | Primary IPv6 DNS Server | [optional] 
**DnsSecondaryIPv6** | Pointer to **string** | Secondary IPv6 DNS Server | [optional] 
**CidrIPv6** | Pointer to **string** | IPv6 Network CIDR | [optional] 
**VlanId** | Pointer to **int64** |  | [optional] 
**Pool** | Pointer to **int64** | Network Pool ID | [optional] 
**PoolIPv6** | Pointer to **int64** | IPv6 Network Pool ID | [optional] 
**AllowStaticOverride** | Pointer to **bool** | Allow IP Override | [optional] 
**AssignPublicIp** | Pointer to **bool** | Assign Public IP | [optional] 
**Active** | Pointer to **bool** | Activate (true) or disable (false) the network | [optional] 
**DhcpServer** | Pointer to **bool** | DHCP Server enabled network | [optional] 
**DhcpServerIPv6** | Pointer to **bool** | IPv6 DHCP Server enabled network | [optional] 
**NetworkDomain** | Pointer to [**CreateNetworksRequestNetworkNetworkDomain**](CreateNetworksRequestNetworkNetworkDomain.md) |  | [optional] 
**SearchDomains** | Pointer to **string** | Search Domains | [optional] 
**NetworkProxy** | Pointer to [**CreateNetworksRequestNetworkNetworkProxy**](CreateNetworksRequestNetworkNetworkProxy.md) |  | [optional] 
**ApplianceUrlProxyBypass** | Pointer to **bool** | Bypass Proxy for Appliance URL | [optional] 
**NoProxy** | Pointer to **string** | Comma-separated list of ip addresses or name servers to exclude proxy traversal for. Typically locally routable servers are excluded. | [optional] 
**Visibility** | Pointer to **string** | Visibility, private or public. | [optional] [default to "private"]
**Config** | Pointer to [**CreateNetworksRequestNetworkConfig**](CreateNetworksRequestNetworkConfig.md) |  | [optional] 
**Tenants** | Pointer to [**[]GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) | Array of tenant account ids that are allowed access | [optional] 
**ResourcePermissions** | Pointer to [**CreateNetworksRequestNetworkResourcePermissions**](CreateNetworksRequestNetworkResourcePermissions.md) |  | [optional] 

## Methods

### NewCreateNetworksRequestNetwork

`func NewCreateNetworksRequestNetwork(name string, site CreateNetworksRequestNetworkSite, zone CreateNetworksRequestNetworkZone, ) *CreateNetworksRequestNetwork`

NewCreateNetworksRequestNetwork instantiates a new CreateNetworksRequestNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworksRequestNetworkWithDefaults

`func NewCreateNetworksRequestNetworkWithDefaults() *CreateNetworksRequestNetwork`

NewCreateNetworksRequestNetworkWithDefaults instantiates a new CreateNetworksRequestNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworksRequestNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworksRequestNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworksRequestNetwork) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *CreateNetworksRequestNetwork) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateNetworksRequestNetwork) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateNetworksRequestNetwork) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateNetworksRequestNetwork) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLabels

`func (o *CreateNetworksRequestNetwork) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateNetworksRequestNetwork) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateNetworksRequestNetwork) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateNetworksRequestNetwork) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetDescription

`func (o *CreateNetworksRequestNetwork) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNetworksRequestNetwork) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNetworksRequestNetwork) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNetworksRequestNetwork) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSite

`func (o *CreateNetworksRequestNetwork) GetSite() CreateNetworksRequestNetworkSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *CreateNetworksRequestNetwork) GetSiteOk() (*CreateNetworksRequestNetworkSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *CreateNetworksRequestNetwork) SetSite(v CreateNetworksRequestNetworkSite)`

SetSite sets Site field to given value.


### GetZone

`func (o *CreateNetworksRequestNetwork) GetZone() CreateNetworksRequestNetworkZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CreateNetworksRequestNetwork) GetZoneOk() (*CreateNetworksRequestNetworkZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CreateNetworksRequestNetwork) SetZone(v CreateNetworksRequestNetworkZone)`

SetZone sets Zone field to given value.


### GetType

`func (o *CreateNetworksRequestNetwork) GetType() CreateNetworksRequestNetworkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNetworksRequestNetwork) GetTypeOk() (*CreateNetworksRequestNetworkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNetworksRequestNetwork) SetType(v CreateNetworksRequestNetworkType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateNetworksRequestNetwork) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIpv4Enabled

`func (o *CreateNetworksRequestNetwork) GetIpv4Enabled() bool`

GetIpv4Enabled returns the Ipv4Enabled field if non-nil, zero value otherwise.

### GetIpv4EnabledOk

`func (o *CreateNetworksRequestNetwork) GetIpv4EnabledOk() (*bool, bool)`

GetIpv4EnabledOk returns a tuple with the Ipv4Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Enabled

`func (o *CreateNetworksRequestNetwork) SetIpv4Enabled(v bool)`

SetIpv4Enabled sets Ipv4Enabled field to given value.

### HasIpv4Enabled

`func (o *CreateNetworksRequestNetwork) HasIpv4Enabled() bool`

HasIpv4Enabled returns a boolean if a field has been set.

### GetIpv6Enabled

`func (o *CreateNetworksRequestNetwork) GetIpv6Enabled() bool`

GetIpv6Enabled returns the Ipv6Enabled field if non-nil, zero value otherwise.

### GetIpv6EnabledOk

`func (o *CreateNetworksRequestNetwork) GetIpv6EnabledOk() (*bool, bool)`

GetIpv6EnabledOk returns a tuple with the Ipv6Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Enabled

`func (o *CreateNetworksRequestNetwork) SetIpv6Enabled(v bool)`

SetIpv6Enabled sets Ipv6Enabled field to given value.

### HasIpv6Enabled

`func (o *CreateNetworksRequestNetwork) HasIpv6Enabled() bool`

HasIpv6Enabled returns a boolean if a field has been set.

### GetCidr

`func (o *CreateNetworksRequestNetwork) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CreateNetworksRequestNetwork) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CreateNetworksRequestNetwork) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *CreateNetworksRequestNetwork) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetGateway

`func (o *CreateNetworksRequestNetwork) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *CreateNetworksRequestNetwork) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *CreateNetworksRequestNetwork) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *CreateNetworksRequestNetwork) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetDnsPrimary

`func (o *CreateNetworksRequestNetwork) GetDnsPrimary() string`

GetDnsPrimary returns the DnsPrimary field if non-nil, zero value otherwise.

### GetDnsPrimaryOk

`func (o *CreateNetworksRequestNetwork) GetDnsPrimaryOk() (*string, bool)`

GetDnsPrimaryOk returns a tuple with the DnsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrimary

`func (o *CreateNetworksRequestNetwork) SetDnsPrimary(v string)`

SetDnsPrimary sets DnsPrimary field to given value.

### HasDnsPrimary

`func (o *CreateNetworksRequestNetwork) HasDnsPrimary() bool`

HasDnsPrimary returns a boolean if a field has been set.

### GetDnsSecondary

`func (o *CreateNetworksRequestNetwork) GetDnsSecondary() string`

GetDnsSecondary returns the DnsSecondary field if non-nil, zero value otherwise.

### GetDnsSecondaryOk

`func (o *CreateNetworksRequestNetwork) GetDnsSecondaryOk() (*string, bool)`

GetDnsSecondaryOk returns a tuple with the DnsSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSecondary

`func (o *CreateNetworksRequestNetwork) SetDnsSecondary(v string)`

SetDnsSecondary sets DnsSecondary field to given value.

### HasDnsSecondary

`func (o *CreateNetworksRequestNetwork) HasDnsSecondary() bool`

HasDnsSecondary returns a boolean if a field has been set.

### GetGatewayIPv6

`func (o *CreateNetworksRequestNetwork) GetGatewayIPv6() string`

GetGatewayIPv6 returns the GatewayIPv6 field if non-nil, zero value otherwise.

### GetGatewayIPv6Ok

`func (o *CreateNetworksRequestNetwork) GetGatewayIPv6Ok() (*string, bool)`

GetGatewayIPv6Ok returns a tuple with the GatewayIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIPv6

`func (o *CreateNetworksRequestNetwork) SetGatewayIPv6(v string)`

SetGatewayIPv6 sets GatewayIPv6 field to given value.

### HasGatewayIPv6

`func (o *CreateNetworksRequestNetwork) HasGatewayIPv6() bool`

HasGatewayIPv6 returns a boolean if a field has been set.

### GetNetmaskIPv6

`func (o *CreateNetworksRequestNetwork) GetNetmaskIPv6() string`

GetNetmaskIPv6 returns the NetmaskIPv6 field if non-nil, zero value otherwise.

### GetNetmaskIPv6Ok

`func (o *CreateNetworksRequestNetwork) GetNetmaskIPv6Ok() (*string, bool)`

GetNetmaskIPv6Ok returns a tuple with the NetmaskIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmaskIPv6

`func (o *CreateNetworksRequestNetwork) SetNetmaskIPv6(v string)`

SetNetmaskIPv6 sets NetmaskIPv6 field to given value.

### HasNetmaskIPv6

`func (o *CreateNetworksRequestNetwork) HasNetmaskIPv6() bool`

HasNetmaskIPv6 returns a boolean if a field has been set.

### GetDnsPrimaryIPv6

`func (o *CreateNetworksRequestNetwork) GetDnsPrimaryIPv6() string`

GetDnsPrimaryIPv6 returns the DnsPrimaryIPv6 field if non-nil, zero value otherwise.

### GetDnsPrimaryIPv6Ok

`func (o *CreateNetworksRequestNetwork) GetDnsPrimaryIPv6Ok() (*string, bool)`

GetDnsPrimaryIPv6Ok returns a tuple with the DnsPrimaryIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrimaryIPv6

`func (o *CreateNetworksRequestNetwork) SetDnsPrimaryIPv6(v string)`

SetDnsPrimaryIPv6 sets DnsPrimaryIPv6 field to given value.

### HasDnsPrimaryIPv6

`func (o *CreateNetworksRequestNetwork) HasDnsPrimaryIPv6() bool`

HasDnsPrimaryIPv6 returns a boolean if a field has been set.

### GetDnsSecondaryIPv6

`func (o *CreateNetworksRequestNetwork) GetDnsSecondaryIPv6() string`

GetDnsSecondaryIPv6 returns the DnsSecondaryIPv6 field if non-nil, zero value otherwise.

### GetDnsSecondaryIPv6Ok

`func (o *CreateNetworksRequestNetwork) GetDnsSecondaryIPv6Ok() (*string, bool)`

GetDnsSecondaryIPv6Ok returns a tuple with the DnsSecondaryIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSecondaryIPv6

`func (o *CreateNetworksRequestNetwork) SetDnsSecondaryIPv6(v string)`

SetDnsSecondaryIPv6 sets DnsSecondaryIPv6 field to given value.

### HasDnsSecondaryIPv6

`func (o *CreateNetworksRequestNetwork) HasDnsSecondaryIPv6() bool`

HasDnsSecondaryIPv6 returns a boolean if a field has been set.

### GetCidrIPv6

`func (o *CreateNetworksRequestNetwork) GetCidrIPv6() string`

GetCidrIPv6 returns the CidrIPv6 field if non-nil, zero value otherwise.

### GetCidrIPv6Ok

`func (o *CreateNetworksRequestNetwork) GetCidrIPv6Ok() (*string, bool)`

GetCidrIPv6Ok returns a tuple with the CidrIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrIPv6

`func (o *CreateNetworksRequestNetwork) SetCidrIPv6(v string)`

SetCidrIPv6 sets CidrIPv6 field to given value.

### HasCidrIPv6

`func (o *CreateNetworksRequestNetwork) HasCidrIPv6() bool`

HasCidrIPv6 returns a boolean if a field has been set.

### GetVlanId

`func (o *CreateNetworksRequestNetwork) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *CreateNetworksRequestNetwork) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *CreateNetworksRequestNetwork) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *CreateNetworksRequestNetwork) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetPool

`func (o *CreateNetworksRequestNetwork) GetPool() int64`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *CreateNetworksRequestNetwork) GetPoolOk() (*int64, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *CreateNetworksRequestNetwork) SetPool(v int64)`

SetPool sets Pool field to given value.

### HasPool

`func (o *CreateNetworksRequestNetwork) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolIPv6

`func (o *CreateNetworksRequestNetwork) GetPoolIPv6() int64`

GetPoolIPv6 returns the PoolIPv6 field if non-nil, zero value otherwise.

### GetPoolIPv6Ok

`func (o *CreateNetworksRequestNetwork) GetPoolIPv6Ok() (*int64, bool)`

GetPoolIPv6Ok returns a tuple with the PoolIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolIPv6

`func (o *CreateNetworksRequestNetwork) SetPoolIPv6(v int64)`

SetPoolIPv6 sets PoolIPv6 field to given value.

### HasPoolIPv6

`func (o *CreateNetworksRequestNetwork) HasPoolIPv6() bool`

HasPoolIPv6 returns a boolean if a field has been set.

### GetAllowStaticOverride

`func (o *CreateNetworksRequestNetwork) GetAllowStaticOverride() bool`

GetAllowStaticOverride returns the AllowStaticOverride field if non-nil, zero value otherwise.

### GetAllowStaticOverrideOk

`func (o *CreateNetworksRequestNetwork) GetAllowStaticOverrideOk() (*bool, bool)`

GetAllowStaticOverrideOk returns a tuple with the AllowStaticOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowStaticOverride

`func (o *CreateNetworksRequestNetwork) SetAllowStaticOverride(v bool)`

SetAllowStaticOverride sets AllowStaticOverride field to given value.

### HasAllowStaticOverride

`func (o *CreateNetworksRequestNetwork) HasAllowStaticOverride() bool`

HasAllowStaticOverride returns a boolean if a field has been set.

### GetAssignPublicIp

`func (o *CreateNetworksRequestNetwork) GetAssignPublicIp() bool`

GetAssignPublicIp returns the AssignPublicIp field if non-nil, zero value otherwise.

### GetAssignPublicIpOk

`func (o *CreateNetworksRequestNetwork) GetAssignPublicIpOk() (*bool, bool)`

GetAssignPublicIpOk returns a tuple with the AssignPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignPublicIp

`func (o *CreateNetworksRequestNetwork) SetAssignPublicIp(v bool)`

SetAssignPublicIp sets AssignPublicIp field to given value.

### HasAssignPublicIp

`func (o *CreateNetworksRequestNetwork) HasAssignPublicIp() bool`

HasAssignPublicIp returns a boolean if a field has been set.

### GetActive

`func (o *CreateNetworksRequestNetwork) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateNetworksRequestNetwork) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateNetworksRequestNetwork) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateNetworksRequestNetwork) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDhcpServer

`func (o *CreateNetworksRequestNetwork) GetDhcpServer() bool`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *CreateNetworksRequestNetwork) GetDhcpServerOk() (*bool, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *CreateNetworksRequestNetwork) SetDhcpServer(v bool)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *CreateNetworksRequestNetwork) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetDhcpServerIPv6

`func (o *CreateNetworksRequestNetwork) GetDhcpServerIPv6() bool`

GetDhcpServerIPv6 returns the DhcpServerIPv6 field if non-nil, zero value otherwise.

### GetDhcpServerIPv6Ok

`func (o *CreateNetworksRequestNetwork) GetDhcpServerIPv6Ok() (*bool, bool)`

GetDhcpServerIPv6Ok returns a tuple with the DhcpServerIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerIPv6

`func (o *CreateNetworksRequestNetwork) SetDhcpServerIPv6(v bool)`

SetDhcpServerIPv6 sets DhcpServerIPv6 field to given value.

### HasDhcpServerIPv6

`func (o *CreateNetworksRequestNetwork) HasDhcpServerIPv6() bool`

HasDhcpServerIPv6 returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *CreateNetworksRequestNetwork) GetNetworkDomain() CreateNetworksRequestNetworkNetworkDomain`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *CreateNetworksRequestNetwork) GetNetworkDomainOk() (*CreateNetworksRequestNetworkNetworkDomain, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *CreateNetworksRequestNetwork) SetNetworkDomain(v CreateNetworksRequestNetworkNetworkDomain)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *CreateNetworksRequestNetwork) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetSearchDomains

`func (o *CreateNetworksRequestNetwork) GetSearchDomains() string`

GetSearchDomains returns the SearchDomains field if non-nil, zero value otherwise.

### GetSearchDomainsOk

`func (o *CreateNetworksRequestNetwork) GetSearchDomainsOk() (*string, bool)`

GetSearchDomainsOk returns a tuple with the SearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomains

`func (o *CreateNetworksRequestNetwork) SetSearchDomains(v string)`

SetSearchDomains sets SearchDomains field to given value.

### HasSearchDomains

`func (o *CreateNetworksRequestNetwork) HasSearchDomains() bool`

HasSearchDomains returns a boolean if a field has been set.

### GetNetworkProxy

`func (o *CreateNetworksRequestNetwork) GetNetworkProxy() CreateNetworksRequestNetworkNetworkProxy`

GetNetworkProxy returns the NetworkProxy field if non-nil, zero value otherwise.

### GetNetworkProxyOk

`func (o *CreateNetworksRequestNetwork) GetNetworkProxyOk() (*CreateNetworksRequestNetworkNetworkProxy, bool)`

GetNetworkProxyOk returns a tuple with the NetworkProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProxy

`func (o *CreateNetworksRequestNetwork) SetNetworkProxy(v CreateNetworksRequestNetworkNetworkProxy)`

SetNetworkProxy sets NetworkProxy field to given value.

### HasNetworkProxy

`func (o *CreateNetworksRequestNetwork) HasNetworkProxy() bool`

HasNetworkProxy returns a boolean if a field has been set.

### GetApplianceUrlProxyBypass

`func (o *CreateNetworksRequestNetwork) GetApplianceUrlProxyBypass() bool`

GetApplianceUrlProxyBypass returns the ApplianceUrlProxyBypass field if non-nil, zero value otherwise.

### GetApplianceUrlProxyBypassOk

`func (o *CreateNetworksRequestNetwork) GetApplianceUrlProxyBypassOk() (*bool, bool)`

GetApplianceUrlProxyBypassOk returns a tuple with the ApplianceUrlProxyBypass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrlProxyBypass

`func (o *CreateNetworksRequestNetwork) SetApplianceUrlProxyBypass(v bool)`

SetApplianceUrlProxyBypass sets ApplianceUrlProxyBypass field to given value.

### HasApplianceUrlProxyBypass

`func (o *CreateNetworksRequestNetwork) HasApplianceUrlProxyBypass() bool`

HasApplianceUrlProxyBypass returns a boolean if a field has been set.

### GetNoProxy

`func (o *CreateNetworksRequestNetwork) GetNoProxy() string`

GetNoProxy returns the NoProxy field if non-nil, zero value otherwise.

### GetNoProxyOk

`func (o *CreateNetworksRequestNetwork) GetNoProxyOk() (*string, bool)`

GetNoProxyOk returns a tuple with the NoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxy

`func (o *CreateNetworksRequestNetwork) SetNoProxy(v string)`

SetNoProxy sets NoProxy field to given value.

### HasNoProxy

`func (o *CreateNetworksRequestNetwork) HasNoProxy() bool`

HasNoProxy returns a boolean if a field has been set.

### GetVisibility

`func (o *CreateNetworksRequestNetwork) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateNetworksRequestNetwork) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateNetworksRequestNetwork) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateNetworksRequestNetwork) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetConfig

`func (o *CreateNetworksRequestNetwork) GetConfig() CreateNetworksRequestNetworkConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateNetworksRequestNetwork) GetConfigOk() (*CreateNetworksRequestNetworkConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateNetworksRequestNetwork) SetConfig(v CreateNetworksRequestNetworkConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateNetworksRequestNetwork) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTenants

`func (o *CreateNetworksRequestNetwork) GetTenants() []GetAlerts200ResponseAllOfChecksInnerAccount`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *CreateNetworksRequestNetwork) GetTenantsOk() (*[]GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *CreateNetworksRequestNetwork) SetTenants(v []GetAlerts200ResponseAllOfChecksInnerAccount)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *CreateNetworksRequestNetwork) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *CreateNetworksRequestNetwork) GetResourcePermissions() CreateNetworksRequestNetworkResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *CreateNetworksRequestNetwork) GetResourcePermissionsOk() (*CreateNetworksRequestNetworkResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *CreateNetworksRequestNetwork) SetResourcePermissions(v CreateNetworksRequestNetworkResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *CreateNetworksRequestNetwork) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


