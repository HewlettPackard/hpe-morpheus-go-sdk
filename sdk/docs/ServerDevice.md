# ServerDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** | (Assignee) Target Server ID | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**DomainId** | Pointer to **int32** |  | [optional] 
**Bus** | Pointer to **int32** |  | [optional] 
**Slot** | Pointer to **int32** |  | [optional] 
**Device** | Pointer to **int32** |  | [optional] 
**VendorId** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**FunctionId** | Pointer to **int32** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**IommuGroup** | Pointer to **int32** |  | [optional] 
**IommuDeviceCount** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to [**ListHostDevices200ResponseDevicesInnerType**](ListHostDevices200ResponseDevicesInnerType.md) |  | [optional] 

## Methods

### NewServerDevice

`func NewServerDevice() *ServerDevice`

NewServerDevice instantiates a new ServerDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerDeviceWithDefaults

`func NewServerDeviceWithDefaults() *ServerDevice`

NewServerDeviceWithDefaults instantiates a new ServerDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerDevice) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerDevice) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerDevice) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ServerDevice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServerDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRefType

`func (o *ServerDevice) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ServerDevice) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ServerDevice) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ServerDevice) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *ServerDevice) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ServerDevice) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ServerDevice) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ServerDevice) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetStatus

`func (o *ServerDevice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerDevice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerDevice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServerDevice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExternalId

`func (o *ServerDevice) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ServerDevice) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ServerDevice) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ServerDevice) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDomainId

`func (o *ServerDevice) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *ServerDevice) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *ServerDevice) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *ServerDevice) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetBus

`func (o *ServerDevice) GetBus() int32`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *ServerDevice) GetBusOk() (*int32, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *ServerDevice) SetBus(v int32)`

SetBus sets Bus field to given value.

### HasBus

`func (o *ServerDevice) HasBus() bool`

HasBus returns a boolean if a field has been set.

### GetSlot

`func (o *ServerDevice) GetSlot() int32`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *ServerDevice) GetSlotOk() (*int32, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *ServerDevice) SetSlot(v int32)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *ServerDevice) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetDevice

`func (o *ServerDevice) GetDevice() int32`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ServerDevice) GetDeviceOk() (*int32, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ServerDevice) SetDevice(v int32)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *ServerDevice) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetVendorId

`func (o *ServerDevice) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *ServerDevice) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *ServerDevice) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *ServerDevice) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetProductId

`func (o *ServerDevice) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ServerDevice) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ServerDevice) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ServerDevice) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetFunctionId

`func (o *ServerDevice) GetFunctionId() int32`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *ServerDevice) GetFunctionIdOk() (*int32, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *ServerDevice) SetFunctionId(v int32)`

SetFunctionId sets FunctionId field to given value.

### HasFunctionId

`func (o *ServerDevice) HasFunctionId() bool`

HasFunctionId returns a boolean if a field has been set.

### GetUniqueId

`func (o *ServerDevice) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *ServerDevice) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *ServerDevice) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *ServerDevice) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetIommuGroup

`func (o *ServerDevice) GetIommuGroup() int32`

GetIommuGroup returns the IommuGroup field if non-nil, zero value otherwise.

### GetIommuGroupOk

`func (o *ServerDevice) GetIommuGroupOk() (*int32, bool)`

GetIommuGroupOk returns a tuple with the IommuGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIommuGroup

`func (o *ServerDevice) SetIommuGroup(v int32)`

SetIommuGroup sets IommuGroup field to given value.

### HasIommuGroup

`func (o *ServerDevice) HasIommuGroup() bool`

HasIommuGroup returns a boolean if a field has been set.

### GetIommuDeviceCount

`func (o *ServerDevice) GetIommuDeviceCount() int32`

GetIommuDeviceCount returns the IommuDeviceCount field if non-nil, zero value otherwise.

### GetIommuDeviceCountOk

`func (o *ServerDevice) GetIommuDeviceCountOk() (*int32, bool)`

GetIommuDeviceCountOk returns a tuple with the IommuDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIommuDeviceCount

`func (o *ServerDevice) SetIommuDeviceCount(v int32)`

SetIommuDeviceCount sets IommuDeviceCount field to given value.

### HasIommuDeviceCount

`func (o *ServerDevice) HasIommuDeviceCount() bool`

HasIommuDeviceCount returns a boolean if a field has been set.

### GetType

`func (o *ServerDevice) GetType() ListHostDevices200ResponseDevicesInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerDevice) GetTypeOk() (*ListHostDevices200ResponseDevicesInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerDevice) SetType(v ListHostDevices200ResponseDevicesInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *ServerDevice) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


