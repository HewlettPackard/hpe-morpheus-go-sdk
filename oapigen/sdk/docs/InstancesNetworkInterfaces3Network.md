# InstancesNetworkInterfaces3Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the network to be used. A network group can be specified instead by prefixing its ID with &#x60;networkGroup-&#x60;. | 
**Pool** | Pointer to [**InstancesNetworkInterfaces3NetworkPool**](InstancesNetworkInterfaces3NetworkPool.md) |  | [optional] 

## Methods

### NewInstancesNetworkInterfaces3Network

`func NewInstancesNetworkInterfaces3Network(id string, ) *InstancesNetworkInterfaces3Network`

NewInstancesNetworkInterfaces3Network instantiates a new InstancesNetworkInterfaces3Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesNetworkInterfaces3NetworkWithDefaults

`func NewInstancesNetworkInterfaces3NetworkWithDefaults() *InstancesNetworkInterfaces3Network`

NewInstancesNetworkInterfaces3NetworkWithDefaults instantiates a new InstancesNetworkInterfaces3Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstancesNetworkInterfaces3Network) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstancesNetworkInterfaces3Network) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstancesNetworkInterfaces3Network) SetId(v string)`

SetId sets Id field to given value.


### GetPool

`func (o *InstancesNetworkInterfaces3Network) GetPool() InstancesNetworkInterfaces3NetworkPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *InstancesNetworkInterfaces3Network) GetPoolOk() (*InstancesNetworkInterfaces3NetworkPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *InstancesNetworkInterfaces3Network) SetPool(v InstancesNetworkInterfaces3NetworkPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *InstancesNetworkInterfaces3Network) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


