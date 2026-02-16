# InstanceContainerServerInterfacesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**PublicIpAddress** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Dhcp** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**PoolAssigned** | Pointer to **bool** |  | [optional] 
**PrimaryInterface** | Pointer to **bool** |  | [optional] 
**Network** | Pointer to [**InstanceContainerServerInterfacesInnerNetwork**](InstanceContainerServerInterfacesInnerNetwork.md) |  | [optional] 
**NetworkGroup** | Pointer to [**InstanceContainerServerInterfacesInnerNetworkGroup**](InstanceContainerServerInterfacesInnerNetworkGroup.md) |  | [optional] 
**NetworkPool** | Pointer to [**InstanceContainerServerInterfacesInnerNetworkPool**](InstanceContainerServerInterfacesInnerNetworkPool.md) |  | [optional] 
**IpMode** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**Interfaces** | Pointer to [**[]InstanceContainerServerInstancesInnerInner**](InstanceContainerServerInstancesInnerInner.md) |  | [optional] 

## Methods

### NewInstanceContainerServerInterfacesInner

`func NewInstanceContainerServerInterfacesInner() *InstanceContainerServerInterfacesInner`

NewInstanceContainerServerInterfacesInner instantiates a new InstanceContainerServerInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceContainerServerInterfacesInnerWithDefaults

`func NewInstanceContainerServerInterfacesInnerWithDefaults() *InstanceContainerServerInterfacesInner`

NewInstanceContainerServerInterfacesInnerWithDefaults instantiates a new InstanceContainerServerInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceContainerServerInterfacesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceContainerServerInterfacesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceContainerServerInterfacesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceContainerServerInterfacesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InstanceContainerServerInterfacesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceContainerServerInterfacesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceContainerServerInterfacesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceContainerServerInterfacesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUniqueId

`func (o *InstanceContainerServerInterfacesInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *InstanceContainerServerInterfacesInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *InstanceContainerServerInterfacesInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *InstanceContainerServerInterfacesInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetPublicIpAddress

`func (o *InstanceContainerServerInterfacesInner) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *InstanceContainerServerInterfacesInner) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *InstanceContainerServerInterfacesInner) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *InstanceContainerServerInterfacesInner) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### GetIpAddress

`func (o *InstanceContainerServerInterfacesInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *InstanceContainerServerInterfacesInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *InstanceContainerServerInterfacesInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *InstanceContainerServerInterfacesInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetDhcp

`func (o *InstanceContainerServerInterfacesInner) GetDhcp() bool`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *InstanceContainerServerInterfacesInner) GetDhcpOk() (*bool, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *InstanceContainerServerInterfacesInner) SetDhcp(v bool)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *InstanceContainerServerInterfacesInner) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### GetActive

`func (o *InstanceContainerServerInterfacesInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InstanceContainerServerInterfacesInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InstanceContainerServerInterfacesInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *InstanceContainerServerInterfacesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPoolAssigned

`func (o *InstanceContainerServerInterfacesInner) GetPoolAssigned() bool`

GetPoolAssigned returns the PoolAssigned field if non-nil, zero value otherwise.

### GetPoolAssignedOk

`func (o *InstanceContainerServerInterfacesInner) GetPoolAssignedOk() (*bool, bool)`

GetPoolAssignedOk returns a tuple with the PoolAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAssigned

`func (o *InstanceContainerServerInterfacesInner) SetPoolAssigned(v bool)`

SetPoolAssigned sets PoolAssigned field to given value.

### HasPoolAssigned

`func (o *InstanceContainerServerInterfacesInner) HasPoolAssigned() bool`

HasPoolAssigned returns a boolean if a field has been set.

### GetPrimaryInterface

`func (o *InstanceContainerServerInterfacesInner) GetPrimaryInterface() bool`

GetPrimaryInterface returns the PrimaryInterface field if non-nil, zero value otherwise.

### GetPrimaryInterfaceOk

`func (o *InstanceContainerServerInterfacesInner) GetPrimaryInterfaceOk() (*bool, bool)`

GetPrimaryInterfaceOk returns a tuple with the PrimaryInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInterface

`func (o *InstanceContainerServerInterfacesInner) SetPrimaryInterface(v bool)`

SetPrimaryInterface sets PrimaryInterface field to given value.

### HasPrimaryInterface

`func (o *InstanceContainerServerInterfacesInner) HasPrimaryInterface() bool`

HasPrimaryInterface returns a boolean if a field has been set.

### GetNetwork

`func (o *InstanceContainerServerInterfacesInner) GetNetwork() InstanceContainerServerInterfacesInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InstanceContainerServerInterfacesInner) GetNetworkOk() (*InstanceContainerServerInterfacesInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InstanceContainerServerInterfacesInner) SetNetwork(v InstanceContainerServerInterfacesInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InstanceContainerServerInterfacesInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkGroup

`func (o *InstanceContainerServerInterfacesInner) GetNetworkGroup() InstanceContainerServerInterfacesInnerNetworkGroup`

GetNetworkGroup returns the NetworkGroup field if non-nil, zero value otherwise.

### GetNetworkGroupOk

`func (o *InstanceContainerServerInterfacesInner) GetNetworkGroupOk() (*InstanceContainerServerInterfacesInnerNetworkGroup, bool)`

GetNetworkGroupOk returns a tuple with the NetworkGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkGroup

`func (o *InstanceContainerServerInterfacesInner) SetNetworkGroup(v InstanceContainerServerInterfacesInnerNetworkGroup)`

SetNetworkGroup sets NetworkGroup field to given value.

### HasNetworkGroup

`func (o *InstanceContainerServerInterfacesInner) HasNetworkGroup() bool`

HasNetworkGroup returns a boolean if a field has been set.

### GetNetworkPool

`func (o *InstanceContainerServerInterfacesInner) GetNetworkPool() InstanceContainerServerInterfacesInnerNetworkPool`

GetNetworkPool returns the NetworkPool field if non-nil, zero value otherwise.

### GetNetworkPoolOk

`func (o *InstanceContainerServerInterfacesInner) GetNetworkPoolOk() (*InstanceContainerServerInterfacesInnerNetworkPool, bool)`

GetNetworkPoolOk returns a tuple with the NetworkPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPool

`func (o *InstanceContainerServerInterfacesInner) SetNetworkPool(v InstanceContainerServerInterfacesInnerNetworkPool)`

SetNetworkPool sets NetworkPool field to given value.

### HasNetworkPool

`func (o *InstanceContainerServerInterfacesInner) HasNetworkPool() bool`

HasNetworkPool returns a boolean if a field has been set.

### GetIpMode

`func (o *InstanceContainerServerInterfacesInner) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *InstanceContainerServerInterfacesInner) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *InstanceContainerServerInterfacesInner) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *InstanceContainerServerInterfacesInner) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetMacAddress

`func (o *InstanceContainerServerInterfacesInner) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *InstanceContainerServerInterfacesInner) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *InstanceContainerServerInterfacesInner) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *InstanceContainerServerInterfacesInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetInterfaces

`func (o *InstanceContainerServerInterfacesInner) GetInterfaces() []InstanceContainerServerInstancesInnerInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *InstanceContainerServerInterfacesInner) GetInterfacesOk() (*[]InstanceContainerServerInstancesInnerInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *InstanceContainerServerInterfacesInner) SetInterfaces(v []InstanceContainerServerInstancesInnerInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *InstanceContainerServerInterfacesInner) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


