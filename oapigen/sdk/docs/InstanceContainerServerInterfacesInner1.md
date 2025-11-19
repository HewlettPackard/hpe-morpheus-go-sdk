# InstanceContainerServerInterfacesInner1

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
**Network** | Pointer to [**AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool**](AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool.md) |  | [optional] 
**NetworkGroup** | Pointer to [**AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool**](AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool.md) |  | [optional] 
**NetworkPool** | Pointer to **int64** |  | [optional] 
**IpMode** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**Interfaces** | Pointer to [**[]AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool**](AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool.md) |  | [optional] 

## Methods

### NewInstanceContainerServerInterfacesInner1

`func NewInstanceContainerServerInterfacesInner1() *InstanceContainerServerInterfacesInner1`

NewInstanceContainerServerInterfacesInner1 instantiates a new InstanceContainerServerInterfacesInner1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceContainerServerInterfacesInner1WithDefaults

`func NewInstanceContainerServerInterfacesInner1WithDefaults() *InstanceContainerServerInterfacesInner1`

NewInstanceContainerServerInterfacesInner1WithDefaults instantiates a new InstanceContainerServerInterfacesInner1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceContainerServerInterfacesInner1) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceContainerServerInterfacesInner1) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceContainerServerInterfacesInner1) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceContainerServerInterfacesInner1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InstanceContainerServerInterfacesInner1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceContainerServerInterfacesInner1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceContainerServerInterfacesInner1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceContainerServerInterfacesInner1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUniqueId

`func (o *InstanceContainerServerInterfacesInner1) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *InstanceContainerServerInterfacesInner1) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *InstanceContainerServerInterfacesInner1) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *InstanceContainerServerInterfacesInner1) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetPublicIpAddress

`func (o *InstanceContainerServerInterfacesInner1) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *InstanceContainerServerInterfacesInner1) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *InstanceContainerServerInterfacesInner1) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *InstanceContainerServerInterfacesInner1) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### GetIpAddress

`func (o *InstanceContainerServerInterfacesInner1) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *InstanceContainerServerInterfacesInner1) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *InstanceContainerServerInterfacesInner1) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *InstanceContainerServerInterfacesInner1) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetDhcp

`func (o *InstanceContainerServerInterfacesInner1) GetDhcp() bool`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *InstanceContainerServerInterfacesInner1) GetDhcpOk() (*bool, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *InstanceContainerServerInterfacesInner1) SetDhcp(v bool)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *InstanceContainerServerInterfacesInner1) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### GetActive

`func (o *InstanceContainerServerInterfacesInner1) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InstanceContainerServerInterfacesInner1) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InstanceContainerServerInterfacesInner1) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *InstanceContainerServerInterfacesInner1) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPoolAssigned

`func (o *InstanceContainerServerInterfacesInner1) GetPoolAssigned() bool`

GetPoolAssigned returns the PoolAssigned field if non-nil, zero value otherwise.

### GetPoolAssignedOk

`func (o *InstanceContainerServerInterfacesInner1) GetPoolAssignedOk() (*bool, bool)`

GetPoolAssignedOk returns a tuple with the PoolAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAssigned

`func (o *InstanceContainerServerInterfacesInner1) SetPoolAssigned(v bool)`

SetPoolAssigned sets PoolAssigned field to given value.

### HasPoolAssigned

`func (o *InstanceContainerServerInterfacesInner1) HasPoolAssigned() bool`

HasPoolAssigned returns a boolean if a field has been set.

### GetPrimaryInterface

`func (o *InstanceContainerServerInterfacesInner1) GetPrimaryInterface() bool`

GetPrimaryInterface returns the PrimaryInterface field if non-nil, zero value otherwise.

### GetPrimaryInterfaceOk

`func (o *InstanceContainerServerInterfacesInner1) GetPrimaryInterfaceOk() (*bool, bool)`

GetPrimaryInterfaceOk returns a tuple with the PrimaryInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInterface

`func (o *InstanceContainerServerInterfacesInner1) SetPrimaryInterface(v bool)`

SetPrimaryInterface sets PrimaryInterface field to given value.

### HasPrimaryInterface

`func (o *InstanceContainerServerInterfacesInner1) HasPrimaryInterface() bool`

HasPrimaryInterface returns a boolean if a field has been set.

### GetNetwork

`func (o *InstanceContainerServerInterfacesInner1) GetNetwork() AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InstanceContainerServerInterfacesInner1) GetNetworkOk() (*AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InstanceContainerServerInterfacesInner1) SetNetwork(v AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InstanceContainerServerInterfacesInner1) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkGroup

`func (o *InstanceContainerServerInterfacesInner1) GetNetworkGroup() AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool`

GetNetworkGroup returns the NetworkGroup field if non-nil, zero value otherwise.

### GetNetworkGroupOk

`func (o *InstanceContainerServerInterfacesInner1) GetNetworkGroupOk() (*AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool, bool)`

GetNetworkGroupOk returns a tuple with the NetworkGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkGroup

`func (o *InstanceContainerServerInterfacesInner1) SetNetworkGroup(v AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool)`

SetNetworkGroup sets NetworkGroup field to given value.

### HasNetworkGroup

`func (o *InstanceContainerServerInterfacesInner1) HasNetworkGroup() bool`

HasNetworkGroup returns a boolean if a field has been set.

### GetNetworkPool

`func (o *InstanceContainerServerInterfacesInner1) GetNetworkPool() int64`

GetNetworkPool returns the NetworkPool field if non-nil, zero value otherwise.

### GetNetworkPoolOk

`func (o *InstanceContainerServerInterfacesInner1) GetNetworkPoolOk() (*int64, bool)`

GetNetworkPoolOk returns a tuple with the NetworkPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPool

`func (o *InstanceContainerServerInterfacesInner1) SetNetworkPool(v int64)`

SetNetworkPool sets NetworkPool field to given value.

### HasNetworkPool

`func (o *InstanceContainerServerInterfacesInner1) HasNetworkPool() bool`

HasNetworkPool returns a boolean if a field has been set.

### GetIpMode

`func (o *InstanceContainerServerInterfacesInner1) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *InstanceContainerServerInterfacesInner1) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *InstanceContainerServerInterfacesInner1) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *InstanceContainerServerInterfacesInner1) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetMacAddress

`func (o *InstanceContainerServerInterfacesInner1) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *InstanceContainerServerInterfacesInner1) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *InstanceContainerServerInterfacesInner1) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *InstanceContainerServerInterfacesInner1) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetInterfaces

`func (o *InstanceContainerServerInterfacesInner1) GetInterfaces() []AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *InstanceContainerServerInterfacesInner1) GetInterfacesOk() (*[]AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *InstanceContainerServerInterfacesInner1) SetInterfaces(v []AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *InstanceContainerServerInterfacesInner1) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


