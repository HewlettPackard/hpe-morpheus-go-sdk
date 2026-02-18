# ClusterServerCreateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultRepoAccount** | Pointer to **NullableInt32** | Default Repo Account is the repository to be used when pulling images.  Default behavior is to be anonymous, which does have limits on allowed image pulls from public Docker Repos. | [optional] 
**ImageServer** | Pointer to **string** | Act as Image Server. Set to on to use the Default Repo Account to pull images. | [optional] 
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

### NewClusterServerCreateConfig

`func NewClusterServerCreateConfig() *ClusterServerCreateConfig`

NewClusterServerCreateConfig instantiates a new ClusterServerCreateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterServerCreateConfigWithDefaults

`func NewClusterServerCreateConfigWithDefaults() *ClusterServerCreateConfig`

NewClusterServerCreateConfigWithDefaults instantiates a new ClusterServerCreateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultRepoAccount

`func (o *ClusterServerCreateConfig) GetDefaultRepoAccount() int32`

GetDefaultRepoAccount returns the DefaultRepoAccount field if non-nil, zero value otherwise.

### GetDefaultRepoAccountOk

`func (o *ClusterServerCreateConfig) GetDefaultRepoAccountOk() (*int32, bool)`

GetDefaultRepoAccountOk returns a tuple with the DefaultRepoAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRepoAccount

`func (o *ClusterServerCreateConfig) SetDefaultRepoAccount(v int32)`

SetDefaultRepoAccount sets DefaultRepoAccount field to given value.

### HasDefaultRepoAccount

`func (o *ClusterServerCreateConfig) HasDefaultRepoAccount() bool`

HasDefaultRepoAccount returns a boolean if a field has been set.

### SetDefaultRepoAccountNil

`func (o *ClusterServerCreateConfig) SetDefaultRepoAccountNil(b bool)`

 SetDefaultRepoAccountNil sets the value for DefaultRepoAccount to be an explicit nil

### UnsetDefaultRepoAccount
`func (o *ClusterServerCreateConfig) UnsetDefaultRepoAccount()`

UnsetDefaultRepoAccount ensures that no value is present for DefaultRepoAccount, not even an explicit nil
### GetImageServer

`func (o *ClusterServerCreateConfig) GetImageServer() string`

GetImageServer returns the ImageServer field if non-nil, zero value otherwise.

### GetImageServerOk

`func (o *ClusterServerCreateConfig) GetImageServerOk() (*string, bool)`

GetImageServerOk returns a tuple with the ImageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServer

`func (o *ClusterServerCreateConfig) SetImageServer(v string)`

SetImageServer sets ImageServer field to given value.

### HasImageServer

`func (o *ClusterServerCreateConfig) HasImageServer() bool`

HasImageServer returns a boolean if a field has been set.

### GetNodeCount

`func (o *ClusterServerCreateConfig) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterServerCreateConfig) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterServerCreateConfig) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterServerCreateConfig) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetCreateUser

`func (o *ClusterServerCreateConfig) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ClusterServerCreateConfig) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ClusterServerCreateConfig) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ClusterServerCreateConfig) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetController

`func (o *ClusterServerCreateConfig) GetController() ClusterServerCreateConfigAnyOfOneOf2Controller`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *ClusterServerCreateConfig) GetControllerOk() (*ClusterServerCreateConfigAnyOfOneOf2Controller, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *ClusterServerCreateConfig) SetController(v ClusterServerCreateConfigAnyOfOneOf2Controller)`

SetController sets Controller field to given value.

### HasController

`func (o *ClusterServerCreateConfig) HasController() bool`

HasController returns a boolean if a field has been set.

### GetWorker

`func (o *ClusterServerCreateConfig) GetWorker() ClusterServerCreateConfigAnyOfOneOf3Worker`

GetWorker returns the Worker field if non-nil, zero value otherwise.

### GetWorkerOk

`func (o *ClusterServerCreateConfig) GetWorkerOk() (*ClusterServerCreateConfigAnyOfOneOf3Worker, bool)`

GetWorkerOk returns a tuple with the Worker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorker

`func (o *ClusterServerCreateConfig) SetWorker(v ClusterServerCreateConfigAnyOfOneOf3Worker)`

SetWorker sets Worker field to given value.

### HasWorker

`func (o *ClusterServerCreateConfig) HasWorker() bool`

HasWorker returns a boolean if a field has been set.

### GetPublicIpType

`func (o *ClusterServerCreateConfig) GetPublicIpType() string`

GetPublicIpType returns the PublicIpType field if non-nil, zero value otherwise.

### GetPublicIpTypeOk

`func (o *ClusterServerCreateConfig) GetPublicIpTypeOk() (*string, bool)`

GetPublicIpTypeOk returns a tuple with the PublicIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpType

`func (o *ClusterServerCreateConfig) SetPublicIpType(v string)`

SetPublicIpType sets PublicIpType field to given value.

### HasPublicIpType

`func (o *ClusterServerCreateConfig) HasPublicIpType() bool`

HasPublicIpType returns a boolean if a field has been set.

### GetGoogleZoneId

`func (o *ClusterServerCreateConfig) GetGoogleZoneId() int64`

GetGoogleZoneId returns the GoogleZoneId field if non-nil, zero value otherwise.

### GetGoogleZoneIdOk

`func (o *ClusterServerCreateConfig) GetGoogleZoneIdOk() (*int64, bool)`

GetGoogleZoneIdOk returns a tuple with the GoogleZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleZoneId

`func (o *ClusterServerCreateConfig) SetGoogleZoneId(v int64)`

SetGoogleZoneId sets GoogleZoneId field to given value.

### HasGoogleZoneId

`func (o *ClusterServerCreateConfig) HasGoogleZoneId() bool`

HasGoogleZoneId returns a boolean if a field has been set.

### GetChannel

`func (o *ClusterServerCreateConfig) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ClusterServerCreateConfig) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ClusterServerCreateConfig) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ClusterServerCreateConfig) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetControlPlaneVersion

`func (o *ClusterServerCreateConfig) GetControlPlaneVersion() string`

GetControlPlaneVersion returns the ControlPlaneVersion field if non-nil, zero value otherwise.

### GetControlPlaneVersionOk

`func (o *ClusterServerCreateConfig) GetControlPlaneVersionOk() (*string, bool)`

GetControlPlaneVersionOk returns a tuple with the ControlPlaneVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneVersion

`func (o *ClusterServerCreateConfig) SetControlPlaneVersion(v string)`

SetControlPlaneVersion sets ControlPlaneVersion field to given value.

### HasControlPlaneVersion

`func (o *ClusterServerCreateConfig) HasControlPlaneVersion() bool`

HasControlPlaneVersion returns a boolean if a field has been set.

### GetPodCidr

`func (o *ClusterServerCreateConfig) GetPodCidr() string`

GetPodCidr returns the PodCidr field if non-nil, zero value otherwise.

### GetPodCidrOk

`func (o *ClusterServerCreateConfig) GetPodCidrOk() (*string, bool)`

GetPodCidrOk returns a tuple with the PodCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCidr

`func (o *ClusterServerCreateConfig) SetPodCidr(v string)`

SetPodCidr sets PodCidr field to given value.

### HasPodCidr

`func (o *ClusterServerCreateConfig) HasPodCidr() bool`

HasPodCidr returns a boolean if a field has been set.

### GetServiceCidr

`func (o *ClusterServerCreateConfig) GetServiceCidr() string`

GetServiceCidr returns the ServiceCidr field if non-nil, zero value otherwise.

### GetServiceCidrOk

`func (o *ClusterServerCreateConfig) GetServiceCidrOk() (*string, bool)`

GetServiceCidrOk returns a tuple with the ServiceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCidr

`func (o *ClusterServerCreateConfig) SetServiceCidr(v string)`

SetServiceCidr sets ServiceCidr field to given value.

### HasServiceCidr

`func (o *ClusterServerCreateConfig) HasServiceCidr() bool`

HasServiceCidr returns a boolean if a field has been set.

### GetCpuArch

`func (o *ClusterServerCreateConfig) GetCpuArch() string`

GetCpuArch returns the CpuArch field if non-nil, zero value otherwise.

### GetCpuArchOk

`func (o *ClusterServerCreateConfig) GetCpuArchOk() (*string, bool)`

GetCpuArchOk returns a tuple with the CpuArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuArch

`func (o *ClusterServerCreateConfig) SetCpuArch(v string)`

SetCpuArch sets CpuArch field to given value.

### HasCpuArch

`func (o *ClusterServerCreateConfig) HasCpuArch() bool`

HasCpuArch returns a boolean if a field has been set.

### GetCpuModel

`func (o *ClusterServerCreateConfig) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *ClusterServerCreateConfig) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *ClusterServerCreateConfig) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *ClusterServerCreateConfig) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetDynamicPlacementMode

`func (o *ClusterServerCreateConfig) GetDynamicPlacementMode() string`

GetDynamicPlacementMode returns the DynamicPlacementMode field if non-nil, zero value otherwise.

### GetDynamicPlacementModeOk

`func (o *ClusterServerCreateConfig) GetDynamicPlacementModeOk() (*string, bool)`

GetDynamicPlacementModeOk returns a tuple with the DynamicPlacementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicPlacementMode

`func (o *ClusterServerCreateConfig) SetDynamicPlacementMode(v string)`

SetDynamicPlacementMode sets DynamicPlacementMode field to given value.

### HasDynamicPlacementMode

`func (o *ClusterServerCreateConfig) HasDynamicPlacementMode() bool`

HasDynamicPlacementMode returns a boolean if a field has been set.

### GetPowerPolicy

`func (o *ClusterServerCreateConfig) GetPowerPolicy() string`

GetPowerPolicy returns the PowerPolicy field if non-nil, zero value otherwise.

### GetPowerPolicyOk

`func (o *ClusterServerCreateConfig) GetPowerPolicyOk() (*string, bool)`

GetPowerPolicyOk returns a tuple with the PowerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPolicy

`func (o *ClusterServerCreateConfig) SetPowerPolicy(v string)`

SetPowerPolicy sets PowerPolicy field to given value.

### HasPowerPolicy

`func (o *ClusterServerCreateConfig) HasPowerPolicy() bool`

HasPowerPolicy returns a boolean if a field has been set.

### GetStorageInterfaceName

`func (o *ClusterServerCreateConfig) GetStorageInterfaceName() string`

GetStorageInterfaceName returns the StorageInterfaceName field if non-nil, zero value otherwise.

### GetStorageInterfaceNameOk

`func (o *ClusterServerCreateConfig) GetStorageInterfaceNameOk() (*string, bool)`

GetStorageInterfaceNameOk returns a tuple with the StorageInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInterfaceName

`func (o *ClusterServerCreateConfig) SetStorageInterfaceName(v string)`

SetStorageInterfaceName sets StorageInterfaceName field to given value.

### HasStorageInterfaceName

`func (o *ClusterServerCreateConfig) HasStorageInterfaceName() bool`

HasStorageInterfaceName returns a boolean if a field has been set.

### GetComputeInterfaceName

`func (o *ClusterServerCreateConfig) GetComputeInterfaceName() string`

GetComputeInterfaceName returns the ComputeInterfaceName field if non-nil, zero value otherwise.

### GetComputeInterfaceNameOk

`func (o *ClusterServerCreateConfig) GetComputeInterfaceNameOk() (*string, bool)`

GetComputeInterfaceNameOk returns a tuple with the ComputeInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeInterfaceName

`func (o *ClusterServerCreateConfig) SetComputeInterfaceName(v string)`

SetComputeInterfaceName sets ComputeInterfaceName field to given value.

### HasComputeInterfaceName

`func (o *ClusterServerCreateConfig) HasComputeInterfaceName() bool`

HasComputeInterfaceName returns a boolean if a field has been set.

### GetComputeVlans

`func (o *ClusterServerCreateConfig) GetComputeVlans() string`

GetComputeVlans returns the ComputeVlans field if non-nil, zero value otherwise.

### GetComputeVlansOk

`func (o *ClusterServerCreateConfig) GetComputeVlansOk() (*string, bool)`

GetComputeVlansOk returns a tuple with the ComputeVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeVlans

`func (o *ClusterServerCreateConfig) SetComputeVlans(v string)`

SetComputeVlans sets ComputeVlans field to given value.

### HasComputeVlans

`func (o *ClusterServerCreateConfig) HasComputeVlans() bool`

HasComputeVlans returns a boolean if a field has been set.

### GetOverlayInterfaceName

`func (o *ClusterServerCreateConfig) GetOverlayInterfaceName() string`

GetOverlayInterfaceName returns the OverlayInterfaceName field if non-nil, zero value otherwise.

### GetOverlayInterfaceNameOk

`func (o *ClusterServerCreateConfig) GetOverlayInterfaceNameOk() (*string, bool)`

GetOverlayInterfaceNameOk returns a tuple with the OverlayInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayInterfaceName

`func (o *ClusterServerCreateConfig) SetOverlayInterfaceName(v string)`

SetOverlayInterfaceName sets OverlayInterfaceName field to given value.

### HasOverlayInterfaceName

`func (o *ClusterServerCreateConfig) HasOverlayInterfaceName() bool`

HasOverlayInterfaceName returns a boolean if a field has been set.

### GetVmwareFolderId

`func (o *ClusterServerCreateConfig) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *ClusterServerCreateConfig) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *ClusterServerCreateConfig) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *ClusterServerCreateConfig) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


