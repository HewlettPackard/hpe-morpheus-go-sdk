# UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Addresses** | Pointer to **[]map[string]interface{}** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**InterfaceId** | Pointer to **string** |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 
**NetworkPool** | Pointer to **map[string]interface{}** |  | [optional] 
**Dhcp** | Pointer to **bool** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**Subnet** | Pointer to **string** |  | [optional] 
**ReplaceHostRecord** | Pointer to **bool** |  | [optional] 
**IpMode** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**IpSubnet** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **string** |  | [optional] 
**PublicIpAddress** | Pointer to **string** |  | [optional] 
**FabricId** | Pointer to **string** |  | [optional] 
**Ipv6Subnet** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**PublicIpv6Address** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**NetworkGroup** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **string** |  | [optional] 
**NetworkDomain** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PrimaryInterface** | Pointer to **bool** |  | [optional] 
**NetworkPoolIPv6** | Pointer to **map[string]interface{}** |  | [optional] 
**Network** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**VlanId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**NetworkPosition** | Pointer to **string** |  | [optional] 
**PoolAssigned** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalType** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface

`func NewUpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface() *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface`

NewUpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface instantiates a new UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterfaceWithDefaults

`func NewUpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterfaceWithDefaults() *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface`

NewUpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterfaceWithDefaults instantiates a new UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddresses

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetAddresses() []map[string]interface{}`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetAddressesOk() (*[]map[string]interface{}, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetAddresses(v []map[string]interface{})`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetInternalId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetInterfaceId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetNetworkPool

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkPool() map[string]interface{}`

GetNetworkPool returns the NetworkPool field if non-nil, zero value otherwise.

### GetNetworkPoolOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkPoolOk() (*map[string]interface{}, bool)`

GetNetworkPoolOk returns a tuple with the NetworkPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPool

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetNetworkPool(v map[string]interface{})`

SetNetworkPool sets NetworkPool field to given value.

### HasNetworkPool

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasNetworkPool() bool`

HasNetworkPool returns a boolean if a field has been set.

### GetDhcp

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetDhcp() bool`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetDhcpOk() (*bool, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetDhcp(v bool)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### GetUuid

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetActive

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUniqueId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetSubnet

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetReplaceHostRecord

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetReplaceHostRecord() bool`

GetReplaceHostRecord returns the ReplaceHostRecord field if non-nil, zero value otherwise.

### GetReplaceHostRecordOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetReplaceHostRecordOk() (*bool, bool)`

GetReplaceHostRecordOk returns a tuple with the ReplaceHostRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceHostRecord

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetReplaceHostRecord(v bool)`

SetReplaceHostRecord sets ReplaceHostRecord field to given value.

### HasReplaceHostRecord

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasReplaceHostRecord() bool`

HasReplaceHostRecord returns a boolean if a field has been set.

### GetIpMode

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetVersion

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIpSubnet

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetIpSubnet() string`

GetIpSubnet returns the IpSubnet field if non-nil, zero value otherwise.

### GetIpSubnetOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetIpSubnetOk() (*string, bool)`

GetIpSubnetOk returns a tuple with the IpSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSubnet

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetIpSubnet(v string)`

SetIpSubnet sets IpSubnet field to given value.

### HasIpSubnet

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasIpSubnet() bool`

HasIpSubnet returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetPublicIpAddress

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### GetFabricId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetFabricId() string`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetFabricIdOk() (*string, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetFabricId(v string)`

SetFabricId sets FabricId field to given value.

### HasFabricId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasFabricId() bool`

HasFabricId returns a boolean if a field has been set.

### GetIpv6Subnet

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetIpv6Subnet() string`

GetIpv6Subnet returns the Ipv6Subnet field if non-nil, zero value otherwise.

### GetIpv6SubnetOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetIpv6SubnetOk() (*string, bool)`

GetIpv6SubnetOk returns a tuple with the Ipv6Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Subnet

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetIpv6Subnet(v string)`

SetIpv6Subnet sets Ipv6Subnet field to given value.

### HasIpv6Subnet

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasIpv6Subnet() bool`

HasIpv6Subnet returns a boolean if a field has been set.

### GetMacAddress

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetPublicIpv6Address

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetPublicIpv6Address() string`

GetPublicIpv6Address returns the PublicIpv6Address field if non-nil, zero value otherwise.

### GetPublicIpv6AddressOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetPublicIpv6AddressOk() (*string, bool)`

GetPublicIpv6AddressOk returns a tuple with the PublicIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpv6Address

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetPublicIpv6Address(v string)`

SetPublicIpv6Address sets PublicIpv6Address field to given value.

### HasPublicIpv6Address

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasPublicIpv6Address() bool`

HasPublicIpv6Address returns a boolean if a field has been set.

### GetRefType

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetNetworkGroup

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkGroup() string`

GetNetworkGroup returns the NetworkGroup field if non-nil, zero value otherwise.

### GetNetworkGroupOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkGroupOk() (*string, bool)`

GetNetworkGroupOk returns a tuple with the NetworkGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkGroup

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetNetworkGroup(v string)`

SetNetworkGroup sets NetworkGroup field to given value.

### HasNetworkGroup

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasNetworkGroup() bool`

HasNetworkGroup returns a boolean if a field has been set.

### GetRefId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkDomain() string`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkDomainOk() (*string, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetNetworkDomain(v string)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetName

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimaryInterface

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetPrimaryInterface() bool`

GetPrimaryInterface returns the PrimaryInterface field if non-nil, zero value otherwise.

### GetPrimaryInterfaceOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetPrimaryInterfaceOk() (*bool, bool)`

GetPrimaryInterfaceOk returns a tuple with the PrimaryInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInterface

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetPrimaryInterface(v bool)`

SetPrimaryInterface sets PrimaryInterface field to given value.

### HasPrimaryInterface

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasPrimaryInterface() bool`

HasPrimaryInterface returns a boolean if a field has been set.

### GetNetworkPoolIPv6

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkPoolIPv6() map[string]interface{}`

GetNetworkPoolIPv6 returns the NetworkPoolIPv6 field if non-nil, zero value otherwise.

### GetNetworkPoolIPv6Ok

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkPoolIPv6Ok() (*map[string]interface{}, bool)`

GetNetworkPoolIPv6Ok returns a tuple with the NetworkPoolIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPoolIPv6

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetNetworkPoolIPv6(v map[string]interface{})`

SetNetworkPoolIPv6 sets NetworkPoolIPv6 field to given value.

### HasNetworkPoolIPv6

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasNetworkPoolIPv6() bool`

HasNetworkPoolIPv6 returns a boolean if a field has been set.

### GetNetwork

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetwork() GetAlerts200ResponseAllOfChecksInnerAccount`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetNetwork(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetVlanId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetVlanId() string`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetVlanIdOk() (*string, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetVlanId(v string)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetType

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetType() GetAlerts200ResponseAllOfChecksInnerAccount`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetTypeOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetType(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetworkPosition

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkPosition() string`

GetNetworkPosition returns the NetworkPosition field if non-nil, zero value otherwise.

### GetNetworkPositionOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetNetworkPositionOk() (*string, bool)`

GetNetworkPositionOk returns a tuple with the NetworkPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPosition

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetNetworkPosition(v string)`

SetNetworkPosition sets NetworkPosition field to given value.

### HasNetworkPosition

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasNetworkPosition() bool`

HasNetworkPosition returns a boolean if a field has been set.

### GetPoolAssigned

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetPoolAssigned() bool`

GetPoolAssigned returns the PoolAssigned field if non-nil, zero value otherwise.

### GetPoolAssignedOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetPoolAssignedOk() (*bool, bool)`

GetPoolAssignedOk returns a tuple with the PoolAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAssigned

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetPoolAssigned(v bool)`

SetPoolAssigned sets PoolAssigned field to given value.

### HasPoolAssigned

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasPoolAssigned() bool`

HasPoolAssigned returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalType

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


