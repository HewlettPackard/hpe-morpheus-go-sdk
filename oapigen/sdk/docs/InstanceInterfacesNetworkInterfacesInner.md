# InstanceInterfacesNetworkInterfacesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**InstanceInterfacesNetworkInterfacesInnerId**](InstanceInterfacesNetworkInterfacesInnerId.md) |  | [optional] 
**Network** | Pointer to [**InstanceInterfacesNetworkInterfacesInnerNetwork**](InstanceInterfacesNetworkInterfacesInnerNetwork.md) |  | [optional] 
**IpAddress** | Pointer to **NullableString** |  | [optional] 
**NetworkInterfaceTypeId** | Pointer to **NullableInt64** |  | [optional] 
**IpMode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInstanceInterfacesNetworkInterfacesInner

`func NewInstanceInterfacesNetworkInterfacesInner() *InstanceInterfacesNetworkInterfacesInner`

NewInstanceInterfacesNetworkInterfacesInner instantiates a new InstanceInterfacesNetworkInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceInterfacesNetworkInterfacesInnerWithDefaults

`func NewInstanceInterfacesNetworkInterfacesInnerWithDefaults() *InstanceInterfacesNetworkInterfacesInner`

NewInstanceInterfacesNetworkInterfacesInnerWithDefaults instantiates a new InstanceInterfacesNetworkInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceInterfacesNetworkInterfacesInner) GetId() InstanceInterfacesNetworkInterfacesInnerId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceInterfacesNetworkInterfacesInner) GetIdOk() (*InstanceInterfacesNetworkInterfacesInnerId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceInterfacesNetworkInterfacesInner) SetId(v InstanceInterfacesNetworkInterfacesInnerId)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceInterfacesNetworkInterfacesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetwork

`func (o *InstanceInterfacesNetworkInterfacesInner) GetNetwork() InstanceInterfacesNetworkInterfacesInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InstanceInterfacesNetworkInterfacesInner) GetNetworkOk() (*InstanceInterfacesNetworkInterfacesInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InstanceInterfacesNetworkInterfacesInner) SetNetwork(v InstanceInterfacesNetworkInterfacesInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InstanceInterfacesNetworkInterfacesInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetIpAddress

`func (o *InstanceInterfacesNetworkInterfacesInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *InstanceInterfacesNetworkInterfacesInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *InstanceInterfacesNetworkInterfacesInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *InstanceInterfacesNetworkInterfacesInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *InstanceInterfacesNetworkInterfacesInner) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *InstanceInterfacesNetworkInterfacesInner) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetNetworkInterfaceTypeId

`func (o *InstanceInterfacesNetworkInterfacesInner) GetNetworkInterfaceTypeId() int64`

GetNetworkInterfaceTypeId returns the NetworkInterfaceTypeId field if non-nil, zero value otherwise.

### GetNetworkInterfaceTypeIdOk

`func (o *InstanceInterfacesNetworkInterfacesInner) GetNetworkInterfaceTypeIdOk() (*int64, bool)`

GetNetworkInterfaceTypeIdOk returns a tuple with the NetworkInterfaceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceTypeId

`func (o *InstanceInterfacesNetworkInterfacesInner) SetNetworkInterfaceTypeId(v int64)`

SetNetworkInterfaceTypeId sets NetworkInterfaceTypeId field to given value.

### HasNetworkInterfaceTypeId

`func (o *InstanceInterfacesNetworkInterfacesInner) HasNetworkInterfaceTypeId() bool`

HasNetworkInterfaceTypeId returns a boolean if a field has been set.

### SetNetworkInterfaceTypeIdNil

`func (o *InstanceInterfacesNetworkInterfacesInner) SetNetworkInterfaceTypeIdNil(b bool)`

 SetNetworkInterfaceTypeIdNil sets the value for NetworkInterfaceTypeId to be an explicit nil

### UnsetNetworkInterfaceTypeId
`func (o *InstanceInterfacesNetworkInterfacesInner) UnsetNetworkInterfaceTypeId()`

UnsetNetworkInterfaceTypeId ensures that no value is present for NetworkInterfaceTypeId, not even an explicit nil
### GetIpMode

`func (o *InstanceInterfacesNetworkInterfacesInner) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *InstanceInterfacesNetworkInterfacesInner) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *InstanceInterfacesNetworkInterfacesInner) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *InstanceInterfacesNetworkInterfacesInner) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### SetIpModeNil

`func (o *InstanceInterfacesNetworkInterfacesInner) SetIpModeNil(b bool)`

 SetIpModeNil sets the value for IpMode to be an explicit nil

### UnsetIpMode
`func (o *InstanceInterfacesNetworkInterfacesInner) UnsetIpMode()`

UnsetIpMode ensures that no value is present for IpMode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


