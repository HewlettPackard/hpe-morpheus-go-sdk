# ListHosts200ResponseAllOfServersInnerInterfacesInner

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
**Network** | Pointer to [**NullableListApprovals200ResponseAllOfApprovalsInnerAccount**](ListApprovals200ResponseAllOfApprovalsInnerAccount.md) |  | [optional] 
**Subnet** | Pointer to **NullableString** |  | [optional] 
**NetworkGroup** | Pointer to **NullableString** |  | [optional] 
**NetworkPosition** | Pointer to **NullableString** |  | [optional] 
**NetworkPool** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**NetworkDomain** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**IpMode** | Pointer to **NullableString** |  | [optional] 
**MacAddress** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListHosts200ResponseAllOfServersInnerInterfacesInner

`func NewListHosts200ResponseAllOfServersInnerInterfacesInner() *ListHosts200ResponseAllOfServersInnerInterfacesInner`

NewListHosts200ResponseAllOfServersInnerInterfacesInner instantiates a new ListHosts200ResponseAllOfServersInnerInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHosts200ResponseAllOfServersInnerInterfacesInnerWithDefaults

`func NewListHosts200ResponseAllOfServersInnerInterfacesInnerWithDefaults() *ListHosts200ResponseAllOfServersInnerInterfacesInner`

NewListHosts200ResponseAllOfServersInnerInterfacesInnerWithDefaults instantiates a new ListHosts200ResponseAllOfServersInnerInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRefType

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetName

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInternalId

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetUniqueId

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### SetUniqueIdNil

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetUniqueIdNil(b bool)`

 SetUniqueIdNil sets the value for UniqueId to be an explicit nil

### UnsetUniqueId
`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) UnsetUniqueId()`

UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
### GetPublicIpAddress

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### SetPublicIpAddressNil

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetPublicIpAddressNil(b bool)`

 SetPublicIpAddressNil sets the value for PublicIpAddress to be an explicit nil

### UnsetPublicIpAddress
`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) UnsetPublicIpAddress()`

UnsetPublicIpAddress ensures that no value is present for PublicIpAddress, not even an explicit nil
### GetPublicIpv6Address

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetPublicIpv6Address() string`

GetPublicIpv6Address returns the PublicIpv6Address field if non-nil, zero value otherwise.

### GetPublicIpv6AddressOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetPublicIpv6AddressOk() (*string, bool)`

GetPublicIpv6AddressOk returns a tuple with the PublicIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpv6Address

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetPublicIpv6Address(v string)`

SetPublicIpv6Address sets PublicIpv6Address field to given value.

### HasPublicIpv6Address

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasPublicIpv6Address() bool`

HasPublicIpv6Address returns a boolean if a field has been set.

### SetPublicIpv6AddressNil

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetPublicIpv6AddressNil(b bool)`

 SetPublicIpv6AddressNil sets the value for PublicIpv6Address to be an explicit nil

### UnsetPublicIpv6Address
`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) UnsetPublicIpv6Address()`

UnsetPublicIpv6Address ensures that no value is present for PublicIpv6Address, not even an explicit nil
### GetIpAddress

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpv6Address

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### SetIpv6AddressNil

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetIpv6AddressNil(b bool)`

 SetIpv6AddressNil sets the value for Ipv6Address to be an explicit nil

### UnsetIpv6Address
`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) UnsetIpv6Address()`

UnsetIpv6Address ensures that no value is present for Ipv6Address, not even an explicit nil
### GetIpSubnet

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIpSubnet() string`

GetIpSubnet returns the IpSubnet field if non-nil, zero value otherwise.

### GetIpSubnetOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIpSubnetOk() (*string, bool)`

GetIpSubnetOk returns a tuple with the IpSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSubnet

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetIpSubnet(v string)`

SetIpSubnet sets IpSubnet field to given value.

### HasIpSubnet

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasIpSubnet() bool`

HasIpSubnet returns a boolean if a field has been set.

### SetIpSubnetNil

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetIpSubnetNil(b bool)`

 SetIpSubnetNil sets the value for IpSubnet to be an explicit nil

### UnsetIpSubnet
`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) UnsetIpSubnet()`

UnsetIpSubnet ensures that no value is present for IpSubnet, not even an explicit nil
### GetIpv6Subnet

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIpv6Subnet() string`

GetIpv6Subnet returns the Ipv6Subnet field if non-nil, zero value otherwise.

### GetIpv6SubnetOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIpv6SubnetOk() (*string, bool)`

GetIpv6SubnetOk returns a tuple with the Ipv6Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Subnet

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetIpv6Subnet(v string)`

SetIpv6Subnet sets Ipv6Subnet field to given value.

### HasIpv6Subnet

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasIpv6Subnet() bool`

HasIpv6Subnet returns a boolean if a field has been set.

### SetIpv6SubnetNil

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetIpv6SubnetNil(b bool)`

 SetIpv6SubnetNil sets the value for Ipv6Subnet to be an explicit nil

### UnsetIpv6Subnet
`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) UnsetIpv6Subnet()`

UnsetIpv6Subnet ensures that no value is present for Ipv6Subnet, not even an explicit nil
### GetDescription

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDhcp

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetDhcp() bool`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetDhcpOk() (*bool, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetDhcp(v bool)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### GetActive

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPoolAssigned

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetPoolAssigned() bool`

GetPoolAssigned returns the PoolAssigned field if non-nil, zero value otherwise.

### GetPoolAssignedOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetPoolAssignedOk() (*bool, bool)`

GetPoolAssignedOk returns a tuple with the PoolAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAssigned

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetPoolAssigned(v bool)`

SetPoolAssigned sets PoolAssigned field to given value.

### HasPoolAssigned

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasPoolAssigned() bool`

HasPoolAssigned returns a boolean if a field has been set.

### GetPrimaryInterface

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetPrimaryInterface() bool`

GetPrimaryInterface returns the PrimaryInterface field if non-nil, zero value otherwise.

### GetPrimaryInterfaceOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetPrimaryInterfaceOk() (*bool, bool)`

GetPrimaryInterfaceOk returns a tuple with the PrimaryInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInterface

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetPrimaryInterface(v bool)`

SetPrimaryInterface sets PrimaryInterface field to given value.

### HasPrimaryInterface

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasPrimaryInterface() bool`

HasPrimaryInterface returns a boolean if a field has been set.

### GetNetwork

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetwork() ListApprovals200ResponseAllOfApprovalsInnerAccount`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkOk() (*ListApprovals200ResponseAllOfApprovalsInnerAccount, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetNetwork(v ListApprovals200ResponseAllOfApprovalsInnerAccount)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### SetNetworkNil

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetSubnet

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetNetworkGroup

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkGroup() string`

GetNetworkGroup returns the NetworkGroup field if non-nil, zero value otherwise.

### GetNetworkGroupOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkGroupOk() (*string, bool)`

GetNetworkGroupOk returns a tuple with the NetworkGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkGroup

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetNetworkGroup(v string)`

SetNetworkGroup sets NetworkGroup field to given value.

### HasNetworkGroup

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasNetworkGroup() bool`

HasNetworkGroup returns a boolean if a field has been set.

### SetNetworkGroupNil

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetNetworkGroupNil(b bool)`

 SetNetworkGroupNil sets the value for NetworkGroup to be an explicit nil

### UnsetNetworkGroup
`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) UnsetNetworkGroup()`

UnsetNetworkGroup ensures that no value is present for NetworkGroup, not even an explicit nil
### GetNetworkPosition

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkPosition() string`

GetNetworkPosition returns the NetworkPosition field if non-nil, zero value otherwise.

### GetNetworkPositionOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkPositionOk() (*string, bool)`

GetNetworkPositionOk returns a tuple with the NetworkPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPosition

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetNetworkPosition(v string)`

SetNetworkPosition sets NetworkPosition field to given value.

### HasNetworkPosition

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasNetworkPosition() bool`

HasNetworkPosition returns a boolean if a field has been set.

### SetNetworkPositionNil

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetNetworkPositionNil(b bool)`

 SetNetworkPositionNil sets the value for NetworkPosition to be an explicit nil

### UnsetNetworkPosition
`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) UnsetNetworkPosition()`

UnsetNetworkPosition ensures that no value is present for NetworkPosition, not even an explicit nil
### GetNetworkPool

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkPool() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetNetworkPool returns the NetworkPool field if non-nil, zero value otherwise.

### GetNetworkPoolOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkPoolOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetNetworkPoolOk returns a tuple with the NetworkPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPool

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetNetworkPool(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetNetworkPool sets NetworkPool field to given value.

### HasNetworkPool

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasNetworkPool() bool`

HasNetworkPool returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkDomain() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkDomainOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetNetworkDomain(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetType

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIpMode

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### SetIpModeNil

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetIpModeNil(b bool)`

 SetIpModeNil sets the value for IpMode to be an explicit nil

### UnsetIpMode
`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) UnsetIpMode()`

UnsetIpMode ensures that no value is present for IpMode, not even an explicit nil
### GetMacAddress

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


