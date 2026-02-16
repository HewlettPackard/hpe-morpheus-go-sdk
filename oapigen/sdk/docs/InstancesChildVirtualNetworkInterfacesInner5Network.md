# InstancesChildVirtualNetworkInterfacesInner5Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the network to be used. A network group can be specified instead by prefixing its ID with &#x60;networkGroup-&#x60;. | 
**Pool** | Pointer to [**InstancesChildVirtualNetworkInterfacesInner5NetworkPool**](InstancesChildVirtualNetworkInterfacesInner5NetworkPool.md) |  | [optional] 

## Methods

### NewInstancesChildVirtualNetworkInterfacesInner5Network

`func NewInstancesChildVirtualNetworkInterfacesInner5Network(id string, ) *InstancesChildVirtualNetworkInterfacesInner5Network`

NewInstancesChildVirtualNetworkInterfacesInner5Network instantiates a new InstancesChildVirtualNetworkInterfacesInner5Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesChildVirtualNetworkInterfacesInner5NetworkWithDefaults

`func NewInstancesChildVirtualNetworkInterfacesInner5NetworkWithDefaults() *InstancesChildVirtualNetworkInterfacesInner5Network`

NewInstancesChildVirtualNetworkInterfacesInner5NetworkWithDefaults instantiates a new InstancesChildVirtualNetworkInterfacesInner5Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstancesChildVirtualNetworkInterfacesInner5Network) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstancesChildVirtualNetworkInterfacesInner5Network) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstancesChildVirtualNetworkInterfacesInner5Network) SetId(v string)`

SetId sets Id field to given value.


### GetPool

`func (o *InstancesChildVirtualNetworkInterfacesInner5Network) GetPool() InstancesChildVirtualNetworkInterfacesInner5NetworkPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *InstancesChildVirtualNetworkInterfacesInner5Network) GetPoolOk() (*InstancesChildVirtualNetworkInterfacesInner5NetworkPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *InstancesChildVirtualNetworkInterfacesInner5Network) SetPool(v InstancesChildVirtualNetworkInterfacesInner5NetworkPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *InstancesChildVirtualNetworkInterfacesInner5Network) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


