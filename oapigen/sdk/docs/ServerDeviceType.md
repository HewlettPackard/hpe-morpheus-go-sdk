# ServerDeviceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Family** | Pointer to **NullableString** |  | [optional] 
**BusType** | Pointer to **NullableString** |  | [optional] 
**Assignable** | Pointer to **bool** |  | [optional] 
**Hotpluggable** | Pointer to **bool** |  | [optional] 
**VendorId** | Pointer to **NullableInt32** |  | [optional] 
**ProductId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewServerDeviceType

`func NewServerDeviceType() *ServerDeviceType`

NewServerDeviceType instantiates a new ServerDeviceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *ServerDeviceType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerDeviceType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerDeviceType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ServerDeviceType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServerDeviceType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerDeviceType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerDeviceType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerDeviceType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFamily

`func (o *ServerDeviceType) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *ServerDeviceType) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *ServerDeviceType) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *ServerDeviceType) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### SetFamilyNil

`func (o *ServerDeviceType) SetFamilyNil(b bool)`

 SetFamilyNil sets the value for Family to be an explicit nil

### UnsetFamily
`func (o *ServerDeviceType) UnsetFamily()`

UnsetFamily ensures that no value is present for Family, not even an explicit nil
### GetBusType

`func (o *ServerDeviceType) GetBusType() string`

GetBusType returns the BusType field if non-nil, zero value otherwise.

### GetBusTypeOk

`func (o *ServerDeviceType) GetBusTypeOk() (*string, bool)`

GetBusTypeOk returns a tuple with the BusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusType

`func (o *ServerDeviceType) SetBusType(v string)`

SetBusType sets BusType field to given value.

### HasBusType

`func (o *ServerDeviceType) HasBusType() bool`

HasBusType returns a boolean if a field has been set.

### SetBusTypeNil

`func (o *ServerDeviceType) SetBusTypeNil(b bool)`

 SetBusTypeNil sets the value for BusType to be an explicit nil

### UnsetBusType
`func (o *ServerDeviceType) UnsetBusType()`

UnsetBusType ensures that no value is present for BusType, not even an explicit nil
### GetAssignable

`func (o *ServerDeviceType) GetAssignable() bool`

GetAssignable returns the Assignable field if non-nil, zero value otherwise.

### GetAssignableOk

`func (o *ServerDeviceType) GetAssignableOk() (*bool, bool)`

GetAssignableOk returns a tuple with the Assignable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignable

`func (o *ServerDeviceType) SetAssignable(v bool)`

SetAssignable sets Assignable field to given value.

### HasAssignable

`func (o *ServerDeviceType) HasAssignable() bool`

HasAssignable returns a boolean if a field has been set.

### GetHotpluggable

`func (o *ServerDeviceType) GetHotpluggable() bool`

GetHotpluggable returns the Hotpluggable field if non-nil, zero value otherwise.

### GetHotpluggableOk

`func (o *ServerDeviceType) GetHotpluggableOk() (*bool, bool)`

GetHotpluggableOk returns a tuple with the Hotpluggable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotpluggable

`func (o *ServerDeviceType) SetHotpluggable(v bool)`

SetHotpluggable sets Hotpluggable field to given value.

### HasHotpluggable

`func (o *ServerDeviceType) HasHotpluggable() bool`

HasHotpluggable returns a boolean if a field has been set.

### GetVendorId

`func (o *ServerDeviceType) GetVendorId() int32`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *ServerDeviceType) GetVendorIdOk() (*int32, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *ServerDeviceType) SetVendorId(v int32)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *ServerDeviceType) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### SetVendorIdNil

`func (o *ServerDeviceType) SetVendorIdNil(b bool)`

 SetVendorIdNil sets the value for VendorId to be an explicit nil

### UnsetVendorId
`func (o *ServerDeviceType) UnsetVendorId()`

UnsetVendorId ensures that no value is present for VendorId, not even an explicit nil
### GetProductId

`func (o *ServerDeviceType) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ServerDeviceType) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ServerDeviceType) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ServerDeviceType) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *ServerDeviceType) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *ServerDeviceType) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


