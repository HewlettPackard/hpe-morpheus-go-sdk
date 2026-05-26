# CreateSubnet200ResponseSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**UniqueId** | Pointer to **NullableString** |  | [optional] 
**AddressPrefix** | Pointer to **NullableString** |  | [optional] 
**Cidr** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to **NullableString** |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**SubnetAddress** | Pointer to **string** |  | [optional] 
**TftpServer** | Pointer to **NullableString** |  | [optional] 
**BootFile** | Pointer to **NullableString** |  | [optional] 
**Pool** | Pointer to [**CreateSubnet200ResponseSubnetPool**](CreateSubnet200ResponseSubnetPool.md) |  | [optional] 
**DhcpServer** | Pointer to **bool** |  | [optional] 
**HasFloatingIps** | Pointer to **bool** |  | [optional] 
**DhcpIp** | Pointer to **NullableString** |  | [optional] 
**DnsPrimary** | Pointer to **NullableString** |  | [optional] 
**DnsSecondary** | Pointer to **NullableString** |  | [optional] 
**DhcpStart** | Pointer to **string** |  | [optional] 
**DhcpEnd** | Pointer to **string** |  | [optional] 
**DhcpRange** | Pointer to **NullableString** |  | [optional] 
**NetworkProxy** | Pointer to [**CreateSubnet200ResponseSubnetNetworkProxy**](CreateSubnet200ResponseSubnetNetworkProxy.md) |  | [optional] 
**NetworkDomain** | Pointer to [**CreateSubnet200ResponseSubnetNetworkDomain**](CreateSubnet200ResponseSubnetNetworkDomain.md) |  | [optional] 
**SearchDomains** | Pointer to **NullableString** |  | [optional] 
**DefaultNetwork** | Pointer to **bool** |  | [optional] 
**AssignPublicIp** | Pointer to **bool** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**CreateSubnet200ResponseSubnetStatus**](CreateSubnet200ResponseSubnetStatus.md) |  | [optional] 
**Network** | Pointer to [**CreateSubnet200ResponseSubnetNetwork**](CreateSubnet200ResponseSubnetNetwork.md) |  | [optional] 
**Zone** | Pointer to [**CreateSubnet200ResponseSubnetZone**](CreateSubnet200ResponseSubnetZone.md) |  | [optional] 
**Type** | Pointer to [**CreateSubnet200ResponseSubnetType**](CreateSubnet200ResponseSubnetType.md) |  | [optional] 
**Account** | Pointer to [**CreateSubnet200ResponseSubnetAccount**](CreateSubnet200ResponseSubnetAccount.md) |  | [optional] 
**SecurityGroups** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tenants** | Pointer to [**[]CreateSubnet200ResponseSubnetTenantsInner**](CreateSubnet200ResponseSubnetTenantsInner.md) |  | [optional] 
**ResourcePermission** | Pointer to [**CreateSubnet200ResponseSubnetResourcePermission**](CreateSubnet200ResponseSubnetResourcePermission.md) |  | [optional] 

## Methods

### NewCreateSubnet200ResponseSubnet

`func NewCreateSubnet200ResponseSubnet() *CreateSubnet200ResponseSubnet`

NewCreateSubnet200ResponseSubnet instantiates a new CreateSubnet200ResponseSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubnet200ResponseSubnetWithDefaults

`func NewCreateSubnet200ResponseSubnetWithDefaults() *CreateSubnet200ResponseSubnet`

NewCreateSubnet200ResponseSubnetWithDefaults instantiates a new CreateSubnet200ResponseSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateSubnet200ResponseSubnet) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateSubnet200ResponseSubnet) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateSubnet200ResponseSubnet) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateSubnet200ResponseSubnet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *CreateSubnet200ResponseSubnet) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateSubnet200ResponseSubnet) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateSubnet200ResponseSubnet) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateSubnet200ResponseSubnet) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *CreateSubnet200ResponseSubnet) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *CreateSubnet200ResponseSubnet) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *CreateSubnet200ResponseSubnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSubnet200ResponseSubnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSubnet200ResponseSubnet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateSubnet200ResponseSubnet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *CreateSubnet200ResponseSubnet) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateSubnet200ResponseSubnet) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateSubnet200ResponseSubnet) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateSubnet200ResponseSubnet) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *CreateSubnet200ResponseSubnet) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *CreateSubnet200ResponseSubnet) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetActive

`func (o *CreateSubnet200ResponseSubnet) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateSubnet200ResponseSubnet) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateSubnet200ResponseSubnet) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateSubnet200ResponseSubnet) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDescription

`func (o *CreateSubnet200ResponseSubnet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSubnet200ResponseSubnet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSubnet200ResponseSubnet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSubnet200ResponseSubnet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateSubnet200ResponseSubnet) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateSubnet200ResponseSubnet) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalId

`func (o *CreateSubnet200ResponseSubnet) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateSubnet200ResponseSubnet) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateSubnet200ResponseSubnet) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateSubnet200ResponseSubnet) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetUniqueId

`func (o *CreateSubnet200ResponseSubnet) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *CreateSubnet200ResponseSubnet) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *CreateSubnet200ResponseSubnet) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *CreateSubnet200ResponseSubnet) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### SetUniqueIdNil

`func (o *CreateSubnet200ResponseSubnet) SetUniqueIdNil(b bool)`

 SetUniqueIdNil sets the value for UniqueId to be an explicit nil

### UnsetUniqueId
`func (o *CreateSubnet200ResponseSubnet) UnsetUniqueId()`

UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
### GetAddressPrefix

`func (o *CreateSubnet200ResponseSubnet) GetAddressPrefix() string`

GetAddressPrefix returns the AddressPrefix field if non-nil, zero value otherwise.

### GetAddressPrefixOk

`func (o *CreateSubnet200ResponseSubnet) GetAddressPrefixOk() (*string, bool)`

GetAddressPrefixOk returns a tuple with the AddressPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPrefix

`func (o *CreateSubnet200ResponseSubnet) SetAddressPrefix(v string)`

SetAddressPrefix sets AddressPrefix field to given value.

### HasAddressPrefix

`func (o *CreateSubnet200ResponseSubnet) HasAddressPrefix() bool`

HasAddressPrefix returns a boolean if a field has been set.

### SetAddressPrefixNil

`func (o *CreateSubnet200ResponseSubnet) SetAddressPrefixNil(b bool)`

 SetAddressPrefixNil sets the value for AddressPrefix to be an explicit nil

### UnsetAddressPrefix
`func (o *CreateSubnet200ResponseSubnet) UnsetAddressPrefix()`

UnsetAddressPrefix ensures that no value is present for AddressPrefix, not even an explicit nil
### GetCidr

`func (o *CreateSubnet200ResponseSubnet) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CreateSubnet200ResponseSubnet) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CreateSubnet200ResponseSubnet) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *CreateSubnet200ResponseSubnet) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetGateway

`func (o *CreateSubnet200ResponseSubnet) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *CreateSubnet200ResponseSubnet) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *CreateSubnet200ResponseSubnet) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *CreateSubnet200ResponseSubnet) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *CreateSubnet200ResponseSubnet) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *CreateSubnet200ResponseSubnet) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetNetmask

`func (o *CreateSubnet200ResponseSubnet) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *CreateSubnet200ResponseSubnet) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *CreateSubnet200ResponseSubnet) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *CreateSubnet200ResponseSubnet) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetSubnetAddress

`func (o *CreateSubnet200ResponseSubnet) GetSubnetAddress() string`

GetSubnetAddress returns the SubnetAddress field if non-nil, zero value otherwise.

### GetSubnetAddressOk

`func (o *CreateSubnet200ResponseSubnet) GetSubnetAddressOk() (*string, bool)`

GetSubnetAddressOk returns a tuple with the SubnetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAddress

`func (o *CreateSubnet200ResponseSubnet) SetSubnetAddress(v string)`

SetSubnetAddress sets SubnetAddress field to given value.

### HasSubnetAddress

`func (o *CreateSubnet200ResponseSubnet) HasSubnetAddress() bool`

HasSubnetAddress returns a boolean if a field has been set.

### GetTftpServer

`func (o *CreateSubnet200ResponseSubnet) GetTftpServer() string`

GetTftpServer returns the TftpServer field if non-nil, zero value otherwise.

### GetTftpServerOk

`func (o *CreateSubnet200ResponseSubnet) GetTftpServerOk() (*string, bool)`

GetTftpServerOk returns a tuple with the TftpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTftpServer

`func (o *CreateSubnet200ResponseSubnet) SetTftpServer(v string)`

SetTftpServer sets TftpServer field to given value.

### HasTftpServer

`func (o *CreateSubnet200ResponseSubnet) HasTftpServer() bool`

HasTftpServer returns a boolean if a field has been set.

### SetTftpServerNil

`func (o *CreateSubnet200ResponseSubnet) SetTftpServerNil(b bool)`

 SetTftpServerNil sets the value for TftpServer to be an explicit nil

### UnsetTftpServer
`func (o *CreateSubnet200ResponseSubnet) UnsetTftpServer()`

UnsetTftpServer ensures that no value is present for TftpServer, not even an explicit nil
### GetBootFile

`func (o *CreateSubnet200ResponseSubnet) GetBootFile() string`

GetBootFile returns the BootFile field if non-nil, zero value otherwise.

### GetBootFileOk

`func (o *CreateSubnet200ResponseSubnet) GetBootFileOk() (*string, bool)`

GetBootFileOk returns a tuple with the BootFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootFile

`func (o *CreateSubnet200ResponseSubnet) SetBootFile(v string)`

SetBootFile sets BootFile field to given value.

### HasBootFile

`func (o *CreateSubnet200ResponseSubnet) HasBootFile() bool`

HasBootFile returns a boolean if a field has been set.

### SetBootFileNil

`func (o *CreateSubnet200ResponseSubnet) SetBootFileNil(b bool)`

 SetBootFileNil sets the value for BootFile to be an explicit nil

### UnsetBootFile
`func (o *CreateSubnet200ResponseSubnet) UnsetBootFile()`

UnsetBootFile ensures that no value is present for BootFile, not even an explicit nil
### GetPool

`func (o *CreateSubnet200ResponseSubnet) GetPool() CreateSubnet200ResponseSubnetPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *CreateSubnet200ResponseSubnet) GetPoolOk() (*CreateSubnet200ResponseSubnetPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *CreateSubnet200ResponseSubnet) SetPool(v CreateSubnet200ResponseSubnetPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *CreateSubnet200ResponseSubnet) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetDhcpServer

`func (o *CreateSubnet200ResponseSubnet) GetDhcpServer() bool`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *CreateSubnet200ResponseSubnet) GetDhcpServerOk() (*bool, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *CreateSubnet200ResponseSubnet) SetDhcpServer(v bool)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *CreateSubnet200ResponseSubnet) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetHasFloatingIps

`func (o *CreateSubnet200ResponseSubnet) GetHasFloatingIps() bool`

GetHasFloatingIps returns the HasFloatingIps field if non-nil, zero value otherwise.

### GetHasFloatingIpsOk

`func (o *CreateSubnet200ResponseSubnet) GetHasFloatingIpsOk() (*bool, bool)`

GetHasFloatingIpsOk returns a tuple with the HasFloatingIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFloatingIps

`func (o *CreateSubnet200ResponseSubnet) SetHasFloatingIps(v bool)`

SetHasFloatingIps sets HasFloatingIps field to given value.

### HasHasFloatingIps

`func (o *CreateSubnet200ResponseSubnet) HasHasFloatingIps() bool`

HasHasFloatingIps returns a boolean if a field has been set.

### GetDhcpIp

`func (o *CreateSubnet200ResponseSubnet) GetDhcpIp() string`

GetDhcpIp returns the DhcpIp field if non-nil, zero value otherwise.

### GetDhcpIpOk

`func (o *CreateSubnet200ResponseSubnet) GetDhcpIpOk() (*string, bool)`

GetDhcpIpOk returns a tuple with the DhcpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpIp

`func (o *CreateSubnet200ResponseSubnet) SetDhcpIp(v string)`

SetDhcpIp sets DhcpIp field to given value.

### HasDhcpIp

`func (o *CreateSubnet200ResponseSubnet) HasDhcpIp() bool`

HasDhcpIp returns a boolean if a field has been set.

### SetDhcpIpNil

`func (o *CreateSubnet200ResponseSubnet) SetDhcpIpNil(b bool)`

 SetDhcpIpNil sets the value for DhcpIp to be an explicit nil

### UnsetDhcpIp
`func (o *CreateSubnet200ResponseSubnet) UnsetDhcpIp()`

UnsetDhcpIp ensures that no value is present for DhcpIp, not even an explicit nil
### GetDnsPrimary

`func (o *CreateSubnet200ResponseSubnet) GetDnsPrimary() string`

GetDnsPrimary returns the DnsPrimary field if non-nil, zero value otherwise.

### GetDnsPrimaryOk

`func (o *CreateSubnet200ResponseSubnet) GetDnsPrimaryOk() (*string, bool)`

GetDnsPrimaryOk returns a tuple with the DnsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrimary

`func (o *CreateSubnet200ResponseSubnet) SetDnsPrimary(v string)`

SetDnsPrimary sets DnsPrimary field to given value.

### HasDnsPrimary

`func (o *CreateSubnet200ResponseSubnet) HasDnsPrimary() bool`

HasDnsPrimary returns a boolean if a field has been set.

### SetDnsPrimaryNil

`func (o *CreateSubnet200ResponseSubnet) SetDnsPrimaryNil(b bool)`

 SetDnsPrimaryNil sets the value for DnsPrimary to be an explicit nil

### UnsetDnsPrimary
`func (o *CreateSubnet200ResponseSubnet) UnsetDnsPrimary()`

UnsetDnsPrimary ensures that no value is present for DnsPrimary, not even an explicit nil
### GetDnsSecondary

`func (o *CreateSubnet200ResponseSubnet) GetDnsSecondary() string`

GetDnsSecondary returns the DnsSecondary field if non-nil, zero value otherwise.

### GetDnsSecondaryOk

`func (o *CreateSubnet200ResponseSubnet) GetDnsSecondaryOk() (*string, bool)`

GetDnsSecondaryOk returns a tuple with the DnsSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSecondary

`func (o *CreateSubnet200ResponseSubnet) SetDnsSecondary(v string)`

SetDnsSecondary sets DnsSecondary field to given value.

### HasDnsSecondary

`func (o *CreateSubnet200ResponseSubnet) HasDnsSecondary() bool`

HasDnsSecondary returns a boolean if a field has been set.

### SetDnsSecondaryNil

`func (o *CreateSubnet200ResponseSubnet) SetDnsSecondaryNil(b bool)`

 SetDnsSecondaryNil sets the value for DnsSecondary to be an explicit nil

### UnsetDnsSecondary
`func (o *CreateSubnet200ResponseSubnet) UnsetDnsSecondary()`

UnsetDnsSecondary ensures that no value is present for DnsSecondary, not even an explicit nil
### GetDhcpStart

`func (o *CreateSubnet200ResponseSubnet) GetDhcpStart() string`

GetDhcpStart returns the DhcpStart field if non-nil, zero value otherwise.

### GetDhcpStartOk

`func (o *CreateSubnet200ResponseSubnet) GetDhcpStartOk() (*string, bool)`

GetDhcpStartOk returns a tuple with the DhcpStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpStart

`func (o *CreateSubnet200ResponseSubnet) SetDhcpStart(v string)`

SetDhcpStart sets DhcpStart field to given value.

### HasDhcpStart

`func (o *CreateSubnet200ResponseSubnet) HasDhcpStart() bool`

HasDhcpStart returns a boolean if a field has been set.

### GetDhcpEnd

`func (o *CreateSubnet200ResponseSubnet) GetDhcpEnd() string`

GetDhcpEnd returns the DhcpEnd field if non-nil, zero value otherwise.

### GetDhcpEndOk

`func (o *CreateSubnet200ResponseSubnet) GetDhcpEndOk() (*string, bool)`

GetDhcpEndOk returns a tuple with the DhcpEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpEnd

`func (o *CreateSubnet200ResponseSubnet) SetDhcpEnd(v string)`

SetDhcpEnd sets DhcpEnd field to given value.

### HasDhcpEnd

`func (o *CreateSubnet200ResponseSubnet) HasDhcpEnd() bool`

HasDhcpEnd returns a boolean if a field has been set.

### GetDhcpRange

`func (o *CreateSubnet200ResponseSubnet) GetDhcpRange() string`

GetDhcpRange returns the DhcpRange field if non-nil, zero value otherwise.

### GetDhcpRangeOk

`func (o *CreateSubnet200ResponseSubnet) GetDhcpRangeOk() (*string, bool)`

GetDhcpRangeOk returns a tuple with the DhcpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRange

`func (o *CreateSubnet200ResponseSubnet) SetDhcpRange(v string)`

SetDhcpRange sets DhcpRange field to given value.

### HasDhcpRange

`func (o *CreateSubnet200ResponseSubnet) HasDhcpRange() bool`

HasDhcpRange returns a boolean if a field has been set.

### SetDhcpRangeNil

`func (o *CreateSubnet200ResponseSubnet) SetDhcpRangeNil(b bool)`

 SetDhcpRangeNil sets the value for DhcpRange to be an explicit nil

### UnsetDhcpRange
`func (o *CreateSubnet200ResponseSubnet) UnsetDhcpRange()`

UnsetDhcpRange ensures that no value is present for DhcpRange, not even an explicit nil
### GetNetworkProxy

`func (o *CreateSubnet200ResponseSubnet) GetNetworkProxy() CreateSubnet200ResponseSubnetNetworkProxy`

GetNetworkProxy returns the NetworkProxy field if non-nil, zero value otherwise.

### GetNetworkProxyOk

`func (o *CreateSubnet200ResponseSubnet) GetNetworkProxyOk() (*CreateSubnet200ResponseSubnetNetworkProxy, bool)`

GetNetworkProxyOk returns a tuple with the NetworkProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProxy

`func (o *CreateSubnet200ResponseSubnet) SetNetworkProxy(v CreateSubnet200ResponseSubnetNetworkProxy)`

SetNetworkProxy sets NetworkProxy field to given value.

### HasNetworkProxy

`func (o *CreateSubnet200ResponseSubnet) HasNetworkProxy() bool`

HasNetworkProxy returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *CreateSubnet200ResponseSubnet) GetNetworkDomain() CreateSubnet200ResponseSubnetNetworkDomain`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *CreateSubnet200ResponseSubnet) GetNetworkDomainOk() (*CreateSubnet200ResponseSubnetNetworkDomain, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *CreateSubnet200ResponseSubnet) SetNetworkDomain(v CreateSubnet200ResponseSubnetNetworkDomain)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *CreateSubnet200ResponseSubnet) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetSearchDomains

`func (o *CreateSubnet200ResponseSubnet) GetSearchDomains() string`

GetSearchDomains returns the SearchDomains field if non-nil, zero value otherwise.

### GetSearchDomainsOk

`func (o *CreateSubnet200ResponseSubnet) GetSearchDomainsOk() (*string, bool)`

GetSearchDomainsOk returns a tuple with the SearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomains

`func (o *CreateSubnet200ResponseSubnet) SetSearchDomains(v string)`

SetSearchDomains sets SearchDomains field to given value.

### HasSearchDomains

`func (o *CreateSubnet200ResponseSubnet) HasSearchDomains() bool`

HasSearchDomains returns a boolean if a field has been set.

### SetSearchDomainsNil

`func (o *CreateSubnet200ResponseSubnet) SetSearchDomainsNil(b bool)`

 SetSearchDomainsNil sets the value for SearchDomains to be an explicit nil

### UnsetSearchDomains
`func (o *CreateSubnet200ResponseSubnet) UnsetSearchDomains()`

UnsetSearchDomains ensures that no value is present for SearchDomains, not even an explicit nil
### GetDefaultNetwork

`func (o *CreateSubnet200ResponseSubnet) GetDefaultNetwork() bool`

GetDefaultNetwork returns the DefaultNetwork field if non-nil, zero value otherwise.

### GetDefaultNetworkOk

`func (o *CreateSubnet200ResponseSubnet) GetDefaultNetworkOk() (*bool, bool)`

GetDefaultNetworkOk returns a tuple with the DefaultNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetwork

`func (o *CreateSubnet200ResponseSubnet) SetDefaultNetwork(v bool)`

SetDefaultNetwork sets DefaultNetwork field to given value.

### HasDefaultNetwork

`func (o *CreateSubnet200ResponseSubnet) HasDefaultNetwork() bool`

HasDefaultNetwork returns a boolean if a field has been set.

### GetAssignPublicIp

`func (o *CreateSubnet200ResponseSubnet) GetAssignPublicIp() bool`

GetAssignPublicIp returns the AssignPublicIp field if non-nil, zero value otherwise.

### GetAssignPublicIpOk

`func (o *CreateSubnet200ResponseSubnet) GetAssignPublicIpOk() (*bool, bool)`

GetAssignPublicIpOk returns a tuple with the AssignPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignPublicIp

`func (o *CreateSubnet200ResponseSubnet) SetAssignPublicIp(v bool)`

SetAssignPublicIp sets AssignPublicIp field to given value.

### HasAssignPublicIp

`func (o *CreateSubnet200ResponseSubnet) HasAssignPublicIp() bool`

HasAssignPublicIp returns a boolean if a field has been set.

### GetVisibility

`func (o *CreateSubnet200ResponseSubnet) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateSubnet200ResponseSubnet) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateSubnet200ResponseSubnet) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateSubnet200ResponseSubnet) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetStatus

`func (o *CreateSubnet200ResponseSubnet) GetStatus() CreateSubnet200ResponseSubnetStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateSubnet200ResponseSubnet) GetStatusOk() (*CreateSubnet200ResponseSubnetStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateSubnet200ResponseSubnet) SetStatus(v CreateSubnet200ResponseSubnetStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateSubnet200ResponseSubnet) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNetwork

`func (o *CreateSubnet200ResponseSubnet) GetNetwork() CreateSubnet200ResponseSubnetNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CreateSubnet200ResponseSubnet) GetNetworkOk() (*CreateSubnet200ResponseSubnetNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CreateSubnet200ResponseSubnet) SetNetwork(v CreateSubnet200ResponseSubnetNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CreateSubnet200ResponseSubnet) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetZone

`func (o *CreateSubnet200ResponseSubnet) GetZone() CreateSubnet200ResponseSubnetZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CreateSubnet200ResponseSubnet) GetZoneOk() (*CreateSubnet200ResponseSubnetZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CreateSubnet200ResponseSubnet) SetZone(v CreateSubnet200ResponseSubnetZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CreateSubnet200ResponseSubnet) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetType

`func (o *CreateSubnet200ResponseSubnet) GetType() CreateSubnet200ResponseSubnetType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateSubnet200ResponseSubnet) GetTypeOk() (*CreateSubnet200ResponseSubnetType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateSubnet200ResponseSubnet) SetType(v CreateSubnet200ResponseSubnetType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateSubnet200ResponseSubnet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccount

`func (o *CreateSubnet200ResponseSubnet) GetAccount() CreateSubnet200ResponseSubnetAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreateSubnet200ResponseSubnet) GetAccountOk() (*CreateSubnet200ResponseSubnetAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreateSubnet200ResponseSubnet) SetAccount(v CreateSubnet200ResponseSubnetAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CreateSubnet200ResponseSubnet) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *CreateSubnet200ResponseSubnet) GetSecurityGroups() []map[string]interface{}`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *CreateSubnet200ResponseSubnet) GetSecurityGroupsOk() (*[]map[string]interface{}, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *CreateSubnet200ResponseSubnet) SetSecurityGroups(v []map[string]interface{})`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *CreateSubnet200ResponseSubnet) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetTenants

`func (o *CreateSubnet200ResponseSubnet) GetTenants() []CreateSubnet200ResponseSubnetTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *CreateSubnet200ResponseSubnet) GetTenantsOk() (*[]CreateSubnet200ResponseSubnetTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *CreateSubnet200ResponseSubnet) SetTenants(v []CreateSubnet200ResponseSubnetTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *CreateSubnet200ResponseSubnet) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *CreateSubnet200ResponseSubnet) GetResourcePermission() CreateSubnet200ResponseSubnetResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *CreateSubnet200ResponseSubnet) GetResourcePermissionOk() (*CreateSubnet200ResponseSubnetResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *CreateSubnet200ResponseSubnet) SetResourcePermission(v CreateSubnet200ResponseSubnetResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *CreateSubnet200ResponseSubnet) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


