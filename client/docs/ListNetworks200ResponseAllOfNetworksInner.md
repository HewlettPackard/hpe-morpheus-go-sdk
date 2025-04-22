# ListNetworks200ResponseAllOfNetworksInner

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
**Code** | Pointer to **string** | Network Code | [optional] 
**Ipv4Enabled** | Pointer to **bool** |  | [optional] 
**Ipv6Enabled** | Pointer to **bool** |  | [optional] 
**Category** | Pointer to **string** | Network Category | [optional] 
**InterfaceName** | Pointer to **string** |  | [optional] 
**BridgeName** | Pointer to **string** |  | [optional] 
**BridgeInterface** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**ExternalType** | Pointer to **string** |  | [optional] 
**RefUrl** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**VlanId** | Pointer to **int64** |  | [optional] 
**VswitchName** | Pointer to **string** |  | [optional] 
**DhcpServer** | Pointer to **bool** |  | [optional] 
**DhcpIp** | Pointer to **string** |  | [optional] 
**DhcpServerIPv6** | Pointer to **bool** |  | [optional] 
**Gateway** | Pointer to **string** | Network Gateway | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Broadcast** | Pointer to **string** |  | [optional] 
**SubnetAddress** | Pointer to **string** |  | [optional] 
**DnsPrimary** | Pointer to **string** | Primary DNS Server | [optional] 
**DnsSecondary** | Pointer to **string** | Secondary DNS Server | [optional] 
**Cidr** | Pointer to **string** | Network CIDR | [optional] 
**GatewayIPv6** | Pointer to **string** | IPv6 Network Gateway | [optional] 
**NetmaskIPv6** | Pointer to **string** |  | [optional] 
**DnsPrimaryIPv6** | Pointer to **string** | Primary IPv6 DNS Server | [optional] 
**DnsSecondaryIPv6** | Pointer to **string** | Secondary IPv6 DNS Server | [optional] 
**CidrIPv6** | Pointer to **string** | IPv6 Network CIDR | [optional] 
**TftpServer** | Pointer to **string** |  | [optional] 
**BootFile** | Pointer to **string** |  | [optional] 
**SwitchId** | Pointer to **string** |  | [optional] 
**FabricId** | Pointer to **string** |  | [optional] 
**NetworkRole** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AvailabilityZone** | Pointer to **string** |  | [optional] 
**Pool** | Pointer to **map[string]interface{}** |  | [optional] 
**PoolIPv6** | Pointer to **map[string]interface{}** |  | [optional] 
**NetworkProxy** | Pointer to [**ListNetworks200ResponseAllOfNetworksInnerNetworkProxy**](ListNetworks200ResponseAllOfNetworksInnerNetworkProxy.md) |  | [optional] 
**NetworkDomain** | Pointer to [**ListNetworks200ResponseAllOfNetworksInnerNetworkDomain**](ListNetworks200ResponseAllOfNetworksInnerNetworkDomain.md) |  | [optional] 
**SearchDomains** | Pointer to **string** |  | [optional] 
**PrefixLength** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**EnableAdmin** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**DefaultNetwork** | Pointer to **bool** |  | [optional] 
**AssignPublicIp** | Pointer to **bool** |  | [optional] 
**NoProxy** | Pointer to **string** |  | [optional] 
**ApplianceUrlProxyBypass** | Pointer to **bool** |  | [optional] 
**ZonePool** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**AllowStaticOverride** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**ListNetworks200ResponseAllOfNetworksInnerConfig**](ListNetworks200ResponseAllOfNetworksInnerConfig.md) |  | [optional] 
**Tenants** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 

## Methods

### NewListNetworks200ResponseAllOfNetworksInner

`func NewListNetworks200ResponseAllOfNetworksInner() *ListNetworks200ResponseAllOfNetworksInner`

NewListNetworks200ResponseAllOfNetworksInner instantiates a new ListNetworks200ResponseAllOfNetworksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworks200ResponseAllOfNetworksInnerWithDefaults

`func NewListNetworks200ResponseAllOfNetworksInnerWithDefaults() *ListNetworks200ResponseAllOfNetworksInner`

NewListNetworks200ResponseAllOfNetworksInnerWithDefaults instantiates a new ListNetworks200ResponseAllOfNetworksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLabels

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetZone

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetZone() ListNetworks200ResponseAllOfNetworksInnerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetZoneOk() (*ListNetworks200ResponseAllOfNetworksInnerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetZone(v ListNetworks200ResponseAllOfNetworksInnerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetType

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetType() ListNetworks200ResponseAllOfNetworksInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetTypeOk() (*ListNetworks200ResponseAllOfNetworksInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetType(v ListNetworks200ResponseAllOfNetworksInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOwner

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetOwner() ListNetworks200ResponseAllOfNetworksInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetOwnerOk() (*ListNetworks200ResponseAllOfNetworksInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetOwner(v ListNetworks200ResponseAllOfNetworksInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCode

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetIpv4Enabled

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetIpv4Enabled() bool`

GetIpv4Enabled returns the Ipv4Enabled field if non-nil, zero value otherwise.

### GetIpv4EnabledOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetIpv4EnabledOk() (*bool, bool)`

GetIpv4EnabledOk returns a tuple with the Ipv4Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Enabled

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetIpv4Enabled(v bool)`

SetIpv4Enabled sets Ipv4Enabled field to given value.

### HasIpv4Enabled

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasIpv4Enabled() bool`

HasIpv4Enabled returns a boolean if a field has been set.

### GetIpv6Enabled

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetIpv6Enabled() bool`

GetIpv6Enabled returns the Ipv6Enabled field if non-nil, zero value otherwise.

### GetIpv6EnabledOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetIpv6EnabledOk() (*bool, bool)`

GetIpv6EnabledOk returns a tuple with the Ipv6Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Enabled

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetIpv6Enabled(v bool)`

SetIpv6Enabled sets Ipv6Enabled field to given value.

### HasIpv6Enabled

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasIpv6Enabled() bool`

HasIpv6Enabled returns a boolean if a field has been set.

### GetCategory

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetInterfaceName

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetBridgeName

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetBridgeName() string`

GetBridgeName returns the BridgeName field if non-nil, zero value otherwise.

### GetBridgeNameOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetBridgeNameOk() (*string, bool)`

GetBridgeNameOk returns a tuple with the BridgeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeName

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetBridgeName(v string)`

SetBridgeName sets BridgeName field to given value.

### HasBridgeName

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasBridgeName() bool`

HasBridgeName returns a boolean if a field has been set.

### GetBridgeInterface

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetBridgeInterface() string`

GetBridgeInterface returns the BridgeInterface field if non-nil, zero value otherwise.

### GetBridgeInterfaceOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetBridgeInterfaceOk() (*string, bool)`

GetBridgeInterfaceOk returns a tuple with the BridgeInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeInterface

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetBridgeInterface(v string)`

SetBridgeInterface sets BridgeInterface field to given value.

### HasBridgeInterface

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasBridgeInterface() bool`

HasBridgeInterface returns a boolean if a field has been set.

### GetDescription

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalId

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInternalId

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetUniqueId

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetExternalType

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### GetRefUrl

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetRefUrl() string`

GetRefUrl returns the RefUrl field if non-nil, zero value otherwise.

### GetRefUrlOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetRefUrlOk() (*string, bool)`

GetRefUrlOk returns a tuple with the RefUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUrl

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetRefUrl(v string)`

SetRefUrl sets RefUrl field to given value.

### HasRefUrl

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasRefUrl() bool`

HasRefUrl returns a boolean if a field has been set.

### GetRefType

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetVlanId

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVswitchName

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetVswitchName() string`

GetVswitchName returns the VswitchName field if non-nil, zero value otherwise.

### GetVswitchNameOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetVswitchNameOk() (*string, bool)`

GetVswitchNameOk returns a tuple with the VswitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchName

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetVswitchName(v string)`

SetVswitchName sets VswitchName field to given value.

### HasVswitchName

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasVswitchName() bool`

HasVswitchName returns a boolean if a field has been set.

### GetDhcpServer

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetDhcpServer() bool`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetDhcpServerOk() (*bool, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetDhcpServer(v bool)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetDhcpIp

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetDhcpIp() string`

GetDhcpIp returns the DhcpIp field if non-nil, zero value otherwise.

### GetDhcpIpOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetDhcpIpOk() (*string, bool)`

GetDhcpIpOk returns a tuple with the DhcpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpIp

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetDhcpIp(v string)`

SetDhcpIp sets DhcpIp field to given value.

### HasDhcpIp

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasDhcpIp() bool`

HasDhcpIp returns a boolean if a field has been set.

### GetDhcpServerIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetDhcpServerIPv6() bool`

GetDhcpServerIPv6 returns the DhcpServerIPv6 field if non-nil, zero value otherwise.

### GetDhcpServerIPv6Ok

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetDhcpServerIPv6Ok() (*bool, bool)`

GetDhcpServerIPv6Ok returns a tuple with the DhcpServerIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetDhcpServerIPv6(v bool)`

SetDhcpServerIPv6 sets DhcpServerIPv6 field to given value.

### HasDhcpServerIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasDhcpServerIPv6() bool`

HasDhcpServerIPv6 returns a boolean if a field has been set.

### GetGateway

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetNetmask

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetBroadcast

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetBroadcast() string`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetBroadcastOk() (*string, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetBroadcast(v string)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.

### GetSubnetAddress

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetSubnetAddress() string`

GetSubnetAddress returns the SubnetAddress field if non-nil, zero value otherwise.

### GetSubnetAddressOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetSubnetAddressOk() (*string, bool)`

GetSubnetAddressOk returns a tuple with the SubnetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAddress

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetSubnetAddress(v string)`

SetSubnetAddress sets SubnetAddress field to given value.

### HasSubnetAddress

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasSubnetAddress() bool`

HasSubnetAddress returns a boolean if a field has been set.

### GetDnsPrimary

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetDnsPrimary() string`

GetDnsPrimary returns the DnsPrimary field if non-nil, zero value otherwise.

### GetDnsPrimaryOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetDnsPrimaryOk() (*string, bool)`

GetDnsPrimaryOk returns a tuple with the DnsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrimary

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetDnsPrimary(v string)`

SetDnsPrimary sets DnsPrimary field to given value.

### HasDnsPrimary

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasDnsPrimary() bool`

HasDnsPrimary returns a boolean if a field has been set.

### GetDnsSecondary

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetDnsSecondary() string`

GetDnsSecondary returns the DnsSecondary field if non-nil, zero value otherwise.

### GetDnsSecondaryOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetDnsSecondaryOk() (*string, bool)`

GetDnsSecondaryOk returns a tuple with the DnsSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSecondary

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetDnsSecondary(v string)`

SetDnsSecondary sets DnsSecondary field to given value.

### HasDnsSecondary

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasDnsSecondary() bool`

HasDnsSecondary returns a boolean if a field has been set.

### GetCidr

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetGatewayIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetGatewayIPv6() string`

GetGatewayIPv6 returns the GatewayIPv6 field if non-nil, zero value otherwise.

### GetGatewayIPv6Ok

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetGatewayIPv6Ok() (*string, bool)`

GetGatewayIPv6Ok returns a tuple with the GatewayIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetGatewayIPv6(v string)`

SetGatewayIPv6 sets GatewayIPv6 field to given value.

### HasGatewayIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasGatewayIPv6() bool`

HasGatewayIPv6 returns a boolean if a field has been set.

### GetNetmaskIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetNetmaskIPv6() string`

GetNetmaskIPv6 returns the NetmaskIPv6 field if non-nil, zero value otherwise.

### GetNetmaskIPv6Ok

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetNetmaskIPv6Ok() (*string, bool)`

GetNetmaskIPv6Ok returns a tuple with the NetmaskIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmaskIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetNetmaskIPv6(v string)`

SetNetmaskIPv6 sets NetmaskIPv6 field to given value.

### HasNetmaskIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasNetmaskIPv6() bool`

HasNetmaskIPv6 returns a boolean if a field has been set.

### GetDnsPrimaryIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetDnsPrimaryIPv6() string`

GetDnsPrimaryIPv6 returns the DnsPrimaryIPv6 field if non-nil, zero value otherwise.

### GetDnsPrimaryIPv6Ok

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetDnsPrimaryIPv6Ok() (*string, bool)`

GetDnsPrimaryIPv6Ok returns a tuple with the DnsPrimaryIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrimaryIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetDnsPrimaryIPv6(v string)`

SetDnsPrimaryIPv6 sets DnsPrimaryIPv6 field to given value.

### HasDnsPrimaryIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasDnsPrimaryIPv6() bool`

HasDnsPrimaryIPv6 returns a boolean if a field has been set.

### GetDnsSecondaryIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetDnsSecondaryIPv6() string`

GetDnsSecondaryIPv6 returns the DnsSecondaryIPv6 field if non-nil, zero value otherwise.

### GetDnsSecondaryIPv6Ok

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetDnsSecondaryIPv6Ok() (*string, bool)`

GetDnsSecondaryIPv6Ok returns a tuple with the DnsSecondaryIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSecondaryIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetDnsSecondaryIPv6(v string)`

SetDnsSecondaryIPv6 sets DnsSecondaryIPv6 field to given value.

### HasDnsSecondaryIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasDnsSecondaryIPv6() bool`

HasDnsSecondaryIPv6 returns a boolean if a field has been set.

### GetCidrIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetCidrIPv6() string`

GetCidrIPv6 returns the CidrIPv6 field if non-nil, zero value otherwise.

### GetCidrIPv6Ok

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetCidrIPv6Ok() (*string, bool)`

GetCidrIPv6Ok returns a tuple with the CidrIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetCidrIPv6(v string)`

SetCidrIPv6 sets CidrIPv6 field to given value.

### HasCidrIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasCidrIPv6() bool`

HasCidrIPv6 returns a boolean if a field has been set.

### GetTftpServer

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetTftpServer() string`

GetTftpServer returns the TftpServer field if non-nil, zero value otherwise.

### GetTftpServerOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetTftpServerOk() (*string, bool)`

GetTftpServerOk returns a tuple with the TftpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTftpServer

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetTftpServer(v string)`

SetTftpServer sets TftpServer field to given value.

### HasTftpServer

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasTftpServer() bool`

HasTftpServer returns a boolean if a field has been set.

### GetBootFile

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetBootFile() string`

GetBootFile returns the BootFile field if non-nil, zero value otherwise.

### GetBootFileOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetBootFileOk() (*string, bool)`

GetBootFileOk returns a tuple with the BootFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootFile

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetBootFile(v string)`

SetBootFile sets BootFile field to given value.

### HasBootFile

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasBootFile() bool`

HasBootFile returns a boolean if a field has been set.

### GetSwitchId

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetFabricId

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetFabricId() string`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetFabricIdOk() (*string, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetFabricId(v string)`

SetFabricId sets FabricId field to given value.

### HasFabricId

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasFabricId() bool`

HasFabricId returns a boolean if a field has been set.

### GetNetworkRole

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetNetworkRole() string`

GetNetworkRole returns the NetworkRole field if non-nil, zero value otherwise.

### GetNetworkRoleOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetNetworkRoleOk() (*string, bool)`

GetNetworkRoleOk returns a tuple with the NetworkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkRole

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetNetworkRole(v string)`

SetNetworkRole sets NetworkRole field to given value.

### HasNetworkRole

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasNetworkRole() bool`

HasNetworkRole returns a boolean if a field has been set.

### GetStatus

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetPool

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetPool() map[string]interface{}`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetPoolOk() (*map[string]interface{}, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetPool(v map[string]interface{})`

SetPool sets Pool field to given value.

### HasPool

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetPoolIPv6() map[string]interface{}`

GetPoolIPv6 returns the PoolIPv6 field if non-nil, zero value otherwise.

### GetPoolIPv6Ok

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetPoolIPv6Ok() (*map[string]interface{}, bool)`

GetPoolIPv6Ok returns a tuple with the PoolIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetPoolIPv6(v map[string]interface{})`

SetPoolIPv6 sets PoolIPv6 field to given value.

### HasPoolIPv6

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasPoolIPv6() bool`

HasPoolIPv6 returns a boolean if a field has been set.

### GetNetworkProxy

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetNetworkProxy() ListNetworks200ResponseAllOfNetworksInnerNetworkProxy`

GetNetworkProxy returns the NetworkProxy field if non-nil, zero value otherwise.

### GetNetworkProxyOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetNetworkProxyOk() (*ListNetworks200ResponseAllOfNetworksInnerNetworkProxy, bool)`

GetNetworkProxyOk returns a tuple with the NetworkProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProxy

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetNetworkProxy(v ListNetworks200ResponseAllOfNetworksInnerNetworkProxy)`

SetNetworkProxy sets NetworkProxy field to given value.

### HasNetworkProxy

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasNetworkProxy() bool`

HasNetworkProxy returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetNetworkDomain() ListNetworks200ResponseAllOfNetworksInnerNetworkDomain`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetNetworkDomainOk() (*ListNetworks200ResponseAllOfNetworksInnerNetworkDomain, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetNetworkDomain(v ListNetworks200ResponseAllOfNetworksInnerNetworkDomain)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetSearchDomains

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetSearchDomains() string`

GetSearchDomains returns the SearchDomains field if non-nil, zero value otherwise.

### GetSearchDomainsOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetSearchDomainsOk() (*string, bool)`

GetSearchDomainsOk returns a tuple with the SearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomains

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetSearchDomains(v string)`

SetSearchDomains sets SearchDomains field to given value.

### HasSearchDomains

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasSearchDomains() bool`

HasSearchDomains returns a boolean if a field has been set.

### GetPrefixLength

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetPrefixLength() string`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetPrefixLengthOk() (*string, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetPrefixLength(v string)`

SetPrefixLength sets PrefixLength field to given value.

### HasPrefixLength

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasPrefixLength() bool`

HasPrefixLength returns a boolean if a field has been set.

### GetVisibility

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEnableAdmin

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetEnableAdmin() bool`

GetEnableAdmin returns the EnableAdmin field if non-nil, zero value otherwise.

### GetEnableAdminOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetEnableAdminOk() (*bool, bool)`

GetEnableAdminOk returns a tuple with the EnableAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdmin

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetEnableAdmin(v bool)`

SetEnableAdmin sets EnableAdmin field to given value.

### HasEnableAdmin

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasEnableAdmin() bool`

HasEnableAdmin returns a boolean if a field has been set.

### GetActive

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDefaultNetwork

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetDefaultNetwork() bool`

GetDefaultNetwork returns the DefaultNetwork field if non-nil, zero value otherwise.

### GetDefaultNetworkOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetDefaultNetworkOk() (*bool, bool)`

GetDefaultNetworkOk returns a tuple with the DefaultNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetwork

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetDefaultNetwork(v bool)`

SetDefaultNetwork sets DefaultNetwork field to given value.

### HasDefaultNetwork

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasDefaultNetwork() bool`

HasDefaultNetwork returns a boolean if a field has been set.

### GetAssignPublicIp

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetAssignPublicIp() bool`

GetAssignPublicIp returns the AssignPublicIp field if non-nil, zero value otherwise.

### GetAssignPublicIpOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetAssignPublicIpOk() (*bool, bool)`

GetAssignPublicIpOk returns a tuple with the AssignPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignPublicIp

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetAssignPublicIp(v bool)`

SetAssignPublicIp sets AssignPublicIp field to given value.

### HasAssignPublicIp

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasAssignPublicIp() bool`

HasAssignPublicIp returns a boolean if a field has been set.

### GetNoProxy

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetNoProxy() string`

GetNoProxy returns the NoProxy field if non-nil, zero value otherwise.

### GetNoProxyOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetNoProxyOk() (*string, bool)`

GetNoProxyOk returns a tuple with the NoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxy

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetNoProxy(v string)`

SetNoProxy sets NoProxy field to given value.

### HasNoProxy

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasNoProxy() bool`

HasNoProxy returns a boolean if a field has been set.

### GetApplianceUrlProxyBypass

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetApplianceUrlProxyBypass() bool`

GetApplianceUrlProxyBypass returns the ApplianceUrlProxyBypass field if non-nil, zero value otherwise.

### GetApplianceUrlProxyBypassOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetApplianceUrlProxyBypassOk() (*bool, bool)`

GetApplianceUrlProxyBypassOk returns a tuple with the ApplianceUrlProxyBypass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrlProxyBypass

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetApplianceUrlProxyBypass(v bool)`

SetApplianceUrlProxyBypass sets ApplianceUrlProxyBypass field to given value.

### HasApplianceUrlProxyBypass

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasApplianceUrlProxyBypass() bool`

HasApplianceUrlProxyBypass returns a boolean if a field has been set.

### GetZonePool

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetZonePool() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetZonePoolOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetZonePool(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetAllowStaticOverride

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetAllowStaticOverride() bool`

GetAllowStaticOverride returns the AllowStaticOverride field if non-nil, zero value otherwise.

### GetAllowStaticOverrideOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetAllowStaticOverrideOk() (*bool, bool)`

GetAllowStaticOverrideOk returns a tuple with the AllowStaticOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowStaticOverride

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetAllowStaticOverride(v bool)`

SetAllowStaticOverride sets AllowStaticOverride field to given value.

### HasAllowStaticOverride

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasAllowStaticOverride() bool`

HasAllowStaticOverride returns a boolean if a field has been set.

### GetConfig

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetConfig() ListNetworks200ResponseAllOfNetworksInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetConfigOk() (*ListNetworks200ResponseAllOfNetworksInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetConfig(v ListNetworks200ResponseAllOfNetworksInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTenants

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetTenants() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ListNetworks200ResponseAllOfNetworksInner) GetTenantsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ListNetworks200ResponseAllOfNetworksInner) SetTenants(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ListNetworks200ResponseAllOfNetworksInner) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


