# AddVDIPools200ResponseAnyOfVdiPoolConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner**](ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner.md) |  | [optional] 
**Cloud** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Instance** | Pointer to [**ListVDIPools200ResponseAllOfVdiPoolsInnerConfigInstance**](ListVDIPools200ResponseAllOfVdiPoolsInnerConfigInstance.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**ListVDIPools200ResponseAllOfVdiPoolsInnerConfigConfig**](ListVDIPools200ResponseAllOfVdiPoolsInnerConfigConfig.md) |  | [optional] 
**Volumes** | Pointer to [**[]ListImageBuilds200ResponseAllOfImageBuildsInnerConfigVolumesInner**](ListImageBuilds200ResponseAllOfImageBuildsInnerConfigVolumesInner.md) |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**Layout** | Pointer to [**ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan**](ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan.md) |  | [optional] 
**StorageControllers** | Pointer to [**[]ListVDIPools200ResponseAllOfVdiPoolsInnerConfigStorageControllersInner**](ListVDIPools200ResponseAllOfVdiPoolsInnerConfigStorageControllersInner.md) |  | [optional] 
**Plan** | Pointer to [**ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan**](ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**NetworkInterfaces** | Pointer to [**[]ListVDIPools200ResponseAllOfVdiPoolsInnerConfigNetworkInterfacesInner**](ListVDIPools200ResponseAllOfVdiPoolsInnerConfigNetworkInterfacesInner.md) |  | [optional] 
**ExecutionId** | Pointer to **string** |  | [optional] 
**Backup** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerConfigBackup**](ListInstances200ResponseAllOfInstancesInnerConfigBackup.md) |  | [optional] 
**LoadBalancer** | Pointer to **[]map[string]interface{}** |  | [optional] 
**HideLock** | Pointer to **bool** |  | [optional] 
**HasNetworks** | Pointer to **bool** |  | [optional] 
**DisplayNetworks** | Pointer to [**[]AddVDIPools200ResponseAnyOfVdiPoolConfigDisplayNetworksInner**](AddVDIPools200ResponseAnyOfVdiPoolConfigDisplayNetworksInner.md) |  | [optional] 
**Copies** | Pointer to **int64** |  | [optional] 
**ShowScale** | Pointer to **bool** |  | [optional] 
**HasPreview** | Pointer to **bool** |  | [optional] 
**VolumesDisplay** | Pointer to [**[]ListVDIPools200ResponseAllOfVdiPoolsInnerConfigVolumesDisplayInner**](ListVDIPools200ResponseAllOfVdiPoolsInnerConfigVolumesDisplayInner.md) |  | [optional] 

## Methods

### NewAddVDIPools200ResponseAnyOfVdiPoolConfig

`func NewAddVDIPools200ResponseAnyOfVdiPoolConfig() *AddVDIPools200ResponseAnyOfVdiPoolConfig`

NewAddVDIPools200ResponseAnyOfVdiPoolConfig instantiates a new AddVDIPools200ResponseAnyOfVdiPoolConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVDIPools200ResponseAnyOfVdiPoolConfigWithDefaults

`func NewAddVDIPools200ResponseAnyOfVdiPoolConfigWithDefaults() *AddVDIPools200ResponseAnyOfVdiPoolConfig`

NewAddVDIPools200ResponseAnyOfVdiPoolConfigWithDefaults instantiates a new AddVDIPools200ResponseAnyOfVdiPoolConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetGroup() ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetGroupOk() (*ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetGroup(v ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCloud

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetCloud() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetCloudOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetCloud(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetType

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInstance

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetInstance() ListVDIPools200ResponseAllOfVdiPoolsInnerConfigInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetInstanceOk() (*ListVDIPools200ResponseAllOfVdiPoolsInnerConfigInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetInstance(v ListVDIPools200ResponseAllOfVdiPoolsInnerConfigInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetName

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnvironment

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetConfig

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetConfig() ListVDIPools200ResponseAllOfVdiPoolsInnerConfigConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetConfigOk() (*ListVDIPools200ResponseAllOfVdiPoolsInnerConfigConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetConfig(v ListVDIPools200ResponseAllOfVdiPoolsInnerConfigConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVolumes

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetVolumes() []ListImageBuilds200ResponseAllOfImageBuildsInnerConfigVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetVolumesOk() (*[]ListImageBuilds200ResponseAllOfImageBuildsInnerConfigVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetVolumes(v []ListImageBuilds200ResponseAllOfImageBuildsInnerConfigVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetHostName

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetLayout

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetLayout() ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetLayoutOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetLayout(v ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetStorageControllers

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetStorageControllers() []ListVDIPools200ResponseAllOfVdiPoolsInnerConfigStorageControllersInner`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetStorageControllersOk() (*[]ListVDIPools200ResponseAllOfVdiPoolsInnerConfigStorageControllersInner, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetStorageControllers(v []ListVDIPools200ResponseAllOfVdiPoolsInnerConfigStorageControllersInner)`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### SetStorageControllersNil

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetStorageControllersNil(b bool)`

 SetStorageControllersNil sets the value for StorageControllers to be an explicit nil

### UnsetStorageControllers
`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) UnsetStorageControllers()`

UnsetStorageControllers ensures that no value is present for StorageControllers, not even an explicit nil
### GetPlan

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetPlan() ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetPlanOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetPlan(v ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetVersion

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetNetworkInterfaces() []ListVDIPools200ResponseAllOfVdiPoolsInnerConfigNetworkInterfacesInner`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetNetworkInterfacesOk() (*[]ListVDIPools200ResponseAllOfVdiPoolsInnerConfigNetworkInterfacesInner, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetNetworkInterfaces(v []ListVDIPools200ResponseAllOfVdiPoolsInnerConfigNetworkInterfacesInner)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetExecutionId

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetBackup

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetBackup() ListInstances200ResponseAllOfInstancesInnerConfigBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetBackupOk() (*ListInstances200ResponseAllOfInstancesInnerConfigBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetBackup(v ListInstances200ResponseAllOfInstancesInnerConfigBackup)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetLoadBalancer() []map[string]interface{}`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetLoadBalancerOk() (*[]map[string]interface{}, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetLoadBalancer(v []map[string]interface{})`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### SetLoadBalancerNil

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetLoadBalancerNil(b bool)`

 SetLoadBalancerNil sets the value for LoadBalancer to be an explicit nil

### UnsetLoadBalancer
`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) UnsetLoadBalancer()`

UnsetLoadBalancer ensures that no value is present for LoadBalancer, not even an explicit nil
### GetHideLock

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetHideLock() bool`

GetHideLock returns the HideLock field if non-nil, zero value otherwise.

### GetHideLockOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetHideLockOk() (*bool, bool)`

GetHideLockOk returns a tuple with the HideLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideLock

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetHideLock(v bool)`

SetHideLock sets HideLock field to given value.

### HasHideLock

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasHideLock() bool`

HasHideLock returns a boolean if a field has been set.

### GetHasNetworks

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetHasNetworks() bool`

GetHasNetworks returns the HasNetworks field if non-nil, zero value otherwise.

### GetHasNetworksOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetHasNetworksOk() (*bool, bool)`

GetHasNetworksOk returns a tuple with the HasNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworks

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetHasNetworks(v bool)`

SetHasNetworks sets HasNetworks field to given value.

### HasHasNetworks

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasHasNetworks() bool`

HasHasNetworks returns a boolean if a field has been set.

### GetDisplayNetworks

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetDisplayNetworks() []AddVDIPools200ResponseAnyOfVdiPoolConfigDisplayNetworksInner`

GetDisplayNetworks returns the DisplayNetworks field if non-nil, zero value otherwise.

### GetDisplayNetworksOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetDisplayNetworksOk() (*[]AddVDIPools200ResponseAnyOfVdiPoolConfigDisplayNetworksInner, bool)`

GetDisplayNetworksOk returns a tuple with the DisplayNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNetworks

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetDisplayNetworks(v []AddVDIPools200ResponseAnyOfVdiPoolConfigDisplayNetworksInner)`

SetDisplayNetworks sets DisplayNetworks field to given value.

### HasDisplayNetworks

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasDisplayNetworks() bool`

HasDisplayNetworks returns a boolean if a field has been set.

### GetCopies

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetCopies() int64`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetCopiesOk() (*int64, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetCopies(v int64)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetShowScale

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetShowScale() bool`

GetShowScale returns the ShowScale field if non-nil, zero value otherwise.

### GetShowScaleOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetShowScaleOk() (*bool, bool)`

GetShowScaleOk returns a tuple with the ShowScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowScale

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetShowScale(v bool)`

SetShowScale sets ShowScale field to given value.

### HasShowScale

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasShowScale() bool`

HasShowScale returns a boolean if a field has been set.

### GetHasPreview

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetHasPreview() bool`

GetHasPreview returns the HasPreview field if non-nil, zero value otherwise.

### GetHasPreviewOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetHasPreviewOk() (*bool, bool)`

GetHasPreviewOk returns a tuple with the HasPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreview

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetHasPreview(v bool)`

SetHasPreview sets HasPreview field to given value.

### HasHasPreview

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasHasPreview() bool`

HasHasPreview returns a boolean if a field has been set.

### GetVolumesDisplay

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetVolumesDisplay() []ListVDIPools200ResponseAllOfVdiPoolsInnerConfigVolumesDisplayInner`

GetVolumesDisplay returns the VolumesDisplay field if non-nil, zero value otherwise.

### GetVolumesDisplayOk

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) GetVolumesDisplayOk() (*[]ListVDIPools200ResponseAllOfVdiPoolsInnerConfigVolumesDisplayInner, bool)`

GetVolumesDisplayOk returns a tuple with the VolumesDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesDisplay

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) SetVolumesDisplay(v []ListVDIPools200ResponseAllOfVdiPoolsInnerConfigVolumesDisplayInner)`

SetVolumesDisplay sets VolumesDisplay field to given value.

### HasVolumesDisplay

`func (o *AddVDIPools200ResponseAnyOfVdiPoolConfig) HasVolumesDisplay() bool`

HasVolumesDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


