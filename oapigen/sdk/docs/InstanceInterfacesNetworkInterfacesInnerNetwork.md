# InstanceInterfacesNetworkInterfacesInnerNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** |  | [optional] 
**Group** | Pointer to **NullableInt64** |  | [optional] 
**Subnet** | Pointer to **NullableString** |  | [optional] 
**DhcpServer** | Pointer to **NullableBool** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Pool** | Pointer to [**InstanceInterfacesNetworkInterfacesInnerNetworkPool**](InstanceInterfacesNetworkInterfacesInnerNetworkPool.md) |  | [optional] 

## Methods

### NewInstanceInterfacesNetworkInterfacesInnerNetwork

`func NewInstanceInterfacesNetworkInterfacesInnerNetwork() *InstanceInterfacesNetworkInterfacesInnerNetwork`

NewInstanceInterfacesNetworkInterfacesInnerNetwork instantiates a new InstanceInterfacesNetworkInterfacesInnerNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceInterfacesNetworkInterfacesInnerNetworkWithDefaults

`func NewInstanceInterfacesNetworkInterfacesInnerNetworkWithDefaults() *InstanceInterfacesNetworkInterfacesInnerNetwork`

NewInstanceInterfacesNetworkInterfacesInnerNetworkWithDefaults instantiates a new InstanceInterfacesNetworkInterfacesInnerNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetGroup

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) GetGroup() int64`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) GetGroupOk() (*int64, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) SetGroup(v int64)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetSubnet

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetDhcpServer

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) GetDhcpServer() bool`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) GetDhcpServerOk() (*bool, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) SetDhcpServer(v bool)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### SetDhcpServerNil

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) SetDhcpServerNil(b bool)`

 SetDhcpServerNil sets the value for DhcpServer to be an explicit nil

### UnsetDhcpServer
`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) UnsetDhcpServer()`

UnsetDhcpServer ensures that no value is present for DhcpServer, not even an explicit nil
### GetName

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPool

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) GetPool() InstanceInterfacesNetworkInterfacesInnerNetworkPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) GetPoolOk() (*InstanceInterfacesNetworkInterfacesInnerNetworkPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) SetPool(v InstanceInterfacesNetworkInterfacesInnerNetworkPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *InstanceInterfacesNetworkInterfacesInnerNetwork) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


