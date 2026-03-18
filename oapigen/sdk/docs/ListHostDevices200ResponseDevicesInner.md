# ListHostDevices200ResponseDevicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableInt64** | (Assignee) Target Server ID | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**DomainId** | Pointer to **NullableInt32** |  | [optional] 
**Bus** | Pointer to **NullableInt32** |  | [optional] 
**Slot** | Pointer to **NullableInt32** |  | [optional] 
**Device** | Pointer to **NullableInt32** |  | [optional] 
**VendorId** | Pointer to **NullableString** |  | [optional] 
**ProductId** | Pointer to **NullableString** |  | [optional] 
**FunctionId** | Pointer to **NullableInt32** |  | [optional] 
**UniqueId** | Pointer to **NullableString** |  | [optional] 
**IommuGroup** | Pointer to **NullableInt32** |  | [optional] 
**IommuDeviceCount** | Pointer to **NullableInt32** |  | [optional] 
**Type** | Pointer to [**ListHostDevices200ResponseDevicesInnerType**](ListHostDevices200ResponseDevicesInnerType.md) |  | [optional] 

## Methods

### NewListHostDevices200ResponseDevicesInner

`func NewListHostDevices200ResponseDevicesInner() *ListHostDevices200ResponseDevicesInner`

NewListHostDevices200ResponseDevicesInner instantiates a new ListHostDevices200ResponseDevicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHostDevices200ResponseDevicesInnerWithDefaults

`func NewListHostDevices200ResponseDevicesInnerWithDefaults() *ListHostDevices200ResponseDevicesInner`

NewListHostDevices200ResponseDevicesInnerWithDefaults instantiates a new ListHostDevices200ResponseDevicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListHostDevices200ResponseDevicesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListHostDevices200ResponseDevicesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListHostDevices200ResponseDevicesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListHostDevices200ResponseDevicesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListHostDevices200ResponseDevicesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListHostDevices200ResponseDevicesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListHostDevices200ResponseDevicesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListHostDevices200ResponseDevicesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRefType

`func (o *ListHostDevices200ResponseDevicesInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListHostDevices200ResponseDevicesInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListHostDevices200ResponseDevicesInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListHostDevices200ResponseDevicesInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *ListHostDevices200ResponseDevicesInner) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *ListHostDevices200ResponseDevicesInner) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *ListHostDevices200ResponseDevicesInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListHostDevices200ResponseDevicesInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListHostDevices200ResponseDevicesInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListHostDevices200ResponseDevicesInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *ListHostDevices200ResponseDevicesInner) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *ListHostDevices200ResponseDevicesInner) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetStatus

`func (o *ListHostDevices200ResponseDevicesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListHostDevices200ResponseDevicesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListHostDevices200ResponseDevicesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListHostDevices200ResponseDevicesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExternalId

`func (o *ListHostDevices200ResponseDevicesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListHostDevices200ResponseDevicesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListHostDevices200ResponseDevicesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListHostDevices200ResponseDevicesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ListHostDevices200ResponseDevicesInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ListHostDevices200ResponseDevicesInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetDomainId

`func (o *ListHostDevices200ResponseDevicesInner) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *ListHostDevices200ResponseDevicesInner) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *ListHostDevices200ResponseDevicesInner) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *ListHostDevices200ResponseDevicesInner) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### SetDomainIdNil

`func (o *ListHostDevices200ResponseDevicesInner) SetDomainIdNil(b bool)`

 SetDomainIdNil sets the value for DomainId to be an explicit nil

### UnsetDomainId
`func (o *ListHostDevices200ResponseDevicesInner) UnsetDomainId()`

UnsetDomainId ensures that no value is present for DomainId, not even an explicit nil
### GetBus

`func (o *ListHostDevices200ResponseDevicesInner) GetBus() int32`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *ListHostDevices200ResponseDevicesInner) GetBusOk() (*int32, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *ListHostDevices200ResponseDevicesInner) SetBus(v int32)`

SetBus sets Bus field to given value.

### HasBus

`func (o *ListHostDevices200ResponseDevicesInner) HasBus() bool`

HasBus returns a boolean if a field has been set.

### SetBusNil

`func (o *ListHostDevices200ResponseDevicesInner) SetBusNil(b bool)`

 SetBusNil sets the value for Bus to be an explicit nil

### UnsetBus
`func (o *ListHostDevices200ResponseDevicesInner) UnsetBus()`

UnsetBus ensures that no value is present for Bus, not even an explicit nil
### GetSlot

`func (o *ListHostDevices200ResponseDevicesInner) GetSlot() int32`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *ListHostDevices200ResponseDevicesInner) GetSlotOk() (*int32, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *ListHostDevices200ResponseDevicesInner) SetSlot(v int32)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *ListHostDevices200ResponseDevicesInner) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### SetSlotNil

`func (o *ListHostDevices200ResponseDevicesInner) SetSlotNil(b bool)`

 SetSlotNil sets the value for Slot to be an explicit nil

### UnsetSlot
`func (o *ListHostDevices200ResponseDevicesInner) UnsetSlot()`

UnsetSlot ensures that no value is present for Slot, not even an explicit nil
### GetDevice

`func (o *ListHostDevices200ResponseDevicesInner) GetDevice() int32`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ListHostDevices200ResponseDevicesInner) GetDeviceOk() (*int32, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ListHostDevices200ResponseDevicesInner) SetDevice(v int32)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *ListHostDevices200ResponseDevicesInner) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *ListHostDevices200ResponseDevicesInner) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *ListHostDevices200ResponseDevicesInner) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetVendorId

`func (o *ListHostDevices200ResponseDevicesInner) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *ListHostDevices200ResponseDevicesInner) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *ListHostDevices200ResponseDevicesInner) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *ListHostDevices200ResponseDevicesInner) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### SetVendorIdNil

`func (o *ListHostDevices200ResponseDevicesInner) SetVendorIdNil(b bool)`

 SetVendorIdNil sets the value for VendorId to be an explicit nil

### UnsetVendorId
`func (o *ListHostDevices200ResponseDevicesInner) UnsetVendorId()`

UnsetVendorId ensures that no value is present for VendorId, not even an explicit nil
### GetProductId

`func (o *ListHostDevices200ResponseDevicesInner) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ListHostDevices200ResponseDevicesInner) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ListHostDevices200ResponseDevicesInner) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ListHostDevices200ResponseDevicesInner) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *ListHostDevices200ResponseDevicesInner) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *ListHostDevices200ResponseDevicesInner) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetFunctionId

`func (o *ListHostDevices200ResponseDevicesInner) GetFunctionId() int32`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *ListHostDevices200ResponseDevicesInner) GetFunctionIdOk() (*int32, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *ListHostDevices200ResponseDevicesInner) SetFunctionId(v int32)`

SetFunctionId sets FunctionId field to given value.

### HasFunctionId

`func (o *ListHostDevices200ResponseDevicesInner) HasFunctionId() bool`

HasFunctionId returns a boolean if a field has been set.

### SetFunctionIdNil

`func (o *ListHostDevices200ResponseDevicesInner) SetFunctionIdNil(b bool)`

 SetFunctionIdNil sets the value for FunctionId to be an explicit nil

### UnsetFunctionId
`func (o *ListHostDevices200ResponseDevicesInner) UnsetFunctionId()`

UnsetFunctionId ensures that no value is present for FunctionId, not even an explicit nil
### GetUniqueId

`func (o *ListHostDevices200ResponseDevicesInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *ListHostDevices200ResponseDevicesInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *ListHostDevices200ResponseDevicesInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *ListHostDevices200ResponseDevicesInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### SetUniqueIdNil

`func (o *ListHostDevices200ResponseDevicesInner) SetUniqueIdNil(b bool)`

 SetUniqueIdNil sets the value for UniqueId to be an explicit nil

### UnsetUniqueId
`func (o *ListHostDevices200ResponseDevicesInner) UnsetUniqueId()`

UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
### GetIommuGroup

`func (o *ListHostDevices200ResponseDevicesInner) GetIommuGroup() int32`

GetIommuGroup returns the IommuGroup field if non-nil, zero value otherwise.

### GetIommuGroupOk

`func (o *ListHostDevices200ResponseDevicesInner) GetIommuGroupOk() (*int32, bool)`

GetIommuGroupOk returns a tuple with the IommuGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIommuGroup

`func (o *ListHostDevices200ResponseDevicesInner) SetIommuGroup(v int32)`

SetIommuGroup sets IommuGroup field to given value.

### HasIommuGroup

`func (o *ListHostDevices200ResponseDevicesInner) HasIommuGroup() bool`

HasIommuGroup returns a boolean if a field has been set.

### SetIommuGroupNil

`func (o *ListHostDevices200ResponseDevicesInner) SetIommuGroupNil(b bool)`

 SetIommuGroupNil sets the value for IommuGroup to be an explicit nil

### UnsetIommuGroup
`func (o *ListHostDevices200ResponseDevicesInner) UnsetIommuGroup()`

UnsetIommuGroup ensures that no value is present for IommuGroup, not even an explicit nil
### GetIommuDeviceCount

`func (o *ListHostDevices200ResponseDevicesInner) GetIommuDeviceCount() int32`

GetIommuDeviceCount returns the IommuDeviceCount field if non-nil, zero value otherwise.

### GetIommuDeviceCountOk

`func (o *ListHostDevices200ResponseDevicesInner) GetIommuDeviceCountOk() (*int32, bool)`

GetIommuDeviceCountOk returns a tuple with the IommuDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIommuDeviceCount

`func (o *ListHostDevices200ResponseDevicesInner) SetIommuDeviceCount(v int32)`

SetIommuDeviceCount sets IommuDeviceCount field to given value.

### HasIommuDeviceCount

`func (o *ListHostDevices200ResponseDevicesInner) HasIommuDeviceCount() bool`

HasIommuDeviceCount returns a boolean if a field has been set.

### SetIommuDeviceCountNil

`func (o *ListHostDevices200ResponseDevicesInner) SetIommuDeviceCountNil(b bool)`

 SetIommuDeviceCountNil sets the value for IommuDeviceCount to be an explicit nil

### UnsetIommuDeviceCount
`func (o *ListHostDevices200ResponseDevicesInner) UnsetIommuDeviceCount()`

UnsetIommuDeviceCount ensures that no value is present for IommuDeviceCount, not even an explicit nil
### GetType

`func (o *ListHostDevices200ResponseDevicesInner) GetType() ListHostDevices200ResponseDevicesInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListHostDevices200ResponseDevicesInner) GetTypeOk() (*ListHostDevices200ResponseDevicesInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListHostDevices200ResponseDevicesInner) SetType(v ListHostDevices200ResponseDevicesInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *ListHostDevices200ResponseDevicesInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


