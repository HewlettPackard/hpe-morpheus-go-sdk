# AddClusterRequestClusterServerConfigAnyOfOneOf6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuArch** | Pointer to **string** |  | [optional] 
**CpuModel** | Pointer to **string** |  | [optional] 
**DynamicPlacementMode** | Pointer to **string** | When enabled, Dynamic Placement will automatically balance VMs across cluster hosts based on resource utilization. When disabled, VMs will only migrate to a new host if they are pinned to a specific host or failed over and not running on the preferred host. | [optional] 
**PowerPolicy** | Pointer to **string** |  | [optional] 
**VcpuPlacementMode** | Pointer to **string** |  | [optional] 
**StorageInterfaceName** | Pointer to **string** |  | [optional] 
**ComputeInterfaceName** | Pointer to **string** |  | [optional] 
**ComputeVlans** | Pointer to **string** |  | [optional] 
**OverlayInterfaceName** | Pointer to **string** |  | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddClusterRequestClusterServerConfigAnyOfOneOf6

`func NewAddClusterRequestClusterServerConfigAnyOfOneOf6() *AddClusterRequestClusterServerConfigAnyOfOneOf6`

NewAddClusterRequestClusterServerConfigAnyOfOneOf6 instantiates a new AddClusterRequestClusterServerConfigAnyOfOneOf6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClusterRequestClusterServerConfigAnyOfOneOf6WithDefaults

`func NewAddClusterRequestClusterServerConfigAnyOfOneOf6WithDefaults() *AddClusterRequestClusterServerConfigAnyOfOneOf6`

NewAddClusterRequestClusterServerConfigAnyOfOneOf6WithDefaults instantiates a new AddClusterRequestClusterServerConfigAnyOfOneOf6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuArch

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) GetCpuArch() string`

GetCpuArch returns the CpuArch field if non-nil, zero value otherwise.

### GetCpuArchOk

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) GetCpuArchOk() (*string, bool)`

GetCpuArchOk returns a tuple with the CpuArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuArch

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) SetCpuArch(v string)`

SetCpuArch sets CpuArch field to given value.

### HasCpuArch

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) HasCpuArch() bool`

HasCpuArch returns a boolean if a field has been set.

### GetCpuModel

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetDynamicPlacementMode

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) GetDynamicPlacementMode() string`

GetDynamicPlacementMode returns the DynamicPlacementMode field if non-nil, zero value otherwise.

### GetDynamicPlacementModeOk

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) GetDynamicPlacementModeOk() (*string, bool)`

GetDynamicPlacementModeOk returns a tuple with the DynamicPlacementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicPlacementMode

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) SetDynamicPlacementMode(v string)`

SetDynamicPlacementMode sets DynamicPlacementMode field to given value.

### HasDynamicPlacementMode

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) HasDynamicPlacementMode() bool`

HasDynamicPlacementMode returns a boolean if a field has been set.

### GetPowerPolicy

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) GetPowerPolicy() string`

GetPowerPolicy returns the PowerPolicy field if non-nil, zero value otherwise.

### GetPowerPolicyOk

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) GetPowerPolicyOk() (*string, bool)`

GetPowerPolicyOk returns a tuple with the PowerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPolicy

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) SetPowerPolicy(v string)`

SetPowerPolicy sets PowerPolicy field to given value.

### HasPowerPolicy

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) HasPowerPolicy() bool`

HasPowerPolicy returns a boolean if a field has been set.

### GetVcpuPlacementMode

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) GetVcpuPlacementMode() string`

GetVcpuPlacementMode returns the VcpuPlacementMode field if non-nil, zero value otherwise.

### GetVcpuPlacementModeOk

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) GetVcpuPlacementModeOk() (*string, bool)`

GetVcpuPlacementModeOk returns a tuple with the VcpuPlacementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpuPlacementMode

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) SetVcpuPlacementMode(v string)`

SetVcpuPlacementMode sets VcpuPlacementMode field to given value.

### HasVcpuPlacementMode

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) HasVcpuPlacementMode() bool`

HasVcpuPlacementMode returns a boolean if a field has been set.

### GetStorageInterfaceName

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) GetStorageInterfaceName() string`

GetStorageInterfaceName returns the StorageInterfaceName field if non-nil, zero value otherwise.

### GetStorageInterfaceNameOk

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) GetStorageInterfaceNameOk() (*string, bool)`

GetStorageInterfaceNameOk returns a tuple with the StorageInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInterfaceName

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) SetStorageInterfaceName(v string)`

SetStorageInterfaceName sets StorageInterfaceName field to given value.

### HasStorageInterfaceName

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) HasStorageInterfaceName() bool`

HasStorageInterfaceName returns a boolean if a field has been set.

### GetComputeInterfaceName

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) GetComputeInterfaceName() string`

GetComputeInterfaceName returns the ComputeInterfaceName field if non-nil, zero value otherwise.

### GetComputeInterfaceNameOk

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) GetComputeInterfaceNameOk() (*string, bool)`

GetComputeInterfaceNameOk returns a tuple with the ComputeInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeInterfaceName

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) SetComputeInterfaceName(v string)`

SetComputeInterfaceName sets ComputeInterfaceName field to given value.

### HasComputeInterfaceName

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) HasComputeInterfaceName() bool`

HasComputeInterfaceName returns a boolean if a field has been set.

### GetComputeVlans

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) GetComputeVlans() string`

GetComputeVlans returns the ComputeVlans field if non-nil, zero value otherwise.

### GetComputeVlansOk

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) GetComputeVlansOk() (*string, bool)`

GetComputeVlansOk returns a tuple with the ComputeVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeVlans

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) SetComputeVlans(v string)`

SetComputeVlans sets ComputeVlans field to given value.

### HasComputeVlans

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) HasComputeVlans() bool`

HasComputeVlans returns a boolean if a field has been set.

### GetOverlayInterfaceName

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) GetOverlayInterfaceName() string`

GetOverlayInterfaceName returns the OverlayInterfaceName field if non-nil, zero value otherwise.

### GetOverlayInterfaceNameOk

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) GetOverlayInterfaceNameOk() (*string, bool)`

GetOverlayInterfaceNameOk returns a tuple with the OverlayInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayInterfaceName

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) SetOverlayInterfaceName(v string)`

SetOverlayInterfaceName sets OverlayInterfaceName field to given value.

### HasOverlayInterfaceName

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) HasOverlayInterfaceName() bool`

HasOverlayInterfaceName returns a boolean if a field has been set.

### GetCreateUser

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf6) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


