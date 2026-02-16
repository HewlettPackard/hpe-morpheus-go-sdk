# InstancesNetworkInterfaces1NetworkInterfacesInnerNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the network to be used. A network group can be specified instead by prefixing its ID with &#x60;networkGroup-&#x60;. | 
**Pool** | Pointer to [**InstancesNetworkInterfaces1NetworkInterfacesInnerNetworkPool**](InstancesNetworkInterfaces1NetworkInterfacesInnerNetworkPool.md) |  | [optional] 

## Methods

### NewInstancesNetworkInterfaces1NetworkInterfacesInnerNetwork

`func NewInstancesNetworkInterfaces1NetworkInterfacesInnerNetwork(id string, ) *InstancesNetworkInterfaces1NetworkInterfacesInnerNetwork`

NewInstancesNetworkInterfaces1NetworkInterfacesInnerNetwork instantiates a new InstancesNetworkInterfaces1NetworkInterfacesInnerNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesNetworkInterfaces1NetworkInterfacesInnerNetworkWithDefaults

`func NewInstancesNetworkInterfaces1NetworkInterfacesInnerNetworkWithDefaults() *InstancesNetworkInterfaces1NetworkInterfacesInnerNetwork`

NewInstancesNetworkInterfaces1NetworkInterfacesInnerNetworkWithDefaults instantiates a new InstancesNetworkInterfaces1NetworkInterfacesInnerNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInnerNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInnerNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInnerNetwork) SetId(v string)`

SetId sets Id field to given value.


### GetPool

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInnerNetwork) GetPool() InstancesNetworkInterfaces1NetworkInterfacesInnerNetworkPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInnerNetwork) GetPoolOk() (*InstancesNetworkInterfaces1NetworkInterfacesInnerNetworkPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInnerNetwork) SetPool(v InstancesNetworkInterfaces1NetworkInterfacesInnerNetworkPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInnerNetwork) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


