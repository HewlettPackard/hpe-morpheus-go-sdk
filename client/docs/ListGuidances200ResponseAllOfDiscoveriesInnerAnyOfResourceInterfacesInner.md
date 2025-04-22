# ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner

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
**NetworkPool** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**NetworkDomain** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**IpMode** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 

## Methods

### NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner

`func NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner() *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner`

NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner instantiates a new ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInnerWithDefaults

`func NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInnerWithDefaults() *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner`

NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInnerWithDefaults instantiates a new ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRefType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetName

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInternalId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetUniqueId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetPublicIpAddress

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### GetPublicIpv6Address

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetPublicIpv6Address() string`

GetPublicIpv6Address returns the PublicIpv6Address field if non-nil, zero value otherwise.

### GetPublicIpv6AddressOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetPublicIpv6AddressOk() (*string, bool)`

GetPublicIpv6AddressOk returns a tuple with the PublicIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpv6Address

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetPublicIpv6Address(v string)`

SetPublicIpv6Address sets PublicIpv6Address field to given value.

### HasPublicIpv6Address

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasPublicIpv6Address() bool`

HasPublicIpv6Address returns a boolean if a field has been set.

### GetIpAddress

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpv6Address

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### GetIpSubnet

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetIpSubnet() string`

GetIpSubnet returns the IpSubnet field if non-nil, zero value otherwise.

### GetIpSubnetOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetIpSubnetOk() (*string, bool)`

GetIpSubnetOk returns a tuple with the IpSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSubnet

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetIpSubnet(v string)`

SetIpSubnet sets IpSubnet field to given value.

### HasIpSubnet

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasIpSubnet() bool`

HasIpSubnet returns a boolean if a field has been set.

### GetIpv6Subnet

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetIpv6Subnet() string`

GetIpv6Subnet returns the Ipv6Subnet field if non-nil, zero value otherwise.

### GetIpv6SubnetOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetIpv6SubnetOk() (*string, bool)`

GetIpv6SubnetOk returns a tuple with the Ipv6Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Subnet

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetIpv6Subnet(v string)`

SetIpv6Subnet sets Ipv6Subnet field to given value.

### HasIpv6Subnet

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasIpv6Subnet() bool`

HasIpv6Subnet returns a boolean if a field has been set.

### GetDescription

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDhcp

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetDhcp() bool`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetDhcpOk() (*bool, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetDhcp(v bool)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### GetActive

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPoolAssigned

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetPoolAssigned() bool`

GetPoolAssigned returns the PoolAssigned field if non-nil, zero value otherwise.

### GetPoolAssignedOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetPoolAssignedOk() (*bool, bool)`

GetPoolAssignedOk returns a tuple with the PoolAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAssigned

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetPoolAssigned(v bool)`

SetPoolAssigned sets PoolAssigned field to given value.

### HasPoolAssigned

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasPoolAssigned() bool`

HasPoolAssigned returns a boolean if a field has been set.

### GetPrimaryInterface

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetPrimaryInterface() bool`

GetPrimaryInterface returns the PrimaryInterface field if non-nil, zero value otherwise.

### GetPrimaryInterfaceOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetPrimaryInterfaceOk() (*bool, bool)`

GetPrimaryInterfaceOk returns a tuple with the PrimaryInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInterface

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetPrimaryInterface(v bool)`

SetPrimaryInterface sets PrimaryInterface field to given value.

### HasPrimaryInterface

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasPrimaryInterface() bool`

HasPrimaryInterface returns a boolean if a field has been set.

### GetNetwork

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetNetwork() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetNetworkOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetNetwork(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSubnet

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetNetworkGroup

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetNetworkGroup() string`

GetNetworkGroup returns the NetworkGroup field if non-nil, zero value otherwise.

### GetNetworkGroupOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetNetworkGroupOk() (*string, bool)`

GetNetworkGroupOk returns a tuple with the NetworkGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkGroup

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetNetworkGroup(v string)`

SetNetworkGroup sets NetworkGroup field to given value.

### HasNetworkGroup

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasNetworkGroup() bool`

HasNetworkGroup returns a boolean if a field has been set.

### GetNetworkPosition

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetNetworkPosition() string`

GetNetworkPosition returns the NetworkPosition field if non-nil, zero value otherwise.

### GetNetworkPositionOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetNetworkPositionOk() (*string, bool)`

GetNetworkPositionOk returns a tuple with the NetworkPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPosition

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetNetworkPosition(v string)`

SetNetworkPosition sets NetworkPosition field to given value.

### HasNetworkPosition

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasNetworkPosition() bool`

HasNetworkPosition returns a boolean if a field has been set.

### GetNetworkPool

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetNetworkPool() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetNetworkPool returns the NetworkPool field if non-nil, zero value otherwise.

### GetNetworkPoolOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetNetworkPoolOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetNetworkPoolOk returns a tuple with the NetworkPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPool

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetNetworkPool(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetNetworkPool sets NetworkPool field to given value.

### HasNetworkPool

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasNetworkPool() bool`

HasNetworkPool returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetNetworkDomain() string`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetNetworkDomainOk() (*string, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetNetworkDomain(v string)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIpMode

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetMacAddress

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


