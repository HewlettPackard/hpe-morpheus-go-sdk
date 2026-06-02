# NetworkInterfaceUpdateSuccessServerInterfacesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Addresses** | Pointer to [**[]NetworkInterfaceUpdateSuccessServerInterfacesInnerAddressesInner**](NetworkInterfaceUpdateSuccessServerInterfacesInnerAddressesInner.md) |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**InterfaceId** | Pointer to **NullableString** |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 
**NetworkPool** | Pointer to **map[string]interface{}** |  | [optional] 
**Dhcp** | Pointer to **bool** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**Subnet** | Pointer to **NullableString** |  | [optional] 
**ReplaceHostRecord** | Pointer to **bool** |  | [optional] 
**IpMode** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**IpSubnet** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to **NullableString** |  | [optional] 
**PublicIpAddress** | Pointer to **string** |  | [optional] 
**FabricId** | Pointer to **NullableString** |  | [optional] 
**Ipv6Subnet** | Pointer to **NullableString** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**PublicIpv6Address** | Pointer to **NullableString** |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**NetworkGroup** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableString** |  | [optional] 
**NetworkDomain** | Pointer to [**NetworkInterfaceUpdateSuccessServerInterfacesInnerNetworkDomain**](NetworkInterfaceUpdateSuccessServerInterfacesInnerNetworkDomain.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PrimaryInterface** | Pointer to **bool** |  | [optional] 
**NetworkPoolIPv6** | Pointer to **map[string]interface{}** |  | [optional] 
**Network** | Pointer to [**NetworkInterfaceUpdateSuccessServerInterfacesInnerNetwork**](NetworkInterfaceUpdateSuccessServerInterfacesInnerNetwork.md) |  | [optional] 
**VlanId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**NetworkInterfaceUpdateSuccessServerInterfacesInnerType**](NetworkInterfaceUpdateSuccessServerInterfacesInnerType.md) |  | [optional] 
**NetworkPosition** | Pointer to **NullableString** |  | [optional] 
**PoolAssigned** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalType** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkInterfaceUpdateSuccessServerInterfacesInner

`func NewNetworkInterfaceUpdateSuccessServerInterfacesInner() *NetworkInterfaceUpdateSuccessServerInterfacesInner`

NewNetworkInterfaceUpdateSuccessServerInterfacesInner instantiates a new NetworkInterfaceUpdateSuccessServerInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddresses

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetAddresses() []NetworkInterfaceUpdateSuccessServerInterfacesInnerAddressesInner`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetAddressesOk() (*[]NetworkInterfaceUpdateSuccessServerInterfacesInnerAddressesInner, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetAddresses(v []NetworkInterfaceUpdateSuccessServerInterfacesInnerAddressesInner)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetInternalId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetInterfaceId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### SetInterfaceIdNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetInterfaceIdNil(b bool)`

 SetInterfaceIdNil sets the value for InterfaceId to be an explicit nil

### UnsetInterfaceId
`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) UnsetInterfaceId()`

UnsetInterfaceId ensures that no value is present for InterfaceId, not even an explicit nil
### GetDisplayOrder

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetNetworkPool

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetNetworkPool() map[string]interface{}`

GetNetworkPool returns the NetworkPool field if non-nil, zero value otherwise.

### GetNetworkPoolOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetNetworkPoolOk() (*map[string]interface{}, bool)`

GetNetworkPoolOk returns a tuple with the NetworkPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPool

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetNetworkPool(v map[string]interface{})`

SetNetworkPool sets NetworkPool field to given value.

### HasNetworkPool

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasNetworkPool() bool`

HasNetworkPool returns a boolean if a field has been set.

### SetNetworkPoolNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetNetworkPoolNil(b bool)`

 SetNetworkPoolNil sets the value for NetworkPool to be an explicit nil

### UnsetNetworkPool
`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) UnsetNetworkPool()`

UnsetNetworkPool ensures that no value is present for NetworkPool, not even an explicit nil
### GetDhcp

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetDhcp() bool`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetDhcpOk() (*bool, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetDhcp(v bool)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### GetUuid

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetActive

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUniqueId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetSubnet

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetReplaceHostRecord

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetReplaceHostRecord() bool`

GetReplaceHostRecord returns the ReplaceHostRecord field if non-nil, zero value otherwise.

### GetReplaceHostRecordOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetReplaceHostRecordOk() (*bool, bool)`

GetReplaceHostRecordOk returns a tuple with the ReplaceHostRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceHostRecord

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetReplaceHostRecord(v bool)`

SetReplaceHostRecord sets ReplaceHostRecord field to given value.

### HasReplaceHostRecord

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasReplaceHostRecord() bool`

HasReplaceHostRecord returns a boolean if a field has been set.

### GetIpMode

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetVersion

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetIpSubnet

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetIpSubnet() string`

GetIpSubnet returns the IpSubnet field if non-nil, zero value otherwise.

### GetIpSubnetOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetIpSubnetOk() (*string, bool)`

GetIpSubnetOk returns a tuple with the IpSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSubnet

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetIpSubnet(v string)`

SetIpSubnet sets IpSubnet field to given value.

### HasIpSubnet

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasIpSubnet() bool`

HasIpSubnet returns a boolean if a field has been set.

### SetIpSubnetNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetIpSubnetNil(b bool)`

 SetIpSubnetNil sets the value for IpSubnet to be an explicit nil

### UnsetIpSubnet
`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) UnsetIpSubnet()`

UnsetIpSubnet ensures that no value is present for IpSubnet, not even an explicit nil
### GetConfig

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetPublicIpAddress

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### GetFabricId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetFabricId() string`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetFabricIdOk() (*string, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetFabricId(v string)`

SetFabricId sets FabricId field to given value.

### HasFabricId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasFabricId() bool`

HasFabricId returns a boolean if a field has been set.

### SetFabricIdNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetFabricIdNil(b bool)`

 SetFabricIdNil sets the value for FabricId to be an explicit nil

### UnsetFabricId
`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) UnsetFabricId()`

UnsetFabricId ensures that no value is present for FabricId, not even an explicit nil
### GetIpv6Subnet

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetIpv6Subnet() string`

GetIpv6Subnet returns the Ipv6Subnet field if non-nil, zero value otherwise.

### GetIpv6SubnetOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetIpv6SubnetOk() (*string, bool)`

GetIpv6SubnetOk returns a tuple with the Ipv6Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Subnet

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetIpv6Subnet(v string)`

SetIpv6Subnet sets Ipv6Subnet field to given value.

### HasIpv6Subnet

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasIpv6Subnet() bool`

HasIpv6Subnet returns a boolean if a field has been set.

### SetIpv6SubnetNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetIpv6SubnetNil(b bool)`

 SetIpv6SubnetNil sets the value for Ipv6Subnet to be an explicit nil

### UnsetIpv6Subnet
`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) UnsetIpv6Subnet()`

UnsetIpv6Subnet ensures that no value is present for Ipv6Subnet, not even an explicit nil
### GetMacAddress

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetPublicIpv6Address

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetPublicIpv6Address() string`

GetPublicIpv6Address returns the PublicIpv6Address field if non-nil, zero value otherwise.

### GetPublicIpv6AddressOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetPublicIpv6AddressOk() (*string, bool)`

GetPublicIpv6AddressOk returns a tuple with the PublicIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpv6Address

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetPublicIpv6Address(v string)`

SetPublicIpv6Address sets PublicIpv6Address field to given value.

### HasPublicIpv6Address

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasPublicIpv6Address() bool`

HasPublicIpv6Address returns a boolean if a field has been set.

### SetPublicIpv6AddressNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetPublicIpv6AddressNil(b bool)`

 SetPublicIpv6AddressNil sets the value for PublicIpv6Address to be an explicit nil

### UnsetPublicIpv6Address
`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) UnsetPublicIpv6Address()`

UnsetPublicIpv6Address ensures that no value is present for PublicIpv6Address, not even an explicit nil
### GetRefType

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetNetworkGroup

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetNetworkGroup() string`

GetNetworkGroup returns the NetworkGroup field if non-nil, zero value otherwise.

### GetNetworkGroupOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetNetworkGroupOk() (*string, bool)`

GetNetworkGroupOk returns a tuple with the NetworkGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkGroup

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetNetworkGroup(v string)`

SetNetworkGroup sets NetworkGroup field to given value.

### HasNetworkGroup

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasNetworkGroup() bool`

HasNetworkGroup returns a boolean if a field has been set.

### SetNetworkGroupNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetNetworkGroupNil(b bool)`

 SetNetworkGroupNil sets the value for NetworkGroup to be an explicit nil

### UnsetNetworkGroup
`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) UnsetNetworkGroup()`

UnsetNetworkGroup ensures that no value is present for NetworkGroup, not even an explicit nil
### GetRefId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetNetworkDomain

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetNetworkDomain() NetworkInterfaceUpdateSuccessServerInterfacesInnerNetworkDomain`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetNetworkDomainOk() (*NetworkInterfaceUpdateSuccessServerInterfacesInnerNetworkDomain, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetNetworkDomain(v NetworkInterfaceUpdateSuccessServerInterfacesInnerNetworkDomain)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetName

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimaryInterface

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetPrimaryInterface() bool`

GetPrimaryInterface returns the PrimaryInterface field if non-nil, zero value otherwise.

### GetPrimaryInterfaceOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetPrimaryInterfaceOk() (*bool, bool)`

GetPrimaryInterfaceOk returns a tuple with the PrimaryInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInterface

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetPrimaryInterface(v bool)`

SetPrimaryInterface sets PrimaryInterface field to given value.

### HasPrimaryInterface

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasPrimaryInterface() bool`

HasPrimaryInterface returns a boolean if a field has been set.

### GetNetworkPoolIPv6

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetNetworkPoolIPv6() map[string]interface{}`

GetNetworkPoolIPv6 returns the NetworkPoolIPv6 field if non-nil, zero value otherwise.

### GetNetworkPoolIPv6Ok

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetNetworkPoolIPv6Ok() (*map[string]interface{}, bool)`

GetNetworkPoolIPv6Ok returns a tuple with the NetworkPoolIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPoolIPv6

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetNetworkPoolIPv6(v map[string]interface{})`

SetNetworkPoolIPv6 sets NetworkPoolIPv6 field to given value.

### HasNetworkPoolIPv6

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasNetworkPoolIPv6() bool`

HasNetworkPoolIPv6 returns a boolean if a field has been set.

### SetNetworkPoolIPv6Nil

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetNetworkPoolIPv6Nil(b bool)`

 SetNetworkPoolIPv6Nil sets the value for NetworkPoolIPv6 to be an explicit nil

### UnsetNetworkPoolIPv6
`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) UnsetNetworkPoolIPv6()`

UnsetNetworkPoolIPv6 ensures that no value is present for NetworkPoolIPv6, not even an explicit nil
### GetNetwork

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetNetwork() NetworkInterfaceUpdateSuccessServerInterfacesInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetNetworkOk() (*NetworkInterfaceUpdateSuccessServerInterfacesInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetNetwork(v NetworkInterfaceUpdateSuccessServerInterfacesInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetVlanId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetVlanId() string`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetVlanIdOk() (*string, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetVlanId(v string)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### SetVlanIdNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetVlanIdNil(b bool)`

 SetVlanIdNil sets the value for VlanId to be an explicit nil

### UnsetVlanId
`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) UnsetVlanId()`

UnsetVlanId ensures that no value is present for VlanId, not even an explicit nil
### GetType

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetType() NetworkInterfaceUpdateSuccessServerInterfacesInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetTypeOk() (*NetworkInterfaceUpdateSuccessServerInterfacesInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetType(v NetworkInterfaceUpdateSuccessServerInterfacesInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetworkPosition

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetNetworkPosition() string`

GetNetworkPosition returns the NetworkPosition field if non-nil, zero value otherwise.

### GetNetworkPositionOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetNetworkPositionOk() (*string, bool)`

GetNetworkPositionOk returns a tuple with the NetworkPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPosition

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetNetworkPosition(v string)`

SetNetworkPosition sets NetworkPosition field to given value.

### HasNetworkPosition

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasNetworkPosition() bool`

HasNetworkPosition returns a boolean if a field has been set.

### SetNetworkPositionNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetNetworkPositionNil(b bool)`

 SetNetworkPositionNil sets the value for NetworkPosition to be an explicit nil

### UnsetNetworkPosition
`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) UnsetNetworkPosition()`

UnsetNetworkPosition ensures that no value is present for NetworkPosition, not even an explicit nil
### GetPoolAssigned

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetPoolAssigned() bool`

GetPoolAssigned returns the PoolAssigned field if non-nil, zero value otherwise.

### GetPoolAssignedOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetPoolAssignedOk() (*bool, bool)`

GetPoolAssignedOk returns a tuple with the PoolAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAssigned

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetPoolAssigned(v bool)`

SetPoolAssigned sets PoolAssigned field to given value.

### HasPoolAssigned

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasPoolAssigned() bool`

HasPoolAssigned returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalType

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### SetExternalTypeNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetExternalTypeNil(b bool)`

 SetExternalTypeNil sets the value for ExternalType to be an explicit nil

### UnsetExternalType
`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) UnsetExternalType()`

UnsetExternalType ensures that no value is present for ExternalType, not even an explicit nil
### GetExternalId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *NetworkInterfaceUpdateSuccessServerInterfacesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


