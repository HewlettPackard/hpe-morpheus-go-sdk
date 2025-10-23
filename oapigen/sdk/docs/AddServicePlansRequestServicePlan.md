# AddServicePlansRequestServicePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Service plan name | 
**Code** | **string** | Service plan code, must be unique | 
**Description** | Pointer to **string** | Service plan description | [optional] 
**RegionCode** | Pointer to **string** | Service plan region code | [optional] 
**Editable** | Pointer to **bool** | Can be used to enable / disable the editability of the service plan. | [optional] [default to true]
**Visibility** | Pointer to **string** | Can be used to enable / disable the visibility of the service plan, defaults to \&quot;public\&quot; unless the account is not a masterAccount in which the default is \&quot;private\&quot; | [optional] 
**MaxStorage** | **int64** | Max storage size in bytes | 
**MaxMemory** | **int64** | Max memory size in bytes | 
**MaxCores** | Pointer to **int64** | Max number of cores | [optional] 
**MaxCpu** | Pointer to **int64** | Max number of CPUs | [optional] 
**CoresPerSocket** | Pointer to **int64** | Number of cores per CPU | [optional] 
**MaxGpus** | Pointer to **int64** | Max number of GPUs | [optional] 
**MaxDisks** | Pointer to **int64** | Max disks allowed | [optional] 
**ProvisionType** | [**AddClusterLayoutsRequestLayoutProvisionType**](AddClusterLayoutsRequestLayoutProvisionType.md) |  | 
**CustomCpu** | Pointer to **bool** | Can be used to enable / disable customizable cpu | [optional] [default to false]
**CustomCores** | Pointer to **bool** | Can be used to enable / disable customizable cores | [optional] [default to false]
**CustomMaxStorage** | Pointer to **bool** | Can be used to enable / disable customizable storage | [optional] [default to false]
**CustomMaxDataStorage** | Pointer to **bool** | Can be used to enable / disable customizable extra volumes. | [optional] [default to false]
**CustomMaxMemory** | Pointer to **bool** | Can be used to enable / disable customizable memory. | [optional] [default to false]
**AddVolumes** | Pointer to **bool** | Can be used to enable / disable ability to add volumes | [optional] [default to false]
**SortOrder** | Pointer to **int64** | Sort order | [optional] 
**PriceSets** | Pointer to [**[]AddServicePlansRequestServicePlanPriceSetsInner**](AddServicePlansRequestServicePlanPriceSetsInner.md) | List of price sets to include in service plan. | [optional] 
**Permissions** | Pointer to [**AddServicePlansRequestServicePlanPermissions**](AddServicePlansRequestServicePlanPermissions.md) |  | [optional] 
**Config** | Pointer to [**AddServicePlansRequestServicePlanConfig**](AddServicePlansRequestServicePlanConfig.md) |  | [optional] 

## Methods

### NewAddServicePlansRequestServicePlan

`func NewAddServicePlansRequestServicePlan(name string, code string, maxStorage int64, maxMemory int64, provisionType AddClusterLayoutsRequestLayoutProvisionType, ) *AddServicePlansRequestServicePlan`

NewAddServicePlansRequestServicePlan instantiates a new AddServicePlansRequestServicePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddServicePlansRequestServicePlanWithDefaults

`func NewAddServicePlansRequestServicePlanWithDefaults() *AddServicePlansRequestServicePlan`

NewAddServicePlansRequestServicePlanWithDefaults instantiates a new AddServicePlansRequestServicePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddServicePlansRequestServicePlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddServicePlansRequestServicePlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddServicePlansRequestServicePlan) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *AddServicePlansRequestServicePlan) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddServicePlansRequestServicePlan) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddServicePlansRequestServicePlan) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *AddServicePlansRequestServicePlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddServicePlansRequestServicePlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddServicePlansRequestServicePlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddServicePlansRequestServicePlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRegionCode

`func (o *AddServicePlansRequestServicePlan) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *AddServicePlansRequestServicePlan) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *AddServicePlansRequestServicePlan) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *AddServicePlansRequestServicePlan) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetEditable

`func (o *AddServicePlansRequestServicePlan) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *AddServicePlansRequestServicePlan) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *AddServicePlansRequestServicePlan) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *AddServicePlansRequestServicePlan) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetVisibility

`func (o *AddServicePlansRequestServicePlan) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddServicePlansRequestServicePlan) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddServicePlansRequestServicePlan) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddServicePlansRequestServicePlan) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetMaxStorage

`func (o *AddServicePlansRequestServicePlan) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *AddServicePlansRequestServicePlan) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *AddServicePlansRequestServicePlan) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.


### GetMaxMemory

`func (o *AddServicePlansRequestServicePlan) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *AddServicePlansRequestServicePlan) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *AddServicePlansRequestServicePlan) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.


### GetMaxCores

`func (o *AddServicePlansRequestServicePlan) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *AddServicePlansRequestServicePlan) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *AddServicePlansRequestServicePlan) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *AddServicePlansRequestServicePlan) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxCpu

`func (o *AddServicePlansRequestServicePlan) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *AddServicePlansRequestServicePlan) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *AddServicePlansRequestServicePlan) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *AddServicePlansRequestServicePlan) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *AddServicePlansRequestServicePlan) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *AddServicePlansRequestServicePlan) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *AddServicePlansRequestServicePlan) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *AddServicePlansRequestServicePlan) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetMaxGpus

`func (o *AddServicePlansRequestServicePlan) GetMaxGpus() int64`

GetMaxGpus returns the MaxGpus field if non-nil, zero value otherwise.

### GetMaxGpusOk

`func (o *AddServicePlansRequestServicePlan) GetMaxGpusOk() (*int64, bool)`

GetMaxGpusOk returns a tuple with the MaxGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGpus

`func (o *AddServicePlansRequestServicePlan) SetMaxGpus(v int64)`

SetMaxGpus sets MaxGpus field to given value.

### HasMaxGpus

`func (o *AddServicePlansRequestServicePlan) HasMaxGpus() bool`

HasMaxGpus returns a boolean if a field has been set.

### GetMaxDisks

`func (o *AddServicePlansRequestServicePlan) GetMaxDisks() int64`

GetMaxDisks returns the MaxDisks field if non-nil, zero value otherwise.

### GetMaxDisksOk

`func (o *AddServicePlansRequestServicePlan) GetMaxDisksOk() (*int64, bool)`

GetMaxDisksOk returns a tuple with the MaxDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisks

`func (o *AddServicePlansRequestServicePlan) SetMaxDisks(v int64)`

SetMaxDisks sets MaxDisks field to given value.

### HasMaxDisks

`func (o *AddServicePlansRequestServicePlan) HasMaxDisks() bool`

HasMaxDisks returns a boolean if a field has been set.

### GetProvisionType

`func (o *AddServicePlansRequestServicePlan) GetProvisionType() AddClusterLayoutsRequestLayoutProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *AddServicePlansRequestServicePlan) GetProvisionTypeOk() (*AddClusterLayoutsRequestLayoutProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *AddServicePlansRequestServicePlan) SetProvisionType(v AddClusterLayoutsRequestLayoutProvisionType)`

SetProvisionType sets ProvisionType field to given value.


### GetCustomCpu

`func (o *AddServicePlansRequestServicePlan) GetCustomCpu() bool`

GetCustomCpu returns the CustomCpu field if non-nil, zero value otherwise.

### GetCustomCpuOk

`func (o *AddServicePlansRequestServicePlan) GetCustomCpuOk() (*bool, bool)`

GetCustomCpuOk returns a tuple with the CustomCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCpu

`func (o *AddServicePlansRequestServicePlan) SetCustomCpu(v bool)`

SetCustomCpu sets CustomCpu field to given value.

### HasCustomCpu

`func (o *AddServicePlansRequestServicePlan) HasCustomCpu() bool`

HasCustomCpu returns a boolean if a field has been set.

### GetCustomCores

`func (o *AddServicePlansRequestServicePlan) GetCustomCores() bool`

GetCustomCores returns the CustomCores field if non-nil, zero value otherwise.

### GetCustomCoresOk

`func (o *AddServicePlansRequestServicePlan) GetCustomCoresOk() (*bool, bool)`

GetCustomCoresOk returns a tuple with the CustomCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCores

`func (o *AddServicePlansRequestServicePlan) SetCustomCores(v bool)`

SetCustomCores sets CustomCores field to given value.

### HasCustomCores

`func (o *AddServicePlansRequestServicePlan) HasCustomCores() bool`

HasCustomCores returns a boolean if a field has been set.

### GetCustomMaxStorage

`func (o *AddServicePlansRequestServicePlan) GetCustomMaxStorage() bool`

GetCustomMaxStorage returns the CustomMaxStorage field if non-nil, zero value otherwise.

### GetCustomMaxStorageOk

`func (o *AddServicePlansRequestServicePlan) GetCustomMaxStorageOk() (*bool, bool)`

GetCustomMaxStorageOk returns a tuple with the CustomMaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxStorage

`func (o *AddServicePlansRequestServicePlan) SetCustomMaxStorage(v bool)`

SetCustomMaxStorage sets CustomMaxStorage field to given value.

### HasCustomMaxStorage

`func (o *AddServicePlansRequestServicePlan) HasCustomMaxStorage() bool`

HasCustomMaxStorage returns a boolean if a field has been set.

### GetCustomMaxDataStorage

`func (o *AddServicePlansRequestServicePlan) GetCustomMaxDataStorage() bool`

GetCustomMaxDataStorage returns the CustomMaxDataStorage field if non-nil, zero value otherwise.

### GetCustomMaxDataStorageOk

`func (o *AddServicePlansRequestServicePlan) GetCustomMaxDataStorageOk() (*bool, bool)`

GetCustomMaxDataStorageOk returns a tuple with the CustomMaxDataStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxDataStorage

`func (o *AddServicePlansRequestServicePlan) SetCustomMaxDataStorage(v bool)`

SetCustomMaxDataStorage sets CustomMaxDataStorage field to given value.

### HasCustomMaxDataStorage

`func (o *AddServicePlansRequestServicePlan) HasCustomMaxDataStorage() bool`

HasCustomMaxDataStorage returns a boolean if a field has been set.

### GetCustomMaxMemory

`func (o *AddServicePlansRequestServicePlan) GetCustomMaxMemory() bool`

GetCustomMaxMemory returns the CustomMaxMemory field if non-nil, zero value otherwise.

### GetCustomMaxMemoryOk

`func (o *AddServicePlansRequestServicePlan) GetCustomMaxMemoryOk() (*bool, bool)`

GetCustomMaxMemoryOk returns a tuple with the CustomMaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxMemory

`func (o *AddServicePlansRequestServicePlan) SetCustomMaxMemory(v bool)`

SetCustomMaxMemory sets CustomMaxMemory field to given value.

### HasCustomMaxMemory

`func (o *AddServicePlansRequestServicePlan) HasCustomMaxMemory() bool`

HasCustomMaxMemory returns a boolean if a field has been set.

### GetAddVolumes

`func (o *AddServicePlansRequestServicePlan) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *AddServicePlansRequestServicePlan) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *AddServicePlansRequestServicePlan) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *AddServicePlansRequestServicePlan) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### GetSortOrder

`func (o *AddServicePlansRequestServicePlan) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AddServicePlansRequestServicePlan) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AddServicePlansRequestServicePlan) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AddServicePlansRequestServicePlan) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetPriceSets

`func (o *AddServicePlansRequestServicePlan) GetPriceSets() []AddServicePlansRequestServicePlanPriceSetsInner`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *AddServicePlansRequestServicePlan) GetPriceSetsOk() (*[]AddServicePlansRequestServicePlanPriceSetsInner, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *AddServicePlansRequestServicePlan) SetPriceSets(v []AddServicePlansRequestServicePlanPriceSetsInner)`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *AddServicePlansRequestServicePlan) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### GetPermissions

`func (o *AddServicePlansRequestServicePlan) GetPermissions() AddServicePlansRequestServicePlanPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AddServicePlansRequestServicePlan) GetPermissionsOk() (*AddServicePlansRequestServicePlanPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AddServicePlansRequestServicePlan) SetPermissions(v AddServicePlansRequestServicePlanPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AddServicePlansRequestServicePlan) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetConfig

`func (o *AddServicePlansRequestServicePlan) GetConfig() AddServicePlansRequestServicePlanConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddServicePlansRequestServicePlan) GetConfigOk() (*AddServicePlansRequestServicePlanConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddServicePlansRequestServicePlan) SetConfig(v AddServicePlansRequestServicePlanConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddServicePlansRequestServicePlan) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


