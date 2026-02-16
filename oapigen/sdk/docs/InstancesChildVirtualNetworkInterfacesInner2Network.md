# InstancesChildVirtualNetworkInterfacesInner2Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the network to be used. A network group can be specified instead by prefixing its ID with &#x60;networkGroup-&#x60;. | 
**Pool** | Pointer to [**InstancesChildVirtualNetworkInterfacesInner2NetworkPool**](InstancesChildVirtualNetworkInterfacesInner2NetworkPool.md) |  | [optional] 

## Methods

### NewInstancesChildVirtualNetworkInterfacesInner2Network

`func NewInstancesChildVirtualNetworkInterfacesInner2Network(id string, ) *InstancesChildVirtualNetworkInterfacesInner2Network`

NewInstancesChildVirtualNetworkInterfacesInner2Network instantiates a new InstancesChildVirtualNetworkInterfacesInner2Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesChildVirtualNetworkInterfacesInner2NetworkWithDefaults

`func NewInstancesChildVirtualNetworkInterfacesInner2NetworkWithDefaults() *InstancesChildVirtualNetworkInterfacesInner2Network`

NewInstancesChildVirtualNetworkInterfacesInner2NetworkWithDefaults instantiates a new InstancesChildVirtualNetworkInterfacesInner2Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstancesChildVirtualNetworkInterfacesInner2Network) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstancesChildVirtualNetworkInterfacesInner2Network) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstancesChildVirtualNetworkInterfacesInner2Network) SetId(v string)`

SetId sets Id field to given value.


### GetPool

`func (o *InstancesChildVirtualNetworkInterfacesInner2Network) GetPool() InstancesChildVirtualNetworkInterfacesInner2NetworkPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *InstancesChildVirtualNetworkInterfacesInner2Network) GetPoolOk() (*InstancesChildVirtualNetworkInterfacesInner2NetworkPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *InstancesChildVirtualNetworkInterfacesInner2Network) SetPool(v InstancesChildVirtualNetworkInterfacesInner2NetworkPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *InstancesChildVirtualNetworkInterfacesInner2Network) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


