# GetImageBuild200ResponseImageBuildConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | Pointer to [**GetImageBuild200ResponseImageBuildConfigInstance**](GetImageBuild200ResponseImageBuildConfigInstance.md) |  | [optional] 
**NetworkInterfaces** | Pointer to [**[]GetImageBuild200ResponseImageBuildConfigNetworkInterfacesInner**](GetImageBuild200ResponseImageBuildConfigNetworkInterfacesInner.md) |  | [optional] 
**Volumes** | Pointer to [**[]GetImageBuild200ResponseImageBuildConfigVolumesInner**](GetImageBuild200ResponseImageBuildConfigVolumesInner.md) |  | [optional] 
**StorageControllers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**Config** | Pointer to [**GetImageBuild200ResponseImageBuildConfigConfig**](GetImageBuild200ResponseImageBuildConfigConfig.md) |  | [optional] 
**Plan** | Pointer to [**GetImageBuild200ResponseImageBuildConfigPlan**](GetImageBuild200ResponseImageBuildConfigPlan.md) |  | [optional] 

## Methods

### NewGetImageBuild200ResponseImageBuildConfig

`func NewGetImageBuild200ResponseImageBuildConfig() *GetImageBuild200ResponseImageBuildConfig`

NewGetImageBuild200ResponseImageBuildConfig instantiates a new GetImageBuild200ResponseImageBuildConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetInstance

`func (o *GetImageBuild200ResponseImageBuildConfig) GetInstance() GetImageBuild200ResponseImageBuildConfigInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *GetImageBuild200ResponseImageBuildConfig) GetInstanceOk() (*GetImageBuild200ResponseImageBuildConfigInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *GetImageBuild200ResponseImageBuildConfig) SetInstance(v GetImageBuild200ResponseImageBuildConfigInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *GetImageBuild200ResponseImageBuildConfig) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *GetImageBuild200ResponseImageBuildConfig) GetNetworkInterfaces() []GetImageBuild200ResponseImageBuildConfigNetworkInterfacesInner`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *GetImageBuild200ResponseImageBuildConfig) GetNetworkInterfacesOk() (*[]GetImageBuild200ResponseImageBuildConfigNetworkInterfacesInner, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *GetImageBuild200ResponseImageBuildConfig) SetNetworkInterfaces(v []GetImageBuild200ResponseImageBuildConfigNetworkInterfacesInner)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *GetImageBuild200ResponseImageBuildConfig) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetVolumes

`func (o *GetImageBuild200ResponseImageBuildConfig) GetVolumes() []GetImageBuild200ResponseImageBuildConfigVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *GetImageBuild200ResponseImageBuildConfig) GetVolumesOk() (*[]GetImageBuild200ResponseImageBuildConfigVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *GetImageBuild200ResponseImageBuildConfig) SetVolumes(v []GetImageBuild200ResponseImageBuildConfigVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *GetImageBuild200ResponseImageBuildConfig) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetStorageControllers

`func (o *GetImageBuild200ResponseImageBuildConfig) GetStorageControllers() []map[string]interface{}`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *GetImageBuild200ResponseImageBuildConfig) GetStorageControllersOk() (*[]map[string]interface{}, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *GetImageBuild200ResponseImageBuildConfig) SetStorageControllers(v []map[string]interface{})`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *GetImageBuild200ResponseImageBuildConfig) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### GetZoneId

`func (o *GetImageBuild200ResponseImageBuildConfig) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *GetImageBuild200ResponseImageBuildConfig) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *GetImageBuild200ResponseImageBuildConfig) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *GetImageBuild200ResponseImageBuildConfig) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetConfig

`func (o *GetImageBuild200ResponseImageBuildConfig) GetConfig() GetImageBuild200ResponseImageBuildConfigConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetImageBuild200ResponseImageBuildConfig) GetConfigOk() (*GetImageBuild200ResponseImageBuildConfigConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetImageBuild200ResponseImageBuildConfig) SetConfig(v GetImageBuild200ResponseImageBuildConfigConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetImageBuild200ResponseImageBuildConfig) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetPlan

`func (o *GetImageBuild200ResponseImageBuildConfig) GetPlan() GetImageBuild200ResponseImageBuildConfigPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *GetImageBuild200ResponseImageBuildConfig) GetPlanOk() (*GetImageBuild200ResponseImageBuildConfigPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *GetImageBuild200ResponseImageBuildConfig) SetPlan(v GetImageBuild200ResponseImageBuildConfigPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *GetImageBuild200ResponseImageBuildConfig) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


