# CreateNetworksRequestNetworkConfigAnyOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mtu** | **string** | GCP MTU | [default to "1460"]
**ZonePool** | [**CreateNetworksRequestNetworkConfigAnyOf2ZonePool**](CreateNetworksRequestNetworkConfigAnyOf2ZonePool.md) |  | 
**AutoCreate** | **bool** | Auto create subnets | [default to true]

## Methods

### NewCreateNetworksRequestNetworkConfigAnyOf2

`func NewCreateNetworksRequestNetworkConfigAnyOf2(mtu string, zonePool CreateNetworksRequestNetworkConfigAnyOf2ZonePool, autoCreate bool, ) *CreateNetworksRequestNetworkConfigAnyOf2`

NewCreateNetworksRequestNetworkConfigAnyOf2 instantiates a new CreateNetworksRequestNetworkConfigAnyOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworksRequestNetworkConfigAnyOf2WithDefaults

`func NewCreateNetworksRequestNetworkConfigAnyOf2WithDefaults() *CreateNetworksRequestNetworkConfigAnyOf2`

NewCreateNetworksRequestNetworkConfigAnyOf2WithDefaults instantiates a new CreateNetworksRequestNetworkConfigAnyOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMtu

`func (o *CreateNetworksRequestNetworkConfigAnyOf2) GetMtu() string`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *CreateNetworksRequestNetworkConfigAnyOf2) GetMtuOk() (*string, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *CreateNetworksRequestNetworkConfigAnyOf2) SetMtu(v string)`

SetMtu sets Mtu field to given value.


### GetZonePool

`func (o *CreateNetworksRequestNetworkConfigAnyOf2) GetZonePool() CreateNetworksRequestNetworkConfigAnyOf2ZonePool`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *CreateNetworksRequestNetworkConfigAnyOf2) GetZonePoolOk() (*CreateNetworksRequestNetworkConfigAnyOf2ZonePool, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *CreateNetworksRequestNetworkConfigAnyOf2) SetZonePool(v CreateNetworksRequestNetworkConfigAnyOf2ZonePool)`

SetZonePool sets ZonePool field to given value.


### GetAutoCreate

`func (o *CreateNetworksRequestNetworkConfigAnyOf2) GetAutoCreate() bool`

GetAutoCreate returns the AutoCreate field if non-nil, zero value otherwise.

### GetAutoCreateOk

`func (o *CreateNetworksRequestNetworkConfigAnyOf2) GetAutoCreateOk() (*bool, bool)`

GetAutoCreateOk returns a tuple with the AutoCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreate

`func (o *CreateNetworksRequestNetworkConfigAnyOf2) SetAutoCreate(v bool)`

SetAutoCreate sets AutoCreate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


