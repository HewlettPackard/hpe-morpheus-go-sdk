# GetInstance200ResponseInstanceInterfacesInner

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

### NewGetInstance200ResponseInstanceInterfacesInner

`func NewGetInstance200ResponseInstanceInterfacesInner() *GetInstance200ResponseInstanceInterfacesInner`

NewGetInstance200ResponseInstanceInterfacesInner instantiates a new GetInstance200ResponseInstanceInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstance200ResponseInstanceInterfacesInnerWithDefaults

`func NewGetInstance200ResponseInstanceInterfacesInnerWithDefaults() *GetInstance200ResponseInstanceInterfacesInner`

NewGetInstance200ResponseInstanceInterfacesInnerWithDefaults instantiates a new GetInstance200ResponseInstanceInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInstance200ResponseInstanceInterfacesInner) GetId() AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInstance200ResponseInstanceInterfacesInner) GetIdOk() (*AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInstance200ResponseInstanceInterfacesInner) SetId(v AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerId)`

SetId sets Id field to given value.

### HasId

`func (o *GetInstance200ResponseInstanceInterfacesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetwork

`func (o *GetInstance200ResponseInstanceInterfacesInner) GetNetwork() AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetInstance200ResponseInstanceInterfacesInner) GetNetworkOk() (*AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetInstance200ResponseInstanceInterfacesInner) SetNetwork(v AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetInstance200ResponseInstanceInterfacesInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetIpAddress

`func (o *GetInstance200ResponseInstanceInterfacesInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *GetInstance200ResponseInstanceInterfacesInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *GetInstance200ResponseInstanceInterfacesInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *GetInstance200ResponseInstanceInterfacesInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *GetInstance200ResponseInstanceInterfacesInner) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *GetInstance200ResponseInstanceInterfacesInner) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetNetworkInterfaceTypeId

`func (o *GetInstance200ResponseInstanceInterfacesInner) GetNetworkInterfaceTypeId() int64`

GetNetworkInterfaceTypeId returns the NetworkInterfaceTypeId field if non-nil, zero value otherwise.

### GetNetworkInterfaceTypeIdOk

`func (o *GetInstance200ResponseInstanceInterfacesInner) GetNetworkInterfaceTypeIdOk() (*int64, bool)`

GetNetworkInterfaceTypeIdOk returns a tuple with the NetworkInterfaceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceTypeId

`func (o *GetInstance200ResponseInstanceInterfacesInner) SetNetworkInterfaceTypeId(v int64)`

SetNetworkInterfaceTypeId sets NetworkInterfaceTypeId field to given value.

### HasNetworkInterfaceTypeId

`func (o *GetInstance200ResponseInstanceInterfacesInner) HasNetworkInterfaceTypeId() bool`

HasNetworkInterfaceTypeId returns a boolean if a field has been set.

### SetNetworkInterfaceTypeIdNil

`func (o *GetInstance200ResponseInstanceInterfacesInner) SetNetworkInterfaceTypeIdNil(b bool)`

 SetNetworkInterfaceTypeIdNil sets the value for NetworkInterfaceTypeId to be an explicit nil

### UnsetNetworkInterfaceTypeId
`func (o *GetInstance200ResponseInstanceInterfacesInner) UnsetNetworkInterfaceTypeId()`

UnsetNetworkInterfaceTypeId ensures that no value is present for NetworkInterfaceTypeId, not even an explicit nil
### GetIpMode

`func (o *GetInstance200ResponseInstanceInterfacesInner) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *GetInstance200ResponseInstanceInterfacesInner) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *GetInstance200ResponseInstanceInterfacesInner) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *GetInstance200ResponseInstanceInterfacesInner) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### SetIpModeNil

`func (o *GetInstance200ResponseInstanceInterfacesInner) SetIpModeNil(b bool)`

 SetIpModeNil sets the value for IpMode to be an explicit nil

### UnsetIpMode
`func (o *GetInstance200ResponseInstanceInterfacesInner) UnsetIpMode()`

UnsetIpMode ensures that no value is present for IpMode, not even an explicit nil
### GetNetworkInterfaces

`func (o *GetInstance200ResponseInstanceInterfacesInner) GetNetworkInterfaces() []InstanceInterfacesNetworkInterfacesInner1`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *GetInstance200ResponseInstanceInterfacesInner) GetNetworkInterfacesOk() (*[]InstanceInterfacesNetworkInterfacesInner1, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *GetInstance200ResponseInstanceInterfacesInner) SetNetworkInterfaces(v []InstanceInterfacesNetworkInterfacesInner1)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *GetInstance200ResponseInstanceInterfacesInner) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


