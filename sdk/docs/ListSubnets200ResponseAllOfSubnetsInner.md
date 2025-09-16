# ListSubnets200ResponseAllOfSubnetsInner

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
**Pool** | Pointer to **NullableString** |  | [optional] 
**DhcpServer** | Pointer to **bool** |  | [optional] 
**HasFloatingIps** | Pointer to **bool** |  | [optional] 
**DhcpIp** | Pointer to **NullableString** |  | [optional] 
**DnsPrimary** | Pointer to **NullableString** |  | [optional] 
**DnsSecondary** | Pointer to **NullableString** |  | [optional] 
**DhcpStart** | Pointer to **string** |  | [optional] 
**DhcpEnd** | Pointer to **string** |  | [optional] 
**DhcpRange** | Pointer to **NullableString** |  | [optional] 
**NetworkProxy** | Pointer to **NullableString** |  | [optional] 
**NetworkDomain** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**SearchDomains** | Pointer to **NullableString** |  | [optional] 
**DefaultNetwork** | Pointer to **bool** |  | [optional] 
**AssignPublicIp** | Pointer to **bool** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**GetAppState200ResponseAllOfInputProvidersInner**](GetAppState200ResponseAllOfInputProvidersInner.md) |  | [optional] 
**Network** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**SecurityGroups** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tenants** | Pointer to [**[]GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**ResourcePermission** | Pointer to [**GetNetworkSubnets200ResponseAllOfSubnetsInnerResourcePermission**](GetNetworkSubnets200ResponseAllOfSubnetsInnerResourcePermission.md) |  | [optional] 

## Methods

### NewListSubnets200ResponseAllOfSubnetsInner

`func NewListSubnets200ResponseAllOfSubnetsInner() *ListSubnets200ResponseAllOfSubnetsInner`

NewListSubnets200ResponseAllOfSubnetsInner instantiates a new ListSubnets200ResponseAllOfSubnetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSubnets200ResponseAllOfSubnetsInnerWithDefaults

`func NewListSubnets200ResponseAllOfSubnetsInnerWithDefaults() *ListSubnets200ResponseAllOfSubnetsInner`

NewListSubnets200ResponseAllOfSubnetsInnerWithDefaults instantiates a new ListSubnets200ResponseAllOfSubnetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ListSubnets200ResponseAllOfSubnetsInner) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ListSubnets200ResponseAllOfSubnetsInner) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetActive

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDescription

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListSubnets200ResponseAllOfSubnetsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalId

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetUniqueId

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### SetUniqueIdNil

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetUniqueIdNil(b bool)`

 SetUniqueIdNil sets the value for UniqueId to be an explicit nil

### UnsetUniqueId
`func (o *ListSubnets200ResponseAllOfSubnetsInner) UnsetUniqueId()`

UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
### GetAddressPrefix

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetAddressPrefix() string`

GetAddressPrefix returns the AddressPrefix field if non-nil, zero value otherwise.

### GetAddressPrefixOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetAddressPrefixOk() (*string, bool)`

GetAddressPrefixOk returns a tuple with the AddressPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPrefix

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetAddressPrefix(v string)`

SetAddressPrefix sets AddressPrefix field to given value.

### HasAddressPrefix

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasAddressPrefix() bool`

HasAddressPrefix returns a boolean if a field has been set.

### SetAddressPrefixNil

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetAddressPrefixNil(b bool)`

 SetAddressPrefixNil sets the value for AddressPrefix to be an explicit nil

### UnsetAddressPrefix
`func (o *ListSubnets200ResponseAllOfSubnetsInner) UnsetAddressPrefix()`

UnsetAddressPrefix ensures that no value is present for AddressPrefix, not even an explicit nil
### GetCidr

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetGateway

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *ListSubnets200ResponseAllOfSubnetsInner) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetNetmask

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetSubnetAddress

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetSubnetAddress() string`

GetSubnetAddress returns the SubnetAddress field if non-nil, zero value otherwise.

### GetSubnetAddressOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetSubnetAddressOk() (*string, bool)`

GetSubnetAddressOk returns a tuple with the SubnetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAddress

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetSubnetAddress(v string)`

SetSubnetAddress sets SubnetAddress field to given value.

### HasSubnetAddress

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasSubnetAddress() bool`

HasSubnetAddress returns a boolean if a field has been set.

### GetTftpServer

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetTftpServer() string`

GetTftpServer returns the TftpServer field if non-nil, zero value otherwise.

### GetTftpServerOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetTftpServerOk() (*string, bool)`

GetTftpServerOk returns a tuple with the TftpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTftpServer

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetTftpServer(v string)`

SetTftpServer sets TftpServer field to given value.

### HasTftpServer

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasTftpServer() bool`

HasTftpServer returns a boolean if a field has been set.

### SetTftpServerNil

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetTftpServerNil(b bool)`

 SetTftpServerNil sets the value for TftpServer to be an explicit nil

### UnsetTftpServer
`func (o *ListSubnets200ResponseAllOfSubnetsInner) UnsetTftpServer()`

UnsetTftpServer ensures that no value is present for TftpServer, not even an explicit nil
### GetBootFile

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetBootFile() string`

GetBootFile returns the BootFile field if non-nil, zero value otherwise.

### GetBootFileOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetBootFileOk() (*string, bool)`

GetBootFileOk returns a tuple with the BootFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootFile

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetBootFile(v string)`

SetBootFile sets BootFile field to given value.

### HasBootFile

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasBootFile() bool`

HasBootFile returns a boolean if a field has been set.

### SetBootFileNil

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetBootFileNil(b bool)`

 SetBootFileNil sets the value for BootFile to be an explicit nil

### UnsetBootFile
`func (o *ListSubnets200ResponseAllOfSubnetsInner) UnsetBootFile()`

UnsetBootFile ensures that no value is present for BootFile, not even an explicit nil
### GetPool

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasPool() bool`

HasPool returns a boolean if a field has been set.

### SetPoolNil

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *ListSubnets200ResponseAllOfSubnetsInner) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil
### GetDhcpServer

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetDhcpServer() bool`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetDhcpServerOk() (*bool, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetDhcpServer(v bool)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetHasFloatingIps

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetHasFloatingIps() bool`

GetHasFloatingIps returns the HasFloatingIps field if non-nil, zero value otherwise.

### GetHasFloatingIpsOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetHasFloatingIpsOk() (*bool, bool)`

GetHasFloatingIpsOk returns a tuple with the HasFloatingIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFloatingIps

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetHasFloatingIps(v bool)`

SetHasFloatingIps sets HasFloatingIps field to given value.

### HasHasFloatingIps

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasHasFloatingIps() bool`

HasHasFloatingIps returns a boolean if a field has been set.

### GetDhcpIp

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetDhcpIp() string`

GetDhcpIp returns the DhcpIp field if non-nil, zero value otherwise.

### GetDhcpIpOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetDhcpIpOk() (*string, bool)`

GetDhcpIpOk returns a tuple with the DhcpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpIp

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetDhcpIp(v string)`

SetDhcpIp sets DhcpIp field to given value.

### HasDhcpIp

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasDhcpIp() bool`

HasDhcpIp returns a boolean if a field has been set.

### SetDhcpIpNil

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetDhcpIpNil(b bool)`

 SetDhcpIpNil sets the value for DhcpIp to be an explicit nil

### UnsetDhcpIp
`func (o *ListSubnets200ResponseAllOfSubnetsInner) UnsetDhcpIp()`

UnsetDhcpIp ensures that no value is present for DhcpIp, not even an explicit nil
### GetDnsPrimary

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetDnsPrimary() string`

GetDnsPrimary returns the DnsPrimary field if non-nil, zero value otherwise.

### GetDnsPrimaryOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetDnsPrimaryOk() (*string, bool)`

GetDnsPrimaryOk returns a tuple with the DnsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrimary

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetDnsPrimary(v string)`

SetDnsPrimary sets DnsPrimary field to given value.

### HasDnsPrimary

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasDnsPrimary() bool`

HasDnsPrimary returns a boolean if a field has been set.

### SetDnsPrimaryNil

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetDnsPrimaryNil(b bool)`

 SetDnsPrimaryNil sets the value for DnsPrimary to be an explicit nil

### UnsetDnsPrimary
`func (o *ListSubnets200ResponseAllOfSubnetsInner) UnsetDnsPrimary()`

UnsetDnsPrimary ensures that no value is present for DnsPrimary, not even an explicit nil
### GetDnsSecondary

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetDnsSecondary() string`

GetDnsSecondary returns the DnsSecondary field if non-nil, zero value otherwise.

### GetDnsSecondaryOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetDnsSecondaryOk() (*string, bool)`

GetDnsSecondaryOk returns a tuple with the DnsSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSecondary

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetDnsSecondary(v string)`

SetDnsSecondary sets DnsSecondary field to given value.

### HasDnsSecondary

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasDnsSecondary() bool`

HasDnsSecondary returns a boolean if a field has been set.

### SetDnsSecondaryNil

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetDnsSecondaryNil(b bool)`

 SetDnsSecondaryNil sets the value for DnsSecondary to be an explicit nil

### UnsetDnsSecondary
`func (o *ListSubnets200ResponseAllOfSubnetsInner) UnsetDnsSecondary()`

UnsetDnsSecondary ensures that no value is present for DnsSecondary, not even an explicit nil
### GetDhcpStart

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetDhcpStart() string`

GetDhcpStart returns the DhcpStart field if non-nil, zero value otherwise.

### GetDhcpStartOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetDhcpStartOk() (*string, bool)`

GetDhcpStartOk returns a tuple with the DhcpStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpStart

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetDhcpStart(v string)`

SetDhcpStart sets DhcpStart field to given value.

### HasDhcpStart

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasDhcpStart() bool`

HasDhcpStart returns a boolean if a field has been set.

### GetDhcpEnd

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetDhcpEnd() string`

GetDhcpEnd returns the DhcpEnd field if non-nil, zero value otherwise.

### GetDhcpEndOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetDhcpEndOk() (*string, bool)`

GetDhcpEndOk returns a tuple with the DhcpEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpEnd

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetDhcpEnd(v string)`

SetDhcpEnd sets DhcpEnd field to given value.

### HasDhcpEnd

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasDhcpEnd() bool`

HasDhcpEnd returns a boolean if a field has been set.

### GetDhcpRange

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetDhcpRange() string`

GetDhcpRange returns the DhcpRange field if non-nil, zero value otherwise.

### GetDhcpRangeOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetDhcpRangeOk() (*string, bool)`

GetDhcpRangeOk returns a tuple with the DhcpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRange

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetDhcpRange(v string)`

SetDhcpRange sets DhcpRange field to given value.

### HasDhcpRange

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasDhcpRange() bool`

HasDhcpRange returns a boolean if a field has been set.

### SetDhcpRangeNil

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetDhcpRangeNil(b bool)`

 SetDhcpRangeNil sets the value for DhcpRange to be an explicit nil

### UnsetDhcpRange
`func (o *ListSubnets200ResponseAllOfSubnetsInner) UnsetDhcpRange()`

UnsetDhcpRange ensures that no value is present for DhcpRange, not even an explicit nil
### GetNetworkProxy

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetNetworkProxy() string`

GetNetworkProxy returns the NetworkProxy field if non-nil, zero value otherwise.

### GetNetworkProxyOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetNetworkProxyOk() (*string, bool)`

GetNetworkProxyOk returns a tuple with the NetworkProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProxy

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetNetworkProxy(v string)`

SetNetworkProxy sets NetworkProxy field to given value.

### HasNetworkProxy

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasNetworkProxy() bool`

HasNetworkProxy returns a boolean if a field has been set.

### SetNetworkProxyNil

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetNetworkProxyNil(b bool)`

 SetNetworkProxyNil sets the value for NetworkProxy to be an explicit nil

### UnsetNetworkProxy
`func (o *ListSubnets200ResponseAllOfSubnetsInner) UnsetNetworkProxy()`

UnsetNetworkProxy ensures that no value is present for NetworkProxy, not even an explicit nil
### GetNetworkDomain

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetNetworkDomain() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetNetworkDomainOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetNetworkDomain(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetSearchDomains

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetSearchDomains() string`

GetSearchDomains returns the SearchDomains field if non-nil, zero value otherwise.

### GetSearchDomainsOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetSearchDomainsOk() (*string, bool)`

GetSearchDomainsOk returns a tuple with the SearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomains

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetSearchDomains(v string)`

SetSearchDomains sets SearchDomains field to given value.

### HasSearchDomains

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasSearchDomains() bool`

HasSearchDomains returns a boolean if a field has been set.

### SetSearchDomainsNil

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetSearchDomainsNil(b bool)`

 SetSearchDomainsNil sets the value for SearchDomains to be an explicit nil

### UnsetSearchDomains
`func (o *ListSubnets200ResponseAllOfSubnetsInner) UnsetSearchDomains()`

UnsetSearchDomains ensures that no value is present for SearchDomains, not even an explicit nil
### GetDefaultNetwork

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetDefaultNetwork() bool`

GetDefaultNetwork returns the DefaultNetwork field if non-nil, zero value otherwise.

### GetDefaultNetworkOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetDefaultNetworkOk() (*bool, bool)`

GetDefaultNetworkOk returns a tuple with the DefaultNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetwork

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetDefaultNetwork(v bool)`

SetDefaultNetwork sets DefaultNetwork field to given value.

### HasDefaultNetwork

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasDefaultNetwork() bool`

HasDefaultNetwork returns a boolean if a field has been set.

### GetAssignPublicIp

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetAssignPublicIp() bool`

GetAssignPublicIp returns the AssignPublicIp field if non-nil, zero value otherwise.

### GetAssignPublicIpOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetAssignPublicIpOk() (*bool, bool)`

GetAssignPublicIpOk returns a tuple with the AssignPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignPublicIp

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetAssignPublicIp(v bool)`

SetAssignPublicIp sets AssignPublicIp field to given value.

### HasAssignPublicIp

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasAssignPublicIp() bool`

HasAssignPublicIp returns a boolean if a field has been set.

### GetVisibility

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetStatus

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetStatus() GetAppState200ResponseAllOfInputProvidersInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetStatusOk() (*GetAppState200ResponseAllOfInputProvidersInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetStatus(v GetAppState200ResponseAllOfInputProvidersInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNetwork

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetNetwork() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetNetworkOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetNetwork(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetType

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccount

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetSecurityGroups() []map[string]interface{}`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetSecurityGroupsOk() (*[]map[string]interface{}, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetSecurityGroups(v []map[string]interface{})`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetTenants

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetTenants() []GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetTenantsOk() (*[]GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetTenants(v []GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetResourcePermission() GetNetworkSubnets200ResponseAllOfSubnetsInnerResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *ListSubnets200ResponseAllOfSubnetsInner) GetResourcePermissionOk() (*GetNetworkSubnets200ResponseAllOfSubnetsInnerResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *ListSubnets200ResponseAllOfSubnetsInner) SetResourcePermission(v GetNetworkSubnets200ResponseAllOfSubnetsInnerResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *ListSubnets200ResponseAllOfSubnetsInner) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


