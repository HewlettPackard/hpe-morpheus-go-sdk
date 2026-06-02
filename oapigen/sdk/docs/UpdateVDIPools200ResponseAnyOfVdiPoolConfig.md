# UpdateVDIPools200ResponseAnyOfVdiPoolConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**UpdateVDIPools200ResponseAnyOfVdiPoolConfigGroup**](UpdateVDIPools200ResponseAnyOfVdiPoolConfigGroup.md) |  | [optional] 
**Cloud** | Pointer to [**UpdateVDIPools200ResponseAnyOfVdiPoolConfigCloud**](UpdateVDIPools200ResponseAnyOfVdiPoolConfigCloud.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Instance** | Pointer to [**UpdateVDIPools200ResponseAnyOfVdiPoolConfigInstance**](UpdateVDIPools200ResponseAnyOfVdiPoolConfigInstance.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**UpdateVDIPools200ResponseAnyOfVdiPoolConfigConfig**](UpdateVDIPools200ResponseAnyOfVdiPoolConfigConfig.md) |  | [optional] 
**Volumes** | Pointer to [**[]UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner**](UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner.md) |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**Layout** | Pointer to [**UpdateVDIPools200ResponseAnyOfVdiPoolConfigLayout**](UpdateVDIPools200ResponseAnyOfVdiPoolConfigLayout.md) |  | [optional] 
**StorageControllers** | Pointer to [**[]UpdateVDIPools200ResponseAnyOfVdiPoolConfigStorageControllersInner**](UpdateVDIPools200ResponseAnyOfVdiPoolConfigStorageControllersInner.md) |  | [optional] 
**Plan** | Pointer to [**UpdateVDIPools200ResponseAnyOfVdiPoolConfigPlan**](UpdateVDIPools200ResponseAnyOfVdiPoolConfigPlan.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**NetworkInterfaces** | Pointer to [**[]UpdateVDIPools200ResponseAnyOfVdiPoolConfigNetworkInterfacesInner**](UpdateVDIPools200ResponseAnyOfVdiPoolConfigNetworkInterfacesInner.md) |  | [optional] 
**ExecutionId** | Pointer to **string** |  | [optional] 
**Backup** | Pointer to [**UpdateVDIPools200ResponseAnyOfVdiPoolConfigBackup**](UpdateVDIPools200ResponseAnyOfVdiPoolConfigBackup.md) |  | [optional] 
**LoadBalancer** | Pointer to **[]map[string]interface{}** |  | [optional] 
**HideLock** | Pointer to **bool** |  | [optional] 
**HasNetworks** | Pointer to **bool** |  | [optional] 
**DisplayNetworks** | Pointer to [**[]UpdateVDIPools200ResponseAnyOfVdiPoolConfigDisplayNetworksInner**](UpdateVDIPools200ResponseAnyOfVdiPoolConfigDisplayNetworksInner.md) |  | [optional] 
**Copies** | Pointer to **int64** |  | [optional] 
**ShowScale** | Pointer to **bool** |  | [optional] 
**HasPreview** | Pointer to **bool** |  | [optional] 
**VolumesDisplay** | Pointer to [**[]UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesDisplayInner**](UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesDisplayInner.md) |  | [optional] 

## Methods

### NewUpdateVDIPools200ResponseAnyOfVdiPoolConfig

`func NewUpdateVDIPools200ResponseAnyOfVdiPoolConfig() *UpdateVDIPools200ResponseAnyOfVdiPoolConfig`

NewUpdateVDIPools200ResponseAnyOfVdiPoolConfig instantiates a new UpdateVDIPools200ResponseAnyOfVdiPoolConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetGroup

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetGroup() UpdateVDIPools200ResponseAnyOfVdiPoolConfigGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetGroupOk() (*UpdateVDIPools200ResponseAnyOfVdiPoolConfigGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetGroup(v UpdateVDIPools200ResponseAnyOfVdiPoolConfigGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCloud

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetCloud() UpdateVDIPools200ResponseAnyOfVdiPoolConfigCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetCloudOk() (*UpdateVDIPools200ResponseAnyOfVdiPoolConfigCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetCloud(v UpdateVDIPools200ResponseAnyOfVdiPoolConfigCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetType

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInstance

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetInstance() UpdateVDIPools200ResponseAnyOfVdiPoolConfigInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetInstanceOk() (*UpdateVDIPools200ResponseAnyOfVdiPoolConfigInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetInstance(v UpdateVDIPools200ResponseAnyOfVdiPoolConfigInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetName

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnvironment

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetConfig() UpdateVDIPools200ResponseAnyOfVdiPoolConfigConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetConfigOk() (*UpdateVDIPools200ResponseAnyOfVdiPoolConfigConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetConfig(v UpdateVDIPools200ResponseAnyOfVdiPoolConfigConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVolumes

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetVolumes() []UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetVolumesOk() (*[]UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetVolumes(v []UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetHostName

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetLayout

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetLayout() UpdateVDIPools200ResponseAnyOfVdiPoolConfigLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetLayoutOk() (*UpdateVDIPools200ResponseAnyOfVdiPoolConfigLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetLayout(v UpdateVDIPools200ResponseAnyOfVdiPoolConfigLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetStorageControllers

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetStorageControllers() []UpdateVDIPools200ResponseAnyOfVdiPoolConfigStorageControllersInner`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetStorageControllersOk() (*[]UpdateVDIPools200ResponseAnyOfVdiPoolConfigStorageControllersInner, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetStorageControllers(v []UpdateVDIPools200ResponseAnyOfVdiPoolConfigStorageControllersInner)`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### SetStorageControllersNil

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetStorageControllersNil(b bool)`

 SetStorageControllersNil sets the value for StorageControllers to be an explicit nil

### UnsetStorageControllers
`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) UnsetStorageControllers()`

UnsetStorageControllers ensures that no value is present for StorageControllers, not even an explicit nil
### GetPlan

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetPlan() UpdateVDIPools200ResponseAnyOfVdiPoolConfigPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetPlanOk() (*UpdateVDIPools200ResponseAnyOfVdiPoolConfigPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetPlan(v UpdateVDIPools200ResponseAnyOfVdiPoolConfigPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetVersion

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetNetworkInterfaces() []UpdateVDIPools200ResponseAnyOfVdiPoolConfigNetworkInterfacesInner`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetNetworkInterfacesOk() (*[]UpdateVDIPools200ResponseAnyOfVdiPoolConfigNetworkInterfacesInner, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetNetworkInterfaces(v []UpdateVDIPools200ResponseAnyOfVdiPoolConfigNetworkInterfacesInner)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetExecutionId

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetBackup

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetBackup() UpdateVDIPools200ResponseAnyOfVdiPoolConfigBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetBackupOk() (*UpdateVDIPools200ResponseAnyOfVdiPoolConfigBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetBackup(v UpdateVDIPools200ResponseAnyOfVdiPoolConfigBackup)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetLoadBalancer() []map[string]interface{}`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetLoadBalancerOk() (*[]map[string]interface{}, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetLoadBalancer(v []map[string]interface{})`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### SetLoadBalancerNil

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetLoadBalancerNil(b bool)`

 SetLoadBalancerNil sets the value for LoadBalancer to be an explicit nil

### UnsetLoadBalancer
`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) UnsetLoadBalancer()`

UnsetLoadBalancer ensures that no value is present for LoadBalancer, not even an explicit nil
### GetHideLock

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetHideLock() bool`

GetHideLock returns the HideLock field if non-nil, zero value otherwise.

### GetHideLockOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetHideLockOk() (*bool, bool)`

GetHideLockOk returns a tuple with the HideLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideLock

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetHideLock(v bool)`

SetHideLock sets HideLock field to given value.

### HasHideLock

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasHideLock() bool`

HasHideLock returns a boolean if a field has been set.

### GetHasNetworks

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetHasNetworks() bool`

GetHasNetworks returns the HasNetworks field if non-nil, zero value otherwise.

### GetHasNetworksOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetHasNetworksOk() (*bool, bool)`

GetHasNetworksOk returns a tuple with the HasNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworks

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetHasNetworks(v bool)`

SetHasNetworks sets HasNetworks field to given value.

### HasHasNetworks

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasHasNetworks() bool`

HasHasNetworks returns a boolean if a field has been set.

### GetDisplayNetworks

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetDisplayNetworks() []UpdateVDIPools200ResponseAnyOfVdiPoolConfigDisplayNetworksInner`

GetDisplayNetworks returns the DisplayNetworks field if non-nil, zero value otherwise.

### GetDisplayNetworksOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetDisplayNetworksOk() (*[]UpdateVDIPools200ResponseAnyOfVdiPoolConfigDisplayNetworksInner, bool)`

GetDisplayNetworksOk returns a tuple with the DisplayNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNetworks

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetDisplayNetworks(v []UpdateVDIPools200ResponseAnyOfVdiPoolConfigDisplayNetworksInner)`

SetDisplayNetworks sets DisplayNetworks field to given value.

### HasDisplayNetworks

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasDisplayNetworks() bool`

HasDisplayNetworks returns a boolean if a field has been set.

### GetCopies

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetCopies() int64`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetCopiesOk() (*int64, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetCopies(v int64)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetShowScale

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetShowScale() bool`

GetShowScale returns the ShowScale field if non-nil, zero value otherwise.

### GetShowScaleOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetShowScaleOk() (*bool, bool)`

GetShowScaleOk returns a tuple with the ShowScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowScale

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetShowScale(v bool)`

SetShowScale sets ShowScale field to given value.

### HasShowScale

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasShowScale() bool`

HasShowScale returns a boolean if a field has been set.

### GetHasPreview

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetHasPreview() bool`

GetHasPreview returns the HasPreview field if non-nil, zero value otherwise.

### GetHasPreviewOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetHasPreviewOk() (*bool, bool)`

GetHasPreviewOk returns a tuple with the HasPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreview

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetHasPreview(v bool)`

SetHasPreview sets HasPreview field to given value.

### HasHasPreview

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasHasPreview() bool`

HasHasPreview returns a boolean if a field has been set.

### GetVolumesDisplay

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetVolumesDisplay() []UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesDisplayInner`

GetVolumesDisplay returns the VolumesDisplay field if non-nil, zero value otherwise.

### GetVolumesDisplayOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) GetVolumesDisplayOk() (*[]UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesDisplayInner, bool)`

GetVolumesDisplayOk returns a tuple with the VolumesDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesDisplay

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) SetVolumesDisplay(v []UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesDisplayInner)`

SetVolumesDisplay sets VolumesDisplay field to given value.

### HasVolumesDisplay

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfig) HasVolumesDisplay() bool`

HasVolumesDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


