# InstancesChildVirtualNetworkInterfacesInner1Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the network to be used. A network group can be specified instead by prefixing its ID with &#x60;networkGroup-&#x60;. | 
**Pool** | Pointer to [**InstancesChildVirtualNetworkInterfacesInner1NetworkPool**](InstancesChildVirtualNetworkInterfacesInner1NetworkPool.md) |  | [optional] 

## Methods

### NewInstancesChildVirtualNetworkInterfacesInner1Network

`func NewInstancesChildVirtualNetworkInterfacesInner1Network(id string, ) *InstancesChildVirtualNetworkInterfacesInner1Network`

NewInstancesChildVirtualNetworkInterfacesInner1Network instantiates a new InstancesChildVirtualNetworkInterfacesInner1Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesChildVirtualNetworkInterfacesInner1NetworkWithDefaults

`func NewInstancesChildVirtualNetworkInterfacesInner1NetworkWithDefaults() *InstancesChildVirtualNetworkInterfacesInner1Network`

NewInstancesChildVirtualNetworkInterfacesInner1NetworkWithDefaults instantiates a new InstancesChildVirtualNetworkInterfacesInner1Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstancesChildVirtualNetworkInterfacesInner1Network) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstancesChildVirtualNetworkInterfacesInner1Network) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstancesChildVirtualNetworkInterfacesInner1Network) SetId(v string)`

SetId sets Id field to given value.


### GetPool

`func (o *InstancesChildVirtualNetworkInterfacesInner1Network) GetPool() InstancesChildVirtualNetworkInterfacesInner1NetworkPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *InstancesChildVirtualNetworkInterfacesInner1Network) GetPoolOk() (*InstancesChildVirtualNetworkInterfacesInner1NetworkPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *InstancesChildVirtualNetworkInterfacesInner1Network) SetPool(v InstancesChildVirtualNetworkInterfacesInner1NetworkPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *InstancesChildVirtualNetworkInterfacesInner1Network) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


