# InstancesNetworkInterfaces4NetworkInterfacesInnerNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the network to be used. A network group can be specified instead by prefixing its ID with &#x60;networkGroup-&#x60;. | 
**Pool** | Pointer to [**InstancesNetworkInterfaces4NetworkInterfacesInnerNetworkPool**](InstancesNetworkInterfaces4NetworkInterfacesInnerNetworkPool.md) |  | [optional] 

## Methods

### NewInstancesNetworkInterfaces4NetworkInterfacesInnerNetwork

`func NewInstancesNetworkInterfaces4NetworkInterfacesInnerNetwork(id string, ) *InstancesNetworkInterfaces4NetworkInterfacesInnerNetwork`

NewInstancesNetworkInterfaces4NetworkInterfacesInnerNetwork instantiates a new InstancesNetworkInterfaces4NetworkInterfacesInnerNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesNetworkInterfaces4NetworkInterfacesInnerNetworkWithDefaults

`func NewInstancesNetworkInterfaces4NetworkInterfacesInnerNetworkWithDefaults() *InstancesNetworkInterfaces4NetworkInterfacesInnerNetwork`

NewInstancesNetworkInterfaces4NetworkInterfacesInnerNetworkWithDefaults instantiates a new InstancesNetworkInterfaces4NetworkInterfacesInnerNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstancesNetworkInterfaces4NetworkInterfacesInnerNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstancesNetworkInterfaces4NetworkInterfacesInnerNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstancesNetworkInterfaces4NetworkInterfacesInnerNetwork) SetId(v string)`

SetId sets Id field to given value.


### GetPool

`func (o *InstancesNetworkInterfaces4NetworkInterfacesInnerNetwork) GetPool() InstancesNetworkInterfaces4NetworkInterfacesInnerNetworkPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *InstancesNetworkInterfaces4NetworkInterfacesInnerNetwork) GetPoolOk() (*InstancesNetworkInterfaces4NetworkInterfacesInnerNetworkPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *InstancesNetworkInterfaces4NetworkInterfacesInnerNetwork) SetPool(v InstancesNetworkInterfaces4NetworkInterfacesInnerNetworkPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *InstancesNetworkInterfaces4NetworkInterfacesInnerNetwork) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


