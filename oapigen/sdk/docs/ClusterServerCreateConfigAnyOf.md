# ClusterServerCreateConfigAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeCount** | Pointer to **int64** | Number of workers or hosts | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 
**Controller** | Pointer to [**ClusterServerCreateConfigAnyOfOneOf2Controller**](ClusterServerCreateConfigAnyOfOneOf2Controller.md) |  | [optional] 
**Worker** | Pointer to [**ClusterServerCreateConfigAnyOfOneOf3Worker**](ClusterServerCreateConfigAnyOfOneOf3Worker.md) |  | [optional] 
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

### NewClusterServerCreateConfigAnyOf

`func NewClusterServerCreateConfigAnyOf() *ClusterServerCreateConfigAnyOf`

NewClusterServerCreateConfigAnyOf instantiates a new ClusterServerCreateConfigAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterServerCreateConfigAnyOfWithDefaults

`func NewClusterServerCreateConfigAnyOfWithDefaults() *ClusterServerCreateConfigAnyOf`

NewClusterServerCreateConfigAnyOfWithDefaults instantiates a new ClusterServerCreateConfigAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeCount

`func (o *ClusterServerCreateConfigAnyOf) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterServerCreateConfigAnyOf) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterServerCreateConfigAnyOf) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterServerCreateConfigAnyOf) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetCreateUser

`func (o *ClusterServerCreateConfigAnyOf) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ClusterServerCreateConfigAnyOf) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ClusterServerCreateConfigAnyOf) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ClusterServerCreateConfigAnyOf) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetController

`func (o *ClusterServerCreateConfigAnyOf) GetController() ClusterServerCreateConfigAnyOfOneOf2Controller`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *ClusterServerCreateConfigAnyOf) GetControllerOk() (*ClusterServerCreateConfigAnyOfOneOf2Controller, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *ClusterServerCreateConfigAnyOf) SetController(v ClusterServerCreateConfigAnyOfOneOf2Controller)`

SetController sets Controller field to given value.

### HasController

`func (o *ClusterServerCreateConfigAnyOf) HasController() bool`

HasController returns a boolean if a field has been set.

### GetWorker

`func (o *ClusterServerCreateConfigAnyOf) GetWorker() ClusterServerCreateConfigAnyOfOneOf3Worker`

GetWorker returns the Worker field if non-nil, zero value otherwise.

### GetWorkerOk

`func (o *ClusterServerCreateConfigAnyOf) GetWorkerOk() (*ClusterServerCreateConfigAnyOfOneOf3Worker, bool)`

GetWorkerOk returns a tuple with the Worker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorker

`func (o *ClusterServerCreateConfigAnyOf) SetWorker(v ClusterServerCreateConfigAnyOfOneOf3Worker)`

SetWorker sets Worker field to given value.

### HasWorker

`func (o *ClusterServerCreateConfigAnyOf) HasWorker() bool`

HasWorker returns a boolean if a field has been set.

### GetPublicIpType

`func (o *ClusterServerCreateConfigAnyOf) GetPublicIpType() string`

GetPublicIpType returns the PublicIpType field if non-nil, zero value otherwise.

### GetPublicIpTypeOk

`func (o *ClusterServerCreateConfigAnyOf) GetPublicIpTypeOk() (*string, bool)`

GetPublicIpTypeOk returns a tuple with the PublicIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpType

`func (o *ClusterServerCreateConfigAnyOf) SetPublicIpType(v string)`

SetPublicIpType sets PublicIpType field to given value.

### HasPublicIpType

`func (o *ClusterServerCreateConfigAnyOf) HasPublicIpType() bool`

HasPublicIpType returns a boolean if a field has been set.

### GetGoogleZoneId

`func (o *ClusterServerCreateConfigAnyOf) GetGoogleZoneId() int64`

GetGoogleZoneId returns the GoogleZoneId field if non-nil, zero value otherwise.

### GetGoogleZoneIdOk

`func (o *ClusterServerCreateConfigAnyOf) GetGoogleZoneIdOk() (*int64, bool)`

GetGoogleZoneIdOk returns a tuple with the GoogleZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleZoneId

`func (o *ClusterServerCreateConfigAnyOf) SetGoogleZoneId(v int64)`

SetGoogleZoneId sets GoogleZoneId field to given value.

### HasGoogleZoneId

`func (o *ClusterServerCreateConfigAnyOf) HasGoogleZoneId() bool`

HasGoogleZoneId returns a boolean if a field has been set.

### GetChannel

`func (o *ClusterServerCreateConfigAnyOf) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ClusterServerCreateConfigAnyOf) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ClusterServerCreateConfigAnyOf) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ClusterServerCreateConfigAnyOf) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetControlPlaneVersion

`func (o *ClusterServerCreateConfigAnyOf) GetControlPlaneVersion() string`

GetControlPlaneVersion returns the ControlPlaneVersion field if non-nil, zero value otherwise.

### GetControlPlaneVersionOk

`func (o *ClusterServerCreateConfigAnyOf) GetControlPlaneVersionOk() (*string, bool)`

GetControlPlaneVersionOk returns a tuple with the ControlPlaneVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneVersion

`func (o *ClusterServerCreateConfigAnyOf) SetControlPlaneVersion(v string)`

SetControlPlaneVersion sets ControlPlaneVersion field to given value.

### HasControlPlaneVersion

`func (o *ClusterServerCreateConfigAnyOf) HasControlPlaneVersion() bool`

HasControlPlaneVersion returns a boolean if a field has been set.

### GetPodCidr

`func (o *ClusterServerCreateConfigAnyOf) GetPodCidr() string`

GetPodCidr returns the PodCidr field if non-nil, zero value otherwise.

### GetPodCidrOk

`func (o *ClusterServerCreateConfigAnyOf) GetPodCidrOk() (*string, bool)`

GetPodCidrOk returns a tuple with the PodCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCidr

`func (o *ClusterServerCreateConfigAnyOf) SetPodCidr(v string)`

SetPodCidr sets PodCidr field to given value.

### HasPodCidr

`func (o *ClusterServerCreateConfigAnyOf) HasPodCidr() bool`

HasPodCidr returns a boolean if a field has been set.

### GetServiceCidr

`func (o *ClusterServerCreateConfigAnyOf) GetServiceCidr() string`

GetServiceCidr returns the ServiceCidr field if non-nil, zero value otherwise.

### GetServiceCidrOk

`func (o *ClusterServerCreateConfigAnyOf) GetServiceCidrOk() (*string, bool)`

GetServiceCidrOk returns a tuple with the ServiceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCidr

`func (o *ClusterServerCreateConfigAnyOf) SetServiceCidr(v string)`

SetServiceCidr sets ServiceCidr field to given value.

### HasServiceCidr

`func (o *ClusterServerCreateConfigAnyOf) HasServiceCidr() bool`

HasServiceCidr returns a boolean if a field has been set.

### GetCpuArch

`func (o *ClusterServerCreateConfigAnyOf) GetCpuArch() string`

GetCpuArch returns the CpuArch field if non-nil, zero value otherwise.

### GetCpuArchOk

`func (o *ClusterServerCreateConfigAnyOf) GetCpuArchOk() (*string, bool)`

GetCpuArchOk returns a tuple with the CpuArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuArch

`func (o *ClusterServerCreateConfigAnyOf) SetCpuArch(v string)`

SetCpuArch sets CpuArch field to given value.

### HasCpuArch

`func (o *ClusterServerCreateConfigAnyOf) HasCpuArch() bool`

HasCpuArch returns a boolean if a field has been set.

### GetCpuModel

`func (o *ClusterServerCreateConfigAnyOf) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *ClusterServerCreateConfigAnyOf) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *ClusterServerCreateConfigAnyOf) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *ClusterServerCreateConfigAnyOf) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetDynamicPlacementMode

`func (o *ClusterServerCreateConfigAnyOf) GetDynamicPlacementMode() string`

GetDynamicPlacementMode returns the DynamicPlacementMode field if non-nil, zero value otherwise.

### GetDynamicPlacementModeOk

`func (o *ClusterServerCreateConfigAnyOf) GetDynamicPlacementModeOk() (*string, bool)`

GetDynamicPlacementModeOk returns a tuple with the DynamicPlacementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicPlacementMode

`func (o *ClusterServerCreateConfigAnyOf) SetDynamicPlacementMode(v string)`

SetDynamicPlacementMode sets DynamicPlacementMode field to given value.

### HasDynamicPlacementMode

`func (o *ClusterServerCreateConfigAnyOf) HasDynamicPlacementMode() bool`

HasDynamicPlacementMode returns a boolean if a field has been set.

### GetPowerPolicy

`func (o *ClusterServerCreateConfigAnyOf) GetPowerPolicy() string`

GetPowerPolicy returns the PowerPolicy field if non-nil, zero value otherwise.

### GetPowerPolicyOk

`func (o *ClusterServerCreateConfigAnyOf) GetPowerPolicyOk() (*string, bool)`

GetPowerPolicyOk returns a tuple with the PowerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPolicy

`func (o *ClusterServerCreateConfigAnyOf) SetPowerPolicy(v string)`

SetPowerPolicy sets PowerPolicy field to given value.

### HasPowerPolicy

`func (o *ClusterServerCreateConfigAnyOf) HasPowerPolicy() bool`

HasPowerPolicy returns a boolean if a field has been set.

### GetStorageInterfaceName

`func (o *ClusterServerCreateConfigAnyOf) GetStorageInterfaceName() string`

GetStorageInterfaceName returns the StorageInterfaceName field if non-nil, zero value otherwise.

### GetStorageInterfaceNameOk

`func (o *ClusterServerCreateConfigAnyOf) GetStorageInterfaceNameOk() (*string, bool)`

GetStorageInterfaceNameOk returns a tuple with the StorageInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInterfaceName

`func (o *ClusterServerCreateConfigAnyOf) SetStorageInterfaceName(v string)`

SetStorageInterfaceName sets StorageInterfaceName field to given value.

### HasStorageInterfaceName

`func (o *ClusterServerCreateConfigAnyOf) HasStorageInterfaceName() bool`

HasStorageInterfaceName returns a boolean if a field has been set.

### GetComputeInterfaceName

`func (o *ClusterServerCreateConfigAnyOf) GetComputeInterfaceName() string`

GetComputeInterfaceName returns the ComputeInterfaceName field if non-nil, zero value otherwise.

### GetComputeInterfaceNameOk

`func (o *ClusterServerCreateConfigAnyOf) GetComputeInterfaceNameOk() (*string, bool)`

GetComputeInterfaceNameOk returns a tuple with the ComputeInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeInterfaceName

`func (o *ClusterServerCreateConfigAnyOf) SetComputeInterfaceName(v string)`

SetComputeInterfaceName sets ComputeInterfaceName field to given value.

### HasComputeInterfaceName

`func (o *ClusterServerCreateConfigAnyOf) HasComputeInterfaceName() bool`

HasComputeInterfaceName returns a boolean if a field has been set.

### GetComputeVlans

`func (o *ClusterServerCreateConfigAnyOf) GetComputeVlans() string`

GetComputeVlans returns the ComputeVlans field if non-nil, zero value otherwise.

### GetComputeVlansOk

`func (o *ClusterServerCreateConfigAnyOf) GetComputeVlansOk() (*string, bool)`

GetComputeVlansOk returns a tuple with the ComputeVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeVlans

`func (o *ClusterServerCreateConfigAnyOf) SetComputeVlans(v string)`

SetComputeVlans sets ComputeVlans field to given value.

### HasComputeVlans

`func (o *ClusterServerCreateConfigAnyOf) HasComputeVlans() bool`

HasComputeVlans returns a boolean if a field has been set.

### GetOverlayInterfaceName

`func (o *ClusterServerCreateConfigAnyOf) GetOverlayInterfaceName() string`

GetOverlayInterfaceName returns the OverlayInterfaceName field if non-nil, zero value otherwise.

### GetOverlayInterfaceNameOk

`func (o *ClusterServerCreateConfigAnyOf) GetOverlayInterfaceNameOk() (*string, bool)`

GetOverlayInterfaceNameOk returns a tuple with the OverlayInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayInterfaceName

`func (o *ClusterServerCreateConfigAnyOf) SetOverlayInterfaceName(v string)`

SetOverlayInterfaceName sets OverlayInterfaceName field to given value.

### HasOverlayInterfaceName

`func (o *ClusterServerCreateConfigAnyOf) HasOverlayInterfaceName() bool`

HasOverlayInterfaceName returns a boolean if a field has been set.

### GetVmwareFolderId

`func (o *ClusterServerCreateConfigAnyOf) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *ClusterServerCreateConfigAnyOf) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *ClusterServerCreateConfigAnyOf) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *ClusterServerCreateConfigAnyOf) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


