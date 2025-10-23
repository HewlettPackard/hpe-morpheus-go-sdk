# ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner

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
**PublicIpAddress** | Pointer to **string** |  | [optional] 
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
**Network** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Subnet** | Pointer to **NullableString** |  | [optional] 
**NetworkGroup** | Pointer to **NullableString** |  | [optional] 
**NetworkPosition** | Pointer to **NullableString** |  | [optional] 
**NetworkPool** | Pointer to **NullableString** |  | [optional] 
**NetworkDomain** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**IpMode** | Pointer to **NullableString** |  | [optional] 
**MacAddress** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner

`func NewListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner() *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner`

NewListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner instantiates a new ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInnerWithDefaults

`func NewListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInnerWithDefaults() *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner`

NewListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInnerWithDefaults instantiates a new ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRefType

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetName

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInternalId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetUniqueId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### SetUniqueIdNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetUniqueIdNil(b bool)`

 SetUniqueIdNil sets the value for UniqueId to be an explicit nil

### UnsetUniqueId
`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) UnsetUniqueId()`

UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
### GetPublicIpAddress

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### GetPublicIpv6Address

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetPublicIpv6Address() string`

GetPublicIpv6Address returns the PublicIpv6Address field if non-nil, zero value otherwise.

### GetPublicIpv6AddressOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetPublicIpv6AddressOk() (*string, bool)`

GetPublicIpv6AddressOk returns a tuple with the PublicIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpv6Address

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetPublicIpv6Address(v string)`

SetPublicIpv6Address sets PublicIpv6Address field to given value.

### HasPublicIpv6Address

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasPublicIpv6Address() bool`

HasPublicIpv6Address returns a boolean if a field has been set.

### SetPublicIpv6AddressNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetPublicIpv6AddressNil(b bool)`

 SetPublicIpv6AddressNil sets the value for PublicIpv6Address to be an explicit nil

### UnsetPublicIpv6Address
`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) UnsetPublicIpv6Address()`

UnsetPublicIpv6Address ensures that no value is present for PublicIpv6Address, not even an explicit nil
### GetIpAddress

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpv6Address

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### SetIpv6AddressNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetIpv6AddressNil(b bool)`

 SetIpv6AddressNil sets the value for Ipv6Address to be an explicit nil

### UnsetIpv6Address
`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) UnsetIpv6Address()`

UnsetIpv6Address ensures that no value is present for Ipv6Address, not even an explicit nil
### GetIpSubnet

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetIpSubnet() string`

GetIpSubnet returns the IpSubnet field if non-nil, zero value otherwise.

### GetIpSubnetOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetIpSubnetOk() (*string, bool)`

GetIpSubnetOk returns a tuple with the IpSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSubnet

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetIpSubnet(v string)`

SetIpSubnet sets IpSubnet field to given value.

### HasIpSubnet

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasIpSubnet() bool`

HasIpSubnet returns a boolean if a field has been set.

### SetIpSubnetNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetIpSubnetNil(b bool)`

 SetIpSubnetNil sets the value for IpSubnet to be an explicit nil

### UnsetIpSubnet
`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) UnsetIpSubnet()`

UnsetIpSubnet ensures that no value is present for IpSubnet, not even an explicit nil
### GetIpv6Subnet

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetIpv6Subnet() string`

GetIpv6Subnet returns the Ipv6Subnet field if non-nil, zero value otherwise.

### GetIpv6SubnetOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetIpv6SubnetOk() (*string, bool)`

GetIpv6SubnetOk returns a tuple with the Ipv6Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Subnet

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetIpv6Subnet(v string)`

SetIpv6Subnet sets Ipv6Subnet field to given value.

### HasIpv6Subnet

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasIpv6Subnet() bool`

HasIpv6Subnet returns a boolean if a field has been set.

### SetIpv6SubnetNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetIpv6SubnetNil(b bool)`

 SetIpv6SubnetNil sets the value for Ipv6Subnet to be an explicit nil

### UnsetIpv6Subnet
`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) UnsetIpv6Subnet()`

UnsetIpv6Subnet ensures that no value is present for Ipv6Subnet, not even an explicit nil
### GetDescription

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDhcp

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetDhcp() bool`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetDhcpOk() (*bool, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetDhcp(v bool)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### GetActive

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPoolAssigned

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetPoolAssigned() bool`

GetPoolAssigned returns the PoolAssigned field if non-nil, zero value otherwise.

### GetPoolAssignedOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetPoolAssignedOk() (*bool, bool)`

GetPoolAssignedOk returns a tuple with the PoolAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAssigned

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetPoolAssigned(v bool)`

SetPoolAssigned sets PoolAssigned field to given value.

### HasPoolAssigned

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasPoolAssigned() bool`

HasPoolAssigned returns a boolean if a field has been set.

### GetPrimaryInterface

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetPrimaryInterface() bool`

GetPrimaryInterface returns the PrimaryInterface field if non-nil, zero value otherwise.

### GetPrimaryInterfaceOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetPrimaryInterfaceOk() (*bool, bool)`

GetPrimaryInterfaceOk returns a tuple with the PrimaryInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInterface

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetPrimaryInterface(v bool)`

SetPrimaryInterface sets PrimaryInterface field to given value.

### HasPrimaryInterface

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasPrimaryInterface() bool`

HasPrimaryInterface returns a boolean if a field has been set.

### GetNetwork

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetNetwork() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetNetworkOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetNetwork(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSubnet

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetNetworkGroup

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetNetworkGroup() string`

GetNetworkGroup returns the NetworkGroup field if non-nil, zero value otherwise.

### GetNetworkGroupOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetNetworkGroupOk() (*string, bool)`

GetNetworkGroupOk returns a tuple with the NetworkGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkGroup

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetNetworkGroup(v string)`

SetNetworkGroup sets NetworkGroup field to given value.

### HasNetworkGroup

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasNetworkGroup() bool`

HasNetworkGroup returns a boolean if a field has been set.

### SetNetworkGroupNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetNetworkGroupNil(b bool)`

 SetNetworkGroupNil sets the value for NetworkGroup to be an explicit nil

### UnsetNetworkGroup
`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) UnsetNetworkGroup()`

UnsetNetworkGroup ensures that no value is present for NetworkGroup, not even an explicit nil
### GetNetworkPosition

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetNetworkPosition() string`

GetNetworkPosition returns the NetworkPosition field if non-nil, zero value otherwise.

### GetNetworkPositionOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetNetworkPositionOk() (*string, bool)`

GetNetworkPositionOk returns a tuple with the NetworkPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPosition

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetNetworkPosition(v string)`

SetNetworkPosition sets NetworkPosition field to given value.

### HasNetworkPosition

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasNetworkPosition() bool`

HasNetworkPosition returns a boolean if a field has been set.

### SetNetworkPositionNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetNetworkPositionNil(b bool)`

 SetNetworkPositionNil sets the value for NetworkPosition to be an explicit nil

### UnsetNetworkPosition
`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) UnsetNetworkPosition()`

UnsetNetworkPosition ensures that no value is present for NetworkPosition, not even an explicit nil
### GetNetworkPool

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetNetworkPool() string`

GetNetworkPool returns the NetworkPool field if non-nil, zero value otherwise.

### GetNetworkPoolOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetNetworkPoolOk() (*string, bool)`

GetNetworkPoolOk returns a tuple with the NetworkPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPool

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetNetworkPool(v string)`

SetNetworkPool sets NetworkPool field to given value.

### HasNetworkPool

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasNetworkPool() bool`

HasNetworkPool returns a boolean if a field has been set.

### SetNetworkPoolNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetNetworkPoolNil(b bool)`

 SetNetworkPoolNil sets the value for NetworkPool to be an explicit nil

### UnsetNetworkPool
`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) UnsetNetworkPool()`

UnsetNetworkPool ensures that no value is present for NetworkPool, not even an explicit nil
### GetNetworkDomain

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetNetworkDomain() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetNetworkDomainOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetNetworkDomain(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetType

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetIpMode

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### SetIpModeNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetIpModeNil(b bool)`

 SetIpModeNil sets the value for IpMode to be an explicit nil

### UnsetIpMode
`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) UnsetIpMode()`

UnsetIpMode ensures that no value is present for IpMode, not even an explicit nil
### GetMacAddress

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


