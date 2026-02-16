# InstancesNetworkInterfaces1Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the network to be used. A network group can be specified instead by prefixing its ID with &#x60;networkGroup-&#x60;. | 
**Pool** | Pointer to [**InstancesNetworkInterfaces1NetworkPool**](InstancesNetworkInterfaces1NetworkPool.md) |  | [optional] 

## Methods

### NewInstancesNetworkInterfaces1Network

`func NewInstancesNetworkInterfaces1Network(id string, ) *InstancesNetworkInterfaces1Network`

NewInstancesNetworkInterfaces1Network instantiates a new InstancesNetworkInterfaces1Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesNetworkInterfaces1NetworkWithDefaults

`func NewInstancesNetworkInterfaces1NetworkWithDefaults() *InstancesNetworkInterfaces1Network`

NewInstancesNetworkInterfaces1NetworkWithDefaults instantiates a new InstancesNetworkInterfaces1Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstancesNetworkInterfaces1Network) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstancesNetworkInterfaces1Network) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstancesNetworkInterfaces1Network) SetId(v string)`

SetId sets Id field to given value.


### GetPool

`func (o *InstancesNetworkInterfaces1Network) GetPool() InstancesNetworkInterfaces1NetworkPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *InstancesNetworkInterfaces1Network) GetPoolOk() (*InstancesNetworkInterfaces1NetworkPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *InstancesNetworkInterfaces1Network) SetPool(v InstancesNetworkInterfaces1NetworkPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *InstancesNetworkInterfaces1Network) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


