# UpdateHostResize200ResponseAllOfServerInterfacesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**UniqueId** | Pointer to **NullableString** |  | [optional] 
**PublicIpAddress** | Pointer to **NullableString** |  | [optional] 
**PublicIpv6Address** | Pointer to **NullableString** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Ipv6Address** | Pointer to **NullableString** |  | [optional] 
**IpSubnet** | Pointer to **NullableString** |  | [optional] 
**Ipv6Subnet** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Dhcp** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**PoolAssigned** | Pointer to **bool** |  | [optional] 
**PrimaryInterface** | Pointer to **bool** |  | [optional] 
**Network** | Pointer to [**UpdateHostResize200ResponseAllOfServerInterfacesInnerNetwork**](UpdateHostResize200ResponseAllOfServerInterfacesInnerNetwork.md) |  | [optional] 
**Subnet** | Pointer to **NullableString** |  | [optional] 
**NetworkGroup** | Pointer to **NullableString** |  | [optional] 
**NetworkPosition** | Pointer to **NullableString** |  | [optional] 
**NetworkPool** | Pointer to [**UpdateHostResize200ResponseAllOfServerInterfacesInnerNetworkPool**](UpdateHostResize200ResponseAllOfServerInterfacesInnerNetworkPool.md) |  | [optional] 
**NetworkDomain** | Pointer to [**UpdateHostResize200ResponseAllOfServerInterfacesInnerNetworkDomain**](UpdateHostResize200ResponseAllOfServerInterfacesInnerNetworkDomain.md) |  | [optional] 
**Type** | Pointer to [**UpdateHostResize200ResponseAllOfServerInterfacesInnerType**](UpdateHostResize200ResponseAllOfServerInterfacesInnerType.md) |  | [optional] 
**IpMode** | Pointer to **NullableString** |  | [optional] 
**MacAddress** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateHostResize200ResponseAllOfServerInterfacesInner

`func NewUpdateHostResize200ResponseAllOfServerInterfacesInner() *UpdateHostResize200ResponseAllOfServerInterfacesInner`

NewUpdateHostResize200ResponseAllOfServerInterfacesInner instantiates a new UpdateHostResize200ResponseAllOfServerInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHostResize200ResponseAllOfServerInterfacesInnerWithDefaults

`func NewUpdateHostResize200ResponseAllOfServerInterfacesInnerWithDefaults() *UpdateHostResize200ResponseAllOfServerInterfacesInner`

NewUpdateHostResize200ResponseAllOfServerInterfacesInnerWithDefaults instantiates a new UpdateHostResize200ResponseAllOfServerInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRefType

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetName

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInternalId

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetUniqueId

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### SetUniqueIdNil

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetUniqueIdNil(b bool)`

 SetUniqueIdNil sets the value for UniqueId to be an explicit nil

### UnsetUniqueId
`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) UnsetUniqueId()`

UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
### GetPublicIpAddress

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### SetPublicIpAddressNil

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetPublicIpAddressNil(b bool)`

 SetPublicIpAddressNil sets the value for PublicIpAddress to be an explicit nil

### UnsetPublicIpAddress
`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) UnsetPublicIpAddress()`

UnsetPublicIpAddress ensures that no value is present for PublicIpAddress, not even an explicit nil
### GetPublicIpv6Address

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetPublicIpv6Address() string`

GetPublicIpv6Address returns the PublicIpv6Address field if non-nil, zero value otherwise.

### GetPublicIpv6AddressOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetPublicIpv6AddressOk() (*string, bool)`

GetPublicIpv6AddressOk returns a tuple with the PublicIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpv6Address

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetPublicIpv6Address(v string)`

SetPublicIpv6Address sets PublicIpv6Address field to given value.

### HasPublicIpv6Address

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasPublicIpv6Address() bool`

HasPublicIpv6Address returns a boolean if a field has been set.

### SetPublicIpv6AddressNil

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetPublicIpv6AddressNil(b bool)`

 SetPublicIpv6AddressNil sets the value for PublicIpv6Address to be an explicit nil

### UnsetPublicIpv6Address
`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) UnsetPublicIpv6Address()`

UnsetPublicIpv6Address ensures that no value is present for PublicIpv6Address, not even an explicit nil
### GetIpAddress

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpv6Address

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### SetIpv6AddressNil

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetIpv6AddressNil(b bool)`

 SetIpv6AddressNil sets the value for Ipv6Address to be an explicit nil

### UnsetIpv6Address
`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) UnsetIpv6Address()`

UnsetIpv6Address ensures that no value is present for Ipv6Address, not even an explicit nil
### GetIpSubnet

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetIpSubnet() string`

GetIpSubnet returns the IpSubnet field if non-nil, zero value otherwise.

### GetIpSubnetOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetIpSubnetOk() (*string, bool)`

GetIpSubnetOk returns a tuple with the IpSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSubnet

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetIpSubnet(v string)`

SetIpSubnet sets IpSubnet field to given value.

### HasIpSubnet

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasIpSubnet() bool`

HasIpSubnet returns a boolean if a field has been set.

### SetIpSubnetNil

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetIpSubnetNil(b bool)`

 SetIpSubnetNil sets the value for IpSubnet to be an explicit nil

### UnsetIpSubnet
`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) UnsetIpSubnet()`

UnsetIpSubnet ensures that no value is present for IpSubnet, not even an explicit nil
### GetIpv6Subnet

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetIpv6Subnet() string`

GetIpv6Subnet returns the Ipv6Subnet field if non-nil, zero value otherwise.

### GetIpv6SubnetOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetIpv6SubnetOk() (*string, bool)`

GetIpv6SubnetOk returns a tuple with the Ipv6Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Subnet

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetIpv6Subnet(v string)`

SetIpv6Subnet sets Ipv6Subnet field to given value.

### HasIpv6Subnet

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasIpv6Subnet() bool`

HasIpv6Subnet returns a boolean if a field has been set.

### SetIpv6SubnetNil

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetIpv6SubnetNil(b bool)`

 SetIpv6SubnetNil sets the value for Ipv6Subnet to be an explicit nil

### UnsetIpv6Subnet
`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) UnsetIpv6Subnet()`

UnsetIpv6Subnet ensures that no value is present for Ipv6Subnet, not even an explicit nil
### GetDescription

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDhcp

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetDhcp() bool`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetDhcpOk() (*bool, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetDhcp(v bool)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### GetActive

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPoolAssigned

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetPoolAssigned() bool`

GetPoolAssigned returns the PoolAssigned field if non-nil, zero value otherwise.

### GetPoolAssignedOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetPoolAssignedOk() (*bool, bool)`

GetPoolAssignedOk returns a tuple with the PoolAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAssigned

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetPoolAssigned(v bool)`

SetPoolAssigned sets PoolAssigned field to given value.

### HasPoolAssigned

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasPoolAssigned() bool`

HasPoolAssigned returns a boolean if a field has been set.

### GetPrimaryInterface

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetPrimaryInterface() bool`

GetPrimaryInterface returns the PrimaryInterface field if non-nil, zero value otherwise.

### GetPrimaryInterfaceOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetPrimaryInterfaceOk() (*bool, bool)`

GetPrimaryInterfaceOk returns a tuple with the PrimaryInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInterface

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetPrimaryInterface(v bool)`

SetPrimaryInterface sets PrimaryInterface field to given value.

### HasPrimaryInterface

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasPrimaryInterface() bool`

HasPrimaryInterface returns a boolean if a field has been set.

### GetNetwork

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetNetwork() UpdateHostResize200ResponseAllOfServerInterfacesInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetNetworkOk() (*UpdateHostResize200ResponseAllOfServerInterfacesInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetNetwork(v UpdateHostResize200ResponseAllOfServerInterfacesInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSubnet

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetNetworkGroup

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetNetworkGroup() string`

GetNetworkGroup returns the NetworkGroup field if non-nil, zero value otherwise.

### GetNetworkGroupOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetNetworkGroupOk() (*string, bool)`

GetNetworkGroupOk returns a tuple with the NetworkGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkGroup

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetNetworkGroup(v string)`

SetNetworkGroup sets NetworkGroup field to given value.

### HasNetworkGroup

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasNetworkGroup() bool`

HasNetworkGroup returns a boolean if a field has been set.

### SetNetworkGroupNil

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetNetworkGroupNil(b bool)`

 SetNetworkGroupNil sets the value for NetworkGroup to be an explicit nil

### UnsetNetworkGroup
`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) UnsetNetworkGroup()`

UnsetNetworkGroup ensures that no value is present for NetworkGroup, not even an explicit nil
### GetNetworkPosition

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetNetworkPosition() string`

GetNetworkPosition returns the NetworkPosition field if non-nil, zero value otherwise.

### GetNetworkPositionOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetNetworkPositionOk() (*string, bool)`

GetNetworkPositionOk returns a tuple with the NetworkPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPosition

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetNetworkPosition(v string)`

SetNetworkPosition sets NetworkPosition field to given value.

### HasNetworkPosition

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasNetworkPosition() bool`

HasNetworkPosition returns a boolean if a field has been set.

### SetNetworkPositionNil

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetNetworkPositionNil(b bool)`

 SetNetworkPositionNil sets the value for NetworkPosition to be an explicit nil

### UnsetNetworkPosition
`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) UnsetNetworkPosition()`

UnsetNetworkPosition ensures that no value is present for NetworkPosition, not even an explicit nil
### GetNetworkPool

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetNetworkPool() UpdateHostResize200ResponseAllOfServerInterfacesInnerNetworkPool`

GetNetworkPool returns the NetworkPool field if non-nil, zero value otherwise.

### GetNetworkPoolOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetNetworkPoolOk() (*UpdateHostResize200ResponseAllOfServerInterfacesInnerNetworkPool, bool)`

GetNetworkPoolOk returns a tuple with the NetworkPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPool

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetNetworkPool(v UpdateHostResize200ResponseAllOfServerInterfacesInnerNetworkPool)`

SetNetworkPool sets NetworkPool field to given value.

### HasNetworkPool

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasNetworkPool() bool`

HasNetworkPool returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetNetworkDomain() UpdateHostResize200ResponseAllOfServerInterfacesInnerNetworkDomain`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetNetworkDomainOk() (*UpdateHostResize200ResponseAllOfServerInterfacesInnerNetworkDomain, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetNetworkDomain(v UpdateHostResize200ResponseAllOfServerInterfacesInnerNetworkDomain)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetType

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetType() UpdateHostResize200ResponseAllOfServerInterfacesInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetTypeOk() (*UpdateHostResize200ResponseAllOfServerInterfacesInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetType(v UpdateHostResize200ResponseAllOfServerInterfacesInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIpMode

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### SetIpModeNil

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetIpModeNil(b bool)`

 SetIpModeNil sets the value for IpMode to be an explicit nil

### UnsetIpMode
`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) UnsetIpMode()`

UnsetIpMode ensures that no value is present for IpMode, not even an explicit nil
### GetMacAddress

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *UpdateHostResize200ResponseAllOfServerInterfacesInner) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


