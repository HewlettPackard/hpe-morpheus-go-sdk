# InstancesChildVirtualNetworkInterfacesInner6Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the network to be used. A network group can be specified instead by prefixing its ID with &#x60;networkGroup-&#x60;. | 
**Pool** | Pointer to [**InstancesChildVirtualNetworkInterfacesInner6NetworkPool**](InstancesChildVirtualNetworkInterfacesInner6NetworkPool.md) |  | [optional] 

## Methods

### NewInstancesChildVirtualNetworkInterfacesInner6Network

`func NewInstancesChildVirtualNetworkInterfacesInner6Network(id string, ) *InstancesChildVirtualNetworkInterfacesInner6Network`

NewInstancesChildVirtualNetworkInterfacesInner6Network instantiates a new InstancesChildVirtualNetworkInterfacesInner6Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesChildVirtualNetworkInterfacesInner6NetworkWithDefaults

`func NewInstancesChildVirtualNetworkInterfacesInner6NetworkWithDefaults() *InstancesChildVirtualNetworkInterfacesInner6Network`

NewInstancesChildVirtualNetworkInterfacesInner6NetworkWithDefaults instantiates a new InstancesChildVirtualNetworkInterfacesInner6Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstancesChildVirtualNetworkInterfacesInner6Network) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstancesChildVirtualNetworkInterfacesInner6Network) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstancesChildVirtualNetworkInterfacesInner6Network) SetId(v string)`

SetId sets Id field to given value.


### GetPool

`func (o *InstancesChildVirtualNetworkInterfacesInner6Network) GetPool() InstancesChildVirtualNetworkInterfacesInner6NetworkPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *InstancesChildVirtualNetworkInterfacesInner6Network) GetPoolOk() (*InstancesChildVirtualNetworkInterfacesInner6NetworkPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *InstancesChildVirtualNetworkInterfacesInner6Network) SetPool(v InstancesChildVirtualNetworkInterfacesInner6NetworkPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *InstancesChildVirtualNetworkInterfacesInner6Network) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


