# InstancesNetworkInterfaces6NetworkInterfacesInnerNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the network to be used. A network group can be specified instead by prefixing its ID with &#x60;networkGroup-&#x60;. | 
**Pool** | Pointer to [**InstancesNetworkInterfaces6NetworkInterfacesInnerNetworkPool**](InstancesNetworkInterfaces6NetworkInterfacesInnerNetworkPool.md) |  | [optional] 

## Methods

### NewInstancesNetworkInterfaces6NetworkInterfacesInnerNetwork

`func NewInstancesNetworkInterfaces6NetworkInterfacesInnerNetwork(id string, ) *InstancesNetworkInterfaces6NetworkInterfacesInnerNetwork`

NewInstancesNetworkInterfaces6NetworkInterfacesInnerNetwork instantiates a new InstancesNetworkInterfaces6NetworkInterfacesInnerNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesNetworkInterfaces6NetworkInterfacesInnerNetworkWithDefaults

`func NewInstancesNetworkInterfaces6NetworkInterfacesInnerNetworkWithDefaults() *InstancesNetworkInterfaces6NetworkInterfacesInnerNetwork`

NewInstancesNetworkInterfaces6NetworkInterfacesInnerNetworkWithDefaults instantiates a new InstancesNetworkInterfaces6NetworkInterfacesInnerNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInnerNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInnerNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInnerNetwork) SetId(v string)`

SetId sets Id field to given value.


### GetPool

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInnerNetwork) GetPool() InstancesNetworkInterfaces6NetworkInterfacesInnerNetworkPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInnerNetwork) GetPoolOk() (*InstancesNetworkInterfaces6NetworkInterfacesInnerNetworkPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInnerNetwork) SetPool(v InstancesNetworkInterfaces6NetworkInterfacesInnerNetworkPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInnerNetwork) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


