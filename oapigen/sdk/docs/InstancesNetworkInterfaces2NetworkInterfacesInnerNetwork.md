# InstancesNetworkInterfaces2NetworkInterfacesInnerNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the network to be used. A network group can be specified instead by prefixing its ID with &#x60;networkGroup-&#x60;. | 
**Pool** | Pointer to [**InstancesNetworkInterfaces2NetworkInterfacesInnerNetworkPool**](InstancesNetworkInterfaces2NetworkInterfacesInnerNetworkPool.md) |  | [optional] 

## Methods

### NewInstancesNetworkInterfaces2NetworkInterfacesInnerNetwork

`func NewInstancesNetworkInterfaces2NetworkInterfacesInnerNetwork(id string, ) *InstancesNetworkInterfaces2NetworkInterfacesInnerNetwork`

NewInstancesNetworkInterfaces2NetworkInterfacesInnerNetwork instantiates a new InstancesNetworkInterfaces2NetworkInterfacesInnerNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesNetworkInterfaces2NetworkInterfacesInnerNetworkWithDefaults

`func NewInstancesNetworkInterfaces2NetworkInterfacesInnerNetworkWithDefaults() *InstancesNetworkInterfaces2NetworkInterfacesInnerNetwork`

NewInstancesNetworkInterfaces2NetworkInterfacesInnerNetworkWithDefaults instantiates a new InstancesNetworkInterfaces2NetworkInterfacesInnerNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstancesNetworkInterfaces2NetworkInterfacesInnerNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstancesNetworkInterfaces2NetworkInterfacesInnerNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstancesNetworkInterfaces2NetworkInterfacesInnerNetwork) SetId(v string)`

SetId sets Id field to given value.


### GetPool

`func (o *InstancesNetworkInterfaces2NetworkInterfacesInnerNetwork) GetPool() InstancesNetworkInterfaces2NetworkInterfacesInnerNetworkPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *InstancesNetworkInterfaces2NetworkInterfacesInnerNetwork) GetPoolOk() (*InstancesNetworkInterfaces2NetworkInterfacesInnerNetworkPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *InstancesNetworkInterfaces2NetworkInterfacesInnerNetwork) SetPool(v InstancesNetworkInterfaces2NetworkInterfacesInnerNetworkPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *InstancesNetworkInterfaces2NetworkInterfacesInnerNetwork) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


