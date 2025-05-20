# GetNetworkSubnets200ResponseAllOfSubnetsInner

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

### NewGetNetworkSubnets200ResponseAllOfSubnetsInner

`func NewGetNetworkSubnets200ResponseAllOfSubnetsInner() *GetNetworkSubnets200ResponseAllOfSubnetsInner`

NewGetNetworkSubnets200ResponseAllOfSubnetsInner instantiates a new GetNetworkSubnets200ResponseAllOfSubnetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSubnets200ResponseAllOfSubnetsInnerWithDefaults

`func NewGetNetworkSubnets200ResponseAllOfSubnetsInnerWithDefaults() *GetNetworkSubnets200ResponseAllOfSubnetsInner`

NewGetNetworkSubnets200ResponseAllOfSubnetsInnerWithDefaults instantiates a new GetNetworkSubnets200ResponseAllOfSubnetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetActive

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDescription

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalId

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetUniqueId

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetAddressPrefix

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetAddressPrefix() string`

GetAddressPrefix returns the AddressPrefix field if non-nil, zero value otherwise.

### GetAddressPrefixOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetAddressPrefixOk() (*string, bool)`

GetAddressPrefixOk returns a tuple with the AddressPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPrefix

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetAddressPrefix(v string)`

SetAddressPrefix sets AddressPrefix field to given value.

### HasAddressPrefix

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasAddressPrefix() bool`

HasAddressPrefix returns a boolean if a field has been set.

### GetCidr

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetGateway

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetNetmask

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetSubnetAddress

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetSubnetAddress() string`

GetSubnetAddress returns the SubnetAddress field if non-nil, zero value otherwise.

### GetSubnetAddressOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetSubnetAddressOk() (*string, bool)`

GetSubnetAddressOk returns a tuple with the SubnetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAddress

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetSubnetAddress(v string)`

SetSubnetAddress sets SubnetAddress field to given value.

### HasSubnetAddress

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasSubnetAddress() bool`

HasSubnetAddress returns a boolean if a field has been set.

### GetTftpServer

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetTftpServer() string`

GetTftpServer returns the TftpServer field if non-nil, zero value otherwise.

### GetTftpServerOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetTftpServerOk() (*string, bool)`

GetTftpServerOk returns a tuple with the TftpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTftpServer

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetTftpServer(v string)`

SetTftpServer sets TftpServer field to given value.

### HasTftpServer

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasTftpServer() bool`

HasTftpServer returns a boolean if a field has been set.

### GetBootFile

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetBootFile() string`

GetBootFile returns the BootFile field if non-nil, zero value otherwise.

### GetBootFileOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetBootFileOk() (*string, bool)`

GetBootFileOk returns a tuple with the BootFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootFile

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetBootFile(v string)`

SetBootFile sets BootFile field to given value.

### HasBootFile

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasBootFile() bool`

HasBootFile returns a boolean if a field has been set.

### GetPool

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetDhcpServer

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetDhcpServer() bool`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetDhcpServerOk() (*bool, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetDhcpServer(v bool)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetHasFloatingIps

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetHasFloatingIps() bool`

GetHasFloatingIps returns the HasFloatingIps field if non-nil, zero value otherwise.

### GetHasFloatingIpsOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetHasFloatingIpsOk() (*bool, bool)`

GetHasFloatingIpsOk returns a tuple with the HasFloatingIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFloatingIps

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetHasFloatingIps(v bool)`

SetHasFloatingIps sets HasFloatingIps field to given value.

### HasHasFloatingIps

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasHasFloatingIps() bool`

HasHasFloatingIps returns a boolean if a field has been set.

### GetDhcpIp

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetDhcpIp() string`

GetDhcpIp returns the DhcpIp field if non-nil, zero value otherwise.

### GetDhcpIpOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetDhcpIpOk() (*string, bool)`

GetDhcpIpOk returns a tuple with the DhcpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpIp

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetDhcpIp(v string)`

SetDhcpIp sets DhcpIp field to given value.

### HasDhcpIp

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasDhcpIp() bool`

HasDhcpIp returns a boolean if a field has been set.

### GetDnsPrimary

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetDnsPrimary() string`

GetDnsPrimary returns the DnsPrimary field if non-nil, zero value otherwise.

### GetDnsPrimaryOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetDnsPrimaryOk() (*string, bool)`

GetDnsPrimaryOk returns a tuple with the DnsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrimary

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetDnsPrimary(v string)`

SetDnsPrimary sets DnsPrimary field to given value.

### HasDnsPrimary

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasDnsPrimary() bool`

HasDnsPrimary returns a boolean if a field has been set.

### GetDnsSecondary

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetDnsSecondary() string`

GetDnsSecondary returns the DnsSecondary field if non-nil, zero value otherwise.

### GetDnsSecondaryOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetDnsSecondaryOk() (*string, bool)`

GetDnsSecondaryOk returns a tuple with the DnsSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSecondary

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetDnsSecondary(v string)`

SetDnsSecondary sets DnsSecondary field to given value.

### HasDnsSecondary

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasDnsSecondary() bool`

HasDnsSecondary returns a boolean if a field has been set.

### GetDhcpStart

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetDhcpStart() string`

GetDhcpStart returns the DhcpStart field if non-nil, zero value otherwise.

### GetDhcpStartOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetDhcpStartOk() (*string, bool)`

GetDhcpStartOk returns a tuple with the DhcpStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpStart

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetDhcpStart(v string)`

SetDhcpStart sets DhcpStart field to given value.

### HasDhcpStart

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasDhcpStart() bool`

HasDhcpStart returns a boolean if a field has been set.

### GetDhcpEnd

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetDhcpEnd() string`

GetDhcpEnd returns the DhcpEnd field if non-nil, zero value otherwise.

### GetDhcpEndOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetDhcpEndOk() (*string, bool)`

GetDhcpEndOk returns a tuple with the DhcpEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpEnd

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetDhcpEnd(v string)`

SetDhcpEnd sets DhcpEnd field to given value.

### HasDhcpEnd

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasDhcpEnd() bool`

HasDhcpEnd returns a boolean if a field has been set.

### GetDhcpRange

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetDhcpRange() string`

GetDhcpRange returns the DhcpRange field if non-nil, zero value otherwise.

### GetDhcpRangeOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetDhcpRangeOk() (*string, bool)`

GetDhcpRangeOk returns a tuple with the DhcpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRange

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetDhcpRange(v string)`

SetDhcpRange sets DhcpRange field to given value.

### HasDhcpRange

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasDhcpRange() bool`

HasDhcpRange returns a boolean if a field has been set.

### GetNetworkProxy

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetNetworkProxy() string`

GetNetworkProxy returns the NetworkProxy field if non-nil, zero value otherwise.

### GetNetworkProxyOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetNetworkProxyOk() (*string, bool)`

GetNetworkProxyOk returns a tuple with the NetworkProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProxy

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetNetworkProxy(v string)`

SetNetworkProxy sets NetworkProxy field to given value.

### HasNetworkProxy

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasNetworkProxy() bool`

HasNetworkProxy returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetNetworkDomain() string`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetNetworkDomainOk() (*string, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetNetworkDomain(v string)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetSearchDomains

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetSearchDomains() string`

GetSearchDomains returns the SearchDomains field if non-nil, zero value otherwise.

### GetSearchDomainsOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetSearchDomainsOk() (*string, bool)`

GetSearchDomainsOk returns a tuple with the SearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomains

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetSearchDomains(v string)`

SetSearchDomains sets SearchDomains field to given value.

### HasSearchDomains

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasSearchDomains() bool`

HasSearchDomains returns a boolean if a field has been set.

### GetDefaultNetwork

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetDefaultNetwork() bool`

GetDefaultNetwork returns the DefaultNetwork field if non-nil, zero value otherwise.

### GetDefaultNetworkOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetDefaultNetworkOk() (*bool, bool)`

GetDefaultNetworkOk returns a tuple with the DefaultNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetwork

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetDefaultNetwork(v bool)`

SetDefaultNetwork sets DefaultNetwork field to given value.

### HasDefaultNetwork

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasDefaultNetwork() bool`

HasDefaultNetwork returns a boolean if a field has been set.

### GetAssignPublicIp

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetAssignPublicIp() bool`

GetAssignPublicIp returns the AssignPublicIp field if non-nil, zero value otherwise.

### GetAssignPublicIpOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetAssignPublicIpOk() (*bool, bool)`

GetAssignPublicIpOk returns a tuple with the AssignPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignPublicIp

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetAssignPublicIp(v bool)`

SetAssignPublicIp sets AssignPublicIp field to given value.

### HasAssignPublicIp

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasAssignPublicIp() bool`

HasAssignPublicIp returns a boolean if a field has been set.

### GetVisibility

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetStatus

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetStatus() GetAppState200ResponseAllOfInputProvidersInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetStatusOk() (*GetAppState200ResponseAllOfInputProvidersInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetStatus(v GetAppState200ResponseAllOfInputProvidersInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNetwork

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetNetwork() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetNetworkOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetNetwork(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetType

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccount

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetSecurityGroups() []map[string]interface{}`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetSecurityGroupsOk() (*[]map[string]interface{}, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetSecurityGroups(v []map[string]interface{})`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetTenants

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetTenants() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetTenantsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetTenants(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetResourcePermission() GetNetworkSubnets200ResponseAllOfSubnetsInnerResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) GetResourcePermissionOk() (*GetNetworkSubnets200ResponseAllOfSubnetsInnerResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) SetResourcePermission(v GetNetworkSubnets200ResponseAllOfSubnetsInnerResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *GetNetworkSubnets200ResponseAllOfSubnetsInner) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


