# GetVDIPools200ResponseVdiPoolConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**GetVDIPools200ResponseVdiPoolConfigGroup**](GetVDIPools200ResponseVdiPoolConfigGroup.md) |  | [optional] 
**Cloud** | Pointer to [**GetVDIPools200ResponseVdiPoolConfigCloud**](GetVDIPools200ResponseVdiPoolConfigCloud.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Instance** | Pointer to [**GetVDIPools200ResponseVdiPoolConfigInstance**](GetVDIPools200ResponseVdiPoolConfigInstance.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**GetVDIPools200ResponseVdiPoolConfigConfig**](GetVDIPools200ResponseVdiPoolConfigConfig.md) |  | [optional] 
**Volumes** | Pointer to [**[]GetVDIPools200ResponseVdiPoolConfigVolumesInner**](GetVDIPools200ResponseVdiPoolConfigVolumesInner.md) |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**Layout** | Pointer to [**GetVDIPools200ResponseVdiPoolConfigLayout**](GetVDIPools200ResponseVdiPoolConfigLayout.md) |  | [optional] 
**StorageControllers** | Pointer to [**[]GetVDIPools200ResponseVdiPoolConfigStorageControllersInner**](GetVDIPools200ResponseVdiPoolConfigStorageControllersInner.md) |  | [optional] 
**Plan** | Pointer to [**GetVDIPools200ResponseVdiPoolConfigPlan**](GetVDIPools200ResponseVdiPoolConfigPlan.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**NetworkInterfaces** | Pointer to [**[]GetVDIPools200ResponseVdiPoolConfigNetworkInterfacesInner**](GetVDIPools200ResponseVdiPoolConfigNetworkInterfacesInner.md) |  | [optional] 
**ExecutionId** | Pointer to **string** |  | [optional] 
**Backup** | Pointer to [**GetVDIPools200ResponseVdiPoolConfigBackup**](GetVDIPools200ResponseVdiPoolConfigBackup.md) |  | [optional] 
**LoadBalancer** | Pointer to **[]map[string]interface{}** |  | [optional] 
**HideLock** | Pointer to **bool** |  | [optional] 
**HasNetworks** | Pointer to **bool** |  | [optional] 
**DisplayNetworks** | Pointer to [**[]GetVDIPools200ResponseVdiPoolConfigDisplayNetworksInner**](GetVDIPools200ResponseVdiPoolConfigDisplayNetworksInner.md) |  | [optional] 
**Copies** | Pointer to **int64** |  | [optional] 
**ShowScale** | Pointer to **bool** |  | [optional] 
**HasPreview** | Pointer to **bool** |  | [optional] 
**VolumesDisplay** | Pointer to [**[]GetVDIPools200ResponseVdiPoolConfigVolumesDisplayInner**](GetVDIPools200ResponseVdiPoolConfigVolumesDisplayInner.md) |  | [optional] 

## Methods

### NewGetVDIPools200ResponseVdiPoolConfig

`func NewGetVDIPools200ResponseVdiPoolConfig() *GetVDIPools200ResponseVdiPoolConfig`

NewGetVDIPools200ResponseVdiPoolConfig instantiates a new GetVDIPools200ResponseVdiPoolConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVDIPools200ResponseVdiPoolConfigWithDefaults

`func NewGetVDIPools200ResponseVdiPoolConfigWithDefaults() *GetVDIPools200ResponseVdiPoolConfig`

NewGetVDIPools200ResponseVdiPoolConfigWithDefaults instantiates a new GetVDIPools200ResponseVdiPoolConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetGroup() GetVDIPools200ResponseVdiPoolConfigGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetGroupOk() (*GetVDIPools200ResponseVdiPoolConfigGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetGroup(v GetVDIPools200ResponseVdiPoolConfigGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCloud

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetCloud() GetVDIPools200ResponseVdiPoolConfigCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetCloudOk() (*GetVDIPools200ResponseVdiPoolConfigCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetCloud(v GetVDIPools200ResponseVdiPoolConfigCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetType

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInstance

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetInstance() GetVDIPools200ResponseVdiPoolConfigInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetInstanceOk() (*GetVDIPools200ResponseVdiPoolConfigInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetInstance(v GetVDIPools200ResponseVdiPoolConfigInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetName

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnvironment

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetConfig

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetConfig() GetVDIPools200ResponseVdiPoolConfigConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetConfigOk() (*GetVDIPools200ResponseVdiPoolConfigConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetConfig(v GetVDIPools200ResponseVdiPoolConfigConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVolumes

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetVolumes() []GetVDIPools200ResponseVdiPoolConfigVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetVolumesOk() (*[]GetVDIPools200ResponseVdiPoolConfigVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetVolumes(v []GetVDIPools200ResponseVdiPoolConfigVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetHostName

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetLayout

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetLayout() GetVDIPools200ResponseVdiPoolConfigLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetLayoutOk() (*GetVDIPools200ResponseVdiPoolConfigLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetLayout(v GetVDIPools200ResponseVdiPoolConfigLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetStorageControllers

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetStorageControllers() []GetVDIPools200ResponseVdiPoolConfigStorageControllersInner`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetStorageControllersOk() (*[]GetVDIPools200ResponseVdiPoolConfigStorageControllersInner, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetStorageControllers(v []GetVDIPools200ResponseVdiPoolConfigStorageControllersInner)`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### SetStorageControllersNil

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetStorageControllersNil(b bool)`

 SetStorageControllersNil sets the value for StorageControllers to be an explicit nil

### UnsetStorageControllers
`func (o *GetVDIPools200ResponseVdiPoolConfig) UnsetStorageControllers()`

UnsetStorageControllers ensures that no value is present for StorageControllers, not even an explicit nil
### GetPlan

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetPlan() GetVDIPools200ResponseVdiPoolConfigPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetPlanOk() (*GetVDIPools200ResponseVdiPoolConfigPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetPlan(v GetVDIPools200ResponseVdiPoolConfigPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetVersion

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetNetworkInterfaces() []GetVDIPools200ResponseVdiPoolConfigNetworkInterfacesInner`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetNetworkInterfacesOk() (*[]GetVDIPools200ResponseVdiPoolConfigNetworkInterfacesInner, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetNetworkInterfaces(v []GetVDIPools200ResponseVdiPoolConfigNetworkInterfacesInner)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetExecutionId

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetBackup

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetBackup() GetVDIPools200ResponseVdiPoolConfigBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetBackupOk() (*GetVDIPools200ResponseVdiPoolConfigBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetBackup(v GetVDIPools200ResponseVdiPoolConfigBackup)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetLoadBalancer() []map[string]interface{}`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetLoadBalancerOk() (*[]map[string]interface{}, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetLoadBalancer(v []map[string]interface{})`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### SetLoadBalancerNil

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetLoadBalancerNil(b bool)`

 SetLoadBalancerNil sets the value for LoadBalancer to be an explicit nil

### UnsetLoadBalancer
`func (o *GetVDIPools200ResponseVdiPoolConfig) UnsetLoadBalancer()`

UnsetLoadBalancer ensures that no value is present for LoadBalancer, not even an explicit nil
### GetHideLock

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetHideLock() bool`

GetHideLock returns the HideLock field if non-nil, zero value otherwise.

### GetHideLockOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetHideLockOk() (*bool, bool)`

GetHideLockOk returns a tuple with the HideLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideLock

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetHideLock(v bool)`

SetHideLock sets HideLock field to given value.

### HasHideLock

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasHideLock() bool`

HasHideLock returns a boolean if a field has been set.

### GetHasNetworks

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetHasNetworks() bool`

GetHasNetworks returns the HasNetworks field if non-nil, zero value otherwise.

### GetHasNetworksOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetHasNetworksOk() (*bool, bool)`

GetHasNetworksOk returns a tuple with the HasNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworks

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetHasNetworks(v bool)`

SetHasNetworks sets HasNetworks field to given value.

### HasHasNetworks

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasHasNetworks() bool`

HasHasNetworks returns a boolean if a field has been set.

### GetDisplayNetworks

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetDisplayNetworks() []GetVDIPools200ResponseVdiPoolConfigDisplayNetworksInner`

GetDisplayNetworks returns the DisplayNetworks field if non-nil, zero value otherwise.

### GetDisplayNetworksOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetDisplayNetworksOk() (*[]GetVDIPools200ResponseVdiPoolConfigDisplayNetworksInner, bool)`

GetDisplayNetworksOk returns a tuple with the DisplayNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNetworks

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetDisplayNetworks(v []GetVDIPools200ResponseVdiPoolConfigDisplayNetworksInner)`

SetDisplayNetworks sets DisplayNetworks field to given value.

### HasDisplayNetworks

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasDisplayNetworks() bool`

HasDisplayNetworks returns a boolean if a field has been set.

### GetCopies

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetCopies() int64`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetCopiesOk() (*int64, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetCopies(v int64)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetShowScale

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetShowScale() bool`

GetShowScale returns the ShowScale field if non-nil, zero value otherwise.

### GetShowScaleOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetShowScaleOk() (*bool, bool)`

GetShowScaleOk returns a tuple with the ShowScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowScale

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetShowScale(v bool)`

SetShowScale sets ShowScale field to given value.

### HasShowScale

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasShowScale() bool`

HasShowScale returns a boolean if a field has been set.

### GetHasPreview

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetHasPreview() bool`

GetHasPreview returns the HasPreview field if non-nil, zero value otherwise.

### GetHasPreviewOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetHasPreviewOk() (*bool, bool)`

GetHasPreviewOk returns a tuple with the HasPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreview

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetHasPreview(v bool)`

SetHasPreview sets HasPreview field to given value.

### HasHasPreview

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasHasPreview() bool`

HasHasPreview returns a boolean if a field has been set.

### GetVolumesDisplay

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetVolumesDisplay() []GetVDIPools200ResponseVdiPoolConfigVolumesDisplayInner`

GetVolumesDisplay returns the VolumesDisplay field if non-nil, zero value otherwise.

### GetVolumesDisplayOk

`func (o *GetVDIPools200ResponseVdiPoolConfig) GetVolumesDisplayOk() (*[]GetVDIPools200ResponseVdiPoolConfigVolumesDisplayInner, bool)`

GetVolumesDisplayOk returns a tuple with the VolumesDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesDisplay

`func (o *GetVDIPools200ResponseVdiPoolConfig) SetVolumesDisplay(v []GetVDIPools200ResponseVdiPoolConfigVolumesDisplayInner)`

SetVolumesDisplay sets VolumesDisplay field to given value.

### HasVolumesDisplay

`func (o *GetVDIPools200ResponseVdiPoolConfig) HasVolumesDisplay() bool`

HasVolumesDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


