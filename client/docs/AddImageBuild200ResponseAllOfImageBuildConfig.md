# AddImageBuild200ResponseAllOfImageBuildConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | Pointer to [**AddImageBuild200ResponseAllOfImageBuildConfigInstance**](AddImageBuild200ResponseAllOfImageBuildConfigInstance.md) |  | [optional] 
**NetworkInterfaces** | Pointer to [**[]AddImageBuild200ResponseAllOfImageBuildConfigNetworkInterfacesInner**](AddImageBuild200ResponseAllOfImageBuildConfigNetworkInterfacesInner.md) |  | [optional] 
**Volumes** | Pointer to [**[]AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner**](AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner.md) |  | [optional] 
**StorageControllers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**Config** | Pointer to [**AddImageBuild200ResponseAllOfImageBuildConfigConfig**](AddImageBuild200ResponseAllOfImageBuildConfigConfig.md) |  | [optional] 
**Plan** | Pointer to [**ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan**](ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan.md) |  | [optional] 

## Methods

### NewAddImageBuild200ResponseAllOfImageBuildConfig

`func NewAddImageBuild200ResponseAllOfImageBuildConfig() *AddImageBuild200ResponseAllOfImageBuildConfig`

NewAddImageBuild200ResponseAllOfImageBuildConfig instantiates a new AddImageBuild200ResponseAllOfImageBuildConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddImageBuild200ResponseAllOfImageBuildConfigWithDefaults

`func NewAddImageBuild200ResponseAllOfImageBuildConfigWithDefaults() *AddImageBuild200ResponseAllOfImageBuildConfig`

NewAddImageBuild200ResponseAllOfImageBuildConfigWithDefaults instantiates a new AddImageBuild200ResponseAllOfImageBuildConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) GetInstance() AddImageBuild200ResponseAllOfImageBuildConfigInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) GetInstanceOk() (*AddImageBuild200ResponseAllOfImageBuildConfigInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) SetInstance(v AddImageBuild200ResponseAllOfImageBuildConfigInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) GetNetworkInterfaces() []AddImageBuild200ResponseAllOfImageBuildConfigNetworkInterfacesInner`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) GetNetworkInterfacesOk() (*[]AddImageBuild200ResponseAllOfImageBuildConfigNetworkInterfacesInner, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) SetNetworkInterfaces(v []AddImageBuild200ResponseAllOfImageBuildConfigNetworkInterfacesInner)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetVolumes

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) GetVolumes() []AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) GetVolumesOk() (*[]AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) SetVolumes(v []AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetStorageControllers

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) GetStorageControllers() []map[string]interface{}`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) GetStorageControllersOk() (*[]map[string]interface{}, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) SetStorageControllers(v []map[string]interface{})`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### GetZoneId

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetConfig

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) GetConfig() AddImageBuild200ResponseAllOfImageBuildConfigConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) GetConfigOk() (*AddImageBuild200ResponseAllOfImageBuildConfigConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) SetConfig(v AddImageBuild200ResponseAllOfImageBuildConfigConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetPlan

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) GetPlan() ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) GetPlanOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) SetPlan(v ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AddImageBuild200ResponseAllOfImageBuildConfig) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


