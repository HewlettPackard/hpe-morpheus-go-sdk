# UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Addresses** | Pointer to **[]map[string]interface{}** |  | [optional] 
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
**NetworkDomain** | Pointer to [**UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterfaceNetworkDomain**](UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterfaceNetworkDomain.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PrimaryInterface** | Pointer to **bool** |  | [optional] 
**NetworkPoolIPv6** | Pointer to **map[string]interface{}** |  | [optional] 
**Network** | Pointer to [**UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterfaceNetwork**](UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterfaceNetwork.md) |  | [optional] 
**VlanId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterfaceType**](UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterfaceType.md) |  | [optional] 
**NetworkPosition** | Pointer to **NullableString** |  | [optional] 
**PoolAssigned** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalType** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface

`func NewUpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface() *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface`

NewUpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface instantiates a new UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterfaceWithDefaults

`func NewUpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterfaceWithDefaults() *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface`

NewUpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterfaceWithDefaults instantiates a new UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddresses

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetAddresses() []map[string]interface{}`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetAddressesOk() (*[]map[string]interface{}, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetAddresses(v []map[string]interface{})`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetInternalId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetInterfaceId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### SetInterfaceIdNil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetInterfaceIdNil(b bool)`

 SetInterfaceIdNil sets the value for InterfaceId to be an explicit nil

### UnsetInterfaceId
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) UnsetInterfaceId()`

UnsetInterfaceId ensures that no value is present for InterfaceId, not even an explicit nil
### GetDisplayOrder

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetNetworkPool

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkPool() map[string]interface{}`

GetNetworkPool returns the NetworkPool field if non-nil, zero value otherwise.

### GetNetworkPoolOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkPoolOk() (*map[string]interface{}, bool)`

GetNetworkPoolOk returns a tuple with the NetworkPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPool

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetNetworkPool(v map[string]interface{})`

SetNetworkPool sets NetworkPool field to given value.

### HasNetworkPool

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasNetworkPool() bool`

HasNetworkPool returns a boolean if a field has been set.

### SetNetworkPoolNil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetNetworkPoolNil(b bool)`

 SetNetworkPoolNil sets the value for NetworkPool to be an explicit nil

### UnsetNetworkPool
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) UnsetNetworkPool()`

UnsetNetworkPool ensures that no value is present for NetworkPool, not even an explicit nil
### GetDhcp

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetDhcp() bool`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetDhcpOk() (*bool, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetDhcp(v bool)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### GetUuid

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetActive

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUniqueId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetSubnet

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetReplaceHostRecord

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetReplaceHostRecord() bool`

GetReplaceHostRecord returns the ReplaceHostRecord field if non-nil, zero value otherwise.

### GetReplaceHostRecordOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetReplaceHostRecordOk() (*bool, bool)`

GetReplaceHostRecordOk returns a tuple with the ReplaceHostRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceHostRecord

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetReplaceHostRecord(v bool)`

SetReplaceHostRecord sets ReplaceHostRecord field to given value.

### HasReplaceHostRecord

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasReplaceHostRecord() bool`

HasReplaceHostRecord returns a boolean if a field has been set.

### GetIpMode

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetVersion

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetIpSubnet

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetIpSubnet() string`

GetIpSubnet returns the IpSubnet field if non-nil, zero value otherwise.

### GetIpSubnetOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetIpSubnetOk() (*string, bool)`

GetIpSubnetOk returns a tuple with the IpSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSubnet

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetIpSubnet(v string)`

SetIpSubnet sets IpSubnet field to given value.

### HasIpSubnet

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasIpSubnet() bool`

HasIpSubnet returns a boolean if a field has been set.

### SetIpSubnetNil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetIpSubnetNil(b bool)`

 SetIpSubnetNil sets the value for IpSubnet to be an explicit nil

### UnsetIpSubnet
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) UnsetIpSubnet()`

UnsetIpSubnet ensures that no value is present for IpSubnet, not even an explicit nil
### GetConfig

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetPublicIpAddress

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### GetFabricId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetFabricId() string`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetFabricIdOk() (*string, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetFabricId(v string)`

SetFabricId sets FabricId field to given value.

### HasFabricId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasFabricId() bool`

HasFabricId returns a boolean if a field has been set.

### SetFabricIdNil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetFabricIdNil(b bool)`

 SetFabricIdNil sets the value for FabricId to be an explicit nil

### UnsetFabricId
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) UnsetFabricId()`

UnsetFabricId ensures that no value is present for FabricId, not even an explicit nil
### GetIpv6Subnet

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetIpv6Subnet() string`

GetIpv6Subnet returns the Ipv6Subnet field if non-nil, zero value otherwise.

### GetIpv6SubnetOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetIpv6SubnetOk() (*string, bool)`

GetIpv6SubnetOk returns a tuple with the Ipv6Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Subnet

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetIpv6Subnet(v string)`

SetIpv6Subnet sets Ipv6Subnet field to given value.

### HasIpv6Subnet

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasIpv6Subnet() bool`

HasIpv6Subnet returns a boolean if a field has been set.

### SetIpv6SubnetNil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetIpv6SubnetNil(b bool)`

 SetIpv6SubnetNil sets the value for Ipv6Subnet to be an explicit nil

### UnsetIpv6Subnet
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) UnsetIpv6Subnet()`

UnsetIpv6Subnet ensures that no value is present for Ipv6Subnet, not even an explicit nil
### GetMacAddress

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetPublicIpv6Address

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetPublicIpv6Address() string`

GetPublicIpv6Address returns the PublicIpv6Address field if non-nil, zero value otherwise.

### GetPublicIpv6AddressOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetPublicIpv6AddressOk() (*string, bool)`

GetPublicIpv6AddressOk returns a tuple with the PublicIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpv6Address

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetPublicIpv6Address(v string)`

SetPublicIpv6Address sets PublicIpv6Address field to given value.

### HasPublicIpv6Address

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasPublicIpv6Address() bool`

HasPublicIpv6Address returns a boolean if a field has been set.

### SetPublicIpv6AddressNil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetPublicIpv6AddressNil(b bool)`

 SetPublicIpv6AddressNil sets the value for PublicIpv6Address to be an explicit nil

### UnsetPublicIpv6Address
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) UnsetPublicIpv6Address()`

UnsetPublicIpv6Address ensures that no value is present for PublicIpv6Address, not even an explicit nil
### GetRefType

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetNetworkGroup

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkGroup() string`

GetNetworkGroup returns the NetworkGroup field if non-nil, zero value otherwise.

### GetNetworkGroupOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkGroupOk() (*string, bool)`

GetNetworkGroupOk returns a tuple with the NetworkGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkGroup

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetNetworkGroup(v string)`

SetNetworkGroup sets NetworkGroup field to given value.

### HasNetworkGroup

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasNetworkGroup() bool`

HasNetworkGroup returns a boolean if a field has been set.

### SetNetworkGroupNil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetNetworkGroupNil(b bool)`

 SetNetworkGroupNil sets the value for NetworkGroup to be an explicit nil

### UnsetNetworkGroup
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) UnsetNetworkGroup()`

UnsetNetworkGroup ensures that no value is present for NetworkGroup, not even an explicit nil
### GetRefId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetNetworkDomain

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkDomain() UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterfaceNetworkDomain`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkDomainOk() (*UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterfaceNetworkDomain, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetNetworkDomain(v UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterfaceNetworkDomain)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetName

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimaryInterface

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetPrimaryInterface() bool`

GetPrimaryInterface returns the PrimaryInterface field if non-nil, zero value otherwise.

### GetPrimaryInterfaceOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetPrimaryInterfaceOk() (*bool, bool)`

GetPrimaryInterfaceOk returns a tuple with the PrimaryInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInterface

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetPrimaryInterface(v bool)`

SetPrimaryInterface sets PrimaryInterface field to given value.

### HasPrimaryInterface

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasPrimaryInterface() bool`

HasPrimaryInterface returns a boolean if a field has been set.

### GetNetworkPoolIPv6

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkPoolIPv6() map[string]interface{}`

GetNetworkPoolIPv6 returns the NetworkPoolIPv6 field if non-nil, zero value otherwise.

### GetNetworkPoolIPv6Ok

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkPoolIPv6Ok() (*map[string]interface{}, bool)`

GetNetworkPoolIPv6Ok returns a tuple with the NetworkPoolIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPoolIPv6

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetNetworkPoolIPv6(v map[string]interface{})`

SetNetworkPoolIPv6 sets NetworkPoolIPv6 field to given value.

### HasNetworkPoolIPv6

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasNetworkPoolIPv6() bool`

HasNetworkPoolIPv6 returns a boolean if a field has been set.

### SetNetworkPoolIPv6Nil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetNetworkPoolIPv6Nil(b bool)`

 SetNetworkPoolIPv6Nil sets the value for NetworkPoolIPv6 to be an explicit nil

### UnsetNetworkPoolIPv6
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) UnsetNetworkPoolIPv6()`

UnsetNetworkPoolIPv6 ensures that no value is present for NetworkPoolIPv6, not even an explicit nil
### GetNetwork

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetwork() UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterfaceNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkOk() (*UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterfaceNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetNetwork(v UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterfaceNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetVlanId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetVlanId() string`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetVlanIdOk() (*string, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetVlanId(v string)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### SetVlanIdNil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetVlanIdNil(b bool)`

 SetVlanIdNil sets the value for VlanId to be an explicit nil

### UnsetVlanId
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) UnsetVlanId()`

UnsetVlanId ensures that no value is present for VlanId, not even an explicit nil
### GetType

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetType() UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterfaceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetTypeOk() (*UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterfaceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetType(v UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterfaceType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetworkPosition

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkPosition() string`

GetNetworkPosition returns the NetworkPosition field if non-nil, zero value otherwise.

### GetNetworkPositionOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkPositionOk() (*string, bool)`

GetNetworkPositionOk returns a tuple with the NetworkPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPosition

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetNetworkPosition(v string)`

SetNetworkPosition sets NetworkPosition field to given value.

### HasNetworkPosition

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasNetworkPosition() bool`

HasNetworkPosition returns a boolean if a field has been set.

### SetNetworkPositionNil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetNetworkPositionNil(b bool)`

 SetNetworkPositionNil sets the value for NetworkPosition to be an explicit nil

### UnsetNetworkPosition
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) UnsetNetworkPosition()`

UnsetNetworkPosition ensures that no value is present for NetworkPosition, not even an explicit nil
### GetPoolAssigned

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetPoolAssigned() bool`

GetPoolAssigned returns the PoolAssigned field if non-nil, zero value otherwise.

### GetPoolAssignedOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetPoolAssignedOk() (*bool, bool)`

GetPoolAssignedOk returns a tuple with the PoolAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAssigned

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetPoolAssigned(v bool)`

SetPoolAssigned sets PoolAssigned field to given value.

### HasPoolAssigned

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasPoolAssigned() bool`

HasPoolAssigned returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalType

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### SetExternalTypeNil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetExternalTypeNil(b bool)`

 SetExternalTypeNil sets the value for ExternalType to be an explicit nil

### UnsetExternalType
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) UnsetExternalType()`

UnsetExternalType ensures that no value is present for ExternalType, not even an explicit nil
### GetExternalId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


