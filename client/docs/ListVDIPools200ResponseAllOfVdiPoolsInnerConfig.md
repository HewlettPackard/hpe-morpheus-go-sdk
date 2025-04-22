# ListVDIPools200ResponseAllOfVdiPoolsInnerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner**](ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner.md) |  | [optional] 
**Cloud** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
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
**DisplayNetworks** | Pointer to [**[]ListVDIPools200ResponseAllOfVdiPoolsInnerConfigDisplayNetworksInner**](ListVDIPools200ResponseAllOfVdiPoolsInnerConfigDisplayNetworksInner.md) |  | [optional] 
**Copies** | Pointer to **int64** |  | [optional] 
**ShowScale** | Pointer to **bool** |  | [optional] 
**HasPreview** | Pointer to **bool** |  | [optional] 
**VolumesDisplay** | Pointer to [**[]ListVDIPools200ResponseAllOfVdiPoolsInnerConfigVolumesDisplayInner**](ListVDIPools200ResponseAllOfVdiPoolsInnerConfigVolumesDisplayInner.md) |  | [optional] 

## Methods

### NewListVDIPools200ResponseAllOfVdiPoolsInnerConfig

`func NewListVDIPools200ResponseAllOfVdiPoolsInnerConfig() *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig`

NewListVDIPools200ResponseAllOfVdiPoolsInnerConfig instantiates a new ListVDIPools200ResponseAllOfVdiPoolsInnerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVDIPools200ResponseAllOfVdiPoolsInnerConfigWithDefaults

`func NewListVDIPools200ResponseAllOfVdiPoolsInnerConfigWithDefaults() *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig`

NewListVDIPools200ResponseAllOfVdiPoolsInnerConfigWithDefaults instantiates a new ListVDIPools200ResponseAllOfVdiPoolsInnerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetGroup() ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetGroupOk() (*ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetGroup(v ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCloud

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetCloud() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetCloudOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetCloud(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetType

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInstance

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetInstance() ListVDIPools200ResponseAllOfVdiPoolsInnerConfigInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetInstanceOk() (*ListVDIPools200ResponseAllOfVdiPoolsInnerConfigInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetInstance(v ListVDIPools200ResponseAllOfVdiPoolsInnerConfigInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetName

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnvironment

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetConfig

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetConfig() ListVDIPools200ResponseAllOfVdiPoolsInnerConfigConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetConfigOk() (*ListVDIPools200ResponseAllOfVdiPoolsInnerConfigConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetConfig(v ListVDIPools200ResponseAllOfVdiPoolsInnerConfigConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVolumes

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetVolumes() []ListImageBuilds200ResponseAllOfImageBuildsInnerConfigVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetVolumesOk() (*[]ListImageBuilds200ResponseAllOfImageBuildsInnerConfigVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetVolumes(v []ListImageBuilds200ResponseAllOfImageBuildsInnerConfigVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetHostName

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetLayout

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetLayout() ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetLayoutOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetLayout(v ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetStorageControllers

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetStorageControllers() []ListVDIPools200ResponseAllOfVdiPoolsInnerConfigStorageControllersInner`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetStorageControllersOk() (*[]ListVDIPools200ResponseAllOfVdiPoolsInnerConfigStorageControllersInner, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetStorageControllers(v []ListVDIPools200ResponseAllOfVdiPoolsInnerConfigStorageControllersInner)`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### GetPlan

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetPlan() ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetPlanOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetPlan(v ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetVersion

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetNetworkInterfaces() []ListVDIPools200ResponseAllOfVdiPoolsInnerConfigNetworkInterfacesInner`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetNetworkInterfacesOk() (*[]ListVDIPools200ResponseAllOfVdiPoolsInnerConfigNetworkInterfacesInner, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetNetworkInterfaces(v []ListVDIPools200ResponseAllOfVdiPoolsInnerConfigNetworkInterfacesInner)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetExecutionId

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetBackup

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetBackup() ListInstances200ResponseAllOfInstancesInnerConfigBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetBackupOk() (*ListInstances200ResponseAllOfInstancesInnerConfigBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetBackup(v ListInstances200ResponseAllOfInstancesInnerConfigBackup)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetLoadBalancer() []map[string]interface{}`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetLoadBalancerOk() (*[]map[string]interface{}, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetLoadBalancer(v []map[string]interface{})`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetHideLock

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetHideLock() bool`

GetHideLock returns the HideLock field if non-nil, zero value otherwise.

### GetHideLockOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetHideLockOk() (*bool, bool)`

GetHideLockOk returns a tuple with the HideLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideLock

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetHideLock(v bool)`

SetHideLock sets HideLock field to given value.

### HasHideLock

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasHideLock() bool`

HasHideLock returns a boolean if a field has been set.

### GetHasNetworks

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetHasNetworks() bool`

GetHasNetworks returns the HasNetworks field if non-nil, zero value otherwise.

### GetHasNetworksOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetHasNetworksOk() (*bool, bool)`

GetHasNetworksOk returns a tuple with the HasNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworks

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetHasNetworks(v bool)`

SetHasNetworks sets HasNetworks field to given value.

### HasHasNetworks

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasHasNetworks() bool`

HasHasNetworks returns a boolean if a field has been set.

### GetDisplayNetworks

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetDisplayNetworks() []ListVDIPools200ResponseAllOfVdiPoolsInnerConfigDisplayNetworksInner`

GetDisplayNetworks returns the DisplayNetworks field if non-nil, zero value otherwise.

### GetDisplayNetworksOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetDisplayNetworksOk() (*[]ListVDIPools200ResponseAllOfVdiPoolsInnerConfigDisplayNetworksInner, bool)`

GetDisplayNetworksOk returns a tuple with the DisplayNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNetworks

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetDisplayNetworks(v []ListVDIPools200ResponseAllOfVdiPoolsInnerConfigDisplayNetworksInner)`

SetDisplayNetworks sets DisplayNetworks field to given value.

### HasDisplayNetworks

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasDisplayNetworks() bool`

HasDisplayNetworks returns a boolean if a field has been set.

### GetCopies

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetCopies() int64`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetCopiesOk() (*int64, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetCopies(v int64)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetShowScale

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetShowScale() bool`

GetShowScale returns the ShowScale field if non-nil, zero value otherwise.

### GetShowScaleOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetShowScaleOk() (*bool, bool)`

GetShowScaleOk returns a tuple with the ShowScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowScale

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetShowScale(v bool)`

SetShowScale sets ShowScale field to given value.

### HasShowScale

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasShowScale() bool`

HasShowScale returns a boolean if a field has been set.

### GetHasPreview

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetHasPreview() bool`

GetHasPreview returns the HasPreview field if non-nil, zero value otherwise.

### GetHasPreviewOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetHasPreviewOk() (*bool, bool)`

GetHasPreviewOk returns a tuple with the HasPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreview

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetHasPreview(v bool)`

SetHasPreview sets HasPreview field to given value.

### HasHasPreview

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasHasPreview() bool`

HasHasPreview returns a boolean if a field has been set.

### GetVolumesDisplay

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetVolumesDisplay() []ListVDIPools200ResponseAllOfVdiPoolsInnerConfigVolumesDisplayInner`

GetVolumesDisplay returns the VolumesDisplay field if non-nil, zero value otherwise.

### GetVolumesDisplayOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) GetVolumesDisplayOk() (*[]ListVDIPools200ResponseAllOfVdiPoolsInnerConfigVolumesDisplayInner, bool)`

GetVolumesDisplayOk returns a tuple with the VolumesDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesDisplay

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) SetVolumesDisplay(v []ListVDIPools200ResponseAllOfVdiPoolsInnerConfigVolumesDisplayInner)`

SetVolumesDisplay sets VolumesDisplay field to given value.

### HasVolumesDisplay

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInnerConfig) HasVolumesDisplay() bool`

HasVolumesDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


