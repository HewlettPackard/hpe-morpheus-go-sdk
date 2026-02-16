# InstanceInterfacesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerId**](AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerId.md) |  | [optional] 
**Network** | Pointer to [**AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetwork**](AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetwork.md) |  | [optional] 
**IpAddress** | Pointer to **NullableString** |  | [optional] 
**NetworkInterfaceTypeId** | Pointer to **NullableInt64** |  | [optional] 
**IpMode** | Pointer to **NullableString** |  | [optional] 
**NetworkInterfaces** | Pointer to [**[]InstanceInterfacesNetworkInterfacesInner1**](InstanceInterfacesNetworkInterfacesInner1.md) |  | [optional] 

## Methods

### NewInstanceInterfacesInner

`func NewInstanceInterfacesInner() *InstanceInterfacesInner`

NewInstanceInterfacesInner instantiates a new InstanceInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceInterfacesInnerWithDefaults

`func NewInstanceInterfacesInnerWithDefaults() *InstanceInterfacesInner`

NewInstanceInterfacesInnerWithDefaults instantiates a new InstanceInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceInterfacesInner) GetId() AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceInterfacesInner) GetIdOk() (*AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceInterfacesInner) SetId(v AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerId)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceInterfacesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetwork

`func (o *InstanceInterfacesInner) GetNetwork() AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InstanceInterfacesInner) GetNetworkOk() (*AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InstanceInterfacesInner) SetNetwork(v AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InstanceInterfacesInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetIpAddress

`func (o *InstanceInterfacesInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *InstanceInterfacesInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *InstanceInterfacesInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *InstanceInterfacesInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *InstanceInterfacesInner) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *InstanceInterfacesInner) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetNetworkInterfaceTypeId

`func (o *InstanceInterfacesInner) GetNetworkInterfaceTypeId() int64`

GetNetworkInterfaceTypeId returns the NetworkInterfaceTypeId field if non-nil, zero value otherwise.

### GetNetworkInterfaceTypeIdOk

`func (o *InstanceInterfacesInner) GetNetworkInterfaceTypeIdOk() (*int64, bool)`

GetNetworkInterfaceTypeIdOk returns a tuple with the NetworkInterfaceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceTypeId

`func (o *InstanceInterfacesInner) SetNetworkInterfaceTypeId(v int64)`

SetNetworkInterfaceTypeId sets NetworkInterfaceTypeId field to given value.

### HasNetworkInterfaceTypeId

`func (o *InstanceInterfacesInner) HasNetworkInterfaceTypeId() bool`

HasNetworkInterfaceTypeId returns a boolean if a field has been set.

### SetNetworkInterfaceTypeIdNil

`func (o *InstanceInterfacesInner) SetNetworkInterfaceTypeIdNil(b bool)`

 SetNetworkInterfaceTypeIdNil sets the value for NetworkInterfaceTypeId to be an explicit nil

### UnsetNetworkInterfaceTypeId
`func (o *InstanceInterfacesInner) UnsetNetworkInterfaceTypeId()`

UnsetNetworkInterfaceTypeId ensures that no value is present for NetworkInterfaceTypeId, not even an explicit nil
### GetIpMode

`func (o *InstanceInterfacesInner) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *InstanceInterfacesInner) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *InstanceInterfacesInner) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *InstanceInterfacesInner) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### SetIpModeNil

`func (o *InstanceInterfacesInner) SetIpModeNil(b bool)`

 SetIpModeNil sets the value for IpMode to be an explicit nil

### UnsetIpMode
`func (o *InstanceInterfacesInner) UnsetIpMode()`

UnsetIpMode ensures that no value is present for IpMode, not even an explicit nil
### GetNetworkInterfaces

`func (o *InstanceInterfacesInner) GetNetworkInterfaces() []InstanceInterfacesNetworkInterfacesInner1`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *InstanceInterfacesInner) GetNetworkInterfacesOk() (*[]InstanceInterfacesNetworkInterfacesInner1, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *InstanceInterfacesInner) SetNetworkInterfaces(v []InstanceInterfacesNetworkInterfacesInner1)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *InstanceInterfacesInner) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


