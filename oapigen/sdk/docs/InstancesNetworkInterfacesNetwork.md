# InstancesNetworkInterfacesNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the network to be used. A network group can be specified instead by prefixing its ID with &#x60;networkGroup-&#x60;. | 
**Pool** | Pointer to [**InstancesNetworkInterfacesNetworkPool**](InstancesNetworkInterfacesNetworkPool.md) |  | [optional] 

## Methods

### NewInstancesNetworkInterfacesNetwork

`func NewInstancesNetworkInterfacesNetwork(id string, ) *InstancesNetworkInterfacesNetwork`

NewInstancesNetworkInterfacesNetwork instantiates a new InstancesNetworkInterfacesNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesNetworkInterfacesNetworkWithDefaults

`func NewInstancesNetworkInterfacesNetworkWithDefaults() *InstancesNetworkInterfacesNetwork`

NewInstancesNetworkInterfacesNetworkWithDefaults instantiates a new InstancesNetworkInterfacesNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstancesNetworkInterfacesNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstancesNetworkInterfacesNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstancesNetworkInterfacesNetwork) SetId(v string)`

SetId sets Id field to given value.


### GetPool

`func (o *InstancesNetworkInterfacesNetwork) GetPool() InstancesNetworkInterfacesNetworkPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *InstancesNetworkInterfacesNetwork) GetPoolOk() (*InstancesNetworkInterfacesNetworkPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *InstancesNetworkInterfacesNetwork) SetPool(v InstancesNetworkInterfacesNetworkPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *InstancesNetworkInterfacesNetwork) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


