# GetNetwork200ResponseNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Network ID | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**DisplayName** | Pointer to **NullableString** | Network Display Name | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Group** | Pointer to [**GetNetwork200ResponseNetworkGroup**](GetNetwork200ResponseNetworkGroup.md) |  | [optional] 
**Zone** | Pointer to [**GetNetwork200ResponseNetworkZone**](GetNetwork200ResponseNetworkZone.md) |  | [optional] 
**Type** | Pointer to [**GetNetwork200ResponseNetworkType**](GetNetwork200ResponseNetworkType.md) |  | [optional] 
**Owner** | Pointer to [**GetNetwork200ResponseNetworkOwner**](GetNetwork200ResponseNetworkOwner.md) |  | [optional] 
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
**Cidr** | Pointer to **NullableString** | Network CIDR | [optional] 
**GatewayIPv6** | Pointer to **NullableString** | IPv6 Network Gateway | [optional] 
**NetmaskIPv6** | Pointer to **NullableString** |  | [optional] 
**DnsPrimaryIPv6** | Pointer to **NullableString** | Primary IPv6 DNS Server | [optional] 
**DnsSecondaryIPv6** | Pointer to **NullableString** | Secondary IPv6 DNS Server | [optional] 
**CidrIPv6** | Pointer to **NullableString** | IPv6 Network CIDR | [optional] 
**TftpServer** | Pointer to **NullableString** |  | [optional] 
**BootFile** | Pointer to **NullableString** |  | [optional] 
**SwitchId** | Pointer to **NullableString** | Network switch identifier | [optional] 
**FabricId** | Pointer to **NullableString** |  | [optional] 
**NetworkRole** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**AvailabilityZone** | Pointer to **NullableString** |  | [optional] 
**Pool** | Pointer to [**GetNetwork200ResponseNetworkPool**](GetNetwork200ResponseNetworkPool.md) |  | [optional] 
**PoolIPv6** | Pointer to [**GetNetwork200ResponseNetworkPoolIPv6**](GetNetwork200ResponseNetworkPoolIPv6.md) |  | [optional] 
**NetworkProxy** | Pointer to [**GetNetwork200ResponseNetworkNetworkProxy**](GetNetwork200ResponseNetworkNetworkProxy.md) |  | [optional] 
**NetworkDomain** | Pointer to [**GetNetwork200ResponseNetworkNetworkDomain**](GetNetwork200ResponseNetworkNetworkDomain.md) |  | [optional] 
**SearchDomains** | Pointer to **NullableString** |  | [optional] 
**PrefixLength** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**EnableAdmin** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**DefaultNetwork** | Pointer to **bool** |  | [optional] 
**AssignPublicIp** | Pointer to **bool** |  | [optional] 
**NoProxy** | Pointer to **NullableString** |  | [optional] 
**ApplianceUrlProxyBypass** | Pointer to **bool** |  | [optional] 
**ZonePool** | Pointer to [**GetNetwork200ResponseNetworkZonePool**](GetNetwork200ResponseNetworkZonePool.md) |  | [optional] 
**AllowStaticOverride** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**GetNetwork200ResponseNetworkConfig**](GetNetwork200ResponseNetworkConfig.md) |  | [optional] 
**Tenants** | Pointer to [**[]GetNetwork200ResponseNetworkTenantsInner**](GetNetwork200ResponseNetworkTenantsInner.md) |  | [optional] 
**ResourcePermission** | Pointer to [**GetNetwork200ResponseNetworkResourcePermission**](GetNetwork200ResponseNetworkResourcePermission.md) |  | [optional] 

## Methods

### NewGetNetwork200ResponseNetwork

`func NewGetNetwork200ResponseNetwork() *GetNetwork200ResponseNetwork`

NewGetNetwork200ResponseNetwork instantiates a new GetNetwork200ResponseNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetwork200ResponseNetworkWithDefaults

`func NewGetNetwork200ResponseNetworkWithDefaults() *GetNetwork200ResponseNetwork`

NewGetNetwork200ResponseNetworkWithDefaults instantiates a new GetNetwork200ResponseNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetwork200ResponseNetwork) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetwork200ResponseNetwork) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetwork200ResponseNetwork) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetwork200ResponseNetwork) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetwork200ResponseNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetwork200ResponseNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetwork200ResponseNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetwork200ResponseNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetNetwork200ResponseNetwork) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetNetwork200ResponseNetwork) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetNetwork200ResponseNetwork) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetNetwork200ResponseNetwork) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *GetNetwork200ResponseNetwork) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *GetNetwork200ResponseNetwork) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLabels

`func (o *GetNetwork200ResponseNetwork) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetNetwork200ResponseNetwork) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetNetwork200ResponseNetwork) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetNetwork200ResponseNetwork) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *GetNetwork200ResponseNetwork) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *GetNetwork200ResponseNetwork) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetGroup

`func (o *GetNetwork200ResponseNetwork) GetGroup() GetNetwork200ResponseNetworkGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GetNetwork200ResponseNetwork) GetGroupOk() (*GetNetwork200ResponseNetworkGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GetNetwork200ResponseNetwork) SetGroup(v GetNetwork200ResponseNetworkGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GetNetwork200ResponseNetwork) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetZone

`func (o *GetNetwork200ResponseNetwork) GetZone() GetNetwork200ResponseNetworkZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetNetwork200ResponseNetwork) GetZoneOk() (*GetNetwork200ResponseNetworkZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetNetwork200ResponseNetwork) SetZone(v GetNetwork200ResponseNetworkZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetNetwork200ResponseNetwork) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetType

`func (o *GetNetwork200ResponseNetwork) GetType() GetNetwork200ResponseNetworkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetwork200ResponseNetwork) GetTypeOk() (*GetNetwork200ResponseNetworkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetwork200ResponseNetwork) SetType(v GetNetwork200ResponseNetworkType)`

SetType sets Type field to given value.

### HasType

`func (o *GetNetwork200ResponseNetwork) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOwner

`func (o *GetNetwork200ResponseNetwork) GetOwner() GetNetwork200ResponseNetworkOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetNetwork200ResponseNetwork) GetOwnerOk() (*GetNetwork200ResponseNetworkOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetNetwork200ResponseNetwork) SetOwner(v GetNetwork200ResponseNetworkOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetNetwork200ResponseNetwork) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCode

`func (o *GetNetwork200ResponseNetwork) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetNetwork200ResponseNetwork) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetNetwork200ResponseNetwork) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetNetwork200ResponseNetwork) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GetNetwork200ResponseNetwork) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GetNetwork200ResponseNetwork) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetIpv4Enabled

`func (o *GetNetwork200ResponseNetwork) GetIpv4Enabled() bool`

GetIpv4Enabled returns the Ipv4Enabled field if non-nil, zero value otherwise.

### GetIpv4EnabledOk

`func (o *GetNetwork200ResponseNetwork) GetIpv4EnabledOk() (*bool, bool)`

GetIpv4EnabledOk returns a tuple with the Ipv4Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Enabled

`func (o *GetNetwork200ResponseNetwork) SetIpv4Enabled(v bool)`

SetIpv4Enabled sets Ipv4Enabled field to given value.

### HasIpv4Enabled

`func (o *GetNetwork200ResponseNetwork) HasIpv4Enabled() bool`

HasIpv4Enabled returns a boolean if a field has been set.

### GetIpv6Enabled

`func (o *GetNetwork200ResponseNetwork) GetIpv6Enabled() bool`

GetIpv6Enabled returns the Ipv6Enabled field if non-nil, zero value otherwise.

### GetIpv6EnabledOk

`func (o *GetNetwork200ResponseNetwork) GetIpv6EnabledOk() (*bool, bool)`

GetIpv6EnabledOk returns a tuple with the Ipv6Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Enabled

`func (o *GetNetwork200ResponseNetwork) SetIpv6Enabled(v bool)`

SetIpv6Enabled sets Ipv6Enabled field to given value.

### HasIpv6Enabled

`func (o *GetNetwork200ResponseNetwork) HasIpv6Enabled() bool`

HasIpv6Enabled returns a boolean if a field has been set.

### GetCategory

`func (o *GetNetwork200ResponseNetwork) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetNetwork200ResponseNetwork) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetNetwork200ResponseNetwork) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetNetwork200ResponseNetwork) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetNetwork200ResponseNetwork) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetNetwork200ResponseNetwork) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetInterfaceName

`func (o *GetNetwork200ResponseNetwork) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *GetNetwork200ResponseNetwork) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *GetNetwork200ResponseNetwork) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *GetNetwork200ResponseNetwork) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### SetInterfaceNameNil

`func (o *GetNetwork200ResponseNetwork) SetInterfaceNameNil(b bool)`

 SetInterfaceNameNil sets the value for InterfaceName to be an explicit nil

### UnsetInterfaceName
`func (o *GetNetwork200ResponseNetwork) UnsetInterfaceName()`

UnsetInterfaceName ensures that no value is present for InterfaceName, not even an explicit nil
### GetBridgeName

`func (o *GetNetwork200ResponseNetwork) GetBridgeName() string`

GetBridgeName returns the BridgeName field if non-nil, zero value otherwise.

### GetBridgeNameOk

`func (o *GetNetwork200ResponseNetwork) GetBridgeNameOk() (*string, bool)`

GetBridgeNameOk returns a tuple with the BridgeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeName

`func (o *GetNetwork200ResponseNetwork) SetBridgeName(v string)`

SetBridgeName sets BridgeName field to given value.

### HasBridgeName

`func (o *GetNetwork200ResponseNetwork) HasBridgeName() bool`

HasBridgeName returns a boolean if a field has been set.

### SetBridgeNameNil

`func (o *GetNetwork200ResponseNetwork) SetBridgeNameNil(b bool)`

 SetBridgeNameNil sets the value for BridgeName to be an explicit nil

### UnsetBridgeName
`func (o *GetNetwork200ResponseNetwork) UnsetBridgeName()`

UnsetBridgeName ensures that no value is present for BridgeName, not even an explicit nil
### GetBridgeInterface

`func (o *GetNetwork200ResponseNetwork) GetBridgeInterface() string`

GetBridgeInterface returns the BridgeInterface field if non-nil, zero value otherwise.

### GetBridgeInterfaceOk

`func (o *GetNetwork200ResponseNetwork) GetBridgeInterfaceOk() (*string, bool)`

GetBridgeInterfaceOk returns a tuple with the BridgeInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeInterface

`func (o *GetNetwork200ResponseNetwork) SetBridgeInterface(v string)`

SetBridgeInterface sets BridgeInterface field to given value.

### HasBridgeInterface

`func (o *GetNetwork200ResponseNetwork) HasBridgeInterface() bool`

HasBridgeInterface returns a boolean if a field has been set.

### SetBridgeInterfaceNil

`func (o *GetNetwork200ResponseNetwork) SetBridgeInterfaceNil(b bool)`

 SetBridgeInterfaceNil sets the value for BridgeInterface to be an explicit nil

### UnsetBridgeInterface
`func (o *GetNetwork200ResponseNetwork) UnsetBridgeInterface()`

UnsetBridgeInterface ensures that no value is present for BridgeInterface, not even an explicit nil
### GetDescription

`func (o *GetNetwork200ResponseNetwork) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetwork200ResponseNetwork) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetwork200ResponseNetwork) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetwork200ResponseNetwork) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetNetwork200ResponseNetwork) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetNetwork200ResponseNetwork) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalId

`func (o *GetNetwork200ResponseNetwork) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetNetwork200ResponseNetwork) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetNetwork200ResponseNetwork) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetNetwork200ResponseNetwork) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetNetwork200ResponseNetwork) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetNetwork200ResponseNetwork) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInternalId

`func (o *GetNetwork200ResponseNetwork) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetNetwork200ResponseNetwork) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetNetwork200ResponseNetwork) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetNetwork200ResponseNetwork) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *GetNetwork200ResponseNetwork) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *GetNetwork200ResponseNetwork) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetUniqueId

`func (o *GetNetwork200ResponseNetwork) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *GetNetwork200ResponseNetwork) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *GetNetwork200ResponseNetwork) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *GetNetwork200ResponseNetwork) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### SetUniqueIdNil

`func (o *GetNetwork200ResponseNetwork) SetUniqueIdNil(b bool)`

 SetUniqueIdNil sets the value for UniqueId to be an explicit nil

### UnsetUniqueId
`func (o *GetNetwork200ResponseNetwork) UnsetUniqueId()`

UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
### GetExternalType

`func (o *GetNetwork200ResponseNetwork) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *GetNetwork200ResponseNetwork) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *GetNetwork200ResponseNetwork) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *GetNetwork200ResponseNetwork) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### GetRefUrl

`func (o *GetNetwork200ResponseNetwork) GetRefUrl() string`

GetRefUrl returns the RefUrl field if non-nil, zero value otherwise.

### GetRefUrlOk

`func (o *GetNetwork200ResponseNetwork) GetRefUrlOk() (*string, bool)`

GetRefUrlOk returns a tuple with the RefUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUrl

`func (o *GetNetwork200ResponseNetwork) SetRefUrl(v string)`

SetRefUrl sets RefUrl field to given value.

### HasRefUrl

`func (o *GetNetwork200ResponseNetwork) HasRefUrl() bool`

HasRefUrl returns a boolean if a field has been set.

### SetRefUrlNil

`func (o *GetNetwork200ResponseNetwork) SetRefUrlNil(b bool)`

 SetRefUrlNil sets the value for RefUrl to be an explicit nil

### UnsetRefUrl
`func (o *GetNetwork200ResponseNetwork) UnsetRefUrl()`

UnsetRefUrl ensures that no value is present for RefUrl, not even an explicit nil
### GetRefType

`func (o *GetNetwork200ResponseNetwork) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetNetwork200ResponseNetwork) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetNetwork200ResponseNetwork) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetNetwork200ResponseNetwork) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *GetNetwork200ResponseNetwork) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetNetwork200ResponseNetwork) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetNetwork200ResponseNetwork) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetNetwork200ResponseNetwork) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetVlanId

`func (o *GetNetwork200ResponseNetwork) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *GetNetwork200ResponseNetwork) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *GetNetwork200ResponseNetwork) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *GetNetwork200ResponseNetwork) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### SetVlanIdNil

`func (o *GetNetwork200ResponseNetwork) SetVlanIdNil(b bool)`

 SetVlanIdNil sets the value for VlanId to be an explicit nil

### UnsetVlanId
`func (o *GetNetwork200ResponseNetwork) UnsetVlanId()`

UnsetVlanId ensures that no value is present for VlanId, not even an explicit nil
### GetVswitchName

`func (o *GetNetwork200ResponseNetwork) GetVswitchName() string`

GetVswitchName returns the VswitchName field if non-nil, zero value otherwise.

### GetVswitchNameOk

`func (o *GetNetwork200ResponseNetwork) GetVswitchNameOk() (*string, bool)`

GetVswitchNameOk returns a tuple with the VswitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchName

`func (o *GetNetwork200ResponseNetwork) SetVswitchName(v string)`

SetVswitchName sets VswitchName field to given value.

### HasVswitchName

`func (o *GetNetwork200ResponseNetwork) HasVswitchName() bool`

HasVswitchName returns a boolean if a field has been set.

### SetVswitchNameNil

`func (o *GetNetwork200ResponseNetwork) SetVswitchNameNil(b bool)`

 SetVswitchNameNil sets the value for VswitchName to be an explicit nil

### UnsetVswitchName
`func (o *GetNetwork200ResponseNetwork) UnsetVswitchName()`

UnsetVswitchName ensures that no value is present for VswitchName, not even an explicit nil
### GetDhcpServer

`func (o *GetNetwork200ResponseNetwork) GetDhcpServer() bool`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *GetNetwork200ResponseNetwork) GetDhcpServerOk() (*bool, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *GetNetwork200ResponseNetwork) SetDhcpServer(v bool)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *GetNetwork200ResponseNetwork) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetDhcpIp

`func (o *GetNetwork200ResponseNetwork) GetDhcpIp() string`

GetDhcpIp returns the DhcpIp field if non-nil, zero value otherwise.

### GetDhcpIpOk

`func (o *GetNetwork200ResponseNetwork) GetDhcpIpOk() (*string, bool)`

GetDhcpIpOk returns a tuple with the DhcpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpIp

`func (o *GetNetwork200ResponseNetwork) SetDhcpIp(v string)`

SetDhcpIp sets DhcpIp field to given value.

### HasDhcpIp

`func (o *GetNetwork200ResponseNetwork) HasDhcpIp() bool`

HasDhcpIp returns a boolean if a field has been set.

### SetDhcpIpNil

`func (o *GetNetwork200ResponseNetwork) SetDhcpIpNil(b bool)`

 SetDhcpIpNil sets the value for DhcpIp to be an explicit nil

### UnsetDhcpIp
`func (o *GetNetwork200ResponseNetwork) UnsetDhcpIp()`

UnsetDhcpIp ensures that no value is present for DhcpIp, not even an explicit nil
### GetDhcpServerIPv6

`func (o *GetNetwork200ResponseNetwork) GetDhcpServerIPv6() bool`

GetDhcpServerIPv6 returns the DhcpServerIPv6 field if non-nil, zero value otherwise.

### GetDhcpServerIPv6Ok

`func (o *GetNetwork200ResponseNetwork) GetDhcpServerIPv6Ok() (*bool, bool)`

GetDhcpServerIPv6Ok returns a tuple with the DhcpServerIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerIPv6

`func (o *GetNetwork200ResponseNetwork) SetDhcpServerIPv6(v bool)`

SetDhcpServerIPv6 sets DhcpServerIPv6 field to given value.

### HasDhcpServerIPv6

`func (o *GetNetwork200ResponseNetwork) HasDhcpServerIPv6() bool`

HasDhcpServerIPv6 returns a boolean if a field has been set.

### GetGateway

`func (o *GetNetwork200ResponseNetwork) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GetNetwork200ResponseNetwork) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GetNetwork200ResponseNetwork) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GetNetwork200ResponseNetwork) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *GetNetwork200ResponseNetwork) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *GetNetwork200ResponseNetwork) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetNetmask

`func (o *GetNetwork200ResponseNetwork) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *GetNetwork200ResponseNetwork) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *GetNetwork200ResponseNetwork) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *GetNetwork200ResponseNetwork) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### SetNetmaskNil

`func (o *GetNetwork200ResponseNetwork) SetNetmaskNil(b bool)`

 SetNetmaskNil sets the value for Netmask to be an explicit nil

### UnsetNetmask
`func (o *GetNetwork200ResponseNetwork) UnsetNetmask()`

UnsetNetmask ensures that no value is present for Netmask, not even an explicit nil
### GetBroadcast

`func (o *GetNetwork200ResponseNetwork) GetBroadcast() string`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *GetNetwork200ResponseNetwork) GetBroadcastOk() (*string, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *GetNetwork200ResponseNetwork) SetBroadcast(v string)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *GetNetwork200ResponseNetwork) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.

### SetBroadcastNil

`func (o *GetNetwork200ResponseNetwork) SetBroadcastNil(b bool)`

 SetBroadcastNil sets the value for Broadcast to be an explicit nil

### UnsetBroadcast
`func (o *GetNetwork200ResponseNetwork) UnsetBroadcast()`

UnsetBroadcast ensures that no value is present for Broadcast, not even an explicit nil
### GetSubnetAddress

`func (o *GetNetwork200ResponseNetwork) GetSubnetAddress() string`

GetSubnetAddress returns the SubnetAddress field if non-nil, zero value otherwise.

### GetSubnetAddressOk

`func (o *GetNetwork200ResponseNetwork) GetSubnetAddressOk() (*string, bool)`

GetSubnetAddressOk returns a tuple with the SubnetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAddress

`func (o *GetNetwork200ResponseNetwork) SetSubnetAddress(v string)`

SetSubnetAddress sets SubnetAddress field to given value.

### HasSubnetAddress

`func (o *GetNetwork200ResponseNetwork) HasSubnetAddress() bool`

HasSubnetAddress returns a boolean if a field has been set.

### SetSubnetAddressNil

`func (o *GetNetwork200ResponseNetwork) SetSubnetAddressNil(b bool)`

 SetSubnetAddressNil sets the value for SubnetAddress to be an explicit nil

### UnsetSubnetAddress
`func (o *GetNetwork200ResponseNetwork) UnsetSubnetAddress()`

UnsetSubnetAddress ensures that no value is present for SubnetAddress, not even an explicit nil
### GetDnsPrimary

`func (o *GetNetwork200ResponseNetwork) GetDnsPrimary() string`

GetDnsPrimary returns the DnsPrimary field if non-nil, zero value otherwise.

### GetDnsPrimaryOk

`func (o *GetNetwork200ResponseNetwork) GetDnsPrimaryOk() (*string, bool)`

GetDnsPrimaryOk returns a tuple with the DnsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrimary

`func (o *GetNetwork200ResponseNetwork) SetDnsPrimary(v string)`

SetDnsPrimary sets DnsPrimary field to given value.

### HasDnsPrimary

`func (o *GetNetwork200ResponseNetwork) HasDnsPrimary() bool`

HasDnsPrimary returns a boolean if a field has been set.

### SetDnsPrimaryNil

`func (o *GetNetwork200ResponseNetwork) SetDnsPrimaryNil(b bool)`

 SetDnsPrimaryNil sets the value for DnsPrimary to be an explicit nil

### UnsetDnsPrimary
`func (o *GetNetwork200ResponseNetwork) UnsetDnsPrimary()`

UnsetDnsPrimary ensures that no value is present for DnsPrimary, not even an explicit nil
### GetDnsSecondary

`func (o *GetNetwork200ResponseNetwork) GetDnsSecondary() string`

GetDnsSecondary returns the DnsSecondary field if non-nil, zero value otherwise.

### GetDnsSecondaryOk

`func (o *GetNetwork200ResponseNetwork) GetDnsSecondaryOk() (*string, bool)`

GetDnsSecondaryOk returns a tuple with the DnsSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSecondary

`func (o *GetNetwork200ResponseNetwork) SetDnsSecondary(v string)`

SetDnsSecondary sets DnsSecondary field to given value.

### HasDnsSecondary

`func (o *GetNetwork200ResponseNetwork) HasDnsSecondary() bool`

HasDnsSecondary returns a boolean if a field has been set.

### SetDnsSecondaryNil

`func (o *GetNetwork200ResponseNetwork) SetDnsSecondaryNil(b bool)`

 SetDnsSecondaryNil sets the value for DnsSecondary to be an explicit nil

### UnsetDnsSecondary
`func (o *GetNetwork200ResponseNetwork) UnsetDnsSecondary()`

UnsetDnsSecondary ensures that no value is present for DnsSecondary, not even an explicit nil
### GetCidr

`func (o *GetNetwork200ResponseNetwork) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *GetNetwork200ResponseNetwork) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *GetNetwork200ResponseNetwork) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *GetNetwork200ResponseNetwork) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### SetCidrNil

`func (o *GetNetwork200ResponseNetwork) SetCidrNil(b bool)`

 SetCidrNil sets the value for Cidr to be an explicit nil

### UnsetCidr
`func (o *GetNetwork200ResponseNetwork) UnsetCidr()`

UnsetCidr ensures that no value is present for Cidr, not even an explicit nil
### GetGatewayIPv6

`func (o *GetNetwork200ResponseNetwork) GetGatewayIPv6() string`

GetGatewayIPv6 returns the GatewayIPv6 field if non-nil, zero value otherwise.

### GetGatewayIPv6Ok

`func (o *GetNetwork200ResponseNetwork) GetGatewayIPv6Ok() (*string, bool)`

GetGatewayIPv6Ok returns a tuple with the GatewayIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIPv6

`func (o *GetNetwork200ResponseNetwork) SetGatewayIPv6(v string)`

SetGatewayIPv6 sets GatewayIPv6 field to given value.

### HasGatewayIPv6

`func (o *GetNetwork200ResponseNetwork) HasGatewayIPv6() bool`

HasGatewayIPv6 returns a boolean if a field has been set.

### SetGatewayIPv6Nil

`func (o *GetNetwork200ResponseNetwork) SetGatewayIPv6Nil(b bool)`

 SetGatewayIPv6Nil sets the value for GatewayIPv6 to be an explicit nil

### UnsetGatewayIPv6
`func (o *GetNetwork200ResponseNetwork) UnsetGatewayIPv6()`

UnsetGatewayIPv6 ensures that no value is present for GatewayIPv6, not even an explicit nil
### GetNetmaskIPv6

`func (o *GetNetwork200ResponseNetwork) GetNetmaskIPv6() string`

GetNetmaskIPv6 returns the NetmaskIPv6 field if non-nil, zero value otherwise.

### GetNetmaskIPv6Ok

`func (o *GetNetwork200ResponseNetwork) GetNetmaskIPv6Ok() (*string, bool)`

GetNetmaskIPv6Ok returns a tuple with the NetmaskIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmaskIPv6

`func (o *GetNetwork200ResponseNetwork) SetNetmaskIPv6(v string)`

SetNetmaskIPv6 sets NetmaskIPv6 field to given value.

### HasNetmaskIPv6

`func (o *GetNetwork200ResponseNetwork) HasNetmaskIPv6() bool`

HasNetmaskIPv6 returns a boolean if a field has been set.

### SetNetmaskIPv6Nil

`func (o *GetNetwork200ResponseNetwork) SetNetmaskIPv6Nil(b bool)`

 SetNetmaskIPv6Nil sets the value for NetmaskIPv6 to be an explicit nil

### UnsetNetmaskIPv6
`func (o *GetNetwork200ResponseNetwork) UnsetNetmaskIPv6()`

UnsetNetmaskIPv6 ensures that no value is present for NetmaskIPv6, not even an explicit nil
### GetDnsPrimaryIPv6

`func (o *GetNetwork200ResponseNetwork) GetDnsPrimaryIPv6() string`

GetDnsPrimaryIPv6 returns the DnsPrimaryIPv6 field if non-nil, zero value otherwise.

### GetDnsPrimaryIPv6Ok

`func (o *GetNetwork200ResponseNetwork) GetDnsPrimaryIPv6Ok() (*string, bool)`

GetDnsPrimaryIPv6Ok returns a tuple with the DnsPrimaryIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrimaryIPv6

`func (o *GetNetwork200ResponseNetwork) SetDnsPrimaryIPv6(v string)`

SetDnsPrimaryIPv6 sets DnsPrimaryIPv6 field to given value.

### HasDnsPrimaryIPv6

`func (o *GetNetwork200ResponseNetwork) HasDnsPrimaryIPv6() bool`

HasDnsPrimaryIPv6 returns a boolean if a field has been set.

### SetDnsPrimaryIPv6Nil

`func (o *GetNetwork200ResponseNetwork) SetDnsPrimaryIPv6Nil(b bool)`

 SetDnsPrimaryIPv6Nil sets the value for DnsPrimaryIPv6 to be an explicit nil

### UnsetDnsPrimaryIPv6
`func (o *GetNetwork200ResponseNetwork) UnsetDnsPrimaryIPv6()`

UnsetDnsPrimaryIPv6 ensures that no value is present for DnsPrimaryIPv6, not even an explicit nil
### GetDnsSecondaryIPv6

`func (o *GetNetwork200ResponseNetwork) GetDnsSecondaryIPv6() string`

GetDnsSecondaryIPv6 returns the DnsSecondaryIPv6 field if non-nil, zero value otherwise.

### GetDnsSecondaryIPv6Ok

`func (o *GetNetwork200ResponseNetwork) GetDnsSecondaryIPv6Ok() (*string, bool)`

GetDnsSecondaryIPv6Ok returns a tuple with the DnsSecondaryIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSecondaryIPv6

`func (o *GetNetwork200ResponseNetwork) SetDnsSecondaryIPv6(v string)`

SetDnsSecondaryIPv6 sets DnsSecondaryIPv6 field to given value.

### HasDnsSecondaryIPv6

`func (o *GetNetwork200ResponseNetwork) HasDnsSecondaryIPv6() bool`

HasDnsSecondaryIPv6 returns a boolean if a field has been set.

### SetDnsSecondaryIPv6Nil

`func (o *GetNetwork200ResponseNetwork) SetDnsSecondaryIPv6Nil(b bool)`

 SetDnsSecondaryIPv6Nil sets the value for DnsSecondaryIPv6 to be an explicit nil

### UnsetDnsSecondaryIPv6
`func (o *GetNetwork200ResponseNetwork) UnsetDnsSecondaryIPv6()`

UnsetDnsSecondaryIPv6 ensures that no value is present for DnsSecondaryIPv6, not even an explicit nil
### GetCidrIPv6

`func (o *GetNetwork200ResponseNetwork) GetCidrIPv6() string`

GetCidrIPv6 returns the CidrIPv6 field if non-nil, zero value otherwise.

### GetCidrIPv6Ok

`func (o *GetNetwork200ResponseNetwork) GetCidrIPv6Ok() (*string, bool)`

GetCidrIPv6Ok returns a tuple with the CidrIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrIPv6

`func (o *GetNetwork200ResponseNetwork) SetCidrIPv6(v string)`

SetCidrIPv6 sets CidrIPv6 field to given value.

### HasCidrIPv6

`func (o *GetNetwork200ResponseNetwork) HasCidrIPv6() bool`

HasCidrIPv6 returns a boolean if a field has been set.

### SetCidrIPv6Nil

`func (o *GetNetwork200ResponseNetwork) SetCidrIPv6Nil(b bool)`

 SetCidrIPv6Nil sets the value for CidrIPv6 to be an explicit nil

### UnsetCidrIPv6
`func (o *GetNetwork200ResponseNetwork) UnsetCidrIPv6()`

UnsetCidrIPv6 ensures that no value is present for CidrIPv6, not even an explicit nil
### GetTftpServer

`func (o *GetNetwork200ResponseNetwork) GetTftpServer() string`

GetTftpServer returns the TftpServer field if non-nil, zero value otherwise.

### GetTftpServerOk

`func (o *GetNetwork200ResponseNetwork) GetTftpServerOk() (*string, bool)`

GetTftpServerOk returns a tuple with the TftpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTftpServer

`func (o *GetNetwork200ResponseNetwork) SetTftpServer(v string)`

SetTftpServer sets TftpServer field to given value.

### HasTftpServer

`func (o *GetNetwork200ResponseNetwork) HasTftpServer() bool`

HasTftpServer returns a boolean if a field has been set.

### SetTftpServerNil

`func (o *GetNetwork200ResponseNetwork) SetTftpServerNil(b bool)`

 SetTftpServerNil sets the value for TftpServer to be an explicit nil

### UnsetTftpServer
`func (o *GetNetwork200ResponseNetwork) UnsetTftpServer()`

UnsetTftpServer ensures that no value is present for TftpServer, not even an explicit nil
### GetBootFile

`func (o *GetNetwork200ResponseNetwork) GetBootFile() string`

GetBootFile returns the BootFile field if non-nil, zero value otherwise.

### GetBootFileOk

`func (o *GetNetwork200ResponseNetwork) GetBootFileOk() (*string, bool)`

GetBootFileOk returns a tuple with the BootFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootFile

`func (o *GetNetwork200ResponseNetwork) SetBootFile(v string)`

SetBootFile sets BootFile field to given value.

### HasBootFile

`func (o *GetNetwork200ResponseNetwork) HasBootFile() bool`

HasBootFile returns a boolean if a field has been set.

### SetBootFileNil

`func (o *GetNetwork200ResponseNetwork) SetBootFileNil(b bool)`

 SetBootFileNil sets the value for BootFile to be an explicit nil

### UnsetBootFile
`func (o *GetNetwork200ResponseNetwork) UnsetBootFile()`

UnsetBootFile ensures that no value is present for BootFile, not even an explicit nil
### GetSwitchId

`func (o *GetNetwork200ResponseNetwork) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *GetNetwork200ResponseNetwork) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *GetNetwork200ResponseNetwork) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *GetNetwork200ResponseNetwork) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### SetSwitchIdNil

`func (o *GetNetwork200ResponseNetwork) SetSwitchIdNil(b bool)`

 SetSwitchIdNil sets the value for SwitchId to be an explicit nil

### UnsetSwitchId
`func (o *GetNetwork200ResponseNetwork) UnsetSwitchId()`

UnsetSwitchId ensures that no value is present for SwitchId, not even an explicit nil
### GetFabricId

`func (o *GetNetwork200ResponseNetwork) GetFabricId() string`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *GetNetwork200ResponseNetwork) GetFabricIdOk() (*string, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *GetNetwork200ResponseNetwork) SetFabricId(v string)`

SetFabricId sets FabricId field to given value.

### HasFabricId

`func (o *GetNetwork200ResponseNetwork) HasFabricId() bool`

HasFabricId returns a boolean if a field has been set.

### SetFabricIdNil

`func (o *GetNetwork200ResponseNetwork) SetFabricIdNil(b bool)`

 SetFabricIdNil sets the value for FabricId to be an explicit nil

### UnsetFabricId
`func (o *GetNetwork200ResponseNetwork) UnsetFabricId()`

UnsetFabricId ensures that no value is present for FabricId, not even an explicit nil
### GetNetworkRole

`func (o *GetNetwork200ResponseNetwork) GetNetworkRole() string`

GetNetworkRole returns the NetworkRole field if non-nil, zero value otherwise.

### GetNetworkRoleOk

`func (o *GetNetwork200ResponseNetwork) GetNetworkRoleOk() (*string, bool)`

GetNetworkRoleOk returns a tuple with the NetworkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkRole

`func (o *GetNetwork200ResponseNetwork) SetNetworkRole(v string)`

SetNetworkRole sets NetworkRole field to given value.

### HasNetworkRole

`func (o *GetNetwork200ResponseNetwork) HasNetworkRole() bool`

HasNetworkRole returns a boolean if a field has been set.

### SetNetworkRoleNil

`func (o *GetNetwork200ResponseNetwork) SetNetworkRoleNil(b bool)`

 SetNetworkRoleNil sets the value for NetworkRole to be an explicit nil

### UnsetNetworkRole
`func (o *GetNetwork200ResponseNetwork) UnsetNetworkRole()`

UnsetNetworkRole ensures that no value is present for NetworkRole, not even an explicit nil
### GetStatus

`func (o *GetNetwork200ResponseNetwork) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetNetwork200ResponseNetwork) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetNetwork200ResponseNetwork) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetNetwork200ResponseNetwork) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GetNetwork200ResponseNetwork) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GetNetwork200ResponseNetwork) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetAvailabilityZone

`func (o *GetNetwork200ResponseNetwork) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *GetNetwork200ResponseNetwork) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *GetNetwork200ResponseNetwork) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *GetNetwork200ResponseNetwork) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *GetNetwork200ResponseNetwork) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *GetNetwork200ResponseNetwork) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetPool

`func (o *GetNetwork200ResponseNetwork) GetPool() GetNetwork200ResponseNetworkPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *GetNetwork200ResponseNetwork) GetPoolOk() (*GetNetwork200ResponseNetworkPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *GetNetwork200ResponseNetwork) SetPool(v GetNetwork200ResponseNetworkPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *GetNetwork200ResponseNetwork) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolIPv6

`func (o *GetNetwork200ResponseNetwork) GetPoolIPv6() GetNetwork200ResponseNetworkPoolIPv6`

GetPoolIPv6 returns the PoolIPv6 field if non-nil, zero value otherwise.

### GetPoolIPv6Ok

`func (o *GetNetwork200ResponseNetwork) GetPoolIPv6Ok() (*GetNetwork200ResponseNetworkPoolIPv6, bool)`

GetPoolIPv6Ok returns a tuple with the PoolIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolIPv6

`func (o *GetNetwork200ResponseNetwork) SetPoolIPv6(v GetNetwork200ResponseNetworkPoolIPv6)`

SetPoolIPv6 sets PoolIPv6 field to given value.

### HasPoolIPv6

`func (o *GetNetwork200ResponseNetwork) HasPoolIPv6() bool`

HasPoolIPv6 returns a boolean if a field has been set.

### GetNetworkProxy

`func (o *GetNetwork200ResponseNetwork) GetNetworkProxy() GetNetwork200ResponseNetworkNetworkProxy`

GetNetworkProxy returns the NetworkProxy field if non-nil, zero value otherwise.

### GetNetworkProxyOk

`func (o *GetNetwork200ResponseNetwork) GetNetworkProxyOk() (*GetNetwork200ResponseNetworkNetworkProxy, bool)`

GetNetworkProxyOk returns a tuple with the NetworkProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProxy

`func (o *GetNetwork200ResponseNetwork) SetNetworkProxy(v GetNetwork200ResponseNetworkNetworkProxy)`

SetNetworkProxy sets NetworkProxy field to given value.

### HasNetworkProxy

`func (o *GetNetwork200ResponseNetwork) HasNetworkProxy() bool`

HasNetworkProxy returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *GetNetwork200ResponseNetwork) GetNetworkDomain() GetNetwork200ResponseNetworkNetworkDomain`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *GetNetwork200ResponseNetwork) GetNetworkDomainOk() (*GetNetwork200ResponseNetworkNetworkDomain, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *GetNetwork200ResponseNetwork) SetNetworkDomain(v GetNetwork200ResponseNetworkNetworkDomain)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *GetNetwork200ResponseNetwork) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetSearchDomains

`func (o *GetNetwork200ResponseNetwork) GetSearchDomains() string`

GetSearchDomains returns the SearchDomains field if non-nil, zero value otherwise.

### GetSearchDomainsOk

`func (o *GetNetwork200ResponseNetwork) GetSearchDomainsOk() (*string, bool)`

GetSearchDomainsOk returns a tuple with the SearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomains

`func (o *GetNetwork200ResponseNetwork) SetSearchDomains(v string)`

SetSearchDomains sets SearchDomains field to given value.

### HasSearchDomains

`func (o *GetNetwork200ResponseNetwork) HasSearchDomains() bool`

HasSearchDomains returns a boolean if a field has been set.

### SetSearchDomainsNil

`func (o *GetNetwork200ResponseNetwork) SetSearchDomainsNil(b bool)`

 SetSearchDomainsNil sets the value for SearchDomains to be an explicit nil

### UnsetSearchDomains
`func (o *GetNetwork200ResponseNetwork) UnsetSearchDomains()`

UnsetSearchDomains ensures that no value is present for SearchDomains, not even an explicit nil
### GetPrefixLength

`func (o *GetNetwork200ResponseNetwork) GetPrefixLength() string`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *GetNetwork200ResponseNetwork) GetPrefixLengthOk() (*string, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *GetNetwork200ResponseNetwork) SetPrefixLength(v string)`

SetPrefixLength sets PrefixLength field to given value.

### HasPrefixLength

`func (o *GetNetwork200ResponseNetwork) HasPrefixLength() bool`

HasPrefixLength returns a boolean if a field has been set.

### SetPrefixLengthNil

`func (o *GetNetwork200ResponseNetwork) SetPrefixLengthNil(b bool)`

 SetPrefixLengthNil sets the value for PrefixLength to be an explicit nil

### UnsetPrefixLength
`func (o *GetNetwork200ResponseNetwork) UnsetPrefixLength()`

UnsetPrefixLength ensures that no value is present for PrefixLength, not even an explicit nil
### GetVisibility

`func (o *GetNetwork200ResponseNetwork) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetNetwork200ResponseNetwork) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetNetwork200ResponseNetwork) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetNetwork200ResponseNetwork) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEnableAdmin

`func (o *GetNetwork200ResponseNetwork) GetEnableAdmin() bool`

GetEnableAdmin returns the EnableAdmin field if non-nil, zero value otherwise.

### GetEnableAdminOk

`func (o *GetNetwork200ResponseNetwork) GetEnableAdminOk() (*bool, bool)`

GetEnableAdminOk returns a tuple with the EnableAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdmin

`func (o *GetNetwork200ResponseNetwork) SetEnableAdmin(v bool)`

SetEnableAdmin sets EnableAdmin field to given value.

### HasEnableAdmin

`func (o *GetNetwork200ResponseNetwork) HasEnableAdmin() bool`

HasEnableAdmin returns a boolean if a field has been set.

### GetActive

`func (o *GetNetwork200ResponseNetwork) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetNetwork200ResponseNetwork) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetNetwork200ResponseNetwork) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetNetwork200ResponseNetwork) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDefaultNetwork

`func (o *GetNetwork200ResponseNetwork) GetDefaultNetwork() bool`

GetDefaultNetwork returns the DefaultNetwork field if non-nil, zero value otherwise.

### GetDefaultNetworkOk

`func (o *GetNetwork200ResponseNetwork) GetDefaultNetworkOk() (*bool, bool)`

GetDefaultNetworkOk returns a tuple with the DefaultNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetwork

`func (o *GetNetwork200ResponseNetwork) SetDefaultNetwork(v bool)`

SetDefaultNetwork sets DefaultNetwork field to given value.

### HasDefaultNetwork

`func (o *GetNetwork200ResponseNetwork) HasDefaultNetwork() bool`

HasDefaultNetwork returns a boolean if a field has been set.

### GetAssignPublicIp

`func (o *GetNetwork200ResponseNetwork) GetAssignPublicIp() bool`

GetAssignPublicIp returns the AssignPublicIp field if non-nil, zero value otherwise.

### GetAssignPublicIpOk

`func (o *GetNetwork200ResponseNetwork) GetAssignPublicIpOk() (*bool, bool)`

GetAssignPublicIpOk returns a tuple with the AssignPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignPublicIp

`func (o *GetNetwork200ResponseNetwork) SetAssignPublicIp(v bool)`

SetAssignPublicIp sets AssignPublicIp field to given value.

### HasAssignPublicIp

`func (o *GetNetwork200ResponseNetwork) HasAssignPublicIp() bool`

HasAssignPublicIp returns a boolean if a field has been set.

### GetNoProxy

`func (o *GetNetwork200ResponseNetwork) GetNoProxy() string`

GetNoProxy returns the NoProxy field if non-nil, zero value otherwise.

### GetNoProxyOk

`func (o *GetNetwork200ResponseNetwork) GetNoProxyOk() (*string, bool)`

GetNoProxyOk returns a tuple with the NoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxy

`func (o *GetNetwork200ResponseNetwork) SetNoProxy(v string)`

SetNoProxy sets NoProxy field to given value.

### HasNoProxy

`func (o *GetNetwork200ResponseNetwork) HasNoProxy() bool`

HasNoProxy returns a boolean if a field has been set.

### SetNoProxyNil

`func (o *GetNetwork200ResponseNetwork) SetNoProxyNil(b bool)`

 SetNoProxyNil sets the value for NoProxy to be an explicit nil

### UnsetNoProxy
`func (o *GetNetwork200ResponseNetwork) UnsetNoProxy()`

UnsetNoProxy ensures that no value is present for NoProxy, not even an explicit nil
### GetApplianceUrlProxyBypass

`func (o *GetNetwork200ResponseNetwork) GetApplianceUrlProxyBypass() bool`

GetApplianceUrlProxyBypass returns the ApplianceUrlProxyBypass field if non-nil, zero value otherwise.

### GetApplianceUrlProxyBypassOk

`func (o *GetNetwork200ResponseNetwork) GetApplianceUrlProxyBypassOk() (*bool, bool)`

GetApplianceUrlProxyBypassOk returns a tuple with the ApplianceUrlProxyBypass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrlProxyBypass

`func (o *GetNetwork200ResponseNetwork) SetApplianceUrlProxyBypass(v bool)`

SetApplianceUrlProxyBypass sets ApplianceUrlProxyBypass field to given value.

### HasApplianceUrlProxyBypass

`func (o *GetNetwork200ResponseNetwork) HasApplianceUrlProxyBypass() bool`

HasApplianceUrlProxyBypass returns a boolean if a field has been set.

### GetZonePool

`func (o *GetNetwork200ResponseNetwork) GetZonePool() GetNetwork200ResponseNetworkZonePool`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *GetNetwork200ResponseNetwork) GetZonePoolOk() (*GetNetwork200ResponseNetworkZonePool, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *GetNetwork200ResponseNetwork) SetZonePool(v GetNetwork200ResponseNetworkZonePool)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *GetNetwork200ResponseNetwork) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetAllowStaticOverride

`func (o *GetNetwork200ResponseNetwork) GetAllowStaticOverride() bool`

GetAllowStaticOverride returns the AllowStaticOverride field if non-nil, zero value otherwise.

### GetAllowStaticOverrideOk

`func (o *GetNetwork200ResponseNetwork) GetAllowStaticOverrideOk() (*bool, bool)`

GetAllowStaticOverrideOk returns a tuple with the AllowStaticOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowStaticOverride

`func (o *GetNetwork200ResponseNetwork) SetAllowStaticOverride(v bool)`

SetAllowStaticOverride sets AllowStaticOverride field to given value.

### HasAllowStaticOverride

`func (o *GetNetwork200ResponseNetwork) HasAllowStaticOverride() bool`

HasAllowStaticOverride returns a boolean if a field has been set.

### GetConfig

`func (o *GetNetwork200ResponseNetwork) GetConfig() GetNetwork200ResponseNetworkConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetNetwork200ResponseNetwork) GetConfigOk() (*GetNetwork200ResponseNetworkConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetNetwork200ResponseNetwork) SetConfig(v GetNetwork200ResponseNetworkConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetNetwork200ResponseNetwork) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTenants

`func (o *GetNetwork200ResponseNetwork) GetTenants() []GetNetwork200ResponseNetworkTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GetNetwork200ResponseNetwork) GetTenantsOk() (*[]GetNetwork200ResponseNetworkTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GetNetwork200ResponseNetwork) SetTenants(v []GetNetwork200ResponseNetworkTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *GetNetwork200ResponseNetwork) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *GetNetwork200ResponseNetwork) GetResourcePermission() GetNetwork200ResponseNetworkResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *GetNetwork200ResponseNetwork) GetResourcePermissionOk() (*GetNetwork200ResponseNetworkResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *GetNetwork200ResponseNetwork) SetResourcePermission(v GetNetwork200ResponseNetworkResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *GetNetwork200ResponseNetwork) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


