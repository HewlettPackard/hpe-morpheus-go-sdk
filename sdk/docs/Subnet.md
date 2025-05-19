# Subnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**AddressPrefix** | Pointer to **string** |  | [optional] 
**Cidr** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**SubnetAddress** | Pointer to **string** |  | [optional] 
**TftpServer** | Pointer to **string** |  | [optional] 
**BootFile** | Pointer to **string** |  | [optional] 
**Pool** | Pointer to **string** |  | [optional] 
**DhcpServer** | Pointer to **bool** |  | [optional] 
**HasFloatingIps** | Pointer to **bool** |  | [optional] 
**DhcpIp** | Pointer to **string** |  | [optional] 
**DnsPrimary** | Pointer to **string** |  | [optional] 
**DnsSecondary** | Pointer to **string** |  | [optional] 
**DhcpStart** | Pointer to **string** |  | [optional] 
**DhcpEnd** | Pointer to **string** |  | [optional] 
**DhcpRange** | Pointer to **string** |  | [optional] 
**NetworkProxy** | Pointer to **string** |  | [optional] 
**NetworkDomain** | Pointer to **string** |  | [optional] 
**SearchDomains** | Pointer to **string** |  | [optional] 
**DefaultNetwork** | Pointer to **bool** |  | [optional] 
**AssignPublicIp** | Pointer to **bool** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**GetAppState200ResponseAllOfInputProvidersInner**](GetAppState200ResponseAllOfInputProvidersInner.md) |  | [optional] 
**Network** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**SecurityGroups** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tenants** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**ResourcePermission** | Pointer to [**GetNetworkSubnets200ResponseAllOfSubnetsInnerResourcePermission**](GetNetworkSubnets200ResponseAllOfSubnetsInnerResourcePermission.md) |  | [optional] 

## Methods

### NewSubnet

`func NewSubnet() *Subnet`

NewSubnet instantiates a new Subnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetWithDefaults

`func NewSubnetWithDefaults() *Subnet`

NewSubnetWithDefaults instantiates a new Subnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subnet) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subnet) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subnet) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Subnet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *Subnet) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Subnet) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Subnet) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Subnet) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *Subnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Subnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Subnet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Subnet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *Subnet) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Subnet) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Subnet) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Subnet) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetActive

`func (o *Subnet) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Subnet) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Subnet) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Subnet) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDescription

`func (o *Subnet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Subnet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Subnet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Subnet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalId

`func (o *Subnet) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Subnet) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Subnet) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Subnet) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetUniqueId

`func (o *Subnet) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *Subnet) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *Subnet) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *Subnet) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetAddressPrefix

`func (o *Subnet) GetAddressPrefix() string`

GetAddressPrefix returns the AddressPrefix field if non-nil, zero value otherwise.

### GetAddressPrefixOk

`func (o *Subnet) GetAddressPrefixOk() (*string, bool)`

GetAddressPrefixOk returns a tuple with the AddressPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPrefix

`func (o *Subnet) SetAddressPrefix(v string)`

SetAddressPrefix sets AddressPrefix field to given value.

### HasAddressPrefix

`func (o *Subnet) HasAddressPrefix() bool`

HasAddressPrefix returns a boolean if a field has been set.

### GetCidr

`func (o *Subnet) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *Subnet) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *Subnet) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *Subnet) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetGateway

`func (o *Subnet) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Subnet) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Subnet) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *Subnet) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetNetmask

`func (o *Subnet) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *Subnet) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *Subnet) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *Subnet) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetSubnetAddress

`func (o *Subnet) GetSubnetAddress() string`

GetSubnetAddress returns the SubnetAddress field if non-nil, zero value otherwise.

### GetSubnetAddressOk

`func (o *Subnet) GetSubnetAddressOk() (*string, bool)`

GetSubnetAddressOk returns a tuple with the SubnetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAddress

`func (o *Subnet) SetSubnetAddress(v string)`

SetSubnetAddress sets SubnetAddress field to given value.

### HasSubnetAddress

`func (o *Subnet) HasSubnetAddress() bool`

HasSubnetAddress returns a boolean if a field has been set.

### GetTftpServer

`func (o *Subnet) GetTftpServer() string`

GetTftpServer returns the TftpServer field if non-nil, zero value otherwise.

### GetTftpServerOk

`func (o *Subnet) GetTftpServerOk() (*string, bool)`

GetTftpServerOk returns a tuple with the TftpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTftpServer

`func (o *Subnet) SetTftpServer(v string)`

SetTftpServer sets TftpServer field to given value.

### HasTftpServer

`func (o *Subnet) HasTftpServer() bool`

HasTftpServer returns a boolean if a field has been set.

### GetBootFile

`func (o *Subnet) GetBootFile() string`

GetBootFile returns the BootFile field if non-nil, zero value otherwise.

### GetBootFileOk

`func (o *Subnet) GetBootFileOk() (*string, bool)`

GetBootFileOk returns a tuple with the BootFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootFile

`func (o *Subnet) SetBootFile(v string)`

SetBootFile sets BootFile field to given value.

### HasBootFile

`func (o *Subnet) HasBootFile() bool`

HasBootFile returns a boolean if a field has been set.

### GetPool

`func (o *Subnet) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *Subnet) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *Subnet) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *Subnet) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetDhcpServer

`func (o *Subnet) GetDhcpServer() bool`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *Subnet) GetDhcpServerOk() (*bool, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *Subnet) SetDhcpServer(v bool)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *Subnet) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetHasFloatingIps

`func (o *Subnet) GetHasFloatingIps() bool`

GetHasFloatingIps returns the HasFloatingIps field if non-nil, zero value otherwise.

### GetHasFloatingIpsOk

`func (o *Subnet) GetHasFloatingIpsOk() (*bool, bool)`

GetHasFloatingIpsOk returns a tuple with the HasFloatingIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFloatingIps

`func (o *Subnet) SetHasFloatingIps(v bool)`

SetHasFloatingIps sets HasFloatingIps field to given value.

### HasHasFloatingIps

`func (o *Subnet) HasHasFloatingIps() bool`

HasHasFloatingIps returns a boolean if a field has been set.

### GetDhcpIp

`func (o *Subnet) GetDhcpIp() string`

GetDhcpIp returns the DhcpIp field if non-nil, zero value otherwise.

### GetDhcpIpOk

`func (o *Subnet) GetDhcpIpOk() (*string, bool)`

GetDhcpIpOk returns a tuple with the DhcpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpIp

`func (o *Subnet) SetDhcpIp(v string)`

SetDhcpIp sets DhcpIp field to given value.

### HasDhcpIp

`func (o *Subnet) HasDhcpIp() bool`

HasDhcpIp returns a boolean if a field has been set.

### GetDnsPrimary

`func (o *Subnet) GetDnsPrimary() string`

GetDnsPrimary returns the DnsPrimary field if non-nil, zero value otherwise.

### GetDnsPrimaryOk

`func (o *Subnet) GetDnsPrimaryOk() (*string, bool)`

GetDnsPrimaryOk returns a tuple with the DnsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrimary

`func (o *Subnet) SetDnsPrimary(v string)`

SetDnsPrimary sets DnsPrimary field to given value.

### HasDnsPrimary

`func (o *Subnet) HasDnsPrimary() bool`

HasDnsPrimary returns a boolean if a field has been set.

### GetDnsSecondary

`func (o *Subnet) GetDnsSecondary() string`

GetDnsSecondary returns the DnsSecondary field if non-nil, zero value otherwise.

### GetDnsSecondaryOk

`func (o *Subnet) GetDnsSecondaryOk() (*string, bool)`

GetDnsSecondaryOk returns a tuple with the DnsSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSecondary

`func (o *Subnet) SetDnsSecondary(v string)`

SetDnsSecondary sets DnsSecondary field to given value.

### HasDnsSecondary

`func (o *Subnet) HasDnsSecondary() bool`

HasDnsSecondary returns a boolean if a field has been set.

### GetDhcpStart

`func (o *Subnet) GetDhcpStart() string`

GetDhcpStart returns the DhcpStart field if non-nil, zero value otherwise.

### GetDhcpStartOk

`func (o *Subnet) GetDhcpStartOk() (*string, bool)`

GetDhcpStartOk returns a tuple with the DhcpStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpStart

`func (o *Subnet) SetDhcpStart(v string)`

SetDhcpStart sets DhcpStart field to given value.

### HasDhcpStart

`func (o *Subnet) HasDhcpStart() bool`

HasDhcpStart returns a boolean if a field has been set.

### GetDhcpEnd

`func (o *Subnet) GetDhcpEnd() string`

GetDhcpEnd returns the DhcpEnd field if non-nil, zero value otherwise.

### GetDhcpEndOk

`func (o *Subnet) GetDhcpEndOk() (*string, bool)`

GetDhcpEndOk returns a tuple with the DhcpEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpEnd

`func (o *Subnet) SetDhcpEnd(v string)`

SetDhcpEnd sets DhcpEnd field to given value.

### HasDhcpEnd

`func (o *Subnet) HasDhcpEnd() bool`

HasDhcpEnd returns a boolean if a field has been set.

### GetDhcpRange

`func (o *Subnet) GetDhcpRange() string`

GetDhcpRange returns the DhcpRange field if non-nil, zero value otherwise.

### GetDhcpRangeOk

`func (o *Subnet) GetDhcpRangeOk() (*string, bool)`

GetDhcpRangeOk returns a tuple with the DhcpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRange

`func (o *Subnet) SetDhcpRange(v string)`

SetDhcpRange sets DhcpRange field to given value.

### HasDhcpRange

`func (o *Subnet) HasDhcpRange() bool`

HasDhcpRange returns a boolean if a field has been set.

### GetNetworkProxy

`func (o *Subnet) GetNetworkProxy() string`

GetNetworkProxy returns the NetworkProxy field if non-nil, zero value otherwise.

### GetNetworkProxyOk

`func (o *Subnet) GetNetworkProxyOk() (*string, bool)`

GetNetworkProxyOk returns a tuple with the NetworkProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProxy

`func (o *Subnet) SetNetworkProxy(v string)`

SetNetworkProxy sets NetworkProxy field to given value.

### HasNetworkProxy

`func (o *Subnet) HasNetworkProxy() bool`

HasNetworkProxy returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *Subnet) GetNetworkDomain() string`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *Subnet) GetNetworkDomainOk() (*string, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *Subnet) SetNetworkDomain(v string)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *Subnet) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetSearchDomains

`func (o *Subnet) GetSearchDomains() string`

GetSearchDomains returns the SearchDomains field if non-nil, zero value otherwise.

### GetSearchDomainsOk

`func (o *Subnet) GetSearchDomainsOk() (*string, bool)`

GetSearchDomainsOk returns a tuple with the SearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomains

`func (o *Subnet) SetSearchDomains(v string)`

SetSearchDomains sets SearchDomains field to given value.

### HasSearchDomains

`func (o *Subnet) HasSearchDomains() bool`

HasSearchDomains returns a boolean if a field has been set.

### GetDefaultNetwork

`func (o *Subnet) GetDefaultNetwork() bool`

GetDefaultNetwork returns the DefaultNetwork field if non-nil, zero value otherwise.

### GetDefaultNetworkOk

`func (o *Subnet) GetDefaultNetworkOk() (*bool, bool)`

GetDefaultNetworkOk returns a tuple with the DefaultNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetwork

`func (o *Subnet) SetDefaultNetwork(v bool)`

SetDefaultNetwork sets DefaultNetwork field to given value.

### HasDefaultNetwork

`func (o *Subnet) HasDefaultNetwork() bool`

HasDefaultNetwork returns a boolean if a field has been set.

### GetAssignPublicIp

`func (o *Subnet) GetAssignPublicIp() bool`

GetAssignPublicIp returns the AssignPublicIp field if non-nil, zero value otherwise.

### GetAssignPublicIpOk

`func (o *Subnet) GetAssignPublicIpOk() (*bool, bool)`

GetAssignPublicIpOk returns a tuple with the AssignPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignPublicIp

`func (o *Subnet) SetAssignPublicIp(v bool)`

SetAssignPublicIp sets AssignPublicIp field to given value.

### HasAssignPublicIp

`func (o *Subnet) HasAssignPublicIp() bool`

HasAssignPublicIp returns a boolean if a field has been set.

### GetVisibility

`func (o *Subnet) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Subnet) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Subnet) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Subnet) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetStatus

`func (o *Subnet) GetStatus() GetAppState200ResponseAllOfInputProvidersInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Subnet) GetStatusOk() (*GetAppState200ResponseAllOfInputProvidersInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Subnet) SetStatus(v GetAppState200ResponseAllOfInputProvidersInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Subnet) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNetwork

`func (o *Subnet) GetNetwork() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Subnet) GetNetworkOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Subnet) SetNetwork(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Subnet) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetType

`func (o *Subnet) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Subnet) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Subnet) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *Subnet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccount

`func (o *Subnet) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Subnet) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Subnet) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Subnet) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *Subnet) GetSecurityGroups() []map[string]interface{}`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *Subnet) GetSecurityGroupsOk() (*[]map[string]interface{}, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *Subnet) SetSecurityGroups(v []map[string]interface{})`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *Subnet) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetTenants

`func (o *Subnet) GetTenants() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *Subnet) GetTenantsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *Subnet) SetTenants(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *Subnet) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *Subnet) GetResourcePermission() GetNetworkSubnets200ResponseAllOfSubnetsInnerResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *Subnet) GetResourcePermissionOk() (*GetNetworkSubnets200ResponseAllOfSubnetsInnerResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *Subnet) SetResourcePermission(v GetNetworkSubnets200ResponseAllOfSubnetsInnerResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *Subnet) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


