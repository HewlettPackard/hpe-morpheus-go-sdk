# AddClusterRequestClusterServerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultRepoAccount** | Pointer to **NullableInt32** | Default Repo Account is the repository to be used when pulling images.  Default behavior is to be anonymous, which does have limits on allowed image pulls from public Docker Repos. | [optional] 
**ImageServer** | Pointer to **string** | Act as Image Server. Set to on to use the Default Repo Account to pull images. | [optional] 
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

### NewAddClusterRequestClusterServerConfig

`func NewAddClusterRequestClusterServerConfig() *AddClusterRequestClusterServerConfig`

NewAddClusterRequestClusterServerConfig instantiates a new AddClusterRequestClusterServerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClusterRequestClusterServerConfigWithDefaults

`func NewAddClusterRequestClusterServerConfigWithDefaults() *AddClusterRequestClusterServerConfig`

NewAddClusterRequestClusterServerConfigWithDefaults instantiates a new AddClusterRequestClusterServerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultRepoAccount

`func (o *AddClusterRequestClusterServerConfig) GetDefaultRepoAccount() int32`

GetDefaultRepoAccount returns the DefaultRepoAccount field if non-nil, zero value otherwise.

### GetDefaultRepoAccountOk

`func (o *AddClusterRequestClusterServerConfig) GetDefaultRepoAccountOk() (*int32, bool)`

GetDefaultRepoAccountOk returns a tuple with the DefaultRepoAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRepoAccount

`func (o *AddClusterRequestClusterServerConfig) SetDefaultRepoAccount(v int32)`

SetDefaultRepoAccount sets DefaultRepoAccount field to given value.

### HasDefaultRepoAccount

`func (o *AddClusterRequestClusterServerConfig) HasDefaultRepoAccount() bool`

HasDefaultRepoAccount returns a boolean if a field has been set.

### SetDefaultRepoAccountNil

`func (o *AddClusterRequestClusterServerConfig) SetDefaultRepoAccountNil(b bool)`

 SetDefaultRepoAccountNil sets the value for DefaultRepoAccount to be an explicit nil

### UnsetDefaultRepoAccount
`func (o *AddClusterRequestClusterServerConfig) UnsetDefaultRepoAccount()`

UnsetDefaultRepoAccount ensures that no value is present for DefaultRepoAccount, not even an explicit nil
### GetImageServer

`func (o *AddClusterRequestClusterServerConfig) GetImageServer() string`

GetImageServer returns the ImageServer field if non-nil, zero value otherwise.

### GetImageServerOk

`func (o *AddClusterRequestClusterServerConfig) GetImageServerOk() (*string, bool)`

GetImageServerOk returns a tuple with the ImageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServer

`func (o *AddClusterRequestClusterServerConfig) SetImageServer(v string)`

SetImageServer sets ImageServer field to given value.

### HasImageServer

`func (o *AddClusterRequestClusterServerConfig) HasImageServer() bool`

HasImageServer returns a boolean if a field has been set.

### GetNodeCount

`func (o *AddClusterRequestClusterServerConfig) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *AddClusterRequestClusterServerConfig) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *AddClusterRequestClusterServerConfig) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *AddClusterRequestClusterServerConfig) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetCreateUser

`func (o *AddClusterRequestClusterServerConfig) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *AddClusterRequestClusterServerConfig) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *AddClusterRequestClusterServerConfig) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *AddClusterRequestClusterServerConfig) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetController

`func (o *AddClusterRequestClusterServerConfig) GetController() AddClusterRequestClusterServerConfigAnyOfOneOf2Controller`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *AddClusterRequestClusterServerConfig) GetControllerOk() (*AddClusterRequestClusterServerConfigAnyOfOneOf2Controller, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *AddClusterRequestClusterServerConfig) SetController(v AddClusterRequestClusterServerConfigAnyOfOneOf2Controller)`

SetController sets Controller field to given value.

### HasController

`func (o *AddClusterRequestClusterServerConfig) HasController() bool`

HasController returns a boolean if a field has been set.

### GetWorker

`func (o *AddClusterRequestClusterServerConfig) GetWorker() AddClusterRequestClusterServerConfigAnyOfOneOf3Worker`

GetWorker returns the Worker field if non-nil, zero value otherwise.

### GetWorkerOk

`func (o *AddClusterRequestClusterServerConfig) GetWorkerOk() (*AddClusterRequestClusterServerConfigAnyOfOneOf3Worker, bool)`

GetWorkerOk returns a tuple with the Worker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorker

`func (o *AddClusterRequestClusterServerConfig) SetWorker(v AddClusterRequestClusterServerConfigAnyOfOneOf3Worker)`

SetWorker sets Worker field to given value.

### HasWorker

`func (o *AddClusterRequestClusterServerConfig) HasWorker() bool`

HasWorker returns a boolean if a field has been set.

### GetPublicIpType

`func (o *AddClusterRequestClusterServerConfig) GetPublicIpType() string`

GetPublicIpType returns the PublicIpType field if non-nil, zero value otherwise.

### GetPublicIpTypeOk

`func (o *AddClusterRequestClusterServerConfig) GetPublicIpTypeOk() (*string, bool)`

GetPublicIpTypeOk returns a tuple with the PublicIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpType

`func (o *AddClusterRequestClusterServerConfig) SetPublicIpType(v string)`

SetPublicIpType sets PublicIpType field to given value.

### HasPublicIpType

`func (o *AddClusterRequestClusterServerConfig) HasPublicIpType() bool`

HasPublicIpType returns a boolean if a field has been set.

### GetGoogleZoneId

`func (o *AddClusterRequestClusterServerConfig) GetGoogleZoneId() int64`

GetGoogleZoneId returns the GoogleZoneId field if non-nil, zero value otherwise.

### GetGoogleZoneIdOk

`func (o *AddClusterRequestClusterServerConfig) GetGoogleZoneIdOk() (*int64, bool)`

GetGoogleZoneIdOk returns a tuple with the GoogleZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleZoneId

`func (o *AddClusterRequestClusterServerConfig) SetGoogleZoneId(v int64)`

SetGoogleZoneId sets GoogleZoneId field to given value.

### HasGoogleZoneId

`func (o *AddClusterRequestClusterServerConfig) HasGoogleZoneId() bool`

HasGoogleZoneId returns a boolean if a field has been set.

### GetChannel

`func (o *AddClusterRequestClusterServerConfig) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *AddClusterRequestClusterServerConfig) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *AddClusterRequestClusterServerConfig) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *AddClusterRequestClusterServerConfig) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetControlPlaneVersion

`func (o *AddClusterRequestClusterServerConfig) GetControlPlaneVersion() string`

GetControlPlaneVersion returns the ControlPlaneVersion field if non-nil, zero value otherwise.

### GetControlPlaneVersionOk

`func (o *AddClusterRequestClusterServerConfig) GetControlPlaneVersionOk() (*string, bool)`

GetControlPlaneVersionOk returns a tuple with the ControlPlaneVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneVersion

`func (o *AddClusterRequestClusterServerConfig) SetControlPlaneVersion(v string)`

SetControlPlaneVersion sets ControlPlaneVersion field to given value.

### HasControlPlaneVersion

`func (o *AddClusterRequestClusterServerConfig) HasControlPlaneVersion() bool`

HasControlPlaneVersion returns a boolean if a field has been set.

### GetPodCidr

`func (o *AddClusterRequestClusterServerConfig) GetPodCidr() string`

GetPodCidr returns the PodCidr field if non-nil, zero value otherwise.

### GetPodCidrOk

`func (o *AddClusterRequestClusterServerConfig) GetPodCidrOk() (*string, bool)`

GetPodCidrOk returns a tuple with the PodCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCidr

`func (o *AddClusterRequestClusterServerConfig) SetPodCidr(v string)`

SetPodCidr sets PodCidr field to given value.

### HasPodCidr

`func (o *AddClusterRequestClusterServerConfig) HasPodCidr() bool`

HasPodCidr returns a boolean if a field has been set.

### GetServiceCidr

`func (o *AddClusterRequestClusterServerConfig) GetServiceCidr() string`

GetServiceCidr returns the ServiceCidr field if non-nil, zero value otherwise.

### GetServiceCidrOk

`func (o *AddClusterRequestClusterServerConfig) GetServiceCidrOk() (*string, bool)`

GetServiceCidrOk returns a tuple with the ServiceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCidr

`func (o *AddClusterRequestClusterServerConfig) SetServiceCidr(v string)`

SetServiceCidr sets ServiceCidr field to given value.

### HasServiceCidr

`func (o *AddClusterRequestClusterServerConfig) HasServiceCidr() bool`

HasServiceCidr returns a boolean if a field has been set.

### GetCpuArch

`func (o *AddClusterRequestClusterServerConfig) GetCpuArch() string`

GetCpuArch returns the CpuArch field if non-nil, zero value otherwise.

### GetCpuArchOk

`func (o *AddClusterRequestClusterServerConfig) GetCpuArchOk() (*string, bool)`

GetCpuArchOk returns a tuple with the CpuArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuArch

`func (o *AddClusterRequestClusterServerConfig) SetCpuArch(v string)`

SetCpuArch sets CpuArch field to given value.

### HasCpuArch

`func (o *AddClusterRequestClusterServerConfig) HasCpuArch() bool`

HasCpuArch returns a boolean if a field has been set.

### GetCpuModel

`func (o *AddClusterRequestClusterServerConfig) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *AddClusterRequestClusterServerConfig) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *AddClusterRequestClusterServerConfig) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *AddClusterRequestClusterServerConfig) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetDynamicPlacementMode

`func (o *AddClusterRequestClusterServerConfig) GetDynamicPlacementMode() string`

GetDynamicPlacementMode returns the DynamicPlacementMode field if non-nil, zero value otherwise.

### GetDynamicPlacementModeOk

`func (o *AddClusterRequestClusterServerConfig) GetDynamicPlacementModeOk() (*string, bool)`

GetDynamicPlacementModeOk returns a tuple with the DynamicPlacementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicPlacementMode

`func (o *AddClusterRequestClusterServerConfig) SetDynamicPlacementMode(v string)`

SetDynamicPlacementMode sets DynamicPlacementMode field to given value.

### HasDynamicPlacementMode

`func (o *AddClusterRequestClusterServerConfig) HasDynamicPlacementMode() bool`

HasDynamicPlacementMode returns a boolean if a field has been set.

### GetPowerPolicy

`func (o *AddClusterRequestClusterServerConfig) GetPowerPolicy() string`

GetPowerPolicy returns the PowerPolicy field if non-nil, zero value otherwise.

### GetPowerPolicyOk

`func (o *AddClusterRequestClusterServerConfig) GetPowerPolicyOk() (*string, bool)`

GetPowerPolicyOk returns a tuple with the PowerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPolicy

`func (o *AddClusterRequestClusterServerConfig) SetPowerPolicy(v string)`

SetPowerPolicy sets PowerPolicy field to given value.

### HasPowerPolicy

`func (o *AddClusterRequestClusterServerConfig) HasPowerPolicy() bool`

HasPowerPolicy returns a boolean if a field has been set.

### GetStorageInterfaceName

`func (o *AddClusterRequestClusterServerConfig) GetStorageInterfaceName() string`

GetStorageInterfaceName returns the StorageInterfaceName field if non-nil, zero value otherwise.

### GetStorageInterfaceNameOk

`func (o *AddClusterRequestClusterServerConfig) GetStorageInterfaceNameOk() (*string, bool)`

GetStorageInterfaceNameOk returns a tuple with the StorageInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInterfaceName

`func (o *AddClusterRequestClusterServerConfig) SetStorageInterfaceName(v string)`

SetStorageInterfaceName sets StorageInterfaceName field to given value.

### HasStorageInterfaceName

`func (o *AddClusterRequestClusterServerConfig) HasStorageInterfaceName() bool`

HasStorageInterfaceName returns a boolean if a field has been set.

### GetComputeInterfaceName

`func (o *AddClusterRequestClusterServerConfig) GetComputeInterfaceName() string`

GetComputeInterfaceName returns the ComputeInterfaceName field if non-nil, zero value otherwise.

### GetComputeInterfaceNameOk

`func (o *AddClusterRequestClusterServerConfig) GetComputeInterfaceNameOk() (*string, bool)`

GetComputeInterfaceNameOk returns a tuple with the ComputeInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeInterfaceName

`func (o *AddClusterRequestClusterServerConfig) SetComputeInterfaceName(v string)`

SetComputeInterfaceName sets ComputeInterfaceName field to given value.

### HasComputeInterfaceName

`func (o *AddClusterRequestClusterServerConfig) HasComputeInterfaceName() bool`

HasComputeInterfaceName returns a boolean if a field has been set.

### GetComputeVlans

`func (o *AddClusterRequestClusterServerConfig) GetComputeVlans() string`

GetComputeVlans returns the ComputeVlans field if non-nil, zero value otherwise.

### GetComputeVlansOk

`func (o *AddClusterRequestClusterServerConfig) GetComputeVlansOk() (*string, bool)`

GetComputeVlansOk returns a tuple with the ComputeVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeVlans

`func (o *AddClusterRequestClusterServerConfig) SetComputeVlans(v string)`

SetComputeVlans sets ComputeVlans field to given value.

### HasComputeVlans

`func (o *AddClusterRequestClusterServerConfig) HasComputeVlans() bool`

HasComputeVlans returns a boolean if a field has been set.

### GetOverlayInterfaceName

`func (o *AddClusterRequestClusterServerConfig) GetOverlayInterfaceName() string`

GetOverlayInterfaceName returns the OverlayInterfaceName field if non-nil, zero value otherwise.

### GetOverlayInterfaceNameOk

`func (o *AddClusterRequestClusterServerConfig) GetOverlayInterfaceNameOk() (*string, bool)`

GetOverlayInterfaceNameOk returns a tuple with the OverlayInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayInterfaceName

`func (o *AddClusterRequestClusterServerConfig) SetOverlayInterfaceName(v string)`

SetOverlayInterfaceName sets OverlayInterfaceName field to given value.

### HasOverlayInterfaceName

`func (o *AddClusterRequestClusterServerConfig) HasOverlayInterfaceName() bool`

HasOverlayInterfaceName returns a boolean if a field has been set.

### GetVmwareFolderId

`func (o *AddClusterRequestClusterServerConfig) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *AddClusterRequestClusterServerConfig) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *AddClusterRequestClusterServerConfig) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *AddClusterRequestClusterServerConfig) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


