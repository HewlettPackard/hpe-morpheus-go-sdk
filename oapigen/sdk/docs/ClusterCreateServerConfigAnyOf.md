# ClusterCreateServerConfigAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeCount** | Pointer to **int64** | Number of workers or hosts | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 
**Controller** | Pointer to [**ClusterCreateServerConfigAnyOfOneOf2Controller**](ClusterCreateServerConfigAnyOfOneOf2Controller.md) |  | [optional] 
**Worker** | Pointer to [**ClusterCreateServerConfigAnyOfOneOf3Worker**](ClusterCreateServerConfigAnyOfOneOf3Worker.md) |  | [optional] 
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

### NewClusterCreateServerConfigAnyOf

`func NewClusterCreateServerConfigAnyOf() *ClusterCreateServerConfigAnyOf`

NewClusterCreateServerConfigAnyOf instantiates a new ClusterCreateServerConfigAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreateServerConfigAnyOfWithDefaults

`func NewClusterCreateServerConfigAnyOfWithDefaults() *ClusterCreateServerConfigAnyOf`

NewClusterCreateServerConfigAnyOfWithDefaults instantiates a new ClusterCreateServerConfigAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeCount

`func (o *ClusterCreateServerConfigAnyOf) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterCreateServerConfigAnyOf) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterCreateServerConfigAnyOf) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterCreateServerConfigAnyOf) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetCreateUser

`func (o *ClusterCreateServerConfigAnyOf) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ClusterCreateServerConfigAnyOf) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ClusterCreateServerConfigAnyOf) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ClusterCreateServerConfigAnyOf) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetController

`func (o *ClusterCreateServerConfigAnyOf) GetController() ClusterCreateServerConfigAnyOfOneOf2Controller`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *ClusterCreateServerConfigAnyOf) GetControllerOk() (*ClusterCreateServerConfigAnyOfOneOf2Controller, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *ClusterCreateServerConfigAnyOf) SetController(v ClusterCreateServerConfigAnyOfOneOf2Controller)`

SetController sets Controller field to given value.

### HasController

`func (o *ClusterCreateServerConfigAnyOf) HasController() bool`

HasController returns a boolean if a field has been set.

### GetWorker

`func (o *ClusterCreateServerConfigAnyOf) GetWorker() ClusterCreateServerConfigAnyOfOneOf3Worker`

GetWorker returns the Worker field if non-nil, zero value otherwise.

### GetWorkerOk

`func (o *ClusterCreateServerConfigAnyOf) GetWorkerOk() (*ClusterCreateServerConfigAnyOfOneOf3Worker, bool)`

GetWorkerOk returns a tuple with the Worker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorker

`func (o *ClusterCreateServerConfigAnyOf) SetWorker(v ClusterCreateServerConfigAnyOfOneOf3Worker)`

SetWorker sets Worker field to given value.

### HasWorker

`func (o *ClusterCreateServerConfigAnyOf) HasWorker() bool`

HasWorker returns a boolean if a field has been set.

### GetPublicIpType

`func (o *ClusterCreateServerConfigAnyOf) GetPublicIpType() string`

GetPublicIpType returns the PublicIpType field if non-nil, zero value otherwise.

### GetPublicIpTypeOk

`func (o *ClusterCreateServerConfigAnyOf) GetPublicIpTypeOk() (*string, bool)`

GetPublicIpTypeOk returns a tuple with the PublicIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpType

`func (o *ClusterCreateServerConfigAnyOf) SetPublicIpType(v string)`

SetPublicIpType sets PublicIpType field to given value.

### HasPublicIpType

`func (o *ClusterCreateServerConfigAnyOf) HasPublicIpType() bool`

HasPublicIpType returns a boolean if a field has been set.

### GetGoogleZoneId

`func (o *ClusterCreateServerConfigAnyOf) GetGoogleZoneId() int64`

GetGoogleZoneId returns the GoogleZoneId field if non-nil, zero value otherwise.

### GetGoogleZoneIdOk

`func (o *ClusterCreateServerConfigAnyOf) GetGoogleZoneIdOk() (*int64, bool)`

GetGoogleZoneIdOk returns a tuple with the GoogleZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleZoneId

`func (o *ClusterCreateServerConfigAnyOf) SetGoogleZoneId(v int64)`

SetGoogleZoneId sets GoogleZoneId field to given value.

### HasGoogleZoneId

`func (o *ClusterCreateServerConfigAnyOf) HasGoogleZoneId() bool`

HasGoogleZoneId returns a boolean if a field has been set.

### GetChannel

`func (o *ClusterCreateServerConfigAnyOf) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ClusterCreateServerConfigAnyOf) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ClusterCreateServerConfigAnyOf) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ClusterCreateServerConfigAnyOf) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetControlPlaneVersion

`func (o *ClusterCreateServerConfigAnyOf) GetControlPlaneVersion() string`

GetControlPlaneVersion returns the ControlPlaneVersion field if non-nil, zero value otherwise.

### GetControlPlaneVersionOk

`func (o *ClusterCreateServerConfigAnyOf) GetControlPlaneVersionOk() (*string, bool)`

GetControlPlaneVersionOk returns a tuple with the ControlPlaneVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneVersion

`func (o *ClusterCreateServerConfigAnyOf) SetControlPlaneVersion(v string)`

SetControlPlaneVersion sets ControlPlaneVersion field to given value.

### HasControlPlaneVersion

`func (o *ClusterCreateServerConfigAnyOf) HasControlPlaneVersion() bool`

HasControlPlaneVersion returns a boolean if a field has been set.

### GetPodCidr

`func (o *ClusterCreateServerConfigAnyOf) GetPodCidr() string`

GetPodCidr returns the PodCidr field if non-nil, zero value otherwise.

### GetPodCidrOk

`func (o *ClusterCreateServerConfigAnyOf) GetPodCidrOk() (*string, bool)`

GetPodCidrOk returns a tuple with the PodCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCidr

`func (o *ClusterCreateServerConfigAnyOf) SetPodCidr(v string)`

SetPodCidr sets PodCidr field to given value.

### HasPodCidr

`func (o *ClusterCreateServerConfigAnyOf) HasPodCidr() bool`

HasPodCidr returns a boolean if a field has been set.

### GetServiceCidr

`func (o *ClusterCreateServerConfigAnyOf) GetServiceCidr() string`

GetServiceCidr returns the ServiceCidr field if non-nil, zero value otherwise.

### GetServiceCidrOk

`func (o *ClusterCreateServerConfigAnyOf) GetServiceCidrOk() (*string, bool)`

GetServiceCidrOk returns a tuple with the ServiceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCidr

`func (o *ClusterCreateServerConfigAnyOf) SetServiceCidr(v string)`

SetServiceCidr sets ServiceCidr field to given value.

### HasServiceCidr

`func (o *ClusterCreateServerConfigAnyOf) HasServiceCidr() bool`

HasServiceCidr returns a boolean if a field has been set.

### GetCpuArch

`func (o *ClusterCreateServerConfigAnyOf) GetCpuArch() string`

GetCpuArch returns the CpuArch field if non-nil, zero value otherwise.

### GetCpuArchOk

`func (o *ClusterCreateServerConfigAnyOf) GetCpuArchOk() (*string, bool)`

GetCpuArchOk returns a tuple with the CpuArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuArch

`func (o *ClusterCreateServerConfigAnyOf) SetCpuArch(v string)`

SetCpuArch sets CpuArch field to given value.

### HasCpuArch

`func (o *ClusterCreateServerConfigAnyOf) HasCpuArch() bool`

HasCpuArch returns a boolean if a field has been set.

### GetCpuModel

`func (o *ClusterCreateServerConfigAnyOf) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *ClusterCreateServerConfigAnyOf) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *ClusterCreateServerConfigAnyOf) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *ClusterCreateServerConfigAnyOf) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetDynamicPlacementMode

`func (o *ClusterCreateServerConfigAnyOf) GetDynamicPlacementMode() string`

GetDynamicPlacementMode returns the DynamicPlacementMode field if non-nil, zero value otherwise.

### GetDynamicPlacementModeOk

`func (o *ClusterCreateServerConfigAnyOf) GetDynamicPlacementModeOk() (*string, bool)`

GetDynamicPlacementModeOk returns a tuple with the DynamicPlacementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicPlacementMode

`func (o *ClusterCreateServerConfigAnyOf) SetDynamicPlacementMode(v string)`

SetDynamicPlacementMode sets DynamicPlacementMode field to given value.

### HasDynamicPlacementMode

`func (o *ClusterCreateServerConfigAnyOf) HasDynamicPlacementMode() bool`

HasDynamicPlacementMode returns a boolean if a field has been set.

### GetPowerPolicy

`func (o *ClusterCreateServerConfigAnyOf) GetPowerPolicy() string`

GetPowerPolicy returns the PowerPolicy field if non-nil, zero value otherwise.

### GetPowerPolicyOk

`func (o *ClusterCreateServerConfigAnyOf) GetPowerPolicyOk() (*string, bool)`

GetPowerPolicyOk returns a tuple with the PowerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPolicy

`func (o *ClusterCreateServerConfigAnyOf) SetPowerPolicy(v string)`

SetPowerPolicy sets PowerPolicy field to given value.

### HasPowerPolicy

`func (o *ClusterCreateServerConfigAnyOf) HasPowerPolicy() bool`

HasPowerPolicy returns a boolean if a field has been set.

### GetStorageInterfaceName

`func (o *ClusterCreateServerConfigAnyOf) GetStorageInterfaceName() string`

GetStorageInterfaceName returns the StorageInterfaceName field if non-nil, zero value otherwise.

### GetStorageInterfaceNameOk

`func (o *ClusterCreateServerConfigAnyOf) GetStorageInterfaceNameOk() (*string, bool)`

GetStorageInterfaceNameOk returns a tuple with the StorageInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInterfaceName

`func (o *ClusterCreateServerConfigAnyOf) SetStorageInterfaceName(v string)`

SetStorageInterfaceName sets StorageInterfaceName field to given value.

### HasStorageInterfaceName

`func (o *ClusterCreateServerConfigAnyOf) HasStorageInterfaceName() bool`

HasStorageInterfaceName returns a boolean if a field has been set.

### GetComputeInterfaceName

`func (o *ClusterCreateServerConfigAnyOf) GetComputeInterfaceName() string`

GetComputeInterfaceName returns the ComputeInterfaceName field if non-nil, zero value otherwise.

### GetComputeInterfaceNameOk

`func (o *ClusterCreateServerConfigAnyOf) GetComputeInterfaceNameOk() (*string, bool)`

GetComputeInterfaceNameOk returns a tuple with the ComputeInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeInterfaceName

`func (o *ClusterCreateServerConfigAnyOf) SetComputeInterfaceName(v string)`

SetComputeInterfaceName sets ComputeInterfaceName field to given value.

### HasComputeInterfaceName

`func (o *ClusterCreateServerConfigAnyOf) HasComputeInterfaceName() bool`

HasComputeInterfaceName returns a boolean if a field has been set.

### GetComputeVlans

`func (o *ClusterCreateServerConfigAnyOf) GetComputeVlans() string`

GetComputeVlans returns the ComputeVlans field if non-nil, zero value otherwise.

### GetComputeVlansOk

`func (o *ClusterCreateServerConfigAnyOf) GetComputeVlansOk() (*string, bool)`

GetComputeVlansOk returns a tuple with the ComputeVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeVlans

`func (o *ClusterCreateServerConfigAnyOf) SetComputeVlans(v string)`

SetComputeVlans sets ComputeVlans field to given value.

### HasComputeVlans

`func (o *ClusterCreateServerConfigAnyOf) HasComputeVlans() bool`

HasComputeVlans returns a boolean if a field has been set.

### GetOverlayInterfaceName

`func (o *ClusterCreateServerConfigAnyOf) GetOverlayInterfaceName() string`

GetOverlayInterfaceName returns the OverlayInterfaceName field if non-nil, zero value otherwise.

### GetOverlayInterfaceNameOk

`func (o *ClusterCreateServerConfigAnyOf) GetOverlayInterfaceNameOk() (*string, bool)`

GetOverlayInterfaceNameOk returns a tuple with the OverlayInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayInterfaceName

`func (o *ClusterCreateServerConfigAnyOf) SetOverlayInterfaceName(v string)`

SetOverlayInterfaceName sets OverlayInterfaceName field to given value.

### HasOverlayInterfaceName

`func (o *ClusterCreateServerConfigAnyOf) HasOverlayInterfaceName() bool`

HasOverlayInterfaceName returns a boolean if a field has been set.

### GetVmwareFolderId

`func (o *ClusterCreateServerConfigAnyOf) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *ClusterCreateServerConfigAnyOf) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *ClusterCreateServerConfigAnyOf) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *ClusterCreateServerConfigAnyOf) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


