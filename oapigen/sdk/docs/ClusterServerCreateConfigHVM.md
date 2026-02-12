# ClusterServerCreateConfigHVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuArch** | Pointer to **string** |  | [optional] 
**CpuModel** | Pointer to **string** |  | [optional] 
**DynamicPlacementMode** | Pointer to **string** |  | [optional] 
**PowerPolicy** | Pointer to **string** |  | [optional] 
**StorageInterfaceName** | Pointer to **string** |  | [optional] 
**ComputeInterfaceName** | Pointer to **string** |  | [optional] 
**ComputeVlans** | Pointer to **string** |  | [optional] 
**OverlayInterfaceName** | Pointer to **string** |  | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterServerCreateConfigHVM

`func NewClusterServerCreateConfigHVM() *ClusterServerCreateConfigHVM`

NewClusterServerCreateConfigHVM instantiates a new ClusterServerCreateConfigHVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterServerCreateConfigHVMWithDefaults

`func NewClusterServerCreateConfigHVMWithDefaults() *ClusterServerCreateConfigHVM`

NewClusterServerCreateConfigHVMWithDefaults instantiates a new ClusterServerCreateConfigHVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuArch

`func (o *ClusterServerCreateConfigHVM) GetCpuArch() string`

GetCpuArch returns the CpuArch field if non-nil, zero value otherwise.

### GetCpuArchOk

`func (o *ClusterServerCreateConfigHVM) GetCpuArchOk() (*string, bool)`

GetCpuArchOk returns a tuple with the CpuArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuArch

`func (o *ClusterServerCreateConfigHVM) SetCpuArch(v string)`

SetCpuArch sets CpuArch field to given value.

### HasCpuArch

`func (o *ClusterServerCreateConfigHVM) HasCpuArch() bool`

HasCpuArch returns a boolean if a field has been set.

### GetCpuModel

`func (o *ClusterServerCreateConfigHVM) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *ClusterServerCreateConfigHVM) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *ClusterServerCreateConfigHVM) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *ClusterServerCreateConfigHVM) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetDynamicPlacementMode

`func (o *ClusterServerCreateConfigHVM) GetDynamicPlacementMode() string`

GetDynamicPlacementMode returns the DynamicPlacementMode field if non-nil, zero value otherwise.

### GetDynamicPlacementModeOk

`func (o *ClusterServerCreateConfigHVM) GetDynamicPlacementModeOk() (*string, bool)`

GetDynamicPlacementModeOk returns a tuple with the DynamicPlacementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicPlacementMode

`func (o *ClusterServerCreateConfigHVM) SetDynamicPlacementMode(v string)`

SetDynamicPlacementMode sets DynamicPlacementMode field to given value.

### HasDynamicPlacementMode

`func (o *ClusterServerCreateConfigHVM) HasDynamicPlacementMode() bool`

HasDynamicPlacementMode returns a boolean if a field has been set.

### GetPowerPolicy

`func (o *ClusterServerCreateConfigHVM) GetPowerPolicy() string`

GetPowerPolicy returns the PowerPolicy field if non-nil, zero value otherwise.

### GetPowerPolicyOk

`func (o *ClusterServerCreateConfigHVM) GetPowerPolicyOk() (*string, bool)`

GetPowerPolicyOk returns a tuple with the PowerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPolicy

`func (o *ClusterServerCreateConfigHVM) SetPowerPolicy(v string)`

SetPowerPolicy sets PowerPolicy field to given value.

### HasPowerPolicy

`func (o *ClusterServerCreateConfigHVM) HasPowerPolicy() bool`

HasPowerPolicy returns a boolean if a field has been set.

### GetStorageInterfaceName

`func (o *ClusterServerCreateConfigHVM) GetStorageInterfaceName() string`

GetStorageInterfaceName returns the StorageInterfaceName field if non-nil, zero value otherwise.

### GetStorageInterfaceNameOk

`func (o *ClusterServerCreateConfigHVM) GetStorageInterfaceNameOk() (*string, bool)`

GetStorageInterfaceNameOk returns a tuple with the StorageInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInterfaceName

`func (o *ClusterServerCreateConfigHVM) SetStorageInterfaceName(v string)`

SetStorageInterfaceName sets StorageInterfaceName field to given value.

### HasStorageInterfaceName

`func (o *ClusterServerCreateConfigHVM) HasStorageInterfaceName() bool`

HasStorageInterfaceName returns a boolean if a field has been set.

### GetComputeInterfaceName

`func (o *ClusterServerCreateConfigHVM) GetComputeInterfaceName() string`

GetComputeInterfaceName returns the ComputeInterfaceName field if non-nil, zero value otherwise.

### GetComputeInterfaceNameOk

`func (o *ClusterServerCreateConfigHVM) GetComputeInterfaceNameOk() (*string, bool)`

GetComputeInterfaceNameOk returns a tuple with the ComputeInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeInterfaceName

`func (o *ClusterServerCreateConfigHVM) SetComputeInterfaceName(v string)`

SetComputeInterfaceName sets ComputeInterfaceName field to given value.

### HasComputeInterfaceName

`func (o *ClusterServerCreateConfigHVM) HasComputeInterfaceName() bool`

HasComputeInterfaceName returns a boolean if a field has been set.

### GetComputeVlans

`func (o *ClusterServerCreateConfigHVM) GetComputeVlans() string`

GetComputeVlans returns the ComputeVlans field if non-nil, zero value otherwise.

### GetComputeVlansOk

`func (o *ClusterServerCreateConfigHVM) GetComputeVlansOk() (*string, bool)`

GetComputeVlansOk returns a tuple with the ComputeVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeVlans

`func (o *ClusterServerCreateConfigHVM) SetComputeVlans(v string)`

SetComputeVlans sets ComputeVlans field to given value.

### HasComputeVlans

`func (o *ClusterServerCreateConfigHVM) HasComputeVlans() bool`

HasComputeVlans returns a boolean if a field has been set.

### GetOverlayInterfaceName

`func (o *ClusterServerCreateConfigHVM) GetOverlayInterfaceName() string`

GetOverlayInterfaceName returns the OverlayInterfaceName field if non-nil, zero value otherwise.

### GetOverlayInterfaceNameOk

`func (o *ClusterServerCreateConfigHVM) GetOverlayInterfaceNameOk() (*string, bool)`

GetOverlayInterfaceNameOk returns a tuple with the OverlayInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayInterfaceName

`func (o *ClusterServerCreateConfigHVM) SetOverlayInterfaceName(v string)`

SetOverlayInterfaceName sets OverlayInterfaceName field to given value.

### HasOverlayInterfaceName

`func (o *ClusterServerCreateConfigHVM) HasOverlayInterfaceName() bool`

HasOverlayInterfaceName returns a boolean if a field has been set.

### GetCreateUser

`func (o *ClusterServerCreateConfigHVM) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ClusterServerCreateConfigHVM) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ClusterServerCreateConfigHVM) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ClusterServerCreateConfigHVM) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


