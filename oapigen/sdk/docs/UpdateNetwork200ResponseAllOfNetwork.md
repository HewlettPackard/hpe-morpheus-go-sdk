# UpdateNetwork200ResponseAllOfNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Network ID | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**DisplayName** | Pointer to **NullableString** | Network Display Name | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Group** | Pointer to [**UpdateNetwork200ResponseAllOfNetworkGroup**](UpdateNetwork200ResponseAllOfNetworkGroup.md) |  | [optional] 
**Zone** | Pointer to [**UpdateNetwork200ResponseAllOfNetworkZone**](UpdateNetwork200ResponseAllOfNetworkZone.md) |  | [optional] 
**Type** | Pointer to [**UpdateNetwork200ResponseAllOfNetworkType**](UpdateNetwork200ResponseAllOfNetworkType.md) |  | [optional] 
**Owner** | Pointer to [**UpdateNetwork200ResponseAllOfNetworkOwner**](UpdateNetwork200ResponseAllOfNetworkOwner.md) |  | [optional] 
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
**Pool** | Pointer to [**UpdateNetwork200ResponseAllOfNetworkPool**](UpdateNetwork200ResponseAllOfNetworkPool.md) |  | [optional] 
**PoolIPv6** | Pointer to [**UpdateNetwork200ResponseAllOfNetworkPoolIPv6**](UpdateNetwork200ResponseAllOfNetworkPoolIPv6.md) |  | [optional] 
**NetworkProxy** | Pointer to [**UpdateNetwork200ResponseAllOfNetworkNetworkProxy**](UpdateNetwork200ResponseAllOfNetworkNetworkProxy.md) |  | [optional] 
**NetworkDomain** | Pointer to [**UpdateNetwork200ResponseAllOfNetworkNetworkDomain**](UpdateNetwork200ResponseAllOfNetworkNetworkDomain.md) |  | [optional] 
**SearchDomains** | Pointer to **NullableString** |  | [optional] 
**PrefixLength** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**EnableAdmin** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**DefaultNetwork** | Pointer to **bool** |  | [optional] 
**AssignPublicIp** | Pointer to **bool** |  | [optional] 
**NoProxy** | Pointer to **NullableString** |  | [optional] 
**ApplianceUrlProxyBypass** | Pointer to **bool** |  | [optional] 
**ZonePool** | Pointer to [**UpdateNetwork200ResponseAllOfNetworkZonePool**](UpdateNetwork200ResponseAllOfNetworkZonePool.md) |  | [optional] 
**AllowStaticOverride** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**UpdateNetwork200ResponseAllOfNetworkConfig**](UpdateNetwork200ResponseAllOfNetworkConfig.md) |  | [optional] 
**Tenants** | Pointer to [**[]UpdateNetwork200ResponseAllOfNetworkTenantsInner**](UpdateNetwork200ResponseAllOfNetworkTenantsInner.md) |  | [optional] 
**ResourcePermission** | Pointer to [**UpdateNetwork200ResponseAllOfNetworkResourcePermission**](UpdateNetwork200ResponseAllOfNetworkResourcePermission.md) |  | [optional] 

## Methods

### NewUpdateNetwork200ResponseAllOfNetwork

`func NewUpdateNetwork200ResponseAllOfNetwork() *UpdateNetwork200ResponseAllOfNetwork`

NewUpdateNetwork200ResponseAllOfNetwork instantiates a new UpdateNetwork200ResponseAllOfNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetwork200ResponseAllOfNetworkWithDefaults

`func NewUpdateNetwork200ResponseAllOfNetworkWithDefaults() *UpdateNetwork200ResponseAllOfNetwork`

NewUpdateNetwork200ResponseAllOfNetworkWithDefaults instantiates a new UpdateNetwork200ResponseAllOfNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLabels

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetGroup

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetGroup() UpdateNetwork200ResponseAllOfNetworkGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetGroupOk() (*UpdateNetwork200ResponseAllOfNetworkGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetGroup(v UpdateNetwork200ResponseAllOfNetworkGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetZone

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetZone() UpdateNetwork200ResponseAllOfNetworkZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetZoneOk() (*UpdateNetwork200ResponseAllOfNetworkZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetZone(v UpdateNetwork200ResponseAllOfNetworkZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetType

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetType() UpdateNetwork200ResponseAllOfNetworkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetTypeOk() (*UpdateNetwork200ResponseAllOfNetworkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetType(v UpdateNetwork200ResponseAllOfNetworkType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOwner

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetOwner() UpdateNetwork200ResponseAllOfNetworkOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetOwnerOk() (*UpdateNetwork200ResponseAllOfNetworkOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetOwner(v UpdateNetwork200ResponseAllOfNetworkOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCode

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetIpv4Enabled

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetIpv4Enabled() bool`

GetIpv4Enabled returns the Ipv4Enabled field if non-nil, zero value otherwise.

### GetIpv4EnabledOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetIpv4EnabledOk() (*bool, bool)`

GetIpv4EnabledOk returns a tuple with the Ipv4Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Enabled

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetIpv4Enabled(v bool)`

SetIpv4Enabled sets Ipv4Enabled field to given value.

### HasIpv4Enabled

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasIpv4Enabled() bool`

HasIpv4Enabled returns a boolean if a field has been set.

### GetIpv6Enabled

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetIpv6Enabled() bool`

GetIpv6Enabled returns the Ipv6Enabled field if non-nil, zero value otherwise.

### GetIpv6EnabledOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetIpv6EnabledOk() (*bool, bool)`

GetIpv6EnabledOk returns a tuple with the Ipv6Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Enabled

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetIpv6Enabled(v bool)`

SetIpv6Enabled sets Ipv6Enabled field to given value.

### HasIpv6Enabled

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasIpv6Enabled() bool`

HasIpv6Enabled returns a boolean if a field has been set.

### GetCategory

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetInterfaceName

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### SetInterfaceNameNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetInterfaceNameNil(b bool)`

 SetInterfaceNameNil sets the value for InterfaceName to be an explicit nil

### UnsetInterfaceName
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetInterfaceName()`

UnsetInterfaceName ensures that no value is present for InterfaceName, not even an explicit nil
### GetBridgeName

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetBridgeName() string`

GetBridgeName returns the BridgeName field if non-nil, zero value otherwise.

### GetBridgeNameOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetBridgeNameOk() (*string, bool)`

GetBridgeNameOk returns a tuple with the BridgeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeName

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetBridgeName(v string)`

SetBridgeName sets BridgeName field to given value.

### HasBridgeName

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasBridgeName() bool`

HasBridgeName returns a boolean if a field has been set.

### SetBridgeNameNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetBridgeNameNil(b bool)`

 SetBridgeNameNil sets the value for BridgeName to be an explicit nil

### UnsetBridgeName
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetBridgeName()`

UnsetBridgeName ensures that no value is present for BridgeName, not even an explicit nil
### GetBridgeInterface

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetBridgeInterface() string`

GetBridgeInterface returns the BridgeInterface field if non-nil, zero value otherwise.

### GetBridgeInterfaceOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetBridgeInterfaceOk() (*string, bool)`

GetBridgeInterfaceOk returns a tuple with the BridgeInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeInterface

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetBridgeInterface(v string)`

SetBridgeInterface sets BridgeInterface field to given value.

### HasBridgeInterface

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasBridgeInterface() bool`

HasBridgeInterface returns a boolean if a field has been set.

### SetBridgeInterfaceNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetBridgeInterfaceNil(b bool)`

 SetBridgeInterfaceNil sets the value for BridgeInterface to be an explicit nil

### UnsetBridgeInterface
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetBridgeInterface()`

UnsetBridgeInterface ensures that no value is present for BridgeInterface, not even an explicit nil
### GetDescription

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalId

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInternalId

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetUniqueId

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### SetUniqueIdNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetUniqueIdNil(b bool)`

 SetUniqueIdNil sets the value for UniqueId to be an explicit nil

### UnsetUniqueId
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetUniqueId()`

UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
### GetExternalType

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### GetRefUrl

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetRefUrl() string`

GetRefUrl returns the RefUrl field if non-nil, zero value otherwise.

### GetRefUrlOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetRefUrlOk() (*string, bool)`

GetRefUrlOk returns a tuple with the RefUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUrl

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetRefUrl(v string)`

SetRefUrl sets RefUrl field to given value.

### HasRefUrl

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasRefUrl() bool`

HasRefUrl returns a boolean if a field has been set.

### SetRefUrlNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetRefUrlNil(b bool)`

 SetRefUrlNil sets the value for RefUrl to be an explicit nil

### UnsetRefUrl
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetRefUrl()`

UnsetRefUrl ensures that no value is present for RefUrl, not even an explicit nil
### GetRefType

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetVlanId

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### SetVlanIdNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetVlanIdNil(b bool)`

 SetVlanIdNil sets the value for VlanId to be an explicit nil

### UnsetVlanId
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetVlanId()`

UnsetVlanId ensures that no value is present for VlanId, not even an explicit nil
### GetVswitchName

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetVswitchName() string`

GetVswitchName returns the VswitchName field if non-nil, zero value otherwise.

### GetVswitchNameOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetVswitchNameOk() (*string, bool)`

GetVswitchNameOk returns a tuple with the VswitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchName

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetVswitchName(v string)`

SetVswitchName sets VswitchName field to given value.

### HasVswitchName

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasVswitchName() bool`

HasVswitchName returns a boolean if a field has been set.

### SetVswitchNameNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetVswitchNameNil(b bool)`

 SetVswitchNameNil sets the value for VswitchName to be an explicit nil

### UnsetVswitchName
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetVswitchName()`

UnsetVswitchName ensures that no value is present for VswitchName, not even an explicit nil
### GetDhcpServer

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetDhcpServer() bool`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetDhcpServerOk() (*bool, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetDhcpServer(v bool)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetDhcpIp

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetDhcpIp() string`

GetDhcpIp returns the DhcpIp field if non-nil, zero value otherwise.

### GetDhcpIpOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetDhcpIpOk() (*string, bool)`

GetDhcpIpOk returns a tuple with the DhcpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpIp

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetDhcpIp(v string)`

SetDhcpIp sets DhcpIp field to given value.

### HasDhcpIp

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasDhcpIp() bool`

HasDhcpIp returns a boolean if a field has been set.

### SetDhcpIpNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetDhcpIpNil(b bool)`

 SetDhcpIpNil sets the value for DhcpIp to be an explicit nil

### UnsetDhcpIp
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetDhcpIp()`

UnsetDhcpIp ensures that no value is present for DhcpIp, not even an explicit nil
### GetDhcpServerIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetDhcpServerIPv6() bool`

GetDhcpServerIPv6 returns the DhcpServerIPv6 field if non-nil, zero value otherwise.

### GetDhcpServerIPv6Ok

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetDhcpServerIPv6Ok() (*bool, bool)`

GetDhcpServerIPv6Ok returns a tuple with the DhcpServerIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetDhcpServerIPv6(v bool)`

SetDhcpServerIPv6 sets DhcpServerIPv6 field to given value.

### HasDhcpServerIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasDhcpServerIPv6() bool`

HasDhcpServerIPv6 returns a boolean if a field has been set.

### GetGateway

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetNetmask

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### SetNetmaskNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetNetmaskNil(b bool)`

 SetNetmaskNil sets the value for Netmask to be an explicit nil

### UnsetNetmask
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetNetmask()`

UnsetNetmask ensures that no value is present for Netmask, not even an explicit nil
### GetBroadcast

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetBroadcast() string`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetBroadcastOk() (*string, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetBroadcast(v string)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.

### SetBroadcastNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetBroadcastNil(b bool)`

 SetBroadcastNil sets the value for Broadcast to be an explicit nil

### UnsetBroadcast
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetBroadcast()`

UnsetBroadcast ensures that no value is present for Broadcast, not even an explicit nil
### GetSubnetAddress

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetSubnetAddress() string`

GetSubnetAddress returns the SubnetAddress field if non-nil, zero value otherwise.

### GetSubnetAddressOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetSubnetAddressOk() (*string, bool)`

GetSubnetAddressOk returns a tuple with the SubnetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAddress

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetSubnetAddress(v string)`

SetSubnetAddress sets SubnetAddress field to given value.

### HasSubnetAddress

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasSubnetAddress() bool`

HasSubnetAddress returns a boolean if a field has been set.

### SetSubnetAddressNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetSubnetAddressNil(b bool)`

 SetSubnetAddressNil sets the value for SubnetAddress to be an explicit nil

### UnsetSubnetAddress
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetSubnetAddress()`

UnsetSubnetAddress ensures that no value is present for SubnetAddress, not even an explicit nil
### GetDnsPrimary

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetDnsPrimary() string`

GetDnsPrimary returns the DnsPrimary field if non-nil, zero value otherwise.

### GetDnsPrimaryOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetDnsPrimaryOk() (*string, bool)`

GetDnsPrimaryOk returns a tuple with the DnsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrimary

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetDnsPrimary(v string)`

SetDnsPrimary sets DnsPrimary field to given value.

### HasDnsPrimary

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasDnsPrimary() bool`

HasDnsPrimary returns a boolean if a field has been set.

### SetDnsPrimaryNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetDnsPrimaryNil(b bool)`

 SetDnsPrimaryNil sets the value for DnsPrimary to be an explicit nil

### UnsetDnsPrimary
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetDnsPrimary()`

UnsetDnsPrimary ensures that no value is present for DnsPrimary, not even an explicit nil
### GetDnsSecondary

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetDnsSecondary() string`

GetDnsSecondary returns the DnsSecondary field if non-nil, zero value otherwise.

### GetDnsSecondaryOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetDnsSecondaryOk() (*string, bool)`

GetDnsSecondaryOk returns a tuple with the DnsSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSecondary

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetDnsSecondary(v string)`

SetDnsSecondary sets DnsSecondary field to given value.

### HasDnsSecondary

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasDnsSecondary() bool`

HasDnsSecondary returns a boolean if a field has been set.

### SetDnsSecondaryNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetDnsSecondaryNil(b bool)`

 SetDnsSecondaryNil sets the value for DnsSecondary to be an explicit nil

### UnsetDnsSecondary
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetDnsSecondary()`

UnsetDnsSecondary ensures that no value is present for DnsSecondary, not even an explicit nil
### GetCidr

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### SetCidrNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetCidrNil(b bool)`

 SetCidrNil sets the value for Cidr to be an explicit nil

### UnsetCidr
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetCidr()`

UnsetCidr ensures that no value is present for Cidr, not even an explicit nil
### GetGatewayIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetGatewayIPv6() string`

GetGatewayIPv6 returns the GatewayIPv6 field if non-nil, zero value otherwise.

### GetGatewayIPv6Ok

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetGatewayIPv6Ok() (*string, bool)`

GetGatewayIPv6Ok returns a tuple with the GatewayIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetGatewayIPv6(v string)`

SetGatewayIPv6 sets GatewayIPv6 field to given value.

### HasGatewayIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasGatewayIPv6() bool`

HasGatewayIPv6 returns a boolean if a field has been set.

### SetGatewayIPv6Nil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetGatewayIPv6Nil(b bool)`

 SetGatewayIPv6Nil sets the value for GatewayIPv6 to be an explicit nil

### UnsetGatewayIPv6
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetGatewayIPv6()`

UnsetGatewayIPv6 ensures that no value is present for GatewayIPv6, not even an explicit nil
### GetNetmaskIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetNetmaskIPv6() string`

GetNetmaskIPv6 returns the NetmaskIPv6 field if non-nil, zero value otherwise.

### GetNetmaskIPv6Ok

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetNetmaskIPv6Ok() (*string, bool)`

GetNetmaskIPv6Ok returns a tuple with the NetmaskIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmaskIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetNetmaskIPv6(v string)`

SetNetmaskIPv6 sets NetmaskIPv6 field to given value.

### HasNetmaskIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasNetmaskIPv6() bool`

HasNetmaskIPv6 returns a boolean if a field has been set.

### SetNetmaskIPv6Nil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetNetmaskIPv6Nil(b bool)`

 SetNetmaskIPv6Nil sets the value for NetmaskIPv6 to be an explicit nil

### UnsetNetmaskIPv6
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetNetmaskIPv6()`

UnsetNetmaskIPv6 ensures that no value is present for NetmaskIPv6, not even an explicit nil
### GetDnsPrimaryIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetDnsPrimaryIPv6() string`

GetDnsPrimaryIPv6 returns the DnsPrimaryIPv6 field if non-nil, zero value otherwise.

### GetDnsPrimaryIPv6Ok

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetDnsPrimaryIPv6Ok() (*string, bool)`

GetDnsPrimaryIPv6Ok returns a tuple with the DnsPrimaryIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrimaryIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetDnsPrimaryIPv6(v string)`

SetDnsPrimaryIPv6 sets DnsPrimaryIPv6 field to given value.

### HasDnsPrimaryIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasDnsPrimaryIPv6() bool`

HasDnsPrimaryIPv6 returns a boolean if a field has been set.

### SetDnsPrimaryIPv6Nil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetDnsPrimaryIPv6Nil(b bool)`

 SetDnsPrimaryIPv6Nil sets the value for DnsPrimaryIPv6 to be an explicit nil

### UnsetDnsPrimaryIPv6
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetDnsPrimaryIPv6()`

UnsetDnsPrimaryIPv6 ensures that no value is present for DnsPrimaryIPv6, not even an explicit nil
### GetDnsSecondaryIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetDnsSecondaryIPv6() string`

GetDnsSecondaryIPv6 returns the DnsSecondaryIPv6 field if non-nil, zero value otherwise.

### GetDnsSecondaryIPv6Ok

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetDnsSecondaryIPv6Ok() (*string, bool)`

GetDnsSecondaryIPv6Ok returns a tuple with the DnsSecondaryIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSecondaryIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetDnsSecondaryIPv6(v string)`

SetDnsSecondaryIPv6 sets DnsSecondaryIPv6 field to given value.

### HasDnsSecondaryIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasDnsSecondaryIPv6() bool`

HasDnsSecondaryIPv6 returns a boolean if a field has been set.

### SetDnsSecondaryIPv6Nil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetDnsSecondaryIPv6Nil(b bool)`

 SetDnsSecondaryIPv6Nil sets the value for DnsSecondaryIPv6 to be an explicit nil

### UnsetDnsSecondaryIPv6
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetDnsSecondaryIPv6()`

UnsetDnsSecondaryIPv6 ensures that no value is present for DnsSecondaryIPv6, not even an explicit nil
### GetCidrIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetCidrIPv6() string`

GetCidrIPv6 returns the CidrIPv6 field if non-nil, zero value otherwise.

### GetCidrIPv6Ok

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetCidrIPv6Ok() (*string, bool)`

GetCidrIPv6Ok returns a tuple with the CidrIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetCidrIPv6(v string)`

SetCidrIPv6 sets CidrIPv6 field to given value.

### HasCidrIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasCidrIPv6() bool`

HasCidrIPv6 returns a boolean if a field has been set.

### SetCidrIPv6Nil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetCidrIPv6Nil(b bool)`

 SetCidrIPv6Nil sets the value for CidrIPv6 to be an explicit nil

### UnsetCidrIPv6
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetCidrIPv6()`

UnsetCidrIPv6 ensures that no value is present for CidrIPv6, not even an explicit nil
### GetTftpServer

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetTftpServer() string`

GetTftpServer returns the TftpServer field if non-nil, zero value otherwise.

### GetTftpServerOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetTftpServerOk() (*string, bool)`

GetTftpServerOk returns a tuple with the TftpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTftpServer

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetTftpServer(v string)`

SetTftpServer sets TftpServer field to given value.

### HasTftpServer

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasTftpServer() bool`

HasTftpServer returns a boolean if a field has been set.

### SetTftpServerNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetTftpServerNil(b bool)`

 SetTftpServerNil sets the value for TftpServer to be an explicit nil

### UnsetTftpServer
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetTftpServer()`

UnsetTftpServer ensures that no value is present for TftpServer, not even an explicit nil
### GetBootFile

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetBootFile() string`

GetBootFile returns the BootFile field if non-nil, zero value otherwise.

### GetBootFileOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetBootFileOk() (*string, bool)`

GetBootFileOk returns a tuple with the BootFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootFile

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetBootFile(v string)`

SetBootFile sets BootFile field to given value.

### HasBootFile

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasBootFile() bool`

HasBootFile returns a boolean if a field has been set.

### SetBootFileNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetBootFileNil(b bool)`

 SetBootFileNil sets the value for BootFile to be an explicit nil

### UnsetBootFile
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetBootFile()`

UnsetBootFile ensures that no value is present for BootFile, not even an explicit nil
### GetSwitchId

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### SetSwitchIdNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetSwitchIdNil(b bool)`

 SetSwitchIdNil sets the value for SwitchId to be an explicit nil

### UnsetSwitchId
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetSwitchId()`

UnsetSwitchId ensures that no value is present for SwitchId, not even an explicit nil
### GetFabricId

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetFabricId() string`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetFabricIdOk() (*string, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetFabricId(v string)`

SetFabricId sets FabricId field to given value.

### HasFabricId

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasFabricId() bool`

HasFabricId returns a boolean if a field has been set.

### SetFabricIdNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetFabricIdNil(b bool)`

 SetFabricIdNil sets the value for FabricId to be an explicit nil

### UnsetFabricId
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetFabricId()`

UnsetFabricId ensures that no value is present for FabricId, not even an explicit nil
### GetNetworkRole

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetNetworkRole() string`

GetNetworkRole returns the NetworkRole field if non-nil, zero value otherwise.

### GetNetworkRoleOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetNetworkRoleOk() (*string, bool)`

GetNetworkRoleOk returns a tuple with the NetworkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkRole

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetNetworkRole(v string)`

SetNetworkRole sets NetworkRole field to given value.

### HasNetworkRole

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasNetworkRole() bool`

HasNetworkRole returns a boolean if a field has been set.

### SetNetworkRoleNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetNetworkRoleNil(b bool)`

 SetNetworkRoleNil sets the value for NetworkRole to be an explicit nil

### UnsetNetworkRole
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetNetworkRole()`

UnsetNetworkRole ensures that no value is present for NetworkRole, not even an explicit nil
### GetStatus

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetAvailabilityZone

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetPool

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetPool() UpdateNetwork200ResponseAllOfNetworkPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetPoolOk() (*UpdateNetwork200ResponseAllOfNetworkPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetPool(v UpdateNetwork200ResponseAllOfNetworkPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetPoolIPv6() UpdateNetwork200ResponseAllOfNetworkPoolIPv6`

GetPoolIPv6 returns the PoolIPv6 field if non-nil, zero value otherwise.

### GetPoolIPv6Ok

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetPoolIPv6Ok() (*UpdateNetwork200ResponseAllOfNetworkPoolIPv6, bool)`

GetPoolIPv6Ok returns a tuple with the PoolIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetPoolIPv6(v UpdateNetwork200ResponseAllOfNetworkPoolIPv6)`

SetPoolIPv6 sets PoolIPv6 field to given value.

### HasPoolIPv6

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasPoolIPv6() bool`

HasPoolIPv6 returns a boolean if a field has been set.

### GetNetworkProxy

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetNetworkProxy() UpdateNetwork200ResponseAllOfNetworkNetworkProxy`

GetNetworkProxy returns the NetworkProxy field if non-nil, zero value otherwise.

### GetNetworkProxyOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetNetworkProxyOk() (*UpdateNetwork200ResponseAllOfNetworkNetworkProxy, bool)`

GetNetworkProxyOk returns a tuple with the NetworkProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProxy

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetNetworkProxy(v UpdateNetwork200ResponseAllOfNetworkNetworkProxy)`

SetNetworkProxy sets NetworkProxy field to given value.

### HasNetworkProxy

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasNetworkProxy() bool`

HasNetworkProxy returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetNetworkDomain() UpdateNetwork200ResponseAllOfNetworkNetworkDomain`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetNetworkDomainOk() (*UpdateNetwork200ResponseAllOfNetworkNetworkDomain, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetNetworkDomain(v UpdateNetwork200ResponseAllOfNetworkNetworkDomain)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetSearchDomains

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetSearchDomains() string`

GetSearchDomains returns the SearchDomains field if non-nil, zero value otherwise.

### GetSearchDomainsOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetSearchDomainsOk() (*string, bool)`

GetSearchDomainsOk returns a tuple with the SearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomains

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetSearchDomains(v string)`

SetSearchDomains sets SearchDomains field to given value.

### HasSearchDomains

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasSearchDomains() bool`

HasSearchDomains returns a boolean if a field has been set.

### SetSearchDomainsNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetSearchDomainsNil(b bool)`

 SetSearchDomainsNil sets the value for SearchDomains to be an explicit nil

### UnsetSearchDomains
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetSearchDomains()`

UnsetSearchDomains ensures that no value is present for SearchDomains, not even an explicit nil
### GetPrefixLength

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetPrefixLength() string`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetPrefixLengthOk() (*string, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetPrefixLength(v string)`

SetPrefixLength sets PrefixLength field to given value.

### HasPrefixLength

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasPrefixLength() bool`

HasPrefixLength returns a boolean if a field has been set.

### SetPrefixLengthNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetPrefixLengthNil(b bool)`

 SetPrefixLengthNil sets the value for PrefixLength to be an explicit nil

### UnsetPrefixLength
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetPrefixLength()`

UnsetPrefixLength ensures that no value is present for PrefixLength, not even an explicit nil
### GetVisibility

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEnableAdmin

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetEnableAdmin() bool`

GetEnableAdmin returns the EnableAdmin field if non-nil, zero value otherwise.

### GetEnableAdminOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetEnableAdminOk() (*bool, bool)`

GetEnableAdminOk returns a tuple with the EnableAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdmin

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetEnableAdmin(v bool)`

SetEnableAdmin sets EnableAdmin field to given value.

### HasEnableAdmin

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasEnableAdmin() bool`

HasEnableAdmin returns a boolean if a field has been set.

### GetActive

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDefaultNetwork

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetDefaultNetwork() bool`

GetDefaultNetwork returns the DefaultNetwork field if non-nil, zero value otherwise.

### GetDefaultNetworkOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetDefaultNetworkOk() (*bool, bool)`

GetDefaultNetworkOk returns a tuple with the DefaultNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetwork

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetDefaultNetwork(v bool)`

SetDefaultNetwork sets DefaultNetwork field to given value.

### HasDefaultNetwork

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasDefaultNetwork() bool`

HasDefaultNetwork returns a boolean if a field has been set.

### GetAssignPublicIp

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetAssignPublicIp() bool`

GetAssignPublicIp returns the AssignPublicIp field if non-nil, zero value otherwise.

### GetAssignPublicIpOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetAssignPublicIpOk() (*bool, bool)`

GetAssignPublicIpOk returns a tuple with the AssignPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignPublicIp

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetAssignPublicIp(v bool)`

SetAssignPublicIp sets AssignPublicIp field to given value.

### HasAssignPublicIp

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasAssignPublicIp() bool`

HasAssignPublicIp returns a boolean if a field has been set.

### GetNoProxy

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetNoProxy() string`

GetNoProxy returns the NoProxy field if non-nil, zero value otherwise.

### GetNoProxyOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetNoProxyOk() (*string, bool)`

GetNoProxyOk returns a tuple with the NoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxy

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetNoProxy(v string)`

SetNoProxy sets NoProxy field to given value.

### HasNoProxy

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasNoProxy() bool`

HasNoProxy returns a boolean if a field has been set.

### SetNoProxyNil

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetNoProxyNil(b bool)`

 SetNoProxyNil sets the value for NoProxy to be an explicit nil

### UnsetNoProxy
`func (o *UpdateNetwork200ResponseAllOfNetwork) UnsetNoProxy()`

UnsetNoProxy ensures that no value is present for NoProxy, not even an explicit nil
### GetApplianceUrlProxyBypass

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetApplianceUrlProxyBypass() bool`

GetApplianceUrlProxyBypass returns the ApplianceUrlProxyBypass field if non-nil, zero value otherwise.

### GetApplianceUrlProxyBypassOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetApplianceUrlProxyBypassOk() (*bool, bool)`

GetApplianceUrlProxyBypassOk returns a tuple with the ApplianceUrlProxyBypass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrlProxyBypass

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetApplianceUrlProxyBypass(v bool)`

SetApplianceUrlProxyBypass sets ApplianceUrlProxyBypass field to given value.

### HasApplianceUrlProxyBypass

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasApplianceUrlProxyBypass() bool`

HasApplianceUrlProxyBypass returns a boolean if a field has been set.

### GetZonePool

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetZonePool() UpdateNetwork200ResponseAllOfNetworkZonePool`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetZonePoolOk() (*UpdateNetwork200ResponseAllOfNetworkZonePool, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetZonePool(v UpdateNetwork200ResponseAllOfNetworkZonePool)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetAllowStaticOverride

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetAllowStaticOverride() bool`

GetAllowStaticOverride returns the AllowStaticOverride field if non-nil, zero value otherwise.

### GetAllowStaticOverrideOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetAllowStaticOverrideOk() (*bool, bool)`

GetAllowStaticOverrideOk returns a tuple with the AllowStaticOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowStaticOverride

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetAllowStaticOverride(v bool)`

SetAllowStaticOverride sets AllowStaticOverride field to given value.

### HasAllowStaticOverride

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasAllowStaticOverride() bool`

HasAllowStaticOverride returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetConfig() UpdateNetwork200ResponseAllOfNetworkConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetConfigOk() (*UpdateNetwork200ResponseAllOfNetworkConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetConfig(v UpdateNetwork200ResponseAllOfNetworkConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTenants

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetTenants() []UpdateNetwork200ResponseAllOfNetworkTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetTenantsOk() (*[]UpdateNetwork200ResponseAllOfNetworkTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetTenants(v []UpdateNetwork200ResponseAllOfNetworkTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetResourcePermission() UpdateNetwork200ResponseAllOfNetworkResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *UpdateNetwork200ResponseAllOfNetwork) GetResourcePermissionOk() (*UpdateNetwork200ResponseAllOfNetworkResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *UpdateNetwork200ResponseAllOfNetwork) SetResourcePermission(v UpdateNetwork200ResponseAllOfNetworkResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *UpdateNetwork200ResponseAllOfNetwork) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


