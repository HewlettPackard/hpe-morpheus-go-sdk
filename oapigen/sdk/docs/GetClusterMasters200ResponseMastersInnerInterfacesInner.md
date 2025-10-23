# GetClusterMasters200ResponseMastersInnerInterfacesInner

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
**Network** | Pointer to [**NullableListApprovals200ResponseAllOfApprovalsInnerAccount**](ListApprovals200ResponseAllOfApprovalsInnerAccount.md) |  | [optional] 
**Subnet** | Pointer to **NullableString** |  | [optional] 
**NetworkGroup** | Pointer to **NullableString** |  | [optional] 
**NetworkPosition** | Pointer to **NullableString** |  | [optional] 
**NetworkPool** | Pointer to **NullableString** |  | [optional] 
**NetworkDomain** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**IpMode** | Pointer to **NullableString** |  | [optional] 
**MacAddress** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetClusterMasters200ResponseMastersInnerInterfacesInner

`func NewGetClusterMasters200ResponseMastersInnerInterfacesInner() *GetClusterMasters200ResponseMastersInnerInterfacesInner`

NewGetClusterMasters200ResponseMastersInnerInterfacesInner instantiates a new GetClusterMasters200ResponseMastersInnerInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterMasters200ResponseMastersInnerInterfacesInnerWithDefaults

`func NewGetClusterMasters200ResponseMastersInnerInterfacesInnerWithDefaults() *GetClusterMasters200ResponseMastersInnerInterfacesInner`

NewGetClusterMasters200ResponseMastersInnerInterfacesInnerWithDefaults instantiates a new GetClusterMasters200ResponseMastersInnerInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRefType

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetName

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInternalId

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetUniqueId

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### SetUniqueIdNil

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetUniqueIdNil(b bool)`

 SetUniqueIdNil sets the value for UniqueId to be an explicit nil

### UnsetUniqueId
`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) UnsetUniqueId()`

UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
### GetPublicIpAddress

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### GetPublicIpv6Address

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetPublicIpv6Address() string`

GetPublicIpv6Address returns the PublicIpv6Address field if non-nil, zero value otherwise.

### GetPublicIpv6AddressOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetPublicIpv6AddressOk() (*string, bool)`

GetPublicIpv6AddressOk returns a tuple with the PublicIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpv6Address

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetPublicIpv6Address(v string)`

SetPublicIpv6Address sets PublicIpv6Address field to given value.

### HasPublicIpv6Address

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasPublicIpv6Address() bool`

HasPublicIpv6Address returns a boolean if a field has been set.

### SetPublicIpv6AddressNil

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetPublicIpv6AddressNil(b bool)`

 SetPublicIpv6AddressNil sets the value for PublicIpv6Address to be an explicit nil

### UnsetPublicIpv6Address
`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) UnsetPublicIpv6Address()`

UnsetPublicIpv6Address ensures that no value is present for PublicIpv6Address, not even an explicit nil
### GetIpAddress

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpv6Address

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### SetIpv6AddressNil

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetIpv6AddressNil(b bool)`

 SetIpv6AddressNil sets the value for Ipv6Address to be an explicit nil

### UnsetIpv6Address
`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) UnsetIpv6Address()`

UnsetIpv6Address ensures that no value is present for Ipv6Address, not even an explicit nil
### GetIpSubnet

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetIpSubnet() string`

GetIpSubnet returns the IpSubnet field if non-nil, zero value otherwise.

### GetIpSubnetOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetIpSubnetOk() (*string, bool)`

GetIpSubnetOk returns a tuple with the IpSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSubnet

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetIpSubnet(v string)`

SetIpSubnet sets IpSubnet field to given value.

### HasIpSubnet

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasIpSubnet() bool`

HasIpSubnet returns a boolean if a field has been set.

### SetIpSubnetNil

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetIpSubnetNil(b bool)`

 SetIpSubnetNil sets the value for IpSubnet to be an explicit nil

### UnsetIpSubnet
`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) UnsetIpSubnet()`

UnsetIpSubnet ensures that no value is present for IpSubnet, not even an explicit nil
### GetIpv6Subnet

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetIpv6Subnet() string`

GetIpv6Subnet returns the Ipv6Subnet field if non-nil, zero value otherwise.

### GetIpv6SubnetOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetIpv6SubnetOk() (*string, bool)`

GetIpv6SubnetOk returns a tuple with the Ipv6Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Subnet

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetIpv6Subnet(v string)`

SetIpv6Subnet sets Ipv6Subnet field to given value.

### HasIpv6Subnet

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasIpv6Subnet() bool`

HasIpv6Subnet returns a boolean if a field has been set.

### SetIpv6SubnetNil

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetIpv6SubnetNil(b bool)`

 SetIpv6SubnetNil sets the value for Ipv6Subnet to be an explicit nil

### UnsetIpv6Subnet
`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) UnsetIpv6Subnet()`

UnsetIpv6Subnet ensures that no value is present for Ipv6Subnet, not even an explicit nil
### GetDescription

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDhcp

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetDhcp() bool`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetDhcpOk() (*bool, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetDhcp(v bool)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### GetActive

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPoolAssigned

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetPoolAssigned() bool`

GetPoolAssigned returns the PoolAssigned field if non-nil, zero value otherwise.

### GetPoolAssignedOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetPoolAssignedOk() (*bool, bool)`

GetPoolAssignedOk returns a tuple with the PoolAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAssigned

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetPoolAssigned(v bool)`

SetPoolAssigned sets PoolAssigned field to given value.

### HasPoolAssigned

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasPoolAssigned() bool`

HasPoolAssigned returns a boolean if a field has been set.

### GetPrimaryInterface

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetPrimaryInterface() bool`

GetPrimaryInterface returns the PrimaryInterface field if non-nil, zero value otherwise.

### GetPrimaryInterfaceOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetPrimaryInterfaceOk() (*bool, bool)`

GetPrimaryInterfaceOk returns a tuple with the PrimaryInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInterface

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetPrimaryInterface(v bool)`

SetPrimaryInterface sets PrimaryInterface field to given value.

### HasPrimaryInterface

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasPrimaryInterface() bool`

HasPrimaryInterface returns a boolean if a field has been set.

### GetNetwork

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetNetwork() ListApprovals200ResponseAllOfApprovalsInnerAccount`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetNetworkOk() (*ListApprovals200ResponseAllOfApprovalsInnerAccount, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetNetwork(v ListApprovals200ResponseAllOfApprovalsInnerAccount)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### SetNetworkNil

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetSubnet

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetNetworkGroup

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetNetworkGroup() string`

GetNetworkGroup returns the NetworkGroup field if non-nil, zero value otherwise.

### GetNetworkGroupOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetNetworkGroupOk() (*string, bool)`

GetNetworkGroupOk returns a tuple with the NetworkGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkGroup

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetNetworkGroup(v string)`

SetNetworkGroup sets NetworkGroup field to given value.

### HasNetworkGroup

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasNetworkGroup() bool`

HasNetworkGroup returns a boolean if a field has been set.

### SetNetworkGroupNil

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetNetworkGroupNil(b bool)`

 SetNetworkGroupNil sets the value for NetworkGroup to be an explicit nil

### UnsetNetworkGroup
`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) UnsetNetworkGroup()`

UnsetNetworkGroup ensures that no value is present for NetworkGroup, not even an explicit nil
### GetNetworkPosition

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetNetworkPosition() string`

GetNetworkPosition returns the NetworkPosition field if non-nil, zero value otherwise.

### GetNetworkPositionOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetNetworkPositionOk() (*string, bool)`

GetNetworkPositionOk returns a tuple with the NetworkPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPosition

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetNetworkPosition(v string)`

SetNetworkPosition sets NetworkPosition field to given value.

### HasNetworkPosition

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasNetworkPosition() bool`

HasNetworkPosition returns a boolean if a field has been set.

### SetNetworkPositionNil

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetNetworkPositionNil(b bool)`

 SetNetworkPositionNil sets the value for NetworkPosition to be an explicit nil

### UnsetNetworkPosition
`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) UnsetNetworkPosition()`

UnsetNetworkPosition ensures that no value is present for NetworkPosition, not even an explicit nil
### GetNetworkPool

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetNetworkPool() string`

GetNetworkPool returns the NetworkPool field if non-nil, zero value otherwise.

### GetNetworkPoolOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetNetworkPoolOk() (*string, bool)`

GetNetworkPoolOk returns a tuple with the NetworkPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPool

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetNetworkPool(v string)`

SetNetworkPool sets NetworkPool field to given value.

### HasNetworkPool

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasNetworkPool() bool`

HasNetworkPool returns a boolean if a field has been set.

### SetNetworkPoolNil

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetNetworkPoolNil(b bool)`

 SetNetworkPoolNil sets the value for NetworkPool to be an explicit nil

### UnsetNetworkPool
`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) UnsetNetworkPool()`

UnsetNetworkPool ensures that no value is present for NetworkPool, not even an explicit nil
### GetNetworkDomain

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetNetworkDomain() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetNetworkDomainOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetNetworkDomain(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetType

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetIpMode

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### SetIpModeNil

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetIpModeNil(b bool)`

 SetIpModeNil sets the value for IpMode to be an explicit nil

### UnsetIpMode
`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) UnsetIpMode()`

UnsetIpMode ensures that no value is present for IpMode, not even an explicit nil
### GetMacAddress

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *GetClusterMasters200ResponseMastersInnerInterfacesInner) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


