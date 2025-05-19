# ListHosts200ResponseAllOfServersInnerInterfacesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**PublicIpAddress** | Pointer to **string** |  | [optional] 
**PublicIpv6Address** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Ipv6Address** | Pointer to **string** |  | [optional] 
**IpSubnet** | Pointer to **string** |  | [optional] 
**Ipv6Subnet** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Dhcp** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**PoolAssigned** | Pointer to **bool** |  | [optional] 
**PrimaryInterface** | Pointer to **bool** |  | [optional] 
**Network** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Subnet** | Pointer to **string** |  | [optional] 
**NetworkGroup** | Pointer to **string** |  | [optional] 
**NetworkPosition** | Pointer to **string** |  | [optional] 
**NetworkPool** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**NetworkDomain** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType**](ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType.md) |  | [optional] 
**IpMode** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 

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

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetwork() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetNetwork(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

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

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkDomain() string`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkDomainOk() (*string, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetNetworkDomain(v string)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetType

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetType() ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetTypeOk() (*ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetType(v ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


