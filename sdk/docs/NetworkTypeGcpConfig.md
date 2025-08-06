# NetworkTypeGcpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mtu** | **string** | GCP MTU | [default to "1460"]
**AutoCreate** | **bool** | Auto create subnets | [default to true]

## Methods

### NewNetworkTypeGcpConfig

`func NewNetworkTypeGcpConfig(mtu string, autoCreate bool, ) *NetworkTypeGcpConfig`

NewNetworkTypeGcpConfig instantiates a new NetworkTypeGcpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTypeGcpConfigWithDefaults

`func NewNetworkTypeGcpConfigWithDefaults() *NetworkTypeGcpConfig`

NewNetworkTypeGcpConfigWithDefaults instantiates a new NetworkTypeGcpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMtu

`func (o *NetworkTypeGcpConfig) GetMtu() string`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *NetworkTypeGcpConfig) GetMtuOk() (*string, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *NetworkTypeGcpConfig) SetMtu(v string)`

SetMtu sets Mtu field to given value.


### GetAutoCreate

`func (o *NetworkTypeGcpConfig) GetAutoCreate() bool`

GetAutoCreate returns the AutoCreate field if non-nil, zero value otherwise.

### GetAutoCreateOk

`func (o *NetworkTypeGcpConfig) GetAutoCreateOk() (*bool, bool)`

GetAutoCreateOk returns a tuple with the AutoCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreate

`func (o *NetworkTypeGcpConfig) SetAutoCreate(v bool)`

SetAutoCreate sets AutoCreate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


