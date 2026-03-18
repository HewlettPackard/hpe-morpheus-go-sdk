# AllocateNetworkFloatingIpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkServerId** | Pointer to **int64** | Id of the network server | [optional] 
**FloatingIpPoolId** | Pointer to **int64** | Id of the network floating ip pool | [optional] 

## Methods

### NewAllocateNetworkFloatingIpRequest

`func NewAllocateNetworkFloatingIpRequest() *AllocateNetworkFloatingIpRequest`

NewAllocateNetworkFloatingIpRequest instantiates a new AllocateNetworkFloatingIpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocateNetworkFloatingIpRequestWithDefaults

`func NewAllocateNetworkFloatingIpRequestWithDefaults() *AllocateNetworkFloatingIpRequest`

NewAllocateNetworkFloatingIpRequestWithDefaults instantiates a new AllocateNetworkFloatingIpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkServerId

`func (o *AllocateNetworkFloatingIpRequest) GetNetworkServerId() int64`

GetNetworkServerId returns the NetworkServerId field if non-nil, zero value otherwise.

### GetNetworkServerIdOk

`func (o *AllocateNetworkFloatingIpRequest) GetNetworkServerIdOk() (*int64, bool)`

GetNetworkServerIdOk returns a tuple with the NetworkServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServerId

`func (o *AllocateNetworkFloatingIpRequest) SetNetworkServerId(v int64)`

SetNetworkServerId sets NetworkServerId field to given value.

### HasNetworkServerId

`func (o *AllocateNetworkFloatingIpRequest) HasNetworkServerId() bool`

HasNetworkServerId returns a boolean if a field has been set.

### GetFloatingIpPoolId

`func (o *AllocateNetworkFloatingIpRequest) GetFloatingIpPoolId() int64`

GetFloatingIpPoolId returns the FloatingIpPoolId field if non-nil, zero value otherwise.

### GetFloatingIpPoolIdOk

`func (o *AllocateNetworkFloatingIpRequest) GetFloatingIpPoolIdOk() (*int64, bool)`

GetFloatingIpPoolIdOk returns a tuple with the FloatingIpPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIpPoolId

`func (o *AllocateNetworkFloatingIpRequest) SetFloatingIpPoolId(v int64)`

SetFloatingIpPoolId sets FloatingIpPoolId field to given value.

### HasFloatingIpPoolId

`func (o *AllocateNetworkFloatingIpRequest) HasFloatingIpPoolId() bool`

HasFloatingIpPoolId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


