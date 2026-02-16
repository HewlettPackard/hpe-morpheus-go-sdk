# InstancesNetworkInterfaces4Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the network to be used. A network group can be specified instead by prefixing its ID with &#x60;networkGroup-&#x60;. | 
**Pool** | Pointer to [**InstancesNetworkInterfaces4NetworkPool**](InstancesNetworkInterfaces4NetworkPool.md) |  | [optional] 

## Methods

### NewInstancesNetworkInterfaces4Network

`func NewInstancesNetworkInterfaces4Network(id string, ) *InstancesNetworkInterfaces4Network`

NewInstancesNetworkInterfaces4Network instantiates a new InstancesNetworkInterfaces4Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesNetworkInterfaces4NetworkWithDefaults

`func NewInstancesNetworkInterfaces4NetworkWithDefaults() *InstancesNetworkInterfaces4Network`

NewInstancesNetworkInterfaces4NetworkWithDefaults instantiates a new InstancesNetworkInterfaces4Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstancesNetworkInterfaces4Network) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstancesNetworkInterfaces4Network) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstancesNetworkInterfaces4Network) SetId(v string)`

SetId sets Id field to given value.


### GetPool

`func (o *InstancesNetworkInterfaces4Network) GetPool() InstancesNetworkInterfaces4NetworkPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *InstancesNetworkInterfaces4Network) GetPoolOk() (*InstancesNetworkInterfaces4NetworkPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *InstancesNetworkInterfaces4Network) SetPool(v InstancesNetworkInterfaces4NetworkPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *InstancesNetworkInterfaces4Network) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


