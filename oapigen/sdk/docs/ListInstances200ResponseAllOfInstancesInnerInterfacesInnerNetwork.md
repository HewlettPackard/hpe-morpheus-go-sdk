# ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** |  | [optional] 
**Group** | Pointer to **NullableInt64** |  | [optional] 
**Subnet** | Pointer to **NullableString** |  | [optional] 
**DhcpServer** | Pointer to **NullableBool** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Pool** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 

## Methods

### NewListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork

`func NewListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork() *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork`

NewListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork instantiates a new ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetworkWithDefaults

`func NewListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetworkWithDefaults() *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork`

NewListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetworkWithDefaults instantiates a new ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetGroup

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) GetGroup() int64`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) GetGroupOk() (*int64, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) SetGroup(v int64)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetSubnet

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetDhcpServer

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) GetDhcpServer() bool`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) GetDhcpServerOk() (*bool, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) SetDhcpServer(v bool)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### SetDhcpServerNil

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) SetDhcpServerNil(b bool)`

 SetDhcpServerNil sets the value for DhcpServer to be an explicit nil

### UnsetDhcpServer
`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) UnsetDhcpServer()`

UnsetDhcpServer ensures that no value is present for DhcpServer, not even an explicit nil
### GetName

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPool

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) GetPool() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) GetPoolOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) SetPool(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetPool sets Pool field to given value.

### HasPool

`func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


