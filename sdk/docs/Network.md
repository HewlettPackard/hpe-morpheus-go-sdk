# Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Network ID | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**DisplayName** | Pointer to **string** | Network Display Name | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Zone** | Pointer to [**ListNetworks200ResponseAllOfNetworksInnerZone**](ListNetworks200ResponseAllOfNetworksInnerZone.md) |  | [optional] 
**Type** | Pointer to [**ListNetworks200ResponseAllOfNetworksInnerType**](ListNetworks200ResponseAllOfNetworksInnerType.md) |  | [optional] 
**Owner** | Pointer to [**ListNetworks200ResponseAllOfNetworksInnerOwner**](ListNetworks200ResponseAllOfNetworksInnerOwner.md) |  | [optional] 
**Code** | Pointer to **NullableString** | Network Code | [optional] 
**Ipv4Enabled** | Pointer to **bool** |  | [optional] 
**Ipv6Enabled** | Pointer to **bool** |  | [optional] 
**Category** | Pointer to **NullableString** | Network Category | [optional] 
**InterfaceName** | Pointer to **NullableString** |  | [optional] 
**BridgeName** | Pointer to **NullableString** |  | [optional] 
**BridgeInterface** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** | Description | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**UniqueId** | Pointer to **NullableString** |  | [optional] 
**ExternalType** | Pointer to **string** |  | [optional] 
**RefUrl** | Pointer to **NullableString** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**VlanId** | Pointer to **NullableInt64** |  | [optional] 
**VswitchName** | Pointer to **NullableString** |  | [optional] 
**DhcpServer** | Pointer to **bool** |  | [optional] 
**DhcpIp** | Pointer to **NullableString** |  | [optional] 
**DhcpServerIPv6** | Pointer to **bool** |  | [optional] 
**Gateway** | Pointer to **NullableString** | Network Gateway | [optional] 
**Netmask** | Pointer to **NullableString** |  | [optional] 
**Broadcast** | Pointer to **NullableString** |  | [optional] 
**SubnetAddress** | Pointer to **NullableString** |  | [optional] 
**DnsPrimary** | Pointer to **NullableString** | Primary DNS Server | [optional] 
**DnsSecondary** | Pointer to **NullableString** | Secondary DNS Server | [optional] 
**Cidr** | Pointer to **string** | Network CIDR | [optional] 
**GatewayIPv6** | Pointer to **NullableString** | IPv6 Network Gateway | [optional] 
**NetmaskIPv6** | Pointer to **NullableString** |  | [optional] 
**DnsPrimaryIPv6** | Pointer to **NullableString** | Primary IPv6 DNS Server | [optional] 
**DnsSecondaryIPv6** | Pointer to **NullableString** | Secondary IPv6 DNS Server | [optional] 
**CidrIPv6** | Pointer to **NullableString** | IPv6 Network CIDR | [optional] 
**TftpServer** | Pointer to **NullableString** |  | [optional] 
**BootFile** | Pointer to **NullableString** |  | [optional] 
**SwitchId** | Pointer to **NullableString** |  | [optional] 
**FabricId** | Pointer to **NullableString** |  | [optional] 
**NetworkRole** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AvailabilityZone** | Pointer to **NullableString** |  | [optional] 
**Pool** | Pointer to **map[string]interface{}** |  | [optional] 
**PoolIPv6** | Pointer to **map[string]interface{}** |  | [optional] 
**NetworkProxy** | Pointer to [**ListNetworks200ResponseAllOfNetworksInnerNetworkProxy**](ListNetworks200ResponseAllOfNetworksInnerNetworkProxy.md) |  | [optional] 
**NetworkDomain** | Pointer to [**ListNetworks200ResponseAllOfNetworksInnerNetworkDomain**](ListNetworks200ResponseAllOfNetworksInnerNetworkDomain.md) |  | [optional] 
**SearchDomains** | Pointer to **NullableString** |  | [optional] 
**PrefixLength** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**EnableAdmin** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**DefaultNetwork** | Pointer to **bool** |  | [optional] 
**AssignPublicIp** | Pointer to **bool** |  | [optional] 
**NoProxy** | Pointer to **NullableString** |  | [optional] 
**ApplianceUrlProxyBypass** | Pointer to **bool** |  | [optional] 
**ZonePool** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**AllowStaticOverride** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**ListNetworks200ResponseAllOfNetworksInnerConfig**](ListNetworks200ResponseAllOfNetworksInnerConfig.md) |  | [optional] 
**Tenants** | Pointer to [**[]GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 

## Methods

### NewNetwork

`func NewNetwork() *Network`

NewNetwork instantiates a new Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkWithDefaults

`func NewNetworkWithDefaults() *Network`

NewNetworkWithDefaults instantiates a new Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Network) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Network) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Network) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Network) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Network) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Network) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Network) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Network) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *Network) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Network) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Network) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Network) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLabels

`func (o *Network) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Network) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Network) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Network) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *Network) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *Network) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetZone

`func (o *Network) GetZone() ListNetworks200ResponseAllOfNetworksInnerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Network) GetZoneOk() (*ListNetworks200ResponseAllOfNetworksInnerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Network) SetZone(v ListNetworks200ResponseAllOfNetworksInnerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Network) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetType

`func (o *Network) GetType() ListNetworks200ResponseAllOfNetworksInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Network) GetTypeOk() (*ListNetworks200ResponseAllOfNetworksInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Network) SetType(v ListNetworks200ResponseAllOfNetworksInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *Network) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOwner

`func (o *Network) GetOwner() ListNetworks200ResponseAllOfNetworksInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Network) GetOwnerOk() (*ListNetworks200ResponseAllOfNetworksInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Network) SetOwner(v ListNetworks200ResponseAllOfNetworksInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Network) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCode

`func (o *Network) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Network) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Network) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Network) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *Network) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *Network) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetIpv4Enabled

`func (o *Network) GetIpv4Enabled() bool`

GetIpv4Enabled returns the Ipv4Enabled field if non-nil, zero value otherwise.

### GetIpv4EnabledOk

`func (o *Network) GetIpv4EnabledOk() (*bool, bool)`

GetIpv4EnabledOk returns a tuple with the Ipv4Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Enabled

`func (o *Network) SetIpv4Enabled(v bool)`

SetIpv4Enabled sets Ipv4Enabled field to given value.

### HasIpv4Enabled

`func (o *Network) HasIpv4Enabled() bool`

HasIpv4Enabled returns a boolean if a field has been set.

### GetIpv6Enabled

`func (o *Network) GetIpv6Enabled() bool`

GetIpv6Enabled returns the Ipv6Enabled field if non-nil, zero value otherwise.

### GetIpv6EnabledOk

`func (o *Network) GetIpv6EnabledOk() (*bool, bool)`

GetIpv6EnabledOk returns a tuple with the Ipv6Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Enabled

`func (o *Network) SetIpv6Enabled(v bool)`

SetIpv6Enabled sets Ipv6Enabled field to given value.

### HasIpv6Enabled

`func (o *Network) HasIpv6Enabled() bool`

HasIpv6Enabled returns a boolean if a field has been set.

### GetCategory

`func (o *Network) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Network) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Network) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Network) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *Network) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *Network) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetInterfaceName

`func (o *Network) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *Network) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *Network) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *Network) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### SetInterfaceNameNil

`func (o *Network) SetInterfaceNameNil(b bool)`

 SetInterfaceNameNil sets the value for InterfaceName to be an explicit nil

### UnsetInterfaceName
`func (o *Network) UnsetInterfaceName()`

UnsetInterfaceName ensures that no value is present for InterfaceName, not even an explicit nil
### GetBridgeName

`func (o *Network) GetBridgeName() string`

GetBridgeName returns the BridgeName field if non-nil, zero value otherwise.

### GetBridgeNameOk

`func (o *Network) GetBridgeNameOk() (*string, bool)`

GetBridgeNameOk returns a tuple with the BridgeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeName

`func (o *Network) SetBridgeName(v string)`

SetBridgeName sets BridgeName field to given value.

### HasBridgeName

`func (o *Network) HasBridgeName() bool`

HasBridgeName returns a boolean if a field has been set.

### SetBridgeNameNil

`func (o *Network) SetBridgeNameNil(b bool)`

 SetBridgeNameNil sets the value for BridgeName to be an explicit nil

### UnsetBridgeName
`func (o *Network) UnsetBridgeName()`

UnsetBridgeName ensures that no value is present for BridgeName, not even an explicit nil
### GetBridgeInterface

`func (o *Network) GetBridgeInterface() string`

GetBridgeInterface returns the BridgeInterface field if non-nil, zero value otherwise.

### GetBridgeInterfaceOk

`func (o *Network) GetBridgeInterfaceOk() (*string, bool)`

GetBridgeInterfaceOk returns a tuple with the BridgeInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeInterface

`func (o *Network) SetBridgeInterface(v string)`

SetBridgeInterface sets BridgeInterface field to given value.

### HasBridgeInterface

`func (o *Network) HasBridgeInterface() bool`

HasBridgeInterface returns a boolean if a field has been set.

### SetBridgeInterfaceNil

`func (o *Network) SetBridgeInterfaceNil(b bool)`

 SetBridgeInterfaceNil sets the value for BridgeInterface to be an explicit nil

### UnsetBridgeInterface
`func (o *Network) UnsetBridgeInterface()`

UnsetBridgeInterface ensures that no value is present for BridgeInterface, not even an explicit nil
### GetDescription

`func (o *Network) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Network) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Network) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Network) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Network) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Network) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalId

`func (o *Network) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Network) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Network) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Network) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *Network) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *Network) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInternalId

`func (o *Network) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *Network) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *Network) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *Network) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *Network) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *Network) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetUniqueId

`func (o *Network) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *Network) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *Network) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *Network) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### SetUniqueIdNil

`func (o *Network) SetUniqueIdNil(b bool)`

 SetUniqueIdNil sets the value for UniqueId to be an explicit nil

### UnsetUniqueId
`func (o *Network) UnsetUniqueId()`

UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
### GetExternalType

`func (o *Network) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *Network) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *Network) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *Network) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### GetRefUrl

`func (o *Network) GetRefUrl() string`

GetRefUrl returns the RefUrl field if non-nil, zero value otherwise.

### GetRefUrlOk

`func (o *Network) GetRefUrlOk() (*string, bool)`

GetRefUrlOk returns a tuple with the RefUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUrl

`func (o *Network) SetRefUrl(v string)`

SetRefUrl sets RefUrl field to given value.

### HasRefUrl

`func (o *Network) HasRefUrl() bool`

HasRefUrl returns a boolean if a field has been set.

### SetRefUrlNil

`func (o *Network) SetRefUrlNil(b bool)`

 SetRefUrlNil sets the value for RefUrl to be an explicit nil

### UnsetRefUrl
`func (o *Network) UnsetRefUrl()`

UnsetRefUrl ensures that no value is present for RefUrl, not even an explicit nil
### GetRefType

`func (o *Network) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *Network) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *Network) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *Network) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *Network) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *Network) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *Network) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *Network) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetVlanId

`func (o *Network) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *Network) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *Network) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *Network) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### SetVlanIdNil

`func (o *Network) SetVlanIdNil(b bool)`

 SetVlanIdNil sets the value for VlanId to be an explicit nil

### UnsetVlanId
`func (o *Network) UnsetVlanId()`

UnsetVlanId ensures that no value is present for VlanId, not even an explicit nil
### GetVswitchName

`func (o *Network) GetVswitchName() string`

GetVswitchName returns the VswitchName field if non-nil, zero value otherwise.

### GetVswitchNameOk

`func (o *Network) GetVswitchNameOk() (*string, bool)`

GetVswitchNameOk returns a tuple with the VswitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchName

`func (o *Network) SetVswitchName(v string)`

SetVswitchName sets VswitchName field to given value.

### HasVswitchName

`func (o *Network) HasVswitchName() bool`

HasVswitchName returns a boolean if a field has been set.

### SetVswitchNameNil

`func (o *Network) SetVswitchNameNil(b bool)`

 SetVswitchNameNil sets the value for VswitchName to be an explicit nil

### UnsetVswitchName
`func (o *Network) UnsetVswitchName()`

UnsetVswitchName ensures that no value is present for VswitchName, not even an explicit nil
### GetDhcpServer

`func (o *Network) GetDhcpServer() bool`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *Network) GetDhcpServerOk() (*bool, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *Network) SetDhcpServer(v bool)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *Network) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetDhcpIp

`func (o *Network) GetDhcpIp() string`

GetDhcpIp returns the DhcpIp field if non-nil, zero value otherwise.

### GetDhcpIpOk

`func (o *Network) GetDhcpIpOk() (*string, bool)`

GetDhcpIpOk returns a tuple with the DhcpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpIp

`func (o *Network) SetDhcpIp(v string)`

SetDhcpIp sets DhcpIp field to given value.

### HasDhcpIp

`func (o *Network) HasDhcpIp() bool`

HasDhcpIp returns a boolean if a field has been set.

### SetDhcpIpNil

`func (o *Network) SetDhcpIpNil(b bool)`

 SetDhcpIpNil sets the value for DhcpIp to be an explicit nil

### UnsetDhcpIp
`func (o *Network) UnsetDhcpIp()`

UnsetDhcpIp ensures that no value is present for DhcpIp, not even an explicit nil
### GetDhcpServerIPv6

`func (o *Network) GetDhcpServerIPv6() bool`

GetDhcpServerIPv6 returns the DhcpServerIPv6 field if non-nil, zero value otherwise.

### GetDhcpServerIPv6Ok

`func (o *Network) GetDhcpServerIPv6Ok() (*bool, bool)`

GetDhcpServerIPv6Ok returns a tuple with the DhcpServerIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerIPv6

`func (o *Network) SetDhcpServerIPv6(v bool)`

SetDhcpServerIPv6 sets DhcpServerIPv6 field to given value.

### HasDhcpServerIPv6

`func (o *Network) HasDhcpServerIPv6() bool`

HasDhcpServerIPv6 returns a boolean if a field has been set.

### GetGateway

`func (o *Network) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Network) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Network) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *Network) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *Network) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *Network) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetNetmask

`func (o *Network) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *Network) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *Network) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *Network) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### SetNetmaskNil

`func (o *Network) SetNetmaskNil(b bool)`

 SetNetmaskNil sets the value for Netmask to be an explicit nil

### UnsetNetmask
`func (o *Network) UnsetNetmask()`

UnsetNetmask ensures that no value is present for Netmask, not even an explicit nil
### GetBroadcast

`func (o *Network) GetBroadcast() string`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *Network) GetBroadcastOk() (*string, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *Network) SetBroadcast(v string)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *Network) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.

### SetBroadcastNil

`func (o *Network) SetBroadcastNil(b bool)`

 SetBroadcastNil sets the value for Broadcast to be an explicit nil

### UnsetBroadcast
`func (o *Network) UnsetBroadcast()`

UnsetBroadcast ensures that no value is present for Broadcast, not even an explicit nil
### GetSubnetAddress

`func (o *Network) GetSubnetAddress() string`

GetSubnetAddress returns the SubnetAddress field if non-nil, zero value otherwise.

### GetSubnetAddressOk

`func (o *Network) GetSubnetAddressOk() (*string, bool)`

GetSubnetAddressOk returns a tuple with the SubnetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAddress

`func (o *Network) SetSubnetAddress(v string)`

SetSubnetAddress sets SubnetAddress field to given value.

### HasSubnetAddress

`func (o *Network) HasSubnetAddress() bool`

HasSubnetAddress returns a boolean if a field has been set.

### SetSubnetAddressNil

`func (o *Network) SetSubnetAddressNil(b bool)`

 SetSubnetAddressNil sets the value for SubnetAddress to be an explicit nil

### UnsetSubnetAddress
`func (o *Network) UnsetSubnetAddress()`

UnsetSubnetAddress ensures that no value is present for SubnetAddress, not even an explicit nil
### GetDnsPrimary

`func (o *Network) GetDnsPrimary() string`

GetDnsPrimary returns the DnsPrimary field if non-nil, zero value otherwise.

### GetDnsPrimaryOk

`func (o *Network) GetDnsPrimaryOk() (*string, bool)`

GetDnsPrimaryOk returns a tuple with the DnsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrimary

`func (o *Network) SetDnsPrimary(v string)`

SetDnsPrimary sets DnsPrimary field to given value.

### HasDnsPrimary

`func (o *Network) HasDnsPrimary() bool`

HasDnsPrimary returns a boolean if a field has been set.

### SetDnsPrimaryNil

`func (o *Network) SetDnsPrimaryNil(b bool)`

 SetDnsPrimaryNil sets the value for DnsPrimary to be an explicit nil

### UnsetDnsPrimary
`func (o *Network) UnsetDnsPrimary()`

UnsetDnsPrimary ensures that no value is present for DnsPrimary, not even an explicit nil
### GetDnsSecondary

`func (o *Network) GetDnsSecondary() string`

GetDnsSecondary returns the DnsSecondary field if non-nil, zero value otherwise.

### GetDnsSecondaryOk

`func (o *Network) GetDnsSecondaryOk() (*string, bool)`

GetDnsSecondaryOk returns a tuple with the DnsSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSecondary

`func (o *Network) SetDnsSecondary(v string)`

SetDnsSecondary sets DnsSecondary field to given value.

### HasDnsSecondary

`func (o *Network) HasDnsSecondary() bool`

HasDnsSecondary returns a boolean if a field has been set.

### SetDnsSecondaryNil

`func (o *Network) SetDnsSecondaryNil(b bool)`

 SetDnsSecondaryNil sets the value for DnsSecondary to be an explicit nil

### UnsetDnsSecondary
`func (o *Network) UnsetDnsSecondary()`

UnsetDnsSecondary ensures that no value is present for DnsSecondary, not even an explicit nil
### GetCidr

`func (o *Network) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *Network) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *Network) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *Network) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetGatewayIPv6

`func (o *Network) GetGatewayIPv6() string`

GetGatewayIPv6 returns the GatewayIPv6 field if non-nil, zero value otherwise.

### GetGatewayIPv6Ok

`func (o *Network) GetGatewayIPv6Ok() (*string, bool)`

GetGatewayIPv6Ok returns a tuple with the GatewayIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIPv6

`func (o *Network) SetGatewayIPv6(v string)`

SetGatewayIPv6 sets GatewayIPv6 field to given value.

### HasGatewayIPv6

`func (o *Network) HasGatewayIPv6() bool`

HasGatewayIPv6 returns a boolean if a field has been set.

### SetGatewayIPv6Nil

`func (o *Network) SetGatewayIPv6Nil(b bool)`

 SetGatewayIPv6Nil sets the value for GatewayIPv6 to be an explicit nil

### UnsetGatewayIPv6
`func (o *Network) UnsetGatewayIPv6()`

UnsetGatewayIPv6 ensures that no value is present for GatewayIPv6, not even an explicit nil
### GetNetmaskIPv6

`func (o *Network) GetNetmaskIPv6() string`

GetNetmaskIPv6 returns the NetmaskIPv6 field if non-nil, zero value otherwise.

### GetNetmaskIPv6Ok

`func (o *Network) GetNetmaskIPv6Ok() (*string, bool)`

GetNetmaskIPv6Ok returns a tuple with the NetmaskIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmaskIPv6

`func (o *Network) SetNetmaskIPv6(v string)`

SetNetmaskIPv6 sets NetmaskIPv6 field to given value.

### HasNetmaskIPv6

`func (o *Network) HasNetmaskIPv6() bool`

HasNetmaskIPv6 returns a boolean if a field has been set.

### SetNetmaskIPv6Nil

`func (o *Network) SetNetmaskIPv6Nil(b bool)`

 SetNetmaskIPv6Nil sets the value for NetmaskIPv6 to be an explicit nil

### UnsetNetmaskIPv6
`func (o *Network) UnsetNetmaskIPv6()`

UnsetNetmaskIPv6 ensures that no value is present for NetmaskIPv6, not even an explicit nil
### GetDnsPrimaryIPv6

`func (o *Network) GetDnsPrimaryIPv6() string`

GetDnsPrimaryIPv6 returns the DnsPrimaryIPv6 field if non-nil, zero value otherwise.

### GetDnsPrimaryIPv6Ok

`func (o *Network) GetDnsPrimaryIPv6Ok() (*string, bool)`

GetDnsPrimaryIPv6Ok returns a tuple with the DnsPrimaryIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrimaryIPv6

`func (o *Network) SetDnsPrimaryIPv6(v string)`

SetDnsPrimaryIPv6 sets DnsPrimaryIPv6 field to given value.

### HasDnsPrimaryIPv6

`func (o *Network) HasDnsPrimaryIPv6() bool`

HasDnsPrimaryIPv6 returns a boolean if a field has been set.

### SetDnsPrimaryIPv6Nil

`func (o *Network) SetDnsPrimaryIPv6Nil(b bool)`

 SetDnsPrimaryIPv6Nil sets the value for DnsPrimaryIPv6 to be an explicit nil

### UnsetDnsPrimaryIPv6
`func (o *Network) UnsetDnsPrimaryIPv6()`

UnsetDnsPrimaryIPv6 ensures that no value is present for DnsPrimaryIPv6, not even an explicit nil
### GetDnsSecondaryIPv6

`func (o *Network) GetDnsSecondaryIPv6() string`

GetDnsSecondaryIPv6 returns the DnsSecondaryIPv6 field if non-nil, zero value otherwise.

### GetDnsSecondaryIPv6Ok

`func (o *Network) GetDnsSecondaryIPv6Ok() (*string, bool)`

GetDnsSecondaryIPv6Ok returns a tuple with the DnsSecondaryIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSecondaryIPv6

`func (o *Network) SetDnsSecondaryIPv6(v string)`

SetDnsSecondaryIPv6 sets DnsSecondaryIPv6 field to given value.

### HasDnsSecondaryIPv6

`func (o *Network) HasDnsSecondaryIPv6() bool`

HasDnsSecondaryIPv6 returns a boolean if a field has been set.

### SetDnsSecondaryIPv6Nil

`func (o *Network) SetDnsSecondaryIPv6Nil(b bool)`

 SetDnsSecondaryIPv6Nil sets the value for DnsSecondaryIPv6 to be an explicit nil

### UnsetDnsSecondaryIPv6
`func (o *Network) UnsetDnsSecondaryIPv6()`

UnsetDnsSecondaryIPv6 ensures that no value is present for DnsSecondaryIPv6, not even an explicit nil
### GetCidrIPv6

`func (o *Network) GetCidrIPv6() string`

GetCidrIPv6 returns the CidrIPv6 field if non-nil, zero value otherwise.

### GetCidrIPv6Ok

`func (o *Network) GetCidrIPv6Ok() (*string, bool)`

GetCidrIPv6Ok returns a tuple with the CidrIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrIPv6

`func (o *Network) SetCidrIPv6(v string)`

SetCidrIPv6 sets CidrIPv6 field to given value.

### HasCidrIPv6

`func (o *Network) HasCidrIPv6() bool`

HasCidrIPv6 returns a boolean if a field has been set.

### SetCidrIPv6Nil

`func (o *Network) SetCidrIPv6Nil(b bool)`

 SetCidrIPv6Nil sets the value for CidrIPv6 to be an explicit nil

### UnsetCidrIPv6
`func (o *Network) UnsetCidrIPv6()`

UnsetCidrIPv6 ensures that no value is present for CidrIPv6, not even an explicit nil
### GetTftpServer

`func (o *Network) GetTftpServer() string`

GetTftpServer returns the TftpServer field if non-nil, zero value otherwise.

### GetTftpServerOk

`func (o *Network) GetTftpServerOk() (*string, bool)`

GetTftpServerOk returns a tuple with the TftpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTftpServer

`func (o *Network) SetTftpServer(v string)`

SetTftpServer sets TftpServer field to given value.

### HasTftpServer

`func (o *Network) HasTftpServer() bool`

HasTftpServer returns a boolean if a field has been set.

### SetTftpServerNil

`func (o *Network) SetTftpServerNil(b bool)`

 SetTftpServerNil sets the value for TftpServer to be an explicit nil

### UnsetTftpServer
`func (o *Network) UnsetTftpServer()`

UnsetTftpServer ensures that no value is present for TftpServer, not even an explicit nil
### GetBootFile

`func (o *Network) GetBootFile() string`

GetBootFile returns the BootFile field if non-nil, zero value otherwise.

### GetBootFileOk

`func (o *Network) GetBootFileOk() (*string, bool)`

GetBootFileOk returns a tuple with the BootFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootFile

`func (o *Network) SetBootFile(v string)`

SetBootFile sets BootFile field to given value.

### HasBootFile

`func (o *Network) HasBootFile() bool`

HasBootFile returns a boolean if a field has been set.

### SetBootFileNil

`func (o *Network) SetBootFileNil(b bool)`

 SetBootFileNil sets the value for BootFile to be an explicit nil

### UnsetBootFile
`func (o *Network) UnsetBootFile()`

UnsetBootFile ensures that no value is present for BootFile, not even an explicit nil
### GetSwitchId

`func (o *Network) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *Network) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *Network) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *Network) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### SetSwitchIdNil

`func (o *Network) SetSwitchIdNil(b bool)`

 SetSwitchIdNil sets the value for SwitchId to be an explicit nil

### UnsetSwitchId
`func (o *Network) UnsetSwitchId()`

UnsetSwitchId ensures that no value is present for SwitchId, not even an explicit nil
### GetFabricId

`func (o *Network) GetFabricId() string`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *Network) GetFabricIdOk() (*string, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *Network) SetFabricId(v string)`

SetFabricId sets FabricId field to given value.

### HasFabricId

`func (o *Network) HasFabricId() bool`

HasFabricId returns a boolean if a field has been set.

### SetFabricIdNil

`func (o *Network) SetFabricIdNil(b bool)`

 SetFabricIdNil sets the value for FabricId to be an explicit nil

### UnsetFabricId
`func (o *Network) UnsetFabricId()`

UnsetFabricId ensures that no value is present for FabricId, not even an explicit nil
### GetNetworkRole

`func (o *Network) GetNetworkRole() string`

GetNetworkRole returns the NetworkRole field if non-nil, zero value otherwise.

### GetNetworkRoleOk

`func (o *Network) GetNetworkRoleOk() (*string, bool)`

GetNetworkRoleOk returns a tuple with the NetworkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkRole

`func (o *Network) SetNetworkRole(v string)`

SetNetworkRole sets NetworkRole field to given value.

### HasNetworkRole

`func (o *Network) HasNetworkRole() bool`

HasNetworkRole returns a boolean if a field has been set.

### SetNetworkRoleNil

`func (o *Network) SetNetworkRoleNil(b bool)`

 SetNetworkRoleNil sets the value for NetworkRole to be an explicit nil

### UnsetNetworkRole
`func (o *Network) UnsetNetworkRole()`

UnsetNetworkRole ensures that no value is present for NetworkRole, not even an explicit nil
### GetStatus

`func (o *Network) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Network) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Network) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Network) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *Network) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *Network) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *Network) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *Network) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *Network) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *Network) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetPool

`func (o *Network) GetPool() map[string]interface{}`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *Network) GetPoolOk() (*map[string]interface{}, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *Network) SetPool(v map[string]interface{})`

SetPool sets Pool field to given value.

### HasPool

`func (o *Network) HasPool() bool`

HasPool returns a boolean if a field has been set.

### SetPoolNil

`func (o *Network) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *Network) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil
### GetPoolIPv6

`func (o *Network) GetPoolIPv6() map[string]interface{}`

GetPoolIPv6 returns the PoolIPv6 field if non-nil, zero value otherwise.

### GetPoolIPv6Ok

`func (o *Network) GetPoolIPv6Ok() (*map[string]interface{}, bool)`

GetPoolIPv6Ok returns a tuple with the PoolIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolIPv6

`func (o *Network) SetPoolIPv6(v map[string]interface{})`

SetPoolIPv6 sets PoolIPv6 field to given value.

### HasPoolIPv6

`func (o *Network) HasPoolIPv6() bool`

HasPoolIPv6 returns a boolean if a field has been set.

### SetPoolIPv6Nil

`func (o *Network) SetPoolIPv6Nil(b bool)`

 SetPoolIPv6Nil sets the value for PoolIPv6 to be an explicit nil

### UnsetPoolIPv6
`func (o *Network) UnsetPoolIPv6()`

UnsetPoolIPv6 ensures that no value is present for PoolIPv6, not even an explicit nil
### GetNetworkProxy

`func (o *Network) GetNetworkProxy() ListNetworks200ResponseAllOfNetworksInnerNetworkProxy`

GetNetworkProxy returns the NetworkProxy field if non-nil, zero value otherwise.

### GetNetworkProxyOk

`func (o *Network) GetNetworkProxyOk() (*ListNetworks200ResponseAllOfNetworksInnerNetworkProxy, bool)`

GetNetworkProxyOk returns a tuple with the NetworkProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProxy

`func (o *Network) SetNetworkProxy(v ListNetworks200ResponseAllOfNetworksInnerNetworkProxy)`

SetNetworkProxy sets NetworkProxy field to given value.

### HasNetworkProxy

`func (o *Network) HasNetworkProxy() bool`

HasNetworkProxy returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *Network) GetNetworkDomain() ListNetworks200ResponseAllOfNetworksInnerNetworkDomain`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *Network) GetNetworkDomainOk() (*ListNetworks200ResponseAllOfNetworksInnerNetworkDomain, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *Network) SetNetworkDomain(v ListNetworks200ResponseAllOfNetworksInnerNetworkDomain)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *Network) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetSearchDomains

`func (o *Network) GetSearchDomains() string`

GetSearchDomains returns the SearchDomains field if non-nil, zero value otherwise.

### GetSearchDomainsOk

`func (o *Network) GetSearchDomainsOk() (*string, bool)`

GetSearchDomainsOk returns a tuple with the SearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomains

`func (o *Network) SetSearchDomains(v string)`

SetSearchDomains sets SearchDomains field to given value.

### HasSearchDomains

`func (o *Network) HasSearchDomains() bool`

HasSearchDomains returns a boolean if a field has been set.

### SetSearchDomainsNil

`func (o *Network) SetSearchDomainsNil(b bool)`

 SetSearchDomainsNil sets the value for SearchDomains to be an explicit nil

### UnsetSearchDomains
`func (o *Network) UnsetSearchDomains()`

UnsetSearchDomains ensures that no value is present for SearchDomains, not even an explicit nil
### GetPrefixLength

`func (o *Network) GetPrefixLength() string`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *Network) GetPrefixLengthOk() (*string, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *Network) SetPrefixLength(v string)`

SetPrefixLength sets PrefixLength field to given value.

### HasPrefixLength

`func (o *Network) HasPrefixLength() bool`

HasPrefixLength returns a boolean if a field has been set.

### SetPrefixLengthNil

`func (o *Network) SetPrefixLengthNil(b bool)`

 SetPrefixLengthNil sets the value for PrefixLength to be an explicit nil

### UnsetPrefixLength
`func (o *Network) UnsetPrefixLength()`

UnsetPrefixLength ensures that no value is present for PrefixLength, not even an explicit nil
### GetVisibility

`func (o *Network) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Network) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Network) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Network) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEnableAdmin

`func (o *Network) GetEnableAdmin() bool`

GetEnableAdmin returns the EnableAdmin field if non-nil, zero value otherwise.

### GetEnableAdminOk

`func (o *Network) GetEnableAdminOk() (*bool, bool)`

GetEnableAdminOk returns a tuple with the EnableAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdmin

`func (o *Network) SetEnableAdmin(v bool)`

SetEnableAdmin sets EnableAdmin field to given value.

### HasEnableAdmin

`func (o *Network) HasEnableAdmin() bool`

HasEnableAdmin returns a boolean if a field has been set.

### GetActive

`func (o *Network) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Network) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Network) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Network) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDefaultNetwork

`func (o *Network) GetDefaultNetwork() bool`

GetDefaultNetwork returns the DefaultNetwork field if non-nil, zero value otherwise.

### GetDefaultNetworkOk

`func (o *Network) GetDefaultNetworkOk() (*bool, bool)`

GetDefaultNetworkOk returns a tuple with the DefaultNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetwork

`func (o *Network) SetDefaultNetwork(v bool)`

SetDefaultNetwork sets DefaultNetwork field to given value.

### HasDefaultNetwork

`func (o *Network) HasDefaultNetwork() bool`

HasDefaultNetwork returns a boolean if a field has been set.

### GetAssignPublicIp

`func (o *Network) GetAssignPublicIp() bool`

GetAssignPublicIp returns the AssignPublicIp field if non-nil, zero value otherwise.

### GetAssignPublicIpOk

`func (o *Network) GetAssignPublicIpOk() (*bool, bool)`

GetAssignPublicIpOk returns a tuple with the AssignPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignPublicIp

`func (o *Network) SetAssignPublicIp(v bool)`

SetAssignPublicIp sets AssignPublicIp field to given value.

### HasAssignPublicIp

`func (o *Network) HasAssignPublicIp() bool`

HasAssignPublicIp returns a boolean if a field has been set.

### GetNoProxy

`func (o *Network) GetNoProxy() string`

GetNoProxy returns the NoProxy field if non-nil, zero value otherwise.

### GetNoProxyOk

`func (o *Network) GetNoProxyOk() (*string, bool)`

GetNoProxyOk returns a tuple with the NoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxy

`func (o *Network) SetNoProxy(v string)`

SetNoProxy sets NoProxy field to given value.

### HasNoProxy

`func (o *Network) HasNoProxy() bool`

HasNoProxy returns a boolean if a field has been set.

### SetNoProxyNil

`func (o *Network) SetNoProxyNil(b bool)`

 SetNoProxyNil sets the value for NoProxy to be an explicit nil

### UnsetNoProxy
`func (o *Network) UnsetNoProxy()`

UnsetNoProxy ensures that no value is present for NoProxy, not even an explicit nil
### GetApplianceUrlProxyBypass

`func (o *Network) GetApplianceUrlProxyBypass() bool`

GetApplianceUrlProxyBypass returns the ApplianceUrlProxyBypass field if non-nil, zero value otherwise.

### GetApplianceUrlProxyBypassOk

`func (o *Network) GetApplianceUrlProxyBypassOk() (*bool, bool)`

GetApplianceUrlProxyBypassOk returns a tuple with the ApplianceUrlProxyBypass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrlProxyBypass

`func (o *Network) SetApplianceUrlProxyBypass(v bool)`

SetApplianceUrlProxyBypass sets ApplianceUrlProxyBypass field to given value.

### HasApplianceUrlProxyBypass

`func (o *Network) HasApplianceUrlProxyBypass() bool`

HasApplianceUrlProxyBypass returns a boolean if a field has been set.

### GetZonePool

`func (o *Network) GetZonePool() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *Network) GetZonePoolOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *Network) SetZonePool(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *Network) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetAllowStaticOverride

`func (o *Network) GetAllowStaticOverride() bool`

GetAllowStaticOverride returns the AllowStaticOverride field if non-nil, zero value otherwise.

### GetAllowStaticOverrideOk

`func (o *Network) GetAllowStaticOverrideOk() (*bool, bool)`

GetAllowStaticOverrideOk returns a tuple with the AllowStaticOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowStaticOverride

`func (o *Network) SetAllowStaticOverride(v bool)`

SetAllowStaticOverride sets AllowStaticOverride field to given value.

### HasAllowStaticOverride

`func (o *Network) HasAllowStaticOverride() bool`

HasAllowStaticOverride returns a boolean if a field has been set.

### GetConfig

`func (o *Network) GetConfig() ListNetworks200ResponseAllOfNetworksInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Network) GetConfigOk() (*ListNetworks200ResponseAllOfNetworksInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Network) SetConfig(v ListNetworks200ResponseAllOfNetworksInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Network) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTenants

`func (o *Network) GetTenants() []GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *Network) GetTenantsOk() (*[]GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *Network) SetTenants(v []GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *Network) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


