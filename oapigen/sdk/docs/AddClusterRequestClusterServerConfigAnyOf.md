# AddClusterRequestClusterServerConfigAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeCount** | Pointer to **int64** | Number of workers or hosts | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 
**Controller** | Pointer to [**AddClusterRequestClusterServerConfigAnyOfOneOf2Controller**](AddClusterRequestClusterServerConfigAnyOfOneOf2Controller.md) |  | [optional] 
**Worker** | Pointer to [**AddClusterRequestClusterServerConfigAnyOfOneOf3Worker**](AddClusterRequestClusterServerConfigAnyOfOneOf3Worker.md) |  | [optional] 
**PublicIpType** | Pointer to **string** |  | [optional] 
**GoogleZoneId** | Pointer to **int64** |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**ControlPlaneVersion** | Pointer to **string** |  | [optional] 
**PodCidr** | Pointer to **string** |  | [optional] 
**ServiceCidr** | Pointer to **string** |  | [optional] 
**CpuArch** | Pointer to **string** |  | [optional] 
**CpuModel** | Pointer to **string** |  | [optional] 
**DynamicPlacementMode** | Pointer to **string** |  | [optional] 
**PowerPolicy** | Pointer to **string** |  | [optional] 
**StorageInterfaceName** | Pointer to **string** |  | [optional] 
**ComputeInterfaceName** | Pointer to **string** |  | [optional] 
**ComputeVlans** | Pointer to **string** |  | [optional] 
**OverlayInterfaceName** | Pointer to **string** |  | [optional] 
**VmwareFolderId** | Pointer to **string** |  | [optional] 

## Methods

### NewAddClusterRequestClusterServerConfigAnyOf

`func NewAddClusterRequestClusterServerConfigAnyOf() *AddClusterRequestClusterServerConfigAnyOf`

NewAddClusterRequestClusterServerConfigAnyOf instantiates a new AddClusterRequestClusterServerConfigAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClusterRequestClusterServerConfigAnyOfWithDefaults

`func NewAddClusterRequestClusterServerConfigAnyOfWithDefaults() *AddClusterRequestClusterServerConfigAnyOf`

NewAddClusterRequestClusterServerConfigAnyOfWithDefaults instantiates a new AddClusterRequestClusterServerConfigAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeCount

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *AddClusterRequestClusterServerConfigAnyOf) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *AddClusterRequestClusterServerConfigAnyOf) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetCreateUser

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *AddClusterRequestClusterServerConfigAnyOf) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *AddClusterRequestClusterServerConfigAnyOf) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetController

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetController() AddClusterRequestClusterServerConfigAnyOfOneOf2Controller`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetControllerOk() (*AddClusterRequestClusterServerConfigAnyOfOneOf2Controller, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *AddClusterRequestClusterServerConfigAnyOf) SetController(v AddClusterRequestClusterServerConfigAnyOfOneOf2Controller)`

SetController sets Controller field to given value.

### HasController

`func (o *AddClusterRequestClusterServerConfigAnyOf) HasController() bool`

HasController returns a boolean if a field has been set.

### GetWorker

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetWorker() AddClusterRequestClusterServerConfigAnyOfOneOf3Worker`

GetWorker returns the Worker field if non-nil, zero value otherwise.

### GetWorkerOk

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetWorkerOk() (*AddClusterRequestClusterServerConfigAnyOfOneOf3Worker, bool)`

GetWorkerOk returns a tuple with the Worker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorker

`func (o *AddClusterRequestClusterServerConfigAnyOf) SetWorker(v AddClusterRequestClusterServerConfigAnyOfOneOf3Worker)`

SetWorker sets Worker field to given value.

### HasWorker

`func (o *AddClusterRequestClusterServerConfigAnyOf) HasWorker() bool`

HasWorker returns a boolean if a field has been set.

### GetPublicIpType

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetPublicIpType() string`

GetPublicIpType returns the PublicIpType field if non-nil, zero value otherwise.

### GetPublicIpTypeOk

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetPublicIpTypeOk() (*string, bool)`

GetPublicIpTypeOk returns a tuple with the PublicIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpType

`func (o *AddClusterRequestClusterServerConfigAnyOf) SetPublicIpType(v string)`

SetPublicIpType sets PublicIpType field to given value.

### HasPublicIpType

`func (o *AddClusterRequestClusterServerConfigAnyOf) HasPublicIpType() bool`

HasPublicIpType returns a boolean if a field has been set.

### GetGoogleZoneId

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetGoogleZoneId() int64`

GetGoogleZoneId returns the GoogleZoneId field if non-nil, zero value otherwise.

### GetGoogleZoneIdOk

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetGoogleZoneIdOk() (*int64, bool)`

GetGoogleZoneIdOk returns a tuple with the GoogleZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleZoneId

`func (o *AddClusterRequestClusterServerConfigAnyOf) SetGoogleZoneId(v int64)`

SetGoogleZoneId sets GoogleZoneId field to given value.

### HasGoogleZoneId

`func (o *AddClusterRequestClusterServerConfigAnyOf) HasGoogleZoneId() bool`

HasGoogleZoneId returns a boolean if a field has been set.

### GetChannel

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *AddClusterRequestClusterServerConfigAnyOf) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *AddClusterRequestClusterServerConfigAnyOf) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetControlPlaneVersion

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetControlPlaneVersion() string`

GetControlPlaneVersion returns the ControlPlaneVersion field if non-nil, zero value otherwise.

### GetControlPlaneVersionOk

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetControlPlaneVersionOk() (*string, bool)`

GetControlPlaneVersionOk returns a tuple with the ControlPlaneVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneVersion

`func (o *AddClusterRequestClusterServerConfigAnyOf) SetControlPlaneVersion(v string)`

SetControlPlaneVersion sets ControlPlaneVersion field to given value.

### HasControlPlaneVersion

`func (o *AddClusterRequestClusterServerConfigAnyOf) HasControlPlaneVersion() bool`

HasControlPlaneVersion returns a boolean if a field has been set.

### GetPodCidr

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetPodCidr() string`

GetPodCidr returns the PodCidr field if non-nil, zero value otherwise.

### GetPodCidrOk

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetPodCidrOk() (*string, bool)`

GetPodCidrOk returns a tuple with the PodCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCidr

`func (o *AddClusterRequestClusterServerConfigAnyOf) SetPodCidr(v string)`

SetPodCidr sets PodCidr field to given value.

### HasPodCidr

`func (o *AddClusterRequestClusterServerConfigAnyOf) HasPodCidr() bool`

HasPodCidr returns a boolean if a field has been set.

### GetServiceCidr

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetServiceCidr() string`

GetServiceCidr returns the ServiceCidr field if non-nil, zero value otherwise.

### GetServiceCidrOk

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetServiceCidrOk() (*string, bool)`

GetServiceCidrOk returns a tuple with the ServiceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCidr

`func (o *AddClusterRequestClusterServerConfigAnyOf) SetServiceCidr(v string)`

SetServiceCidr sets ServiceCidr field to given value.

### HasServiceCidr

`func (o *AddClusterRequestClusterServerConfigAnyOf) HasServiceCidr() bool`

HasServiceCidr returns a boolean if a field has been set.

### GetCpuArch

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetCpuArch() string`

GetCpuArch returns the CpuArch field if non-nil, zero value otherwise.

### GetCpuArchOk

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetCpuArchOk() (*string, bool)`

GetCpuArchOk returns a tuple with the CpuArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuArch

`func (o *AddClusterRequestClusterServerConfigAnyOf) SetCpuArch(v string)`

SetCpuArch sets CpuArch field to given value.

### HasCpuArch

`func (o *AddClusterRequestClusterServerConfigAnyOf) HasCpuArch() bool`

HasCpuArch returns a boolean if a field has been set.

### GetCpuModel

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *AddClusterRequestClusterServerConfigAnyOf) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *AddClusterRequestClusterServerConfigAnyOf) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetDynamicPlacementMode

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetDynamicPlacementMode() string`

GetDynamicPlacementMode returns the DynamicPlacementMode field if non-nil, zero value otherwise.

### GetDynamicPlacementModeOk

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetDynamicPlacementModeOk() (*string, bool)`

GetDynamicPlacementModeOk returns a tuple with the DynamicPlacementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicPlacementMode

`func (o *AddClusterRequestClusterServerConfigAnyOf) SetDynamicPlacementMode(v string)`

SetDynamicPlacementMode sets DynamicPlacementMode field to given value.

### HasDynamicPlacementMode

`func (o *AddClusterRequestClusterServerConfigAnyOf) HasDynamicPlacementMode() bool`

HasDynamicPlacementMode returns a boolean if a field has been set.

### GetPowerPolicy

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetPowerPolicy() string`

GetPowerPolicy returns the PowerPolicy field if non-nil, zero value otherwise.

### GetPowerPolicyOk

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetPowerPolicyOk() (*string, bool)`

GetPowerPolicyOk returns a tuple with the PowerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPolicy

`func (o *AddClusterRequestClusterServerConfigAnyOf) SetPowerPolicy(v string)`

SetPowerPolicy sets PowerPolicy field to given value.

### HasPowerPolicy

`func (o *AddClusterRequestClusterServerConfigAnyOf) HasPowerPolicy() bool`

HasPowerPolicy returns a boolean if a field has been set.

### GetStorageInterfaceName

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetStorageInterfaceName() string`

GetStorageInterfaceName returns the StorageInterfaceName field if non-nil, zero value otherwise.

### GetStorageInterfaceNameOk

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetStorageInterfaceNameOk() (*string, bool)`

GetStorageInterfaceNameOk returns a tuple with the StorageInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInterfaceName

`func (o *AddClusterRequestClusterServerConfigAnyOf) SetStorageInterfaceName(v string)`

SetStorageInterfaceName sets StorageInterfaceName field to given value.

### HasStorageInterfaceName

`func (o *AddClusterRequestClusterServerConfigAnyOf) HasStorageInterfaceName() bool`

HasStorageInterfaceName returns a boolean if a field has been set.

### GetComputeInterfaceName

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetComputeInterfaceName() string`

GetComputeInterfaceName returns the ComputeInterfaceName field if non-nil, zero value otherwise.

### GetComputeInterfaceNameOk

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetComputeInterfaceNameOk() (*string, bool)`

GetComputeInterfaceNameOk returns a tuple with the ComputeInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeInterfaceName

`func (o *AddClusterRequestClusterServerConfigAnyOf) SetComputeInterfaceName(v string)`

SetComputeInterfaceName sets ComputeInterfaceName field to given value.

### HasComputeInterfaceName

`func (o *AddClusterRequestClusterServerConfigAnyOf) HasComputeInterfaceName() bool`

HasComputeInterfaceName returns a boolean if a field has been set.

### GetComputeVlans

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetComputeVlans() string`

GetComputeVlans returns the ComputeVlans field if non-nil, zero value otherwise.

### GetComputeVlansOk

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetComputeVlansOk() (*string, bool)`

GetComputeVlansOk returns a tuple with the ComputeVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeVlans

`func (o *AddClusterRequestClusterServerConfigAnyOf) SetComputeVlans(v string)`

SetComputeVlans sets ComputeVlans field to given value.

### HasComputeVlans

`func (o *AddClusterRequestClusterServerConfigAnyOf) HasComputeVlans() bool`

HasComputeVlans returns a boolean if a field has been set.

### GetOverlayInterfaceName

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetOverlayInterfaceName() string`

GetOverlayInterfaceName returns the OverlayInterfaceName field if non-nil, zero value otherwise.

### GetOverlayInterfaceNameOk

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetOverlayInterfaceNameOk() (*string, bool)`

GetOverlayInterfaceNameOk returns a tuple with the OverlayInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayInterfaceName

`func (o *AddClusterRequestClusterServerConfigAnyOf) SetOverlayInterfaceName(v string)`

SetOverlayInterfaceName sets OverlayInterfaceName field to given value.

### HasOverlayInterfaceName

`func (o *AddClusterRequestClusterServerConfigAnyOf) HasOverlayInterfaceName() bool`

HasOverlayInterfaceName returns a boolean if a field has been set.

### GetVmwareFolderId

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *AddClusterRequestClusterServerConfigAnyOf) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *AddClusterRequestClusterServerConfigAnyOf) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *AddClusterRequestClusterServerConfigAnyOf) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


